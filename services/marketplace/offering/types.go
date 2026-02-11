package offering

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type MarketplaceOfferingCreateRequest struct {
}

type MarketplaceOfferingResponse struct {
	UUID *string `json:"uuid"`

	BackendId                 *string                                     `json:"backend_id" tfsdk:"backend_id"`
	Billable                  *bool                                       `json:"billable" tfsdk:"billable"`
	BillingTypeClassification *string                                     `json:"billing_type_classification" tfsdk:"billing_type_classification"`
	Category                  *string                                     `json:"category" tfsdk:"category"`
	CategoryTitle             *string                                     `json:"category_title" tfsdk:"category_title"`
	CategoryUuid              *string                                     `json:"category_uuid" tfsdk:"category_uuid"`
	CitationCount             *int64                                      `json:"citation_count" tfsdk:"citation_count"`
	ComplianceChecklist       *string                                     `json:"compliance_checklist" tfsdk:"compliance_checklist"`
	Components                *[]common.OfferingComponent                 `json:"components" tfsdk:"components"`
	Country                   *string                                     `json:"country" tfsdk:"country"`
	Customer                  *string                                     `json:"customer" tfsdk:"customer"`
	DataciteDoi               *string                                     `json:"datacite_doi" tfsdk:"datacite_doi"`
	Description               *string                                     `json:"description" tfsdk:"description"`
	Endpoints                 *[]common.NestedEndpoint                    `json:"endpoints" tfsdk:"endpoints"`
	Files                     *[]common.NestedOfferingFile                `json:"files" tfsdk:"files"`
	FullDescription           *string                                     `json:"full_description" tfsdk:"full_description"`
	GettingStarted            *string                                     `json:"getting_started" tfsdk:"getting_started"`
	GoogleCalendarIsPublic    *bool                                       `json:"google_calendar_is_public" tfsdk:"google_calendar_is_public"`
	GoogleCalendarLink        *string                                     `json:"google_calendar_link" tfsdk:"google_calendar_link"`
	HasComplianceRequirements *bool                                       `json:"has_compliance_requirements" tfsdk:"has_compliance_requirements"`
	Image                     *string                                     `json:"image" tfsdk:"image"`
	IntegrationGuide          *string                                     `json:"integration_guide" tfsdk:"integration_guide"`
	IsAccessible              *bool                                       `json:"is_accessible" tfsdk:"is_accessible"`
	Latitude                  common.FlexibleNumber                       `json:"latitude" tfsdk:"latitude"`
	Longitude                 common.FlexibleNumber                       `json:"longitude" tfsdk:"longitude"`
	Name                      *string                                     `json:"name" tfsdk:"name"`
	Options                   *MarketplaceOfferingOptionsResponse         `json:"options" tfsdk:"options"`
	OrderCount                *int64                                      `json:"order_count" tfsdk:"order_count"`
	OrganizationGroups        *[]common.OrganizationGroup                 `json:"organization_groups" tfsdk:"organization_groups"`
	ParentDescription         *string                                     `json:"parent_description" tfsdk:"parent_description"`
	ParentName                *string                                     `json:"parent_name" tfsdk:"parent_name"`
	ParentUuid                *string                                     `json:"parent_uuid" tfsdk:"parent_uuid"`
	Partitions                *[]common.NestedPartition                   `json:"partitions" tfsdk:"partitions"`
	PausedReason              *string                                     `json:"paused_reason" tfsdk:"paused_reason"`
	Plans                     *[]common.BasePublicPlan                    `json:"plans" tfsdk:"plans"`
	PluginOptions             *MarketplaceOfferingPluginOptionsResponse   `json:"plugin_options" tfsdk:"plugin_options"`
	PrivacyPolicyLink         *string                                     `json:"privacy_policy_link" tfsdk:"privacy_policy_link"`
	Project                   *string                                     `json:"project" tfsdk:"project"`
	PromotionCampaigns        *[]common.NestedCampaign                    `json:"promotion_campaigns" tfsdk:"promotion_campaigns"`
	Quotas                    *[]common.Quota                             `json:"quotas" tfsdk:"quotas"`
	ResourceOptions           *MarketplaceOfferingResourceOptionsResponse `json:"resource_options" tfsdk:"resource_options"`
	Roles                     *[]common.NestedRole                        `json:"roles" tfsdk:"roles"`
	Scope                     *string                                     `json:"scope" tfsdk:"scope"`
	ScopeErrorMessage         *string                                     `json:"scope_error_message" tfsdk:"scope_error_message"`
	ScopeName                 *string                                     `json:"scope_name" tfsdk:"scope_name"`
	ScopeState                *string                                     `json:"scope_state" tfsdk:"scope_state"`
	ScopeUuid                 *string                                     `json:"scope_uuid" tfsdk:"scope_uuid"`
	Screenshots               *[]common.NestedScreenshot                  `json:"screenshots" tfsdk:"screenshots"`
	Shared                    *bool                                       `json:"shared" tfsdk:"shared"`
	Slug                      *string                                     `json:"slug" tfsdk:"slug"`
	SoftwareCatalogs          *[]common.NestedSoftwareCatalog             `json:"software_catalogs" tfsdk:"software_catalogs"`
	State                     *string                                     `json:"state" tfsdk:"state"`
	Tags                      *[]common.NestedTag                         `json:"tags" tfsdk:"tags"`
	Thumbnail                 *string                                     `json:"thumbnail" tfsdk:"thumbnail"`
	TotalCost                 *int64                                      `json:"total_cost" tfsdk:"total_cost"`
	TotalCostEstimated        *int64                                      `json:"total_cost_estimated" tfsdk:"total_cost_estimated"`
	TotalCustomers            *int64                                      `json:"total_customers" tfsdk:"total_customers"`
	Type                      *string                                     `json:"type" tfsdk:"type"`
	Url                       *string                                     `json:"url" tfsdk:"url"`
	UserHasConsent            *bool                                       `json:"user_has_consent" tfsdk:"user_has_consent"`
	VendorDetails             *string                                     `json:"vendor_details" tfsdk:"vendor_details"`
}

