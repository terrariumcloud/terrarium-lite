package github

// GithubSource is a generic struct containing repository properties. This also satisfies the types.SourceData
// interface.
type GithubSource struct {
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Username    string                   `json:"username"`
	RepoURI     string                   `json:"repo_uri"`
	Tags        []*GithubTagDownloadPair `json:"tags"`
}

// GithubTagDownloadPair maps a tag to it's corresponding commit or data. This can then be inspected by
// Terrarium and returned in response to `terraform init` call for download
type GithubTagDownloadPair struct {
	Tag           string `json:"tag"`
	SSHCloneURI   string `json:"ssh_clone_uri"`
	HTTPSCloneURI string `json:"https_clone_uri"`
	Commit        string `json:"commit"`
}

// GetRepoName returns the Github repo name
func (g *GithubSource) GetRepoName() string {
	return g.Name
}

// GetRepoName returns the Github repo description
func (g *GithubSource) GetRepoDescription() string {
	return g.Description
}

// GetRepoOwner returns the Github username that owns the repo
func (g *GithubSource) GetRepoOwner() string {
	return g.Username
}
