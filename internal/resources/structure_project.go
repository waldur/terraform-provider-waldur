package resources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-waldur-provider/internal/client"
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
		MarkdownDescription: "StructureProject resource",

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
				MarkdownDescription: "",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "",
			},
			"customer_display_billing_info_in_projects": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"end_date": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"end_date_requested_by": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"grace_period_days": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Number of extra days after project end date before resources are terminated. Overrides customer-level setting.",
			},
			"image": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"is_industry": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"is_removed": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"kind": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"max_service_accounts": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Maximum number of service accounts allowed",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "",
			},
			"oecd_fos_2007_code": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"oecd_fos_2007_label": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"project_credit": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"resources_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"slug": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"staff_notes": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"start_date": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"type_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"type_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
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

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

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

	// Call Waldur API to create resource
	var result map[string]interface{}
	err := r.client.Create(ctx, "/api/projects/", requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create StructureProject",
			"An error occurred while creating the structure_project: "+err.Error(),
		)
		return
	}
	// Extract UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
	// Map response fields to data model
	_ = sourceMap
	if val, ok := sourceMap["backend_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
		}
	} else {
		if data.BackendId.IsUnknown() {
			data.BackendId = types.StringNull()
		}
	}
	if val, ok := sourceMap["created"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
	} else {
		if data.Created.IsUnknown() {
			data.Created = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Customer = types.StringValue(str)
		}
	} else {
		if data.Customer.IsUnknown() {
			data.Customer = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_display_billing_info_in_projects"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.CustomerDisplayBillingInfoInProjects = types.BoolValue(b)
		}
	} else {
		if data.CustomerDisplayBillingInfoInProjects.IsUnknown() {
			data.CustomerDisplayBillingInfoInProjects = types.BoolNull()
		}
	}
	if val, ok := sourceMap["customer_slug"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CustomerSlug = types.StringValue(str)
		}
	} else {
		if data.CustomerSlug.IsUnknown() {
			data.CustomerSlug = types.StringNull()
		}
	}
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
		}
	}
	if val, ok := sourceMap["end_date"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.EndDate = types.StringValue(str)
		}
	} else {
		if data.EndDate.IsUnknown() {
			data.EndDate = types.StringNull()
		}
	}
	if val, ok := sourceMap["end_date_requested_by"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.EndDateRequestedBy = types.StringValue(str)
		}
	} else {
		if data.EndDateRequestedBy.IsUnknown() {
			data.EndDateRequestedBy = types.StringNull()
		}
	}
	if val, ok := sourceMap["grace_period_days"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.GracePeriodDays = types.Int64Value(int64(num))
		}
	} else {
		if data.GracePeriodDays.IsUnknown() {
			data.GracePeriodDays = types.Int64Null()
		}
	}
	if val, ok := sourceMap["image"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Image = types.StringValue(str)
		}
	} else {
		if data.Image.IsUnknown() {
			data.Image = types.StringNull()
		}
	}
	if val, ok := sourceMap["is_industry"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.IsIndustry = types.BoolValue(b)
		}
	} else {
		if data.IsIndustry.IsUnknown() {
			data.IsIndustry = types.BoolNull()
		}
	}
	if val, ok := sourceMap["is_removed"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.IsRemoved = types.BoolValue(b)
		}
	} else {
		if data.IsRemoved.IsUnknown() {
			data.IsRemoved = types.BoolNull()
		}
	}
	if val, ok := sourceMap["kind"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Kind = types.StringValue(str)
		}
	} else {
		if data.Kind.IsUnknown() {
			data.Kind = types.StringNull()
		}
	}
	if val, ok := sourceMap["max_service_accounts"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.MaxServiceAccounts = types.Int64Value(int64(num))
		}
	} else {
		if data.MaxServiceAccounts.IsUnknown() {
			data.MaxServiceAccounts = types.Int64Null()
		}
	}
	if val, ok := sourceMap["oecd_fos_2007_code"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.OecdFos2007Code = types.StringValue(str)
		}
	} else {
		if data.OecdFos2007Code.IsUnknown() {
			data.OecdFos2007Code = types.StringNull()
		}
	}
	if val, ok := sourceMap["oecd_fos_2007_label"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.OecdFos2007Label = types.StringValue(str)
		}
	} else {
		if data.OecdFos2007Label.IsUnknown() {
			data.OecdFos2007Label = types.StringNull()
		}
	}
	if val, ok := sourceMap["project_credit"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.ProjectCredit = types.Float64Value(num)
		}
	} else {
		if data.ProjectCredit.IsUnknown() {
			data.ProjectCredit = types.Float64Null()
		}
	}
	if val, ok := sourceMap["resources_count"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.ResourcesCount = types.Int64Value(int64(num))
		}
	} else {
		if data.ResourcesCount.IsUnknown() {
			data.ResourcesCount = types.Int64Null()
		}
	}
	if val, ok := sourceMap["slug"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Slug = types.StringValue(str)
		}
	} else {
		if data.Slug.IsUnknown() {
			data.Slug = types.StringNull()
		}
	}
	if val, ok := sourceMap["staff_notes"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.StaffNotes = types.StringValue(str)
		}
	} else {
		if data.StaffNotes.IsUnknown() {
			data.StaffNotes = types.StringNull()
		}
	}
	if val, ok := sourceMap["start_date"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.StartDate = types.StringValue(str)
		}
	} else {
		if data.StartDate.IsUnknown() {
			data.StartDate = types.StringNull()
		}
	}
	if val, ok := sourceMap["type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Type = types.StringValue(str)
		}
	} else {
		if data.Type.IsUnknown() {
			data.Type = types.StringNull()
		}
	}
	if val, ok := sourceMap["type_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.TypeName = types.StringValue(str)
		}
	} else {
		if data.TypeName.IsUnknown() {
			data.TypeName = types.StringNull()
		}
	}
	if val, ok := sourceMap["type_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.TypeUuid = types.StringValue(str)
		}
	} else {
		if data.TypeUuid.IsUnknown() {
			data.TypeUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Url = types.StringValue(str)
		}
	} else {
		if data.Url.IsUnknown() {
			data.Url = types.StringNull()
		}
	}

	// Map filter parameters from response if available

	// Save data into Terraform state
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
	var result map[string]interface{}

	retrievePath := strings.Replace("/api/projects/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read StructureProject",
			"An error occurred while reading the structure_project: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
	// Map response fields to data model
	_ = sourceMap
	if val, ok := sourceMap["backend_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
		}
	} else {
		if data.BackendId.IsUnknown() {
			data.BackendId = types.StringNull()
		}
	}
	if val, ok := sourceMap["created"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
	} else {
		if data.Created.IsUnknown() {
			data.Created = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Customer = types.StringValue(str)
		}
	} else {
		if data.Customer.IsUnknown() {
			data.Customer = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_display_billing_info_in_projects"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.CustomerDisplayBillingInfoInProjects = types.BoolValue(b)
		}
	} else {
		if data.CustomerDisplayBillingInfoInProjects.IsUnknown() {
			data.CustomerDisplayBillingInfoInProjects = types.BoolNull()
		}
	}
	if val, ok := sourceMap["customer_slug"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CustomerSlug = types.StringValue(str)
		}
	} else {
		if data.CustomerSlug.IsUnknown() {
			data.CustomerSlug = types.StringNull()
		}
	}
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
		}
	}
	if val, ok := sourceMap["end_date"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.EndDate = types.StringValue(str)
		}
	} else {
		if data.EndDate.IsUnknown() {
			data.EndDate = types.StringNull()
		}
	}
	if val, ok := sourceMap["end_date_requested_by"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.EndDateRequestedBy = types.StringValue(str)
		}
	} else {
		if data.EndDateRequestedBy.IsUnknown() {
			data.EndDateRequestedBy = types.StringNull()
		}
	}
	if val, ok := sourceMap["grace_period_days"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.GracePeriodDays = types.Int64Value(int64(num))
		}
	} else {
		if data.GracePeriodDays.IsUnknown() {
			data.GracePeriodDays = types.Int64Null()
		}
	}
	if val, ok := sourceMap["image"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Image = types.StringValue(str)
		}
	} else {
		if data.Image.IsUnknown() {
			data.Image = types.StringNull()
		}
	}
	if val, ok := sourceMap["is_industry"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.IsIndustry = types.BoolValue(b)
		}
	} else {
		if data.IsIndustry.IsUnknown() {
			data.IsIndustry = types.BoolNull()
		}
	}
	if val, ok := sourceMap["is_removed"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.IsRemoved = types.BoolValue(b)
		}
	} else {
		if data.IsRemoved.IsUnknown() {
			data.IsRemoved = types.BoolNull()
		}
	}
	if val, ok := sourceMap["kind"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Kind = types.StringValue(str)
		}
	} else {
		if data.Kind.IsUnknown() {
			data.Kind = types.StringNull()
		}
	}
	if val, ok := sourceMap["max_service_accounts"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.MaxServiceAccounts = types.Int64Value(int64(num))
		}
	} else {
		if data.MaxServiceAccounts.IsUnknown() {
			data.MaxServiceAccounts = types.Int64Null()
		}
	}
	if val, ok := sourceMap["oecd_fos_2007_code"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.OecdFos2007Code = types.StringValue(str)
		}
	} else {
		if data.OecdFos2007Code.IsUnknown() {
			data.OecdFos2007Code = types.StringNull()
		}
	}
	if val, ok := sourceMap["oecd_fos_2007_label"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.OecdFos2007Label = types.StringValue(str)
		}
	} else {
		if data.OecdFos2007Label.IsUnknown() {
			data.OecdFos2007Label = types.StringNull()
		}
	}
	if val, ok := sourceMap["project_credit"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.ProjectCredit = types.Float64Value(num)
		}
	} else {
		if data.ProjectCredit.IsUnknown() {
			data.ProjectCredit = types.Float64Null()
		}
	}
	if val, ok := sourceMap["resources_count"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.ResourcesCount = types.Int64Value(int64(num))
		}
	} else {
		if data.ResourcesCount.IsUnknown() {
			data.ResourcesCount = types.Int64Null()
		}
	}
	if val, ok := sourceMap["slug"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Slug = types.StringValue(str)
		}
	} else {
		if data.Slug.IsUnknown() {
			data.Slug = types.StringNull()
		}
	}
	if val, ok := sourceMap["staff_notes"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.StaffNotes = types.StringValue(str)
		}
	} else {
		if data.StaffNotes.IsUnknown() {
			data.StaffNotes = types.StringNull()
		}
	}
	if val, ok := sourceMap["start_date"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.StartDate = types.StringValue(str)
		}
	} else {
		if data.StartDate.IsUnknown() {
			data.StartDate = types.StringNull()
		}
	}
	if val, ok := sourceMap["type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Type = types.StringValue(str)
		}
	} else {
		if data.Type.IsUnknown() {
			data.Type = types.StringNull()
		}
	}
	if val, ok := sourceMap["type_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.TypeName = types.StringValue(str)
		}
	} else {
		if data.TypeName.IsUnknown() {
			data.TypeName = types.StringNull()
		}
	}
	if val, ok := sourceMap["type_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.TypeUuid = types.StringValue(str)
		}
	} else {
		if data.TypeUuid.IsUnknown() {
			data.TypeUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Url = types.StringValue(str)
		}
	} else {
		if data.Url.IsUnknown() {
			data.Url = types.StringNull()
		}
	}

	// Map filter parameters from response if available

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StructureProjectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data StructureProjectResourceModel
	var state StructureProjectResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	// Read current state to get the UUID (which is computed and not in plan)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Use UUID from state
	data.UUID = state.UUID
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
	var result map[string]interface{}

	err := r.client.Update(ctx, "/api/projects/{uuid}/", data.UUID.ValueString(), requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update StructureProject",
			"An error occurred while updating the structure_project: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
	// Map response fields to data model
	_ = sourceMap
	if val, ok := sourceMap["backend_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
		}
	} else {
		if data.BackendId.IsUnknown() {
			data.BackendId = types.StringNull()
		}
	}
	if val, ok := sourceMap["created"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
	} else {
		if data.Created.IsUnknown() {
			data.Created = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Customer = types.StringValue(str)
		}
	} else {
		if data.Customer.IsUnknown() {
			data.Customer = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_display_billing_info_in_projects"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.CustomerDisplayBillingInfoInProjects = types.BoolValue(b)
		}
	} else {
		if data.CustomerDisplayBillingInfoInProjects.IsUnknown() {
			data.CustomerDisplayBillingInfoInProjects = types.BoolNull()
		}
	}
	if val, ok := sourceMap["customer_slug"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CustomerSlug = types.StringValue(str)
		}
	} else {
		if data.CustomerSlug.IsUnknown() {
			data.CustomerSlug = types.StringNull()
		}
	}
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
		}
	}
	if val, ok := sourceMap["end_date"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.EndDate = types.StringValue(str)
		}
	} else {
		if data.EndDate.IsUnknown() {
			data.EndDate = types.StringNull()
		}
	}
	if val, ok := sourceMap["end_date_requested_by"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.EndDateRequestedBy = types.StringValue(str)
		}
	} else {
		if data.EndDateRequestedBy.IsUnknown() {
			data.EndDateRequestedBy = types.StringNull()
		}
	}
	if val, ok := sourceMap["grace_period_days"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.GracePeriodDays = types.Int64Value(int64(num))
		}
	} else {
		if data.GracePeriodDays.IsUnknown() {
			data.GracePeriodDays = types.Int64Null()
		}
	}
	if val, ok := sourceMap["image"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Image = types.StringValue(str)
		}
	} else {
		if data.Image.IsUnknown() {
			data.Image = types.StringNull()
		}
	}
	if val, ok := sourceMap["is_industry"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.IsIndustry = types.BoolValue(b)
		}
	} else {
		if data.IsIndustry.IsUnknown() {
			data.IsIndustry = types.BoolNull()
		}
	}
	if val, ok := sourceMap["is_removed"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.IsRemoved = types.BoolValue(b)
		}
	} else {
		if data.IsRemoved.IsUnknown() {
			data.IsRemoved = types.BoolNull()
		}
	}
	if val, ok := sourceMap["kind"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Kind = types.StringValue(str)
		}
	} else {
		if data.Kind.IsUnknown() {
			data.Kind = types.StringNull()
		}
	}
	if val, ok := sourceMap["max_service_accounts"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.MaxServiceAccounts = types.Int64Value(int64(num))
		}
	} else {
		if data.MaxServiceAccounts.IsUnknown() {
			data.MaxServiceAccounts = types.Int64Null()
		}
	}
	if val, ok := sourceMap["oecd_fos_2007_code"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.OecdFos2007Code = types.StringValue(str)
		}
	} else {
		if data.OecdFos2007Code.IsUnknown() {
			data.OecdFos2007Code = types.StringNull()
		}
	}
	if val, ok := sourceMap["oecd_fos_2007_label"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.OecdFos2007Label = types.StringValue(str)
		}
	} else {
		if data.OecdFos2007Label.IsUnknown() {
			data.OecdFos2007Label = types.StringNull()
		}
	}
	if val, ok := sourceMap["project_credit"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.ProjectCredit = types.Float64Value(num)
		}
	} else {
		if data.ProjectCredit.IsUnknown() {
			data.ProjectCredit = types.Float64Null()
		}
	}
	if val, ok := sourceMap["resources_count"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.ResourcesCount = types.Int64Value(int64(num))
		}
	} else {
		if data.ResourcesCount.IsUnknown() {
			data.ResourcesCount = types.Int64Null()
		}
	}
	if val, ok := sourceMap["slug"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Slug = types.StringValue(str)
		}
	} else {
		if data.Slug.IsUnknown() {
			data.Slug = types.StringNull()
		}
	}
	if val, ok := sourceMap["staff_notes"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.StaffNotes = types.StringValue(str)
		}
	} else {
		if data.StaffNotes.IsUnknown() {
			data.StaffNotes = types.StringNull()
		}
	}
	if val, ok := sourceMap["start_date"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.StartDate = types.StringValue(str)
		}
	} else {
		if data.StartDate.IsUnknown() {
			data.StartDate = types.StringNull()
		}
	}
	if val, ok := sourceMap["type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Type = types.StringValue(str)
		}
	} else {
		if data.Type.IsUnknown() {
			data.Type = types.StringNull()
		}
	}
	if val, ok := sourceMap["type_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.TypeName = types.StringValue(str)
		}
	} else {
		if data.TypeName.IsUnknown() {
			data.TypeName = types.StringNull()
		}
	}
	if val, ok := sourceMap["type_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.TypeUuid = types.StringValue(str)
		}
	} else {
		if data.TypeUuid.IsUnknown() {
			data.TypeUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Url = types.StringValue(str)
		}
	} else {
		if data.Url.IsUnknown() {
			data.Url = types.StringNull()
		}
	}

	// Map filter parameters from response if available

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StructureProjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data StructureProjectResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Call Waldur API to delete resource
	err := r.client.DeleteByUUID(ctx, "/api/projects/{uuid}/", data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete StructureProject",
			"An error occurred while deleting the structure_project: "+err.Error(),
		)
		return
	}
}

func (r *StructureProjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
