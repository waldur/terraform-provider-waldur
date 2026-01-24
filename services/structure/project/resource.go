package project

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
	UUID                                 types.String   `tfsdk:"id"`
	BackendId                            types.String   `tfsdk:"backend_id"`
	Created                              types.String   `tfsdk:"created"`
	Customer                             types.String   `tfsdk:"customer"`
	CustomerDisplayBillingInfoInProjects types.Bool     `tfsdk:"customer_display_billing_info_in_projects"`
	CustomerSlug                         types.String   `tfsdk:"customer_slug"`
	Description                          types.String   `tfsdk:"description"`
	EndDate                              types.String   `tfsdk:"end_date"`
	EndDateRequestedBy                   types.String   `tfsdk:"end_date_requested_by"`
	GracePeriodDays                      types.Int64    `tfsdk:"grace_period_days"`
	Image                                types.String   `tfsdk:"image"`
	IsIndustry                           types.Bool     `tfsdk:"is_industry"`
	IsRemoved                            types.Bool     `tfsdk:"is_removed"`
	Kind                                 types.String   `tfsdk:"kind"`
	MaxServiceAccounts                   types.Int64    `tfsdk:"max_service_accounts"`
	Name                                 types.String   `tfsdk:"name"`
	OecdFos2007Code                      types.String   `tfsdk:"oecd_fos_2007_code"`
	OecdFos2007Label                     types.String   `tfsdk:"oecd_fos_2007_label"`
	ProjectCredit                        types.Float64  `tfsdk:"project_credit"`
	ResourcesCount                       types.Int64    `tfsdk:"resources_count"`
	Slug                                 types.String   `tfsdk:"slug"`
	StaffNotes                           types.String   `tfsdk:"staff_notes"`
	StartDate                            types.String   `tfsdk:"start_date"`
	Type                                 types.String   `tfsdk:"type"`
	TypeName                             types.String   `tfsdk:"type_name"`
	TypeUuid                             types.String   `tfsdk:"type_uuid"`
	Url                                  types.String   `tfsdk:"url"`
	Timeouts                             timeouts.Value `tfsdk:"timeouts"`
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
				Computed: true,
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

	createTimeout, diags := data.Timeouts.Create(ctx, 30*time.Minute)
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

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

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

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

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

	updateTimeout, diags := data.Timeouts.Update(ctx, 30*time.Minute)
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

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

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

	deleteTimeout, diags := data.Timeouts.Delete(ctx, 10*time.Minute)
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
	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StructureProjectResource) mapResponseToModel(ctx context.Context, apiResp StructureProjectResponse, model *StructureProjectResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Customer = types.StringPointerValue(apiResp.Customer)
	model.CustomerDisplayBillingInfoInProjects = types.BoolPointerValue(apiResp.CustomerDisplayBillingInfoInProjects)
	model.CustomerSlug = types.StringPointerValue(apiResp.CustomerSlug)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.EndDate = types.StringPointerValue(apiResp.EndDate)
	model.EndDateRequestedBy = types.StringPointerValue(apiResp.EndDateRequestedBy)
	model.GracePeriodDays = types.Int64PointerValue(apiResp.GracePeriodDays)
	model.Image = types.StringPointerValue(apiResp.Image)
	model.IsIndustry = types.BoolPointerValue(apiResp.IsIndustry)
	model.IsRemoved = types.BoolPointerValue(apiResp.IsRemoved)
	model.Kind = types.StringPointerValue(apiResp.Kind)
	model.MaxServiceAccounts = types.Int64PointerValue(apiResp.MaxServiceAccounts)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.OecdFos2007Code = types.StringPointerValue(apiResp.OecdFos2007Code)
	model.OecdFos2007Label = types.StringPointerValue(apiResp.OecdFos2007Label)
	model.ProjectCredit = types.Float64PointerValue(apiResp.ProjectCredit)
	model.ResourcesCount = types.Int64PointerValue(apiResp.ResourcesCount)
	model.Slug = types.StringPointerValue(apiResp.Slug)
	model.StaffNotes = types.StringPointerValue(apiResp.StaffNotes)
	model.StartDate = types.StringPointerValue(apiResp.StartDate)
	model.Type = types.StringPointerValue(apiResp.Type)
	model.TypeName = types.StringPointerValue(apiResp.TypeName)
	model.TypeUuid = types.StringPointerValue(apiResp.TypeUuid)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
