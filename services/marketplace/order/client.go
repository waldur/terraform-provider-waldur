package order

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type MarketplaceOrderClient struct {
	Client *client.Client
}

func NewMarketplaceOrderClient(c *client.Client) *MarketplaceOrderClient {
	return &MarketplaceOrderClient{Client: c}
}

func (c *MarketplaceOrderClient) Configure(ctx context.Context, providerData interface{}) error {
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

func (c *MarketplaceOrderClient) Create(ctx context.Context, req *MarketplaceOrderCreateRequest) (*MarketplaceOrderResponse, error) {
	var apiResp MarketplaceOrderResponse
	path := "/api/marketplace-orders/"

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *MarketplaceOrderClient) Get(ctx context.Context, id string) (*MarketplaceOrderResponse, error) {
	var apiResp MarketplaceOrderResponse
	err := c.Client.Get(ctx, "/api/marketplace-orders/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *MarketplaceOrderClient) Update(ctx context.Context, id string, req *MarketplaceOrderUpdateRequest) (*MarketplaceOrderResponse, error) {
	var apiResp MarketplaceOrderResponse
	err := c.Client.Update(ctx, "/api/marketplace-orders/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
func (c *MarketplaceOrderClient) Delete(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/marketplace-orders/{uuid}/", id)
}

func (c *MarketplaceOrderClient) List(ctx context.Context, filter map[string]string) ([]MarketplaceOrderResponse, error) {
	var listResult []MarketplaceOrderResponse
	err := c.Client.List(ctx, "/api/marketplace-orders/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
