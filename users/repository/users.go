package repository

import (
	"github.com/electronlabs/vibes-api/users/model"
	"github.com/satori/go.uuid"
)

// UsersRepository provides an abstraction on top of user data source
type UsersRepository interface {
	Create(*users.User) error
	ReadOne(ID uuid.UUID) (*users.User, error)
	ReadAll() ([]users.User, error)
	Update(ID uuid.UUID, updates map[string]interface{}) (*users.User, error)
	Delete(ID uuid.UUID)
}
