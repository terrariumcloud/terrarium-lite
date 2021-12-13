// Package types provides interfaces to implement Terrarium and allow
// extensibility by 3rd parties
package types

import (
	"context"
)

// TerrariumDriver is a generic database interface to allow further database implementations for Terrarium
// if you would like to implement a different database beyond the core drivers this interface should be implemented
type TerrariumDatabaseDriver interface {
	Connect(ctx context.Context) error
	Organizations() OrganizationStore
	VCSConnections() VCSSConnectionStore
}

// TerrariumSourceDriver is a generic VCS source interface to allow further version control backend implementations for Terrarium
// if you would like to implement a different version control beyond the core drivers this interface should be extended
type TerrariumSourceDriver interface {
	GithubSources() SourceProvider
}

// TerrariumStorageDriver is a generic interface for interacting with underlying module storage.
type TerrariumStorageDriver interface {
	Init() error
	FetchModuleSource(ctx context.Context, bucket string, key string) error
}
