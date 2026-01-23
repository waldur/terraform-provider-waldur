package volume

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

// OpenstackVolume Operations

func (c *Client) CreateOpenstackVolumeOrder(ctx context.Context, req *OpenstackVolumeCreateRequest) (*common.OrderDetails, error) {
	var apiResp common.OrderDetails
	err := c.Client.Post(ctx, "/api/marketplace-orders/", req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) TerminateOpenstackVolume(ctx context.Context, id string, req map[string]interface{}) (string, error) {
	res, err := c.ExecuteAction(ctx, "/api/marketplace-resources/{uuid}/terminate/", id, req)
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
	err := c.Client.GetByUUID(ctx, "/api/openstack-volumes/{uuid}/", id, &apiResp)
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
	path := "/api/openstack-volumes/"
	err := c.Client.ListWithFilter(ctx, path, filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *Client) OpenstackVolumeExtend(ctx context.Context, id string, req *OpenstackVolumeExtendActionRequest) error {
	path := "/api/openstack-volumes/{uuid}/extend/"
	_, err := c.ExecuteAction(ctx, path, id, req)
	return err
}
func (c *Client) OpenstackVolumeRetype(ctx context.Context, id string, req *OpenstackVolumeRetypeActionRequest) error {
	path := "/api/openstack-volumes/{uuid}/retype/"
	_, err := c.ExecuteAction(ctx, path, id, req)
	return err
}
func (c *Client) OpenstackVolumePull(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-volumes/{uuid}/pull/", id, nil)
	return err
}
func (c *Client) OpenstackVolumeUnlink(ctx context.Context, id string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-volumes/{uuid}/unlink/", id, nil)
	return err
}
