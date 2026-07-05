package permission

type ProjectPermissionCreateRequest struct {
	ExpirationTime *string `json:"expiration_time,omitempty" tfsdk:"expiration_time"`

	Project *string `json:"project" tfsdk:"project"`

	Role *string `json:"role" tfsdk:"role"`

	User *string `json:"user" tfsdk:"user"`
}

type ProjectPermissionUpdateRequest struct {
	ExpirationTime *string `json:"expiration_time,omitempty" tfsdk:"expiration_time"`

	Project *string `json:"project" tfsdk:"project"`

	Role *string `json:"role" tfsdk:"role"`

	User *string `json:"user" tfsdk:"user"`
}

type ProjectPermissionResponse struct {
	UUID *string `json:"uuid"`

	ExpirationTime *string `json:"expiration_time,omitempty" tfsdk:"expiration_time"`

	Project *string `json:"project" tfsdk:"project"`

	Role *string `json:"role" tfsdk:"role"`

	User *string `json:"user" tfsdk:"user"`
}

func (r *ProjectPermissionResponse) GetState() string {
	return "OK"
}

func (r *ProjectPermissionResponse) GetErrorMessage() string {
	return ""
}
