package resources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
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

// OpenstackSecurityGroupApiResponse is the API response model.
type OpenstackSecurityGroupApiResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl      *string                               `json:"access_url" tfsdk:"access_url"`
	BackendId      *string                               `json:"backend_id" tfsdk:"backend_id"`
	Created        *string                               `json:"created" tfsdk:"created"`
	Description    *string                               `json:"description" tfsdk:"description"`
	ErrorMessage   *string                               `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback *string                               `json:"error_traceback" tfsdk:"error_traceback"`
	Modified       *string                               `json:"modified" tfsdk:"modified"`
	Name           *string                               `json:"name" tfsdk:"name"`
	ResourceType   *string                               `json:"resource_type" tfsdk:"resource_type"`
	Rules          []OpenstackSecurityGroupRulesResponse `json:"rules" tfsdk:"rules"`
	State          *string                               `json:"state" tfsdk:"state"`
	Tenant         *string                               `json:"tenant" tfsdk:"tenant"`
	TenantName     *string                               `json:"tenant_name" tfsdk:"tenant_name"`
	TenantUuid     *string                               `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Url            *string                               `json:"url" tfsdk:"url"`
}

type OpenstackSecurityGroupRulesResponse struct {
	Cidr        *string `json:"cidr" tfsdk:"cidr"`
	Description *string `json:"description" tfsdk:"description"`
	Direction   *string `json:"direction" tfsdk:"direction"`
	Ethertype   *string `json:"ethertype" tfsdk:"ethertype"`
	FromPort    *int64  `json:"from_port" tfsdk:"from_port"`
	Protocol    *string `json:"protocol" tfsdk:"protocol"`
	RemoteGroup *string `json:"remote_group" tfsdk:"remote_group"`
	ToPort      *int64  `json:"to_port" tfsdk:"to_port"`
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
		MarkdownDescription: "Openstack Security Group resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"access_url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Access url",
			},
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "ID of the backend",
			},
			"created": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Created",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Description of the resource",
			},
			"error_message": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Error message",
			},
			"error_traceback": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Error traceback",
			},
			"modified": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name of the resource",
			},
			"resource_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Resource type",
			},
			"rules": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cidr": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "CIDR notation for the source/destination network address range",
						},
						"description": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Description of the resource",
						},
						"direction": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Traffic direction - either 'ingress' (incoming) or 'egress' (outgoing)",
						},
						"ethertype": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "IP protocol version - either 'IPv4' or 'IPv6'",
						},
						"from_port": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Starting port number in the range (1-65535)",
						},
						"protocol": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "The network protocol (TCP, UDP, ICMP, or empty for any protocol)",
						},
						"remote_group": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Remote security group that this rule references, if any",
						},
						"to_port": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Ending port number in the range (1-65535)",
						},
					},
				},
				Required:            true,
				MarkdownDescription: "Rules",
			},
			"state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "State",
			},
			"tenant": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Required path parameter for resource creation",
			},
			"tenant_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the tenant",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the tenant",
			},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Url",
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

	// Call Waldur API to create resource
	var apiResp OpenstackSecurityGroupApiResponse
	// Custom create operation via parent resource
	createPath := "/api/openstack-tenants/{uuid}/create_security_group/"
	createPath = strings.Replace(createPath, "{uuid}", data.Tenant.ValueString(), 1) // Prepare request body
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
	err := r.client.Post(ctx, createPath, requestBody, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Openstack Security Group",
			"An error occurred while creating the Openstack Security Group: "+err.Error(),
		)
		return
	}
	// Extract UUID from response
	if apiResp.UUID != nil {
		data.UUID = types.StringPointerValue(apiResp.UUID)
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

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

	retrievePath := strings.Replace("/api/openstack-security-groups/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	var apiResp OpenstackSecurityGroupApiResponse
	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &apiResp)
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Openstack Security Group",
			"An error occurred while reading the Openstack Security Group: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

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
	var apiResp OpenstackSecurityGroupApiResponse

	err := r.client.Update(ctx, "/api/openstack-security-groups/{uuid}/", data.UUID.ValueString(), requestBody, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Openstack Security Group",
			"An error occurred while updating the Openstack Security Group: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if apiResp.UUID != nil {
		data.UUID = types.StringPointerValue(apiResp.UUID)
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

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
			"Unable to Delete Openstack Security Group",
			"An error occurred while deleting the Openstack Security Group: "+err.Error(),
		)
		return
	}
}

func (r *OpenstackSecurityGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *OpenstackSecurityGroupResource) mapResponseToModel(ctx context.Context, apiResp OpenstackSecurityGroupApiResponse, model *OpenstackSecurityGroupResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	listValRules, listDiagsRules := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"cidr":         types.StringType,
		"description":  types.StringType,
		"direction":    types.StringType,
		"ethertype":    types.StringType,
		"from_port":    types.Int64Type,
		"protocol":     types.StringType,
		"remote_group": types.StringType,
		"to_port":      types.Int64Type,
	}}, apiResp.Rules)
	diags.Append(listDiagsRules...)
	model.Rules = listValRules
	model.State = types.StringPointerValue(apiResp.State)
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantName = types.StringPointerValue(apiResp.TenantName)
	model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
