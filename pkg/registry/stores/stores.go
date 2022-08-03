package stores

import (
	"github.com/terrariumcloud/terrarium/pkg/registry/data/modules"
)

type ModuleStore interface {
	Init() error
	ReadAll(limit int, offset int) ([]*modules.Module, error)
	ReadOrganizationModules(orgName string, limit int, offset int) ([]*modules.Module, error)
	ReadModuleVersions(orgName string, moduleName string, providerName string) ([]*modules.Module, error)
	ReadOne(orgName string, moduleName string, providerName string) (*modules.Module, error)
	ReadModuleVersionSource(orgName string, moduleName string, providerName string, version string) (string, error)
}
