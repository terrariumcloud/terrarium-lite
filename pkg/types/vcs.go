package types

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// VCSStore is a generic data interface for implementaing database operations relating to VCS OAuth Connections
type VCSStore interface {
	Init() error
	Create(orgID string, orgName string, link *VCSOAuthClientLink) (*VCS, error)
	ReadAll(limit int, offset int) ([]*VCS, error)
	ReadOne(id string) (*VCS, error)
	Update(name string, orgName string, serviceProvider string, httpURI string, apiURI string, clientID string, clientSecret string, callback string) (*VCS, error)
	Delete(name string) error
}

// VCS represents the VCS data structure stored in the database
type VCS struct {
	ID           primitive.ObjectID   `json:"id" bson:"_id"`
	Organization *VCSOrganizationLink `json:"organization" bson:"organization"`
	OAuth        *VCSOAuthClientLink  `json:"oauth" bson:"oauth"`
}

type VCSOrganizationLink struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Link string             `json:"link" bson:"link"`
}

type VCSOAuthClientLink struct {
	ServiceProvider string    `json:"service_provider" bson:"service_provider"`
	HTTPURI         string    `json:"http_uri" bson:"http_uri"`
	APIURI          string    `json:"api_uri" bson:"api_uri"`
	ClientID        string    `json:"client_id" bson:"client_id"`
	ClientSecret    string    `json:"client_secret" bson:"client_secret"`
	CallbackURI     string    `json:"callback_uri" bson:"callback_uri"`
	Token           *VCSToken `json:"token" bson:"token"`
}

type VCSToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func (v *VCSOAuthClientLink) Validate() error {
	if v.ServiceProvider == "" {
		return errors.New("service_provider missing. Supported values are: 'github'")
	}
	if v.HTTPURI == "" {
		return errors.New("http_uri missing")
	}
	if v.APIURI == "" {
		return errors.New("api_uri missing")
	}
	if v.ClientID == "" {
		return errors.New("client_id missing")
	}
	if v.ClientSecret == "" {
		return errors.New("client_secret missing")
	}
	if v.CallbackURI == "" {
		return errors.New("callback_uri missing")
	}
	return nil
}
