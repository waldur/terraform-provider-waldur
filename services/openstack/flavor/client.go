package flavor

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackFlavorClient struct {
	Client *client.Client
}

func NewOpenstackFlavorClient(c *client.Client) *OpenstackFlavorClient {
	return &OpenstackFlavorClient{Client: c}
}

func (c *OpenstackFlavorClient) Configure(ctx context.Context, providerData interface{}) error {
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

func (c *OpenstackFlavorClient) Get(ctx context.Context, id string) (*OpenstackFlavorResponse, error) {
	var apiResp OpenstackFlavorResponse
	err := c.Client.Get(ctx, "/api/openstack-flavors/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *OpenstackFlavorClient) List(ctx context.Context, filter map[string]string) ([]OpenstackFlavorResponse, error) {
	var listResult []OpenstackFlavorResponse
	err := c.Client.List(ctx, "/api/openstack-flavors/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
