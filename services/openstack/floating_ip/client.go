package floating_ip

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type Client struct {
	Client *client.Client
}

func NewClient(c *client.Client) *Client {
	return &Client{Client: c}
}

func (c *Client) Configure(ctx context.Context, providerData interface{}) error {
	if providerData == nil {
		return nil
	}

	raw, ok := providerData.(*client.Client)
	if !ok {
		return fmt.Errorf("unexpected provider data type: %T", providerData)
	}

	c.Client = raw
	return nil
}

func IsNotFoundError(err error) bool {
	return common.IsNotFoundError(err)
}

func (c *Client) CreateOpenstackFloatingIp(ctx context.Context, tenant string, req *OpenstackFloatingIpCreateRequest) (*OpenstackFloatingIpResponse, error) {
	var apiResp OpenstackFloatingIpResponse
	err := c.Client.ExecuteAction(ctx, "/api/openstack-tenants/{uuid}/create_floating_ip/", tenant, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) GetOpenstackFloatingIp(ctx context.Context, id string) (*OpenstackFloatingIpResponse, error) {
	var apiResp OpenstackFloatingIpResponse
	err := c.Client.Get(ctx, "/api/openstack-floating-ips/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) DeleteOpenstackFloatingIp(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/openstack-floating-ips/{uuid}/", id)
}

func (c *Client) ListOpenstackFloatingIp(ctx context.Context, filter map[string]string) ([]OpenstackFloatingIpResponse, error) {
	var listResult []OpenstackFloatingIpResponse
	err := c.Client.List(ctx, "/api/openstack-floating-ips/", filter, &listResult)
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
