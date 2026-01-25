package offering

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

var _ list.ListResource = &MarketplaceOfferingList{}

type MarketplaceOfferingList struct {
	client *Client
}

func NewMarketplaceOfferingList() list.ListResource {
	return &MarketplaceOfferingList{}
}

func (l *MarketplaceOfferingList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_offering"
}

func (l *MarketplaceOfferingList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"filters": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Filter parameters for querying Marketplace Offering",
				Attributes: map[string]schema.Attribute{
					"accessible_via_calls": schema.BoolAttribute{
						Description: "Accessible via calls",
						Optional:    true,
					},
					"allowed_customer_uuid": schema.StringAttribute{
						Description: "Allowed customer UUID",
						Optional:    true,
					},
					"attributes": schema.StringAttribute{
						Description: "Offering attributes (JSON)",
						Optional:    true,
					},
					"billable": schema.BoolAttribute{
						Description: "Billable",
						Optional:    true,
					},
					"can_create_offering_user": schema.BoolAttribute{
						Description: "",
						Optional:    true,
					},
					"category_group_uuid": schema.StringAttribute{
						Description: "Category group UUID",
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
					"customer": schema.StringAttribute{
						Description: "Customer URL",
						Optional:    true,
					},
					"customer_uuid": schema.StringAttribute{
						Description: "Customer UUID",
						Optional:    true,
					},
					"description": schema.StringAttribute{
						Description: "Description contains",
						Optional:    true,
					},
					"has_active_terms_of_service": schema.BoolAttribute{
						Description: "Has Active Terms of Service",
						Optional:    true,
					},
					"has_terms_of_service": schema.BoolAttribute{
						Description: "Has Terms of Service",
						Optional:    true,
					},
					"keyword": schema.StringAttribute{
						Description: "Keyword",
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
					"page": schema.Int64Attribute{
						Description: "A page number within the paginated result set.",
						Optional:    true,
					},
					"page_size": schema.Int64Attribute{
						Description: "Number of results to return per page.",
						Optional:    true,
					},
					"parent_uuid": schema.StringAttribute{
						Description: "Parent offering UUID",
						Optional:    true,
					},
					"project_uuid": schema.StringAttribute{
						Description: "Project UUID",
						Optional:    true,
					},
					"query": schema.StringAttribute{
						Description: "Search by offering name, slug or description",
						Optional:    true,
					},
					"resource_customer_uuid": schema.StringAttribute{
						Description: "Resource customer UUID",
						Optional:    true,
					},
					"resource_project_uuid": schema.StringAttribute{
						Description: "Resource project UUID",
						Optional:    true,
					},
					"scope_uuid": schema.StringAttribute{
						Description: "Scope UUID",
						Optional:    true,
					},
					"service_manager_uuid": schema.StringAttribute{
						Description: "Service manager UUID",
						Optional:    true,
					},
					"shared": schema.BoolAttribute{
						Description: "Shared",
						Optional:    true,
					},
					"user_has_consent": schema.BoolAttribute{
						Description: "User Has Consent",
						Optional:    true,
					},
					"user_has_offering_user": schema.BoolAttribute{
						Description: "User Has Offering User",
						Optional:    true,
					},
					"uuid_list": schema.StringAttribute{
						Description: "Comma-separated offering UUIDs",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (l *MarketplaceOfferingList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type MarketplaceOfferingListModel struct {
	Filters *MarketplaceOfferingFiltersModel `tfsdk:"filters"`
}

func (l *MarketplaceOfferingList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config MarketplaceOfferingListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := common.BuildQueryFilters(config.Filters)

	// Call API
	listResult, err := l.client.ListMarketplaceOffering(ctx, filters)
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
			var data MarketplaceOfferingResourceModel
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
