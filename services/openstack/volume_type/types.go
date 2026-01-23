package volume_type

// Shared Structs

// Resource Structs

// OpenstackVolumeType Structs

// Create Request
type OpenstackVolumeTypeCreateRequest struct {
}

// Update Request

// Update Actions Structs

// Response Struct
type OpenstackVolumeTypeResponse struct {
	UUID *string `json:"uuid"`

	Description *string `json:"description" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
	Settings    *string `json:"settings" tfsdk:"settings"`
	Url         *string `json:"url" tfsdk:"url"`
}
