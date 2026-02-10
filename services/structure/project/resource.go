package project

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &StructureProjectResource{}
var _ resource.ResourceWithImportState = &StructureProjectResource{}

func NewStructureProjectResource() resource.Resource {
	return &StructureProjectResource{}
}

// StructureProjectResource defines the resource implementation.
type StructureProjectResource struct {
	client *StructureProjectClient
}

// StructureProjectResourceModel describes the resource data model.
type StructureProjectResourceModel struct {
	StructureProjectModel
	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

func (r *StructureProjectResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_structure_project"
}

func (r *StructureProjectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Structure Project resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Structure Project UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"backend_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Backend Id",
			},
			"billing_price_estimate": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"current": schema.Float64Attribute{
						Computed: true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Current",
					},
					"tax": schema.Float64Attribute{
						Computed: true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Tax",
					},
					"tax_current": schema.Float64Attribute{
						Computed: true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Tax Current",
					},
					"total": schema.Float64Attribute{
						Computed: true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Total",
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Billing Price Estimate",
			},
			"customer": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Customer",
			},
			"customer_display_billing_info_in_projects": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Customer Display Billing Info In Projects",
			},
			"customer_slug": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Customer Slug",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Project description (HTML content will be sanitized)",
			},
			"end_date": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Project end date. Setting this field requires DELETE_PROJECT permission.",
			},
			"end_date_requested_by": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "End Date Requested By",
			},
			"grace_period_days": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Number of extra days after project end date before resources are terminated. Overrides customer-level setting.",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
					int64validator.AtMost(2147483647),
				},
			},
			"image": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Image",
			},
			"is_industry": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Is Industry",
			},
			"is_removed": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Is Removed",
			},
			"kind": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Kind",
			},
			"max_service_accounts": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Maximum number of service accounts allowed",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
					int64validator.AtMost(32767),
				},
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name",
			},
			"oecd_fos_2007_code": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Oecd Fos 2007 Code",
			},
			"oecd_fos_2007_label": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Human-readable label for the OECD FOS 2007 classification code",
			},
			"project_credit": schema.Float64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Project Credit",
			},
			"resources_count": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Number of active resources in this project",
			},
			"slug": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
				},
			},
			"staff_notes": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Internal notes visible only to staff and support users (HTML content will be sanitized)",
			},
			"start_date": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Project start date. Cannot be edited after the start date has arrived.",
			},
			"type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Type",
			},
			"type_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Type Name",
			},
			"type_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Type Uuid",
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

func (r *StructureProjectResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.client = &StructureProjectClient{}
	if err := r.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			err.Error(),
		)
		return
	}
}

