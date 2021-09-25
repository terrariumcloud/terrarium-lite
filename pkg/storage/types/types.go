package types

type OrganizationStore interface {
	Init() error
	Create(name string, email string) (*Organization, error)
	ReadAll() ([]Organization, error)
	ReadOne(id string) (*Organization, error)
	Update(id string, name string, email string) (*Organization, error)
	Delete(id string) error
}

type Organization struct {
	ID        string
	Name      string
	Email     string
	CreatedOn string
}
