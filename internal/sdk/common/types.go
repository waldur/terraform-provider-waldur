package common

import (
	"os"
	"time"
)

var (
	DefaultCreateTimeout  = 15 * time.Minute
	DefaultUpdateTimeout  = 15 * time.Minute
	DefaultDeleteTimeout  = 15 * time.Minute
	DefaultActionTimeout  = 15 * time.Minute
	DefaultPollDelay      = 10 * time.Second
	DefaultPollMinTimeout = 5 * time.Second
)

func init() {
	if delay := os.Getenv("WALDUR_POLL_DELAY"); delay != "" {
		if d, err := time.ParseDuration(delay); err == nil {
			DefaultPollDelay = d
		}
	}
	if timeout := os.Getenv("WALDUR_POLL_MIN_TIMEOUT"); timeout != "" {
		if t, err := time.ParseDuration(timeout); err == nil {
			DefaultPollMinTimeout = t
		}
	}
}

type BasePublicPlan struct {
	Archived           *bool                  `json:"archived,omitempty" tfsdk:"archived"`
	ArticleCode        *string                `json:"article_code,omitempty" tfsdk:"article_code"`
	BackendId          *string                `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Components         *[]NestedPlanComponent `json:"components,omitempty" tfsdk:"components"`
	Description        *string                `json:"description,omitempty" tfsdk:"description"`
	InitPrice          *float64               `json:"init_price,omitempty" tfsdk:"init_price"`
	IsActive           *bool                  `json:"is_active,omitempty" tfsdk:"is_active"`
	MaxAmount          *int64                 `json:"max_amount,omitempty" tfsdk:"max_amount"`
	MinimalPrice       *float64               `json:"minimal_price,omitempty" tfsdk:"minimal_price"`
	Name               *string                `json:"name,omitempty" tfsdk:"name"`
	OrganizationGroups *[]OrganizationGroup   `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	PlanType           *string                `json:"plan_type,omitempty" tfsdk:"plan_type"`
	ResourcesCount     *int64                 `json:"resources_count,omitempty" tfsdk:"resources_count"`
	SwitchPrice        *float64               `json:"switch_price,omitempty" tfsdk:"switch_price"`
	Unit               *string                `json:"unit,omitempty" tfsdk:"unit"`
	UnitPrice          *string                `json:"unit_price,omitempty" tfsdk:"unit_price"`
	Url                *string                `json:"url,omitempty" tfsdk:"url"`
	Uuid               *string                `json:"uuid,omitempty" tfsdk:"uuid"`
}

