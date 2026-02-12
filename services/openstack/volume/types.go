package volume

type OpenstackVolumeCreateRequest struct {
	EndDate    *string                         `json:"end_date,omitempty"`
	Limits     map[string]float64              `json:"limits,omitempty"`
	Offering   *string                         `json:"offering"`
	Plan       *string                         `json:"plan,omitempty"`
	Project    *string                         `json:"project"`
	StartDate  *string                         `json:"start_date,omitempty"`
	Attributes OpenstackVolumeCreateAttributes `json:"attributes"`
}
type OpenstackVolumeCreateAttributes struct {
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	Description *string `json:"description,omitempty"`

	EndDate *string `json:"end_date,omitempty"`

	Image *string `json:"image,omitempty"`

	Limits map[string]float64 `json:"limits,omitempty"`

	Name *string `json:"name"`

	Plan *string `json:"plan,omitempty"`

	Size *int64 `json:"size,omitempty"`

	StartDate *string `json:"start_date,omitempty"`

	Type *string `json:"type,omitempty"`
}

type OpenstackVolumeCreateLimitsRequest struct {
}

type OpenstackVolumeUpdateRequest struct {
	Bootable *bool `json:"bootable,omitempty" tfsdk:"bootable"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	Name *string `json:"name,omitempty" tfsdk:"name"`
}

type OpenstackVolumeExtendActionRequest struct {
	Size *int64 `json:"size,omitempty"`
}

type OpenstackVolumeRetypeActionRequest struct {
	Type *string `json:"type,omitempty"`
}

type OpenstackVolumeResponse struct {
	UUID *string `json:"uuid"`

	AvailabilityZone *string `json:"availability_zone,omitempty" tfsdk:"availability_zone"`

	AvailabilityZoneName *string `json:"availability_zone_name,omitempty" tfsdk:"availability_zone_name"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Bootable *bool `json:"bootable,omitempty" tfsdk:"bootable"`

	Customer *string `json:"customer,omitempty" tfsdk:"customer"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	Device *string `json:"device,omitempty" tfsdk:"device"`

	ErrorMessage *string `json:"error_message,omitempty" tfsdk:"error_message"`

	ExtendEnabled *bool `json:"extend_enabled,omitempty" tfsdk:"extend_enabled"`

	Image *string `json:"image,omitempty" tfsdk:"image"`

	ImageMetadata *string `json:"image_metadata,omitempty" tfsdk:"image_metadata"`

	ImageName *string `json:"image_name,omitempty" tfsdk:"image_name"`

	Instance *string `json:"instance,omitempty" tfsdk:"instance"`

	InstanceMarketplaceUuid *string `json:"instance_marketplace_uuid,omitempty" tfsdk:"instance_marketplace_uuid"`

	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`

	Name *string `json:"name" tfsdk:"name"`

	Project *string `json:"project" tfsdk:"project"`

	ResourceType *string `json:"resource_type,omitempty" tfsdk:"resource_type"`

	RuntimeState *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`

	Size *int64 `json:"size,omitempty" tfsdk:"size"`

	SourceSnapshot *string `json:"source_snapshot,omitempty" tfsdk:"source_snapshot"`

	State *string `json:"state,omitempty" tfsdk:"state"`

	Tenant *string `json:"tenant,omitempty" tfsdk:"tenant"`

	TenantUuid *string `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`

	Type *string `json:"type,omitempty" tfsdk:"type"`

	TypeName *string `json:"type_name,omitempty" tfsdk:"type_name"`

	Url *string `json:"url,omitempty" tfsdk:"url"`
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
