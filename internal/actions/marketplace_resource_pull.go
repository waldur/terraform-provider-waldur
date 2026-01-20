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
var _ action.Action = &MarketplaceResourcePullAction{}
var _ action.ActionWithConfigure = &MarketplaceResourcePullAction{}

type MarketplaceResourcePullAction struct {
	client *client.Client
}

func NewMarketplaceResourcePullAction() action.Action {
	return &MarketplaceResourcePullAction{}
}

func (a *MarketplaceResourcePullAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_resource_pull"
}
func (a *MarketplaceResourcePullAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Perform pull action on marketplace_resource",
		Attributes: map[string]schema.Attribute{
			"uuid": schema.StringAttribute{
				Description: "The UUID of the marketplace_resource",
				Required:    true,
			},
		},
	}
}

func (a *MarketplaceResourcePullAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

type MarketplaceResourcePullModel struct {
	Uuid types.String `tfsdk:"uuid"`
}

func (a *MarketplaceResourcePullAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data MarketplaceResourcePullModel
	// Read validation data
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	uuid := data.Uuid.ValueString()
	// Perform the action
	path := "/api/marketplace-resources/{uuid}/pull/"
	path = strings.Replace(path, "{uuid}", uuid, 1)

	err := a.client.Post(ctx, path, nil, nil)

	if err != nil {
		resp.Diagnostics.AddError(
			"Action Failed",
			fmt.Sprintf("Failed to perform pull on %s: %s", uuid, err),
		)
		return
	}
}
