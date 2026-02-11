package order

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

func IssueType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"key":  types.StringType,
		"uuid": types.StringType,
	}}
}

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

func (m *MarketplaceOrderFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Marketplace Order",
		Attributes: map[string]schema.Attribute{
			"can_approve_as_consumer": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Can approve as consumer",
			},
			"can_approve_as_provider": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Can approve as provider",
			},
			"category_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Category UUID",
			},
			"created": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Created after",
			},
			"customer_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer UUID",
			},
			"modified": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Modified after",
			},
			"offering": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Offering",
			},
			"offering_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Offering UUID",
			},
			"parent_offering_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "UUID of the parent offering",
			},
			"project_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project UUID",
			},
			"provider_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Provider UUID",
			},
			"query": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Search by order UUID, slug, project name or resource name",
			},
			"resource": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Resource URL",
			},
			"resource_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Resource name",
			},
			"resource_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Resource UUID",
			},
			"service_manager_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Service manager UUID",
			},
		},
	}
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
	CreatedByCivilNumber       types.String      `tfsdk:"created_by_civil_number"`
	CreatedByFullName          types.String      `tfsdk:"created_by_full_name"`
	CreatedByUsername          types.String      `tfsdk:"created_by_username"`
	CustomerSlug               types.String      `tfsdk:"customer_slug"`
	ErrorMessage               types.String      `tfsdk:"error_message"`
	FixedPrice                 types.Float64     `tfsdk:"fixed_price"`
	Issue                      types.Object      `tfsdk:"issue"`
	MarketplaceResourceUuid    types.String      `tfsdk:"marketplace_resource_uuid"`
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
	ProjectSlug                types.String      `tfsdk:"project_slug"`
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

	model.ActivationPrice = types.Float64PointerValue(apiResp.ActivationPrice.Float64Ptr())

	model.Attachment = common.StringPointerValue(apiResp.Attachment)

	model.BackendId = common.StringPointerValue(apiResp.BackendId)

	model.CallbackUrl = common.StringPointerValue(apiResp.CallbackUrl)

	model.CanTerminate = types.BoolPointerValue(apiResp.CanTerminate)

	model.CategoryIcon = common.StringPointerValue(apiResp.CategoryIcon)

	model.CategoryTitle = common.StringPointerValue(apiResp.CategoryTitle)

	model.CategoryUuid = common.StringPointerValue(apiResp.CategoryUuid)

	valCompletedAt, diagsCompletedAt := timetypes.NewRFC3339PointerValue(apiResp.CompletedAt)
	diags.Append(diagsCompletedAt...)
	model.CompletedAt = valCompletedAt

	valConsumerReviewedAt, diagsConsumerReviewedAt := timetypes.NewRFC3339PointerValue(apiResp.ConsumerReviewedAt)
	diags.Append(diagsConsumerReviewedAt...)
	model.ConsumerReviewedAt = valConsumerReviewedAt

	model.ConsumerReviewedBy = common.StringPointerValue(apiResp.ConsumerReviewedBy)

	model.ConsumerReviewedByFullName = common.StringPointerValue(apiResp.ConsumerReviewedByFullName)

	model.ConsumerReviewedByUsername = common.StringPointerValue(apiResp.ConsumerReviewedByUsername)

	model.Cost = common.StringPointerValue(apiResp.Cost)

	model.CreatedByCivilNumber = common.StringPointerValue(apiResp.CreatedByCivilNumber)

	model.CreatedByFullName = common.StringPointerValue(apiResp.CreatedByFullName)

	model.CreatedByUsername = common.StringPointerValue(apiResp.CreatedByUsername)

	model.CustomerSlug = common.StringPointerValue(apiResp.CustomerSlug)

	model.ErrorMessage = common.StringPointerValue(apiResp.ErrorMessage)

	model.FixedPrice = types.Float64PointerValue(apiResp.FixedPrice.Float64Ptr())

	if apiResp.Issue != nil {
		valIssue, diagsIssue := types.ObjectValueFrom(ctx, IssueType().AttrTypes, *apiResp.Issue)
		diags.Append(diagsIssue...)
		model.Issue = valIssue
	} else {
		model.Issue = types.ObjectNull(IssueType().AttrTypes)
	}

	model.MarketplaceResourceUuid = common.StringPointerValue(apiResp.MarketplaceResourceUuid)

	model.NewCostEstimate = common.StringPointerValue(apiResp.NewCostEstimate)

	model.NewPlanName = common.StringPointerValue(apiResp.NewPlanName)

	model.NewPlanUuid = common.StringPointerValue(apiResp.NewPlanUuid)

	model.Offering = common.StringPointerValue(apiResp.Offering)

	model.OfferingBillable = types.BoolPointerValue(apiResp.OfferingBillable)

	model.OfferingDescription = common.StringPointerValue(apiResp.OfferingDescription)

	model.OfferingImage = common.StringPointerValue(apiResp.OfferingImage)

	model.OfferingName = common.StringPointerValue(apiResp.OfferingName)

	model.OfferingShared = types.BoolPointerValue(apiResp.OfferingShared)

	model.OfferingThumbnail = common.StringPointerValue(apiResp.OfferingThumbnail)

	model.OfferingType = common.StringPointerValue(apiResp.OfferingType)

	model.OfferingUuid = common.StringPointerValue(apiResp.OfferingUuid)

	model.OldCostEstimate = types.Float64PointerValue(apiResp.OldCostEstimate.Float64Ptr())

	model.OldPlanName = common.StringPointerValue(apiResp.OldPlanName)

	model.OldPlanUuid = common.StringPointerValue(apiResp.OldPlanUuid)

	model.OrderSubtype = common.StringPointerValue(apiResp.OrderSubtype)

	model.Output = common.StringPointerValue(apiResp.Output)

	model.Plan = common.StringPointerValue(apiResp.Plan)

	model.PlanDescription = common.StringPointerValue(apiResp.PlanDescription)

	model.PlanName = common.StringPointerValue(apiResp.PlanName)

	model.PlanUnit = common.StringPointerValue(apiResp.PlanUnit)

	model.PlanUuid = common.StringPointerValue(apiResp.PlanUuid)

	model.ProjectDescription = common.StringPointerValue(apiResp.ProjectDescription)

	model.ProjectSlug = common.StringPointerValue(apiResp.ProjectSlug)

	model.ProviderName = common.StringPointerValue(apiResp.ProviderName)

	valProviderReviewedAt, diagsProviderReviewedAt := timetypes.NewRFC3339PointerValue(apiResp.ProviderReviewedAt)
	diags.Append(diagsProviderReviewedAt...)
	model.ProviderReviewedAt = valProviderReviewedAt

	model.ProviderReviewedBy = common.StringPointerValue(apiResp.ProviderReviewedBy)

	model.ProviderReviewedByFullName = common.StringPointerValue(apiResp.ProviderReviewedByFullName)

	model.ProviderReviewedByUsername = common.StringPointerValue(apiResp.ProviderReviewedByUsername)

	model.ProviderSlug = common.StringPointerValue(apiResp.ProviderSlug)

	model.ProviderUuid = common.StringPointerValue(apiResp.ProviderUuid)

	model.RequestComment = common.StringPointerValue(apiResp.RequestComment)

	model.ResourceName = common.StringPointerValue(apiResp.ResourceName)

	model.ResourceType = common.StringPointerValue(apiResp.ResourceType)

	model.ResourceUuid = common.StringPointerValue(apiResp.ResourceUuid)

	model.Slug = common.StringPointerValue(apiResp.Slug)

	model.StartDate = common.StringPointerValue(apiResp.StartDate)

	model.State = common.StringPointerValue(apiResp.State)

	model.TerminationComment = common.StringPointerValue(apiResp.TerminationComment)

	model.Type = common.StringPointerValue(apiResp.Type)

	model.Url = common.StringPointerValue(apiResp.Url)

	return diags
}
