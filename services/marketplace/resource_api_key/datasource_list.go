package resource_api_key

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &MarketplaceResourceApiKeyListDataSource{}

func NewMarketplaceResourceApiKeyListDataSource() datasource.DataSource {
	return &MarketplaceResourceApiKeyListDataSource{}
}

// MarketplaceResourceApiKeyListDataSource is the plural form of the marketplace_resource_api_key data
// source: it returns every match of its filters, across all pages, where the
// singular data source requires exactly one.
type MarketplaceResourceApiKeyListDataSource struct {
	client *MarketplaceResourceApiKeyClient
}

type MarketplaceResourceApiKeyListDataSourceModel struct {
	Filters *MarketplaceResourceApiKeyFiltersModel `tfsdk:"filters"`
	Items   []MarketplaceResourceApiKeyModel       `tfsdk:"items"`
}

func (d *MarketplaceResourceApiKeyListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_resource_api_keys"
}

func (d *MarketplaceResourceApiKeyListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "List of Marketplace Resource Api Key matching the filters",

		Attributes: map[string]schema.Attribute{
			"filters": (&MarketplaceResourceApiKeyFiltersModel{}).GetSchema(),
			"items": schema.ListNestedAttribute{
				Computed:            true,
				MarkdownDescription: "Marketplace Resource Api Key objects matching the filters, in the order the API returns them",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Marketplace Resource Api Key UUID",
						},
						"allowed_models": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true, MarkdownDescription: "The key's own model list, as chosen."},
						"client_id": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Client Id"},
						"description": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Description"},
						"effective_models": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true, MarkdownDescription: "The models the key can call: the resource's selection, then the key's list. The site agent applies this one."},
						"error_message": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Error Message"},
						"estimated_cost": schema.StringAttribute{
							Computed: true, MarkdownDescription: "What the key's usage this month costs at the unit prices of the resource's plan, without discounts or credits. Zero until the first usage report of the month; recomputed on every report.",
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,12}(?:\.\d{0,10})?$`), ""),
							}},
						"issued_at": schema.StringAttribute{
							CustomType: timetypes.RFC3339Type{},
							Computed:   true, MarkdownDescription: "When the agent last stored a value for this key: its creation or its latest rotation. Unlike modified, a pause or an edit leaves it alone."},
						"limit_resumes_at": schema.StringAttribute{
							CustomType: timetypes.RFC3339Type{},
							Computed:   true, MarkdownDescription: "When every limit that paused the key has reset, so that it resumes without a new usage report. Empty while a reached total limit holds the key: a total never resets, so the key resumes only when that limit is raised or cleared."},
						"limits": schema.MapAttribute{
							ElementType: types.Float64Type,
							Computed:    true, MarkdownDescription: "The month limits of spending_limits, per component type, for clients that know only monthly limits."},
						"offering_uuid": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Offering Uuid"},
						"paused_by_limit": schema.BoolAttribute{
							Computed: true, MarkdownDescription: "Waldur paused the key because its usage reached a limit. Such a key is resumed automatically once its usage is under its limits again: when the period of the limit resets, or when a limit is raised. A total limit never resets, so only raising or clearing it does."},
						"pending_action": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Pending Action"},
						"resource_backend_id": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Resource Backend Id"},
						"resource_name": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Resource Name"},
						"resource_plan_uuid": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Resource Plan Uuid"},
						"resource_spending_limits": schema.MapAttribute{
							ElementType: types.StringType,
							Computed:    true, MarkdownDescription: "The resource's spending limits, per component and period, which cap every key of the resource."},
						"resource_uuid": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Resource Uuid"},
						"spending_limit_status": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"component_type": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Component Type"},
									"limit": schema.Float64Attribute{
										Computed: true, MarkdownDescription: "Limit"},
									"period": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Period",
										Validators: []validator.String{
											stringvalidator.OneOf("day", "week", "month", "quarter", "year", "total"),
										}},
									"reached": schema.BoolAttribute{
										Computed: true, MarkdownDescription: "Reached"},
									"resets_at": schema.StringAttribute{
										CustomType: timetypes.RFC3339Type{},
										Computed:   true, MarkdownDescription: "When the current period ends and its usage starts from zero. Null for total, which never resets."},
									"used": schema.Float64Attribute{
										Computed: true, MarkdownDescription: "Usage reported in the current period."},
								},
							},
							Computed: true, MarkdownDescription: "Spending Limit Status",
						},
						"spending_limits": schema.MapAttribute{
							ElementType: types.StringType,
							Computed:    true, MarkdownDescription: "Limits per component type and period."},
						"state": schema.StringAttribute{
							Computed: true, MarkdownDescription: "State",
							Validators: []validator.String{
								stringvalidator.OneOf("Creating", "OK", "Updating", "Paused", "Deleting", "Deleted", "Erred"),
							}},
						"usage_period": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Usage Period"},
						"user_full_name": schema.StringAttribute{
							Computed: true, MarkdownDescription: "User Full Name"},
						"user_uuid": schema.StringAttribute{
							Computed: true, MarkdownDescription: "User Uuid"},
					},
				},
			},
		},
	}
}

func (d *MarketplaceResourceApiKeyListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	d.client = &MarketplaceResourceApiKeyClient{}
	if err := d.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			err.Error(),
		)
		return
	}
}

func (d *MarketplaceResourceApiKeyListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data MarketplaceResourceApiKeyListDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	filters := common.BuildQueryFilters(data.Filters)

	results, err := d.client.ListAll(ctx, filters)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to List Marketplace Resource Api Key",
			"An error occurred while listing Marketplace Resource Api Key: "+err.Error(),
		)
		return
	}

	data.Items = make([]MarketplaceResourceApiKeyModel, len(results))
	for i := range results {
		resp.Diagnostics.Append(data.Items[i].CopyFrom(ctx, results[i])...)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
