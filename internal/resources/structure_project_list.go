package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
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
				Description: "Customer abbreviation",
				Optional:    true,
			},
			"customer_name": schema.StringAttribute{
				Description: "Customer name",
				Optional:    true,
			},
			"customer_native_name": schema.StringAttribute{
				Description: "Customer native name",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description",
				Optional:    true,
			},
			"include_terminated": schema.BoolAttribute{
				Description: "Include soft-deleted (terminated) projects. Only available to staff and support users, or users with organizational roles who can see their terminated projects.",
				Optional:    true,
			},
			"is_removed": schema.BoolAttribute{
				Description: "Is removed",
				Optional:    true,
			},
			"modified": schema.StringAttribute{
				Description: "Modified after",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name",
				Optional:    true,
			},
			"name_exact": schema.StringAttribute{
				Description: "Name (exact)",
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
				Description: "Slug",
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
	var listResult []StructureProjectApiResponse
	err := l.client.ListWithFilter(ctx, "/api/projects/", filters, &listResult)
	if err != nil {
		// Return error diagnostics
		resp.AddError("Failed to list resources", err.Error())
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Stream results
	stream.Results = func(push func(list.ListResult) bool) {
		for _, apiResp := range listResult {
			result := req.NewListResult(ctx)

			// Map item to model
			var data StructureProjectResourceModel
			model := &data

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

			// Set the resource state
			// For ListResource, we generally return the "Resource" state matching the main resource schema.
			// This allows `terraform plan` to correlate if using `terraform query`.

			setDiags := result.Resource.Set(ctx, &data)
			diags.Append(setDiags...)
			result.Diagnostics.Append(diags...)

			// Set Identity if possible (usually UUID)
			if !data.UUID.IsNull() && !data.UUID.IsUnknown() {
				result.Diagnostics.Append(result.Identity.Set(ctx, data.UUID.ValueString())...)
			}

			if !push(result) {
				return
			}
		}
	}
}
