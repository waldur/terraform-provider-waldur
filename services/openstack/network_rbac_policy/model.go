package network_rbac_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackNetworkRbacPolicyFiltersModel struct {
	Network          types.String `tfsdk:"network"`
	NetworkUuid      types.String `tfsdk:"network_uuid"`
	PolicyType       types.String `tfsdk:"policy_type"`
	TargetTenant     types.String `tfsdk:"target_tenant"`
	TargetTenantUuid types.String `tfsdk:"target_tenant_uuid"`
	Tenant           types.String `tfsdk:"tenant"`
	TenantUuid       types.String `tfsdk:"tenant_uuid"`
}

func (m *OpenstackNetworkRbacPolicyFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Openstack Network Rbac Policy",
		Attributes: map[string]schema.Attribute{
			"network": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Network URL",
			},
			"network_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Network UUID",
			},
			"policy_type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Type of access granted - either shared access or external network access Allowed values: `access_as_external`, `access_as_shared`.",
				Validators: []validator.String{
					stringvalidator.OneOf("access_as_external", "access_as_shared"),
				},
			},
			"target_tenant": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Target tenant URL",
			},
			"target_tenant_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Target tenant UUID",
			},
			"tenant": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Tenant URL",
			},
			"tenant_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Tenant UUID",
			},
		},
	}
}

type OpenstackNetworkRbacPolicyModel struct {
	UUID             types.String `tfsdk:"id"`
	BackendId        types.String `tfsdk:"backend_id"`
	Network          types.String `tfsdk:"network"`
	NetworkName      types.String `tfsdk:"network_name"`
	PolicyType       types.String `tfsdk:"policy_type"`
	TargetTenant     types.String `tfsdk:"target_tenant"`
	TargetTenantName types.String `tfsdk:"target_tenant_name"`
	Url              types.String `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackNetworkRbacPolicyModel) CopyFrom(ctx context.Context, apiResp OpenstackNetworkRbacPolicyResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.BackendId = common.StringPointerValue(apiResp.BackendId)
	model.Network = common.StringPointerValue(apiResp.Network)
	model.NetworkName = common.StringPointerValue(apiResp.NetworkName)
	model.PolicyType = common.StringPointerValue(apiResp.PolicyType)
	model.TargetTenant = common.StringPointerValue(apiResp.TargetTenant)
	model.TargetTenantName = common.StringPointerValue(apiResp.TargetTenantName)
	model.Url = common.StringPointerValue(apiResp.Url)

	return diags
}
