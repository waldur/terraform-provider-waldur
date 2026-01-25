package instance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// OpenstackInstanceFiltersModel contains the filter parameters for querying.
type OpenstackInstanceFiltersModel struct {
	AttachVolumeUuid     types.String `tfsdk:"attach_volume_uuid"`
	AvailabilityZoneName types.String `tfsdk:"availability_zone_name"`
	BackendId            types.String `tfsdk:"backend_id"`
	CanManage            types.Bool   `tfsdk:"can_manage"`
	Customer             types.String `tfsdk:"customer"`
	CustomerAbbreviation types.String `tfsdk:"customer_abbreviation"`
	CustomerName         types.String `tfsdk:"customer_name"`
	CustomerNativeName   types.String `tfsdk:"customer_native_name"`
	CustomerUuid         types.String `tfsdk:"customer_uuid"`
	Description          types.String `tfsdk:"description"`
	ExternalIp           types.String `tfsdk:"external_ip"`
	Name                 types.String `tfsdk:"name"`
	NameExact            types.String `tfsdk:"name_exact"`
	Project              types.String `tfsdk:"project"`
	ProjectName          types.String `tfsdk:"project_name"`
	ProjectUuid          types.String `tfsdk:"project_uuid"`
	Query                types.String `tfsdk:"query"`
	RuntimeState         types.String `tfsdk:"runtime_state"`
	ServiceSettingsName  types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid  types.String `tfsdk:"service_settings_uuid"`
	Tenant               types.String `tfsdk:"tenant"`
	TenantUuid           types.String `tfsdk:"tenant_uuid"`
	Uuid                 types.String `tfsdk:"uuid"`
}

