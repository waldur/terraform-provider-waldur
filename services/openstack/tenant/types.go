package tenant

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackTenantCreateRequest struct {
	Project    *string                         `json:"project"`
	Offering   *string                         `json:"offering"`
	Plan       *string                         `json:"plan,omitempty"`
	Limits     map[string]float64              `json:"limits,omitempty"`
	Attributes OpenstackTenantCreateAttributes `json:"attributes"`
}
type OpenstackTenantCreateAttributes struct {
	AvailabilityZone            *string                                       `json:"availability_zone,omitempty"`
	Description                 *string                                       `json:"description,omitempty"`
	Name                        *string                                       `json:"name,omitempty"`
	SecurityGroups              *[]common.OpenStackTenantSecurityGroupRequest `json:"security_groups,omitempty"`
	SkipConnectionExtnet        *bool                                         `json:"skip_connection_extnet,omitempty"`
	SkipCreationOfDefaultRouter *bool                                         `json:"skip_creation_of_default_router,omitempty"`
	SkipCreationOfDefaultSubnet *bool                                         `json:"skip_creation_of_default_subnet,omitempty"`
	SubnetCidr                  *string                                       `json:"subnet_cidr,omitempty"`
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
	SecurityGroups []common.OpenStackTenantSecurityGroupRequest `json:"security_groups"`
}

type OpenstackTenantResponse struct {
	UUID *string `json:"uuid"`

	AvailabilityZone            *string         `json:"availability_zone" tfsdk:"availability_zone"`
	BackendId                   *string         `json:"backend_id" tfsdk:"backend_id"`
	Created                     *string         `json:"created" tfsdk:"created"`
	Customer                    *string         `json:"customer" tfsdk:"customer"`
	DefaultVolumeTypeName       *string         `json:"default_volume_type_name" tfsdk:"default_volume_type_name"`
	Description                 *string         `json:"description" tfsdk:"description"`
	ErrorMessage                *string         `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback              *string         `json:"error_traceback" tfsdk:"error_traceback"`
	ExternalNetworkId           *string         `json:"external_network_id" tfsdk:"external_network_id"`
	InternalNetworkId           *string         `json:"internal_network_id" tfsdk:"internal_network_id"`
	MarketplaceResourceUuid     *string         `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string         `json:"modified" tfsdk:"modified"`
	Name                        *string         `json:"name" tfsdk:"name"`
	Project                     *string         `json:"project" tfsdk:"project"`
	Quotas                      *[]common.Quota `json:"quotas" tfsdk:"quotas"`
	ResourceType                *string         `json:"resource_type" tfsdk:"resource_type"`
	SkipCreationOfDefaultRouter *bool           `json:"skip_creation_of_default_router" tfsdk:"skip_creation_of_default_router"`
	State                       *string         `json:"state" tfsdk:"state"`
	Url                         *string         `json:"url" tfsdk:"url"`
	UserPassword                *string         `json:"user_password" tfsdk:"user_password"`
	UserUsername                *string         `json:"user_username" tfsdk:"user_username"`
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
