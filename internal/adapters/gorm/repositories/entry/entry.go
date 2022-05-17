package entry

import (
	"crypto/sha256"
	"encoding/hex"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"jackBlack/internal/adapters/gorm/entities/balance"
	"jackBlack/internal/adapters/gorm/entities/player"
	"jackBlack/internal/adapters/gorm/entities/session"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) HasPlayer(login string) (bool, error) {
	var record player.Entity
	err := r.db.Where("login = ?", login).Take(&record).Error

	if err == nil {
		return true, nil
	}

	if err == gorm.ErrRecordNotFound {
		return false, nil
	}

	return false, err
}

func hash(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

func (r *Repository) PlayerAuthorized(login, password string) (bool, error) {
	var record player.Entity
	err := r.db.Where("login = ?", login).Where("password_hash = ?", hash(password)).Take(&record).Error

	if err == nil {
		return true, nil
	}

	if err == gorm.ErrRecordNotFound {
		return false, nil
	}

	return false, err
}

func (r *Repository) CreatePlayer(login, password string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		playerRecord := player.Entity{
			Login:        login,
			PasswordHash: hash(password),
		}

		if err := r.db.Create(&playerRecord).Error; err != nil {
			return err
		}

		balanceRecord := balance.Entity{
			PlayerID: playerRecord.PlayerID,
			Balance:  1000,
		}

		if err := r.db.Create(&balanceRecord).Error; err != nil {
			return err
		}

		sessionRecord := session.Entity{
			PlayerID:      playerRecord.PlayerID,
			CurrentAction: "",
			LastMessageID: 0,
			PlayerHand:    datatypes.JSON("[]"),
			DealerHand:    datatypes.JSON("[]"),
			Deck:          datatypes.JSON("[]"),
			Bet:           0,
		}

		if err := r.db.Create(&sessionRecord).Error; err != nil {
			return err
		}

		return nil
	})
}
