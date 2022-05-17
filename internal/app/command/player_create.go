package command

import (
	"context"
	"jackBlack/internal/adapters/gorm/repositories/entry"
	"jackBlack/internal/adapters/gorm/repositories/player"
	"jackBlack/internal/common/generic_errors"
	"jackBlack/internal/ports/http_api"
)

type PlayerCreateHandler struct {
	entryRepo  *entry.Repository
	playerRepo *player.Repository
}

func NewPlayerCreateHandler(entryRepo *entry.Repository, playerRepo *player.Repository) *PlayerCreateHandler {
	return &PlayerCreateHandler{entryRepo: entryRepo, playerRepo: playerRepo}
}

func (h *PlayerCreateHandler) Handle(ctx context.Context, body *http_api.PlayerCreatePayload) (player *http_api.Player, err error) {
	var (
		playerExists bool
	)

	if playerExists, err = h.entryRepo.HasPlayer(body.Login); err != nil {
		return nil, err
	}

	if playerExists {
		return nil, generic_errors.ErrLoginAlreadyTaken
	}

	if err = h.entryRepo.CreatePlayer(body.Login, body.Password); err != nil {
		return nil, err
	}

	player = &http_api.Player{
		Balance: 0,
	}

	return
}
