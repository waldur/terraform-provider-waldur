package order

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// MarketplaceOrderFiltersModel contains the filter parameters for querying.
type MarketplaceOrderFiltersModel struct {
	CanApproveAsConsumer types.Bool   `tfsdk:"can_approve_as_consumer"`
	CanApproveAsProvider types.Bool   `tfsdk:"can_approve_as_provider"`
	CategoryUuid         types.String `tfsdk:"category_uuid"`
	Created              types.String `tfsdk:"created"`
	CustomerUuid         types.String `tfsdk:"customer_uuid"`
	Modified             types.String `tfsdk:"modified"`
	Offering             types.String `tfsdk:"offering"`
	OfferingUuid         types.String `tfsdk:"offering_uuid"`
	ParentOfferingUuid   types.String `tfsdk:"parent_offering_uuid"`
	ProjectUuid          types.String `tfsdk:"project_uuid"`
	ProviderUuid         types.String `tfsdk:"provider_uuid"`
	Query                types.String `tfsdk:"query"`
	Resource             types.String `tfsdk:"resource"`
	ResourceName         types.String `tfsdk:"resource_name"`
	ResourceUuid         types.String `tfsdk:"resource_uuid"`
	ServiceManagerUuid   types.String `tfsdk:"service_manager_uuid"`
}

