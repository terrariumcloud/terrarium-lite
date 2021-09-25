package terrariumpsql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type TerrariumPostgres struct {
	Host       string
	User       string
	Password   string
	Database   string
	SSLMode    string
	connection *sql.Conn
	database   *sql.DB
}

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

func (p *TerrariumPostgres) Organizations() *OrganizationBackend {
	return &OrganizationBackend{
		db: p.database,
	}
}

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
	return &TerrariumPostgres{
		Host:     host,
		User:     user,
		Password: password,
		Database: database,
		SSLMode:  sslmode,
	}, nil
}
