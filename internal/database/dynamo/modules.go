package dynamo

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/dylanrhysscott/terrarium/pkg/registry/data/modules"
	"github.com/dylanrhysscott/terrarium/pkg/registry/stores"
)

const orgModulesIndex string = "organization_module_index"
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
		},
		KeySchema: []dynamodbtypes.KeySchemaElement{
			{
				AttributeName: aws.String("_id"),
				KeyType:       "HASH",
			},
			{
				AttributeName: aws.String("_organization_id"),
				KeyType:       "RANGE",
			},
		},
		GlobalSecondaryIndexes: []dynamodbtypes.GlobalSecondaryIndex{
			{
				IndexName: aws.String(allModuleVersionIndex),
				KeySchema: []dynamodbtypes.KeySchemaElement{
					{
						AttributeName: aws.String("organization"),
						KeyType:       "HASH",
					},
					{
						AttributeName: aws.String("name"),
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
				IndexName: aws.String(orgModulesIndex),
				KeySchema: []dynamodbtypes.KeySchemaElement{
					{
						AttributeName: aws.String("organization"),
						KeyType:       "HASH",
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
	ctx := context.TODO()
	p := dynamodb.NewScanPaginator(m.Client, &dynamodb.ScanInput{
		TableName: aws.String(m.TableName),
	})
	var terraformModules []*modules.Module
	for p.HasMorePages() {
		out, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		var moduleList []*modules.Module

		err = attributevalue.UnmarshalListOfMaps(out.Items, &moduleList)
		if err != nil {
			return nil, err
		}
		terraformModules = append(terraformModules, moduleList...)
	}
	var finalModuleList []*modules.Module = terraformModules
	if offset+limit < len(terraformModules) {
		finalModuleList = terraformModules[offset:limit]
	}
	return finalModuleList, nil
}

// ReadOne Returns a single module from the Modules table
func (m *ModuleBackend) ReadOne(orgName string) (*modules.Module, error) {
	return nil, nil
}

// ReadModuleVersions Returns all versions of a given module from the Modules table
func (m *ModuleBackend) ReadModuleVersions(orgName string, moduleName string, providerName string) ([]*modules.Module, error) {
	ctx := context.TODO()
	p := dynamodb.NewQueryPaginator(m.Client, &dynamodb.QueryInput{
		TableName:              aws.String(m.TableName),
		IndexName:              aws.String(orgModulesIndex),
		KeyConditionExpression: aws.String("#o = :o AND #n = :n AND #p = :p"),
		ExpressionAttributeNames: map[string]string{
			"#o": "organization",
			"#n": "name",
			"#p": "provider",
		},
		ExpressionAttributeValues: map[string]dynamodbtypes.AttributeValue{
			":o": &dynamodbtypes.AttributeValueMemberS{
				Value: orgName,
			},
			":n": &dynamodbtypes.AttributeValueMemberS{
				Value: moduleName,
			},
			":p": &dynamodbtypes.AttributeValueMemberS{
				Value: providerName,
			},
		},
	})
	var terraformModules []*modules.Module
	for p.HasMorePages() {
		out, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		var moduleList []*modules.Module

		err = attributevalue.UnmarshalListOfMaps(out.Items, &moduleList)
		if err != nil {
			return nil, err
		}
		terraformModules = append(terraformModules, moduleList...)
	}

	return terraformModules, nil
}

// ReadOrganizationModules Returns a list of organization modules
func (m *ModuleBackend) ReadOrganizationModules(orgName string, limit int, offset int) ([]*modules.Module, error) {
	ctx := context.TODO()
	p := dynamodb.NewQueryPaginator(m.Client, &dynamodb.QueryInput{
		TableName:              aws.String(m.TableName),
		IndexName:              aws.String(orgModulesIndex),
		KeyConditionExpression: aws.String("#n = :o"),
		ExpressionAttributeNames: map[string]string{
			"#n": "organization",
		},
		ExpressionAttributeValues: map[string]dynamodbtypes.AttributeValue{
			":o": &dynamodbtypes.AttributeValueMemberS{
				Value: orgName,
			},
		},
	})
	var terraformModules []*modules.Module
	for p.HasMorePages() {
		out, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		var moduleList []*modules.Module

		err = attributevalue.UnmarshalListOfMaps(out.Items, &moduleList)
		if err != nil {
			return nil, err
		}
		terraformModules = append(terraformModules, moduleList...)
	}
	var finalModuleList []*modules.Module = terraformModules
	if offset+limit < len(terraformModules) {
		finalModuleList = terraformModules[offset:limit]
	}
	return finalModuleList, nil
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
