package balance

type Entity struct {
	BalanceID int64 `gorm:"primaryKey;column:balance_id"`
	PlayerID  int64 `gorm:"column:player_id"`
	Balance   int64 `gorm:"column:balance"`
}

func (*Entity) TableName() string {
	return "jack_black.balances"
}
