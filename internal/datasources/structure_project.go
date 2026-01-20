package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &StructureProjectDataSource{}

func NewStructureProjectDataSource() datasource.DataSource {
	return &StructureProjectDataSource{}
}

// StructureProjectDataSource defines the data source implementation.
type StructureProjectDataSource struct {
	client *client.Client
}

// StructureProjectDataSourceModel describes the data source data model.
type StructureProjectDataSourceModel struct {
	UUID                                 types.String  `tfsdk:"id"`
	BackendId                            types.String  `tfsdk:"backend_id"`
	CanAdmin                             types.Bool    `tfsdk:"can_admin"`
	CanManage                            types.Bool    `tfsdk:"can_manage"`
	ConcealFinishedProjects              types.Bool    `tfsdk:"conceal_finished_projects"`
	Created                              types.String  `tfsdk:"created"`
	Customer                             types.String  `tfsdk:"customer"`
	CustomerAbbreviation                 types.String  `tfsdk:"customer_abbreviation"`
	CustomerName                         types.String  `tfsdk:"customer_name"`
	CustomerNativeName                   types.String  `tfsdk:"customer_native_name"`
	Description                          types.String  `tfsdk:"description"`
	IncludeTerminated                    types.Bool    `tfsdk:"include_terminated"`
	IsRemoved                            types.Bool    `tfsdk:"is_removed"`
	Modified                             types.String  `tfsdk:"modified"`
	Name                                 types.String  `tfsdk:"name"`
	NameExact                            types.String  `tfsdk:"name_exact"`
	Query                                types.String  `tfsdk:"query"`
	Slug                                 types.String  `tfsdk:"slug"`
	CustomerDisplayBillingInfoInProjects types.Bool    `tfsdk:"customer_display_billing_info_in_projects"`
	CustomerSlug                         types.String  `tfsdk:"customer_slug"`
	EndDate                              types.String  `tfsdk:"end_date"`
	EndDateRequestedBy                   types.String  `tfsdk:"end_date_requested_by"`
	GracePeriodDays                      types.Int64   `tfsdk:"grace_period_days"`
	Image                                types.String  `tfsdk:"image"`
	IsIndustry                           types.Bool    `tfsdk:"is_industry"`
	Kind                                 types.String  `tfsdk:"kind"`
	MaxServiceAccounts                   types.Int64   `tfsdk:"max_service_accounts"`
	OecdFos2007Code                      types.String  `tfsdk:"oecd_fos_2007_code"`
	OecdFos2007Label                     types.String  `tfsdk:"oecd_fos_2007_label"`
	ProjectCredit                        types.Float64 `tfsdk:"project_credit"`
	ResourcesCount                       types.Int64   `tfsdk:"resources_count"`
	StaffNotes                           types.String  `tfsdk:"staff_notes"`
	StartDate                            types.String  `tfsdk:"start_date"`
	Type                                 types.String  `tfsdk:"type"`
	TypeName                             types.String  `tfsdk:"type_name"`
	TypeUuid                             types.String  `tfsdk:"type_uuid"`
	Url                                  types.String  `tfsdk:"url"`
}

func (d *StructureProjectDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_structure_project"
}

func (d *StructureProjectDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Structure Project data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"can_admin": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Return a list of projects where current user is admin.",
			},
			"can_manage": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Return a list of projects where current user is manager or a customer owner.",
			},
			"conceal_finished_projects": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Conceal finished projects",
			},
			"created": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Created after",
			},
			"customer": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Multiple values may be separated by commas.",
			},
			"customer_abbreviation": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"customer_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"customer_native_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"include_terminated": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Include soft-deleted (terminated) projects. Only available to staff and support users, or users with organizational roles who can see their terminated projects.",
			},
			"is_removed": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"modified": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Modified after",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"query": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Filter by name, slug, UUID, backend ID or resource effective ID",
			},
			"slug": schema.StringAttribute{
				Optional:            true,
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
			"end_date": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"end_date_requested_by": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"grace_period_days": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of extra days after project end date before resources are terminated. Overrides customer-level setting.",
			},
			"image": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"is_industry": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"kind": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"max_service_accounts": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Maximum number of service accounts allowed",
			},
			"oecd_fos_2007_code": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"oecd_fos_2007_label": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"project_credit": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"resources_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"staff_notes": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"start_date": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"type": schema.StringAttribute{
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
	}
}

