package terrariummongo

import (
	"github.com/dylanrhysscott/terrarium/pkg/types"
)

// OrganizationBackend is a struct that implements Mongo operations for organizations
type OrganizationBackend struct {
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
	return nil, nil
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
