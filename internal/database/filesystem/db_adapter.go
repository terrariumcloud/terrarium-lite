package filesystem

import (
	"context"
	"fmt"
	"github.com/terrariumcloud/terrarium-lite/pkg/registry/data/modules"
	"github.com/terrariumcloud/terrarium-lite/pkg/registry/stores"
	"log"
	"os"
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

	matches, _ := filepath.Glob(fmt.Sprintf("%s/*/*/*/*.zip", modulesPath))

	for _, name := range matches {
		sourcePath, err := filepath.Rel(modulesPath, name)
		if err != nil {
			return nil, err
		}

		elements := strings.Split(sourcePath, string(os.PathSeparator))
		if len(elements) == 4 {
			module := modules.Module{
				Name:         elements[1],
				Organization: elements[0],
				Provider:     elements[2],
				Version:      elements[3][:len(elements[3])-4], // remove .zip from the version name.
			}
			allModules = append(allModules, &module)
			log.Printf("INFO: Added module %s", name)
		} else {
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
