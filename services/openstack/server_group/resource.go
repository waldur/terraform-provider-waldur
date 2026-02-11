package server_group

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackServerGroupResource{}
var _ resource.ResourceWithImportState = &OpenstackServerGroupResource{}

func NewOpenstackServerGroupResource() resource.Resource {
	return &OpenstackServerGroupResource{}
}

// OpenstackServerGroupResource defines the resource implementation.
type OpenstackServerGroupResource struct {
	client *OpenstackServerGroupClient
}

// OpenstackServerGroupResourceModel describes the resource data model.
type OpenstackServerGroupResourceModel struct {
	OpenstackServerGroupModel
	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackServerGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_server_group"
}

func (r *OpenstackServerGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Server Group resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Openstack Server Group UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"backend_id": schema.StringAttribute{

				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Backend Id",
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
			"display_name": schema.StringAttribute{

				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Display Name",
			},
			"error_message": schema.StringAttribute{

				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Error Message",
			},
			"instances": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"backend_id": schema.StringAttribute{

							Optional: true,
							Computed: true,

							PlanModifiers: []planmodifier.String{

								stringplanmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "Instance ID in the OpenStack backend",
						},
						"name": schema.StringAttribute{

							Optional: true,
							Computed: true,

							PlanModifiers: []planmodifier.String{

								stringplanmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "Name",
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

				Computed: true,

				PlanModifiers: []planmodifier.List{

					listplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Instances",
			},
			"marketplace_resource_uuid": schema.StringAttribute{

				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Marketplace Resource Uuid",
			},
			"name": schema.StringAttribute{

				Required: true,

				MarkdownDescription: "Name",
			},
			"policy": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Server group policy determining the rules for scheduling servers in this group",
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

				MarkdownDescription: "Tenant",
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

func (r *OpenstackServerGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.client = &OpenstackServerGroupClient{}
	if err := r.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			err.Error(),
		)
		return
	}
}

func (r *OpenstackServerGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var data OpenstackServerGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := OpenstackServerGroupCreateRequest{}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {

		requestBody.Description = data.Description.ValueStringPointer()
	}

	requestBody.Name = data.Name.ValueStringPointer()
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {

		requestBody.Policy = data.Policy.ValueStringPointer()
	}

	apiResp, err := r.client.Create(ctx, data.Tenant.ValueString(), &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Openstack Server Group",
			"An error occurred while creating the Openstack Server Group: "+err.Error(),
		)
		return
	}
	data.UUID = types.StringPointerValue(apiResp.UUID)
	createTimeout, diags := data.Timeouts.Create(ctx, common.DefaultCreateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackServerGroupResponse, error) {
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

func (r *OpenstackServerGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackServerGroupResourceModel

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
			"Unable to Read Openstack Server Group",
			"An error occurred while reading the Openstack Server Group: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

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

	var apiResp *OpenstackServerGroupResponse
	updateTimeout, diags := data.Timeouts.Update(ctx, common.DefaultUpdateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	_ = updateTimeout
	anyChanges := false
	requestBody := OpenstackServerGroupUpdateRequest{}
	if !data.Description.IsNull() && !data.Description.IsUnknown() && !data.Description.Equal(state.Description) {
		anyChanges = true

		requestBody.Description = data.Description.ValueStringPointer()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() && !data.Name.Equal(state.Name) {
		anyChanges = true

		requestBody.Name = data.Name.ValueStringPointer()
	}
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() && !data.Policy.Equal(state.Policy) {
		anyChanges = true

		requestBody.Policy = data.Policy.ValueStringPointer()
	}

	if anyChanges {
		var err error
		apiResp, err = r.client.Update(ctx, data.UUID.ValueString(), &requestBody)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Update Openstack Server Group",
				"An error occurred while updating the Openstack Server Group: "+err.Error(),
			)
			return
		}
		// Wait for the resource to return to OK state
		newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackServerGroupResponse, error) {
			return r.client.Get(ctx, data.UUID.ValueString())
		}, updateTimeout)
		if err != nil {
			resp.Diagnostics.AddError("Wait for update failed", err.Error())
			return
		}
		apiResp = newResp
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

func (r *OpenstackServerGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var data OpenstackServerGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Openstack Server Group",
			"An error occurred while deleting the Openstack Server Group: "+err.Error(),
		)
		return
	}
	deleteTimeout, diags := data.Timeouts.Delete(ctx, common.DefaultDeleteTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err = common.WaitForDeletion(ctx, func(ctx context.Context) (*OpenstackServerGroupResponse, error) {
		return r.client.Get(ctx, data.UUID.ValueString())
	}, deleteTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource deletion", err.Error())
		return
	}
}

func (r *OpenstackServerGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	uuid := req.ID
	if uuid == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Import ID cannot be empty. Please provide the UUID of the Openstack Server Group.",
		)
		return
	}

	tflog.Info(ctx, "Importing Openstack Server Group", map[string]interface{}{
		"uuid": uuid,
	})

	apiResp, err := r.client.Get(ctx, uuid)
	if err != nil {
		if IsNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("Openstack Server Group with UUID '%s' does not exist or is not accessible.", uuid),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to Import Openstack Server Group",
			fmt.Sprintf("An error occurred while fetching the Openstack Server Group: %s", err.Error()),
		)
		return
	}

	var data OpenstackServerGroupResourceModel
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
