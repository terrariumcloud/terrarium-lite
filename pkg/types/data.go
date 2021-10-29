package types

import (
	"github.com/dylanrhysscott/terrarium/internal/terrariummongo/vcsconn"
)

// Module represents the module data structure stored in the database
type Module struct {
	ID            string                `json:"_id"`
	Owner         string                `json:"owner"`
	Name          string                `json:"name"`
	Namespace     string                `json:"namespace"`
	Version       string                `json:"version"`
	Provider      string                `json:"provider"`
	Description   string                `json:"description"`
	Source        string                `json:"source"`
	Tag           string                `json:"tag"`
	Organization  *vcsconn.ResourceLink `json:"organization"`
	VCSConnection *vcsconn.ResourceLink `json:"vcs_connection"`
	VCSRepo       SourceData            `json:"vcs_repo"`
}

type ModuleVersion struct {
	Source   string     `json:"source"`
	Versions []*Version `json:"version"`
}

type Version struct {
	Version     string `json:"version"`
	DownloadURI string `json:"download_uri"`
}
