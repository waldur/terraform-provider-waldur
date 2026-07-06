package permission

type CustomerPermissionCreateRequest struct {
	Customer *string `json:"customer" tfsdk:"customer"`

	ExpirationTime *string `json:"expiration_time,omitempty" tfsdk:"expiration_time"`

	Role *string `json:"role" tfsdk:"role"`

	User *string `json:"user" tfsdk:"user"`
}

type CustomerPermissionUpdateRequest struct {
	Customer *string `json:"customer" tfsdk:"customer"`

	ExpirationTime *string `json:"expiration_time,omitempty" tfsdk:"expiration_time"`

	Role *string `json:"role" tfsdk:"role"`

	User *string `json:"user" tfsdk:"user"`
}

type CustomerPermissionResponse struct {
	UUID *string `json:"uuid"`

	Customer *string `json:"customer" tfsdk:"customer"`

	ExpirationTime *string `json:"expiration_time,omitempty" tfsdk:"expiration_time"`

	Role *string `json:"role" tfsdk:"role"`

	User *string `json:"user" tfsdk:"user"`
}

func (r *CustomerPermissionResponse) GetState() string {
	return "OK"
}

func (r *CustomerPermissionResponse) GetErrorMessage() string {
	return ""
}
