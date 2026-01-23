package customer

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

// StructureCustomer Operations

func (c *Client) CreateStructureCustomer(ctx context.Context, req *StructureCustomerCreateRequest) (*StructureCustomerResponse, error) {
	var apiResp StructureCustomerResponse
	path := "/api/customers/"

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) GetStructureCustomer(ctx context.Context, id string) (*StructureCustomerResponse, error) {
	var apiResp StructureCustomerResponse
	err := c.Client.GetByUUID(ctx, "/api/customers/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) UpdateStructureCustomer(ctx context.Context, id string, req *StructureCustomerUpdateRequest) (*StructureCustomerResponse, error) {
	var apiResp StructureCustomerResponse
	err := c.Client.Update(ctx, "/api/customers/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
func (c *Client) DeleteStructureCustomer(ctx context.Context, id string) error {
	return c.Client.DeleteByUUID(ctx, "/api/customers/{uuid}/", id)
}

func (c *Client) ListStructureCustomer(ctx context.Context, filter map[string]string) ([]StructureCustomerResponse, error) {
	var listResult []StructureCustomerResponse
	path := "/api/customers/"
	err := c.Client.ListWithFilter(ctx, path, filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
