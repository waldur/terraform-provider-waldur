package flavor

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

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

func (m *OpenstackFlavorFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Openstack Flavor",
		Attributes: map[string]schema.Attribute{
			"cores": schema.Int64Attribute{
				Optional: true,
			},
			"cores__gte": schema.Int64Attribute{
				Optional: true,
			},
			"cores__lte": schema.Int64Attribute{
				Optional: true,
			},
			"disk": schema.Int64Attribute{
				Optional: true,
			},
			"disk__gte": schema.Int64Attribute{
				Optional: true,
			},
			"disk__lte": schema.Int64Attribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"name_exact": schema.StringAttribute{
				Optional: true,
			},
			"name_iregex": schema.StringAttribute{
				Optional: true,
			},
			"offering_uuid": schema.StringAttribute{
				Optional: true,
			},
			"ram": schema.Int64Attribute{
				Optional: true,
			},
			"ram__gte": schema.Int64Attribute{
				Optional: true,
			},
			"ram__lte": schema.Int64Attribute{
				Optional: true,
			},
			"settings": schema.StringAttribute{
				Optional: true,
			},
			"settings_uuid": schema.StringAttribute{
				Optional: true,
			},
			"tenant": schema.StringAttribute{
				Optional: true,
			},
			"tenant_uuid": schema.StringAttribute{
				Optional: true,
			},
		},
	}
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
	model.BackendId = common.StringPointerValue(apiResp.BackendId)
	model.Cores = types.Int64PointerValue(apiResp.Cores)
	model.Disk = types.Int64PointerValue(apiResp.Disk)
	model.DisplayName = common.StringPointerValue(apiResp.DisplayName)
	model.Name = common.StringPointerValue(apiResp.Name)
	model.Ram = types.Int64PointerValue(apiResp.Ram)
	model.Settings = common.StringPointerValue(apiResp.Settings)
	model.Url = common.StringPointerValue(apiResp.Url)

	return diags
}
