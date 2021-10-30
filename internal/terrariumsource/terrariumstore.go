// Package terrariumsource is an internal package providing the ability to interact with various source
// providers such as Github. This is used to populate the registry with modules backed by version control
// This package is internal and not intended for use outside of the core project. See the "types" package
// for extending Terrarium
package terrariumsorce

import (
	"github.com/dylanrhysscott/terrarium/internal/terrariumsource/github"
	"github.com/dylanrhysscott/terrarium/pkg/types"
)

// TerrariumSourceDriver implements the types.TerrariumSourceDriver interface for generic multi provider interactions.
// Further source providers maybe added in future
type TerrariumSourceDriver struct{}

// GithubSources returns a Github backend which can be used to fetch data
// via the Github API
func (t *TerrariumSourceDriver) GithubSources() types.SourceProvider {
	return github.NewGithubBackend()
}

// NewTerrariumSourceDriver creates a new instance of the TerrariumSourceDriver ready for use
// via API calls
func NewTerrariumSourceDriver() *TerrariumSourceDriver {
	return &TerrariumSourceDriver{}
}
