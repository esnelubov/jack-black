package deck

import (
	"github.com/stretchr/testify/assert"
	"jackBlack/internal/domain/game/deck"
	"jackBlack/internal/domain/game/deck/card"
	"testing"
)

func TestTotals(t *testing.T) {
	hand := deck.Model{&card.Model{Values: card.PossibleValues{1, 10}}, &card.Model{Values: card.PossibleValues{2}}}

	assert.Equal(t, card.PossibleValues{3, 12}, hand.Totals())
}

func TestHasBlackjack(t *testing.T) {
	hand := deck.Model{&card.Model{Values: card.PossibleValues{1, 10}}, &card.Model{Values: card.PossibleValues{20}}}

	assert.True(t, hand.HasBlackjack())

	hand = deck.Model{&card.Model{Values: card.PossibleValues{1, 10}}, &card.Model{Values: card.PossibleValues{19}}}

	assert.False(t, hand.HasBlackjack())
}

func TestBusts(t *testing.T) {
	hand := deck.Model{&card.Model{Values: card.PossibleValues{1, 10}}, &card.Model{Values: card.PossibleValues{20}}}

	assert.False(t, hand.Busts())

	hand = deck.Model{&card.Model{Values: card.PossibleValues{1, 10}}, &card.Model{Values: card.PossibleValues{21}}}

	assert.True(t, hand.Busts())
}

func TestResolveDealersHand(t *testing.T) {
	d := &deck.Model{
		&card.Model{Values: card.PossibleValues{2}},
		&card.Model{Values: card.PossibleValues{1}},
	}
	hand := &deck.Model{&card.Model{Values: card.PossibleValues{1, 16}}}
	_ = hand.ResolveDealersHand(d)

	assert.Equal(t, &deck.Model{
		&card.Model{Values: card.PossibleValues{2}},
	}, d)

	assert.Equal(t, &deck.Model{
		&card.Model{Values: card.PossibleValues{1, 16}},
		&card.Model{Values: card.PossibleValues{1}},
	}, hand)
}
