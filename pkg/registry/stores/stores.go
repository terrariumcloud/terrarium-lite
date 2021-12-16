package stores

import (
	"github.com/dylanrhysscott/terrarium/pkg/registry/data/modules"
	"github.com/dylanrhysscott/terrarium/pkg/registry/data/organizations"
	"github.com/dylanrhysscott/terrarium/pkg/registry/data/vcs"
)

// OrganizationStore is a generic data interface for implementaing database operations relating to organizations
// An organization in Terrarium is a logical grouping or "namespace" under which modules can be stored.
// This interface will provide CRUD related operations for interacting with the organization object
type OrganizationStore interface {
	Init() error
	Create(name string, email string) (*organizations.Organization, error)
	ReadAll(limit int, offset int) ([]*organizations.Organization, error)
	ReadOne(name string) (*organizations.Organization, error)
	Update(name string, email string) (*organizations.Organization, error)
	Delete(name string) error
	GetBackendType() string
}

type ModuleStore interface {
	Init() error
	ReadAll(limit int, offset int) ([]*modules.Module, error)
	ReadOrganizationModules(orgName string, limit int, offset int) ([]*modules.Module, error)
	ReadModuleVersions(orgName string, moduleName string, providerName string) ([]*modules.Module, error)
	ReadOne(orgName string, moduleName string, providerName string, version string) (*modules.Module, error)
	ReadModuleVersionSource(orgName string, moduleName string, providerName string, version string) (string, error)
}

// VCSStore is a generic data interface for implementaing database operations relating to VCS OAuth Connections
// A VCSOAuthConnection is used to implement OAuth operations when interacting with a source provider. This auth
// information needs to be persisted in the database to allow Terrarium to interact with the source provider
// on behalf of the user. This interface will provide CRUD related operations for interacting with the persisted data
type VCSSConnectionStore interface {
	Init() error
	Create(orgID string, orgName string, link *vcs.VCSOAuthClientLink) (*vcs.VCS, error)
	ReadAll(limit int, offset int) ([]*vcs.VCS, error)
	ReadOne(id string, showTokens bool) (*vcs.VCS, error)
	Update(orgID string, orgName string, link *vcs.VCSOAuthClientLink) (*vcs.VCS, error)
	UpdateVCSToken(clientID string, token *vcs.VCSToken) error
	Delete(name string) error
}
