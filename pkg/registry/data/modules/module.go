package modules

import (
	"github.com/dylanrhysscott/terrarium/pkg/registry/data/relationships"
)

type Module struct {
	ID           string                      `json:"_id"`
	Name         string                      `json:"name"`
	Provider     string                      `json:"provider"`
	Version      string                      `json:"version"`
	Description  string                      `json:"description"`
	Source       string                      `json:"source"`
	Organization *relationships.ResourceLink `json:"organization"`
}

type ModuleVersionItem struct {
	Version string `json:"version"`
}

type ModuleVersions struct {
	Versions []*ModuleVersionItem `json:"versions"`
}

type ModuleVersionResponse struct {
	Modules []*ModuleVersions `json:"modules"`
}
