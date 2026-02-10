package resource

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure implementation satisfies interfaces.
var _ action.Action = &MarketplaceResourceUnlinkAction{}
var _ action.ActionWithConfigure = &MarketplaceResourceUnlinkAction{}

type MarketplaceResourceUnlinkModel struct {
	Uuid    types.String `tfsdk:"uuid"`
	Timeout types.String `tfsdk:"timeout"`
}

type MarketplaceResourceUnlinkAction struct {
	client *MarketplaceResourceClient
}

func NewMarketplaceResourceUnlinkAction() action.Action {
	return &MarketplaceResourceUnlinkAction{}
}

func (a *MarketplaceResourceUnlinkAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_resource_unlink"
}

func (a *MarketplaceResourceUnlinkAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "",
		Attributes: map[string]schema.Attribute{
			"uuid": schema.StringAttribute{
				Description: "UUID of the resource",
				Required:    true,
			},
			"timeout": schema.StringAttribute{
				Description: "Timeout for the action execution and state stabilization (e.g. '10m').",
				Optional:    true,
			},
		},
	}
}

func (a *MarketplaceResourceUnlinkAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	a.client = &MarketplaceResourceClient{}
	if err := a.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Action Configure Type",
			err.Error(),
		)
		return
	}
}

func (a *MarketplaceResourceUnlinkAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data MarketplaceResourceUnlinkModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	uuid := data.Uuid.ValueString()
	err := a.client.Unlink(ctx, uuid)

	if err != nil {
		if !common.IsNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Action Failed",
				fmt.Sprintf("Failed to perform unlink on %s: %s", uuid, err),
			)
			return
		}
	}

	// Wait for resource to stabilize
	timeout := common.DefaultActionTimeout
	if !data.Timeout.IsNull() {
		var err error
		timeout, err = time.ParseDuration(data.Timeout.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Invalid timeout", "Failed to parse timeout: "+err.Error())
			return
		}
	}
	err = common.WaitForDeletion(ctx, func(ctx context.Context) (*MarketplaceResourceResponse, error) {
		return a.client.Get(ctx, uuid)
	}, timeout)
	if err != nil {
		resp.Diagnostics.AddWarning("Resource deletion check failed", err.Error())
	}
}
