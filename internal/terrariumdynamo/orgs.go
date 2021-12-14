package terrariumdynamo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"

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
		},
		KeySchema: []dynamodbtypes.KeySchemaElement{
			{
				AttributeName: aws.String("_id"),
				KeyType:       "HASH",
			},
			{
				AttributeName: aws.String("name"),
				KeyType:       "RANGE",
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
		var notFoundErr *dynamodbtypes.ResourceNotFoundException
		if errors.As(err, &notFoundErr) {
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
	id := uuid.NewString()
	org := &types.Organization{
		ID:        id,
		Name:      name,
		Email:     email,
		CreatedOn: time.Now().UTC().String(),
	}
	ctx := context.TODO()
	_, err := o.Client.PutItem(ctx, &dynamodb.PutItemInput{
		Item: map[string]dynamodbtypes.AttributeValue{
			"_id": &dynamodbtypes.AttributeValueMemberS{
				Value: id,
			},
			"name": &dynamodbtypes.AttributeValueMemberS{
				Value: org.Name,
			},
			"email": &dynamodbtypes.AttributeValueMemberS{
				Value: org.Email,
			},
			"created_on": &dynamodbtypes.AttributeValueMemberS{
				Value: org.CreatedOn,
			},
		},
		TableName: aws.String(o.TableName),
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// ReadAll Returns all organizations from the organizations table
func (o *OrganizationBackend) ReadAll(limit int, offset int) ([]*types.Organization, error) {
	// https://dynobase.dev/dynamodb-golang-query-examples/#pagination
	ctx := context.TODO()
	p := dynamodb.NewScanPaginator(o.Client, &dynamodb.ScanInput{
		TableName: aws.String(o.TableName),
	})
	var orgs []*types.Organization
	for p.HasMorePages() {
		out, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		var orgList []*types.Organization

		err = attributevalue.UnmarshalListOfMaps(out.Items, &orgList)
		if err != nil {
			return nil, err
		}
		orgs = append(orgs, orgList...)
	}
	var finalOrgList []*types.Organization = orgs
	if offset+limit < len(orgs) {
		finalOrgList = orgs[offset:limit]
	}
	return finalOrgList, nil
}

// ReadOne Returns a single organization from the organizations table
func (o *OrganizationBackend) ReadOne(orgName string) (*types.Organization, error) {
	ctx := context.TODO()
	org, err := o.Client.Query(ctx, &dynamodb.QueryInput{
		KeyConditionExpression: aws.String(fmt.Sprintf("#n = %s", orgName)),
		Limit:                  aws.Int32(int32(1)),
		TableName:              &o.TableName,
		ExpressionAttributeNames: map[string]string{
			"#n": orgName,
		},
	})
	if err != nil {
		return nil, err
	}

	if org.Count > 0 {
		finalOrg := &types.Organization{}
		err = attributevalue.UnmarshalMap(org.Items[0], &finalOrg)
		if err != nil {
			return nil, err
		}
		return finalOrg, nil
	}

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
