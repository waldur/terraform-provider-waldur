package image

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackImageClient struct {
	Client *client.Client
}

func NewOpenstackImageClient(c *client.Client) *OpenstackImageClient {
	return &OpenstackImageClient{Client: c}
}

func (c *OpenstackImageClient) Configure(ctx context.Context, providerData interface{}) error {
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

func (c *OpenstackImageClient) Get(ctx context.Context, id string) (*OpenstackImageResponse, error) {
	var apiResp OpenstackImageResponse
	err := c.Client.Get(ctx, "/api/openstack-images/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *OpenstackImageClient) List(ctx context.Context, filter map[string]string) ([]OpenstackImageResponse, error) {
	var listResult []OpenstackImageResponse
	err := c.Client.List(ctx, "/api/openstack-images/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
