package terrariummongo

import (
	"context"

	"github.com/dylanrhysscott/terrarium/pkg/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ModuleBackend is a struct that implements Mongo operations for modules
type ModuleBackend struct {
	CollectionName string
	Database       string
	client         *mongo.Client
}

// Init initializes the Organizations table
func (m *ModuleBackend) Init() error {
	collection := m.client.Database(m.Database).Collection(m.CollectionName)
	// Ensures unique email and name combination
	_, err := collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "name", Value: 1},
			{Key: "namespace", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}
	return nil
}

// Create Adds a new module to the modules table
func (m *ModuleBackend) Create() (*types.Module, error) {
	return nil, nil
}

// ReadAll Returns all organizations from the organizations table
func (m *ModuleBackend) ReadAll(limit int, offset int) ([]*types.Module, error) {
	return nil, nil
}

// ReadOne Returns a single organization from the organizations table
func (m *ModuleBackend) ReadOne() (*types.Module, error) {
	return nil, nil
}

// Update Updates an organization in the organization table
func (m *ModuleBackend) Update() (*types.Module, error) {
	return nil, nil
}

// Delete Removes an organization from the organization table
func (m *ModuleBackend) Delete() error {
	return nil
}
