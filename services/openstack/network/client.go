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

// OpenstackNetwork Operations

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

func (c *Client) GetOpenstackNetwork(ctx context.Context, id string) (*OpenstackNetworkResponse, error) {
	var apiResp OpenstackNetworkResponse
	err := c.Client.GetByUUID(ctx, "/api/openstack-networks/{uuid}/", id, &apiResp)
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
	return c.Client.DeleteByUUID(ctx, "/api/openstack-networks/{uuid}/", id)
}

func (c *Client) ListOpenstackNetwork(ctx context.Context, filter map[string]string) ([]OpenstackNetworkResponse, error) {
	var listResult []OpenstackNetworkResponse
	path := "/api/openstack-networks/"
	err := c.Client.ListWithFilter(ctx, path, filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *Client) OpenstackNetworkSetMtu(ctx context.Context, id string, req *OpenstackNetworkSetMtuActionRequest) error {
	path := "/api/openstack-networks/{uuid}/set_mtu/"
	_, err := c.ExecuteAction(ctx, path, id, req)
	return err
}
func (c *Client) OpenstackNetworkPull(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-networks/{uuid}/pull/", id, nil)
	return err
}
func (c *Client) OpenstackNetworkUnlink(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-networks/{uuid}/unlink/", id, nil)
	return err
}
