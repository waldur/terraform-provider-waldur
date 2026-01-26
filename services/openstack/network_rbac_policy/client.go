package network_rbac_policy

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

func (c *Client) CreateOpenstackNetworkRbacPolicy(ctx context.Context, req *OpenstackNetworkRbacPolicyCreateRequest) (*OpenstackNetworkRbacPolicyResponse, error) {
	var apiResp OpenstackNetworkRbacPolicyResponse
	path := "/api/openstack-network-rbac-policies/"

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) GetOpenstackNetworkRbacPolicy(ctx context.Context, id string) (*OpenstackNetworkRbacPolicyResponse, error) {
	var apiResp OpenstackNetworkRbacPolicyResponse
	err := c.Client.Get(ctx, "/api/openstack-network-rbac-policies/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) UpdateOpenstackNetworkRbacPolicy(ctx context.Context, id string, req *OpenstackNetworkRbacPolicyUpdateRequest) (*OpenstackNetworkRbacPolicyResponse, error) {
	var apiResp OpenstackNetworkRbacPolicyResponse
	err := c.Client.Update(ctx, "/api/openstack-network-rbac-policies/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
func (c *Client) DeleteOpenstackNetworkRbacPolicy(ctx context.Context, id string) error {
	return c.Client.Delete(ctx, "/api/openstack-network-rbac-policies/{uuid}/", id)
}

func (c *Client) ListOpenstackNetworkRbacPolicy(ctx context.Context, filter map[string]string) ([]OpenstackNetworkRbacPolicyResponse, error) {
	var listResult []OpenstackNetworkRbacPolicyResponse
	err := c.Client.List(ctx, "/api/openstack-network-rbac-policies/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
