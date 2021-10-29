package types

import (
	"github.com/dylanrhysscott/terrarium/internal/terrariummongo/sources"
	"github.com/dylanrhysscott/terrarium/internal/terrariummongo/vcsconn"
)

// Module represents the module data structure stored in the database
type Module struct {
	ID           string                 `json:"_id"`
	Name         string                 `json:"name"`
	Namespace    string                 `json:"namespace"`
	Provider     string                 `json:"provider"`
	Organization *vcsconn.ResourceLink  `json:"organization"`
	VCS          *vcsconn.ResourceLink  `json:"vcs"`
	VCSRepo      *sources.SourceVCSRepo `json:"vcs_repo"`
}
