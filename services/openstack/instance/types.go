package instance

import (
	"encoding/json"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackInstanceCreateRequest struct {
	EndDate    *string                           `json:"end_date,omitempty"`
	Limits     map[string]float64                `json:"limits,omitempty"`
	Offering   *string                           `json:"offering"`
	Plan       *string                           `json:"plan,omitempty"`
	Project    *string                           `json:"project"`
	StartDate  *string                           `json:"start_date,omitempty"`
	Attributes OpenstackInstanceCreateAttributes `json:"attributes"`
}
type OpenstackInstanceCreateAttributes struct {
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	ConnectDirectlyToExternalNetwork *bool `json:"connect_directly_to_external_network,omitempty"`

	DataVolumeSize *int64 `json:"data_volume_size,omitempty"`

	DataVolumeType *string `json:"data_volume_type,omitempty"`

	DataVolumes *[]common.OpenStackDataVolumeRequest `json:"data_volumes,omitempty"`

	Description *string `json:"description,omitempty"`

	EndDate *string `json:"end_date,omitempty"`

	Flavor *string `json:"flavor"`

	FloatingIps *[]common.OpenStackCreateFloatingIPRequest `json:"floating_ips,omitempty"`

	Image *string `json:"image"`

	Limits map[string]float64 `json:"limits,omitempty"`

	Name *string `json:"name"`

	Plan *string `json:"plan,omitempty"`

	Ports []common.OpenStackCreateInstancePortRequest `json:"ports"`

	SecurityGroups *[]common.OpenStackSecurityGroupHyperlinkRequest `json:"security_groups,omitempty"`

	ServerGroup *OpenstackInstanceCreateServerGroupRequest `json:"server_group,omitempty"`

	SshPublicKey *string `json:"ssh_public_key,omitempty"`

	StartDate *string `json:"start_date,omitempty"`

	SystemVolumeSize *int64 `json:"system_volume_size"`

	SystemVolumeType *string `json:"system_volume_type,omitempty"`

	UserData *string `json:"user_data,omitempty"`
}

type OpenstackInstanceCreateLimitsRequest struct {
}

type OpenstackInstanceCreateServerGroupRequest struct {
	Url *string `json:"url" tfsdk:"url"`
}

type OpenstackInstanceUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`

	Name *string `json:"name,omitempty" tfsdk:"name"`
}

type OpenstackInstanceUpdateFloatingIpsActionRequest struct {
	FloatingIps []common.OpenStackCreateFloatingIPRequest `json:"-"`
}

func (r OpenstackInstanceUpdateFloatingIpsActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.FloatingIps)
}

type OpenstackInstanceUpdatePortsActionRequest struct {
	Ports []common.OpenStackCreateInstancePortRequest `json:"-"`
}

func (r OpenstackInstanceUpdatePortsActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Ports)
}

type OpenstackInstanceUpdateSecurityGroupsActionRequest struct {
	SecurityGroups []common.OpenStackSecurityGroupHyperlinkRequest `json:"-"`
}

func (r OpenstackInstanceUpdateSecurityGroupsActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.SecurityGroups)
}

type OpenstackInstanceResponse struct {
	UUID *string `json:"uuid"`

	Action *string `json:"action,omitempty" tfsdk:"action"`

	AvailabilityZone *string `json:"availability_zone,omitempty" tfsdk:"availability_zone"`

	AvailabilityZoneName *string `json:"availability_zone_name,omitempty" tfsdk:"availability_zone_name"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	ConnectDirectlyToExternalNetwork *bool `json:"connect_directly_to_external_network,omitempty" tfsdk:"connect_directly_to_external_network"`

	Cores *int64 `json:"cores,omitempty" tfsdk:"cores"`

	Customer *string `json:"customer,omitempty" tfsdk:"customer"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	Disk *int64 `json:"disk,omitempty" tfsdk:"disk"`

	ErrorMessage *string `json:"error_message,omitempty" tfsdk:"error_message"`

	ExternalAddress *[]string `json:"external_address,omitempty" tfsdk:"external_address"`

	ExternalIps *[]string `json:"external_ips,omitempty" tfsdk:"external_ips"`

	FlavorDisk *int64 `json:"flavor_disk,omitempty" tfsdk:"flavor_disk"`

	FlavorName *string `json:"flavor_name,omitempty" tfsdk:"flavor_name"`

	FloatingIps *[]common.OpenStackCreateFloatingIPRequest `json:"floating_ips,omitempty" tfsdk:"floating_ips"`

	HypervisorHostname *string `json:"hypervisor_hostname,omitempty" tfsdk:"hypervisor_hostname"`

	ImageName *string `json:"image_name,omitempty" tfsdk:"image_name"`

	InternalIps *[]string `json:"internal_ips,omitempty" tfsdk:"internal_ips"`

	KeyFingerprint *string `json:"key_fingerprint,omitempty" tfsdk:"key_fingerprint"`

	KeyName *string `json:"key_name,omitempty" tfsdk:"key_name"`

	Latitude common.FlexibleNumber `json:"latitude,omitempty" tfsdk:"latitude"`

	Longitude common.FlexibleNumber `json:"longitude,omitempty" tfsdk:"longitude"`

	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`

	MinDisk *int64 `json:"min_disk,omitempty" tfsdk:"min_disk"`

	MinRam *int64 `json:"min_ram,omitempty" tfsdk:"min_ram"`

	Name *string `json:"name" tfsdk:"name"`

	Ports []common.OpenStackCreateInstancePortRequest `json:"ports" tfsdk:"ports"`

	Project *string `json:"project" tfsdk:"project"`

	Ram *int64 `json:"ram,omitempty" tfsdk:"ram"`

	RancherCluster *OpenstackInstanceRancherClusterResponse `json:"rancher_cluster,omitempty" tfsdk:"rancher_cluster"`

	ResourceType *string `json:"resource_type,omitempty" tfsdk:"resource_type"`

	RuntimeState *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`

	SecurityGroups *[]common.OpenStackSecurityGroupHyperlinkRequest `json:"security_groups,omitempty" tfsdk:"security_groups"`

	ServerGroup *OpenstackInstanceServerGroupResponse `json:"server_group,omitempty" tfsdk:"server_group"`

	StartTime *string `json:"start_time,omitempty" tfsdk:"start_time"`

	State *string `json:"state,omitempty" tfsdk:"state"`

	Tenant *string `json:"tenant,omitempty" tfsdk:"tenant"`

	TenantUuid *string `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	UserData *string `json:"user_data,omitempty" tfsdk:"user_data"`

	Volumes *[]common.OpenStackNestedVolume `json:"volumes,omitempty" tfsdk:"volumes"`
}

