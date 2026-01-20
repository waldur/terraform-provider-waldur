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
var _ resource.Resource = &OpenstackSecurityGroupResource{}
var _ resource.ResourceWithImportState = &OpenstackSecurityGroupResource{}

func NewOpenstackSecurityGroupResource() resource.Resource {
	return &OpenstackSecurityGroupResource{}
}

// OpenstackSecurityGroupResource defines the resource implementation.
type OpenstackSecurityGroupResource struct {
	client *client.Client
}

// OpenstackSecurityGroupResourceModel describes the resource data model.
type OpenstackSecurityGroupResourceModel struct {
	UUID           types.String   `tfsdk:"id"`
	AccessUrl      types.String   `tfsdk:"access_url"`
	BackendId      types.String   `tfsdk:"backend_id"`
	Created        types.String   `tfsdk:"created"`
	Description    types.String   `tfsdk:"description"`
	ErrorMessage   types.String   `tfsdk:"error_message"`
	ErrorTraceback types.String   `tfsdk:"error_traceback"`
	Modified       types.String   `tfsdk:"modified"`
	Name           types.String   `tfsdk:"name"`
	ResourceType   types.String   `tfsdk:"resource_type"`
	Rules          types.List     `tfsdk:"rules"`
	State          types.String   `tfsdk:"state"`
	Tenant         types.String   `tfsdk:"tenant"`
	TenantName     types.String   `tfsdk:"tenant_name"`
	TenantUuid     types.String   `tfsdk:"tenant_uuid"`
	Url            types.String   `tfsdk:"url"`
	Timeouts       timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackSecurityGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_security_group"
}

func (r *OpenstackSecurityGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "OpenstackSecurityGroup resource",

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
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: " ",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"rules": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cidr": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "CIDR notation for the source/destination network address range",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"direction": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Traffic direction - either 'ingress' (incoming) or 'egress' (outgoing)",
						},
						"ethertype": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "IP protocol version - either 'IPv4' or 'IPv6'",
						},
						"from_port": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Starting port number in the range (1-65535)",
						},
						"protocol": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "The network protocol (TCP, UDP, ICMP, or empty for any protocol)",
						},
						"remote_group": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Remote security group that this rule references, if any",
						},
						"to_port": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Ending port number in the range (1-65535)",
						},
					},
				},
				Required:            true,
				MarkdownDescription: " ",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"tenant": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Required path parameter for resource creation",
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

