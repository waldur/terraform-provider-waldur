package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// MarketplaceResourceFiltersModel contains the filter parameters for querying.
type MarketplaceResourceFiltersModel struct {
	BackendId            types.String  `tfsdk:"backend_id"`
	CategoryUuid         types.String  `tfsdk:"category_uuid"`
	ComponentCount       types.Float64 `tfsdk:"component_count"`
	Created              types.String  `tfsdk:"created"`
	Customer             types.String  `tfsdk:"customer"`
	CustomerUuid         types.String  `tfsdk:"customer_uuid"`
	Downscaled           types.Bool    `tfsdk:"downscaled"`
	HasTerminateDate     types.Bool    `tfsdk:"has_terminate_date"`
	IsAttached           types.Bool    `tfsdk:"is_attached"`
	LexisLinksSupported  types.Bool    `tfsdk:"lexis_links_supported"`
	LimitBased           types.Bool    `tfsdk:"limit_based"`
	LimitComponentCount  types.Float64 `tfsdk:"limit_component_count"`
	Modified             types.String  `tfsdk:"modified"`
	Name                 types.String  `tfsdk:"name"`
	NameExact            types.String  `tfsdk:"name_exact"`
	Offering             types.String  `tfsdk:"offering"`
	OfferingBillable     types.Bool    `tfsdk:"offering_billable"`
	OfferingShared       types.Bool    `tfsdk:"offering_shared"`
	OfferingType         types.String  `tfsdk:"offering_type"`
	OnlyLimitBased       types.Bool    `tfsdk:"only_limit_based"`
	OnlyUsageBased       types.Bool    `tfsdk:"only_usage_based"`
	ParentOfferingUuid   types.String  `tfsdk:"parent_offering_uuid"`
	Paused               types.Bool    `tfsdk:"paused"`
	PlanUuid             types.String  `tfsdk:"plan_uuid"`
	ProjectName          types.String  `tfsdk:"project_name"`
	ProjectUuid          types.String  `tfsdk:"project_uuid"`
	ProviderUuid         types.String  `tfsdk:"provider_uuid"`
	Query                types.String  `tfsdk:"query"`
	RestrictMemberAccess types.Bool    `tfsdk:"restrict_member_access"`
	RuntimeState         types.String  `tfsdk:"runtime_state"`
	ServiceManagerUuid   types.String  `tfsdk:"service_manager_uuid"`
	UsageBased           types.Bool    `tfsdk:"usage_based"`
	VisibleToProviders   types.Bool    `tfsdk:"visible_to_providers"`
	VisibleToUsername    types.String  `tfsdk:"visible_to_username"`
}

