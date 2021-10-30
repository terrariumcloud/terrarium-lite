package modules

import (
	"github.com/dylanrhysscott/terrarium/internal/terrariummongo/relationships"
	"github.com/dylanrhysscott/terrarium/pkg/types"
)

// Module represents the module data structure stored in the database
type VCSModule struct {
	ID            string                      `json:"_id"`
	Name          string                      `json:"name"`
	Provider      string                      `json:"provider"`
	Description   string                      `json:"description"`
	VCSConnection *relationships.ResourceLink `json:"vcs_connection"`
	Organization  *relationships.ResourceLink `json:"organization"`
	VCSRepo       types.SourceData            `json:"vcs_repo"`
}
