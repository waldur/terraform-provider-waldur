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

func (c *Client) CreateOpenstackFloatingIp(ctx context.Context, req *OpenstackFloatingIpCreateRequest) (*OpenstackFloatingIpResponse, error) {
	var apiResp OpenstackFloatingIpResponse
	path := "/api/openstack-tenants/{uuid}/create_floating_ip/"
	path = strings.Replace(path, "{uuid}", *req.Tenant, 1)

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) GetOpenstackFloatingIp(ctx context.Context, id string) (*OpenstackFloatingIpResponse, error) {
	var apiResp OpenstackFloatingIpResponse
	err := c.Client.GetByUUID(ctx, "/api/openstack-floating-ips/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) DeleteOpenstackFloatingIp(ctx context.Context, id string) error {
	return c.Client.DeleteByUUID(ctx, "/api/openstack-floating-ips/{uuid}/", id)
}

func (c *Client) ListOpenstackFloatingIp(ctx context.Context, filter map[string]string) ([]OpenstackFloatingIpResponse, error) {
	var listResult []OpenstackFloatingIpResponse
	err := c.Client.ListWithFilter(ctx, "/api/openstack-floating-ips/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *Client) OpenstackFloatingIpUpdateDescription(ctx context.Context, id string, req *OpenstackFloatingIpUpdateDescriptionActionRequest) error {
	path := "/api/openstack-floating-ips/{uuid}/update_description/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}
