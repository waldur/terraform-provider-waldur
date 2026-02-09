package volume

type OpenstackVolumeCreateRequest struct {
	Project    *string                         `json:"project"`
	Offering   *string                         `json:"offering"`
	Plan       *string                         `json:"plan,omitempty"`
	Limits     map[string]float64              `json:"limits,omitempty"`
	Attributes OpenstackVolumeCreateAttributes `json:"attributes"`
}
type OpenstackVolumeCreateAttributes struct {
	AvailabilityZone *string `json:"availability_zone,omitempty"`
	Description      *string `json:"description,omitempty"`
	Image            *string `json:"image,omitempty"`
	Name             *string `json:"name,omitempty"`
	Size             *int64  `json:"size,omitempty"`
	Type             *string `json:"type,omitempty"`
}

type OpenstackVolumeUpdateRequest struct {
	Bootable    *bool   `json:"bootable,omitempty" tfsdk:"bootable"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type OpenstackVolumeExtendActionRequest struct {
	Size *int64 `json:"size"`
}
type OpenstackVolumeRetypeActionRequest struct {
	Type *string `json:"type"`
}

type OpenstackVolumeResponse struct {
	UUID *string `json:"uuid"`

	Action                  *string `json:"action" tfsdk:"action"`
	AvailabilityZone        *string `json:"availability_zone" tfsdk:"availability_zone"`
	AvailabilityZoneName    *string `json:"availability_zone_name" tfsdk:"availability_zone_name"`
	BackendId               *string `json:"backend_id" tfsdk:"backend_id"`
	Bootable                *bool   `json:"bootable" tfsdk:"bootable"`
	Created                 *string `json:"created" tfsdk:"created"`
	Customer                *string `json:"customer" tfsdk:"customer"`
	Description             *string `json:"description" tfsdk:"description"`
	Device                  *string `json:"device" tfsdk:"device"`
	ErrorMessage            *string `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback          *string `json:"error_traceback" tfsdk:"error_traceback"`
	ExtendEnabled           *bool   `json:"extend_enabled" tfsdk:"extend_enabled"`
	Image                   *string `json:"image" tfsdk:"image"`
	ImageMetadata           *string `json:"image_metadata" tfsdk:"image_metadata"`
	ImageName               *string `json:"image_name" tfsdk:"image_name"`
	Instance                *string `json:"instance" tfsdk:"instance"`
	InstanceMarketplaceUuid *string `json:"instance_marketplace_uuid" tfsdk:"instance_marketplace_uuid"`
	InstanceName            *string `json:"instance_name" tfsdk:"instance_name"`
	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                *string `json:"modified" tfsdk:"modified"`
	Name                    *string `json:"name" tfsdk:"name"`
	Project                 *string `json:"project" tfsdk:"project"`
	ResourceType            *string `json:"resource_type" tfsdk:"resource_type"`
	RuntimeState            *string `json:"runtime_state" tfsdk:"runtime_state"`
	Size                    *int64  `json:"size" tfsdk:"size"`
	SourceSnapshot          *string `json:"source_snapshot" tfsdk:"source_snapshot"`
	State                   *string `json:"state" tfsdk:"state"`
	Tenant                  *string `json:"tenant" tfsdk:"tenant"`
	TenantUuid              *string `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Type                    *string `json:"type" tfsdk:"type"`
	TypeName                *string `json:"type_name" tfsdk:"type_name"`
	Url                     *string `json:"url" tfsdk:"url"`
}

func (r *OpenstackVolumeResponse) GetState() string {
	if r.State != nil {
		return *r.State
	}
	return "OK"
}

func (r *OpenstackVolumeResponse) GetErrorMessage() string {
	if r.ErrorMessage != nil {
		return *r.ErrorMessage
	}
	return ""
}
