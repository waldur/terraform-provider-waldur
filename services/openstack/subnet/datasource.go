package subnet

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackSubnetDataSource{}

func NewOpenstackSubnetDataSource() datasource.DataSource {
	return &OpenstackSubnetDataSource{}
}

// OpenstackSubnetDataSource defines the data source implementation.
type OpenstackSubnetDataSource struct {
	client *Client
}

// OpenstackSubnetFiltersModel contains the filter parameters for querying.
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
	State                types.String `tfsdk:"state"`
	Tenant               types.String `tfsdk:"tenant"`
	TenantUuid           types.String `tfsdk:"tenant_uuid"`
	Uuid                 types.String `tfsdk:"uuid"`
}

type OpenstackSubnetDataSourceModel struct {
	UUID            types.String                 `tfsdk:"id"`
	Filters         *OpenstackSubnetFiltersModel `tfsdk:"filters"`
	AccessUrl       types.String                 `tfsdk:"access_url"`
	AllocationPools types.List                   `tfsdk:"allocation_pools"`
	BackendId       types.String                 `tfsdk:"backend_id"`
	Cidr            types.String                 `tfsdk:"cidr"`
	Created         types.String                 `tfsdk:"created"`
	Description     types.String                 `tfsdk:"description"`
	DisableGateway  types.Bool                   `tfsdk:"disable_gateway"`
	DnsNameservers  types.List                   `tfsdk:"dns_nameservers"`
	EnableDhcp      types.Bool                   `tfsdk:"enable_dhcp"`
	ErrorMessage    types.String                 `tfsdk:"error_message"`
	ErrorTraceback  types.String                 `tfsdk:"error_traceback"`
	GatewayIp       types.String                 `tfsdk:"gateway_ip"`
	HostRoutes      types.List                   `tfsdk:"host_routes"`
	IpVersion       types.Int64                  `tfsdk:"ip_version"`
	IsConnected     types.Bool                   `tfsdk:"is_connected"`
	Modified        types.String                 `tfsdk:"modified"`
	Name            types.String                 `tfsdk:"name"`
	Network         types.String                 `tfsdk:"network"`
	NetworkName     types.String                 `tfsdk:"network_name"`
	ResourceType    types.String                 `tfsdk:"resource_type"`
	State           types.String                 `tfsdk:"state"`
	Tenant          types.String                 `tfsdk:"tenant"`
	TenantName      types.String                 `tfsdk:"tenant_name"`
	Url             types.String                 `tfsdk:"url"`
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
			"filters": schema.SingleNestedAttribute{
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
					"state": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "State Allowed values: `CREATING`, `CREATION_SCHEDULED`, `DELETING`, `DELETION_SCHEDULED`, `ERRED`, `OK`, `UPDATE_SCHEDULED`, `UPDATING`.",
						Validators: []validator.String{
							stringvalidator.OneOf("CREATING", "CREATION_SCHEDULED", "DELETING", "DELETION_SCHEDULED", "ERRED", "OK", "UPDATE_SCHEDULED", "UPDATING"),
						},
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
			},
			"access_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Access url",
			},
			"allocation_pools": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"end": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "An IPv4 or IPv6 address.",
						},
						"start": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "An IPv4 or IPv6 address.",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Allocation pools",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of the backend",
			},
			"cidr": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Cidr",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Created",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Description of the resource",
			},
			"disable_gateway": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "If True, no gateway IP address will be allocated",
			},
			"dns_nameservers": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "Dns nameservers",
			},
			"enable_dhcp": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "If True, DHCP service will be enabled on this subnet",
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
			"host_routes": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"destination": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Destination",
						},
						"nexthop": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "An IPv4 or IPv6 address.",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Host routes",
			},
			"ip_version": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "IP protocol version (4 or 6)",
			},
			"is_connected": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Is subnet connected to the default tenant router.",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the resource",
			},
			"network": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Network to which this subnet belongs",
			},
			"network_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the network",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource type",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "State",
			},
			"tenant": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Tenant",
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

	rawClient, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = NewClient(rawClient)
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
		apiResp, err := d.client.GetOpenstackSubnet(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Subnet",
				"An error occurred while reading the Openstack Subnet by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, *apiResp, &data)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_subnet.",
			)
			return
		}

		results, err := d.client.ListOpenstackSubnet(ctx, filters)
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

func (d *OpenstackSubnetDataSource) mapResponseToModel(ctx context.Context, apiResp OpenstackSubnetResponse, model *OpenstackSubnetDataSourceModel) diag.Diagnostics {
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
	model.Created = types.StringPointerValue(apiResp.Created)
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
	model.Modified = types.StringPointerValue(apiResp.Modified)
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
