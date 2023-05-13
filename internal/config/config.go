package config

import (
	"fmt"
	"github.com/inczero/token-janitor/internal/secret"
	"os"
)

const (
	SecretManagerCredential = "SM_CREDENTIAL"
	FirebaseDatabaseURL     = "FIREBASE_DB_URL"
	GoogleCloudProjectId    = "GCP_PROJECT_ID"

	// SMSecretFirebaseServiceAccount - name of Firebase service account credential secret inside Secret Manager
	SMSecretFirebaseServiceAccount = "firebase-service-account-credentials"
)

type Config struct {
	FirebaseDbURL  string // Firebase Realtime Database URL
	FirebaseSACred []byte // Firebase Service Account Credential JSON
}

func NewConfig() (*Config, error) {
	smCred, isSet1 := os.LookupEnv(SecretManagerCredential)
	if !isSet1 {
		return nil, fmt.Errorf("environment variable '%s' not set", SecretManagerCredential)
	}

	firebaseDbURL, isSet2 := os.LookupEnv(FirebaseDatabaseURL)
	if !isSet2 {
		return nil, fmt.Errorf("environment variable '%s' not set", FirebaseDatabaseURL)
	}

	gcpProjectId, isSet3 := os.LookupEnv(GoogleCloudProjectId)
	if !isSet3 {
		return nil, fmt.Errorf("environment variable '%s' not set", GoogleCloudProjectId)
	}

	smClient, newErr := secret.NewClient(gcpProjectId, smCred)
	if newErr != nil {
		return nil, newErr
	}

	firebaseCred, getErr := smClient.GetSecretDataLatest(SMSecretFirebaseServiceAccount)
	if getErr != nil {
		return nil, getErr
	}

	config := Config{
		FirebaseDbURL:  firebaseDbURL,
		FirebaseSACred: firebaseCred,
	}

	return &config, nil
}
