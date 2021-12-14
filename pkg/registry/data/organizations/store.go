package organizations

// OrganizationStore is a generic data interface for implementaing database operations relating to organizations
// An organization in Terrarium is a logical grouping or "namespace" under which modules can be stored.
// This interface will provide CRUD related operations for interacting with the organization object
type OrganizationStore interface {
	Init() error
	Create(name string, email string) (*Organization, error)
	ReadAll(limit int, offset int) ([]*Organization, error)
	ReadOne(name string) (*Organization, error)
	Update(name string, email string) (*Organization, error)
	Delete(name string) error
	GetBackendType() string
}
