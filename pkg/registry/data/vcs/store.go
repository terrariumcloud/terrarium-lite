package vcs

// VCSStore is a generic data interface for implementaing database operations relating to VCS OAuth Connections
// A VCSOAuthConnection is used to implement OAuth operations when interacting with a source provider. This auth
// information needs to be persisted in the database to allow Terrarium to interact with the source provider
// on behalf of the user. This interface will provide CRUD related operations for interacting with the persisted data
type VCSSConnectionStore interface {
	Init() error
	Create(orgID string, orgName string, link *VCSOAuthClientLink) (*VCS, error)
	ReadAll(limit int, offset int) ([]*VCS, error)
	ReadOne(id string, showTokens bool) (*VCS, error)
	Update(orgID string, orgName string, link *VCSOAuthClientLink) (*VCS, error)
	UpdateVCSToken(clientID string, token *VCSToken) error
	Delete(name string) error
}

// SourceProvider is a generic version control provider interface for implementaing source control operations
// FetchVCSSource should accept an OAuth token and a repo name returning some form of source data. If there
// is an error this should be returned and source data is expected to be nil
type SourceProvider interface {
	FetchVCSSource(token string, vcsRepoName string) (SourceData, error)
}

// SourceData is an interface to standardise the kind of data returned by the SourceProvider interface. This
// interface should provide the ability to return repo names, descriptions and owners information for the piece
// of source data. Typically the owner will be the username on the source provider
type SourceData interface {
	GetRepoName() string
	GetRepoDescription() string
	GetRepoOwner() string
}
