package money

import "github.com/pkg/errors"

type Model int64

func New(amount int64) *Model {
	m := Model(amount)
	return &m
}

func Zero() *Model {
	zero := Model(0)
	return &zero
}

func (m *Model) Deposit(sum Model) {
	*m += sum
}

func (m *Model) Spend(sum Model) error {
	if sum > *m {
		return ErrOverdraw
	}

	*m -= sum

	return nil
}

func (m *Model) Transfer(sum Model, to *Model) (err error) {
	if err = m.Spend(sum); err != nil {
		return err
	}
	to.Deposit(sum)
	return
}

var ErrOverdraw = errors.New("you can't spend more money than you have")
