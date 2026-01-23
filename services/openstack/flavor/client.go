package flavor

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

// GetOpenstackFlavor retrieves a resource by its UUID.
func (c *Client) GetOpenstackFlavor(ctx context.Context, id string) (*OpenstackFlavorResponse, error) {
	var apiResp OpenstackFlavorResponse
	err := c.Client.Get(ctx, "/api/openstack-flavors/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// ListOpenstackFlavor retrieves a list of resources with optional filtering.
func (c *Client) ListOpenstackFlavor(ctx context.Context, filter map[string]string) ([]OpenstackFlavorResponse, error) {
	var listResult []OpenstackFlavorResponse
	err := c.Client.List(ctx, "/api/openstack-flavors/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
