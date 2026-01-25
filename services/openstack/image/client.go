package image

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

func (c *Client) GetOpenstackImage(ctx context.Context, id string) (*OpenstackImageResponse, error) {
	var apiResp OpenstackImageResponse
	err := c.Client.Get(ctx, "/api/openstack-images/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) ListOpenstackImage(ctx context.Context, filter map[string]string) ([]OpenstackImageResponse, error) {
	var listResult []OpenstackImageResponse
	err := c.Client.List(ctx, "/api/openstack-images/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
