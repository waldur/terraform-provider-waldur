package port

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackPortCreateRequest struct {
	AllowedAddressPairs *[]common.OpenStackAllowedAddressPairRequest      `json:"allowed_address_pairs,omitempty" tfsdk:"allowed_address_pairs"`
	Description         *string                                           `json:"description,omitempty" tfsdk:"description"`
	FixedIps            *[]common.OpenStackFixedIpRequest                 `json:"fixed_ips,omitempty" tfsdk:"fixed_ips"`
	MacAddress          *string                                           `json:"mac_address,omitempty" tfsdk:"mac_address"`
	Name                *string                                           `json:"name" tfsdk:"name"`
	Network             *string                                           `json:"network,omitempty" tfsdk:"network"`
	PortSecurityEnabled *bool                                             `json:"port_security_enabled,omitempty" tfsdk:"port_security_enabled"`
	SecurityGroups      *[]common.OpenStackPortNestedSecurityGroupRequest `json:"security_groups,omitempty" tfsdk:"security_groups"`
	TargetTenant        *string                                           `json:"target_tenant,omitempty" tfsdk:"target_tenant"`
}

type OpenstackPortUpdateRequest struct {
	Description    *string                                           `json:"description,omitempty" tfsdk:"description"`
	Name           *string                                           `json:"name,omitempty" tfsdk:"name"`
	SecurityGroups *[]common.OpenStackPortNestedSecurityGroupRequest `json:"security_groups,omitempty" tfsdk:"security_groups"`
	TargetTenant   *string                                           `json:"target_tenant,omitempty" tfsdk:"target_tenant"`
}

type OpenstackPortUpdateSecurityGroupsActionRequest struct {
	SecurityGroups []common.OpenStackPortNestedSecurityGroupRequest `json:"security_groups" tfsdk:"security_groups"`
}

type OpenstackPortResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl                   *string                                    `json:"access_url" tfsdk:"access_url"`
	AdminStateUp                *bool                                      `json:"admin_state_up" tfsdk:"admin_state_up"`
	AllowedAddressPairs         []OpenstackPortAllowedAddressPairsResponse `json:"allowed_address_pairs" tfsdk:"allowed_address_pairs"`
	BackendId                   *string                                    `json:"backend_id" tfsdk:"backend_id"`
	Created                     *string                                    `json:"created" tfsdk:"created"`
	Customer                    *string                                    `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation        *string                                    `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName                *string                                    `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName          *string                                    `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerUuid                *string                                    `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description                 *string                                    `json:"description" tfsdk:"description"`
	DeviceId                    *string                                    `json:"device_id" tfsdk:"device_id"`
	DeviceOwner                 *string                                    `json:"device_owner" tfsdk:"device_owner"`
	ErrorMessage                *string                                    `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback              *string                                    `json:"error_traceback" tfsdk:"error_traceback"`
	FixedIps                    []OpenstackPortFixedIpsResponse            `json:"fixed_ips" tfsdk:"fixed_ips"`
	FloatingIps                 []string                                   `json:"floating_ips" tfsdk:"floating_ips"`
	IsLimitBased                *bool                                      `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                                      `json:"is_usage_based" tfsdk:"is_usage_based"`
	MacAddress                  *string                                    `json:"mac_address" tfsdk:"mac_address"`
	MarketplaceCategoryName     *string                                    `json:"marketplace_category_name" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                                    `json:"marketplace_category_uuid" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                                    `json:"marketplace_offering_name" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                                    `json:"marketplace_offering_uuid" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                                    `json:"marketplace_plan_uuid" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                                    `json:"marketplace_resource_state" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                                    `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                                    `json:"modified" tfsdk:"modified"`
	Name                        *string                                    `json:"name" tfsdk:"name"`
	Network                     *string                                    `json:"network" tfsdk:"network"`
	NetworkName                 *string                                    `json:"network_name" tfsdk:"network_name"`
	NetworkUuid                 *string                                    `json:"network_uuid" tfsdk:"network_uuid"`
	PortSecurityEnabled         *bool                                      `json:"port_security_enabled" tfsdk:"port_security_enabled"`
	Project                     *string                                    `json:"project" tfsdk:"project"`
	ProjectName                 *string                                    `json:"project_name" tfsdk:"project_name"`
	ProjectUuid                 *string                                    `json:"project_uuid" tfsdk:"project_uuid"`
	ResourceType                *string                                    `json:"resource_type" tfsdk:"resource_type"`
	SecurityGroups              []OpenstackPortSecurityGroupsResponse      `json:"security_groups" tfsdk:"security_groups"`
	ServiceName                 *string                                    `json:"service_name" tfsdk:"service_name"`
	ServiceSettings             *string                                    `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                                    `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                                    `json:"service_settings_state" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                                    `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	State                       *string                                    `json:"state" tfsdk:"state"`
	Status                      *string                                    `json:"status" tfsdk:"status"`
	Tenant                      *string                                    `json:"tenant" tfsdk:"tenant"`
	TenantName                  *string                                    `json:"tenant_name" tfsdk:"tenant_name"`
	TenantUuid                  *string                                    `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Url                         *string                                    `json:"url" tfsdk:"url"`
}

type OpenstackPortAllowedAddressPairsResponse struct {
	IpAddress  *string `json:"ip_address" tfsdk:"ip_address"`
	MacAddress *string `json:"mac_address" tfsdk:"mac_address"`
}

type OpenstackPortFixedIpsResponse struct {
	IpAddress *string `json:"ip_address" tfsdk:"ip_address"`
	SubnetId  *string `json:"subnet_id" tfsdk:"subnet_id"`
}

type OpenstackPortSecurityGroupsResponse struct {
	Name *string `json:"name" tfsdk:"name"`
}

func (r *OpenstackPortResponse) GetState() string {
	if r.State != nil {
		return *r.State
	}
	return "OK"
}

func (r *OpenstackPortResponse) GetErrorMessage() string {
	if r.ErrorMessage != nil {
		return *r.ErrorMessage
	}
	return ""
}
