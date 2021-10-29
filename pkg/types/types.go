// Package types provides interfaces and structs to implement Terrarium and allow
// extensibility by 3rd parties
package types

import (
	"context"
	"net/http"

	"github.com/dylanrhysscott/terrarium/internal/terrariummongo/orgs"
	vcs "github.com/dylanrhysscott/terrarium/internal/terrariummongo/vcsconn"
)

// TerrariumDriver is a generic database interface to allow further database implementations for Terrarium
// if you would like to implement a different database beyond the core drivers this interface should be implemented
type TerrariumDriver interface {
	Connect(ctx context.Context) error
	Organizations() OrganizationStore
	VCSConnections() VCSSConnectionStore
	GithubSources() SourceStore
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

type APIResponseWriter interface {
	Write(rw http.ResponseWriter, data interface{}, statusCode int)
	Redirect(rw http.ResponseWriter, r *http.Request, uri string)
}

type APIErrorWriter interface {
	Write(rw http.ResponseWriter, err error, statusCode int)
}

// VCSStore is a generic data interface for implementaing database operations relating to VCS OAuth Connections
type VCSSConnectionStore interface {
	Init() error
	Create(orgID string, orgName string, link *vcs.VCSOAuthClientLink) (*vcs.VCS, error)
	ReadAll(limit int, offset int) ([]*vcs.VCS, error)
	ReadOne(id string) (*vcs.VCS, error)
	Update(orgID string, orgName string, link *vcs.VCSOAuthClientLink) (*vcs.VCS, error)
	UpdateVCSToken(clientID string, token *vcs.VCSToken) error
	Delete(name string) error
}

// SourceStore is a generic data interface for implementaing database operations relating to modules
type SourceStore interface {
	FetchVCSSources(token string)
	FetchVCSSource(token string, vcsRepoName string)
}