func (r *OpenstackSecurityGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OpenstackSecurityGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackSecurityGroupResourceModel
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
	requestBody["name"] = data.Name.ValueString()
	if v := ConvertTFValue(data.Rules); v != nil {
		requestBody["rules"] = v
	}

	// Call Waldur API to create resource
	var result map[string]interface{}
	// Custom create operation via parent resource
	createPath := "/api/openstack-tenants/{uuid}/create_security_group/"
	createPath = strings.Replace(createPath, "{uuid}", data.Tenant.ValueString(), 1)
	err := r.client.Post(ctx, createPath, requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create OpenstackSecurityGroup",
			"An error occurred while creating the openstack_security_group: "+err.Error(),
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
	if val, ok := sourceMap["modified"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Modified = types.StringValue(str)
		}
	} else {
		if data.Modified.IsUnknown() {
			data.Modified = types.StringNull()
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
	if val, ok := sourceMap["rules"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"cidr":         types.StringType,
						"description":  types.StringType,
						"direction":    types.StringType,
						"ethertype":    types.StringType,
						"from_port":    types.Int64Type,
						"protocol":     types.StringType,
						"remote_group": types.StringType,
						"to_port":      types.Int64Type,
					}
					attrValues := map[string]attr.Value{
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
						"direction": func() attr.Value {
							if v, ok := objMap["direction"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"ethertype": func() attr.Value {
							if v, ok := objMap["ethertype"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"from_port": func() attr.Value {
							if v, ok := objMap["from_port"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"protocol": func() attr.Value {
							if v, ok := objMap["protocol"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"remote_group": func() attr.Value {
							if v, ok := objMap["remote_group"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"to_port": func() attr.Value {
							if v, ok := objMap["to_port"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"cidr":         types.StringType,
				"description":  types.StringType,
				"direction":    types.StringType,
				"ethertype":    types.StringType,
				"from_port":    types.Int64Type,
				"protocol":     types.StringType,
				"remote_group": types.StringType,
				"to_port":      types.Int64Type,
			}}, items)
			data.Rules = listVal
		}
	} else {
		if data.Rules.IsUnknown() {
			data.Rules = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"cidr":         types.StringType,
				"description":  types.StringType,
				"direction":    types.StringType,
				"ethertype":    types.StringType,
				"from_port":    types.Int64Type,
				"protocol":     types.StringType,
				"remote_group": types.StringType,
				"to_port":      types.Int64Type,
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackSecurityGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackSecurityGroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource
	var result map[string]interface{}

	retrievePath := strings.Replace("/api/openstack-security-groups/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read OpenstackSecurityGroup",
			"An error occurred while reading the openstack_security_group: "+err.Error(),
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
	if val, ok := sourceMap["modified"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Modified = types.StringValue(str)
		}
	} else {
		if data.Modified.IsUnknown() {
			data.Modified = types.StringNull()
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
	if val, ok := sourceMap["rules"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"cidr":         types.StringType,
						"description":  types.StringType,
						"direction":    types.StringType,
						"ethertype":    types.StringType,
						"from_port":    types.Int64Type,
						"protocol":     types.StringType,
						"remote_group": types.StringType,
						"to_port":      types.Int64Type,
					}
					attrValues := map[string]attr.Value{
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
						"direction": func() attr.Value {
							if v, ok := objMap["direction"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"ethertype": func() attr.Value {
							if v, ok := objMap["ethertype"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"from_port": func() attr.Value {
							if v, ok := objMap["from_port"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"protocol": func() attr.Value {
							if v, ok := objMap["protocol"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"remote_group": func() attr.Value {
							if v, ok := objMap["remote_group"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"to_port": func() attr.Value {
							if v, ok := objMap["to_port"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"cidr":         types.StringType,
				"description":  types.StringType,
				"direction":    types.StringType,
				"ethertype":    types.StringType,
				"from_port":    types.Int64Type,
				"protocol":     types.StringType,
				"remote_group": types.StringType,
				"to_port":      types.Int64Type,
			}}, items)
			data.Rules = listVal
		}
	} else {
		if data.Rules.IsUnknown() {
			data.Rules = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"cidr":         types.StringType,
				"description":  types.StringType,
				"direction":    types.StringType,
				"ethertype":    types.StringType,
				"from_port":    types.Int64Type,
				"protocol":     types.StringType,
				"remote_group": types.StringType,
				"to_port":      types.Int64Type,
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

func (r *OpenstackSecurityGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackSecurityGroupResourceModel
	var state OpenstackSecurityGroupResourceModel
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

	// Call Waldur API to update resource
	var result map[string]interface{}

	err := r.client.Update(ctx, "/api/openstack-security-groups/{uuid}/", data.UUID.ValueString(), requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update OpenstackSecurityGroup",
			"An error occurred while updating the openstack_security_group: "+err.Error(),
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
	if val, ok := sourceMap["modified"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Modified = types.StringValue(str)
		}
	} else {
		if data.Modified.IsUnknown() {
			data.Modified = types.StringNull()
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
	if val, ok := sourceMap["rules"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"cidr":         types.StringType,
						"description":  types.StringType,
						"direction":    types.StringType,
						"ethertype":    types.StringType,
						"from_port":    types.Int64Type,
						"protocol":     types.StringType,
						"remote_group": types.StringType,
						"to_port":      types.Int64Type,
					}
					attrValues := map[string]attr.Value{
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
						"direction": func() attr.Value {
							if v, ok := objMap["direction"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"ethertype": func() attr.Value {
							if v, ok := objMap["ethertype"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"from_port": func() attr.Value {
							if v, ok := objMap["from_port"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"protocol": func() attr.Value {
							if v, ok := objMap["protocol"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"remote_group": func() attr.Value {
							if v, ok := objMap["remote_group"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"to_port": func() attr.Value {
							if v, ok := objMap["to_port"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"cidr":         types.StringType,
				"description":  types.StringType,
				"direction":    types.StringType,
				"ethertype":    types.StringType,
				"from_port":    types.Int64Type,
				"protocol":     types.StringType,
				"remote_group": types.StringType,
				"to_port":      types.Int64Type,
			}}, items)
			data.Rules = listVal
		}
	} else {
		if data.Rules.IsUnknown() {
			data.Rules = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"cidr":         types.StringType,
				"description":  types.StringType,
				"direction":    types.StringType,
				"ethertype":    types.StringType,
				"from_port":    types.Int64Type,
				"protocol":     types.StringType,
				"remote_group": types.StringType,
				"to_port":      types.Int64Type,
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackSecurityGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackSecurityGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to delete resource
	err := r.client.DeleteByUUID(ctx, "/api/openstack-security-groups/{uuid}/", data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete OpenstackSecurityGroup",
			"An error occurred while deleting the openstack_security_group: "+err.Error(),
		)
		return
	}
}

func (r *OpenstackSecurityGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
