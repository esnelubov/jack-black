package card

import (
	"fmt"
	"github.com/pkg/errors"
)

type Suit int
type Rank int

const (
	Clubs Suit = iota + 1
	Diamonds
	Hearts
	Spades
)

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type PossibleValues []int

func (this PossibleValues) SumWith(other PossibleValues) PossibleValues {
	result := PossibleValues{}

	for _, v := range this {
		for _, ov := range other {
			sum := v + ov
			result = append(result, sum)
		}
	}

	return result
}

var ErrNoValueFound = errors.New("no value found")

func (this PossibleValues) MaxLTE(limit int) (max int, err error) {
	max = -1
	for _, v := range this {
		if v > max && v <= limit {
			max = v
		}
	}

	if max == -1 {
		err = ErrNoValueFound
	}

	return
}

type Model struct {
	Suit   Suit
	Rank   Rank
	Values PossibleValues
}

func (m Model) String() string {
	rank := ""
	switch m.Rank {
	case 1:
		rank = "Ace"
	case 2:
		rank = "2"
	case 3:
		rank = "3"
	case 4:
		rank = "4"
	case 5:
		rank = "5"
	case 6:
		rank = "6"
	case 7:
		rank = "7"
	case 8:
		rank = "8"
	case 9:
		rank = "9"
	case 10:
		rank = "10"
	case 11:
		rank = "Jack"
	case 12:
		rank = "Queen"
	case 13:
		rank = "King"
	}

	suit := ""
	switch m.Suit {
	case 1:
		suit = "Clubs"
	case 2:
		suit = "Diamonds"
	case 3:
		suit = "Hearts"
	case 4:
		suit = "Spades"
	}

	return fmt.Sprintf("%s of %s (%v points)", rank, suit, m.Values)
}
