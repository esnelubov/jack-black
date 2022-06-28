package query

import (
	"context"
	"encoding/json"
	"jackBlack/internal/adapters/gorm/repositories/entry"
	"jackBlack/internal/adapters/gorm/repositories/player"
	"jackBlack/internal/adapters/gorm/repositories/session"
	"jackBlack/internal/common/generic_errors"
	sm "jackBlack/internal/domain/game/session"
	pm "jackBlack/internal/domain/player"
	"jackBlack/internal/ports/http_api"
)

type GameStateHandler struct {
	entryRepo   *entry.Repository
	playerRepo  *player.Repository
	sessionRepo *session.Repository
}

func NewGameGetStateHandler(entryRepo *entry.Repository, playerRepo *player.Repository, sessionRepo *session.Repository) *GameStateHandler {
	return &GameStateHandler{entryRepo: entryRepo, playerRepo: playerRepo, sessionRepo: sessionRepo}
}

func (h *GameStateHandler) Handle(ctx context.Context, password string, params *http_api.GameStateParams) (gameState *http_api.GameState, err error) {
	var (
		playerAuthorized bool
		playerModel      *pm.Model
		sessionModel     *sm.Model
	)

	if playerAuthorized, err = h.entryRepo.PlayerAuthorized(params.Login, password); err != nil {
		return nil, err
	}

	if !playerAuthorized {
		return nil, generic_errors.ErrAuthorization
	}

	if playerModel, err = h.playerRepo.Load(params.Login); err != nil {
		return nil, err
	}

	if sessionModel, err = sm.LoadForPlayer(playerModel, h.sessionRepo); err != nil {
		return nil, err
	}

	gameState, err = makeGameStateResponseFrom(sessionModel)
	return
}

func makeGameStateResponseFrom(sessionModel *sm.Model) (gameState *http_api.GameState, err error) {
	var (
		description []byte
		actions     = []http_api.GameAction{}
	)

	description, err = json.Marshal(sessionModel.CurrentAction().DescribeState(sessionModel))

	for _, action := range sessionModel.CurrentAction().NextActions(sessionModel).Items() {
		actions = append(actions, http_api.GameAction(action))
	}

	gameState = &http_api.GameState{
		AllowedActions:  actions,
		DescriptionJson: string(description),
		SerialId:        sessionModel.LastMessageID,
	}

	return
}
