package entry

type Repository interface {
	HasPlayer(login string) (bool, error)
	PlayerAuthorized(login, password string) (bool, error)
	CreatePlayer(login, password string) error
}
