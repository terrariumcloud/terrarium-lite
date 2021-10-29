package types

// ModuleStore is a generic data interface for implementaing database operations relating to modules
type ModuleStore interface {
	Init() error
	Create() (*Module, error)
	ReadAll(limit int, offset int) ([]*Module, error)
	ReadOne() (*Module, error)
	Update() (*Module, error)
	Delete() error
}

// Module represents the module data structure stored in the database
type Module struct {
	ID           string         `json:"_id"`
	Name         string         `json:"name"`
	Namespace    string         `json:"namespace"`
	Provider     string         `json:"provider"`
	Organization *ResourceLink  `json:"organization"`
	VCS          *ResourceLink  `json:"vcs"`
	VCSRepo      *ModuleVCSRepo `json:"vcs_repo"`
}

type ModuleVCSRepo struct {
	VCSUser string `json:"vcs_user"`
	Branch  string `json:"branch"`
	Repo    string `json:"repo"`
}
