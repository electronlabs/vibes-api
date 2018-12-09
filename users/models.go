package users

import (
	"github.com/satori/go.uuid"
)

// User struct contains information about a user.
type User struct {
	ID        uuid.UUID `validate:"required" json:"id"`
	FirstName string    `validate:"required" json:"firstName"`
	LastName  string    `validate:"required" json:"lastName"`
	Email     string    `validate:"required" json:"email"`
}
