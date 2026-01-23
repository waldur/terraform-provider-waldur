package image

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

// GetOpenstackImage retrieves a resource by its UUID.
func (c *Client) GetOpenstackImage(ctx context.Context, id string) (*OpenstackImageResponse, error) {
	var apiResp OpenstackImageResponse
	err := c.Client.Get(ctx, "/api/openstack-images/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// ListOpenstackImage retrieves a list of resources with optional filtering.
func (c *Client) ListOpenstackImage(ctx context.Context, filter map[string]string) ([]OpenstackImageResponse, error) {
	var listResult []OpenstackImageResponse
	err := c.Client.List(ctx, "/api/openstack-images/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
