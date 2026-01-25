package image

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// OpenstackImageFiltersModel contains the filter parameters for querying.
type OpenstackImageFiltersModel struct {
	Name         types.String `tfsdk:"name"`
	NameExact    types.String `tfsdk:"name_exact"`
	OfferingUuid types.String `tfsdk:"offering_uuid"`
	Settings     types.String `tfsdk:"settings"`
	SettingsUuid types.String `tfsdk:"settings_uuid"`
	Tenant       types.String `tfsdk:"tenant"`
	TenantUuid   types.String `tfsdk:"tenant_uuid"`
}

type OpenstackImageModel struct {
	UUID      types.String `tfsdk:"id"`
	BackendId types.String `tfsdk:"backend_id"`
	MinDisk   types.Int64  `tfsdk:"min_disk"`
	MinRam    types.Int64  `tfsdk:"min_ram"`
	Name      types.String `tfsdk:"name"`
	Settings  types.String `tfsdk:"settings"`
	Url       types.String `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackImageModel) CopyFrom(ctx context.Context, apiResp OpenstackImageResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)

	return diags
}
