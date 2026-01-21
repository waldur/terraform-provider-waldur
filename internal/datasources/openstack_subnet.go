package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackSubnetDataSource{}

func NewOpenstackSubnetDataSource() datasource.DataSource {
	return &OpenstackSubnetDataSource{}
}

// OpenstackSubnetDataSource defines the data source implementation.
type OpenstackSubnetDataSource struct {
	client *client.Client
}

// OpenstackSubnetApiResponse is the API response model.
type OpenstackSubnetApiResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl       *string                                  `json:"access_url" tfsdk:"access_url"`
	AllocationPools []OpenstackSubnetAllocationPoolsResponse `json:"allocation_pools" tfsdk:"allocation_pools"`
	Cidr            *string                                  `json:"cidr" tfsdk:"cidr"`
	Created         *string                                  `json:"created" tfsdk:"created"`
	DisableGateway  *bool                                    `json:"disable_gateway" tfsdk:"disable_gateway"`
	DnsNameservers  []string                                 `json:"dns_nameservers" tfsdk:"dns_nameservers"`
	ErrorMessage    *string                                  `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback  *string                                  `json:"error_traceback" tfsdk:"error_traceback"`
	GatewayIp       *string                                  `json:"gateway_ip" tfsdk:"gateway_ip"`
	HostRoutes      []OpenstackSubnetHostRoutesResponse      `json:"host_routes" tfsdk:"host_routes"`
	IsConnected     *bool                                    `json:"is_connected" tfsdk:"is_connected"`
	Modified        *string                                  `json:"modified" tfsdk:"modified"`
	NetworkName     *string                                  `json:"network_name" tfsdk:"network_name"`
	ResourceType    *string                                  `json:"resource_type" tfsdk:"resource_type"`
	TenantName      *string                                  `json:"tenant_name" tfsdk:"tenant_name"`
	Url             *string                                  `json:"url" tfsdk:"url"`
}

type OpenstackSubnetAllocationPoolsResponse struct {
	End   *string `json:"end" tfsdk:"end"`
	Start *string `json:"start" tfsdk:"start"`
}

type OpenstackSubnetHostRoutesResponse struct {
	Destination *string `json:"destination" tfsdk:"destination"`
	Nexthop     *string `json:"nexthop" tfsdk:"nexthop"`
}

var openstacksubnet_allocation_poolsAttrTypes = map[string]attr.Type{
	"end":   types.StringType,
	"start": types.StringType,
}
var openstacksubnet_allocation_poolsObjectType = types.ObjectType{
	AttrTypes: openstacksubnet_allocation_poolsAttrTypes,
}

var openstacksubnet_host_routesAttrTypes = map[string]attr.Type{
	"destination": types.StringType,
	"nexthop":     types.StringType,
}
var openstacksubnet_host_routesObjectType = types.ObjectType{
	AttrTypes: openstacksubnet_host_routesAttrTypes,
}

// OpenstackSubnetDataSourceModel describes the data source data model.
type OpenstackSubnetDataSourceModel struct {
	UUID                 types.String `tfsdk:"id"`
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
	State                types.String `tfsdk:"state"`
	Tenant               types.String `tfsdk:"tenant"`
	TenantUuid           types.String `tfsdk:"tenant_uuid"`
	Uuid                 types.String `tfsdk:"uuid"`
	AccessUrl            types.String `tfsdk:"access_url"`
	AllocationPools      types.List   `tfsdk:"allocation_pools"`
	Cidr                 types.String `tfsdk:"cidr"`
	Created              types.String `tfsdk:"created"`
	DisableGateway       types.Bool   `tfsdk:"disable_gateway"`
	DnsNameservers       types.List   `tfsdk:"dns_nameservers"`
	ErrorMessage         types.String `tfsdk:"error_message"`
	ErrorTraceback       types.String `tfsdk:"error_traceback"`
	GatewayIp            types.String `tfsdk:"gateway_ip"`
	HostRoutes           types.List   `tfsdk:"host_routes"`
	IsConnected          types.Bool   `tfsdk:"is_connected"`
	Modified             types.String `tfsdk:"modified"`
	NetworkName          types.String `tfsdk:"network_name"`
	ResourceType         types.String `tfsdk:"resource_type"`
	TenantName           types.String `tfsdk:"tenant_name"`
	Url                  types.String `tfsdk:"url"`
}

func (d *OpenstackSubnetDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_subnet"
}

func (d *OpenstackSubnetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Subnet data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
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
			"direct_only": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Direct only",
			},
			"enable_dhcp": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"external_ip": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "External IP",
			},
			"ip_version": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: " ",
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
			"state": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "State",
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
			"access_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"allocation_pools": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"end":   types.StringType,
					"start": types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"cidr": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"disable_gateway": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "If True, no gateway IP address will be allocated",
			},
			"dns_nameservers": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"gateway_ip": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "IP address of the gateway for this subnet",
			},
			"host_routes": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"destination": types.StringType,
					"nexthop":     types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"is_connected": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Is subnet connected to the default tenant router.",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"network_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"tenant_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
		},
	}
}

func (d *OpenstackSubnetDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
}

func (d *OpenstackSubnetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackSubnetDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var apiResp OpenstackSubnetApiResponse

		err := d.client.GetByUUID(ctx, "/api/openstack-subnets/{uuid}/", data.UUID.ValueString(), &apiResp)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Subnet",
				"An error occurred while reading the Openstack Subnet by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, apiResp, &data)...)

	} else {
		// Filter by provided parameters
		var results []OpenstackSubnetApiResponse

		filters := map[string]string{}
		if !data.BackendId.IsNull() {
			filters["backend_id"] = data.BackendId.ValueString()
		}
		if !data.CanManage.IsNull() {
			filters["can_manage"] = fmt.Sprintf("%t", data.CanManage.ValueBool())
		}
		if !data.Customer.IsNull() {
			filters["customer"] = data.Customer.ValueString()
		}
		if !data.CustomerAbbreviation.IsNull() {
			filters["customer_abbreviation"] = data.CustomerAbbreviation.ValueString()
		}
		if !data.CustomerName.IsNull() {
			filters["customer_name"] = data.CustomerName.ValueString()
		}
		if !data.CustomerNativeName.IsNull() {
			filters["customer_native_name"] = data.CustomerNativeName.ValueString()
		}
		if !data.CustomerUuid.IsNull() {
			filters["customer_uuid"] = data.CustomerUuid.ValueString()
		}
		if !data.Description.IsNull() {
			filters["description"] = data.Description.ValueString()
		}
		if !data.DirectOnly.IsNull() {
			filters["direct_only"] = fmt.Sprintf("%t", data.DirectOnly.ValueBool())
		}
		if !data.EnableDhcp.IsNull() {
			filters["enable_dhcp"] = fmt.Sprintf("%t", data.EnableDhcp.ValueBool())
		}
		if !data.ExternalIp.IsNull() {
			filters["external_ip"] = data.ExternalIp.ValueString()
		}
		if !data.IpVersion.IsNull() {
			filters["ip_version"] = fmt.Sprintf("%d", data.IpVersion.ValueInt64())
		}
		if !data.Name.IsNull() {
			filters["name"] = data.Name.ValueString()
		}
		if !data.NameExact.IsNull() {
			filters["name_exact"] = data.NameExact.ValueString()
		}
		if !data.Network.IsNull() {
			filters["network"] = data.Network.ValueString()
		}
		if !data.NetworkUuid.IsNull() {
			filters["network_uuid"] = data.NetworkUuid.ValueString()
		}
		if !data.Project.IsNull() {
			filters["project"] = data.Project.ValueString()
		}
		if !data.ProjectName.IsNull() {
			filters["project_name"] = data.ProjectName.ValueString()
		}
		if !data.ProjectUuid.IsNull() {
			filters["project_uuid"] = data.ProjectUuid.ValueString()
		}
		if !data.RbacOnly.IsNull() {
			filters["rbac_only"] = fmt.Sprintf("%t", data.RbacOnly.ValueBool())
		}
		if !data.ServiceSettingsName.IsNull() {
			filters["service_settings_name"] = data.ServiceSettingsName.ValueString()
		}
		if !data.ServiceSettingsUuid.IsNull() {
			filters["service_settings_uuid"] = data.ServiceSettingsUuid.ValueString()
		}
		if !data.State.IsNull() {
			filters["state"] = data.State.ValueString()
		}
		if !data.Tenant.IsNull() {
			filters["tenant"] = data.Tenant.ValueString()
		}
		if !data.TenantUuid.IsNull() {
			filters["tenant_uuid"] = data.TenantUuid.ValueString()
		}
		if !data.Uuid.IsNull() {
			filters["uuid"] = data.Uuid.ValueString()
		}

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_subnet.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/openstack-subnets/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Openstack Subnet",
				"An error occurred while filtering Openstack Subnet: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Subnet Not Found",
				"No Openstack Subnet found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Subnets Found",
				fmt.Sprintf("Found %d Openstack Subnets with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *OpenstackSubnetDataSource) mapResponseToModel(ctx context.Context, apiResp OpenstackSubnetApiResponse, model *OpenstackSubnetDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	listValAllocationPools, listDiagsAllocationPools := types.ListValueFrom(ctx, openstacksubnet_allocation_poolsObjectType, apiResp.AllocationPools)
	diags.Append(listDiagsAllocationPools...)
	model.AllocationPools = listValAllocationPools
	model.Cidr = types.StringPointerValue(apiResp.Cidr)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.DisableGateway = types.BoolPointerValue(apiResp.DisableGateway)
	model.DnsNameservers, _ = types.ListValueFrom(ctx, types.StringType, apiResp.DnsNameservers)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.GatewayIp = types.StringPointerValue(apiResp.GatewayIp)
	listValHostRoutes, listDiagsHostRoutes := types.ListValueFrom(ctx, openstacksubnet_host_routesObjectType, apiResp.HostRoutes)
	diags.Append(listDiagsHostRoutes...)
	model.HostRoutes = listValHostRoutes
	model.IsConnected = types.BoolPointerValue(apiResp.IsConnected)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.NetworkName = types.StringPointerValue(apiResp.NetworkName)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.TenantName = types.StringPointerValue(apiResp.TenantName)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
