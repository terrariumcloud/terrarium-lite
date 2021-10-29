package terrariumsorce

import (
	"github.com/dylanrhysscott/terrarium/internal/terrariumsource/github"
	"github.com/dylanrhysscott/terrarium/pkg/types"
)

type TerrariumSourceDriver struct{}

func (t *TerrariumSourceDriver) GithubSources() types.SourceProvider {
	return github.NewGithubBackend()
}

func NewTerrariumSourceDriver() *TerrariumSourceDriver {
	return &TerrariumSourceDriver{}
}
