package stats

import (
	"jackBlack/internal/domain/game/money"
	"time"
)

type Model struct {
	PlayerID     int64
	TotalWins    int64
	TotalLoses   int64
	GamesHistory []*GameHistoryRecord
}

type GameHistoryRecord struct {
	Time    time.Time
	Result  Result
	Balance *money.Model
}
