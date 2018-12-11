package users

// UsersRepository provides an abstraction on top of user data source
type Repository interface {
	GetAll() ([]User, error)
}
