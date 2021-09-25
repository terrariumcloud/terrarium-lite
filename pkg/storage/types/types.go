package types

type OrganizationStore interface {
	Init() error
	Create(name string, email string) (*Organization, error)
	ReadAll() ([]Organization, error)
	ReadOne()
	Update()
	Delete()
}

type Organization struct {
	ID        string
	Name      string
	Email     string
	CreatedOn string
}
