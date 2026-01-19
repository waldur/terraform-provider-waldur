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

	"github.com/waldur/terraform-waldur-provider/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackFloatingIpResource{}
var _ resource.ResourceWithImportState = &OpenstackFloatingIpResource{}

func NewOpenstackFloatingIpResource() resource.Resource {
	return &OpenstackFloatingIpResource{}
}

// OpenstackFloatingIpResource defines the resource implementation.
type OpenstackFloatingIpResource struct {
	client *client.Client
}

// OpenstackFloatingIpResourceModel describes the resource data model.
type OpenstackFloatingIpResourceModel struct {
	UUID             types.String   `tfsdk:"id"`
	AccessUrl        types.String   `tfsdk:"access_url"`
	Address          types.String   `tfsdk:"address"`
	BackendId        types.String   `tfsdk:"backend_id"`
	BackendNetworkId types.String   `tfsdk:"backend_network_id"`
	Created          types.String   `tfsdk:"created"`
	Description      types.String   `tfsdk:"description"`
	ErrorMessage     types.String   `tfsdk:"error_message"`
	ErrorTraceback   types.String   `tfsdk:"error_traceback"`
	ExternalAddress  types.String   `tfsdk:"external_address"`
	InstanceName     types.String   `tfsdk:"instance_name"`
	InstanceUrl      types.String   `tfsdk:"instance_url"`
	InstanceUuid     types.String   `tfsdk:"instance_uuid"`
	Modified         types.String   `tfsdk:"modified"`
	Name             types.String   `tfsdk:"name"`
	Port             types.String   `tfsdk:"port"`
	PortFixedIps     types.List     `tfsdk:"port_fixed_ips"`
	ResourceType     types.String   `tfsdk:"resource_type"`
	RuntimeState     types.String   `tfsdk:"runtime_state"`
	State            types.String   `tfsdk:"state"`
	Tenant           types.String   `tfsdk:"tenant"`
	TenantName       types.String   `tfsdk:"tenant_name"`
	TenantUuid       types.String   `tfsdk:"tenant_uuid"`
	Url              types.String   `tfsdk:"url"`
	Timeouts         timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackFloatingIpResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_floating_ip"
}

func (r *OpenstackFloatingIpResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "OpenstackFloatingIp resource",

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
				MarkdownDescription: "",
			},
			"address": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The public IPv4 address of the floating IP",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"backend_network_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of network in OpenStack where this floating IP is allocated",
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
			"external_address": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Optional address that maps to floating IP's address in external networks",
			},
			"instance_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"instance_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"instance_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"port": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"port_fixed_ips": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "IP address to assign to the port",
						},
						"subnet_id": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "ID of the subnet in which to assign the IP address",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"runtime_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"tenant": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "OpenStack tenant this floating IP belongs to",
			},
			"tenant_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
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

func (r *OpenstackFloatingIpResource) convertTFValue(v attr.Value) interface{} {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	switch val := v.(type) {
	case types.String:
		return val.ValueString()
	case types.Int64:
		return val.ValueInt64()
	case types.Bool:
		return val.ValueBool()
	case types.Float64:
		return val.ValueFloat64()
	case types.List:
		items := make([]interface{}, len(val.Elements()))
		for i, elem := range val.Elements() {
			items[i] = r.convertTFValue(elem)
		}
		return items
	case types.Object:
		obj := make(map[string]interface{})
		for k, attr := range val.Attributes() {
			if converted := r.convertTFValue(attr); converted != nil {
				obj[k] = converted
			}
		}
		return obj
	}
	return nil
}

func (r *OpenstackFloatingIpResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OpenstackFloatingIpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackFloatingIpResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Prepare request body
	requestBody := map[string]interface{}{}

	// Call Waldur API to create resource
	var result map[string]interface{}
	// Custom create operation via parent resource
	createPath := "/api/openstack-tenants/{uuid}/create_floating_ip/"
	createPath = strings.Replace(createPath, "{uuid}", data.Tenant.ValueString(), 1)
	err := r.client.Post(ctx, createPath, requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create OpenstackFloatingIp",
			"An error occurred while creating the openstack_floating_ip: "+err.Error(),
		)
		return
	}
	// Extract UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
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

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackFloatingIpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackFloatingIpResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource
	var result map[string]interface{}

	retrievePath := strings.Replace("/api/openstack-floating-ips/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read OpenstackFloatingIp",
			"An error occurred while reading the openstack_floating_ip: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
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

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackFloatingIpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackFloatingIpResourceModel
	var state OpenstackFloatingIpResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	// Read current state to get the UUID (which is computed and not in plan)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Use UUID from state
	data.UUID = state.UUID
	// Prepare request body
	requestBody := map[string]interface{}{}

	// Call Waldur API to update resource
	var result map[string]interface{}

	err := r.client.Update(ctx, "<no value>", data.UUID.ValueString(), requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update OpenstackFloatingIp",
			"An error occurred while updating the openstack_floating_ip: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
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

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackFloatingIpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackFloatingIpResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Call Waldur API to delete resource
	err := r.client.DeleteByUUID(ctx, "/api/openstack-floating-ips/{uuid}/", data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete OpenstackFloatingIp",
			"An error occurred while deleting the openstack_floating_ip: "+err.Error(),
		)
		return
	}
}

func (r *OpenstackFloatingIpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
