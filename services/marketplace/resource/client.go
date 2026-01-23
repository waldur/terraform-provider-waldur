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

func (c *Client) Create(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.Client.Create(ctx, path, body, result)
}

func (c *Client) GetByUUID(ctx context.Context, path string, uuid string, result interface{}) error {
	return c.Client.GetByUUID(ctx, path, uuid, result)
}

func (c *Client) Update(ctx context.Context, path string, uuid string, body interface{}, result interface{}) error {
	return c.Client.Update(ctx, path, uuid, body, result)
}

func (c *Client) DeleteByUUID(ctx context.Context, path string, uuid string) error {
	return c.Client.DeleteByUUID(ctx, path, uuid)
}

func (c *Client) Post(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.Client.Post(ctx, path, body, result)
}

func (c *Client) ListWithFilter(ctx context.Context, path string, filter map[string]string, result interface{}) error {
	return c.Client.ListWithFilter(ctx, path, filter, result)
}

func (c *Client) ExecuteAction(ctx context.Context, pathTemplate string, uuid string, body interface{}) (map[string]interface{}, error) {
	path := strings.Replace(pathTemplate, "{uuid}", uuid, 1)
	var res map[string]interface{}
	err := c.Client.Post(ctx, path, body, &res)
	return res, err
}

// MarketplaceResource Operations

func (c *Client) GetMarketplaceResource(ctx context.Context, id string) (*MarketplaceResourceResponse, error) {
	var apiResp MarketplaceResourceResponse
	err := c.Client.GetByUUID(ctx, "/api/marketplace-resources/{uuid}/", id, &apiResp)
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
	path := "/api/marketplace-resources/"
	err := c.Client.ListWithFilter(ctx, path, filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *Client) MarketplaceResourceUpdateLimits(ctx context.Context, id string, req *MarketplaceResourceUpdateLimitsActionRequest) error {
	path := "/api/marketplace-resources/{uuid}/update_limits/"
	_, err := c.ExecuteAction(ctx, path, id, req)
	return err
}
func (c *Client) MarketplaceResourceUpdateOptions(ctx context.Context, id string, req *MarketplaceResourceUpdateOptionsActionRequest) error {
	path := "/api/marketplace-resources/{uuid}/update_options/"
	_, err := c.ExecuteAction(ctx, path, id, req)
	return err
}
func (c *Client) MarketplaceResourcePull(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/marketplace-resources/{uuid}/pull/", id, nil)
	return err
}
func (c *Client) MarketplaceResourceTerminate(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/marketplace-resources/{uuid}/terminate/", id, nil)
	return err
}
func (c *Client) MarketplaceResourceUnlink(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/marketplace-resources/{uuid}/unlink/", id, nil)
	return err
}
