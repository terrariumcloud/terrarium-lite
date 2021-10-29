package types

// SourceStore is a generic data interface for implementaing database operations relating to modules
type SourceStore interface {
	FetchVCSSources()
}

// Module represents the module data structure stored in the database
type Module struct {
	ID           string         `json:"_id"`
	Name         string         `json:"name"`
	Namespace    string         `json:"namespace"`
	Provider     string         `json:"provider"`
	Organization *ResourceLink  `json:"organization"`
	VCS          *ResourceLink  `json:"vcs"`
	VCSRepo      *SourceVCSRepo `json:"vcs_repo"`
}

type SourceVCSRepo struct {
	VCSUser string `json:"vcs_user"`
	Branch  string `json:"branch"`
	Repo    string `json:"repo"`
}
