package session

import (
	"jackBlack/internal/domain/game/action"
	"jackBlack/internal/domain/game/actions"
	"jackBlack/internal/domain/game/constants"
	"jackBlack/internal/domain/game/deck"
	"jackBlack/internal/domain/game/money"
	"jackBlack/internal/domain/player"
)

type Model struct {
	LastMessageID int64
	currentAction action.SomeAction
	player        *player.Model
	playerHand    *deck.Model
	dealerHand    *deck.Model
	deck          *deck.Model
	bet           *money.Model
	repository    Repository
}

func (s *Model) SetCurrentAction(action action.SomeAction) {
	s.currentAction = action
}

func (s *Model) SetPlayer(player *player.Model) {
	s.player = player
}

func (s *Model) SetPlayerHand(hand *deck.Model) {
	s.playerHand = hand
}

func (s *Model) SetDealerHand(hand *deck.Model) {
	s.dealerHand = hand
}

func (s *Model) SetDeck(deck *deck.Model) {
	s.deck = deck
}

func (s *Model) SetBet(money *money.Model) {
	s.bet = money
}

func (s *Model) CurrentAction() action.SomeAction {
	return s.currentAction
}

func (s *Model) Player() *player.Model {
	return s.player
}

func (s *Model) PlayerHand() *deck.Model {
	return s.playerHand
}

func (s *Model) DealerHand() *deck.Model {
	return s.dealerHand
}

func (s *Model) Deck() *deck.Model {
	return s.deck
}

func (s *Model) Bet() *money.Model {
	return s.bet
}

func LoadForPlayer(player *player.Model, rep Repository) (session *Model, err error) {
	actions.InitActionsMap()

	if session, err = rep.LoadForPlayer(player); err != nil {
		return
	}

	session.repository = rep

	if session.currentAction == nil {
		session.SetCurrentAction(action.WithName[constants.ActionEnter])
		_ = session.currentAction.Perform(session, map[string]interface{}{})
	}

	return
}

func (s *Model) Store() (err error) {
	return s.repository.Store(s)
}
