package resources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackPortResource{}
var _ resource.ResourceWithImportState = &OpenstackPortResource{}

func NewOpenstackPortResource() resource.Resource {
	return &OpenstackPortResource{}
}

// OpenstackPortResource defines the resource implementation.
type OpenstackPortResource struct {
	client *client.Client
}

// OpenstackPortResourceModel describes the resource data model.
type OpenstackPortResourceModel struct {
	UUID                types.String   `tfsdk:"id"`
	AccessUrl           types.String   `tfsdk:"access_url"`
	AdminStateUp        types.Bool     `tfsdk:"admin_state_up"`
	AllowedAddressPairs types.List     `tfsdk:"allowed_address_pairs"`
	BackendId           types.String   `tfsdk:"backend_id"`
	Created             types.String   `tfsdk:"created"`
	Description         types.String   `tfsdk:"description"`
	DeviceId            types.String   `tfsdk:"device_id"`
	DeviceOwner         types.String   `tfsdk:"device_owner"`
	ErrorMessage        types.String   `tfsdk:"error_message"`
	ErrorTraceback      types.String   `tfsdk:"error_traceback"`
	FixedIps            types.List     `tfsdk:"fixed_ips"`
	FloatingIps         types.List     `tfsdk:"floating_ips"`
	MacAddress          types.String   `tfsdk:"mac_address"`
	Modified            types.String   `tfsdk:"modified"`
	Name                types.String   `tfsdk:"name"`
	Network             types.String   `tfsdk:"network"`
	NetworkName         types.String   `tfsdk:"network_name"`
	NetworkUuid         types.String   `tfsdk:"network_uuid"`
	PortSecurityEnabled types.Bool     `tfsdk:"port_security_enabled"`
	ResourceType        types.String   `tfsdk:"resource_type"`
	SecurityGroups      types.List     `tfsdk:"security_groups"`
	State               types.String   `tfsdk:"state"`
	Status              types.String   `tfsdk:"status"`
	TargetTenant        types.String   `tfsdk:"target_tenant"`
	Tenant              types.String   `tfsdk:"tenant"`
	TenantName          types.String   `tfsdk:"tenant_name"`
	TenantUuid          types.String   `tfsdk:"tenant_uuid"`
	Url                 types.String   `tfsdk:"url"`
	Timeouts            timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackPortResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_port"
}

func (r *OpenstackPortResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Port resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"access_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"admin_state_up": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Administrative state of the port. If down, port does not forward packets",
			},
			"allowed_address_pairs": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"mac_address": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Port ID in OpenStack",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"device_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of device (instance, router etc) to which this port is connected",
			},
			"device_owner": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Entity that uses this port (e.g. network:router_interface)",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"fixed_ips": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "IP address to assign to the port",
						},
						"subnet_id": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "ID of the subnet in which to assign the IP address",
						},
					},
				},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"floating_ips": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"mac_address": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "MAC address of the port",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: " ",
			},
			"network": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Network to which this port belongs",
			},
			"network_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"network_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"port_security_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "If True, security groups and rules will be applied to this port",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"security_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"status": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Port status in OpenStack (e.g. ACTIVE, DOWN)",
			},
			"target_tenant": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Target tenant for shared network port creation. If not specified, defaults to network's tenant.",
			},
			"tenant": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "OpenStack tenant this port belongs to",
			},
			"tenant_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
		},

		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx, timeouts.Opts{
				Create: true,
				Update: true,
				Delete: true,
			}),
		},
	}
}

func (r *OpenstackPortResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	r.client = client
}

func (r *OpenstackPortResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackPortResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Prepare request body
	requestBody := map[string]interface{}{}
	if !data.AllowedAddressPairs.IsNull() && !data.AllowedAddressPairs.IsUnknown() {
		if v := ConvertTFValue(data.AllowedAddressPairs); v != nil {
			requestBody["allowed_address_pairs"] = v
		}
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		if v := data.Description.ValueString(); v != "" {
			requestBody["description"] = v
		}
	}
	if !data.FixedIps.IsNull() && !data.FixedIps.IsUnknown() {
		if v := ConvertTFValue(data.FixedIps); v != nil {
			requestBody["fixed_ips"] = v
		}
	}
	if !data.MacAddress.IsNull() && !data.MacAddress.IsUnknown() {
		if v := data.MacAddress.ValueString(); v != "" {
			requestBody["mac_address"] = v
		}
	}
	requestBody["name"] = data.Name.ValueString()
	if !data.Network.IsNull() && !data.Network.IsUnknown() {
		if v := data.Network.ValueString(); v != "" {
			requestBody["network"] = v
		}
	}
	if !data.PortSecurityEnabled.IsNull() && !data.PortSecurityEnabled.IsUnknown() {
		requestBody["port_security_enabled"] = data.PortSecurityEnabled.ValueBool()
	}
	if !data.SecurityGroups.IsNull() && !data.SecurityGroups.IsUnknown() {
		if v := ConvertTFValue(data.SecurityGroups); v != nil {
			requestBody["security_groups"] = v
		}
	}
	if !data.TargetTenant.IsNull() && !data.TargetTenant.IsUnknown() {
		if v := data.TargetTenant.ValueString(); v != "" {
			requestBody["target_tenant"] = v
		}
	}

	// Call Waldur API to create resource
	var result map[string]interface{}
	err := r.client.Create(ctx, "/api/openstack-ports/", requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Openstack Port",
			"An error occurred while creating the Openstack Port: "+err.Error(),
		)
		return
	}
	// Extract UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	r.updateFromValue(ctx, &data, result)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackPortResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackPortResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource
	var result map[string]interface{}

	retrievePath := strings.Replace("/api/openstack-ports/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Openstack Port",
			"An error occurred while reading the Openstack Port: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	r.updateFromValue(ctx, &data, result)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackPortResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackPortResourceModel
	var state OpenstackPortResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.UUID = state.UUID

	// Prepare request body
	requestBody := map[string]interface{}{}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		if v := data.Description.ValueString(); v != "" {
			requestBody["description"] = v
		}
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		if v := data.Name.ValueString(); v != "" {
			requestBody["name"] = v
		}
	}
	if !data.SecurityGroups.IsNull() && !data.SecurityGroups.IsUnknown() {
		if v := ConvertTFValue(data.SecurityGroups); v != nil {
			requestBody["security_groups"] = v
		}
	}
	if !data.TargetTenant.IsNull() && !data.TargetTenant.IsUnknown() {
		if v := data.TargetTenant.ValueString(); v != "" {
			requestBody["target_tenant"] = v
		}
	}

	// Call Waldur API to update resource
	var result map[string]interface{}

	err := r.client.Update(ctx, "/api/openstack-ports/{uuid}/", data.UUID.ValueString(), requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Openstack Port",
			"An error occurred while updating the Openstack Port: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	r.updateFromValue(ctx, &data, result)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackPortResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackPortResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to delete resource
	err := r.client.DeleteByUUID(ctx, "/api/openstack-ports/{uuid}/", data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Openstack Port",
			"An error occurred while deleting the Openstack Port: "+err.Error(),
		)
		return
	}
}

func (r *OpenstackPortResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *OpenstackPortResource) updateFromValue(ctx context.Context, data *OpenstackPortResourceModel, sourceMap map[string]interface{}) {
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
}
