package flavor

type OpenstackFlavorCreateRequest struct {
}

type OpenstackFlavorResponse struct {
	UUID *string `json:"uuid"`

	BackendId   *string `json:"backend_id" tfsdk:"backend_id"`
	Cores       *int64  `json:"cores" tfsdk:"cores"`
	Disk        *int64  `json:"disk" tfsdk:"disk"`
	DisplayName *string `json:"display_name" tfsdk:"display_name"`
	Name        *string `json:"name" tfsdk:"name"`
	Ram         *int64  `json:"ram" tfsdk:"ram"`
	Settings    *string `json:"settings" tfsdk:"settings"`
	Url         *string `json:"url" tfsdk:"url"`
}

func (r *OpenstackFlavorResponse) GetState() string {
	return "OK"
}

func (r *OpenstackFlavorResponse) GetErrorMessage() string {
	return ""
}
