package slack

import (
	"fmt"
	"strings"

	slack "github.com/monochromegane/slack-incoming-webhooks"
)

func NewSlackClient(webhookURL, postTo string) SlackClient {
	return SlackClient{
		slack: slack.Client{
			WebhookURL: webhookURL,
		},
		postTo: postTo,
	}
}

type SlackClient struct {
	slack  slack.Client
	postTo string
}

func (s SlackClient) Send(reviewers []string, title, message string) error {
	rs := strings.Join(reviewers, ",")
	t := fmt.Sprintf("%s, %s", rs, message)
	p := slack.Payload{
		Channel:   s.postTo,
		Username:  "ReviewNotification",
		IconEmoji: ":innocent:",
		Attachments: []*slack.Attachment{
			&slack.Attachment{
				Color: "#0000FF",
				Title: title,
				Text:  t,
			},
		},
	}
	return s.slack.Post(&p)
}
