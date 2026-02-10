package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

func BackendMetadataType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"action":        types.StringType,
		"instance_name": types.StringType,
		"runtime_state": types.StringType,
		"state":         types.StringType,
	}}
}
func CreationOrderType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"activation_price":               types.Float64Type,
		"attachment":                     types.StringType,
		"backend_id":                     types.StringType,
		"callback_url":                   types.StringType,
		"can_terminate":                  types.BoolType,
		"category_icon":                  types.StringType,
		"category_title":                 types.StringType,
		"category_uuid":                  types.StringType,
		"completed_at":                   types.StringType,
		"consumer_reviewed_at":           types.StringType,
		"consumer_reviewed_by":           types.StringType,
		"consumer_reviewed_by_full_name": types.StringType,
		"consumer_reviewed_by_username":  types.StringType,
		"cost":                           types.StringType,
		"created_by_civil_number":        types.StringType,
		"created_by_full_name":           types.StringType,
		"created_by_username":            types.StringType,
		"customer_slug":                  types.StringType,
		"error_message":                  types.StringType,
		"error_traceback":                types.StringType,
		"fixed_price":                    types.Float64Type,
		"issue":                          CreationOrderIssueType(),
		"marketplace_resource_uuid":      types.StringType,
		"new_cost_estimate":              types.StringType,
		"new_plan_name":                  types.StringType,
		"new_plan_uuid":                  types.StringType,
		"offering":                       types.StringType,
		"offering_billable":              types.BoolType,
		"offering_description":           types.StringType,
		"offering_image":                 types.StringType,
		"offering_name":                  types.StringType,
		"offering_shared":                types.BoolType,
		"offering_thumbnail":             types.StringType,
		"offering_type":                  types.StringType,
		"offering_uuid":                  types.StringType,
		"old_cost_estimate":              types.Float64Type,
		"old_plan_name":                  types.StringType,
		"old_plan_uuid":                  types.StringType,
		"order_subtype":                  types.StringType,
		"output":                         types.StringType,
		"plan":                           types.StringType,
		"plan_description":               types.StringType,
		"plan_name":                      types.StringType,
		"plan_unit":                      types.StringType,
		"plan_uuid":                      types.StringType,
		"project_description":            types.StringType,
		"project_slug":                   types.StringType,
		"provider_name":                  types.StringType,
		"provider_reviewed_at":           types.StringType,
		"provider_reviewed_by":           types.StringType,
		"provider_reviewed_by_full_name": types.StringType,
		"provider_reviewed_by_username":  types.StringType,
		"provider_slug":                  types.StringType,
		"provider_uuid":                  types.StringType,
		"request_comment":                types.StringType,
		"resource_name":                  types.StringType,
		"resource_type":                  types.StringType,
		"resource_uuid":                  types.StringType,
		"slug":                           types.StringType,
		"start_date":                     types.StringType,
		"state":                          types.StringType,
		"termination_comment":            types.StringType,
		"type":                           types.StringType,
		"url":                            types.StringType,
		"uuid":                           types.StringType,
	}}
}
func CreationOrderIssueType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"key":  types.StringType,
		"uuid": types.StringType,
	}}
}
func NestedEndpointType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"name": types.StringType,
		"url":  types.StringType,
		"uuid": types.StringType,
	}}
}
func OfferingComponentType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
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
		"uuid":                 types.StringType,
	}}
}
func ReportSectionType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"body":   types.StringType,
		"header": types.StringType,
	}}
}

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

