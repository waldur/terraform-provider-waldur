package resource

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

func (c *Client) GetMarketplaceResource(ctx context.Context, id string) (*MarketplaceResourceResponse, error) {
	var apiResp MarketplaceResourceResponse
	err := c.Client.Get(ctx, "/api/marketplace-resources/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) UpdateMarketplaceResource(ctx context.Context, id string, req *MarketplaceResourceUpdateRequest) (*MarketplaceResourceResponse, error) {
	var apiResp MarketplaceResourceResponse
	err := c.Client.Update(ctx, "/api/marketplace-resources/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) ListMarketplaceResource(ctx context.Context, filter map[string]string) ([]MarketplaceResourceResponse, error) {
	var listResult []MarketplaceResourceResponse
	err := c.Client.List(ctx, "/api/marketplace-resources/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *Client) MarketplaceResourceUpdateLimits(ctx context.Context, id string, req *MarketplaceResourceUpdateLimitsActionRequest) error {
	path := "/api/marketplace-resources/{uuid}/update_limits/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}
func (c *Client) MarketplaceResourceUpdateOptions(ctx context.Context, id string, req *MarketplaceResourceUpdateOptionsActionRequest) error {
	path := "/api/marketplace-resources/{uuid}/update_options/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}
func (c *Client) MarketplaceResourcePull(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/marketplace-resources/{uuid}/pull/", id, nil, nil)
	return err
}
func (c *Client) MarketplaceResourceTerminate(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/marketplace-resources/{uuid}/terminate/", id, nil, nil)
	return err
}
func (c *Client) MarketplaceResourceUnlink(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/marketplace-resources/{uuid}/unlink/", id, nil, nil)
	return err
}
