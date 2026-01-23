package instance

import (
	"context"
	"fmt"
	"strings"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
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

// OpenstackInstance Operations

func (c *Client) CreateOpenstackInstanceOrder(ctx context.Context, req *OpenstackInstanceCreateRequest) (*common.OrderDetails, error) {
	var apiResp common.OrderDetails
	err := c.Client.Post(ctx, "/api/marketplace-orders/", req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) TerminateOpenstackInstance(ctx context.Context, id string, req map[string]interface{}) (string, error) {
	res, err := c.ExecuteAction(ctx, "/api/marketplace-resources/{uuid}/terminate/", id, req)
	if err != nil {
		return "", err
	}
	if uuid, ok := res["uuid"].(string); ok {
		return uuid, nil
	}
	return "", nil
}

func (c *Client) GetOpenstackInstance(ctx context.Context, id string) (*OpenstackInstanceResponse, error) {
	var apiResp OpenstackInstanceResponse
	err := c.Client.GetByUUID(ctx, "/api/openstack-instances/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) UpdateOpenstackInstance(ctx context.Context, id string, req *OpenstackInstanceUpdateRequest) (*OpenstackInstanceResponse, error) {
	var apiResp OpenstackInstanceResponse
	err := c.Client.Update(ctx, "/api/openstack-instances/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) ListOpenstackInstance(ctx context.Context, filter map[string]string) ([]OpenstackInstanceResponse, error) {
	var listResult []OpenstackInstanceResponse
	path := "/api/openstack-instances/"
	err := c.Client.ListWithFilter(ctx, path, filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *Client) OpenstackInstanceUpdateFloatingIps(ctx context.Context, id string, req *OpenstackInstanceUpdateFloatingIpsActionRequest) error {
	path := "/api/openstack-instances/{uuid}/update_floating_ips/"
	_, err := c.ExecuteAction(ctx, path, id, req)
	return err
}
func (c *Client) OpenstackInstanceUpdatePorts(ctx context.Context, id string, req *OpenstackInstanceUpdatePortsActionRequest) error {
	path := "/api/openstack-instances/{uuid}/update_ports/"
	_, err := c.ExecuteAction(ctx, path, id, req)
	return err
}
func (c *Client) OpenstackInstanceUpdateSecurityGroups(ctx context.Context, id string, req *OpenstackInstanceUpdateSecurityGroupsActionRequest) error {
	path := "/api/openstack-instances/{uuid}/update_security_groups/"
	_, err := c.ExecuteAction(ctx, path, id, req)
	return err
}
func (c *Client) OpenstackInstanceStart(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-instances/{uuid}/start/", id, nil)
	return err
}
func (c *Client) OpenstackInstanceStop(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-instances/{uuid}/stop/", id, nil)
	return err
}
func (c *Client) OpenstackInstanceRestart(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-instances/{uuid}/restart/", id, nil)
	return err
}
func (c *Client) OpenstackInstancePull(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-instances/{uuid}/pull/", id, nil)
	return err
}
func (c *Client) OpenstackInstanceUnlink(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-instances/{uuid}/unlink/", id, nil)
	return err
}
