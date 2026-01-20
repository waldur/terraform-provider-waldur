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
var _ resource.Resource = &OpenstackSubnetResource{}
var _ resource.ResourceWithImportState = &OpenstackSubnetResource{}

func NewOpenstackSubnetResource() resource.Resource {
	return &OpenstackSubnetResource{}
}

// OpenstackSubnetResource defines the resource implementation.
type OpenstackSubnetResource struct {
	client *client.Client
}

// OpenstackSubnetResourceModel describes the resource data model.
type OpenstackSubnetResourceModel struct {
	UUID            types.String   `tfsdk:"id"`
	AccessUrl       types.String   `tfsdk:"access_url"`
	AllocationPools types.List     `tfsdk:"allocation_pools"`
	BackendId       types.String   `tfsdk:"backend_id"`
	Cidr            types.String   `tfsdk:"cidr"`
	Created         types.String   `tfsdk:"created"`
	Description     types.String   `tfsdk:"description"`
	DisableGateway  types.Bool     `tfsdk:"disable_gateway"`
	DnsNameservers  types.List     `tfsdk:"dns_nameservers"`
	EnableDhcp      types.Bool     `tfsdk:"enable_dhcp"`
	ErrorMessage    types.String   `tfsdk:"error_message"`
	ErrorTraceback  types.String   `tfsdk:"error_traceback"`
	GatewayIp       types.String   `tfsdk:"gateway_ip"`
	HostRoutes      types.List     `tfsdk:"host_routes"`
	IpVersion       types.Int64    `tfsdk:"ip_version"`
	IsConnected     types.Bool     `tfsdk:"is_connected"`
	Modified        types.String   `tfsdk:"modified"`
	Name            types.String   `tfsdk:"name"`
	Network         types.String   `tfsdk:"network"`
	NetworkName     types.String   `tfsdk:"network_name"`
	ResourceType    types.String   `tfsdk:"resource_type"`
	State           types.String   `tfsdk:"state"`
	Tenant          types.String   `tfsdk:"tenant"`
	TenantName      types.String   `tfsdk:"tenant_name"`
	Url             types.String   `tfsdk:"url"`
	Timeouts        timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackSubnetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_subnet"
}

func (r *OpenstackSubnetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Subnet resource",

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
			"allocation_pools": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"end": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "An IPv4 or IPv6 address.",
						},
						"start": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "An IPv4 or IPv6 address.",
						},
					},
				},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"cidr": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
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
			"disable_gateway": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "If True, no gateway IP address will be allocated",
			},
			"dns_nameservers": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"enable_dhcp": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "If True, DHCP service will be enabled on this subnet",
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
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "IP address of the gateway for this subnet",
			},
			"host_routes": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"destination": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: " ",
						},
						"nexthop": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "An IPv4 or IPv6 address.",
						},
					},
				},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"ip_version": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "IP protocol version (4 or 6)",
			},
			"is_connected": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Is subnet connected to the default tenant router.",
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
				Required:            true,
				MarkdownDescription: "Required path parameter for resource creation",
			},
			"network_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"tenant": schema.StringAttribute{
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

		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx, timeouts.Opts{
				Create: true,
				Update: true,
				Delete: true,
			}),
		},
	}
}

