package auth

import (
	"github.com/pkg/errors"
	"jackBlack/internal/domain/entry"
)

type Type struct {
	Repository entry.Repository
}

var ErrPlayerNotAuthorized = errors.New("player not authorized")

func (l *Type) Authorize(login, password string) (err error) {
	var authorized bool

	if authorized, err = l.Repository.PlayerAuthorized(login, password); err != nil {
		return err
	}

	if !authorized {
		err = ErrPlayerNotAuthorized
	}

	return
}
