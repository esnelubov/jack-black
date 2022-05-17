package player

import (
	"jackBlack/internal/domain/game/money"
	"jackBlack/internal/domain/player/stats"
	"time"
)

type Model struct {
	ID            int64
	Login         string
	Balance       *money.Model
	RecentHistory []*stats.GameHistoryRecord
}

func (p *Model) WriteWin() {
	p.RecentHistory = append(p.RecentHistory, &stats.GameHistoryRecord{
		Time:    time.Now(),
		Result:  stats.Win,
		Balance: p.Balance,
	})
}

func (p *Model) WriteLose() {
	p.RecentHistory = append(p.RecentHistory, &stats.GameHistoryRecord{
		Time:    time.Now(),
		Result:  stats.Lose,
		Balance: p.Balance,
	})
}

func (p *Model) WriteTie() {
	p.RecentHistory = append(p.RecentHistory, &stats.GameHistoryRecord{
		Time:    time.Now(),
		Result:  stats.Tie,
		Balance: p.Balance,
	})
}
