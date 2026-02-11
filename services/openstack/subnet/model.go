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

	if apiResp.AllocationPools != nil {
		valAllocationPools, diagsAllocationPools := types.ListValueFrom(ctx, OpenStackSubNetAllocationPoolRequestType(), apiResp.AllocationPools)
		diags.Append(diagsAllocationPools...)
		model.AllocationPools = valAllocationPools
	} else {
		model.AllocationPools = types.ListNull(OpenStackSubNetAllocationPoolRequestType())
	}

	model.BackendId = common.StringPointerValue(apiResp.BackendId)

	model.Cidr = common.StringPointerValue(apiResp.Cidr)

	model.Customer = common.StringPointerValue(apiResp.Customer)

	model.Description = common.StringPointerValue(apiResp.Description)

	model.DisableGateway = types.BoolPointerValue(apiResp.DisableGateway)

	if apiResp.DnsNameservers != nil {
		valDnsNameservers, diagsDnsNameservers := types.ListValueFrom(ctx, types.StringType, apiResp.DnsNameservers)
		diags.Append(diagsDnsNameservers...)
		model.DnsNameservers = valDnsNameservers
	} else {
		model.DnsNameservers = types.ListNull(types.StringType)
	}

	model.EnableDhcp = types.BoolPointerValue(apiResp.EnableDhcp)

	model.ErrorMessage = common.StringPointerValue(apiResp.ErrorMessage)

	model.GatewayIp = common.StringPointerValue(apiResp.GatewayIp)

	if apiResp.HostRoutes != nil {
		valHostRoutes, diagsHostRoutes := types.ListValueFrom(ctx, OpenStackStaticRouteRequestType(), apiResp.HostRoutes)
		diags.Append(diagsHostRoutes...)
		model.HostRoutes = valHostRoutes
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
