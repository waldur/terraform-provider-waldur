package resource_api_key_ephemeral

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ ephemeral.EphemeralResource = &MarketplaceResourceApiKeyEphemeralResource{}
var _ ephemeral.EphemeralResourceWithConfigure = &MarketplaceResourceApiKeyEphemeralResource{}

func NewMarketplaceResourceApiKeyEphemeralResource() ephemeral.EphemeralResource {
	return &MarketplaceResourceApiKeyEphemeralResource{}
}

// MarketplaceResourceApiKeyEphemeralResource calls the API each time Terraform opens
// it. Its result lives for one run only: Terraform does not write it to the
// plan or the state.
type MarketplaceResourceApiKeyEphemeralResource struct {
	client *MarketplaceResourceApiKeyClient
}

type MarketplaceResourceApiKeyEphemeralModel struct {
	MarketplaceResourceApiKeyModel
}

func (e *MarketplaceResourceApiKeyEphemeralResource) Metadata(ctx context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_resource_api_key"
}

func (e *MarketplaceResourceApiKeyEphemeralResource) Schema(ctx context.Context, req ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Marketplace Resource Api Key ephemeral resource. Its values are fetched on every run and never stored in the Terraform plan or state.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Marketplace Resource Api Key UUID",
			},
			"allowed_models": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true, MarkdownDescription: "The key's own model list, as chosen."},
			"api_key": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Api Key",
				Sensitive: true},
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

func (e *MarketplaceResourceApiKeyEphemeralResource) Configure(ctx context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	e.client = &MarketplaceResourceApiKeyClient{}
	if err := e.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Ephemeral Resource Configure Type",
			err.Error(),
		)
		return
	}
}

func (e *MarketplaceResourceApiKeyEphemeralResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var data MarketplaceResourceApiKeyEphemeralModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	configured := data

	apiResp, err := e.client.Open(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Open Marketplace Resource Api Key",
			"An error occurred while reading the Marketplace Resource Api Key: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	// The arguments stay as configured, whatever the response carries.
	data.UUID = configured.UUID

	resp.Diagnostics.Append(resp.Result.Set(ctx, &data)...)
}
