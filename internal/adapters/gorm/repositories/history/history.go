package history

import (
	"gorm.io/gorm"
	"jackBlack/internal/adapters/gorm/entities/history"
	"jackBlack/internal/domain/game/money"
	"jackBlack/internal/domain/player/stats"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) Load(playerID int64) (*stats.Model, error) {
	var historyRecords []*history.Entity

	if err := r.db.Where("player_id = ?", playerID).Find(&historyRecords).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &stats.Model{
				PlayerID:     playerID,
				TotalWins:    0,
				TotalLoses:   0,
				GamesHistory: []*stats.GameHistoryRecord{},
			}, nil
		}

		return nil, err
	}

	statsModel := stats.Model{}

	for _, historyRecord := range historyRecords {
		gameResult := stats.Result(historyRecord.Result)

		switch gameResult {
		case stats.Win:
			statsModel.TotalWins++
		case stats.Lose:
			statsModel.TotalLoses++
		}

		statsModel.GamesHistory = append(statsModel.GamesHistory, &stats.GameHistoryRecord{
			Time:    historyRecord.ResultTime,
			Result:  gameResult,
			Balance: money.New(historyRecord.Balance),
		})
	}

	return &statsModel, nil
}
