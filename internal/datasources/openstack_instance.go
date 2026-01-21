package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackInstanceDataSource{}

func NewOpenstackInstanceDataSource() datasource.DataSource {
	return &OpenstackInstanceDataSource{}
}

// OpenstackInstanceDataSource defines the data source implementation.
type OpenstackInstanceDataSource struct {
	client *client.Client
}

// OpenstackInstanceApiResponse is the API response model.
type OpenstackInstanceApiResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl                        *string                                   `json:"access_url" tfsdk:"access_url"`
	Action                           *string                                   `json:"action" tfsdk:"action"`
	AvailabilityZone                 *string                                   `json:"availability_zone" tfsdk:"availability_zone"`
	AvailabilityZoneName             *string                                   `json:"availability_zone_name" tfsdk:"availability_zone_name"`
	BackendId                        *string                                   `json:"backend_id" tfsdk:"backend_id"`
	ConnectDirectlyToExternalNetwork *bool                                     `json:"connect_directly_to_external_network" tfsdk:"connect_directly_to_external_network"`
	Cores                            *int64                                    `json:"cores" tfsdk:"cores"`
	Created                          *string                                   `json:"created" tfsdk:"created"`
	Customer                         *string                                   `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation             *string                                   `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName                     *string                                   `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName               *string                                   `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerUuid                     *string                                   `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description                      *string                                   `json:"description" tfsdk:"description"`
	Disk                             *int64                                    `json:"disk" tfsdk:"disk"`
	ErrorMessage                     *string                                   `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback                   *string                                   `json:"error_traceback" tfsdk:"error_traceback"`
	ExternalAddress                  []string                                  `json:"external_address" tfsdk:"external_address"`
	ExternalIps                      []string                                  `json:"external_ips" tfsdk:"external_ips"`
	FlavorDisk                       *int64                                    `json:"flavor_disk" tfsdk:"flavor_disk"`
	FlavorName                       *string                                   `json:"flavor_name" tfsdk:"flavor_name"`
	FloatingIps                      []OpenstackInstanceFloatingIpsResponse    `json:"floating_ips" tfsdk:"floating_ips"`
	HypervisorHostname               *string                                   `json:"hypervisor_hostname" tfsdk:"hypervisor_hostname"`
	ImageName                        *string                                   `json:"image_name" tfsdk:"image_name"`
	InternalIps                      []string                                  `json:"internal_ips" tfsdk:"internal_ips"`
	IsLimitBased                     *bool                                     `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased                     *bool                                     `json:"is_usage_based" tfsdk:"is_usage_based"`
	KeyFingerprint                   *string                                   `json:"key_fingerprint" tfsdk:"key_fingerprint"`
	KeyName                          *string                                   `json:"key_name" tfsdk:"key_name"`
	Latitude                         *float64                                  `json:"latitude" tfsdk:"latitude"`
	Longitude                        *float64                                  `json:"longitude" tfsdk:"longitude"`
	MarketplaceCategoryName          *string                                   `json:"marketplace_category_name" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid          *string                                   `json:"marketplace_category_uuid" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName          *string                                   `json:"marketplace_offering_name" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid          *string                                   `json:"marketplace_offering_uuid" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid              *string                                   `json:"marketplace_plan_uuid" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState         *string                                   `json:"marketplace_resource_state" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid          *string                                   `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	MinDisk                          *int64                                    `json:"min_disk" tfsdk:"min_disk"`
	MinRam                           *int64                                    `json:"min_ram" tfsdk:"min_ram"`
	Modified                         *string                                   `json:"modified" tfsdk:"modified"`
	Name                             *string                                   `json:"name" tfsdk:"name"`
	Ports                            []OpenstackInstancePortsResponse          `json:"ports" tfsdk:"ports"`
	Project                          *string                                   `json:"project" tfsdk:"project"`
	ProjectName                      *string                                   `json:"project_name" tfsdk:"project_name"`
	ProjectUuid                      *string                                   `json:"project_uuid" tfsdk:"project_uuid"`
	Ram                              *int64                                    `json:"ram" tfsdk:"ram"`
	ResourceType                     *string                                   `json:"resource_type" tfsdk:"resource_type"`
	RuntimeState                     *string                                   `json:"runtime_state" tfsdk:"runtime_state"`
	SecurityGroups                   []OpenstackInstanceSecurityGroupsResponse `json:"security_groups" tfsdk:"security_groups"`
	ServerGroup                      *OpenstackInstanceServerGroupResponse     `json:"server_group" tfsdk:"server_group"`
	ServiceName                      *string                                   `json:"service_name" tfsdk:"service_name"`
	ServiceSettings                  *string                                   `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage      *string                                   `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState             *string                                   `json:"service_settings_state" tfsdk:"service_settings_state"`
	ServiceSettingsUuid              *string                                   `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	StartTime                        *string                                   `json:"start_time" tfsdk:"start_time"`
	State                            *string                                   `json:"state" tfsdk:"state"`
	Tenant                           *string                                   `json:"tenant" tfsdk:"tenant"`
	TenantUuid                       *string                                   `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Url                              *string                                   `json:"url" tfsdk:"url"`
	UserData                         *string                                   `json:"user_data" tfsdk:"user_data"`
	Volumes                          []OpenstackInstanceVolumesResponse        `json:"volumes" tfsdk:"volumes"`
}

