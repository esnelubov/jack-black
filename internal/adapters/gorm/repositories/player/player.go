package player

import (
	"gorm.io/gorm"
	"jackBlack/internal/adapters/gorm/entities/balance"
	player2 "jackBlack/internal/adapters/gorm/entities/player"
	"jackBlack/internal/domain/game/money"
	"jackBlack/internal/domain/player"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Load(login string) (*player.Model, error) {
	var (
		playerRecord  player2.Entity
		balanceRecord balance.Entity
	)

	if err := r.db.Where("login = ?", login).Take(&playerRecord).Error; err != nil {
		return nil, err
	}

	if err := r.db.Where("player_id = ?", playerRecord.PlayerID).Take(&balanceRecord).Error; err != nil {
		return nil, err
	}

	playerModel := &player.Model{
		ID:      playerRecord.PlayerID,
		Login:   playerRecord.Login,
		Balance: money.New(balanceRecord.Balance),
	}

	return playerModel, nil
}
