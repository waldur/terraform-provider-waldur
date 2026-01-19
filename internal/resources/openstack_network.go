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
var _ resource.Resource = &OpenstackNetworkResource{}
var _ resource.ResourceWithImportState = &OpenstackNetworkResource{}

func NewOpenstackNetworkResource() resource.Resource {
	return &OpenstackNetworkResource{}
}

// OpenstackNetworkResource defines the resource implementation.
type OpenstackNetworkResource struct {
	client *client.Client
}

// OpenstackNetworkResourceModel describes the resource data model.
type OpenstackNetworkResourceModel struct {
	UUID           types.String   `tfsdk:"id"`
	AccessUrl      types.String   `tfsdk:"access_url"`
	BackendId      types.String   `tfsdk:"backend_id"`
	Created        types.String   `tfsdk:"created"`
	Description    types.String   `tfsdk:"description"`
	ErrorMessage   types.String   `tfsdk:"error_message"`
	ErrorTraceback types.String   `tfsdk:"error_traceback"`
	IsExternal     types.Bool     `tfsdk:"is_external"`
	Modified       types.String   `tfsdk:"modified"`
	Mtu            types.Int64    `tfsdk:"mtu"`
	Name           types.String   `tfsdk:"name"`
	RbacPolicies   types.List     `tfsdk:"rbac_policies"`
	ResourceType   types.String   `tfsdk:"resource_type"`
	SegmentationId types.Int64    `tfsdk:"segmentation_id"`
	State          types.String   `tfsdk:"state"`
	Subnets        types.List     `tfsdk:"subnets"`
	Tenant         types.String   `tfsdk:"tenant"`
	TenantName     types.String   `tfsdk:"tenant_name"`
	TenantUuid     types.String   `tfsdk:"tenant_uuid"`
	Type           types.String   `tfsdk:"type"`
	Url            types.String   `tfsdk:"url"`
	Timeouts       timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackNetworkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_network"
}

func (r *OpenstackNetworkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "OpenstackNetwork resource",

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
			"backend_id": schema.StringAttribute{
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
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"is_external": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Defines whether this network is external (public) or internal (private)",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"mtu": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The maximum transmission unit (MTU) value to address fragmentation.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: " ",
			},
			"rbac_policies": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"backend_id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"created": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"network": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"network_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"policy_type": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Type of access granted - either shared access or external network access",
						},
						"target_tenant": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"target_tenant_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"segmentation_id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "VLAN ID for VLAN networks or tunnel ID for VXLAN/GRE networks",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"subnets": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allocation_pools": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"end": schema.StringAttribute{
										Optional:            true,
										Computed:            true,
										MarkdownDescription: "An IPv4 or IPv6 address.",
									},
									"start": schema.StringAttribute{
										Optional:            true,
										Computed:            true,
										MarkdownDescription: "An IPv4 or IPv6 address.",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: " ",
						},
						"cidr": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "IPv4 network address in CIDR format (e.g. 192.168.0.0/24)",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"enable_dhcp": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "If True, DHCP service will be enabled on this subnet",
						},
						"gateway_ip": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "IP address of the gateway for this subnet",
						},
						"ip_version": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "IP protocol version (4 or 6)",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"tenant": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "OpenStack tenant this network belongs to",
			},
			"tenant_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Network type, such as local, flat, vlan, vxlan, or gre",
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