type MarketplaceOfferingComponentsResponse struct {
	ArticleCode        *string `json:"article_code" tfsdk:"article_code"`
	BillingType        *string `json:"billing_type" tfsdk:"billing_type"`
	DefaultLimit       *int64  `json:"default_limit" tfsdk:"default_limit"`
	Description        *string `json:"description" tfsdk:"description"`
	Factor             *int64  `json:"factor" tfsdk:"factor"`
	IsBoolean          *bool   `json:"is_boolean" tfsdk:"is_boolean"`
	IsBuiltin          *bool   `json:"is_builtin" tfsdk:"is_builtin"`
	IsPrepaid          *bool   `json:"is_prepaid" tfsdk:"is_prepaid"`
	LimitAmount        *int64  `json:"limit_amount" tfsdk:"limit_amount"`
	LimitPeriod        *string `json:"limit_period" tfsdk:"limit_period"`
	MaxAvailableLimit  *int64  `json:"max_available_limit" tfsdk:"max_available_limit"`
	MaxPrepaidDuration *int64  `json:"max_prepaid_duration" tfsdk:"max_prepaid_duration"`
	MaxValue           *int64  `json:"max_value" tfsdk:"max_value"`
	MeasuredUnit       *string `json:"measured_unit" tfsdk:"measured_unit"`
	MinPrepaidDuration *int64  `json:"min_prepaid_duration" tfsdk:"min_prepaid_duration"`
	MinValue           *int64  `json:"min_value" tfsdk:"min_value"`
	Name               *string `json:"name" tfsdk:"name"`
	OverageComponent   *string `json:"overage_component" tfsdk:"overage_component"`
	Type               *string `json:"type" tfsdk:"type"`
	UnitFactor         *int64  `json:"unit_factor" tfsdk:"unit_factor"`
	Uuid               *string `json:"uuid" tfsdk:"uuid"`
}

type MarketplaceOfferingEndpointsResponse struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
	Uuid *string `json:"uuid" tfsdk:"uuid"`
}

type MarketplaceOfferingFilesResponse struct {
	File *string `json:"file" tfsdk:"file"`
	Name *string `json:"name" tfsdk:"name"`
}

type MarketplaceOfferingOptionsResponse struct {
	Order *[]string `json:"order" tfsdk:"order"`
}

type MarketplaceOfferingOrganizationGroupsResponse struct {
	CustomersCount *int64  `json:"customers_count" tfsdk:"customers_count"`
	Name           *string `json:"name" tfsdk:"name"`
	Parent         *string `json:"parent" tfsdk:"parent"`
	ParentName     *string `json:"parent_name" tfsdk:"parent_name"`
	ParentUuid     *string `json:"parent_uuid" tfsdk:"parent_uuid"`
	Url            *string `json:"url" tfsdk:"url"`
	Uuid           *string `json:"uuid" tfsdk:"uuid"`
}

