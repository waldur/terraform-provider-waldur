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

// CreateOpenstackNetwork creates a new resource.
func (c *Client) CreateOpenstackNetwork(ctx context.Context, req *OpenstackNetworkCreateRequest) (*OpenstackNetworkResponse, error) {
	var apiResp OpenstackNetworkResponse
	path := "/api/openstack-tenants/{uuid}/create_network/"
	path = strings.Replace(path, "{uuid}", *req.Tenant, 1)

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// GetOpenstackNetwork retrieves a resource by its UUID.
func (c *Client) GetOpenstackNetwork(ctx context.Context, id string) (*OpenstackNetworkResponse, error) {
	var apiResp OpenstackNetworkResponse
	err := c.Client.Get(ctx, "/api/openstack-networks/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// UpdateOpenstackNetwork updates an existing resource.
func (c *Client) UpdateOpenstackNetwork(ctx context.Context, id string, req *OpenstackNetworkUpdateRequest) (*OpenstackNetworkResponse, error) {
	var apiResp OpenstackNetworkResponse
	err := c.Client.Update(ctx, "/api/openstack-networks/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// DeleteOpenstackNetwork deletes a resource.
func (c *Client) DeleteOpenstackNetwork(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/openstack-networks/{uuid}/", id)
}

// ListOpenstackNetwork retrieves a list of resources with optional filtering.
func (c *Client) ListOpenstackNetwork(ctx context.Context, filter map[string]string) ([]OpenstackNetworkResponse, error) {
	var listResult []OpenstackNetworkResponse
	err := c.Client.List(ctx, "/api/openstack-networks/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

// OpenstackNetworkSetMtu executes the set_mtu action.
func (c *Client) OpenstackNetworkSetMtu(ctx context.Context, id string, req *OpenstackNetworkSetMtuActionRequest) error {
	path := "/api/openstack-networks/{uuid}/set_mtu/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}

// OpenstackNetworkPull executes the pull action.
func (c *Client) OpenstackNetworkPull(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-networks/{uuid}/pull/", id, nil, nil)
	return err
}

// OpenstackNetworkUnlink executes the unlink action.
func (c *Client) OpenstackNetworkUnlink(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-networks/{uuid}/unlink/", id, nil, nil)
	return err
}
