package port

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackPortClient struct {
	Client *client.Client
}

func NewOpenstackPortClient(c *client.Client) *OpenstackPortClient {
	return &OpenstackPortClient{Client: c}
}

func (c *OpenstackPortClient) Configure(ctx context.Context, providerData interface{}) error {
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

func (c *OpenstackPortClient) Create(ctx context.Context, req *OpenstackPortCreateRequest) (*OpenstackPortResponse, error) {
	var apiResp OpenstackPortResponse
	path := "/api/openstack-ports/"

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *OpenstackPortClient) Get(ctx context.Context, id string) (*OpenstackPortResponse, error) {
	var apiResp OpenstackPortResponse
	err := c.Client.Get(ctx, "/api/openstack-ports/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *OpenstackPortClient) Update(ctx context.Context, id string, req *OpenstackPortUpdateRequest) (*OpenstackPortResponse, error) {
	var apiResp OpenstackPortResponse
	err := c.Client.Update(ctx, "/api/openstack-ports/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
func (c *OpenstackPortClient) Delete(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/openstack-ports/{uuid}/", id)
}

func (c *OpenstackPortClient) List(ctx context.Context, filter map[string]string) ([]OpenstackPortResponse, error) {
	var listResult []OpenstackPortResponse
	err := c.Client.List(ctx, "/api/openstack-ports/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *OpenstackPortClient) UpdateSecurityGroups(ctx context.Context, id string, req *OpenstackPortUpdateSecurityGroupsActionRequest) error {
	path := "/api/openstack-ports/{uuid}/update_security_groups/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}
func (c *OpenstackPortClient) EnablePort(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/enable_port/", id, nil, nil)
	return err
}
func (c *OpenstackPortClient) DisablePort(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/disable_port/", id, nil, nil)
	return err
}
func (c *OpenstackPortClient) EnablePortSecurity(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/enable_port_security/", id, nil, nil)
	return err
}
func (c *OpenstackPortClient) DisablePortSecurity(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/disable_port_security/", id, nil, nil)
	return err
}
func (c *OpenstackPortClient) Pull(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/pull/", id, nil, nil)
	return err
}
func (c *OpenstackPortClient) Unlink(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-ports/{uuid}/unlink/", id, nil, nil)
	return err
}