type MarketplaceOfferingPartitionsResponse struct {
	CpuBind          *int64  `json:"cpu_bind" tfsdk:"cpu_bind"`
	DefCpuPerGpu     *int64  `json:"def_cpu_per_gpu" tfsdk:"def_cpu_per_gpu"`
	DefMemPerCpu     *int64  `json:"def_mem_per_cpu" tfsdk:"def_mem_per_cpu"`
	DefMemPerGpu     *int64  `json:"def_mem_per_gpu" tfsdk:"def_mem_per_gpu"`
	DefMemPerNode    *int64  `json:"def_mem_per_node" tfsdk:"def_mem_per_node"`
	DefaultTime      *int64  `json:"default_time" tfsdk:"default_time"`
	ExclusiveTopo    *bool   `json:"exclusive_topo" tfsdk:"exclusive_topo"`
	ExclusiveUser    *bool   `json:"exclusive_user" tfsdk:"exclusive_user"`
	GraceTime        *int64  `json:"grace_time" tfsdk:"grace_time"`
	MaxCpusPerNode   *int64  `json:"max_cpus_per_node" tfsdk:"max_cpus_per_node"`
	MaxCpusPerSocket *int64  `json:"max_cpus_per_socket" tfsdk:"max_cpus_per_socket"`
	MaxMemPerCpu     *int64  `json:"max_mem_per_cpu" tfsdk:"max_mem_per_cpu"`
	MaxMemPerNode    *int64  `json:"max_mem_per_node" tfsdk:"max_mem_per_node"`
	MaxNodes         *int64  `json:"max_nodes" tfsdk:"max_nodes"`
	MaxTime          *int64  `json:"max_time" tfsdk:"max_time"`
	MinNodes         *int64  `json:"min_nodes" tfsdk:"min_nodes"`
	PartitionName    *string `json:"partition_name" tfsdk:"partition_name"`
	PriorityTier     *int64  `json:"priority_tier" tfsdk:"priority_tier"`
	Qos              *string `json:"qos" tfsdk:"qos"`
	ReqResv          *bool   `json:"req_resv" tfsdk:"req_resv"`
	Uuid             *string `json:"uuid" tfsdk:"uuid"`
}

type MarketplaceOfferingPlansResponse struct {
	Archived           *bool                         `json:"archived" tfsdk:"archived"`
	ArticleCode        *string                       `json:"article_code" tfsdk:"article_code"`
	BackendId          *string                       `json:"backend_id" tfsdk:"backend_id"`
	Components         *[]common.NestedPlanComponent `json:"components" tfsdk:"components"`
	Description        *string                       `json:"description" tfsdk:"description"`
	InitPrice          common.FlexibleNumber         `json:"init_price" tfsdk:"init_price"`
	IsActive           *bool                         `json:"is_active" tfsdk:"is_active"`
	MaxAmount          *int64                        `json:"max_amount" tfsdk:"max_amount"`
	MinimalPrice       common.FlexibleNumber         `json:"minimal_price" tfsdk:"minimal_price"`
	Name               *string                       `json:"name" tfsdk:"name"`
	OrganizationGroups *[]common.OrganizationGroup   `json:"organization_groups" tfsdk:"organization_groups"`
	PlanType           *string                       `json:"plan_type" tfsdk:"plan_type"`
	ResourcesCount     *int64                        `json:"resources_count" tfsdk:"resources_count"`
	SwitchPrice        common.FlexibleNumber         `json:"switch_price" tfsdk:"switch_price"`
	Unit               *string                       `json:"unit" tfsdk:"unit"`
	UnitPrice          *string                       `json:"unit_price" tfsdk:"unit_price"`
	Url                *string                       `json:"url" tfsdk:"url"`
	Uuid               *string                       `json:"uuid" tfsdk:"uuid"`
}

type MarketplaceOfferingPlansComponentsResponse struct {
	Amount            *int64  `json:"amount" tfsdk:"amount"`
	DiscountRate      *int64  `json:"discount_rate" tfsdk:"discount_rate"`
	DiscountThreshold *int64  `json:"discount_threshold" tfsdk:"discount_threshold"`
	FuturePrice       *string `json:"future_price" tfsdk:"future_price"`
	MeasuredUnit      *string `json:"measured_unit" tfsdk:"measured_unit"`
	Name              *string `json:"name" tfsdk:"name"`
	Price             *string `json:"price" tfsdk:"price"`
	Type              *string `json:"type" tfsdk:"type"`
}

