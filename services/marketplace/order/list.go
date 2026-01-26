package order

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

var _ list.ListResource = &MarketplaceOrderList{}

type MarketplaceOrderList struct {
	client *Client
}

func NewMarketplaceOrderList() list.ListResource {
	return &MarketplaceOrderList{}
}

func (l *MarketplaceOrderList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_order"
}

func (l *MarketplaceOrderList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"filters": (&MarketplaceOrderFiltersModel{}).GetSchema(),
		},
	}
}

func (l *MarketplaceOrderList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	l.client = &Client{}
	if err := l.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected List Configure Type",
			err.Error(),
		)
		return
	}
}

type MarketplaceOrderListModel struct {
	Filters *MarketplaceOrderFiltersModel `tfsdk:"filters"`
}

func (l *MarketplaceOrderList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config MarketplaceOrderListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := common.BuildQueryFilters(config.Filters)

	// Call API
	listResult, err := l.client.ListMarketplaceOrder(ctx, filters)
	if err != nil {
		// Return error diagnostics
		resp.AddError("Failed to list resources", err.Error())
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Stream results
	stream.Results = func(push func(list.ListResult) bool) {
		for _, apiResp := range listResult {
			result := req.NewListResult(ctx)

			// Map item to model
			var data MarketplaceOrderResourceModel
			model := &data

			var diags diag.Diagnostics

			diags.Append(model.CopyFrom(ctx, apiResp)...)

			// Set the resource state
			// For ListResource, we generally return the "Resource" state matching the main resource schema.
			// This allows `terraform plan` to correlate if using `terraform query`.

			setDiags := result.Resource.Set(ctx, &data)
			diags.Append(setDiags...)
			result.Diagnostics.Append(diags...)

			// Set Identity if possible (usually UUID)
			if !data.UUID.IsNull() && !data.UUID.IsUnknown() {
				result.Diagnostics.Append(result.Identity.Set(ctx, data.UUID.ValueString())...)
			}

			if !push(result) {
				return
			}
		}
	}
}
