package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ara-ta3/reviewer-notification/github"
	"github.com/ara-ta3/reviewer-notification/slack"
)

var logger = log.New(os.Stdout, "", log.Ldate+log.Ltime+log.Lshortfile)

func main() {
	u := os.Getenv("SLACK_WEBHOOK_URL")
	token := os.Getenv("TOKEN")
	labels := strings.Split(os.Getenv("TARGET_LABELS"), ",")

	h := GithubNotificationHandler{
		n: ReviewerNotification{
			s: slack.SlackClient{
				WebhookURL: u,
			},
			token:  token,
			labels: labels,
			logger: *logger,
		},
	}
	http.Handle("/", h)
	http.ListenAndServe(":8080", nil)
}

type GithubNotificationHandler struct {
	n ReviewerNotification
}

func (h GithubNotificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	g := github.GithubEvent{}
	e := h.n.Notify(g)
	if e != nil {
		fmt.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type ReviewerNotification struct {
	s      slack.SlackClient
	token  string
	labels []string
	logger log.Logger
}

func (n ReviewerNotification) Notify(g github.GithubEvent) error {
	n.logger.Printf("Action: %s\n", g.Action)
	if g.Action != "labeled" {
		return nil
	}
	n.logger.Printf("%#v", g.Issue.Assignee)
	r := "ara-ta3"
	message := "hogehoge"
	return n.s.Send(r, message)
}
