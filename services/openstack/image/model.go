package image

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackImageFiltersModel struct {
	Name               types.String `tfsdk:"name"`
	NameExact          types.String `tfsdk:"name_exact"`
	OfferingUuid       types.String `tfsdk:"offering_uuid"`
	Settings           types.String `tfsdk:"settings"`
	SettingsUuid       types.String `tfsdk:"settings_uuid"`
	ShowDuplicateNames types.Bool   `tfsdk:"show_duplicate_names"`
	Tenant             types.String `tfsdk:"tenant"`
	TenantUuid         types.String `tfsdk:"tenant_uuid"`
}

func (m *OpenstackImageFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Openstack Image",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name (exact)",
			},
			"offering_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Offering UUID",
			},
			"settings": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Settings URL",
			},
			"settings_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Settings UUID",
			},
			"show_duplicate_names": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Show duplicate image names",
			},
			"tenant": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Tenant URL",
			},
			"tenant_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Tenant UUID",
			},
		},
	}
}

type OpenstackImageModel struct {
	UUID             types.String      `tfsdk:"id"`
	BackendCreatedAt timetypes.RFC3339 `tfsdk:"backend_created_at"`
	BackendId        types.String      `tfsdk:"backend_id"`
	MinDisk          types.Int64       `tfsdk:"min_disk"`
	MinRam           types.Int64       `tfsdk:"min_ram"`
	Name             types.String      `tfsdk:"name"`
	Settings         types.String      `tfsdk:"settings"`
	Url              types.String      `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackImageModel) CopyFrom(ctx context.Context, apiResp OpenstackImageResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	valBackendCreatedAt, diagsBackendCreatedAt := timetypes.NewRFC3339PointerValue(apiResp.BackendCreatedAt)
	diags.Append(diagsBackendCreatedAt...)
	model.BackendCreatedAt = valBackendCreatedAt
	model.BackendId = common.StringPointerValue(apiResp.BackendId)
	model.MinDisk = types.Int64PointerValue(apiResp.MinDisk)
	model.MinRam = types.Int64PointerValue(apiResp.MinRam)
	model.Name = common.StringPointerValue(apiResp.Name)
	model.Settings = common.StringPointerValue(apiResp.Settings)
	model.Url = common.StringPointerValue(apiResp.Url)

	return diags
}
