package resource

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

var _ list.ListResource = &MarketplaceResourceList{}

type MarketplaceResourceList struct {
	client *Client
}

func NewMarketplaceResourceList() list.ListResource {
	return &MarketplaceResourceList{}
}

func (l *MarketplaceResourceList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_resource"
}

func (l *MarketplaceResourceList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"filters": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Filter parameters for querying Marketplace Resource",
				Attributes: map[string]schema.Attribute{
					"backend_id": schema.StringAttribute{
						Description: "Backend ID",
						Optional:    true,
					},
					"category_uuid": schema.StringAttribute{
						Description: "Category UUID",
						Optional:    true,
					},
					"component_count": schema.Float64Attribute{
						Description: "Filter by exact number of components",
						Optional:    true,
					},
					"created": schema.StringAttribute{
						Description: "Created after",
						Optional:    true,
					},
					"customer": schema.StringAttribute{
						Description: "Customer URL",
						Optional:    true,
					},
					"customer_uuid": schema.StringAttribute{
						Description: "Customer UUID",
						Optional:    true,
					},
					"downscaled": schema.BoolAttribute{
						Description: "Downscaled",
						Optional:    true,
					},
					"has_terminate_date": schema.BoolAttribute{
						Description: "Has termination date",
						Optional:    true,
					},
					"is_attached": schema.BoolAttribute{
						Description: "Filter by attached state",
						Optional:    true,
					},
					"lexis_links_supported": schema.BoolAttribute{
						Description: "LEXIS links supported",
						Optional:    true,
					},
					"limit_based": schema.BoolAttribute{
						Description: "Filter by limit-based offerings",
						Optional:    true,
					},
					"limit_component_count": schema.Float64Attribute{
						Description: "Filter by exact number of limit-based components",
						Optional:    true,
					},
					"modified": schema.StringAttribute{
						Description: "Modified after",
						Optional:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name",
						Optional:    true,
					},
					"name_exact": schema.StringAttribute{
						Description: "Name (exact)",
						Optional:    true,
					},
					"offering": schema.StringAttribute{
						Description: "",
						Optional:    true,
					},
					"offering_billable": schema.BoolAttribute{
						Description: "Offering billable",
						Optional:    true,
					},
					"offering_shared": schema.BoolAttribute{
						Description: "Offering shared",
						Optional:    true,
					},
					"offering_type": schema.StringAttribute{
						Description: "Offering type",
						Optional:    true,
					},
					"only_limit_based": schema.BoolAttribute{
						Description: "Filter resources with only limit-based components",
						Optional:    true,
					},
					"only_usage_based": schema.BoolAttribute{
						Description: "Filter resources with only usage-based components",
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
					"paused": schema.BoolAttribute{
						Description: "Paused",
						Optional:    true,
					},
					"plan_uuid": schema.StringAttribute{
						Description: "Plan UUID",
						Optional:    true,
					},
					"project_name": schema.StringAttribute{
						Description: "Project name",
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
						Description: "Search by resource UUID, name, slug, backend ID, effective ID, IPs or hypervisor",
						Optional:    true,
					},
					"restrict_member_access": schema.BoolAttribute{
						Description: "Restrict member access",
						Optional:    true,
					},
					"runtime_state": schema.StringAttribute{
						Description: "Runtime state",
						Optional:    true,
					},
					"service_manager_uuid": schema.StringAttribute{
						Description: "Service manager UUID",
						Optional:    true,
					},
					"usage_based": schema.BoolAttribute{
						Description: "Filter by usage-based offerings",
						Optional:    true,
					},
					"visible_to_providers": schema.BoolAttribute{
						Description: "Include only resources visible to service providers",
						Optional:    true,
					},
					"visible_to_username": schema.StringAttribute{
						Description: "Visible to username",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (l *MarketplaceResourceList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type MarketplaceResourceListModel struct {
	Filters *MarketplaceResourceFiltersModel `tfsdk:"filters"`
}

func (l *MarketplaceResourceList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config MarketplaceResourceListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := common.BuildQueryFilters(config.Filters)

	// Call API
	listResult, err := l.client.ListMarketplaceResource(ctx, filters)
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
			var data MarketplaceResourceResourceModel
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
