package project

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

var _ list.ListResource = &StructureProjectList{}

type StructureProjectList struct {
	client *Client
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
			"filters": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Filter parameters for querying Structure Project",
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

	l.client = NewClient(client)
}

type StructureProjectListModel struct {
	Filters *StructureProjectFiltersModel `tfsdk:"filters"`
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
	filters := common.BuildQueryFilters(config.Filters)

	// Call API
	listResult, err := l.client.ListStructureProject(ctx, filters)
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

			diags.Append(model.CopyFrom(ctx, apiResp)...)

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
