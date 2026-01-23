package port

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

// CreateOpenstackPort creates a new resource.
func (c *Client) CreateOpenstackPort(ctx context.Context, req *OpenstackPortCreateRequest) (*OpenstackPortResponse, error) {
	var apiResp OpenstackPortResponse
	path := "/api/openstack-ports/"

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// GetOpenstackPort retrieves a resource by its UUID.
func (c *Client) GetOpenstackPort(ctx context.Context, id string) (*OpenstackPortResponse, error) {
	var apiResp OpenstackPortResponse
	err := c.Client.Get(ctx, "/api/openstack-ports/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// UpdateOpenstackPort updates an existing resource.
func (c *Client) UpdateOpenstackPort(ctx context.Context, id string, req *OpenstackPortUpdateRequest) (*OpenstackPortResponse, error) {
	var apiResp OpenstackPortResponse
	err := c.Client.Update(ctx, "/api/openstack-ports/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// DeleteOpenstackPort deletes a resource.
func (c *Client) DeleteOpenstackPort(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/openstack-ports/{uuid}/", id)
}

// ListOpenstackPort retrieves a list of resources with optional filtering.
func (c *Client) ListOpenstackPort(ctx context.Context, filter map[string]string) ([]OpenstackPortResponse, error) {
	var listResult []OpenstackPortResponse
	err := c.Client.List(ctx, "/api/openstack-ports/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

// OpenstackPortUpdateSecurityGroups executes the update_security_groups action.
func (c *Client) OpenstackPortUpdateSecurityGroups(ctx context.Context, id string, req *OpenstackPortUpdateSecurityGroupsActionRequest) error {
	path := "/api/openstack-ports/{uuid}/update_security_groups/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}

// OpenstackPortEnablePort executes the enable_port action.
func (c *Client) OpenstackPortEnablePort(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/enable_port/", id, nil, nil)
	return err
}

// OpenstackPortDisablePort executes the disable_port action.
func (c *Client) OpenstackPortDisablePort(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/disable_port/", id, nil, nil)
	return err
}

// OpenstackPortEnablePortSecurity executes the enable_port_security action.
func (c *Client) OpenstackPortEnablePortSecurity(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/enable_port_security/", id, nil, nil)
	return err
}

// OpenstackPortDisablePortSecurity executes the disable_port_security action.
func (c *Client) OpenstackPortDisablePortSecurity(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/disable_port_security/", id, nil, nil)
	return err
}

// OpenstackPortPull executes the pull action.
func (c *Client) OpenstackPortPull(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/pull/", id, nil, nil)
	return err
}

// OpenstackPortUnlink executes the unlink action.
func (c *Client) OpenstackPortUnlink(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/unlink/", id, nil, nil)
	return err
}
