package server_group

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

func (c *Client) CreateOpenstackServerGroup(ctx context.Context, tenant string, req *OpenstackServerGroupCreateRequest) (*OpenstackServerGroupResponse, error) {
	var apiResp OpenstackServerGroupResponse
	err := c.Client.ExecuteAction(ctx, "/api/openstack-tenants/{uuid}/create_server_group/", tenant, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) GetOpenstackServerGroup(ctx context.Context, id string) (*OpenstackServerGroupResponse, error) {
	var apiResp OpenstackServerGroupResponse
	err := c.Client.Get(ctx, "/api/openstack-server-groups/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) UpdateOpenstackServerGroup(ctx context.Context, id string, req *OpenstackServerGroupUpdateRequest) (*OpenstackServerGroupResponse, error) {
	var apiResp OpenstackServerGroupResponse
	err := c.Client.Update(ctx, "/api/openstack-server-groups/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
func (c *Client) DeleteOpenstackServerGroup(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/openstack-server-groups/{uuid}/", id)
}

func (c *Client) ListOpenstackServerGroup(ctx context.Context, filter map[string]string) ([]OpenstackServerGroupResponse, error) {
	var listResult []OpenstackServerGroupResponse
	err := c.Client.List(ctx, "/api/openstack-server-groups/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
