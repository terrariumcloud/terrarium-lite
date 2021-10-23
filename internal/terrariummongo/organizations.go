package terrariummongo

import (
	"context"

	"github.com/dylanrhysscott/terrarium/pkg/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionName string = "organizations"

// OrganizationBackend is a struct that implements Mongo operations for organizations
type OrganizationBackend struct {
	Database string
	client   *mongo.Client
}

// Init initializes the Organizations table
func (o *OrganizationBackend) Init() error {
	return nil
}

// Create Adds a new organization to the organizations table
func (o *OrganizationBackend) Create(name string, email string) error {
	return nil
}

// ReadAll Returns all organizations from the organizations table
func (o *OrganizationBackend) ReadAll(limit int, offset int) ([]*types.Organization, error) {
	ctx := context.TODO()
	limitOpt := options.Find().SetLimit(int64(limit))
	skipOpt := options.Find().SetSkip(int64(offset))
	cur, err := o.client.Database(o.Database).Collection(collectionName).Find(ctx, bson.D{}, limitOpt, skipOpt)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var organizationList []*types.Organization = []*types.Organization{}
	for cur.Next(ctx) {
		result := &types.Organization{}
		err := cur.Decode(result)
		if err != nil {
			return nil, err
		}
		organizationList = append(organizationList, result)
	}
	return organizationList, nil
}

// ReadOne Returns a single organization from the organizations table
func (o *OrganizationBackend) ReadOne(id string) (*types.Organization, error) {
	return nil, nil
}

// Update Updates an organization in the organization table
func (o *OrganizationBackend) Update(id string, name string, email string) (*types.Organization, error) {
	return nil, nil
}

// Delete Removes an organization from the organization table
func (o *OrganizationBackend) Delete(id string) error {
	return nil
}
