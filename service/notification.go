package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/ara-ta3/reviewer-notification/github"
	"github.com/ara-ta3/reviewer-notification/slack"
)

func NewReviewerNotification(
	client slack.SlackClient,
	token string,
	labels []string,
	logger *log.Logger,
	accountMap map[string]string,
	service *github.Service,
) ReviewerNotification {
	return ReviewerNotification{
		s:             client,
		token:         token,
		labels:        labels,
		logger:        *logger,
		accountMap:    accountMap,
		githubService: service,
	}
}

type ReviewerNotification struct {
	s             slack.SlackClient
	token         string
	labels        []string
	logger        log.Logger
	accountMap    map[string]string
	githubService *github.Service
}

func (n ReviewerNotification) NotifyWithRequestBody(body io.ReadCloser, event string) error {
	// TODO token verification
	bs, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if event == "issue" {
		e := new(github.IssueEvent)
		err = json.Unmarshal(bs, e)
		if err != nil {
			return err
		}
		return n.NotifyIssue(e)
	} else if event == "pull_request" {
		e := new(github.PullRequestEvent)
		err = json.Unmarshal(bs, e)
		if err != nil {
			return err
		}
		err := n.NotifyPullRequest(e)
		if n.githubService == nil || err != nil {
			return err
		}

		return n.Assign(e)
	}
	return fmt.Errorf("notification for %s is not implemented", event)
}

func (n ReviewerNotification) Assign(g *github.PullRequestEvent) error {
	as := g.GetAssigneeNames()
	return n.githubService.Assign(
		as,
		string(g.Repository.Owner.ID),
		g.Repository.Name,
		g.PullRequest.Number,
	)
}

func (n ReviewerNotification) NotifyIssue(g *github.IssueEvent) error {
	n.logger.Printf("Action: %s\n", g.Action)
	if g.Action != "labeled" {
		return nil
	}
	if !n.hasTargetLabel(g.Issue.Labels) {
		return nil
	}
	as := g.GetAssigneeNames()
	replaced := n.replaceAccountName(as)

	t := fmt.Sprintf("<%s|%s>", g.Issue.HTMLURL, g.Issue.Title)
	return n.s.Send(replaced, t, "Review Request!!")
}

func (n ReviewerNotification) NotifyPullRequest(e *github.PullRequestEvent) error {
	n.logger.Printf("Action: %s\n", e.Action)
	if e.Action != "labeled" {
		return nil
	}
	if !Include(n.labels, e.Label.Name) {
		return nil
	}
	as := e.GetAssigneeNames()
	replaced := n.replaceAccountName(as)

	t := fmt.Sprintf("<%s|%s>", e.PullRequest.HTMLURL, e.PullRequest.Title)
	return n.s.Send(replaced, t, "Review Request!!")
}

func (n ReviewerNotification) replaceAccountName(ns []string) []string {
	ss := []string{}
	for _, x := range ns {
		r, ok := n.accountMap[x]
		if ok {
			ss = append(ss, r)
		} else {
			ss = append(ss, x)
		}
	}
	return ss
}

func (n ReviewerNotification) hasTargetLabel(labels []github.Label) bool {
	for _, l := range labels {
		if Include(n.labels, l.Name) {
			return true
		}
	}
	return false
}

func Include(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}
