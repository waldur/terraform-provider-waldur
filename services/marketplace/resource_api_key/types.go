package resource_api_key

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type MarketplaceResourceApiKeyCreateRequest struct {
}

type MarketplaceResourceApiKeyResponse struct {
	UUID *string `json:"uuid"`

	AllowedModels []string `json:"allowed_models" tfsdk:"allowed_models"`

	ClientId *string `json:"client_id,omitempty" tfsdk:"client_id"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	EffectiveModels []string `json:"effective_models" tfsdk:"effective_models"`

	ErrorMessage *string `json:"error_message,omitempty" tfsdk:"error_message"`

	EstimatedCost *string `json:"estimated_cost" tfsdk:"estimated_cost"`

	IssuedAt *string `json:"issued_at,omitempty" tfsdk:"issued_at"`

	LimitResumesAt *string `json:"limit_resumes_at" tfsdk:"limit_resumes_at"`

	Limits map[string]float64 `json:"limits" tfsdk:"limits"`

	OfferingUuid *string `json:"offering_uuid" tfsdk:"offering_uuid"`

	PausedByLimit *bool `json:"paused_by_limit" tfsdk:"paused_by_limit"`

	PendingAction *string `json:"pending_action" tfsdk:"pending_action"`

	ResourceBackendId *string `json:"resource_backend_id" tfsdk:"resource_backend_id"`

	ResourceName *string `json:"resource_name" tfsdk:"resource_name"`

	ResourcePlanUuid *string `json:"resource_plan_uuid" tfsdk:"resource_plan_uuid"`

	ResourceSpendingLimits common.JSONStringMap `json:"resource_spending_limits" tfsdk:"resource_spending_limits"`

	ResourceUuid *string `json:"resource_uuid" tfsdk:"resource_uuid"`

	SpendingLimitStatus []common.SpendingLimitUsage `json:"spending_limit_status" tfsdk:"spending_limit_status"`

	SpendingLimits common.JSONStringMap `json:"spending_limits" tfsdk:"spending_limits"`

	State *string `json:"state,omitempty" tfsdk:"state"`

	UsagePeriod *string `json:"usage_period" tfsdk:"usage_period"`

	UserFullName *string `json:"user_full_name" tfsdk:"user_full_name"`

	UserUuid *string `json:"user_uuid" tfsdk:"user_uuid"`
}

type MarketplaceResourceApiKeyLimitsResponse struct {
}

type MarketplaceResourceApiKeyResourceSpendingLimitsResponse struct {
}

type MarketplaceResourceApiKeySpendingLimitStatusResponse struct {
	ComponentType *string `json:"component_type" tfsdk:"component_type"`

	Limit common.FlexibleNumber `json:"limit" tfsdk:"limit"`

	Period *string `json:"period" tfsdk:"period"`

	Reached *bool `json:"reached" tfsdk:"reached"`

	ResetsAt *string `json:"resets_at" tfsdk:"resets_at"`

	Used common.FlexibleNumber `json:"used" tfsdk:"used"`
}

type MarketplaceResourceApiKeySpendingLimitsResponse struct {
}

func (r *MarketplaceResourceApiKeyResponse) GetState() string {
	if r.State != nil {
		return *r.State
	}
	return "OK"
}

func (r *MarketplaceResourceApiKeyResponse) GetErrorMessage() string {
	if r.ErrorMessage != nil {
		return *r.ErrorMessage
	}
	return ""
}
