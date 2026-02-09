package subnet

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackSubnetCreateRequest struct {
	AllocationPools *[]common.OpenStackSubNetAllocationPoolRequest `json:"allocation_pools,omitempty" tfsdk:"allocation_pools"`
	Cidr            *string                                        `json:"cidr,omitempty" tfsdk:"cidr"`
	Description     *string                                        `json:"description,omitempty" tfsdk:"description"`
	DisableGateway  *bool                                          `json:"disable_gateway,omitempty" tfsdk:"disable_gateway"`
	DnsNameservers  *[]string                                      `json:"dns_nameservers,omitempty" tfsdk:"dns_nameservers"`
	GatewayIp       *string                                        `json:"gateway_ip,omitempty" tfsdk:"gateway_ip"`
	HostRoutes      *[]common.OpenStackStaticRouteRequest          `json:"host_routes,omitempty" tfsdk:"host_routes"`
	Name            *string                                        `json:"name" tfsdk:"name"`
}

type OpenstackSubnetUpdateRequest struct {
	AllocationPools *[]common.OpenStackSubNetAllocationPoolRequest `json:"allocation_pools,omitempty" tfsdk:"allocation_pools"`
	Cidr            *string                                        `json:"cidr,omitempty" tfsdk:"cidr"`
	Description     *string                                        `json:"description,omitempty" tfsdk:"description"`
	DisableGateway  *bool                                          `json:"disable_gateway,omitempty" tfsdk:"disable_gateway"`
	DnsNameservers  *[]string                                      `json:"dns_nameservers,omitempty" tfsdk:"dns_nameservers"`
	GatewayIp       *string                                        `json:"gateway_ip,omitempty" tfsdk:"gateway_ip"`
	HostRoutes      *[]common.OpenStackStaticRouteRequest          `json:"host_routes,omitempty" tfsdk:"host_routes"`
	Name            *string                                        `json:"name,omitempty" tfsdk:"name"`
}

type OpenstackSubnetResponse struct {
	UUID *string `json:"uuid"`

	AllocationPools         *[]common.OpenStackSubNetAllocationPoolRequest `json:"allocation_pools" tfsdk:"allocation_pools"`
	BackendId               *string                                        `json:"backend_id" tfsdk:"backend_id"`
	Cidr                    *string                                        `json:"cidr" tfsdk:"cidr"`
	Customer                *string                                        `json:"customer" tfsdk:"customer"`
	Description             *string                                        `json:"description" tfsdk:"description"`
	DisableGateway          *bool                                          `json:"disable_gateway" tfsdk:"disable_gateway"`
	DnsNameservers          *[]string                                      `json:"dns_nameservers" tfsdk:"dns_nameservers"`
	EnableDhcp              *bool                                          `json:"enable_dhcp" tfsdk:"enable_dhcp"`
	ErrorMessage            *string                                        `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback          *string                                        `json:"error_traceback" tfsdk:"error_traceback"`
	GatewayIp               *string                                        `json:"gateway_ip" tfsdk:"gateway_ip"`
	HostRoutes              *[]common.OpenStackStaticRouteRequest          `json:"host_routes" tfsdk:"host_routes"`
	IpVersion               *int64                                         `json:"ip_version" tfsdk:"ip_version"`
	IsConnected             *bool                                          `json:"is_connected" tfsdk:"is_connected"`
	MarketplaceResourceUuid *string                                        `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Name                    *string                                        `json:"name" tfsdk:"name"`
	Network                 *string                                        `json:"network" tfsdk:"network"`
	NetworkName             *string                                        `json:"network_name" tfsdk:"network_name"`
	Project                 *string                                        `json:"project" tfsdk:"project"`
	ResourceType            *string                                        `json:"resource_type" tfsdk:"resource_type"`
	State                   *string                                        `json:"state" tfsdk:"state"`
	Tenant                  *string                                        `json:"tenant" tfsdk:"tenant"`
	TenantName              *string                                        `json:"tenant_name" tfsdk:"tenant_name"`
	Url                     *string                                        `json:"url" tfsdk:"url"`
}

type OpenstackSubnetAllocationPoolsResponse struct {
	End   *string `json:"end" tfsdk:"end"`
	Start *string `json:"start" tfsdk:"start"`
}

type OpenstackSubnetHostRoutesResponse struct {
	Destination *string `json:"destination" tfsdk:"destination"`
	Nexthop     *string `json:"nexthop" tfsdk:"nexthop"`
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
