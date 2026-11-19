package resource_api_key

import (
	"context"
	"fmt"
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
var _ datasource.DataSource = &MarketplaceResourceApiKeyDataSource{}

func NewMarketplaceResourceApiKeyDataSource() datasource.DataSource {
	return &MarketplaceResourceApiKeyDataSource{}
}

type MarketplaceResourceApiKeyDataSource struct {
	client *MarketplaceResourceApiKeyClient
}

type MarketplaceResourceApiKeyDataSourceModel struct {
	MarketplaceResourceApiKeyModel
	Filters *MarketplaceResourceApiKeyFiltersModel `tfsdk:"filters"`
}

func (d *MarketplaceResourceApiKeyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_resource_api_key"
}

func (d *MarketplaceResourceApiKeyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Marketplace Resource Api Key data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Marketplace Resource Api Key UUID",
			},
			"filters": (&MarketplaceResourceApiKeyFiltersModel{}).GetSchema(),
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
	}
}

func (d *MarketplaceResourceApiKeyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *MarketplaceResourceApiKeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data MarketplaceResourceApiKeyDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		apiResp, err := d.client.Get(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Marketplace Resource Api Key",
				"An error occurred while reading the Marketplace Resource Api Key by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup marketplace_resource_api_key.",
			)
			return
		}

		results, err := d.client.List(ctx, filters)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Marketplace Resource Api Key",
				"An error occurred while filtering Marketplace Resource Api Key: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Marketplace Resource Api Key Not Found",
				"No Marketplace Resource Api Key found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Marketplace Resource Api Keys Found",
				fmt.Sprintf("Found %d Marketplace Resource Api Keys with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, results[0])...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
