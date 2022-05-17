package session

import (
	"github.com/pkg/errors"
	"jackBlack/internal/domain/player"
)

var ErrSessionNotExist = errors.New("session doesn't exist")

type Repository interface {
	LoadForPlayer(player *player.Model) (gameData *Model, err error)
	Store(gameData *Model) (err error)
}
