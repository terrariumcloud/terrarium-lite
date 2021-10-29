package sources

import "errors"

type SourceVCSRepoBody struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Provider    string   `json:"provider"`
	Repo        string   `json:"repo"`
	Owner       string   `json:"owner"`
	Tags        []string `json:"tags"`
}

func (s *SourceVCSRepoBody) Validate() error {
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.Provider == "" {
		return errors.New("provider is required")
	}
	if s.Repo == "" {
		return errors.New("repo is required")
	}
	if s.Owner == "" {
		return errors.New("owner is required")
	}
	return nil
}
