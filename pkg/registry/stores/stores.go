package stores

import (
	"github.com/terrariumcloud/terrarium-lite/pkg/registry/data/modules"
)

type ModuleStore interface {
	Init() error
	ReadModuleVersions(orgName string, moduleName string, providerName string) ([]*modules.Module, error)
	ReadModuleVersionSource(orgName string, moduleName string, providerName string, version string) (string, error)
}
