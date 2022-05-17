package history

import "time"

type Entity struct {
	HistoryID  int64     `gorm:"primaryKey;column:history_id"`
	PlayerID   int64     `gorm:"column:player_id"`
	ResultTime time.Time `gorm:"column:result_time"`
	Balance    int64     `gorm:"column:balance"`
	Result     string    `gorm:"column:result"`
}

func (*Entity) TableName() string {
	return "jack_black.history"
}
