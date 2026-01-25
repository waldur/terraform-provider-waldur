package port

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

var _ list.ListResource = &OpenstackPortList{}

type OpenstackPortList struct {
	client *Client
}

func NewOpenstackPortList() list.ListResource {
	return &OpenstackPortList{}
}

func (l *OpenstackPortList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_port"
}

func (l *OpenstackPortList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"filters": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Filter parameters for querying Openstack Port",
				Attributes: map[string]schema.Attribute{
					"admin_state_up": schema.BoolAttribute{
						Description: "",
						Optional:    true,
					},
					"backend_id": schema.StringAttribute{
						Description: "",
						Optional:    true,
					},
					"device_id": schema.StringAttribute{
						Description: "",
						Optional:    true,
					},
					"device_owner": schema.StringAttribute{
						Description: "",
						Optional:    true,
					},
					"exclude_subnet_uuids": schema.StringAttribute{
						Description: "Exclude Subnet UUIDs (comma-separated)",
						Optional:    true,
					},
					"fixed_ips": schema.StringAttribute{
						Description: "Search by fixed IP",
						Optional:    true,
					},
					"has_device_owner": schema.BoolAttribute{
						Description: "Has device owner",
						Optional:    true,
					},
					"mac_address": schema.StringAttribute{
						Description: "",
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
					"network_name": schema.StringAttribute{
						Description: "Search by network name",
						Optional:    true,
					},
					"network_uuid": schema.StringAttribute{
						Description: "Search by network UUID",
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
						Description: "Search by name, MAC address or backend ID",
						Optional:    true,
					},
					"status": schema.StringAttribute{
						Description: "",
						Optional:    true,
					},
					"tenant": schema.StringAttribute{
						Description: "Tenant URL",
						Optional:    true,
					},
					"tenant_uuid": schema.StringAttribute{
						Description: "Tenant UUID",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (l *OpenstackPortList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type OpenstackPortListModel struct {
	Filters *OpenstackPortFiltersModel `tfsdk:"filters"`
}

func (l *OpenstackPortList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config OpenstackPortListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := common.BuildQueryFilters(config.Filters)

	// Call API
	listResult, err := l.client.ListOpenstackPort(ctx, filters)
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
			var data OpenstackPortResourceModel
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
