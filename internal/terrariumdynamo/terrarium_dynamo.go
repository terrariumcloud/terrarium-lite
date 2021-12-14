package terrariumdynamo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/dylanrhysscott/terrarium/pkg/types"
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
func (d *TerrariumDynamoDB) Organizations() types.OrganizationStore {
	return nil
}

// VCSConnections returns a DynamoDB compatible VCSConnection store which implements the VCSConnectionsStore interface
func (d *TerrariumDynamoDB) VCSConnections() types.VCSSConnectionStore {
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
	return driver, nil
}
