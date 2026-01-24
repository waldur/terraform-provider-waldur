package network

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

func (c *Client) CreateOpenstackNetwork(ctx context.Context, tenant string, req *OpenstackNetworkCreateRequest) (*OpenstackNetworkResponse, error) {
	var apiResp OpenstackNetworkResponse
	err := c.Client.ExecuteAction(ctx, "/api/openstack-tenants/{uuid}/create_network/", tenant, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) GetOpenstackNetwork(ctx context.Context, id string) (*OpenstackNetworkResponse, error) {
	var apiResp OpenstackNetworkResponse
	err := c.Client.Get(ctx, "/api/openstack-networks/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) UpdateOpenstackNetwork(ctx context.Context, id string, req *OpenstackNetworkUpdateRequest) (*OpenstackNetworkResponse, error) {
	var apiResp OpenstackNetworkResponse
	err := c.Client.Update(ctx, "/api/openstack-networks/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
func (c *Client) DeleteOpenstackNetwork(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/openstack-networks/{uuid}/", id)
}

func (c *Client) ListOpenstackNetwork(ctx context.Context, filter map[string]string) ([]OpenstackNetworkResponse, error) {
	var listResult []OpenstackNetworkResponse
	err := c.Client.List(ctx, "/api/openstack-networks/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *Client) OpenstackNetworkSetMtu(ctx context.Context, id string, req *OpenstackNetworkSetMtuActionRequest) error {
	path := "/api/openstack-networks/{uuid}/set_mtu/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}
func (c *Client) OpenstackNetworkPull(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-networks/{uuid}/pull/", id, nil, nil)
	return err
}
func (c *Client) OpenstackNetworkUnlink(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-networks/{uuid}/unlink/", id, nil, nil)
	return err
}
