package tenant

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackTenantClient struct {
	Client *client.Client
}

func NewOpenstackTenantClient(c *client.Client) *OpenstackTenantClient {
	return &OpenstackTenantClient{Client: c}
}

func (c *OpenstackTenantClient) Configure(ctx context.Context, providerData interface{}) error {
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

// CreateOrder creates a marketplace order for this resource.
func (c *OpenstackTenantClient) CreateOrder(ctx context.Context, req *OpenstackTenantCreateRequest) (*common.OrderDetails, error) {
	var apiResp common.OrderDetails
	err := c.Client.Post(ctx, "/api/marketplace-orders/", req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// Terminate terminates the resource via marketplace.
func (c *OpenstackTenantClient) Terminate(ctx context.Context, id string, req map[string]interface{}) (string, error) {
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

func (c *OpenstackTenantClient) Get(ctx context.Context, id string) (*OpenstackTenantResponse, error) {
	var apiResp OpenstackTenantResponse
	err := c.Client.Get(ctx, "/api/openstack-tenants/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *OpenstackTenantClient) Update(ctx context.Context, id string, req *OpenstackTenantUpdateRequest) (*OpenstackTenantResponse, error) {
	var apiResp OpenstackTenantResponse
	err := c.Client.Update(ctx, "/api/openstack-tenants/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *OpenstackTenantClient) List(ctx context.Context, filter map[string]string) ([]OpenstackTenantResponse, error) {
	var listResult []OpenstackTenantResponse
	err := c.Client.List(ctx, "/api/openstack-tenants/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *OpenstackTenantClient) PushSecurityGroups(ctx context.Context, id string, req *OpenstackTenantPushSecurityGroupsActionRequest) error {
	path := "/api/openstack-tenants/{uuid}/push_security_groups/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}
func (c *OpenstackTenantClient) Pull(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-tenants/{uuid}/pull/", id, nil, nil)
	return err
}
func (c *OpenstackTenantClient) Unlink(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-tenants/{uuid}/unlink/", id, nil, nil)
	return err
}
