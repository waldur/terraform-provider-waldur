package permission

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type CustomerPermissionClient struct {
	Client *client.Client
}

func NewCustomerPermissionClient(c *client.Client) *CustomerPermissionClient {
	return &CustomerPermissionClient{Client: c}
}

func (c *CustomerPermissionClient) Configure(ctx context.Context, providerData interface{}) error {
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

func (c *CustomerPermissionClient) Create(ctx context.Context, req *CustomerPermissionCreateRequest) (*CustomerPermissionResponse, error) {
	var apiResp CustomerPermissionResponse
	path := "/api/customers/"

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *CustomerPermissionClient) Get(ctx context.Context, id string) (*CustomerPermissionResponse, error) {
	var apiResp CustomerPermissionResponse
	err := c.Client.Get(ctx, "<no value>", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *CustomerPermissionClient) List(ctx context.Context, filter map[string]string) ([]CustomerPermissionResponse, error) {
	var listResult []CustomerPermissionResponse
	err := c.Client.List(ctx, "<no value>", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
