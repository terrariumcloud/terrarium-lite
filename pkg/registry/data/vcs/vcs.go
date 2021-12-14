package vcs

import (
	"github.com/dylanrhysscott/terrarium/pkg/registry/data/relationships"
	"github.com/dylanrhysscott/terrarium/pkg/registry/sources"
)

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

type VCSModule struct {
	ID            string                      `json:"_id"`
	Name          string                      `json:"name"`
	Provider      string                      `json:"provider"`
	Description   string                      `json:"description"`
	VCSConnection *relationships.ResourceLink `json:"vcs_connection"`
	Organization  *relationships.ResourceLink `json:"organization"`
	VCSRepo       sources.SourceData          `json:"vcs_repo"`
}
