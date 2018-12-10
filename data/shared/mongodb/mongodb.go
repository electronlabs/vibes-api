package mongodb

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// Connect creates a MongoDB client from a connection string and connects to the database.
func Connect(uri string) (*mongo.Client, error) {
	client, err := mongo.NewClient(uri)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Disconnect disconnects from the database
func Disconnect(client *mongo.Client) error {
	err := client.Disconnect(context.Background())
	if err != nil {
		return err
	}

	return nil
}
