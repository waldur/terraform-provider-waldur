package subnet

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackSubnetCreateRequest struct {
	AllocationPools *[]common.OpenStackSubNetAllocationPoolRequest `json:"allocation_pools,omitempty" tfsdk:"allocation_pools"`

	Cidr *string `json:"cidr,omitempty" tfsdk:"cidr"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	DisableGateway *bool `json:"disable_gateway,omitempty" tfsdk:"disable_gateway"`

	DnsNameservers *[]string `json:"dns_nameservers,omitempty" tfsdk:"dns_nameservers"`

	GatewayIp *string `json:"gateway_ip,omitempty" tfsdk:"gateway_ip"`

	HostRoutes *[]common.OpenStackStaticRouteRequest `json:"host_routes,omitempty" tfsdk:"host_routes"`

	Name *string `json:"name" tfsdk:"name"`

	Router *string `json:"router,omitempty" tfsdk:"router"`

	SkipRouterConnection *bool `json:"skip_router_connection,omitempty" tfsdk:"skip_router_connection"`
}

type OpenstackSubnetUpdateRequest struct {
	AllocationPools *[]common.OpenStackSubNetAllocationPoolRequest `json:"allocation_pools,omitempty" tfsdk:"allocation_pools"`

	Cidr *string `json:"cidr,omitempty" tfsdk:"cidr"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	DisableGateway *bool `json:"disable_gateway,omitempty" tfsdk:"disable_gateway"`

	DnsNameservers *[]string `json:"dns_nameservers,omitempty" tfsdk:"dns_nameservers"`

	GatewayIp *string `json:"gateway_ip,omitempty" tfsdk:"gateway_ip"`

	HostRoutes *[]common.OpenStackStaticRouteRequest `json:"host_routes,omitempty" tfsdk:"host_routes"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Router *string `json:"router,omitempty" tfsdk:"router"`

	SkipRouterConnection *bool `json:"skip_router_connection,omitempty" tfsdk:"skip_router_connection"`
}

type OpenstackSubnetResponse struct {
	UUID *string `json:"uuid"`

	AllocationPools *[]common.OpenStackSubNetAllocationPoolRequest `json:"allocation_pools,omitempty" tfsdk:"allocation_pools"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Cidr *string `json:"cidr,omitempty" tfsdk:"cidr"`

	Customer *string `json:"customer,omitempty" tfsdk:"customer"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	DisableGateway *bool `json:"disable_gateway,omitempty" tfsdk:"disable_gateway"`

	DnsNameservers *[]string `json:"dns_nameservers,omitempty" tfsdk:"dns_nameservers"`

	EnableDhcp *bool `json:"enable_dhcp,omitempty" tfsdk:"enable_dhcp"`

	ErrorMessage *string `json:"error_message,omitempty" tfsdk:"error_message"`

	GatewayIp *string `json:"gateway_ip,omitempty" tfsdk:"gateway_ip"`

	HostRoutes *[]common.OpenStackStaticRouteRequest `json:"host_routes,omitempty" tfsdk:"host_routes"`

	IpVersion *int64 `json:"ip_version,omitempty" tfsdk:"ip_version"`

	IsConnected *bool `json:"is_connected,omitempty" tfsdk:"is_connected"`

	MarketplaceOfferingType *string `json:"marketplace_offering_type,omitempty" tfsdk:"marketplace_offering_type"`

	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`

	Name *string `json:"name" tfsdk:"name"`

	Network *string `json:"network" tfsdk:"network"`

	NetworkName *string `json:"network_name,omitempty" tfsdk:"network_name"`

	PortSecurityEnabled *bool `json:"port_security_enabled,omitempty" tfsdk:"port_security_enabled"`

	Project *string `json:"project,omitempty" tfsdk:"project"`

	ResourceType *string `json:"resource_type,omitempty" tfsdk:"resource_type"`

	Router *string `json:"router,omitempty" tfsdk:"router"`

	RouterName *string `json:"router_name,omitempty" tfsdk:"router_name"`

	RouterUuid *string `json:"router_uuid,omitempty" tfsdk:"router_uuid"`

	State *string `json:"state,omitempty" tfsdk:"state"`

	Tenant *string `json:"tenant,omitempty" tfsdk:"tenant"`

	TenantName *string `json:"tenant_name,omitempty" tfsdk:"tenant_name"`

	Url *string `json:"url,omitempty" tfsdk:"url"`
}

type OpenstackSubnetAllocationPoolsResponse struct {
	End *string `json:"end" tfsdk:"end"`

	Start *string `json:"start" tfsdk:"start"`
}

type OpenstackSubnetHostRoutesResponse struct {
	Destination *string `json:"destination" tfsdk:"destination"`

	Nexthop *string `json:"nexthop" tfsdk:"nexthop"`
}

func (r *OpenstackSubnetResponse) GetState() string {
	if r.State != nil {
		return *r.State
	}
	return "OK"
}

func (r *OpenstackSubnetResponse) GetErrorMessage() string {
	if r.ErrorMessage != nil {
		return *r.ErrorMessage
	}
	return ""
}
