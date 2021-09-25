package terrariumpsql

import (
	"database/sql"

	"github.com/dylanrhysscott/terrarium/pkg/storage/types"
)

type OrganizationStore interface {
	Init() error
	Create()
	ReadAll() ([]*types.Organization, error)
	ReadOne()
	Update()
	Delete()
}

type OrganizationBackend struct {
	db *sql.DB
}

func (o *OrganizationBackend) Init() error {
	query := `CREATE TABLE IF NOT EXISTS organizations (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) UNIQUE NOT NULL,
		created_on TIMESTAMP NOT NULL	
	);`
	_, err := o.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrganizationBackend) Create() {

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
		err := result.Scan(&org.ID, &org.Name, &org.CreatedOn)
		if err != nil {
			return nil, err
		}
		organizations = append(organizations, org)
	}
	return organizations, nil
}

func (o *OrganizationBackend) Update() {

}

func (o *OrganizationBackend) Delete() {

}
