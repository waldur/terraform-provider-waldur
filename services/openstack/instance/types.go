package instance

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// OpenstackInstance Structs

type OpenstackInstanceCreateRequest struct {
	Project    *string                           `json:"project" tfsdk:"project"`
	Offering   *string                           `json:"offering" tfsdk:"offering"`
	Attributes OpenstackInstanceCreateAttributes `json:"attributes" tfsdk:"attributes"`
}
type OpenstackInstanceCreateAttributes struct {
	AvailabilityZone                 *string                                         `json:"availability_zone,omitempty" tfsdk:"availability_zone"`
	ConnectDirectlyToExternalNetwork *bool                                           `json:"connect_directly_to_external_network,omitempty" tfsdk:"connect_directly_to_external_network"`
	DataVolumeSize                   *int64                                          `json:"data_volume_size,omitempty" tfsdk:"data_volume_size"`
	DataVolumeType                   *string                                         `json:"data_volume_type,omitempty" tfsdk:"data_volume_type"`
	DataVolumes                      []common.OpenStackDataVolumeRequest             `json:"data_volumes,omitempty" tfsdk:"data_volumes"`
	Description                      *string                                         `json:"description,omitempty" tfsdk:"description"`
	Flavor                           *string                                         `json:"flavor,omitempty" tfsdk:"flavor"`
	FloatingIps                      []common.OpenStackCreateFloatingIPRequest       `json:"floating_ips,omitempty" tfsdk:"floating_ips"`
	Image                            *string                                         `json:"image,omitempty" tfsdk:"image"`
	Name                             *string                                         `json:"name,omitempty" tfsdk:"name"`
	Ports                            []common.OpenStackCreateInstancePortRequest     `json:"ports,omitempty" tfsdk:"ports"`
	SecurityGroups                   []common.OpenStackSecurityGroupHyperlinkRequest `json:"security_groups,omitempty" tfsdk:"security_groups"`
	SshPublicKey                     *string                                         `json:"ssh_public_key,omitempty" tfsdk:"ssh_public_key"`
	SystemVolumeSize                 *int64                                          `json:"system_volume_size,omitempty" tfsdk:"system_volume_size"`
	SystemVolumeType                 *string                                         `json:"system_volume_type,omitempty" tfsdk:"system_volume_type"`
	UserData                         *string                                         `json:"user_data,omitempty" tfsdk:"user_data"`
}

type OpenstackInstanceUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type OpenstackInstanceUpdateFloatingIpsActionRequest struct {
	FloatingIps []common.OpenStackCreateFloatingIPRequest `json:"floating_ips" tfsdk:"floating_ips"`
}
type OpenstackInstanceUpdatePortsActionRequest struct {
	Ports []common.OpenStackCreateInstancePortRequest `json:"ports" tfsdk:"ports"`
}
type OpenstackInstanceUpdateSecurityGroupsActionRequest struct {
	SecurityGroups []common.OpenStackSecurityGroupHyperlinkRequest `json:"security_groups" tfsdk:"security_groups"`
}

type OpenstackInstanceResponse struct {
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
