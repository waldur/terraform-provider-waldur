package subnet

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

func OpenStackStaticRouteRequestType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"destination": types.StringType,
		"nexthop":     types.StringType,
	}}
}
func OpenStackSubNetAllocationPoolRequestType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"end":   types.StringType,
		"start": types.StringType,
	}}
}

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
			"direct_only": schema.BoolAttribute{
				Optional: true,
			},
			"enable_dhcp": schema.BoolAttribute{
				Optional: true,
			},
			"external_ip": schema.StringAttribute{
				Optional: true,
			},
			"ip_version": schema.Int64Attribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"name_exact": schema.StringAttribute{
				Optional: true,
			},
			"network": schema.StringAttribute{
				Optional: true,
			},
			"network_uuid": schema.StringAttribute{
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
			"rbac_only": schema.BoolAttribute{
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

type OpenstackSubnetModel struct {
	UUID                    types.String `tfsdk:"id"`
	AllocationPools         types.List   `tfsdk:"allocation_pools"`
	BackendId               types.String `tfsdk:"backend_id"`
	Cidr                    types.String `tfsdk:"cidr"`
	Customer                types.String `tfsdk:"customer"`
	Description             types.String `tfsdk:"description"`
	DisableGateway          types.Bool   `tfsdk:"disable_gateway"`
	DnsNameservers          types.List   `tfsdk:"dns_nameservers"`
	EnableDhcp              types.Bool   `tfsdk:"enable_dhcp"`
	ErrorMessage            types.String `tfsdk:"error_message"`
	ErrorTraceback          types.String `tfsdk:"error_traceback"`
	GatewayIp               types.String `tfsdk:"gateway_ip"`
	HostRoutes              types.List   `tfsdk:"host_routes"`
	IpVersion               types.Int64  `tfsdk:"ip_version"`
	IsConnected             types.Bool   `tfsdk:"is_connected"`
	MarketplaceResourceUuid types.String `tfsdk:"marketplace_resource_uuid"`
	Name                    types.String `tfsdk:"name"`
	Network                 types.String `tfsdk:"network"`
	NetworkName             types.String `tfsdk:"network_name"`
	Project                 types.String `tfsdk:"project"`
	ResourceType            types.String `tfsdk:"resource_type"`
	State                   types.String `tfsdk:"state"`
	Tenant                  types.String `tfsdk:"tenant"`
	TenantName              types.String `tfsdk:"tenant_name"`
	Url                     types.String `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackSubnetModel) CopyFrom(ctx context.Context, apiResp OpenstackSubnetResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)

	if apiResp.AllocationPools != nil && len(*apiResp.AllocationPools) > 0 {
		listValAllocationPools, listDiagsAllocationPools := types.ListValueFrom(ctx, OpenStackSubNetAllocationPoolRequestType(), apiResp.AllocationPools)
		diags.Append(listDiagsAllocationPools...)
		model.AllocationPools = listValAllocationPools
	} else {
		model.AllocationPools = types.ListNull(OpenStackSubNetAllocationPoolRequestType())
	}
	model.BackendId = common.StringPointerValue(apiResp.BackendId)
	model.Cidr = common.StringPointerValue(apiResp.Cidr)
	model.Customer = common.StringPointerValue(apiResp.Customer)
	model.Description = common.StringPointerValue(apiResp.Description)
	model.DisableGateway = types.BoolPointerValue(apiResp.DisableGateway)
	listValDnsNameservers, listDiagsDnsNameservers := types.ListValueFrom(ctx, types.StringType, apiResp.DnsNameservers)
	model.DnsNameservers = listValDnsNameservers
	diags.Append(listDiagsDnsNameservers...)
	model.EnableDhcp = types.BoolPointerValue(apiResp.EnableDhcp)
	model.ErrorMessage = common.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = common.StringPointerValue(apiResp.ErrorTraceback)
	model.GatewayIp = common.StringPointerValue(apiResp.GatewayIp)

	if apiResp.HostRoutes != nil && len(*apiResp.HostRoutes) > 0 {
		listValHostRoutes, listDiagsHostRoutes := types.ListValueFrom(ctx, OpenStackStaticRouteRequestType(), apiResp.HostRoutes)
		diags.Append(listDiagsHostRoutes...)
		model.HostRoutes = listValHostRoutes
	} else {
		model.HostRoutes = types.ListNull(OpenStackStaticRouteRequestType())
	}
	model.IpVersion = types.Int64PointerValue(apiResp.IpVersion)
	model.IsConnected = types.BoolPointerValue(apiResp.IsConnected)
	model.MarketplaceResourceUuid = common.StringPointerValue(apiResp.MarketplaceResourceUuid)
	model.Name = common.StringPointerValue(apiResp.Name)
	model.Network = common.StringPointerValue(apiResp.Network)
	model.NetworkName = common.StringPointerValue(apiResp.NetworkName)
	model.Project = common.StringPointerValue(apiResp.Project)
	model.ResourceType = common.StringPointerValue(apiResp.ResourceType)
	model.State = common.StringPointerValue(apiResp.State)
	model.Tenant = common.StringPointerValue(apiResp.Tenant)
	model.TenantName = common.StringPointerValue(apiResp.TenantName)
	model.Url = common.StringPointerValue(apiResp.Url)

	return diags
}
