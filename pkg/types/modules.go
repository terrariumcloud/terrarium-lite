package types

// VCSModule represents a VCS Backed Module stored in the database
type VCSModule struct {
	ID            string        `json:"_id"`
	Name          string        `json:"name"`
	Provider      string        `json:"provider"`
	Description   string        `json:"description"`
	VCSConnection *ResourceLink `json:"vcs_connection"`
	Organization  *ResourceLink `json:"organization"`
	VCSRepo       SourceData    `json:"vcs_repo"`
}

type Module struct {
	ID           string        `json:"_id"`
	Name         string        `json:"name"`
	Provider     string        `json:"provider"`
	Version      string        `json:"version"`
	Description  string        `json:"description"`
	Source       string        `json:"source"`
	Organization *ResourceLink `json:"organization"`
}
