package instance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

func NestedSecurityGroupRuleType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
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
	}}
}
func OpenStackAllowedAddressPairType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"mac_address": types.StringType,
	}}
}
func OpenStackCreateFloatingIPRequestType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"ip_address":         types.StringType,
		"subnet":             types.StringType,
		"url":                types.StringType,
		"address":            types.StringType,
		"port_fixed_ips":     types.ListType{ElemType: OpenStackFixedIpType()},
		"port_mac_address":   types.StringType,
		"subnet_cidr":        types.StringType,
		"subnet_description": types.StringType,
		"subnet_name":        types.StringType,
		"subnet_uuid":        types.StringType,
		"uuid":               types.StringType,
	}}
}
func OpenStackCreateInstancePortRequestType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"fixed_ips":             types.ListType{ElemType: OpenStackFixedIpType()},
		"port":                  types.StringType,
		"subnet":                types.StringType,
		"allowed_address_pairs": types.ListType{ElemType: OpenStackAllowedAddressPairType()},
		"device_id":             types.StringType,
		"device_owner":          types.StringType,
		"mac_address":           types.StringType,
		"security_groups":       types.SetType{ElemType: OpenStackSecurityGroupType()},
		"subnet_cidr":           types.StringType,
		"subnet_description":    types.StringType,
		"subnet_name":           types.StringType,
		"subnet_uuid":           types.StringType,
		"url":                   types.StringType,
	}}
}
func OpenStackDataVolumeRequestType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"size":        types.Int64Type,
		"volume_type": types.StringType,
	}}
}
func OpenStackFixedIpType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"ip_address": types.StringType,
		"subnet_id":  types.StringType,
	}}
}
func OpenStackNestedVolumeType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
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
		"uuid":                      types.StringType,
	}}
}
func OpenStackSecurityGroupType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"backend_id":                types.StringType,
		"customer":                  types.StringType,
		"description":               types.StringType,
		"error_message":             types.StringType,
		"marketplace_resource_uuid": types.StringType,
		"name":                      types.StringType,
		"project":                   types.StringType,
		"resource_type":             types.StringType,
		"rules":                     types.ListType{ElemType: OpenStackSecurityGroupRuleCreateType()},
		"state":                     types.StringType,
		"tenant":                    types.StringType,
		"tenant_name":               types.StringType,
		"tenant_uuid":               types.StringType,
		"url":                       types.StringType,
		"uuid":                      types.StringType,
	}}
}
func OpenStackSecurityGroupHyperlinkRequestType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"url":         types.StringType,
		"description": types.StringType,
		"name":        types.StringType,
		"rules":       types.ListType{ElemType: NestedSecurityGroupRuleType()},
		"state":       types.StringType,
	}}
}
func OpenStackSecurityGroupRuleCreateType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
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
	}}
}
func RancherClusterType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"marketplace_uuid": types.StringType,
		"name":             types.StringType,
		"uuid":             types.StringType,
	}}
}
func ServerGroupType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"url":    types.StringType,
		"name":   types.StringType,
		"policy": types.StringType,
		"state":  types.StringType,
	}}
}

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

func (m *OpenstackInstanceFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Openstack Instance",
		Attributes: map[string]schema.Attribute{
			"attach_volume_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Filter for attachment to volume UUID",
			},
			"availability_zone_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Availability zone name",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Backend ID",
			},
			"can_manage": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Can manage",
			},
			"customer": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer UUID",
			},
			"customer_abbreviation": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer abbreviation",
			},
			"customer_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer name",
			},
			"customer_native_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer native name",
			},
			"customer_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer UUID",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description",
			},
			"external_ip": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "External IP",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name (exact)",
			},
			"project": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project UUID",
			},
			"project_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project name",
			},
			"project_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project UUID",
			},
			"query": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Search by name, internal IP, or external IP",
			},
			"runtime_state": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Runtime state",
			},
			"service_settings_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Service settings name",
			},
			"service_settings_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Service settings UUID",
			},
			"tenant": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Tenant URL",
			},
			"tenant_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Tenant UUID",
			},
			"uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "UUID",
			},
		},
	}
}

