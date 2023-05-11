package firebase

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/iterator"
)

type User struct {
	UID string
}

func (c *Client) GetAllUsers() ([]User, error) {
	var allExportedUsers []auth.ExportedUserRecord

	pager := iterator.NewPager(c.auth.Users(context.Background(), ""), 100, "")

	for {
		var exportedUsers []*auth.ExportedUserRecord

		nextPageToken, err := pager.NextPage(&exportedUsers)
		if err != nil {
			return nil, err
		}

		for _, u := range exportedUsers {
			allExportedUsers = append(allExportedUsers, *u)
		}

		if nextPageToken == "" {
			break
		}
	}

	var users []User

	for _, exportedUser := range allExportedUsers {
		user := User{
			UID: exportedUser.UID,
		}

		users = append(users, user)
	}

	return users, nil
}
