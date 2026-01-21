package resources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &StructureProjectResource{}
var _ resource.ResourceWithImportState = &StructureProjectResource{}

func NewStructureProjectResource() resource.Resource {
	return &StructureProjectResource{}
}

// StructureProjectResource defines the resource implementation.
type StructureProjectResource struct {
	client *client.Client
}

// StructureProjectApiResponse is the API response model.
type StructureProjectApiResponse struct {
	UUID *string `json:"uuid"`

	BackendId                            *string  `json:"backend_id" tfsdk:"backend_id"`
	Created                              *string  `json:"created" tfsdk:"created"`
	Customer                             *string  `json:"customer" tfsdk:"customer"`
	CustomerDisplayBillingInfoInProjects *bool    `json:"customer_display_billing_info_in_projects" tfsdk:"customer_display_billing_info_in_projects"`
	CustomerSlug                         *string  `json:"customer_slug" tfsdk:"customer_slug"`
	Description                          *string  `json:"description" tfsdk:"description"`
	EndDate                              *string  `json:"end_date" tfsdk:"end_date"`
	EndDateRequestedBy                   *string  `json:"end_date_requested_by" tfsdk:"end_date_requested_by"`
	GracePeriodDays                      *int64   `json:"grace_period_days" tfsdk:"grace_period_days"`
	Image                                *string  `json:"image" tfsdk:"image"`
	IsIndustry                           *bool    `json:"is_industry" tfsdk:"is_industry"`
	IsRemoved                            *bool    `json:"is_removed" tfsdk:"is_removed"`
	Kind                                 *string  `json:"kind" tfsdk:"kind"`
	MaxServiceAccounts                   *int64   `json:"max_service_accounts" tfsdk:"max_service_accounts"`
	Name                                 *string  `json:"name" tfsdk:"name"`
	OecdFos2007Code                      *string  `json:"oecd_fos_2007_code" tfsdk:"oecd_fos_2007_code"`
	OecdFos2007Label                     *string  `json:"oecd_fos_2007_label" tfsdk:"oecd_fos_2007_label"`
	ProjectCredit                        *float64 `json:"project_credit" tfsdk:"project_credit"`
	ResourcesCount                       *int64   `json:"resources_count" tfsdk:"resources_count"`
	Slug                                 *string  `json:"slug" tfsdk:"slug"`
	StaffNotes                           *string  `json:"staff_notes" tfsdk:"staff_notes"`
	StartDate                            *string  `json:"start_date" tfsdk:"start_date"`
	Type                                 *string  `json:"type" tfsdk:"type"`
	TypeName                             *string  `json:"type_name" tfsdk:"type_name"`
	TypeUuid                             *string  `json:"type_uuid" tfsdk:"type_uuid"`
	Url                                  *string  `json:"url" tfsdk:"url"`
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
				Computed:            true,
				MarkdownDescription: " ",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"customer": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: " ",
			},
			"customer_display_billing_info_in_projects": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"customer_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project description (HTML content will be sanitized)",
			},
			"end_date": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project end date. Setting this field requires DELETE_PROJECT permission.",
			},
			"end_date_requested_by": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"grace_period_days": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Number of extra days after project end date before resources are terminated. Overrides customer-level setting.",
			},
			"image": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"is_industry": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"is_removed": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"kind": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"max_service_accounts": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Maximum number of service accounts allowed",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: " ",
			},
			"oecd_fos_2007_code": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"oecd_fos_2007_label": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Human-readable label for the OECD FOS 2007 classification code",
			},
			"project_credit": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"resources_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of active resources in this project",
			},
			"slug": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
			},
			"staff_notes": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Internal notes visible only to staff and support users (HTML content will be sanitized)",
			},
			"start_date": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project start date. Cannot be edited after the start date has arrived.",
			},
			"type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"type_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"type_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
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

	r.client = client
}

