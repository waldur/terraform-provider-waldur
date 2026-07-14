package resource

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type MarketplaceResourceClient struct {
	Client *client.Client
}

func NewMarketplaceResourceClient(c *client.Client) *MarketplaceResourceClient {
	return &MarketplaceResourceClient{Client: c}
}

func (c *MarketplaceResourceClient) Configure(ctx context.Context, providerData interface{}) error {
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

func (c *MarketplaceResourceClient) Get(ctx context.Context, id string) (*MarketplaceResourceResponse, error) {
	var apiResp MarketplaceResourceResponse
	err := c.Client.Get(ctx, "/api/marketplace-resources/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *MarketplaceResourceClient) Update(ctx context.Context, id string, req *MarketplaceResourceUpdateRequest) (*MarketplaceResourceResponse, error) {
	var apiResp MarketplaceResourceResponse
	err := c.Client.Update(ctx, "/api/marketplace-resources/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *MarketplaceResourceClient) List(ctx context.Context, filter map[string]string) ([]MarketplaceResourceResponse, error) {
	var listResult []MarketplaceResourceResponse
	err := c.Client.List(ctx, "/api/marketplace-resources/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *MarketplaceResourceClient) UpdateLimits(ctx context.Context, id string, req *MarketplaceResourceUpdateLimitsActionRequest) error {
	path := "/api/marketplace-resources/{uuid}/update_limits/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}
func (c *MarketplaceResourceClient) UpdateOptions(ctx context.Context, id string, req *MarketplaceResourceUpdateOptionsActionRequest) error {
	path := "/api/marketplace-resources/{uuid}/update_options/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}
func (c *MarketplaceResourceClient) Pull(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/marketplace-resources/{uuid}/pull/", id, nil, nil)
	return err
}
func (c *MarketplaceResourceClient) Terminate(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/marketplace-resources/{uuid}/terminate/", id, nil, nil)
	return err
}
func (c *MarketplaceResourceClient) Unlink(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/marketplace-resources/{uuid}/unlink/", id, nil, nil)
	return err
}
