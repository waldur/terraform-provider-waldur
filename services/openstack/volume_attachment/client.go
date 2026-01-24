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

func (c *Client) GetOpenstackVolumeAttachment(ctx context.Context, id string) (*OpenstackVolumeAttachmentResponse, error) {
	var apiResp OpenstackVolumeAttachmentResponse
	err := c.Client.Get(ctx, "/api/openstack-volumes/{uuid}/", id, &apiResp)
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
	err := c.Client.List(ctx, "/api/openstack-volumes/", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

func (c *Client) LinkOpenstackVolumeAttachment(ctx context.Context, req *OpenstackVolumeAttachmentCreateRequest) (*OpenstackVolumeAttachmentResponse, error) {
	sourceUUID := *req.Volume

	var apiResp OpenstackVolumeAttachmentResponse
	err := c.Client.ExecuteAction(ctx, "/api/openstack-volumes/{uuid}/attach/", sourceUUID, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
func (c *Client) UnlinkOpenstackVolumeAttachment(ctx context.Context, sourceUUID string) error {
	err := c.Client.ExecuteAction(ctx, "/api/openstack-volumes/{uuid}/detach/", sourceUUID, nil, nil)
	return err
}
