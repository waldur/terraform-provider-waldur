package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-provider-waldur/internal/client"
)

var _ list.ListResource = &StructureProjectList{}

type StructureProjectList struct {
	client *client.Client
}

func NewStructureProjectList() list.ListResource {
	return &StructureProjectList{}
}

func (l *StructureProjectList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_structure_project"
}

func (l *StructureProjectList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"backend_id": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"can_admin": schema.BoolAttribute{
				Description: "Return a list of projects where current user is admin.",
				Optional:    true,
			},
			"can_manage": schema.BoolAttribute{
				Description: "Return a list of projects where current user is manager or a customer owner.",
				Optional:    true,
			},
			"conceal_finished_projects": schema.BoolAttribute{
				Description: "Conceal finished projects",
				Optional:    true,
			},
			"created": schema.StringAttribute{
				Description: "Created after",
				Optional:    true,
			},
			"customer_abbreviation": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"customer_name": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"customer_native_name": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"include_terminated": schema.BoolAttribute{
				Description: "Include soft-deleted (terminated) projects. Only available to staff and support users, or users with organizational roles who can see their terminated projects.",
				Optional:    true,
			},
			"is_removed": schema.BoolAttribute{
				Description: "",
				Optional:    true,
			},
			"modified": schema.StringAttribute{
				Description: "Modified after",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"name_exact": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"page": schema.Int64Attribute{
				Description: "A page number within the paginated result set.",
				Optional:    true,
			},
			"page_size": schema.Int64Attribute{
				Description: "Number of results to return per page.",
				Optional:    true,
			},
			"query": schema.StringAttribute{
				Description: "Filter by name, slug, UUID, backend ID or resource effective ID",
				Optional:    true,
			},
			"slug": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
		},
	}
}

func (l *StructureProjectList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected List Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T", req.ProviderData),
		)
		return
	}

	l.client = client
}

type StructureProjectListModel struct {
	BackendId               types.String `tfsdk:"backend_id"`
	CanAdmin                types.Bool   `tfsdk:"can_admin"`
	CanManage               types.Bool   `tfsdk:"can_manage"`
	ConcealFinishedProjects types.Bool   `tfsdk:"conceal_finished_projects"`
	Created                 types.String `tfsdk:"created"`
	CustomerAbbreviation    types.String `tfsdk:"customer_abbreviation"`
	CustomerName            types.String `tfsdk:"customer_name"`
	CustomerNativeName      types.String `tfsdk:"customer_native_name"`
	Description             types.String `tfsdk:"description"`
	IncludeTerminated       types.Bool   `tfsdk:"include_terminated"`
	IsRemoved               types.Bool   `tfsdk:"is_removed"`
	Modified                types.String `tfsdk:"modified"`
	Name                    types.String `tfsdk:"name"`
	NameExact               types.String `tfsdk:"name_exact"`
	Page                    types.Int64  `tfsdk:"page"`
	PageSize                types.Int64  `tfsdk:"page_size"`
	Query                   types.String `tfsdk:"query"`
	Slug                    types.String `tfsdk:"slug"`
}

