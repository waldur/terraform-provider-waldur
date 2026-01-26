package volume

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

// CreateOpenstackVolumeOrder creates a marketplace order for this resource.
func (c *Client) CreateOpenstackVolumeOrder(ctx context.Context, req *OpenstackVolumeCreateRequest) (*common.OrderDetails, error) {
	var apiResp common.OrderDetails
	err := c.Client.Post(ctx, "/api/marketplace-orders/", req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

// TerminateOpenstackVolume terminates the resource via marketplace.
func (c *Client) TerminateOpenstackVolume(ctx context.Context, id string, req map[string]interface{}) (string, error) {
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

func (c *Client) GetOpenstackVolume(ctx context.Context, id string) (*OpenstackVolumeResponse, error) {
	var apiResp OpenstackVolumeResponse
	err := c.Client.Get(ctx, "/api/openstack-volumes/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) UpdateOpenstackVolume(ctx context.Context, id string, req *OpenstackVolumeUpdateRequest) (*OpenstackVolumeResponse, error) {
	var apiResp OpenstackVolumeResponse
	err := c.Client.Update(ctx, "/api/openstack-volumes/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) ListOpenstackVolume(ctx context.Context, filter map[string]string) ([]OpenstackVolumeResponse, error) {
	var listResult []OpenstackVolumeResponse
	err := c.Client.List(ctx, "/api/openstack-volumes/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *Client) OpenstackVolumeExtend(ctx context.Context, id string, req *OpenstackVolumeExtendActionRequest) error {
	path := "/api/openstack-volumes/{uuid}/extend/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}
func (c *Client) OpenstackVolumeRetype(ctx context.Context, id string, req *OpenstackVolumeRetypeActionRequest) error {
	path := "/api/openstack-volumes/{uuid}/retype/"
	err := c.Client.ExecuteAction(ctx, path, id, req, nil)
	return err
}
func (c *Client) OpenstackVolumePull(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-volumes/{uuid}/pull/", id, nil, nil)
	return err
}
func (c *Client) OpenstackVolumeUnlink(ctx context.Context, id string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-volumes/{uuid}/unlink/", id, nil, nil)
	return err
}
