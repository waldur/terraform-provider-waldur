package volume_attachment

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

// OpenstackVolumeAttachment Operations

func (c *Client) GetOpenstackVolumeAttachment(ctx context.Context, id string) (*OpenstackVolumeAttachmentResponse, error) {
	var apiResp OpenstackVolumeAttachmentResponse
	err := c.Client.GetByUUID(ctx, "/api/openstack-volumes/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) UpdateOpenstackVolumeAttachment(ctx context.Context, id string, req *OpenstackVolumeAttachmentUpdateRequest) (*OpenstackVolumeAttachmentResponse, error) {
	var apiResp OpenstackVolumeAttachmentResponse
	err := c.Client.Update(ctx, "/api/openstack-volumes/{uuid}/", id, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *Client) ListOpenstackVolumeAttachment(ctx context.Context, filter map[string]string) ([]OpenstackVolumeAttachmentResponse, error) {
	var listResult []OpenstackVolumeAttachmentResponse
	path := "/api/openstack-volumes/"
	err := c.Client.ListWithFilter(ctx, path, filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *Client) LinkOpenstackVolumeAttachment(ctx context.Context, req *OpenstackVolumeAttachmentCreateRequest) (*OpenstackVolumeAttachmentResponse, error) {
	sourceUUID := *req.Volume
	path := strings.Replace("/api/openstack-volumes/{uuid}/attach/", "{uuid}", sourceUUID, 1)

	var apiResp OpenstackVolumeAttachmentResponse
	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
func (c *Client) UnlinkOpenstackVolumeAttachment(ctx context.Context, sourceUUID string) error {
	_, err := c.ExecuteAction(ctx, "/api/openstack-volumes/{uuid}/detach/", sourceUUID, nil)
	return err
}
