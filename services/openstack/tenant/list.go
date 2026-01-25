package tenant

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

var _ list.ListResource = &OpenstackTenantList{}

type OpenstackTenantList struct {
	client *Client
}

func NewOpenstackTenantList() list.ListResource {
	return &OpenstackTenantList{}
}

func (l *OpenstackTenantList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_tenant"
}

func (l *OpenstackTenantList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"filters": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Filter parameters for querying Openstack Tenant",
				Attributes: map[string]schema.Attribute{
					"backend_id": schema.StringAttribute{
						Description: "Backend ID",
						Optional:    true,
					},
					"can_manage": schema.BoolAttribute{
						Description: "Can manage",
						Optional:    true,
					},
					"customer": schema.StringAttribute{
						Description: "Customer UUID",
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
					"customer_uuid": schema.StringAttribute{
						Description: "Customer UUID",
						Optional:    true,
					},
					"description": schema.StringAttribute{
						Description: "Description",
						Optional:    true,
					},
					"external_ip": schema.StringAttribute{
						Description: "External IP",
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
					"project": schema.StringAttribute{
						Description: "Project UUID",
						Optional:    true,
					},
					"project_name": schema.StringAttribute{
						Description: "Project name",
						Optional:    true,
					},
					"project_uuid": schema.StringAttribute{
						Description: "Project UUID",
						Optional:    true,
					},
					"service_settings_name": schema.StringAttribute{
						Description: "Service settings name",
						Optional:    true,
					},
					"service_settings_uuid": schema.StringAttribute{
						Description: "Service settings UUID",
						Optional:    true,
					},
					"uuid": schema.StringAttribute{
						Description: "UUID",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (l *OpenstackTenantList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type OpenstackTenantListModel struct {
	Filters *OpenstackTenantFiltersModel `tfsdk:"filters"`
}

func (l *OpenstackTenantList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config OpenstackTenantListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := common.BuildQueryFilters(config.Filters)

	// Call API
	listResult, err := l.client.ListOpenstackTenant(ctx, filters)
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
			var data OpenstackTenantResourceModel
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
