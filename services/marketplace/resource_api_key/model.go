package resource_api_key

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type MarketplaceResourceApiKeyFiltersModel struct {
	CustomerUuid     types.String `tfsdk:"customer_uuid"`
	HasPendingAction types.Bool   `tfsdk:"has_pending_action"`
	ModifiedBefore   types.String `tfsdk:"modified_before"`
	OfferingUuid     types.String `tfsdk:"offering_uuid"`
	PausedByLimit    types.Bool   `tfsdk:"paused_by_limit"`
	PendingAction    types.String `tfsdk:"pending_action"`
	ProjectUuid      types.String `tfsdk:"project_uuid"`
	ProviderUuid     types.String `tfsdk:"provider_uuid"`
	ResourceUuid     types.String `tfsdk:"resource_uuid"`
	State            types.String `tfsdk:"state"`
	UserUuid         types.String `tfsdk:"user_uuid"`
}

func (m *MarketplaceResourceApiKeyFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Marketplace Resource Api Key",
		Attributes: map[string]schema.Attribute{
			"customer_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Consumer organization UUID",
			},
			"has_pending_action": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Has a pending command",
			},
			"modified_before": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Modified before",
			},
			"offering_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Offering UUID",
			},
			"paused_by_limit": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Paused by a spending limit",
			},
			"pending_action": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Pending command",

				Validators: []validator.String{
					stringvalidator.OneOf("create", "rotate", "pause", "resume", "delete", "update"),
				},
			},
			"project_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project UUID",
			},
			"provider_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Keys of resources of the offerings of this service provider organization, to callers the provider resource list shows those resources to.",
			},
			"resource_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Resource UUID",
			},
			"state": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "API key state",

				Validators: []validator.String{
					stringvalidator.OneOf("Creating", "OK", "Updating", "Paused", "Deleting", "Deleted", "Erred"),
				},
			},
			"user_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Assignee UUID",
			},
		},
	}
}

type MarketplaceResourceApiKeyModel struct {
	UUID                   types.String      `tfsdk:"id"`
	AllowedModels          types.List        `tfsdk:"allowed_models"`
	ClientId               types.String      `tfsdk:"client_id"`
	Description            types.String      `tfsdk:"description"`
	EffectiveModels        types.List        `tfsdk:"effective_models"`
	ErrorMessage           types.String      `tfsdk:"error_message"`
	EstimatedCost          types.String      `tfsdk:"estimated_cost"`
	IssuedAt               timetypes.RFC3339 `tfsdk:"issued_at"`
	LimitResumesAt         timetypes.RFC3339 `tfsdk:"limit_resumes_at"`
	Limits                 types.Map         `tfsdk:"limits"`
	OfferingUuid           types.String      `tfsdk:"offering_uuid"`
	PausedByLimit          types.Bool        `tfsdk:"paused_by_limit"`
	PendingAction          types.String      `tfsdk:"pending_action"`
	ResourceBackendId      types.String      `tfsdk:"resource_backend_id"`
	ResourceName           types.String      `tfsdk:"resource_name"`
	ResourcePlanUuid       types.String      `tfsdk:"resource_plan_uuid"`
	ResourceSpendingLimits types.Map         `tfsdk:"resource_spending_limits"`
	ResourceUuid           types.String      `tfsdk:"resource_uuid"`
	SpendingLimitStatus    types.List        `tfsdk:"spending_limit_status"`
	SpendingLimits         types.Map         `tfsdk:"spending_limits"`
	State                  types.String      `tfsdk:"state"`
	UsagePeriod            types.String      `tfsdk:"usage_period"`
	UserFullName           types.String      `tfsdk:"user_full_name"`
	UserUuid               types.String      `tfsdk:"user_uuid"`
}

