package github

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dylanrhysscott/terrarium/pkg/types"
	ghlib "github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

// Github is a struct that implements source interactions for module documents
type GithubBackend struct {
	client *ghlib.Client
}

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

func (g *GithubBackend) FetchVCSSource(token string, vcsRepoName string, vcsRepoOwner string) (types.SourceData, error) {
	ctx := context.TODO()
	g.init(ctx, token)
	repo, _, err := g.client.Repositories.Get(ctx, vcsRepoOwner, vcsRepoName)
	if err != nil {
		return nil, err
	}
	tags, resp, err := g.client.Repositories.ListTags(ctx, vcsRepoOwner, vcsRepoName, &ghlib.ListOptions{})
	if resp.StatusCode != http.StatusNotFound && err != nil {
		return nil, err
	}
	var finalTags []*GithubTagDownloadPair = []*GithubTagDownloadPair{}
	for _, tag := range tags {
		pair := &GithubTagDownloadPair{
			Tag:            tag.GetName(),
			TarDownloadURI: fmt.Sprintf("%s/%s", tag.GetTarballURL(), "/*?archive=tar.gz"),
			ZipDownloadURI: fmt.Sprintf("%s/%s", tag.GetZipballURL(), "/*?archive=zip"),
			Commit:         tag.Commit.GetSHA(),
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

func (g *GithubBackend) FetchVCSSources(token string) {
	ctx := context.TODO()
	g.init(ctx, token)
	repos, _, err := g.client.Repositories.List(ctx, "", &ghlib.RepositoryListOptions{})

	if err != nil {
		log.Println(err.Error())
	}
	for _, repo := range repos {
		log.Println(repo.GetName())
		tags, resp, err := g.client.Repositories.ListTags(ctx, "", *repo.Name, &ghlib.ListOptions{})
		if resp.StatusCode != http.StatusNotFound {
			if err != nil {
				log.Println(err.Error())
			}
		}

		if len(tags) == 0 {
			log.Println("No tags")
		} else {
			for _, tag := range tags {
				log.Print(tag.GetName())
			}
		}
	}
}

func NewGithubBackend() *GithubBackend {
	return &GithubBackend{}
}
