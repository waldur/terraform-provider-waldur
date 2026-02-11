package ssh_public_key

type CoreSshPublicKeyCreateRequest struct {
}

type CoreSshPublicKeyResponse struct {
	UUID *string `json:"uuid"`

	FingerprintMd5 *string `json:"fingerprint_md5,omitempty" tfsdk:"fingerprint_md5"`

	FingerprintSha256 *string `json:"fingerprint_sha256,omitempty" tfsdk:"fingerprint_sha256"`

	FingerprintSha512 *string `json:"fingerprint_sha512,omitempty" tfsdk:"fingerprint_sha512"`

	IsShared *bool `json:"is_shared,omitempty" tfsdk:"is_shared"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	PublicKey *string `json:"public_key,omitempty" tfsdk:"public_key"`

	Type *string `json:"type,omitempty" tfsdk:"type"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	UserUuid *string `json:"user_uuid,omitempty" tfsdk:"user_uuid"`
}

func (r *CoreSshPublicKeyResponse) GetState() string {
	return "OK"
}

func (r *CoreSshPublicKeyResponse) GetErrorMessage() string {
	return ""
}
