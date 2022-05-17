package money

import (
	"github.com/stretchr/testify/assert"
	"jackBlack/internal/domain/game/money"
	"testing"
)

func TestDeposit(t *testing.T) {
	m := money.New(100)
	bet := money.New(50)
	m.Deposit(*bet)

	assert.Equal(t, money.New(150), m)
}

func TestSpend(t *testing.T) {
	m := money.New(100)
	bet := money.New(50)
	_ = m.Spend(*bet)

	assert.Equal(t, money.New(50), m)
}

func TestOverdraw(t *testing.T) {
	m := money.New(100)
	bet := money.New(150)
	err := m.Spend(*bet)

	assert.ErrorIs(t, err, money.ErrOverdraw)
}

func TestTransfer(t *testing.T) {
	from := money.New(100)
	to := money.New(40)
	_ = from.Transfer(money.Model(25), to)

	assert.Equal(t, money.New(75), from)
	assert.Equal(t, money.New(65), to)
}
