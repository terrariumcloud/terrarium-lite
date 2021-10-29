package types

import (
	"github.com/dylanrhysscott/terrarium/internal/terrariummongo/orgs"
	"github.com/dylanrhysscott/terrarium/internal/terrariummongo/vcsconn"
)

// VCSStore is a generic data interface for implementaing database operations relating to VCS OAuth Connections
type VCSSConnectionStore interface {
	Init() error
	Create(orgID string, orgName string, link *vcsconn.VCSOAuthClientLink) (*vcsconn.VCS, error)
	ReadAll(limit int, offset int) ([]*vcsconn.VCS, error)
	ReadOne(id string) (*vcsconn.VCS, error)
	Update(orgID string, orgName string, link *vcsconn.VCSOAuthClientLink) (*vcsconn.VCS, error)
	UpdateVCSToken(clientID string, token *vcsconn.VCSToken) error
	Delete(name string) error
}

// OrganizationStore is a generic data interface for implementaing database operations relating to organizations
type OrganizationStore interface {
	Init() error
	Create(name string, email string) (*orgs.Organization, error)
	ReadAll(limit int, offset int) ([]*orgs.Organization, error)
	ReadOne(name string) (*orgs.Organization, error)
	Update(name string, email string) (*orgs.Organization, error)
	Delete(name string) error
	GetBackendType() string
}
