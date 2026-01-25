package flavor

import (
	"context"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

type Client struct {
	Client *client.Client
}

func NewClient(c *client.Client) *Client {
	return &Client{Client: c}
}

func (c *Client) GetOpenstackFlavor(ctx context.Context, id string) (*OpenstackFlavorResponse, error) {
	var apiResp OpenstackFlavorResponse
	err := c.Client.Get(ctx, "/api/openstack-flavors/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) ListOpenstackFlavor(ctx context.Context, filter map[string]string) ([]OpenstackFlavorResponse, error) {
	var listResult []OpenstackFlavorResponse
	err := c.Client.List(ctx, "/api/openstack-flavors/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