func (l *StructureProjectList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config StructureProjectListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := make(map[string]string)
	if !config.BackendId.IsNull() && !config.BackendId.IsUnknown() {
		filters["backend_id"] = config.BackendId.ValueString()
	}
	if !config.CanAdmin.IsNull() && !config.CanAdmin.IsUnknown() {
		filters["can_admin"] = fmt.Sprintf("%t", config.CanAdmin.ValueBool())
	}
	if !config.CanManage.IsNull() && !config.CanManage.IsUnknown() {
		filters["can_manage"] = fmt.Sprintf("%t", config.CanManage.ValueBool())
	}
	if !config.ConcealFinishedProjects.IsNull() && !config.ConcealFinishedProjects.IsUnknown() {
		filters["conceal_finished_projects"] = fmt.Sprintf("%t", config.ConcealFinishedProjects.ValueBool())
	}
	if !config.Created.IsNull() && !config.Created.IsUnknown() {
		filters["created"] = config.Created.ValueString()
	}
	if !config.CustomerAbbreviation.IsNull() && !config.CustomerAbbreviation.IsUnknown() {
		filters["customer_abbreviation"] = config.CustomerAbbreviation.ValueString()
	}
	if !config.CustomerName.IsNull() && !config.CustomerName.IsUnknown() {
		filters["customer_name"] = config.CustomerName.ValueString()
	}
	if !config.CustomerNativeName.IsNull() && !config.CustomerNativeName.IsUnknown() {
		filters["customer_native_name"] = config.CustomerNativeName.ValueString()
	}
	if !config.Description.IsNull() && !config.Description.IsUnknown() {
		filters["description"] = config.Description.ValueString()
	}
	if !config.IncludeTerminated.IsNull() && !config.IncludeTerminated.IsUnknown() {
		filters["include_terminated"] = fmt.Sprintf("%t", config.IncludeTerminated.ValueBool())
	}
	if !config.IsRemoved.IsNull() && !config.IsRemoved.IsUnknown() {
		filters["is_removed"] = fmt.Sprintf("%t", config.IsRemoved.ValueBool())
	}
	if !config.Modified.IsNull() && !config.Modified.IsUnknown() {
		filters["modified"] = config.Modified.ValueString()
	}
	if !config.Name.IsNull() && !config.Name.IsUnknown() {
		filters["name"] = config.Name.ValueString()
	}
	if !config.NameExact.IsNull() && !config.NameExact.IsUnknown() {
		filters["name_exact"] = config.NameExact.ValueString()
	}
	if !config.Page.IsNull() && !config.Page.IsUnknown() {
		filters["page"] = fmt.Sprintf("%d", config.Page.ValueInt64())
	}
	if !config.PageSize.IsNull() && !config.PageSize.IsUnknown() {
		filters["page_size"] = fmt.Sprintf("%d", config.PageSize.ValueInt64())
	}
	if !config.Query.IsNull() && !config.Query.IsUnknown() {
		filters["query"] = config.Query.ValueString()
	}
	if !config.Slug.IsNull() && !config.Slug.IsUnknown() {
		filters["slug"] = config.Slug.ValueString()
	}

	// Call API
	var listResult []map[string]interface{}
	err := l.client.ListWithFilter(ctx, "/api/projects/", filters, &listResult)
	if err != nil {
		// Return error diagnostics
		resp.AddError("Failed to list resources", err.Error())
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Stream results
	stream.Results = func(push func(list.ListResult) bool) {
		for _, item := range listResult {
			result := req.NewListResult(ctx)

			// Map item to model
			var data StructureProjectResourceModel

			// Initialize lists and objects to empty/null if needed (or rely on map logic)

			sourceMap := item

			// Reuse the mapResponseFields logic (embedded here or via shared template)
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

			// Set the resource state
			// For ListResource, we generally return the "Resource" state matching the main resource schema.
			// This allows `terraform plan` to correlate if using `terraform query`.

			diags := result.Resource.Set(ctx, &data)
			result.Diagnostics.Append(diags...)

			// Set Identity if possible (usually UUID)
			if !data.UUID.IsNull() && !data.UUID.IsUnknown() {
				// Identity value must match what the resource uses for Import?
				// Typically implicit. For now just setting Resource is key.
				// result.Identity.Set(ctx, data.UUID.ValueString())
				// The doc says: "An error is returned if a list result in the stream contains a null identity"
				result.Diagnostics.Append(result.Identity.Set(ctx, data.UUID.ValueString())...)
			} else {
				// Try to fallback to "uuid" from map if model failed
				if uuid, ok := item["uuid"].(string); ok {
					result.Diagnostics.Append(result.Identity.Set(ctx, uuid)...)
				}
			}

			if !push(result) {
				return
			}
		}
	}
}
