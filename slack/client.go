package slack

type SlackClient struct {
	WebhookURL string
	PostTo     string
}

func (s SlackClient) Send(reviewers []string, message string) error {
	return nil
}
