package network

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackNetworkCreateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
}

type OpenstackNetworkUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type OpenstackNetworkSetMtuActionRequest struct {
	Mtu *int64 `json:"mtu"`
}

type OpenstackNetworkResponse struct {
	UUID *string `json:"uuid"`

	BackendId               *string                         `json:"backend_id" tfsdk:"backend_id"`
	Created                 *string                         `json:"created" tfsdk:"created"`
	Customer                *string                         `json:"customer" tfsdk:"customer"`
	Description             *string                         `json:"description" tfsdk:"description"`
	ErrorMessage            *string                         `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback          *string                         `json:"error_traceback" tfsdk:"error_traceback"`
	IsExternal              *bool                           `json:"is_external" tfsdk:"is_external"`
	MarketplaceResourceUuid *string                         `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                *string                         `json:"modified" tfsdk:"modified"`
	Mtu                     *int64                          `json:"mtu" tfsdk:"mtu"`
	Name                    *string                         `json:"name" tfsdk:"name"`
	Project                 *string                         `json:"project" tfsdk:"project"`
	RbacPolicies            *[]common.NetworkRBACPolicy     `json:"rbac_policies" tfsdk:"rbac_policies"`
	ResourceType            *string                         `json:"resource_type" tfsdk:"resource_type"`
	SegmentationId          *int64                          `json:"segmentation_id" tfsdk:"segmentation_id"`
	State                   *string                         `json:"state" tfsdk:"state"`
	Subnets                 *[]common.OpenStackNestedSubNet `json:"subnets" tfsdk:"subnets"`
	Tenant                  *string                         `json:"tenant" tfsdk:"tenant"`
	TenantName              *string                         `json:"tenant_name" tfsdk:"tenant_name"`
	TenantUuid              *string                         `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Type                    *string                         `json:"type" tfsdk:"type"`
	Url                     *string                         `json:"url" tfsdk:"url"`
}

type OpenstackNetworkRbacPoliciesResponse struct {
	BackendId        *string `json:"backend_id" tfsdk:"backend_id"`
	Created          *string `json:"created" tfsdk:"created"`
	Network          *string `json:"network" tfsdk:"network"`
	NetworkName      *string `json:"network_name" tfsdk:"network_name"`
	PolicyType       *string `json:"policy_type" tfsdk:"policy_type"`
	TargetTenant     *string `json:"target_tenant" tfsdk:"target_tenant"`
	TargetTenantName *string `json:"target_tenant_name" tfsdk:"target_tenant_name"`
	Url              *string `json:"url" tfsdk:"url"`
	Uuid             *string `json:"uuid" tfsdk:"uuid"`
}

type OpenstackNetworkSubnetsResponse struct {
	AllocationPools *[]common.OpenStackSubNetAllocationPool `json:"allocation_pools" tfsdk:"allocation_pools"`
	Cidr            *string                                 `json:"cidr" tfsdk:"cidr"`
	Description     *string                                 `json:"description" tfsdk:"description"`
	EnableDhcp      *bool                                   `json:"enable_dhcp" tfsdk:"enable_dhcp"`
	GatewayIp       *string                                 `json:"gateway_ip" tfsdk:"gateway_ip"`
	IpVersion       *int64                                  `json:"ip_version" tfsdk:"ip_version"`
	Name            *string                                 `json:"name" tfsdk:"name"`
	Uuid            *string                                 `json:"uuid" tfsdk:"uuid"`
}

type OpenstackNetworkSubnetsAllocationPoolsResponse struct {
	End   *string `json:"end" tfsdk:"end"`
	Start *string `json:"start" tfsdk:"start"`
}

func (r *OpenstackNetworkResponse) GetState() string {
	if r.State != nil {
		return *r.State
	}
	return "OK"
}

func (r *OpenstackNetworkResponse) GetErrorMessage() string {
	if r.ErrorMessage != nil {
		return *r.ErrorMessage
	}
	return ""
}
