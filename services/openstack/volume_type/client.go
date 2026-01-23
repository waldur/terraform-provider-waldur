package volume_type

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

// GetOpenstackVolumeType retrieves a resource by its UUID.
func (c *Client) GetOpenstackVolumeType(ctx context.Context, id string) (*OpenstackVolumeTypeResponse, error) {
	var apiResp OpenstackVolumeTypeResponse
	err := c.Client.Get(ctx, "/api/openstack-volume-types/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// ListOpenstackVolumeType retrieves a list of resources with optional filtering.
func (c *Client) ListOpenstackVolumeType(ctx context.Context, filter map[string]string) ([]OpenstackVolumeTypeResponse, error) {
	var listResult []OpenstackVolumeTypeResponse
	err := c.Client.List(ctx, "/api/openstack-volume-types/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