func (d *StructureProjectDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
}

func (d *StructureProjectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data StructureProjectDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var item map[string]interface{}

		err := d.client.GetByUUID(ctx, "/api/projects/{uuid}/", data.UUID.ValueString(), &item)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Project",
				"An error occurred while reading the structure_project by UUID: "+err.Error(),
			)
			return
		}

		// Extract data from single result
		if uuid, ok := item["uuid"].(string); ok {
			data.UUID = types.StringValue(uuid)
		}

		sourceMap := item
		// Map response fields to data model
		_ = sourceMap
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
		if val, ok := sourceMap["backend_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BackendId = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["can_admin"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanAdmin = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["can_manage"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanManage = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["conceal_finished_projects"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.ConcealFinishedProjects = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["created"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Created = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Customer = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_abbreviation"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerAbbreviation = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_native_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerNativeName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Description = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["include_terminated"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IncludeTerminated = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["is_removed"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsRemoved = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["modified"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Modified = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Name = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name_exact"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NameExact = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["query"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Query = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["slug"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Slug = types.StringValue(str)
			}
		}

	} else {
		// Filter by provided parameters
		var results []map[string]interface{}

		filters := map[string]string{}
		if !data.BackendId.IsNull() {
			filters["backend_id"] = data.BackendId.ValueString()
		}
		if !data.CanAdmin.IsNull() {
			filters["can_admin"] = fmt.Sprintf("%t", data.CanAdmin.ValueBool())
		}
		if !data.CanManage.IsNull() {
			filters["can_manage"] = fmt.Sprintf("%t", data.CanManage.ValueBool())
		}
		if !data.ConcealFinishedProjects.IsNull() {
			filters["conceal_finished_projects"] = fmt.Sprintf("%t", data.ConcealFinishedProjects.ValueBool())
		}
		if !data.Created.IsNull() {
			filters["created"] = data.Created.ValueString()
		}
		if !data.Customer.IsNull() {
			filters["customer"] = data.Customer.ValueString()
		}
		if !data.CustomerAbbreviation.IsNull() {
			filters["customer_abbreviation"] = data.CustomerAbbreviation.ValueString()
		}
		if !data.CustomerName.IsNull() {
			filters["customer_name"] = data.CustomerName.ValueString()
		}
		if !data.CustomerNativeName.IsNull() {
			filters["customer_native_name"] = data.CustomerNativeName.ValueString()
		}
		if !data.Description.IsNull() {
			filters["description"] = data.Description.ValueString()
		}
		if !data.IncludeTerminated.IsNull() {
			filters["include_terminated"] = fmt.Sprintf("%t", data.IncludeTerminated.ValueBool())
		}
		if !data.IsRemoved.IsNull() {
			filters["is_removed"] = fmt.Sprintf("%t", data.IsRemoved.ValueBool())
		}
		if !data.Modified.IsNull() {
			filters["modified"] = data.Modified.ValueString()
		}
		if !data.Name.IsNull() {
			filters["name"] = data.Name.ValueString()
		}
		if !data.NameExact.IsNull() {
			filters["name_exact"] = data.NameExact.ValueString()
		}
		if !data.Query.IsNull() {
			filters["query"] = data.Query.ValueString()
		}
		if !data.Slug.IsNull() {
			filters["slug"] = data.Slug.ValueString()
		}

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup structure_project.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/projects/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Project",
				"An error occurred while filtering structure_project: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Project Not Found",
				"No structure_project found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Projects Found",
				fmt.Sprintf("Found %d structure_projects with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		// Extract data from the single result
		if uuid, ok := results[0]["uuid"].(string); ok {
			data.UUID = types.StringValue(uuid)
		}
		sourceMap := results[0]
		// Map response fields to data model
		_ = sourceMap
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
		if val, ok := sourceMap["backend_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BackendId = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["can_admin"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanAdmin = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["can_manage"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanManage = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["conceal_finished_projects"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.ConcealFinishedProjects = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["created"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Created = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Customer = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_abbreviation"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerAbbreviation = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_native_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerNativeName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Description = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["include_terminated"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IncludeTerminated = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["is_removed"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsRemoved = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["modified"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Modified = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Name = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name_exact"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NameExact = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["query"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Query = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["slug"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Slug = types.StringValue(str)
			}
		}
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
