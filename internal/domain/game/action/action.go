package action

import (
	"jackBlack/internal/common/set_of_strings"
	"jackBlack/internal/domain/game/deck"
	"jackBlack/internal/domain/game/money"
	"jackBlack/internal/domain/player"
)

type SomeAction interface {
	Name() string

	// AllowedFor session?
	AllowedFor(session GameData) bool

	// Perform this action. In case of an error state of the game should be reverted
	Perform(session GameData, args map[string]interface{}) (err error)

	// DescribeState of the game. Should be called after this action is Perform'ed
	DescribeState(session GameData) (description map[string]interface{})

	// NextActions that user can Perform at the current state of the game
	NextActions(session GameData) (allowedActions set_of_strings.Type)

	// Transit to the one of NextActions and Perform it. In case of an error should cancel Transit.
	// CurrentAction in session should be modified accordingly
	Transit(to string, session GameData, args map[string]interface{}) (err error)
}

type GameData interface {
	CurrentAction() SomeAction
	SetCurrentAction(action SomeAction)

	Player() *player.Model
	SetPlayer(player *player.Model)

	PlayerHand() *deck.Model
	SetPlayerHand(hand *deck.Model)

	DealerHand() *deck.Model
	SetDealerHand(hand *deck.Model)

	Deck() *deck.Model
	SetDeck(deck *deck.Model)

	Bet() *money.Model
	SetBet(money *money.Model)
}
