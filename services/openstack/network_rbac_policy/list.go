package network_rbac_policy

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

var _ list.ListResource = &OpenstackNetworkRbacPolicyList{}

type OpenstackNetworkRbacPolicyList struct {
	client *Client
}

func NewOpenstackNetworkRbacPolicyList() list.ListResource {
	return &OpenstackNetworkRbacPolicyList{}
}

func (l *OpenstackNetworkRbacPolicyList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_network_rbac_policy"
}

func (l *OpenstackNetworkRbacPolicyList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"filters": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Filter parameters for querying Openstack Network Rbac Policy",
				Attributes: map[string]schema.Attribute{
					"network": schema.StringAttribute{
						Description: "Network URL",
						Optional:    true,
					},
					"network_uuid": schema.StringAttribute{
						Description: "Network UUID",
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
					"policy_type": schema.StringAttribute{
						Description: "Type of access granted - either shared access or external network access",
						Optional:    true,
					},
					"target_tenant": schema.StringAttribute{
						Description: "Target tenant URL",
						Optional:    true,
					},
					"target_tenant_uuid": schema.StringAttribute{
						Description: "Target tenant UUID",
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

func (l *OpenstackNetworkRbacPolicyList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type OpenstackNetworkRbacPolicyListModel struct {
	Filters *OpenstackNetworkRbacPolicyFiltersModel `tfsdk:"filters"`
}

func (l *OpenstackNetworkRbacPolicyList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config OpenstackNetworkRbacPolicyListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := common.BuildQueryFilters(config.Filters)

	// Call API
	listResult, err := l.client.ListOpenstackNetworkRbacPolicy(ctx, filters)
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
			var data OpenstackNetworkRbacPolicyResourceModel
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
