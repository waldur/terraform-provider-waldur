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

func (c *Client) GetOpenstackVolumeType(ctx context.Context, id string) (*OpenstackVolumeTypeResponse, error) {
	var apiResp OpenstackVolumeTypeResponse
	err := c.Client.GetByUUID(ctx, "/api/openstack-volume-types/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) ListOpenstackVolumeType(ctx context.Context, filter map[string]string) ([]OpenstackVolumeTypeResponse, error) {
	var listResult []OpenstackVolumeTypeResponse
	err := c.Client.ListWithFilter(ctx, "/api/openstack-volume-types/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