type MarketplaceResourceModel struct {
	UUID                      types.String      `tfsdk:"id"`
	AvailableActions          types.List        `tfsdk:"available_actions"`
	BackendId                 types.String      `tfsdk:"backend_id"`
	CanTerminate              types.Bool        `tfsdk:"can_terminate"`
	CategoryIcon              types.String      `tfsdk:"category_icon"`
	CategoryTitle             types.String      `tfsdk:"category_title"`
	CategoryUuid              types.String      `tfsdk:"category_uuid"`
	Created                   timetypes.RFC3339 `tfsdk:"created"`
	CustomerName              types.String      `tfsdk:"customer_name"`
	CustomerSlug              types.String      `tfsdk:"customer_slug"`
	CustomerUuid              types.String      `tfsdk:"customer_uuid"`
	Description               types.String      `tfsdk:"description"`
	Downscaled                types.Bool        `tfsdk:"downscaled"`
	EffectiveId               types.String      `tfsdk:"effective_id"`
	EndDate                   types.String      `tfsdk:"end_date"`
	EndDateRequestedBy        types.String      `tfsdk:"end_date_requested_by"`
	Endpoints                 types.List        `tfsdk:"endpoints"`
	ErrorMessage              types.String      `tfsdk:"error_message"`
	ErrorTraceback            types.String      `tfsdk:"error_traceback"`
	IsLimitBased              types.Bool        `tfsdk:"is_limit_based"`
	IsUsageBased              types.Bool        `tfsdk:"is_usage_based"`
	LastSync                  timetypes.RFC3339 `tfsdk:"last_sync"`
	Modified                  timetypes.RFC3339 `tfsdk:"modified"`
	Name                      types.String      `tfsdk:"name"`
	Offering                  types.String      `tfsdk:"offering"`
	OfferingBillable          types.Bool        `tfsdk:"offering_billable"`
	OfferingComponents        types.List        `tfsdk:"offering_components"`
	OfferingDescription       types.String      `tfsdk:"offering_description"`
	OfferingImage             types.String      `tfsdk:"offering_image"`
	OfferingName              types.String      `tfsdk:"offering_name"`
	OfferingShared            types.Bool        `tfsdk:"offering_shared"`
	OfferingSlug              types.String      `tfsdk:"offering_slug"`
	OfferingState             types.String      `tfsdk:"offering_state"`
	OfferingThumbnail         types.String      `tfsdk:"offering_thumbnail"`
	OfferingType              types.String      `tfsdk:"offering_type"`
	OfferingUuid              types.String      `tfsdk:"offering_uuid"`
	ParentName                types.String      `tfsdk:"parent_name"`
	ParentOfferingName        types.String      `tfsdk:"parent_offering_name"`
	ParentOfferingSlug        types.String      `tfsdk:"parent_offering_slug"`
	ParentOfferingUuid        types.String      `tfsdk:"parent_offering_uuid"`
	ParentUuid                types.String      `tfsdk:"parent_uuid"`
	Paused                    types.Bool        `tfsdk:"paused"`
	Plan                      types.String      `tfsdk:"plan"`
	PlanDescription           types.String      `tfsdk:"plan_description"`
	PlanName                  types.String      `tfsdk:"plan_name"`
	PlanUnit                  types.String      `tfsdk:"plan_unit"`
	PlanUuid                  types.String      `tfsdk:"plan_uuid"`
	Project                   types.String      `tfsdk:"project"`
	ProjectDescription        types.String      `tfsdk:"project_description"`
	ProjectEndDate            types.String      `tfsdk:"project_end_date"`
	ProjectEndDateRequestedBy types.String      `tfsdk:"project_end_date_requested_by"`
	ProjectName               types.String      `tfsdk:"project_name"`
	ProjectSlug               types.String      `tfsdk:"project_slug"`
	ProjectUuid               types.String      `tfsdk:"project_uuid"`
	ProviderName              types.String      `tfsdk:"provider_name"`
	ProviderSlug              types.String      `tfsdk:"provider_slug"`
	ProviderUuid              types.String      `tfsdk:"provider_uuid"`
	Report                    types.List        `tfsdk:"report"`
	ResourceType              types.String      `tfsdk:"resource_type"`
	ResourceUuid              types.String      `tfsdk:"resource_uuid"`
	RestrictMemberAccess      types.Bool        `tfsdk:"restrict_member_access"`
	Scope                     types.String      `tfsdk:"scope"`
	ServiceSettingsUuid       types.String      `tfsdk:"service_settings_uuid"`
	Slug                      types.String      `tfsdk:"slug"`
	State                     types.String      `tfsdk:"state"`
	Url                       types.String      `tfsdk:"url"`
	UserRequiresReconsent     types.Bool        `tfsdk:"user_requires_reconsent"`
	Username                  types.String      `tfsdk:"username"`
}

