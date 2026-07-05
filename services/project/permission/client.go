package permission

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type ProjectPermissionClient struct {
	Client *client.Client
}

func NewProjectPermissionClient(c *client.Client) *ProjectPermissionClient {
	return &ProjectPermissionClient{Client: c}
}

func (c *ProjectPermissionClient) Configure(ctx context.Context, providerData interface{}) error {
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

func (c *ProjectPermissionClient) Create(ctx context.Context, req *ProjectPermissionCreateRequest) (*ProjectPermissionResponse, error) {
	var apiResp ProjectPermissionResponse
	path := "/api/projects/"

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *ProjectPermissionClient) Get(ctx context.Context, id string) (*ProjectPermissionResponse, error) {
	var apiResp ProjectPermissionResponse
	err := c.Client.Get(ctx, "<no value>", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *ProjectPermissionClient) List(ctx context.Context, filter map[string]string) ([]ProjectPermissionResponse, error) {
	var listResult []ProjectPermissionResponse
	err := c.Client.List(ctx, "<no value>", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
