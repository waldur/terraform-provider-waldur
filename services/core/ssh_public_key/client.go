package ssh_public_key

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type CoreSshPublicKeyClient struct {
	Client *client.Client
}

func NewCoreSshPublicKeyClient(c *client.Client) *CoreSshPublicKeyClient {
	return &CoreSshPublicKeyClient{Client: c}
}

func (c *CoreSshPublicKeyClient) Configure(ctx context.Context, providerData interface{}) error {
	if providerData == nil {
		return nil
	}

	raw, ok := providerData.(*client.Client)
	if !ok {
		return fmt.Errorf("unexpected provider data type: %T", providerData)
	}

	c.Client = raw
	return nil
}

func IsNotFoundError(err error) bool {
	return common.IsNotFoundError(err)
}

func (c *CoreSshPublicKeyClient) Get(ctx context.Context, id string) (*CoreSshPublicKeyResponse, error) {
	var apiResp CoreSshPublicKeyResponse
	err := c.Client.Get(ctx, "/api/keys/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *CoreSshPublicKeyClient) List(ctx context.Context, filter map[string]string) ([]CoreSshPublicKeyResponse, error) {
	var listResult []CoreSshPublicKeyResponse
	err := c.Client.List(ctx, "/api/keys/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
