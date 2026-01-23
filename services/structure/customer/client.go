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
	err := c.Client.ListWithFilter(ctx, "/api/customers/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
