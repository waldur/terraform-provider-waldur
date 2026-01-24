package volume_type

// OpenstackVolumeType Structs

type OpenstackVolumeTypeCreateRequest struct {
}

type OpenstackVolumeTypeResponse struct {
	UUID *string `json:"uuid"`

	Description *string `json:"description" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
	Settings    *string `json:"settings" tfsdk:"settings"`
	Url         *string `json:"url" tfsdk:"url"`
}

func (r *OpenstackVolumeTypeResponse) GetState() string {
	return "OK"
}

func (r *OpenstackVolumeTypeResponse) GetErrorMessage() string {
	return ""
}
