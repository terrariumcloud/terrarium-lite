package vcs

// VCS represents the VCS data structure stored in the database
type VCS struct {
	ID           interface{}         `json:"id" bson:"_id"`
	Organization *ResourceLink       `json:"organization" bson:"organization"`
	OAuth        *VCSOAuthClientLink `json:"oauth" bson:"oauth"`
}

type ResourceLink struct {
	ID   interface{} `json:"id" bson:"_id"`
	Link string      `json:"link" bson:"link"`
}

type VCSToken struct {
	AccessToken           string `json:"access_token" bson:"access_token"`
	ExpiresIn             int    `json:"expires_in" bson:"expires_in"`
	RefreshToken          string `json:"refresh_token" bson:"refresh_token"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in" bson:"refresh_token_expires_in"`
	TokenType             string `json:"token_type" bson:"token_type"`
	Scope                 string `json:"scope" scope:"scope"`
}
