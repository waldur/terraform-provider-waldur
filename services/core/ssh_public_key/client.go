package ssh_public_key

import (
	"context"
	"fmt"
	"strings"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

type Client struct {
	Client *client.Client
}

func NewClient(c *client.Client) *Client {
	_ = fmt.Errorf
	_ = strings.Split
	return &Client{Client: c}
}

func (c *Client) GetCoreSshPublicKey(ctx context.Context, id string) (*CoreSshPublicKeyResponse, error) {
	var apiResp CoreSshPublicKeyResponse
	err := c.Client.Get(ctx, "/api/keys/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) ListCoreSshPublicKey(ctx context.Context, filter map[string]string) ([]CoreSshPublicKeyResponse, error) {
	var listResult []CoreSshPublicKeyResponse
	err := c.Client.List(ctx, "/api/keys/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
