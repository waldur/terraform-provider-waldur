package security_group

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackSecurityGroupClient struct {
	Client *client.Client
}

func NewOpenstackSecurityGroupClient(c *client.Client) *OpenstackSecurityGroupClient {
	return &OpenstackSecurityGroupClient{Client: c}
}

func (c *OpenstackSecurityGroupClient) Configure(ctx context.Context, providerData interface{}) error {
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

func (c *OpenstackSecurityGroupClient) Create(ctx context.Context, tenant string, req *OpenstackSecurityGroupCreateRequest) (*OpenstackSecurityGroupResponse, error) {
	var apiResp OpenstackSecurityGroupResponse
	tenantUUID := common.ExtractUUIDFromURL(tenant)
	err := c.Client.ExecuteAction(ctx, "/api/openstack-tenants/{uuid}/create_security_group/", tenantUUID, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *OpenstackSecurityGroupClient) Get(ctx context.Context, id string) (*OpenstackSecurityGroupResponse, error) {
	var apiResp OpenstackSecurityGroupResponse
	err := c.Client.Get(ctx, "/api/openstack-security-groups/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *OpenstackSecurityGroupClient) Update(ctx context.Context, id string, req *OpenstackSecurityGroupUpdateRequest) (*OpenstackSecurityGroupResponse, error) {
	var apiResp OpenstackSecurityGroupResponse
	err := c.Client.Update(ctx, "/api/openstack-security-groups/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
func (c *OpenstackSecurityGroupClient) Delete(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/openstack-security-groups/{uuid}/", id)
}

func (c *OpenstackSecurityGroupClient) List(ctx context.Context, filter map[string]string) ([]OpenstackSecurityGroupResponse, error) {
	var listResult []OpenstackSecurityGroupResponse
	err := c.Client.List(ctx, "/api/openstack-security-groups/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *OpenstackSecurityGroupClient) SetRules(ctx context.Context, id string, req *OpenstackSecurityGroupSetRulesActionRequest) error {
	path := "/api/openstack-security-groups/{uuid}/set_rules/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}
