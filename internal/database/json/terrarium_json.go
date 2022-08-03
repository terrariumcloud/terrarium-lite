package json

import (
	"context"
	json_decoder "encoding/json"
	"github.com/terrariumcloud/terrarium/pkg/registry/data/modules"
	"github.com/terrariumcloud/terrarium/pkg/registry/stores"
	"os"
)

type jsonMetadata struct {
	modules []*modules.Module `json:"modules"`
}

type TerrariumJson struct {
	moduleBackend jsonModuleBackend
}

func (m *TerrariumJson) Connect(_ context.Context) error {
	return nil
}

func (m *TerrariumJson) Modules() stores.ModuleStore {
	return &m.moduleBackend
}

func loadMetadataFromJsonFile(metadataPath string) (*jsonMetadata, error) {
	metadataFile, err := os.Open(metadataPath)
	if err != nil {
		return nil, err
	}
	defer metadataFile.Close()
	decoder := json_decoder.NewDecoder(metadataFile)
	var metadata jsonMetadata
	err = decoder.Decode(&metadata)
	if err != nil {
		return nil, err
	}
	return &metadata, nil
}

func New(metadataPath string) (*TerrariumJson, error) {
	if readOnlyData, err := loadMetadataFromJsonFile(metadataPath); err != nil {
		return nil, err
	} else {
		driver := &TerrariumJson{
			moduleBackend: jsonModuleBackend{
				modules: readOnlyData.modules,
			},
		}
		return driver, nil
	}
}