type MarketplaceOfferingPlansOrganizationGroupsResponse struct {
	CustomersCount *int64  `json:"customers_count" tfsdk:"customers_count"`
	Name           *string `json:"name" tfsdk:"name"`
	Parent         *string `json:"parent" tfsdk:"parent"`
	ParentName     *string `json:"parent_name" tfsdk:"parent_name"`
	ParentUuid     *string `json:"parent_uuid" tfsdk:"parent_uuid"`
	Url            *string `json:"url" tfsdk:"url"`
	Uuid           *string `json:"uuid" tfsdk:"uuid"`
}

type MarketplaceOfferingPluginOptionsResponse struct {
	AutoApproveInServiceProviderProjects           *bool     `json:"auto_approve_in_service_provider_projects" tfsdk:"auto_approve_in_service_provider_projects"`
	AutoApproveMarketplaceScript                   *bool     `json:"auto_approve_marketplace_script" tfsdk:"auto_approve_marketplace_script"`
	AutoApproveRemoteOrders                        *bool     `json:"auto_approve_remote_orders" tfsdk:"auto_approve_remote_orders"`
	BackendIdDisplayLabel                          *string   `json:"backend_id_display_label" tfsdk:"backend_id_display_label"`
	CanRestoreResource                             *bool     `json:"can_restore_resource" tfsdk:"can_restore_resource"`
	ConcealBillingData                             *bool     `json:"conceal_billing_data" tfsdk:"conceal_billing_data"`
	CreateOrdersOnResourceOptionChange             *bool     `json:"create_orders_on_resource_option_change" tfsdk:"create_orders_on_resource_option_change"`
	DefaultInternalNetworkMtu                      *int64    `json:"default_internal_network_mtu" tfsdk:"default_internal_network_mtu"`
	DefaultResourceTerminationOffsetInDays         *int64    `json:"default_resource_termination_offset_in_days" tfsdk:"default_resource_termination_offset_in_days"`
	DeploymentMode                                 *string   `json:"deployment_mode" tfsdk:"deployment_mode"`
	DisableAutoapprove                             *bool     `json:"disable_autoapprove" tfsdk:"disable_autoapprove"`
	EnableDisplayOfOrderActionsForServiceProvider  *bool     `json:"enable_display_of_order_actions_for_service_provider" tfsdk:"enable_display_of_order_actions_for_service_provider"`
	EnableIssuesForMembershipChanges               *bool     `json:"enable_issues_for_membership_changes" tfsdk:"enable_issues_for_membership_changes"`
	EnablePurchaseOrderUpload                      *bool     `json:"enable_purchase_order_upload" tfsdk:"enable_purchase_order_upload"`
	FlavorsRegex                                   *string   `json:"flavors_regex" tfsdk:"flavors_regex"`
	HeappeClusterId                                *string   `json:"heappe_cluster_id" tfsdk:"heappe_cluster_id"`
	HeappeLocalBasePath                            *string   `json:"heappe_local_base_path" tfsdk:"heappe_local_base_path"`
	HeappeUrl                                      *string   `json:"heappe_url" tfsdk:"heappe_url"`
	HeappeUsername                                 *string   `json:"heappe_username" tfsdk:"heappe_username"`
	HighlightBackendIdDisplay                      *bool     `json:"highlight_backend_id_display" tfsdk:"highlight_backend_id_display"`
	HomedirPrefix                                  *string   `json:"homedir_prefix" tfsdk:"homedir_prefix"`
	InitialPrimarygroupNumber                      *int64    `json:"initial_primarygroup_number" tfsdk:"initial_primarygroup_number"`
	InitialUidnumber                               *int64    `json:"initial_uidnumber" tfsdk:"initial_uidnumber"`
	InitialUsergroupNumber                         *int64    `json:"initial_usergroup_number" tfsdk:"initial_usergroup_number"`
	IsResourceTerminationDateRequired              *bool     `json:"is_resource_termination_date_required" tfsdk:"is_resource_termination_date_required"`
	LatestDateForResourceTermination               *string   `json:"latest_date_for_resource_termination" tfsdk:"latest_date_for_resource_termination"`
	ManagedRancherLoadBalancerDataVolumeSizeGb     *int64    `json:"managed_rancher_load_balancer_data_volume_size_gb" tfsdk:"managed_rancher_load_balancer_data_volume_size_gb"`
	ManagedRancherLoadBalancerDataVolumeTypeName   *string   `json:"managed_rancher_load_balancer_data_volume_type_name" tfsdk:"managed_rancher_load_balancer_data_volume_type_name"`
	ManagedRancherLoadBalancerFlavorName           *string   `json:"managed_rancher_load_balancer_flavor_name" tfsdk:"managed_rancher_load_balancer_flavor_name"`
	ManagedRancherLoadBalancerSystemVolumeSizeGb   *int64    `json:"managed_rancher_load_balancer_system_volume_size_gb" tfsdk:"managed_rancher_load_balancer_system_volume_size_gb"`
	ManagedRancherLoadBalancerSystemVolumeTypeName *string   `json:"managed_rancher_load_balancer_system_volume_type_name" tfsdk:"managed_rancher_load_balancer_system_volume_type_name"`
	ManagedRancherServerDataVolumeSizeGb           *int64    `json:"managed_rancher_server_data_volume_size_gb" tfsdk:"managed_rancher_server_data_volume_size_gb"`
	ManagedRancherServerDataVolumeTypeName         *string   `json:"managed_rancher_server_data_volume_type_name" tfsdk:"managed_rancher_server_data_volume_type_name"`
	ManagedRancherServerFlavorName                 *string   `json:"managed_rancher_server_flavor_name" tfsdk:"managed_rancher_server_flavor_name"`
	ManagedRancherServerSystemVolumeSizeGb         *int64    `json:"managed_rancher_server_system_volume_size_gb" tfsdk:"managed_rancher_server_system_volume_size_gb"`
	ManagedRancherServerSystemVolumeTypeName       *string   `json:"managed_rancher_server_system_volume_type_name" tfsdk:"managed_rancher_server_system_volume_type_name"`
	ManagedRancherTenantMaxCpu                     *int64    `json:"managed_rancher_tenant_max_cpu" tfsdk:"managed_rancher_tenant_max_cpu"`
	ManagedRancherTenantMaxDisk                    *int64    `json:"managed_rancher_tenant_max_disk" tfsdk:"managed_rancher_tenant_max_disk"`
	ManagedRancherTenantMaxRam                     *int64    `json:"managed_rancher_tenant_max_ram" tfsdk:"managed_rancher_tenant_max_ram"`
	ManagedRancherWorkerSystemVolumeSizeGb         *int64    `json:"managed_rancher_worker_system_volume_size_gb" tfsdk:"managed_rancher_worker_system_volume_size_gb"`
	ManagedRancherWorkerSystemVolumeTypeName       *string   `json:"managed_rancher_worker_system_volume_type_name" tfsdk:"managed_rancher_worker_system_volume_type_name"`
	MaxInstances                                   *int64    `json:"max_instances" tfsdk:"max_instances"`
	MaxResourceTerminationOffsetInDays             *int64    `json:"max_resource_termination_offset_in_days" tfsdk:"max_resource_termination_offset_in_days"`
	MaxSecurityGroups                              *int64    `json:"max_security_groups" tfsdk:"max_security_groups"`
	MaxVolumes                                     *int64    `json:"max_volumes" tfsdk:"max_volumes"`
	MaximalResourceCountPerProject                 *int64    `json:"maximal_resource_count_per_project" tfsdk:"maximal_resource_count_per_project"`
	MinimalTeamCountForProvisioning                *int64    `json:"minimal_team_count_for_provisioning" tfsdk:"minimal_team_count_for_provisioning"`
	OpenstackOfferingUuidList                      *[]string `json:"openstack_offering_uuid_list" tfsdk:"openstack_offering_uuid_list"`
	ProjectPermanentDirectory                      *string   `json:"project_permanent_directory" tfsdk:"project_permanent_directory"`
	RequirePurchaseOrderUpload                     *bool     `json:"require_purchase_order_upload" tfsdk:"require_purchase_order_upload"`
	RequiredTeamRoleForProvisioning                *string   `json:"required_team_role_for_provisioning" tfsdk:"required_team_role_for_provisioning"`
	ResourceExpirationThreshold                    *int64    `json:"resource_expiration_threshold" tfsdk:"resource_expiration_threshold"`
	ScratchProjectDirectory                        *string   `json:"scratch_project_directory" tfsdk:"scratch_project_directory"`
	ServiceProviderCanCreateOfferingUser           *bool     `json:"service_provider_can_create_offering_user" tfsdk:"service_provider_can_create_offering_user"`
	SlurmPeriodicPolicyEnabled                     *bool     `json:"slurm_periodic_policy_enabled" tfsdk:"slurm_periodic_policy_enabled"`
	SnapshotSizeLimitGb                            *int64    `json:"snapshot_size_limit_gb" tfsdk:"snapshot_size_limit_gb"`
	StorageMode                                    *string   `json:"storage_mode" tfsdk:"storage_mode"`
	SupportsDownscaling                            *bool     `json:"supports_downscaling" tfsdk:"supports_downscaling"`
	SupportsPausing                                *bool     `json:"supports_pausing" tfsdk:"supports_pausing"`
	UniqueResourcePerAttribute                     *string   `json:"unique_resource_per_attribute" tfsdk:"unique_resource_per_attribute"`
	UsernameAnonymizedPrefix                       *string   `json:"username_anonymized_prefix" tfsdk:"username_anonymized_prefix"`
	UsernameGenerationPolicy                       *string   `json:"username_generation_policy" tfsdk:"username_generation_policy"`
}

