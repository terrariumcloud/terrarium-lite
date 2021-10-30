package github

type GithubSource struct {
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Username    string                   `json:"username"`
	RepoURI     string                   `json:"repo_uri"`
	Tags        []*GithubTagDownloadPair `json:"tags"`
}

type GithubTagDownloadPair struct {
	Tag            string `json:"tag"`
	TarDownloadURI string `json:"tar_download_uri"`
	ZipDownloadURI string `json:"zip_download_uri"`
	Commit         string `json:"commit"`
}

func (g *GithubSource) GetRepoName() string {
	return g.Name
}

func (g *GithubSource) GetRepoDescription() string {
	return g.Description
}

func (g *GithubSource) GetRepoOwner() string {
	return g.Username
}
