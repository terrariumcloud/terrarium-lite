package filesystem

import (
	"context"
	"fmt"
	"github.com/terrariumcloud/terrarium/pkg/registry/data/modules"
	"github.com/terrariumcloud/terrarium/pkg/registry/stores"
	"log"
	"path/filepath"
	"strings"
)

type adapter struct {
	moduleBackend fsModuleBackend
}

func (m *adapter) Connect(_ context.Context) error {
	return nil
}

func (m *adapter) Modules() stores.ModuleStore {
	return &m.moduleBackend
}

func loadFromPath(modulesPath string) ([]*modules.Module, error) {
	allModules := make([]*modules.Module, 0)

	matches, _ := filepath.Glob(fmt.Sprintf("%s/**.zip", modulesPath))

	for _, name := range matches {
		sourcePath, err := filepath.Rel(modulesPath, name)
		if err != nil {
			return nil, err
		}

		elements := filepath.SplitList(sourcePath)
		switch len(elements) {
		case 4: // organization / module name / provider / version.zip
			module := modules.Module{
				ID:             strings.Join(elements, "/"),
				OrganizationID: elements[0],
				Name:           elements[1],
				Organization:   elements[0],
				Provider:       elements[2],
				Version:        elements[3][:len(elements[3])-4],
				Description:    "",
				Source:         sourcePath,
			}
			allModules = append(allModules, &module)
			log.Printf("INFO: Added module %s", name)
			break
		case 3: // organization / module name / version.zip
			module := modules.Module{
				ID:             strings.Join(elements, "/"),
				OrganizationID: elements[0],
				Name:           elements[1],
				Organization:   elements[0],
				Provider:       "",
				Version:        elements[2][:len(elements[2])-4],
				Description:    "",
				Source:         name,
			}
			allModules = append(allModules, &module)
			log.Printf("INFO: Added module %s", name)
			break
		default:
			log.Printf("WARN: Ignoring invalid module path: %s", name)
		}
	}
	return allModules, nil
}

func New(modulesPath string) (*adapter, error) {
	if allModules, err := loadFromPath(modulesPath); err != nil {
		return nil, err
	} else {
		driver := &adapter{
			moduleBackend: fsModuleBackend{
				modules: allModules,
			},
		}
		return driver, nil
	}
}
