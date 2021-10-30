package types

import (
	"github.com/dylanrhysscott/terrarium/internal/terrariummongo/orgs"
	"github.com/dylanrhysscott/terrarium/internal/terrariummongo/vcsconn"
)

// VCSStore is a generic data interface for implementaing database operations relating to VCS OAuth Connections
// A VCSOAuthConnection is used to implement OAuth operations when interacting with a source provider. This auth
// information needs to be persisted in the database to allow Terrarium to interact with the source provider
// on behalf of the user. This interface will provide CRUD related operations for interacting with the persisted data
type VCSSConnectionStore interface {
	Init() error
	Create(orgID string, orgName string, link *vcsconn.VCSOAuthClientLink) (*vcsconn.VCS, error)
	ReadAll(limit int, offset int) ([]*vcsconn.VCS, error)
	ReadOne(id string, showTokens bool) (*vcsconn.VCS, error)
	Update(orgID string, orgName string, link *vcsconn.VCSOAuthClientLink) (*vcsconn.VCS, error)
	UpdateVCSToken(clientID string, token *vcsconn.VCSToken) error
	Delete(name string) error
}

// OrganizationStore is a generic data interface for implementaing database operations relating to organizations
// An organization in Terrarium is a logical grouping or "namespace" under which modules can be stored.
// This interface will provide CRUD related operations for interacting with the organization object
type OrganizationStore interface {
	Init() error
	Create(name string, email string) (*orgs.Organization, error)
	ReadAll(limit int, offset int) ([]*orgs.Organization, error)
	ReadOne(name string) (*orgs.Organization, error)
	Update(name string, email string) (*orgs.Organization, error)
	Delete(name string) error
	GetBackendType() string
}
