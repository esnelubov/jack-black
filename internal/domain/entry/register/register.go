package register

import (
	"github.com/pkg/errors"
	"jackBlack/internal/domain/entry"
)

type Type struct {
	Repository entry.Repository
}

var ErrPlayerAlreadyExists = errors.New("player already exists")

func (r *Type) Register(name, password string) (err error) {
	var (
		hasPlayer bool
	)

	if hasPlayer, err = r.Repository.HasPlayer(name); err != nil {
		return
	}

	if hasPlayer {
		return ErrPlayerAlreadyExists
	}

	return r.Repository.CreatePlayer(name, password)
}
