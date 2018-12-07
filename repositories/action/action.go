package user

import (
	"github.com/electronlabs/vibes-api/models/action"
	uuid "github.com/satori/go.uuid"
)

// Repository provides an abstraction on top of the action data source
type Repository interface {
	Create(*action.Action) error
	ReadOne(ID uuid.UUID) (*action.Action, error)
	ReadAll() ([]action.Action, error)
	Read(center []float64, radius float64) ([]action.Action, error)
}
