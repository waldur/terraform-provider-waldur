package ssh_public_key

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

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

func (m *CoreSshPublicKeyFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Core Ssh Public Key",
		Attributes: map[string]schema.Attribute{
			"created": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Created after",
			},
			"fingerprint_md5": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"fingerprint_sha256": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"fingerprint_sha512": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"is_shared": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"modified": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Modified after",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name (exact)",
			},
			"user_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "User UUID",
			},
			"uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "UUID",
			},
		},
	}
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
	model.FingerprintMd5 = common.StringPointerValue(apiResp.FingerprintMd5)
	model.FingerprintSha256 = common.StringPointerValue(apiResp.FingerprintSha256)
	model.FingerprintSha512 = common.StringPointerValue(apiResp.FingerprintSha512)
	model.IsShared = types.BoolPointerValue(apiResp.IsShared)
	model.Name = common.StringPointerValue(apiResp.Name)
	model.PublicKey = common.StringPointerValue(apiResp.PublicKey)
	model.Type = common.StringPointerValue(apiResp.Type)
	model.Url = common.StringPointerValue(apiResp.Url)
	model.UserUuid = common.StringPointerValue(apiResp.UserUuid)

	return diags
}
