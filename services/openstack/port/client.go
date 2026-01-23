package port

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

// OpenstackPort Operations

func (c *Client) CreateOpenstackPort(ctx context.Context, req *OpenstackPortCreateRequest) (*OpenstackPortResponse, error) {
	var apiResp OpenstackPortResponse
	path := "/api/openstack-ports/"

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) GetOpenstackPort(ctx context.Context, id string) (*OpenstackPortResponse, error) {
	var apiResp OpenstackPortResponse
	err := c.Client.GetByUUID(ctx, "/api/openstack-ports/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) UpdateOpenstackPort(ctx context.Context, id string, req *OpenstackPortUpdateRequest) (*OpenstackPortResponse, error) {
	var apiResp OpenstackPortResponse
	err := c.Client.Update(ctx, "/api/openstack-ports/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
func (c *Client) DeleteOpenstackPort(ctx context.Context, id string) error {
	return c.Client.DeleteByUUID(ctx, "/api/openstack-ports/{uuid}/", id)
}

func (c *Client) ListOpenstackPort(ctx context.Context, filter map[string]string) ([]OpenstackPortResponse, error) {
	var listResult []OpenstackPortResponse
	path := "/api/openstack-ports/"
	err := c.Client.ListWithFilter(ctx, path, filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *Client) OpenstackPortUpdateSecurityGroups(ctx context.Context, id string, req *OpenstackPortUpdateSecurityGroupsActionRequest) error {
	path := "/api/openstack-ports/{uuid}/update_security_groups/"
	_, err := c.ExecuteAction(ctx, path, id, req)
	return err
}
func (c *Client) OpenstackPortEnablePort(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/enable_port/", id, nil)
	return err
}
func (c *Client) OpenstackPortDisablePort(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/disable_port/", id, nil)
	return err
}
func (c *Client) OpenstackPortEnablePortSecurity(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/enable_port_security/", id, nil)
	return err
}
func (c *Client) OpenstackPortDisablePortSecurity(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/disable_port_security/", id, nil)
	return err
}
func (c *Client) OpenstackPortPull(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/pull/", id, nil)
	return err
}
func (c *Client) OpenstackPortUnlink(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/unlink/", id, nil)
	return err
}
