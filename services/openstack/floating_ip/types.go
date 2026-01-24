package floating_ip

// OpenstackFloatingIp Structs

type OpenstackFloatingIpCreateRequest struct {
}

type OpenstackFloatingIpUpdateDescriptionActionRequest struct {
	Description *string `json:"description" tfsdk:"description"`
}

type OpenstackFloatingIpResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl                   *string                                   `json:"access_url" tfsdk:"access_url"`
	Address                     *string                                   `json:"address" tfsdk:"address"`
	BackendId                   *string                                   `json:"backend_id" tfsdk:"backend_id"`
	BackendNetworkId            *string                                   `json:"backend_network_id" tfsdk:"backend_network_id"`
	Created                     *string                                   `json:"created" tfsdk:"created"`
	Customer                    *string                                   `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation        *string                                   `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName                *string                                   `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName          *string                                   `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerUuid                *string                                   `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description                 *string                                   `json:"description" tfsdk:"description"`
	ErrorMessage                *string                                   `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback              *string                                   `json:"error_traceback" tfsdk:"error_traceback"`
	ExternalAddress             *string                                   `json:"external_address" tfsdk:"external_address"`
	InstanceName                *string                                   `json:"instance_name" tfsdk:"instance_name"`
	InstanceUrl                 *string                                   `json:"instance_url" tfsdk:"instance_url"`
	InstanceUuid                *string                                   `json:"instance_uuid" tfsdk:"instance_uuid"`
	IsLimitBased                *bool                                     `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                                     `json:"is_usage_based" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string                                   `json:"marketplace_category_name" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                                   `json:"marketplace_category_uuid" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                                   `json:"marketplace_offering_name" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                                   `json:"marketplace_offering_uuid" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                                   `json:"marketplace_plan_uuid" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                                   `json:"marketplace_resource_state" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                                   `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                                   `json:"modified" tfsdk:"modified"`
	Name                        *string                                   `json:"name" tfsdk:"name"`
	Port                        *string                                   `json:"port" tfsdk:"port"`
	PortFixedIps                []OpenstackFloatingIpPortFixedIpsResponse `json:"port_fixed_ips" tfsdk:"port_fixed_ips"`
	Project                     *string                                   `json:"project" tfsdk:"project"`
	ProjectName                 *string                                   `json:"project_name" tfsdk:"project_name"`
	ProjectUuid                 *string                                   `json:"project_uuid" tfsdk:"project_uuid"`
	ResourceType                *string                                   `json:"resource_type" tfsdk:"resource_type"`
	RuntimeState                *string                                   `json:"runtime_state" tfsdk:"runtime_state"`
	ServiceName                 *string                                   `json:"service_name" tfsdk:"service_name"`
	ServiceSettings             *string                                   `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                                   `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                                   `json:"service_settings_state" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                                   `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	State                       *string                                   `json:"state" tfsdk:"state"`
	Tenant                      *string                                   `json:"tenant" tfsdk:"tenant"`
	TenantName                  *string                                   `json:"tenant_name" tfsdk:"tenant_name"`
	TenantUuid                  *string                                   `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Url                         *string                                   `json:"url" tfsdk:"url"`
}

type OpenstackFloatingIpPortFixedIpsResponse struct {
	IpAddress *string `json:"ip_address" tfsdk:"ip_address"`
	SubnetId  *string `json:"subnet_id" tfsdk:"subnet_id"`
}
