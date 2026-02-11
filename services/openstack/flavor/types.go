package flavor

type OpenstackFlavorCreateRequest struct {
}

type OpenstackFlavorResponse struct {
	UUID *string `json:"uuid"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Cores *int64 `json:"cores,omitempty" tfsdk:"cores"`

	Disk *int64 `json:"disk,omitempty" tfsdk:"disk"`

	DisplayName *string `json:"display_name,omitempty" tfsdk:"display_name"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Ram *int64 `json:"ram,omitempty" tfsdk:"ram"`

	Settings *string `json:"settings,omitempty" tfsdk:"settings"`

	Url *string `json:"url,omitempty" tfsdk:"url"`
}

func (r *OpenstackFlavorResponse) GetState() string {
	return "OK"
}

func (r *OpenstackFlavorResponse) GetErrorMessage() string {
	return ""
}
