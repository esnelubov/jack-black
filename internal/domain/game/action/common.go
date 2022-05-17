package action

import (
	"fmt"
	"github.com/pkg/errors"
	"jackBlack/internal/common/set_of_strings"
	"jackBlack/internal/domain/game/constants"
)

var ErrActionNotAllowed = errors.New("action not allowed")

func MakeErrArgNorProvided(argName string) error {
	return fmt.Errorf("required argument '%s' is not provided", argName)
}

func MakeErrArgParse(arg interface{}, expectedType string) error {
	return fmt.Errorf("can't parse arg %+v as %s", arg, expectedType)
}

func TransitCommon(from SomeAction, to string, session GameData, args map[string]interface{}) (err error) {
	var nextAction SomeAction

	if !from.NextActions(session).Has(to) {
		return ErrActionNotAllowed
	}

	nextAction = WithName[to]

	if err = nextAction.Perform(session, args); err != nil {
		return err
	}

	session.SetCurrentAction(nextAction)

	return
}

func NextActionsCommon(session GameData) (allowedActions set_of_strings.Type) {
	allowedActions = set_of_strings.New()

	if WithName[constants.ActionLose].AllowedFor(session) {
		allowedActions.Add(constants.ActionLose)
		return
	}

	if WithName[constants.ActionWin].AllowedFor(session) {
		allowedActions.Add(constants.ActionWin)
		return
	}

	if WithName[constants.ActionTie].AllowedFor(session) {
		allowedActions.Add(constants.ActionTie)
		return
	}

	if WithName[constants.ActionDoubleDown].AllowedFor(session) {
		allowedActions.Add(constants.ActionDoubleDown)
	}

	allowedActions.Add(constants.ActionHit)
	allowedActions.Add(constants.ActionStand)
	return
}
