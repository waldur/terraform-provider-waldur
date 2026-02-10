package server_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

func OpenStackNestedInstanceType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"backend_id": types.StringType,
		"name":       types.StringType,
		"uuid":       types.StringType,
	}}
}

type OpenstackServerGroupFiltersModel struct {
	BackendId            types.String `tfsdk:"backend_id"`
	CanManage            types.Bool   `tfsdk:"can_manage"`
	Customer             types.String `tfsdk:"customer"`
	CustomerAbbreviation types.String `tfsdk:"customer_abbreviation"`
	CustomerName         types.String `tfsdk:"customer_name"`
	CustomerNativeName   types.String `tfsdk:"customer_native_name"`
	CustomerUuid         types.String `tfsdk:"customer_uuid"`
	Description          types.String `tfsdk:"description"`
	ExternalIp           types.String `tfsdk:"external_ip"`
	Name                 types.String `tfsdk:"name"`
	NameExact            types.String `tfsdk:"name_exact"`
	Project              types.String `tfsdk:"project"`
	ProjectName          types.String `tfsdk:"project_name"`
	ProjectUuid          types.String `tfsdk:"project_uuid"`
	ServiceSettingsName  types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid  types.String `tfsdk:"service_settings_uuid"`
	Tenant               types.String `tfsdk:"tenant"`
	TenantUuid           types.String `tfsdk:"tenant_uuid"`
	Uuid                 types.String `tfsdk:"uuid"`
}

func (m *OpenstackServerGroupFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Openstack Server Group",
		Attributes: map[string]schema.Attribute{
			"backend_id": schema.StringAttribute{
				Optional: true,
			},
			"can_manage": schema.BoolAttribute{
				Optional: true,
			},
			"customer": schema.StringAttribute{
				Optional: true,
			},
			"customer_abbreviation": schema.StringAttribute{
				Optional: true,
			},
			"customer_name": schema.StringAttribute{
				Optional: true,
			},
			"customer_native_name": schema.StringAttribute{
				Optional: true,
			},
			"customer_uuid": schema.StringAttribute{
				Optional: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"external_ip": schema.StringAttribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"name_exact": schema.StringAttribute{
				Optional: true,
			},
			"project": schema.StringAttribute{
				Optional: true,
			},
			"project_name": schema.StringAttribute{
				Optional: true,
			},
			"project_uuid": schema.StringAttribute{
				Optional: true,
			},
			"service_settings_name": schema.StringAttribute{
				Optional: true,
			},
			"service_settings_uuid": schema.StringAttribute{
				Optional: true,
			},
			"tenant": schema.StringAttribute{
				Optional: true,
			},
			"tenant_uuid": schema.StringAttribute{
				Optional: true,
			},
			"uuid": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

type OpenstackServerGroupModel struct {
	UUID                    types.String `tfsdk:"id"`
	BackendId               types.String `tfsdk:"backend_id"`
	Customer                types.String `tfsdk:"customer"`
	Description             types.String `tfsdk:"description"`
	DisplayName             types.String `tfsdk:"display_name"`
	ErrorMessage            types.String `tfsdk:"error_message"`
	ErrorTraceback          types.String `tfsdk:"error_traceback"`
	Instances               types.List   `tfsdk:"instances"`
	MarketplaceResourceUuid types.String `tfsdk:"marketplace_resource_uuid"`
	Name                    types.String `tfsdk:"name"`
	Policy                  types.String `tfsdk:"policy"`
	Project                 types.String `tfsdk:"project"`
	ResourceType            types.String `tfsdk:"resource_type"`
	State                   types.String `tfsdk:"state"`
	Tenant                  types.String `tfsdk:"tenant"`
	TenantName              types.String `tfsdk:"tenant_name"`
	TenantUuid              types.String `tfsdk:"tenant_uuid"`
	Url                     types.String `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackServerGroupModel) CopyFrom(ctx context.Context, apiResp OpenstackServerGroupResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.BackendId = common.StringPointerValue(apiResp.BackendId)
	model.Customer = common.StringPointerValue(apiResp.Customer)
	model.Description = common.StringPointerValue(apiResp.Description)
	model.DisplayName = common.StringPointerValue(apiResp.DisplayName)
	model.ErrorMessage = common.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = common.StringPointerValue(apiResp.ErrorTraceback)

	if apiResp.Instances != nil && len(*apiResp.Instances) > 0 {
		listValInstances, listDiagsInstances := types.ListValueFrom(ctx, OpenStackNestedInstanceType(), apiResp.Instances)
		diags.Append(listDiagsInstances...)
		model.Instances = listValInstances
	} else {
		model.Instances = types.ListNull(OpenStackNestedInstanceType())
	}
	model.MarketplaceResourceUuid = common.StringPointerValue(apiResp.MarketplaceResourceUuid)
	model.Name = common.StringPointerValue(apiResp.Name)
	model.Policy = common.StringPointerValue(apiResp.Policy)
	model.Project = common.StringPointerValue(apiResp.Project)
	model.ResourceType = common.StringPointerValue(apiResp.ResourceType)
	model.State = common.StringPointerValue(apiResp.State)
	model.Tenant = common.StringPointerValue(apiResp.Tenant)
	model.TenantName = common.StringPointerValue(apiResp.TenantName)
	model.TenantUuid = common.StringPointerValue(apiResp.TenantUuid)
	model.Url = common.StringPointerValue(apiResp.Url)

	return diags
}
