package hit

import (
	"jackBlack/internal/common/set_of_strings"
	"jackBlack/internal/domain/game"
	"jackBlack/internal/domain/game/action"
)

type Action string

func (this Action) AllowedFor(session action.GameData) bool {
	return true
}

func (this Action) Perform(session action.GameData, args map[string]interface{}) (err error) {
	_ = session.Deck().DealIntoEach(1, session.PlayerHand())
	return
}

func (this Action) DescribeState(session action.GameData) (description map[string]interface{}) {
	return map[string]interface{}{
		"balance":     session.Player().Balance,
		"bet":         session.Bet(),
		"player_hand": session.PlayerHand(),
		"dealer_card": game.ShowDealersFirstCard(session.DealerHand()),
	}
}

func (this Action) NextActions(session action.GameData) (allowedActions set_of_strings.Type) {
	return action.NextActionsCommon(session)
}

func (from Action) Transit(to string, session action.GameData, args map[string]interface{}) (err error) {
	return action.TransitCommon(from, to, session, args)
}

func (this Action) Name() string {
	return string(this)
}
