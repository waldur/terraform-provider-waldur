package bridge_link

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type IdentityBridgeLinkModel struct {
	UUID          types.String `tfsdk:"id"`
	IsCreated     types.Bool   `tfsdk:"is_created"`
	UpdatedFields types.List   `tfsdk:"updated_fields"`
	UserUuid      types.String `tfsdk:"user_uuid"`
}

// CopyFrom maps the API response to the model fields.
func (model *IdentityBridgeLinkModel) CopyFrom(ctx context.Context, apiResp IdentityBridgeLinkResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)

	model.IsCreated = types.BoolPointerValue(apiResp.IsCreated)

	if apiResp.UpdatedFields != nil {
		valUpdatedFields, diagsUpdatedFields := types.ListValueFrom(ctx, types.StringType, apiResp.UpdatedFields)
		diags.Append(diagsUpdatedFields...)
		model.UpdatedFields = valUpdatedFields
	} else {
		model.UpdatedFields = types.ListNull(types.StringType)
	}

	model.UserUuid = common.StringPointerValue(apiResp.UserUuid)

	return diags
}
