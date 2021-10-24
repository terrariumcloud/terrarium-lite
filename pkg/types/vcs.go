package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// VCSStore is a generic data interface for implementaing database operations relating to VCS OAuth Connections
type VCSStore interface {
	Init() error
	Create(name string, orgID string, serviceProvider string, httpURI string, apiURI string, clientID string, clientSecret string, callback string) (*VCS, error)
	ReadAll(limit int, offset int) ([]*VCS, error)
	ReadOne(name string) (*VCS, error)
	Update(name string, orgID string, serviceProvider string, httpURI string, apiURI string, clientID string, clientSecret string, callback string) (*VCS, error)
	Delete(name string) error
}

// VCS represents the VCS data structure stored in the database
type VCS struct {
	ID           primitive.ObjectID   `json:"id" bson:"_id"`
	Name         string               `json:"name" bson:"name"`
	Organization *VCSOrganizationLink `json:"organization" bson:"organization"`
	OAuth        *VCSOAuthClientLink  `json:"oauth" bson:"oauth"`
}

type VCSOrganizationLink struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Link string             `json:"link" bson:"link"`
}

type VCSOAuthClientLink struct {
	ServiceProvider string `json:"service_provider" bson:"service_provider"`
	HTTPURI         string `json:"http_uri" bson:"http_uri"`
	APIURI          string `json:"api_uri" bson:"api_uri"`
	ClientID        string `json:"client_id" bson:"client_id"`
	ClientSecret    string `json:"client_secret" bson:"client_secret"`
	CallbackURI     string `json:"callback_uri" bson:"callback_uri"`
}