type OpenstackInstanceFloatingIpsResponse struct {
	Address           *string                                            `json:"address" tfsdk:"address"`
	PortFixedIps      []OpenstackInstanceFloatingIpsPortFixedIpsResponse `json:"port_fixed_ips" tfsdk:"port_fixed_ips"`
	PortMacAddress    *string                                            `json:"port_mac_address" tfsdk:"port_mac_address"`
	Subnet            *string                                            `json:"subnet" tfsdk:"subnet"`
	SubnetCidr        *string                                            `json:"subnet_cidr" tfsdk:"subnet_cidr"`
	SubnetDescription *string                                            `json:"subnet_description" tfsdk:"subnet_description"`
	SubnetName        *string                                            `json:"subnet_name" tfsdk:"subnet_name"`
	SubnetUuid        *string                                            `json:"subnet_uuid" tfsdk:"subnet_uuid"`
	Url               *string                                            `json:"url" tfsdk:"url"`
}

type OpenstackInstanceFloatingIpsPortFixedIpsResponse struct {
	IpAddress *string `json:"ip_address" tfsdk:"ip_address"`
	SubnetId  *string `json:"subnet_id" tfsdk:"subnet_id"`
}

type OpenstackInstancePortsResponse struct {
	AllowedAddressPairs []OpenstackInstancePortsAllowedAddressPairsResponse `json:"allowed_address_pairs" tfsdk:"allowed_address_pairs"`
	DeviceId            *string                                             `json:"device_id" tfsdk:"device_id"`
	DeviceOwner         *string                                             `json:"device_owner" tfsdk:"device_owner"`
	FixedIps            []OpenstackInstancePortsFixedIpsResponse            `json:"fixed_ips" tfsdk:"fixed_ips"`
	MacAddress          *string                                             `json:"mac_address" tfsdk:"mac_address"`
	SecurityGroups      []OpenstackInstancePortsSecurityGroupsResponse      `json:"security_groups" tfsdk:"security_groups"`
	Subnet              *string                                             `json:"subnet" tfsdk:"subnet"`
	SubnetCidr          *string                                             `json:"subnet_cidr" tfsdk:"subnet_cidr"`
	SubnetDescription   *string                                             `json:"subnet_description" tfsdk:"subnet_description"`
	SubnetName          *string                                             `json:"subnet_name" tfsdk:"subnet_name"`
	SubnetUuid          *string                                             `json:"subnet_uuid" tfsdk:"subnet_uuid"`
	Url                 *string                                             `json:"url" tfsdk:"url"`
}

type OpenstackInstancePortsAllowedAddressPairsResponse struct {
	MacAddress *string `json:"mac_address" tfsdk:"mac_address"`
}

type OpenstackInstancePortsFixedIpsResponse struct {
	IpAddress *string `json:"ip_address" tfsdk:"ip_address"`
	SubnetId  *string `json:"subnet_id" tfsdk:"subnet_id"`
}

