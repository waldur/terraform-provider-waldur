package resource_api_key

import (
	"context"
	"fmt"

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

func (c *MarketplaceResourceApiKeyClient) Get(ctx context.Context, id string) (*MarketplaceResourceApiKeyResponse, error) {
	var apiResp MarketplaceResourceApiKeyResponse
	err := c.Client.Get(ctx, "/api/marketplace-resource-api-keys/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *MarketplaceResourceApiKeyClient) List(ctx context.Context, filter map[string]string) ([]MarketplaceResourceApiKeyResponse, error) {
	var listResult []MarketplaceResourceApiKeyResponse
	err := c.Client.List(ctx, "/api/marketplace-resource-api-keys/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

// ListAll returns every match of the filter, across all pages.
func (c *MarketplaceResourceApiKeyClient) ListAll(ctx context.Context, filter map[string]string) ([]MarketplaceResourceApiKeyResponse, error) {
	var listResult []MarketplaceResourceApiKeyResponse
	err := c.Client.ListAll(ctx, "/api/marketplace-resource-api-keys/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
