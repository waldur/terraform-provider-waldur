package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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

// StructureProjectApiResponse is the API response model.
type StructureProjectApiResponse struct {
	UUID *string `json:"uuid"`

	BackendId                            *string  `json:"backend_id" tfsdk:"backend_id"`
	Created                              *string  `json:"created" tfsdk:"created"`
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

// StructureProjectFiltersModel contains the filter parameters for querying.
type StructureProjectFiltersModel struct {
	BackendId               types.String `tfsdk:"backend_id"`
	CanAdmin                types.Bool   `tfsdk:"can_admin"`
	CanManage               types.Bool   `tfsdk:"can_manage"`
	ConcealFinishedProjects types.Bool   `tfsdk:"conceal_finished_projects"`
	Created                 types.String `tfsdk:"created"`
	Customer                types.String `tfsdk:"customer"`
	CustomerAbbreviation    types.String `tfsdk:"customer_abbreviation"`
	CustomerName            types.String `tfsdk:"customer_name"`
	CustomerNativeName      types.String `tfsdk:"customer_native_name"`
	Description             types.String `tfsdk:"description"`
	IncludeTerminated       types.Bool   `tfsdk:"include_terminated"`
	IsRemoved               types.Bool   `tfsdk:"is_removed"`
	Modified                types.String `tfsdk:"modified"`
	Name                    types.String `tfsdk:"name"`
	NameExact               types.String `tfsdk:"name_exact"`
	Query                   types.String `tfsdk:"query"`
	Slug                    types.String `tfsdk:"slug"`
}

// StructureProjectDataSourceModel describes the data source data model.
type StructureProjectDataSourceModel struct {
	UUID                                 types.String                  `tfsdk:"id"`
	Filters                              *StructureProjectFiltersModel `tfsdk:"filters"`
	BackendId                            types.String                  `tfsdk:"backend_id"`
	Created                              types.String                  `tfsdk:"created"`
	CustomerDisplayBillingInfoInProjects types.Bool                    `tfsdk:"customer_display_billing_info_in_projects"`
	CustomerSlug                         types.String                  `tfsdk:"customer_slug"`
	Description                          types.String                  `tfsdk:"description"`
	EndDate                              types.String                  `tfsdk:"end_date"`
	EndDateRequestedBy                   types.String                  `tfsdk:"end_date_requested_by"`
	GracePeriodDays                      types.Int64                   `tfsdk:"grace_period_days"`
	Image                                types.String                  `tfsdk:"image"`
	IsIndustry                           types.Bool                    `tfsdk:"is_industry"`
	IsRemoved                            types.Bool                    `tfsdk:"is_removed"`
	Kind                                 types.String                  `tfsdk:"kind"`
	MaxServiceAccounts                   types.Int64                   `tfsdk:"max_service_accounts"`
	Name                                 types.String                  `tfsdk:"name"`
	OecdFos2007Code                      types.String                  `tfsdk:"oecd_fos_2007_code"`
	OecdFos2007Label                     types.String                  `tfsdk:"oecd_fos_2007_label"`
	ProjectCredit                        types.Float64                 `tfsdk:"project_credit"`
	ResourcesCount                       types.Int64                   `tfsdk:"resources_count"`
	Slug                                 types.String                  `tfsdk:"slug"`
	StaffNotes                           types.String                  `tfsdk:"staff_notes"`
	StartDate                            types.String                  `tfsdk:"start_date"`
	Type                                 types.String                  `tfsdk:"type"`
	TypeName                             types.String                  `tfsdk:"type_name"`
	TypeUuid                             types.String                  `tfsdk:"type_uuid"`
	Url                                  types.String                  `tfsdk:"url"`
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
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Filter parameters for querying Structure Project",
				Attributes: map[string]schema.Attribute{
					"backend_id": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "ID of the backend",
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
						MarkdownDescription: "Customer abbreviation",
					},
					"customer_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Customer name",
					},
					"customer_native_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Customer native name",
					},
					"description": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Description",
					},
					"include_terminated": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Include soft-deleted (terminated) projects. Only available to staff and support users, or users with organizational roles who can see their terminated projects.",
					},
					"is_removed": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Is removed",
					},
					"modified": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Modified after",
					},
					"name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name",
					},
					"name_exact": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name (exact)",
					},
					"query": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Filter by name, slug, UUID, backend ID or resource effective ID",
					},
					"slug": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Slug",
					},
				},
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of the backend",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Created",
			},
			"customer_display_billing_info_in_projects": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Customer display billing info in projects",
			},
			"customer_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Customer slug",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project description (HTML content will be sanitized)",
			},
			"end_date": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project end date. Setting this field requires DELETE_PROJECT permission.",
			},
			"end_date_requested_by": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "End date requested by",
			},
			"grace_period_days": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of extra days after project end date before resources are terminated. Overrides customer-level setting.",
			},
			"image": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Image",
			},
			"is_industry": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Is industry",
			},
			"is_removed": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Is removed",
			},
			"kind": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Kind",
			},
			"max_service_accounts": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Maximum number of service accounts allowed",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the resource",
			},
			"oecd_fos_2007_code": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Oecd fos 2007 code",
			},
			"oecd_fos_2007_label": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Human-readable label for the OECD FOS 2007 classification code",
			},
			"project_credit": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "Project credit",
			},
			"resources_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of active resources in this project",
			},
			"slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
			},
			"staff_notes": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Internal notes visible only to staff and support users (HTML content will be sanitized)",
			},
			"start_date": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project start date. Cannot be edited after the start date has arrived.",
			},
			"type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Type",
			},
			"type_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the type",
			},
			"type_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the type",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Url",
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
		var apiResp StructureProjectApiResponse

		err := d.client.GetByUUID(ctx, "/api/projects/{uuid}/", data.UUID.ValueString(), &apiResp)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Structure Project",
				"An error occurred while reading the Structure Project by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, apiResp, &data)...)

	} else {
		// Filter by provided parameters
		var results []StructureProjectApiResponse

		filters := make(map[string]string)
		if data.Filters != nil {
			type filterDef struct {
				name string
				val  attr.Value
			}
			filterDefs := []filterDef{
				{"backend_id", data.Filters.BackendId},
				{"can_admin", data.Filters.CanAdmin},
				{"can_manage", data.Filters.CanManage},
				{"conceal_finished_projects", data.Filters.ConcealFinishedProjects},
				{"created", data.Filters.Created},
				{"customer", data.Filters.Customer},
				{"customer_abbreviation", data.Filters.CustomerAbbreviation},
				{"customer_name", data.Filters.CustomerName},
				{"customer_native_name", data.Filters.CustomerNativeName},
				{"description", data.Filters.Description},
				{"include_terminated", data.Filters.IncludeTerminated},
				{"is_removed", data.Filters.IsRemoved},
				{"modified", data.Filters.Modified},
				{"name", data.Filters.Name},
				{"name_exact", data.Filters.NameExact},
				{"query", data.Filters.Query},
				{"slug", data.Filters.Slug},
			}

			for _, fd := range filterDefs {
				if fd.val.IsNull() || fd.val.IsUnknown() {
					continue
				}
				switch v := fd.val.(type) {
				case types.String:
					filters[fd.name] = v.ValueString()
				case types.Int64:
					filters[fd.name] = fmt.Sprintf("%d", v.ValueInt64())
				case types.Bool:
					filters[fd.name] = fmt.Sprintf("%t", v.ValueBool())
				case types.Float64:
					filters[fd.name] = fmt.Sprintf("%f", v.ValueFloat64())
				}
			}
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
				"Unable to List Structure Project",
				"An error occurred while filtering Structure Project: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Structure Project Not Found",
				"No Structure Project found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Structure Projects Found",
				fmt.Sprintf("Found %d Structure Projects with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *StructureProjectDataSource) mapResponseToModel(ctx context.Context, apiResp StructureProjectApiResponse, model *StructureProjectDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Created = types.StringPointerValue(apiResp.Created)
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
