package firebase

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

const (
	PathToRegistrationTokens          = "/users/%s/fcmRegistrationTokens"               // %s - user's id
	PathToRegistrationTokenDeprecated = "/users/%s/fcmRegistrationTokens/%s/deprecated" // %s - user's id, %s - number of token
	PathToRegistrationTokenRotated    = "/users/%s/fcmRegistrationTokens/%s/rotated"    // %s - user's id, %s - number of token
)

type RegistrationToken struct {
	CreatedOn  time.Time
	Deprecated bool
	Rotated    bool
	Token      string
}

func (c *Client) GetUserRTs(uid string) (map[int]RegistrationToken, error) {
	path := fmt.Sprintf(PathToRegistrationTokens, uid)
	ref := c.db.NewRef(path)

	var tokens map[int]RegistrationToken
	if err := ref.Get(context.Background(), &tokens); err != nil {
		return nil, err
	}

	return tokens, nil
}

func (c *Client) SetUserRTDeprecated(uid string, number int, deprecated bool) error {
	path := fmt.Sprintf(PathToRegistrationTokenDeprecated, uid, strconv.Itoa(number))
	ref := c.db.NewRef(path)

	if err := ref.Set(context.Background(), deprecated); err != nil {
		return err
	}

	return nil
}

func (c *Client) SetUserRTRotated(uid string, number int, rotated bool) error {
	path := fmt.Sprintf(PathToRegistrationTokenRotated, uid, strconv.Itoa(number))
	ref := c.db.NewRef(path)

	if err := ref.Set(context.Background(), rotated); err != nil {
		return err
	}

	return nil
}
