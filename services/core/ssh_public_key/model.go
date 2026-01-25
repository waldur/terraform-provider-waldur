package ssh_public_key

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// CoreSshPublicKeyFiltersModel contains the filter parameters for querying.
type CoreSshPublicKeyFiltersModel struct {
	Created           types.String `tfsdk:"created"`
	FingerprintMd5    types.String `tfsdk:"fingerprint_md5"`
	FingerprintSha256 types.String `tfsdk:"fingerprint_sha256"`
	FingerprintSha512 types.String `tfsdk:"fingerprint_sha512"`
	IsShared          types.Bool   `tfsdk:"is_shared"`
	Modified          types.String `tfsdk:"modified"`
	Name              types.String `tfsdk:"name"`
	NameExact         types.String `tfsdk:"name_exact"`
	UserUuid          types.String `tfsdk:"user_uuid"`
	Uuid              types.String `tfsdk:"uuid"`
}

type CoreSshPublicKeyModel struct {
	UUID              types.String `tfsdk:"id"`
	FingerprintMd5    types.String `tfsdk:"fingerprint_md5"`
	FingerprintSha256 types.String `tfsdk:"fingerprint_sha256"`
	FingerprintSha512 types.String `tfsdk:"fingerprint_sha512"`
	IsShared          types.Bool   `tfsdk:"is_shared"`
	Name              types.String `tfsdk:"name"`
	PublicKey         types.String `tfsdk:"public_key"`
	Type              types.String `tfsdk:"type"`
	Url               types.String `tfsdk:"url"`
	UserUuid          types.String `tfsdk:"user_uuid"`
}

// CopyFrom maps the API response to the model fields.
func (model *CoreSshPublicKeyModel) CopyFrom(ctx context.Context, apiResp CoreSshPublicKeyResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)

	return diags
}
