package player

type Entity struct {
	PlayerID     int64  `gorm:"primaryKey;column:player_id"`
	Login        string `gorm:"column:login"`
	PasswordHash string `gorm:"column:password_hash"`
}

func (*Entity) TableName() string {
	return "jack_black.players"
}
