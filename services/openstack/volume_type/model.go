package volume_type

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OpenstackVolumeTypeFiltersModel struct {
	Name         types.String `tfsdk:"name"`
	NameExact    types.String `tfsdk:"name_exact"`
	OfferingUuid types.String `tfsdk:"offering_uuid"`
	Settings     types.String `tfsdk:"settings"`
	SettingsUuid types.String `tfsdk:"settings_uuid"`
	Tenant       types.String `tfsdk:"tenant"`
	TenantUuid   types.String `tfsdk:"tenant_uuid"`
}

type OpenstackVolumeTypeModel struct {
	UUID        types.String `tfsdk:"id"`
	Description types.String `tfsdk:"description"`
	Name        types.String `tfsdk:"name"`
	Settings    types.String `tfsdk:"settings"`
	Url         types.String `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackVolumeTypeModel) CopyFrom(ctx context.Context, apiResp OpenstackVolumeTypeResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)

	return diags
}