// CopyFrom maps the API response to the model fields.
func (model *MarketplaceResourceApiKeyModel) CopyFrom(ctx context.Context, apiResp MarketplaceResourceApiKeyResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)

	if apiResp.AllowedModels != nil {
		valAllowedModels, diagsAllowedModels := types.ListValueFrom(ctx, types.StringType, apiResp.AllowedModels)
		diags.Append(diagsAllowedModels...)
		model.AllowedModels = valAllowedModels
	} else {
		model.AllowedModels = types.ListNull(types.StringType)
	}

	model.ClientId = common.StringPointerValue(apiResp.ClientId)

	model.Description = common.StringPointerValue(apiResp.Description)

	if apiResp.EffectiveModels != nil {
		valEffectiveModels, diagsEffectiveModels := types.ListValueFrom(ctx, types.StringType, apiResp.EffectiveModels)
		diags.Append(diagsEffectiveModels...)
		model.EffectiveModels = valEffectiveModels
	} else {
		model.EffectiveModels = types.ListNull(types.StringType)
	}

	model.ErrorMessage = common.StringPointerValue(apiResp.ErrorMessage)

	model.EstimatedCost = common.StringPointerValue(apiResp.EstimatedCost)

	valIssuedAt, diagsIssuedAt := timetypes.NewRFC3339PointerValue(apiResp.IssuedAt)
	diags.Append(diagsIssuedAt...)
	model.IssuedAt = valIssuedAt

	valLimitResumesAt, diagsLimitResumesAt := timetypes.NewRFC3339PointerValue(apiResp.LimitResumesAt)
	diags.Append(diagsLimitResumesAt...)
	model.LimitResumesAt = valLimitResumesAt

	if apiResp.Limits != nil {
		valLimits, diagsLimits := types.MapValueFrom(ctx, types.Float64Type, apiResp.Limits)
		diags.Append(diagsLimits...)
		model.Limits = valLimits
	} else {
		model.Limits = types.MapNull(types.Float64Type)
	}

	model.OfferingUuid = common.StringPointerValue(apiResp.OfferingUuid)

	model.PausedByLimit = types.BoolPointerValue(apiResp.PausedByLimit)

	model.PendingAction = common.StringPointerValue(apiResp.PendingAction)

	model.ResourceBackendId = common.StringPointerValue(apiResp.ResourceBackendId)

	model.ResourceName = common.StringPointerValue(apiResp.ResourceName)

	model.ResourcePlanUuid = common.StringPointerValue(apiResp.ResourcePlanUuid)

	if apiResp.ResourceSpendingLimits != nil {
		valResourceSpendingLimits, diagsResourceSpendingLimits := types.MapValueFrom(ctx, types.StringType, apiResp.ResourceSpendingLimits)
		diags.Append(diagsResourceSpendingLimits...)
		model.ResourceSpendingLimits = valResourceSpendingLimits
	} else {
		model.ResourceSpendingLimits = types.MapNull(types.StringType)
	}

	model.ResourceUuid = common.StringPointerValue(apiResp.ResourceUuid)

	if apiResp.SpendingLimitStatus != nil {
		valSpendingLimitStatus, diagsSpendingLimitStatus := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"component_type": types.StringType,
			"limit":          types.Float64Type,
			"period":         types.StringType,
			"reached":        types.BoolType,
			"resets_at":      types.StringType,
			"used":           types.Float64Type,
		}}, apiResp.SpendingLimitStatus)
		diags.Append(diagsSpendingLimitStatus...)
		model.SpendingLimitStatus = valSpendingLimitStatus
	} else {
		model.SpendingLimitStatus = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
			"component_type": types.StringType,
			"limit":          types.Float64Type,
			"period":         types.StringType,
			"reached":        types.BoolType,
			"resets_at":      types.StringType,
			"used":           types.Float64Type,
		}})
	}

	if apiResp.SpendingLimits != nil {
		valSpendingLimits, diagsSpendingLimits := types.MapValueFrom(ctx, types.StringType, apiResp.SpendingLimits)
		diags.Append(diagsSpendingLimits...)
		model.SpendingLimits = valSpendingLimits
	} else {
		model.SpendingLimits = types.MapNull(types.StringType)
	}

	model.State = common.StringPointerValue(apiResp.State)

	model.UsagePeriod = common.StringPointerValue(apiResp.UsagePeriod)

	model.UserFullName = common.StringPointerValue(apiResp.UserFullName)

	model.UserUuid = common.StringPointerValue(apiResp.UserUuid)

	return diags
}