type MarketplaceOfferingPromotionCampaignsResponse struct {
	Description     *string `json:"description" tfsdk:"description"`
	Discount        *int64  `json:"discount" tfsdk:"discount"`
	DiscountType    *string `json:"discount_type" tfsdk:"discount_type"`
	EndDate         *string `json:"end_date" tfsdk:"end_date"`
	Months          *int64  `json:"months" tfsdk:"months"`
	Name            *string `json:"name" tfsdk:"name"`
	ServiceProvider *string `json:"service_provider" tfsdk:"service_provider"`
	StartDate       *string `json:"start_date" tfsdk:"start_date"`
	Stock           *int64  `json:"stock" tfsdk:"stock"`
	Uuid            *string `json:"uuid" tfsdk:"uuid"`
}

type MarketplaceOfferingQuotasResponse struct {
	Limit *int64  `json:"limit" tfsdk:"limit"`
	Name  *string `json:"name" tfsdk:"name"`
	Usage *int64  `json:"usage" tfsdk:"usage"`
}

type MarketplaceOfferingResourceOptionsResponse struct {
	Order *[]string `json:"order" tfsdk:"order"`
}

type MarketplaceOfferingRolesResponse struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
	Uuid *string `json:"uuid" tfsdk:"uuid"`
}

type MarketplaceOfferingScreenshotsResponse struct {
	Description *string `json:"description" tfsdk:"description"`
	Image       *string `json:"image" tfsdk:"image"`
	Name        *string `json:"name" tfsdk:"name"`
	Thumbnail   *string `json:"thumbnail" tfsdk:"thumbnail"`
	Uuid        *string `json:"uuid" tfsdk:"uuid"`
}