type OpenstackInstancePortsSecurityGroupsResponse struct {
	AccessUrl                   *string                                             `json:"access_url" tfsdk:"access_url"`
	BackendId                   *string                                             `json:"backend_id" tfsdk:"backend_id"`
	Created                     *string                                             `json:"created" tfsdk:"created"`
	Customer                    *string                                             `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation        *string                                             `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName                *string                                             `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName          *string                                             `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerUuid                *string                                             `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description                 *string                                             `json:"description" tfsdk:"description"`
	ErrorMessage                *string                                             `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback              *string                                             `json:"error_traceback" tfsdk:"error_traceback"`
	IsLimitBased                *bool                                               `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                                               `json:"is_usage_based" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string                                             `json:"marketplace_category_name" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                                             `json:"marketplace_category_uuid" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                                             `json:"marketplace_offering_name" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                                             `json:"marketplace_offering_uuid" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                                             `json:"marketplace_plan_uuid" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                                             `json:"marketplace_resource_state" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                                             `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                                             `json:"modified" tfsdk:"modified"`
	Name                        *string                                             `json:"name" tfsdk:"name"`
	Project                     *string                                             `json:"project" tfsdk:"project"`
	ProjectName                 *string                                             `json:"project_name" tfsdk:"project_name"`
	ProjectUuid                 *string                                             `json:"project_uuid" tfsdk:"project_uuid"`
	ResourceType                *string                                             `json:"resource_type" tfsdk:"resource_type"`
	Rules                       []OpenstackInstancePortsSecurityGroupsRulesResponse `json:"rules" tfsdk:"rules"`
	ServiceName                 *string                                             `json:"service_name" tfsdk:"service_name"`
	ServiceSettings             *string                                             `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                                             `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                                             `json:"service_settings_state" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                                             `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	State                       *string                                             `json:"state" tfsdk:"state"`
	Tenant                      *string                                             `json:"tenant" tfsdk:"tenant"`
	TenantName                  *string                                             `json:"tenant_name" tfsdk:"tenant_name"`
	TenantUuid                  *string                                             `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Url                         *string                                             `json:"url" tfsdk:"url"`
}

type OpenstackInstancePortsSecurityGroupsRulesResponse struct {
	Cidr            *string `json:"cidr" tfsdk:"cidr"`
	Description     *string `json:"description" tfsdk:"description"`
	Direction       *string `json:"direction" tfsdk:"direction"`
	Ethertype       *string `json:"ethertype" tfsdk:"ethertype"`
	FromPort        *int64  `json:"from_port" tfsdk:"from_port"`
	Id              *int64  `json:"id" tfsdk:"id"`
	Protocol        *string `json:"protocol" tfsdk:"protocol"`
	RemoteGroup     *string `json:"remote_group" tfsdk:"remote_group"`
	RemoteGroupName *string `json:"remote_group_name" tfsdk:"remote_group_name"`
	RemoteGroupUuid *string `json:"remote_group_uuid" tfsdk:"remote_group_uuid"`
	ToPort          *int64  `json:"to_port" tfsdk:"to_port"`
}

type OpenstackInstanceSecurityGroupsResponse struct {
	Description *string                                        `json:"description" tfsdk:"description"`
	Name        *string                                        `json:"name" tfsdk:"name"`
	Rules       []OpenstackInstanceSecurityGroupsRulesResponse `json:"rules" tfsdk:"rules"`
	State       *string                                        `json:"state" tfsdk:"state"`
	Url         *string                                        `json:"url" tfsdk:"url"`
}

type OpenstackInstanceSecurityGroupsRulesResponse struct {
	Cidr            *string `json:"cidr" tfsdk:"cidr"`
	Description     *string `json:"description" tfsdk:"description"`
	Direction       *string `json:"direction" tfsdk:"direction"`
	Ethertype       *string `json:"ethertype" tfsdk:"ethertype"`
	FromPort        *int64  `json:"from_port" tfsdk:"from_port"`
	Id              *int64  `json:"id" tfsdk:"id"`
	Protocol        *string `json:"protocol" tfsdk:"protocol"`
	RemoteGroupName *string `json:"remote_group_name" tfsdk:"remote_group_name"`
	RemoteGroupUuid *string `json:"remote_group_uuid" tfsdk:"remote_group_uuid"`
	ToPort          *int64  `json:"to_port" tfsdk:"to_port"`
}

