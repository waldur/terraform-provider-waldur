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
var _ datasource.DataSource = &OpenstackPortDataSource{}

func NewOpenstackPortDataSource() datasource.DataSource {
	return &OpenstackPortDataSource{}
}

// OpenstackPortDataSource defines the data source implementation.
type OpenstackPortDataSource struct {
	client *client.Client
}

// OpenstackPortApiResponse is the API response model.
type OpenstackPortApiResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl           *string                                    `json:"access_url" tfsdk:"access_url"`
	AllowedAddressPairs []OpenstackPortAllowedAddressPairsResponse `json:"allowed_address_pairs" tfsdk:"allowed_address_pairs"`
	Created             *string                                    `json:"created" tfsdk:"created"`
	Description         *string                                    `json:"description" tfsdk:"description"`
	ErrorMessage        *string                                    `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback      *string                                    `json:"error_traceback" tfsdk:"error_traceback"`
	FloatingIps         []string                                   `json:"floating_ips" tfsdk:"floating_ips"`
	Modified            *string                                    `json:"modified" tfsdk:"modified"`
	Network             *string                                    `json:"network" tfsdk:"network"`
	PortSecurityEnabled *bool                                      `json:"port_security_enabled" tfsdk:"port_security_enabled"`
	ResourceType        *string                                    `json:"resource_type" tfsdk:"resource_type"`
	SecurityGroups      []OpenstackPortSecurityGroupsResponse      `json:"security_groups" tfsdk:"security_groups"`
	State               *string                                    `json:"state" tfsdk:"state"`
	TenantName          *string                                    `json:"tenant_name" tfsdk:"tenant_name"`
	Url                 *string                                    `json:"url" tfsdk:"url"`
}

type OpenstackPortAllowedAddressPairsResponse struct {
	MacAddress *string `json:"mac_address" tfsdk:"mac_address"`
}

type OpenstackPortSecurityGroupsResponse struct {
	Url *string `json:"url" tfsdk:"url"`
}

var openstackport_allowed_address_pairsAttrTypes = map[string]attr.Type{
	"mac_address": types.StringType,
}
var openstackport_allowed_address_pairsObjectType = types.ObjectType{
	AttrTypes: openstackport_allowed_address_pairsAttrTypes,
}

var openstackport_security_groupsAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"url":  types.StringType,
}
var openstackport_security_groupsObjectType = types.ObjectType{
	AttrTypes: openstackport_security_groupsAttrTypes,
}

// OpenstackPortDataSourceModel describes the data source data model.
type OpenstackPortDataSourceModel struct {
	UUID                types.String `tfsdk:"id"`
	AdminStateUp        types.Bool   `tfsdk:"admin_state_up"`
	BackendId           types.String `tfsdk:"backend_id"`
	DeviceId            types.String `tfsdk:"device_id"`
	DeviceOwner         types.String `tfsdk:"device_owner"`
	ExcludeSubnetUuids  types.String `tfsdk:"exclude_subnet_uuids"`
	FixedIps            types.String `tfsdk:"fixed_ips"`
	HasDeviceOwner      types.Bool   `tfsdk:"has_device_owner"`
	MacAddress          types.String `tfsdk:"mac_address"`
	Name                types.String `tfsdk:"name"`
	NameExact           types.String `tfsdk:"name_exact"`
	NetworkName         types.String `tfsdk:"network_name"`
	NetworkUuid         types.String `tfsdk:"network_uuid"`
	Query               types.String `tfsdk:"query"`
	Status              types.String `tfsdk:"status"`
	Tenant              types.String `tfsdk:"tenant"`
	TenantUuid          types.String `tfsdk:"tenant_uuid"`
	AccessUrl           types.String `tfsdk:"access_url"`
	AllowedAddressPairs types.List   `tfsdk:"allowed_address_pairs"`
	Created             types.String `tfsdk:"created"`
	Description         types.String `tfsdk:"description"`
	ErrorMessage        types.String `tfsdk:"error_message"`
	ErrorTraceback      types.String `tfsdk:"error_traceback"`
	FloatingIps         types.List   `tfsdk:"floating_ips"`
	Modified            types.String `tfsdk:"modified"`
	Network             types.String `tfsdk:"network"`
	PortSecurityEnabled types.Bool   `tfsdk:"port_security_enabled"`
	ResourceType        types.String `tfsdk:"resource_type"`
	SecurityGroups      types.List   `tfsdk:"security_groups"`
	State               types.String `tfsdk:"state"`
	TenantName          types.String `tfsdk:"tenant_name"`
	Url                 types.String `tfsdk:"url"`
}

func (d *OpenstackPortDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_port"
}

func (d *OpenstackPortDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Port data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"admin_state_up": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"device_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"device_owner": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"exclude_subnet_uuids": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Exclude Subnet UUIDs (comma-separated)",
			},
			"fixed_ips": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Search by fixed IP",
			},
			"has_device_owner": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Has device owner",
			},
			"mac_address": schema.StringAttribute{
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
			"network_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Search by network name",
			},
			"network_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Search by network UUID",
			},
			"query": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Search by name, MAC address or backend ID",
			},
			"status": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"tenant": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Tenant URL",
			},
			"tenant_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Tenant UUID",
			},
			"access_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"allowed_address_pairs": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"mac_address": types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"description": schema.StringAttribute{
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
			"floating_ips": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"network": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Network to which this port belongs",
			},
			"port_security_enabled": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "If True, security groups and rules will be applied to this port",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"security_groups": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"name": types.StringType,
					"url":  types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"state": schema.StringAttribute{
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

func (d *OpenstackPortDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OpenstackPortDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackPortDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var apiResp OpenstackPortApiResponse

		err := d.client.GetByUUID(ctx, "/api/openstack-ports/{uuid}/", data.UUID.ValueString(), &apiResp)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Port",
				"An error occurred while reading the Openstack Port by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, apiResp, &data)...)

	} else {
		// Filter by provided parameters
		var results []OpenstackPortApiResponse

		filters := map[string]string{}
		if !data.AdminStateUp.IsNull() {
			filters["admin_state_up"] = fmt.Sprintf("%t", data.AdminStateUp.ValueBool())
		}
		if !data.BackendId.IsNull() {
			filters["backend_id"] = data.BackendId.ValueString()
		}
		if !data.DeviceId.IsNull() {
			filters["device_id"] = data.DeviceId.ValueString()
		}
		if !data.DeviceOwner.IsNull() {
			filters["device_owner"] = data.DeviceOwner.ValueString()
		}
		if !data.ExcludeSubnetUuids.IsNull() {
			filters["exclude_subnet_uuids"] = data.ExcludeSubnetUuids.ValueString()
		}
		if !data.FixedIps.IsNull() {
			filters["fixed_ips"] = data.FixedIps.ValueString()
		}
		if !data.HasDeviceOwner.IsNull() {
			filters["has_device_owner"] = fmt.Sprintf("%t", data.HasDeviceOwner.ValueBool())
		}
		if !data.MacAddress.IsNull() {
			filters["mac_address"] = data.MacAddress.ValueString()
		}
		if !data.Name.IsNull() {
			filters["name"] = data.Name.ValueString()
		}
		if !data.NameExact.IsNull() {
			filters["name_exact"] = data.NameExact.ValueString()
		}
		if !data.NetworkName.IsNull() {
			filters["network_name"] = data.NetworkName.ValueString()
		}
		if !data.NetworkUuid.IsNull() {
			filters["network_uuid"] = data.NetworkUuid.ValueString()
		}
		if !data.Query.IsNull() {
			filters["query"] = data.Query.ValueString()
		}
		if !data.Status.IsNull() {
			filters["status"] = data.Status.ValueString()
		}
		if !data.Tenant.IsNull() {
			filters["tenant"] = data.Tenant.ValueString()
		}
		if !data.TenantUuid.IsNull() {
			filters["tenant_uuid"] = data.TenantUuid.ValueString()
		}

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_port.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/openstack-ports/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Openstack Port",
				"An error occurred while filtering Openstack Port: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Port Not Found",
				"No Openstack Port found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Ports Found",
				fmt.Sprintf("Found %d Openstack Ports with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *OpenstackPortDataSource) mapResponseToModel(ctx context.Context, apiResp OpenstackPortApiResponse, model *OpenstackPortDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	listValAllowedAddressPairs, listDiagsAllowedAddressPairs := types.ListValueFrom(ctx, openstackport_allowed_address_pairsObjectType, apiResp.AllowedAddressPairs)
	diags.Append(listDiagsAllowedAddressPairs...)
	model.AllowedAddressPairs = listValAllowedAddressPairs
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.FloatingIps, _ = types.ListValueFrom(ctx, types.StringType, apiResp.FloatingIps)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.Network = types.StringPointerValue(apiResp.Network)
	model.PortSecurityEnabled = types.BoolPointerValue(apiResp.PortSecurityEnabled)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	listValSecurityGroups, listDiagsSecurityGroups := types.ListValueFrom(ctx, openstackport_security_groupsObjectType, apiResp.SecurityGroups)
	diags.Append(listDiagsSecurityGroups...)
	model.SecurityGroups = listValSecurityGroups
	model.State = types.StringPointerValue(apiResp.State)
	model.TenantName = types.StringPointerValue(apiResp.TenantName)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
