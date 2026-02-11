package volume_attachment

type OpenstackVolumeAttachmentCreateRequest struct {
	Device *string `json:"device,omitempty" tfsdk:"device"`

	Instance *string `json:"instance" tfsdk:"instance"`

	Volume *string `json:"volume" tfsdk:"volume"`
}

type OpenstackVolumeAttachmentResponse struct {
	UUID *string `json:"uuid"`

	Action *string `json:"action,omitempty" tfsdk:"action"`

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

	Instance *string `json:"instance" tfsdk:"instance"`

	InstanceMarketplaceUuid *string `json:"instance_marketplace_uuid,omitempty" tfsdk:"instance_marketplace_uuid"`

	InstanceName *string `json:"instance_name,omitempty" tfsdk:"instance_name"`

	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Project *string `json:"project,omitempty" tfsdk:"project"`

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

func (r *OpenstackVolumeAttachmentResponse) GetState() string {
	if r.State != nil {
		return *r.State
	}
	return "OK"
}

func (r *OpenstackVolumeAttachmentResponse) GetErrorMessage() string {
	if r.ErrorMessage != nil {
		return *r.ErrorMessage
	}
	return ""
}
