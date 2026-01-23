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

// CreateOpenstackInstanceOrder creates a marketplace order for this resource.
func (c *Client) CreateOpenstackInstanceOrder(ctx context.Context, req *OpenstackInstanceCreateRequest) (*common.OrderDetails, error) {
	var apiResp common.OrderDetails
	err := c.Client.Post(ctx, "/api/marketplace-orders/", req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// TerminateOpenstackInstance terminates the resource via marketplace.
func (c *Client) TerminateOpenstackInstance(ctx context.Context, id string, req map[string]interface{}) (string, error) {
	var res map[string]interface{}
	err := c.Client.ExecuteAction(ctx, "/api/marketplace-resources/{uuid}/terminate/", id, req, &res)
	if err != nil {
		return "", err
	}
	if uuid, ok := res["uuid"].(string); ok {
		return uuid, nil
	}
	return "", nil
}

// GetOpenstackInstance retrieves a resource by its UUID.
func (c *Client) GetOpenstackInstance(ctx context.Context, id string) (*OpenstackInstanceResponse, error) {
	var apiResp OpenstackInstanceResponse
	err := c.Client.Get(ctx, "/api/openstack-instances/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// UpdateOpenstackInstance updates an existing resource.
func (c *Client) UpdateOpenstackInstance(ctx context.Context, id string, req *OpenstackInstanceUpdateRequest) (*OpenstackInstanceResponse, error) {
	var apiResp OpenstackInstanceResponse
	err := c.Client.Update(ctx, "/api/openstack-instances/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// ListOpenstackInstance retrieves a list of resources with optional filtering.
func (c *Client) ListOpenstackInstance(ctx context.Context, filter map[string]string) ([]OpenstackInstanceResponse, error) {
	var listResult []OpenstackInstanceResponse
	err := c.Client.List(ctx, "/api/openstack-instances/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

// OpenstackInstanceUpdateFloatingIps executes the update_floating_ips action.
func (c *Client) OpenstackInstanceUpdateFloatingIps(ctx context.Context, id string, req *OpenstackInstanceUpdateFloatingIpsActionRequest) error {
	path := "/api/openstack-instances/{uuid}/update_floating_ips/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}

// OpenstackInstanceUpdatePorts executes the update_ports action.
func (c *Client) OpenstackInstanceUpdatePorts(ctx context.Context, id string, req *OpenstackInstanceUpdatePortsActionRequest) error {
	path := "/api/openstack-instances/{uuid}/update_ports/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}

// OpenstackInstanceUpdateSecurityGroups executes the update_security_groups action.
func (c *Client) OpenstackInstanceUpdateSecurityGroups(ctx context.Context, id string, req *OpenstackInstanceUpdateSecurityGroupsActionRequest) error {
	path := "/api/openstack-instances/{uuid}/update_security_groups/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}

// OpenstackInstanceStart executes the start action.
func (c *Client) OpenstackInstanceStart(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-instances/{uuid}/start/", id, nil, nil)
	return err
}

// OpenstackInstanceStop executes the stop action.
func (c *Client) OpenstackInstanceStop(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-instances/{uuid}/stop/", id, nil, nil)
	return err
}

// OpenstackInstanceRestart executes the restart action.
func (c *Client) OpenstackInstanceRestart(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-instances/{uuid}/restart/", id, nil, nil)
	return err
}

// OpenstackInstancePull executes the pull action.
func (c *Client) OpenstackInstancePull(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-instances/{uuid}/pull/", id, nil, nil)
	return err
}

// OpenstackInstanceUnlink executes the unlink action.
func (c *Client) OpenstackInstanceUnlink(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-instances/{uuid}/unlink/", id, nil, nil)
	return err
}
