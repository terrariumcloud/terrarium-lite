package sources

type SourceVCSRepo struct {
	VCSUser string `json:"vcs_user"`
	Branch  string `json:"branch"`
	Repo    string `json:"repo"`
}

// SourceBackend is a struct that implements source interactions for module documents
type SourceBackend struct{}

func (s *SourceBackend) FetchVCSSources() {

}
