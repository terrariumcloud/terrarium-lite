package dynamo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/terrariumcloud/terrarium/pkg/registry/stores"
)

// TerrariumDynamoDB implements DynamoDB support for Terrarium for all API's
type TerrariumDynamoDB struct {
	Region  string
	Service *dynamodb.Client
	config  aws.Config
}

func (d *TerrariumDynamoDB) Connect(ctx context.Context) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(d.Region))
	if err != nil {
		return err
	}
	d.Service = dynamodb.NewFromConfig(cfg)
	d.config = cfg
	return nil
}

// Organizations returns a DynamoDB compatible organization store which implements the OrganizationStore interface
func (d *TerrariumDynamoDB) Organizations() stores.OrganizationStore {
	return &OrganizationBackend{
		TableName: "terrarium_organizations",
		Client:    d.Service,
	}
}

// Modules returns a DynamoDB compatible module store which implements the ModuleStore interface
func (d *TerrariumDynamoDB) Modules() stores.ModuleStore {
	return &ModuleBackend{
		TableName:           "terrarium_modules",
		Client:              d.Service,
		OrganizationBackend: d.Organizations(),
	}
}

// VCSConnections returns a DynamoDB compatible VCSConnection store which implements the VCSConnectionsStore interface
func (d *TerrariumDynamoDB) VCSConnections() stores.VCSSConnectionStore {
	return nil
}

// New creates a TerrariumDynamoDB driver
func New(region string) (*TerrariumDynamoDB, error) {
	driver := &TerrariumDynamoDB{
		Region: region,
	}
	err := driver.Connect(context.TODO())
	if err != nil {
		return nil, err
	}
	err = driver.Organizations().Init()
	if err != nil {
		return nil, err
	}
	err = driver.Modules().Init()
	if err != nil {
		return nil, err
	}
	return driver, nil
}
