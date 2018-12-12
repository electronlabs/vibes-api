package actions

import (
	"github.com/electronlabs/vibes-api/domain/actions"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Store struct manages interactions with actions store
type Store struct {
	client *mongo.Client
}

// New creates a new Database struct
func New(client *mongo.Client) *Store {
	return &Store{
		client: client,
	}
}

// ListActions gets all actions from the database
func (s *Store) ListActions() ([]actions.Action, error) {
	// TODO
	return []actions.Action{}, nil
}

// GetAction gets all actions from the database
func (s *Store) GetAction() (actions.Action, error) {
	// TODO
	return actions.Action{}, nil
}
