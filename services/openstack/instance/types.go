package instance

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackInstanceCreateRequest struct {
	Project    *string                           `json:"project"`
	Offering   *string                           `json:"offering"`
	Plan       *string                           `json:"plan,omitempty"`
	Limits     map[string]float64                `json:"limits,omitempty"`
	Attributes OpenstackInstanceCreateAttributes `json:"attributes"`
}
type OpenstackInstanceCreateAttributes struct {
	AvailabilityZone                 *string                                          `json:"availability_zone,omitempty"`
	ConnectDirectlyToExternalNetwork *bool                                            `json:"connect_directly_to_external_network,omitempty"`
	DataVolumeSize                   *int64                                           `json:"data_volume_size,omitempty"`
	DataVolumeType                   *string                                          `json:"data_volume_type,omitempty"`
	DataVolumes                      *[]common.OpenStackDataVolumeRequest             `json:"data_volumes,omitempty"`
	Description                      *string                                          `json:"description,omitempty"`
	Flavor                           *string                                          `json:"flavor,omitempty"`
	FloatingIps                      *[]common.OpenStackCreateFloatingIPRequest       `json:"floating_ips,omitempty"`
	Image                            *string                                          `json:"image,omitempty"`
	Name                             *string                                          `json:"name,omitempty"`
	Ports                            *[]common.OpenStackCreateInstancePortRequest     `json:"ports,omitempty"`
	SecurityGroups                   *[]common.OpenStackSecurityGroupHyperlinkRequest `json:"security_groups,omitempty"`
	ServerGroup                      *OpenstackInstanceCreateServerGroupRequest       `json:"server_group,omitempty"`
	SshPublicKey                     *string                                          `json:"ssh_public_key,omitempty"`
	SystemVolumeSize                 *int64                                           `json:"system_volume_size,omitempty"`
	SystemVolumeType                 *string                                          `json:"system_volume_type,omitempty"`
	UserData                         *string                                          `json:"user_data,omitempty"`
}

type OpenstackInstanceCreateServerGroupRequest struct {
	Url *string `json:"url" tfsdk:"url"`
}

type OpenstackInstanceUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type OpenstackInstanceUpdateSecurityGroupsActionRequest struct {
	SecurityGroups []common.OpenStackSecurityGroupHyperlinkRequest `json:"security_groups"`
}
type OpenstackInstanceUpdateFloatingIpsActionRequest struct {
	FloatingIps []common.OpenStackCreateFloatingIPRequest `json:"floating_ips"`
}
type OpenstackInstanceUpdatePortsActionRequest struct {
	Ports []common.OpenStackCreateInstancePortRequest `json:"ports"`
}

