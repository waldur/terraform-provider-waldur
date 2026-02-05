package network

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackNetworkClient struct {
	Client *client.Client
}

func NewOpenstackNetworkClient(c *client.Client) *OpenstackNetworkClient {
	return &OpenstackNetworkClient{Client: c}
}

func (c *OpenstackNetworkClient) Configure(ctx context.Context, providerData interface{}) error {
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

func (c *OpenstackNetworkClient) Create(ctx context.Context, tenant string, req *OpenstackNetworkCreateRequest) (*OpenstackNetworkResponse, error) {
	var apiResp OpenstackNetworkResponse
	err := c.Client.ExecuteAction(ctx, "/api/openstack-tenants/{uuid}/create_network/", tenant, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *OpenstackNetworkClient) Get(ctx context.Context, id string) (*OpenstackNetworkResponse, error) {
	var apiResp OpenstackNetworkResponse
	err := c.Client.Get(ctx, "/api/openstack-networks/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *OpenstackNetworkClient) Update(ctx context.Context, id string, req *OpenstackNetworkUpdateRequest) (*OpenstackNetworkResponse, error) {
	var apiResp OpenstackNetworkResponse
	err := c.Client.Update(ctx, "/api/openstack-networks/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
func (c *OpenstackNetworkClient) Delete(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/openstack-networks/{uuid}/", id)
}

func (c *OpenstackNetworkClient) List(ctx context.Context, filter map[string]string) ([]OpenstackNetworkResponse, error) {
	var listResult []OpenstackNetworkResponse
	err := c.Client.List(ctx, "/api/openstack-networks/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *OpenstackNetworkClient) SetMtu(ctx context.Context, id string, req *OpenstackNetworkSetMtuActionRequest) error {
	path := "/api/openstack-networks/{uuid}/set_mtu/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}
func (c *OpenstackNetworkClient) Pull(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-networks/{uuid}/pull/", id, nil, nil)
	return err
}
func (c *OpenstackNetworkClient) Unlink(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-networks/{uuid}/unlink/", id, nil, nil)
	return err
}
