package image

type OpenstackImageCreateRequest struct {
}

type OpenstackImageResponse struct {
	UUID *string `json:"uuid"`

	BackendCreatedAt *string `json:"backend_created_at,omitempty" tfsdk:"backend_created_at"`

	BackendId *string `json:"backend_id" tfsdk:"backend_id"`

	HwRescueBus *string `json:"hw_rescue_bus,omitempty" tfsdk:"hw_rescue_bus"`

	HwRescueDevice *string `json:"hw_rescue_device,omitempty" tfsdk:"hw_rescue_device"`

	IsRescueImage *bool `json:"is_rescue_image" tfsdk:"is_rescue_image"`

	MinDisk *int64 `json:"min_disk,omitempty" tfsdk:"min_disk"`

	MinRam *int64 `json:"min_ram,omitempty" tfsdk:"min_ram"`

	Name *string `json:"name" tfsdk:"name"`

	Settings *string `json:"settings" tfsdk:"settings"`

	Url *string `json:"url" tfsdk:"url"`
}

func (r *OpenstackImageResponse) GetState() string {
	return "OK"
}

func (r *OpenstackImageResponse) GetErrorMessage() string {
	return ""
}
