package floating_ip

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackFloatingIpResource{}
var _ resource.ResourceWithImportState = &OpenstackFloatingIpResource{}

func NewOpenstackFloatingIpResource() resource.Resource {
	return &OpenstackFloatingIpResource{}
}

// OpenstackFloatingIpResource defines the resource implementation.
type OpenstackFloatingIpResource struct {
	client *Client
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
		MarkdownDescription: "Openstack Floating Ip resource",

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
			"address": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The public IPv4 address of the floating IP",
			},
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "ID of the backend",
			},
			"backend_network_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "ID of network in OpenStack where this floating IP is allocated",
			},
			"created": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Created",
			},
			"description": schema.StringAttribute{
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
			"external_address": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Optional address that maps to floating IP's address in external networks",
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the instance",
			},
			"instance_url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Instance url",
			},
			"instance_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the instance",
			},
			"modified": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the resource",
			},
			"port": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Port",
			},
			"port_fixed_ips": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "IP address to assign to the port",
						},
						"subnet_id": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "ID of the subnet in which to assign the IP address",
						},
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Port fixed ips",
			},
			"resource_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Resource type",
			},
			"runtime_state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Runtime state",
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

	r.client = NewClient(client)
}

func (r *OpenstackFloatingIpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackFloatingIpResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := OpenstackFloatingIpCreateRequest{}

	apiResp, err := r.client.CreateOpenstackFloatingIp(ctx, data.Tenant.ValueString(), &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Openstack Floating Ip",
			"An error occurred while creating the Openstack Floating Ip: "+err.Error(),
		)
		return
	}
	data.UUID = types.StringPointerValue(apiResp.UUID)

	createTimeout, diags := data.Timeouts.Create(ctx, 30*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackFloatingIpResponse, error) {
		return r.client.GetOpenstackFloatingIp(ctx, data.UUID.ValueString())
	}, createTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource creation", err.Error())
		return
	}
	apiResp = newResp

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

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

	apiResp, err := r.client.GetOpenstackFloatingIp(ctx, data.UUID.ValueString())
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Openstack Floating Ip",
			"An error occurred while reading the Openstack Floating Ip: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackFloatingIpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update Not Supported", "This resource cannot be updated via the API.")
}

func (r *OpenstackFloatingIpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackFloatingIpResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteOpenstackFloatingIp(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Openstack Floating Ip",
			"An error occurred while deleting the Openstack Floating Ip: "+err.Error(),
		)
		return
	}

	deleteTimeout, diags := data.Timeouts.Delete(ctx, 10*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err = common.WaitForDeletion(ctx, func(ctx context.Context) (*OpenstackFloatingIpResponse, error) {
		return r.client.GetOpenstackFloatingIp(ctx, data.UUID.ValueString())
	}, deleteTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource deletion", err.Error())
		return
	}
}

func (r *OpenstackFloatingIpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *OpenstackFloatingIpResource) mapResponseToModel(ctx context.Context, apiResp OpenstackFloatingIpResponse, model *OpenstackFloatingIpResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.Address = types.StringPointerValue(apiResp.Address)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.BackendNetworkId = types.StringPointerValue(apiResp.BackendNetworkId)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.ExternalAddress = types.StringPointerValue(apiResp.ExternalAddress)
	model.InstanceName = types.StringPointerValue(apiResp.InstanceName)
	model.InstanceUrl = types.StringPointerValue(apiResp.InstanceUrl)
	model.InstanceUuid = types.StringPointerValue(apiResp.InstanceUuid)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.Port = types.StringPointerValue(apiResp.Port)

	{
		listValPortFixedIps, listDiagsPortFixedIps := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"ip_address": types.StringType,
			"subnet_id":  types.StringType,
		}}, apiResp.PortFixedIps)
		diags.Append(listDiagsPortFixedIps...)
		model.PortFixedIps = listValPortFixedIps
	}
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.RuntimeState = types.StringPointerValue(apiResp.RuntimeState)
	model.State = types.StringPointerValue(apiResp.State)
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantName = types.StringPointerValue(apiResp.TenantName)
	model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
