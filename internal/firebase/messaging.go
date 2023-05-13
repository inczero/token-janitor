package firebase

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"fmt"
)

const (
	// Firebase Cloud Messaging messages
	FCMTokenActive               = "FCM registration token is active"
	FCMTokenUnregisteredError    = "FCM registration token was unregistered"
	FCMTokenInvalidArgumentError = "FCM registration token is invalid"
)

// sendMessageDryRun function sends a message in dry run mode (every operation is done on the backend except the
// message sending at the end) to the client app instance from which the Firebase Cloud Messaging Registration
// Token was generated.
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

// DetectInvalidToken function determines if a Firebase Cloud Messaging Registration Token is still valid by sending
// a message in dry run mode.
func (c *Client) DetectInvalidToken(token string) error {
	err := c.sendMessageDryRun(token)
	if err != nil {
		if messaging.IsInvalidArgument(err) {
			return fmt.Errorf(FCMTokenInvalidArgumentError)
		}

		if messaging.IsUnregistered(err) {
			return fmt.Errorf(FCMTokenUnregisteredError)
		}

		return err
	}

	return nil
}
