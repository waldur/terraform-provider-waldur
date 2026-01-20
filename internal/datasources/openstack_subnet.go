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
var _ datasource.DataSource = &OpenstackSubnetDataSource{}

func NewOpenstackSubnetDataSource() datasource.DataSource {
	return &OpenstackSubnetDataSource{}
}

// OpenstackSubnetDataSource defines the data source implementation.
type OpenstackSubnetDataSource struct {
	client *client.Client
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
		var item map[string]interface{}

		err := d.client.GetByUUID(ctx, "/api/openstack-subnets/{uuid}/", data.UUID.ValueString(), &item)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Subnet",
				"An error occurred while reading the Openstack Subnet by UUID: "+err.Error(),
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
		if val, ok := sourceMap["allocation_pools"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"end":   types.StringType,
							"start": types.StringType,
						}
						attrValues := map[string]attr.Value{
							"end": func() attr.Value {
								if v, ok := objMap["end"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"start": func() attr.Value {
								if v, ok := objMap["start"].(string); ok {
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
					"end":   types.StringType,
					"start": types.StringType,
				}}, items)
				data.AllocationPools = listVal
			}
		} else {
			if data.AllocationPools.IsUnknown() {
				data.AllocationPools = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"end":   types.StringType,
					"start": types.StringType,
				}})
			}
		}
		if val, ok := sourceMap["cidr"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Cidr = types.StringValue(str)
			}
		} else {
			if data.Cidr.IsUnknown() {
				data.Cidr = types.StringNull()
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
		if val, ok := sourceMap["disable_gateway"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.DisableGateway = types.BoolValue(b)
			}
		} else {
			if data.DisableGateway.IsUnknown() {
				data.DisableGateway = types.BoolNull()
			}
		}
		if val, ok := sourceMap["dns_nameservers"]; ok && val != nil {
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
				data.DnsNameservers = listVal
			}
		} else {
			if data.DnsNameservers.IsUnknown() {
				data.DnsNameservers = types.ListNull(types.StringType)
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
		if val, ok := sourceMap["gateway_ip"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.GatewayIp = types.StringValue(str)
			}
		} else {
			if data.GatewayIp.IsUnknown() {
				data.GatewayIp = types.StringNull()
			}
		}
		if val, ok := sourceMap["host_routes"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"destination": types.StringType,
							"nexthop":     types.StringType,
						}
						attrValues := map[string]attr.Value{
							"destination": func() attr.Value {
								if v, ok := objMap["destination"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"nexthop": func() attr.Value {
								if v, ok := objMap["nexthop"].(string); ok {
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
					"destination": types.StringType,
					"nexthop":     types.StringType,
				}}, items)
				data.HostRoutes = listVal
			}
		} else {
			if data.HostRoutes.IsUnknown() {
				data.HostRoutes = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"destination": types.StringType,
					"nexthop":     types.StringType,
				}})
			}
		}
		if val, ok := sourceMap["is_connected"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsConnected = types.BoolValue(b)
			}
		} else {
			if data.IsConnected.IsUnknown() {
				data.IsConnected = types.BoolNull()
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
		if val, ok := sourceMap["network_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NetworkName = types.StringValue(str)
			}
		} else {
			if data.NetworkName.IsUnknown() {
				data.NetworkName = types.StringNull()
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
		if val, ok := sourceMap["direct_only"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.DirectOnly = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["enable_dhcp"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.EnableDhcp = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["external_ip"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ExternalIp = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["ip_version"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.IpVersion = types.Int64Value(int64(num))
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
		if val, ok := sourceMap["network"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Network = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["network_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NetworkUuid = types.StringValue(str)
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
		if val, ok := sourceMap["rbac_only"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.RbacOnly = types.BoolValue(b)
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
		if val, ok := sourceMap["allocation_pools"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"end":   types.StringType,
							"start": types.StringType,
						}
						attrValues := map[string]attr.Value{
							"end": func() attr.Value {
								if v, ok := objMap["end"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"start": func() attr.Value {
								if v, ok := objMap["start"].(string); ok {
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
					"end":   types.StringType,
					"start": types.StringType,
				}}, items)
				data.AllocationPools = listVal
			}
		} else {
			if data.AllocationPools.IsUnknown() {
				data.AllocationPools = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"end":   types.StringType,
					"start": types.StringType,
				}})
			}
		}
		if val, ok := sourceMap["cidr"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Cidr = types.StringValue(str)
			}
		} else {
			if data.Cidr.IsUnknown() {
				data.Cidr = types.StringNull()
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
		if val, ok := sourceMap["disable_gateway"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.DisableGateway = types.BoolValue(b)
			}
		} else {
			if data.DisableGateway.IsUnknown() {
				data.DisableGateway = types.BoolNull()
			}
		}
		if val, ok := sourceMap["dns_nameservers"]; ok && val != nil {
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
				data.DnsNameservers = listVal
			}
		} else {
			if data.DnsNameservers.IsUnknown() {
				data.DnsNameservers = types.ListNull(types.StringType)
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
		if val, ok := sourceMap["gateway_ip"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.GatewayIp = types.StringValue(str)
			}
		} else {
			if data.GatewayIp.IsUnknown() {
				data.GatewayIp = types.StringNull()
			}
		}
		if val, ok := sourceMap["host_routes"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"destination": types.StringType,
							"nexthop":     types.StringType,
						}
						attrValues := map[string]attr.Value{
							"destination": func() attr.Value {
								if v, ok := objMap["destination"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"nexthop": func() attr.Value {
								if v, ok := objMap["nexthop"].(string); ok {
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
					"destination": types.StringType,
					"nexthop":     types.StringType,
				}}, items)
				data.HostRoutes = listVal
			}
		} else {
			if data.HostRoutes.IsUnknown() {
				data.HostRoutes = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"destination": types.StringType,
					"nexthop":     types.StringType,
				}})
			}
		}
		if val, ok := sourceMap["is_connected"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsConnected = types.BoolValue(b)
			}
		} else {
			if data.IsConnected.IsUnknown() {
				data.IsConnected = types.BoolNull()
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
		if val, ok := sourceMap["network_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NetworkName = types.StringValue(str)
			}
		} else {
			if data.NetworkName.IsUnknown() {
				data.NetworkName = types.StringNull()
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
		if val, ok := sourceMap["direct_only"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.DirectOnly = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["enable_dhcp"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.EnableDhcp = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["external_ip"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ExternalIp = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["ip_version"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.IpVersion = types.Int64Value(int64(num))
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
		if val, ok := sourceMap["network"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Network = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["network_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NetworkUuid = types.StringValue(str)
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
		if val, ok := sourceMap["rbac_only"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.RbacOnly = types.BoolValue(b)
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