type OpenstackInstanceFloatingIpsResponse struct {
	IpAddress *string `json:"ip_address,omitempty" tfsdk:"ip_address"`

	Subnet *string `json:"subnet" tfsdk:"subnet"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	Address *string `json:"address,omitempty" tfsdk:"address"`

	PortFixedIps *[]common.OpenStackFixedIp `json:"port_fixed_ips,omitempty" tfsdk:"port_fixed_ips"`

	PortMacAddress *string `json:"port_mac_address,omitempty" tfsdk:"port_mac_address"`

	SubnetCidr *string `json:"subnet_cidr,omitempty" tfsdk:"subnet_cidr"`

	SubnetDescription *string `json:"subnet_description,omitempty" tfsdk:"subnet_description"`

	SubnetName *string `json:"subnet_name,omitempty" tfsdk:"subnet_name"`

	SubnetUuid *string `json:"subnet_uuid,omitempty" tfsdk:"subnet_uuid"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type OpenstackInstanceFloatingIpsPortFixedIpsResponse struct {
	IpAddress *string `json:"ip_address,omitempty" tfsdk:"ip_address"`

	SubnetId *string `json:"subnet_id,omitempty" tfsdk:"subnet_id"`
}

type OpenstackInstancePortsResponse struct {
	FixedIps *[]common.OpenStackFixedIpRequest `json:"fixed_ips,omitempty" tfsdk:"fixed_ips"`

	Port *string `json:"port,omitempty" tfsdk:"port"`

	Subnet *string `json:"subnet,omitempty" tfsdk:"subnet"`

	AllowedAddressPairs *[]common.OpenStackAllowedAddressPair `json:"allowed_address_pairs,omitempty" tfsdk:"allowed_address_pairs"`

	DeviceId *string `json:"device_id,omitempty" tfsdk:"device_id"`

	DeviceOwner *string `json:"device_owner,omitempty" tfsdk:"device_owner"`

	MacAddress *string `json:"mac_address,omitempty" tfsdk:"mac_address"`

	SecurityGroups *[]common.OpenStackSecurityGroup `json:"security_groups,omitempty" tfsdk:"security_groups"`

	SubnetCidr *string `json:"subnet_cidr,omitempty" tfsdk:"subnet_cidr"`

	SubnetDescription *string `json:"subnet_description,omitempty" tfsdk:"subnet_description"`

	SubnetName *string `json:"subnet_name,omitempty" tfsdk:"subnet_name"`

	SubnetUuid *string `json:"subnet_uuid,omitempty" tfsdk:"subnet_uuid"`

	Url *string `json:"url,omitempty" tfsdk:"url"`
}

type OpenstackInstancePortsFixedIpsResponse struct {
	IpAddress *string `json:"ip_address" tfsdk:"ip_address"`

	SubnetId *string `json:"subnet_id" tfsdk:"subnet_id"`
}

type OpenstackInstancePortsAllowedAddressPairsResponse struct {
	MacAddress *string `json:"mac_address,omitempty" tfsdk:"mac_address"`
}

type OpenstackInstancePortsSecurityGroupsResponse struct {
	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Customer *string `json:"customer,omitempty" tfsdk:"customer"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	ErrorMessage *string `json:"error_message,omitempty" tfsdk:"error_message"`

	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Project *string `json:"project,omitempty" tfsdk:"project"`

	ResourceType *string `json:"resource_type,omitempty" tfsdk:"resource_type"`

	Rules *[]common.OpenStackSecurityGroupRuleCreate `json:"rules,omitempty" tfsdk:"rules"`

	State *string `json:"state,omitempty" tfsdk:"state"`

	Tenant *string `json:"tenant,omitempty" tfsdk:"tenant"`

	TenantName *string `json:"tenant_name,omitempty" tfsdk:"tenant_name"`

	TenantUuid *string `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type OpenstackInstancePortsSecurityGroupsRulesResponse struct {
	Cidr *string `json:"cidr,omitempty" tfsdk:"cidr"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	Direction *string `json:"direction,omitempty" tfsdk:"direction"`

	Ethertype *string `json:"ethertype,omitempty" tfsdk:"ethertype"`

	FromPort *int64 `json:"from_port,omitempty" tfsdk:"from_port"`

	Id *int64 `json:"id,omitempty" tfsdk:"id"`

	Protocol *string `json:"protocol,omitempty" tfsdk:"protocol"`

	RemoteGroup *string `json:"remote_group,omitempty" tfsdk:"remote_group"`

	RemoteGroupName *string `json:"remote_group_name,omitempty" tfsdk:"remote_group_name"`

	RemoteGroupUuid *string `json:"remote_group_uuid,omitempty" tfsdk:"remote_group_uuid"`

	ToPort *int64 `json:"to_port,omitempty" tfsdk:"to_port"`
}

type OpenstackInstanceRancherClusterResponse struct {
	MarketplaceUuid *string `json:"marketplace_uuid,omitempty" tfsdk:"marketplace_uuid"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type OpenstackInstanceSecurityGroupsResponse struct {
	Url *string `json:"url" tfsdk:"url"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Rules *[]common.NestedSecurityGroupRule `json:"rules,omitempty" tfsdk:"rules"`

	State *string `json:"state,omitempty" tfsdk:"state"`
}

type OpenstackInstanceSecurityGroupsRulesResponse struct {
	Cidr *string `json:"cidr,omitempty" tfsdk:"cidr"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	Direction *string `json:"direction,omitempty" tfsdk:"direction"`

	Ethertype *string `json:"ethertype,omitempty" tfsdk:"ethertype"`

	FromPort *int64 `json:"from_port,omitempty" tfsdk:"from_port"`

	Id *int64 `json:"id,omitempty" tfsdk:"id"`

	Protocol *string `json:"protocol,omitempty" tfsdk:"protocol"`

	RemoteGroupName *string `json:"remote_group_name,omitempty" tfsdk:"remote_group_name"`

	RemoteGroupUuid *string `json:"remote_group_uuid,omitempty" tfsdk:"remote_group_uuid"`

	ToPort *int64 `json:"to_port,omitempty" tfsdk:"to_port"`
}

type OpenstackInstanceServerGroupResponse struct {
	Url *string `json:"url" tfsdk:"url"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Policy *string `json:"policy,omitempty" tfsdk:"policy"`

	State *string `json:"state,omitempty" tfsdk:"state"`
}

type OpenstackInstanceVolumesResponse struct {
	Bootable *bool `json:"bootable,omitempty" tfsdk:"bootable"`

	Device *string `json:"device,omitempty" tfsdk:"device"`

	ImageName *string `json:"image_name,omitempty" tfsdk:"image_name"`

	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	ResourceType *string `json:"resource_type,omitempty" tfsdk:"resource_type"`

	Size *int64 `json:"size,omitempty" tfsdk:"size"`

	State *string `json:"state,omitempty" tfsdk:"state"`

	Type *string `json:"type,omitempty" tfsdk:"type"`

	TypeName *string `json:"type_name,omitempty" tfsdk:"type_name"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
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
