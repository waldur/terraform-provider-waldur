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

// CreateOpenstackSecurityGroup creates a new resource.
func (c *Client) CreateOpenstackSecurityGroup(ctx context.Context, req *OpenstackSecurityGroupCreateRequest) (*OpenstackSecurityGroupResponse, error) {
	var apiResp OpenstackSecurityGroupResponse
	err := c.Client.ExecuteAction(ctx, "/api/openstack-tenants/{uuid}/create_security_group/", *req.Tenant, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// GetOpenstackSecurityGroup retrieves a resource by its UUID.
func (c *Client) GetOpenstackSecurityGroup(ctx context.Context, id string) (*OpenstackSecurityGroupResponse, error) {
	var apiResp OpenstackSecurityGroupResponse
	err := c.Client.Get(ctx, "/api/openstack-security-groups/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// UpdateOpenstackSecurityGroup updates an existing resource.
func (c *Client) UpdateOpenstackSecurityGroup(ctx context.Context, id string, req *OpenstackSecurityGroupUpdateRequest) (*OpenstackSecurityGroupResponse, error) {
	var apiResp OpenstackSecurityGroupResponse
	err := c.Client.Update(ctx, "/api/openstack-security-groups/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// DeleteOpenstackSecurityGroup deletes a resource.
func (c *Client) DeleteOpenstackSecurityGroup(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/openstack-security-groups/{uuid}/", id)
}

// ListOpenstackSecurityGroup retrieves a list of resources with optional filtering.
func (c *Client) ListOpenstackSecurityGroup(ctx context.Context, filter map[string]string) ([]OpenstackSecurityGroupResponse, error) {
	var listResult []OpenstackSecurityGroupResponse
	err := c.Client.List(ctx, "/api/openstack-security-groups/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

// OpenstackSecurityGroupSetRules executes the set_rules action.
func (c *Client) OpenstackSecurityGroupSetRules(ctx context.Context, id string, req *OpenstackSecurityGroupSetRulesActionRequest) error {
	path := "/api/openstack-security-groups/{uuid}/set_rules/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}
