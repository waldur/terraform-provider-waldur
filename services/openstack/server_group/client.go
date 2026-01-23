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

func (c *Client) CreateOpenstackServerGroup(ctx context.Context, req *OpenstackServerGroupCreateRequest) (*OpenstackServerGroupResponse, error) {
	var apiResp OpenstackServerGroupResponse
	path := "/api/openstack-tenants/{uuid}/create_server_group/"
	path = strings.Replace(path, "{uuid}", *req.Tenant, 1)

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) GetOpenstackServerGroup(ctx context.Context, id string) (*OpenstackServerGroupResponse, error) {
	var apiResp OpenstackServerGroupResponse
	err := c.Client.GetByUUID(ctx, "/api/openstack-server-groups/{uuid}/", id, &apiResp)
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
	return c.Client.DeleteByUUID(ctx, "/api/openstack-server-groups/{uuid}/", id)
}

func (c *Client) ListOpenstackServerGroup(ctx context.Context, filter map[string]string) ([]OpenstackServerGroupResponse, error) {
	var listResult []OpenstackServerGroupResponse
	err := c.Client.ListWithFilter(ctx, "/api/openstack-server-groups/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