func (r *StructureProjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data StructureProjectResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to create resource
	var apiResp StructureProjectApiResponse // Prepare request body
	requestBody := map[string]interface{}{}
	if !data.BackendId.IsNull() && !data.BackendId.IsUnknown() {
		if v := data.BackendId.ValueString(); v != "" {
			requestBody["backend_id"] = v
		}
	}
	requestBody["customer"] = data.Customer.ValueString()
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		if v := data.Description.ValueString(); v != "" {
			requestBody["description"] = v
		}
	}
	if !data.EndDate.IsNull() && !data.EndDate.IsUnknown() {
		if v := data.EndDate.ValueString(); v != "" {
			requestBody["end_date"] = v
		}
	}
	if !data.GracePeriodDays.IsNull() && !data.GracePeriodDays.IsUnknown() {
		requestBody["grace_period_days"] = data.GracePeriodDays.ValueInt64()
	}
	if !data.Image.IsNull() && !data.Image.IsUnknown() {
		if v := data.Image.ValueString(); v != "" {
			requestBody["image"] = v
		}
	}
	if !data.IsIndustry.IsNull() && !data.IsIndustry.IsUnknown() {
		requestBody["is_industry"] = data.IsIndustry.ValueBool()
	}
	if !data.Kind.IsNull() && !data.Kind.IsUnknown() {
		if v := data.Kind.ValueString(); v != "" {
			requestBody["kind"] = v
		}
	}
	requestBody["name"] = data.Name.ValueString()
	if !data.OecdFos2007Code.IsNull() && !data.OecdFos2007Code.IsUnknown() {
		if v := data.OecdFos2007Code.ValueString(); v != "" {
			requestBody["oecd_fos_2007_code"] = v
		}
	}
	if !data.Slug.IsNull() && !data.Slug.IsUnknown() {
		if v := data.Slug.ValueString(); v != "" {
			requestBody["slug"] = v
		}
	}
	if !data.StaffNotes.IsNull() && !data.StaffNotes.IsUnknown() {
		if v := data.StaffNotes.ValueString(); v != "" {
			requestBody["staff_notes"] = v
		}
	}
	if !data.StartDate.IsNull() && !data.StartDate.IsUnknown() {
		if v := data.StartDate.ValueString(); v != "" {
			requestBody["start_date"] = v
		}
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		if v := data.Type.ValueString(); v != "" {
			requestBody["type"] = v
		}
	}
	err := r.client.Create(ctx, "/api/projects/", requestBody, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Structure Project",
			"An error occurred while creating the Structure Project: "+err.Error(),
		)
		return
	}
	// Extract UUID from response
	if apiResp.UUID != nil {
		data.UUID = types.StringPointerValue(apiResp.UUID)
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

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

	retrievePath := strings.Replace("/api/projects/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	var apiResp StructureProjectApiResponse
	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Structure Project",
			"An error occurred while reading the Structure Project: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

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

	// Prepare request body
	requestBody := map[string]interface{}{}
	if !data.BackendId.IsNull() && !data.BackendId.IsUnknown() {
		if v := data.BackendId.ValueString(); v != "" {
			requestBody["backend_id"] = v
		}
	}
	if !data.Customer.IsNull() && !data.Customer.IsUnknown() {
		if v := data.Customer.ValueString(); v != "" {
			requestBody["customer"] = v
		}
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		if v := data.Description.ValueString(); v != "" {
			requestBody["description"] = v
		}
	}
	if !data.EndDate.IsNull() && !data.EndDate.IsUnknown() {
		if v := data.EndDate.ValueString(); v != "" {
			requestBody["end_date"] = v
		}
	}
	if !data.GracePeriodDays.IsNull() && !data.GracePeriodDays.IsUnknown() {
		requestBody["grace_period_days"] = data.GracePeriodDays.ValueInt64()
	}
	if !data.Image.IsNull() && !data.Image.IsUnknown() {
		if v := data.Image.ValueString(); v != "" {
			requestBody["image"] = v
		}
	}
	if !data.IsIndustry.IsNull() && !data.IsIndustry.IsUnknown() {
		requestBody["is_industry"] = data.IsIndustry.ValueBool()
	}
	if !data.Kind.IsNull() && !data.Kind.IsUnknown() {
		if v := data.Kind.ValueString(); v != "" {
			requestBody["kind"] = v
		}
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		if v := data.Name.ValueString(); v != "" {
			requestBody["name"] = v
		}
	}
	if !data.OecdFos2007Code.IsNull() && !data.OecdFos2007Code.IsUnknown() {
		if v := data.OecdFos2007Code.ValueString(); v != "" {
			requestBody["oecd_fos_2007_code"] = v
		}
	}
	if !data.Slug.IsNull() && !data.Slug.IsUnknown() {
		if v := data.Slug.ValueString(); v != "" {
			requestBody["slug"] = v
		}
	}
	if !data.StaffNotes.IsNull() && !data.StaffNotes.IsUnknown() {
		if v := data.StaffNotes.ValueString(); v != "" {
			requestBody["staff_notes"] = v
		}
	}
	if !data.StartDate.IsNull() && !data.StartDate.IsUnknown() {
		if v := data.StartDate.ValueString(); v != "" {
			requestBody["start_date"] = v
		}
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		if v := data.Type.ValueString(); v != "" {
			requestBody["type"] = v
		}
	}

	// Call Waldur API to update resource
	var apiResp StructureProjectApiResponse

	err := r.client.Update(ctx, "/api/projects/{uuid}/", data.UUID.ValueString(), requestBody, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Structure Project",
			"An error occurred while updating the Structure Project: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if apiResp.UUID != nil {
		data.UUID = types.StringPointerValue(apiResp.UUID)
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StructureProjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data StructureProjectResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to delete resource
	err := r.client.DeleteByUUID(ctx, "/api/projects/{uuid}/", data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Structure Project",
			"An error occurred while deleting the Structure Project: "+err.Error(),
		)
		return
	}
}

func (r *StructureProjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *StructureProjectResource) mapResponseToModel(ctx context.Context, apiResp StructureProjectApiResponse, model *StructureProjectResourceModel) diag.Diagnostics {
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
