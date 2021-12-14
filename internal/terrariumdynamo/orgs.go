package terrariumdynamo

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/dylanrhysscott/terrarium/pkg/types"
)

// OrganizationBackend is a struct that implements Mongo operations for organizations
type OrganizationBackend struct {
	TableName string
	Client    *dynamodb.Client
}

func getTableSchema(table string) *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		AttributeDefinitions: []dynamodbtypes.AttributeDefinition{
			{
				AttributeName: aws.String("_id"),
				AttributeType: dynamodbtypes.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("name"),
				AttributeType: dynamodbtypes.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("email"),
				AttributeType: dynamodbtypes.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("created_on"),
				AttributeType: dynamodbtypes.ScalarAttributeTypeS,
			},
		},
		KeySchema: []dynamodbtypes.KeySchemaElement{
			{
				AttributeName: aws.String("_id"),
				KeyType:       "HASH",
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

// Init initializes the Organizations table
func (o *OrganizationBackend) Init() error {
	ctx := context.TODO()
	_, err := o.Client.DescribeTable(ctx, &dynamodb.DescribeTableInput{
		TableName: aws.String(o.TableName),
	})
	if err != nil {
		var tableExistsErr *dynamodbtypes.TableNotFoundException
		if errors.As(err, &tableExistsErr) {
			log.Printf("Creating DynamoDB Table: %s", o.TableName)
			_, err = o.Client.CreateTable(ctx, getTableSchema(o.TableName))
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	return nil
}

// Create Adds a new organization to the organizations table
func (o *OrganizationBackend) Create(name string, email string) (*types.Organization, error) {
	return nil, nil
}

// ReadAll Returns all organizations from the organizations table
func (o *OrganizationBackend) ReadAll(limit int, offset int) ([]*types.Organization, error) {
	return nil, nil
}

// ReadOne Returns a single organization from the organizations table
func (o *OrganizationBackend) ReadOne(orgName string) (*types.Organization, error) {
	return nil, nil
}

// Update Updates an organization in the organization table
func (o *OrganizationBackend) Update(name string, email string) (*types.Organization, error) {
	return nil, nil
}

// Delete Removes an organization from the organization table
func (o *OrganizationBackend) Delete(name string) error {
	return nil
}

// GetBackendType Returns the type of backend used
func (o *OrganizationBackend) GetBackendType() string {
	return "dynamodb"
}
