package json

import (
	"errors"
	"fmt"
	"github.com/terrariumcloud/terrarium/pkg/registry/data/organizations"
)

// jsonOrganizationBackend is a struct that implements Mongo operations for organizations
type jsonOrganizationBackend struct {
	organizations []*organizations.Organization
}

// Init initializes the Organizations table
func (o *jsonOrganizationBackend) Init() error {
	return nil
}

// Create Adds a new organization to the organizations table
func (o *jsonOrganizationBackend) Create(_ string, _ string) (*organizations.Organization, error) {
	return nil, errors.New("Operation not supported on Json Backend")
}

// ReadAll Returns all organizations from the organizations table
func (o *jsonOrganizationBackend) ReadAll(limit int, offset int) ([]*organizations.Organization, error) {
	count := len(o.organizations)
	if offset >= count {
		return []*organizations.Organization{}, nil
	}
	if offset+limit >= count {
		limit = count - offset
	}
	return o.organizations[offset:limit], nil

}

// ReadOne Returns a single organization from the organizations table
func (o *jsonOrganizationBackend) ReadOne(orgName string) (*organizations.Organization, error) {
	for _, organisation := range o.organizations {
		if organisation.Name == orgName {
			return organisation, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Organization '%s' not found", orgName))
}

// Update Updates an organization in the organization table
func (o *jsonOrganizationBackend) Update(_ string, _ string) (*organizations.Organization, error) {
	return nil, errors.New("Operation not supported on Json Backend")
}

// Delete Removes an organization from the organization table
func (o *jsonOrganizationBackend) Delete(_ string) error {
	return errors.New("Operation not supported on Json Backend")
}

// GetBackendType Returns the type of backend used
func (o *jsonOrganizationBackend) GetBackendType() string {
	return "json"
}
