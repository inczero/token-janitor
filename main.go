package main

import (
	"fmt"
	"github.com/inczero/token-janitor/internal/firebase"
	"log"
)

const (
	FirebasePathToKey   = "path/to/key"
	FirebaseDatabaseURL = "db-url"
)

func main() {
	fmt.Println("Starting token-janitor...")

	client, initErr := firebase.InitClient(FirebaseDatabaseURL, make([]byte, 0))
	if initErr != nil {
		log.Fatalln(initErr)
	}
}
