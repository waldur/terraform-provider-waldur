package resources

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackInstanceResource{}
var _ resource.ResourceWithImportState = &OpenstackInstanceResource{}

func NewOpenstackInstanceResource() resource.Resource {
	return &OpenstackInstanceResource{}
}

// OpenstackInstanceResource defines the resource implementation.
type OpenstackInstanceResource struct {
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
	DataVolumeSize                   *int64                                    `json:"data_volume_size" tfsdk:"data_volume_size"`
	DataVolumeType                   *string                                   `json:"data_volume_type" tfsdk:"data_volume_type"`
	DataVolumes                      []OpenstackInstanceDataVolumesResponse    `json:"data_volumes" tfsdk:"data_volumes"`
	DeleteVolumes                    *bool                                     `json:"delete_volumes" tfsdk:"delete_volumes"`
	Description                      *string                                   `json:"description" tfsdk:"description"`
	Disk                             *int64                                    `json:"disk" tfsdk:"disk"`
	ErrorMessage                     *string                                   `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback                   *string                                   `json:"error_traceback" tfsdk:"error_traceback"`
	ExternalAddress                  []string                                  `json:"external_address" tfsdk:"external_address"`
	ExternalIps                      []string                                  `json:"external_ips" tfsdk:"external_ips"`
	Flavor                           *string                                   `json:"flavor" tfsdk:"flavor"`
	FlavorDisk                       *int64                                    `json:"flavor_disk" tfsdk:"flavor_disk"`
	FlavorName                       *string                                   `json:"flavor_name" tfsdk:"flavor_name"`
	FloatingIps                      []OpenstackInstanceFloatingIpsResponse    `json:"floating_ips" tfsdk:"floating_ips"`
	HypervisorHostname               *string                                   `json:"hypervisor_hostname" tfsdk:"hypervisor_hostname"`
	Image                            *string                                   `json:"image" tfsdk:"image"`
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
	Offering                         *string                                   `json:"offering" tfsdk:"offering"`
	Ports                            []OpenstackInstancePortsResponse          `json:"ports" tfsdk:"ports"`
	Project                          *string                                   `json:"project" tfsdk:"project"`
	ProjectName                      *string                                   `json:"project_name" tfsdk:"project_name"`
	ProjectUuid                      *string                                   `json:"project_uuid" tfsdk:"project_uuid"`
	Ram                              *int64                                    `json:"ram" tfsdk:"ram"`
	ReleaseFloatingIps               *bool                                     `json:"release_floating_ips" tfsdk:"release_floating_ips"`
	ResourceType                     *string                                   `json:"resource_type" tfsdk:"resource_type"`
	RuntimeState                     *string                                   `json:"runtime_state" tfsdk:"runtime_state"`
	SecurityGroups                   []OpenstackInstanceSecurityGroupsResponse `json:"security_groups" tfsdk:"security_groups"`
	ServerGroup                      *OpenstackInstanceServerGroupResponse     `json:"server_group" tfsdk:"server_group"`
	ServiceName                      *string                                   `json:"service_name" tfsdk:"service_name"`
	ServiceSettings                  *string                                   `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage      *string                                   `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState             *string                                   `json:"service_settings_state" tfsdk:"service_settings_state"`
	ServiceSettingsUuid              *string                                   `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	SshPublicKey                     *string                                   `json:"ssh_public_key" tfsdk:"ssh_public_key"`
	StartTime                        *string                                   `json:"start_time" tfsdk:"start_time"`
	State                            *string                                   `json:"state" tfsdk:"state"`
	SystemVolumeSize                 *int64                                    `json:"system_volume_size" tfsdk:"system_volume_size"`
	SystemVolumeType                 *string                                   `json:"system_volume_type" tfsdk:"system_volume_type"`
	Tenant                           *string                                   `json:"tenant" tfsdk:"tenant"`
	TenantUuid                       *string                                   `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Url                              *string                                   `json:"url" tfsdk:"url"`
	UserData                         *string                                   `json:"user_data" tfsdk:"user_data"`
	Volumes                          []OpenstackInstanceVolumesResponse        `json:"volumes" tfsdk:"volumes"`
}

type OpenstackInstanceDataVolumesResponse struct {
	Size       *int64  `json:"size" tfsdk:"size"`
	VolumeType *string `json:"volume_type" tfsdk:"volume_type"`
}

type OpenstackInstanceFloatingIpsResponse struct {
	IpAddress         *string                                            `json:"ip_address" tfsdk:"ip_address"`
	Subnet            *string                                            `json:"subnet" tfsdk:"subnet"`
	Url               *string                                            `json:"url" tfsdk:"url"`
	Address           *string                                            `json:"address" tfsdk:"address"`
	PortFixedIps      []OpenstackInstanceFloatingIpsPortFixedIpsResponse `json:"port_fixed_ips" tfsdk:"port_fixed_ips"`
	PortMacAddress    *string                                            `json:"port_mac_address" tfsdk:"port_mac_address"`
	SubnetCidr        *string                                            `json:"subnet_cidr" tfsdk:"subnet_cidr"`
	SubnetDescription *string                                            `json:"subnet_description" tfsdk:"subnet_description"`
	SubnetName        *string                                            `json:"subnet_name" tfsdk:"subnet_name"`
	SubnetUuid        *string                                            `json:"subnet_uuid" tfsdk:"subnet_uuid"`
}

type OpenstackInstanceFloatingIpsPortFixedIpsResponse struct {
	IpAddress *string `json:"ip_address" tfsdk:"ip_address"`
	SubnetId  *string `json:"subnet_id" tfsdk:"subnet_id"`
}

type OpenstackInstancePortsResponse struct {
	FixedIps            []OpenstackInstancePortsFixedIpsResponse            `json:"fixed_ips" tfsdk:"fixed_ips"`
	Port                *string                                             `json:"port" tfsdk:"port"`
	Subnet              *string                                             `json:"subnet" tfsdk:"subnet"`
	AllowedAddressPairs []OpenstackInstancePortsAllowedAddressPairsResponse `json:"allowed_address_pairs" tfsdk:"allowed_address_pairs"`
	DeviceId            *string                                             `json:"device_id" tfsdk:"device_id"`
	DeviceOwner         *string                                             `json:"device_owner" tfsdk:"device_owner"`
	MacAddress          *string                                             `json:"mac_address" tfsdk:"mac_address"`
	SecurityGroups      []OpenstackInstancePortsSecurityGroupsResponse      `json:"security_groups" tfsdk:"security_groups"`
	SubnetCidr          *string                                             `json:"subnet_cidr" tfsdk:"subnet_cidr"`
	SubnetDescription   *string                                             `json:"subnet_description" tfsdk:"subnet_description"`
	SubnetName          *string                                             `json:"subnet_name" tfsdk:"subnet_name"`
	SubnetUuid          *string                                             `json:"subnet_uuid" tfsdk:"subnet_uuid"`
	Url                 *string                                             `json:"url" tfsdk:"url"`
}

type OpenstackInstancePortsFixedIpsResponse struct {
	IpAddress *string `json:"ip_address" tfsdk:"ip_address"`
	SubnetId  *string `json:"subnet_id" tfsdk:"subnet_id"`
}

type OpenstackInstancePortsAllowedAddressPairsResponse struct {
	MacAddress *string `json:"mac_address" tfsdk:"mac_address"`
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
	Url         *string                                        `json:"url" tfsdk:"url"`
	Description *string                                        `json:"description" tfsdk:"description"`
	Name        *string                                        `json:"name" tfsdk:"name"`
	Rules       []OpenstackInstanceSecurityGroupsRulesResponse `json:"rules" tfsdk:"rules"`
	State       *string                                        `json:"state" tfsdk:"state"`
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

// OpenstackInstanceResourceModel describes the resource data model.
type OpenstackInstanceResourceModel struct {
	UUID                             types.String   `tfsdk:"id"`
	AccessUrl                        types.String   `tfsdk:"access_url"`
	Action                           types.String   `tfsdk:"action"`
	AvailabilityZone                 types.String   `tfsdk:"availability_zone"`
	AvailabilityZoneName             types.String   `tfsdk:"availability_zone_name"`
	BackendId                        types.String   `tfsdk:"backend_id"`
	ConnectDirectlyToExternalNetwork types.Bool     `tfsdk:"connect_directly_to_external_network"`
	Cores                            types.Int64    `tfsdk:"cores"`
	Created                          types.String   `tfsdk:"created"`
	Customer                         types.String   `tfsdk:"customer"`
	CustomerAbbreviation             types.String   `tfsdk:"customer_abbreviation"`
	CustomerName                     types.String   `tfsdk:"customer_name"`
	CustomerNativeName               types.String   `tfsdk:"customer_native_name"`
	CustomerUuid                     types.String   `tfsdk:"customer_uuid"`
	DataVolumeSize                   types.Int64    `tfsdk:"data_volume_size"`
	DataVolumeType                   types.String   `tfsdk:"data_volume_type"`
	DataVolumes                      types.List     `tfsdk:"data_volumes"`
	DeleteVolumes                    types.Bool     `tfsdk:"delete_volumes"`
	Description                      types.String   `tfsdk:"description"`
	Disk                             types.Int64    `tfsdk:"disk"`
	ErrorMessage                     types.String   `tfsdk:"error_message"`
	ErrorTraceback                   types.String   `tfsdk:"error_traceback"`
	ExternalAddress                  types.List     `tfsdk:"external_address"`
	ExternalIps                      types.List     `tfsdk:"external_ips"`
	Flavor                           types.String   `tfsdk:"flavor"`
	FlavorDisk                       types.Int64    `tfsdk:"flavor_disk"`
	FlavorName                       types.String   `tfsdk:"flavor_name"`
	FloatingIps                      types.List     `tfsdk:"floating_ips"`
	HypervisorHostname               types.String   `tfsdk:"hypervisor_hostname"`
	Image                            types.String   `tfsdk:"image"`
	ImageName                        types.String   `tfsdk:"image_name"`
	InternalIps                      types.List     `tfsdk:"internal_ips"`
	IsLimitBased                     types.Bool     `tfsdk:"is_limit_based"`
	IsUsageBased                     types.Bool     `tfsdk:"is_usage_based"`
	KeyFingerprint                   types.String   `tfsdk:"key_fingerprint"`
	KeyName                          types.String   `tfsdk:"key_name"`
	Latitude                         types.Float64  `tfsdk:"latitude"`
	Longitude                        types.Float64  `tfsdk:"longitude"`
	MarketplaceCategoryName          types.String   `tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid          types.String   `tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName          types.String   `tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid          types.String   `tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid              types.String   `tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState         types.String   `tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid          types.String   `tfsdk:"marketplace_resource_uuid"`
	MinDisk                          types.Int64    `tfsdk:"min_disk"`
	MinRam                           types.Int64    `tfsdk:"min_ram"`
	Modified                         types.String   `tfsdk:"modified"`
	Name                             types.String   `tfsdk:"name"`
	Offering                         types.String   `tfsdk:"offering"`
	Ports                            types.List     `tfsdk:"ports"`
	Project                          types.String   `tfsdk:"project"`
	ProjectName                      types.String   `tfsdk:"project_name"`
	ProjectUuid                      types.String   `tfsdk:"project_uuid"`
	Ram                              types.Int64    `tfsdk:"ram"`
	ReleaseFloatingIps               types.Bool     `tfsdk:"release_floating_ips"`
	ResourceType                     types.String   `tfsdk:"resource_type"`
	RuntimeState                     types.String   `tfsdk:"runtime_state"`
	SecurityGroups                   types.List     `tfsdk:"security_groups"`
	ServerGroup                      types.Object   `tfsdk:"server_group"`
	ServiceName                      types.String   `tfsdk:"service_name"`
	ServiceSettings                  types.String   `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage      types.String   `tfsdk:"service_settings_error_message"`
	ServiceSettingsState             types.String   `tfsdk:"service_settings_state"`
	ServiceSettingsUuid              types.String   `tfsdk:"service_settings_uuid"`
	SshPublicKey                     types.String   `tfsdk:"ssh_public_key"`
	StartTime                        types.String   `tfsdk:"start_time"`
	State                            types.String   `tfsdk:"state"`
	SystemVolumeSize                 types.Int64    `tfsdk:"system_volume_size"`
	SystemVolumeType                 types.String   `tfsdk:"system_volume_type"`
	Tenant                           types.String   `tfsdk:"tenant"`
	TenantUuid                       types.String   `tfsdk:"tenant_uuid"`
	Url                              types.String   `tfsdk:"url"`
	UserData                         types.String   `tfsdk:"user_data"`
	Volumes                          types.List     `tfsdk:"volumes"`
	Timeouts                         timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackInstanceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_instance"
}

