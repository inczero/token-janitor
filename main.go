package main

import (
	"fmt"
	"github.com/inczero/token-janitor/internal/config"
	"github.com/inczero/token-janitor/internal/firebase"
	"log"
	"time"
)

func main() {
	fmt.Printf("Starting token-janitor...\n\n")

	fmt.Println("Setting up configuration...")
	conf, newErr := config.NewConfig()
	if newErr != nil {
		log.Fatalln(newErr)
	}
	fmt.Printf("OK\n\n")

	fmt.Println("Initializing Firebase clients...")
	client, initErr := firebase.InitClient(conf.FirebaseDbURL, conf.FirebaseSACred)
	if initErr != nil {
		log.Fatalln(initErr)
	}
	fmt.Printf("OK\n\n")

	fmt.Println("Getting users...")
	users, getErr := client.GetAllUsers()
	if getErr != nil {
		log.Fatalln(getErr)
	}
	fmt.Printf("OK\n\n")

	fmt.Println("Checking registration tokens...")
	for _, user := range users {
		tokens, err := client.GetUserRTs(user.UID)
		if err != nil {
			log.Fatalln(err)
		}

		if checkErr := checkTokens(client, user.UID, tokens); checkErr != nil {
			log.Fatalln(checkErr)
		}
	}
	fmt.Printf("OK\n\n")

	fmt.Printf("Finished at %s\n", time.Now())
}

func checkTokens(client *firebase.Client, uid string, tokens map[string]firebase.RegistrationToken) error {
	// TODO: fix logging messages + handle different type of errors
	for id, token := range tokens {
		if token.Rotated {
			if token.Deprecated {
				// rotated && deprecated -> check if active -> delete
				active, err := client.IsTokenActive(token.Token)
				if err != nil {
					log.Printf("error happened - log it")
				}

				if active {
					// log error
					log.Printf("deprecated and rotated token is still active")
				} else {
					// delete the deprecated+rotated+inactive token
					if delErr := client.DeleteUserRT(uid, id); delErr != nil {
						log.Printf("error happened during deletion - log it")
					} else {
						log.Printf("deleted successfully")
					}
				}
			} else {
				// rotated && !deprecated -> log error
				log.Printf("token was rotated but it isn't deprecated")
			}
		} else {
			if token.Deprecated {
				// !rotated && deprecated -> log that the token needs to be rotated
				log.Printf("token is older than 30 days - needs to be rotated")
			} else {
				// !rotated && !deprecated -> Older than 30 days? -> check if active -> log it
				if isOlderThan30Days(token.CreatedOn) {
					if setErr := client.SetUserRTDeprecated(uid, id, true); setErr != nil {
						log.Printf("token couldn't be set deprecated - error happened")
					}
				} else {
					active, err := client.IsTokenActive(token.Token)
					if err != nil {
						log.Printf("error happened - log it")
					} else {
						if active {
							log.Printf("token is active - everything is cool")
						} else {
							log.Printf("token is inactive - deleting...")

							if delErr := client.DeleteUserRT(uid, id); delErr != nil {
								log.Printf("error happened during deletion - log it")
							}
						}
					}
				}
			}
		}
	}

	return nil
}

func isOlderThan30Days(createdOn int64) bool {
	creationTime := time.Unix(createdOn, 0)
	deprecationTime := creationTime.AddDate(0, 0, 30)

	return deprecationTime.Before(time.Now())
}
