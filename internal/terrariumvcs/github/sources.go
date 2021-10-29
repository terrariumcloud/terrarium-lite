package github

import "log"

// Github is a struct that implements source interactions for module documents
type GithubBackend struct{}

func (s *GithubBackend) FetchVCSSources() {
	log.Println("Fetch GH Sources")
}

func NewGithubBackend() *GithubBackend {
	return &GithubBackend{}
}
