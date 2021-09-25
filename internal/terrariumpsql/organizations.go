package terrariumpsql

import (
	"database/sql"
	"time"

	"github.com/dylanrhysscott/terrarium/pkg/storage/types"
)

type OrganizationBackend struct {
	db *sql.DB
}

func (o *OrganizationBackend) Init() error {
	query := `CREATE TABLE IF NOT EXISTS organizations (
		id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		email VARCHAR(120) NOT NULL,
		created_on TIMESTAMP NOT NULL,
		UNIQUE(name, email)
	);`
	_, err := o.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrganizationBackend) Create(name string, email string) error {
	query := `INSERT INTO organizations (name, email, created_on) VALUES ($1, $2, $3)`
	stmt, err := o.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, email, time.Now().UTC())
	if err != nil {
		return err
	}
	return nil
}

func (o *OrganizationBackend) ReadAll() ([]*types.Organization, error) {
	query := `SELECT * FROM organizations;`
	result, err := o.db.Query(query)
	if err != nil {
		return nil, err
	}
	organizations := []*types.Organization{}
	for result.Next() {
		org := &types.Organization{}
		err := result.Scan(&org.ID, &org.Name, &org.Email, &org.CreatedOn)
		if err != nil {
			return nil, err
		}
		organizations = append(organizations, org)
	}
	return organizations, nil
}

func (o *OrganizationBackend) ReadOne(id string) (*types.Organization, error) {
	query := `SELECT * FROM organizations WHERE id = $1;`
	stmt, err := o.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	result, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	if result.Next() {
		org := &types.Organization{}
		err = result.Scan(&org.ID, &org.Name, &org.Email, &org.CreatedOn)
		if err != nil {
			return nil, err
		}
		return org, nil
	}
	return nil, nil
}

func (o *OrganizationBackend) Update() {

}

func (o *OrganizationBackend) Delete(id string) error {
	query := `DELETE FROM organizations WHERE id = $1`
	stmt, err := o.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
