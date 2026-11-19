package resource_api_key_ephemeral

import (
	"context"
	"fmt"
	"strings"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type MarketplaceResourceApiKeyClient struct {
	Client *client.Client
}

func NewMarketplaceResourceApiKeyClient(c *client.Client) *MarketplaceResourceApiKeyClient {
	return &MarketplaceResourceApiKeyClient{Client: c}
}

func (c *MarketplaceResourceApiKeyClient) Configure(ctx context.Context, providerData interface{}) error {
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

// Open calls the operation behind the ephemeral resource.
func (c *MarketplaceResourceApiKeyClient) Open(ctx context.Context, uuidParam string) (*MarketplaceResourceApiKeyResponse, error) {
	path := "/api/marketplace-resource-api-keys/{uuid}/reveal/"
	path = strings.Replace(path, "{uuid}", uuidParam, 1)
	var apiResp MarketplaceResourceApiKeyResponse
	err := c.Client.Post(ctx, path, nil, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
