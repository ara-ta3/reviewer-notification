package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ara-ta3/reviewer-notification/github"
	"github.com/ara-ta3/reviewer-notification/slack"
)

var logger = log.New(os.Stdout, "", log.Ldate+log.Ltime+log.Lshortfile)

// func readRequestBody() {
//     body, _ := ioutil.ReadAll(resp.Body)
//     defer resp.Body.Close()
//     r := response{}
//     json.Unmarshal(body, &r)
//     c := 0
// }

func main() {
	n := ReviewerNotification{
		s: slack.SlackClient{
			token: "",
		},
		logger: *logger,
	}
	h := GithubNotificationHandler{n: n}
	http.Handle("/", h)
	http.ListenAndServe(":8080", nil)
}

type GithubNotificationHandler struct {
	n ReviewerNotification
}

func (h GithubNotificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	g := GithubEvent{}
	e := h.n.Notify(g)
	if e != nil {
		fmt.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type ReviewerNotification struct {
	s      slack.SlackClient
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
