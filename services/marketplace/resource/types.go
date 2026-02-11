package resource

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type MarketplaceResourceCreateRequest struct {
}

type MarketplaceResourceUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`

	EndDate *string `json:"end_date,omitempty" tfsdk:"end_date"`

	Name *string `json:"name,omitempty" tfsdk:"name"`
}

type MarketplaceResourceUpdateLimitsActionRequest struct {
	Limits map[string]interface{} `json:"limits"`
}

type MarketplaceResourceCreateLimitsRequest struct {
}

type MarketplaceResourceResponse struct {
	UUID *string `json:"uuid"`

	Attributes map[string]interface{} `json:"attributes,omitempty" tfsdk:"attributes"`

	AvailableActions *[]string `json:"available_actions,omitempty" tfsdk:"available_actions"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	BackendMetadata *MarketplaceResourceBackendMetadataResponse `json:"backend_metadata,omitempty" tfsdk:"backend_metadata"`

	CanTerminate *bool `json:"can_terminate,omitempty" tfsdk:"can_terminate"`

	CategoryIcon *string `json:"category_icon,omitempty" tfsdk:"category_icon"`

	CategoryTitle *string `json:"category_title,omitempty" tfsdk:"category_title"`

	CategoryUuid *string `json:"category_uuid,omitempty" tfsdk:"category_uuid"`

	CreationOrder *MarketplaceResourceCreationOrderResponse `json:"creation_order,omitempty" tfsdk:"creation_order"`

	CurrentUsages map[string]int64 `json:"current_usages,omitempty" tfsdk:"current_usages"`

	CustomerSlug *string `json:"customer_slug,omitempty" tfsdk:"customer_slug"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	Downscaled *bool `json:"downscaled,omitempty" tfsdk:"downscaled"`

	EffectiveId *string `json:"effective_id,omitempty" tfsdk:"effective_id"`

	EndDate *string `json:"end_date,omitempty" tfsdk:"end_date"`

	EndDateRequestedBy *string `json:"end_date_requested_by,omitempty" tfsdk:"end_date_requested_by"`

	Endpoints *[]common.NestedEndpoint `json:"endpoints,omitempty" tfsdk:"endpoints"`

	ErrorMessage *string `json:"error_message,omitempty" tfsdk:"error_message"`

	LastSync *string `json:"last_sync,omitempty" tfsdk:"last_sync"`

	LimitUsage map[string]float64 `json:"limit_usage,omitempty" tfsdk:"limit_usage"`

	Limits map[string]int64 `json:"limits,omitempty" tfsdk:"limits"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Offering *string `json:"offering,omitempty" tfsdk:"offering"`

	OfferingBackendId *string `json:"offering_backend_id,omitempty" tfsdk:"offering_backend_id"`

	OfferingBillable *bool `json:"offering_billable,omitempty" tfsdk:"offering_billable"`

	OfferingComponents *[]common.OfferingComponent `json:"offering_components,omitempty" tfsdk:"offering_components"`

	OfferingDescription *string `json:"offering_description,omitempty" tfsdk:"offering_description"`

	OfferingImage *string `json:"offering_image,omitempty" tfsdk:"offering_image"`

	OfferingName *string `json:"offering_name,omitempty" tfsdk:"offering_name"`

	OfferingShared *bool `json:"offering_shared,omitempty" tfsdk:"offering_shared"`

	OfferingSlug *string `json:"offering_slug,omitempty" tfsdk:"offering_slug"`

	OfferingState *string `json:"offering_state,omitempty" tfsdk:"offering_state"`

	OfferingThumbnail *string `json:"offering_thumbnail,omitempty" tfsdk:"offering_thumbnail"`

	OfferingType *string `json:"offering_type,omitempty" tfsdk:"offering_type"`

	OfferingUuid *string `json:"offering_uuid,omitempty" tfsdk:"offering_uuid"`

	OrderInProgress *MarketplaceResourceOrderInProgressResponse `json:"order_in_progress,omitempty" tfsdk:"order_in_progress"`

	ParentName *string `json:"parent_name,omitempty" tfsdk:"parent_name"`

	ParentOfferingName *string `json:"parent_offering_name,omitempty" tfsdk:"parent_offering_name"`

	ParentOfferingSlug *string `json:"parent_offering_slug,omitempty" tfsdk:"parent_offering_slug"`

	ParentOfferingUuid *string `json:"parent_offering_uuid,omitempty" tfsdk:"parent_offering_uuid"`

	ParentUuid *string `json:"parent_uuid,omitempty" tfsdk:"parent_uuid"`

	Paused *bool `json:"paused,omitempty" tfsdk:"paused"`

	Plan *string `json:"plan,omitempty" tfsdk:"plan"`

	PlanDescription *string `json:"plan_description,omitempty" tfsdk:"plan_description"`

	PlanName *string `json:"plan_name,omitempty" tfsdk:"plan_name"`

	PlanUnit *string `json:"plan_unit,omitempty" tfsdk:"plan_unit"`

	PlanUuid *string `json:"plan_uuid,omitempty" tfsdk:"plan_uuid"`

	Project *string `json:"project,omitempty" tfsdk:"project"`

	ProjectDescription *string `json:"project_description,omitempty" tfsdk:"project_description"`

	ProjectEndDate *string `json:"project_end_date,omitempty" tfsdk:"project_end_date"`

	ProjectEndDateRequestedBy *string `json:"project_end_date_requested_by,omitempty" tfsdk:"project_end_date_requested_by"`

	ProjectSlug *string `json:"project_slug,omitempty" tfsdk:"project_slug"`

	ProviderName *string `json:"provider_name,omitempty" tfsdk:"provider_name"`

	ProviderSlug *string `json:"provider_slug,omitempty" tfsdk:"provider_slug"`

	ProviderUuid *string `json:"provider_uuid,omitempty" tfsdk:"provider_uuid"`

	RenewalDate map[string]string `json:"renewal_date,omitempty" tfsdk:"renewal_date"`

	Report *[]common.ReportSection `json:"report,omitempty" tfsdk:"report"`

	ResourceType *string `json:"resource_type,omitempty" tfsdk:"resource_type"`

	ResourceUuid *string `json:"resource_uuid,omitempty" tfsdk:"resource_uuid"`

	RestrictMemberAccess *bool `json:"restrict_member_access,omitempty" tfsdk:"restrict_member_access"`

	Scope *string `json:"scope,omitempty" tfsdk:"scope"`

	Slug *string `json:"slug,omitempty" tfsdk:"slug"`

	State *string `json:"state,omitempty" tfsdk:"state"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	UserRequiresReconsent *bool `json:"user_requires_reconsent,omitempty" tfsdk:"user_requires_reconsent"`

	Username *string `json:"username,omitempty" tfsdk:"username"`
}

type MarketplaceResourceAttributesResponse struct {
}

type MarketplaceResourceBackendMetadataResponse struct {
	Action *string `json:"action,omitempty" tfsdk:"action"`

	InstanceName *string `json:"instance_name,omitempty" tfsdk:"instance_name"`

	RuntimeState *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`

	State *string `json:"state,omitempty" tfsdk:"state"`
}