func (r *StructureProjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var data StructureProjectResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := StructureProjectCreateRequest{}
	if !data.BackendId.IsNull() && !data.BackendId.IsUnknown() {

		requestBody.BackendId = data.BackendId.ValueStringPointer()
	}

	requestBody.Customer = data.Customer.ValueStringPointer()
	if !data.Description.IsNull() && !data.Description.IsUnknown() {

		requestBody.Description = data.Description.ValueStringPointer()
	}
	if !data.EndDate.IsNull() && !data.EndDate.IsUnknown() {

		requestBody.EndDate = data.EndDate.ValueStringPointer()
	}
	if !data.GracePeriodDays.IsNull() && !data.GracePeriodDays.IsUnknown() {

		requestBody.GracePeriodDays = data.GracePeriodDays.ValueInt64Pointer()
	}
	if !data.Image.IsNull() && !data.Image.IsUnknown() {

		requestBody.Image = data.Image.ValueStringPointer()
	}
	if !data.IsIndustry.IsNull() && !data.IsIndustry.IsUnknown() {

		requestBody.IsIndustry = data.IsIndustry.ValueBoolPointer()
	}
	if !data.Kind.IsNull() && !data.Kind.IsUnknown() {

		requestBody.Kind = data.Kind.ValueStringPointer()
	}

	requestBody.Name = data.Name.ValueStringPointer()
	if !data.OecdFos2007Code.IsNull() && !data.OecdFos2007Code.IsUnknown() {

		requestBody.OecdFos2007Code = data.OecdFos2007Code.ValueStringPointer()
	}
	if !data.Slug.IsNull() && !data.Slug.IsUnknown() {

		requestBody.Slug = data.Slug.ValueStringPointer()
	}
	if !data.StaffNotes.IsNull() && !data.StaffNotes.IsUnknown() {

		requestBody.StaffNotes = data.StaffNotes.ValueStringPointer()
	}
	if !data.StartDate.IsNull() && !data.StartDate.IsUnknown() {

		requestBody.StartDate = data.StartDate.ValueStringPointer()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {

		requestBody.Type = data.Type.ValueStringPointer()
	}

	apiResp, err := r.client.Create(ctx, &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Structure Project",
			"An error occurred while creating the Structure Project: "+err.Error(),
		)
		return
	}
	data.UUID = types.StringPointerValue(apiResp.UUID)

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StructureProjectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data StructureProjectResourceModel

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
			"Unable to Read Structure Project",
			"An error occurred while reading the Structure Project: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StructureProjectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var data StructureProjectResourceModel
	var state StructureProjectResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp *StructureProjectResponse
	anyChanges := false
	requestBody := StructureProjectUpdateRequest{}
	if !data.BackendId.IsNull() && !data.BackendId.IsUnknown() && !data.BackendId.Equal(state.BackendId) {
		anyChanges = true

		requestBody.BackendId = data.BackendId.ValueStringPointer()
	}
	if !data.Customer.IsNull() && !data.Customer.IsUnknown() && !data.Customer.Equal(state.Customer) {
		anyChanges = true

		requestBody.Customer = data.Customer.ValueStringPointer()
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() && !data.Description.Equal(state.Description) {
		anyChanges = true

		requestBody.Description = data.Description.ValueStringPointer()
	}
	if !data.EndDate.IsNull() && !data.EndDate.IsUnknown() && !data.EndDate.Equal(state.EndDate) {
		anyChanges = true

		requestBody.EndDate = data.EndDate.ValueStringPointer()
	}
	if !data.GracePeriodDays.IsNull() && !data.GracePeriodDays.IsUnknown() && !data.GracePeriodDays.Equal(state.GracePeriodDays) {
		anyChanges = true

		requestBody.GracePeriodDays = data.GracePeriodDays.ValueInt64Pointer()
	}
	if !data.Image.IsNull() && !data.Image.IsUnknown() && !data.Image.Equal(state.Image) {
		anyChanges = true

		requestBody.Image = data.Image.ValueStringPointer()
	}
	if !data.IsIndustry.IsNull() && !data.IsIndustry.IsUnknown() && !data.IsIndustry.Equal(state.IsIndustry) {
		anyChanges = true

		requestBody.IsIndustry = data.IsIndustry.ValueBoolPointer()
	}
	if !data.Kind.IsNull() && !data.Kind.IsUnknown() && !data.Kind.Equal(state.Kind) {
		anyChanges = true

		requestBody.Kind = data.Kind.ValueStringPointer()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() && !data.Name.Equal(state.Name) {
		anyChanges = true

		requestBody.Name = data.Name.ValueStringPointer()
	}
	if !data.OecdFos2007Code.IsNull() && !data.OecdFos2007Code.IsUnknown() && !data.OecdFos2007Code.Equal(state.OecdFos2007Code) {
		anyChanges = true

		requestBody.OecdFos2007Code = data.OecdFos2007Code.ValueStringPointer()
	}
	if !data.Slug.IsNull() && !data.Slug.IsUnknown() && !data.Slug.Equal(state.Slug) {
		anyChanges = true

		requestBody.Slug = data.Slug.ValueStringPointer()
	}
	if !data.StaffNotes.IsNull() && !data.StaffNotes.IsUnknown() && !data.StaffNotes.Equal(state.StaffNotes) {
		anyChanges = true

		requestBody.StaffNotes = data.StaffNotes.ValueStringPointer()
	}
	if !data.StartDate.IsNull() && !data.StartDate.IsUnknown() && !data.StartDate.Equal(state.StartDate) {
		anyChanges = true

		requestBody.StartDate = data.StartDate.ValueStringPointer()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() && !data.Type.Equal(state.Type) {
		anyChanges = true

		requestBody.Type = data.Type.ValueStringPointer()
	}

	if anyChanges {
		var err error
		apiResp, err = r.client.Update(ctx, data.UUID.ValueString(), &requestBody)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Update Structure Project",
				"An error occurred while updating the Structure Project: "+err.Error(),
			)
			return
		}
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

func (r *StructureProjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var data StructureProjectResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Structure Project",
			"An error occurred while deleting the Structure Project: "+err.Error(),
		)
		return
	}
}

func (r *StructureProjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	uuid := req.ID
	if uuid == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Import ID cannot be empty. Please provide the UUID of the Structure Project.",
		)
		return
	}

	tflog.Info(ctx, "Importing Structure Project", map[string]interface{}{
		"uuid": uuid,
	})

	apiResp, err := r.client.Get(ctx, uuid)
	if err != nil {
		if IsNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("Structure Project with UUID '%s' does not exist or is not accessible.", uuid),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to Import Structure Project",
			fmt.Sprintf("An error occurred while fetching the Structure Project: %s", err.Error()),
		)
		return
	}

	var data StructureProjectResourceModel
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
