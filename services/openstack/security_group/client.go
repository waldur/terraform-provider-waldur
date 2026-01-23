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
	err := c.Client.ListWithFilter(ctx, "/api/openstack-security-groups/", filter, &listResult)
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
