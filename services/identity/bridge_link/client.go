package bridge_link

import (
	"context"
	"fmt"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type IdentityBridgeLinkClient struct {
	Client *client.Client
}

func NewIdentityBridgeLinkClient(c *client.Client) *IdentityBridgeLinkClient {
	return &IdentityBridgeLinkClient{Client: c}
}

func (c *IdentityBridgeLinkClient) Configure(ctx context.Context, providerData interface{}) error {
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

func (c *IdentityBridgeLinkClient) Create(ctx context.Context, req *IdentityBridgeLinkCreateRequest) (*IdentityBridgeLinkResponse, error) {
	var apiResp IdentityBridgeLinkResponse
	path := "/api/identity-bridge/"

	err := c.Client.Post(ctx, path, req, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *IdentityBridgeLinkClient) Get(ctx context.Context, id string) (*IdentityBridgeLinkResponse, error) {
	var apiResp IdentityBridgeLinkResponse
	err := c.Client.Get(ctx, "<no value>", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *IdentityBridgeLinkClient) List(ctx context.Context, filter map[string]string) ([]IdentityBridgeLinkResponse, error) {
	var listResult []IdentityBridgeLinkResponse
	err := c.Client.List(ctx, "<no value>", filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}
