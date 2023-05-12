package firebase

import (
	"context"
	"firebase.google.com/go/v4/messaging"
)

func (c *Client) SendMessageDryRun(token string) error {
	message := &messaging.Message{
		Data: map[string]string{
			"test": "dry-run",
		},
		Token: token,
	}

	// TODO: do some validation on the response
	_, err := c.msg.SendDryRun(context.Background(), message)
	if err != nil {
		return err
	}

	return nil
}
