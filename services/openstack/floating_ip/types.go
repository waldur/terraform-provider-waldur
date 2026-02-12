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

	Address *string `json:"address,omitempty" tfsdk:"address"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	BackendNetworkId *string `json:"backend_network_id,omitempty" tfsdk:"backend_network_id"`

	Customer *string `json:"customer,omitempty" tfsdk:"customer"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	ErrorMessage *string `json:"error_message,omitempty" tfsdk:"error_message"`

	ExternalAddress *string `json:"external_address,omitempty" tfsdk:"external_address"`

	InstanceUrl *string `json:"instance_url,omitempty" tfsdk:"instance_url"`

	InstanceUuid *string `json:"instance_uuid,omitempty" tfsdk:"instance_uuid"`

	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Port *string `json:"port,omitempty" tfsdk:"port"`

	PortFixedIps *[]common.OpenStackFixedIp `json:"port_fixed_ips,omitempty" tfsdk:"port_fixed_ips"`

	Project *string `json:"project,omitempty" tfsdk:"project"`

	ResourceType *string `json:"resource_type,omitempty" tfsdk:"resource_type"`

	RuntimeState *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`

	State *string `json:"state,omitempty" tfsdk:"state"`

	Tenant *string `json:"tenant" tfsdk:"tenant"`

	TenantName *string `json:"tenant_name,omitempty" tfsdk:"tenant_name"`

	TenantUuid *string `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`

	Url *string `json:"url,omitempty" tfsdk:"url"`
}

type OpenstackFloatingIpPortFixedIpsResponse struct {
	IpAddress *string `json:"ip_address,omitempty" tfsdk:"ip_address"`

	SubnetId *string `json:"subnet_id,omitempty" tfsdk:"subnet_id"`
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
