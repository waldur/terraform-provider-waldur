package floating_ip

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

// CreateOpenstackFloatingIp creates a new resource.
func (c *Client) CreateOpenstackFloatingIp(ctx context.Context, req *OpenstackFloatingIpCreateRequest) (*OpenstackFloatingIpResponse, error) {
	var apiResp OpenstackFloatingIpResponse
	err := c.Client.ExecuteAction(ctx, "/api/openstack-tenants/{uuid}/create_floating_ip/", *req.Tenant, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// GetOpenstackFloatingIp retrieves a resource by its UUID.
func (c *Client) GetOpenstackFloatingIp(ctx context.Context, id string) (*OpenstackFloatingIpResponse, error) {
	var apiResp OpenstackFloatingIpResponse
	err := c.Client.Get(ctx, "/api/openstack-floating-ips/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// DeleteOpenstackFloatingIp deletes a resource.
func (c *Client) DeleteOpenstackFloatingIp(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/openstack-floating-ips/{uuid}/", id)
}

// ListOpenstackFloatingIp retrieves a list of resources with optional filtering.
func (c *Client) ListOpenstackFloatingIp(ctx context.Context, filter map[string]string) ([]OpenstackFloatingIpResponse, error) {
	var listResult []OpenstackFloatingIpResponse
	err := c.Client.List(ctx, "/api/openstack-floating-ips/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

// OpenstackFloatingIpUpdateDescription executes the update_description action.
func (c *Client) OpenstackFloatingIpUpdateDescription(ctx context.Context, id string, req *OpenstackFloatingIpUpdateDescriptionActionRequest) error {
	path := "/api/openstack-floating-ips/{uuid}/update_description/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}
