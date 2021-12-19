// Package drivers provides interfaces to implement new database backends, storage backends or VCS backends
// for Terrarium. Terrarium has been designed for extensibility by 3rd parties in mind. If you would like to add
// further support please implement an interface. Existing implementations can be found internally within Terrarium
// for reference
package drivers

import (
	"context"

	"github.com/terrariumcloud/terrarium/pkg/registry/sources"
	"github.com/terrariumcloud/terrarium/pkg/registry/stores"
)

// TerrariumDriver is a generic database interface to allow further database implementations for Terrarium
// if you would like to implement a different database beyond the core drivers this interface should be implemented
// Terrarium currently supports DynamoDB and MongoDB (WIP)
type TerrariumDatabaseDriver interface {
	Connect(ctx context.Context) error
	Organizations() stores.OrganizationStore
	Modules() stores.ModuleStore
	VCSConnections() stores.VCSSConnectionStore
}

// TerrariumSourceDriver is a generic VCS source interface to allow further version control backend implementations for Terrarium
// if you would like to implement a different version control beyond the core drivers this interface should be extended
// Terrarium currently supports Github (WIP)
type TerrariumSourceDriver interface {
	GithubSources() sources.SourceProvider
}

// TerrariumStorageDriver is a generic interface for interacting with underlying module storage and
// to allow further module storage backend implementations for Terrarium. If you would like to implement a
// different module storage beyond the core drivers this interface should be extended
// Terrarium currently supports S3 (WIP)
type TerrariumStorageDriver interface {
	GetBackingStoreName() string
	FetchModuleSource(ctx context.Context, key string) ([]byte, error)
}
