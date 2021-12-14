package vcs

import "errors"

type VCSOAuthClientLink struct {
	ServiceProvider string    `json:"service_provider" bson:"service_provider"`
	HTTPURI         string    `json:"http_uri" bson:"http_uri"`
	APIURI          string    `json:"api_uri" bson:"api_uri"`
	ClientID        string    `json:"client_id" bson:"client_id"`
	ClientSecret    string    `json:"client_secret,omitempty" bson:"client_secret"`
	CallbackURI     string    `json:"callback_uri" bson:"callback_uri"`
	Token           *VCSToken `json:"token,omitempty" bson:"token"`
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
	return nil
}
