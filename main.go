package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"google.golang.org/api/option"
	"log"
)

const (
	FirebasePathToKey   = "path/to/file"
	FirebaseDatabaseURL = "dburl"
)

func main() {
	cont := context.Background()

	conf := &firebase.Config{
		DatabaseURL: FirebaseDatabaseURL,
	}

	opt := option.WithCredentialsFile(FirebasePathToKey)

	app, newErr := firebase.NewApp(cont, conf, opt)
	if newErr != nil {
		log.Fatal(newErr)
	}

	client, dbErr := app.Database(cont)
	if dbErr != nil {
		log.Fatalln("Error initializing database client:", dbErr)
	}

	ref := client.NewRef("/ref/path")
	var data map[string]interface{}
	if err := ref.Get(cont, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	fmt.Println(data)
}
