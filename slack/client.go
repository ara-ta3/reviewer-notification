package slack

type SlackClient struct {
	token string
}

func (s SlackClient) Send(reviewer, message string) error {
	return nil
}
