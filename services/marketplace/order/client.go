package order

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

func (c *Client) CreateMarketplaceOrder(ctx context.Context, req *MarketplaceOrderCreateRequest) (*MarketplaceOrderResponse, error) {
	var apiResp MarketplaceOrderResponse
	path := "/api/marketplace-orders/"

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) GetMarketplaceOrder(ctx context.Context, id string) (*MarketplaceOrderResponse, error) {
	var apiResp MarketplaceOrderResponse
	err := c.Client.GetByUUID(ctx, "/api/marketplace-orders/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) DeleteMarketplaceOrder(ctx context.Context, id string) error {
	return c.Client.DeleteByUUID(ctx, "/api/marketplace-orders/{uuid}/", id)
}

func (c *Client) ListMarketplaceOrder(ctx context.Context, filter map[string]string) ([]MarketplaceOrderResponse, error) {
	var listResult []MarketplaceOrderResponse
	err := c.Client.ListWithFilter(ctx, "/api/marketplace-orders/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
