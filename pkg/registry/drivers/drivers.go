// Package drivers provides interfaces to implement Terrarium and allow
// extensibility by 3rd parties
package drivers

import (
	"context"

	"github.com/dylanrhysscott/terrarium/pkg/registry/sources"
	"github.com/dylanrhysscott/terrarium/pkg/registry/stores"
)

// TerrariumDriver is a generic database interface to allow further database implementations for Terrarium
// if you would like to implement a different database beyond the core drivers this interface should be implemented
type TerrariumDatabaseDriver interface {
	Connect(ctx context.Context) error
	Organizations() stores.OrganizationStore
	VCSConnections() stores.VCSSConnectionStore
}

// TerrariumSourceDriver is a generic VCS source interface to allow further version control backend implementations for Terrarium
// if you would like to implement a different version control beyond the core drivers this interface should be extended
type TerrariumSourceDriver interface {
	GithubSources() sources.SourceProvider
}

// TerrariumStorageDriver is a generic interface for interacting with underlying module storage.
type TerrariumStorageDriver interface {
	Init() error
	FetchModuleSource(ctx context.Context, bucket string, key string) ([]byte, error)
}
