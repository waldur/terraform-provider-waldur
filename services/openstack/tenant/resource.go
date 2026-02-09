package tenant

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackTenantResource{}
var _ resource.ResourceWithImportState = &OpenstackTenantResource{}

func NewOpenstackTenantResource() resource.Resource {
	return &OpenstackTenantResource{}
}

// OpenstackTenantResource defines the resource implementation.
type OpenstackTenantResource struct {
	client *OpenstackTenantClient
}

// OpenstackTenantResourceModel describes the resource data model.
type OpenstackTenantResourceModel struct {
	OpenstackTenantModel
	Limits                      types.Map      `tfsdk:"limits"`
	Offering                    types.String   `tfsdk:"offering"`
	Plan                        types.String   `tfsdk:"plan"`
	SecurityGroups              types.Set      `tfsdk:"security_groups"`
	SkipConnectionExtnet        types.Bool     `tfsdk:"skip_connection_extnet"`
	SkipCreationOfDefaultSubnet types.Bool     `tfsdk:"skip_creation_of_default_subnet"`
	SubnetCidr                  types.String   `tfsdk:"subnet_cidr"`
	Timeouts                    timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackTenantResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_tenant"
}

func (r *OpenstackTenantResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Tenant resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Openstack Tenant UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"availability_zone": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Optional availability group. Will be used for all instances provisioned in this tenant",
			},
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "ID of tenant in the OpenStack backend",
			},
			"customer": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Customer",
			},
			"default_volume_type_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Volume type name to use when creating volumes.",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Description of the Openstack Tenant",
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
			"external_network_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "ID of external network connected to OpenStack tenant",
			},
			"internal_network_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "ID of internal network in OpenStack tenant",
			},
			"limits": schema.MapAttribute{
				ElementType: types.Float64Type,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.RequiresReplace(),
					mapplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Resource limits",
			},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the marketplace resource",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the Openstack Tenant",
			},
			"offering": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Offering URL",
			},
			"plan": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Plan URL",
			},
			"project": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Project URL",
			},
			"quotas": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"limit": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Limit",
						},
						"name": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Name of the Openstack Tenant",
						},
						"usage": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Usage",
						},
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Quotas",
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
						"description": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Description of the Openstack Tenant",
						},
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Name of the Openstack Tenant",
						},
						"rules": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"cidr": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "CIDR notation for the source/destination network address range",
									},
									"description": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Description of the Openstack Tenant",
									},
									"direction": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Traffic direction - either 'ingress' (incoming) or 'egress' (outgoing)",
									},
									"ethertype": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "IP protocol version - either 'IPv4' or 'IPv6'",
									},
									"from_port": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "Starting port number in the range (1-65535)",
										Validators: []validator.Int64{
											int64validator.AtLeast(-2147483648),
											int64validator.AtMost(65535),
										},
									},
									"protocol": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "The network protocol (TCP, UDP, ICMP, or empty for any protocol)",
									},
									"remote_group": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Remote security group that this rule references, if any",
									},
									"to_port": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "Ending port number in the range (1-65535)",
										Validators: []validator.Int64{
											int64validator.AtLeast(-2147483648),
											int64validator.AtMost(65535),
										},
									},
								},
							},
							Optional:            true,
							MarkdownDescription: "Rules",
						},
					},
				},
				Optional:            true,
				MarkdownDescription: "Security groups",
			},
			"skip_connection_extnet": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Skip connection extnet",
			},
			"skip_creation_of_default_router": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Skip creation of default router",
			},
			"skip_creation_of_default_subnet": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Skip creation of default subnet",
			},
			"state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "State",
			},
			"subnet_cidr": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Subnet cidr",
			},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Url",
			},
			"user_password": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Password of the tenant user",
				Sensitive:           true,
			},
			"user_username": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Username of the tenant user",
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

func (r *OpenstackTenantResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.client = &OpenstackTenantClient{}
	if err := r.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			err.Error(),
		)
		return
	}
}

