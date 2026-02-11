package port

import (
	"encoding/json"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackPortCreateRequest struct {
	AllowedAddressPairs *[]common.OpenStackAllowedAddressPairRequest `json:"allowed_address_pairs,omitempty" tfsdk:"allowed_address_pairs"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	FixedIps *[]common.OpenStackFixedIpRequest `json:"fixed_ips,omitempty" tfsdk:"fixed_ips"`

	MacAddress *string `json:"mac_address,omitempty" tfsdk:"mac_address"`

	Name *string `json:"name" tfsdk:"name"`

	Network *string `json:"network,omitempty" tfsdk:"network"`

	PortSecurityEnabled *bool `json:"port_security_enabled,omitempty" tfsdk:"port_security_enabled"`

	SecurityGroups *[]common.OpenStackPortNestedSecurityGroupRequest `json:"security_groups,omitempty" tfsdk:"security_groups"`

	TargetTenant *string `json:"target_tenant,omitempty" tfsdk:"target_tenant"`
}

type OpenstackPortUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	SecurityGroups *[]common.OpenStackPortNestedSecurityGroupRequest `json:"security_groups,omitempty" tfsdk:"security_groups"`

	TargetTenant *string `json:"target_tenant,omitempty" tfsdk:"target_tenant"`
}

type OpenstackPortUpdateSecurityGroupsActionRequest struct {
	SecurityGroups []common.OpenStackPortNestedSecurityGroupRequest `json:"-"`
}

func (r OpenstackPortUpdateSecurityGroupsActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.SecurityGroups)
}

type OpenstackPortResponse struct {
	UUID *string `json:"uuid"`

	AdminStateUp *bool `json:"admin_state_up,omitempty" tfsdk:"admin_state_up"`

	AllowedAddressPairs *[]common.OpenStackAllowedAddressPairRequest `json:"allowed_address_pairs,omitempty" tfsdk:"allowed_address_pairs"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Customer *string `json:"customer,omitempty" tfsdk:"customer"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	DeviceId *string `json:"device_id,omitempty" tfsdk:"device_id"`

	DeviceOwner *string `json:"device_owner,omitempty" tfsdk:"device_owner"`

	ErrorMessage *string `json:"error_message,omitempty" tfsdk:"error_message"`

	FixedIps *[]common.OpenStackFixedIpRequest `json:"fixed_ips,omitempty" tfsdk:"fixed_ips"`

	FloatingIps *[]string `json:"floating_ips,omitempty" tfsdk:"floating_ips"`

	MacAddress *string `json:"mac_address,omitempty" tfsdk:"mac_address"`

	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`

	Name *string `json:"name" tfsdk:"name"`

	Network *string `json:"network,omitempty" tfsdk:"network"`

	NetworkName *string `json:"network_name,omitempty" tfsdk:"network_name"`

	NetworkUuid *string `json:"network_uuid,omitempty" tfsdk:"network_uuid"`

	PortSecurityEnabled *bool `json:"port_security_enabled,omitempty" tfsdk:"port_security_enabled"`

	Project *string `json:"project,omitempty" tfsdk:"project"`

	ResourceType *string `json:"resource_type,omitempty" tfsdk:"resource_type"`

	SecurityGroups *[]common.OpenStackPortNestedSecurityGroupRequest `json:"security_groups,omitempty" tfsdk:"security_groups"`

	State *string `json:"state,omitempty" tfsdk:"state"`

	Status *string `json:"status,omitempty" tfsdk:"status"`

	Tenant *string `json:"tenant,omitempty" tfsdk:"tenant"`

	TenantName *string `json:"tenant_name,omitempty" tfsdk:"tenant_name"`

	TenantUuid *string `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`

	Url *string `json:"url,omitempty" tfsdk:"url"`
}

type OpenstackPortAllowedAddressPairsResponse struct {
	IpAddress *string `json:"ip_address,omitempty" tfsdk:"ip_address"`

	MacAddress *string `json:"mac_address,omitempty" tfsdk:"mac_address"`
}

type OpenstackPortFixedIpsResponse struct {
	IpAddress *string `json:"ip_address" tfsdk:"ip_address"`

	SubnetId *string `json:"subnet_id" tfsdk:"subnet_id"`
}

type OpenstackPortSecurityGroupsResponse struct {
	Name *string `json:"name" tfsdk:"name"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
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