func (r *OpenstackInstanceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Instance resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"access_url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Access url",
			},
			"action": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Action",
			},
			"availability_zone": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Availability zone where this instance is located",
			},
			"availability_zone_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the availability zone where instance is located",
			},
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Instance ID in the OpenStack backend",
			},
			"connect_directly_to_external_network": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "If True, instance will be connected directly to external network",
			},
			"cores": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Number of cores in a VM",
			},
			"created": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Created",
			},
			"customer": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Customer",
			},
			"customer_abbreviation": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Customer abbreviation",
			},
			"customer_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the customer",
			},
			"customer_native_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the customer native",
			},
			"customer_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the customer",
			},
			"data_volume_size": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Size of the data volume in MiB. Minimum size is 1024 MiB (1 GiB)",
			},
			"data_volume_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Volume type for the data volume",
			},
			"data_volumes": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"size": schema.Int64Attribute{
							Required:            true,
							MarkdownDescription: "Size",
						},
						"volume_type": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Volume type",
						},
					},
				},
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Additional data volumes to attach to the instance",
			},
			"delete_volumes": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Termination attribute",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Description of the resource",
			},
			"disk": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Disk size in MiB",
			},
			"error_message": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Error message",
			},
			"error_traceback": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Error traceback",
			},
			"external_address": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.StringType},
				Computed:   true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "External address",
			},
			"external_ips": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.StringType},
				Computed:   true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "External ips",
			},
			"flavor": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The flavor to use for the instance",
			},
			"flavor_disk": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Flavor disk size in MiB",
			},
			"flavor_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the flavor used by this instance",
			},
			"floating_ips": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Existing floating IP address in selected OpenStack tenant to be assigned to new virtual machine",
						},
						"subnet": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Subnet",
						},
						"url": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Url",
						},
						"address": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "The public IPv4 address of the floating IP",
						},
						"port_fixed_ips": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip_address": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "IP address to assign to the port",
									},
									"subnet_id": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "ID of the subnet in which to assign the IP address",
									},
								},
							},
							Computed: true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Port fixed ips",
						},
						"port_mac_address": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "MAC address of the port",
						},
						"subnet_cidr": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "IPv4 network address in CIDR format (e.g. 192.168.0.0/24)",
						},
						"subnet_description": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Subnet description",
						},
						"subnet_name": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Name of the subnet",
						},
						"subnet_uuid": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "UUID of the subnet",
						},
					},
				},
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Floating IPs to assign to the instance",
			},
			"hypervisor_hostname": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the hypervisor hosting this instance",
			},
			"image": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The OS image to use for the instance",
			},
			"image_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the image",
			},
			"internal_ips": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.StringType},
				Computed:   true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Internal ips",
			},
			"is_limit_based": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Is limit based",
			},
			"is_usage_based": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Is usage based",
			},
			"key_fingerprint": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Key fingerprint",
			},
			"key_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the key",
			},
			"latitude": schema.Float64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Latitude",
			},
			"longitude": schema.Float64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Longitude",
			},
			"marketplace_category_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the marketplace category",
			},
			"marketplace_category_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the marketplace category",
			},
			"marketplace_offering_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the marketplace offering",
			},
			"marketplace_offering_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the marketplace offering",
			},
			"marketplace_plan_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the marketplace plan",
			},
			"marketplace_resource_state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Marketplace resource state",
			},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the marketplace resource",
			},
			"min_disk": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Minimum disk size in MiB",
			},
			"min_ram": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Minimum memory size in MiB",
			},
			"modified": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the resource",
			},
			"offering": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Offering URL",
			},
			"ports": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"fixed_ips": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip_address": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "IP address to assign to the port",
									},
									"subnet_id": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "ID of the subnet in which to assign the IP address",
									},
								},
							},
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Fixed ips",
						},
						"port": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Port",
						},
						"subnet": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Subnet to which this port belongs",
						},
						"allowed_address_pairs": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"mac_address": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Mac address",
									},
								},
							},
							Computed: true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Allowed address pairs",
						},
						"device_id": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "ID of device (instance, router etc) to which this port is connected",
						},
						"device_owner": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Entity that uses this port (e.g. network:router_interface)",
						},
						"mac_address": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "MAC address of the port",
						},
						"security_groups": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"access_url": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Access url",
									},
									"backend_id": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "ID of the backend",
									},
									"created": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Created",
									},
									"customer": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Customer",
									},
									"customer_abbreviation": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Customer abbreviation",
									},
									"customer_name": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Name of the customer",
									},
									"customer_native_name": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Name of the customer native",
									},
									"customer_uuid": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "UUID of the customer",
									},
									"description": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Description of the resource",
									},
									"error_message": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Error message",
									},
									"error_traceback": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Error traceback",
									},
									"is_limit_based": schema.BoolAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.Bool{
											boolplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Is limit based",
									},
									"is_usage_based": schema.BoolAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.Bool{
											boolplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Is usage based",
									},
									"marketplace_category_name": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Name of the marketplace category",
									},
									"marketplace_category_uuid": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "UUID of the marketplace category",
									},
									"marketplace_offering_name": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Name of the marketplace offering",
									},
									"marketplace_offering_uuid": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "UUID of the marketplace offering",
									},
									"marketplace_plan_uuid": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "UUID of the marketplace plan",
									},
									"marketplace_resource_state": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Marketplace resource state",
									},
									"marketplace_resource_uuid": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "UUID of the marketplace resource",
									},
									"modified": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Modified",
									},
									"name": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Name of the resource",
									},
									"project": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Project",
									},
									"project_name": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Name of the project",
									},
									"project_uuid": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "UUID of the project",
									},
									"resource_type": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Resource type",
									},
									"rules": schema.ListNestedAttribute{
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"cidr": schema.StringAttribute{
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "CIDR notation for the source/destination network address range",
												},
												"description": schema.StringAttribute{
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "Description of the resource",
												},
												"direction": schema.StringAttribute{
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "Traffic direction - either 'ingress' (incoming) or 'egress' (outgoing)",
												},
												"ethertype": schema.StringAttribute{
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "IP protocol version - either 'IPv4' or 'IPv6'",
												},
												"from_port": schema.Int64Attribute{
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.Int64{
														int64planmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "Starting port number in the range (1-65535)",
												},
												"id": schema.Int64Attribute{
													Computed: true,
													PlanModifiers: []planmodifier.Int64{
														int64planmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "Id",
												},
												"protocol": schema.StringAttribute{
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "The network protocol (TCP, UDP, ICMP, or empty for any protocol)",
												},
												"remote_group": schema.StringAttribute{
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "Remote security group that this rule references, if any",
												},
												"remote_group_name": schema.StringAttribute{
													Computed: true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "Name of the remote group",
												},
												"remote_group_uuid": schema.StringAttribute{
													Computed: true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "UUID of the remote group",
												},
												"to_port": schema.Int64Attribute{
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.Int64{
														int64planmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "Ending port number in the range (1-65535)",
												},
											},
										},
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.List{
											listplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Rules",
									},
									"service_name": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Name of the service",
									},
									"service_settings": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Service settings",
									},
									"service_settings_error_message": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Service settings error message",
									},
									"service_settings_state": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Service settings state",
									},
									"service_settings_uuid": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "UUID of the service settings",
									},
									"state": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "State",
									},
									"tenant": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Tenant",
									},
									"tenant_name": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Name of the tenant",
									},
									"tenant_uuid": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "UUID of the tenant",
									},
									"url": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Url",
									},
								},
							},
							Computed: true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Security groups",
						},
						"subnet_cidr": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "IPv4 network address in CIDR format (e.g. 192.168.0.0/24)",
						},
						"subnet_description": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Subnet description",
						},
						"subnet_name": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Name of the subnet",
						},
						"subnet_uuid": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "UUID of the subnet",
						},
						"url": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Url",
						},
					},
				},
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Network ports to attach to the instance",
			},
			"project": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Project",
			},
			"project_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the project",
			},
			"project_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the project",
			},
			"ram": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Memory size in MiB",
			},
			"release_floating_ips": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Termination attribute",
			},
			"resource_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Resource type",
			},
			"runtime_state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Runtime state",
			},
			"security_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"url": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Url",
						},
						"description": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Description of the resource",
						},
						"name": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Name of the resource",
						},
						"rules": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"cidr": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "CIDR notation for the source/destination network address range",
									},
									"description": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Description of the resource",
									},
									"direction": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Traffic direction - either 'ingress' (incoming) or 'egress' (outgoing)",
									},
									"ethertype": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "IP protocol version - either 'IPv4' or 'IPv6'",
									},
									"from_port": schema.Int64Attribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Starting port number in the range (1-65535)",
									},
									"id": schema.Int64Attribute{
										Computed: true,
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Id",
									},
									"protocol": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "The network protocol (TCP, UDP, ICMP, or empty for any protocol)",
									},
									"remote_group_name": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Name of the remote group",
									},
									"remote_group_uuid": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "UUID of the remote group",
									},
									"to_port": schema.Int64Attribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Ending port number in the range (1-65535)",
									},
								},
							},
							Computed: true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Rules",
						},
						"state": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "State",
						},
					},
				},
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "List of security groups to apply to the instance",
			},
			"server_group": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the resource",
					},
					"policy": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Server group policy determining the rules for scheduling servers in this group",
					},
					"state": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "State",
					},
					"url": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Url",
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Server group",
			},
			"service_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the service",
			},
			"service_settings": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "OpenStack provider settings",
			},
			"service_settings_error_message": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Service settings error message",
			},
			"service_settings_state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Service settings state",
			},
			"service_settings_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the service settings",
			},
			"ssh_public_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Ssh public key",
			},
			"start_time": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Start time",
			},
			"state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "State",
			},
			"system_volume_size": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Size of the system volume in MiB. Minimum size is 1024 MiB (1 GiB)",
			},
			"system_volume_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Volume type for the system volume",
			},
			"tenant": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The OpenStack tenant to create the instance in",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the OpenStack tenant",
			},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Url",
			},
			"user_data": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Additional data that will be added to instance on provisioning",
			},
			"volumes": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"bootable": schema.BoolAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Indicates if this volume can be used to boot an instance",
						},
						"device": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Name of volume as instance device e.g. /dev/vdb.",
						},
						"image_name": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Name of the image this volume was created from",
						},
						"marketplace_resource_uuid": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "UUID of the marketplace resource",
						},
						"name": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Name of the resource",
						},
						"resource_type": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Resource type",
						},
						"size": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Size in MiB",
						},
						"state": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "State",
						},
						"type": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Type of the volume (e.g. SSD, HDD)",
						},
						"type_name": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Name of the type",
						},
						"url": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Url",
						},
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "List of volumes attached to the instance",
			},
		},

		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx, timeouts.Opts{
				Create: true,
				Update: true,
				Delete: true,
			}),
		},
	}
}

