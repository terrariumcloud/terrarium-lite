// Package drivers provides interfaces to implement new database backends or storage backends
// for Terrarium.
package drivers

import (
	"context"

	"github.com/terrariumcloud/terrarium-lite/pkg/registry/stores"
)

type TerrariumDatabaseDriver interface {
	Connect(ctx context.Context) error
	Modules() stores.ModuleStore
}

type TerrariumStorageDriver interface {
	GetBackingStoreName() string
	FetchModuleSource(ctx context.Context, key string) ([]byte, error)
}
