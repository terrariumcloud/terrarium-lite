package types

import "errors"

type SourceVCSRepoBody struct {
	Name     string `json:"name"`
	Provider string `json:"provider"`
}

func (s *SourceVCSRepoBody) Validate() error {
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.Provider == "" {
		return errors.New("provider is required")
	}
	return nil
}
