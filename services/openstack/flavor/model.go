package flavor

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// OpenstackFlavorFiltersModel contains the filter parameters for querying.
type OpenstackFlavorFiltersModel struct {
	Cores        types.Int64  `tfsdk:"cores"`
	CoresGte     types.Int64  `tfsdk:"cores__gte"`
	CoresLte     types.Int64  `tfsdk:"cores__lte"`
	Disk         types.Int64  `tfsdk:"disk"`
	DiskGte      types.Int64  `tfsdk:"disk__gte"`
	DiskLte      types.Int64  `tfsdk:"disk__lte"`
	Name         types.String `tfsdk:"name"`
	NameExact    types.String `tfsdk:"name_exact"`
	NameIregex   types.String `tfsdk:"name_iregex"`
	OfferingUuid types.String `tfsdk:"offering_uuid"`
	Ram          types.Int64  `tfsdk:"ram"`
	RamGte       types.Int64  `tfsdk:"ram__gte"`
	RamLte       types.Int64  `tfsdk:"ram__lte"`
	Settings     types.String `tfsdk:"settings"`
	SettingsUuid types.String `tfsdk:"settings_uuid"`
	Tenant       types.String `tfsdk:"tenant"`
	TenantUuid   types.String `tfsdk:"tenant_uuid"`
}

type OpenstackFlavorModel struct {
	UUID        types.String `tfsdk:"id"`
	BackendId   types.String `tfsdk:"backend_id"`
	Cores       types.Int64  `tfsdk:"cores"`
	Disk        types.Int64  `tfsdk:"disk"`
	DisplayName types.String `tfsdk:"display_name"`
	Name        types.String `tfsdk:"name"`
	Ram         types.Int64  `tfsdk:"ram"`
	Settings    types.String `tfsdk:"settings"`
	Url         types.String `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackFlavorModel) CopyFrom(ctx context.Context, apiResp OpenstackFlavorResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)

	return diags
}
