package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dylanrhysscott/terrarium/pkg/registry/sources"
	ghlib "github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

// GithubBackend implements the types.SourceProvider and provides source interactions via the Github API.
// It is an abstraction over various API calls required by Terrarium to populate a VCS backed module
// into the registry
type GithubBackend struct {
	client *ghlib.Client
}

// init Accepts a context and token arguments to create an authenticated OAuth 2 client. This is then
// passed into the Github client to ensure any API calls are correctly authenticated
func (g *GithubBackend) init(ctx context.Context, token string) {
	if g.client == nil {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{
				AccessToken: token,
			},
		)
		tc := oauth2.NewClient(ctx, ts)
		ghc := ghlib.NewClient(tc)
		g.client = ghc
	}
}

// FetchVCSSource Accepts a Github OAuth token and a valid Github repo name. The Github API is queried
// and a struct implementing the types.SourceData interface is returned. Typically this will then be
// transformed into a Terraform compliant module document for storage in the database.
func (g *GithubBackend) FetchVCSSource(token string, vcsRepoName string) (sources.SourceData, error) {
	ctx := context.TODO()
	g.init(ctx, token)
	authenticated, _, err := g.client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}
	repo, _, err := g.client.Repositories.Get(ctx, authenticated.GetLogin(), vcsRepoName)
	if err != nil {
		return nil, err
	}
	tags, resp, err := g.client.Repositories.ListTags(ctx, authenticated.GetLogin(), vcsRepoName, &ghlib.ListOptions{})
	if resp.StatusCode != http.StatusNotFound && err != nil {
		return nil, err
	}
	var finalTags []*GithubTagDownloadPair = []*GithubTagDownloadPair{}
	for _, tag := range tags {
		pair := &GithubTagDownloadPair{
			Tag:           tag.GetName(),
			SSHCloneURI:   fmt.Sprintf("git::%s?ref=%s", repo.GetSSHURL(), tag.GetName()),
			HTTPSCloneURI: fmt.Sprintf("git::%s?ref=%s", repo.GetCloneURL(), tag.GetName()),
			Commit:        tag.Commit.GetSHA(),
		}
		finalTags = append(finalTags, pair)
	}
	return &GithubSource{
		Name:        repo.GetName(),
		Description: repo.GetDescription(),
		Username:    repo.GetOwner().GetLogin(),
		RepoURI:     repo.GetHTMLURL(),
		Tags:        finalTags,
	}, nil
}

// NewGithubBacked Creates a new Github backend struct implementing the types.SourceProvider interface
func NewGithubBackend() *GithubBackend {
	return &GithubBackend{}
}
