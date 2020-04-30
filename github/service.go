package github

import (
	"context"
)
import "golang.org/x/oauth2"
import "github.com/google/go-github/v31/github"

type Service struct {
	ctx    context.Context
	client *github.Client
}

func NewGithubService(token string) *Service {
	if token == "" {
		return nil
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return &Service{
		ctx:    ctx,
		client: client,
	}
}

func (s Service) Assign(targetGithubUserIds []string, owner, repository string, number int) error {
	_, _, e := s.client.PullRequests.RequestReviewers(
		s.ctx,
		owner,
		repository,
		number,
		github.ReviewersRequest{
			Reviewers: targetGithubUserIds,
		},
	)
	return e
}
