package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-provider-waldur/internal/client"
)

var _ list.ListResource = &OpenstackPortList{}

type OpenstackPortList struct {
	client *client.Client
}

func NewOpenstackPortList() list.ListResource {
	return &OpenstackPortList{}
}

func (l *OpenstackPortList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_port"
}

func (l *OpenstackPortList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"admin_state_up": schema.BoolAttribute{
				Description: "",
				Optional:    true,
			},
			"backend_id": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"device_id": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"device_owner": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"exclude_subnet_uuids": schema.StringAttribute{
				Description: "Exclude Subnet UUIDs (comma-separated)",
				Optional:    true,
			},
			"fixed_ips": schema.StringAttribute{
				Description: "Search by fixed IP",
				Optional:    true,
			},
			"has_device_owner": schema.BoolAttribute{
				Description: "Has device owner",
				Optional:    true,
			},
			"mac_address": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"name_exact": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"network_name": schema.StringAttribute{
				Description: "Search by network name",
				Optional:    true,
			},
			"network_uuid": schema.StringAttribute{
				Description: "Search by network UUID",
				Optional:    true,
			},
			"page": schema.Int64Attribute{
				Description: "A page number within the paginated result set.",
				Optional:    true,
			},
			"page_size": schema.Int64Attribute{
				Description: "Number of results to return per page.",
				Optional:    true,
			},
			"query": schema.StringAttribute{
				Description: "Search by name, MAC address or backend ID",
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"tenant": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"tenant_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
		},
	}
}

func (l *OpenstackPortList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected List Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T", req.ProviderData),
		)
		return
	}

	l.client = client
}

type OpenstackPortListModel struct {
	AdminStateUp       types.Bool   `tfsdk:"admin_state_up"`
	BackendId          types.String `tfsdk:"backend_id"`
	DeviceId           types.String `tfsdk:"device_id"`
	DeviceOwner        types.String `tfsdk:"device_owner"`
	ExcludeSubnetUuids types.String `tfsdk:"exclude_subnet_uuids"`
	FixedIps           types.String `tfsdk:"fixed_ips"`
	HasDeviceOwner     types.Bool   `tfsdk:"has_device_owner"`
	MacAddress         types.String `tfsdk:"mac_address"`
	Name               types.String `tfsdk:"name"`
	NameExact          types.String `tfsdk:"name_exact"`
	NetworkName        types.String `tfsdk:"network_name"`
	NetworkUuid        types.String `tfsdk:"network_uuid"`
	Page               types.Int64  `tfsdk:"page"`
	PageSize           types.Int64  `tfsdk:"page_size"`
	Query              types.String `tfsdk:"query"`
	Status             types.String `tfsdk:"status"`
	Tenant             types.String `tfsdk:"tenant"`
	TenantUuid         types.String `tfsdk:"tenant_uuid"`
}

