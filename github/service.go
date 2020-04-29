package github

import (
	"context"
)
import "golang.org/x/oauth2"
import "github.com/google/go-github/v31/github"

type Service struct {
	client *github.Client
}

func NewGithubService(token string) Service {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return Service{
		client: client,
	}
}

func (s Service) assign(targetGithubUserId string) error {
	return nil
}
