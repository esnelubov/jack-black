package command

import (
	"context"
	"encoding/json"
	"jackBlack/internal/adapters/gorm/repositories/entry"
	"jackBlack/internal/adapters/gorm/repositories/history"
	"jackBlack/internal/adapters/gorm/repositories/player"
	"jackBlack/internal/adapters/gorm/repositories/session"
	"jackBlack/internal/common/generic_errors"
	sm "jackBlack/internal/domain/game/session"
	pm "jackBlack/internal/domain/player"
	"jackBlack/internal/ports/http_api"
)

type GameMakeActionHandler struct {
	entryRepo   *entry.Repository
	playerRepo  *player.Repository
	sessionRepo *session.Repository
	historyRepo *history.Repository
}

func NewGameMakeActionHandler(entryRepo *entry.Repository, playerRepo *player.Repository, sessionRepo *session.Repository, historyRepo *history.Repository) *GameMakeActionHandler {
	return &GameMakeActionHandler{entryRepo: entryRepo, playerRepo: playerRepo, sessionRepo: sessionRepo, historyRepo: historyRepo}
}

func (h *GameMakeActionHandler) Handle(ctx context.Context, password string, body *http_api.GameMakeActionPayload) (err error) {
	var (
		playerAuthorized bool
		playerModel      *pm.Model
		sessionModel     *sm.Model
		args             map[string]interface{}
	)

	if playerAuthorized, err = h.entryRepo.PlayerAuthorized(body.Login, password); err != nil {
		return
	}

	if !playerAuthorized {
		return generic_errors.ErrAuthorization
	}

	if playerModel, err = h.playerRepo.Load(body.Login); err != nil {
		return
	}

	if sessionModel, err = sm.LoadForPlayer(playerModel, h.sessionRepo); err != nil {
		return err
	}

	// Some previous message suddenly arrived
	if body.SerialId <= sessionModel.LastMessageID {
		return
	} else {
		sessionModel.LastMessageID = body.SerialId
	}

	if body.ArgsJson != nil && *body.ArgsJson != "" {
		if err = json.Unmarshal([]byte(*body.ArgsJson), &args); err != nil {
			return
		}
	}

	if err = sessionModel.CurrentAction().Transit(string(body.Action), sessionModel, args); err != nil {
		return
	}

	if err = sessionModel.Store(); err != nil {
		return
	}

	return
}
