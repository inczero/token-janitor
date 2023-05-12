package firebase

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"fmt"
)

const (
	// Firebase Cloud Messaging error messages
	FCMTokenUnregisteredError    = "FCM registration token was unregistered"
	FCMTokenInvalidArgumentError = "FCM registration token is invalid"
)

func (c *Client) sendMessageDryRun(token string) error {
	message := &messaging.Message{
		Data: map[string]string{
			"test": "dry-run",
		},
		Token: token,
	}

	_, err := c.msg.SendDryRun(context.Background(), message)
	return err
}

func (c *Client) IsTokenActive(token string) (bool, error) {
	err := c.sendMessageDryRun(token)
	if err != nil {
		if messaging.IsInvalidArgument(err) {
			return false, fmt.Errorf(FCMTokenInvalidArgumentError)
		}

		if messaging.IsUnregistered(err) {
			return false, fmt.Errorf(FCMTokenUnregisteredError)
		}

		return false, err
	}

	return true, nil
}
