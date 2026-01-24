package security_group

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// OpenstackSecurityGroup Structs

type OpenstackSecurityGroupCreateRequest struct {
	Description *string                                          `json:"description,omitempty" tfsdk:"description"`
	Name        *string                                          `json:"name" tfsdk:"name"`
	Rules       []common.OpenStackSecurityGroupRuleCreateRequest `json:"rules" tfsdk:"rules"`
}

type OpenstackSecurityGroupUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type OpenstackSecurityGroupSetRulesActionRequest struct {
	Rules []common.OpenStackSecurityGroupRuleCreateRequest `json:"rules" tfsdk:"rules"`
}

type OpenstackSecurityGroupResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl                   *string                               `json:"access_url" tfsdk:"access_url"`
	BackendId                   *string                               `json:"backend_id" tfsdk:"backend_id"`
	Created                     *string                               `json:"created" tfsdk:"created"`
	Customer                    *string                               `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation        *string                               `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName                *string                               `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName          *string                               `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerUuid                *string                               `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description                 *string                               `json:"description" tfsdk:"description"`
	ErrorMessage                *string                               `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback              *string                               `json:"error_traceback" tfsdk:"error_traceback"`
	IsLimitBased                *bool                                 `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                                 `json:"is_usage_based" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string                               `json:"marketplace_category_name" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                               `json:"marketplace_category_uuid" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                               `json:"marketplace_offering_name" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                               `json:"marketplace_offering_uuid" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                               `json:"marketplace_plan_uuid" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                               `json:"marketplace_resource_state" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                               `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                               `json:"modified" tfsdk:"modified"`
	Name                        *string                               `json:"name" tfsdk:"name"`
	Project                     *string                               `json:"project" tfsdk:"project"`
	ProjectName                 *string                               `json:"project_name" tfsdk:"project_name"`
	ProjectUuid                 *string                               `json:"project_uuid" tfsdk:"project_uuid"`
	ResourceType                *string                               `json:"resource_type" tfsdk:"resource_type"`
	Rules                       []OpenstackSecurityGroupRulesResponse `json:"rules" tfsdk:"rules"`
	ServiceName                 *string                               `json:"service_name" tfsdk:"service_name"`
	ServiceSettings             *string                               `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                               `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                               `json:"service_settings_state" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                               `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	State                       *string                               `json:"state" tfsdk:"state"`
	Tenant                      *string                               `json:"tenant" tfsdk:"tenant"`
	TenantName                  *string                               `json:"tenant_name" tfsdk:"tenant_name"`
	TenantUuid                  *string                               `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Url                         *string                               `json:"url" tfsdk:"url"`
}

type OpenstackSecurityGroupRulesResponse struct {
	Cidr        *string `json:"cidr" tfsdk:"cidr"`
	Description *string `json:"description" tfsdk:"description"`
	Direction   *string `json:"direction" tfsdk:"direction"`
	Ethertype   *string `json:"ethertype" tfsdk:"ethertype"`
	FromPort    *int64  `json:"from_port" tfsdk:"from_port"`
	Protocol    *string `json:"protocol" tfsdk:"protocol"`
	RemoteGroup *string `json:"remote_group" tfsdk:"remote_group"`
	ToPort      *int64  `json:"to_port" tfsdk:"to_port"`
}

func (r *OpenstackSecurityGroupResponse) GetState() string {
	if r.State != nil {
		return *r.State
	}
	return "OK"
}

func (r *OpenstackSecurityGroupResponse) GetErrorMessage() string {
	if r.ErrorMessage != nil {
		return *r.ErrorMessage
	}
	return ""
}
