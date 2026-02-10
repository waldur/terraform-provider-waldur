package floating_ip

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

func OpenStackFixedIpType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"ip_address": types.StringType,
		"subnet_id":  types.StringType,
	}}
}

type OpenstackFloatingIpFiltersModel struct {
	Address              types.String `tfsdk:"address"`
	BackendId            types.String `tfsdk:"backend_id"`
	CanManage            types.Bool   `tfsdk:"can_manage"`
	Customer             types.String `tfsdk:"customer"`
	CustomerAbbreviation types.String `tfsdk:"customer_abbreviation"`
	CustomerName         types.String `tfsdk:"customer_name"`
	CustomerNativeName   types.String `tfsdk:"customer_native_name"`
	CustomerUuid         types.String `tfsdk:"customer_uuid"`
	Description          types.String `tfsdk:"description"`
	ExternalIp           types.String `tfsdk:"external_ip"`
	Free                 types.Bool   `tfsdk:"free"`
	Name                 types.String `tfsdk:"name"`
	NameExact            types.String `tfsdk:"name_exact"`
	Project              types.String `tfsdk:"project"`
	ProjectName          types.String `tfsdk:"project_name"`
	ProjectUuid          types.String `tfsdk:"project_uuid"`
	RuntimeState         types.String `tfsdk:"runtime_state"`
	ServiceSettingsName  types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid  types.String `tfsdk:"service_settings_uuid"`
	Tenant               types.String `tfsdk:"tenant"`
	TenantUuid           types.String `tfsdk:"tenant_uuid"`
	Uuid                 types.String `tfsdk:"uuid"`
}

func (m *OpenstackFloatingIpFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Openstack Floating Ip",
		Attributes: map[string]schema.Attribute{
			"address": schema.StringAttribute{
				Optional: true,
			},
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
			"free": schema.BoolAttribute{
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
			"runtime_state": schema.StringAttribute{
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

type OpenstackFloatingIpModel struct {
	UUID                    types.String `tfsdk:"id"`
	Address                 types.String `tfsdk:"address"`
	BackendId               types.String `tfsdk:"backend_id"`
	BackendNetworkId        types.String `tfsdk:"backend_network_id"`
	Customer                types.String `tfsdk:"customer"`
	Description             types.String `tfsdk:"description"`
	ErrorMessage            types.String `tfsdk:"error_message"`
	ErrorTraceback          types.String `tfsdk:"error_traceback"`
	ExternalAddress         types.String `tfsdk:"external_address"`
	InstanceName            types.String `tfsdk:"instance_name"`
	InstanceUrl             types.String `tfsdk:"instance_url"`
	InstanceUuid            types.String `tfsdk:"instance_uuid"`
	MarketplaceResourceUuid types.String `tfsdk:"marketplace_resource_uuid"`
	Name                    types.String `tfsdk:"name"`
	Port                    types.String `tfsdk:"port"`
	PortFixedIps            types.List   `tfsdk:"port_fixed_ips"`
	Project                 types.String `tfsdk:"project"`
	ResourceType            types.String `tfsdk:"resource_type"`
	RuntimeState            types.String `tfsdk:"runtime_state"`
	State                   types.String `tfsdk:"state"`
	Tenant                  types.String `tfsdk:"tenant"`
	TenantName              types.String `tfsdk:"tenant_name"`
	TenantUuid              types.String `tfsdk:"tenant_uuid"`
	Url                     types.String `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackFloatingIpModel) CopyFrom(ctx context.Context, apiResp OpenstackFloatingIpResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.Address = common.StringPointerValue(apiResp.Address)
	model.BackendId = common.StringPointerValue(apiResp.BackendId)
	model.BackendNetworkId = common.StringPointerValue(apiResp.BackendNetworkId)
	model.Customer = common.StringPointerValue(apiResp.Customer)
	model.Description = common.StringPointerValue(apiResp.Description)
	model.ErrorMessage = common.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = common.StringPointerValue(apiResp.ErrorTraceback)
	model.ExternalAddress = common.StringPointerValue(apiResp.ExternalAddress)
	model.InstanceName = common.StringPointerValue(apiResp.InstanceName)
	model.InstanceUrl = common.StringPointerValue(apiResp.InstanceUrl)
	model.InstanceUuid = common.StringPointerValue(apiResp.InstanceUuid)
	model.MarketplaceResourceUuid = common.StringPointerValue(apiResp.MarketplaceResourceUuid)
	model.Name = common.StringPointerValue(apiResp.Name)
	model.Port = common.StringPointerValue(apiResp.Port)

	if apiResp.PortFixedIps != nil && len(*apiResp.PortFixedIps) > 0 {
		listValPortFixedIps, listDiagsPortFixedIps := types.ListValueFrom(ctx, OpenStackFixedIpType(), apiResp.PortFixedIps)
		diags.Append(listDiagsPortFixedIps...)
		model.PortFixedIps = listValPortFixedIps
	} else {
		model.PortFixedIps = types.ListNull(OpenStackFixedIpType())
	}
	model.Project = common.StringPointerValue(apiResp.Project)
	model.ResourceType = common.StringPointerValue(apiResp.ResourceType)
	model.RuntimeState = common.StringPointerValue(apiResp.RuntimeState)
	model.State = common.StringPointerValue(apiResp.State)
	model.Tenant = common.StringPointerValue(apiResp.Tenant)
	model.TenantName = common.StringPointerValue(apiResp.TenantName)
	model.TenantUuid = common.StringPointerValue(apiResp.TenantUuid)
	model.Url = common.StringPointerValue(apiResp.Url)

	return diags
}
