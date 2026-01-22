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

	AccessUrl                   *string                                    `json:"access_url" tfsdk:"access_url"`
	AdminStateUp                *bool                                      `json:"admin_state_up" tfsdk:"admin_state_up"`
	AllowedAddressPairs         []OpenstackPortAllowedAddressPairsResponse `json:"allowed_address_pairs" tfsdk:"allowed_address_pairs"`
	BackendId                   *string                                    `json:"backend_id" tfsdk:"backend_id"`
	Created                     *string                                    `json:"created" tfsdk:"created"`
	Customer                    *string                                    `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation        *string                                    `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName                *string                                    `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName          *string                                    `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerUuid                *string                                    `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description                 *string                                    `json:"description" tfsdk:"description"`
	DeviceId                    *string                                    `json:"device_id" tfsdk:"device_id"`
	DeviceOwner                 *string                                    `json:"device_owner" tfsdk:"device_owner"`
	ErrorMessage                *string                                    `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback              *string                                    `json:"error_traceback" tfsdk:"error_traceback"`
	FixedIps                    []OpenstackPortFixedIpsResponse            `json:"fixed_ips" tfsdk:"fixed_ips"`
	FloatingIps                 []string                                   `json:"floating_ips" tfsdk:"floating_ips"`
	IsLimitBased                *bool                                      `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                                      `json:"is_usage_based" tfsdk:"is_usage_based"`
	MacAddress                  *string                                    `json:"mac_address" tfsdk:"mac_address"`
	MarketplaceCategoryName     *string                                    `json:"marketplace_category_name" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                                    `json:"marketplace_category_uuid" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                                    `json:"marketplace_offering_name" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                                    `json:"marketplace_offering_uuid" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                                    `json:"marketplace_plan_uuid" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                                    `json:"marketplace_resource_state" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                                    `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                                    `json:"modified" tfsdk:"modified"`
	Name                        *string                                    `json:"name" tfsdk:"name"`
	Network                     *string                                    `json:"network" tfsdk:"network"`
	NetworkName                 *string                                    `json:"network_name" tfsdk:"network_name"`
	NetworkUuid                 *string                                    `json:"network_uuid" tfsdk:"network_uuid"`
	PortSecurityEnabled         *bool                                      `json:"port_security_enabled" tfsdk:"port_security_enabled"`
	Project                     *string                                    `json:"project" tfsdk:"project"`
	ProjectName                 *string                                    `json:"project_name" tfsdk:"project_name"`
	ProjectUuid                 *string                                    `json:"project_uuid" tfsdk:"project_uuid"`
	ResourceType                *string                                    `json:"resource_type" tfsdk:"resource_type"`
	SecurityGroups              []OpenstackPortSecurityGroupsResponse      `json:"security_groups" tfsdk:"security_groups"`
	ServiceName                 *string                                    `json:"service_name" tfsdk:"service_name"`
	ServiceSettings             *string                                    `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                                    `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                                    `json:"service_settings_state" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                                    `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	State                       *string                                    `json:"state" tfsdk:"state"`
	Status                      *string                                    `json:"status" tfsdk:"status"`
	Tenant                      *string                                    `json:"tenant" tfsdk:"tenant"`
	TenantName                  *string                                    `json:"tenant_name" tfsdk:"tenant_name"`
	TenantUuid                  *string                                    `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Url                         *string                                    `json:"url" tfsdk:"url"`
}

type OpenstackPortAllowedAddressPairsResponse struct {
	MacAddress *string `json:"mac_address" tfsdk:"mac_address"`
}

type OpenstackPortFixedIpsResponse struct {
	IpAddress *string `json:"ip_address" tfsdk:"ip_address"`
	SubnetId  *string `json:"subnet_id" tfsdk:"subnet_id"`
}

type OpenstackPortSecurityGroupsResponse struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
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
				MarkdownDescription: "Admin state up",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "ID of the backend",
			},
			"device_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "ID of the device",
			},
			"device_owner": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Device owner",
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
				MarkdownDescription: "Mac address",
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
				MarkdownDescription: "Status",
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
				MarkdownDescription: "Access url",
			},
			"allowed_address_pairs": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"mac_address": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Mac address",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Allowed address pairs",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Created",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Description of the resource",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error message",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error traceback",
			},
			"floating_ips": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "Floating ips",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Modified",
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
				MarkdownDescription: "Resource type",
			},
			"security_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the resource",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Url",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Security groups",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "State",
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

		type filterDef struct {
			name string
			val  attr.Value
		}
		filterDefs := []filterDef{
			{"admin_state_up", data.AdminStateUp},
			{"backend_id", data.BackendId},
			{"device_id", data.DeviceId},
			{"device_owner", data.DeviceOwner},
			{"exclude_subnet_uuids", data.ExcludeSubnetUuids},
			{"fixed_ips", data.FixedIps},
			{"has_device_owner", data.HasDeviceOwner},
			{"mac_address", data.MacAddress},
			{"name", data.Name},
			{"name_exact", data.NameExact},
			{"network_name", data.NetworkName},
			{"network_uuid", data.NetworkUuid},
			{"query", data.Query},
			{"status", data.Status},
			{"tenant", data.Tenant},
			{"tenant_uuid", data.TenantUuid},
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
	model.AdminStateUp = types.BoolPointerValue(apiResp.AdminStateUp)
	listValAllowedAddressPairs, listDiagsAllowedAddressPairs := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"mac_address": types.StringType,
	}}, apiResp.AllowedAddressPairs)
	diags.Append(listDiagsAllowedAddressPairs...)
	model.AllowedAddressPairs = listValAllowedAddressPairs
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.DeviceId = types.StringPointerValue(apiResp.DeviceId)
	model.DeviceOwner = types.StringPointerValue(apiResp.DeviceOwner)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.FloatingIps, _ = types.ListValueFrom(ctx, types.StringType, apiResp.FloatingIps)
	model.MacAddress = types.StringPointerValue(apiResp.MacAddress)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.Network = types.StringPointerValue(apiResp.Network)
	model.NetworkName = types.StringPointerValue(apiResp.NetworkName)
	model.NetworkUuid = types.StringPointerValue(apiResp.NetworkUuid)
	model.PortSecurityEnabled = types.BoolPointerValue(apiResp.PortSecurityEnabled)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	listValSecurityGroups, listDiagsSecurityGroups := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"name": types.StringType,
		"url":  types.StringType,
	}}, apiResp.SecurityGroups)
	diags.Append(listDiagsSecurityGroups...)
	model.SecurityGroups = listValSecurityGroups
	model.State = types.StringPointerValue(apiResp.State)
	model.Status = types.StringPointerValue(apiResp.Status)
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantName = types.StringPointerValue(apiResp.TenantName)
	model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