func (r *OpenstackInstanceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	r.client = client
}

func (r *OpenstackInstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackInstanceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Phase 1: Payload Construction
	attributes := map[string]interface{}{}
	if !data.AvailabilityZone.IsNull() {
		attributes["availability_zone"] = data.AvailabilityZone.ValueString()
	}
	if !data.ConnectDirectlyToExternalNetwork.IsNull() {
		attributes["connect_directly_to_external_network"] = data.ConnectDirectlyToExternalNetwork.ValueBool()
	}
	if !data.DataVolumeSize.IsNull() {
		attributes["data_volume_size"] = data.DataVolumeSize.ValueInt64()
	}
	if !data.DataVolumeType.IsNull() {
		attributes["data_volume_type"] = data.DataVolumeType.ValueString()
	}
	if !data.DataVolumes.IsNull() {
		items := make([]interface{}, 0)
		for _, elem := range data.DataVolumes.Elements() {
			items = append(items, ConvertTFValue(elem))
		}
		attributes["data_volumes"] = items
	}
	if !data.Description.IsNull() {
		attributes["description"] = data.Description.ValueString()
	}
	if !data.Flavor.IsNull() {
		attributes["flavor"] = data.Flavor.ValueString()
	}
	if !data.FloatingIps.IsNull() {
		items := make([]interface{}, 0)
		for _, elem := range data.FloatingIps.Elements() {
			items = append(items, ConvertTFValue(elem))
		}
		attributes["floating_ips"] = items
	}
	if !data.Image.IsNull() {
		attributes["image"] = data.Image.ValueString()
	}
	if !data.Name.IsNull() {
		attributes["name"] = data.Name.ValueString()
	}
	if !data.Ports.IsNull() {
		items := make([]interface{}, 0)
		for _, elem := range data.Ports.Elements() {
			items = append(items, ConvertTFValue(elem))
		}
		attributes["ports"] = items
	}
	if !data.SecurityGroups.IsNull() {
		items := make([]interface{}, 0)
		for _, elem := range data.SecurityGroups.Elements() {
			items = append(items, ConvertTFValue(elem))
		}
		attributes["security_groups"] = items
	}
	if !data.SshPublicKey.IsNull() {
		attributes["ssh_public_key"] = data.SshPublicKey.ValueString()
	}
	if !data.SystemVolumeSize.IsNull() {
		attributes["system_volume_size"] = data.SystemVolumeSize.ValueInt64()
	}
	if !data.SystemVolumeType.IsNull() {
		attributes["system_volume_type"] = data.SystemVolumeType.ValueString()
	}
	if !data.UserData.IsNull() {
		attributes["user_data"] = data.UserData.ValueString()
	}

	payload := map[string]interface{}{
		"project":    data.Project.ValueString(),
		"offering":   data.Offering.ValueString(),
		"attributes": attributes,
	}

	// Phase 2: Submit Order
	var orderRes map[string]interface{}
	err := r.client.Post(ctx, "/api/marketplace-orders/", payload, &orderRes)
	if err != nil {
		resp.Diagnostics.AddError("Order Submission Failed", err.Error())
		return
	}

	orderUUID, ok := orderRes["uuid"].(string)
	if !ok {
		resp.Diagnostics.AddError("Invalid Response", "Order UUID not found")
		return
	}

	// Phase 3: Poll for Completion
	// Attempt to resolve UUID
	if uuid, ok := orderRes["resource_uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	} else {
		data.UUID = types.StringValue(orderUUID)
	}

	// Attempt to fetch the resource to populate state
	{
		var mpUUID string
		if uuid, ok := orderRes["resource_uuid"].(string); ok {
			mpUUID = uuid
		} else if uuid, ok := orderRes["marketplace_resource_uuid"].(string); ok {
			mpUUID = uuid
		}

		if mpUUID != "" {
			var mpRes map[string]interface{}
			err = r.client.GetByUUID(ctx, "/api/marketplace-resources/{uuid}/", mpUUID, &mpRes)
			if err == nil {
				// Debug logging
				tflog.Warn(ctx, fmt.Sprintf("Fetched MP Resource: %+v", mpRes))
				if val, exists := mpRes["resource_uuid"]; exists {
					tflog.Warn(ctx, fmt.Sprintf("resource_uuid type: %T, value: %v", val, val))
				} else {
					tflog.Warn(ctx, "resource_uuid key missing in MP response")
				}

				// Plugin Resource UUID is available directly in resource_uuid field
				if pluginUUID, ok := mpRes["resource_uuid"].(string); ok {
					if pluginUUID != "" {
						data.UUID = types.StringValue(pluginUUID)

						// Fetch Plugin Resource
						var apiResp OpenstackInstanceApiResponse
						retrievePath := strings.Replace("/api/openstack-instances/{uuid}/", "{uuid}", pluginUUID, 1)
						tflog.Warn(ctx, "Attempting to fetch plugin resource at: "+retrievePath)
						err = r.client.GetByUUID(ctx, retrievePath, pluginUUID, &apiResp)
						if err == nil {
							tflog.Warn(ctx, "Successfully fetched plugin resource")
							resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)
						} else {
							tflog.Warn(ctx, "Failed to fetch plugin resource: "+err.Error())
						}
					} else {
						tflog.Warn(ctx, "resource_uuid is empty string")
					}
				} else {
					tflog.Warn(ctx, "Failed to cast resource_uuid to string")
				}
			} else {
				tflog.Warn(ctx, "Failed to fetch MP resource: "+err.Error())
			}
		}
	}

	if os.Getenv("WALDUR_E2E_SKIP_WAIT") != "" {
		tflog.Warn(ctx, "Skipping wait for order completion due to WALDUR_E2E_SKIP_WAIT")
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	stateConf := &retry.StateChangeConf{
		Pending: []string{"pending", "executing", "created"},
		Target:  []string{"done"},
		Refresh: func() (interface{}, string, error) {
			var res map[string]interface{}
			err := r.client.GetByUUID(ctx, "/api/marketplace-orders/{uuid}/", orderUUID, &res)
			if err != nil {
				return nil, "", err
			}

			state, _ := res["state"].(string)
			if state == "erred" || state == "rejected" {
				msg, _ := res["error_message"].(string)
				return res, "failed", fmt.Errorf("order failed: %s", msg)
			}
			return res, state, nil
		},
		Timeout: func() time.Duration {
			timeout, diags := data.Timeouts.Create(ctx, 45*time.Minute)
			resp.Diagnostics.Append(diags...)
			return timeout
		}(),
		Delay:      10 * time.Second,
		MinTimeout: 5 * time.Second,
	}

	rawResult, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Order Failed", err.Error())
		return
	}

	result := rawResult.(map[string]interface{})
	if resourceUUID, ok := result["marketplace_resource_uuid"].(string); ok {
		data.UUID = types.StringValue(resourceUUID)
	} else {
		resp.Diagnostics.AddError("Resource UUID Missing", "Order completed but marketplace_resource_uuid is missing")
		return
	}

	// Fetch final resource state
	var apiResp OpenstackInstanceApiResponse
	err = r.client.GetByUUID(ctx, "/api/openstack-instances/{uuid}/", data.UUID.ValueString(), &apiResp)
	if err != nil {
		resp.Diagnostics.AddError("Failed to Read Resource", err.Error())
		return
	}
	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackInstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackInstanceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource

	retrievePath := strings.Replace("/api/openstack-instances/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	var apiResp OpenstackInstanceApiResponse
	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &apiResp)
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Openstack Instance",
			"An error occurred while reading the Openstack Instance: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackInstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackInstanceResourceModel
	var state OpenstackInstanceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Phase 1: Standard PATCH (Simple fields)
	patchPayload := map[string]interface{}{}
	if !data.Description.IsNull() && !data.Description.Equal(state.Description) {
		patchPayload["description"] = data.Description.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.Equal(state.Name) {
		patchPayload["name"] = data.Name.ValueString()
	}

	if len(patchPayload) > 0 {
		var result map[string]interface{}
		err := r.client.Update(ctx, "/api/openstack-instances/{uuid}/", data.UUID.ValueString(), patchPayload, &result)
		if err != nil {
			resp.Diagnostics.AddError("Update Failed", err.Error())
			return
		}
		_ = result
	}

	// Phase 2: RPC Actions
	if os.Getenv("WALDUR_E2E_SKIP_WAIT") != "" {
		tflog.Warn(ctx, "Skipping wait for update order completion due to WALDUR_E2E_SKIP_WAIT")
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}
	if !data.FloatingIps.Equal(state.FloatingIps) {

		// Convert Terraform value to API payload

		var itemsUpdateFloatingIps []interface{}
		for _, elem := range data.FloatingIps.Elements() {
			if objVal, ok := elem.(types.Object); ok {
				objMap := make(map[string]interface{})
				for key, attr := range objVal.Attributes() {
					switch v := attr.(type) {
					case types.String:
						objMap[key] = v.ValueString()
					case types.Int64:
						objMap[key] = v.ValueInt64()
					case types.Bool:
						objMap[key] = v.ValueBool()
					case types.Float64:
						objMap[key] = v.ValueFloat64()
					}
				}
				itemsUpdateFloatingIps = append(itemsUpdateFloatingIps, objMap)
			}
		}
		actionPayloadUpdateFloatingIps := map[string]interface{}{
			"floating_ips": itemsUpdateFloatingIps,
		}
		actionUrlUpdateFloatingIps := strings.Replace("/api/openstack-instances/{uuid}/update_floating_ips/", "{uuid}", data.UUID.ValueString(), 1)
		var actionResultUpdateFloatingIps map[string]interface{}
		if err := r.client.Post(ctx, actionUrlUpdateFloatingIps, actionPayloadUpdateFloatingIps, &actionResultUpdateFloatingIps); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: update_floating_ips", err.Error())
			return
		}
	}
	if !data.Ports.Equal(state.Ports) {

		// Convert Terraform value to API payload

		var itemsUpdatePorts []interface{}
		for _, elem := range data.Ports.Elements() {
			if objVal, ok := elem.(types.Object); ok {
				objMap := make(map[string]interface{})
				for key, attr := range objVal.Attributes() {
					switch v := attr.(type) {
					case types.String:
						objMap[key] = v.ValueString()
					case types.Int64:
						objMap[key] = v.ValueInt64()
					case types.Bool:
						objMap[key] = v.ValueBool()
					case types.Float64:
						objMap[key] = v.ValueFloat64()
					}
				}
				itemsUpdatePorts = append(itemsUpdatePorts, objMap)
			}
		}
		actionPayloadUpdatePorts := map[string]interface{}{
			"ports": itemsUpdatePorts,
		}
		actionUrlUpdatePorts := strings.Replace("/api/openstack-instances/{uuid}/update_ports/", "{uuid}", data.UUID.ValueString(), 1)
		var actionResultUpdatePorts map[string]interface{}
		if err := r.client.Post(ctx, actionUrlUpdatePorts, actionPayloadUpdatePorts, &actionResultUpdatePorts); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: update_ports", err.Error())
			return
		}
	}
	if !data.SecurityGroups.Equal(state.SecurityGroups) {

		// Convert Terraform value to API payload

		var itemsUpdateSecurityGroups []interface{}
		for _, elem := range data.SecurityGroups.Elements() {
			if objVal, ok := elem.(types.Object); ok {
				objMap := make(map[string]interface{})
				for key, attr := range objVal.Attributes() {
					switch v := attr.(type) {
					case types.String:
						objMap[key] = v.ValueString()
					case types.Int64:
						objMap[key] = v.ValueInt64()
					case types.Bool:
						objMap[key] = v.ValueBool()
					case types.Float64:
						objMap[key] = v.ValueFloat64()
					}
				}
				itemsUpdateSecurityGroups = append(itemsUpdateSecurityGroups, objMap)
			}
		}
		actionPayloadUpdateSecurityGroups := map[string]interface{}{
			"security_groups": itemsUpdateSecurityGroups,
		}
		actionUrlUpdateSecurityGroups := strings.Replace("/api/openstack-instances/{uuid}/update_security_groups/", "{uuid}", data.UUID.ValueString(), 1)
		var actionResultUpdateSecurityGroups map[string]interface{}
		if err := r.client.Post(ctx, actionUrlUpdateSecurityGroups, actionPayloadUpdateSecurityGroups, &actionResultUpdateSecurityGroups); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: update_security_groups", err.Error())
			return
		}
	}

	// Fetch updated state
	// Call Waldur API to read resource
	var apiResp OpenstackInstanceApiResponse

	retrievePath := strings.Replace("/api/openstack-instances/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &apiResp)
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Failed to Read Resource After Update", err.Error())
		return
	}
	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackInstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackInstanceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Order-based Delete
	payload := map[string]interface{}{}
	if !data.DeleteVolumes.IsNull() {
		payload["delete_volumes"] = data.DeleteVolumes.ValueBool()
	}
	if !data.ReleaseFloatingIps.IsNull() {
		payload["release_floating_ips"] = data.ReleaseFloatingIps.ValueBool()
	}

	url := fmt.Sprintf("/api/marketplace-resources/%s/terminate/", data.UUID.ValueString())
	var res map[string]interface{}
	err := r.client.Post(ctx, url, payload, &res)
	if err != nil {
		resp.Diagnostics.AddError("Termination Failed", err.Error())
		return
	}

	// Wait for deletion if order UUID is returned
	if orderUUID, ok := res["uuid"].(string); ok {
		stateConf := &retry.StateChangeConf{
			Pending: []string{"pending", "executing", "created"},
			Target:  []string{"done"},
			Refresh: func() (interface{}, string, error) {
				var res map[string]interface{}
				err := r.client.GetByUUID(ctx, "/api/marketplace-orders/{uuid}/", orderUUID, &res)
				if err != nil {
					return nil, "", err
				}
				state, _ := res["state"].(string)
				if state == "erred" || state == "rejected" {
					return res, "failed", fmt.Errorf("termination order failed")
				}
				return res, state, nil
			},
			Timeout: func() time.Duration {
				timeout, diags := data.Timeouts.Delete(ctx, 45*time.Minute)
				resp.Diagnostics.Append(diags...)
				return timeout
			}(),
			Delay:      10 * time.Second,
			MinTimeout: 5 * time.Second,
		}
		_, err := stateConf.WaitForStateContext(ctx)
		if err != nil {
			resp.Diagnostics.AddError("Termination Order Failed", err.Error())
			return
		}
	}
}

