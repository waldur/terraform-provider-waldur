package resource

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type MarketplaceResourceCreateRequest struct {
}

type MarketplaceResourceUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	EndDate     *string `json:"end_date,omitempty" tfsdk:"end_date"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type MarketplaceResourceResponse struct {
	UUID *string `json:"uuid"`

	AvailableActions          *[]string                                   `json:"available_actions" tfsdk:"available_actions"`
	BackendId                 *string                                     `json:"backend_id" tfsdk:"backend_id"`
	BackendMetadata           *MarketplaceResourceBackendMetadataResponse `json:"backend_metadata" tfsdk:"backend_metadata"`
	CanTerminate              *bool                                       `json:"can_terminate" tfsdk:"can_terminate"`
	CategoryIcon              *string                                     `json:"category_icon" tfsdk:"category_icon"`
	CategoryTitle             *string                                     `json:"category_title" tfsdk:"category_title"`
	CategoryUuid              *string                                     `json:"category_uuid" tfsdk:"category_uuid"`
	CreationOrder             *MarketplaceResourceCreationOrderResponse   `json:"creation_order" tfsdk:"creation_order"`
	CustomerSlug              *string                                     `json:"customer_slug" tfsdk:"customer_slug"`
	Description               *string                                     `json:"description" tfsdk:"description"`
	Downscaled                *bool                                       `json:"downscaled" tfsdk:"downscaled"`
	EffectiveId               *string                                     `json:"effective_id" tfsdk:"effective_id"`
	EndDate                   *string                                     `json:"end_date" tfsdk:"end_date"`
	EndDateRequestedBy        *string                                     `json:"end_date_requested_by" tfsdk:"end_date_requested_by"`
	Endpoints                 *[]common.NestedEndpoint                    `json:"endpoints" tfsdk:"endpoints"`
	ErrorMessage              *string                                     `json:"error_message" tfsdk:"error_message"`
	LastSync                  *string                                     `json:"last_sync" tfsdk:"last_sync"`
	Name                      *string                                     `json:"name" tfsdk:"name"`
	Offering                  *string                                     `json:"offering" tfsdk:"offering"`
	OfferingBackendId         *string                                     `json:"offering_backend_id" tfsdk:"offering_backend_id"`
	OfferingBillable          *bool                                       `json:"offering_billable" tfsdk:"offering_billable"`
	OfferingComponents        *[]common.OfferingComponent                 `json:"offering_components" tfsdk:"offering_components"`
	OfferingDescription       *string                                     `json:"offering_description" tfsdk:"offering_description"`
	OfferingImage             *string                                     `json:"offering_image" tfsdk:"offering_image"`
	OfferingName              *string                                     `json:"offering_name" tfsdk:"offering_name"`
	OfferingShared            *bool                                       `json:"offering_shared" tfsdk:"offering_shared"`
	OfferingSlug              *string                                     `json:"offering_slug" tfsdk:"offering_slug"`
	OfferingState             *string                                     `json:"offering_state" tfsdk:"offering_state"`
	OfferingThumbnail         *string                                     `json:"offering_thumbnail" tfsdk:"offering_thumbnail"`
	OfferingType              *string                                     `json:"offering_type" tfsdk:"offering_type"`
	OfferingUuid              *string                                     `json:"offering_uuid" tfsdk:"offering_uuid"`
	OrderInProgress           *MarketplaceResourceOrderInProgressResponse `json:"order_in_progress" tfsdk:"order_in_progress"`
	ParentName                *string                                     `json:"parent_name" tfsdk:"parent_name"`
	ParentOfferingName        *string                                     `json:"parent_offering_name" tfsdk:"parent_offering_name"`
	ParentOfferingSlug        *string                                     `json:"parent_offering_slug" tfsdk:"parent_offering_slug"`
	ParentOfferingUuid        *string                                     `json:"parent_offering_uuid" tfsdk:"parent_offering_uuid"`
	ParentUuid                *string                                     `json:"parent_uuid" tfsdk:"parent_uuid"`
	Paused                    *bool                                       `json:"paused" tfsdk:"paused"`
	Plan                      *string                                     `json:"plan" tfsdk:"plan"`
	PlanDescription           *string                                     `json:"plan_description" tfsdk:"plan_description"`
	PlanName                  *string                                     `json:"plan_name" tfsdk:"plan_name"`
	PlanUnit                  *string                                     `json:"plan_unit" tfsdk:"plan_unit"`
	PlanUuid                  *string                                     `json:"plan_uuid" tfsdk:"plan_uuid"`
	Project                   *string                                     `json:"project" tfsdk:"project"`
	ProjectDescription        *string                                     `json:"project_description" tfsdk:"project_description"`
	ProjectEndDate            *string                                     `json:"project_end_date" tfsdk:"project_end_date"`
	ProjectEndDateRequestedBy *string                                     `json:"project_end_date_requested_by" tfsdk:"project_end_date_requested_by"`
	ProjectSlug               *string                                     `json:"project_slug" tfsdk:"project_slug"`
	ProviderName              *string                                     `json:"provider_name" tfsdk:"provider_name"`
	ProviderSlug              *string                                     `json:"provider_slug" tfsdk:"provider_slug"`
	ProviderUuid              *string                                     `json:"provider_uuid" tfsdk:"provider_uuid"`
	Report                    *[]common.ReportSection                     `json:"report" tfsdk:"report"`
	ResourceType              *string                                     `json:"resource_type" tfsdk:"resource_type"`
	ResourceUuid              *string                                     `json:"resource_uuid" tfsdk:"resource_uuid"`
	RestrictMemberAccess      *bool                                       `json:"restrict_member_access" tfsdk:"restrict_member_access"`
	Scope                     *string                                     `json:"scope" tfsdk:"scope"`
	Slug                      *string                                     `json:"slug" tfsdk:"slug"`
	State                     *string                                     `json:"state" tfsdk:"state"`
	Url                       *string                                     `json:"url" tfsdk:"url"`
	UserRequiresReconsent     *bool                                       `json:"user_requires_reconsent" tfsdk:"user_requires_reconsent"`
	Username                  *string                                     `json:"username" tfsdk:"username"`
}

type MarketplaceResourceBackendMetadataResponse struct {
	Action       *string `json:"action" tfsdk:"action"`
	InstanceName *string `json:"instance_name" tfsdk:"instance_name"`
	RuntimeState *string `json:"runtime_state" tfsdk:"runtime_state"`
	State        *string `json:"state" tfsdk:"state"`
}

type MarketplaceResourceCreationOrderResponse struct {
	ActivationPrice            *common.FlexibleNumber                         `json:"activation_price" tfsdk:"activation_price"`
	Attachment                 *string                                        `json:"attachment" tfsdk:"attachment"`
	BackendId                  *string                                        `json:"backend_id" tfsdk:"backend_id"`
	CallbackUrl                *string                                        `json:"callback_url" tfsdk:"callback_url"`
	CanTerminate               *bool                                          `json:"can_terminate" tfsdk:"can_terminate"`
	CategoryIcon               *string                                        `json:"category_icon" tfsdk:"category_icon"`
	CategoryTitle              *string                                        `json:"category_title" tfsdk:"category_title"`
	CategoryUuid               *string                                        `json:"category_uuid" tfsdk:"category_uuid"`
	CompletedAt                *string                                        `json:"completed_at" tfsdk:"completed_at"`
	ConsumerReviewedAt         *string                                        `json:"consumer_reviewed_at" tfsdk:"consumer_reviewed_at"`
	ConsumerReviewedBy         *string                                        `json:"consumer_reviewed_by" tfsdk:"consumer_reviewed_by"`
	ConsumerReviewedByFullName *string                                        `json:"consumer_reviewed_by_full_name" tfsdk:"consumer_reviewed_by_full_name"`
	ConsumerReviewedByUsername *string                                        `json:"consumer_reviewed_by_username" tfsdk:"consumer_reviewed_by_username"`
	Cost                       *string                                        `json:"cost" tfsdk:"cost"`
	CreatedByCivilNumber       *string                                        `json:"created_by_civil_number" tfsdk:"created_by_civil_number"`
	CreatedByFullName          *string                                        `json:"created_by_full_name" tfsdk:"created_by_full_name"`
	CreatedByUsername          *string                                        `json:"created_by_username" tfsdk:"created_by_username"`
	CustomerSlug               *string                                        `json:"customer_slug" tfsdk:"customer_slug"`
	ErrorMessage               *string                                        `json:"error_message" tfsdk:"error_message"`
	FixedPrice                 *common.FlexibleNumber                         `json:"fixed_price" tfsdk:"fixed_price"`
	Issue                      *MarketplaceResourceCreationOrderIssueResponse `json:"issue" tfsdk:"issue"`
	MarketplaceResourceUuid    *string                                        `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	NewCostEstimate            *string                                        `json:"new_cost_estimate" tfsdk:"new_cost_estimate"`
	NewPlanName                *string                                        `json:"new_plan_name" tfsdk:"new_plan_name"`
	NewPlanUuid                *string                                        `json:"new_plan_uuid" tfsdk:"new_plan_uuid"`
	Offering                   *string                                        `json:"offering" tfsdk:"offering"`
	OfferingBillable           *bool                                          `json:"offering_billable" tfsdk:"offering_billable"`
	OfferingDescription        *string                                        `json:"offering_description" tfsdk:"offering_description"`
	OfferingImage              *string                                        `json:"offering_image" tfsdk:"offering_image"`
	OfferingName               *string                                        `json:"offering_name" tfsdk:"offering_name"`
	OfferingShared             *bool                                          `json:"offering_shared" tfsdk:"offering_shared"`
	OfferingThumbnail          *string                                        `json:"offering_thumbnail" tfsdk:"offering_thumbnail"`
	OfferingType               *string                                        `json:"offering_type" tfsdk:"offering_type"`
	OfferingUuid               *string                                        `json:"offering_uuid" tfsdk:"offering_uuid"`
	OldCostEstimate            *common.FlexibleNumber                         `json:"old_cost_estimate" tfsdk:"old_cost_estimate"`
	OldPlanName                *string                                        `json:"old_plan_name" tfsdk:"old_plan_name"`
	OldPlanUuid                *string                                        `json:"old_plan_uuid" tfsdk:"old_plan_uuid"`
	OrderSubtype               *string                                        `json:"order_subtype" tfsdk:"order_subtype"`
	Output                     *string                                        `json:"output" tfsdk:"output"`
	Plan                       *string                                        `json:"plan" tfsdk:"plan"`
	PlanDescription            *string                                        `json:"plan_description" tfsdk:"plan_description"`
	PlanName                   *string                                        `json:"plan_name" tfsdk:"plan_name"`
	PlanUnit                   *string                                        `json:"plan_unit" tfsdk:"plan_unit"`
	PlanUuid                   *string                                        `json:"plan_uuid" tfsdk:"plan_uuid"`
	ProjectDescription         *string                                        `json:"project_description" tfsdk:"project_description"`
	ProjectSlug                *string                                        `json:"project_slug" tfsdk:"project_slug"`
	ProviderName               *string                                        `json:"provider_name" tfsdk:"provider_name"`
	ProviderReviewedAt         *string                                        `json:"provider_reviewed_at" tfsdk:"provider_reviewed_at"`
	ProviderReviewedBy         *string                                        `json:"provider_reviewed_by" tfsdk:"provider_reviewed_by"`
	ProviderReviewedByFullName *string                                        `json:"provider_reviewed_by_full_name" tfsdk:"provider_reviewed_by_full_name"`
	ProviderReviewedByUsername *string                                        `json:"provider_reviewed_by_username" tfsdk:"provider_reviewed_by_username"`
	ProviderSlug               *string                                        `json:"provider_slug" tfsdk:"provider_slug"`
	ProviderUuid               *string                                        `json:"provider_uuid" tfsdk:"provider_uuid"`
	RequestComment             *string                                        `json:"request_comment" tfsdk:"request_comment"`
	ResourceName               *string                                        `json:"resource_name" tfsdk:"resource_name"`
	ResourceType               *string                                        `json:"resource_type" tfsdk:"resource_type"`
	ResourceUuid               *string                                        `json:"resource_uuid" tfsdk:"resource_uuid"`
	Slug                       *string                                        `json:"slug" tfsdk:"slug"`
	StartDate                  *string                                        `json:"start_date" tfsdk:"start_date"`
	State                      *string                                        `json:"state" tfsdk:"state"`
	TerminationComment         *string                                        `json:"termination_comment" tfsdk:"termination_comment"`
	Type                       *string                                        `json:"type" tfsdk:"type"`
	Url                        *string                                        `json:"url" tfsdk:"url"`
	Uuid                       *string                                        `json:"uuid" tfsdk:"uuid"`
}

