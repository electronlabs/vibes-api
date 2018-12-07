package actions

import (
	actions "github.com/electronlabs/vibes-api/actions/model"
	uuid "github.com/satori/go.uuid"
)

// Repository provides an abstraction on top of the action data source
type Repository interface {
	Create(*actions.Action) error
	ReadOne(ID uuid.UUID) (*actions.Action, error)
	ReadAll() ([]actions.Action, error)
	Read(center []float64, radius float64) ([]actions.Action, error)
}
