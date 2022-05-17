package deck

import (
	"crypto/rand"
	"github.com/pkg/errors"
	"jackBlack/internal/domain/game/deck/card"
	"math/big"
)

type Model []*card.Model

func New() *Model {
	result := make(Model, 52)
	for s := 1; s <= 4; s++ {
		for r := 1; r <= 13; r++ {
			c := card.Model{
				Suit:   card.Suit(s),
				Rank:   card.Rank(r),
				Values: card.PossibleValues{r},
			}

			switch c.Rank {
			case card.Jack, card.Queen, card.King:
				c.Values = card.PossibleValues{10}
			case card.Ace:
				c.Values = append(c.Values, 11)
			}

			result[(13*(s-1))+r-1] = &c
		}
	}

	return &result
}

func NewEmpty() *Model {
	result := make(Model, 0, 52)
	return &result
}

func (d Model) Totals() card.PossibleValues {
	result := card.PossibleValues{0}

	for _, c := range d {
		result = result.SumWith(c.Values)
	}

	return result
}

func (d Model) HasBlackjack() bool {
	totals := d.Totals()
	for _, v := range totals {
		if v == 21 {
			return true
		}
	}

	return false
}

func (d Model) Busts() bool {
	totals := d.Totals()
	for _, v := range totals {
		if v <= 21 {
			return false
		}
	}

	return true
}

func (d *Model) ResolveDealersHand(deck *Model) (err error) {
	var thisMax int

	for {
		if thisMax, err = d.Totals().MaxLTE(21); err != nil {
			break
		}

		if thisMax >= 17 {
			break
		}

		if err = deck.DealIntoEach(1, d); err != nil {
			break
		}
	}

	// Dealer hand busts
	if err == card.ErrNoValueFound {
		err = nil
	}

	return
}

func (d Model) GreaterThan(other *Model) (bool, error) {
	var (
		thisMax  int
		otherMax int
		err      error
	)

	if thisMax, err = d.Totals().MaxLTE(21); err != nil {
		return false, err
	}

	if otherMax, err = other.Totals().MaxLTE(21); err != nil {
		return false, err
	}

	return thisMax > otherMax, nil
}

func (d Model) Equal(other *Model) (bool, error) {
	var (
		thisMax  int
		otherMax int
		err      error
	)

	if thisMax, err = d.Totals().MaxLTE(21); err != nil {
		return false, err
	}

	if otherMax, err = other.Totals().MaxLTE(21); err != nil {
		return false, err
	}

	return thisMax == otherMax, nil
}

func (d Model) Shuffled() *Model {
	tmp := make(Model, len(d))
	copy(tmp, d)

	for i := range tmp {
		jBig, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			panic(err)
		}

		j := jBig.Int64()

		tmp[i], tmp[j] = tmp[j], tmp[i]
	}

	return &tmp
}

var ErrDeckEmpty = errors.New("deck is empty")
var ErrNotEnoughCards = errors.New("not enough cards in the deck")

func (d *Model) TakeCard() (c *card.Model, err error) {
	if len(*d) == 0 {
		return nil, ErrDeckEmpty
	}

	c, newDeck := (*d)[len(*d)-1], (*d)[:len(*d)-1]
	*d = newDeck

	return
}

func (d *Model) PushCard(card *card.Model) {
	*d = append(*d, card)
}

func (d *Model) DealIntoEach(numberOfCards int, hands ...*Model) error {
	if numberOfCards*len(hands) > len(*d) {
		return ErrNotEnoughCards
	}

	for _, hand := range hands {
		for i := 0; i < numberOfCards; i++ {
			c, _ := d.TakeCard()
			hand.PushCard(c)
		}
	}

	return nil
}

func (d Model) ShowFirstCard() (c *card.Model, err error) {
	if len(d) == 0 {
		return nil, ErrDeckEmpty
	}

	return d[0], nil
}
