// Package terrariumsql provides Postgres support for Terrarium by implementing the TerrariumDriver interface
package terrariumpsql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/dylanrhysscott/terrarium/pkg/types"
	_ "github.com/lib/pq"
)

// TerrariumPostgres implements Postgres support for Terrarium for all API's
type TerrariumPostgres struct {
	Host       string
	User       string
	Password   string
	Database   string
	SSLMode    string
	connection *sql.Conn
	database   *sql.DB
}

// Connect iniitialises a database connection to Postgres
func (p *TerrariumPostgres) Connect(ctx context.Context) error {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", p.User, p.Password, p.Host, p.Database, p.SSLMode)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	conn, err := db.Conn(ctx)
	if err != nil {
		return err
	}
	p.connection = conn
	p.database = db
	return nil
}

// Organizations returns a Postgres compatible organization store which implements the OrganizationStore interface
func (p *TerrariumPostgres) Organizations() types.OrganizationStore {
	return &OrganizationBackend{
		db: p.database,
	}
}

// New creates a TerrariumPostgres driver
func New(host string, user string, password string, database string, sslmode string) (*TerrariumPostgres, error) {
	if sslmode == "" {
		sslmode = "disable"
	}
	if host == "" {
		return nil, errors.New("postgres host cannot be empty")
	}
	if user == "" {
		return nil, errors.New("postgres user cannot be empty")
	}
	if password == "" {
		return nil, errors.New("postgres password cannot be empty")
	}
	driver := &TerrariumPostgres{
		Host:     host,
		User:     user,
		Password: password,
		Database: database,
		SSLMode:  sslmode,
	}
	err := driver.Connect(context.TODO())
	if err != nil {
		return nil, err
	}
	err = driver.Organizations().Init()
	if err != nil {
		return nil, err
	}
	return driver, nil
}
