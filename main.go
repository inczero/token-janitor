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
	i := 1

	for id, token := range tokens {
		fmt.Printf("%d. token with id '%s':\n", i, id)

		if err := client.DetectInvalidToken(token.Token); err == nil {
			fmt.Printf("\tstatus: %s\n", firebase.FCMTokenActive)

			if isOld := isOlderThan30Days(token.CreatedOn); isOld {
				if token.Deprecated {
					if token.Rotated {
						fmt.Printf("\ttoken was rotated, but it is still active - needs checking\n")
					} else {
						fmt.Printf("\ttoken already marked as deprecated\n")
					}
				} else {
					fmt.Printf("\ttoken was created more than 30 days ago\n")
					fmt.Printf("\tmarking as deprecated...\n")

					if setErr := client.SetUserRTDeprecated(uid, id, true); setErr != nil {
						fmt.Printf("\tcould not mark token as deprecated - error: %s\n", setErr.Error())
					} else {
						fmt.Printf("\tdone\n")
					}
				}
			} else {
				fmt.Printf("\ttoken was created in the last 30 days\n")
			}
		} else {
			if err.Error() == firebase.FCMTokenInvalidArgumentError || err.Error() == firebase.FCMTokenUnregisteredError {
				fmt.Printf("\tstatus: %s\n", err.Error())
				fmt.Printf("\tinfo: deprecated-%t rotated-%t\n", token.Deprecated, token.Rotated)
				fmt.Printf("\tdeleting...\n")

				if delErr := client.DeleteUserRT(uid, id); delErr != nil {
					fmt.Printf("\tcould not delete token - error: %s\n", delErr.Error())
				} else {
					fmt.Printf("\tdone\n")
				}
			} else {
				return err
			}
		}

		i++
	}

	return nil
}

func isOlderThan30Days(createdOn int64) bool {
	creationTime := time.Unix(createdOn, 0)
	deprecationTime := creationTime.AddDate(0, 0, 30)

	return deprecationTime.Before(time.Now())
}
