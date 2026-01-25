package port

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

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
	OpenstackPortModel
	TargetTenant types.String   `tfsdk:"target_tenant"`
	Timeouts     timeouts.Value `tfsdk:"timeouts"`
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
				CustomType: timetypes.RFC3339Type{},
				Computed:   true,
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
			"floating_ips": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
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
				CustomType: timetypes.RFC3339Type{},
				Computed:   true,
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
			"security_groups": schema.SetNestedAttribute{
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

	requestBody := OpenstackPortCreateRequest{
		Description:         data.Description.ValueStringPointer(),
		MacAddress:          data.MacAddress.ValueStringPointer(),
		Name:                data.Name.ValueStringPointer(),
		Network:             data.Network.ValueStringPointer(),
		PortSecurityEnabled: data.PortSecurityEnabled.ValueBoolPointer(),
		TargetTenant:        data.TargetTenant.ValueStringPointer(),
	}
	{
		var items []common.OpenStackAllowedAddressPairRequest
		diags := data.AllowedAddressPairs.ElementsAs(ctx, &items, false)
		resp.Diagnostics.Append(diags...)
		if !diags.HasError() {
			if !data.AllowedAddressPairs.IsNull() && !data.AllowedAddressPairs.IsUnknown() {
				requestBody.AllowedAddressPairs = &items
			}
		}
	}
	{
		var items []common.OpenStackFixedIpRequest
		diags := data.FixedIps.ElementsAs(ctx, &items, false)
		resp.Diagnostics.Append(diags...)
		if !diags.HasError() {
			if !data.FixedIps.IsNull() && !data.FixedIps.IsUnknown() {
				requestBody.FixedIps = &items
			}
		}
	}
	{
		var items []common.OpenStackPortNestedSecurityGroupRequest
		diags := data.SecurityGroups.ElementsAs(ctx, &items, false)
		resp.Diagnostics.Append(diags...)
		if !diags.HasError() {
			if !data.SecurityGroups.IsNull() && !data.SecurityGroups.IsUnknown() {
				requestBody.SecurityGroups = &items
			}
		}
	}

	apiResp, err := r.client.CreateOpenstackPort(ctx, &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Openstack Port",
			"An error occurred while creating the Openstack Port: "+err.Error(),
		)
		return
	}
	data.UUID = types.StringPointerValue(apiResp.UUID)

	createTimeout, diags := data.Timeouts.Create(ctx, common.DefaultCreateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackPortResponse, error) {
		return r.client.GetOpenstackPort(ctx, data.UUID.ValueString())
	}, createTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource creation", err.Error())
		return
	}
	apiResp = newResp

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

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

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

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

	requestBody := OpenstackPortUpdateRequest{
		Description:  data.Description.ValueStringPointer(),
		Name:         data.Name.ValueStringPointer(),
		TargetTenant: data.TargetTenant.ValueStringPointer(),
	}
	{
		var items []common.OpenStackPortNestedSecurityGroupRequest
		diags := data.SecurityGroups.ElementsAs(ctx, &items, false)
		resp.Diagnostics.Append(diags...)
		if !diags.HasError() {
			if !data.SecurityGroups.IsNull() && !data.SecurityGroups.IsUnknown() {
				requestBody.SecurityGroups = &items
			}
		}
	}

	apiResp, err := r.client.UpdateOpenstackPort(ctx, data.UUID.ValueString(), &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Openstack Port",
			"An error occurred while updating the Openstack Port: "+err.Error(),
		)
		return
	}

	updateTimeout, diags := data.Timeouts.Update(ctx, common.DefaultUpdateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackPortResponse, error) {
		return r.client.GetOpenstackPort(ctx, data.UUID.ValueString())
	}, updateTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource update", err.Error())
		return
	}
	apiResp = newResp

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

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

	deleteTimeout, diags := data.Timeouts.Delete(ctx, common.DefaultDeleteTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err = common.WaitForDeletion(ctx, func(ctx context.Context) (*OpenstackPortResponse, error) {
		return r.client.GetOpenstackPort(ctx, data.UUID.ValueString())
	}, deleteTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource deletion", err.Error())
		return
	}
}

func (r *OpenstackPortResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	uuid := req.ID
	if uuid == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Import ID cannot be empty. Please provide the UUID of the Openstack Port.",
		)
		return
	}

	tflog.Info(ctx, "Importing Openstack Port", map[string]interface{}{
		"uuid": uuid,
	})

	apiResp, err := r.client.GetOpenstackPort(ctx, uuid)
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("Openstack Port with UUID '%s' does not exist or is not accessible.", uuid),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to Import Openstack Port",
			fmt.Sprintf("An error occurred while fetching the Openstack Port: %s", err.Error()),
		)
		return
	}

	var data OpenstackPortResourceModel
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
