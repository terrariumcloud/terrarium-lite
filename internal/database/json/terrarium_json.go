package json

import (
	"context"
	json_decoder "encoding/json"
	"github.com/terrariumcloud/terrarium/pkg/registry/data/modules"
	"github.com/terrariumcloud/terrarium/pkg/registry/data/organizations"
	"github.com/terrariumcloud/terrarium/pkg/registry/stores"
	"os"
)

type jsonMetadata struct {
	modules       []*modules.Module             `json:"modules"`
	organizations []*organizations.Organization `json:"organizations"`
}

// TerrariumMongo implements Mongo support for Terrarium for all API's
type TerrariumJson struct {
	moduleBackend       jsonModuleBackend
	organizationBackend jsonOrganizationBackend
}

// Connect initialises a database connection to mongo
func (m *TerrariumJson) Connect(_ context.Context) error {
	return nil
}

// Organizations returns a compatible organization store which implements the OrganizationStore interface
func (m *TerrariumJson) Organizations() stores.OrganizationStore {
	return &m.organizationBackend
}

// Modules returns a compatible ModuleBackend store which implements the VCSConnectionsStore interface
func (m *TerrariumJson) Modules() stores.ModuleStore {
	return &m.moduleBackend
}

// VCSConnections returns nil as there is not use case to support a VCS connection.
func (m *TerrariumJson) VCSConnections() stores.VCSSConnectionStore {
	return nil
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

// New creates a TerrariumMongo driver
func New(metadataPath string) (*TerrariumJson, error) {
	if readOnlyData, err := loadMetadataFromJsonFile(metadataPath); err != nil {
		return nil, err
	} else {
		driver := &TerrariumJson{
			moduleBackend: jsonModuleBackend{
				modules: readOnlyData.modules,
			},
			organizationBackend: jsonOrganizationBackend{
				organizations: readOnlyData.organizations,
			},
		}
		return driver, nil
	}
}
