package offering

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

func (c *Client) CreateMarketplaceOffering(ctx context.Context, req *MarketplaceOfferingCreateRequest) (*MarketplaceOfferingResponse, error) {
	var apiResp MarketplaceOfferingResponse
	path := "/api/marketplace-provider-offerings/"

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) GetMarketplaceOffering(ctx context.Context, id string) (*MarketplaceOfferingResponse, error) {
	var apiResp MarketplaceOfferingResponse
	err := c.Client.GetByUUID(ctx, "/api/marketplace-provider-offerings/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) DeleteMarketplaceOffering(ctx context.Context, id string) error {
	return c.Client.DeleteByUUID(ctx, "/api/marketplace-provider-offerings/{uuid}/", id)
}

func (c *Client) ListMarketplaceOffering(ctx context.Context, filter map[string]string) ([]MarketplaceOfferingResponse, error) {
	var listResult []MarketplaceOfferingResponse
	err := c.Client.ListWithFilter(ctx, "/api/marketplace-provider-offerings/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
