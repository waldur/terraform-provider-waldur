package network

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

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

type OpenstackNetworkModel struct {
	UUID                        types.String      `tfsdk:"id"`
	AccessUrl                   types.String      `tfsdk:"access_url"`
	BackendId                   types.String      `tfsdk:"backend_id"`
	Created                     timetypes.RFC3339 `tfsdk:"created"`
	Customer                    types.String      `tfsdk:"customer"`
	CustomerAbbreviation        types.String      `tfsdk:"customer_abbreviation"`
	CustomerName                types.String      `tfsdk:"customer_name"`
	CustomerNativeName          types.String      `tfsdk:"customer_native_name"`
	CustomerUuid                types.String      `tfsdk:"customer_uuid"`
	Description                 types.String      `tfsdk:"description"`
	ErrorMessage                types.String      `tfsdk:"error_message"`
	ErrorTraceback              types.String      `tfsdk:"error_traceback"`
	IsExternal                  types.Bool        `tfsdk:"is_external"`
	IsLimitBased                types.Bool        `tfsdk:"is_limit_based"`
	IsUsageBased                types.Bool        `tfsdk:"is_usage_based"`
	MarketplaceCategoryName     types.String      `tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     types.String      `tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     types.String      `tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     types.String      `tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         types.String      `tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    types.String      `tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     types.String      `tfsdk:"marketplace_resource_uuid"`
	Modified                    timetypes.RFC3339 `tfsdk:"modified"`
	Mtu                         types.Int64       `tfsdk:"mtu"`
	Name                        types.String      `tfsdk:"name"`
	Project                     types.String      `tfsdk:"project"`
	ProjectName                 types.String      `tfsdk:"project_name"`
	ProjectUuid                 types.String      `tfsdk:"project_uuid"`
	RbacPolicies                types.List        `tfsdk:"rbac_policies"`
	ResourceType                types.String      `tfsdk:"resource_type"`
	SegmentationId              types.Int64       `tfsdk:"segmentation_id"`
	ServiceName                 types.String      `tfsdk:"service_name"`
	ServiceSettings             types.String      `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage types.String      `tfsdk:"service_settings_error_message"`
	ServiceSettingsState        types.String      `tfsdk:"service_settings_state"`
	ServiceSettingsUuid         types.String      `tfsdk:"service_settings_uuid"`
	State                       types.String      `tfsdk:"state"`
	Subnets                     types.List        `tfsdk:"subnets"`
	Tenant                      types.String      `tfsdk:"tenant"`
	TenantName                  types.String      `tfsdk:"tenant_name"`
	TenantUuid                  types.String      `tfsdk:"tenant_uuid"`
	Type                        types.String      `tfsdk:"type"`
	Url                         types.String      `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackNetworkModel) CopyFrom(ctx context.Context, apiResp OpenstackNetworkResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	valCreated, diagsCreated := timetypes.NewRFC3339PointerValue(apiResp.Created)
	diags.Append(diagsCreated...)
	model.Created = valCreated
	model.Description = types.StringPointerValue(apiResp.Description)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.IsExternal = types.BoolPointerValue(apiResp.IsExternal)
	valModified, diagsModified := timetypes.NewRFC3339PointerValue(apiResp.Modified)
	diags.Append(diagsModified...)
	model.Modified = valModified
	model.Mtu = types.Int64PointerValue(apiResp.Mtu)
	model.Name = types.StringPointerValue(apiResp.Name)

	{
		listValRbacPolicies, listDiagsRbacPolicies := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"backend_id":         types.StringType,
			"created":            types.StringType,
			"network":            types.StringType,
			"network_name":       types.StringType,
			"policy_type":        types.StringType,
			"target_tenant":      types.StringType,
			"target_tenant_name": types.StringType,
			"url":                types.StringType,
		}}, apiResp.RbacPolicies)
		diags.Append(listDiagsRbacPolicies...)
		model.RbacPolicies = listValRbacPolicies
	}
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.SegmentationId = types.Int64PointerValue(apiResp.SegmentationId)
	model.State = types.StringPointerValue(apiResp.State)

	{
		listValSubnets, listDiagsSubnets := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"end":   types.StringType,
				"start": types.StringType,
			}}},
			"cidr":        types.StringType,
			"description": types.StringType,
			"enable_dhcp": types.BoolType,
			"gateway_ip":  types.StringType,
			"ip_version":  types.Int64Type,
			"name":        types.StringType,
		}}, apiResp.Subnets)
		diags.Append(listDiagsSubnets...)
		model.Subnets = listValSubnets
	}
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantName = types.StringPointerValue(apiResp.TenantName)
	model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
	model.Type = types.StringPointerValue(apiResp.Type)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
