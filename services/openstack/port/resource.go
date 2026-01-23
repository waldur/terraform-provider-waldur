package port

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackPortResource{}
var _ resource.ResourceWithImportState = &OpenstackPortResource{}

func NewOpenstackPortResource() resource.Resource {
	return &OpenstackPortResource{}
}

// OpenstackPortResource defines the resource implementation.
type OpenstackPortResource struct {
	client *Client
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
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Access url",
			},
			"admin_state_up": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Administrative state of the port. If down, port does not forward packets",
			},
			"allowed_address_pairs": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Ip address",
						},
						"mac_address": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Mac address",
						},
					},
				},
				Optional: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Allowed address pairs",
			},
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Port ID in OpenStack",
			},
			"created": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Created",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description of the resource",
			},
			"device_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "ID of device (instance, router etc) to which this port is connected",
			},
			"device_owner": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Entity that uses this port (e.g. network:router_interface)",
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
				Optional: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Fixed ips",
			},
			"floating_ips": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.StringType},
				Computed:   true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Floating ips",
			},
			"mac_address": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "MAC address of the port",
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
			"network": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Network to which this port belongs",
			},
			"network_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the network",
			},
			"network_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the network",
			},
			"port_security_enabled": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "If True, security groups and rules will be applied to this port",
			},
			"resource_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Resource type",
			},
			"security_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Name of the resource",
						},
					},
				},
				Optional:            true,
				MarkdownDescription: "Security groups",
			},
			"state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "State",
			},
			"status": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Port status in OpenStack (e.g. ACTIVE, DOWN)",
			},
			"target_tenant": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Target tenant for shared network port creation. If not specified, defaults to network's tenant.",
			},
			"tenant": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "OpenStack tenant this port belongs to",
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

	r.client = NewClient(client)
}

func (r *OpenstackPortResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackPortResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var requestBody OpenstackPortCreateRequest // Prepare request body
	{
		var items []common.OpenStackAllowedAddressPairRequest
		if diags := data.AllowedAddressPairs.ElementsAs(ctx, &items, false); !diags.HasError() && len(items) > 0 {
			requestBody.AllowedAddressPairs = items
		}
	}
	requestBody.Description = data.Description.ValueStringPointer()
	{
		var items []common.OpenStackFixedIpRequest
		if diags := data.FixedIps.ElementsAs(ctx, &items, false); !diags.HasError() && len(items) > 0 {
			requestBody.FixedIps = items
		}
	}
	requestBody.MacAddress = data.MacAddress.ValueStringPointer()
	requestBody.Name = data.Name.ValueStringPointer()
	requestBody.Network = data.Network.ValueStringPointer()
	requestBody.PortSecurityEnabled = data.PortSecurityEnabled.ValueBoolPointer()
	{
		var items []common.OpenStackPortNestedSecurityGroupRequest
		if diags := data.SecurityGroups.ElementsAs(ctx, &items, false); !diags.HasError() && len(items) > 0 {
			requestBody.SecurityGroups = items
		}
	}
	requestBody.TargetTenant = data.TargetTenant.ValueStringPointer()

	apiResp, err := r.client.CreateOpenstackPort(ctx, &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Openstack Port",
			"An error occurred while creating the Openstack Port: "+err.Error(),
		)
		return
	}
	data.UUID = types.StringPointerValue(apiResp.UUID)

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

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

	apiResp, err := r.client.GetOpenstackPort(ctx, data.UUID.ValueString())
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Openstack Port",
			"An error occurred while reading the Openstack Port: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

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

	var requestBody OpenstackPortUpdateRequest // Prepare request body
	requestBody.Description = data.Description.ValueStringPointer()
	requestBody.Name = data.Name.ValueStringPointer()
	{
		var items []common.OpenStackPortNestedSecurityGroupRequest
		if diags := data.SecurityGroups.ElementsAs(ctx, &items, false); !diags.HasError() && len(items) > 0 {
			requestBody.SecurityGroups = items
		}
	}
	requestBody.TargetTenant = data.TargetTenant.ValueStringPointer()

	apiResp, err := r.client.UpdateOpenstackPort(ctx, data.UUID.ValueString(), &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Openstack Port",
			"An error occurred while updating the Openstack Port: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if apiResp.UUID != nil {
		data.UUID = types.StringPointerValue(apiResp.UUID)
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackPortResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackPortResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteOpenstackPort(ctx, data.UUID.ValueString())
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

func (r *OpenstackPortResource) mapResponseToModel(ctx context.Context, apiResp OpenstackPortResponse, model *OpenstackPortResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.AdminStateUp = types.BoolPointerValue(apiResp.AdminStateUp)
	listValAllowedAddressPairs, listDiagsAllowedAddressPairs := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"ip_address":  types.StringType,
		"mac_address": types.StringType,
	}}, apiResp.AllowedAddressPairs)
	diags.Append(listDiagsAllowedAddressPairs...)
	model.AllowedAddressPairs = listValAllowedAddressPairs
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.DeviceId = types.StringPointerValue(apiResp.DeviceId)
	model.DeviceOwner = types.StringPointerValue(apiResp.DeviceOwner)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	listValFixedIps, listDiagsFixedIps := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"ip_address": types.StringType,
		"subnet_id":  types.StringType,
	}}, apiResp.FixedIps)
	diags.Append(listDiagsFixedIps...)
	model.FixedIps = listValFixedIps
	model.FloatingIps, _ = types.ListValueFrom(ctx, types.StringType, apiResp.FloatingIps)
	model.MacAddress = types.StringPointerValue(apiResp.MacAddress)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.Network = types.StringPointerValue(apiResp.Network)
	model.NetworkName = types.StringPointerValue(apiResp.NetworkName)
	model.NetworkUuid = types.StringPointerValue(apiResp.NetworkUuid)
	model.PortSecurityEnabled = types.BoolPointerValue(apiResp.PortSecurityEnabled)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	listValSecurityGroups, listDiagsSecurityGroups := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"name": types.StringType,
	}}, apiResp.SecurityGroups)
	diags.Append(listDiagsSecurityGroups...)
	model.SecurityGroups = listValSecurityGroups
	model.State = types.StringPointerValue(apiResp.State)
	model.Status = types.StringPointerValue(apiResp.Status)
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantName = types.StringPointerValue(apiResp.TenantName)
	model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
