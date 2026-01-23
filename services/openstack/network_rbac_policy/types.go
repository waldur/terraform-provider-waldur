package network_rbac_policy

// Shared Structs

// Resource Structs

// OpenstackNetworkRbacPolicy Structs

// Create Request
type OpenstackNetworkRbacPolicyCreateRequest struct {
	Network      *string `json:"network" tfsdk:"network"`
	PolicyType   *string `json:"policy_type,omitempty" tfsdk:"policy_type"`
	TargetTenant *string `json:"target_tenant" tfsdk:"target_tenant"`
}

// Update Request
type OpenstackNetworkRbacPolicyUpdateRequest struct {
	Network      *string `json:"network,omitempty" tfsdk:"network"`
	PolicyType   *string `json:"policy_type,omitempty" tfsdk:"policy_type"`
	TargetTenant *string `json:"target_tenant,omitempty" tfsdk:"target_tenant"`
}

// Update Actions Structs

// Response Struct
type OpenstackNetworkRbacPolicyResponse struct {
	UUID *string `json:"uuid"`

	BackendId        *string `json:"backend_id" tfsdk:"backend_id"`
	Created          *string `json:"created" tfsdk:"created"`
	Network          *string `json:"network" tfsdk:"network"`
	NetworkName      *string `json:"network_name" tfsdk:"network_name"`
	PolicyType       *string `json:"policy_type" tfsdk:"policy_type"`
	TargetTenant     *string `json:"target_tenant" tfsdk:"target_tenant"`
	TargetTenantName *string `json:"target_tenant_name" tfsdk:"target_tenant_name"`
	Url              *string `json:"url" tfsdk:"url"`
}
