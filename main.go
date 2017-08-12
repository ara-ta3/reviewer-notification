package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	p := os.Getenv("PORT")
	if p == "" {
		p = "80"
	}
	port := fmt.Sprintf(":%s", p)

	h := GithubNotificationHandler{
		NotificationService: ReviewerNotification{
			s: slack.SlackClient{
				WebhookURL: u,
			},
			token:  token,
			labels: labels,
			logger: *logger,
		},
	}
	http.Handle("/", h)
	http.ListenAndServe(port, nil)
}

type GithubNotificationHandler struct {
	NotificationService ReviewerNotification
}

func (h GithubNotificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	e := h.NotificationService.NotifyWithRequestBody(r.Body)
	if e != nil {
		logger.Printf("%#v\n", e)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type ReviewerNotification struct {
	s      slack.SlackClient
	token  string
	labels []string
	logger log.Logger
}

func (n ReviewerNotification) NotifyWithRequestBody(body io.ReadCloser) error {
	bs, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	g := new(github.GithubEvent)
	err = json.Unmarshal(bs, g)
	if err != nil {
		return err
	}
	n.logger.Printf("%#v\n", g)
	return n.Notify(g)
}

func (n ReviewerNotification) Notify(g *github.GithubEvent) error {
	n.logger.Printf("Action: %s\n", g.Action)
	if g.Action != "labeled" {
		return nil
	}
	r := "ara-ta3"
	message := "hogehoge"
	return n.s.Send(r, message)
}
