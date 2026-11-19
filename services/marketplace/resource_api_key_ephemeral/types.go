package resource_api_key_ephemeral

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type MarketplaceResourceApiKeyCreateRequest struct {
}

type MarketplaceResourceApiKeyResponse struct {
	UUID *string `json:"uuid"`

	AllowedModels []string `json:"allowed_models,omitempty" tfsdk:"allowed_models"`

	ApiKey *string `json:"api_key,omitempty" tfsdk:"api_key"`

	ClientId *string `json:"client_id,omitempty" tfsdk:"client_id"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	EffectiveModels []string `json:"effective_models,omitempty" tfsdk:"effective_models"`

	ErrorMessage *string `json:"error_message,omitempty" tfsdk:"error_message"`

	EstimatedCost *string `json:"estimated_cost,omitempty" tfsdk:"estimated_cost"`

	IssuedAt *string `json:"issued_at,omitempty" tfsdk:"issued_at"`

	LimitResumesAt *string `json:"limit_resumes_at,omitempty" tfsdk:"limit_resumes_at"`

	Limits map[string]float64 `json:"limits,omitempty" tfsdk:"limits"`

	OfferingUuid *string `json:"offering_uuid,omitempty" tfsdk:"offering_uuid"`

	PausedByLimit *bool `json:"paused_by_limit,omitempty" tfsdk:"paused_by_limit"`

	PendingAction *string `json:"pending_action,omitempty" tfsdk:"pending_action"`

	ResourceBackendId *string `json:"resource_backend_id,omitempty" tfsdk:"resource_backend_id"`

	ResourceName *string `json:"resource_name,omitempty" tfsdk:"resource_name"`

	ResourcePlanUuid *string `json:"resource_plan_uuid,omitempty" tfsdk:"resource_plan_uuid"`

	ResourceSpendingLimits common.JSONStringMap `json:"resource_spending_limits,omitempty" tfsdk:"resource_spending_limits"`

	ResourceUuid *string `json:"resource_uuid,omitempty" tfsdk:"resource_uuid"`

	SpendingLimitStatus []common.SpendingLimitUsage `json:"spending_limit_status,omitempty" tfsdk:"spending_limit_status"`

	SpendingLimits common.JSONStringMap `json:"spending_limits,omitempty" tfsdk:"spending_limits"`

	State *string `json:"state,omitempty" tfsdk:"state"`

	UsagePeriod *string `json:"usage_period,omitempty" tfsdk:"usage_period"`

	UserFullName *string `json:"user_full_name,omitempty" tfsdk:"user_full_name"`

	UserUuid *string `json:"user_uuid,omitempty" tfsdk:"user_uuid"`
}

type MarketplaceResourceApiKeyLimitsResponse struct {
}

type MarketplaceResourceApiKeyResourceSpendingLimitsResponse struct {
}

type MarketplaceResourceApiKeySpendingLimitStatusResponse struct {
	ComponentType *string `json:"component_type,omitempty" tfsdk:"component_type"`

	Limit common.FlexibleNumber `json:"limit,omitempty" tfsdk:"limit"`

	Period *string `json:"period,omitempty" tfsdk:"period"`

	Reached *bool `json:"reached,omitempty" tfsdk:"reached"`

	ResetsAt *string `json:"resets_at,omitempty" tfsdk:"resets_at"`

	Used common.FlexibleNumber `json:"used,omitempty" tfsdk:"used"`
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