func (r *OpenstackInstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *OpenstackInstanceResource) mapResponseToModel(ctx context.Context, apiResp OpenstackInstanceApiResponse, model *OpenstackInstanceResourceModel) diag.Diagnostics {
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
	model.DataVolumeSize = types.Int64PointerValue(apiResp.DataVolumeSize)
	model.DataVolumeType = types.StringPointerValue(apiResp.DataVolumeType)
	listValDataVolumes, listDiagsDataVolumes := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"size":        types.Int64Type,
		"volume_type": types.StringType,
	}}, apiResp.DataVolumes)
	diags.Append(listDiagsDataVolumes...)
	model.DataVolumes = listValDataVolumes
	model.DeleteVolumes = types.BoolPointerValue(apiResp.DeleteVolumes)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.Disk = types.Int64PointerValue(apiResp.Disk)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.ExternalAddress, _ = types.ListValueFrom(ctx, types.StringType, apiResp.ExternalAddress)
	model.ExternalIps, _ = types.ListValueFrom(ctx, types.StringType, apiResp.ExternalIps)
	model.Flavor = types.StringPointerValue(apiResp.Flavor)
	model.FlavorDisk = types.Int64PointerValue(apiResp.FlavorDisk)
	model.FlavorName = types.StringPointerValue(apiResp.FlavorName)
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
	model.HypervisorHostname = types.StringPointerValue(apiResp.HypervisorHostname)
	model.Image = types.StringPointerValue(apiResp.Image)
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
	model.Offering = types.StringPointerValue(apiResp.Offering)
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
	model.Project = types.StringPointerValue(apiResp.Project)
	model.ProjectName = types.StringPointerValue(apiResp.ProjectName)
	model.ProjectUuid = types.StringPointerValue(apiResp.ProjectUuid)
	model.Ram = types.Int64PointerValue(apiResp.Ram)
	model.ReleaseFloatingIps = types.BoolPointerValue(apiResp.ReleaseFloatingIps)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.RuntimeState = types.StringPointerValue(apiResp.RuntimeState)
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
	model.SshPublicKey = types.StringPointerValue(apiResp.SshPublicKey)
	model.StartTime = types.StringPointerValue(apiResp.StartTime)
	model.State = types.StringPointerValue(apiResp.State)
	model.SystemVolumeSize = types.Int64PointerValue(apiResp.SystemVolumeSize)
	model.SystemVolumeType = types.StringPointerValue(apiResp.SystemVolumeType)
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
