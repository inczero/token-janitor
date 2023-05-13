package firebase

import (
	"context"
	"fmt"
)

const (
	PathToRegistrationTokens          = "/users/%s/fcmRegistrationTokens"               // %s - user's id
	PathToRegistrationTokenDeprecated = "/users/%s/fcmRegistrationTokens/%s/deprecated" // %s - user's id, %s - number of token
	PathToRegistrationTokenRotated    = "/users/%s/fcmRegistrationTokens/%s/rotated"    // %s - user's id, %s - number of token
	PathToRegistrationToken           = "/users/%s/fcmRegistrationTokens/%s"            // %s - user's id, %s - number of token
)

type RegistrationToken struct {
	CreatedOn  int64  `json:"createdOn"`
	Deprecated bool   `json:"deprecated"`
	Rotated    bool   `json:"rotated"`
	Token      string `json:"token"`
}

// GetUserRTs function returns a map which contains the Firebase Cloud Messaging Registration Tokens that belong to a
// user along with other metadata. The keys are UUIDs generated when the tokens are added to the database.
func (c *Client) GetUserRTs(uid string) (map[string]RegistrationToken, error) {
	path := fmt.Sprintf(PathToRegistrationTokens, uid)
	ref := c.db.NewRef(path)

	var tokens map[string]RegistrationToken
	if err := ref.Get(context.Background(), &tokens); err != nil {
		return nil, err
	}

	return tokens, nil
}

// SetUserRTDeprecated function sets the deprecated field of a user's Firebase Cloud Messaging Registration Token. By
// setting this to true, it marks the token for rotation which will be done by application on the user's smartphone.
func (c *Client) SetUserRTDeprecated(uid string, id string, deprecated bool) error {
	path := fmt.Sprintf(PathToRegistrationTokenDeprecated, uid, id)
	ref := c.db.NewRef(path)

	if err := ref.Set(context.Background(), deprecated); err != nil {
		return err
	}

	return nil
}

// SetUserRTRotated function sets the rotated field of a user's Firebase Cloud Messaging Registration Token.
func (c *Client) SetUserRTRotated(uid string, id string, rotated bool) error {
	path := fmt.Sprintf(PathToRegistrationTokenRotated, uid, id)
	ref := c.db.NewRef(path)

	if err := ref.Set(context.Background(), rotated); err != nil {
		return err
	}

	return nil
}

// DeleteUserRT function deletes the Firebase Cloud Messaging Registration Token entry from the database.
func (c *Client) DeleteUserRT(uid string, id string) error {
	path := fmt.Sprintf(PathToRegistrationToken, uid, id)
	ref := c.db.NewRef(path)

	if err := ref.Delete(context.Background()); err != nil {
		return err
	}

	return nil
}
