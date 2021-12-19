// Package sources is an internal package providing the ability to interact with various source
// providers such as Github. This is used to populate the registry with modules backed by version control
// This package is internal and not intended for use outside of the core project. See the "types" package
// for extending Terrarium
package sourcecontrol

import (
	"github.com/terrariumcloud/terrarium/internal/sourcecontrol/github"
	"github.com/terrariumcloud/terrarium/pkg/registry/sources"
)

// TerrariumSourceControl implements the types.TerrariumSourceDriver interface for generic multi provider interactions.
// Further source providers maybe added in future
type TerrariumSourceControl struct{}

// GithubSources returns a Github backend which can be used to fetch data
// via the Github API
func (t *TerrariumSourceControl) GithubSources() sources.SourceProvider {
	return github.NewGithubBackend()
}
