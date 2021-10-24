package types

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrganizationStore is a generic data interface for implementaing database operations relating to organizations
type OrganizationStore interface {
	Init() error
	Create(name string, email string) (*Organization, error)
	ReadAll(limit int, offset int) ([]*Organization, error)
	ReadOne(name string) (*Organization, error)
	Update(name string, email string) (*Organization, error)
	Delete(name string) error
}

// Organization represents the organization data structure stored in the database
type Organization struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	CreatedOn string             `json:"created_on" bson:"created_on"`
}

func (o *Organization) Validate() error {
	if o.Name == "" {
		return errors.New("missing organization name")
	}

	if o.Email == "" {
		return errors.New("missing organization email")
	}

	return nil
}
