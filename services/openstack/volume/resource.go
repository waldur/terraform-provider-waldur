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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

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
	client *OpenstackVolumeClient
}

// OpenstackVolumeResourceModel describes the resource data model.
type OpenstackVolumeResourceModel struct {
	OpenstackVolumeModel
	EndDate   timetypes.RFC3339 `tfsdk:"end_date"`
	Limits    types.Map         `tfsdk:"limits"`
	Offering  types.String      `tfsdk:"offering"`
	Plan      types.String      `tfsdk:"plan"`
	StartDate timetypes.RFC3339 `tfsdk:"start_date"`
	Timeouts  timeouts.Value    `tfsdk:"timeouts"`
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
				MarkdownDescription: "Openstack Volume UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"action": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Action"},
			"availability_zone": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Availability zone where this volume is located"},
			"availability_zone_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Availability Zone Name"},
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Volume ID in the OpenStack backend"},
			"bootable": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{

					boolplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Indicates if this volume can be used to boot an instance"},
			"customer": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Customer"},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Description"},
			"device": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Name of volume as instance device e.g. /dev/vdb."},
			"end_date": schema.StringAttribute{
				CustomType: timetypes.RFC3339Type{},
				Optional:   true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.RequiresReplace(),
				}, MarkdownDescription: "Order end date"},
			"error_message": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Error Message"},
			"extend_enabled": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{

					boolplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Extend Enabled"},
			"image": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Image that this volume was created from, if any"},
			"image_metadata": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Metadata of the image this volume was created from"},
			"image_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Name of the image this volume was created from"},
			"instance": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Instance that this volume is attached to, if any"},
			"instance_marketplace_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Instance Marketplace Uuid"},
			"instance_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Instance Name"},
			"limits": schema.MapAttribute{
				ElementType: types.Float64Type,
				Optional:    true,
				PlanModifiers: []planmodifier.Map{

					mapplanmodifier.RequiresReplace(),
				}, MarkdownDescription: "Resource limits"},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Marketplace Resource Uuid"},
			"name": schema.StringAttribute{
				Required: true, MarkdownDescription: "Name"},
			"offering": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.RequiresReplace(),
				}, MarkdownDescription: "Offering URL"},
			"plan": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.RequiresReplace(),
				}, MarkdownDescription: "Plan URL"},
			"project": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.RequiresReplace(),
				}, MarkdownDescription: "Project URL"},
			"resource_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Resource Type"},
			"runtime_state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Runtime State"},
			"size": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{

					int64planmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Size in MiB",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
					int64validator.AtMost(2147483647),
				}},
			"source_snapshot": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Snapshot that this volume was created from, if any"},
			"start_date": schema.StringAttribute{
				CustomType: timetypes.RFC3339Type{},
				Optional:   true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.RequiresReplace(),
				}, MarkdownDescription: "Order start date"},
			"state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "State"},
			"tenant": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Tenant"},
			"tenant_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Tenant Uuid"},
			"type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Type of the volume (e.g. SSD, HDD)"},
			"type_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Type Name"},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Url"},
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

	r.client = &OpenstackVolumeClient{}
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
func (r *OpenstackVolumeResource) resolveUnknownAttributes(data *OpenstackVolumeResourceModel) {
	// Iterate over all model fields to handle Unknown values
	if data.Action.IsUnknown() {
		data.Action = types.StringNull()
	}
	if data.AvailabilityZone.IsUnknown() {
		data.AvailabilityZone = types.StringNull()
	}
	if data.AvailabilityZoneName.IsUnknown() {
		data.AvailabilityZoneName = types.StringNull()
	}
	if data.BackendId.IsUnknown() {
		data.BackendId = types.StringNull()
	}
	if data.Bootable.IsUnknown() {
		data.Bootable = types.BoolNull()
	}
	if data.Customer.IsUnknown() {
		data.Customer = types.StringNull()
	}
	if data.Description.IsUnknown() {
		data.Description = types.StringNull()
	}
	if data.Device.IsUnknown() {
		data.Device = types.StringNull()
	}
	if data.EndDate.IsUnknown() {
		data.EndDate = timetypes.NewRFC3339Null()
	}
	if data.ErrorMessage.IsUnknown() {
		data.ErrorMessage = types.StringNull()
	}
	if data.ExtendEnabled.IsUnknown() {
		data.ExtendEnabled = types.BoolNull()
	}
	if data.Image.IsUnknown() {
		data.Image = types.StringNull()
	}
	if data.ImageMetadata.IsUnknown() {
		data.ImageMetadata = types.StringNull()
	}
	if data.ImageName.IsUnknown() {
		data.ImageName = types.StringNull()
	}
	if data.Instance.IsUnknown() {
		data.Instance = types.StringNull()
	}
	if data.InstanceMarketplaceUuid.IsUnknown() {
		data.InstanceMarketplaceUuid = types.StringNull()
	}
	if data.InstanceName.IsUnknown() {
		data.InstanceName = types.StringNull()
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
	if data.ResourceType.IsUnknown() {
		data.ResourceType = types.StringNull()
	}
	if data.RuntimeState.IsUnknown() {
		data.RuntimeState = types.StringNull()
	}
	if data.Size.IsUnknown() {
		data.Size = types.Int64Null()
	}
	if data.SourceSnapshot.IsUnknown() {
		data.SourceSnapshot = types.StringNull()
	}
	if data.StartDate.IsUnknown() {
		data.StartDate = timetypes.NewRFC3339Null()
	}
	if data.State.IsUnknown() {
		data.State = types.StringNull()
	}
	if data.Tenant.IsUnknown() {
		data.Tenant = types.StringNull()
	}
	if data.TenantUuid.IsUnknown() {
		data.TenantUuid = types.StringNull()
	}
	if data.Type.IsUnknown() {
		data.Type = types.StringNull()
	}
	if data.TypeName.IsUnknown() {
		data.TypeName = types.StringNull()
	}
	if data.Url.IsUnknown() {
		data.Url = types.StringNull()
	}
}

func (r *OpenstackVolumeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var data OpenstackVolumeResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Phase 1: Payload Construction
	// We map the Terraform schema fields to the 'attributes' map required by the Marketplace Order API.
	attributes := OpenstackVolumeCreateAttributes{}
	if !data.AvailabilityZone.IsNull() && !data.AvailabilityZone.IsUnknown() {
		attributes.AvailabilityZone = data.AvailabilityZone.ValueStringPointer()
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		attributes.Description = data.Description.ValueStringPointer()
	}
	if !data.Image.IsNull() && !data.Image.IsUnknown() {
		attributes.Image = data.Image.ValueStringPointer()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		attributes.Name = data.Name.ValueStringPointer()
	}
	if !data.Size.IsNull() && !data.Size.IsUnknown() {
		attributes.Size = data.Size.ValueInt64Pointer()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		attributes.Type = data.Type.ValueStringPointer()
	}

	// Construct the Create Order Request
	payload := OpenstackVolumeCreateRequest{
		Attributes: attributes,
	}
	if !data.EndDate.IsNull() && !data.EndDate.IsUnknown() {
		payload.EndDate = data.EndDate.ValueStringPointer()
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
	payload.Offering = data.Offering.ValueStringPointer()
	if !data.Plan.IsNull() && !data.Plan.IsUnknown() {
		payload.Plan = data.Plan.ValueStringPointer()
	}
	payload.Project = data.Project.ValueStringPointer()
	if !data.StartDate.IsNull() && !data.StartDate.IsUnknown() {
		payload.StartDate = data.StartDate.ValueStringPointer()
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

func (r *OpenstackVolumeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackVolumeResourceModel

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
			"Unable to Read Openstack Volume",
			"An error occurred while reading the Openstack Volume: "+err.Error(),
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
	anyChanges := false
	var patchPayload OpenstackVolumeUpdateRequest
	if !data.Bootable.IsNull() && !data.Bootable.Equal(state.Bootable) {
		anyChanges = true
		patchPayload.Bootable = data.Bootable.ValueBoolPointer()
	}
	if !data.Description.IsNull() && !data.Description.Equal(state.Description) {
		anyChanges = true
		patchPayload.Description = data.Description.ValueStringPointer()
	}
	if !data.Name.IsNull() && !data.Name.Equal(state.Name) {
		anyChanges = true
		patchPayload.Name = data.Name.ValueStringPointer()
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
	if !data.Size.Equal(state.Size) {
		// Convert Terraform value to API payload for the specific action
		var req OpenstackVolumeExtendActionRequest
		req.Size = data.Size.ValueInt64Pointer()

		// Execute the Action
		if err := r.client.Extend(ctx, data.UUID.ValueString(), &req); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: extend", err.Error())
			return
		}

		// Wait for the resource to return to OK state
		apiResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackVolumeResponse, error) {
			return r.client.Get(ctx, data.UUID.ValueString())
		}, updateTimeout)
		if err != nil {
			resp.Diagnostics.AddError("Wait for RPC action failed", err.Error())
			return
		}
		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
		state = data // Update local state to avoid repeated action calls if multiple fields changed (though actions are usually 1-to-1)
	}
	if !data.Type.Equal(state.Type) {
		// Convert Terraform value to API payload for the specific action
		var req OpenstackVolumeRetypeActionRequest
		req.Type = data.Type.ValueStringPointer()

		// Execute the Action
		if err := r.client.Retype(ctx, data.UUID.ValueString(), &req); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: retype", err.Error())
			return
		}

		// Wait for the resource to return to OK state
		apiResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackVolumeResponse, error) {
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

func (r *OpenstackVolumeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var data OpenstackVolumeResourceModel
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

	apiResp, err := r.client.Get(ctx, uuid)
	if err != nil {
		if IsNotFoundError(err) {
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
