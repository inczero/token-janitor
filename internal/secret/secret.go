package secret

import (
	sm "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"context"
	"encoding/base64"
	"fmt"
	"google.golang.org/api/option"
)

type Client struct {
	client *sm.Client
}

func NewClient(credentialsBase64 string) (*Client, error) {
	// decode credentials JSON
	credJSON, decErr := base64.StdEncoding.DecodeString(credentialsBase64)
	if decErr != nil {
		return nil, decErr
	}

	opt := option.WithCredentialsJSON(credJSON)

	smClient, err := sm.NewClient(context.Background(), opt)
	if err != nil {
		return nil, err
	}

	client := new(Client)
	client.client = smClient

	return client, nil
}

func (c *Client) GetSecretData(secretName string) (string, error) {
	// TODO: maybe a different name is needed -> documentation takes version.name
	versionReq := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secretName,
	}

	result, err := c.client.AccessSecretVersion(context.Background(), versionReq)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", result.Payload.Data), nil
}
