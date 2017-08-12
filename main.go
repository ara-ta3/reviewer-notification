package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ara-ta3/reviewer-notification/service"
	"github.com/ara-ta3/reviewer-notification/slack"
)

var logger = log.New(os.Stdout, "", log.Ldate+log.Ltime+log.Lshortfile)

func main() {
	u := os.Getenv("SLACK_WEBHOOK_URL")
	token := os.Getenv("TOKEN")
	labels := strings.Split(os.Getenv("TARGET_LABELS"), ",")
	p := os.Getenv("PORT")
	accountMap := parseAccountMap(os.Getenv("ACCOUNT_MAP"))
	slackChannel := os.Getenv("SLACK_CHANNEL")
	if p == "" {
		p = "80"
	}
	logger.Printf("target labels: %#v\n", labels)
	logger.Printf("port: %#v\n", p)
	logger.Printf("slack channel id: %#v\n", slackChannel)
	logger.Printf("account map: %#v\n", accountMap)

	h := service.GithubNotificationHandler{
		NotificationService: service.NewReviewerNotification(
			slack.NewSlackClient(u, slackChannel),
			token,
			labels,
			logger,
			accountMap,
		),
		Logger: *logger,
	}
	http.Handle("/", h)
	http.ListenAndServe(fmt.Sprintf(":%s", p), nil)
}

func parseAccountMap(s string) map[string]string {
	ms := strings.Split(s, ",")
	r := map[string]string{}
	for _, m := range ms {
		x := strings.Split(m, ":")
		r[x[0]] = x[1]
	}
	return r
}