type NestedCampaign struct {
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Discount        *int64  `json:"discount,omitempty" tfsdk:"discount"`
	DiscountType    *string `json:"discount_type,omitempty" tfsdk:"discount_type"`
	EndDate         *string `json:"end_date,omitempty" tfsdk:"end_date"`
	Months          *int64  `json:"months,omitempty" tfsdk:"months"`
	Name            *string `json:"name,omitempty" tfsdk:"name"`
	ServiceProvider *string `json:"service_provider,omitempty" tfsdk:"service_provider"`
	StartDate       *string `json:"start_date,omitempty" tfsdk:"start_date"`
	Stock           *int64  `json:"stock,omitempty" tfsdk:"stock"`
	Uuid            *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type NestedEndpoint struct {
	Name *string `json:"name,omitempty" tfsdk:"name"`
	Url  *string `json:"url,omitempty" tfsdk:"url"`
	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type NestedOfferingFile struct {
	File *string `json:"file,omitempty" tfsdk:"file"`
	Name *string `json:"name,omitempty" tfsdk:"name"`
}

type NestedPartition struct {
	CpuBind          *int64  `json:"cpu_bind,omitempty" tfsdk:"cpu_bind"`
	DefCpuPerGpu     *int64  `json:"def_cpu_per_gpu,omitempty" tfsdk:"def_cpu_per_gpu"`
	DefMemPerCpu     *int64  `json:"def_mem_per_cpu,omitempty" tfsdk:"def_mem_per_cpu"`
	DefMemPerGpu     *int64  `json:"def_mem_per_gpu,omitempty" tfsdk:"def_mem_per_gpu"`
	DefMemPerNode    *int64  `json:"def_mem_per_node,omitempty" tfsdk:"def_mem_per_node"`
	DefaultTime      *int64  `json:"default_time,omitempty" tfsdk:"default_time"`
	ExclusiveTopo    *bool   `json:"exclusive_topo,omitempty" tfsdk:"exclusive_topo"`
	ExclusiveUser    *bool   `json:"exclusive_user,omitempty" tfsdk:"exclusive_user"`
	GraceTime        *int64  `json:"grace_time,omitempty" tfsdk:"grace_time"`
	MaxCpusPerNode   *int64  `json:"max_cpus_per_node,omitempty" tfsdk:"max_cpus_per_node"`
	MaxCpusPerSocket *int64  `json:"max_cpus_per_socket,omitempty" tfsdk:"max_cpus_per_socket"`
	MaxMemPerCpu     *int64  `json:"max_mem_per_cpu,omitempty" tfsdk:"max_mem_per_cpu"`
	MaxMemPerNode    *int64  `json:"max_mem_per_node,omitempty" tfsdk:"max_mem_per_node"`
	MaxNodes         *int64  `json:"max_nodes,omitempty" tfsdk:"max_nodes"`
	MaxTime          *int64  `json:"max_time,omitempty" tfsdk:"max_time"`
	MinNodes         *int64  `json:"min_nodes,omitempty" tfsdk:"min_nodes"`
	PartitionName    *string `json:"partition_name,omitempty" tfsdk:"partition_name"`
	PriorityTier     *int64  `json:"priority_tier,omitempty" tfsdk:"priority_tier"`
	Qos              *string `json:"qos,omitempty" tfsdk:"qos"`
	ReqResv          *bool   `json:"req_resv,omitempty" tfsdk:"req_resv"`
	Uuid             *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type NestedPlanComponent struct {
	Amount            *int64  `json:"amount,omitempty" tfsdk:"amount"`
	DiscountRate      *int64  `json:"discount_rate,omitempty" tfsdk:"discount_rate"`
	DiscountThreshold *int64  `json:"discount_threshold,omitempty" tfsdk:"discount_threshold"`
	FuturePrice       *string `json:"future_price,omitempty" tfsdk:"future_price"`
	MeasuredUnit      *string `json:"measured_unit,omitempty" tfsdk:"measured_unit"`
	Name              *string `json:"name,omitempty" tfsdk:"name"`
	Price             *string `json:"price,omitempty" tfsdk:"price"`
	Type              *string `json:"type,omitempty" tfsdk:"type"`
}

type NestedRole struct {
	Name *string `json:"name,omitempty" tfsdk:"name"`
	Url  *string `json:"url,omitempty" tfsdk:"url"`
	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type NestedScreenshot struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image,omitempty" tfsdk:"image"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
	Thumbnail   *string `json:"thumbnail,omitempty" tfsdk:"thumbnail"`
	Uuid        *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type NestedSecurityGroupRule struct {
	Cidr            *string `json:"cidr,omitempty" tfsdk:"cidr"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Direction       *string `json:"direction,omitempty" tfsdk:"direction"`
	Ethertype       *string `json:"ethertype,omitempty" tfsdk:"ethertype"`
	FromPort        *int64  `json:"from_port,omitempty" tfsdk:"from_port"`
	Id              *int64  `json:"id,omitempty" tfsdk:"id"`
	Protocol        *string `json:"protocol,omitempty" tfsdk:"protocol"`
	RemoteGroupName *string `json:"remote_group_name,omitempty" tfsdk:"remote_group_name"`
	RemoteGroupUuid *string `json:"remote_group_uuid,omitempty" tfsdk:"remote_group_uuid"`
	ToPort          *int64  `json:"to_port,omitempty" tfsdk:"to_port"`
}

type NestedSoftwareCatalog struct {
	Catalog      *NestedSoftwareCatalogCatalog   `json:"catalog,omitempty" tfsdk:"catalog"`
	PackageCount *int64                          `json:"package_count,omitempty" tfsdk:"package_count"`
	Partition    *NestedSoftwareCatalogPartition `json:"partition,omitempty" tfsdk:"partition"`
	Uuid         *string                         `json:"uuid,omitempty" tfsdk:"uuid"`
}
type NestedSoftwareCatalogCatalog struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
	Uuid        *string `json:"uuid,omitempty" tfsdk:"uuid"`
	Version     *string `json:"version,omitempty" tfsdk:"version"`
}
type NestedSoftwareCatalogPartition struct {
	PartitionName *string `json:"partition_name,omitempty" tfsdk:"partition_name"`
	PriorityTier  *int64  `json:"priority_tier,omitempty" tfsdk:"priority_tier"`
	Qos           *string `json:"qos,omitempty" tfsdk:"qos"`
	Uuid          *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type NestedTag struct {
	Name *string `json:"name,omitempty" tfsdk:"name"`
	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type NetworkRBACPolicy struct {
	BackendId        *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Network          *string `json:"network,omitempty" tfsdk:"network"`
	NetworkName      *string `json:"network_name,omitempty" tfsdk:"network_name"`
	PolicyType       *string `json:"policy_type,omitempty" tfsdk:"policy_type"`
	TargetTenant     *string `json:"target_tenant,omitempty" tfsdk:"target_tenant"`
	TargetTenantName *string `json:"target_tenant_name,omitempty" tfsdk:"target_tenant_name"`
	Url              *string `json:"url,omitempty" tfsdk:"url"`
	Uuid             *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type OfferingComponent struct {
	ArticleCode        *string `json:"article_code,omitempty" tfsdk:"article_code"`
	BillingType        *string `json:"billing_type,omitempty" tfsdk:"billing_type"`
	DefaultLimit       *int64  `json:"default_limit,omitempty" tfsdk:"default_limit"`
	Description        *string `json:"description,omitempty" tfsdk:"description"`
	Factor             *int64  `json:"factor,omitempty" tfsdk:"factor"`
	IsBoolean          *bool   `json:"is_boolean,omitempty" tfsdk:"is_boolean"`
	IsBuiltin          *bool   `json:"is_builtin,omitempty" tfsdk:"is_builtin"`
	IsPrepaid          *bool   `json:"is_prepaid,omitempty" tfsdk:"is_prepaid"`
	LimitAmount        *int64  `json:"limit_amount,omitempty" tfsdk:"limit_amount"`
	LimitPeriod        *string `json:"limit_period,omitempty" tfsdk:"limit_period"`
	MaxAvailableLimit  *int64  `json:"max_available_limit,omitempty" tfsdk:"max_available_limit"`
	MaxPrepaidDuration *int64  `json:"max_prepaid_duration,omitempty" tfsdk:"max_prepaid_duration"`
	MaxValue           *int64  `json:"max_value,omitempty" tfsdk:"max_value"`
	MeasuredUnit       *string `json:"measured_unit,omitempty" tfsdk:"measured_unit"`
	MinPrepaidDuration *int64  `json:"min_prepaid_duration,omitempty" tfsdk:"min_prepaid_duration"`
	MinValue           *int64  `json:"min_value,omitempty" tfsdk:"min_value"`
	Name               *string `json:"name,omitempty" tfsdk:"name"`
	OverageComponent   *string `json:"overage_component,omitempty" tfsdk:"overage_component"`
	Type               *string `json:"type,omitempty" tfsdk:"type"`
	UnitFactor         *int64  `json:"unit_factor,omitempty" tfsdk:"unit_factor"`
	Uuid               *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type OpenStackAllowedAddressPair struct {
	MacAddress *string `json:"mac_address,omitempty" tfsdk:"mac_address"`
}

type OpenStackAllowedAddressPairRequest struct {
	IpAddress  *string `json:"ip_address,omitempty" tfsdk:"ip_address"`
	MacAddress *string `json:"mac_address,omitempty" tfsdk:"mac_address"`
}

type OpenStackCreateFloatingIPRequest struct {
	Address           *string             `json:"address,omitempty" tfsdk:"address"`
	IpAddress         *string             `json:"ip_address,omitempty" tfsdk:"ip_address"`
	PortFixedIps      *[]OpenStackFixedIp `json:"port_fixed_ips,omitempty" tfsdk:"port_fixed_ips"`
	PortMacAddress    *string             `json:"port_mac_address,omitempty" tfsdk:"port_mac_address"`
	Subnet            *string             `json:"subnet" tfsdk:"subnet"`
	SubnetCidr        *string             `json:"subnet_cidr,omitempty" tfsdk:"subnet_cidr"`
	SubnetDescription *string             `json:"subnet_description,omitempty" tfsdk:"subnet_description"`
	SubnetName        *string             `json:"subnet_name,omitempty" tfsdk:"subnet_name"`
	SubnetUuid        *string             `json:"subnet_uuid,omitempty" tfsdk:"subnet_uuid"`
	Url               *string             `json:"url,omitempty" tfsdk:"url"`
	Uuid              *string             `json:"uuid,omitempty" tfsdk:"uuid"`
}

type OpenStackCreateInstancePortRequest struct {
	AllowedAddressPairs *[]OpenStackAllowedAddressPair `json:"allowed_address_pairs,omitempty" tfsdk:"allowed_address_pairs"`
	DeviceId            *string                        `json:"device_id,omitempty" tfsdk:"device_id"`
	DeviceOwner         *string                        `json:"device_owner,omitempty" tfsdk:"device_owner"`
	FixedIps            *[]OpenStackFixedIpRequest     `json:"fixed_ips,omitempty" tfsdk:"fixed_ips"`
	MacAddress          *string                        `json:"mac_address,omitempty" tfsdk:"mac_address"`
	Port                *string                        `json:"port,omitempty" tfsdk:"port"`
	SecurityGroups      *[]OpenStackSecurityGroup      `json:"security_groups,omitempty" tfsdk:"security_groups"`
	Subnet              *string                        `json:"subnet,omitempty" tfsdk:"subnet"`
	SubnetCidr          *string                        `json:"subnet_cidr,omitempty" tfsdk:"subnet_cidr"`
	SubnetDescription   *string                        `json:"subnet_description,omitempty" tfsdk:"subnet_description"`
	SubnetName          *string                        `json:"subnet_name,omitempty" tfsdk:"subnet_name"`
	SubnetUuid          *string                        `json:"subnet_uuid,omitempty" tfsdk:"subnet_uuid"`
	Url                 *string                        `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackDataVolumeRequest struct {
	Size       *int64  `json:"size" tfsdk:"size"`
	VolumeType *string `json:"volume_type,omitempty" tfsdk:"volume_type"`
}

type OpenStackFixedIp struct {
	IpAddress *string `json:"ip_address,omitempty" tfsdk:"ip_address"`
	SubnetId  *string `json:"subnet_id,omitempty" tfsdk:"subnet_id"`
}

type OpenStackFixedIpRequest struct {
	IpAddress *string `json:"ip_address" tfsdk:"ip_address"`
	SubnetId  *string `json:"subnet_id" tfsdk:"subnet_id"`
}

type OpenStackNestedFloatingIP struct {
	Address           *string             `json:"address,omitempty" tfsdk:"address"`
	PortFixedIps      *[]OpenStackFixedIp `json:"port_fixed_ips,omitempty" tfsdk:"port_fixed_ips"`
	PortMacAddress    *string             `json:"port_mac_address,omitempty" tfsdk:"port_mac_address"`
	Subnet            *string             `json:"subnet,omitempty" tfsdk:"subnet"`
	SubnetCidr        *string             `json:"subnet_cidr,omitempty" tfsdk:"subnet_cidr"`
	SubnetDescription *string             `json:"subnet_description,omitempty" tfsdk:"subnet_description"`
	SubnetName        *string             `json:"subnet_name,omitempty" tfsdk:"subnet_name"`
	SubnetUuid        *string             `json:"subnet_uuid,omitempty" tfsdk:"subnet_uuid"`
	Url               *string             `json:"url,omitempty" tfsdk:"url"`
	Uuid              *string             `json:"uuid,omitempty" tfsdk:"uuid"`
}

type OpenStackNestedInstance struct {
	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Name      *string `json:"name,omitempty" tfsdk:"name"`
	Uuid      *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type OpenStackNestedPort struct {
	AllowedAddressPairs *[]OpenStackAllowedAddressPair `json:"allowed_address_pairs,omitempty" tfsdk:"allowed_address_pairs"`
	DeviceId            *string                        `json:"device_id,omitempty" tfsdk:"device_id"`
	DeviceOwner         *string                        `json:"device_owner,omitempty" tfsdk:"device_owner"`
	FixedIps            *[]OpenStackFixedIp            `json:"fixed_ips,omitempty" tfsdk:"fixed_ips"`
	MacAddress          *string                        `json:"mac_address,omitempty" tfsdk:"mac_address"`
	SecurityGroups      *[]OpenStackSecurityGroup      `json:"security_groups,omitempty" tfsdk:"security_groups"`
	Subnet              *string                        `json:"subnet,omitempty" tfsdk:"subnet"`
	SubnetCidr          *string                        `json:"subnet_cidr,omitempty" tfsdk:"subnet_cidr"`
	SubnetDescription   *string                        `json:"subnet_description,omitempty" tfsdk:"subnet_description"`
	SubnetName          *string                        `json:"subnet_name,omitempty" tfsdk:"subnet_name"`
	SubnetUuid          *string                        `json:"subnet_uuid,omitempty" tfsdk:"subnet_uuid"`
	Url                 *string                        `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackNestedSecurityGroup struct {
	Description *string                    `json:"description,omitempty" tfsdk:"description"`
	Name        *string                    `json:"name,omitempty" tfsdk:"name"`
	Rules       *[]NestedSecurityGroupRule `json:"rules,omitempty" tfsdk:"rules"`
	State       *string                    `json:"state,omitempty" tfsdk:"state"`
	Url         *string                    `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackNestedServerGroup struct {
	Name   *string `json:"name,omitempty" tfsdk:"name"`
	Policy *string `json:"policy,omitempty" tfsdk:"policy"`
	State  *string `json:"state,omitempty" tfsdk:"state"`
	Url    *string `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackNestedSubNet struct {
	AllocationPools *[]OpenStackSubNetAllocationPool `json:"allocation_pools,omitempty" tfsdk:"allocation_pools"`
	Cidr            *string                          `json:"cidr,omitempty" tfsdk:"cidr"`
	Description     *string                          `json:"description,omitempty" tfsdk:"description"`
	EnableDhcp      *bool                            `json:"enable_dhcp,omitempty" tfsdk:"enable_dhcp"`
	GatewayIp       *string                          `json:"gateway_ip,omitempty" tfsdk:"gateway_ip"`
	IpVersion       *int64                           `json:"ip_version,omitempty" tfsdk:"ip_version"`
	Name            *string                          `json:"name,omitempty" tfsdk:"name"`
	Uuid            *string                          `json:"uuid,omitempty" tfsdk:"uuid"`
}

type OpenStackNestedVolume struct {
	Bootable                *bool   `json:"bootable,omitempty" tfsdk:"bootable"`
	Device                  *string `json:"device,omitempty" tfsdk:"device"`
	ImageName               *string `json:"image_name,omitempty" tfsdk:"image_name"`
	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Name                    *string `json:"name,omitempty" tfsdk:"name"`
	ResourceType            *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	Size                    *int64  `json:"size,omitempty" tfsdk:"size"`
	State                   *string `json:"state,omitempty" tfsdk:"state"`
	Type                    *string `json:"type,omitempty" tfsdk:"type"`
	TypeName                *string `json:"type_name,omitempty" tfsdk:"type_name"`
	Url                     *string `json:"url,omitempty" tfsdk:"url"`
	Uuid                    *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type OpenStackPortNestedSecurityGroup struct {
	Name *string `json:"name,omitempty" tfsdk:"name"`
	Url  *string `json:"url,omitempty" tfsdk:"url"`
	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type OpenStackPortNestedSecurityGroupRequest struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url,omitempty" tfsdk:"url"`
	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type OpenStackSecurityGroup struct {
	BackendId               *string                             `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Customer                *string                             `json:"customer,omitempty" tfsdk:"customer"`
	Description             *string                             `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage            *string                             `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback          *string                             `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	MarketplaceResourceUuid *string                             `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Name                    *string                             `json:"name,omitempty" tfsdk:"name"`
	Project                 *string                             `json:"project,omitempty" tfsdk:"project"`
	ResourceType            *string                             `json:"resource_type,omitempty" tfsdk:"resource_type"`
	Rules                   *[]OpenStackSecurityGroupRuleCreate `json:"rules,omitempty" tfsdk:"rules"`
	State                   *string                             `json:"state,omitempty" tfsdk:"state"`
	Tenant                  *string                             `json:"tenant,omitempty" tfsdk:"tenant"`
	TenantName              *string                             `json:"tenant_name,omitempty" tfsdk:"tenant_name"`
	TenantUuid              *string                             `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`
	Url                     *string                             `json:"url,omitempty" tfsdk:"url"`
	Uuid                    *string                             `json:"uuid,omitempty" tfsdk:"uuid"`
}

type OpenStackSecurityGroupHyperlinkRequest struct {
	Description *string                    `json:"description,omitempty" tfsdk:"description"`
	Name        *string                    `json:"name,omitempty" tfsdk:"name"`
	Rules       *[]NestedSecurityGroupRule `json:"rules,omitempty" tfsdk:"rules"`
	State       *string                    `json:"state,omitempty" tfsdk:"state"`
	Url         *string                    `json:"url" tfsdk:"url"`
}

type OpenStackSecurityGroupRuleCreate struct {
	Cidr            *string `json:"cidr,omitempty" tfsdk:"cidr"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Direction       *string `json:"direction,omitempty" tfsdk:"direction"`
	Ethertype       *string `json:"ethertype,omitempty" tfsdk:"ethertype"`
	FromPort        *int64  `json:"from_port,omitempty" tfsdk:"from_port"`
	Id              *int64  `json:"id,omitempty" tfsdk:"id"`
	Protocol        *string `json:"protocol,omitempty" tfsdk:"protocol"`
	RemoteGroup     *string `json:"remote_group,omitempty" tfsdk:"remote_group"`
	RemoteGroupName *string `json:"remote_group_name,omitempty" tfsdk:"remote_group_name"`
	RemoteGroupUuid *string `json:"remote_group_uuid,omitempty" tfsdk:"remote_group_uuid"`
	ToPort          *int64  `json:"to_port,omitempty" tfsdk:"to_port"`
}

type OpenStackSecurityGroupRuleCreateRequest struct {
	Cidr            *string `json:"cidr,omitempty" tfsdk:"cidr"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Direction       *string `json:"direction,omitempty" tfsdk:"direction"`
	Ethertype       *string `json:"ethertype,omitempty" tfsdk:"ethertype"`
	FromPort        *int64  `json:"from_port,omitempty" tfsdk:"from_port"`
	Id              *int64  `json:"id,omitempty" tfsdk:"id"`
	Protocol        *string `json:"protocol,omitempty" tfsdk:"protocol"`
	RemoteGroup     *string `json:"remote_group,omitempty" tfsdk:"remote_group"`
	RemoteGroupName *string `json:"remote_group_name,omitempty" tfsdk:"remote_group_name"`
	RemoteGroupUuid *string `json:"remote_group_uuid,omitempty" tfsdk:"remote_group_uuid"`
	ToPort          *int64  `json:"to_port,omitempty" tfsdk:"to_port"`
}

type OpenStackStaticRoute struct {
	Destination *string `json:"destination,omitempty" tfsdk:"destination"`
	Nexthop     *string `json:"nexthop,omitempty" tfsdk:"nexthop"`
}

type OpenStackStaticRouteRequest struct {
	Destination *string `json:"destination" tfsdk:"destination"`
	Nexthop     *string `json:"nexthop" tfsdk:"nexthop"`
}

type OpenStackSubNetAllocationPool struct {
	End   *string `json:"end,omitempty" tfsdk:"end"`
	Start *string `json:"start,omitempty" tfsdk:"start"`
}

type OpenStackSubNetAllocationPoolRequest struct {
	End   *string `json:"end" tfsdk:"end"`
	Start *string `json:"start" tfsdk:"start"`
}

type OpenStackTenantSecurityGroupRequest struct {
	Description *string                                    `json:"description,omitempty" tfsdk:"description"`
	Name        *string                                    `json:"name" tfsdk:"name"`
	Rules       *[]OpenStackSecurityGroupRuleCreateRequest `json:"rules,omitempty" tfsdk:"rules"`
}

type OrderDetails struct {
	ActivationPrice            *float64           `json:"activation_price,omitempty" tfsdk:"activation_price"`
	Attachment                 *string            `json:"attachment,omitempty" tfsdk:"attachment"`
	BackendId                  *string            `json:"backend_id,omitempty" tfsdk:"backend_id"`
	CallbackUrl                *string            `json:"callback_url,omitempty" tfsdk:"callback_url"`
	CanTerminate               *bool              `json:"can_terminate,omitempty" tfsdk:"can_terminate"`
	CategoryIcon               *string            `json:"category_icon,omitempty" tfsdk:"category_icon"`
	CategoryTitle              *string            `json:"category_title,omitempty" tfsdk:"category_title"`
	CategoryUuid               *string            `json:"category_uuid,omitempty" tfsdk:"category_uuid"`
	CompletedAt                *string            `json:"completed_at,omitempty" tfsdk:"completed_at"`
	ConsumerReviewedAt         *string            `json:"consumer_reviewed_at,omitempty" tfsdk:"consumer_reviewed_at"`
	ConsumerReviewedBy         *string            `json:"consumer_reviewed_by,omitempty" tfsdk:"consumer_reviewed_by"`
	ConsumerReviewedByFullName *string            `json:"consumer_reviewed_by_full_name,omitempty" tfsdk:"consumer_reviewed_by_full_name"`
	ConsumerReviewedByUsername *string            `json:"consumer_reviewed_by_username,omitempty" tfsdk:"consumer_reviewed_by_username"`
	Cost                       *string            `json:"cost,omitempty" tfsdk:"cost"`
	CreatedByCivilNumber       *string            `json:"created_by_civil_number,omitempty" tfsdk:"created_by_civil_number"`
	CreatedByFullName          *string            `json:"created_by_full_name,omitempty" tfsdk:"created_by_full_name"`
	CreatedByUsername          *string            `json:"created_by_username,omitempty" tfsdk:"created_by_username"`
	CustomerSlug               *string            `json:"customer_slug,omitempty" tfsdk:"customer_slug"`
	ErrorMessage               *string            `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback             *string            `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	FixedPrice                 *float64           `json:"fixed_price,omitempty" tfsdk:"fixed_price"`
	Issue                      *OrderDetailsIssue `json:"issue,omitempty" tfsdk:"issue"`
	MarketplaceResourceUuid    *string            `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	NewCostEstimate            *string            `json:"new_cost_estimate,omitempty" tfsdk:"new_cost_estimate"`
	NewPlanName                *string            `json:"new_plan_name,omitempty" tfsdk:"new_plan_name"`
	NewPlanUuid                *string            `json:"new_plan_uuid,omitempty" tfsdk:"new_plan_uuid"`
	Offering                   *string            `json:"offering,omitempty" tfsdk:"offering"`
	OfferingBillable           *bool              `json:"offering_billable,omitempty" tfsdk:"offering_billable"`
	OfferingDescription        *string            `json:"offering_description,omitempty" tfsdk:"offering_description"`
	OfferingImage              *string            `json:"offering_image,omitempty" tfsdk:"offering_image"`
	OfferingName               *string            `json:"offering_name,omitempty" tfsdk:"offering_name"`
	OfferingShared             *bool              `json:"offering_shared,omitempty" tfsdk:"offering_shared"`
	OfferingThumbnail          *string            `json:"offering_thumbnail,omitempty" tfsdk:"offering_thumbnail"`
	OfferingType               *string            `json:"offering_type,omitempty" tfsdk:"offering_type"`
	OfferingUuid               *string            `json:"offering_uuid,omitempty" tfsdk:"offering_uuid"`
	OldCostEstimate            *float64           `json:"old_cost_estimate,omitempty" tfsdk:"old_cost_estimate"`
	OldPlanName                *string            `json:"old_plan_name,omitempty" tfsdk:"old_plan_name"`
	OldPlanUuid                *string            `json:"old_plan_uuid,omitempty" tfsdk:"old_plan_uuid"`
	OrderSubtype               *string            `json:"order_subtype,omitempty" tfsdk:"order_subtype"`
	Output                     *string            `json:"output,omitempty" tfsdk:"output"`
	Plan                       *string            `json:"plan,omitempty" tfsdk:"plan"`
	PlanDescription            *string            `json:"plan_description,omitempty" tfsdk:"plan_description"`
	PlanName                   *string            `json:"plan_name,omitempty" tfsdk:"plan_name"`
	PlanUnit                   *string            `json:"plan_unit,omitempty" tfsdk:"plan_unit"`
	PlanUuid                   *string            `json:"plan_uuid,omitempty" tfsdk:"plan_uuid"`
	ProjectDescription         *string            `json:"project_description,omitempty" tfsdk:"project_description"`
	ProjectSlug                *string            `json:"project_slug,omitempty" tfsdk:"project_slug"`
	ProviderName               *string            `json:"provider_name,omitempty" tfsdk:"provider_name"`
	ProviderReviewedAt         *string            `json:"provider_reviewed_at,omitempty" tfsdk:"provider_reviewed_at"`
	ProviderReviewedBy         *string            `json:"provider_reviewed_by,omitempty" tfsdk:"provider_reviewed_by"`
	ProviderReviewedByFullName *string            `json:"provider_reviewed_by_full_name,omitempty" tfsdk:"provider_reviewed_by_full_name"`
	ProviderReviewedByUsername *string            `json:"provider_reviewed_by_username,omitempty" tfsdk:"provider_reviewed_by_username"`
	ProviderSlug               *string            `json:"provider_slug,omitempty" tfsdk:"provider_slug"`
	ProviderUuid               *string            `json:"provider_uuid,omitempty" tfsdk:"provider_uuid"`
	RequestComment             *string            `json:"request_comment,omitempty" tfsdk:"request_comment"`
	ResourceName               *string            `json:"resource_name,omitempty" tfsdk:"resource_name"`
	ResourceType               *string            `json:"resource_type,omitempty" tfsdk:"resource_type"`
	ResourceUuid               *string            `json:"resource_uuid,omitempty" tfsdk:"resource_uuid"`
	Slug                       *string            `json:"slug,omitempty" tfsdk:"slug"`
	StartDate                  *string            `json:"start_date,omitempty" tfsdk:"start_date"`
	State                      *string            `json:"state,omitempty" tfsdk:"state"`
	TerminationComment         *string            `json:"termination_comment,omitempty" tfsdk:"termination_comment"`
	Type                       *string            `json:"type,omitempty" tfsdk:"type"`
	Url                        *string            `json:"url,omitempty" tfsdk:"url"`
	Uuid                       *string            `json:"uuid,omitempty" tfsdk:"uuid"`
}
type OrderDetailsIssue struct {
	Key  *string `json:"key,omitempty" tfsdk:"key"`
	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type OrganizationGroup struct {
	CustomersCount *int64  `json:"customers_count,omitempty" tfsdk:"customers_count"`
	Name           *string `json:"name,omitempty" tfsdk:"name"`
	Parent         *string `json:"parent,omitempty" tfsdk:"parent"`
	ParentName     *string `json:"parent_name,omitempty" tfsdk:"parent_name"`
	ParentUuid     *string `json:"parent_uuid,omitempty" tfsdk:"parent_uuid"`
	Url            *string `json:"url,omitempty" tfsdk:"url"`
	Uuid           *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type PaymentProfile struct {
	Attributes         *PaymentProfileAttributes `json:"attributes,omitempty" tfsdk:"attributes"`
	IsActive           *bool                     `json:"is_active,omitempty" tfsdk:"is_active"`
	Name               *string                   `json:"name,omitempty" tfsdk:"name"`
	Organization       *string                   `json:"organization,omitempty" tfsdk:"organization"`
	OrganizationUuid   *string                   `json:"organization_uuid,omitempty" tfsdk:"organization_uuid"`
	PaymentType        *string                   `json:"payment_type,omitempty" tfsdk:"payment_type"`
	PaymentTypeDisplay *string                   `json:"payment_type_display,omitempty" tfsdk:"payment_type_display"`
	Url                *string                   `json:"url,omitempty" tfsdk:"url"`
	Uuid               *string                   `json:"uuid,omitempty" tfsdk:"uuid"`
}

type PaymentProfileAttributes struct {
	AgreementNumber *string `json:"agreement_number,omitempty" tfsdk:"agreement_number"`
	ContractSum     *int64  `json:"contract_sum,omitempty" tfsdk:"contract_sum"`
	EndDate         *string `json:"end_date,omitempty" tfsdk:"end_date"`
}

type Quota struct {
	Limit *int64  `json:"limit,omitempty" tfsdk:"limit"`
	Name  *string `json:"name,omitempty" tfsdk:"name"`
	Usage *int64  `json:"usage,omitempty" tfsdk:"usage"`
}

type ReportSection struct {
	Body   *string `json:"body,omitempty" tfsdk:"body"`
	Header *string `json:"header,omitempty" tfsdk:"header"`
}
