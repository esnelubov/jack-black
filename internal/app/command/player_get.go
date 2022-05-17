package command

import (
	"context"
	"jackBlack/internal/adapters/gorm/repositories/entry"
	"jackBlack/internal/adapters/gorm/repositories/player"
	"jackBlack/internal/common/generic_errors"
	pm "jackBlack/internal/domain/player"
	"jackBlack/internal/ports/http_api"
)

type PlayerGetHandler struct {
	entryRepo  *entry.Repository
	playerRepo *player.Repository
}

func NewPlayerGetHandler(entryRepo *entry.Repository, playerRepo *player.Repository) *PlayerGetHandler {
	return &PlayerGetHandler{entryRepo: entryRepo, playerRepo: playerRepo}
}

func (h *PlayerGetHandler) Handle(ctx context.Context, password string, params *http_api.PlayerGetParams) (player *http_api.Player, err error) {
	var (
		playerExists     bool
		playerAuthorized bool
		playerModel      *pm.Model
	)

	if playerExists, err = h.entryRepo.HasPlayer(params.Login); err != nil {
		return nil, err
	}

	if !playerExists {
		return nil, generic_errors.ErrPlayerNotExists
	}

	if playerAuthorized, err = h.entryRepo.PlayerAuthorized(params.Login, password); err != nil {
		return nil, err
	}

	if !playerAuthorized {
		return nil, generic_errors.ErrAuthorization
	}

	if playerModel, err = h.playerRepo.Load(params.Login); err != nil {
		return nil, err
	}

	player = &http_api.Player{
		Balance: int64(*playerModel.Balance),
	}

	return
}
