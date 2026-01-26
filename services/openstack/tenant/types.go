package tenant

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackTenantCreateRequest struct {
	Project    *string                         `json:"project" tfsdk:"project"`
	Offering   *string                         `json:"offering" tfsdk:"offering"`
	Attributes OpenstackTenantCreateAttributes `json:"attributes" tfsdk:"attributes"`
}
type OpenstackTenantCreateAttributes struct {
	AvailabilityZone            *string                                       `json:"availability_zone,omitempty" tfsdk:"availability_zone"`
	Description                 *string                                       `json:"description,omitempty" tfsdk:"description"`
	Name                        *string                                       `json:"name,omitempty" tfsdk:"name"`
	SecurityGroups              *[]common.OpenStackTenantSecurityGroupRequest `json:"security_groups,omitempty" tfsdk:"security_groups"`
	SkipConnectionExtnet        *bool                                         `json:"skip_connection_extnet,omitempty" tfsdk:"skip_connection_extnet"`
	SkipCreationOfDefaultRouter *bool                                         `json:"skip_creation_of_default_router,omitempty" tfsdk:"skip_creation_of_default_router"`
	SkipCreationOfDefaultSubnet *bool                                         `json:"skip_creation_of_default_subnet,omitempty" tfsdk:"skip_creation_of_default_subnet"`
	SubnetCidr                  *string                                       `json:"subnet_cidr,omitempty" tfsdk:"subnet_cidr"`
}

type OpenstackTenantUpdateRequest struct {
	AvailabilityZone            *string                                       `json:"availability_zone,omitempty" tfsdk:"availability_zone"`
	DefaultVolumeTypeName       *string                                       `json:"default_volume_type_name,omitempty" tfsdk:"default_volume_type_name"`
	Description                 *string                                       `json:"description,omitempty" tfsdk:"description"`
	Name                        *string                                       `json:"name,omitempty" tfsdk:"name"`
	SecurityGroups              *[]common.OpenStackTenantSecurityGroupRequest `json:"security_groups,omitempty" tfsdk:"security_groups"`
	SkipCreationOfDefaultRouter *bool                                         `json:"skip_creation_of_default_router,omitempty" tfsdk:"skip_creation_of_default_router"`
	SkipCreationOfDefaultSubnet *bool                                         `json:"skip_creation_of_default_subnet,omitempty" tfsdk:"skip_creation_of_default_subnet"`
}

type OpenstackTenantPushSecurityGroupsActionRequest struct {
	SecurityGroups []common.OpenStackTenantSecurityGroupRequest `json:"security_groups" tfsdk:"security_groups"`
}

type OpenstackTenantResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl                   *string                         `json:"access_url" tfsdk:"access_url"`
	AvailabilityZone            *string                         `json:"availability_zone" tfsdk:"availability_zone"`
	BackendId                   *string                         `json:"backend_id" tfsdk:"backend_id"`
	Created                     *string                         `json:"created" tfsdk:"created"`
	Customer                    *string                         `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation        *string                         `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName                *string                         `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName          *string                         `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerUuid                *string                         `json:"customer_uuid" tfsdk:"customer_uuid"`
	DefaultVolumeTypeName       *string                         `json:"default_volume_type_name" tfsdk:"default_volume_type_name"`
	Description                 *string                         `json:"description" tfsdk:"description"`
	ErrorMessage                *string                         `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback              *string                         `json:"error_traceback" tfsdk:"error_traceback"`
	ExternalNetworkId           *string                         `json:"external_network_id" tfsdk:"external_network_id"`
	InternalNetworkId           *string                         `json:"internal_network_id" tfsdk:"internal_network_id"`
	IsLimitBased                *bool                           `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                           `json:"is_usage_based" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string                         `json:"marketplace_category_name" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                         `json:"marketplace_category_uuid" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                         `json:"marketplace_offering_name" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                         `json:"marketplace_offering_uuid" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                         `json:"marketplace_plan_uuid" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                         `json:"marketplace_resource_state" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                         `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                         `json:"modified" tfsdk:"modified"`
	Name                        *string                         `json:"name" tfsdk:"name"`
	Project                     *string                         `json:"project" tfsdk:"project"`
	ProjectName                 *string                         `json:"project_name" tfsdk:"project_name"`
	ProjectUuid                 *string                         `json:"project_uuid" tfsdk:"project_uuid"`
	Quotas                      []OpenstackTenantQuotasResponse `json:"quotas" tfsdk:"quotas"`
	ResourceType                *string                         `json:"resource_type" tfsdk:"resource_type"`
	ServiceName                 *string                         `json:"service_name" tfsdk:"service_name"`
	ServiceSettings             *string                         `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                         `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                         `json:"service_settings_state" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                         `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	SkipCreationOfDefaultRouter *bool                           `json:"skip_creation_of_default_router" tfsdk:"skip_creation_of_default_router"`
	State                       *string                         `json:"state" tfsdk:"state"`
	Url                         *string                         `json:"url" tfsdk:"url"`
	UserPassword                *string                         `json:"user_password" tfsdk:"user_password"`
	UserUsername                *string                         `json:"user_username" tfsdk:"user_username"`
}

type OpenstackTenantQuotasResponse struct {
	Limit *int64  `json:"limit" tfsdk:"limit"`
	Name  *string `json:"name" tfsdk:"name"`
	Usage *int64  `json:"usage" tfsdk:"usage"`
}

func (r *OpenstackTenantResponse) GetState() string {
	if r.State != nil {
		return *r.State
	}
	return "OK"
}

func (r *OpenstackTenantResponse) GetErrorMessage() string {
	if r.ErrorMessage != nil {
		return *r.ErrorMessage
	}
	return ""
}
