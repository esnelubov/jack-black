package lose

import (
	"jackBlack/internal/common/set_of_strings"
	"jackBlack/internal/domain/game/action"
	"jackBlack/internal/domain/game/constants"
)

type Action string

func (this Action) AllowedFor(session action.GameData) bool {
	if session.CurrentAction() == action.WithName[constants.ActionStand] {
		handGreater, err := session.DealerHand().GreaterThan(session.PlayerHand())

		return (handGreater && err == nil) || session.PlayerHand().Busts()
	}

	return session.PlayerHand().Busts()
}

func (this Action) Perform(session action.GameData, args map[string]interface{}) (err error) {
	session.Player().WriteLose()

	return
}

func (this Action) DescribeState(session action.GameData) (description map[string]interface{}) {
	return map[string]interface{}{
		"balance": session.Player().Balance,
	}
}

func (this Action) NextActions(session action.GameData) (allowedActions set_of_strings.Type) {
	allowedActions = set_of_strings.New()

	allowedActions.Add(constants.ActionEnter)
	return
}

func (from Action) Transit(to string, session action.GameData, args map[string]interface{}) (err error) {
	return action.TransitCommon(from, to, session, args)
}

func (this Action) Name() string {
	return string(this)
}
