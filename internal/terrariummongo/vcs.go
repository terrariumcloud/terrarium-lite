package terrariummongo

import (
	"context"
	"fmt"

	"github.com/dylanrhysscott/terrarium/pkg/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// VCSBackend is a struct that implements Mongo operations for organizations
type VCSBackend struct {
	CollectionName string
	Database       string
	client         *mongo.Client
}

// Init initializes the VCS table
func (v *VCSBackend) Init() error {
	return nil
}

// Create Adds a new VCS to the VCS table
func (v *VCSBackend) Create(orgID string, orgName string, link *types.VCSOAuthClientLink) (*types.VCS, error) {
	ctx := context.TODO()
	oid, err := primitive.ObjectIDFromHex(orgID)
	if err != nil {
		return nil, err
	}
	vcsConnection := &types.VCS{
		ID: primitive.NewObjectID(),
		Organization: &types.VCSOrganizationLink{
			ID:   oid,
			Link: fmt.Sprintf("/v1/organizations/%s", orgName),
		},
		OAuth: link,
	}
	_, err = v.client.Database(v.Database).Collection(v.CollectionName).InsertOne(ctx, vcsConnection, options.InsertOne())
	if err != nil {
		return nil, err
	}
	return vcsConnection, nil
}

// ReadAll Returns all VCSs from the VCS table
func (v *VCSBackend) ReadAll(limit int, offset int) ([]*types.VCS, error) {
	// ctx := context.TODO()
	// limitOpt := options.Find().SetLimit(int64(limit))
	// skipOpt := options.Find().SetSkip(int64(offset))
	// cur, err := v.client.Database(v.Database).Collection(v.CollectionName).Find(ctx, bson.D{}, limitOpt, skipOpt)
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

// ReadOne Returns a single VCS from the VCSs table
func (v *VCSBackend) ReadOne(orgName string) (*types.VCS, error) {
	return nil, nil
}

// Update Updates an VCS in the VCS table
func (v *VCSBackend) Update(name string, orgName string, serviceProvider string, httpURI string, apiURI string, clientID string, clientSecret string, callback string) (*types.VCS, error) {
	return nil, nil
}

// Delete Removes an VCS from the VCS table
func (v *VCSBackend) Delete(name string) error {
	return nil
}
