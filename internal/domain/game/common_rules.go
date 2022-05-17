package game

import (
	"jackBlack/internal/domain/game/deck"
	"jackBlack/internal/domain/game/deck/card"
	"log"
)

const MinimalBet = 2

func ShowDealersFirstCard(deck *deck.Model) (c *card.Model) {
	var err error
	if c, err = deck.ShowFirstCard(); err != nil {
		log.Fatalf("You're trying to show card from the dealer's hand. But the hand is empty.")
	}
	return
}
