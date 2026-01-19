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
var _ datasource.DataSource = &OpenstackPortDataSource{}

func NewOpenstackPortDataSource() datasource.DataSource {
	return &OpenstackPortDataSource{}
}

// OpenstackPortDataSource defines the data source implementation.
type OpenstackPortDataSource struct {
	client *client.Client
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
		MarkdownDescription: "OpenstackPort data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"admin_state_up": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"device_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"device_owner": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
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
				MarkdownDescription: "",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
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
				MarkdownDescription: "",
			},
			"tenant": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"tenant_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"access_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"allowed_address_pairs": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"mac_address": types.StringType}}},
				Computed:            true,
				MarkdownDescription: "",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"floating_ips": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Computed:            true,
				MarkdownDescription: "",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
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
				MarkdownDescription: "",
			},
			"security_groups": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType, "url": types.StringType}}},
				Computed:            true,
				MarkdownDescription: "",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"tenant_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
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
		var item map[string]interface{}

		err := d.client.GetByUUID(ctx, "/api/openstack-ports/{uuid}/", data.UUID.ValueString(), &item)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Port",
				"An error occurred while reading the openstack_port by UUID: "+err.Error(),
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
		if val, ok := sourceMap["allowed_address_pairs"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"mac_address": types.StringType,
						}
						attrValues := map[string]attr.Value{
							"mac_address": func() attr.Value {
								if v, ok := objMap["mac_address"].(string); ok {
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
					"mac_address": types.StringType,
				}}, items)
				data.AllowedAddressPairs = listVal
			}
		} else {
			if data.AllowedAddressPairs.IsUnknown() {
				data.AllowedAddressPairs = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"mac_address": types.StringType,
				}})
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
		if val, ok := sourceMap["description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Description = types.StringValue(str)
			}
		} else {
			if data.Description.IsUnknown() {
				data.Description = types.StringNull()
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
		if val, ok := sourceMap["floating_ips"]; ok && val != nil {
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
				data.FloatingIps = listVal
			}
		} else {
			if data.FloatingIps.IsUnknown() {
				data.FloatingIps = types.ListNull(types.StringType)
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
		if val, ok := sourceMap["network"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Network = types.StringValue(str)
			}
		} else {
			if data.Network.IsUnknown() {
				data.Network = types.StringNull()
			}
		}
		if val, ok := sourceMap["port_security_enabled"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.PortSecurityEnabled = types.BoolValue(b)
			}
		} else {
			if data.PortSecurityEnabled.IsUnknown() {
				data.PortSecurityEnabled = types.BoolNull()
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
							"name": types.StringType,
							"url":  types.StringType,
						}
						attrValues := map[string]attr.Value{
							"name": func() attr.Value {
								if v, ok := objMap["name"].(string); ok {
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
					"name": types.StringType,
					"url":  types.StringType,
				}}, items)
				data.SecurityGroups = listVal
			}
		} else {
			if data.SecurityGroups.IsUnknown() {
				data.SecurityGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"name": types.StringType,
					"url":  types.StringType,
				}})
			}
		}
		if val, ok := sourceMap["state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.State = types.StringValue(str)
			}
		} else {
			if data.State.IsUnknown() {
				data.State = types.StringNull()
			}
		}
		if val, ok := sourceMap["tenant_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.TenantName = types.StringValue(str)
			}
		} else {
			if data.TenantName.IsUnknown() {
				data.TenantName = types.StringNull()
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

		// Map filter parameters from response if available
		if val, ok := sourceMap["admin_state_up"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.AdminStateUp = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["backend_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BackendId = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["device_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.DeviceId = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["device_owner"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.DeviceOwner = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["exclude_subnet_uuids"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ExcludeSubnetUuids = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["fixed_ips"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.FixedIps = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["has_device_owner"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.HasDeviceOwner = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["mac_address"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MacAddress = types.StringValue(str)
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
		if val, ok := sourceMap["network_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NetworkName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["network_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NetworkUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["query"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Query = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["status"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Status = types.StringValue(str)
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

	} else {
		// Filter by provided parameters
		var results []map[string]interface{}

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
				"Unable to List Port",
				"An error occurred while filtering openstack_port: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Port Not Found",
				"No openstack_port found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Ports Found",
				fmt.Sprintf("Found %d openstack_ports with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
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
		if val, ok := sourceMap["allowed_address_pairs"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"mac_address": types.StringType,
						}
						attrValues := map[string]attr.Value{
							"mac_address": func() attr.Value {
								if v, ok := objMap["mac_address"].(string); ok {
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
					"mac_address": types.StringType,
				}}, items)
				data.AllowedAddressPairs = listVal
			}
		} else {
			if data.AllowedAddressPairs.IsUnknown() {
				data.AllowedAddressPairs = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"mac_address": types.StringType,
				}})
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
		if val, ok := sourceMap["description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Description = types.StringValue(str)
			}
		} else {
			if data.Description.IsUnknown() {
				data.Description = types.StringNull()
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
		if val, ok := sourceMap["floating_ips"]; ok && val != nil {
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
				data.FloatingIps = listVal
			}
		} else {
			if data.FloatingIps.IsUnknown() {
				data.FloatingIps = types.ListNull(types.StringType)
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
		if val, ok := sourceMap["network"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Network = types.StringValue(str)
			}
		} else {
			if data.Network.IsUnknown() {
				data.Network = types.StringNull()
			}
		}
		if val, ok := sourceMap["port_security_enabled"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.PortSecurityEnabled = types.BoolValue(b)
			}
		} else {
			if data.PortSecurityEnabled.IsUnknown() {
				data.PortSecurityEnabled = types.BoolNull()
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
							"name": types.StringType,
							"url":  types.StringType,
						}
						attrValues := map[string]attr.Value{
							"name": func() attr.Value {
								if v, ok := objMap["name"].(string); ok {
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
					"name": types.StringType,
					"url":  types.StringType,
				}}, items)
				data.SecurityGroups = listVal
			}
		} else {
			if data.SecurityGroups.IsUnknown() {
				data.SecurityGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"name": types.StringType,
					"url":  types.StringType,
				}})
			}
		}
		if val, ok := sourceMap["state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.State = types.StringValue(str)
			}
		} else {
			if data.State.IsUnknown() {
				data.State = types.StringNull()
			}
		}
		if val, ok := sourceMap["tenant_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.TenantName = types.StringValue(str)
			}
		} else {
			if data.TenantName.IsUnknown() {
				data.TenantName = types.StringNull()
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

		// Map filter parameters from response if available
		if val, ok := sourceMap["admin_state_up"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.AdminStateUp = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["backend_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BackendId = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["device_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.DeviceId = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["device_owner"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.DeviceOwner = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["exclude_subnet_uuids"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ExcludeSubnetUuids = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["fixed_ips"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.FixedIps = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["has_device_owner"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.HasDeviceOwner = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["mac_address"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MacAddress = types.StringValue(str)
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
		if val, ok := sourceMap["network_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NetworkName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["network_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NetworkUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["query"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Query = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["status"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Status = types.StringValue(str)
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
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
