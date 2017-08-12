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
) ReviewerNotification {
	return ReviewerNotification{
		s:          client,
		token:      token,
		labels:     labels,
		logger:     *logger,
		accountMap: accountMap,
	}
}

type ReviewerNotification struct {
	s          slack.SlackClient
	token      string
	labels     []string
	logger     log.Logger
	accountMap map[string]string
}

func (n ReviewerNotification) NotifyWithRequestBody(body io.ReadCloser) error {
	// TODO token verification
	bs, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	g := new(github.GithubEvent)
	err = json.Unmarshal(bs, g)
	if err != nil {
		return err
	}
	return n.Notify(g)
}

func (n ReviewerNotification) Notify(g *github.GithubEvent) error {
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
