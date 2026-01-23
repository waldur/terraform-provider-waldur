package project

import (
	"context"
	"fmt"
	"strings"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

type Client struct {
	Client *client.Client
}

func NewClient(c *client.Client) *Client {
	_ = fmt.Errorf
	_ = strings.Split
	return &Client{Client: c}
}

func (c *Client) Create(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.Client.Create(ctx, path, body, result)
}

func (c *Client) GetByUUID(ctx context.Context, path string, uuid string, result interface{}) error {
	return c.Client.GetByUUID(ctx, path, uuid, result)
}

func (c *Client) Update(ctx context.Context, path string, uuid string, body interface{}, result interface{}) error {
	return c.Client.Update(ctx, path, uuid, body, result)
}

func (c *Client) DeleteByUUID(ctx context.Context, path string, uuid string) error {
	return c.Client.DeleteByUUID(ctx, path, uuid)
}

func (c *Client) Post(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.Client.Post(ctx, path, body, result)
}

func (c *Client) ListWithFilter(ctx context.Context, path string, filter map[string]string, result interface{}) error {
	return c.Client.ListWithFilter(ctx, path, filter, result)
}

func (c *Client) ExecuteAction(ctx context.Context, pathTemplate string, uuid string, body interface{}) (map[string]interface{}, error) {
	path := strings.Replace(pathTemplate, "{uuid}", uuid, 1)
	var res map[string]interface{}
	err := c.Client.Post(ctx, path, body, &res)
	return res, err
}

// StructureProject Operations

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
	err := c.Client.GetByUUID(ctx, "/api/projects/{uuid}/", id, &apiResp)
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
	return c.Client.DeleteByUUID(ctx, "/api/projects/{uuid}/", id)
}

func (c *Client) ListStructureProject(ctx context.Context, filter map[string]string) ([]StructureProjectResponse, error) {
	var listResult []StructureProjectResponse
	path := "/api/projects/"
	err := c.Client.ListWithFilter(ctx, path, filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