type MarketplaceResourceCreationOrderIssueResponse struct {
	Key  *string `json:"key" tfsdk:"key"`
	Uuid *string `json:"uuid" tfsdk:"uuid"`
}

type MarketplaceResourceEndpointsResponse struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
	Uuid *string `json:"uuid" tfsdk:"uuid"`
}

type MarketplaceResourceOfferingComponentsResponse struct {
	ArticleCode        *string `json:"article_code" tfsdk:"article_code"`
	BillingType        *string `json:"billing_type" tfsdk:"billing_type"`
	DefaultLimit       *int64  `json:"default_limit" tfsdk:"default_limit"`
	Description        *string `json:"description" tfsdk:"description"`
	Factor             *int64  `json:"factor" tfsdk:"factor"`
	IsBoolean          *bool   `json:"is_boolean" tfsdk:"is_boolean"`
	IsBuiltin          *bool   `json:"is_builtin" tfsdk:"is_builtin"`
	IsPrepaid          *bool   `json:"is_prepaid" tfsdk:"is_prepaid"`
	LimitAmount        *int64  `json:"limit_amount" tfsdk:"limit_amount"`
	LimitPeriod        *string `json:"limit_period" tfsdk:"limit_period"`
	MaxAvailableLimit  *int64  `json:"max_available_limit" tfsdk:"max_available_limit"`
	MaxPrepaidDuration *int64  `json:"max_prepaid_duration" tfsdk:"max_prepaid_duration"`
	MaxValue           *int64  `json:"max_value" tfsdk:"max_value"`
	MeasuredUnit       *string `json:"measured_unit" tfsdk:"measured_unit"`
	MinPrepaidDuration *int64  `json:"min_prepaid_duration" tfsdk:"min_prepaid_duration"`
	MinValue           *int64  `json:"min_value" tfsdk:"min_value"`
	Name               *string `json:"name" tfsdk:"name"`
	OverageComponent   *string `json:"overage_component" tfsdk:"overage_component"`
	Type               *string `json:"type" tfsdk:"type"`
	UnitFactor         *int64  `json:"unit_factor" tfsdk:"unit_factor"`
	Uuid               *string `json:"uuid" tfsdk:"uuid"`
}

