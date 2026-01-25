package subnet

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

func (c *Client) CreateOpenstackSubnet(ctx context.Context, network string, req *OpenstackSubnetCreateRequest) (*OpenstackSubnetResponse, error) {
	var apiResp OpenstackSubnetResponse
	err := c.Client.ExecuteAction(ctx, "/api/openstack-networks/{uuid}/create_subnet/", network, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) GetOpenstackSubnet(ctx context.Context, id string) (*OpenstackSubnetResponse, error) {
	var apiResp OpenstackSubnetResponse
	err := c.Client.Get(ctx, "/api/openstack-subnets/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) UpdateOpenstackSubnet(ctx context.Context, id string, req *OpenstackSubnetUpdateRequest) (*OpenstackSubnetResponse, error) {
	var apiResp OpenstackSubnetResponse
	err := c.Client.Update(ctx, "/api/openstack-subnets/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
func (c *Client) DeleteOpenstackSubnet(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/openstack-subnets/{uuid}/", id)
}

func (c *Client) ListOpenstackSubnet(ctx context.Context, filter map[string]string) ([]OpenstackSubnetResponse, error) {
	var listResult []OpenstackSubnetResponse
	err := c.Client.List(ctx, "/api/openstack-subnets/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *Client) OpenstackSubnetConnect(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-subnets/{uuid}/connect/", id, nil, nil)
	return err
}
func (c *Client) OpenstackSubnetDisconnect(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-subnets/{uuid}/disconnect/", id, nil, nil)
	return err
}
func (c *Client) OpenstackSubnetPull(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-subnets/{uuid}/pull/", id, nil, nil)
	return err
}
func (c *Client) OpenstackSubnetUnlink(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-subnets/{uuid}/unlink/", id, nil, nil)
	return err
}
