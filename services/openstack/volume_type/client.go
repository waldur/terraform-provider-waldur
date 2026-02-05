package volume_type

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackVolumeTypeClient struct {
	Client *client.Client
}

func NewOpenstackVolumeTypeClient(c *client.Client) *OpenstackVolumeTypeClient {
	return &OpenstackVolumeTypeClient{Client: c}
}

func (c *OpenstackVolumeTypeClient) Configure(ctx context.Context, providerData interface{}) error {
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

func (c *OpenstackVolumeTypeClient) Get(ctx context.Context, id string) (*OpenstackVolumeTypeResponse, error) {
	var apiResp OpenstackVolumeTypeResponse
	err := c.Client.Get(ctx, "/api/openstack-volume-types/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *OpenstackVolumeTypeClient) List(ctx context.Context, filter map[string]string) ([]OpenstackVolumeTypeResponse, error) {
	var listResult []OpenstackVolumeTypeResponse
	err := c.Client.List(ctx, "/api/openstack-volume-types/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