type MarketplaceOfferingSoftwareCatalogsResponse struct {
	Catalog      *MarketplaceOfferingSoftwareCatalogsCatalogResponse   `json:"catalog" tfsdk:"catalog"`
	PackageCount *int64                                                `json:"package_count" tfsdk:"package_count"`
	Partition    *MarketplaceOfferingSoftwareCatalogsPartitionResponse `json:"partition" tfsdk:"partition"`
	Uuid         *string                                               `json:"uuid" tfsdk:"uuid"`
}

type MarketplaceOfferingSoftwareCatalogsCatalogResponse struct {
	Description *string `json:"description" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
	Uuid        *string `json:"uuid" tfsdk:"uuid"`
	Version     *string `json:"version" tfsdk:"version"`
}

type MarketplaceOfferingSoftwareCatalogsPartitionResponse struct {
	PartitionName *string `json:"partition_name" tfsdk:"partition_name"`
	PriorityTier  *int64  `json:"priority_tier" tfsdk:"priority_tier"`
	Qos           *string `json:"qos" tfsdk:"qos"`
	Uuid          *string `json:"uuid" tfsdk:"uuid"`
}

type MarketplaceOfferingTagsResponse struct {
	Name *string `json:"name" tfsdk:"name"`
	Uuid *string `json:"uuid" tfsdk:"uuid"`
}

func (r *MarketplaceOfferingResponse) GetState() string {
	if r.State != nil {
		return *r.State
	}
	return "OK"
}

func (r *MarketplaceOfferingResponse) GetErrorMessage() string {
	return ""
}