type MarketplaceResourceOrderInProgressResponse struct {
	ActivationPrice            *common.FlexibleNumber                           `json:"activation_price" tfsdk:"activation_price"`
	Attachment                 *string                                          `json:"attachment" tfsdk:"attachment"`
	BackendId                  *string                                          `json:"backend_id" tfsdk:"backend_id"`
	CallbackUrl                *string                                          `json:"callback_url" tfsdk:"callback_url"`
	CanTerminate               *bool                                            `json:"can_terminate" tfsdk:"can_terminate"`
	CategoryIcon               *string                                          `json:"category_icon" tfsdk:"category_icon"`
	CategoryTitle              *string                                          `json:"category_title" tfsdk:"category_title"`
	CategoryUuid               *string                                          `json:"category_uuid" tfsdk:"category_uuid"`
	CompletedAt                *string                                          `json:"completed_at" tfsdk:"completed_at"`
	ConsumerReviewedAt         *string                                          `json:"consumer_reviewed_at" tfsdk:"consumer_reviewed_at"`
	ConsumerReviewedBy         *string                                          `json:"consumer_reviewed_by" tfsdk:"consumer_reviewed_by"`
	ConsumerReviewedByFullName *string                                          `json:"consumer_reviewed_by_full_name" tfsdk:"consumer_reviewed_by_full_name"`
	ConsumerReviewedByUsername *string                                          `json:"consumer_reviewed_by_username" tfsdk:"consumer_reviewed_by_username"`
	Cost                       *string                                          `json:"cost" tfsdk:"cost"`
	CreatedByCivilNumber       *string                                          `json:"created_by_civil_number" tfsdk:"created_by_civil_number"`
	CreatedByFullName          *string                                          `json:"created_by_full_name" tfsdk:"created_by_full_name"`
	CreatedByUsername          *string                                          `json:"created_by_username" tfsdk:"created_by_username"`
	CustomerSlug               *string                                          `json:"customer_slug" tfsdk:"customer_slug"`
	ErrorMessage               *string                                          `json:"error_message" tfsdk:"error_message"`
	FixedPrice                 *common.FlexibleNumber                           `json:"fixed_price" tfsdk:"fixed_price"`
	Issue                      *MarketplaceResourceOrderInProgressIssueResponse `json:"issue" tfsdk:"issue"`
	MarketplaceResourceUuid    *string                                          `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	NewCostEstimate            *string                                          `json:"new_cost_estimate" tfsdk:"new_cost_estimate"`
	NewPlanName                *string                                          `json:"new_plan_name" tfsdk:"new_plan_name"`
	NewPlanUuid                *string                                          `json:"new_plan_uuid" tfsdk:"new_plan_uuid"`
	Offering                   *string                                          `json:"offering" tfsdk:"offering"`
	OfferingBillable           *bool                                            `json:"offering_billable" tfsdk:"offering_billable"`
	OfferingDescription        *string                                          `json:"offering_description" tfsdk:"offering_description"`
	OfferingImage              *string                                          `json:"offering_image" tfsdk:"offering_image"`
	OfferingName               *string                                          `json:"offering_name" tfsdk:"offering_name"`
	OfferingShared             *bool                                            `json:"offering_shared" tfsdk:"offering_shared"`
	OfferingThumbnail          *string                                          `json:"offering_thumbnail" tfsdk:"offering_thumbnail"`
	OfferingType               *string                                          `json:"offering_type" tfsdk:"offering_type"`
	OfferingUuid               *string                                          `json:"offering_uuid" tfsdk:"offering_uuid"`
	OldCostEstimate            *common.FlexibleNumber                           `json:"old_cost_estimate" tfsdk:"old_cost_estimate"`
	OldPlanName                *string                                          `json:"old_plan_name" tfsdk:"old_plan_name"`
	OldPlanUuid                *string                                          `json:"old_plan_uuid" tfsdk:"old_plan_uuid"`
	OrderSubtype               *string                                          `json:"order_subtype" tfsdk:"order_subtype"`
	Output                     *string                                          `json:"output" tfsdk:"output"`
	Plan                       *string                                          `json:"plan" tfsdk:"plan"`
	PlanDescription            *string                                          `json:"plan_description" tfsdk:"plan_description"`
	PlanName                   *string                                          `json:"plan_name" tfsdk:"plan_name"`
	PlanUnit                   *string                                          `json:"plan_unit" tfsdk:"plan_unit"`
	PlanUuid                   *string                                          `json:"plan_uuid" tfsdk:"plan_uuid"`
	ProjectDescription         *string                                          `json:"project_description" tfsdk:"project_description"`
	ProjectSlug                *string                                          `json:"project_slug" tfsdk:"project_slug"`
	ProviderName               *string                                          `json:"provider_name" tfsdk:"provider_name"`
	ProviderReviewedAt         *string                                          `json:"provider_reviewed_at" tfsdk:"provider_reviewed_at"`
	ProviderReviewedBy         *string                                          `json:"provider_reviewed_by" tfsdk:"provider_reviewed_by"`
	ProviderReviewedByFullName *string                                          `json:"provider_reviewed_by_full_name" tfsdk:"provider_reviewed_by_full_name"`
	ProviderReviewedByUsername *string                                          `json:"provider_reviewed_by_username" tfsdk:"provider_reviewed_by_username"`
	ProviderSlug               *string                                          `json:"provider_slug" tfsdk:"provider_slug"`
	ProviderUuid               *string                                          `json:"provider_uuid" tfsdk:"provider_uuid"`
	RequestComment             *string                                          `json:"request_comment" tfsdk:"request_comment"`
	ResourceName               *string                                          `json:"resource_name" tfsdk:"resource_name"`
	ResourceType               *string                                          `json:"resource_type" tfsdk:"resource_type"`
	ResourceUuid               *string                                          `json:"resource_uuid" tfsdk:"resource_uuid"`
	Slug                       *string                                          `json:"slug" tfsdk:"slug"`
	StartDate                  *string                                          `json:"start_date" tfsdk:"start_date"`
	State                      *string                                          `json:"state" tfsdk:"state"`
	TerminationComment         *string                                          `json:"termination_comment" tfsdk:"termination_comment"`
	Type                       *string                                          `json:"type" tfsdk:"type"`
	Url                        *string                                          `json:"url" tfsdk:"url"`
	Uuid                       *string                                          `json:"uuid" tfsdk:"uuid"`
}

type MarketplaceResourceOrderInProgressIssueResponse struct {
	Key  *string `json:"key" tfsdk:"key"`
	Uuid *string `json:"uuid" tfsdk:"uuid"`
}

type MarketplaceResourceReportResponse struct {
	Body   *string `json:"body" tfsdk:"body"`
	Header *string `json:"header" tfsdk:"header"`
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
