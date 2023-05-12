package config

import (
	"fmt"
	"github.com/inczero/token-janitor/internal/secret"
	"os"
)

const (
	SecretManagerCredential = "SM_CREDENTIAL"
	FirebaseDatabaseURL     = "FIREBASE_DB_URL"

	// SMSecretFirebaseServiceAccount - name of Firebase service account credential secret inside Secret Manager
	SMSecretFirebaseServiceAccount = "firebase-service-account-credentials"
)

type Config struct {
	FirebaseDbURL  string
	FirebaseSACred string // Firebase Service Account Credential JSON
}

func NewConfig() (*Config, error) {
	smCredBase64, isSet1 := os.LookupEnv(SecretManagerCredential)
	if !isSet1 {
		return nil, fmt.Errorf("environment variable '%s' not set", SecretManagerCredential)
	}

	firebaseDbURL, isSet2 := os.LookupEnv(FirebaseDatabaseURL)
	if !isSet2 {
		return nil, fmt.Errorf("environment variable '%s' not set", FirebaseDatabaseURL)
	}

	smClient, newErr := secret.NewClient(smCredBase64)
	if newErr != nil {
		return nil, newErr
	}

	firebaseCred, getErr := smClient.GetSecretData(SMSecretFirebaseServiceAccount)
	if getErr != nil {
		return nil, getErr
	}

	config := Config{
		FirebaseDbURL:  firebaseDbURL,
		FirebaseSACred: firebaseCred,
	}

	return &config, nil
}
