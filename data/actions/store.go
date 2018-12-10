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

// GetAll gets all actions from the database
func (s *Store) GetAll() ([]actions.Action, error) {
	// TODO
	return []actions.Action{}, nil
}
