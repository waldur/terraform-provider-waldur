package network

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackNetworkCreateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`

	Name *string `json:"name" tfsdk:"name"`
}

type OpenstackNetworkUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`

	Name *string `json:"name,omitempty" tfsdk:"name"`
}

type OpenstackNetworkSetMtuActionRequest struct {
	Mtu *int64 `json:"mtu,omitempty"`
}

type OpenstackNetworkResponse struct {
	UUID *string `json:"uuid"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Customer *string `json:"customer,omitempty" tfsdk:"customer"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	ErrorMessage *string `json:"error_message,omitempty" tfsdk:"error_message"`

	IsExternal *bool `json:"is_external,omitempty" tfsdk:"is_external"`

	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`

	Mtu *int64 `json:"mtu,omitempty" tfsdk:"mtu"`

	Name *string `json:"name" tfsdk:"name"`

	Project *string `json:"project,omitempty" tfsdk:"project"`

	RbacPolicies *[]common.NetworkRBACPolicy `json:"rbac_policies,omitempty" tfsdk:"rbac_policies"`

	ResourceType *string `json:"resource_type,omitempty" tfsdk:"resource_type"`

	SegmentationId *int64 `json:"segmentation_id,omitempty" tfsdk:"segmentation_id"`

	State *string `json:"state,omitempty" tfsdk:"state"`

	Subnets *[]common.OpenStackNestedSubNet `json:"subnets,omitempty" tfsdk:"subnets"`

	Tenant *string `json:"tenant" tfsdk:"tenant"`

	TenantName *string `json:"tenant_name,omitempty" tfsdk:"tenant_name"`

	TenantUuid *string `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`

	Type *string `json:"type,omitempty" tfsdk:"type"`

	Url *string `json:"url,omitempty" tfsdk:"url"`
}

type OpenstackNetworkRbacPoliciesResponse struct {
	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Network *string `json:"network,omitempty" tfsdk:"network"`

	NetworkName *string `json:"network_name,omitempty" tfsdk:"network_name"`

	PolicyType *string `json:"policy_type,omitempty" tfsdk:"policy_type"`

	TargetTenant *string `json:"target_tenant,omitempty" tfsdk:"target_tenant"`

	TargetTenantName *string `json:"target_tenant_name,omitempty" tfsdk:"target_tenant_name"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type OpenstackNetworkSubnetsResponse struct {
	AllocationPools *[]common.OpenStackSubNetAllocationPool `json:"allocation_pools,omitempty" tfsdk:"allocation_pools"`

	Cidr *string `json:"cidr,omitempty" tfsdk:"cidr"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	EnableDhcp *bool `json:"enable_dhcp,omitempty" tfsdk:"enable_dhcp"`

	GatewayIp *string `json:"gateway_ip,omitempty" tfsdk:"gateway_ip"`

	IpVersion *int64 `json:"ip_version,omitempty" tfsdk:"ip_version"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type OpenstackNetworkSubnetsAllocationPoolsResponse struct {
	End *string `json:"end,omitempty" tfsdk:"end"`

	Start *string `json:"start,omitempty" tfsdk:"start"`
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
