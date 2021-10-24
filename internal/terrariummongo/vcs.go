package terrariummongo

import (
	"github.com/dylanrhysscott/terrarium/pkg/types"
	"go.mongodb.org/mongo-driver/mongo"
)

// VCSBackend is a struct that implements Mongo operations for organizations
type VCSBackend struct {
	CollectionName string
	Database       string
	client         *mongo.Client
}

// Init initializes the VCS table
func (o *VCSBackend) Init() error {
	return nil
}

// Create Adds a new VCS to the VCS table
func (o *VCSBackend) Create(name string, orgID string, serviceProvider string, httpURI string, apiURI string, clientID string, clientSecret string, callback string) (*types.VCS, error) {

	return nil, nil
}

// ReadAll Returns all VCSs from the VCS table
func (o *VCSBackend) ReadAll(limit int, offset int) ([]*types.VCS, error) {
	return nil, nil
}

// ReadOne Returns a single VCS from the VCSs table
func (o *VCSBackend) ReadOne(orgName string) (*types.VCS, error) {
	return nil, nil
}

// Update Updates an VCS in the VCS table
func (o *VCSBackend) Update(name string, orgID string, serviceProvider string, httpURI string, apiURI string, clientID string, clientSecret string, callback string) (*types.VCS, error) {
	return nil, nil
}

// Delete Removes an VCS from the VCS table
func (o *VCSBackend) Delete(name string) error {
	return nil
}
