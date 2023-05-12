package firebase

import (
	"context"
	"fmt"
)

const (
	PathToRegistrationTokens          = "/users/%s/fcmRegistrationTokens"               // %s - user's id
	PathToRegistrationTokenDeprecated = "/users/%s/fcmRegistrationTokens/%s/deprecated" // %s - user's id, %s - number of token
	PathToRegistrationTokenRotated    = "/users/%s/fcmRegistrationTokens/%s/rotated"    // %s - user's id, %s - number of token
)

type RegistrationToken struct {
	CreatedOn  int64  `json:"createdOn"`
	Deprecated bool   `json:"deprecated"`
	Rotated    bool   `json:"rotated"`
	Token      string `json:"token"`
}

func (c *Client) GetUserRTs(uid string) (map[string]RegistrationToken, error) {
	path := fmt.Sprintf(PathToRegistrationTokens, uid)
	ref := c.db.NewRef(path)

	var tokens map[string]RegistrationToken
	if err := ref.Get(context.Background(), &tokens); err != nil {
		return nil, err
	}

	return tokens, nil
}

func (c *Client) SetUserRTDeprecated(uid string, id string, deprecated bool) error {
	path := fmt.Sprintf(PathToRegistrationTokenDeprecated, uid, id)
	ref := c.db.NewRef(path)

	if err := ref.Set(context.Background(), deprecated); err != nil {
		return err
	}

	return nil
}

func (c *Client) SetUserRTRotated(uid string, id string, rotated bool) error {
	path := fmt.Sprintf(PathToRegistrationTokenRotated, uid, id)
	ref := c.db.NewRef(path)

	if err := ref.Set(context.Background(), rotated); err != nil {
		return err
	}

	return nil
}
