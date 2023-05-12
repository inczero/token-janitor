package firebase

import (
	"context"
	"firebase.google.com/go/v4/messaging"
)

func (c *Client) sendMessageDryRun(token string) error {
	message := &messaging.Message{
		Data: map[string]string{
			"test": "dry-run",
		},
		Token: token,
	}

	// TODO: response and error will need some additional checking
	_, err := c.msg.SendDryRun(context.Background(), message)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) IsTokenActive(token string) (bool, error) {
	err := c.sendMessageDryRun(token)
	if err != nil {
		// TODO: check error message to determine if it needs to be returned
		return false, nil
	}

	return true, nil
}
