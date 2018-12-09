package repository

import (
	"github.com/electronlabs/vibes-api/actions/model"
)

// ActionsRepository provides an abstraction on top of the action data source
type ActionsRepository interface {
	GetAll() ([]model.Action, error)
}
