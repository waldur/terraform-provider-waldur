package security_group

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

func (c *Client) CreateOpenstackSecurityGroup(ctx context.Context, tenant string, req *OpenstackSecurityGroupCreateRequest) (*OpenstackSecurityGroupResponse, error) {
	var apiResp OpenstackSecurityGroupResponse
	err := c.Client.ExecuteAction(ctx, "/api/openstack-tenants/{uuid}/create_security_group/", tenant, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) GetOpenstackSecurityGroup(ctx context.Context, id string) (*OpenstackSecurityGroupResponse, error) {
	var apiResp OpenstackSecurityGroupResponse
	err := c.Client.Get(ctx, "/api/openstack-security-groups/{uuid}/", id, &apiResp)
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
	return c.Client.Delete(ctx, "/api/openstack-security-groups/{uuid}/", id)
}

func (c *Client) ListOpenstackSecurityGroup(ctx context.Context, filter map[string]string) ([]OpenstackSecurityGroupResponse, error) {
	var listResult []OpenstackSecurityGroupResponse
	err := c.Client.List(ctx, "/api/openstack-security-groups/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *Client) OpenstackSecurityGroupSetRules(ctx context.Context, id string, req *OpenstackSecurityGroupSetRulesActionRequest) error {
	path := "/api/openstack-security-groups/{uuid}/set_rules/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}
