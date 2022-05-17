package session

import (
	"encoding/json"
	"gorm.io/gorm"
	"jackBlack/internal/adapters/gorm/entities/balance"
	"jackBlack/internal/adapters/gorm/entities/history"
	player2 "jackBlack/internal/adapters/gorm/entities/player"
	session2 "jackBlack/internal/adapters/gorm/entities/session"
	"jackBlack/internal/domain/game/action"
	"jackBlack/internal/domain/game/deck"
	"jackBlack/internal/domain/game/money"
	"jackBlack/internal/domain/game/session"
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

func (r Repository) LoadForPlayer(player *player.Model) (gameData *session.Model, err error) {
	var (
		sessionRecord session2.Entity
		playerHand    *deck.Model
		dealerHand    *deck.Model
		cardDeck      *deck.Model
	)

	if err = r.db.Where("player_id = ?", player.ID).Take(&sessionRecord).Error; err != nil {
		return nil, err
	}

	gameData = &session.Model{
		LastMessageID: sessionRecord.LastMessageID,
	}

	if sessionRecord.CurrentAction != "" {
		gameData.SetCurrentAction(action.WithName[sessionRecord.CurrentAction])
	}

	gameData.SetPlayer(player)

	if err = json.Unmarshal(sessionRecord.PlayerHand, &playerHand); err != nil {
		return nil, err
	}
	gameData.SetPlayerHand(playerHand)

	if err = json.Unmarshal(sessionRecord.DealerHand, &dealerHand); err != nil {
		return nil, err
	}
	gameData.SetDealerHand(dealerHand)

	if err = json.Unmarshal(sessionRecord.Deck, &cardDeck); err != nil {
		return nil, err
	}
	gameData.SetDeck(cardDeck)

	gameData.SetBet(money.New(sessionRecord.Bet))

	return
}

func (r Repository) Store(gameData *session.Model) (err error) {
	var (
		sessionRecord  *session2.Entity
		playerRecord   *player2.Entity
		balanceRecord  *balance.Entity
		historyRecords []*history.Entity
	)

	if sessionRecord, err = r.makeSessionRecordFrom(gameData); err != nil {
		return
	}

	if err = r.db.Where("login = ?", gameData.Player().Login).Take(&playerRecord).Error; err != nil {
		return err
	}

	if err = r.db.Where("player_id = ?", playerRecord.PlayerID).Take(&balanceRecord).Error; err != nil {
		return err
	}

	balanceRecord.Balance = int64(*gameData.Player().Balance)

	historyRecords = r.makeHistoryRecordsFrom(playerRecord.PlayerID, gameData)

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Save(&sessionRecord).Error; err != nil {
			return err
		}

		if err = tx.Save(&balanceRecord).Error; err != nil {
			return err
		}

		if len(historyRecords) > 0 {
			if err = tx.Create(&historyRecords).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r Repository) makeSessionRecordFrom(gameData *session.Model) (sessionRecord *session2.Entity, err error) {
	var (
		playerHand []byte
		dealerHand []byte
		cardDeck   []byte
	)

	if err = r.db.Where("player_id = ?", gameData.Player().ID).Take(&sessionRecord).Error; err != nil {
		return
	}

	sessionRecord.CurrentAction = gameData.CurrentAction().Name()
	sessionRecord.LastMessageID = gameData.LastMessageID

	if playerHand, err = json.Marshal(gameData.PlayerHand()); err != nil {
		return
	}
	sessionRecord.PlayerHand = playerHand

	if dealerHand, err = json.Marshal(gameData.DealerHand()); err != nil {
		return
	}
	sessionRecord.DealerHand = dealerHand

	if cardDeck, err = json.Marshal(gameData.Deck()); err != nil {
		return
	}
	sessionRecord.Deck = cardDeck

	sessionRecord.Bet = int64(*gameData.Bet())

	return
}

func (r Repository) makeHistoryRecordsFrom(playerID int64, gameData *session.Model) (historyRecords []*history.Entity) {
	historyRecords = make([]*history.Entity, 0, len(gameData.Player().RecentHistory))

	for _, gameRecord := range gameData.Player().RecentHistory {
		historyRecords = append(historyRecords, &history.Entity{
			PlayerID:   playerID,
			ResultTime: gameRecord.Time,
			Balance:    int64(*gameRecord.Balance),
			Result:     string(gameRecord.Result),
		})
	}

	return
}
