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

func (c *Client) Create(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.Client.Create(ctx, path, body, result)
}

func (c *Client) GetByUUID(ctx context.Context, path string, uuid string, result interface{}) error {
	return c.Client.GetByUUID(ctx, path, uuid, result)
}

func (c *Client) Update(ctx context.Context, path string, uuid string, body interface{}, result interface{}) error {
	return c.Client.Update(ctx, path, uuid, body, result)
}

func (c *Client) DeleteByUUID(ctx context.Context, path string, uuid string) error {
	return c.Client.DeleteByUUID(ctx, path, uuid)
}

func (c *Client) Post(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.Client.Post(ctx, path, body, result)
}

func (c *Client) ListWithFilter(ctx context.Context, path string, filter map[string]string, result interface{}) error {
	return c.Client.ListWithFilter(ctx, path, filter, result)
}

func (c *Client) ExecuteAction(ctx context.Context, pathTemplate string, uuid string, body interface{}) (map[string]interface{}, error) {
	path := strings.Replace(pathTemplate, "{uuid}", uuid, 1)
	var res map[string]interface{}
	err := c.Client.Post(ctx, path, body, &res)
	return res, err
}

// OpenstackSubnet Operations

func (c *Client) CreateOpenstackSubnet(ctx context.Context, req *OpenstackSubnetCreateRequest) (*OpenstackSubnetResponse, error) {
	var apiResp OpenstackSubnetResponse
	path := "/api/openstack-networks/{uuid}/create_subnet/"
	path = strings.Replace(path, "{uuid}", *req.Network, 1)

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) GetOpenstackSubnet(ctx context.Context, id string) (*OpenstackSubnetResponse, error) {
	var apiResp OpenstackSubnetResponse
	err := c.Client.GetByUUID(ctx, "/api/openstack-subnets/{uuid}/", id, &apiResp)
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
	return c.Client.DeleteByUUID(ctx, "/api/openstack-subnets/{uuid}/", id)
}

func (c *Client) ListOpenstackSubnet(ctx context.Context, filter map[string]string) ([]OpenstackSubnetResponse, error) {
	var listResult []OpenstackSubnetResponse
	path := "/api/openstack-subnets/"
	err := c.Client.ListWithFilter(ctx, path, filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *Client) OpenstackSubnetConnect(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-subnets/{uuid}/connect/", id, nil)
	return err
}
func (c *Client) OpenstackSubnetDisconnect(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-subnets/{uuid}/disconnect/", id, nil)
	return err
}
func (c *Client) OpenstackSubnetPull(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-subnets/{uuid}/pull/", id, nil)
	return err
}
func (c *Client) OpenstackSubnetUnlink(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-subnets/{uuid}/unlink/", id, nil)
	return err
}