type OpenstackInstanceModel struct {
	UUID                             types.String      `tfsdk:"id"`
	AccessUrl                        types.String      `tfsdk:"access_url"`
	Action                           types.String      `tfsdk:"action"`
	AvailabilityZone                 types.String      `tfsdk:"availability_zone"`
	AvailabilityZoneName             types.String      `tfsdk:"availability_zone_name"`
	BackendId                        types.String      `tfsdk:"backend_id"`
	ConnectDirectlyToExternalNetwork types.Bool        `tfsdk:"connect_directly_to_external_network"`
	Cores                            types.Int64       `tfsdk:"cores"`
	Created                          timetypes.RFC3339 `tfsdk:"created"`
	Customer                         types.String      `tfsdk:"customer"`
	CustomerAbbreviation             types.String      `tfsdk:"customer_abbreviation"`
	CustomerName                     types.String      `tfsdk:"customer_name"`
	CustomerNativeName               types.String      `tfsdk:"customer_native_name"`
	CustomerUuid                     types.String      `tfsdk:"customer_uuid"`
	Description                      types.String      `tfsdk:"description"`
	Disk                             types.Int64       `tfsdk:"disk"`
	ErrorMessage                     types.String      `tfsdk:"error_message"`
	ErrorTraceback                   types.String      `tfsdk:"error_traceback"`
	ExternalAddress                  types.List        `tfsdk:"external_address"`
	ExternalIps                      types.List        `tfsdk:"external_ips"`
	FlavorDisk                       types.Int64       `tfsdk:"flavor_disk"`
	FlavorName                       types.String      `tfsdk:"flavor_name"`
	FloatingIps                      types.List        `tfsdk:"floating_ips"`
	HypervisorHostname               types.String      `tfsdk:"hypervisor_hostname"`
	ImageName                        types.String      `tfsdk:"image_name"`
	InternalIps                      types.List        `tfsdk:"internal_ips"`
	IsLimitBased                     types.Bool        `tfsdk:"is_limit_based"`
	IsUsageBased                     types.Bool        `tfsdk:"is_usage_based"`
	KeyFingerprint                   types.String      `tfsdk:"key_fingerprint"`
	KeyName                          types.String      `tfsdk:"key_name"`
	Latitude                         types.Float64     `tfsdk:"latitude"`
	Longitude                        types.Float64     `tfsdk:"longitude"`
	MarketplaceCategoryName          types.String      `tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid          types.String      `tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName          types.String      `tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid          types.String      `tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid              types.String      `tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState         types.String      `tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid          types.String      `tfsdk:"marketplace_resource_uuid"`
	MinDisk                          types.Int64       `tfsdk:"min_disk"`
	MinRam                           types.Int64       `tfsdk:"min_ram"`
	Modified                         timetypes.RFC3339 `tfsdk:"modified"`
	Name                             types.String      `tfsdk:"name"`
	Ports                            types.List        `tfsdk:"ports"`
	Project                          types.String      `tfsdk:"project"`
	ProjectName                      types.String      `tfsdk:"project_name"`
	ProjectUuid                      types.String      `tfsdk:"project_uuid"`
	Ram                              types.Int64       `tfsdk:"ram"`
	ResourceType                     types.String      `tfsdk:"resource_type"`
	RuntimeState                     types.String      `tfsdk:"runtime_state"`
	SecurityGroups                   types.List        `tfsdk:"security_groups"`
	ServerGroup                      types.Object      `tfsdk:"server_group"`
	ServiceName                      types.String      `tfsdk:"service_name"`
	ServiceSettings                  types.String      `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage      types.String      `tfsdk:"service_settings_error_message"`
	ServiceSettingsState             types.String      `tfsdk:"service_settings_state"`
	ServiceSettingsUuid              types.String      `tfsdk:"service_settings_uuid"`
	StartTime                        timetypes.RFC3339 `tfsdk:"start_time"`
	State                            types.String      `tfsdk:"state"`
	Tenant                           types.String      `tfsdk:"tenant"`
	TenantUuid                       types.String      `tfsdk:"tenant_uuid"`
	Url                              types.String      `tfsdk:"url"`
	UserData                         types.String      `tfsdk:"user_data"`
	Volumes                          types.List        `tfsdk:"volumes"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackInstanceModel) CopyFrom(ctx context.Context, apiResp OpenstackInstanceResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.Action = types.StringPointerValue(apiResp.Action)
	model.AvailabilityZone = types.StringPointerValue(apiResp.AvailabilityZone)
	model.AvailabilityZoneName = types.StringPointerValue(apiResp.AvailabilityZoneName)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.ConnectDirectlyToExternalNetwork = types.BoolPointerValue(apiResp.ConnectDirectlyToExternalNetwork)
	model.Cores = types.Int64PointerValue(apiResp.Cores)
	valCreated, diagsCreated := timetypes.NewRFC3339PointerValue(apiResp.Created)
	diags.Append(diagsCreated...)
	model.Created = valCreated
	model.Customer = types.StringPointerValue(apiResp.Customer)
	model.CustomerAbbreviation = types.StringPointerValue(apiResp.CustomerAbbreviation)
	model.CustomerName = types.StringPointerValue(apiResp.CustomerName)
	model.CustomerNativeName = types.StringPointerValue(apiResp.CustomerNativeName)
	model.CustomerUuid = types.StringPointerValue(apiResp.CustomerUuid)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.Disk = types.Int64PointerValue(apiResp.Disk)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	listValExternalAddress, listDiagsExternalAddress := types.ListValueFrom(ctx, types.StringType, apiResp.ExternalAddress)
	model.ExternalAddress = listValExternalAddress
	diags.Append(listDiagsExternalAddress...)
	listValExternalIps, listDiagsExternalIps := types.ListValueFrom(ctx, types.StringType, apiResp.ExternalIps)
	model.ExternalIps = listValExternalIps
	diags.Append(listDiagsExternalIps...)
	model.FlavorDisk = types.Int64PointerValue(apiResp.FlavorDisk)
	model.FlavorName = types.StringPointerValue(apiResp.FlavorName)

	{
		listValFloatingIps, listDiagsFloatingIps := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"ip_address": types.StringType,
			"subnet":     types.StringType,
			"url":        types.StringType,
			"address":    types.StringType,
			"port_fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"ip_address": types.StringType,
				"subnet_id":  types.StringType,
			}}},
			"port_mac_address":   types.StringType,
			"subnet_cidr":        types.StringType,
			"subnet_description": types.StringType,
			"subnet_name":        types.StringType,
			"subnet_uuid":        types.StringType,
		}}, apiResp.FloatingIps)
		diags.Append(listDiagsFloatingIps...)
		model.FloatingIps = listValFloatingIps
	}
	model.HypervisorHostname = types.StringPointerValue(apiResp.HypervisorHostname)
	model.ImageName = types.StringPointerValue(apiResp.ImageName)
	listValInternalIps, listDiagsInternalIps := types.ListValueFrom(ctx, types.StringType, apiResp.InternalIps)
	model.InternalIps = listValInternalIps
	diags.Append(listDiagsInternalIps...)
	model.IsLimitBased = types.BoolPointerValue(apiResp.IsLimitBased)
	model.IsUsageBased = types.BoolPointerValue(apiResp.IsUsageBased)
	model.KeyFingerprint = types.StringPointerValue(apiResp.KeyFingerprint)
	model.KeyName = types.StringPointerValue(apiResp.KeyName)
	model.Latitude = types.Float64PointerValue(apiResp.Latitude)
	model.Longitude = types.Float64PointerValue(apiResp.Longitude)
	model.MarketplaceCategoryName = types.StringPointerValue(apiResp.MarketplaceCategoryName)
	model.MarketplaceCategoryUuid = types.StringPointerValue(apiResp.MarketplaceCategoryUuid)
	model.MarketplaceOfferingName = types.StringPointerValue(apiResp.MarketplaceOfferingName)
	model.MarketplaceOfferingUuid = types.StringPointerValue(apiResp.MarketplaceOfferingUuid)
	model.MarketplacePlanUuid = types.StringPointerValue(apiResp.MarketplacePlanUuid)
	model.MarketplaceResourceState = types.StringPointerValue(apiResp.MarketplaceResourceState)
	model.MarketplaceResourceUuid = types.StringPointerValue(apiResp.MarketplaceResourceUuid)
	model.MinDisk = types.Int64PointerValue(apiResp.MinDisk)
	model.MinRam = types.Int64PointerValue(apiResp.MinRam)
	valModified, diagsModified := timetypes.NewRFC3339PointerValue(apiResp.Modified)
	diags.Append(diagsModified...)
	model.Modified = valModified
	model.Name = types.StringPointerValue(apiResp.Name)

	{
		listValPorts, listDiagsPorts := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"ip_address": types.StringType,
				"subnet_id":  types.StringType,
			}}},
			"port":   types.StringType,
			"subnet": types.StringType,
			"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"mac_address": types.StringType,
			}}},
			"device_id":    types.StringType,
			"device_owner": types.StringType,
			"mac_address":  types.StringType,
			"security_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"access_url":                 types.StringType,
				"backend_id":                 types.StringType,
				"created":                    types.StringType,
				"customer":                   types.StringType,
				"customer_abbreviation":      types.StringType,
				"customer_name":              types.StringType,
				"customer_native_name":       types.StringType,
				"customer_uuid":              types.StringType,
				"description":                types.StringType,
				"error_message":              types.StringType,
				"error_traceback":            types.StringType,
				"is_limit_based":             types.BoolType,
				"is_usage_based":             types.BoolType,
				"marketplace_category_name":  types.StringType,
				"marketplace_category_uuid":  types.StringType,
				"marketplace_offering_name":  types.StringType,
				"marketplace_offering_uuid":  types.StringType,
				"marketplace_plan_uuid":      types.StringType,
				"marketplace_resource_state": types.StringType,
				"marketplace_resource_uuid":  types.StringType,
				"modified":                   types.StringType,
				"name":                       types.StringType,
				"project":                    types.StringType,
				"project_name":               types.StringType,
				"project_uuid":               types.StringType,
				"resource_type":              types.StringType,
				"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"cidr":              types.StringType,
					"description":       types.StringType,
					"direction":         types.StringType,
					"ethertype":         types.StringType,
					"from_port":         types.Int64Type,
					"id":                types.Int64Type,
					"protocol":          types.StringType,
					"remote_group":      types.StringType,
					"remote_group_name": types.StringType,
					"remote_group_uuid": types.StringType,
					"to_port":           types.Int64Type,
				}}},
				"service_name":                   types.StringType,
				"service_settings":               types.StringType,
				"service_settings_error_message": types.StringType,
				"service_settings_state":         types.StringType,
				"service_settings_uuid":          types.StringType,
				"state":                          types.StringType,
				"tenant":                         types.StringType,
				"tenant_name":                    types.StringType,
				"tenant_uuid":                    types.StringType,
				"url":                            types.StringType,
			}}},
			"subnet_cidr":        types.StringType,
			"subnet_description": types.StringType,
			"subnet_name":        types.StringType,
			"subnet_uuid":        types.StringType,
			"url":                types.StringType,
		}}, apiResp.Ports)
		diags.Append(listDiagsPorts...)
		model.Ports = listValPorts
	}
	model.Project = types.StringPointerValue(apiResp.Project)
	model.ProjectName = types.StringPointerValue(apiResp.ProjectName)
	model.ProjectUuid = types.StringPointerValue(apiResp.ProjectUuid)
	model.Ram = types.Int64PointerValue(apiResp.Ram)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.RuntimeState = types.StringPointerValue(apiResp.RuntimeState)

	{
		listValSecurityGroups, listDiagsSecurityGroups := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"url":         types.StringType,
			"description": types.StringType,
			"name":        types.StringType,
			"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"cidr":              types.StringType,
				"description":       types.StringType,
				"direction":         types.StringType,
				"ethertype":         types.StringType,
				"from_port":         types.Int64Type,
				"id":                types.Int64Type,
				"protocol":          types.StringType,
				"remote_group_name": types.StringType,
				"remote_group_uuid": types.StringType,
				"to_port":           types.Int64Type,
			}}},
			"state": types.StringType,
		}}, apiResp.SecurityGroups)
		diags.Append(listDiagsSecurityGroups...)
		model.SecurityGroups = listValSecurityGroups
	}
	if apiResp.ServerGroup != nil {
		objValServerGroup, objDiagsServerGroup := types.ObjectValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"name":   types.StringType,
			"policy": types.StringType,
			"state":  types.StringType,
			"url":    types.StringType,
		}}.AttrTypes, *apiResp.ServerGroup)
		diags.Append(objDiagsServerGroup...)
		model.ServerGroup = objValServerGroup
	} else {
		model.ServerGroup = types.ObjectNull(types.ObjectType{AttrTypes: map[string]attr.Type{
			"name":   types.StringType,
			"policy": types.StringType,
			"state":  types.StringType,
			"url":    types.StringType,
		}}.AttrTypes)
	}
	model.ServiceName = types.StringPointerValue(apiResp.ServiceName)
	model.ServiceSettings = types.StringPointerValue(apiResp.ServiceSettings)
	model.ServiceSettingsErrorMessage = types.StringPointerValue(apiResp.ServiceSettingsErrorMessage)
	model.ServiceSettingsState = types.StringPointerValue(apiResp.ServiceSettingsState)
	model.ServiceSettingsUuid = types.StringPointerValue(apiResp.ServiceSettingsUuid)
	valStartTime, diagsStartTime := timetypes.NewRFC3339PointerValue(apiResp.StartTime)
	diags.Append(diagsStartTime...)
	model.StartTime = valStartTime
	model.State = types.StringPointerValue(apiResp.State)
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
	model.Url = types.StringPointerValue(apiResp.Url)
	model.UserData = types.StringPointerValue(apiResp.UserData)

	{
		listValVolumes, listDiagsVolumes := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"bootable":                  types.BoolType,
			"device":                    types.StringType,
			"image_name":                types.StringType,
			"marketplace_resource_uuid": types.StringType,
			"name":                      types.StringType,
			"resource_type":             types.StringType,
			"size":                      types.Int64Type,
			"state":                     types.StringType,
			"type":                      types.StringType,
			"type_name":                 types.StringType,
			"url":                       types.StringType,
		}}, apiResp.Volumes)
		diags.Append(listDiagsVolumes...)
		model.Volumes = listValVolumes
	}

	return diags
}
