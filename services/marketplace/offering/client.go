package offering

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type Client struct {
	Client *client.Client
}

func NewClient(c *client.Client) *Client {
	return &Client{Client: c}
}

func (c *Client) Configure(ctx context.Context, providerData interface{}) error {
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

func (c *Client) GetMarketplaceOffering(ctx context.Context, id string) (*MarketplaceOfferingResponse, error) {
	var apiResp MarketplaceOfferingResponse
	err := c.Client.Get(ctx, "/api/marketplace-public-offerings/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) ListMarketplaceOffering(ctx context.Context, filter map[string]string) ([]MarketplaceOfferingResponse, error) {
	var listResult []MarketplaceOfferingResponse
	err := c.Client.List(ctx, "/api/marketplace-public-offerings/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
