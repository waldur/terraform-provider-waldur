package port

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OpenstackPortFiltersModel struct {
	AdminStateUp       types.Bool   `tfsdk:"admin_state_up"`
	BackendId          types.String `tfsdk:"backend_id"`
	DeviceId           types.String `tfsdk:"device_id"`
	DeviceOwner        types.String `tfsdk:"device_owner"`
	ExcludeSubnetUuids types.String `tfsdk:"exclude_subnet_uuids"`
	FixedIps           types.String `tfsdk:"fixed_ips"`
	HasDeviceOwner     types.Bool   `tfsdk:"has_device_owner"`
	MacAddress         types.String `tfsdk:"mac_address"`
	Name               types.String `tfsdk:"name"`
	NameExact          types.String `tfsdk:"name_exact"`
	NetworkName        types.String `tfsdk:"network_name"`
	NetworkUuid        types.String `tfsdk:"network_uuid"`
	Query              types.String `tfsdk:"query"`
	Status             types.String `tfsdk:"status"`
	Tenant             types.String `tfsdk:"tenant"`
	TenantUuid         types.String `tfsdk:"tenant_uuid"`
}

type OpenstackPortModel struct {
	UUID                        types.String      `tfsdk:"id"`
	AccessUrl                   types.String      `tfsdk:"access_url"`
	AdminStateUp                types.Bool        `tfsdk:"admin_state_up"`
	AllowedAddressPairs         types.List        `tfsdk:"allowed_address_pairs"`
	BackendId                   types.String      `tfsdk:"backend_id"`
	Created                     timetypes.RFC3339 `tfsdk:"created"`
	Customer                    types.String      `tfsdk:"customer"`
	CustomerAbbreviation        types.String      `tfsdk:"customer_abbreviation"`
	CustomerName                types.String      `tfsdk:"customer_name"`
	CustomerNativeName          types.String      `tfsdk:"customer_native_name"`
	CustomerUuid                types.String      `tfsdk:"customer_uuid"`
	Description                 types.String      `tfsdk:"description"`
	DeviceId                    types.String      `tfsdk:"device_id"`
	DeviceOwner                 types.String      `tfsdk:"device_owner"`
	ErrorMessage                types.String      `tfsdk:"error_message"`
	ErrorTraceback              types.String      `tfsdk:"error_traceback"`
	FixedIps                    types.List        `tfsdk:"fixed_ips"`
	FloatingIps                 types.Set         `tfsdk:"floating_ips"`
	IsLimitBased                types.Bool        `tfsdk:"is_limit_based"`
	IsUsageBased                types.Bool        `tfsdk:"is_usage_based"`
	MacAddress                  types.String      `tfsdk:"mac_address"`
	MarketplaceCategoryName     types.String      `tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     types.String      `tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     types.String      `tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     types.String      `tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         types.String      `tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    types.String      `tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     types.String      `tfsdk:"marketplace_resource_uuid"`
	Modified                    timetypes.RFC3339 `tfsdk:"modified"`
	Name                        types.String      `tfsdk:"name"`
	Network                     types.String      `tfsdk:"network"`
	NetworkName                 types.String      `tfsdk:"network_name"`
	NetworkUuid                 types.String      `tfsdk:"network_uuid"`
	PortSecurityEnabled         types.Bool        `tfsdk:"port_security_enabled"`
	Project                     types.String      `tfsdk:"project"`
	ProjectName                 types.String      `tfsdk:"project_name"`
	ProjectUuid                 types.String      `tfsdk:"project_uuid"`
	ResourceType                types.String      `tfsdk:"resource_type"`
	SecurityGroups              types.Set         `tfsdk:"security_groups"`
	ServiceName                 types.String      `tfsdk:"service_name"`
	ServiceSettings             types.String      `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage types.String      `tfsdk:"service_settings_error_message"`
	ServiceSettingsState        types.String      `tfsdk:"service_settings_state"`
	ServiceSettingsUuid         types.String      `tfsdk:"service_settings_uuid"`
	State                       types.String      `tfsdk:"state"`
	Status                      types.String      `tfsdk:"status"`
	Tenant                      types.String      `tfsdk:"tenant"`
	TenantName                  types.String      `tfsdk:"tenant_name"`
	TenantUuid                  types.String      `tfsdk:"tenant_uuid"`
	Url                         types.String      `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackPortModel) CopyFrom(ctx context.Context, apiResp OpenstackPortResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.AdminStateUp = types.BoolPointerValue(apiResp.AdminStateUp)

	{
		listValAllowedAddressPairs, listDiagsAllowedAddressPairs := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"ip_address":  types.StringType,
			"mac_address": types.StringType,
		}}, apiResp.AllowedAddressPairs)
		diags.Append(listDiagsAllowedAddressPairs...)
		model.AllowedAddressPairs = listValAllowedAddressPairs
	}
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	valCreated, diagsCreated := timetypes.NewRFC3339PointerValue(apiResp.Created)
	diags.Append(diagsCreated...)
	model.Created = valCreated
	model.Description = types.StringPointerValue(apiResp.Description)
	model.DeviceId = types.StringPointerValue(apiResp.DeviceId)
	model.DeviceOwner = types.StringPointerValue(apiResp.DeviceOwner)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)

	{
		listValFixedIps, listDiagsFixedIps := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"ip_address": types.StringType,
			"subnet_id":  types.StringType,
		}}, apiResp.FixedIps)
		diags.Append(listDiagsFixedIps...)
		model.FixedIps = listValFixedIps
	}
	setValFloatingIps, setDiagsFloatingIps := types.SetValueFrom(ctx, types.StringType, apiResp.FloatingIps)
	model.FloatingIps = setValFloatingIps
	diags.Append(setDiagsFloatingIps...)
	model.MacAddress = types.StringPointerValue(apiResp.MacAddress)
	valModified, diagsModified := timetypes.NewRFC3339PointerValue(apiResp.Modified)
	diags.Append(diagsModified...)
	model.Modified = valModified
	model.Name = types.StringPointerValue(apiResp.Name)
	model.Network = types.StringPointerValue(apiResp.Network)
	model.NetworkName = types.StringPointerValue(apiResp.NetworkName)
	model.NetworkUuid = types.StringPointerValue(apiResp.NetworkUuid)
	model.PortSecurityEnabled = types.BoolPointerValue(apiResp.PortSecurityEnabled)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	{
		setValSecurityGroups, setDiagsSecurityGroups := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		}}, apiResp.SecurityGroups)
		diags.Append(setDiagsSecurityGroups...)
		model.SecurityGroups = setValSecurityGroups
	}
	model.State = types.StringPointerValue(apiResp.State)
	model.Status = types.StringPointerValue(apiResp.Status)
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantName = types.StringPointerValue(apiResp.TenantName)
	model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
