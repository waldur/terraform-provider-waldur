package order

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type MarketplaceOrderCreateRequest struct {
	AcceptingTermsOfService *bool `json:"accepting_terms_of_service,omitempty" tfsdk:"accepting_terms_of_service"`

	Attributes map[string]string `json:"attributes,omitempty" tfsdk:"attributes"`

	CallbackUrl *string `json:"callback_url,omitempty" tfsdk:"callback_url"`

	Limits map[string]int64 `json:"limits,omitempty" tfsdk:"limits"`

	Offering *string `json:"offering" tfsdk:"offering"`

	Plan *string `json:"plan,omitempty" tfsdk:"plan"`

	Project *string `json:"project" tfsdk:"project"`

	RequestComment *string `json:"request_comment,omitempty" tfsdk:"request_comment"`

	Slug *string `json:"slug,omitempty" tfsdk:"slug"`

	StartDate *string `json:"start_date,omitempty" tfsdk:"start_date"`

	Type *string `json:"type,omitempty" tfsdk:"type"`
}

type MarketplaceOrderCreateAttributesRequest struct {
}

type MarketplaceOrderCreateLimitsRequest struct {
}

type MarketplaceOrderUpdateRequest struct {
	Limits map[string]int64 `json:"limits,omitempty" tfsdk:"limits"`

	StartDate *string `json:"start_date,omitempty" tfsdk:"start_date"`
}

type MarketplaceOrderUpdateLimitsRequest struct {
}

type MarketplaceOrderResponse struct {
	UUID *string `json:"uuid"`

	ActivationPrice common.FlexibleNumber `json:"activation_price,omitempty" tfsdk:"activation_price"`

	Attachment *string `json:"attachment,omitempty" tfsdk:"attachment"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	CallbackUrl *string `json:"callback_url,omitempty" tfsdk:"callback_url"`

	CanTerminate *bool `json:"can_terminate,omitempty" tfsdk:"can_terminate"`

	CategoryIcon *string `json:"category_icon,omitempty" tfsdk:"category_icon"`

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

	Issue *MarketplaceOrderIssueResponse `json:"issue,omitempty" tfsdk:"issue"`

	Limits map[string]int64 `json:"limits,omitempty" tfsdk:"limits"`

	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`

	NewCostEstimate *string `json:"new_cost_estimate,omitempty" tfsdk:"new_cost_estimate"`

	NewPlanName *string `json:"new_plan_name,omitempty" tfsdk:"new_plan_name"`

	NewPlanUuid *string `json:"new_plan_uuid,omitempty" tfsdk:"new_plan_uuid"`

	Offering *string `json:"offering" tfsdk:"offering"`

	OfferingBillable *bool `json:"offering_billable,omitempty" tfsdk:"offering_billable"`

	OfferingDescription *string `json:"offering_description,omitempty" tfsdk:"offering_description"`

	OfferingImage *string `json:"offering_image,omitempty" tfsdk:"offering_image"`

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
}

type MarketplaceOrderIssueResponse struct {
	Key *string `json:"key,omitempty" tfsdk:"key"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceOrderLimitsResponse struct {
}

func (r *MarketplaceOrderResponse) GetState() string {
	if r.State != nil {
		return *r.State
	}
	return "OK"
}

func (r *MarketplaceOrderResponse) GetErrorMessage() string {
	if r.ErrorMessage != nil {
		return *r.ErrorMessage
	}
	return ""
}
