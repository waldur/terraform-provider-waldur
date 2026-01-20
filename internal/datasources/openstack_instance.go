package datasources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackInstanceDataSource{}

func NewOpenstackInstanceDataSource() datasource.DataSource {
	return &OpenstackInstanceDataSource{}
}

// OpenstackInstanceDataSource defines the data source implementation.
type OpenstackInstanceDataSource struct {
	client *client.Client
}

// OpenstackInstanceDataSourceModel describes the data source data model.
type OpenstackInstanceDataSourceModel struct {
	UUID                             types.String  `tfsdk:"id"`
	AttachVolumeUuid                 types.String  `tfsdk:"attach_volume_uuid"`
	AvailabilityZoneName             types.String  `tfsdk:"availability_zone_name"`
	BackendId                        types.String  `tfsdk:"backend_id"`
	CanManage                        types.Bool    `tfsdk:"can_manage"`
	Customer                         types.String  `tfsdk:"customer"`
	CustomerAbbreviation             types.String  `tfsdk:"customer_abbreviation"`
	CustomerName                     types.String  `tfsdk:"customer_name"`
	CustomerNativeName               types.String  `tfsdk:"customer_native_name"`
	CustomerUuid                     types.String  `tfsdk:"customer_uuid"`
	Description                      types.String  `tfsdk:"description"`
	ExternalIp                       types.String  `tfsdk:"external_ip"`
	Name                             types.String  `tfsdk:"name"`
	NameExact                        types.String  `tfsdk:"name_exact"`
	Project                          types.String  `tfsdk:"project"`
	ProjectName                      types.String  `tfsdk:"project_name"`
	ProjectUuid                      types.String  `tfsdk:"project_uuid"`
	Query                            types.String  `tfsdk:"query"`
	RuntimeState                     types.String  `tfsdk:"runtime_state"`
	ServiceSettingsName              types.String  `tfsdk:"service_settings_name"`
	ServiceSettingsUuid              types.String  `tfsdk:"service_settings_uuid"`
	State                            types.String  `tfsdk:"state"`
	Tenant                           types.String  `tfsdk:"tenant"`
	TenantUuid                       types.String  `tfsdk:"tenant_uuid"`
	Uuid                             types.String  `tfsdk:"uuid"`
	AccessUrl                        types.String  `tfsdk:"access_url"`
	Action                           types.String  `tfsdk:"action"`
	AvailabilityZone                 types.String  `tfsdk:"availability_zone"`
	ConnectDirectlyToExternalNetwork types.Bool    `tfsdk:"connect_directly_to_external_network"`
	Cores                            types.Int64   `tfsdk:"cores"`
	Created                          types.String  `tfsdk:"created"`
	Disk                             types.Int64   `tfsdk:"disk"`
	ErrorMessage                     types.String  `tfsdk:"error_message"`
	ErrorTraceback                   types.String  `tfsdk:"error_traceback"`
	ExternalAddress                  types.List    `tfsdk:"external_address"`
	ExternalIps                      types.List    `tfsdk:"external_ips"`
	FlavorDisk                       types.Int64   `tfsdk:"flavor_disk"`
	FlavorName                       types.String  `tfsdk:"flavor_name"`
	FloatingIps                      types.List    `tfsdk:"floating_ips"`
	HypervisorHostname               types.String  `tfsdk:"hypervisor_hostname"`
	ImageName                        types.String  `tfsdk:"image_name"`
	InternalIps                      types.List    `tfsdk:"internal_ips"`
	IsLimitBased                     types.Bool    `tfsdk:"is_limit_based"`
	IsUsageBased                     types.Bool    `tfsdk:"is_usage_based"`
	KeyFingerprint                   types.String  `tfsdk:"key_fingerprint"`
	KeyName                          types.String  `tfsdk:"key_name"`
	Latitude                         types.Float64 `tfsdk:"latitude"`
	Longitude                        types.Float64 `tfsdk:"longitude"`
	MarketplaceCategoryName          types.String  `tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid          types.String  `tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName          types.String  `tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid          types.String  `tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid              types.String  `tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState         types.String  `tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid          types.String  `tfsdk:"marketplace_resource_uuid"`
	MinDisk                          types.Int64   `tfsdk:"min_disk"`
	MinRam                           types.Int64   `tfsdk:"min_ram"`
	Modified                         types.String  `tfsdk:"modified"`
	Ports                            types.List    `tfsdk:"ports"`
	Ram                              types.Int64   `tfsdk:"ram"`
	ResourceType                     types.String  `tfsdk:"resource_type"`
	SecurityGroups                   types.List    `tfsdk:"security_groups"`
	ServerGroup                      types.Object  `tfsdk:"server_group"`
	ServiceName                      types.String  `tfsdk:"service_name"`
	ServiceSettings                  types.String  `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage      types.String  `tfsdk:"service_settings_error_message"`
	ServiceSettingsState             types.String  `tfsdk:"service_settings_state"`
	StartTime                        types.String  `tfsdk:"start_time"`
	Url                              types.String  `tfsdk:"url"`
	UserData                         types.String  `tfsdk:"user_data"`
	Volumes                          types.List    `tfsdk:"volumes"`
}

func (d *OpenstackInstanceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_instance"
}

func (d *OpenstackInstanceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Instance data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"attach_volume_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Filter for attachment to volume UUID",
			},
			"availability_zone_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Availability zone name",
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
			"query": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Search by name, internal IP, or external IP",
			},
			"runtime_state": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
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
			"action": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"availability_zone": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Availability zone where this instance is located",
			},
			"connect_directly_to_external_network": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "If True, instance will be connected directly to external network",
			},
			"cores": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of cores in a VM",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"disk": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Disk size in MiB",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"external_address": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"external_ips": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"flavor_disk": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Flavor disk size in MiB",
			},
			"flavor_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the flavor used by this instance",
			},
			"floating_ips": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"address": types.StringType,
					"port_fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"ip_address": types.StringType,
						"subnet_id":  types.StringType,
					}}},
					"port_mac_address":   types.StringType,
					"subnet":             types.StringType,
					"subnet_cidr":        types.StringType,
					"subnet_description": types.StringType,
					"subnet_name":        types.StringType,
					"subnet_uuid":        types.StringType,
					"url":                types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"hypervisor_hostname": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the hypervisor hosting this instance",
			},
			"image_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"internal_ips": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"is_limit_based": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"is_usage_based": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"key_fingerprint": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"key_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"latitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"longitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_category_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_category_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_offering_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_offering_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_resource_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"min_disk": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Minimum disk size in MiB",
			},
			"min_ram": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Minimum memory size in MiB",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"ports": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"mac_address": types.StringType,
					}}},
					"device_id":    types.StringType,
					"device_owner": types.StringType,
					"fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"ip_address": types.StringType,
						"subnet_id":  types.StringType,
					}}},
					"mac_address": types.StringType,
					"security_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"access_url":                 types.StringType,
						"backend_id":                 types.StringType,
						"created":                    types.StringType,
						"customer":                   types.StringType,
						"customer_abbreviation":      types.StringType,
						"customer_name":              types.StringType,
						"customer_native_name":       types.StringType,
						"customer_uuid":              types.StringType,
						"description":                types.StringType,
						"error_message":              types.StringType,
						"error_traceback":            types.StringType,
						"is_limit_based":             types.BoolType,
						"is_usage_based":             types.BoolType,
						"marketplace_category_name":  types.StringType,
						"marketplace_category_uuid":  types.StringType,
						"marketplace_offering_name":  types.StringType,
						"marketplace_offering_uuid":  types.StringType,
						"marketplace_plan_uuid":      types.StringType,
						"marketplace_resource_state": types.StringType,
						"marketplace_resource_uuid":  types.StringType,
						"modified":                   types.StringType,
						"name":                       types.StringType,
						"project":                    types.StringType,
						"project_name":               types.StringType,
						"project_uuid":               types.StringType,
						"resource_type":              types.StringType,
						"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"cidr":              types.StringType,
							"description":       types.StringType,
							"direction":         types.StringType,
							"ethertype":         types.StringType,
							"from_port":         types.Int64Type,
							"id":                types.Int64Type,
							"protocol":          types.StringType,
							"remote_group":      types.StringType,
							"remote_group_name": types.StringType,
							"remote_group_uuid": types.StringType,
							"to_port":           types.Int64Type,
						}}},
						"service_name":                   types.StringType,
						"service_settings":               types.StringType,
						"service_settings_error_message": types.StringType,
						"service_settings_state":         types.StringType,
						"service_settings_uuid":          types.StringType,
						"state":                          types.StringType,
						"tenant":                         types.StringType,
						"tenant_name":                    types.StringType,
						"tenant_uuid":                    types.StringType,
						"url":                            types.StringType,
					}}},
					"subnet":             types.StringType,
					"subnet_cidr":        types.StringType,
					"subnet_description": types.StringType,
					"subnet_name":        types.StringType,
					"subnet_uuid":        types.StringType,
					"url":                types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"ram": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Memory size in MiB",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"security_groups": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"description": types.StringType,
					"name":        types.StringType,
					"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"cidr":              types.StringType,
						"description":       types.StringType,
						"direction":         types.StringType,
						"ethertype":         types.StringType,
						"from_port":         types.Int64Type,
						"id":                types.Int64Type,
						"protocol":          types.StringType,
						"remote_group_name": types.StringType,
						"remote_group_uuid": types.StringType,
						"to_port":           types.Int64Type,
					}}},
					"state": types.StringType,
					"url":   types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"server_group": schema.ObjectAttribute{
				CustomType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"name":   types.StringType,
					"policy": types.StringType,
					"state":  types.StringType,
					"url":    types.StringType,
				}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_settings": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "OpenStack provider settings",
			},
			"service_settings_error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_settings_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"start_time": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"user_data": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Additional data that will be added to instance on provisioning",
			},
			"volumes": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"bootable":                  types.BoolType,
					"device":                    types.StringType,
					"image_name":                types.StringType,
					"marketplace_resource_uuid": types.StringType,
					"name":                      types.StringType,
					"resource_type":             types.StringType,
					"size":                      types.Int64Type,
					"state":                     types.StringType,
					"type":                      types.StringType,
					"type_name":                 types.StringType,
					"url":                       types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: "List of volumes attached to the instance",
			},
		},
	}
}

func (d *OpenstackInstanceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OpenstackInstanceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackInstanceDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var item map[string]interface{}

		err := d.client.GetByUUID(ctx, "/api/openstack-instances/{uuid}/", data.UUID.ValueString(), &item)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Instance",
				"An error occurred while reading the Openstack Instance by UUID: "+err.Error(),
			)
			return
		}

		// Extract data from single result
		if uuid, ok := item["uuid"].(string); ok {
			data.UUID = types.StringValue(uuid)
		}

		sourceMap := item
		// Map response fields to data model
		_ = sourceMap
		if val, ok := sourceMap["access_url"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AccessUrl = types.StringValue(str)
			}
		} else {
			if data.AccessUrl.IsUnknown() {
				data.AccessUrl = types.StringNull()
			}
		}
		if val, ok := sourceMap["action"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Action = types.StringValue(str)
			}
		} else {
			if data.Action.IsUnknown() {
				data.Action = types.StringNull()
			}
		}
		if val, ok := sourceMap["availability_zone"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AvailabilityZone = types.StringValue(str)
			}
		} else {
			if data.AvailabilityZone.IsUnknown() {
				data.AvailabilityZone = types.StringNull()
			}
		}
		if val, ok := sourceMap["connect_directly_to_external_network"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.ConnectDirectlyToExternalNetwork = types.BoolValue(b)
			}
		} else {
			if data.ConnectDirectlyToExternalNetwork.IsUnknown() {
				data.ConnectDirectlyToExternalNetwork = types.BoolNull()
			}
		}
		if val, ok := sourceMap["cores"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Cores = types.Int64Value(int64(num))
			}
		} else {
			if data.Cores.IsUnknown() {
				data.Cores = types.Int64Null()
			}
		}
		if val, ok := sourceMap["created"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Created = types.StringValue(str)
			}
		} else {
			if data.Created.IsUnknown() {
				data.Created = types.StringNull()
			}
		}
		if val, ok := sourceMap["disk"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Disk = types.Int64Value(int64(num))
			}
		} else {
			if data.Disk.IsUnknown() {
				data.Disk = types.Int64Null()
			}
		}
		if val, ok := sourceMap["error_message"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ErrorMessage = types.StringValue(str)
			}
		} else {
			if data.ErrorMessage.IsUnknown() {
				data.ErrorMessage = types.StringNull()
			}
		}
		if val, ok := sourceMap["error_traceback"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ErrorTraceback = types.StringValue(str)
			}
		} else {
			if data.ErrorTraceback.IsUnknown() {
				data.ErrorTraceback = types.StringNull()
			}
		}
		if val, ok := sourceMap["external_address"]; ok && val != nil {
			// List of strings (or flattened objects)
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if str, ok := item.(string); ok {
						items = append(items, types.StringValue(str))
					} else if obj, ok := item.(map[string]interface{}); ok {
						// Flattening logic: extract URL or UUID
						if url, ok := obj["url"].(string); ok {
							parts := strings.Split(strings.TrimRight(url, "/"), "/")
							uuid := parts[len(parts)-1]
							items = append(items, types.StringValue(uuid))
						} else if uuid, ok := obj["uuid"].(string); ok {
							items = append(items, types.StringValue(uuid))
						} else if name, ok := obj["name"].(string); ok {
							items = append(items, types.StringValue(name))
						}
					}
				}
				listVal, _ := types.ListValue(types.StringType, items)
				data.ExternalAddress = listVal
			}
		} else {
			if data.ExternalAddress.IsUnknown() {
				data.ExternalAddress = types.ListNull(types.StringType)
			}
		}
		if val, ok := sourceMap["external_ips"]; ok && val != nil {
			// List of strings (or flattened objects)
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if str, ok := item.(string); ok {
						items = append(items, types.StringValue(str))
					} else if obj, ok := item.(map[string]interface{}); ok {
						// Flattening logic: extract URL or UUID
						if url, ok := obj["url"].(string); ok {
							parts := strings.Split(strings.TrimRight(url, "/"), "/")
							uuid := parts[len(parts)-1]
							items = append(items, types.StringValue(uuid))
						} else if uuid, ok := obj["uuid"].(string); ok {
							items = append(items, types.StringValue(uuid))
						} else if name, ok := obj["name"].(string); ok {
							items = append(items, types.StringValue(name))
						}
					}
				}
				listVal, _ := types.ListValue(types.StringType, items)
				data.ExternalIps = listVal
			}
		} else {
			if data.ExternalIps.IsUnknown() {
				data.ExternalIps = types.ListNull(types.StringType)
			}
		}
		if val, ok := sourceMap["flavor_disk"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.FlavorDisk = types.Int64Value(int64(num))
			}
		} else {
			if data.FlavorDisk.IsUnknown() {
				data.FlavorDisk = types.Int64Null()
			}
		}
		if val, ok := sourceMap["flavor_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.FlavorName = types.StringValue(str)
			}
		} else {
			if data.FlavorName.IsUnknown() {
				data.FlavorName = types.StringNull()
			}
		}
		if val, ok := sourceMap["floating_ips"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"address": types.StringType,
							"port_fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"ip_address": types.StringType,
								"subnet_id":  types.StringType,
							}}},
							"port_mac_address":   types.StringType,
							"subnet":             types.StringType,
							"subnet_cidr":        types.StringType,
							"subnet_description": types.StringType,
							"subnet_name":        types.StringType,
							"subnet_uuid":        types.StringType,
							"url":                types.StringType,
						}
						attrValues := map[string]attr.Value{
							"address": func() attr.Value {
								if v, ok := objMap["address"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"port_fixed_ips": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"ip_address": types.StringType,
								"subnet_id":  types.StringType,
							}}}.ElemType),
							"port_mac_address": func() attr.Value {
								if v, ok := objMap["port_mac_address"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet": func() attr.Value {
								if v, ok := objMap["subnet"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet_cidr": func() attr.Value {
								if v, ok := objMap["subnet_cidr"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet_description": func() attr.Value {
								if v, ok := objMap["subnet_description"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet_name": func() attr.Value {
								if v, ok := objMap["subnet_name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet_uuid": func() attr.Value {
								if v, ok := objMap["subnet_uuid"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"url": func() attr.Value {
								if v, ok := objMap["url"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
						}
						objVal, _ := types.ObjectValue(attrTypes, attrValues)
						items = append(items, objVal)
					}
				}
				listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
					"address": types.StringType,
					"port_fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"ip_address": types.StringType,
						"subnet_id":  types.StringType,
					}}},
					"port_mac_address":   types.StringType,
					"subnet":             types.StringType,
					"subnet_cidr":        types.StringType,
					"subnet_description": types.StringType,
					"subnet_name":        types.StringType,
					"subnet_uuid":        types.StringType,
					"url":                types.StringType,
				}}, items)
				data.FloatingIps = listVal
			}
		} else {
			if data.FloatingIps.IsUnknown() {
				data.FloatingIps = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"address": types.StringType,
					"port_fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"ip_address": types.StringType,
						"subnet_id":  types.StringType,
					}}},
					"port_mac_address":   types.StringType,
					"subnet":             types.StringType,
					"subnet_cidr":        types.StringType,
					"subnet_description": types.StringType,
					"subnet_name":        types.StringType,
					"subnet_uuid":        types.StringType,
					"url":                types.StringType,
				}})
			}
		}
		if val, ok := sourceMap["hypervisor_hostname"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.HypervisorHostname = types.StringValue(str)
			}
		} else {
			if data.HypervisorHostname.IsUnknown() {
				data.HypervisorHostname = types.StringNull()
			}
		}
		if val, ok := sourceMap["image_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ImageName = types.StringValue(str)
			}
		} else {
			if data.ImageName.IsUnknown() {
				data.ImageName = types.StringNull()
			}
		}
		if val, ok := sourceMap["internal_ips"]; ok && val != nil {
			// List of strings (or flattened objects)
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if str, ok := item.(string); ok {
						items = append(items, types.StringValue(str))
					} else if obj, ok := item.(map[string]interface{}); ok {
						// Flattening logic: extract URL or UUID
						if url, ok := obj["url"].(string); ok {
							parts := strings.Split(strings.TrimRight(url, "/"), "/")
							uuid := parts[len(parts)-1]
							items = append(items, types.StringValue(uuid))
						} else if uuid, ok := obj["uuid"].(string); ok {
							items = append(items, types.StringValue(uuid))
						} else if name, ok := obj["name"].(string); ok {
							items = append(items, types.StringValue(name))
						}
					}
				}
				listVal, _ := types.ListValue(types.StringType, items)
				data.InternalIps = listVal
			}
		} else {
			if data.InternalIps.IsUnknown() {
				data.InternalIps = types.ListNull(types.StringType)
			}
		}
		if val, ok := sourceMap["is_limit_based"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsLimitBased = types.BoolValue(b)
			}
		} else {
			if data.IsLimitBased.IsUnknown() {
				data.IsLimitBased = types.BoolNull()
			}
		}
		if val, ok := sourceMap["is_usage_based"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsUsageBased = types.BoolValue(b)
			}
		} else {
			if data.IsUsageBased.IsUnknown() {
				data.IsUsageBased = types.BoolNull()
			}
		}
		if val, ok := sourceMap["key_fingerprint"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.KeyFingerprint = types.StringValue(str)
			}
		} else {
			if data.KeyFingerprint.IsUnknown() {
				data.KeyFingerprint = types.StringNull()
			}
		}
		if val, ok := sourceMap["key_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.KeyName = types.StringValue(str)
			}
		} else {
			if data.KeyName.IsUnknown() {
				data.KeyName = types.StringNull()
			}
		}
		if val, ok := sourceMap["latitude"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Latitude = types.Float64Value(num)
			}
		} else {
			if data.Latitude.IsUnknown() {
				data.Latitude = types.Float64Null()
			}
		}
		if val, ok := sourceMap["longitude"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Longitude = types.Float64Value(num)
			}
		} else {
			if data.Longitude.IsUnknown() {
				data.Longitude = types.Float64Null()
			}
		}
		if val, ok := sourceMap["marketplace_category_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceCategoryName = types.StringValue(str)
			}
		} else {
			if data.MarketplaceCategoryName.IsUnknown() {
				data.MarketplaceCategoryName = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_category_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceCategoryUuid = types.StringValue(str)
			}
		} else {
			if data.MarketplaceCategoryUuid.IsUnknown() {
				data.MarketplaceCategoryUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_offering_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceOfferingName = types.StringValue(str)
			}
		} else {
			if data.MarketplaceOfferingName.IsUnknown() {
				data.MarketplaceOfferingName = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_offering_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceOfferingUuid = types.StringValue(str)
			}
		} else {
			if data.MarketplaceOfferingUuid.IsUnknown() {
				data.MarketplaceOfferingUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_plan_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplacePlanUuid = types.StringValue(str)
			}
		} else {
			if data.MarketplacePlanUuid.IsUnknown() {
				data.MarketplacePlanUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_resource_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceResourceState = types.StringValue(str)
			}
		} else {
			if data.MarketplaceResourceState.IsUnknown() {
				data.MarketplaceResourceState = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_resource_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceResourceUuid = types.StringValue(str)
			}
		} else {
			if data.MarketplaceResourceUuid.IsUnknown() {
				data.MarketplaceResourceUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["min_disk"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.MinDisk = types.Int64Value(int64(num))
			}
		} else {
			if data.MinDisk.IsUnknown() {
				data.MinDisk = types.Int64Null()
			}
		}
		if val, ok := sourceMap["min_ram"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.MinRam = types.Int64Value(int64(num))
			}
		} else {
			if data.MinRam.IsUnknown() {
				data.MinRam = types.Int64Null()
			}
		}
		if val, ok := sourceMap["modified"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Modified = types.StringValue(str)
			}
		} else {
			if data.Modified.IsUnknown() {
				data.Modified = types.StringNull()
			}
		}
		if val, ok := sourceMap["ports"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"mac_address": types.StringType,
							}}},
							"device_id":    types.StringType,
							"device_owner": types.StringType,
							"fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"ip_address": types.StringType,
								"subnet_id":  types.StringType,
							}}},
							"mac_address": types.StringType,
							"security_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"access_url":                 types.StringType,
								"backend_id":                 types.StringType,
								"created":                    types.StringType,
								"customer":                   types.StringType,
								"customer_abbreviation":      types.StringType,
								"customer_name":              types.StringType,
								"customer_native_name":       types.StringType,
								"customer_uuid":              types.StringType,
								"description":                types.StringType,
								"error_message":              types.StringType,
								"error_traceback":            types.StringType,
								"is_limit_based":             types.BoolType,
								"is_usage_based":             types.BoolType,
								"marketplace_category_name":  types.StringType,
								"marketplace_category_uuid":  types.StringType,
								"marketplace_offering_name":  types.StringType,
								"marketplace_offering_uuid":  types.StringType,
								"marketplace_plan_uuid":      types.StringType,
								"marketplace_resource_state": types.StringType,
								"marketplace_resource_uuid":  types.StringType,
								"modified":                   types.StringType,
								"name":                       types.StringType,
								"project":                    types.StringType,
								"project_name":               types.StringType,
								"project_uuid":               types.StringType,
								"resource_type":              types.StringType,
								"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"cidr":              types.StringType,
									"description":       types.StringType,
									"direction":         types.StringType,
									"ethertype":         types.StringType,
									"from_port":         types.Int64Type,
									"id":                types.Int64Type,
									"protocol":          types.StringType,
									"remote_group":      types.StringType,
									"remote_group_name": types.StringType,
									"remote_group_uuid": types.StringType,
									"to_port":           types.Int64Type,
								}}},
								"service_name":                   types.StringType,
								"service_settings":               types.StringType,
								"service_settings_error_message": types.StringType,
								"service_settings_state":         types.StringType,
								"service_settings_uuid":          types.StringType,
								"state":                          types.StringType,
								"tenant":                         types.StringType,
								"tenant_name":                    types.StringType,
								"tenant_uuid":                    types.StringType,
								"url":                            types.StringType,
							}}},
							"subnet":             types.StringType,
							"subnet_cidr":        types.StringType,
							"subnet_description": types.StringType,
							"subnet_name":        types.StringType,
							"subnet_uuid":        types.StringType,
							"url":                types.StringType,
						}
						attrValues := map[string]attr.Value{
							"allowed_address_pairs": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"mac_address": types.StringType,
							}}}.ElemType),
							"device_id": func() attr.Value {
								if v, ok := objMap["device_id"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"device_owner": func() attr.Value {
								if v, ok := objMap["device_owner"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"fixed_ips": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"ip_address": types.StringType,
								"subnet_id":  types.StringType,
							}}}.ElemType),
							"mac_address": func() attr.Value {
								if v, ok := objMap["mac_address"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"security_groups": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"access_url":                 types.StringType,
								"backend_id":                 types.StringType,
								"created":                    types.StringType,
								"customer":                   types.StringType,
								"customer_abbreviation":      types.StringType,
								"customer_name":              types.StringType,
								"customer_native_name":       types.StringType,
								"customer_uuid":              types.StringType,
								"description":                types.StringType,
								"error_message":              types.StringType,
								"error_traceback":            types.StringType,
								"is_limit_based":             types.BoolType,
								"is_usage_based":             types.BoolType,
								"marketplace_category_name":  types.StringType,
								"marketplace_category_uuid":  types.StringType,
								"marketplace_offering_name":  types.StringType,
								"marketplace_offering_uuid":  types.StringType,
								"marketplace_plan_uuid":      types.StringType,
								"marketplace_resource_state": types.StringType,
								"marketplace_resource_uuid":  types.StringType,
								"modified":                   types.StringType,
								"name":                       types.StringType,
								"project":                    types.StringType,
								"project_name":               types.StringType,
								"project_uuid":               types.StringType,
								"resource_type":              types.StringType,
								"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"cidr":              types.StringType,
									"description":       types.StringType,
									"direction":         types.StringType,
									"ethertype":         types.StringType,
									"from_port":         types.Int64Type,
									"id":                types.Int64Type,
									"protocol":          types.StringType,
									"remote_group":      types.StringType,
									"remote_group_name": types.StringType,
									"remote_group_uuid": types.StringType,
									"to_port":           types.Int64Type,
								}}},
								"service_name":                   types.StringType,
								"service_settings":               types.StringType,
								"service_settings_error_message": types.StringType,
								"service_settings_state":         types.StringType,
								"service_settings_uuid":          types.StringType,
								"state":                          types.StringType,
								"tenant":                         types.StringType,
								"tenant_name":                    types.StringType,
								"tenant_uuid":                    types.StringType,
								"url":                            types.StringType,
							}}}.ElemType),
							"subnet": func() attr.Value {
								if v, ok := objMap["subnet"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet_cidr": func() attr.Value {
								if v, ok := objMap["subnet_cidr"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet_description": func() attr.Value {
								if v, ok := objMap["subnet_description"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet_name": func() attr.Value {
								if v, ok := objMap["subnet_name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet_uuid": func() attr.Value {
								if v, ok := objMap["subnet_uuid"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"url": func() attr.Value {
								if v, ok := objMap["url"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
						}
						objVal, _ := types.ObjectValue(attrTypes, attrValues)
						items = append(items, objVal)
					}
				}
				listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
					"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"mac_address": types.StringType,
					}}},
					"device_id":    types.StringType,
					"device_owner": types.StringType,
					"fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"ip_address": types.StringType,
						"subnet_id":  types.StringType,
					}}},
					"mac_address": types.StringType,
					"security_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"access_url":                 types.StringType,
						"backend_id":                 types.StringType,
						"created":                    types.StringType,
						"customer":                   types.StringType,
						"customer_abbreviation":      types.StringType,
						"customer_name":              types.StringType,
						"customer_native_name":       types.StringType,
						"customer_uuid":              types.StringType,
						"description":                types.StringType,
						"error_message":              types.StringType,
						"error_traceback":            types.StringType,
						"is_limit_based":             types.BoolType,
						"is_usage_based":             types.BoolType,
						"marketplace_category_name":  types.StringType,
						"marketplace_category_uuid":  types.StringType,
						"marketplace_offering_name":  types.StringType,
						"marketplace_offering_uuid":  types.StringType,
						"marketplace_plan_uuid":      types.StringType,
						"marketplace_resource_state": types.StringType,
						"marketplace_resource_uuid":  types.StringType,
						"modified":                   types.StringType,
						"name":                       types.StringType,
						"project":                    types.StringType,
						"project_name":               types.StringType,
						"project_uuid":               types.StringType,
						"resource_type":              types.StringType,
						"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"cidr":              types.StringType,
							"description":       types.StringType,
							"direction":         types.StringType,
							"ethertype":         types.StringType,
							"from_port":         types.Int64Type,
							"id":                types.Int64Type,
							"protocol":          types.StringType,
							"remote_group":      types.StringType,
							"remote_group_name": types.StringType,
							"remote_group_uuid": types.StringType,
							"to_port":           types.Int64Type,
						}}},
						"service_name":                   types.StringType,
						"service_settings":               types.StringType,
						"service_settings_error_message": types.StringType,
						"service_settings_state":         types.StringType,
						"service_settings_uuid":          types.StringType,
						"state":                          types.StringType,
						"tenant":                         types.StringType,
						"tenant_name":                    types.StringType,
						"tenant_uuid":                    types.StringType,
						"url":                            types.StringType,
					}}},
					"subnet":             types.StringType,
					"subnet_cidr":        types.StringType,
					"subnet_description": types.StringType,
					"subnet_name":        types.StringType,
					"subnet_uuid":        types.StringType,
					"url":                types.StringType,
				}}, items)
				data.Ports = listVal
			}
		} else {
			if data.Ports.IsUnknown() {
				data.Ports = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"mac_address": types.StringType,
					}}},
					"device_id":    types.StringType,
					"device_owner": types.StringType,
					"fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"ip_address": types.StringType,
						"subnet_id":  types.StringType,
					}}},
					"mac_address": types.StringType,
					"security_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"access_url":                 types.StringType,
						"backend_id":                 types.StringType,
						"created":                    types.StringType,
						"customer":                   types.StringType,
						"customer_abbreviation":      types.StringType,
						"customer_name":              types.StringType,
						"customer_native_name":       types.StringType,
						"customer_uuid":              types.StringType,
						"description":                types.StringType,
						"error_message":              types.StringType,
						"error_traceback":            types.StringType,
						"is_limit_based":             types.BoolType,
						"is_usage_based":             types.BoolType,
						"marketplace_category_name":  types.StringType,
						"marketplace_category_uuid":  types.StringType,
						"marketplace_offering_name":  types.StringType,
						"marketplace_offering_uuid":  types.StringType,
						"marketplace_plan_uuid":      types.StringType,
						"marketplace_resource_state": types.StringType,
						"marketplace_resource_uuid":  types.StringType,
						"modified":                   types.StringType,
						"name":                       types.StringType,
						"project":                    types.StringType,
						"project_name":               types.StringType,
						"project_uuid":               types.StringType,
						"resource_type":              types.StringType,
						"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"cidr":              types.StringType,
							"description":       types.StringType,
							"direction":         types.StringType,
							"ethertype":         types.StringType,
							"from_port":         types.Int64Type,
							"id":                types.Int64Type,
							"protocol":          types.StringType,
							"remote_group":      types.StringType,
							"remote_group_name": types.StringType,
							"remote_group_uuid": types.StringType,
							"to_port":           types.Int64Type,
						}}},
						"service_name":                   types.StringType,
						"service_settings":               types.StringType,
						"service_settings_error_message": types.StringType,
						"service_settings_state":         types.StringType,
						"service_settings_uuid":          types.StringType,
						"state":                          types.StringType,
						"tenant":                         types.StringType,
						"tenant_name":                    types.StringType,
						"tenant_uuid":                    types.StringType,
						"url":                            types.StringType,
					}}},
					"subnet":             types.StringType,
					"subnet_cidr":        types.StringType,
					"subnet_description": types.StringType,
					"subnet_name":        types.StringType,
					"subnet_uuid":        types.StringType,
					"url":                types.StringType,
				}})
			}
		}
		if val, ok := sourceMap["ram"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Ram = types.Int64Value(int64(num))
			}
		} else {
			if data.Ram.IsUnknown() {
				data.Ram = types.Int64Null()
			}
		}
		if val, ok := sourceMap["resource_type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ResourceType = types.StringValue(str)
			}
		} else {
			if data.ResourceType.IsUnknown() {
				data.ResourceType = types.StringNull()
			}
		}
		if val, ok := sourceMap["security_groups"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"description": types.StringType,
							"name":        types.StringType,
							"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"cidr":              types.StringType,
								"description":       types.StringType,
								"direction":         types.StringType,
								"ethertype":         types.StringType,
								"from_port":         types.Int64Type,
								"id":                types.Int64Type,
								"protocol":          types.StringType,
								"remote_group_name": types.StringType,
								"remote_group_uuid": types.StringType,
								"to_port":           types.Int64Type,
							}}},
							"state": types.StringType,
							"url":   types.StringType,
						}
						attrValues := map[string]attr.Value{
							"description": func() attr.Value {
								if v, ok := objMap["description"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"name": func() attr.Value {
								if v, ok := objMap["name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"rules": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"cidr":              types.StringType,
								"description":       types.StringType,
								"direction":         types.StringType,
								"ethertype":         types.StringType,
								"from_port":         types.Int64Type,
								"id":                types.Int64Type,
								"protocol":          types.StringType,
								"remote_group_name": types.StringType,
								"remote_group_uuid": types.StringType,
								"to_port":           types.Int64Type,
							}}}.ElemType),
							"state": func() attr.Value {
								if v, ok := objMap["state"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"url": func() attr.Value {
								if v, ok := objMap["url"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
						}
						objVal, _ := types.ObjectValue(attrTypes, attrValues)
						items = append(items, objVal)
					}
				}
				listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
					"description": types.StringType,
					"name":        types.StringType,
					"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"cidr":              types.StringType,
						"description":       types.StringType,
						"direction":         types.StringType,
						"ethertype":         types.StringType,
						"from_port":         types.Int64Type,
						"id":                types.Int64Type,
						"protocol":          types.StringType,
						"remote_group_name": types.StringType,
						"remote_group_uuid": types.StringType,
						"to_port":           types.Int64Type,
					}}},
					"state": types.StringType,
					"url":   types.StringType,
				}}, items)
				data.SecurityGroups = listVal
			}
		} else {
			if data.SecurityGroups.IsUnknown() {
				data.SecurityGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"description": types.StringType,
					"name":        types.StringType,
					"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"cidr":              types.StringType,
						"description":       types.StringType,
						"direction":         types.StringType,
						"ethertype":         types.StringType,
						"from_port":         types.Int64Type,
						"id":                types.Int64Type,
						"protocol":          types.StringType,
						"remote_group_name": types.StringType,
						"remote_group_uuid": types.StringType,
						"to_port":           types.Int64Type,
					}}},
					"state": types.StringType,
					"url":   types.StringType,
				}})
			}
		}
		if val, ok := sourceMap["server_group"]; ok && val != nil {
			// Nested object
			if objMap, ok := val.(map[string]interface{}); ok {
				_ = objMap // Avoid unused variable if properties are empty
				attrTypes := map[string]attr.Type{
					"name":   types.StringType,
					"policy": types.StringType,
					"state":  types.StringType,
					"url":    types.StringType,
				}
				attrValues := map[string]attr.Value{
					"name": func() attr.Value {
						if v, ok := objMap["name"].(string); ok {
							return types.StringValue(v)
						}
						return types.StringNull()
					}(),
					"policy": func() attr.Value {
						if v, ok := objMap["policy"].(string); ok {
							return types.StringValue(v)
						}
						return types.StringNull()
					}(),
					"state": func() attr.Value {
						if v, ok := objMap["state"].(string); ok {
							return types.StringValue(v)
						}
						return types.StringNull()
					}(),
					"url": func() attr.Value {
						if v, ok := objMap["url"].(string); ok {
							return types.StringValue(v)
						}
						return types.StringNull()
					}(),
				}
				objVal, _ := types.ObjectValue(attrTypes, attrValues)
				data.ServerGroup = objVal
			}
		} else {
			if data.ServerGroup.IsUnknown() {
				data.ServerGroup = types.ObjectNull(map[string]attr.Type{
					"name":   types.StringType,
					"policy": types.StringType,
					"state":  types.StringType,
					"url":    types.StringType,
				})
			}
		}
		if val, ok := sourceMap["service_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceName = types.StringValue(str)
			}
		} else {
			if data.ServiceName.IsUnknown() {
				data.ServiceName = types.StringNull()
			}
		}
		if val, ok := sourceMap["service_settings"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettings = types.StringValue(str)
			}
		} else {
			if data.ServiceSettings.IsUnknown() {
				data.ServiceSettings = types.StringNull()
			}
		}
		if val, ok := sourceMap["service_settings_error_message"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsErrorMessage = types.StringValue(str)
			}
		} else {
			if data.ServiceSettingsErrorMessage.IsUnknown() {
				data.ServiceSettingsErrorMessage = types.StringNull()
			}
		}
		if val, ok := sourceMap["service_settings_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsState = types.StringValue(str)
			}
		} else {
			if data.ServiceSettingsState.IsUnknown() {
				data.ServiceSettingsState = types.StringNull()
			}
		}
		if val, ok := sourceMap["start_time"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.StartTime = types.StringValue(str)
			}
		} else {
			if data.StartTime.IsUnknown() {
				data.StartTime = types.StringNull()
			}
		}
		if val, ok := sourceMap["url"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Url = types.StringValue(str)
			}
		} else {
			if data.Url.IsUnknown() {
				data.Url = types.StringNull()
			}
		}
		if val, ok := sourceMap["user_data"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.UserData = types.StringValue(str)
			}
		} else {
			if data.UserData.IsUnknown() {
				data.UserData = types.StringNull()
			}
		}
		if val, ok := sourceMap["volumes"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"bootable":                  types.BoolType,
							"device":                    types.StringType,
							"image_name":                types.StringType,
							"marketplace_resource_uuid": types.StringType,
							"name":                      types.StringType,
							"resource_type":             types.StringType,
							"size":                      types.Int64Type,
							"state":                     types.StringType,
							"type":                      types.StringType,
							"type_name":                 types.StringType,
							"url":                       types.StringType,
						}
						attrValues := map[string]attr.Value{
							"bootable": func() attr.Value {
								if v, ok := objMap["bootable"].(bool); ok {
									return types.BoolValue(v)
								}
								return types.BoolNull()
							}(),
							"device": func() attr.Value {
								if v, ok := objMap["device"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"image_name": func() attr.Value {
								if v, ok := objMap["image_name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"marketplace_resource_uuid": func() attr.Value {
								if v, ok := objMap["marketplace_resource_uuid"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"name": func() attr.Value {
								if v, ok := objMap["name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"resource_type": func() attr.Value {
								if v, ok := objMap["resource_type"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"size": func() attr.Value {
								if v, ok := objMap["size"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"state": func() attr.Value {
								if v, ok := objMap["state"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"type": func() attr.Value {
								if v, ok := objMap["type"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"type_name": func() attr.Value {
								if v, ok := objMap["type_name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"url": func() attr.Value {
								if v, ok := objMap["url"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
						}
						objVal, _ := types.ObjectValue(attrTypes, attrValues)
						items = append(items, objVal)
					}
				}
				listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
					"bootable":                  types.BoolType,
					"device":                    types.StringType,
					"image_name":                types.StringType,
					"marketplace_resource_uuid": types.StringType,
					"name":                      types.StringType,
					"resource_type":             types.StringType,
					"size":                      types.Int64Type,
					"state":                     types.StringType,
					"type":                      types.StringType,
					"type_name":                 types.StringType,
					"url":                       types.StringType,
				}}, items)
				data.Volumes = listVal
			}
		} else {
			if data.Volumes.IsUnknown() {
				data.Volumes = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"bootable":                  types.BoolType,
					"device":                    types.StringType,
					"image_name":                types.StringType,
					"marketplace_resource_uuid": types.StringType,
					"name":                      types.StringType,
					"resource_type":             types.StringType,
					"size":                      types.Int64Type,
					"state":                     types.StringType,
					"type":                      types.StringType,
					"type_name":                 types.StringType,
					"url":                       types.StringType,
				}})
			}
		}
		if val, ok := sourceMap["attach_volume_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AttachVolumeUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["availability_zone_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AvailabilityZoneName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["backend_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BackendId = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["can_manage"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanManage = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["customer"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Customer = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_abbreviation"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerAbbreviation = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_native_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerNativeName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Description = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["external_ip"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ExternalIp = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Name = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name_exact"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NameExact = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Project = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["query"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Query = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["runtime_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.RuntimeState = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_settings_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_settings_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.State = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["tenant"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Tenant = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["tenant_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.TenantUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Uuid = types.StringValue(str)
			}
		}

	} else {
		// Filter by provided parameters
		var results []map[string]interface{}

		filters := map[string]string{}
		if !data.AttachVolumeUuid.IsNull() {
			filters["attach_volume_uuid"] = data.AttachVolumeUuid.ValueString()
		}
		if !data.AvailabilityZoneName.IsNull() {
			filters["availability_zone_name"] = data.AvailabilityZoneName.ValueString()
		}
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
		if !data.ExternalIp.IsNull() {
			filters["external_ip"] = data.ExternalIp.ValueString()
		}
		if !data.Name.IsNull() {
			filters["name"] = data.Name.ValueString()
		}
		if !data.NameExact.IsNull() {
			filters["name_exact"] = data.NameExact.ValueString()
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
		if !data.Query.IsNull() {
			filters["query"] = data.Query.ValueString()
		}
		if !data.RuntimeState.IsNull() {
			filters["runtime_state"] = data.RuntimeState.ValueString()
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
				"At least one filter parameter (or 'id') must be provided to lookup openstack_instance.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/openstack-instances/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Openstack Instance",
				"An error occurred while filtering Openstack Instance: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Instance Not Found",
				"No Openstack Instance found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Instances Found",
				fmt.Sprintf("Found %d Openstack Instances with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		// Extract data from the single result
		if uuid, ok := results[0]["uuid"].(string); ok {
			data.UUID = types.StringValue(uuid)
		}
		sourceMap := results[0]
		// Map response fields to data model
		_ = sourceMap
		if val, ok := sourceMap["access_url"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AccessUrl = types.StringValue(str)
			}
		} else {
			if data.AccessUrl.IsUnknown() {
				data.AccessUrl = types.StringNull()
			}
		}
		if val, ok := sourceMap["action"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Action = types.StringValue(str)
			}
		} else {
			if data.Action.IsUnknown() {
				data.Action = types.StringNull()
			}
		}
		if val, ok := sourceMap["availability_zone"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AvailabilityZone = types.StringValue(str)
			}
		} else {
			if data.AvailabilityZone.IsUnknown() {
				data.AvailabilityZone = types.StringNull()
			}
		}
		if val, ok := sourceMap["connect_directly_to_external_network"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.ConnectDirectlyToExternalNetwork = types.BoolValue(b)
			}
		} else {
			if data.ConnectDirectlyToExternalNetwork.IsUnknown() {
				data.ConnectDirectlyToExternalNetwork = types.BoolNull()
			}
		}
		if val, ok := sourceMap["cores"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Cores = types.Int64Value(int64(num))
			}
		} else {
			if data.Cores.IsUnknown() {
				data.Cores = types.Int64Null()
			}
		}
		if val, ok := sourceMap["created"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Created = types.StringValue(str)
			}
		} else {
			if data.Created.IsUnknown() {
				data.Created = types.StringNull()
			}
		}
		if val, ok := sourceMap["disk"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Disk = types.Int64Value(int64(num))
			}
		} else {
			if data.Disk.IsUnknown() {
				data.Disk = types.Int64Null()
			}
		}
		if val, ok := sourceMap["error_message"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ErrorMessage = types.StringValue(str)
			}
		} else {
			if data.ErrorMessage.IsUnknown() {
				data.ErrorMessage = types.StringNull()
			}
		}
		if val, ok := sourceMap["error_traceback"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ErrorTraceback = types.StringValue(str)
			}
		} else {
			if data.ErrorTraceback.IsUnknown() {
				data.ErrorTraceback = types.StringNull()
			}
		}
		if val, ok := sourceMap["external_address"]; ok && val != nil {
			// List of strings (or flattened objects)
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if str, ok := item.(string); ok {
						items = append(items, types.StringValue(str))
					} else if obj, ok := item.(map[string]interface{}); ok {
						// Flattening logic: extract URL or UUID
						if url, ok := obj["url"].(string); ok {
							parts := strings.Split(strings.TrimRight(url, "/"), "/")
							uuid := parts[len(parts)-1]
							items = append(items, types.StringValue(uuid))
						} else if uuid, ok := obj["uuid"].(string); ok {
							items = append(items, types.StringValue(uuid))
						} else if name, ok := obj["name"].(string); ok {
							items = append(items, types.StringValue(name))
						}
					}
				}
				listVal, _ := types.ListValue(types.StringType, items)
				data.ExternalAddress = listVal
			}
		} else {
			if data.ExternalAddress.IsUnknown() {
				data.ExternalAddress = types.ListNull(types.StringType)
			}
		}
		if val, ok := sourceMap["external_ips"]; ok && val != nil {
			// List of strings (or flattened objects)
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if str, ok := item.(string); ok {
						items = append(items, types.StringValue(str))
					} else if obj, ok := item.(map[string]interface{}); ok {
						// Flattening logic: extract URL or UUID
						if url, ok := obj["url"].(string); ok {
							parts := strings.Split(strings.TrimRight(url, "/"), "/")
							uuid := parts[len(parts)-1]
							items = append(items, types.StringValue(uuid))
						} else if uuid, ok := obj["uuid"].(string); ok {
							items = append(items, types.StringValue(uuid))
						} else if name, ok := obj["name"].(string); ok {
							items = append(items, types.StringValue(name))
						}
					}
				}
				listVal, _ := types.ListValue(types.StringType, items)
				data.ExternalIps = listVal
			}
		} else {
			if data.ExternalIps.IsUnknown() {
				data.ExternalIps = types.ListNull(types.StringType)
			}
		}
		if val, ok := sourceMap["flavor_disk"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.FlavorDisk = types.Int64Value(int64(num))
			}
		} else {
			if data.FlavorDisk.IsUnknown() {
				data.FlavorDisk = types.Int64Null()
			}
		}
		if val, ok := sourceMap["flavor_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.FlavorName = types.StringValue(str)
			}
		} else {
			if data.FlavorName.IsUnknown() {
				data.FlavorName = types.StringNull()
			}
		}
		if val, ok := sourceMap["floating_ips"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"address": types.StringType,
							"port_fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"ip_address": types.StringType,
								"subnet_id":  types.StringType,
							}}},
							"port_mac_address":   types.StringType,
							"subnet":             types.StringType,
							"subnet_cidr":        types.StringType,
							"subnet_description": types.StringType,
							"subnet_name":        types.StringType,
							"subnet_uuid":        types.StringType,
							"url":                types.StringType,
						}
						attrValues := map[string]attr.Value{
							"address": func() attr.Value {
								if v, ok := objMap["address"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"port_fixed_ips": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"ip_address": types.StringType,
								"subnet_id":  types.StringType,
							}}}.ElemType),
							"port_mac_address": func() attr.Value {
								if v, ok := objMap["port_mac_address"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet": func() attr.Value {
								if v, ok := objMap["subnet"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet_cidr": func() attr.Value {
								if v, ok := objMap["subnet_cidr"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet_description": func() attr.Value {
								if v, ok := objMap["subnet_description"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet_name": func() attr.Value {
								if v, ok := objMap["subnet_name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet_uuid": func() attr.Value {
								if v, ok := objMap["subnet_uuid"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"url": func() attr.Value {
								if v, ok := objMap["url"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
						}
						objVal, _ := types.ObjectValue(attrTypes, attrValues)
						items = append(items, objVal)
					}
				}
				listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
					"address": types.StringType,
					"port_fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"ip_address": types.StringType,
						"subnet_id":  types.StringType,
					}}},
					"port_mac_address":   types.StringType,
					"subnet":             types.StringType,
					"subnet_cidr":        types.StringType,
					"subnet_description": types.StringType,
					"subnet_name":        types.StringType,
					"subnet_uuid":        types.StringType,
					"url":                types.StringType,
				}}, items)
				data.FloatingIps = listVal
			}
		} else {
			if data.FloatingIps.IsUnknown() {
				data.FloatingIps = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"address": types.StringType,
					"port_fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"ip_address": types.StringType,
						"subnet_id":  types.StringType,
					}}},
					"port_mac_address":   types.StringType,
					"subnet":             types.StringType,
					"subnet_cidr":        types.StringType,
					"subnet_description": types.StringType,
					"subnet_name":        types.StringType,
					"subnet_uuid":        types.StringType,
					"url":                types.StringType,
				}})
			}
		}
		if val, ok := sourceMap["hypervisor_hostname"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.HypervisorHostname = types.StringValue(str)
			}
		} else {
			if data.HypervisorHostname.IsUnknown() {
				data.HypervisorHostname = types.StringNull()
			}
		}
		if val, ok := sourceMap["image_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ImageName = types.StringValue(str)
			}
		} else {
			if data.ImageName.IsUnknown() {
				data.ImageName = types.StringNull()
			}
		}
		if val, ok := sourceMap["internal_ips"]; ok && val != nil {
			// List of strings (or flattened objects)
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if str, ok := item.(string); ok {
						items = append(items, types.StringValue(str))
					} else if obj, ok := item.(map[string]interface{}); ok {
						// Flattening logic: extract URL or UUID
						if url, ok := obj["url"].(string); ok {
							parts := strings.Split(strings.TrimRight(url, "/"), "/")
							uuid := parts[len(parts)-1]
							items = append(items, types.StringValue(uuid))
						} else if uuid, ok := obj["uuid"].(string); ok {
							items = append(items, types.StringValue(uuid))
						} else if name, ok := obj["name"].(string); ok {
							items = append(items, types.StringValue(name))
						}
					}
				}
				listVal, _ := types.ListValue(types.StringType, items)
				data.InternalIps = listVal
			}
		} else {
			if data.InternalIps.IsUnknown() {
				data.InternalIps = types.ListNull(types.StringType)
			}
		}
		if val, ok := sourceMap["is_limit_based"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsLimitBased = types.BoolValue(b)
			}
		} else {
			if data.IsLimitBased.IsUnknown() {
				data.IsLimitBased = types.BoolNull()
			}
		}
		if val, ok := sourceMap["is_usage_based"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsUsageBased = types.BoolValue(b)
			}
		} else {
			if data.IsUsageBased.IsUnknown() {
				data.IsUsageBased = types.BoolNull()
			}
		}
		if val, ok := sourceMap["key_fingerprint"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.KeyFingerprint = types.StringValue(str)
			}
		} else {
			if data.KeyFingerprint.IsUnknown() {
				data.KeyFingerprint = types.StringNull()
			}
		}
		if val, ok := sourceMap["key_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.KeyName = types.StringValue(str)
			}
		} else {
			if data.KeyName.IsUnknown() {
				data.KeyName = types.StringNull()
			}
		}
		if val, ok := sourceMap["latitude"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Latitude = types.Float64Value(num)
			}
		} else {
			if data.Latitude.IsUnknown() {
				data.Latitude = types.Float64Null()
			}
		}
		if val, ok := sourceMap["longitude"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Longitude = types.Float64Value(num)
			}
		} else {
			if data.Longitude.IsUnknown() {
				data.Longitude = types.Float64Null()
			}
		}
		if val, ok := sourceMap["marketplace_category_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceCategoryName = types.StringValue(str)
			}
		} else {
			if data.MarketplaceCategoryName.IsUnknown() {
				data.MarketplaceCategoryName = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_category_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceCategoryUuid = types.StringValue(str)
			}
		} else {
			if data.MarketplaceCategoryUuid.IsUnknown() {
				data.MarketplaceCategoryUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_offering_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceOfferingName = types.StringValue(str)
			}
		} else {
			if data.MarketplaceOfferingName.IsUnknown() {
				data.MarketplaceOfferingName = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_offering_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceOfferingUuid = types.StringValue(str)
			}
		} else {
			if data.MarketplaceOfferingUuid.IsUnknown() {
				data.MarketplaceOfferingUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_plan_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplacePlanUuid = types.StringValue(str)
			}
		} else {
			if data.MarketplacePlanUuid.IsUnknown() {
				data.MarketplacePlanUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_resource_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceResourceState = types.StringValue(str)
			}
		} else {
			if data.MarketplaceResourceState.IsUnknown() {
				data.MarketplaceResourceState = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_resource_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceResourceUuid = types.StringValue(str)
			}
		} else {
			if data.MarketplaceResourceUuid.IsUnknown() {
				data.MarketplaceResourceUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["min_disk"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.MinDisk = types.Int64Value(int64(num))
			}
		} else {
			if data.MinDisk.IsUnknown() {
				data.MinDisk = types.Int64Null()
			}
		}
		if val, ok := sourceMap["min_ram"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.MinRam = types.Int64Value(int64(num))
			}
		} else {
			if data.MinRam.IsUnknown() {
				data.MinRam = types.Int64Null()
			}
		}
		if val, ok := sourceMap["modified"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Modified = types.StringValue(str)
			}
		} else {
			if data.Modified.IsUnknown() {
				data.Modified = types.StringNull()
			}
		}
		if val, ok := sourceMap["ports"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"mac_address": types.StringType,
							}}},
							"device_id":    types.StringType,
							"device_owner": types.StringType,
							"fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"ip_address": types.StringType,
								"subnet_id":  types.StringType,
							}}},
							"mac_address": types.StringType,
							"security_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"access_url":                 types.StringType,
								"backend_id":                 types.StringType,
								"created":                    types.StringType,
								"customer":                   types.StringType,
								"customer_abbreviation":      types.StringType,
								"customer_name":              types.StringType,
								"customer_native_name":       types.StringType,
								"customer_uuid":              types.StringType,
								"description":                types.StringType,
								"error_message":              types.StringType,
								"error_traceback":            types.StringType,
								"is_limit_based":             types.BoolType,
								"is_usage_based":             types.BoolType,
								"marketplace_category_name":  types.StringType,
								"marketplace_category_uuid":  types.StringType,
								"marketplace_offering_name":  types.StringType,
								"marketplace_offering_uuid":  types.StringType,
								"marketplace_plan_uuid":      types.StringType,
								"marketplace_resource_state": types.StringType,
								"marketplace_resource_uuid":  types.StringType,
								"modified":                   types.StringType,
								"name":                       types.StringType,
								"project":                    types.StringType,
								"project_name":               types.StringType,
								"project_uuid":               types.StringType,
								"resource_type":              types.StringType,
								"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"cidr":              types.StringType,
									"description":       types.StringType,
									"direction":         types.StringType,
									"ethertype":         types.StringType,
									"from_port":         types.Int64Type,
									"id":                types.Int64Type,
									"protocol":          types.StringType,
									"remote_group":      types.StringType,
									"remote_group_name": types.StringType,
									"remote_group_uuid": types.StringType,
									"to_port":           types.Int64Type,
								}}},
								"service_name":                   types.StringType,
								"service_settings":               types.StringType,
								"service_settings_error_message": types.StringType,
								"service_settings_state":         types.StringType,
								"service_settings_uuid":          types.StringType,
								"state":                          types.StringType,
								"tenant":                         types.StringType,
								"tenant_name":                    types.StringType,
								"tenant_uuid":                    types.StringType,
								"url":                            types.StringType,
							}}},
							"subnet":             types.StringType,
							"subnet_cidr":        types.StringType,
							"subnet_description": types.StringType,
							"subnet_name":        types.StringType,
							"subnet_uuid":        types.StringType,
							"url":                types.StringType,
						}
						attrValues := map[string]attr.Value{
							"allowed_address_pairs": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"mac_address": types.StringType,
							}}}.ElemType),
							"device_id": func() attr.Value {
								if v, ok := objMap["device_id"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"device_owner": func() attr.Value {
								if v, ok := objMap["device_owner"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"fixed_ips": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"ip_address": types.StringType,
								"subnet_id":  types.StringType,
							}}}.ElemType),
							"mac_address": func() attr.Value {
								if v, ok := objMap["mac_address"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"security_groups": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"access_url":                 types.StringType,
								"backend_id":                 types.StringType,
								"created":                    types.StringType,
								"customer":                   types.StringType,
								"customer_abbreviation":      types.StringType,
								"customer_name":              types.StringType,
								"customer_native_name":       types.StringType,
								"customer_uuid":              types.StringType,
								"description":                types.StringType,
								"error_message":              types.StringType,
								"error_traceback":            types.StringType,
								"is_limit_based":             types.BoolType,
								"is_usage_based":             types.BoolType,
								"marketplace_category_name":  types.StringType,
								"marketplace_category_uuid":  types.StringType,
								"marketplace_offering_name":  types.StringType,
								"marketplace_offering_uuid":  types.StringType,
								"marketplace_plan_uuid":      types.StringType,
								"marketplace_resource_state": types.StringType,
								"marketplace_resource_uuid":  types.StringType,
								"modified":                   types.StringType,
								"name":                       types.StringType,
								"project":                    types.StringType,
								"project_name":               types.StringType,
								"project_uuid":               types.StringType,
								"resource_type":              types.StringType,
								"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"cidr":              types.StringType,
									"description":       types.StringType,
									"direction":         types.StringType,
									"ethertype":         types.StringType,
									"from_port":         types.Int64Type,
									"id":                types.Int64Type,
									"protocol":          types.StringType,
									"remote_group":      types.StringType,
									"remote_group_name": types.StringType,
									"remote_group_uuid": types.StringType,
									"to_port":           types.Int64Type,
								}}},
								"service_name":                   types.StringType,
								"service_settings":               types.StringType,
								"service_settings_error_message": types.StringType,
								"service_settings_state":         types.StringType,
								"service_settings_uuid":          types.StringType,
								"state":                          types.StringType,
								"tenant":                         types.StringType,
								"tenant_name":                    types.StringType,
								"tenant_uuid":                    types.StringType,
								"url":                            types.StringType,
							}}}.ElemType),
							"subnet": func() attr.Value {
								if v, ok := objMap["subnet"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet_cidr": func() attr.Value {
								if v, ok := objMap["subnet_cidr"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet_description": func() attr.Value {
								if v, ok := objMap["subnet_description"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet_name": func() attr.Value {
								if v, ok := objMap["subnet_name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"subnet_uuid": func() attr.Value {
								if v, ok := objMap["subnet_uuid"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"url": func() attr.Value {
								if v, ok := objMap["url"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
						}
						objVal, _ := types.ObjectValue(attrTypes, attrValues)
						items = append(items, objVal)
					}
				}
				listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
					"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"mac_address": types.StringType,
					}}},
					"device_id":    types.StringType,
					"device_owner": types.StringType,
					"fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"ip_address": types.StringType,
						"subnet_id":  types.StringType,
					}}},
					"mac_address": types.StringType,
					"security_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"access_url":                 types.StringType,
						"backend_id":                 types.StringType,
						"created":                    types.StringType,
						"customer":                   types.StringType,
						"customer_abbreviation":      types.StringType,
						"customer_name":              types.StringType,
						"customer_native_name":       types.StringType,
						"customer_uuid":              types.StringType,
						"description":                types.StringType,
						"error_message":              types.StringType,
						"error_traceback":            types.StringType,
						"is_limit_based":             types.BoolType,
						"is_usage_based":             types.BoolType,
						"marketplace_category_name":  types.StringType,
						"marketplace_category_uuid":  types.StringType,
						"marketplace_offering_name":  types.StringType,
						"marketplace_offering_uuid":  types.StringType,
						"marketplace_plan_uuid":      types.StringType,
						"marketplace_resource_state": types.StringType,
						"marketplace_resource_uuid":  types.StringType,
						"modified":                   types.StringType,
						"name":                       types.StringType,
						"project":                    types.StringType,
						"project_name":               types.StringType,
						"project_uuid":               types.StringType,
						"resource_type":              types.StringType,
						"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"cidr":              types.StringType,
							"description":       types.StringType,
							"direction":         types.StringType,
							"ethertype":         types.StringType,
							"from_port":         types.Int64Type,
							"id":                types.Int64Type,
							"protocol":          types.StringType,
							"remote_group":      types.StringType,
							"remote_group_name": types.StringType,
							"remote_group_uuid": types.StringType,
							"to_port":           types.Int64Type,
						}}},
						"service_name":                   types.StringType,
						"service_settings":               types.StringType,
						"service_settings_error_message": types.StringType,
						"service_settings_state":         types.StringType,
						"service_settings_uuid":          types.StringType,
						"state":                          types.StringType,
						"tenant":                         types.StringType,
						"tenant_name":                    types.StringType,
						"tenant_uuid":                    types.StringType,
						"url":                            types.StringType,
					}}},
					"subnet":             types.StringType,
					"subnet_cidr":        types.StringType,
					"subnet_description": types.StringType,
					"subnet_name":        types.StringType,
					"subnet_uuid":        types.StringType,
					"url":                types.StringType,
				}}, items)
				data.Ports = listVal
			}
		} else {
			if data.Ports.IsUnknown() {
				data.Ports = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"mac_address": types.StringType,
					}}},
					"device_id":    types.StringType,
					"device_owner": types.StringType,
					"fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"ip_address": types.StringType,
						"subnet_id":  types.StringType,
					}}},
					"mac_address": types.StringType,
					"security_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"access_url":                 types.StringType,
						"backend_id":                 types.StringType,
						"created":                    types.StringType,
						"customer":                   types.StringType,
						"customer_abbreviation":      types.StringType,
						"customer_name":              types.StringType,
						"customer_native_name":       types.StringType,
						"customer_uuid":              types.StringType,
						"description":                types.StringType,
						"error_message":              types.StringType,
						"error_traceback":            types.StringType,
						"is_limit_based":             types.BoolType,
						"is_usage_based":             types.BoolType,
						"marketplace_category_name":  types.StringType,
						"marketplace_category_uuid":  types.StringType,
						"marketplace_offering_name":  types.StringType,
						"marketplace_offering_uuid":  types.StringType,
						"marketplace_plan_uuid":      types.StringType,
						"marketplace_resource_state": types.StringType,
						"marketplace_resource_uuid":  types.StringType,
						"modified":                   types.StringType,
						"name":                       types.StringType,
						"project":                    types.StringType,
						"project_name":               types.StringType,
						"project_uuid":               types.StringType,
						"resource_type":              types.StringType,
						"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"cidr":              types.StringType,
							"description":       types.StringType,
							"direction":         types.StringType,
							"ethertype":         types.StringType,
							"from_port":         types.Int64Type,
							"id":                types.Int64Type,
							"protocol":          types.StringType,
							"remote_group":      types.StringType,
							"remote_group_name": types.StringType,
							"remote_group_uuid": types.StringType,
							"to_port":           types.Int64Type,
						}}},
						"service_name":                   types.StringType,
						"service_settings":               types.StringType,
						"service_settings_error_message": types.StringType,
						"service_settings_state":         types.StringType,
						"service_settings_uuid":          types.StringType,
						"state":                          types.StringType,
						"tenant":                         types.StringType,
						"tenant_name":                    types.StringType,
						"tenant_uuid":                    types.StringType,
						"url":                            types.StringType,
					}}},
					"subnet":             types.StringType,
					"subnet_cidr":        types.StringType,
					"subnet_description": types.StringType,
					"subnet_name":        types.StringType,
					"subnet_uuid":        types.StringType,
					"url":                types.StringType,
				}})
			}
		}
		if val, ok := sourceMap["ram"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Ram = types.Int64Value(int64(num))
			}
		} else {
			if data.Ram.IsUnknown() {
				data.Ram = types.Int64Null()
			}
		}
		if val, ok := sourceMap["resource_type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ResourceType = types.StringValue(str)
			}
		} else {
			if data.ResourceType.IsUnknown() {
				data.ResourceType = types.StringNull()
			}
		}
		if val, ok := sourceMap["security_groups"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"description": types.StringType,
							"name":        types.StringType,
							"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"cidr":              types.StringType,
								"description":       types.StringType,
								"direction":         types.StringType,
								"ethertype":         types.StringType,
								"from_port":         types.Int64Type,
								"id":                types.Int64Type,
								"protocol":          types.StringType,
								"remote_group_name": types.StringType,
								"remote_group_uuid": types.StringType,
								"to_port":           types.Int64Type,
							}}},
							"state": types.StringType,
							"url":   types.StringType,
						}
						attrValues := map[string]attr.Value{
							"description": func() attr.Value {
								if v, ok := objMap["description"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"name": func() attr.Value {
								if v, ok := objMap["name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"rules": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"cidr":              types.StringType,
								"description":       types.StringType,
								"direction":         types.StringType,
								"ethertype":         types.StringType,
								"from_port":         types.Int64Type,
								"id":                types.Int64Type,
								"protocol":          types.StringType,
								"remote_group_name": types.StringType,
								"remote_group_uuid": types.StringType,
								"to_port":           types.Int64Type,
							}}}.ElemType),
							"state": func() attr.Value {
								if v, ok := objMap["state"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"url": func() attr.Value {
								if v, ok := objMap["url"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
						}
						objVal, _ := types.ObjectValue(attrTypes, attrValues)
						items = append(items, objVal)
					}
				}
				listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
					"description": types.StringType,
					"name":        types.StringType,
					"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"cidr":              types.StringType,
						"description":       types.StringType,
						"direction":         types.StringType,
						"ethertype":         types.StringType,
						"from_port":         types.Int64Type,
						"id":                types.Int64Type,
						"protocol":          types.StringType,
						"remote_group_name": types.StringType,
						"remote_group_uuid": types.StringType,
						"to_port":           types.Int64Type,
					}}},
					"state": types.StringType,
					"url":   types.StringType,
				}}, items)
				data.SecurityGroups = listVal
			}
		} else {
			if data.SecurityGroups.IsUnknown() {
				data.SecurityGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"description": types.StringType,
					"name":        types.StringType,
					"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"cidr":              types.StringType,
						"description":       types.StringType,
						"direction":         types.StringType,
						"ethertype":         types.StringType,
						"from_port":         types.Int64Type,
						"id":                types.Int64Type,
						"protocol":          types.StringType,
						"remote_group_name": types.StringType,
						"remote_group_uuid": types.StringType,
						"to_port":           types.Int64Type,
					}}},
					"state": types.StringType,
					"url":   types.StringType,
				}})
			}
		}
		if val, ok := sourceMap["server_group"]; ok && val != nil {
			// Nested object
			if objMap, ok := val.(map[string]interface{}); ok {
				_ = objMap // Avoid unused variable if properties are empty
				attrTypes := map[string]attr.Type{
					"name":   types.StringType,
					"policy": types.StringType,
					"state":  types.StringType,
					"url":    types.StringType,
				}
				attrValues := map[string]attr.Value{
					"name": func() attr.Value {
						if v, ok := objMap["name"].(string); ok {
							return types.StringValue(v)
						}
						return types.StringNull()
					}(),
					"policy": func() attr.Value {
						if v, ok := objMap["policy"].(string); ok {
							return types.StringValue(v)
						}
						return types.StringNull()
					}(),
					"state": func() attr.Value {
						if v, ok := objMap["state"].(string); ok {
							return types.StringValue(v)
						}
						return types.StringNull()
					}(),
					"url": func() attr.Value {
						if v, ok := objMap["url"].(string); ok {
							return types.StringValue(v)
						}
						return types.StringNull()
					}(),
				}
				objVal, _ := types.ObjectValue(attrTypes, attrValues)
				data.ServerGroup = objVal
			}
		} else {
			if data.ServerGroup.IsUnknown() {
				data.ServerGroup = types.ObjectNull(map[string]attr.Type{
					"name":   types.StringType,
					"policy": types.StringType,
					"state":  types.StringType,
					"url":    types.StringType,
				})
			}
		}
		if val, ok := sourceMap["service_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceName = types.StringValue(str)
			}
		} else {
			if data.ServiceName.IsUnknown() {
				data.ServiceName = types.StringNull()
			}
		}
		if val, ok := sourceMap["service_settings"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettings = types.StringValue(str)
			}
		} else {
			if data.ServiceSettings.IsUnknown() {
				data.ServiceSettings = types.StringNull()
			}
		}
		if val, ok := sourceMap["service_settings_error_message"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsErrorMessage = types.StringValue(str)
			}
		} else {
			if data.ServiceSettingsErrorMessage.IsUnknown() {
				data.ServiceSettingsErrorMessage = types.StringNull()
			}
		}
		if val, ok := sourceMap["service_settings_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsState = types.StringValue(str)
			}
		} else {
			if data.ServiceSettingsState.IsUnknown() {
				data.ServiceSettingsState = types.StringNull()
			}
		}
		if val, ok := sourceMap["start_time"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.StartTime = types.StringValue(str)
			}
		} else {
			if data.StartTime.IsUnknown() {
				data.StartTime = types.StringNull()
			}
		}
		if val, ok := sourceMap["url"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Url = types.StringValue(str)
			}
		} else {
			if data.Url.IsUnknown() {
				data.Url = types.StringNull()
			}
		}
		if val, ok := sourceMap["user_data"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.UserData = types.StringValue(str)
			}
		} else {
			if data.UserData.IsUnknown() {
				data.UserData = types.StringNull()
			}
		}
		if val, ok := sourceMap["volumes"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"bootable":                  types.BoolType,
							"device":                    types.StringType,
							"image_name":                types.StringType,
							"marketplace_resource_uuid": types.StringType,
							"name":                      types.StringType,
							"resource_type":             types.StringType,
							"size":                      types.Int64Type,
							"state":                     types.StringType,
							"type":                      types.StringType,
							"type_name":                 types.StringType,
							"url":                       types.StringType,
						}
						attrValues := map[string]attr.Value{
							"bootable": func() attr.Value {
								if v, ok := objMap["bootable"].(bool); ok {
									return types.BoolValue(v)
								}
								return types.BoolNull()
							}(),
							"device": func() attr.Value {
								if v, ok := objMap["device"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"image_name": func() attr.Value {
								if v, ok := objMap["image_name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"marketplace_resource_uuid": func() attr.Value {
								if v, ok := objMap["marketplace_resource_uuid"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"name": func() attr.Value {
								if v, ok := objMap["name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"resource_type": func() attr.Value {
								if v, ok := objMap["resource_type"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"size": func() attr.Value {
								if v, ok := objMap["size"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"state": func() attr.Value {
								if v, ok := objMap["state"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"type": func() attr.Value {
								if v, ok := objMap["type"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"type_name": func() attr.Value {
								if v, ok := objMap["type_name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"url": func() attr.Value {
								if v, ok := objMap["url"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
						}
						objVal, _ := types.ObjectValue(attrTypes, attrValues)
						items = append(items, objVal)
					}
				}
				listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
					"bootable":                  types.BoolType,
					"device":                    types.StringType,
					"image_name":                types.StringType,
					"marketplace_resource_uuid": types.StringType,
					"name":                      types.StringType,
					"resource_type":             types.StringType,
					"size":                      types.Int64Type,
					"state":                     types.StringType,
					"type":                      types.StringType,
					"type_name":                 types.StringType,
					"url":                       types.StringType,
				}}, items)
				data.Volumes = listVal
			}
		} else {
			if data.Volumes.IsUnknown() {
				data.Volumes = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"bootable":                  types.BoolType,
					"device":                    types.StringType,
					"image_name":                types.StringType,
					"marketplace_resource_uuid": types.StringType,
					"name":                      types.StringType,
					"resource_type":             types.StringType,
					"size":                      types.Int64Type,
					"state":                     types.StringType,
					"type":                      types.StringType,
					"type_name":                 types.StringType,
					"url":                       types.StringType,
				}})
			}
		}
		if val, ok := sourceMap["attach_volume_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AttachVolumeUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["availability_zone_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AvailabilityZoneName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["backend_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BackendId = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["can_manage"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanManage = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["customer"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Customer = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_abbreviation"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerAbbreviation = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_native_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerNativeName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Description = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["external_ip"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ExternalIp = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Name = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name_exact"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NameExact = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Project = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["query"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Query = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["runtime_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.RuntimeState = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_settings_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_settings_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.State = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["tenant"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Tenant = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["tenant_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.TenantUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Uuid = types.StringValue(str)
			}
		}
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
