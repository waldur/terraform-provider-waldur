package project

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type StructureProjectClient struct {
	Client *client.Client
}

func NewStructureProjectClient(c *client.Client) *StructureProjectClient {
	return &StructureProjectClient{Client: c}
}

func (c *StructureProjectClient) Configure(ctx context.Context, providerData interface{}) error {
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

func (c *StructureProjectClient) Create(ctx context.Context, req *StructureProjectCreateRequest) (*StructureProjectResponse, error) {
	var apiResp StructureProjectResponse
	path := "/api/projects/"

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *StructureProjectClient) Get(ctx context.Context, id string) (*StructureProjectResponse, error) {
	var apiResp StructureProjectResponse
	err := c.Client.Get(ctx, "/api/projects/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *StructureProjectClient) Update(ctx context.Context, id string, req *StructureProjectUpdateRequest) (*StructureProjectResponse, error) {
	var apiResp StructureProjectResponse
	err := c.Client.Update(ctx, "/api/projects/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
func (c *StructureProjectClient) Delete(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/projects/{uuid}/", id)
}

func (c *StructureProjectClient) List(ctx context.Context, filter map[string]string) ([]StructureProjectResponse, error) {
	var listResult []StructureProjectResponse
	err := c.Client.List(ctx, "/api/projects/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
