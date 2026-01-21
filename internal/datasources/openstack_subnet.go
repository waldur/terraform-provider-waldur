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

	AccessUrl                   *string                                  `json:"access_url" tfsdk:"access_url"`
	AllocationPools             []OpenstackSubnetAllocationPoolsResponse `json:"allocation_pools" tfsdk:"allocation_pools"`
	BackendId                   *string                                  `json:"backend_id" tfsdk:"backend_id"`
	Cidr                        *string                                  `json:"cidr" tfsdk:"cidr"`
	Created                     *string                                  `json:"created" tfsdk:"created"`
	Customer                    *string                                  `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation        *string                                  `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName                *string                                  `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName          *string                                  `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerUuid                *string                                  `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description                 *string                                  `json:"description" tfsdk:"description"`
	DisableGateway              *bool                                    `json:"disable_gateway" tfsdk:"disable_gateway"`
	DnsNameservers              []string                                 `json:"dns_nameservers" tfsdk:"dns_nameservers"`
	EnableDhcp                  *bool                                    `json:"enable_dhcp" tfsdk:"enable_dhcp"`
	ErrorMessage                *string                                  `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback              *string                                  `json:"error_traceback" tfsdk:"error_traceback"`
	GatewayIp                   *string                                  `json:"gateway_ip" tfsdk:"gateway_ip"`
	HostRoutes                  []OpenstackSubnetHostRoutesResponse      `json:"host_routes" tfsdk:"host_routes"`
	IpVersion                   *int64                                   `json:"ip_version" tfsdk:"ip_version"`
	IsConnected                 *bool                                    `json:"is_connected" tfsdk:"is_connected"`
	IsLimitBased                *bool                                    `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                                    `json:"is_usage_based" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string                                  `json:"marketplace_category_name" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                                  `json:"marketplace_category_uuid" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                                  `json:"marketplace_offering_name" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                                  `json:"marketplace_offering_uuid" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                                  `json:"marketplace_plan_uuid" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                                  `json:"marketplace_resource_state" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                                  `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                                  `json:"modified" tfsdk:"modified"`
	Name                        *string                                  `json:"name" tfsdk:"name"`
	Network                     *string                                  `json:"network" tfsdk:"network"`
	NetworkName                 *string                                  `json:"network_name" tfsdk:"network_name"`
	Project                     *string                                  `json:"project" tfsdk:"project"`
	ProjectName                 *string                                  `json:"project_name" tfsdk:"project_name"`
	ProjectUuid                 *string                                  `json:"project_uuid" tfsdk:"project_uuid"`
	ResourceType                *string                                  `json:"resource_type" tfsdk:"resource_type"`
	ServiceName                 *string                                  `json:"service_name" tfsdk:"service_name"`
	ServiceSettings             *string                                  `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                                  `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                                  `json:"service_settings_state" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                                  `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	State                       *string                                  `json:"state" tfsdk:"state"`
	Tenant                      *string                                  `json:"tenant" tfsdk:"tenant"`
	TenantName                  *string                                  `json:"tenant_name" tfsdk:"tenant_name"`
	Url                         *string                                  `json:"url" tfsdk:"url"`
}

type OpenstackSubnetAllocationPoolsResponse struct {
	End   *string `json:"end" tfsdk:"end"`
	Start *string `json:"start" tfsdk:"start"`
}

type OpenstackSubnetHostRoutesResponse struct {
	Destination *string `json:"destination" tfsdk:"destination"`
	Nexthop     *string `json:"nexthop" tfsdk:"nexthop"`
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
				MarkdownDescription: "Access url",
			},
			"allocation_pools": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"end":   types.StringType,
					"start": types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: "Allocation pools",
			},
			"cidr": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Cidr",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Created",
			},
			"disable_gateway": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "If True, no gateway IP address will be allocated",
			},
			"dns_nameservers": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Computed:            true,
				MarkdownDescription: "Dns nameservers",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error message",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error traceback",
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
				MarkdownDescription: "Host routes",
			},
			"is_connected": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Is subnet connected to the default tenant router.",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Modified",
			},
			"network_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the network",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource type",
			},
			"tenant_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the tenant",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Url",
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

		type filterDef struct {
			name string
			val  attr.Value
		}
		filterDefs := []filterDef{
			{"backend_id", data.BackendId},
			{"can_manage", data.CanManage},
			{"customer", data.Customer},
			{"customer_abbreviation", data.CustomerAbbreviation},
			{"customer_name", data.CustomerName},
			{"customer_native_name", data.CustomerNativeName},
			{"customer_uuid", data.CustomerUuid},
			{"description", data.Description},
			{"direct_only", data.DirectOnly},
			{"enable_dhcp", data.EnableDhcp},
			{"external_ip", data.ExternalIp},
			{"ip_version", data.IpVersion},
			{"name", data.Name},
			{"name_exact", data.NameExact},
			{"network", data.Network},
			{"network_uuid", data.NetworkUuid},
			{"project", data.Project},
			{"project_name", data.ProjectName},
			{"project_uuid", data.ProjectUuid},
			{"rbac_only", data.RbacOnly},
			{"service_settings_name", data.ServiceSettingsName},
			{"service_settings_uuid", data.ServiceSettingsUuid},
			{"state", data.State},
			{"tenant", data.Tenant},
			{"tenant_uuid", data.TenantUuid},
			{"uuid", data.Uuid},
		}

		filters := make(map[string]string)
		for _, fd := range filterDefs {
			if fd.val.IsNull() || fd.val.IsUnknown() {
				continue
			}
			switch v := fd.val.(type) {
			case types.String:
				filters[fd.name] = v.ValueString()
			case types.Int64:
				filters[fd.name] = fmt.Sprintf("%d", v.ValueInt64())
			case types.Bool:
				filters[fd.name] = fmt.Sprintf("%t", v.ValueBool())
			case types.Float64:
				filters[fd.name] = fmt.Sprintf("%f", v.ValueFloat64())
			}
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
	listValAllocationPools, listDiagsAllocationPools := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"end":   types.StringType,
		"start": types.StringType,
	}}, apiResp.AllocationPools)
	diags.Append(listDiagsAllocationPools...)
	model.AllocationPools = listValAllocationPools
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Cidr = types.StringPointerValue(apiResp.Cidr)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Customer = types.StringPointerValue(apiResp.Customer)
	model.CustomerAbbreviation = types.StringPointerValue(apiResp.CustomerAbbreviation)
	model.CustomerName = types.StringPointerValue(apiResp.CustomerName)
	model.CustomerNativeName = types.StringPointerValue(apiResp.CustomerNativeName)
	model.CustomerUuid = types.StringPointerValue(apiResp.CustomerUuid)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.DisableGateway = types.BoolPointerValue(apiResp.DisableGateway)
	model.DnsNameservers, _ = types.ListValueFrom(ctx, types.StringType, apiResp.DnsNameservers)
	model.EnableDhcp = types.BoolPointerValue(apiResp.EnableDhcp)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.GatewayIp = types.StringPointerValue(apiResp.GatewayIp)
	listValHostRoutes, listDiagsHostRoutes := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"destination": types.StringType,
		"nexthop":     types.StringType,
	}}, apiResp.HostRoutes)
	diags.Append(listDiagsHostRoutes...)
	model.HostRoutes = listValHostRoutes
	model.IpVersion = types.Int64PointerValue(apiResp.IpVersion)
	model.IsConnected = types.BoolPointerValue(apiResp.IsConnected)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.Network = types.StringPointerValue(apiResp.Network)
	model.NetworkName = types.StringPointerValue(apiResp.NetworkName)
	model.Project = types.StringPointerValue(apiResp.Project)
	model.ProjectName = types.StringPointerValue(apiResp.ProjectName)
	model.ProjectUuid = types.StringPointerValue(apiResp.ProjectUuid)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.State = types.StringPointerValue(apiResp.State)
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantName = types.StringPointerValue(apiResp.TenantName)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
