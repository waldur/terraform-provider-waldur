package volume

// OpenstackVolume Structs

type OpenstackVolumeCreateRequest struct {
	Project    *string                         `json:"project" tfsdk:"project"`
	Offering   *string                         `json:"offering" tfsdk:"offering"`
	Attributes OpenstackVolumeCreateAttributes `json:"attributes" tfsdk:"attributes"`
}
type OpenstackVolumeCreateAttributes struct {
	AvailabilityZone *string `json:"availability_zone,omitempty" tfsdk:"availability_zone"`
	Description      *string `json:"description,omitempty" tfsdk:"description"`
	Image            *string `json:"image,omitempty" tfsdk:"image"`
	Name             *string `json:"name,omitempty" tfsdk:"name"`
	Size             *int64  `json:"size,omitempty" tfsdk:"size"`
	Type             *string `json:"type,omitempty" tfsdk:"type"`
}

type OpenstackVolumeUpdateRequest struct {
	Bootable    *bool   `json:"bootable,omitempty" tfsdk:"bootable"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type OpenstackVolumeExtendActionRequest struct {
	Size *int64 `json:"size" tfsdk:"size"`
}
type OpenstackVolumeRetypeActionRequest struct {
	Type *string `json:"type" tfsdk:"type"`
}

type OpenstackVolumeResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl                   *string `json:"access_url" tfsdk:"access_url"`
	Action                      *string `json:"action" tfsdk:"action"`
	AvailabilityZone            *string `json:"availability_zone" tfsdk:"availability_zone"`
	AvailabilityZoneName        *string `json:"availability_zone_name" tfsdk:"availability_zone_name"`
	BackendId                   *string `json:"backend_id" tfsdk:"backend_id"`
	Bootable                    *bool   `json:"bootable" tfsdk:"bootable"`
	Created                     *string `json:"created" tfsdk:"created"`
	Customer                    *string `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description                 *string `json:"description" tfsdk:"description"`
	Device                      *string `json:"device" tfsdk:"device"`
	ErrorMessage                *string `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback" tfsdk:"error_traceback"`
	ExtendEnabled               *bool   `json:"extend_enabled" tfsdk:"extend_enabled"`
	Image                       *string `json:"image" tfsdk:"image"`
	ImageMetadata               *string `json:"image_metadata" tfsdk:"image_metadata"`
	ImageName                   *string `json:"image_name" tfsdk:"image_name"`
	Instance                    *string `json:"instance" tfsdk:"instance"`
	InstanceMarketplaceUuid     *string `json:"instance_marketplace_uuid" tfsdk:"instance_marketplace_uuid"`
	InstanceName                *string `json:"instance_name" tfsdk:"instance_name"`
	IsLimitBased                *bool   `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased                *bool   `json:"is_usage_based" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string `json:"marketplace_category_name" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string `json:"marketplace_category_uuid" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string `json:"marketplace_offering_name" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string `json:"marketplace_offering_uuid" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string `json:"marketplace_plan_uuid" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string `json:"marketplace_resource_state" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string `json:"modified" tfsdk:"modified"`
	Name                        *string `json:"name" tfsdk:"name"`
	Project                     *string `json:"project" tfsdk:"project"`
	ProjectName                 *string `json:"project_name" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid" tfsdk:"project_uuid"`
	ResourceType                *string `json:"resource_type" tfsdk:"resource_type"`
	RuntimeState                *string `json:"runtime_state" tfsdk:"runtime_state"`
	ServiceName                 *string `json:"service_name" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	Size                        *int64  `json:"size" tfsdk:"size"`
	SourceSnapshot              *string `json:"source_snapshot" tfsdk:"source_snapshot"`
	State                       *string `json:"state" tfsdk:"state"`
	Tenant                      *string `json:"tenant" tfsdk:"tenant"`
	TenantUuid                  *string `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Type                        *string `json:"type" tfsdk:"type"`
	TypeName                    *string `json:"type_name" tfsdk:"type_name"`
	Url                         *string `json:"url" tfsdk:"url"`
}
