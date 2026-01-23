package security_group

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

// OpenstackSecurityGroup Operations

func (c *Client) CreateOpenstackSecurityGroup(ctx context.Context, req *OpenstackSecurityGroupCreateRequest) (*OpenstackSecurityGroupResponse, error) {
	var apiResp OpenstackSecurityGroupResponse
	path := "/api/openstack-tenants/{uuid}/create_security_group/"
	path = strings.Replace(path, "{uuid}", *req.Tenant, 1)

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) GetOpenstackSecurityGroup(ctx context.Context, id string) (*OpenstackSecurityGroupResponse, error) {
	var apiResp OpenstackSecurityGroupResponse
	err := c.Client.GetByUUID(ctx, "/api/openstack-security-groups/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) UpdateOpenstackSecurityGroup(ctx context.Context, id string, req *OpenstackSecurityGroupUpdateRequest) (*OpenstackSecurityGroupResponse, error) {
	var apiResp OpenstackSecurityGroupResponse
	err := c.Client.Update(ctx, "/api/openstack-security-groups/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
func (c *Client) DeleteOpenstackSecurityGroup(ctx context.Context, id string) error {
	return c.Client.DeleteByUUID(ctx, "/api/openstack-security-groups/{uuid}/", id)
}

func (c *Client) ListOpenstackSecurityGroup(ctx context.Context, filter map[string]string) ([]OpenstackSecurityGroupResponse, error) {
	var listResult []OpenstackSecurityGroupResponse
	path := "/api/openstack-security-groups/"
	err := c.Client.ListWithFilter(ctx, path, filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *Client) OpenstackSecurityGroupSetRules(ctx context.Context, id string, req *OpenstackSecurityGroupSetRulesActionRequest) error {
	path := "/api/openstack-security-groups/{uuid}/set_rules/"
	_, err := c.ExecuteAction(ctx, path, id, req)
	return err
}
