package repository

import (
	"github.com/electronlabs/vibes-api/actions/model"
)

// Repository provides an abstraction on top of the action data source
type Repository interface {
	GetAll() ([]model.Action, error)
}
