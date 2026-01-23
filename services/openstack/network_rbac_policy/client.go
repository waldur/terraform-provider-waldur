package network_rbac_policy

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

// CreateOpenstackNetworkRbacPolicy creates a new resource.
func (c *Client) CreateOpenstackNetworkRbacPolicy(ctx context.Context, req *OpenstackNetworkRbacPolicyCreateRequest) (*OpenstackNetworkRbacPolicyResponse, error) {
	var apiResp OpenstackNetworkRbacPolicyResponse
	path := "/api/openstack-network-rbac-policies/"

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// GetOpenstackNetworkRbacPolicy retrieves a resource by its UUID.
func (c *Client) GetOpenstackNetworkRbacPolicy(ctx context.Context, id string) (*OpenstackNetworkRbacPolicyResponse, error) {
	var apiResp OpenstackNetworkRbacPolicyResponse
	err := c.Client.Get(ctx, "/api/openstack-network-rbac-policies/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// UpdateOpenstackNetworkRbacPolicy updates an existing resource.
func (c *Client) UpdateOpenstackNetworkRbacPolicy(ctx context.Context, id string, req *OpenstackNetworkRbacPolicyUpdateRequest) (*OpenstackNetworkRbacPolicyResponse, error) {
	var apiResp OpenstackNetworkRbacPolicyResponse
	err := c.Client.Update(ctx, "/api/openstack-network-rbac-policies/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// DeleteOpenstackNetworkRbacPolicy deletes a resource.
func (c *Client) DeleteOpenstackNetworkRbacPolicy(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/openstack-network-rbac-policies/{uuid}/", id)
}

// ListOpenstackNetworkRbacPolicy retrieves a list of resources with optional filtering.
func (c *Client) ListOpenstackNetworkRbacPolicy(ctx context.Context, filter map[string]string) ([]OpenstackNetworkRbacPolicyResponse, error) {
	var listResult []OpenstackNetworkRbacPolicyResponse
	err := c.Client.List(ctx, "/api/openstack-network-rbac-policies/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