// CopyFrom maps the API response to the model fields.
func (model *MarketplaceResourceModel) CopyFrom(ctx context.Context, apiResp MarketplaceResourceResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	listValAvailableActions, listDiagsAvailableActions := types.ListValueFrom(ctx, types.StringType, apiResp.AvailableActions)
	model.AvailableActions = listValAvailableActions
	diags.Append(listDiagsAvailableActions...)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.CanTerminate = types.BoolPointerValue(apiResp.CanTerminate)
	model.CategoryIcon = types.StringPointerValue(apiResp.CategoryIcon)
	model.CategoryTitle = types.StringPointerValue(apiResp.CategoryTitle)
	model.CategoryUuid = types.StringPointerValue(apiResp.CategoryUuid)
	valCreated, diagsCreated := timetypes.NewRFC3339PointerValue(apiResp.Created)
	diags.Append(diagsCreated...)
	model.Created = valCreated
	model.CustomerSlug = types.StringPointerValue(apiResp.CustomerSlug)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.Downscaled = types.BoolPointerValue(apiResp.Downscaled)
	model.EffectiveId = types.StringPointerValue(apiResp.EffectiveId)
	model.EndDate = types.StringPointerValue(apiResp.EndDate)
	model.EndDateRequestedBy = types.StringPointerValue(apiResp.EndDateRequestedBy)

	{
		listValEndpoints, listDiagsEndpoints := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"url":  types.StringType,
		}}, apiResp.Endpoints)
		diags.Append(listDiagsEndpoints...)
		model.Endpoints = listValEndpoints
	}
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	valLastSync, diagsLastSync := timetypes.NewRFC3339PointerValue(apiResp.LastSync)
	diags.Append(diagsLastSync...)
	model.LastSync = valLastSync
	valModified, diagsModified := timetypes.NewRFC3339PointerValue(apiResp.Modified)
	diags.Append(diagsModified...)
	model.Modified = valModified
	model.Name = types.StringPointerValue(apiResp.Name)
	model.Offering = types.StringPointerValue(apiResp.Offering)
	model.OfferingBillable = types.BoolPointerValue(apiResp.OfferingBillable)

	{
		listValOfferingComponents, listDiagsOfferingComponents := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"article_code":         types.StringType,
			"billing_type":         types.StringType,
			"default_limit":        types.Int64Type,
			"description":          types.StringType,
			"factor":               types.Int64Type,
			"is_boolean":           types.BoolType,
			"is_builtin":           types.BoolType,
			"is_prepaid":           types.BoolType,
			"limit_amount":         types.Int64Type,
			"limit_period":         types.StringType,
			"max_available_limit":  types.Int64Type,
			"max_prepaid_duration": types.Int64Type,
			"max_value":            types.Int64Type,
			"measured_unit":        types.StringType,
			"min_prepaid_duration": types.Int64Type,
			"min_value":            types.Int64Type,
			"name":                 types.StringType,
			"overage_component":    types.StringType,
			"type":                 types.StringType,
			"unit_factor":          types.Int64Type,
		}}, apiResp.OfferingComponents)
		diags.Append(listDiagsOfferingComponents...)
		model.OfferingComponents = listValOfferingComponents
	}
	model.OfferingDescription = types.StringPointerValue(apiResp.OfferingDescription)
	model.OfferingImage = types.StringPointerValue(apiResp.OfferingImage)
	model.OfferingName = types.StringPointerValue(apiResp.OfferingName)
	model.OfferingShared = types.BoolPointerValue(apiResp.OfferingShared)
	model.OfferingSlug = types.StringPointerValue(apiResp.OfferingSlug)
	model.OfferingState = types.StringPointerValue(apiResp.OfferingState)
	model.OfferingThumbnail = types.StringPointerValue(apiResp.OfferingThumbnail)
	model.OfferingType = types.StringPointerValue(apiResp.OfferingType)
	model.OfferingUuid = types.StringPointerValue(apiResp.OfferingUuid)
	model.ParentName = types.StringPointerValue(apiResp.ParentName)
	model.ParentOfferingName = types.StringPointerValue(apiResp.ParentOfferingName)
	model.ParentOfferingSlug = types.StringPointerValue(apiResp.ParentOfferingSlug)
	model.ParentOfferingUuid = types.StringPointerValue(apiResp.ParentOfferingUuid)
	model.ParentUuid = types.StringPointerValue(apiResp.ParentUuid)
	model.Paused = types.BoolPointerValue(apiResp.Paused)
	model.Plan = types.StringPointerValue(apiResp.Plan)
	model.PlanDescription = types.StringPointerValue(apiResp.PlanDescription)
	model.PlanName = types.StringPointerValue(apiResp.PlanName)
	model.PlanUnit = types.StringPointerValue(apiResp.PlanUnit)
	model.PlanUuid = types.StringPointerValue(apiResp.PlanUuid)
	model.ProjectDescription = types.StringPointerValue(apiResp.ProjectDescription)
	model.ProjectEndDate = types.StringPointerValue(apiResp.ProjectEndDate)
	model.ProjectEndDateRequestedBy = types.StringPointerValue(apiResp.ProjectEndDateRequestedBy)
	model.ProjectSlug = types.StringPointerValue(apiResp.ProjectSlug)
	model.ProviderName = types.StringPointerValue(apiResp.ProviderName)
	model.ProviderSlug = types.StringPointerValue(apiResp.ProviderSlug)
	model.ProviderUuid = types.StringPointerValue(apiResp.ProviderUuid)

	{
		listValReport, listDiagsReport := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"body":   types.StringType,
			"header": types.StringType,
		}}, apiResp.Report)
		diags.Append(listDiagsReport...)
		model.Report = listValReport
	}
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.ResourceUuid = types.StringPointerValue(apiResp.ResourceUuid)
	model.RestrictMemberAccess = types.BoolPointerValue(apiResp.RestrictMemberAccess)
	model.Scope = types.StringPointerValue(apiResp.Scope)
	model.Slug = types.StringPointerValue(apiResp.Slug)
	model.State = types.StringPointerValue(apiResp.State)
	model.Url = types.StringPointerValue(apiResp.Url)
	model.UserRequiresReconsent = types.BoolPointerValue(apiResp.UserRequiresReconsent)
	model.Username = types.StringPointerValue(apiResp.Username)

	return diags
}