type OpenstackInstanceResponse struct {
	UUID *string `json:"uuid"`

	Action                           *string                                          `json:"action" tfsdk:"action"`
	AvailabilityZone                 *string                                          `json:"availability_zone" tfsdk:"availability_zone"`
	AvailabilityZoneName             *string                                          `json:"availability_zone_name" tfsdk:"availability_zone_name"`
	BackendId                        *string                                          `json:"backend_id" tfsdk:"backend_id"`
	ConnectDirectlyToExternalNetwork *bool                                            `json:"connect_directly_to_external_network" tfsdk:"connect_directly_to_external_network"`
	Cores                            *int64                                           `json:"cores" tfsdk:"cores"`
	Customer                         *string                                          `json:"customer" tfsdk:"customer"`
	Description                      *string                                          `json:"description" tfsdk:"description"`
	Disk                             *int64                                           `json:"disk" tfsdk:"disk"`
	ErrorMessage                     *string                                          `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback                   *string                                          `json:"error_traceback" tfsdk:"error_traceback"`
	ExternalAddress                  *[]string                                        `json:"external_address" tfsdk:"external_address"`
	ExternalIps                      *[]string                                        `json:"external_ips" tfsdk:"external_ips"`
	FlavorDisk                       *int64                                           `json:"flavor_disk" tfsdk:"flavor_disk"`
	FlavorName                       *string                                          `json:"flavor_name" tfsdk:"flavor_name"`
	FloatingIps                      *[]common.OpenStackCreateFloatingIPRequest       `json:"floating_ips" tfsdk:"floating_ips"`
	HypervisorHostname               *string                                          `json:"hypervisor_hostname" tfsdk:"hypervisor_hostname"`
	ImageName                        *string                                          `json:"image_name" tfsdk:"image_name"`
	InternalIps                      *[]string                                        `json:"internal_ips" tfsdk:"internal_ips"`
	KeyFingerprint                   *string                                          `json:"key_fingerprint" tfsdk:"key_fingerprint"`
	KeyName                          *string                                          `json:"key_name" tfsdk:"key_name"`
	Latitude                         *common.FlexibleNumber                           `json:"latitude" tfsdk:"latitude"`
	Longitude                        *common.FlexibleNumber                           `json:"longitude" tfsdk:"longitude"`
	MarketplaceResourceUuid          *string                                          `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	MinDisk                          *int64                                           `json:"min_disk" tfsdk:"min_disk"`
	MinRam                           *int64                                           `json:"min_ram" tfsdk:"min_ram"`
	Name                             *string                                          `json:"name" tfsdk:"name"`
	Ports                            *[]common.OpenStackCreateInstancePortRequest     `json:"ports" tfsdk:"ports"`
	Project                          *string                                          `json:"project" tfsdk:"project"`
	Ram                              *int64                                           `json:"ram" tfsdk:"ram"`
	RancherCluster                   *OpenstackInstanceRancherClusterResponse         `json:"rancher_cluster" tfsdk:"rancher_cluster"`
	ResourceType                     *string                                          `json:"resource_type" tfsdk:"resource_type"`
	RuntimeState                     *string                                          `json:"runtime_state" tfsdk:"runtime_state"`
	SecurityGroups                   *[]common.OpenStackSecurityGroupHyperlinkRequest `json:"security_groups" tfsdk:"security_groups"`
	ServerGroup                      *OpenstackInstanceServerGroupResponse            `json:"server_group" tfsdk:"server_group"`
	StartTime                        *string                                          `json:"start_time" tfsdk:"start_time"`
	State                            *string                                          `json:"state" tfsdk:"state"`
	Tenant                           *string                                          `json:"tenant" tfsdk:"tenant"`
	TenantUuid                       *string                                          `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Url                              *string                                          `json:"url" tfsdk:"url"`
	UserData                         *string                                          `json:"user_data" tfsdk:"user_data"`
	Volumes                          *[]common.OpenStackNestedVolume                  `json:"volumes" tfsdk:"volumes"`
}

type OpenstackInstanceFloatingIpsResponse struct {
	IpAddress         *string                    `json:"ip_address" tfsdk:"ip_address"`
	Subnet            *string                    `json:"subnet" tfsdk:"subnet"`
	Url               *string                    `json:"url" tfsdk:"url"`
	Address           *string                    `json:"address" tfsdk:"address"`
	PortFixedIps      *[]common.OpenStackFixedIp `json:"port_fixed_ips" tfsdk:"port_fixed_ips"`
	PortMacAddress    *string                    `json:"port_mac_address" tfsdk:"port_mac_address"`
	SubnetCidr        *string                    `json:"subnet_cidr" tfsdk:"subnet_cidr"`
	SubnetDescription *string                    `json:"subnet_description" tfsdk:"subnet_description"`
	SubnetName        *string                    `json:"subnet_name" tfsdk:"subnet_name"`
	SubnetUuid        *string                    `json:"subnet_uuid" tfsdk:"subnet_uuid"`
	Uuid              *string                    `json:"uuid" tfsdk:"uuid"`
}

type OpenstackInstanceFloatingIpsPortFixedIpsResponse struct {
	IpAddress *string `json:"ip_address" tfsdk:"ip_address"`
	SubnetId  *string `json:"subnet_id" tfsdk:"subnet_id"`
}

type OpenstackInstancePortsResponse struct {
	FixedIps            *[]common.OpenStackFixedIpRequest     `json:"fixed_ips" tfsdk:"fixed_ips"`
	Port                *string                               `json:"port" tfsdk:"port"`
	Subnet              *string                               `json:"subnet" tfsdk:"subnet"`
	AllowedAddressPairs *[]common.OpenStackAllowedAddressPair `json:"allowed_address_pairs" tfsdk:"allowed_address_pairs"`
	DeviceId            *string                               `json:"device_id" tfsdk:"device_id"`
	DeviceOwner         *string                               `json:"device_owner" tfsdk:"device_owner"`
	MacAddress          *string                               `json:"mac_address" tfsdk:"mac_address"`
	SecurityGroups      *[]common.OpenStackSecurityGroup      `json:"security_groups" tfsdk:"security_groups"`
	SubnetCidr          *string                               `json:"subnet_cidr" tfsdk:"subnet_cidr"`
	SubnetDescription   *string                               `json:"subnet_description" tfsdk:"subnet_description"`
	SubnetName          *string                               `json:"subnet_name" tfsdk:"subnet_name"`
	SubnetUuid          *string                               `json:"subnet_uuid" tfsdk:"subnet_uuid"`
	Url                 *string                               `json:"url" tfsdk:"url"`
}

type OpenstackInstancePortsFixedIpsResponse struct {
	IpAddress *string `json:"ip_address" tfsdk:"ip_address"`
	SubnetId  *string `json:"subnet_id" tfsdk:"subnet_id"`
}

type OpenstackInstancePortsAllowedAddressPairsResponse struct {
	MacAddress *string `json:"mac_address" tfsdk:"mac_address"`
}

type OpenstackInstancePortsSecurityGroupsResponse struct {
	BackendId               *string                                    `json:"backend_id" tfsdk:"backend_id"`
	Customer                *string                                    `json:"customer" tfsdk:"customer"`
	Description             *string                                    `json:"description" tfsdk:"description"`
	ErrorMessage            *string                                    `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback          *string                                    `json:"error_traceback" tfsdk:"error_traceback"`
	MarketplaceResourceUuid *string                                    `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Name                    *string                                    `json:"name" tfsdk:"name"`
	Project                 *string                                    `json:"project" tfsdk:"project"`
	ResourceType            *string                                    `json:"resource_type" tfsdk:"resource_type"`
	Rules                   *[]common.OpenStackSecurityGroupRuleCreate `json:"rules" tfsdk:"rules"`
	State                   *string                                    `json:"state" tfsdk:"state"`
	Tenant                  *string                                    `json:"tenant" tfsdk:"tenant"`
	TenantName              *string                                    `json:"tenant_name" tfsdk:"tenant_name"`
	TenantUuid              *string                                    `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Url                     *string                                    `json:"url" tfsdk:"url"`
	Uuid                    *string                                    `json:"uuid" tfsdk:"uuid"`
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

type OpenstackInstanceRancherClusterResponse struct {
	MarketplaceUuid *string `json:"marketplace_uuid" tfsdk:"marketplace_uuid"`
	Name            *string `json:"name" tfsdk:"name"`
	Uuid            *string `json:"uuid" tfsdk:"uuid"`
}

type OpenstackInstanceSecurityGroupsResponse struct {
	Url         *string                           `json:"url" tfsdk:"url"`
	Description *string                           `json:"description" tfsdk:"description"`
	Name        *string                           `json:"name" tfsdk:"name"`
	Rules       *[]common.NestedSecurityGroupRule `json:"rules" tfsdk:"rules"`
	State       *string                           `json:"state" tfsdk:"state"`
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
	Url    *string `json:"url" tfsdk:"url"`
	Name   *string `json:"name" tfsdk:"name"`
	Policy *string `json:"policy" tfsdk:"policy"`
	State  *string `json:"state" tfsdk:"state"`
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
	Uuid                    *string `json:"uuid" tfsdk:"uuid"`
}

func (r *OpenstackInstanceResponse) GetState() string {
	if r.State != nil {
		return *r.State
	}
	return "OK"
}

func (r *OpenstackInstanceResponse) GetErrorMessage() string {
	if r.ErrorMessage != nil {
		return *r.ErrorMessage
	}
	return ""
}
