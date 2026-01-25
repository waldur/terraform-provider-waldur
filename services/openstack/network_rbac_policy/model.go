package network_rbac_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

type OpenstackNetworkRbacPolicyModel struct {
	UUID             types.String      `tfsdk:"id"`
	BackendId        types.String      `tfsdk:"backend_id"`
	Created          timetypes.RFC3339 `tfsdk:"created"`
	Network          types.String      `tfsdk:"network"`
	NetworkName      types.String      `tfsdk:"network_name"`
	PolicyType       types.String      `tfsdk:"policy_type"`
	TargetTenant     types.String      `tfsdk:"target_tenant"`
	TargetTenantName types.String      `tfsdk:"target_tenant_name"`
	Url              types.String      `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackNetworkRbacPolicyModel) CopyFrom(ctx context.Context, apiResp OpenstackNetworkRbacPolicyResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	valCreated, diagsCreated := timetypes.NewRFC3339PointerValue(apiResp.Created)
	diags.Append(diagsCreated...)
	model.Created = valCreated
	model.Network = types.StringPointerValue(apiResp.Network)
	model.NetworkName = types.StringPointerValue(apiResp.NetworkName)
	model.PolicyType = types.StringPointerValue(apiResp.PolicyType)
	model.TargetTenant = types.StringPointerValue(apiResp.TargetTenant)
	model.TargetTenantName = types.StringPointerValue(apiResp.TargetTenantName)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
