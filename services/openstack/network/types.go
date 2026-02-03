package network

type OpenstackNetworkCreateRequest struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name"`
}

type OpenstackNetworkUpdateRequest struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type OpenstackNetworkSetMtuActionRequest struct {
	Mtu *int64 `json:"mtu"`
}

type OpenstackNetworkResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl                   *string                                `json:"access_url" tfsdk:"access_url"`
	BackendId                   *string                                `json:"backend_id" tfsdk:"backend_id"`
	Created                     *string                                `json:"created" tfsdk:"created"`
	Customer                    *string                                `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation        *string                                `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName                *string                                `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName          *string                                `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerUuid                *string                                `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description                 *string                                `json:"description" tfsdk:"description"`
	ErrorMessage                *string                                `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback              *string                                `json:"error_traceback" tfsdk:"error_traceback"`
	IsExternal                  *bool                                  `json:"is_external" tfsdk:"is_external"`
	IsLimitBased                *bool                                  `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                                  `json:"is_usage_based" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string                                `json:"marketplace_category_name" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                                `json:"marketplace_category_uuid" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                                `json:"marketplace_offering_name" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                                `json:"marketplace_offering_uuid" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                                `json:"marketplace_plan_uuid" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                                `json:"marketplace_resource_state" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                                `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                                `json:"modified" tfsdk:"modified"`
	Mtu                         *int64                                 `json:"mtu" tfsdk:"mtu"`
	Name                        *string                                `json:"name" tfsdk:"name"`
	Project                     *string                                `json:"project" tfsdk:"project"`
	ProjectName                 *string                                `json:"project_name" tfsdk:"project_name"`
	ProjectUuid                 *string                                `json:"project_uuid" tfsdk:"project_uuid"`
	RbacPolicies                []OpenstackNetworkRbacPoliciesResponse `json:"rbac_policies" tfsdk:"rbac_policies"`
	ResourceType                *string                                `json:"resource_type" tfsdk:"resource_type"`
	SegmentationId              *int64                                 `json:"segmentation_id" tfsdk:"segmentation_id"`
	ServiceName                 *string                                `json:"service_name" tfsdk:"service_name"`
	ServiceSettings             *string                                `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                                `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                                `json:"service_settings_state" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                                `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	State                       *string                                `json:"state" tfsdk:"state"`
	Subnets                     []OpenstackNetworkSubnetsResponse      `json:"subnets" tfsdk:"subnets"`
	Tenant                      *string                                `json:"tenant" tfsdk:"tenant"`
	TenantName                  *string                                `json:"tenant_name" tfsdk:"tenant_name"`
	TenantUuid                  *string                                `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Type                        *string                                `json:"type" tfsdk:"type"`
	Url                         *string                                `json:"url" tfsdk:"url"`
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
	AllocationPools []OpenstackNetworkSubnetsAllocationPoolsResponse `json:"allocation_pools" tfsdk:"allocation_pools"`
	Cidr            *string                                          `json:"cidr" tfsdk:"cidr"`
	Description     *string                                          `json:"description" tfsdk:"description"`
	EnableDhcp      *bool                                            `json:"enable_dhcp" tfsdk:"enable_dhcp"`
	GatewayIp       *string                                          `json:"gateway_ip" tfsdk:"gateway_ip"`
	IpVersion       *int64                                           `json:"ip_version" tfsdk:"ip_version"`
	Name            *string                                          `json:"name" tfsdk:"name"`
	Uuid            *string                                          `json:"uuid" tfsdk:"uuid"`
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
