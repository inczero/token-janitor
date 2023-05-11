package firebase

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
	"log"
)

type Client struct {
	app  *firebase.App
	auth *auth.Client
	db   *db.Client
}

func InitClient(databaseURL string, credentialsJSON []byte) (*Client, error) {
	client := new(Client)

	conf := &firebase.Config{
		DatabaseURL: databaseURL,
	}

	opt := option.WithCredentialsJSON(credentialsJSON)

	app, newErr := firebase.NewApp(context.Background(), conf, opt)
	if newErr != nil {
		log.Fatal(newErr)
	}

	client.app = app

	authClient, authErr := app.Auth(context.Background())
	if authErr != nil {
		return nil, authErr
	}

	client.auth = authClient

	dbClient, dbErr := app.Database(context.Background())
	if dbErr != nil {
		return nil, dbErr
	}

	client.db = dbClient

	return client, nil
}
