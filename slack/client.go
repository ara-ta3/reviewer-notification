package slack

type SlackClient struct {
	WebhookURL string
}

func (s SlackClient) Send(reviewer, message string) error {
	return nil
}
