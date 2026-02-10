package network

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

func NetworkRBACPolicyType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"backend_id":         types.StringType,
		"network":            types.StringType,
		"network_name":       types.StringType,
		"policy_type":        types.StringType,
		"target_tenant":      types.StringType,
		"target_tenant_name": types.StringType,
		"url":                types.StringType,
		"uuid":               types.StringType,
	}}
}
func OpenStackNestedSubNetType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"allocation_pools": types.ListType{ElemType: OpenStackSubNetAllocationPoolType()},
		"cidr":             types.StringType,
		"description":      types.StringType,
		"enable_dhcp":      types.BoolType,
		"gateway_ip":       types.StringType,
		"ip_version":       types.Int64Type,
		"name":             types.StringType,
		"uuid":             types.StringType,
	}}
}
func OpenStackSubNetAllocationPoolType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"end":   types.StringType,
		"start": types.StringType,
	}}
}

type OpenstackNetworkFiltersModel struct {
	BackendId            types.String `tfsdk:"backend_id"`
	CanManage            types.Bool   `tfsdk:"can_manage"`
	Customer             types.String `tfsdk:"customer"`
	CustomerAbbreviation types.String `tfsdk:"customer_abbreviation"`
	CustomerName         types.String `tfsdk:"customer_name"`
	CustomerNativeName   types.String `tfsdk:"customer_native_name"`
	CustomerUuid         types.String `tfsdk:"customer_uuid"`
	Description          types.String `tfsdk:"description"`
	DirectOnly           types.Bool   `tfsdk:"direct_only"`
	ExternalIp           types.String `tfsdk:"external_ip"`
	IsExternal           types.Bool   `tfsdk:"is_external"`
	Name                 types.String `tfsdk:"name"`
	NameExact            types.String `tfsdk:"name_exact"`
	Project              types.String `tfsdk:"project"`
	ProjectName          types.String `tfsdk:"project_name"`
	ProjectUuid          types.String `tfsdk:"project_uuid"`
	RbacOnly             types.Bool   `tfsdk:"rbac_only"`
	ServiceSettingsName  types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid  types.String `tfsdk:"service_settings_uuid"`
	Tenant               types.String `tfsdk:"tenant"`
	TenantUuid           types.String `tfsdk:"tenant_uuid"`
	Type                 types.String `tfsdk:"type"`
	Uuid                 types.String `tfsdk:"uuid"`
}

func (m *OpenstackNetworkFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Openstack Network",
		Attributes: map[string]schema.Attribute{
			"backend_id": schema.StringAttribute{
				Optional: true,
			},
			"can_manage": schema.BoolAttribute{
				Optional: true,
			},
			"customer": schema.StringAttribute{
				Optional: true,
			},
			"customer_abbreviation": schema.StringAttribute{
				Optional: true,
			},
			"customer_name": schema.StringAttribute{
				Optional: true,
			},
			"customer_native_name": schema.StringAttribute{
				Optional: true,
			},
			"customer_uuid": schema.StringAttribute{
				Optional: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"direct_only": schema.BoolAttribute{
				Optional: true,
			},
			"external_ip": schema.StringAttribute{
				Optional: true,
			},
			"is_external": schema.BoolAttribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"name_exact": schema.StringAttribute{
				Optional: true,
			},
			"project": schema.StringAttribute{
				Optional: true,
			},
			"project_name": schema.StringAttribute{
				Optional: true,
			},
			"project_uuid": schema.StringAttribute{
				Optional: true,
			},
			"rbac_only": schema.BoolAttribute{
				Optional: true,
			},
			"service_settings_name": schema.StringAttribute{
				Optional: true,
			},
			"service_settings_uuid": schema.StringAttribute{
				Optional: true,
			},
			"tenant": schema.StringAttribute{
				Optional: true,
			},
			"tenant_uuid": schema.StringAttribute{
				Optional: true,
			},
			"type": schema.StringAttribute{
				Optional: true,
			},
			"uuid": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

type OpenstackNetworkModel struct {
	UUID                    types.String `tfsdk:"id"`
	BackendId               types.String `tfsdk:"backend_id"`
	Customer                types.String `tfsdk:"customer"`
	Description             types.String `tfsdk:"description"`
	ErrorMessage            types.String `tfsdk:"error_message"`
	ErrorTraceback          types.String `tfsdk:"error_traceback"`
	IsExternal              types.Bool   `tfsdk:"is_external"`
	MarketplaceResourceUuid types.String `tfsdk:"marketplace_resource_uuid"`
	Mtu                     types.Int64  `tfsdk:"mtu"`
	Name                    types.String `tfsdk:"name"`
	Project                 types.String `tfsdk:"project"`
	RbacPolicies            types.List   `tfsdk:"rbac_policies"`
	ResourceType            types.String `tfsdk:"resource_type"`
	SegmentationId          types.Int64  `tfsdk:"segmentation_id"`
	State                   types.String `tfsdk:"state"`
	Subnets                 types.List   `tfsdk:"subnets"`
	Tenant                  types.String `tfsdk:"tenant"`
	TenantName              types.String `tfsdk:"tenant_name"`
	TenantUuid              types.String `tfsdk:"tenant_uuid"`
	Type                    types.String `tfsdk:"type"`
	Url                     types.String `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackNetworkModel) CopyFrom(ctx context.Context, apiResp OpenstackNetworkResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.BackendId = common.StringPointerValue(apiResp.BackendId)
	model.Customer = common.StringPointerValue(apiResp.Customer)
	model.Description = common.StringPointerValue(apiResp.Description)
	model.ErrorMessage = common.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = common.StringPointerValue(apiResp.ErrorTraceback)
	model.IsExternal = types.BoolPointerValue(apiResp.IsExternal)
	model.MarketplaceResourceUuid = common.StringPointerValue(apiResp.MarketplaceResourceUuid)
	model.Mtu = types.Int64PointerValue(apiResp.Mtu)
	model.Name = common.StringPointerValue(apiResp.Name)
	model.Project = common.StringPointerValue(apiResp.Project)

	if apiResp.RbacPolicies != nil && len(*apiResp.RbacPolicies) > 0 {
		listValRbacPolicies, listDiagsRbacPolicies := types.ListValueFrom(ctx, NetworkRBACPolicyType(), apiResp.RbacPolicies)
		diags.Append(listDiagsRbacPolicies...)
		model.RbacPolicies = listValRbacPolicies
	} else {
		model.RbacPolicies = types.ListNull(NetworkRBACPolicyType())
	}
	model.ResourceType = common.StringPointerValue(apiResp.ResourceType)
	model.SegmentationId = types.Int64PointerValue(apiResp.SegmentationId)
	model.State = common.StringPointerValue(apiResp.State)

	if apiResp.Subnets != nil && len(*apiResp.Subnets) > 0 {
		listValSubnets, listDiagsSubnets := types.ListValueFrom(ctx, OpenStackNestedSubNetType(), apiResp.Subnets)
		diags.Append(listDiagsSubnets...)
		model.Subnets = listValSubnets
	} else {
		model.Subnets = types.ListNull(OpenStackNestedSubNetType())
	}
	model.Tenant = common.StringPointerValue(apiResp.Tenant)
	model.TenantName = common.StringPointerValue(apiResp.TenantName)
	model.TenantUuid = common.StringPointerValue(apiResp.TenantUuid)
	model.Type = common.StringPointerValue(apiResp.Type)
	model.Url = common.StringPointerValue(apiResp.Url)

	return diags
}
