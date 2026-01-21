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

var _ list.ListResource = &OpenstackNetworkRbacPolicyList{}

type OpenstackNetworkRbacPolicyList struct {
	client *client.Client
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

	l.client = client
}

type OpenstackNetworkRbacPolicyListModel struct {
	Network          types.String `tfsdk:"network"`
	NetworkUuid      types.String `tfsdk:"network_uuid"`
	Page             types.Int64  `tfsdk:"page"`
	PageSize         types.Int64  `tfsdk:"page_size"`
	PolicyType       types.String `tfsdk:"policy_type"`
	TargetTenant     types.String `tfsdk:"target_tenant"`
	TargetTenantUuid types.String `tfsdk:"target_tenant_uuid"`
	Tenant           types.String `tfsdk:"tenant"`
	TenantUuid       types.String `tfsdk:"tenant_uuid"`
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
	filters := make(map[string]string)
	if !config.Network.IsNull() && !config.Network.IsUnknown() {
		filters["network"] = config.Network.ValueString()
	}
	if !config.NetworkUuid.IsNull() && !config.NetworkUuid.IsUnknown() {
		filters["network_uuid"] = config.NetworkUuid.ValueString()
	}
	if !config.Page.IsNull() && !config.Page.IsUnknown() {
		filters["page"] = fmt.Sprintf("%d", config.Page.ValueInt64())
	}
	if !config.PageSize.IsNull() && !config.PageSize.IsUnknown() {
		filters["page_size"] = fmt.Sprintf("%d", config.PageSize.ValueInt64())
	}
	if !config.PolicyType.IsNull() && !config.PolicyType.IsUnknown() {
		filters["policy_type"] = config.PolicyType.ValueString()
	}
	if !config.TargetTenant.IsNull() && !config.TargetTenant.IsUnknown() {
		filters["target_tenant"] = config.TargetTenant.ValueString()
	}
	if !config.TargetTenantUuid.IsNull() && !config.TargetTenantUuid.IsUnknown() {
		filters["target_tenant_uuid"] = config.TargetTenantUuid.ValueString()
	}
	if !config.Tenant.IsNull() && !config.Tenant.IsUnknown() {
		filters["tenant"] = config.Tenant.ValueString()
	}
	if !config.TenantUuid.IsNull() && !config.TenantUuid.IsUnknown() {
		filters["tenant_uuid"] = config.TenantUuid.ValueString()
	}

	// Call API
	var listResult []OpenstackNetworkRbacPolicyApiResponse
	err := l.client.ListWithFilter(ctx, "/api/openstack-network-rbac-policies/", filters, &listResult)
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

			model.UUID = types.StringPointerValue(apiResp.UUID)
			model.BackendId = types.StringPointerValue(apiResp.BackendId)
			model.Created = types.StringPointerValue(apiResp.Created)
			model.Network = types.StringPointerValue(apiResp.Network)
			model.NetworkName = types.StringPointerValue(apiResp.NetworkName)
			model.PolicyType = types.StringPointerValue(apiResp.PolicyType)
			model.TargetTenant = types.StringPointerValue(apiResp.TargetTenant)
			model.TargetTenantName = types.StringPointerValue(apiResp.TargetTenantName)
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
