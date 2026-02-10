package volume_type

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
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

func (m *OpenstackVolumeTypeFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Openstack Volume Type",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Optional: true,
			},
			"name_exact": schema.StringAttribute{
				Optional: true,
			},
			"offering_uuid": schema.StringAttribute{
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
	model.Description = common.StringPointerValue(apiResp.Description)
	model.Name = common.StringPointerValue(apiResp.Name)
	model.Settings = common.StringPointerValue(apiResp.Settings)
	model.Url = common.StringPointerValue(apiResp.Url)

	return diags
}
