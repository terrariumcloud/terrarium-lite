package terrariummongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/dylanrhysscott/terrarium/pkg/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TerrariumMongo implements Mongo support for Terrarium for all API's
type TerrariumMongo struct {
	Host     string
	User     string
	Password string
	Database string
	client   *mongo.Client
}

// Connect iniitialises a database connection to mongo
func (m *TerrariumMongo) Connect(ctx context.Context) error {
	connectionStr := fmt.Sprintf("mongodb://%s", m.Host)
	if m.User != "" && m.Password != "" {
		connectionStr = fmt.Sprintf("mongodb://%s:%s@%s", m.User, m.Password, m.Host)
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionStr))
	if err != nil {
		return err
	}
	m.client = client
	return nil
}

// Organizations returns a Mongo compatible organization store which implements the OrganizationStore interface
func (m *TerrariumMongo) Organizations() types.OrganizationStore {
	return &OrganizationBackend{
		CollectionName: "organizations",
		client:         m.client,
		Database:       m.Database,
	}
}

// VCS returns a Mongo compatible VCS store which implements the VCSStore interface
func (m *TerrariumMongo) VCS() types.VCSStore {
	return &VCSBackend{
		CollectionName: "vcs",
		client:         m.client,
		Database:       m.Database,
	}
}

// VCS returns a Mongo compatible VCS store which implements the ModuleStore interface
func (m *TerrariumMongo) Modules() types.ModuleStore {
	return &ModuleBackend{
		CollectionName: "modules",
		client:         m.client,
		Database:       m.Database,
	}
}

// New creates a TerrariumMongo driver
func New(host string, user string, password string, database string) (*TerrariumMongo, error) {
	if host == "" {
		return nil, errors.New("mongo host cannot be empty")
	}
	driver := &TerrariumMongo{
		Host:     host,
		User:     user,
		Password: password,
		Database: database,
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
