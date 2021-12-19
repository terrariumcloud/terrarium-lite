package json

import (
	"errors"
	"github.com/terrariumcloud/terrarium/pkg/registry/data/modules"
)

// jsonModuleBackend is a struct that implements Mongo operations for Modules
type jsonModuleBackend struct {
	modules []*modules.Module
}

// Init initializes the Modules table
func (m *jsonModuleBackend) Init() error {
	return nil
}

func isModuleMatching(module *modules.Module, orgName string, moduleName string, providerName string) bool {
	return module.Organization == orgName && module.Name == moduleName && module.Provider == providerName
}

func (m *jsonModuleBackend) filterModules(orgName string, moduleName string, providerName string) []*modules.Module {
	result := make([]*modules.Module, 0, len(m.modules))
	for _, module := range m.modules {
		if isModuleMatching(module, orgName, moduleName, providerName) {
			result = append(result, module)
		}
	}
	return result
}

func (m *jsonModuleBackend) findModuleByVersion(orgName string, moduleName string, providerName string, version string) *modules.Module {
	for _, module := range m.modules {
		if isModuleMatching(module, orgName, moduleName, providerName) && module.Version == version {
			return module
		}
	}
	return nil
}

func (m *jsonModuleBackend) findModulesByOrganization(orgName string) []*modules.Module {
	result := make([]*modules.Module, 0, len(m.modules))
	for _, module := range m.modules {
		if module.Organization == orgName {
			result = append(result, module)
		}
	}
	return result
}

// Create Adds a new module to the Modules table
func (m *jsonModuleBackend) Create(name string, email string) (*modules.Module, error) {
	return nil, errors.New("Operation not supported on Json Backend")
}

// ReadAll Returns all Modules from the Modules table
func (m *jsonModuleBackend) ReadAll(limit int, offset int) ([]*modules.Module, error) {
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
func (m *jsonModuleBackend) ReadOne(orgName string, moduleName string, providerName string) (*modules.Module, error) {
	matchingModules := m.filterModules(orgName, moduleName, providerName)
	if len(matchingModules) < 1 {
		return nil, errors.New("No matching modules found")
	}
	return matchingModules[0], nil
}

// ReadModuleVersions Returns all versions of a given module from the Modules table
func (m *jsonModuleBackend) ReadModuleVersions(orgName string, moduleName string, providerName string) ([]*modules.Module, error) {
	matchingModules := m.filterModules(orgName, moduleName, providerName)
	if len(matchingModules) < 1 {
		return nil, errors.New("No matching modules found")
	}
	return matchingModules, nil
}

// ReadOrganizationModules Returns a list of organization modules
func (m *jsonModuleBackend) ReadOrganizationModules(orgName string, limit int, offset int) ([]*modules.Module, error) {
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

func (m *jsonModuleBackend) ReadModuleVersionSource(orgName string, moduleName string, providerName string, version string) (string, error) {
	module := m.findModuleByVersion(orgName, moduleName, providerName, version)
	if nil != module {
		return module.Source, nil
	}
	return "", errors.New("No module found for the specified version")
}

// Update Updates an module in the module table
func (m *jsonModuleBackend) Update(name string, email string) (*modules.Module, error) {
	return nil, errors.New("Operation not supported on Json Backend")
}

// Delete Removes an module from the module table
func (m *jsonModuleBackend) Delete(name string) error {
	return errors.New("Operation not supported on Json Backend")
}

// GetBackendType Returns the type of backend used
func (m *jsonModuleBackend) GetBackendType() string {
	return "json"
}