type MarketplaceOrderModel struct {
	UUID                       types.String      `tfsdk:"id"`
	ActivationPrice            types.Float64     `tfsdk:"activation_price"`
	Attachment                 types.String      `tfsdk:"attachment"`
	BackendId                  types.String      `tfsdk:"backend_id"`
	CallbackUrl                types.String      `tfsdk:"callback_url"`
	CanTerminate               types.Bool        `tfsdk:"can_terminate"`
	CategoryIcon               types.String      `tfsdk:"category_icon"`
	CategoryTitle              types.String      `tfsdk:"category_title"`
	CategoryUuid               types.String      `tfsdk:"category_uuid"`
	CompletedAt                timetypes.RFC3339 `tfsdk:"completed_at"`
	ConsumerReviewedAt         timetypes.RFC3339 `tfsdk:"consumer_reviewed_at"`
	ConsumerReviewedBy         types.String      `tfsdk:"consumer_reviewed_by"`
	ConsumerReviewedByFullName types.String      `tfsdk:"consumer_reviewed_by_full_name"`
	ConsumerReviewedByUsername types.String      `tfsdk:"consumer_reviewed_by_username"`
	Cost                       types.String      `tfsdk:"cost"`
	Created                    timetypes.RFC3339 `tfsdk:"created"`
	CreatedByCivilNumber       types.String      `tfsdk:"created_by_civil_number"`
	CreatedByFullName          types.String      `tfsdk:"created_by_full_name"`
	CreatedByUsername          types.String      `tfsdk:"created_by_username"`
	CustomerName               types.String      `tfsdk:"customer_name"`
	CustomerSlug               types.String      `tfsdk:"customer_slug"`
	CustomerUuid               types.String      `tfsdk:"customer_uuid"`
	ErrorMessage               types.String      `tfsdk:"error_message"`
	ErrorTraceback             types.String      `tfsdk:"error_traceback"`
	FixedPrice                 types.Float64     `tfsdk:"fixed_price"`
	MarketplaceResourceUuid    types.String      `tfsdk:"marketplace_resource_uuid"`
	Modified                   timetypes.RFC3339 `tfsdk:"modified"`
	NewCostEstimate            types.String      `tfsdk:"new_cost_estimate"`
	NewPlanName                types.String      `tfsdk:"new_plan_name"`
	NewPlanUuid                types.String      `tfsdk:"new_plan_uuid"`
	Offering                   types.String      `tfsdk:"offering"`
	OfferingBillable           types.Bool        `tfsdk:"offering_billable"`
	OfferingDescription        types.String      `tfsdk:"offering_description"`
	OfferingImage              types.String      `tfsdk:"offering_image"`
	OfferingName               types.String      `tfsdk:"offering_name"`
	OfferingShared             types.Bool        `tfsdk:"offering_shared"`
	OfferingThumbnail          types.String      `tfsdk:"offering_thumbnail"`
	OfferingType               types.String      `tfsdk:"offering_type"`
	OfferingUuid               types.String      `tfsdk:"offering_uuid"`
	OldCostEstimate            types.Float64     `tfsdk:"old_cost_estimate"`
	OldPlanName                types.String      `tfsdk:"old_plan_name"`
	OldPlanUuid                types.String      `tfsdk:"old_plan_uuid"`
	OrderSubtype               types.String      `tfsdk:"order_subtype"`
	Output                     types.String      `tfsdk:"output"`
	Plan                       types.String      `tfsdk:"plan"`
	PlanDescription            types.String      `tfsdk:"plan_description"`
	PlanName                   types.String      `tfsdk:"plan_name"`
	PlanUnit                   types.String      `tfsdk:"plan_unit"`
	PlanUuid                   types.String      `tfsdk:"plan_uuid"`
	ProjectDescription         types.String      `tfsdk:"project_description"`
	ProjectName                types.String      `tfsdk:"project_name"`
	ProjectSlug                types.String      `tfsdk:"project_slug"`
	ProjectUuid                types.String      `tfsdk:"project_uuid"`
	ProviderName               types.String      `tfsdk:"provider_name"`
	ProviderReviewedAt         timetypes.RFC3339 `tfsdk:"provider_reviewed_at"`
	ProviderReviewedBy         types.String      `tfsdk:"provider_reviewed_by"`
	ProviderReviewedByFullName types.String      `tfsdk:"provider_reviewed_by_full_name"`
	ProviderReviewedByUsername types.String      `tfsdk:"provider_reviewed_by_username"`
	ProviderSlug               types.String      `tfsdk:"provider_slug"`
	ProviderUuid               types.String      `tfsdk:"provider_uuid"`
	RequestComment             types.String      `tfsdk:"request_comment"`
	ResourceName               types.String      `tfsdk:"resource_name"`
	ResourceType               types.String      `tfsdk:"resource_type"`
	ResourceUuid               types.String      `tfsdk:"resource_uuid"`
	Slug                       types.String      `tfsdk:"slug"`
	StartDate                  types.String      `tfsdk:"start_date"`
	State                      types.String      `tfsdk:"state"`
	TerminationComment         types.String      `tfsdk:"termination_comment"`
	Type                       types.String      `tfsdk:"type"`
	Url                        types.String      `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *MarketplaceOrderModel) CopyFrom(ctx context.Context, apiResp MarketplaceOrderResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.ActivationPrice = types.Float64PointerValue(apiResp.ActivationPrice)
	model.Attachment = types.StringPointerValue(apiResp.Attachment)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.CallbackUrl = types.StringPointerValue(apiResp.CallbackUrl)
	model.CanTerminate = types.BoolPointerValue(apiResp.CanTerminate)
	model.CategoryIcon = types.StringPointerValue(apiResp.CategoryIcon)
	model.CategoryTitle = types.StringPointerValue(apiResp.CategoryTitle)
	model.CategoryUuid = types.StringPointerValue(apiResp.CategoryUuid)
	valCompletedAt, diagsCompletedAt := timetypes.NewRFC3339PointerValue(apiResp.CompletedAt)
	diags.Append(diagsCompletedAt...)
	model.CompletedAt = valCompletedAt
	valConsumerReviewedAt, diagsConsumerReviewedAt := timetypes.NewRFC3339PointerValue(apiResp.ConsumerReviewedAt)
	diags.Append(diagsConsumerReviewedAt...)
	model.ConsumerReviewedAt = valConsumerReviewedAt
	model.ConsumerReviewedBy = types.StringPointerValue(apiResp.ConsumerReviewedBy)
	model.ConsumerReviewedByFullName = types.StringPointerValue(apiResp.ConsumerReviewedByFullName)
	model.ConsumerReviewedByUsername = types.StringPointerValue(apiResp.ConsumerReviewedByUsername)
	model.Cost = types.StringPointerValue(apiResp.Cost)
	valCreated, diagsCreated := timetypes.NewRFC3339PointerValue(apiResp.Created)
	diags.Append(diagsCreated...)
	model.Created = valCreated
	model.CreatedByCivilNumber = types.StringPointerValue(apiResp.CreatedByCivilNumber)
	model.CreatedByFullName = types.StringPointerValue(apiResp.CreatedByFullName)
	model.CreatedByUsername = types.StringPointerValue(apiResp.CreatedByUsername)
	model.CustomerSlug = types.StringPointerValue(apiResp.CustomerSlug)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.FixedPrice = types.Float64PointerValue(apiResp.FixedPrice)
	valModified, diagsModified := timetypes.NewRFC3339PointerValue(apiResp.Modified)
	diags.Append(diagsModified...)
	model.Modified = valModified
	model.NewCostEstimate = types.StringPointerValue(apiResp.NewCostEstimate)
	model.NewPlanName = types.StringPointerValue(apiResp.NewPlanName)
	model.NewPlanUuid = types.StringPointerValue(apiResp.NewPlanUuid)
	model.Offering = types.StringPointerValue(apiResp.Offering)
	model.OfferingBillable = types.BoolPointerValue(apiResp.OfferingBillable)
	model.OfferingDescription = types.StringPointerValue(apiResp.OfferingDescription)
	model.OfferingImage = types.StringPointerValue(apiResp.OfferingImage)
	model.OfferingName = types.StringPointerValue(apiResp.OfferingName)
	model.OfferingShared = types.BoolPointerValue(apiResp.OfferingShared)
	model.OfferingThumbnail = types.StringPointerValue(apiResp.OfferingThumbnail)
	model.OfferingType = types.StringPointerValue(apiResp.OfferingType)
	model.OfferingUuid = types.StringPointerValue(apiResp.OfferingUuid)
	model.OldCostEstimate = types.Float64PointerValue(apiResp.OldCostEstimate)
	model.OldPlanName = types.StringPointerValue(apiResp.OldPlanName)
	model.OldPlanUuid = types.StringPointerValue(apiResp.OldPlanUuid)
	model.OrderSubtype = types.StringPointerValue(apiResp.OrderSubtype)
	model.Output = types.StringPointerValue(apiResp.Output)
	model.Plan = types.StringPointerValue(apiResp.Plan)
	model.PlanDescription = types.StringPointerValue(apiResp.PlanDescription)
	model.PlanName = types.StringPointerValue(apiResp.PlanName)
	model.PlanUnit = types.StringPointerValue(apiResp.PlanUnit)
	model.PlanUuid = types.StringPointerValue(apiResp.PlanUuid)
	model.ProjectDescription = types.StringPointerValue(apiResp.ProjectDescription)
	model.ProjectSlug = types.StringPointerValue(apiResp.ProjectSlug)
	model.ProviderName = types.StringPointerValue(apiResp.ProviderName)
	valProviderReviewedAt, diagsProviderReviewedAt := timetypes.NewRFC3339PointerValue(apiResp.ProviderReviewedAt)
	diags.Append(diagsProviderReviewedAt...)
	model.ProviderReviewedAt = valProviderReviewedAt
	model.ProviderReviewedBy = types.StringPointerValue(apiResp.ProviderReviewedBy)
	model.ProviderReviewedByFullName = types.StringPointerValue(apiResp.ProviderReviewedByFullName)
	model.ProviderReviewedByUsername = types.StringPointerValue(apiResp.ProviderReviewedByUsername)
	model.ProviderSlug = types.StringPointerValue(apiResp.ProviderSlug)
	model.ProviderUuid = types.StringPointerValue(apiResp.ProviderUuid)
	model.RequestComment = types.StringPointerValue(apiResp.RequestComment)
	model.ResourceName = types.StringPointerValue(apiResp.ResourceName)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.ResourceUuid = types.StringPointerValue(apiResp.ResourceUuid)
	model.Slug = types.StringPointerValue(apiResp.Slug)
	model.StartDate = types.StringPointerValue(apiResp.StartDate)
	model.State = types.StringPointerValue(apiResp.State)
	model.TerminationComment = types.StringPointerValue(apiResp.TerminationComment)
	model.Type = types.StringPointerValue(apiResp.Type)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
