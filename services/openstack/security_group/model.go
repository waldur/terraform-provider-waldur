package security_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

func OpenStackSecurityGroupRuleCreateRequestType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"cidr":              types.StringType,
		"description":       types.StringType,
		"direction":         types.StringType,
		"ethertype":         types.StringType,
		"from_port":         types.Int64Type,
		"protocol":          types.StringType,
		"remote_group":      types.StringType,
		"to_port":           types.Int64Type,
		"id":                types.Int64Type,
		"remote_group_name": types.StringType,
		"remote_group_uuid": types.StringType,
	}}
}

type OpenstackSecurityGroupFiltersModel struct {
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
	Query                types.String `tfsdk:"query"`
	ServiceSettingsName  types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid  types.String `tfsdk:"service_settings_uuid"`
	Tenant               types.String `tfsdk:"tenant"`
	TenantUuid           types.String `tfsdk:"tenant_uuid"`
	Uuid                 types.String `tfsdk:"uuid"`
}

func (m *OpenstackSecurityGroupFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Openstack Security Group",
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
			"query": schema.StringAttribute{
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

type OpenstackSecurityGroupModel struct {
	UUID                    types.String `tfsdk:"id"`
	BackendId               types.String `tfsdk:"backend_id"`
	Customer                types.String `tfsdk:"customer"`
	Description             types.String `tfsdk:"description"`
	ErrorMessage            types.String `tfsdk:"error_message"`
	ErrorTraceback          types.String `tfsdk:"error_traceback"`
	MarketplaceResourceUuid types.String `tfsdk:"marketplace_resource_uuid"`
	Name                    types.String `tfsdk:"name"`
	Project                 types.String `tfsdk:"project"`
	ResourceType            types.String `tfsdk:"resource_type"`
	Rules                   types.List   `tfsdk:"rules"`
	State                   types.String `tfsdk:"state"`
	Tenant                  types.String `tfsdk:"tenant"`
	TenantName              types.String `tfsdk:"tenant_name"`
	TenantUuid              types.String `tfsdk:"tenant_uuid"`
	Url                     types.String `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackSecurityGroupModel) CopyFrom(ctx context.Context, apiResp OpenstackSecurityGroupResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.BackendId = common.StringPointerValue(apiResp.BackendId)
	model.Customer = common.StringPointerValue(apiResp.Customer)
	model.Description = common.StringPointerValue(apiResp.Description)
	model.ErrorMessage = common.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = common.StringPointerValue(apiResp.ErrorTraceback)
	model.MarketplaceResourceUuid = common.StringPointerValue(apiResp.MarketplaceResourceUuid)
	model.Name = common.StringPointerValue(apiResp.Name)
	model.Project = common.StringPointerValue(apiResp.Project)
	model.ResourceType = common.StringPointerValue(apiResp.ResourceType)

	if len(apiResp.Rules) > 0 {
		listValRules, listDiagsRules := types.ListValueFrom(ctx, OpenStackSecurityGroupRuleCreateRequestType(), apiResp.Rules)
		diags.Append(listDiagsRules...)
		model.Rules = listValRules
	} else {
		model.Rules = types.ListNull(OpenStackSecurityGroupRuleCreateRequestType())
	}
	model.State = common.StringPointerValue(apiResp.State)
	model.Tenant = common.StringPointerValue(apiResp.Tenant)
	model.TenantName = common.StringPointerValue(apiResp.TenantName)
	model.TenantUuid = common.StringPointerValue(apiResp.TenantUuid)
	model.Url = common.StringPointerValue(apiResp.Url)

	return diags
}