type MarketplaceResourceCreationOrderResponse struct {
	ActivationPrice common.FlexibleNumber `json:"activation_price,omitempty" tfsdk:"activation_price"`

	Attachment *string `json:"attachment,omitempty" tfsdk:"attachment"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	CallbackUrl *string `json:"callback_url,omitempty" tfsdk:"callback_url"`

	CanTerminate *bool `json:"can_terminate,omitempty" tfsdk:"can_terminate"`

	CategoryIcon *string `json:"category_icon,omitempty" tfsdk:"category_icon"`

	CategoryTitle *string `json:"category_title,omitempty" tfsdk:"category_title"`

	CategoryUuid *string `json:"category_uuid,omitempty" tfsdk:"category_uuid"`

	CompletedAt *string `json:"completed_at,omitempty" tfsdk:"completed_at"`

	ConsumerReviewedAt *string `json:"consumer_reviewed_at,omitempty" tfsdk:"consumer_reviewed_at"`

	ConsumerReviewedBy *string `json:"consumer_reviewed_by,omitempty" tfsdk:"consumer_reviewed_by"`

	ConsumerReviewedByFullName *string `json:"consumer_reviewed_by_full_name,omitempty" tfsdk:"consumer_reviewed_by_full_name"`

	ConsumerReviewedByUsername *string `json:"consumer_reviewed_by_username,omitempty" tfsdk:"consumer_reviewed_by_username"`

	Cost *string `json:"cost,omitempty" tfsdk:"cost"`

	CreatedByCivilNumber *string `json:"created_by_civil_number,omitempty" tfsdk:"created_by_civil_number"`

	CreatedByFullName *string `json:"created_by_full_name,omitempty" tfsdk:"created_by_full_name"`

	CreatedByUsername *string `json:"created_by_username,omitempty" tfsdk:"created_by_username"`

	CustomerSlug *string `json:"customer_slug,omitempty" tfsdk:"customer_slug"`

	ErrorMessage *string `json:"error_message,omitempty" tfsdk:"error_message"`

	FixedPrice common.FlexibleNumber `json:"fixed_price,omitempty" tfsdk:"fixed_price"`

	Issue *MarketplaceResourceCreationOrderIssueResponse `json:"issue,omitempty" tfsdk:"issue"`

	Limits map[string]int64 `json:"limits,omitempty" tfsdk:"limits"`

	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`

	NewCostEstimate *string `json:"new_cost_estimate,omitempty" tfsdk:"new_cost_estimate"`

	NewPlanName *string `json:"new_plan_name,omitempty" tfsdk:"new_plan_name"`

	NewPlanUuid *string `json:"new_plan_uuid,omitempty" tfsdk:"new_plan_uuid"`

	Offering *string `json:"offering,omitempty" tfsdk:"offering"`

	OfferingBillable *bool `json:"offering_billable,omitempty" tfsdk:"offering_billable"`

	OfferingDescription *string `json:"offering_description,omitempty" tfsdk:"offering_description"`

	OfferingImage *string `json:"offering_image,omitempty" tfsdk:"offering_image"`

	OfferingName *string `json:"offering_name,omitempty" tfsdk:"offering_name"`

	OfferingShared *bool `json:"offering_shared,omitempty" tfsdk:"offering_shared"`

	OfferingThumbnail *string `json:"offering_thumbnail,omitempty" tfsdk:"offering_thumbnail"`

	OfferingType *string `json:"offering_type,omitempty" tfsdk:"offering_type"`

	OfferingUuid *string `json:"offering_uuid,omitempty" tfsdk:"offering_uuid"`

	OldCostEstimate common.FlexibleNumber `json:"old_cost_estimate,omitempty" tfsdk:"old_cost_estimate"`

	OldPlanName *string `json:"old_plan_name,omitempty" tfsdk:"old_plan_name"`

	OldPlanUuid *string `json:"old_plan_uuid,omitempty" tfsdk:"old_plan_uuid"`

	OrderSubtype *string `json:"order_subtype,omitempty" tfsdk:"order_subtype"`

	Output *string `json:"output,omitempty" tfsdk:"output"`

	Plan *string `json:"plan,omitempty" tfsdk:"plan"`

	PlanDescription *string `json:"plan_description,omitempty" tfsdk:"plan_description"`

	PlanName *string `json:"plan_name,omitempty" tfsdk:"plan_name"`

	PlanUnit *string `json:"plan_unit,omitempty" tfsdk:"plan_unit"`

	PlanUuid *string `json:"plan_uuid,omitempty" tfsdk:"plan_uuid"`

	ProjectDescription *string `json:"project_description,omitempty" tfsdk:"project_description"`

	ProjectSlug *string `json:"project_slug,omitempty" tfsdk:"project_slug"`

	ProviderName *string `json:"provider_name,omitempty" tfsdk:"provider_name"`

	ProviderReviewedAt *string `json:"provider_reviewed_at,omitempty" tfsdk:"provider_reviewed_at"`

	ProviderReviewedBy *string `json:"provider_reviewed_by,omitempty" tfsdk:"provider_reviewed_by"`

	ProviderReviewedByFullName *string `json:"provider_reviewed_by_full_name,omitempty" tfsdk:"provider_reviewed_by_full_name"`

	ProviderReviewedByUsername *string `json:"provider_reviewed_by_username,omitempty" tfsdk:"provider_reviewed_by_username"`

	ProviderSlug *string `json:"provider_slug,omitempty" tfsdk:"provider_slug"`

	ProviderUuid *string `json:"provider_uuid,omitempty" tfsdk:"provider_uuid"`

	RequestComment *string `json:"request_comment,omitempty" tfsdk:"request_comment"`

	ResourceName *string `json:"resource_name,omitempty" tfsdk:"resource_name"`

	ResourceType *string `json:"resource_type,omitempty" tfsdk:"resource_type"`

	ResourceUuid *string `json:"resource_uuid,omitempty" tfsdk:"resource_uuid"`

	Slug *string `json:"slug,omitempty" tfsdk:"slug"`

	StartDate *string `json:"start_date,omitempty" tfsdk:"start_date"`

	State *string `json:"state,omitempty" tfsdk:"state"`

	TerminationComment *string `json:"termination_comment,omitempty" tfsdk:"termination_comment"`

	Type *string `json:"type,omitempty" tfsdk:"type"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceResourceCreationOrderIssueResponse struct {
	Key *string `json:"key,omitempty" tfsdk:"key"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceResourceCreationOrderLimitsResponse struct {
}

type MarketplaceResourceCurrentUsagesResponse struct {
}

type MarketplaceResourceEndpointsResponse struct {
	Name *string `json:"name,omitempty" tfsdk:"name"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceResourceLimitUsageResponse struct {
}

type MarketplaceResourceLimitsResponse struct {
}

type MarketplaceResourceOfferingComponentsResponse struct {
	ArticleCode *string `json:"article_code,omitempty" tfsdk:"article_code"`

	BillingType *string `json:"billing_type,omitempty" tfsdk:"billing_type"`

	DefaultLimit *int64 `json:"default_limit,omitempty" tfsdk:"default_limit"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	Factor *int64 `json:"factor,omitempty" tfsdk:"factor"`

	IsBoolean *bool `json:"is_boolean,omitempty" tfsdk:"is_boolean"`

	IsBuiltin *bool `json:"is_builtin,omitempty" tfsdk:"is_builtin"`

	IsPrepaid *bool `json:"is_prepaid,omitempty" tfsdk:"is_prepaid"`

	LimitAmount *int64 `json:"limit_amount,omitempty" tfsdk:"limit_amount"`

	LimitPeriod *string `json:"limit_period,omitempty" tfsdk:"limit_period"`

	MaxAvailableLimit *int64 `json:"max_available_limit,omitempty" tfsdk:"max_available_limit"`

	MaxPrepaidDuration *int64 `json:"max_prepaid_duration,omitempty" tfsdk:"max_prepaid_duration"`

	MaxValue *int64 `json:"max_value,omitempty" tfsdk:"max_value"`

	MeasuredUnit *string `json:"measured_unit,omitempty" tfsdk:"measured_unit"`

	MinPrepaidDuration *int64 `json:"min_prepaid_duration,omitempty" tfsdk:"min_prepaid_duration"`

	MinValue *int64 `json:"min_value,omitempty" tfsdk:"min_value"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	OverageComponent *string `json:"overage_component,omitempty" tfsdk:"overage_component"`

	Type *string `json:"type,omitempty" tfsdk:"type"`

	UnitFactor *int64 `json:"unit_factor,omitempty" tfsdk:"unit_factor"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceResourceOrderInProgressResponse struct {
	ActivationPrice common.FlexibleNumber `json:"activation_price,omitempty" tfsdk:"activation_price"`

	Attachment *string `json:"attachment,omitempty" tfsdk:"attachment"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	CallbackUrl *string `json:"callback_url,omitempty" tfsdk:"callback_url"`

	CanTerminate *bool `json:"can_terminate,omitempty" tfsdk:"can_terminate"`

	CategoryIcon *string `json:"category_icon,omitempty" tfsdk:"category_icon"`

	CategoryTitle *string `json:"category_title,omitempty" tfsdk:"category_title"`

	CategoryUuid *string `json:"category_uuid,omitempty" tfsdk:"category_uuid"`

	CompletedAt *string `json:"completed_at,omitempty" tfsdk:"completed_at"`

	ConsumerReviewedAt *string `json:"consumer_reviewed_at,omitempty" tfsdk:"consumer_reviewed_at"`

	ConsumerReviewedBy *string `json:"consumer_reviewed_by,omitempty" tfsdk:"consumer_reviewed_by"`

	ConsumerReviewedByFullName *string `json:"consumer_reviewed_by_full_name,omitempty" tfsdk:"consumer_reviewed_by_full_name"`

	ConsumerReviewedByUsername *string `json:"consumer_reviewed_by_username,omitempty" tfsdk:"consumer_reviewed_by_username"`

	Cost *string `json:"cost,omitempty" tfsdk:"cost"`

	CreatedByCivilNumber *string `json:"created_by_civil_number,omitempty" tfsdk:"created_by_civil_number"`

	CreatedByFullName *string `json:"created_by_full_name,omitempty" tfsdk:"created_by_full_name"`

	CreatedByUsername *string `json:"created_by_username,omitempty" tfsdk:"created_by_username"`

	CustomerSlug *string `json:"customer_slug,omitempty" tfsdk:"customer_slug"`

	ErrorMessage *string `json:"error_message,omitempty" tfsdk:"error_message"`

	FixedPrice common.FlexibleNumber `json:"fixed_price,omitempty" tfsdk:"fixed_price"`

	Issue *MarketplaceResourceOrderInProgressIssueResponse `json:"issue,omitempty" tfsdk:"issue"`

	Limits map[string]int64 `json:"limits,omitempty" tfsdk:"limits"`

	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`

	NewCostEstimate *string `json:"new_cost_estimate,omitempty" tfsdk:"new_cost_estimate"`

	NewPlanName *string `json:"new_plan_name,omitempty" tfsdk:"new_plan_name"`

	NewPlanUuid *string `json:"new_plan_uuid,omitempty" tfsdk:"new_plan_uuid"`

	Offering *string `json:"offering,omitempty" tfsdk:"offering"`

	OfferingBillable *bool `json:"offering_billable,omitempty" tfsdk:"offering_billable"`

	OfferingDescription *string `json:"offering_description,omitempty" tfsdk:"offering_description"`

	OfferingImage *string `json:"offering_image,omitempty" tfsdk:"offering_image"`

	OfferingName *string `json:"offering_name,omitempty" tfsdk:"offering_name"`

	OfferingShared *bool `json:"offering_shared,omitempty" tfsdk:"offering_shared"`

	OfferingThumbnail *string `json:"offering_thumbnail,omitempty" tfsdk:"offering_thumbnail"`

	OfferingType *string `json:"offering_type,omitempty" tfsdk:"offering_type"`

	OfferingUuid *string `json:"offering_uuid,omitempty" tfsdk:"offering_uuid"`

	OldCostEstimate common.FlexibleNumber `json:"old_cost_estimate,omitempty" tfsdk:"old_cost_estimate"`

	OldPlanName *string `json:"old_plan_name,omitempty" tfsdk:"old_plan_name"`

	OldPlanUuid *string `json:"old_plan_uuid,omitempty" tfsdk:"old_plan_uuid"`

	OrderSubtype *string `json:"order_subtype,omitempty" tfsdk:"order_subtype"`

	Output *string `json:"output,omitempty" tfsdk:"output"`

	Plan *string `json:"plan,omitempty" tfsdk:"plan"`

	PlanDescription *string `json:"plan_description,omitempty" tfsdk:"plan_description"`

	PlanName *string `json:"plan_name,omitempty" tfsdk:"plan_name"`

	PlanUnit *string `json:"plan_unit,omitempty" tfsdk:"plan_unit"`

	PlanUuid *string `json:"plan_uuid,omitempty" tfsdk:"plan_uuid"`

	ProjectDescription *string `json:"project_description,omitempty" tfsdk:"project_description"`

	ProjectSlug *string `json:"project_slug,omitempty" tfsdk:"project_slug"`

	ProviderName *string `json:"provider_name,omitempty" tfsdk:"provider_name"`

	ProviderReviewedAt *string `json:"provider_reviewed_at,omitempty" tfsdk:"provider_reviewed_at"`

	ProviderReviewedBy *string `json:"provider_reviewed_by,omitempty" tfsdk:"provider_reviewed_by"`

	ProviderReviewedByFullName *string `json:"provider_reviewed_by_full_name,omitempty" tfsdk:"provider_reviewed_by_full_name"`

	ProviderReviewedByUsername *string `json:"provider_reviewed_by_username,omitempty" tfsdk:"provider_reviewed_by_username"`

	ProviderSlug *string `json:"provider_slug,omitempty" tfsdk:"provider_slug"`

	ProviderUuid *string `json:"provider_uuid,omitempty" tfsdk:"provider_uuid"`

	RequestComment *string `json:"request_comment,omitempty" tfsdk:"request_comment"`

	ResourceName *string `json:"resource_name,omitempty" tfsdk:"resource_name"`

	ResourceType *string `json:"resource_type,omitempty" tfsdk:"resource_type"`

	ResourceUuid *string `json:"resource_uuid,omitempty" tfsdk:"resource_uuid"`

	Slug *string `json:"slug,omitempty" tfsdk:"slug"`

	StartDate *string `json:"start_date,omitempty" tfsdk:"start_date"`

	State *string `json:"state,omitempty" tfsdk:"state"`

	TerminationComment *string `json:"termination_comment,omitempty" tfsdk:"termination_comment"`

	Type *string `json:"type,omitempty" tfsdk:"type"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceResourceOrderInProgressIssueResponse struct {
	Key *string `json:"key,omitempty" tfsdk:"key"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceResourceOrderInProgressLimitsResponse struct {
}

type MarketplaceResourceRenewalDateResponse struct {
}

type MarketplaceResourceReportResponse struct {
	Body *string `json:"body,omitempty" tfsdk:"body"`

	Header *string `json:"header,omitempty" tfsdk:"header"`
}

func (r *MarketplaceResourceResponse) GetState() string {
	if r.State != nil {
		return *r.State
	}
	return "OK"
}

func (r *MarketplaceResourceResponse) GetErrorMessage() string {
	if r.ErrorMessage != nil {
		return *r.ErrorMessage
	}
	return ""
}
