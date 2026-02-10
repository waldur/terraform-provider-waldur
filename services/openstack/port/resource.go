package port

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

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
	client *OpenstackPortClient
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
				MarkdownDescription: "Openstack Port UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
							MarkdownDescription: "Ip Address",
						},
						"mac_address": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Mac Address",
						},
					},
				},
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Allowed Address Pairs",
			},
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Port ID in OpenStack",
			},
			"customer": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Customer",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Description",
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
				MarkdownDescription: "Error Message",
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
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Fixed Ips",
			},
			"floating_ips": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Floating Ips",
			},
			"mac_address": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "MAC address of the port",
			},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Marketplace Resource Uuid",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name",
			},
			"network": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Network to which this port belongs",
			},
			"network_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Network Name",
			},
			"network_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Network Uuid",
			},
			"port_security_enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "If True, security groups and rules will be applied to this port",
			},
			"project": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Project",
			},
			"resource_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Resource Type",
			},
			"security_groups": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Name",
						},
						"url": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Url",
						},
						"uuid": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Uuid",
						},
					},
				},
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Security Groups",
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
				MarkdownDescription: "Tenant Name",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Tenant Uuid",
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

	r.client = &OpenstackPortClient{}
	if err := r.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			err.Error(),
		)
		return
	}
}

func (r *OpenstackPortResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var data OpenstackPortResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := OpenstackPortCreateRequest{}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {

		requestBody.Description = data.Description.ValueStringPointer()
	}
	if !data.MacAddress.IsNull() && !data.MacAddress.IsUnknown() {

		requestBody.MacAddress = data.MacAddress.ValueStringPointer()
	}

	requestBody.Name = data.Name.ValueStringPointer()
	if !data.Network.IsNull() && !data.Network.IsUnknown() {

		requestBody.Network = data.Network.ValueStringPointer()
	}
	if !data.PortSecurityEnabled.IsNull() && !data.PortSecurityEnabled.IsUnknown() {

		requestBody.PortSecurityEnabled = data.PortSecurityEnabled.ValueBoolPointer()
	}
	if !data.TargetTenant.IsNull() && !data.TargetTenant.IsUnknown() {

		requestBody.TargetTenant = data.TargetTenant.ValueStringPointer()
	}
	resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.AllowedAddressPairs, &requestBody.AllowedAddressPairs)...)
	resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.FixedIps, &requestBody.FixedIps)...)
	resp.Diagnostics.Append(common.PopulateOptionalSetField(ctx, data.SecurityGroups, &requestBody.SecurityGroups)...)

	apiResp, err := r.client.Create(ctx, &requestBody)
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
		return r.client.Get(ctx, data.UUID.ValueString())
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

	apiResp, err := r.client.Get(ctx, data.UUID.ValueString())
	if err != nil {
		if IsNotFoundError(err) {
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

	var apiResp *OpenstackPortResponse
	updateTimeout, diags := data.Timeouts.Update(ctx, common.DefaultUpdateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	_ = updateTimeout
	anyChanges := false
	requestBody := OpenstackPortUpdateRequest{}
	if !data.Description.IsNull() && !data.Description.IsUnknown() && !data.Description.Equal(state.Description) {
		anyChanges = true

		requestBody.Description = data.Description.ValueStringPointer()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() && !data.Name.Equal(state.Name) {
		anyChanges = true

		requestBody.Name = data.Name.ValueStringPointer()
	}
	if !data.TargetTenant.IsNull() && !data.TargetTenant.IsUnknown() && !data.TargetTenant.Equal(state.TargetTenant) {
		anyChanges = true

		requestBody.TargetTenant = data.TargetTenant.ValueStringPointer()
	}

	resp.Diagnostics.Append(common.PopulateOptionalSetField(ctx, data.SecurityGroups, &requestBody.SecurityGroups)...)

	if anyChanges {
		var err error
		apiResp, err = r.client.Update(ctx, data.UUID.ValueString(), &requestBody)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Update Openstack Port",
				"An error occurred while updating the Openstack Port: "+err.Error(),
			)
			return
		}
		// Wait for the resource to return to OK state
		newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackPortResponse, error) {
			return r.client.Get(ctx, data.UUID.ValueString())
		}, updateTimeout)
		if err != nil {
			resp.Diagnostics.AddError("Wait for update failed", err.Error())
			return
		}
		apiResp = newResp
	}
	if !data.SecurityGroups.Equal(state.SecurityGroups) {
		// Convert Terraform value to API payload for the specific action
		var req OpenstackPortUpdateSecurityGroupsActionRequest
		resp.Diagnostics.Append(common.PopulateSetField(ctx, data.SecurityGroups, &req.SecurityGroups)...)

		// Execute the Action
		if err := r.client.UpdateSecurityGroups(ctx, data.UUID.ValueString(), &req); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: update_security_groups", err.Error())
			return
		}
		// Wait for the resource to return to OK state
		_, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackPortResponse, error) {
			return r.client.Get(ctx, data.UUID.ValueString())
		}, updateTimeout)
		if err != nil {
			resp.Diagnostics.AddError("Wait for RPC action failed", err.Error())
			return
		}
		state = data
	}

	newResp, err := r.client.Get(ctx, data.UUID.ValueString())
	if err != nil {
		if IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Failed to Read Resource After Update", err.Error())
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

	err := r.client.Delete(ctx, data.UUID.ValueString())
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
		return r.client.Get(ctx, data.UUID.ValueString())
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

	apiResp, err := r.client.Get(ctx, uuid)
	if err != nil {
		if IsNotFoundError(err) {
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
