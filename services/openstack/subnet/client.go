package subnet

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

// CreateOpenstackSubnet creates a new resource.
func (c *Client) CreateOpenstackSubnet(ctx context.Context, req *OpenstackSubnetCreateRequest) (*OpenstackSubnetResponse, error) {
	var apiResp OpenstackSubnetResponse
	err := c.Client.ExecuteAction(ctx, "/api/openstack-networks/{uuid}/create_subnet/", *req.Network, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// GetOpenstackSubnet retrieves a resource by its UUID.
func (c *Client) GetOpenstackSubnet(ctx context.Context, id string) (*OpenstackSubnetResponse, error) {
	var apiResp OpenstackSubnetResponse
	err := c.Client.Get(ctx, "/api/openstack-subnets/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// UpdateOpenstackSubnet updates an existing resource.
func (c *Client) UpdateOpenstackSubnet(ctx context.Context, id string, req *OpenstackSubnetUpdateRequest) (*OpenstackSubnetResponse, error) {
	var apiResp OpenstackSubnetResponse
	err := c.Client.Update(ctx, "/api/openstack-subnets/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// DeleteOpenstackSubnet deletes a resource.
func (c *Client) DeleteOpenstackSubnet(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/openstack-subnets/{uuid}/", id)
}

// ListOpenstackSubnet retrieves a list of resources with optional filtering.
func (c *Client) ListOpenstackSubnet(ctx context.Context, filter map[string]string) ([]OpenstackSubnetResponse, error) {
	var listResult []OpenstackSubnetResponse
	err := c.Client.List(ctx, "/api/openstack-subnets/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

// OpenstackSubnetConnect executes the connect action.
func (c *Client) OpenstackSubnetConnect(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-subnets/{uuid}/connect/", id, nil, nil)
	return err
}

// OpenstackSubnetDisconnect executes the disconnect action.
func (c *Client) OpenstackSubnetDisconnect(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-subnets/{uuid}/disconnect/", id, nil, nil)
	return err
}

// OpenstackSubnetPull executes the pull action.
func (c *Client) OpenstackSubnetPull(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-subnets/{uuid}/pull/", id, nil, nil)
	return err
}

// OpenstackSubnetUnlink executes the unlink action.
func (c *Client) OpenstackSubnetUnlink(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-subnets/{uuid}/unlink/", id, nil, nil)
	return err
}