type OpenstackInstanceServerGroupResponse struct {
	Name   *string `json:"name" tfsdk:"name"`
	Policy *string `json:"policy" tfsdk:"policy"`
	State  *string `json:"state" tfsdk:"state"`
	Url    *string `json:"url" tfsdk:"url"`
}

type OpenstackInstanceVolumesResponse struct {
	Bootable                *bool   `json:"bootable" tfsdk:"bootable"`
	Device                  *string `json:"device" tfsdk:"device"`
	ImageName               *string `json:"image_name" tfsdk:"image_name"`
	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Name                    *string `json:"name" tfsdk:"name"`
	ResourceType            *string `json:"resource_type" tfsdk:"resource_type"`
	Size                    *int64  `json:"size" tfsdk:"size"`
	State                   *string `json:"state" tfsdk:"state"`
	Type                    *string `json:"type" tfsdk:"type"`
	TypeName                *string `json:"type_name" tfsdk:"type_name"`
	Url                     *string `json:"url" tfsdk:"url"`
}

// OpenstackInstanceDataSourceModel describes the data source data model.
type OpenstackInstanceDataSourceModel struct {
	UUID                             types.String  `tfsdk:"id"`
	AttachVolumeUuid                 types.String  `tfsdk:"attach_volume_uuid"`
	AvailabilityZoneName             types.String  `tfsdk:"availability_zone_name"`
	BackendId                        types.String  `tfsdk:"backend_id"`
	CanManage                        types.Bool    `tfsdk:"can_manage"`
	Customer                         types.String  `tfsdk:"customer"`
	CustomerAbbreviation             types.String  `tfsdk:"customer_abbreviation"`
	CustomerName                     types.String  `tfsdk:"customer_name"`
	CustomerNativeName               types.String  `tfsdk:"customer_native_name"`
	CustomerUuid                     types.String  `tfsdk:"customer_uuid"`
	Description                      types.String  `tfsdk:"description"`
	ExternalIp                       types.String  `tfsdk:"external_ip"`
	Name                             types.String  `tfsdk:"name"`
	NameExact                        types.String  `tfsdk:"name_exact"`
	Project                          types.String  `tfsdk:"project"`
	ProjectName                      types.String  `tfsdk:"project_name"`
	ProjectUuid                      types.String  `tfsdk:"project_uuid"`
	Query                            types.String  `tfsdk:"query"`
	RuntimeState                     types.String  `tfsdk:"runtime_state"`
	ServiceSettingsName              types.String  `tfsdk:"service_settings_name"`
	ServiceSettingsUuid              types.String  `tfsdk:"service_settings_uuid"`
	State                            types.String  `tfsdk:"state"`
	Tenant                           types.String  `tfsdk:"tenant"`
	TenantUuid                       types.String  `tfsdk:"tenant_uuid"`
	Uuid                             types.String  `tfsdk:"uuid"`
	AccessUrl                        types.String  `tfsdk:"access_url"`
	Action                           types.String  `tfsdk:"action"`
	AvailabilityZone                 types.String  `tfsdk:"availability_zone"`
	ConnectDirectlyToExternalNetwork types.Bool    `tfsdk:"connect_directly_to_external_network"`
	Cores                            types.Int64   `tfsdk:"cores"`
	Created                          types.String  `tfsdk:"created"`
	Disk                             types.Int64   `tfsdk:"disk"`
	ErrorMessage                     types.String  `tfsdk:"error_message"`
	ErrorTraceback                   types.String  `tfsdk:"error_traceback"`
	ExternalAddress                  types.List    `tfsdk:"external_address"`
	ExternalIps                      types.List    `tfsdk:"external_ips"`
	FlavorDisk                       types.Int64   `tfsdk:"flavor_disk"`
	FlavorName                       types.String  `tfsdk:"flavor_name"`
	FloatingIps                      types.List    `tfsdk:"floating_ips"`
	HypervisorHostname               types.String  `tfsdk:"hypervisor_hostname"`
	ImageName                        types.String  `tfsdk:"image_name"`
	InternalIps                      types.List    `tfsdk:"internal_ips"`
	IsLimitBased                     types.Bool    `tfsdk:"is_limit_based"`
	IsUsageBased                     types.Bool    `tfsdk:"is_usage_based"`
	KeyFingerprint                   types.String  `tfsdk:"key_fingerprint"`
	KeyName                          types.String  `tfsdk:"key_name"`
	Latitude                         types.Float64 `tfsdk:"latitude"`
	Longitude                        types.Float64 `tfsdk:"longitude"`
	MarketplaceCategoryName          types.String  `tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid          types.String  `tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName          types.String  `tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid          types.String  `tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid              types.String  `tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState         types.String  `tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid          types.String  `tfsdk:"marketplace_resource_uuid"`
	MinDisk                          types.Int64   `tfsdk:"min_disk"`
	MinRam                           types.Int64   `tfsdk:"min_ram"`
	Modified                         types.String  `tfsdk:"modified"`
	Ports                            types.List    `tfsdk:"ports"`
	Ram                              types.Int64   `tfsdk:"ram"`
	ResourceType                     types.String  `tfsdk:"resource_type"`
	SecurityGroups                   types.List    `tfsdk:"security_groups"`
	ServerGroup                      types.Object  `tfsdk:"server_group"`
	ServiceName                      types.String  `tfsdk:"service_name"`
	ServiceSettings                  types.String  `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage      types.String  `tfsdk:"service_settings_error_message"`
	ServiceSettingsState             types.String  `tfsdk:"service_settings_state"`
	StartTime                        types.String  `tfsdk:"start_time"`
	Url                              types.String  `tfsdk:"url"`
	UserData                         types.String  `tfsdk:"user_data"`
	Volumes                          types.List    `tfsdk:"volumes"`
}

func (d *OpenstackInstanceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_instance"
}

func (d *OpenstackInstanceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Instance data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"attach_volume_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Filter for attachment to volume UUID",
			},
			"availability_zone_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Availability zone name",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Backend ID",
			},
			"can_manage": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Can manage",
			},
			"customer": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer UUID",
			},
			"customer_abbreviation": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer abbreviation",
			},
			"customer_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer name",
			},
			"customer_native_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer native name",
			},
			"customer_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer UUID",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Description",
			},
			"external_ip": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "External IP",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Name",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Name (exact)",
			},
			"project": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project UUID",
			},
			"project_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project name",
			},
			"project_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project UUID",
			},
			"query": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Search by name, internal IP, or external IP",
			},
			"runtime_state": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_settings_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Service settings name",
			},
			"service_settings_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Service settings UUID",
			},
			"state": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "State",
			},
			"tenant": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Tenant URL",
			},
			"tenant_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Tenant UUID",
			},
			"uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "UUID",
			},
			"access_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"action": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"availability_zone": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Availability zone where this instance is located",
			},
			"connect_directly_to_external_network": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "If True, instance will be connected directly to external network",
			},
			"cores": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of cores in a VM",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"disk": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Disk size in MiB",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"external_address": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"external_ips": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"flavor_disk": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Flavor disk size in MiB",
			},
			"flavor_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the flavor used by this instance",
			},
			"floating_ips": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"address": types.StringType,
					"port_fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"ip_address": types.StringType,
						"subnet_id":  types.StringType,
					}}},
					"port_mac_address":   types.StringType,
					"subnet":             types.StringType,
					"subnet_cidr":        types.StringType,
					"subnet_description": types.StringType,
					"subnet_name":        types.StringType,
					"subnet_uuid":        types.StringType,
					"url":                types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"hypervisor_hostname": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the hypervisor hosting this instance",
			},
			"image_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"internal_ips": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"is_limit_based": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"is_usage_based": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"key_fingerprint": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"key_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"latitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"longitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_category_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_category_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_offering_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_offering_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_resource_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"min_disk": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Minimum disk size in MiB",
			},
			"min_ram": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Minimum memory size in MiB",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"ports": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"mac_address": types.StringType,
					}}},
					"device_id":    types.StringType,
					"device_owner": types.StringType,
					"fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"ip_address": types.StringType,
						"subnet_id":  types.StringType,
					}}},
					"mac_address": types.StringType,
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
					"subnet":             types.StringType,
					"subnet_cidr":        types.StringType,
					"subnet_description": types.StringType,
					"subnet_name":        types.StringType,
					"subnet_uuid":        types.StringType,
					"url":                types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"ram": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Memory size in MiB",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"security_groups": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
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
					"url":   types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"server_group": schema.ObjectAttribute{
				CustomType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"name":   types.StringType,
					"policy": types.StringType,
					"state":  types.StringType,
					"url":    types.StringType,
				}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_settings": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "OpenStack provider settings",
			},
			"service_settings_error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_settings_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"start_time": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"user_data": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Additional data that will be added to instance on provisioning",
			},
			"volumes": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
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
				}}},
				Computed:            true,
				MarkdownDescription: "List of volumes attached to the instance",
			},
		},
	}
}

func (d *OpenstackInstanceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
}

func (d *OpenstackInstanceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackInstanceDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var apiResp OpenstackInstanceApiResponse

		err := d.client.GetByUUID(ctx, "/api/openstack-instances/{uuid}/", data.UUID.ValueString(), &apiResp)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Instance",
				"An error occurred while reading the Openstack Instance by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, apiResp, &data)...)

	} else {
		// Filter by provided parameters
		var results []OpenstackInstanceApiResponse

		filters := map[string]string{}
		if !data.AttachVolumeUuid.IsNull() {
			filters["attach_volume_uuid"] = data.AttachVolumeUuid.ValueString()
		}
		if !data.AvailabilityZoneName.IsNull() {
			filters["availability_zone_name"] = data.AvailabilityZoneName.ValueString()
		}
		if !data.BackendId.IsNull() {
			filters["backend_id"] = data.BackendId.ValueString()
		}
		if !data.CanManage.IsNull() {
			filters["can_manage"] = fmt.Sprintf("%t", data.CanManage.ValueBool())
		}
		if !data.Customer.IsNull() {
			filters["customer"] = data.Customer.ValueString()
		}
		if !data.CustomerAbbreviation.IsNull() {
			filters["customer_abbreviation"] = data.CustomerAbbreviation.ValueString()
		}
		if !data.CustomerName.IsNull() {
			filters["customer_name"] = data.CustomerName.ValueString()
		}
		if !data.CustomerNativeName.IsNull() {
			filters["customer_native_name"] = data.CustomerNativeName.ValueString()
		}
		if !data.CustomerUuid.IsNull() {
			filters["customer_uuid"] = data.CustomerUuid.ValueString()
		}
		if !data.Description.IsNull() {
			filters["description"] = data.Description.ValueString()
		}
		if !data.ExternalIp.IsNull() {
			filters["external_ip"] = data.ExternalIp.ValueString()
		}
		if !data.Name.IsNull() {
			filters["name"] = data.Name.ValueString()
		}
		if !data.NameExact.IsNull() {
			filters["name_exact"] = data.NameExact.ValueString()
		}
		if !data.Project.IsNull() {
			filters["project"] = data.Project.ValueString()
		}
		if !data.ProjectName.IsNull() {
			filters["project_name"] = data.ProjectName.ValueString()
		}
		if !data.ProjectUuid.IsNull() {
			filters["project_uuid"] = data.ProjectUuid.ValueString()
		}
		if !data.Query.IsNull() {
			filters["query"] = data.Query.ValueString()
		}
		if !data.RuntimeState.IsNull() {
			filters["runtime_state"] = data.RuntimeState.ValueString()
		}
		if !data.ServiceSettingsName.IsNull() {
			filters["service_settings_name"] = data.ServiceSettingsName.ValueString()
		}
		if !data.ServiceSettingsUuid.IsNull() {
			filters["service_settings_uuid"] = data.ServiceSettingsUuid.ValueString()
		}
		if !data.State.IsNull() {
			filters["state"] = data.State.ValueString()
		}
		if !data.Tenant.IsNull() {
			filters["tenant"] = data.Tenant.ValueString()
		}
		if !data.TenantUuid.IsNull() {
			filters["tenant_uuid"] = data.TenantUuid.ValueString()
		}
		if !data.Uuid.IsNull() {
			filters["uuid"] = data.Uuid.ValueString()
		}

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_instance.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/openstack-instances/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Openstack Instance",
				"An error occurred while filtering Openstack Instance: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Instance Not Found",
				"No Openstack Instance found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Instances Found",
				fmt.Sprintf("Found %d Openstack Instances with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *OpenstackInstanceDataSource) mapResponseToModel(ctx context.Context, apiResp OpenstackInstanceApiResponse, model *OpenstackInstanceDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.Action = types.StringPointerValue(apiResp.Action)
	model.AvailabilityZone = types.StringPointerValue(apiResp.AvailabilityZone)
	model.AvailabilityZoneName = types.StringPointerValue(apiResp.AvailabilityZoneName)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.ConnectDirectlyToExternalNetwork = types.BoolPointerValue(apiResp.ConnectDirectlyToExternalNetwork)
	model.Cores = types.Int64PointerValue(apiResp.Cores)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Customer = types.StringPointerValue(apiResp.Customer)
	model.CustomerAbbreviation = types.StringPointerValue(apiResp.CustomerAbbreviation)
	model.CustomerName = types.StringPointerValue(apiResp.CustomerName)
	model.CustomerNativeName = types.StringPointerValue(apiResp.CustomerNativeName)
	model.CustomerUuid = types.StringPointerValue(apiResp.CustomerUuid)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.Disk = types.Int64PointerValue(apiResp.Disk)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.ExternalAddress, _ = types.ListValueFrom(ctx, types.StringType, apiResp.ExternalAddress)
	model.ExternalIps, _ = types.ListValueFrom(ctx, types.StringType, apiResp.ExternalIps)
	model.FlavorDisk = types.Int64PointerValue(apiResp.FlavorDisk)
	model.FlavorName = types.StringPointerValue(apiResp.FlavorName)
	listValFloatingIps, listDiagsFloatingIps := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"address": types.StringType,
		"port_fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
			"ip_address": types.StringType,
			"subnet_id":  types.StringType,
		}}},
		"port_mac_address":   types.StringType,
		"subnet":             types.StringType,
		"subnet_cidr":        types.StringType,
		"subnet_description": types.StringType,
		"subnet_name":        types.StringType,
		"subnet_uuid":        types.StringType,
		"url":                types.StringType,
	}}, apiResp.FloatingIps)
	diags.Append(listDiagsFloatingIps...)
	model.FloatingIps = listValFloatingIps
	model.HypervisorHostname = types.StringPointerValue(apiResp.HypervisorHostname)
	model.ImageName = types.StringPointerValue(apiResp.ImageName)
	model.InternalIps, _ = types.ListValueFrom(ctx, types.StringType, apiResp.InternalIps)
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
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.Name = types.StringPointerValue(apiResp.Name)
	listValPorts, listDiagsPorts := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
			"mac_address": types.StringType,
		}}},
		"device_id":    types.StringType,
		"device_owner": types.StringType,
		"fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
			"ip_address": types.StringType,
			"subnet_id":  types.StringType,
		}}},
		"mac_address": types.StringType,
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
		"subnet":             types.StringType,
		"subnet_cidr":        types.StringType,
		"subnet_description": types.StringType,
		"subnet_name":        types.StringType,
		"subnet_uuid":        types.StringType,
		"url":                types.StringType,
	}}, apiResp.Ports)
	diags.Append(listDiagsPorts...)
	model.Ports = listValPorts
	model.Project = types.StringPointerValue(apiResp.Project)
	model.ProjectName = types.StringPointerValue(apiResp.ProjectName)
	model.ProjectUuid = types.StringPointerValue(apiResp.ProjectUuid)
	model.Ram = types.Int64PointerValue(apiResp.Ram)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.RuntimeState = types.StringPointerValue(apiResp.RuntimeState)
	listValSecurityGroups, listDiagsSecurityGroups := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
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
		"url":   types.StringType,
	}}, apiResp.SecurityGroups)
	diags.Append(listDiagsSecurityGroups...)
	model.SecurityGroups = listValSecurityGroups
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
	model.StartTime = types.StringPointerValue(apiResp.StartTime)
	model.State = types.StringPointerValue(apiResp.State)
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
	model.Url = types.StringPointerValue(apiResp.Url)
	model.UserData = types.StringPointerValue(apiResp.UserData)
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

	return diags
}
