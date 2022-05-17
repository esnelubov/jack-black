package generic_errors

import "github.com/pkg/errors"

type ErrorType string

var (
	ErrAuthorization      = errors.New("authorization information is missing or invalid")
	ErrPlayerNotExists    = errors.New("a player with the specified login was not found")
	ErrLoginAlreadyTaken  = errors.New("a player with the given login is already exists")
	ErrActionIsNotAllowed = errors.New("action is not allowed at the current state of the game")
)