// resolveUnknownAttributes ensures that fields not returned by the Waldur GET API
// are set to explicit null values instead of remaining "Unknown".
func (r *OpenstackTenantResource) resolveUnknownAttributes(data *OpenstackTenantResourceModel) {
	// Iterate over all model fields to handle Unknown values
	if data.AvailabilityZone.IsUnknown() {
		data.AvailabilityZone = types.StringNull()
	}
	if data.BackendId.IsUnknown() {
		data.BackendId = types.StringNull()
	}
	if data.Customer.IsUnknown() {
		data.Customer = types.StringNull()
	}
	if data.DefaultVolumeTypeName.IsUnknown() {
		data.DefaultVolumeTypeName = types.StringNull()
	}
	if data.Description.IsUnknown() {
		data.Description = types.StringNull()
	}
	if data.ErrorMessage.IsUnknown() {
		data.ErrorMessage = types.StringNull()
	}
	if data.ErrorTraceback.IsUnknown() {
		data.ErrorTraceback = types.StringNull()
	}
	if data.ExternalNetworkId.IsUnknown() {
		data.ExternalNetworkId = types.StringNull()
	}
	if data.InternalNetworkId.IsUnknown() {
		data.InternalNetworkId = types.StringNull()
	}
	if data.Limits.IsUnknown() {
		data.Limits = types.MapNull(types.Float64Type)
	}
	if data.MarketplaceResourceUuid.IsUnknown() {
		data.MarketplaceResourceUuid = types.StringNull()
	}
	if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if data.Offering.IsUnknown() {
		data.Offering = types.StringNull()
	}
	if data.Plan.IsUnknown() {
		data.Plan = types.StringNull()
	}
	if data.Project.IsUnknown() {
		data.Project = types.StringNull()
	}
	if data.Quotas.IsUnknown() {
		data.Quotas = types.ListNull(QuotaType())
	}
	if data.ResourceType.IsUnknown() {
		data.ResourceType = types.StringNull()
	}
	if data.SecurityGroups.IsUnknown() {
		data.SecurityGroups = types.SetNull(OpenStackTenantSecurityGroupRequestType())
	}
	if data.SkipConnectionExtnet.IsUnknown() {
		data.SkipConnectionExtnet = types.BoolNull()
	}
	if data.SkipCreationOfDefaultRouter.IsUnknown() {
		data.SkipCreationOfDefaultRouter = types.BoolNull()
	}
	if data.SkipCreationOfDefaultSubnet.IsUnknown() {
		data.SkipCreationOfDefaultSubnet = types.BoolNull()
	}
	if data.State.IsUnknown() {
		data.State = types.StringNull()
	}
	if data.SubnetCidr.IsUnknown() {
		data.SubnetCidr = types.StringNull()
	}
	if data.Url.IsUnknown() {
		data.Url = types.StringNull()
	}
	if data.UserPassword.IsUnknown() {
		data.UserPassword = types.StringNull()
	}
	if data.UserUsername.IsUnknown() {
		data.UserUsername = types.StringNull()
	}
}

