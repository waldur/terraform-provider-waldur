package network_rbac_policy

type OpenstackNetworkRbacPolicyCreateRequest struct {
	Network *string `json:"network" tfsdk:"network"`

	PolicyType *string `json:"policy_type,omitempty" tfsdk:"policy_type"`

	TargetTenant *string `json:"target_tenant" tfsdk:"target_tenant"`
}

type OpenstackNetworkRbacPolicyUpdateRequest struct {
	Network *string `json:"network,omitempty" tfsdk:"network"`

	PolicyType *string `json:"policy_type,omitempty" tfsdk:"policy_type"`

	TargetTenant *string `json:"target_tenant,omitempty" tfsdk:"target_tenant"`
}

type OpenstackNetworkRbacPolicyResponse struct {
	UUID *string `json:"uuid"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Network *string `json:"network" tfsdk:"network"`

	NetworkName *string `json:"network_name,omitempty" tfsdk:"network_name"`

	PolicyType *string `json:"policy_type,omitempty" tfsdk:"policy_type"`

	TargetTenant *string `json:"target_tenant" tfsdk:"target_tenant"`

	TargetTenantName *string `json:"target_tenant_name,omitempty" tfsdk:"target_tenant_name"`

	Url *string `json:"url,omitempty" tfsdk:"url"`
}

func (r *OpenstackNetworkRbacPolicyResponse) GetState() string {
	return "OK"
}

func (r *OpenstackNetworkRbacPolicyResponse) GetErrorMessage() string {
	return ""
}
