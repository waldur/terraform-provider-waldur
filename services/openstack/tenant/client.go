package tenant

import (
	"context"
	"fmt"
	"strings"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type Client struct {
	Client *client.Client
}

func NewClient(c *client.Client) *Client {
	_ = fmt.Errorf
	_ = strings.Split
	return &Client{Client: c}
}

// CreateOpenstackTenantOrder creates a marketplace order for this resource.
func (c *Client) CreateOpenstackTenantOrder(ctx context.Context, req *OpenstackTenantCreateRequest) (*common.OrderDetails, error) {
	var apiResp common.OrderDetails
	err := c.Client.Post(ctx, "/api/marketplace-orders/", req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// TerminateOpenstackTenant terminates the resource via marketplace.
func (c *Client) TerminateOpenstackTenant(ctx context.Context, id string, req map[string]interface{}) (string, error) {
	var res map[string]interface{}
	err := c.Client.ExecuteAction(ctx, "/api/marketplace-resources/{uuid}/terminate/", id, req, &res)
	if err != nil {
		return "", err
	}
	if uuid, ok := res["uuid"].(string); ok {
		return uuid, nil
	}
	return "", nil
}

// GetOpenstackTenant retrieves a resource by its UUID.
func (c *Client) GetOpenstackTenant(ctx context.Context, id string) (*OpenstackTenantResponse, error) {
	var apiResp OpenstackTenantResponse
	err := c.Client.Get(ctx, "/api/openstack-tenants/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// UpdateOpenstackTenant updates an existing resource.
func (c *Client) UpdateOpenstackTenant(ctx context.Context, id string, req *OpenstackTenantUpdateRequest) (*OpenstackTenantResponse, error) {
	var apiResp OpenstackTenantResponse
	err := c.Client.Update(ctx, "/api/openstack-tenants/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// ListOpenstackTenant retrieves a list of resources with optional filtering.
func (c *Client) ListOpenstackTenant(ctx context.Context, filter map[string]string) ([]OpenstackTenantResponse, error) {
	var listResult []OpenstackTenantResponse
	err := c.Client.List(ctx, "/api/openstack-tenants/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

// OpenstackTenantPushSecurityGroups executes the push_security_groups action.
func (c *Client) OpenstackTenantPushSecurityGroups(ctx context.Context, id string, req *OpenstackTenantPushSecurityGroupsActionRequest) error {
	path := "/api/openstack-tenants/{uuid}/push_security_groups/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}

// OpenstackTenantPull executes the pull action.
func (c *Client) OpenstackTenantPull(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-tenants/{uuid}/pull/", id, nil, nil)
	return err
}

// OpenstackTenantUnlink executes the unlink action.
func (c *Client) OpenstackTenantUnlink(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-tenants/{uuid}/unlink/", id, nil, nil)
	return err
}
