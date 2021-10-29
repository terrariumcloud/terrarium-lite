package github

import "github.com/dylanrhysscott/terrarium/pkg/types"

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

func (g *GithubSource) ToModuleDocument() *types.Module {
	return nil
}

func (g *GithubSource) GetVersionList() *types.ModuleVersion {
	return nil
}
