package subnet

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackSubnetCreateRequest struct {
	AllocationPools *[]common.OpenStackSubNetAllocationPoolRequest `json:"allocation_pools,omitempty"`
	Cidr            *string                                        `json:"cidr,omitempty"`
	Description     *string                                        `json:"description,omitempty"`
	DisableGateway  *bool                                          `json:"disable_gateway,omitempty"`
	DnsNameservers  *[]string                                      `json:"dns_nameservers,omitempty"`
	GatewayIp       *string                                        `json:"gateway_ip,omitempty"`
	HostRoutes      *[]common.OpenStackStaticRouteRequest          `json:"host_routes,omitempty"`
	Name            *string                                        `json:"name"`
}

type OpenstackSubnetUpdateRequest struct {
	AllocationPools *[]common.OpenStackSubNetAllocationPoolRequest `json:"allocation_pools,omitempty"`
	Cidr            *string                                        `json:"cidr,omitempty"`
	Description     *string                                        `json:"description,omitempty"`
	DisableGateway  *bool                                          `json:"disable_gateway,omitempty"`
	DnsNameservers  *[]string                                      `json:"dns_nameservers,omitempty"`
	GatewayIp       *string                                        `json:"gateway_ip,omitempty"`
	HostRoutes      *[]common.OpenStackStaticRouteRequest          `json:"host_routes,omitempty"`
	Name            *string                                        `json:"name,omitempty"`
}

type OpenstackSubnetResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl                   *string                                        `json:"access_url" tfsdk:"access_url"`
	AllocationPools             *[]common.OpenStackSubNetAllocationPoolRequest `json:"allocation_pools" tfsdk:"allocation_pools"`
	BackendId                   *string                                        `json:"backend_id" tfsdk:"backend_id"`
	Cidr                        *string                                        `json:"cidr" tfsdk:"cidr"`
	Created                     *string                                        `json:"created" tfsdk:"created"`
	Customer                    *string                                        `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation        *string                                        `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName                *string                                        `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName          *string                                        `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerUuid                *string                                        `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description                 *string                                        `json:"description" tfsdk:"description"`
	DisableGateway              *bool                                          `json:"disable_gateway" tfsdk:"disable_gateway"`
	DnsNameservers              *[]string                                      `json:"dns_nameservers" tfsdk:"dns_nameservers"`
	EnableDhcp                  *bool                                          `json:"enable_dhcp" tfsdk:"enable_dhcp"`
	ErrorMessage                *string                                        `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback              *string                                        `json:"error_traceback" tfsdk:"error_traceback"`
	GatewayIp                   *string                                        `json:"gateway_ip" tfsdk:"gateway_ip"`
	HostRoutes                  *[]common.OpenStackStaticRouteRequest          `json:"host_routes" tfsdk:"host_routes"`
	IpVersion                   *int64                                         `json:"ip_version" tfsdk:"ip_version"`
	IsConnected                 *bool                                          `json:"is_connected" tfsdk:"is_connected"`
	IsLimitBased                *bool                                          `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                                          `json:"is_usage_based" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string                                        `json:"marketplace_category_name" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                                        `json:"marketplace_category_uuid" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                                        `json:"marketplace_offering_name" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                                        `json:"marketplace_offering_uuid" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                                        `json:"marketplace_plan_uuid" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                                        `json:"marketplace_resource_state" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                                        `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                                        `json:"modified" tfsdk:"modified"`
	Name                        *string                                        `json:"name" tfsdk:"name"`
	Network                     *string                                        `json:"network" tfsdk:"network"`
	NetworkName                 *string                                        `json:"network_name" tfsdk:"network_name"`
	Project                     *string                                        `json:"project" tfsdk:"project"`
	ProjectName                 *string                                        `json:"project_name" tfsdk:"project_name"`
	ProjectUuid                 *string                                        `json:"project_uuid" tfsdk:"project_uuid"`
	ResourceType                *string                                        `json:"resource_type" tfsdk:"resource_type"`
	ServiceName                 *string                                        `json:"service_name" tfsdk:"service_name"`
	ServiceSettings             *string                                        `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                                        `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                                        `json:"service_settings_state" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                                        `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	State                       *string                                        `json:"state" tfsdk:"state"`
	Tenant                      *string                                        `json:"tenant" tfsdk:"tenant"`
	TenantName                  *string                                        `json:"tenant_name" tfsdk:"tenant_name"`
	Url                         *string                                        `json:"url" tfsdk:"url"`
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
