package project

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &StructureProjectResource{}
var _ resource.ResourceWithImportState = &StructureProjectResource{}

func NewStructureProjectResource() resource.Resource {
	return &StructureProjectResource{}
}

// StructureProjectResource defines the resource implementation.
type StructureProjectResource struct {
	client *Client
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
				MarkdownDescription: "Resource UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "ID of the backend",
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
				Required:            true,
				MarkdownDescription: "Customer",
			},
			"customer_display_billing_info_in_projects": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Customer display billing info in projects",
			},
			"customer_slug": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Customer slug",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project description (HTML content will be sanitized)",
			},
			"end_date": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project end date. Setting this field requires DELETE_PROJECT permission.",
			},
			"end_date_requested_by": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "End date requested by",
			},
			"grace_period_days": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Number of extra days after project end date before resources are terminated. Overrides customer-level setting.",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
					int64validator.AtMost(2147483647),
				},
			},
			"image": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Image",
			},
			"is_industry": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Is industry",
			},
			"is_removed": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Is removed",
			},
			"kind": schema.StringAttribute{
				Optional:            true,
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
				MarkdownDescription: "Name of the resource",
			},
			"oecd_fos_2007_code": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Oecd fos 2007 code",
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
				MarkdownDescription: "Project credit",
			},
			"resources_count": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Number of active resources in this project",
			},
			"slug": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
				},
			},
			"staff_notes": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Internal notes visible only to staff and support users (HTML content will be sanitized)",
			},
			"start_date": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project start date. Cannot be edited after the start date has arrived.",
			},
			"type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Type",
			},
			"type_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the type",
			},
			"type_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the type",
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

func (r *StructureProjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data StructureProjectResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := StructureProjectCreateRequest{
		BackendId:       data.BackendId.ValueStringPointer(),
		Customer:        data.Customer.ValueStringPointer(),
		Description:     data.Description.ValueStringPointer(),
		EndDate:         data.EndDate.ValueStringPointer(),
		GracePeriodDays: data.GracePeriodDays.ValueInt64Pointer(),
		Image:           data.Image.ValueStringPointer(),
		IsIndustry:      data.IsIndustry.ValueBoolPointer(),
		Kind:            data.Kind.ValueStringPointer(),
		Name:            data.Name.ValueStringPointer(),
		OecdFos2007Code: data.OecdFos2007Code.ValueStringPointer(),
		Slug:            data.Slug.ValueStringPointer(),
		StaffNotes:      data.StaffNotes.ValueStringPointer(),
		StartDate:       data.StartDate.ValueStringPointer(),
		Type:            data.Type.ValueStringPointer(),
	}

	apiResp, err := r.client.CreateStructureProject(ctx, &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Structure Project",
			"An error occurred while creating the Structure Project: "+err.Error(),
		)
		return
	}
	data.UUID = types.StringPointerValue(apiResp.UUID)

	createTimeout, diags := data.Timeouts.Create(ctx, common.DefaultCreateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*StructureProjectResponse, error) {
		return r.client.GetStructureProject(ctx, data.UUID.ValueString())
	}, createTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource creation", err.Error())
		return
	}
	apiResp = newResp

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

	apiResp, err := r.client.GetStructureProject(ctx, data.UUID.ValueString())
	if err != nil {
		if client.IsNotFoundError(err) {
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

	requestBody := StructureProjectUpdateRequest{
		BackendId:       data.BackendId.ValueStringPointer(),
		Customer:        data.Customer.ValueStringPointer(),
		Description:     data.Description.ValueStringPointer(),
		EndDate:         data.EndDate.ValueStringPointer(),
		GracePeriodDays: data.GracePeriodDays.ValueInt64Pointer(),
		Image:           data.Image.ValueStringPointer(),
		IsIndustry:      data.IsIndustry.ValueBoolPointer(),
		Kind:            data.Kind.ValueStringPointer(),
		Name:            data.Name.ValueStringPointer(),
		OecdFos2007Code: data.OecdFos2007Code.ValueStringPointer(),
		Slug:            data.Slug.ValueStringPointer(),
		StaffNotes:      data.StaffNotes.ValueStringPointer(),
		StartDate:       data.StartDate.ValueStringPointer(),
		Type:            data.Type.ValueStringPointer(),
	}

	apiResp, err := r.client.UpdateStructureProject(ctx, data.UUID.ValueString(), &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Structure Project",
			"An error occurred while updating the Structure Project: "+err.Error(),
		)
		return
	}

	updateTimeout, diags := data.Timeouts.Update(ctx, common.DefaultUpdateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*StructureProjectResponse, error) {
		return r.client.GetStructureProject(ctx, data.UUID.ValueString())
	}, updateTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource update", err.Error())
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

	err := r.client.DeleteStructureProject(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Structure Project",
			"An error occurred while deleting the Structure Project: "+err.Error(),
		)
		return
	}

	deleteTimeout, diags := data.Timeouts.Delete(ctx, common.DefaultDeleteTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err = common.WaitForDeletion(ctx, func(ctx context.Context) (*StructureProjectResponse, error) {
		return r.client.GetStructureProject(ctx, data.UUID.ValueString())
	}, deleteTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource deletion", err.Error())
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

	apiResp, err := r.client.GetStructureProject(ctx, uuid)
	if err != nil {
		if client.IsNotFoundError(err) {
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
