package actions

import (
	"github.com/satori/go.uuid"
)

// Action struct contains information about the main application action.
type Action struct {
	ID      uuid.UUID `validate:"required" json:"id"`
	Point   []float64 `validate:"required" json:"point"`
	ActorID uuid.UUID `validate:"required" json:"actorId"`
}
