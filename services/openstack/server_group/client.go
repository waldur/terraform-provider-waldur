package server_group

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

// CreateOpenstackServerGroup creates a new resource.
func (c *Client) CreateOpenstackServerGroup(ctx context.Context, req *OpenstackServerGroupCreateRequest) (*OpenstackServerGroupResponse, error) {
	var apiResp OpenstackServerGroupResponse
	err := c.Client.ExecuteAction(ctx, "/api/openstack-tenants/{uuid}/create_server_group/", *req.Tenant, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// GetOpenstackServerGroup retrieves a resource by its UUID.
func (c *Client) GetOpenstackServerGroup(ctx context.Context, id string) (*OpenstackServerGroupResponse, error) {
	var apiResp OpenstackServerGroupResponse
	err := c.Client.Get(ctx, "/api/openstack-server-groups/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// UpdateOpenstackServerGroup updates an existing resource.
func (c *Client) UpdateOpenstackServerGroup(ctx context.Context, id string, req *OpenstackServerGroupUpdateRequest) (*OpenstackServerGroupResponse, error) {
	var apiResp OpenstackServerGroupResponse
	err := c.Client.Update(ctx, "/api/openstack-server-groups/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// DeleteOpenstackServerGroup deletes a resource.
func (c *Client) DeleteOpenstackServerGroup(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/openstack-server-groups/{uuid}/", id)
}

// ListOpenstackServerGroup retrieves a list of resources with optional filtering.
func (c *Client) ListOpenstackServerGroup(ctx context.Context, filter map[string]string) ([]OpenstackServerGroupResponse, error) {
	var listResult []OpenstackServerGroupResponse
	err := c.Client.List(ctx, "/api/openstack-server-groups/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
