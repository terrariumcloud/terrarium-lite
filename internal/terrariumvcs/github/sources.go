package github

import (
	"context"
	"log"

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
		repos, _, err := g.client.Repositories.List(ctx, "", &ghlib.RepositoryListOptions{
			Visibility: "all",
		})
		if err != nil {
			log.Println(err.Error())
			return
		}
		for _, repo := range repos {
			log.Println(repo.GetName())
		}
	}
}

func NewGithubBackend() *GithubBackend {
	return &GithubBackend{}
}