type OpenstackInstanceModel struct {
	UUID                             types.String      `tfsdk:"id"`
	Action                           types.String      `tfsdk:"action"`
	AvailabilityZone                 types.String      `tfsdk:"availability_zone"`
	AvailabilityZoneName             types.String      `tfsdk:"availability_zone_name"`
	BackendId                        types.String      `tfsdk:"backend_id"`
	ConnectDirectlyToExternalNetwork types.Bool        `tfsdk:"connect_directly_to_external_network"`
	Cores                            types.Int64       `tfsdk:"cores"`
	Customer                         types.String      `tfsdk:"customer"`
	Description                      types.String      `tfsdk:"description"`
	Disk                             types.Int64       `tfsdk:"disk"`
	ErrorMessage                     types.String      `tfsdk:"error_message"`
	ExternalAddress                  types.List        `tfsdk:"external_address"`
	ExternalIps                      types.List        `tfsdk:"external_ips"`
	FlavorDisk                       types.Int64       `tfsdk:"flavor_disk"`
	FlavorName                       types.String      `tfsdk:"flavor_name"`
	FloatingIps                      types.Set         `tfsdk:"floating_ips"`
	HypervisorHostname               types.String      `tfsdk:"hypervisor_hostname"`
	ImageName                        types.String      `tfsdk:"image_name"`
	InternalIps                      types.List        `tfsdk:"internal_ips"`
	KeyFingerprint                   types.String      `tfsdk:"key_fingerprint"`
	KeyName                          types.String      `tfsdk:"key_name"`
	Latitude                         types.Float64     `tfsdk:"latitude"`
	Longitude                        types.Float64     `tfsdk:"longitude"`
	MarketplaceResourceUuid          types.String      `tfsdk:"marketplace_resource_uuid"`
	MinDisk                          types.Int64       `tfsdk:"min_disk"`
	MinRam                           types.Int64       `tfsdk:"min_ram"`
	Name                             types.String      `tfsdk:"name"`
	Ports                            types.List        `tfsdk:"ports"`
	Project                          types.String      `tfsdk:"project"`
	Ram                              types.Int64       `tfsdk:"ram"`
	RancherCluster                   types.Object      `tfsdk:"rancher_cluster"`
	ResourceType                     types.String      `tfsdk:"resource_type"`
	RuntimeState                     types.String      `tfsdk:"runtime_state"`
	SecurityGroups                   types.Set         `tfsdk:"security_groups"`
	ServerGroup                      types.Object      `tfsdk:"server_group"`
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
	model.Action = common.StringPointerValue(apiResp.Action)
	model.AvailabilityZone = common.StringPointerValue(apiResp.AvailabilityZone)
	model.AvailabilityZoneName = common.StringPointerValue(apiResp.AvailabilityZoneName)
	model.BackendId = common.StringPointerValue(apiResp.BackendId)
	model.ConnectDirectlyToExternalNetwork = types.BoolPointerValue(apiResp.ConnectDirectlyToExternalNetwork)
	model.Cores = types.Int64PointerValue(apiResp.Cores)
	model.Customer = common.StringPointerValue(apiResp.Customer)
	model.Description = common.StringPointerValue(apiResp.Description)
	model.Disk = types.Int64PointerValue(apiResp.Disk)
	model.ErrorMessage = common.StringPointerValue(apiResp.ErrorMessage)
	listValExternalAddress, listDiagsExternalAddress := types.ListValueFrom(ctx, types.StringType, apiResp.ExternalAddress)
	model.ExternalAddress = listValExternalAddress
	diags.Append(listDiagsExternalAddress...)
	listValExternalIps, listDiagsExternalIps := types.ListValueFrom(ctx, types.StringType, apiResp.ExternalIps)
	model.ExternalIps = listValExternalIps
	diags.Append(listDiagsExternalIps...)
	model.FlavorDisk = types.Int64PointerValue(apiResp.FlavorDisk)
	model.FlavorName = common.StringPointerValue(apiResp.FlavorName)
	if apiResp.FloatingIps != nil && len(*apiResp.FloatingIps) > 0 {
		setValFloatingIps, setDiagsFloatingIps := types.SetValueFrom(ctx, OpenStackCreateFloatingIPRequestType(), apiResp.FloatingIps)
		diags.Append(setDiagsFloatingIps...)
		model.FloatingIps = setValFloatingIps
	} else {
		model.FloatingIps = types.SetNull(OpenStackCreateFloatingIPRequestType())
	}
	model.HypervisorHostname = common.StringPointerValue(apiResp.HypervisorHostname)
	model.ImageName = common.StringPointerValue(apiResp.ImageName)
	listValInternalIps, listDiagsInternalIps := types.ListValueFrom(ctx, types.StringType, apiResp.InternalIps)
	model.InternalIps = listValInternalIps
	diags.Append(listDiagsInternalIps...)
	model.KeyFingerprint = common.StringPointerValue(apiResp.KeyFingerprint)
	model.KeyName = common.StringPointerValue(apiResp.KeyName)
	model.Latitude = types.Float64PointerValue(apiResp.Latitude.Float64Ptr())
	model.Longitude = types.Float64PointerValue(apiResp.Longitude.Float64Ptr())
	model.MarketplaceResourceUuid = common.StringPointerValue(apiResp.MarketplaceResourceUuid)
	model.MinDisk = types.Int64PointerValue(apiResp.MinDisk)
	model.MinRam = types.Int64PointerValue(apiResp.MinRam)
	model.Name = common.StringPointerValue(apiResp.Name)

	listValPorts, listDiagsPorts := types.ListValueFrom(ctx, OpenStackCreateInstancePortRequestType(), apiResp.Ports)
	diags.Append(listDiagsPorts...)
	model.Ports = listValPorts
	model.Project = common.StringPointerValue(apiResp.Project)
	model.Ram = types.Int64PointerValue(apiResp.Ram)
	if apiResp.RancherCluster != nil {
		objValRancherCluster, objDiagsRancherCluster := types.ObjectValueFrom(ctx, RancherClusterType().AttrTypes, *apiResp.RancherCluster)
		diags.Append(objDiagsRancherCluster...)
		model.RancherCluster = objValRancherCluster
	} else {
		model.RancherCluster = types.ObjectNull(RancherClusterType().AttrTypes)
	}
	model.ResourceType = common.StringPointerValue(apiResp.ResourceType)
	model.RuntimeState = common.StringPointerValue(apiResp.RuntimeState)
	if apiResp.SecurityGroups != nil && len(*apiResp.SecurityGroups) > 0 {
		setValSecurityGroups, setDiagsSecurityGroups := types.SetValueFrom(ctx, OpenStackSecurityGroupHyperlinkRequestType(), apiResp.SecurityGroups)
		diags.Append(setDiagsSecurityGroups...)
		model.SecurityGroups = setValSecurityGroups
	} else {
		model.SecurityGroups = types.SetNull(OpenStackSecurityGroupHyperlinkRequestType())
	}
	if apiResp.ServerGroup != nil {
		objValServerGroup, objDiagsServerGroup := types.ObjectValueFrom(ctx, ServerGroupType().AttrTypes, *apiResp.ServerGroup)
		diags.Append(objDiagsServerGroup...)
		model.ServerGroup = objValServerGroup
	} else {
		model.ServerGroup = types.ObjectNull(ServerGroupType().AttrTypes)
	}
	valStartTime, diagsStartTime := timetypes.NewRFC3339PointerValue(apiResp.StartTime)
	diags.Append(diagsStartTime...)
	model.StartTime = valStartTime
	model.State = common.StringPointerValue(apiResp.State)
	model.Tenant = common.StringPointerValue(apiResp.Tenant)
	model.TenantUuid = common.StringPointerValue(apiResp.TenantUuid)
	model.Url = common.StringPointerValue(apiResp.Url)
	model.UserData = common.StringPointerValue(apiResp.UserData)

	if apiResp.Volumes != nil {
		listValVolumes, listDiagsVolumes := types.ListValueFrom(ctx, OpenStackNestedVolumeType(), apiResp.Volumes)
		diags.Append(listDiagsVolumes...)
		model.Volumes = listValVolumes
	} else {
		model.Volumes = types.ListNull(OpenStackNestedVolumeType())
	}

	return diags
}
