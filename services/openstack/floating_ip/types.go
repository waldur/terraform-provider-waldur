package floating_ip

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackFloatingIpCreateRequest struct {
	Router *string `json:"router,omitempty" tfsdk:"router"`
}

type OpenstackFloatingIpUpdateDescriptionActionRequest struct {
	Description *string `json:"description,omitempty"`
}

type OpenstackFloatingIpResponse struct {
	UUID *string `json:"uuid"`

	Address                 *string                    `json:"address" tfsdk:"address"`
	BackendId               *string                    `json:"backend_id" tfsdk:"backend_id"`
	BackendNetworkId        *string                    `json:"backend_network_id" tfsdk:"backend_network_id"`
	Customer                *string                    `json:"customer" tfsdk:"customer"`
	Description             *string                    `json:"description" tfsdk:"description"`
	ErrorMessage            *string                    `json:"error_message" tfsdk:"error_message"`
	ExternalAddress         *string                    `json:"external_address" tfsdk:"external_address"`
	InstanceName            *string                    `json:"instance_name" tfsdk:"instance_name"`
	InstanceUrl             *string                    `json:"instance_url" tfsdk:"instance_url"`
	InstanceUuid            *string                    `json:"instance_uuid" tfsdk:"instance_uuid"`
	MarketplaceResourceUuid *string                    `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Name                    *string                    `json:"name" tfsdk:"name"`
	Port                    *string                    `json:"port" tfsdk:"port"`
	PortFixedIps            *[]common.OpenStackFixedIp `json:"port_fixed_ips" tfsdk:"port_fixed_ips"`
	Project                 *string                    `json:"project" tfsdk:"project"`
	ResourceType            *string                    `json:"resource_type" tfsdk:"resource_type"`
	RuntimeState            *string                    `json:"runtime_state" tfsdk:"runtime_state"`
	State                   *string                    `json:"state" tfsdk:"state"`
	Tenant                  *string                    `json:"tenant" tfsdk:"tenant"`
	TenantName              *string                    `json:"tenant_name" tfsdk:"tenant_name"`
	TenantUuid              *string                    `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Url                     *string                    `json:"url" tfsdk:"url"`
}

type OpenstackFloatingIpPortFixedIpsResponse struct {
	IpAddress *string `json:"ip_address" tfsdk:"ip_address"`
	SubnetId  *string `json:"subnet_id" tfsdk:"subnet_id"`
}

func (r *OpenstackFloatingIpResponse) GetState() string {
	if r.State != nil {
		return *r.State
	}
	return "OK"
}

func (r *OpenstackFloatingIpResponse) GetErrorMessage() string {
	if r.ErrorMessage != nil {
		return *r.ErrorMessage
	}
	return ""
}
