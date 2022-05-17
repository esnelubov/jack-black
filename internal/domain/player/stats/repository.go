package stats

type Repository interface {
	Load(playerID int64) (*Model, error)
}
