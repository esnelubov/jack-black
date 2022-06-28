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

func (h *PlayerCreateHandler) Handle(ctx context.Context, body *http_api.PlayerCreatePayload) (err error) {
	var (
		playerExists bool
	)

	if playerExists, err = h.entryRepo.HasPlayer(body.Login); err != nil {
		return err
	}

	if playerExists {
		return generic_errors.ErrLoginAlreadyTaken
	}

	if err = h.entryRepo.CreatePlayer(body.Login, body.Password); err != nil {
		return err
	}

	return
}
