package volume

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackVolumeResource{}
var _ resource.ResourceWithImportState = &OpenstackVolumeResource{}

func NewOpenstackVolumeResource() resource.Resource {
	return &OpenstackVolumeResource{}
}

// OpenstackVolumeResource defines the resource implementation.
type OpenstackVolumeResource struct {
	client *Client
}

// OpenstackVolumeResourceModel describes the resource data model.
type OpenstackVolumeResourceModel struct {
	OpenstackVolumeModel
	Offering types.String   `tfsdk:"offering"`
	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackVolumeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_volume"
}

func (r *OpenstackVolumeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Volume resource",

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
			"action": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Action",
			},
			"availability_zone": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Availability zone where this volume is located",
			},
			"availability_zone_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the availability zone",
			},
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Volume ID in the OpenStack backend",
			},
			"bootable": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Indicates if this volume can be used to boot an instance",
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
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description of the resource",
			},
			"device": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of volume as instance device e.g. /dev/vdb.",
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
			"extend_enabled": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Extend enabled",
			},
			"image": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Image that this volume was created from, if any",
			},
			"image_metadata": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Metadata of the image this volume was created from",
			},
			"image_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the image this volume was created from",
			},
			"instance": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Instance that this volume is attached to, if any",
			},
			"instance_marketplace_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the instance marketplace",
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the instance",
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
				MarkdownDescription: "Name of the resource",
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
			"size": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Size in MiB",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
					int64validator.AtMost(2147483647),
				},
			},
			"source_snapshot": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Snapshot that this volume was created from, if any",
			},
			"state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "State",
			},
			"tenant": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Tenant",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the tenant",
			},
			"type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Type of the volume (e.g. SSD, HDD)",
			},
			"type_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the type",
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

func (r *OpenstackVolumeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OpenstackVolumeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackVolumeResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Phase 1: Payload Construction
	// We map the Terraform schema fields to the 'attributes' map required by the Marketplace Order API.
	attributes := OpenstackVolumeCreateAttributes{
		AvailabilityZone: data.AvailabilityZone.ValueStringPointer(),
		Description:      data.Description.ValueStringPointer(),
		Image:            data.Image.ValueStringPointer(),
		Name:             data.Name.ValueStringPointer(),
		Size:             data.Size.ValueInt64Pointer(),
		Type:             data.Type.ValueStringPointer(),
	}

	// Construct the Create Order Request
	payload := OpenstackVolumeCreateRequest{
		Project:    data.Project.ValueStringPointer(),
		Offering:   data.Offering.ValueStringPointer(),
		Attributes: attributes,
	}

	// Phase 2: Submit Order
	orderRes, err := r.client.CreateOpenstackVolumeOrder(ctx, &payload)
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
	apiResp, err := r.client.GetOpenstackVolume(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to Read Resource", err.Error())
		return
	}
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackVolumeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackVolumeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource

	apiResp, err := r.client.GetOpenstackVolume(ctx, data.UUID.ValueString())
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Openstack Volume",
			"An error occurred while reading the Openstack Volume: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackVolumeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackVolumeResourceModel
	var state OpenstackVolumeResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Phase 1: Standard PATCH (Simple fields)
	// We compare the plan (data) with the state (state) to determine which fields changed.
	var patchPayload OpenstackVolumeUpdateRequest
	if !data.Bootable.IsNull() && !data.Bootable.Equal(state.Bootable) {
		patchPayload.Bootable = data.Bootable.ValueBoolPointer()
	}
	if !data.Description.IsNull() && !data.Description.Equal(state.Description) {
		patchPayload.Description = data.Description.ValueStringPointer()
	}
	if !data.Name.IsNull() && !data.Name.Equal(state.Name) {
		patchPayload.Name = data.Name.ValueStringPointer()
	}

	{
		// Execute the PATCH request
		_, err := r.client.UpdateOpenstackVolume(ctx, data.UUID.ValueString(), &patchPayload)
		if err != nil {
			resp.Diagnostics.AddError("Update Failed", err.Error())
			return
		}
	}

	// Phase 2: RPC Actions
	// These actions are triggered when their corresponding specific fields change.
	if !data.Size.Equal(state.Size) {
		// Convert Terraform value to API payload for the specific action
		var req OpenstackVolumeExtendActionRequest
		req.Size = data.Size.ValueInt64Pointer()

		// Execute the Action
		if err := r.client.OpenstackVolumeExtend(ctx, data.UUID.ValueString(), &req); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: extend", err.Error())
			return
		}
	}
	if !data.Type.Equal(state.Type) {
		// Convert Terraform value to API payload for the specific action
		var req OpenstackVolumeRetypeActionRequest
		req.Type = data.Type.ValueStringPointer()

		// Execute the Action
		if err := r.client.OpenstackVolumeRetype(ctx, data.UUID.ValueString(), &req); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: retype", err.Error())
			return
		}
	}

	// Fetch updated state after all changes
	apiResp, err := r.client.GetOpenstackVolume(ctx, data.UUID.ValueString())
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Failed to Read Resource After Update", err.Error())
		return
	}
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackVolumeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackVolumeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Order-based Delete
	payload := map[string]interface{}{}

	// Submit termination order
	orderUUID, err := r.client.TerminateOpenstackVolume(ctx, data.UUID.ValueString(), payload)
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

func (r *OpenstackVolumeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	uuid := req.ID
	if uuid == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Import ID cannot be empty. Please provide the UUID of the Openstack Volume.",
		)
		return
	}

	tflog.Info(ctx, "Importing Openstack Volume", map[string]interface{}{
		"uuid": uuid,
	})

	apiResp, err := r.client.GetOpenstackVolume(ctx, uuid)
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("Openstack Volume with UUID '%s' does not exist or is not accessible.", uuid),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to Import Openstack Volume",
			fmt.Sprintf("An error occurred while fetching the Openstack Volume: %s", err.Error()),
		)
		return
	}

	var data OpenstackVolumeResourceModel
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