func (m *MarketplaceResourceFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Marketplace Resource",
		Attributes: map[string]schema.Attribute{
			"backend_id": schema.StringAttribute{
				Optional: true,
			},
			"category_uuid": schema.StringAttribute{
				Optional: true,
			},
			"component_count": schema.Float64Attribute{
				Optional: true,
			},
			"created": schema.StringAttribute{
				Optional: true,
			},
			"customer": schema.StringAttribute{
				Optional: true,
			},
			"customer_uuid": schema.StringAttribute{
				Optional: true,
			},
			"downscaled": schema.BoolAttribute{
				Optional: true,
			},
			"has_terminate_date": schema.BoolAttribute{
				Optional: true,
			},
			"is_attached": schema.BoolAttribute{
				Optional: true,
			},
			"lexis_links_supported": schema.BoolAttribute{
				Optional: true,
			},
			"limit_based": schema.BoolAttribute{
				Optional: true,
			},
			"limit_component_count": schema.Float64Attribute{
				Optional: true,
			},
			"modified": schema.StringAttribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"name_exact": schema.StringAttribute{
				Optional: true,
			},
			"offering": schema.StringAttribute{
				Optional: true,
			},
			"offering_billable": schema.BoolAttribute{
				Optional: true,
			},
			"offering_shared": schema.BoolAttribute{
				Optional: true,
			},
			"offering_type": schema.StringAttribute{
				Optional: true,
			},
			"only_limit_based": schema.BoolAttribute{
				Optional: true,
			},
			"only_usage_based": schema.BoolAttribute{
				Optional: true,
			},
			"parent_offering_uuid": schema.StringAttribute{
				Optional: true,
			},
			"paused": schema.BoolAttribute{
				Optional: true,
			},
			"plan_uuid": schema.StringAttribute{
				Optional: true,
			},
			"project_name": schema.StringAttribute{
				Optional: true,
			},
			"project_uuid": schema.StringAttribute{
				Optional: true,
			},
			"provider_uuid": schema.StringAttribute{
				Optional: true,
			},
			"query": schema.StringAttribute{
				Optional: true,
			},
			"restrict_member_access": schema.BoolAttribute{
				Optional: true,
			},
			"runtime_state": schema.StringAttribute{
				Optional: true,
			},
			"service_manager_uuid": schema.StringAttribute{
				Optional: true,
			},
			"usage_based": schema.BoolAttribute{
				Optional: true,
			},
			"visible_to_providers": schema.BoolAttribute{
				Optional: true,
			},
			"visible_to_username": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

type MarketplaceResourceModel struct {
	UUID                      types.String      `tfsdk:"id"`
	AvailableActions          types.List        `tfsdk:"available_actions"`
	BackendId                 types.String      `tfsdk:"backend_id"`
	BackendMetadata           types.Object      `tfsdk:"backend_metadata"`
	CanTerminate              types.Bool        `tfsdk:"can_terminate"`
	CategoryIcon              types.String      `tfsdk:"category_icon"`
	CategoryTitle             types.String      `tfsdk:"category_title"`
	CategoryUuid              types.String      `tfsdk:"category_uuid"`
	CreationOrder             types.Object      `tfsdk:"creation_order"`
	CustomerSlug              types.String      `tfsdk:"customer_slug"`
	Description               types.String      `tfsdk:"description"`
	Downscaled                types.Bool        `tfsdk:"downscaled"`
	EffectiveId               types.String      `tfsdk:"effective_id"`
	EndDate                   types.String      `tfsdk:"end_date"`
	EndDateRequestedBy        types.String      `tfsdk:"end_date_requested_by"`
	Endpoints                 types.List        `tfsdk:"endpoints"`
	ErrorMessage              types.String      `tfsdk:"error_message"`
	ErrorTraceback            types.String      `tfsdk:"error_traceback"`
	LastSync                  timetypes.RFC3339 `tfsdk:"last_sync"`
	Name                      types.String      `tfsdk:"name"`
	Offering                  types.String      `tfsdk:"offering"`
	OfferingBackendId         types.String      `tfsdk:"offering_backend_id"`
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
	OrderInProgress           types.Object      `tfsdk:"order_in_progress"`
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
	ProjectSlug               types.String      `tfsdk:"project_slug"`
	ProviderName              types.String      `tfsdk:"provider_name"`
	ProviderSlug              types.String      `tfsdk:"provider_slug"`
	ProviderUuid              types.String      `tfsdk:"provider_uuid"`
	Report                    types.List        `tfsdk:"report"`
	ResourceType              types.String      `tfsdk:"resource_type"`
	ResourceUuid              types.String      `tfsdk:"resource_uuid"`
	RestrictMemberAccess      types.Bool        `tfsdk:"restrict_member_access"`
	Scope                     types.String      `tfsdk:"scope"`
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
	model.BackendId = common.StringPointerValue(apiResp.BackendId)
	if apiResp.BackendMetadata != nil {
		objValBackendMetadata, objDiagsBackendMetadata := types.ObjectValueFrom(ctx, BackendMetadataType().AttrTypes, *apiResp.BackendMetadata)
		diags.Append(objDiagsBackendMetadata...)
		model.BackendMetadata = objValBackendMetadata
	} else {
		model.BackendMetadata = types.ObjectNull(BackendMetadataType().AttrTypes)
	}
	model.CanTerminate = types.BoolPointerValue(apiResp.CanTerminate)
	model.CategoryIcon = common.StringPointerValue(apiResp.CategoryIcon)
	model.CategoryTitle = common.StringPointerValue(apiResp.CategoryTitle)
	model.CategoryUuid = common.StringPointerValue(apiResp.CategoryUuid)
	if apiResp.CreationOrder != nil {
		objValCreationOrder, objDiagsCreationOrder := types.ObjectValueFrom(ctx, CreationOrderType().AttrTypes, *apiResp.CreationOrder)
		diags.Append(objDiagsCreationOrder...)
		model.CreationOrder = objValCreationOrder
	} else {
		model.CreationOrder = types.ObjectNull(CreationOrderType().AttrTypes)
	}
	model.CustomerSlug = common.StringPointerValue(apiResp.CustomerSlug)
	model.Description = common.StringPointerValue(apiResp.Description)
	model.Downscaled = types.BoolPointerValue(apiResp.Downscaled)
	model.EffectiveId = common.StringPointerValue(apiResp.EffectiveId)
	model.EndDate = common.StringPointerValue(apiResp.EndDate)
	model.EndDateRequestedBy = common.StringPointerValue(apiResp.EndDateRequestedBy)

	if apiResp.Endpoints != nil && len(*apiResp.Endpoints) > 0 {
		listValEndpoints, listDiagsEndpoints := types.ListValueFrom(ctx, NestedEndpointType(), apiResp.Endpoints)
		diags.Append(listDiagsEndpoints...)
		model.Endpoints = listValEndpoints
	} else {
		model.Endpoints = types.ListNull(NestedEndpointType())
	}
	model.ErrorMessage = common.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = common.StringPointerValue(apiResp.ErrorTraceback)
	valLastSync, diagsLastSync := timetypes.NewRFC3339PointerValue(apiResp.LastSync)
	diags.Append(diagsLastSync...)
	model.LastSync = valLastSync
	model.Name = common.StringPointerValue(apiResp.Name)
	model.Offering = common.StringPointerValue(apiResp.Offering)
	model.OfferingBackendId = common.StringPointerValue(apiResp.OfferingBackendId)
	model.OfferingBillable = types.BoolPointerValue(apiResp.OfferingBillable)

	if apiResp.OfferingComponents != nil && len(*apiResp.OfferingComponents) > 0 {
		listValOfferingComponents, listDiagsOfferingComponents := types.ListValueFrom(ctx, OfferingComponentType(), apiResp.OfferingComponents)
		diags.Append(listDiagsOfferingComponents...)
		model.OfferingComponents = listValOfferingComponents
	} else {
		model.OfferingComponents = types.ListNull(OfferingComponentType())
	}
	model.OfferingDescription = common.StringPointerValue(apiResp.OfferingDescription)
	model.OfferingImage = common.StringPointerValue(apiResp.OfferingImage)
	model.OfferingName = common.StringPointerValue(apiResp.OfferingName)
	model.OfferingShared = types.BoolPointerValue(apiResp.OfferingShared)
	model.OfferingSlug = common.StringPointerValue(apiResp.OfferingSlug)
	model.OfferingState = common.StringPointerValue(apiResp.OfferingState)
	model.OfferingThumbnail = common.StringPointerValue(apiResp.OfferingThumbnail)
	model.OfferingType = common.StringPointerValue(apiResp.OfferingType)
	model.OfferingUuid = common.StringPointerValue(apiResp.OfferingUuid)
	if apiResp.OrderInProgress != nil {
		objValOrderInProgress, objDiagsOrderInProgress := types.ObjectValueFrom(ctx, CreationOrderType().AttrTypes, *apiResp.OrderInProgress)
		diags.Append(objDiagsOrderInProgress...)
		model.OrderInProgress = objValOrderInProgress
	} else {
		model.OrderInProgress = types.ObjectNull(CreationOrderType().AttrTypes)
	}
	model.ParentName = common.StringPointerValue(apiResp.ParentName)
	model.ParentOfferingName = common.StringPointerValue(apiResp.ParentOfferingName)
	model.ParentOfferingSlug = common.StringPointerValue(apiResp.ParentOfferingSlug)
	model.ParentOfferingUuid = common.StringPointerValue(apiResp.ParentOfferingUuid)
	model.ParentUuid = common.StringPointerValue(apiResp.ParentUuid)
	model.Paused = types.BoolPointerValue(apiResp.Paused)
	model.Plan = common.StringPointerValue(apiResp.Plan)
	model.PlanDescription = common.StringPointerValue(apiResp.PlanDescription)
	model.PlanName = common.StringPointerValue(apiResp.PlanName)
	model.PlanUnit = common.StringPointerValue(apiResp.PlanUnit)
	model.PlanUuid = common.StringPointerValue(apiResp.PlanUuid)
	model.Project = common.StringPointerValue(apiResp.Project)
	model.ProjectDescription = common.StringPointerValue(apiResp.ProjectDescription)
	model.ProjectEndDate = common.StringPointerValue(apiResp.ProjectEndDate)
	model.ProjectEndDateRequestedBy = common.StringPointerValue(apiResp.ProjectEndDateRequestedBy)
	model.ProjectSlug = common.StringPointerValue(apiResp.ProjectSlug)
	model.ProviderName = common.StringPointerValue(apiResp.ProviderName)
	model.ProviderSlug = common.StringPointerValue(apiResp.ProviderSlug)
	model.ProviderUuid = common.StringPointerValue(apiResp.ProviderUuid)

	if apiResp.Report != nil && len(*apiResp.Report) > 0 {
		listValReport, listDiagsReport := types.ListValueFrom(ctx, ReportSectionType(), apiResp.Report)
		diags.Append(listDiagsReport...)
		model.Report = listValReport
	} else {
		model.Report = types.ListNull(ReportSectionType())
	}
	model.ResourceType = common.StringPointerValue(apiResp.ResourceType)
	model.ResourceUuid = common.StringPointerValue(apiResp.ResourceUuid)
	model.RestrictMemberAccess = types.BoolPointerValue(apiResp.RestrictMemberAccess)
	model.Scope = common.StringPointerValue(apiResp.Scope)
	model.Slug = common.StringPointerValue(apiResp.Slug)
	model.State = common.StringPointerValue(apiResp.State)
	model.Url = common.StringPointerValue(apiResp.Url)
	model.UserRequiresReconsent = types.BoolPointerValue(apiResp.UserRequiresReconsent)
	model.Username = common.StringPointerValue(apiResp.Username)

	return diags
}
