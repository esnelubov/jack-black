package session

import (
	"gorm.io/datatypes"
)

type Entity struct {
	SessionID     int64          `gorm:"primaryKey;column:session_id"`
	PlayerID      int64          `gorm:"column:player_id"`
	CurrentAction string         `gorm:"column:current_action"`
	LastMessageID int64          `gorm:"column:last_message_id"`
	PlayerHand    datatypes.JSON `gorm:"column:player_hand"`
	DealerHand    datatypes.JSON `gorm:"column:dealer_hand"`
	Deck          datatypes.JSON `gorm:"column:deck"`
	Bet           int64          `gorm:"column:bet"`
}

func (*Entity) TableName() string {
	return "jack_black.sessions"
}