func (l *OpenstackPortList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config OpenstackPortListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := make(map[string]string)
	if !config.AdminStateUp.IsNull() && !config.AdminStateUp.IsUnknown() {
		filters["admin_state_up"] = fmt.Sprintf("%t", config.AdminStateUp.ValueBool())
	}
	if !config.BackendId.IsNull() && !config.BackendId.IsUnknown() {
		filters["backend_id"] = config.BackendId.ValueString()
	}
	if !config.DeviceId.IsNull() && !config.DeviceId.IsUnknown() {
		filters["device_id"] = config.DeviceId.ValueString()
	}
	if !config.DeviceOwner.IsNull() && !config.DeviceOwner.IsUnknown() {
		filters["device_owner"] = config.DeviceOwner.ValueString()
	}
	if !config.ExcludeSubnetUuids.IsNull() && !config.ExcludeSubnetUuids.IsUnknown() {
		filters["exclude_subnet_uuids"] = config.ExcludeSubnetUuids.ValueString()
	}
	if !config.FixedIps.IsNull() && !config.FixedIps.IsUnknown() {
		filters["fixed_ips"] = config.FixedIps.ValueString()
	}
	if !config.HasDeviceOwner.IsNull() && !config.HasDeviceOwner.IsUnknown() {
		filters["has_device_owner"] = fmt.Sprintf("%t", config.HasDeviceOwner.ValueBool())
	}
	if !config.MacAddress.IsNull() && !config.MacAddress.IsUnknown() {
		filters["mac_address"] = config.MacAddress.ValueString()
	}
	if !config.Name.IsNull() && !config.Name.IsUnknown() {
		filters["name"] = config.Name.ValueString()
	}
	if !config.NameExact.IsNull() && !config.NameExact.IsUnknown() {
		filters["name_exact"] = config.NameExact.ValueString()
	}
	if !config.NetworkName.IsNull() && !config.NetworkName.IsUnknown() {
		filters["network_name"] = config.NetworkName.ValueString()
	}
	if !config.NetworkUuid.IsNull() && !config.NetworkUuid.IsUnknown() {
		filters["network_uuid"] = config.NetworkUuid.ValueString()
	}
	if !config.Page.IsNull() && !config.Page.IsUnknown() {
		filters["page"] = fmt.Sprintf("%d", config.Page.ValueInt64())
	}
	if !config.PageSize.IsNull() && !config.PageSize.IsUnknown() {
		filters["page_size"] = fmt.Sprintf("%d", config.PageSize.ValueInt64())
	}
	if !config.Query.IsNull() && !config.Query.IsUnknown() {
		filters["query"] = config.Query.ValueString()
	}
	if !config.Status.IsNull() && !config.Status.IsUnknown() {
		filters["status"] = config.Status.ValueString()
	}
	if !config.Tenant.IsNull() && !config.Tenant.IsUnknown() {
		filters["tenant"] = config.Tenant.ValueString()
	}
	if !config.TenantUuid.IsNull() && !config.TenantUuid.IsUnknown() {
		filters["tenant_uuid"] = config.TenantUuid.ValueString()
	}

	// Call API
	var listResult []map[string]interface{}
	err := l.client.ListWithFilter(ctx, "/api/openstack-ports/", filters, &listResult)
	if err != nil {
		// Return error diagnostics
		resp.AddError("Failed to list resources", err.Error())
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Stream results
	stream.Results = func(push func(list.ListResult) bool) {
		for _, item := range listResult {
			result := req.NewListResult(ctx)

			// Map item to model
			var data OpenstackPortResourceModel

			// Initialize lists and objects to empty/null if needed (or rely on map logic)

			sourceMap := item

			// Reuse the mapResponseFields logic (embedded here or via shared template)
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
			if val, ok := sourceMap["admin_state_up"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.AdminStateUp = types.BoolValue(b)
				}
			} else {
				if data.AdminStateUp.IsUnknown() {
					data.AdminStateUp = types.BoolNull()
				}
			}
			if val, ok := sourceMap["allowed_address_pairs"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"ip_address":  types.StringType,
								"mac_address": types.StringType,
							}
							attrValues := map[string]attr.Value{
								"ip_address": func() attr.Value {
									if v, ok := objMap["ip_address"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
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
						"ip_address":  types.StringType,
						"mac_address": types.StringType,
					}}, items)
					data.AllowedAddressPairs = listVal
				}
			} else {
				if data.AllowedAddressPairs.IsUnknown() {
					data.AllowedAddressPairs = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"ip_address":  types.StringType,
						"mac_address": types.StringType,
					}})
				}
			}
			if val, ok := sourceMap["backend_id"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.BackendId = types.StringValue(str)
				}
			} else {
				if data.BackendId.IsUnknown() {
					data.BackendId = types.StringNull()
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
			if val, ok := sourceMap["device_id"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.DeviceId = types.StringValue(str)
				}
			} else {
				if data.DeviceId.IsUnknown() {
					data.DeviceId = types.StringNull()
				}
			}
			if val, ok := sourceMap["device_owner"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.DeviceOwner = types.StringValue(str)
				}
			} else {
				if data.DeviceOwner.IsUnknown() {
					data.DeviceOwner = types.StringNull()
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
			if val, ok := sourceMap["fixed_ips"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"ip_address": types.StringType,
								"subnet_id":  types.StringType,
							}
							attrValues := map[string]attr.Value{
								"ip_address": func() attr.Value {
									if v, ok := objMap["ip_address"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"subnet_id": func() attr.Value {
									if v, ok := objMap["subnet_id"].(string); ok {
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
						"ip_address": types.StringType,
						"subnet_id":  types.StringType,
					}}, items)
					data.FixedIps = listVal
				}
			} else {
				if data.FixedIps.IsUnknown() {
					data.FixedIps = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"ip_address": types.StringType,
						"subnet_id":  types.StringType,
					}})
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
			if val, ok := sourceMap["mac_address"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.MacAddress = types.StringValue(str)
				}
			} else {
				if data.MacAddress.IsUnknown() {
					data.MacAddress = types.StringNull()
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
			if val, ok := sourceMap["network_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.NetworkName = types.StringValue(str)
				}
			} else {
				if data.NetworkName.IsUnknown() {
					data.NetworkName = types.StringNull()
				}
			}
			if val, ok := sourceMap["network_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.NetworkUuid = types.StringValue(str)
				}
			} else {
				if data.NetworkUuid.IsUnknown() {
					data.NetworkUuid = types.StringNull()
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
							}
							attrValues := map[string]attr.Value{
								"name": func() attr.Value {
									if v, ok := objMap["name"].(string); ok {
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
					}}, items)
					data.SecurityGroups = listVal
				}
			} else {
				if data.SecurityGroups.IsUnknown() {
					data.SecurityGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"name": types.StringType,
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
			if val, ok := sourceMap["status"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Status = types.StringValue(str)
				}
			} else {
				if data.Status.IsUnknown() {
					data.Status = types.StringNull()
				}
			}
			if val, ok := sourceMap["target_tenant"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.TargetTenant = types.StringValue(str)
				}
			} else {
				if data.TargetTenant.IsUnknown() {
					data.TargetTenant = types.StringNull()
				}
			}
			if val, ok := sourceMap["tenant"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Tenant = types.StringValue(str)
				}
			} else {
				if data.Tenant.IsUnknown() {
					data.Tenant = types.StringNull()
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
			if val, ok := sourceMap["tenant_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.TenantUuid = types.StringValue(str)
				}
			} else {
				if data.TenantUuid.IsUnknown() {
					data.TenantUuid = types.StringNull()
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

			// Set the resource state
			// For ListResource, we generally return the "Resource" state matching the main resource schema.
			// This allows `terraform plan` to correlate if using `terraform query`.

			diags := result.Resource.Set(ctx, &data)
			result.Diagnostics.Append(diags...)

			// Set Identity if possible (usually UUID)
			if !data.UUID.IsNull() && !data.UUID.IsUnknown() {
				// Identity value must match what the resource uses for Import?
				// Typically implicit. For now just setting Resource is key.
				// result.Identity.Set(ctx, data.UUID.ValueString())
				// The doc says: "An error is returned if a list result in the stream contains a null identity"
				result.Diagnostics.Append(result.Identity.Set(ctx, data.UUID.ValueString())...)
			} else {
				// Try to fallback to "uuid" from map if model failed
				if uuid, ok := item["uuid"].(string); ok {
					result.Diagnostics.Append(result.Identity.Set(ctx, uuid)...)
				}
			}

			if !push(result) {
				return
			}
		}
	}
}
