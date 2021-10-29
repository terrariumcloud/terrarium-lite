package github

import (
	"context"
	"log"
	"net/http"

	ghlib "github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

// Github is a struct that implements source interactions for module documents
type GithubBackend struct {
	client *ghlib.Client
}

func (g *GithubBackend) FetchVCSSources(token string) {
	ctx := context.TODO()
	log.Println("Token:", token)
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
