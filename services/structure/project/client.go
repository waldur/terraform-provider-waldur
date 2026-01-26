package project

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type Client struct {
	Client *client.Client
}

func NewClient(c *client.Client) *Client {
	return &Client{Client: c}
}

func (c *Client) Configure(ctx context.Context, providerData interface{}) error {
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

func (c *Client) CreateStructureProject(ctx context.Context, req *StructureProjectCreateRequest) (*StructureProjectResponse, error) {
	var apiResp StructureProjectResponse
	path := "/api/projects/"

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) GetStructureProject(ctx context.Context, id string) (*StructureProjectResponse, error) {
	var apiResp StructureProjectResponse
	err := c.Client.Get(ctx, "/api/projects/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) UpdateStructureProject(ctx context.Context, id string, req *StructureProjectUpdateRequest) (*StructureProjectResponse, error) {
	var apiResp StructureProjectResponse
	err := c.Client.Update(ctx, "/api/projects/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
func (c *Client) DeleteStructureProject(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/projects/{uuid}/", id)
}

func (c *Client) ListStructureProject(ctx context.Context, filter map[string]string) ([]StructureProjectResponse, error) {
	var listResult []StructureProjectResponse
	err := c.Client.List(ctx, "/api/projects/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
