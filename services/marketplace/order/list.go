package order

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/waldur/terraform-provider-waldur/internal/client"
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
			"filters": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Filter parameters for querying Marketplace Order",
				Attributes: map[string]schema.Attribute{
					"can_approve_as_consumer": schema.BoolAttribute{
						Description: "Can approve as consumer",
						Optional:    true,
					},
					"can_approve_as_provider": schema.BoolAttribute{
						Description: "Can approve as provider",
						Optional:    true,
					},
					"category_uuid": schema.StringAttribute{
						Description: "Category UUID",
						Optional:    true,
					},
					"created": schema.StringAttribute{
						Description: "Created after",
						Optional:    true,
					},
					"customer_uuid": schema.StringAttribute{
						Description: "Customer UUID",
						Optional:    true,
					},
					"modified": schema.StringAttribute{
						Description: "Modified after",
						Optional:    true,
					},
					"offering": schema.StringAttribute{
						Description: "",
						Optional:    true,
					},
					"offering_uuid": schema.StringAttribute{
						Description: "Offering UUID",
						Optional:    true,
					},
					"page": schema.Int64Attribute{
						Description: "A page number within the paginated result set.",
						Optional:    true,
					},
					"page_size": schema.Int64Attribute{
						Description: "Number of results to return per page.",
						Optional:    true,
					},
					"parent_offering_uuid": schema.StringAttribute{
						Description: "",
						Optional:    true,
					},
					"project_uuid": schema.StringAttribute{
						Description: "Project UUID",
						Optional:    true,
					},
					"provider_uuid": schema.StringAttribute{
						Description: "Provider UUID",
						Optional:    true,
					},
					"query": schema.StringAttribute{
						Description: "Search by order UUID, slug, project name or resource name",
						Optional:    true,
					},
					"resource": schema.StringAttribute{
						Description: "Resource URL",
						Optional:    true,
					},
					"resource_name": schema.StringAttribute{
						Description: "Resource name",
						Optional:    true,
					},
					"resource_uuid": schema.StringAttribute{
						Description: "Resource UUID",
						Optional:    true,
					},
					"service_manager_uuid": schema.StringAttribute{
						Description: "Service manager UUID",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (l *MarketplaceOrderList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected List Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T", req.ProviderData),
		)
		return
	}

	l.client = NewClient(client)
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
