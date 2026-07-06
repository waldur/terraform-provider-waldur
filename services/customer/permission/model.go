package permission

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type CustomerPermissionModel struct {
	UUID           types.String `tfsdk:"id"`
	Customer       types.String `tfsdk:"customer"`
	ExpirationTime types.String `tfsdk:"expiration_time"`
	Role           types.String `tfsdk:"role"`
	User           types.String `tfsdk:"user"`
}

// CopyFrom maps the API response to the model fields.
func (model *CustomerPermissionModel) CopyFrom(ctx context.Context, apiResp CustomerPermissionResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)

	model.Customer = common.StringPointerValue(apiResp.Customer)

	model.ExpirationTime = common.StringPointerValue(apiResp.ExpirationTime)

	model.Role = common.StringPointerValue(apiResp.Role)

	model.User = common.StringPointerValue(apiResp.User)

	return diags
}
