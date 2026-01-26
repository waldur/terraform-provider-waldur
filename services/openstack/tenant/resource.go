package tenant

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
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
	client *Client
}

// OpenstackTenantResourceModel describes the resource data model.
type OpenstackTenantResourceModel struct {
	OpenstackTenantModel
	Offering                    types.String   `tfsdk:"offering"`
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
			"access_url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Access url",
			},
			"availability_zone": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Optional availability group. Will be used for all instances provisioned in this tenant",
			},
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "ID of tenant in the OpenStack backend",
			},
			"created": schema.StringAttribute{
				CustomType: timetypes.RFC3339Type{},
				Computed:   true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Created",
			},
			"customer": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Customer",
			},
			"customer_abbreviation": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Customer abbreviation",
			},
			"customer_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the customer",
			},
			"customer_native_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the customer native",
			},
			"customer_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the customer",
			},
			"default_volume_type_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Volume type name to use when creating volumes.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
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
			"is_limit_based": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Is limit based",
			},
			"is_usage_based": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Is usage based",
			},
			"marketplace_category_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the marketplace category",
			},
			"marketplace_category_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the marketplace category",
			},
			"marketplace_offering_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the marketplace offering",
			},
			"marketplace_offering_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the marketplace offering",
			},
			"marketplace_plan_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the marketplace plan",
			},
			"marketplace_resource_state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Marketplace resource state",
			},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the marketplace resource",
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
				Optional:            true,
				MarkdownDescription: "Name of the Openstack Tenant",
			},
			"offering": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Offering URL",
			},
			"project": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Project",
			},
			"project_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the project",
			},
			"project_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the project",
			},
			"quotas": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"limit": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Limit",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the Openstack Tenant",
						},
						"usage": schema.Int64Attribute{
							Optional:            true,
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
			"service_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the service",
			},
			"service_settings": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Service settings",
			},
			"service_settings_error_message": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Service settings error message",
			},
			"service_settings_state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Service settings state",
			},
			"service_settings_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the service settings",
			},
			"skip_connection_extnet": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Skip connection extnet",
			},
			"skip_creation_of_default_router": schema.BoolAttribute{
				Optional:            true,
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

	r.client = &Client{}
	if err := r.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			err.Error(),
		)
		return
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
	attributes := OpenstackTenantCreateAttributes{
		AvailabilityZone:            data.AvailabilityZone.ValueStringPointer(),
		Description:                 data.Description.ValueStringPointer(),
		Name:                        data.Name.ValueStringPointer(),
		SkipConnectionExtnet:        data.SkipConnectionExtnet.ValueBoolPointer(),
		SkipCreationOfDefaultRouter: data.SkipCreationOfDefaultRouter.ValueBoolPointer(),
		SkipCreationOfDefaultSubnet: data.SkipCreationOfDefaultSubnet.ValueBoolPointer(),
		SubnetCidr:                  data.SubnetCidr.ValueStringPointer(),
	}
	{
		var items []common.OpenStackTenantSecurityGroupRequest
		diags := data.SecurityGroups.ElementsAs(ctx, &items, false)
		resp.Diagnostics.Append(diags...)
		if !diags.HasError() {
			if !data.SecurityGroups.IsNull() && !data.SecurityGroups.IsUnknown() {
				attributes.SecurityGroups = &items
			}
		}
	}

	// Construct the Create Order Request
	payload := OpenstackTenantCreateRequest{
		Project:    data.Project.ValueStringPointer(),
		Offering:   data.Offering.ValueStringPointer(),
		Attributes: attributes,
	}

	// Phase 2: Submit Order
	orderRes, err := r.client.CreateOpenstackTenantOrder(ctx, &payload)
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
	apiResp, err := r.client.GetOpenstackTenant(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to Read Resource", err.Error())
		return
	}
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

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

	apiResp, err := r.client.GetOpenstackTenant(ctx, data.UUID.ValueString())
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
	var patchPayload OpenstackTenantUpdateRequest
	if !data.AvailabilityZone.IsNull() && !data.AvailabilityZone.Equal(state.AvailabilityZone) {
		patchPayload.AvailabilityZone = data.AvailabilityZone.ValueStringPointer()
	}
	if !data.DefaultVolumeTypeName.IsNull() && !data.DefaultVolumeTypeName.Equal(state.DefaultVolumeTypeName) {
		patchPayload.DefaultVolumeTypeName = data.DefaultVolumeTypeName.ValueStringPointer()
	}
	if !data.Description.IsNull() && !data.Description.Equal(state.Description) {
		patchPayload.Description = data.Description.ValueStringPointer()
	}
	if !data.Name.IsNull() && !data.Name.Equal(state.Name) {
		patchPayload.Name = data.Name.ValueStringPointer()
	}
	if !data.SkipCreationOfDefaultRouter.IsNull() && !data.SkipCreationOfDefaultRouter.Equal(state.SkipCreationOfDefaultRouter) {
		patchPayload.SkipCreationOfDefaultRouter = data.SkipCreationOfDefaultRouter.ValueBoolPointer()
	}
	if !data.SkipCreationOfDefaultSubnet.IsNull() && !data.SkipCreationOfDefaultSubnet.Equal(state.SkipCreationOfDefaultSubnet) {
		patchPayload.SkipCreationOfDefaultSubnet = data.SkipCreationOfDefaultSubnet.ValueBoolPointer()
	}

	{
		// Execute the PATCH request
		_, err := r.client.UpdateOpenstackTenant(ctx, data.UUID.ValueString(), &patchPayload)
		if err != nil {
			resp.Diagnostics.AddError("Update Failed", err.Error())
			return
		}
	}

	// Phase 2: RPC Actions
	// These actions are triggered when their corresponding specific fields change.
	if !data.SecurityGroups.Equal(state.SecurityGroups) {
		// Convert Terraform value to API payload for the specific action
		var req OpenstackTenantPushSecurityGroupsActionRequest
		resp.Diagnostics.Append(common.PopulateSetField(ctx, data.SecurityGroups, &req.SecurityGroups)...)

		// Execute the Action
		if err := r.client.OpenstackTenantPushSecurityGroups(ctx, data.UUID.ValueString(), &req); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: push_security_groups", err.Error())
			return
		}
	}

	// Fetch updated state after all changes
	apiResp, err := r.client.GetOpenstackTenant(ctx, data.UUID.ValueString())
	if err != nil {
		if IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Failed to Read Resource After Update", err.Error())
		return
	}
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

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
	orderUUID, err := r.client.TerminateOpenstackTenant(ctx, data.UUID.ValueString(), payload)
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

	apiResp, err := r.client.GetOpenstackTenant(ctx, uuid)
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