func (r *OpenstackNetworkResource) convertTFValue(v attr.Value) interface{} {
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

func (r *OpenstackNetworkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OpenstackNetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackNetworkResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
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

	// Call Waldur API to create resource
	var result map[string]interface{}
	// Custom create operation via parent resource
	createPath := "/api/openstack-tenants/{uuid}/create_network/"
	createPath = strings.Replace(createPath, "{uuid}", data.Tenant.ValueString(), 1)
	err := r.client.Post(ctx, createPath, requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create OpenstackNetwork",
			"An error occurred while creating the openstack_network: "+err.Error(),
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
						"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"end": types.StringType, "start": types.StringType}}},
						"cidr":             types.StringType,
						"description":      types.StringType,
						"enable_dhcp":      types.BoolType,
						"gateway_ip":       types.StringType,
						"ip_version":       types.Int64Type,
						"name":             types.StringType,
					}
					attrValues := map[string]attr.Value{
						"allocation_pools": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"end": types.StringType, "start": types.StringType}}}.ElemType),
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
				"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"end": types.StringType, "start": types.StringType}}},
				"cidr":             types.StringType,
				"description":      types.StringType,
				"enable_dhcp":      types.BoolType,
				"gateway_ip":       types.StringType,
				"ip_version":       types.Int64Type,
				"name":             types.StringType,
			}}, items)
			data.Subnets = listVal
		}
	} else {
		if data.Subnets.IsUnknown() {
			data.Subnets = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"end": types.StringType, "start": types.StringType}}},
				"cidr":             types.StringType,
				"description":      types.StringType,
				"enable_dhcp":      types.BoolType,
				"gateway_ip":       types.StringType,
				"ip_version":       types.Int64Type,
				"name":             types.StringType,
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

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackNetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackNetworkResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource
	var result map[string]interface{}

	retrievePath := strings.Replace("/api/openstack-networks/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read OpenstackNetwork",
			"An error occurred while reading the openstack_network: "+err.Error(),
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
						"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"end": types.StringType, "start": types.StringType}}},
						"cidr":             types.StringType,
						"description":      types.StringType,
						"enable_dhcp":      types.BoolType,
						"gateway_ip":       types.StringType,
						"ip_version":       types.Int64Type,
						"name":             types.StringType,
					}
					attrValues := map[string]attr.Value{
						"allocation_pools": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"end": types.StringType, "start": types.StringType}}}.ElemType),
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
				"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"end": types.StringType, "start": types.StringType}}},
				"cidr":             types.StringType,
				"description":      types.StringType,
				"enable_dhcp":      types.BoolType,
				"gateway_ip":       types.StringType,
				"ip_version":       types.Int64Type,
				"name":             types.StringType,
			}}, items)
			data.Subnets = listVal
		}
	} else {
		if data.Subnets.IsUnknown() {
			data.Subnets = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"end": types.StringType, "start": types.StringType}}},
				"cidr":             types.StringType,
				"description":      types.StringType,
				"enable_dhcp":      types.BoolType,
				"gateway_ip":       types.StringType,
				"ip_version":       types.Int64Type,
				"name":             types.StringType,
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

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackNetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackNetworkResourceModel
	var state OpenstackNetworkResourceModel

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

	// Call Waldur API to update resource
	var result map[string]interface{}

	err := r.client.Update(ctx, "/api/openstack-networks/{uuid}/", data.UUID.ValueString(), requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update OpenstackNetwork",
			"An error occurred while updating the openstack_network: "+err.Error(),
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
						"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"end": types.StringType, "start": types.StringType}}},
						"cidr":             types.StringType,
						"description":      types.StringType,
						"enable_dhcp":      types.BoolType,
						"gateway_ip":       types.StringType,
						"ip_version":       types.Int64Type,
						"name":             types.StringType,
					}
					attrValues := map[string]attr.Value{
						"allocation_pools": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"end": types.StringType, "start": types.StringType}}}.ElemType),
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
				"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"end": types.StringType, "start": types.StringType}}},
				"cidr":             types.StringType,
				"description":      types.StringType,
				"enable_dhcp":      types.BoolType,
				"gateway_ip":       types.StringType,
				"ip_version":       types.Int64Type,
				"name":             types.StringType,
			}}, items)
			data.Subnets = listVal
		}
	} else {
		if data.Subnets.IsUnknown() {
			data.Subnets = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"end": types.StringType, "start": types.StringType}}},
				"cidr":             types.StringType,
				"description":      types.StringType,
				"enable_dhcp":      types.BoolType,
				"gateway_ip":       types.StringType,
				"ip_version":       types.Int64Type,
				"name":             types.StringType,
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

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackNetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackNetworkResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Call Waldur API to delete resource
	err := r.client.DeleteByUUID(ctx, "/api/openstack-networks/{uuid}/", data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete OpenstackNetwork",
			"An error occurred while deleting the openstack_network: "+err.Error(),
		)
		return
	}
}

func (r *OpenstackNetworkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
