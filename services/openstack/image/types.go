package image

// Shared Structs

// Resource Structs

// OpenstackImage Structs

// Create Request
type OpenstackImageCreateRequest struct {
}

// Update Request

// Update Actions Structs

// Response Struct
type OpenstackImageResponse struct {
	UUID *string `json:"uuid"`

	BackendId *string `json:"backend_id" tfsdk:"backend_id"`
	MinDisk   *int64  `json:"min_disk" tfsdk:"min_disk"`
	MinRam    *int64  `json:"min_ram" tfsdk:"min_ram"`
	Name      *string `json:"name" tfsdk:"name"`
	Settings  *string `json:"settings" tfsdk:"settings"`
	Url       *string `json:"url" tfsdk:"url"`
}
