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
var _ resource.Resource = &OpenstackServerGroupResource{}
var _ resource.ResourceWithImportState = &OpenstackServerGroupResource{}

func NewOpenstackServerGroupResource() resource.Resource {
	return &OpenstackServerGroupResource{}
}

// OpenstackServerGroupResource defines the resource implementation.
type OpenstackServerGroupResource struct {
	client *client.Client
}

// OpenstackServerGroupResourceModel describes the resource data model.
type OpenstackServerGroupResourceModel struct {
	UUID           types.String   `tfsdk:"id"`
	AccessUrl      types.String   `tfsdk:"access_url"`
	BackendId      types.String   `tfsdk:"backend_id"`
	Created        types.String   `tfsdk:"created"`
	Description    types.String   `tfsdk:"description"`
	DisplayName    types.String   `tfsdk:"display_name"`
	ErrorMessage   types.String   `tfsdk:"error_message"`
	ErrorTraceback types.String   `tfsdk:"error_traceback"`
	Instances      types.List     `tfsdk:"instances"`
	Modified       types.String   `tfsdk:"modified"`
	Name           types.String   `tfsdk:"name"`
	Policy         types.String   `tfsdk:"policy"`
	ResourceType   types.String   `tfsdk:"resource_type"`
	State          types.String   `tfsdk:"state"`
	Tenant         types.String   `tfsdk:"tenant"`
	TenantName     types.String   `tfsdk:"tenant_name"`
	TenantUuid     types.String   `tfsdk:"tenant_uuid"`
	Url            types.String   `tfsdk:"url"`
	Timeouts       timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackServerGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_server_group"
}

func (r *OpenstackServerGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "OpenstackServerGroup resource",

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
			"display_name": schema.StringAttribute{
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
			"instances": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"backend_id": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Instance ID in the OpenStack backend",
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
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: " ",
			},
			"policy": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Server group policy determining the rules for scheduling servers in this group",
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

func (r *OpenstackServerGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OpenstackServerGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackServerGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Prepare request body
	requestBody := map[string]interface{}{}
	// Check if this field is a path param (skip adding to body)
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		if v := data.Description.ValueString(); v != "" {
			requestBody["description"] = v
		}
	}
	// Check if this field is a path param (skip adding to body)
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		if v := data.Name.ValueString(); v != "" {
			requestBody["name"] = v
		}
	}
	// Check if this field is a path param (skip adding to body)
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {
		if v := data.Policy.ValueString(); v != "" {
			requestBody["policy"] = v
		}
	}
	// Check if this field is a path param (skip adding to body)

	// Call Waldur API to create resource
	var result map[string]interface{}
	// Custom create operation via parent resource
	createPath := "/api/openstack-tenants/{uuid}/create_server_group/"
	createPath = strings.Replace(createPath, "{uuid}", data.Tenant.ValueString(), 1)
	err := r.client.Post(ctx, createPath, requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create OpenstackServerGroup",
			"An error occurred while creating the openstack_server_group: "+err.Error(),
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
	if val, ok := sourceMap["display_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.DisplayName = types.StringValue(str)
		}
	} else {
		if data.DisplayName.IsUnknown() {
			data.DisplayName = types.StringNull()
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
	if val, ok := sourceMap["instances"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"backend_id": types.StringType,
						"name":       types.StringType,
					}
					attrValues := map[string]attr.Value{
						"backend_id": func() attr.Value {
							if v, ok := objMap["backend_id"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
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
				"backend_id": types.StringType,
				"name":       types.StringType,
			}}, items)
			data.Instances = listVal
		}
	} else {
		if data.Instances.IsUnknown() {
			data.Instances = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"backend_id": types.StringType,
				"name":       types.StringType,
			}})
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
	if val, ok := sourceMap["policy"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Policy = types.StringValue(str)
		}
	} else {
		if data.Policy.IsUnknown() {
			data.Policy = types.StringNull()
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

func (r *OpenstackServerGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackServerGroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource
	var result map[string]interface{}

	retrievePath := strings.Replace("/api/openstack-server-groups/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read OpenstackServerGroup",
			"An error occurred while reading the openstack_server_group: "+err.Error(),
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
	if val, ok := sourceMap["display_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.DisplayName = types.StringValue(str)
		}
	} else {
		if data.DisplayName.IsUnknown() {
			data.DisplayName = types.StringNull()
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
	if val, ok := sourceMap["instances"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"backend_id": types.StringType,
						"name":       types.StringType,
					}
					attrValues := map[string]attr.Value{
						"backend_id": func() attr.Value {
							if v, ok := objMap["backend_id"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
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
				"backend_id": types.StringType,
				"name":       types.StringType,
			}}, items)
			data.Instances = listVal
		}
	} else {
		if data.Instances.IsUnknown() {
			data.Instances = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"backend_id": types.StringType,
				"name":       types.StringType,
			}})
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
	if val, ok := sourceMap["policy"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Policy = types.StringValue(str)
		}
	} else {
		if data.Policy.IsUnknown() {
			data.Policy = types.StringNull()
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

func (r *OpenstackServerGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackServerGroupResourceModel
	var state OpenstackServerGroupResourceModel
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
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {
		if v := data.Policy.ValueString(); v != "" {
			requestBody["policy"] = v
		}
	}

	// Call Waldur API to update resource
	var result map[string]interface{}

	err := r.client.Update(ctx, "/api/openstack-server-groups/{uuid}/", data.UUID.ValueString(), requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update OpenstackServerGroup",
			"An error occurred while updating the openstack_server_group: "+err.Error(),
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
	if val, ok := sourceMap["display_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.DisplayName = types.StringValue(str)
		}
	} else {
		if data.DisplayName.IsUnknown() {
			data.DisplayName = types.StringNull()
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
	if val, ok := sourceMap["instances"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"backend_id": types.StringType,
						"name":       types.StringType,
					}
					attrValues := map[string]attr.Value{
						"backend_id": func() attr.Value {
							if v, ok := objMap["backend_id"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
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
				"backend_id": types.StringType,
				"name":       types.StringType,
			}}, items)
			data.Instances = listVal
		}
	} else {
		if data.Instances.IsUnknown() {
			data.Instances = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"backend_id": types.StringType,
				"name":       types.StringType,
			}})
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
	if val, ok := sourceMap["policy"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Policy = types.StringValue(str)
		}
	} else {
		if data.Policy.IsUnknown() {
			data.Policy = types.StringNull()
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

func (r *OpenstackServerGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackServerGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to delete resource
	err := r.client.DeleteByUUID(ctx, "/api/openstack-server-groups/{uuid}/", data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete OpenstackServerGroup",
			"An error occurred while deleting the openstack_server_group: "+err.Error(),
		)
		return
	}
}

func (r *OpenstackServerGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
