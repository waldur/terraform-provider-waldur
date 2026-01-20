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

var _ list.ListResource = &OpenstackFloatingIpList{}

type OpenstackFloatingIpList struct {
	client *client.Client
}

func NewOpenstackFloatingIpList() list.ListResource {
	return &OpenstackFloatingIpList{}
}

func (l *OpenstackFloatingIpList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_floating_ip"
}

func (l *OpenstackFloatingIpList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		// Filter parameters can be added here if needed
		Attributes: map[string]schema.Attribute{},
	}
}

func (l *OpenstackFloatingIpList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type OpenstackFloatingIpListModel struct {
	// Add filter fields here if added to schema
}

func (l *OpenstackFloatingIpList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config OpenstackFloatingIpListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Call API
	var listResult []map[string]interface{}
	err := l.client.List(ctx, "/api/openstack-floating-ips/", &listResult)
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
			var data OpenstackFloatingIpResourceModel

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
			if val, ok := sourceMap["address"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Address = types.StringValue(str)
				}
			} else {
				if data.Address.IsUnknown() {
					data.Address = types.StringNull()
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
			if val, ok := sourceMap["backend_network_id"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.BackendNetworkId = types.StringValue(str)
				}
			} else {
				if data.BackendNetworkId.IsUnknown() {
					data.BackendNetworkId = types.StringNull()
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
			if val, ok := sourceMap["external_address"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ExternalAddress = types.StringValue(str)
				}
			} else {
				if data.ExternalAddress.IsUnknown() {
					data.ExternalAddress = types.StringNull()
				}
			}
			if val, ok := sourceMap["instance_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.InstanceName = types.StringValue(str)
				}
			} else {
				if data.InstanceName.IsUnknown() {
					data.InstanceName = types.StringNull()
				}
			}
			if val, ok := sourceMap["instance_url"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.InstanceUrl = types.StringValue(str)
				}
			} else {
				if data.InstanceUrl.IsUnknown() {
					data.InstanceUrl = types.StringNull()
				}
			}
			if val, ok := sourceMap["instance_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.InstanceUuid = types.StringValue(str)
				}
			} else {
				if data.InstanceUuid.IsUnknown() {
					data.InstanceUuid = types.StringNull()
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
			if val, ok := sourceMap["port"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Port = types.StringValue(str)
				}
			} else {
				if data.Port.IsUnknown() {
					data.Port = types.StringNull()
				}
			}
			if val, ok := sourceMap["port_fixed_ips"]; ok && val != nil {
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
					data.PortFixedIps = listVal
				}
			} else {
				if data.PortFixedIps.IsUnknown() {
					data.PortFixedIps = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"ip_address": types.StringType,
						"subnet_id":  types.StringType,
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
			if val, ok := sourceMap["runtime_state"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.RuntimeState = types.StringValue(str)
				}
			} else {
				if data.RuntimeState.IsUnknown() {
					data.RuntimeState = types.StringNull()
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
