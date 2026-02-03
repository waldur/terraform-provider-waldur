package floating_ip

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

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
				Optional:            true,
				MarkdownDescription: "Address",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Backend ID",
			},
			"can_manage": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Can manage",
			},
			"customer": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer UUID",
			},
			"customer_abbreviation": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer abbreviation",
			},
			"customer_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer name",
			},
			"customer_native_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer native name",
			},
			"customer_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer UUID",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description",
			},
			"external_ip": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "External IP",
			},
			"free": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Is free",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name (exact)",
			},
			"project": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project UUID",
			},
			"project_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project name",
			},
			"project_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project UUID",
			},
			"runtime_state": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Runtime state",
			},
			"service_settings_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Service settings name",
			},
			"service_settings_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Service settings UUID",
			},
			"tenant": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Tenant URL",
			},
			"tenant_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Tenant UUID",
			},
			"uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "UUID",
			},
		},
	}
}

type OpenstackFloatingIpModel struct {
	UUID                        types.String      `tfsdk:"id"`
	AccessUrl                   types.String      `tfsdk:"access_url"`
	Address                     types.String      `tfsdk:"address"`
	BackendId                   types.String      `tfsdk:"backend_id"`
	BackendNetworkId            types.String      `tfsdk:"backend_network_id"`
	Created                     timetypes.RFC3339 `tfsdk:"created"`
	Customer                    types.String      `tfsdk:"customer"`
	CustomerAbbreviation        types.String      `tfsdk:"customer_abbreviation"`
	CustomerName                types.String      `tfsdk:"customer_name"`
	CustomerNativeName          types.String      `tfsdk:"customer_native_name"`
	CustomerUuid                types.String      `tfsdk:"customer_uuid"`
	Description                 types.String      `tfsdk:"description"`
	ErrorMessage                types.String      `tfsdk:"error_message"`
	ErrorTraceback              types.String      `tfsdk:"error_traceback"`
	ExternalAddress             types.String      `tfsdk:"external_address"`
	InstanceName                types.String      `tfsdk:"instance_name"`
	InstanceUrl                 types.String      `tfsdk:"instance_url"`
	InstanceUuid                types.String      `tfsdk:"instance_uuid"`
	IsLimitBased                types.Bool        `tfsdk:"is_limit_based"`
	IsUsageBased                types.Bool        `tfsdk:"is_usage_based"`
	MarketplaceCategoryName     types.String      `tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     types.String      `tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     types.String      `tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     types.String      `tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         types.String      `tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    types.String      `tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     types.String      `tfsdk:"marketplace_resource_uuid"`
	Modified                    timetypes.RFC3339 `tfsdk:"modified"`
	Name                        types.String      `tfsdk:"name"`
	Port                        types.String      `tfsdk:"port"`
	PortFixedIps                types.List        `tfsdk:"port_fixed_ips"`
	Project                     types.String      `tfsdk:"project"`
	ProjectName                 types.String      `tfsdk:"project_name"`
	ProjectUuid                 types.String      `tfsdk:"project_uuid"`
	ResourceType                types.String      `tfsdk:"resource_type"`
	Router                      types.String      `tfsdk:"router"`
	RuntimeState                types.String      `tfsdk:"runtime_state"`
	ServiceName                 types.String      `tfsdk:"service_name"`
	ServiceSettings             types.String      `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage types.String      `tfsdk:"service_settings_error_message"`
	ServiceSettingsState        types.String      `tfsdk:"service_settings_state"`
	ServiceSettingsUuid         types.String      `tfsdk:"service_settings_uuid"`
	State                       types.String      `tfsdk:"state"`
	Tenant                      types.String      `tfsdk:"tenant"`
	TenantName                  types.String      `tfsdk:"tenant_name"`
	TenantUuid                  types.String      `tfsdk:"tenant_uuid"`
	Url                         types.String      `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackFloatingIpModel) CopyFrom(ctx context.Context, apiResp OpenstackFloatingIpResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = common.StringPointerValue(apiResp.AccessUrl)
	model.Address = common.StringPointerValue(apiResp.Address)
	model.BackendId = common.StringPointerValue(apiResp.BackendId)
	model.BackendNetworkId = common.StringPointerValue(apiResp.BackendNetworkId)
	valCreated, diagsCreated := timetypes.NewRFC3339PointerValue(apiResp.Created)
	diags.Append(diagsCreated...)
	model.Created = valCreated
	model.Customer = common.StringPointerValue(apiResp.Customer)
	model.CustomerAbbreviation = common.StringPointerValue(apiResp.CustomerAbbreviation)
	model.CustomerName = common.StringPointerValue(apiResp.CustomerName)
	model.CustomerNativeName = common.StringPointerValue(apiResp.CustomerNativeName)
	model.CustomerUuid = common.StringPointerValue(apiResp.CustomerUuid)
	model.Description = common.StringPointerValue(apiResp.Description)
	model.ErrorMessage = common.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = common.StringPointerValue(apiResp.ErrorTraceback)
	model.ExternalAddress = common.StringPointerValue(apiResp.ExternalAddress)
	model.InstanceName = common.StringPointerValue(apiResp.InstanceName)
	model.InstanceUrl = common.StringPointerValue(apiResp.InstanceUrl)
	model.InstanceUuid = common.StringPointerValue(apiResp.InstanceUuid)
	model.IsLimitBased = types.BoolPointerValue(apiResp.IsLimitBased)
	model.IsUsageBased = types.BoolPointerValue(apiResp.IsUsageBased)
	model.MarketplaceCategoryName = common.StringPointerValue(apiResp.MarketplaceCategoryName)
	model.MarketplaceCategoryUuid = common.StringPointerValue(apiResp.MarketplaceCategoryUuid)
	model.MarketplaceOfferingName = common.StringPointerValue(apiResp.MarketplaceOfferingName)
	model.MarketplaceOfferingUuid = common.StringPointerValue(apiResp.MarketplaceOfferingUuid)
	model.MarketplacePlanUuid = common.StringPointerValue(apiResp.MarketplacePlanUuid)
	model.MarketplaceResourceState = common.StringPointerValue(apiResp.MarketplaceResourceState)
	model.MarketplaceResourceUuid = common.StringPointerValue(apiResp.MarketplaceResourceUuid)
	valModified, diagsModified := timetypes.NewRFC3339PointerValue(apiResp.Modified)
	diags.Append(diagsModified...)
	model.Modified = valModified
	model.Name = common.StringPointerValue(apiResp.Name)
	model.Port = common.StringPointerValue(apiResp.Port)

	{
		listValPortFixedIps, listDiagsPortFixedIps := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"ip_address": types.StringType,
			"subnet_id":  types.StringType,
		}}, apiResp.PortFixedIps)
		diags.Append(listDiagsPortFixedIps...)
		model.PortFixedIps = listValPortFixedIps
	}
	model.Project = common.StringPointerValue(apiResp.Project)
	model.ProjectName = common.StringPointerValue(apiResp.ProjectName)
	model.ProjectUuid = common.StringPointerValue(apiResp.ProjectUuid)
	model.ResourceType = common.StringPointerValue(apiResp.ResourceType)
	model.RuntimeState = common.StringPointerValue(apiResp.RuntimeState)
	model.ServiceName = common.StringPointerValue(apiResp.ServiceName)
	model.ServiceSettings = common.StringPointerValue(apiResp.ServiceSettings)
	model.ServiceSettingsErrorMessage = common.StringPointerValue(apiResp.ServiceSettingsErrorMessage)
	model.ServiceSettingsState = common.StringPointerValue(apiResp.ServiceSettingsState)
	model.ServiceSettingsUuid = common.StringPointerValue(apiResp.ServiceSettingsUuid)
	model.State = common.StringPointerValue(apiResp.State)
	model.Tenant = common.StringPointerValue(apiResp.Tenant)
	model.TenantName = common.StringPointerValue(apiResp.TenantName)
	model.TenantUuid = common.StringPointerValue(apiResp.TenantUuid)
	model.Url = common.StringPointerValue(apiResp.Url)

	return diags
}
