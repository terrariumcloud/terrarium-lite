package types

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

// OrganizationStore is a generic data interface for implementaing database operations relating to organizations
// An organization in Terrarium is a logical grouping or "namespace" under which modules can be stored.
// This interface will provide CRUD related operations for interacting with the organization object
type OrganizationStore interface {
	Init() error
	Create(name string, email string) (*Organization, error)
	ReadAll(limit int, offset int) ([]*Organization, error)
	ReadOne(name string) (*Organization, error)
	Update(name string, email string) (*Organization, error)
	Delete(name string) error
	GetBackendType() string
}

// ModuleStore is a generic data interface for implementaing database operations relating to modules
// This interface will provide CRUD related operations for interacting with the module object

// TODO: Expand Interface for full CRUD
type ModuleStore interface {
	Init() error
	ReadAll(limit int, offset int) ([]*Module, error)
	ReadOne(name string) (*Module, error)
}
