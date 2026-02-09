package subnet

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackSubnetClient struct {
	Client *client.Client
}

func NewOpenstackSubnetClient(c *client.Client) *OpenstackSubnetClient {
	return &OpenstackSubnetClient{Client: c}
}

func (c *OpenstackSubnetClient) Configure(ctx context.Context, providerData interface{}) error {
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

func (c *OpenstackSubnetClient) Create(ctx context.Context, network string, req *OpenstackSubnetCreateRequest) (*OpenstackSubnetResponse, error) {
	var apiResp OpenstackSubnetResponse
	networkUUID := common.ExtractUUIDFromURL(network)
	err := c.Client.ExecuteAction(ctx, "/api/openstack-networks/{uuid}/create_subnet/", networkUUID, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *OpenstackSubnetClient) Get(ctx context.Context, id string) (*OpenstackSubnetResponse, error) {
	var apiResp OpenstackSubnetResponse
	err := c.Client.Get(ctx, "/api/openstack-subnets/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *OpenstackSubnetClient) Update(ctx context.Context, id string, req *OpenstackSubnetUpdateRequest) (*OpenstackSubnetResponse, error) {
	var apiResp OpenstackSubnetResponse
	err := c.Client.Update(ctx, "/api/openstack-subnets/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
func (c *OpenstackSubnetClient) Delete(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/openstack-subnets/{uuid}/", id)
}

func (c *OpenstackSubnetClient) List(ctx context.Context, filter map[string]string) ([]OpenstackSubnetResponse, error) {
	var listResult []OpenstackSubnetResponse
	err := c.Client.List(ctx, "/api/openstack-subnets/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *OpenstackSubnetClient) Connect(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-subnets/{uuid}/connect/", id, nil, nil)
	return err
}
func (c *OpenstackSubnetClient) Disconnect(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-subnets/{uuid}/disconnect/", id, nil, nil)
	return err
}
func (c *OpenstackSubnetClient) Pull(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-subnets/{uuid}/pull/", id, nil, nil)
	return err
}
func (c *OpenstackSubnetClient) Unlink(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-subnets/{uuid}/unlink/", id, nil, nil)
	return err
}
