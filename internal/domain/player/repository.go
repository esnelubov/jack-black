package player

type Repository interface {
	Load(login string) (*Model, error)
}
