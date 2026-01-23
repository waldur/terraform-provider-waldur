package ssh_public_key

// CoreSshPublicKey Structs

type CoreSshPublicKeyCreateRequest struct {
}

type CoreSshPublicKeyResponse struct {
	UUID *string `json:"uuid"`

	FingerprintMd5    *string `json:"fingerprint_md5" tfsdk:"fingerprint_md5"`
	FingerprintSha256 *string `json:"fingerprint_sha256" tfsdk:"fingerprint_sha256"`
	FingerprintSha512 *string `json:"fingerprint_sha512" tfsdk:"fingerprint_sha512"`
	IsShared          *bool   `json:"is_shared" tfsdk:"is_shared"`
	Name              *string `json:"name" tfsdk:"name"`
	PublicKey         *string `json:"public_key" tfsdk:"public_key"`
	Type              *string `json:"type" tfsdk:"type"`
	Url               *string `json:"url" tfsdk:"url"`
	UserUuid          *string `json:"user_uuid" tfsdk:"user_uuid"`
}
