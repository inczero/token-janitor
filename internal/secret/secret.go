package secret

import (
	sm "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"context"
	"fmt"
	"google.golang.org/api/option"
)

const (
	// SMSecretPath - Secret Manager path to latest secret version
	SMSecretPath = "projects/%s/secrets/%s/versions/latest" // %s - project id, %s - secret name
)

type Client struct {
	client      *sm.Client
	projectName string
}

func NewClient(projectName string, credentials string) (*Client, error) {
	opt := option.WithCredentialsJSON([]byte(credentials))

	smClient, err := sm.NewClient(context.Background(), opt)
	if err != nil {
		return nil, err
	}

	client := Client{
		client:      smClient,
		projectName: projectName,
	}

	return &client, nil
}

func (c *Client) GetSecretDataLatest(secretName string) ([]byte, error) {
	secretPath := fmt.Sprintf(SMSecretPath, c.projectName, secretName)

	versionReq := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secretPath,
	}

	result, err := c.client.AccessSecretVersion(context.Background(), versionReq)
	if err != nil {
		return nil, err
	}

	return result.Payload.Data, nil
}
