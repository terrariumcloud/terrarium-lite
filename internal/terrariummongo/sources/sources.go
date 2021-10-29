package sources

type SourceVCSRepo struct {
	VCSUser string `json:"vcs_user"`
	Branch  string `json:"branch"`
	Repo    string `json:"repo"`
}