func (r *OpenstackTenantResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackTenantResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Phase 1: Payload Construction
	// We map the Terraform schema fields to the 'attributes' map required by the Marketplace Order API.
	attributes := OpenstackTenantCreateAttributes{}
	if !data.AvailabilityZone.IsNull() && !data.AvailabilityZone.IsUnknown() {
		attributes.AvailabilityZone = data.AvailabilityZone.ValueStringPointer()
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		attributes.Description = data.Description.ValueStringPointer()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		attributes.Name = data.Name.ValueStringPointer()
	}
	if !data.SkipConnectionExtnet.IsNull() && !data.SkipConnectionExtnet.IsUnknown() {
		attributes.SkipConnectionExtnet = data.SkipConnectionExtnet.ValueBoolPointer()
	}
	if !data.SkipCreationOfDefaultRouter.IsNull() && !data.SkipCreationOfDefaultRouter.IsUnknown() {
		attributes.SkipCreationOfDefaultRouter = data.SkipCreationOfDefaultRouter.ValueBoolPointer()
	}
	if !data.SkipCreationOfDefaultSubnet.IsNull() && !data.SkipCreationOfDefaultSubnet.IsUnknown() {
		attributes.SkipCreationOfDefaultSubnet = data.SkipCreationOfDefaultSubnet.ValueBoolPointer()
	}
	if !data.SubnetCidr.IsNull() && !data.SubnetCidr.IsUnknown() {
		attributes.SubnetCidr = data.SubnetCidr.ValueStringPointer()
	}
	resp.Diagnostics.Append(common.PopulateOptionalSetField(ctx, data.SecurityGroups, &attributes.SecurityGroups)...)

	// Construct the Create Order Request
	payload := OpenstackTenantCreateRequest{
		Project:    data.Project.ValueStringPointer(),
		Offering:   data.Offering.ValueStringPointer(),
		Attributes: attributes,
	}

	if !data.Plan.IsNull() && !data.Plan.IsUnknown() {
		payload.Plan = data.Plan.ValueStringPointer()
	}

	if !data.Limits.IsNull() && !data.Limits.IsUnknown() {
		limits := make(map[string]float64)
		diags := data.Limits.ElementsAs(ctx, &limits, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		payload.Limits = limits
	}

	// Phase 2: Submit Order
	orderRes, err := r.client.CreateOrder(ctx, &payload)
	if err != nil {
		resp.Diagnostics.AddError("Order Submission Failed", err.Error())
		return
	}

	// Phase 3: Poll for Completion
	// We use the 'time' package to handle the timeout specified in the TF config or default to global default.
	timeout, diags := data.Timeouts.Create(ctx, common.DefaultCreateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Wait for the order to reach a terminal state (done/erred)
	finalOrder, err := common.WaitForOrder(ctx, r.client.Client, *orderRes.Uuid, timeout)
	if err != nil {
		resp.Diagnostics.AddError("Order Failed", err.Error())
		return
	}

	// Resolve the created Resource UUID from the completed order
	if uuid := common.ResolveResourceUUID(finalOrder); uuid != "" {
		data.UUID = types.StringValue(uuid)
	} else {
		resp.Diagnostics.AddError("Resource UUID Missing", "Order completed but resource UUID is missing")
		return
	}

	// Fetch final resource state to ensure Terraform state matches reality
	apiResp, err := r.client.Get(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to Read Resource", err.Error())
		return
	}
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	// Resolve unknown attributes to explicit null values
	r.resolveUnknownAttributes(&data)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackTenantResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackTenantResourceModel

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
			"Unable to Read Openstack Tenant",
			"An error occurred while reading the Openstack Tenant: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Resolve unknown attributes to explicit null values
	r.resolveUnknownAttributes(&data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackTenantResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackTenantResourceModel
	var state OpenstackTenantResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Phase 1: Standard PATCH (Simple fields)
	// We compare the plan (data) with the state (state) to determine which fields changed.
	anyChanges := false
	var patchPayload OpenstackTenantUpdateRequest
	if !data.AvailabilityZone.IsNull() && !data.AvailabilityZone.Equal(state.AvailabilityZone) {
		anyChanges = true
		patchPayload.AvailabilityZone = data.AvailabilityZone.ValueStringPointer()
	}
	if !data.DefaultVolumeTypeName.IsNull() && !data.DefaultVolumeTypeName.Equal(state.DefaultVolumeTypeName) {
		anyChanges = true
		patchPayload.DefaultVolumeTypeName = data.DefaultVolumeTypeName.ValueStringPointer()
	}
	if !data.Description.IsNull() && !data.Description.Equal(state.Description) {
		anyChanges = true
		patchPayload.Description = data.Description.ValueStringPointer()
	}
	if !data.Name.IsNull() && !data.Name.Equal(state.Name) {
		anyChanges = true
		patchPayload.Name = data.Name.ValueStringPointer()
	}
	if !data.SkipCreationOfDefaultRouter.IsNull() && !data.SkipCreationOfDefaultRouter.Equal(state.SkipCreationOfDefaultRouter) {
		anyChanges = true
		patchPayload.SkipCreationOfDefaultRouter = data.SkipCreationOfDefaultRouter.ValueBoolPointer()
	}
	if !data.SkipCreationOfDefaultSubnet.IsNull() && !data.SkipCreationOfDefaultSubnet.Equal(state.SkipCreationOfDefaultSubnet) {
		anyChanges = true
		patchPayload.SkipCreationOfDefaultSubnet = data.SkipCreationOfDefaultSubnet.ValueBoolPointer()
	}

	if anyChanges {
		// Execute the PATCH request
		_, err := r.client.Update(ctx, data.UUID.ValueString(), &patchPayload)
		if err != nil {
			resp.Diagnostics.AddError("Update Failed", err.Error())
			return
		}
	}

	updateTimeout, diags := data.Timeouts.Update(ctx, common.DefaultUpdateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Phase 2: RPC Actions
	// These actions are triggered when their corresponding specific fields change.
	if !data.SecurityGroups.Equal(state.SecurityGroups) {
		// Convert Terraform value to API payload for the specific action
		var req OpenstackTenantPushSecurityGroupsActionRequest
		resp.Diagnostics.Append(common.PopulateSetField(ctx, data.SecurityGroups, &req.SecurityGroups)...)

		// Execute the Action
		if err := r.client.PushSecurityGroups(ctx, data.UUID.ValueString(), &req); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: push_security_groups", err.Error())
			return
		}

		// Wait for the resource to return to OK state
		apiResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackTenantResponse, error) {
			return r.client.Get(ctx, data.UUID.ValueString())
		}, updateTimeout)
		if err != nil {
			resp.Diagnostics.AddError("Wait for RPC action failed", err.Error())
			return
		}
		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
		state = data // Update local state to avoid repeated action calls if multiple fields changed (though actions are usually 1-to-1)
	}

	// Fetch updated state after all changes
	apiResp, err := r.client.Get(ctx, data.UUID.ValueString())
	if err != nil {
		if IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Failed to Read Resource After Update", err.Error())
		return
	}
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	// Resolve unknown attributes to explicit null values
	r.resolveUnknownAttributes(&data)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackTenantResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackTenantResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Order-based Delete

	payload := map[string]interface{}{}

	// Submit termination order
	resourceID := data.UUID.ValueString()
	if !data.MarketplaceResourceUuid.IsNull() && !data.MarketplaceResourceUuid.IsUnknown() {
		resourceID = data.MarketplaceResourceUuid.ValueString()
	}
	orderUUID, err := r.client.Terminate(ctx, resourceID, payload)
	if err != nil {
		resp.Diagnostics.AddError("Termination Failed", err.Error())
		return
	}

	// Wait for deletion if order UUID is returned
	if orderUUID != "" {
		timeout, diags := data.Timeouts.Delete(ctx, common.DefaultDeleteTimeout)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		_, err := common.WaitForOrder(ctx, r.client.Client, orderUUID, timeout)
		if err != nil {
			resp.Diagnostics.AddError("Termination Order Failed", err.Error())
			return
		}
	}
}

func (r *OpenstackTenantResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	uuid := req.ID
	if uuid == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Import ID cannot be empty. Please provide the UUID of the Openstack Tenant.",
		)
		return
	}

	tflog.Info(ctx, "Importing Openstack Tenant", map[string]interface{}{
		"uuid": uuid,
	})

	apiResp, err := r.client.Get(ctx, uuid)
	if err != nil {
		if IsNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("Openstack Tenant with UUID '%s' does not exist or is not accessible.", uuid),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to Import Openstack Tenant",
			fmt.Sprintf("An error occurred while fetching the Openstack Tenant: %s", err.Error()),
		)
		return
	}

	var data OpenstackTenantResourceModel
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
