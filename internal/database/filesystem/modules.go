package filesystem

import (
	"errors"
	"github.com/terrariumcloud/terrarium-lite/pkg/registry/data/modules"
)

// fsModuleBackend is a struct that implements Mongo operations for Modules
type fsModuleBackend struct {
	modules []*modules.Module
}

// Init initializes the Modules table
func (m *fsModuleBackend) Init() error {
	return nil
}

func isModuleMatching(module *modules.Module, orgName string, moduleName string, providerName string) bool {
	return module.Organization == orgName && module.Name == moduleName && module.Provider == providerName
}

func (m *fsModuleBackend) filterModules(orgName string, moduleName string, providerName string) []*modules.Module {
	result := make([]*modules.Module, 0, len(m.modules))
	for _, module := range m.modules {
		if isModuleMatching(module, orgName, moduleName, providerName) {
			result = append(result, module)
		}
	}
	return result
}

func (m *fsModuleBackend) findModuleByVersion(orgName string, moduleName string, providerName string, version string) *modules.Module {
	for _, module := range m.modules {
		if isModuleMatching(module, orgName, moduleName, providerName) && module.Version == version {
			return module
		}
	}
	return nil
}

// ReadModuleVersions Returns all versions of a given module from the Modules table
func (m *fsModuleBackend) ReadModuleVersions(orgName string, moduleName string, providerName string) ([]*modules.Module, error) {
	matchingModules := m.filterModules(orgName, moduleName, providerName)
	if len(matchingModules) < 1 {
		return nil, errors.New("No matching modules found")
	}
	return matchingModules, nil
}

func (m *fsModuleBackend) ReadModuleVersionSource(orgName string, moduleName string, providerName string, version string) (string, error) {
	module := m.findModuleByVersion(orgName, moduleName, providerName, version)
	if nil != module {
		return module.Source, nil
	}
	return "", errors.New("No module found for the specified version")
}

// GetBackendType Returns the type of backend used
func (m *fsModuleBackend) GetBackendType() string {
	return "filesystem"
}
