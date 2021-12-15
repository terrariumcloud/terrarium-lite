package dynamo

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/dylanrhysscott/terrarium/pkg/registry/data/modules"
	"github.com/dylanrhysscott/terrarium/pkg/registry/stores"
)

const versionIndex string = "module_version_index"
const allModuleVersionIndex string = "all_module_versions_index"

// ModuleBackend is a struct that implements Mongo operations for Modules
type ModuleBackend struct {
	TableName           string
	OrganizationBackend stores.OrganizationStore
	Client              *dynamodb.Client
}

func (m *ModuleBackend) getTableSchema(table string) *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		AttributeDefinitions: []dynamodbtypes.AttributeDefinition{
			{
				AttributeName: aws.String("_id"),
				AttributeType: dynamodbtypes.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("_organization_id"),
				AttributeType: dynamodbtypes.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("organization"),
				AttributeType: dynamodbtypes.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("name"),
				AttributeType: dynamodbtypes.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("provider"),
				AttributeType: dynamodbtypes.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("version"),
				AttributeType: dynamodbtypes.ScalarAttributeTypeS,
			},
		},
		KeySchema: []dynamodbtypes.KeySchemaElement{
			{
				AttributeName: aws.String("_id"),
				KeyType:       "HASH",
			},
			{
				AttributeName: aws.String("_organization_id"),
				KeyType:       "HASH",
			},
		},
		GlobalSecondaryIndexes: []dynamodbtypes.GlobalSecondaryIndex{
			{
				IndexName: aws.String(allModuleVersionIndex),
				KeySchema: []dynamodbtypes.KeySchemaElement{
					{
						AttributeName: aws.String("organization"),
						KeyType:       "RANGE",
					},
					{
						AttributeName: aws.String("name"),
						KeyType:       "RANGE",
					},
					{
						AttributeName: aws.String("provider"),
						KeyType:       "RANGE",
					},
				},
				Projection: &dynamodbtypes.Projection{
					ProjectionType: dynamodbtypes.ProjectionTypeAll,
				},
				ProvisionedThroughput: &dynamodbtypes.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(1),
					WriteCapacityUnits: aws.Int64(1),
				},
			},
			{
				IndexName: aws.String(versionIndex),
				KeySchema: []dynamodbtypes.KeySchemaElement{
					{
						AttributeName: aws.String("organization"),
						KeyType:       "RANGE",
					},
					{
						AttributeName: aws.String("name"),
						KeyType:       "RANGE",
					},
					{
						AttributeName: aws.String("provider"),
						KeyType:       "RANGE",
					},
					{
						AttributeName: aws.String("version"),
						KeyType:       "RANGE",
					},
				},
				Projection: &dynamodbtypes.Projection{
					ProjectionType: dynamodbtypes.ProjectionTypeAll,
				},
				ProvisionedThroughput: &dynamodbtypes.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(1),
					WriteCapacityUnits: aws.Int64(1),
				},
			},
		},
		TableName:   aws.String(table),
		BillingMode: dynamodbtypes.BillingModeProvisioned,
		ProvisionedThroughput: &dynamodbtypes.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	}
}

// Init initializes the Modules table
func (m *ModuleBackend) Init() error {
	ctx := context.TODO()
	_, err := m.Client.DescribeTable(ctx, &dynamodb.DescribeTableInput{
		TableName: aws.String(m.TableName),
	})
	if err != nil {
		var notFoundErr *dynamodbtypes.ResourceNotFoundException
		if errors.As(err, &notFoundErr) {
			log.Printf("Creating DynamoDB Table: %s", m.TableName)
			_, err = m.Client.CreateTable(ctx, m.getTableSchema(m.TableName))
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	return nil
}

// Create Adds a new module to the Modules table
func (m *ModuleBackend) Create(name string, email string) (*modules.Module, error) {
	return nil, nil
}

// ReadAll Returns all Modules from the Modules table
func (m *ModuleBackend) ReadAll(limit int, offset int) ([]*modules.Module, error) {
	return nil, nil
}

// ReadOne Returns a single module from the Modules table
func (m *ModuleBackend) ReadOne(orgName string) (*modules.Module, error) {
	return nil, nil
}

// Update Updates an module in the module table
func (m *ModuleBackend) Update(name string, email string) (*modules.Module, error) {
	return nil, nil
}

// Delete Removes an module from the module table
func (m *ModuleBackend) Delete(name string) error {
	return nil
}

// GetBackendType Returns the type of backend used
func (m *ModuleBackend) GetBackendType() string {
	return "dynamodb"
}
