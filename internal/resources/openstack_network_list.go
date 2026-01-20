package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-provider-waldur/internal/client"
)

var _ list.ListResource = &OpenstackNetworkList{}

type OpenstackNetworkList struct {
	client *client.Client
}

func NewOpenstackNetworkList() list.ListResource {
	return &OpenstackNetworkList{}
}

func (l *OpenstackNetworkList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_network"
}

func (l *OpenstackNetworkList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
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
			"external_ip": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"is_external": schema.BoolAttribute{
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
			"type": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
		},
	}
}

func (l *OpenstackNetworkList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type OpenstackNetworkListModel struct {
	BackendId            types.String `tfsdk:"backend_id"`
	CanManage            types.Bool   `tfsdk:"can_manage"`
	Customer             types.String `tfsdk:"customer"`
	CustomerAbbreviation types.String `tfsdk:"customer_abbreviation"`
	CustomerName         types.String `tfsdk:"customer_name"`
	CustomerNativeName   types.String `tfsdk:"customer_native_name"`
	CustomerUuid         types.String `tfsdk:"customer_uuid"`
	Description          types.String `tfsdk:"description"`
	DirectOnly           types.Bool   `tfsdk:"direct_only"`
	ExternalIp           types.String `tfsdk:"external_ip"`
	IsExternal           types.Bool   `tfsdk:"is_external"`
	Name                 types.String `tfsdk:"name"`
	NameExact            types.String `tfsdk:"name_exact"`
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
	Type                 types.String `tfsdk:"type"`
	Uuid                 types.String `tfsdk:"uuid"`
}

func (l *OpenstackNetworkList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config OpenstackNetworkListModel

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
	if !config.ExternalIp.IsNull() && !config.ExternalIp.IsUnknown() {
		filters["external_ip"] = config.ExternalIp.ValueString()
	}
	if !config.IsExternal.IsNull() && !config.IsExternal.IsUnknown() {
		filters["is_external"] = fmt.Sprintf("%t", config.IsExternal.ValueBool())
	}
	if !config.Name.IsNull() && !config.Name.IsUnknown() {
		filters["name"] = config.Name.ValueString()
	}
	if !config.NameExact.IsNull() && !config.NameExact.IsUnknown() {
		filters["name_exact"] = config.NameExact.ValueString()
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
	if !config.Type.IsNull() && !config.Type.IsUnknown() {
		filters["type"] = config.Type.ValueString()
	}
	if !config.Uuid.IsNull() && !config.Uuid.IsUnknown() {
		filters["uuid"] = config.Uuid.ValueString()
	}

	// Call API
	var listResult []map[string]interface{}
	err := l.client.ListWithFilter(ctx, "/api/openstack-networks/", filters, &listResult)
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
			var data OpenstackNetworkResourceModel

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
			if val, ok := sourceMap["is_external"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.IsExternal = types.BoolValue(b)
				}
			} else {
				if data.IsExternal.IsUnknown() {
					data.IsExternal = types.BoolNull()
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
			if val, ok := sourceMap["mtu"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.Mtu = types.Int64Value(int64(num))
				}
			} else {
				if data.Mtu.IsUnknown() {
					data.Mtu = types.Int64Null()
				}
			}
			if val, ok := sourceMap["rbac_policies"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"backend_id":         types.StringType,
								"created":            types.StringType,
								"network":            types.StringType,
								"network_name":       types.StringType,
								"policy_type":        types.StringType,
								"target_tenant":      types.StringType,
								"target_tenant_name": types.StringType,
								"url":                types.StringType,
							}
							attrValues := map[string]attr.Value{
								"backend_id": func() attr.Value {
									if v, ok := objMap["backend_id"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"created": func() attr.Value {
									if v, ok := objMap["created"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"network": func() attr.Value {
									if v, ok := objMap["network"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"network_name": func() attr.Value {
									if v, ok := objMap["network_name"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"policy_type": func() attr.Value {
									if v, ok := objMap["policy_type"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"target_tenant": func() attr.Value {
									if v, ok := objMap["target_tenant"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"target_tenant_name": func() attr.Value {
									if v, ok := objMap["target_tenant_name"].(string); ok {
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
						"backend_id":         types.StringType,
						"created":            types.StringType,
						"network":            types.StringType,
						"network_name":       types.StringType,
						"policy_type":        types.StringType,
						"target_tenant":      types.StringType,
						"target_tenant_name": types.StringType,
						"url":                types.StringType,
					}}, items)
					data.RbacPolicies = listVal
				}
			} else {
				if data.RbacPolicies.IsUnknown() {
					data.RbacPolicies = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"backend_id":         types.StringType,
						"created":            types.StringType,
						"network":            types.StringType,
						"network_name":       types.StringType,
						"policy_type":        types.StringType,
						"target_tenant":      types.StringType,
						"target_tenant_name": types.StringType,
						"url":                types.StringType,
					}})
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
			if val, ok := sourceMap["segmentation_id"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.SegmentationId = types.Int64Value(int64(num))
				}
			} else {
				if data.SegmentationId.IsUnknown() {
					data.SegmentationId = types.Int64Null()
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
			if val, ok := sourceMap["subnets"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"end":   types.StringType,
									"start": types.StringType,
								}}},
								"cidr":        types.StringType,
								"description": types.StringType,
								"enable_dhcp": types.BoolType,
								"gateway_ip":  types.StringType,
								"ip_version":  types.Int64Type,
								"name":        types.StringType,
							}
							attrValues := map[string]attr.Value{
								"allocation_pools": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"end":   types.StringType,
									"start": types.StringType,
								}}}.ElemType),
								"cidr": func() attr.Value {
									if v, ok := objMap["cidr"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"description": func() attr.Value {
									if v, ok := objMap["description"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"enable_dhcp": func() attr.Value {
									if v, ok := objMap["enable_dhcp"].(bool); ok {
										return types.BoolValue(v)
									}
									return types.BoolNull()
								}(),
								"gateway_ip": func() attr.Value {
									if v, ok := objMap["gateway_ip"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"ip_version": func() attr.Value {
									if v, ok := objMap["ip_version"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
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
						"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"end":   types.StringType,
							"start": types.StringType,
						}}},
						"cidr":        types.StringType,
						"description": types.StringType,
						"enable_dhcp": types.BoolType,
						"gateway_ip":  types.StringType,
						"ip_version":  types.Int64Type,
						"name":        types.StringType,
					}}, items)
					data.Subnets = listVal
				}
			} else {
				if data.Subnets.IsUnknown() {
					data.Subnets = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"end":   types.StringType,
							"start": types.StringType,
						}}},
						"cidr":        types.StringType,
						"description": types.StringType,
						"enable_dhcp": types.BoolType,
						"gateway_ip":  types.StringType,
						"ip_version":  types.Int64Type,
						"name":        types.StringType,
					}})
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
			if val, ok := sourceMap["type"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Type = types.StringValue(str)
				}
			} else {
				if data.Type.IsUnknown() {
					data.Type = types.StringNull()
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
