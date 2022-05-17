package enter

import (
	"jackBlack/internal/common/set_of_strings"
	"jackBlack/internal/domain/game/action"
	"jackBlack/internal/domain/game/constants"
	"jackBlack/internal/domain/game/deck"
	"jackBlack/internal/domain/game/money"
)

type Action string

func (this Action) AllowedFor(session action.GameData) bool {
	return true
}

func (this Action) Perform(session action.GameData, args map[string]interface{}) (err error) {
	session.SetDeck(deck.New().Shuffled())
	session.SetPlayerHand(deck.NewEmpty())
	session.SetDealerHand(deck.NewEmpty())
	session.SetBet(money.Zero())
	return
}

func (this Action) DescribeState(session action.GameData) map[string]interface{} {
	return map[string]interface{}{
		"balance":   session.Player().Balance,
		"deck_size": len(*session.Deck()),
	}
}

func (this Action) NextActions(session action.GameData) (allowedActions set_of_strings.Type) {
	allowedActions = set_of_strings.New()

	if action.WithName[constants.ActionBet].AllowedFor(session) {
		allowedActions.Add(constants.ActionBet)
	}
	return
}

func (from Action) Transit(to string, session action.GameData, args map[string]interface{}) (err error) {
	return action.TransitCommon(from, to, session, args)
}

func (this Action) Name() string {
	return string(this)
}
