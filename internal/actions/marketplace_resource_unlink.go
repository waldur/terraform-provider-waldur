package actions

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure implementation satisfies interfaces.
var _ action.Action = &MarketplaceResourceUnlinkAction{}
var _ action.ActionWithConfigure = &MarketplaceResourceUnlinkAction{}

type MarketplaceResourceUnlinkAction struct {
	client *client.Client
}

func NewMarketplaceResourceUnlinkAction() action.Action {
	return &MarketplaceResourceUnlinkAction{}
}

func (a *MarketplaceResourceUnlinkAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_resource_unlink"
}
func (a *MarketplaceResourceUnlinkAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Perform unlink action on marketplace_resource",
		Attributes: map[string]schema.Attribute{
			"uuid": schema.StringAttribute{
				Description: "The UUID of the marketplace_resource",
				Required:    true,
			},
		},
	}
}

func (a *MarketplaceResourceUnlinkAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Action Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T", req.ProviderData),
		)
		return
	}

	a.client = client
}

type MarketplaceResourceUnlinkModel struct {
	Uuid types.String `tfsdk:"uuid"`
}

func (a *MarketplaceResourceUnlinkAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data MarketplaceResourceUnlinkModel
	// Read validation data
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	uuid := data.Uuid.ValueString()
	// Perform the action
	path := "/api/marketplace-resources/{uuid}/unlink/"
	path = strings.Replace(path, "{uuid}", uuid, 1)

	err := a.client.Post(ctx, path, nil, nil)

	if err != nil {
		resp.Diagnostics.AddError(
			"Action Failed",
			fmt.Sprintf("Failed to perform unlink on %s: %s", uuid, err),
		)
		return
	}
}
