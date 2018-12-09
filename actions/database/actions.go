package database

import (
	"github.com/electronlabs/vibes-api/actions/model"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Database struct manages interactions with actions database
type Database struct {
	client *mongo.Client
}

// New creates a new Database struct
func New(client *mongo.Client) *Database {
	return &Database{
		client: client,
	}
}

// GetAll gets all actions from the database
func (db *Database) GetAll() ([]model.Action, error) {
	// TODO
	return []model.Action{}, nil
}
