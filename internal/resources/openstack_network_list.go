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
		// Filter parameters can be added here if needed
		Attributes: map[string]schema.Attribute{},
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
	// Add filter fields here if added to schema
}

func (l *OpenstackNetworkList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config OpenstackNetworkListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Call API
	var listResult []map[string]interface{}
	err := l.client.List(ctx, "/api/openstack-networks/", &listResult)
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
