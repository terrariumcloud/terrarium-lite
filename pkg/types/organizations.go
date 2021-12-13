package types

import "errors"

// Organization represents the organization data structure stored in the database
type Organization struct {
	ID        interface{} `json:"id" bson:"_id"`
	Name      string      `json:"name" bson:"name"`
	Email     string      `json:"email" bson:"email"`
	CreatedOn string      `json:"created_on" bson:"created_on"`
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
