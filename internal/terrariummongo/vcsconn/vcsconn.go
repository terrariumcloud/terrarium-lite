package vcsconn

import "github.com/dylanrhysscott/terrarium/internal/terrariummongo/relationships"

// VCS represents the VCS data structure stored in the database
type VCS struct {
	ID           interface{}                 `json:"id" bson:"_id"`
	Organization *relationships.ResourceLink `json:"organization" bson:"organization"`
	OAuth        *VCSOAuthClientLink         `json:"oauth" bson:"oauth"`
}

type VCSToken struct {
	AccessToken           string `json:"access_token,omitempty" bson:"access_token"`
	ExpiresIn             int    `json:"expires_in,omitempty" bson:"expires_in,omitempty"`
	RefreshToken          string `json:"refresh_token,omitempty" bson:"refresh_token,omitempty"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in,omitempty" bson:"refresh_token_expires_in,omitempty"`
	TokenType             string `json:"token_type" bson:"token_type"`
	Scope                 string `json:"scope" scope:"scope"`
}
