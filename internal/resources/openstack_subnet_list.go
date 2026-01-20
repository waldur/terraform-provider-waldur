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

var _ list.ListResource = &OpenstackSubnetList{}

type OpenstackSubnetList struct {
	client *client.Client
}

func NewOpenstackSubnetList() list.ListResource {
	return &OpenstackSubnetList{}
}

func (l *OpenstackSubnetList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_subnet"
}

func (l *OpenstackSubnetList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"backend_id": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"can_manage": schema.BoolAttribute{
				Description: "Can manage",
				Optional:    true,
			},
			"customer": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"customer_abbreviation": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"customer_name": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"customer_native_name": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"customer_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"direct_only": schema.BoolAttribute{
				Description: "Direct only",
				Optional:    true,
			},
			"enable_dhcp": schema.BoolAttribute{
				Description: "",
				Optional:    true,
			},
			"external_ip": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"ip_version": schema.Int64Attribute{
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
			"network": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"network_uuid": schema.StringAttribute{
				Description: "",
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
			"project": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"project_name": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"project_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"rbac_only": schema.BoolAttribute{
				Description: "RBAC only",
				Optional:    true,
			},
			"service_settings_name": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"service_settings_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"tenant": schema.StringAttribute{
				Description: "Tenant URL",
				Optional:    true,
			},
			"tenant_uuid": schema.StringAttribute{
				Description: "Tenant UUID",
				Optional:    true,
			},
			"uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
		},
	}
}

func (l *OpenstackSubnetList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type OpenstackSubnetListModel struct {
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
	Page                 types.Int64  `tfsdk:"page"`
	PageSize             types.Int64  `tfsdk:"page_size"`
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

func (l *OpenstackSubnetList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config OpenstackSubnetListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := make(map[string]string)
	if !config.BackendId.IsNull() && !config.BackendId.IsUnknown() {
		filters["backend_id"] = config.BackendId.ValueString()
	}
	if !config.CanManage.IsNull() && !config.CanManage.IsUnknown() {
		filters["can_manage"] = fmt.Sprintf("%t", config.CanManage.ValueBool())
	}
	if !config.Customer.IsNull() && !config.Customer.IsUnknown() {
		filters["customer"] = config.Customer.ValueString()
	}
	if !config.CustomerAbbreviation.IsNull() && !config.CustomerAbbreviation.IsUnknown() {
		filters["customer_abbreviation"] = config.CustomerAbbreviation.ValueString()
	}
	if !config.CustomerName.IsNull() && !config.CustomerName.IsUnknown() {
		filters["customer_name"] = config.CustomerName.ValueString()
	}
	if !config.CustomerNativeName.IsNull() && !config.CustomerNativeName.IsUnknown() {
		filters["customer_native_name"] = config.CustomerNativeName.ValueString()
	}
	if !config.CustomerUuid.IsNull() && !config.CustomerUuid.IsUnknown() {
		filters["customer_uuid"] = config.CustomerUuid.ValueString()
	}
	if !config.Description.IsNull() && !config.Description.IsUnknown() {
		filters["description"] = config.Description.ValueString()
	}
	if !config.DirectOnly.IsNull() && !config.DirectOnly.IsUnknown() {
		filters["direct_only"] = fmt.Sprintf("%t", config.DirectOnly.ValueBool())
	}
	if !config.EnableDhcp.IsNull() && !config.EnableDhcp.IsUnknown() {
		filters["enable_dhcp"] = fmt.Sprintf("%t", config.EnableDhcp.ValueBool())
	}
	if !config.ExternalIp.IsNull() && !config.ExternalIp.IsUnknown() {
		filters["external_ip"] = config.ExternalIp.ValueString()
	}
	if !config.IpVersion.IsNull() && !config.IpVersion.IsUnknown() {
		filters["ip_version"] = fmt.Sprintf("%d", config.IpVersion.ValueInt64())
	}
	if !config.Name.IsNull() && !config.Name.IsUnknown() {
		filters["name"] = config.Name.ValueString()
	}
	if !config.NameExact.IsNull() && !config.NameExact.IsUnknown() {
		filters["name_exact"] = config.NameExact.ValueString()
	}
	if !config.Network.IsNull() && !config.Network.IsUnknown() {
		filters["network"] = config.Network.ValueString()
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
	if !config.Project.IsNull() && !config.Project.IsUnknown() {
		filters["project"] = config.Project.ValueString()
	}
	if !config.ProjectName.IsNull() && !config.ProjectName.IsUnknown() {
		filters["project_name"] = config.ProjectName.ValueString()
	}
	if !config.ProjectUuid.IsNull() && !config.ProjectUuid.IsUnknown() {
		filters["project_uuid"] = config.ProjectUuid.ValueString()
	}
	if !config.RbacOnly.IsNull() && !config.RbacOnly.IsUnknown() {
		filters["rbac_only"] = fmt.Sprintf("%t", config.RbacOnly.ValueBool())
	}
	if !config.ServiceSettingsName.IsNull() && !config.ServiceSettingsName.IsUnknown() {
		filters["service_settings_name"] = config.ServiceSettingsName.ValueString()
	}
	if !config.ServiceSettingsUuid.IsNull() && !config.ServiceSettingsUuid.IsUnknown() {
		filters["service_settings_uuid"] = config.ServiceSettingsUuid.ValueString()
	}
	if !config.Tenant.IsNull() && !config.Tenant.IsUnknown() {
		filters["tenant"] = config.Tenant.ValueString()
	}
	if !config.TenantUuid.IsNull() && !config.TenantUuid.IsUnknown() {
		filters["tenant_uuid"] = config.TenantUuid.ValueString()
	}
	if !config.Uuid.IsNull() && !config.Uuid.IsUnknown() {
		filters["uuid"] = config.Uuid.ValueString()
	}

	// Call API
	var listResult []map[string]interface{}
	err := l.client.ListWithFilter(ctx, "/api/openstack-subnets/", filters, &listResult)
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
			var data OpenstackSubnetResourceModel

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
			if val, ok := sourceMap["backend_id"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.BackendId = types.StringValue(str)
				}
			} else {
				if data.BackendId.IsUnknown() {
					data.BackendId = types.StringNull()
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
			if val, ok := sourceMap["description"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Description = types.StringValue(str)
				}
			} else {
				if data.Description.IsUnknown() {
					data.Description = types.StringNull()
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
			if val, ok := sourceMap["enable_dhcp"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.EnableDhcp = types.BoolValue(b)
				}
			} else {
				if data.EnableDhcp.IsUnknown() {
					data.EnableDhcp = types.BoolNull()
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
			if val, ok := sourceMap["ip_version"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.IpVersion = types.Int64Value(int64(num))
				}
			} else {
				if data.IpVersion.IsUnknown() {
					data.IpVersion = types.Int64Null()
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
			if val, ok := sourceMap["resource_type"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ResourceType = types.StringValue(str)
				}
			} else {
				if data.ResourceType.IsUnknown() {
					data.ResourceType = types.StringNull()
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
