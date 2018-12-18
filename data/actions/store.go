package actions

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/electronlabs/vibes-api/domain/actions"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Store struct manages interactions with actions store
type Store struct {
	client       *mongo.Client
	databaseName string
}

func (s *Store) getCollection(databaseName string) *mongo.Collection {
	collection := s.client.Database(databaseName).Collection("collection")
	return collection
}

// New creates a new Database struct
func New(client *mongo.Client, databaseName string) *Store {
	return &Store{
		client:       client,
		databaseName: databaseName,
	}
}

// ListActions gets all actions from the database
func (s *Store) ListActions() ([]actions.Action, error) {
	collection := s.getCollection(s.databaseName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var results []actions.Action
	actionDoc := actions.Action{}
	for cur.Next(ctx) {
		err := cur.Decode(&actionDoc)
		spew.Dump(actionDoc)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, actionDoc)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	spew.Dump(results)
	return results, nil
}

// GetAction gets all actions from the database
func (s *Store) GetAction(actionId string) (actions.Action, error) {
	var result actions.Action
	filter := bson.M{"id": actionId}
	collection := s.getCollection(s.databaseName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		spew.Dump(err)
		return result, errors.New("Action doc not found")
	}
	return result, err
}
