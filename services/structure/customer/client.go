package customer

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type StructureCustomerClient struct {
	Client *client.Client
}

func NewStructureCustomerClient(c *client.Client) *StructureCustomerClient {
	return &StructureCustomerClient{Client: c}
}

func (c *StructureCustomerClient) Configure(ctx context.Context, providerData interface{}) error {
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

func (c *StructureCustomerClient) Create(ctx context.Context, req *StructureCustomerCreateRequest) (*StructureCustomerResponse, error) {
	var apiResp StructureCustomerResponse
	path := "/api/customers/"

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *StructureCustomerClient) Get(ctx context.Context, id string) (*StructureCustomerResponse, error) {
	var apiResp StructureCustomerResponse
	err := c.Client.Get(ctx, "/api/customers/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *StructureCustomerClient) Update(ctx context.Context, id string, req *StructureCustomerUpdateRequest) (*StructureCustomerResponse, error) {
	var apiResp StructureCustomerResponse
	err := c.Client.Update(ctx, "/api/customers/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
func (c *StructureCustomerClient) Delete(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/customers/{uuid}/", id)
}

func (c *StructureCustomerClient) List(ctx context.Context, filter map[string]string) ([]StructureCustomerResponse, error) {
	var listResult []StructureCustomerResponse
	err := c.Client.List(ctx, "/api/customers/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
