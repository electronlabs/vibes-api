package user

import (
	"github.com/electronlabs/vibes-api/models/user"
	uuid "github.com/satori/go.uuid"
)

// Repository provides an abstraction on top of user data source
type Repository interface {
	Create(*user.User) error
	ReadOne(ID uuid.UUID) (*user.User, error)
	ReadAll() ([]user.User, error)
	Update(ID uuid.UUID, updates map[string]interface{}) (*user.User, error)
	Delete(ID uuid.UUID)
}