func (r *OpenstackSubnetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OpenstackSubnetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackSubnetResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Prepare request body
	requestBody := map[string]interface{}{}
	if !data.AllocationPools.IsNull() && !data.AllocationPools.IsUnknown() {
		if v := ConvertTFValue(data.AllocationPools); v != nil {
			requestBody["allocation_pools"] = v
		}
	}
	if !data.Cidr.IsNull() && !data.Cidr.IsUnknown() {
		if v := data.Cidr.ValueString(); v != "" {
			requestBody["cidr"] = v
		}
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		if v := data.Description.ValueString(); v != "" {
			requestBody["description"] = v
		}
	}
	if !data.DisableGateway.IsNull() && !data.DisableGateway.IsUnknown() {
		requestBody["disable_gateway"] = data.DisableGateway.ValueBool()
	}
	if !data.DnsNameservers.IsNull() && !data.DnsNameservers.IsUnknown() {
		if v := ConvertTFValue(data.DnsNameservers); v != nil {
			requestBody["dns_nameservers"] = v
		}
	}
	if !data.GatewayIp.IsNull() && !data.GatewayIp.IsUnknown() {
		if v := data.GatewayIp.ValueString(); v != "" {
			requestBody["gateway_ip"] = v
		}
	}
	if !data.HostRoutes.IsNull() && !data.HostRoutes.IsUnknown() {
		if v := ConvertTFValue(data.HostRoutes); v != nil {
			requestBody["host_routes"] = v
		}
	}
	requestBody["name"] = data.Name.ValueString()

	// Call Waldur API to create resource
	var result map[string]interface{}
	// Custom create operation via parent resource
	createPath := "/api/openstack-networks/{uuid}/create_subnet/"
	createPath = strings.Replace(createPath, "{uuid}", data.Network.ValueString(), 1)
	err := r.client.Post(ctx, createPath, requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Openstack Subnet",
			"An error occurred while creating the Openstack Subnet: "+err.Error(),
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

func (r *OpenstackSubnetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackSubnetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource
	var result map[string]interface{}

	retrievePath := strings.Replace("/api/openstack-subnets/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Openstack Subnet",
			"An error occurred while reading the Openstack Subnet: "+err.Error(),
		)
		return
	}

	r.updateFromValue(ctx, &data, result)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackSubnetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackSubnetResourceModel
	var state OpenstackSubnetResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.UUID = state.UUID

	// Prepare request body
	requestBody := map[string]interface{}{}
	if !data.AllocationPools.IsNull() && !data.AllocationPools.IsUnknown() {
		if v := ConvertTFValue(data.AllocationPools); v != nil {
			requestBody["allocation_pools"] = v
		}
	}
	if !data.Cidr.IsNull() && !data.Cidr.IsUnknown() {
		if v := data.Cidr.ValueString(); v != "" {
			requestBody["cidr"] = v
		}
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		if v := data.Description.ValueString(); v != "" {
			requestBody["description"] = v
		}
	}
	if !data.DisableGateway.IsNull() && !data.DisableGateway.IsUnknown() {
		requestBody["disable_gateway"] = data.DisableGateway.ValueBool()
	}
	if !data.DnsNameservers.IsNull() && !data.DnsNameservers.IsUnknown() {
		if v := ConvertTFValue(data.DnsNameservers); v != nil {
			requestBody["dns_nameservers"] = v
		}
	}
	if !data.GatewayIp.IsNull() && !data.GatewayIp.IsUnknown() {
		if v := data.GatewayIp.ValueString(); v != "" {
			requestBody["gateway_ip"] = v
		}
	}
	if !data.HostRoutes.IsNull() && !data.HostRoutes.IsUnknown() {
		if v := ConvertTFValue(data.HostRoutes); v != nil {
			requestBody["host_routes"] = v
		}
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		if v := data.Name.ValueString(); v != "" {
			requestBody["name"] = v
		}
	}

	// Call Waldur API to update resource
	var result map[string]interface{}

	err := r.client.Update(ctx, "/api/openstack-subnets/{uuid}/", data.UUID.ValueString(), requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Openstack Subnet",
			"An error occurred while updating the Openstack Subnet: "+err.Error(),
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

func (r *OpenstackSubnetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackSubnetResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to delete resource
	err := r.client.DeleteByUUID(ctx, "/api/openstack-subnets/{uuid}/", data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Openstack Subnet",
			"An error occurred while deleting the Openstack Subnet: "+err.Error(),
		)
		return
	}
}

func (r *OpenstackSubnetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *OpenstackSubnetResource) updateFromValue(ctx context.Context, data *OpenstackSubnetResourceModel, sourceMap map[string]interface{}) {
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
}
