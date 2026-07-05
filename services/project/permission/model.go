package permission

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type ProjectPermissionModel struct {
	UUID           types.String `tfsdk:"id"`
	ExpirationTime types.String `tfsdk:"expiration_time"`
	Project        types.String `tfsdk:"project"`
	Role           types.String `tfsdk:"role"`
	User           types.String `tfsdk:"user"`
}

// CopyFrom maps the API response to the model fields.
func (model *ProjectPermissionModel) CopyFrom(ctx context.Context, apiResp ProjectPermissionResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)

	model.ExpirationTime = common.StringPointerValue(apiResp.ExpirationTime)

	model.Project = common.StringPointerValue(apiResp.Project)

	model.Role = common.StringPointerValue(apiResp.Role)

	model.User = common.StringPointerValue(apiResp.User)

	return diags
}
