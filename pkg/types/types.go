// Package types provides interfaces and structs to implement Terrarium and allow
// extensibility by 3rd parties
package types

import "context"

// OrganizationStore is a generic data interface for implementaing database operations relating to organizations
type OrganizationStore interface {
	Init() error
	Create(name string, email string) error
	ReadAll() ([]*Organization, error)
	ReadOne(id string) (*Organization, error)
	Update(id string, name string, email string) (*Organization, error)
	Delete(id string) error
}

// Organization represents the organization data structure stored in the database
type Organization struct {
	ID        string
	Name      string
	Email     string
	CreatedOn string
}

// TerrariumDriver is a generic database interface to allow further database implementations for Terrarium
// if you would like to implement a different database beyond the core drivers this interface should be implemented
type TerrariumDriver interface {
	Connect(ctx context.Context) error
	Organizations() OrganizationStore
}
