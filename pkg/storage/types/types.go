package types

type OrganizationStore interface {
	Init() error
	Create()
	ReadAll() ([]Organization, error)
	ReadOne()
	Update()
	Delete()
}

type Organization struct {
	ID        string
	Name      string
	CreatedOn string
}
