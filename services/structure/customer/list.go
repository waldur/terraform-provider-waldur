package customer

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

var _ list.ListResource = &StructureCustomerList{}

type StructureCustomerList struct {
	client *Client
}

func NewStructureCustomerList() list.ListResource {
	return &StructureCustomerList{}
}

func (l *StructureCustomerList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_structure_customer"
}

func (l *StructureCustomerList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"filters": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Filter parameters for querying Structure Customer",
				Attributes: map[string]schema.Attribute{
					"abbreviation": schema.StringAttribute{
						Description: "Abbreviation",
						Optional:    true,
					},
					"agreement_number": schema.StringAttribute{
						Description: "",
						Optional:    true,
					},
					"archived": schema.BoolAttribute{
						Description: "",
						Optional:    true,
					},
					"backend_id": schema.StringAttribute{
						Description: "",
						Optional:    true,
					},
					"contact_details": schema.StringAttribute{
						Description: "Contact details",
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
					"native_name": schema.StringAttribute{
						Description: "Native name",
						Optional:    true,
					},
					"o": schema.StringAttribute{
						Description: "Which field to use when ordering the results.",
						Optional:    true,
					},
					"organization_group_name": schema.StringAttribute{
						Description: "Organization group name",
						Optional:    true,
					},
					"owned_by_current_user": schema.BoolAttribute{
						Description: "Return a list of customers where current user is owner.",
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
						Description: "Filter by name, native name, abbreviation, domain, UUID, registration code or agreement number",
						Optional:    true,
					},
					"registration_code": schema.StringAttribute{
						Description: "",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (l *StructureCustomerList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type StructureCustomerListModel struct {
	Filters *StructureCustomerFiltersModel `tfsdk:"filters"`
}

func (l *StructureCustomerList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config StructureCustomerListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := common.BuildQueryFilters(config.Filters)

	// Call API
	listResult, err := l.client.ListStructureCustomer(ctx, filters)
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
			var data StructureCustomerResourceModel
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
