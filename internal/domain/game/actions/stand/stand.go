package stand

import (
	"jackBlack/internal/common/set_of_strings"
	"jackBlack/internal/domain/game/action"
	"jackBlack/internal/domain/game/constants"
)

type Action string

func (this Action) AllowedFor(session action.GameData) bool {
	return true
}

func (this Action) Perform(session action.GameData, args map[string]interface{}) (err error) {
	return session.DealerHand().ResolveDealersHand(session.Deck())
}

func (this Action) DescribeState(session action.GameData) (description map[string]interface{}) {
	return map[string]interface{}{
		"bet":         session.Bet(),
		"player_hand": session.PlayerHand(),
		"dealer_hand": session.DealerHand(),
	}
}

func (this Action) NextActions(session action.GameData) (allowedActions set_of_strings.Type) {
	allowedActions = set_of_strings.New()

	if action.WithName[constants.ActionLose].AllowedFor(session) {
		allowedActions.Add(constants.ActionLose)
		return
	}

	if action.WithName[constants.ActionWin].AllowedFor(session) {
		allowedActions.Add(constants.ActionWin)
		return
	}

	if action.WithName[constants.ActionTie].AllowedFor(session) {
		allowedActions.Add(constants.ActionTie)
		return
	}

	return
}

func (from Action) Transit(to string, session action.GameData, args map[string]interface{}) (err error) {
	return action.TransitCommon(from, to, session, args)
}

func (this Action) Name() string {
	return string(this)
}
