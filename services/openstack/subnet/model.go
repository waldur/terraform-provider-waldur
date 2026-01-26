package subnet

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OpenstackSubnetFiltersModel struct {
	BackendId            types.String `tfsdk:"backend_id"`
	CanManage            types.Bool   `tfsdk:"can_manage"`
	Customer             types.String `tfsdk:"customer"`
	CustomerAbbreviation types.String `tfsdk:"customer_abbreviation"`
	CustomerName         types.String `tfsdk:"customer_name"`
	CustomerNativeName   types.String `tfsdk:"customer_native_name"`
	CustomerUuid         types.String `tfsdk:"customer_uuid"`
	Description          types.String `tfsdk:"description"`
	DirectOnly           types.Bool   `tfsdk:"direct_only"`
	EnableDhcp           types.Bool   `tfsdk:"enable_dhcp"`
	ExternalIp           types.String `tfsdk:"external_ip"`
	IpVersion            types.Int64  `tfsdk:"ip_version"`
	Name                 types.String `tfsdk:"name"`
	NameExact            types.String `tfsdk:"name_exact"`
	Network              types.String `tfsdk:"network"`
	NetworkUuid          types.String `tfsdk:"network_uuid"`
	Project              types.String `tfsdk:"project"`
	ProjectName          types.String `tfsdk:"project_name"`
	ProjectUuid          types.String `tfsdk:"project_uuid"`
	RbacOnly             types.Bool   `tfsdk:"rbac_only"`
	ServiceSettingsName  types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid  types.String `tfsdk:"service_settings_uuid"`
	Tenant               types.String `tfsdk:"tenant"`
	TenantUuid           types.String `tfsdk:"tenant_uuid"`
	Uuid                 types.String `tfsdk:"uuid"`
}

func (m *OpenstackSubnetFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Openstack Subnet",
		Attributes: map[string]schema.Attribute{
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
			"direct_only": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Direct only",
			},
			"enable_dhcp": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Enable dhcp",
			},
			"external_ip": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "External IP",
			},
			"ip_version": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Ip version",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name (exact)",
			},
			"network": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Network URL",
			},
			"network_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Network UUID",
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
			"rbac_only": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "RBAC only",
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

type OpenstackSubnetModel struct {
	UUID                        types.String      `tfsdk:"id"`
	AccessUrl                   types.String      `tfsdk:"access_url"`
	AllocationPools             types.List        `tfsdk:"allocation_pools"`
	BackendId                   types.String      `tfsdk:"backend_id"`
	Cidr                        types.String      `tfsdk:"cidr"`
	Created                     timetypes.RFC3339 `tfsdk:"created"`
	Customer                    types.String      `tfsdk:"customer"`
	CustomerAbbreviation        types.String      `tfsdk:"customer_abbreviation"`
	CustomerName                types.String      `tfsdk:"customer_name"`
	CustomerNativeName          types.String      `tfsdk:"customer_native_name"`
	CustomerUuid                types.String      `tfsdk:"customer_uuid"`
	Description                 types.String      `tfsdk:"description"`
	DisableGateway              types.Bool        `tfsdk:"disable_gateway"`
	DnsNameservers              types.List        `tfsdk:"dns_nameservers"`
	EnableDhcp                  types.Bool        `tfsdk:"enable_dhcp"`
	ErrorMessage                types.String      `tfsdk:"error_message"`
	ErrorTraceback              types.String      `tfsdk:"error_traceback"`
	GatewayIp                   types.String      `tfsdk:"gateway_ip"`
	HostRoutes                  types.List        `tfsdk:"host_routes"`
	IpVersion                   types.Int64       `tfsdk:"ip_version"`
	IsConnected                 types.Bool        `tfsdk:"is_connected"`
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
	Network                     types.String      `tfsdk:"network"`
	NetworkName                 types.String      `tfsdk:"network_name"`
	Project                     types.String      `tfsdk:"project"`
	ProjectName                 types.String      `tfsdk:"project_name"`
	ProjectUuid                 types.String      `tfsdk:"project_uuid"`
	ResourceType                types.String      `tfsdk:"resource_type"`
	ServiceName                 types.String      `tfsdk:"service_name"`
	ServiceSettings             types.String      `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage types.String      `tfsdk:"service_settings_error_message"`
	ServiceSettingsState        types.String      `tfsdk:"service_settings_state"`
	ServiceSettingsUuid         types.String      `tfsdk:"service_settings_uuid"`
	State                       types.String      `tfsdk:"state"`
	Tenant                      types.String      `tfsdk:"tenant"`
	TenantName                  types.String      `tfsdk:"tenant_name"`
	Url                         types.String      `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackSubnetModel) CopyFrom(ctx context.Context, apiResp OpenstackSubnetResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)

	{
		listValAllocationPools, listDiagsAllocationPools := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"end":   types.StringType,
			"start": types.StringType,
		}}, apiResp.AllocationPools)
		diags.Append(listDiagsAllocationPools...)
		model.AllocationPools = listValAllocationPools
	}
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Cidr = types.StringPointerValue(apiResp.Cidr)
	valCreated, diagsCreated := timetypes.NewRFC3339PointerValue(apiResp.Created)
	diags.Append(diagsCreated...)
	model.Created = valCreated
	model.Description = types.StringPointerValue(apiResp.Description)
	model.DisableGateway = types.BoolPointerValue(apiResp.DisableGateway)
	listValDnsNameservers, listDiagsDnsNameservers := types.ListValueFrom(ctx, types.StringType, apiResp.DnsNameservers)
	model.DnsNameservers = listValDnsNameservers
	diags.Append(listDiagsDnsNameservers...)
	model.EnableDhcp = types.BoolPointerValue(apiResp.EnableDhcp)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.GatewayIp = types.StringPointerValue(apiResp.GatewayIp)

	{
		listValHostRoutes, listDiagsHostRoutes := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
			"nexthop":     types.StringType,
		}}, apiResp.HostRoutes)
		diags.Append(listDiagsHostRoutes...)
		model.HostRoutes = listValHostRoutes
	}
	model.IpVersion = types.Int64PointerValue(apiResp.IpVersion)
	model.IsConnected = types.BoolPointerValue(apiResp.IsConnected)
	valModified, diagsModified := timetypes.NewRFC3339PointerValue(apiResp.Modified)
	diags.Append(diagsModified...)
	model.Modified = valModified
	model.Name = types.StringPointerValue(apiResp.Name)
	model.Network = types.StringPointerValue(apiResp.Network)
	model.NetworkName = types.StringPointerValue(apiResp.NetworkName)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.State = types.StringPointerValue(apiResp.State)
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantName = types.StringPointerValue(apiResp.TenantName)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
