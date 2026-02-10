package security_group

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackSecurityGroupResource{}
var _ resource.ResourceWithImportState = &OpenstackSecurityGroupResource{}

func NewOpenstackSecurityGroupResource() resource.Resource {
	return &OpenstackSecurityGroupResource{}
}

// OpenstackSecurityGroupResource defines the resource implementation.
type OpenstackSecurityGroupResource struct {
	client *OpenstackSecurityGroupClient
}

// OpenstackSecurityGroupResourceModel describes the resource data model.
type OpenstackSecurityGroupResourceModel struct {
	OpenstackSecurityGroupModel
	Timeouts timeouts.Value `tfsdk:"timeouts"`
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
				MarkdownDescription: "Openstack Security Group UUID (used as Terraform ID)",
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
			"error_message": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Error Message",
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
								common.UnknownIfNullModifier{},
							},
							MarkdownDescription: "Description",
						},
						"direction": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								common.UnknownIfNullModifier{},
							},
							MarkdownDescription: "Traffic direction - either 'ingress' (incoming) or 'egress' (outgoing)",
						},
						"ethertype": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								common.UnknownIfNullModifier{},
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
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(65535),
							},
						},
						"protocol": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								common.UnknownIfNullModifier{},
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
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(65535),
							},
						},
						"id": schema.Int64Attribute{
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
								common.UnknownIfNullModifier{},
							},
							MarkdownDescription: "Id",
						},
						"remote_group_name": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Remote Group Name",
						},
						"remote_group_uuid": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Remote Group Uuid",
						},
					},
				},
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
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

func (r *OpenstackSecurityGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.client = &OpenstackSecurityGroupClient{}
	if err := r.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			err.Error(),
		)
		return
	}
}

func (r *OpenstackSecurityGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var data OpenstackSecurityGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := OpenstackSecurityGroupCreateRequest{}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {

		requestBody.Description = data.Description.ValueStringPointer()
	}

	requestBody.Name = data.Name.ValueStringPointer()
	resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.Rules, &requestBody.Rules)...)

	apiResp, err := r.client.Create(ctx, data.Tenant.ValueString(), &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Openstack Security Group",
			"An error occurred while creating the Openstack Security Group: "+err.Error(),
		)
		return
	}
	data.UUID = types.StringPointerValue(apiResp.UUID)
	createTimeout, diags := data.Timeouts.Create(ctx, common.DefaultCreateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackSecurityGroupResponse, error) {
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

func (r *OpenstackSecurityGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackSecurityGroupResourceModel

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
			"Unable to Read Openstack Security Group",
			"An error occurred while reading the Openstack Security Group: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

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

	var apiResp *OpenstackSecurityGroupResponse
	updateTimeout, diags := data.Timeouts.Update(ctx, common.DefaultUpdateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	_ = updateTimeout
	anyChanges := false
	requestBody := OpenstackSecurityGroupUpdateRequest{}
	if !data.Description.IsNull() && !data.Description.IsUnknown() && !data.Description.Equal(state.Description) {
		anyChanges = true

		requestBody.Description = data.Description.ValueStringPointer()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() && !data.Name.Equal(state.Name) {
		anyChanges = true

		requestBody.Name = data.Name.ValueStringPointer()
	}

	if anyChanges {
		var err error
		apiResp, err = r.client.Update(ctx, data.UUID.ValueString(), &requestBody)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Update Openstack Security Group",
				"An error occurred while updating the Openstack Security Group: "+err.Error(),
			)
			return
		}
		// Wait for the resource to return to OK state
		newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackSecurityGroupResponse, error) {
			return r.client.Get(ctx, data.UUID.ValueString())
		}, updateTimeout)
		if err != nil {
			resp.Diagnostics.AddError("Wait for update failed", err.Error())
			return
		}
		apiResp = newResp
	}
	if !data.Rules.Equal(state.Rules) {
		// Convert Terraform value to API payload for the specific action
		var req OpenstackSecurityGroupSetRulesActionRequest
		resp.Diagnostics.Append(common.PopulateSliceField(ctx, data.Rules, &req.Rules)...)

		// Execute the Action
		if err := r.client.SetRules(ctx, data.UUID.ValueString(), &req); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: set_rules", err.Error())
			return
		}
		// Wait for the resource to return to OK state
		_, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackSecurityGroupResponse, error) {
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

func (r *OpenstackSecurityGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var data OpenstackSecurityGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Openstack Security Group",
			"An error occurred while deleting the Openstack Security Group: "+err.Error(),
		)
		return
	}
	deleteTimeout, diags := data.Timeouts.Delete(ctx, common.DefaultDeleteTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err = common.WaitForDeletion(ctx, func(ctx context.Context) (*OpenstackSecurityGroupResponse, error) {
		return r.client.Get(ctx, data.UUID.ValueString())
	}, deleteTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource deletion", err.Error())
		return
	}
}

func (r *OpenstackSecurityGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	uuid := req.ID
	if uuid == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Import ID cannot be empty. Please provide the UUID of the Openstack Security Group.",
		)
		return
	}

	tflog.Info(ctx, "Importing Openstack Security Group", map[string]interface{}{
		"uuid": uuid,
	})

	apiResp, err := r.client.Get(ctx, uuid)
	if err != nil {
		if IsNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("Openstack Security Group with UUID '%s' does not exist or is not accessible.", uuid),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to Import Openstack Security Group",
			fmt.Sprintf("An error occurred while fetching the Openstack Security Group: %s", err.Error()),
		)
		return
	}

	var data OpenstackSecurityGroupResourceModel
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
