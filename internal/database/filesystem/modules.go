package filesystem

import (
	"errors"
	"github.com/terrariumcloud/terrarium/pkg/registry/data/modules"
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

func (m *fsModuleBackend) findModulesByOrganization(orgName string) []*modules.Module {
	result := make([]*modules.Module, 0, len(m.modules))
	for _, module := range m.modules {
		if module.Organization == orgName {
			result = append(result, module)
		}
	}
	return result
}

// ReadAll Returns all Modules from the Modules table
func (m *fsModuleBackend) ReadAll(limit int, offset int) ([]*modules.Module, error) {
	count := len(m.modules)
	if offset >= count {
		return []*modules.Module{}, nil
	}
	if offset+limit >= count {
		limit = count - offset
	}
	return m.modules[offset:limit], nil
}

// ReadOne Returns a single module from the Modules table
func (m *fsModuleBackend) ReadOne(orgName string, moduleName string, providerName string) (*modules.Module, error) {
	matchingModules := m.filterModules(orgName, moduleName, providerName)
	if len(matchingModules) < 1 {
		return nil, errors.New("No matching modules found")
	}
	return matchingModules[0], nil
}

// ReadModuleVersions Returns all versions of a given module from the Modules table
func (m *fsModuleBackend) ReadModuleVersions(orgName string, moduleName string, providerName string) ([]*modules.Module, error) {
	matchingModules := m.filterModules(orgName, moduleName, providerName)
	if len(matchingModules) < 1 {
		return nil, errors.New("No matching modules found")
	}
	return matchingModules, nil
}

// ReadOrganizationModules Returns a list of organization modules
func (m *fsModuleBackend) ReadOrganizationModules(orgName string, limit int, offset int) ([]*modules.Module, error) {
	modulesForOrganization := m.findModulesByOrganization(orgName)
	count := len(modulesForOrganization)
	if offset >= count {
		return []*modules.Module{}, nil
	}
	if offset+limit >= count {
		limit = count - offset
	}
	return modulesForOrganization[offset:limit], nil
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
