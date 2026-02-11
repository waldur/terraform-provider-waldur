package offering

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type MarketplaceOfferingCreateRequest struct {
}

type MarketplaceOfferingResponse struct {
	UUID *string `json:"uuid"`

	Attributes map[string]interface{} `json:"attributes,omitempty" tfsdk:"attributes"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Billable *bool `json:"billable,omitempty" tfsdk:"billable"`

	BillingTypeClassification *string `json:"billing_type_classification,omitempty" tfsdk:"billing_type_classification"`

	Category *string `json:"category,omitempty" tfsdk:"category"`

	CategoryTitle *string `json:"category_title,omitempty" tfsdk:"category_title"`

	CategoryUuid *string `json:"category_uuid,omitempty" tfsdk:"category_uuid"`

	CitationCount *int64 `json:"citation_count,omitempty" tfsdk:"citation_count"`

	ComplianceChecklist *string `json:"compliance_checklist,omitempty" tfsdk:"compliance_checklist"`

	Components *[]common.OfferingComponent `json:"components,omitempty" tfsdk:"components"`

	Country *string `json:"country,omitempty" tfsdk:"country"`

	Customer *string `json:"customer,omitempty" tfsdk:"customer"`

	DataciteDoi *string `json:"datacite_doi,omitempty" tfsdk:"datacite_doi"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	Endpoints *[]common.NestedEndpoint `json:"endpoints,omitempty" tfsdk:"endpoints"`

	Files *[]common.NestedOfferingFile `json:"files,omitempty" tfsdk:"files"`

	FullDescription *string `json:"full_description,omitempty" tfsdk:"full_description"`

	GettingStarted *string `json:"getting_started,omitempty" tfsdk:"getting_started"`

	GoogleCalendarIsPublic *bool `json:"google_calendar_is_public,omitempty" tfsdk:"google_calendar_is_public"`

	GoogleCalendarLink *string `json:"google_calendar_link,omitempty" tfsdk:"google_calendar_link"`

	HasComplianceRequirements *bool `json:"has_compliance_requirements,omitempty" tfsdk:"has_compliance_requirements"`

	Image *string `json:"image,omitempty" tfsdk:"image"`

	IntegrationGuide *string `json:"integration_guide,omitempty" tfsdk:"integration_guide"`

	IsAccessible *bool `json:"is_accessible,omitempty" tfsdk:"is_accessible"`

	Latitude common.FlexibleNumber `json:"latitude,omitempty" tfsdk:"latitude"`

	Longitude common.FlexibleNumber `json:"longitude,omitempty" tfsdk:"longitude"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Options *MarketplaceOfferingOptionsResponse `json:"options,omitempty" tfsdk:"options"`

	OrderCount *int64 `json:"order_count,omitempty" tfsdk:"order_count"`

	OrganizationGroups *[]common.OrganizationGroup `json:"organization_groups,omitempty" tfsdk:"organization_groups"`

	ParentDescription *string `json:"parent_description,omitempty" tfsdk:"parent_description"`

	ParentName *string `json:"parent_name,omitempty" tfsdk:"parent_name"`

	ParentUuid *string `json:"parent_uuid,omitempty" tfsdk:"parent_uuid"`

	Partitions *[]common.NestedPartition `json:"partitions,omitempty" tfsdk:"partitions"`

	PausedReason *string `json:"paused_reason,omitempty" tfsdk:"paused_reason"`

	Plans *[]common.BasePublicPlan `json:"plans,omitempty" tfsdk:"plans"`

	PluginOptions *MarketplaceOfferingPluginOptionsResponse `json:"plugin_options,omitempty" tfsdk:"plugin_options"`

	PrivacyPolicyLink *string `json:"privacy_policy_link,omitempty" tfsdk:"privacy_policy_link"`

	Project *string `json:"project,omitempty" tfsdk:"project"`

	PromotionCampaigns *[]common.NestedCampaign `json:"promotion_campaigns,omitempty" tfsdk:"promotion_campaigns"`

	Quotas *[]common.Quota `json:"quotas,omitempty" tfsdk:"quotas"`

	ResourceOptions *MarketplaceOfferingResourceOptionsResponse `json:"resource_options,omitempty" tfsdk:"resource_options"`

	Roles *[]common.NestedRole `json:"roles,omitempty" tfsdk:"roles"`

	Scope *string `json:"scope,omitempty" tfsdk:"scope"`

	ScopeErrorMessage *string `json:"scope_error_message,omitempty" tfsdk:"scope_error_message"`

	ScopeName *string `json:"scope_name,omitempty" tfsdk:"scope_name"`

	ScopeState *string `json:"scope_state,omitempty" tfsdk:"scope_state"`

	ScopeUuid *string `json:"scope_uuid,omitempty" tfsdk:"scope_uuid"`

	Screenshots *[]common.NestedScreenshot `json:"screenshots,omitempty" tfsdk:"screenshots"`

	Shared *bool `json:"shared,omitempty" tfsdk:"shared"`

	Slug *string `json:"slug,omitempty" tfsdk:"slug"`

	SoftwareCatalogs *[]common.NestedSoftwareCatalog `json:"software_catalogs,omitempty" tfsdk:"software_catalogs"`

	State *string `json:"state,omitempty" tfsdk:"state"`

	Tags *[]common.NestedTag `json:"tags,omitempty" tfsdk:"tags"`

	Thumbnail *string `json:"thumbnail,omitempty" tfsdk:"thumbnail"`

	TotalCost *int64 `json:"total_cost,omitempty" tfsdk:"total_cost"`

	TotalCostEstimated *int64 `json:"total_cost_estimated,omitempty" tfsdk:"total_cost_estimated"`

	TotalCustomers *int64 `json:"total_customers,omitempty" tfsdk:"total_customers"`

	Type *string `json:"type,omitempty" tfsdk:"type"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	UserHasConsent *bool `json:"user_has_consent,omitempty" tfsdk:"user_has_consent"`

	VendorDetails *string `json:"vendor_details,omitempty" tfsdk:"vendor_details"`
}

type MarketplaceOfferingAttributesResponse struct {
}

type MarketplaceOfferingComponentsResponse struct {
	ArticleCode *string `json:"article_code,omitempty" tfsdk:"article_code"`

	BillingType *string `json:"billing_type,omitempty" tfsdk:"billing_type"`

	DefaultLimit *int64 `json:"default_limit,omitempty" tfsdk:"default_limit"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	Factor *int64 `json:"factor,omitempty" tfsdk:"factor"`

	IsBoolean *bool `json:"is_boolean,omitempty" tfsdk:"is_boolean"`

	IsBuiltin *bool `json:"is_builtin,omitempty" tfsdk:"is_builtin"`

	IsPrepaid *bool `json:"is_prepaid,omitempty" tfsdk:"is_prepaid"`

	LimitAmount *int64 `json:"limit_amount,omitempty" tfsdk:"limit_amount"`

	LimitPeriod *string `json:"limit_period,omitempty" tfsdk:"limit_period"`

	MaxAvailableLimit *int64 `json:"max_available_limit,omitempty" tfsdk:"max_available_limit"`

	MaxPrepaidDuration *int64 `json:"max_prepaid_duration,omitempty" tfsdk:"max_prepaid_duration"`

	MaxValue *int64 `json:"max_value,omitempty" tfsdk:"max_value"`

	MeasuredUnit *string `json:"measured_unit,omitempty" tfsdk:"measured_unit"`

	MinPrepaidDuration *int64 `json:"min_prepaid_duration,omitempty" tfsdk:"min_prepaid_duration"`

	MinValue *int64 `json:"min_value,omitempty" tfsdk:"min_value"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	OverageComponent *string `json:"overage_component,omitempty" tfsdk:"overage_component"`

	Type *string `json:"type,omitempty" tfsdk:"type"`

	UnitFactor *int64 `json:"unit_factor,omitempty" tfsdk:"unit_factor"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceOfferingEndpointsResponse struct {
	Name *string `json:"name,omitempty" tfsdk:"name"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceOfferingFilesResponse struct {
	File *string `json:"file,omitempty" tfsdk:"file"`

	Name *string `json:"name,omitempty" tfsdk:"name"`
}

type MarketplaceOfferingOptionsResponse struct {
	Options map[string]interface{} `json:"options,omitempty" tfsdk:"options"`

	Order *[]string `json:"order,omitempty" tfsdk:"order"`
}

type MarketplaceOfferingOptionsOptionsResponse struct {
}

type MarketplaceOfferingOrganizationGroupsResponse struct {
	CustomersCount *int64 `json:"customers_count,omitempty" tfsdk:"customers_count"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Parent *string `json:"parent,omitempty" tfsdk:"parent"`

	ParentName *string `json:"parent_name,omitempty" tfsdk:"parent_name"`

	ParentUuid *string `json:"parent_uuid,omitempty" tfsdk:"parent_uuid"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceOfferingPartitionsResponse struct {
	CpuBind *int64 `json:"cpu_bind,omitempty" tfsdk:"cpu_bind"`

	DefCpuPerGpu *int64 `json:"def_cpu_per_gpu,omitempty" tfsdk:"def_cpu_per_gpu"`

	DefMemPerCpu *int64 `json:"def_mem_per_cpu,omitempty" tfsdk:"def_mem_per_cpu"`

	DefMemPerGpu *int64 `json:"def_mem_per_gpu,omitempty" tfsdk:"def_mem_per_gpu"`

	DefMemPerNode *int64 `json:"def_mem_per_node,omitempty" tfsdk:"def_mem_per_node"`

	DefaultTime *int64 `json:"default_time,omitempty" tfsdk:"default_time"`

	ExclusiveTopo *bool `json:"exclusive_topo,omitempty" tfsdk:"exclusive_topo"`

	ExclusiveUser *bool `json:"exclusive_user,omitempty" tfsdk:"exclusive_user"`

	GraceTime *int64 `json:"grace_time,omitempty" tfsdk:"grace_time"`

	MaxCpusPerNode *int64 `json:"max_cpus_per_node,omitempty" tfsdk:"max_cpus_per_node"`

	MaxCpusPerSocket *int64 `json:"max_cpus_per_socket,omitempty" tfsdk:"max_cpus_per_socket"`

	MaxMemPerCpu *int64 `json:"max_mem_per_cpu,omitempty" tfsdk:"max_mem_per_cpu"`

	MaxMemPerNode *int64 `json:"max_mem_per_node,omitempty" tfsdk:"max_mem_per_node"`

	MaxNodes *int64 `json:"max_nodes,omitempty" tfsdk:"max_nodes"`

	MaxTime *int64 `json:"max_time,omitempty" tfsdk:"max_time"`

	MinNodes *int64 `json:"min_nodes,omitempty" tfsdk:"min_nodes"`

	PartitionName *string `json:"partition_name,omitempty" tfsdk:"partition_name"`

	PriorityTier *int64 `json:"priority_tier,omitempty" tfsdk:"priority_tier"`

	Qos *string `json:"qos,omitempty" tfsdk:"qos"`

	ReqResv *bool `json:"req_resv,omitempty" tfsdk:"req_resv"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceOfferingPlansResponse struct {
	Archived *bool `json:"archived,omitempty" tfsdk:"archived"`

	ArticleCode *string `json:"article_code,omitempty" tfsdk:"article_code"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Components *[]common.NestedPlanComponent `json:"components,omitempty" tfsdk:"components"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	FuturePrices map[string]float64 `json:"future_prices,omitempty" tfsdk:"future_prices"`

	InitPrice common.FlexibleNumber `json:"init_price,omitempty" tfsdk:"init_price"`

	IsActive *bool `json:"is_active,omitempty" tfsdk:"is_active"`

	MaxAmount *int64 `json:"max_amount,omitempty" tfsdk:"max_amount"`

	MinimalPrice common.FlexibleNumber `json:"minimal_price,omitempty" tfsdk:"minimal_price"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	OrganizationGroups *[]common.OrganizationGroup `json:"organization_groups,omitempty" tfsdk:"organization_groups"`

	PlanType *string `json:"plan_type,omitempty" tfsdk:"plan_type"`

	Prices map[string]string `json:"prices,omitempty" tfsdk:"prices"`

	Quotas map[string]float64 `json:"quotas,omitempty" tfsdk:"quotas"`

	ResourcesCount *int64 `json:"resources_count,omitempty" tfsdk:"resources_count"`

	SwitchPrice common.FlexibleNumber `json:"switch_price,omitempty" tfsdk:"switch_price"`

	Unit *string `json:"unit,omitempty" tfsdk:"unit"`

	UnitPrice *string `json:"unit_price,omitempty" tfsdk:"unit_price"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceOfferingPlansComponentsResponse struct {
	Amount *int64 `json:"amount,omitempty" tfsdk:"amount"`

	DiscountRate *int64 `json:"discount_rate,omitempty" tfsdk:"discount_rate"`

	DiscountThreshold *int64 `json:"discount_threshold,omitempty" tfsdk:"discount_threshold"`

	FuturePrice *string `json:"future_price,omitempty" tfsdk:"future_price"`

	MeasuredUnit *string `json:"measured_unit,omitempty" tfsdk:"measured_unit"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Price *string `json:"price,omitempty" tfsdk:"price"`

	Type *string `json:"type,omitempty" tfsdk:"type"`
}

type MarketplaceOfferingPlansFuturePricesResponse struct {
}

type MarketplaceOfferingPlansOrganizationGroupsResponse struct {
	CustomersCount *int64 `json:"customers_count,omitempty" tfsdk:"customers_count"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Parent *string `json:"parent,omitempty" tfsdk:"parent"`

	ParentName *string `json:"parent_name,omitempty" tfsdk:"parent_name"`

	ParentUuid *string `json:"parent_uuid,omitempty" tfsdk:"parent_uuid"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceOfferingPlansPricesResponse struct {
}

type MarketplaceOfferingPlansQuotasResponse struct {
}

type MarketplaceOfferingPluginOptionsResponse struct {
	AutoApproveInServiceProviderProjects *bool `json:"auto_approve_in_service_provider_projects,omitempty" tfsdk:"auto_approve_in_service_provider_projects"`

	AutoApproveMarketplaceScript *bool `json:"auto_approve_marketplace_script,omitempty" tfsdk:"auto_approve_marketplace_script"`

	AutoApproveRemoteOrders *bool `json:"auto_approve_remote_orders,omitempty" tfsdk:"auto_approve_remote_orders"`

	BackendIdDisplayLabel *string `json:"backend_id_display_label,omitempty" tfsdk:"backend_id_display_label"`

	CanRestoreResource *bool `json:"can_restore_resource,omitempty" tfsdk:"can_restore_resource"`

	ConcealBillingData *bool `json:"conceal_billing_data,omitempty" tfsdk:"conceal_billing_data"`

	CreateOrdersOnResourceOptionChange *bool `json:"create_orders_on_resource_option_change,omitempty" tfsdk:"create_orders_on_resource_option_change"`

	DefaultInternalNetworkMtu *int64 `json:"default_internal_network_mtu,omitempty" tfsdk:"default_internal_network_mtu"`

	DefaultResourceTerminationOffsetInDays *int64 `json:"default_resource_termination_offset_in_days,omitempty" tfsdk:"default_resource_termination_offset_in_days"`

	DeploymentMode *string `json:"deployment_mode,omitempty" tfsdk:"deployment_mode"`

	DisableAutoapprove *bool `json:"disable_autoapprove,omitempty" tfsdk:"disable_autoapprove"`

	EnableDisplayOfOrderActionsForServiceProvider *bool `json:"enable_display_of_order_actions_for_service_provider,omitempty" tfsdk:"enable_display_of_order_actions_for_service_provider"`

	EnableIssuesForMembershipChanges *bool `json:"enable_issues_for_membership_changes,omitempty" tfsdk:"enable_issues_for_membership_changes"`

	EnablePurchaseOrderUpload *bool `json:"enable_purchase_order_upload,omitempty" tfsdk:"enable_purchase_order_upload"`

	FlavorsRegex *string `json:"flavors_regex,omitempty" tfsdk:"flavors_regex"`

	HeappeClusterId *string `json:"heappe_cluster_id,omitempty" tfsdk:"heappe_cluster_id"`

	HeappeLocalBasePath *string `json:"heappe_local_base_path,omitempty" tfsdk:"heappe_local_base_path"`

	HeappeUrl *string `json:"heappe_url,omitempty" tfsdk:"heappe_url"`

	HeappeUsername *string `json:"heappe_username,omitempty" tfsdk:"heappe_username"`

	HighlightBackendIdDisplay *bool `json:"highlight_backend_id_display,omitempty" tfsdk:"highlight_backend_id_display"`

	HomedirPrefix *string `json:"homedir_prefix,omitempty" tfsdk:"homedir_prefix"`

	InitialPrimarygroupNumber *int64 `json:"initial_primarygroup_number,omitempty" tfsdk:"initial_primarygroup_number"`

	InitialUidnumber *int64 `json:"initial_uidnumber,omitempty" tfsdk:"initial_uidnumber"`

	InitialUsergroupNumber *int64 `json:"initial_usergroup_number,omitempty" tfsdk:"initial_usergroup_number"`

	IsResourceTerminationDateRequired *bool `json:"is_resource_termination_date_required,omitempty" tfsdk:"is_resource_termination_date_required"`

	LatestDateForResourceTermination *string `json:"latest_date_for_resource_termination,omitempty" tfsdk:"latest_date_for_resource_termination"`

	ManagedRancherLoadBalancerDataVolumeSizeGb *int64 `json:"managed_rancher_load_balancer_data_volume_size_gb,omitempty" tfsdk:"managed_rancher_load_balancer_data_volume_size_gb"`

	ManagedRancherLoadBalancerDataVolumeTypeName *string `json:"managed_rancher_load_balancer_data_volume_type_name,omitempty" tfsdk:"managed_rancher_load_balancer_data_volume_type_name"`

	ManagedRancherLoadBalancerFlavorName *string `json:"managed_rancher_load_balancer_flavor_name,omitempty" tfsdk:"managed_rancher_load_balancer_flavor_name"`

	ManagedRancherLoadBalancerSystemVolumeSizeGb *int64 `json:"managed_rancher_load_balancer_system_volume_size_gb,omitempty" tfsdk:"managed_rancher_load_balancer_system_volume_size_gb"`

	ManagedRancherLoadBalancerSystemVolumeTypeName *string `json:"managed_rancher_load_balancer_system_volume_type_name,omitempty" tfsdk:"managed_rancher_load_balancer_system_volume_type_name"`

	ManagedRancherServerDataVolumeSizeGb *int64 `json:"managed_rancher_server_data_volume_size_gb,omitempty" tfsdk:"managed_rancher_server_data_volume_size_gb"`

	ManagedRancherServerDataVolumeTypeName *string `json:"managed_rancher_server_data_volume_type_name,omitempty" tfsdk:"managed_rancher_server_data_volume_type_name"`

	ManagedRancherServerFlavorName *string `json:"managed_rancher_server_flavor_name,omitempty" tfsdk:"managed_rancher_server_flavor_name"`

	ManagedRancherServerSystemVolumeSizeGb *int64 `json:"managed_rancher_server_system_volume_size_gb,omitempty" tfsdk:"managed_rancher_server_system_volume_size_gb"`

	ManagedRancherServerSystemVolumeTypeName *string `json:"managed_rancher_server_system_volume_type_name,omitempty" tfsdk:"managed_rancher_server_system_volume_type_name"`

	ManagedRancherTenantMaxCpu *int64 `json:"managed_rancher_tenant_max_cpu,omitempty" tfsdk:"managed_rancher_tenant_max_cpu"`

	ManagedRancherTenantMaxDisk *int64 `json:"managed_rancher_tenant_max_disk,omitempty" tfsdk:"managed_rancher_tenant_max_disk"`

	ManagedRancherTenantMaxRam *int64 `json:"managed_rancher_tenant_max_ram,omitempty" tfsdk:"managed_rancher_tenant_max_ram"`

	ManagedRancherWorkerSystemVolumeSizeGb *int64 `json:"managed_rancher_worker_system_volume_size_gb,omitempty" tfsdk:"managed_rancher_worker_system_volume_size_gb"`

	ManagedRancherWorkerSystemVolumeTypeName *string `json:"managed_rancher_worker_system_volume_type_name,omitempty" tfsdk:"managed_rancher_worker_system_volume_type_name"`

	MaxInstances *int64 `json:"max_instances,omitempty" tfsdk:"max_instances"`

	MaxResourceTerminationOffsetInDays *int64 `json:"max_resource_termination_offset_in_days,omitempty" tfsdk:"max_resource_termination_offset_in_days"`

	MaxSecurityGroups *int64 `json:"max_security_groups,omitempty" tfsdk:"max_security_groups"`

	MaxVolumes *int64 `json:"max_volumes,omitempty" tfsdk:"max_volumes"`

	MaximalResourceCountPerProject *int64 `json:"maximal_resource_count_per_project,omitempty" tfsdk:"maximal_resource_count_per_project"`

	MinimalTeamCountForProvisioning *int64 `json:"minimal_team_count_for_provisioning,omitempty" tfsdk:"minimal_team_count_for_provisioning"`

	OpenstackOfferingUuidList *[]string `json:"openstack_offering_uuid_list,omitempty" tfsdk:"openstack_offering_uuid_list"`

	ProjectPermanentDirectory *string `json:"project_permanent_directory,omitempty" tfsdk:"project_permanent_directory"`

	RequirePurchaseOrderUpload *bool `json:"require_purchase_order_upload,omitempty" tfsdk:"require_purchase_order_upload"`

	RequiredTeamRoleForProvisioning *string `json:"required_team_role_for_provisioning,omitempty" tfsdk:"required_team_role_for_provisioning"`

	ResourceExpirationThreshold *int64 `json:"resource_expiration_threshold,omitempty" tfsdk:"resource_expiration_threshold"`

	ScratchProjectDirectory *string `json:"scratch_project_directory,omitempty" tfsdk:"scratch_project_directory"`

	ServiceProviderCanCreateOfferingUser *bool `json:"service_provider_can_create_offering_user,omitempty" tfsdk:"service_provider_can_create_offering_user"`

	SlurmPeriodicPolicyEnabled *bool `json:"slurm_periodic_policy_enabled,omitempty" tfsdk:"slurm_periodic_policy_enabled"`

	SnapshotSizeLimitGb *int64 `json:"snapshot_size_limit_gb,omitempty" tfsdk:"snapshot_size_limit_gb"`

	StorageMode *string `json:"storage_mode,omitempty" tfsdk:"storage_mode"`

	SupportsDownscaling *bool `json:"supports_downscaling,omitempty" tfsdk:"supports_downscaling"`

	SupportsPausing *bool `json:"supports_pausing,omitempty" tfsdk:"supports_pausing"`

	UniqueResourcePerAttribute *string `json:"unique_resource_per_attribute,omitempty" tfsdk:"unique_resource_per_attribute"`

	UsernameAnonymizedPrefix *string `json:"username_anonymized_prefix,omitempty" tfsdk:"username_anonymized_prefix"`

	UsernameGenerationPolicy *string `json:"username_generation_policy,omitempty" tfsdk:"username_generation_policy"`
}

type MarketplaceOfferingPromotionCampaignsResponse struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`

	Discount *int64 `json:"discount,omitempty" tfsdk:"discount"`

	DiscountType *string `json:"discount_type,omitempty" tfsdk:"discount_type"`

	EndDate *string `json:"end_date,omitempty" tfsdk:"end_date"`

	Months *int64 `json:"months,omitempty" tfsdk:"months"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	ServiceProvider *string `json:"service_provider,omitempty" tfsdk:"service_provider"`

	StartDate *string `json:"start_date,omitempty" tfsdk:"start_date"`

	Stock *int64 `json:"stock,omitempty" tfsdk:"stock"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceOfferingQuotasResponse struct {
	Limit *int64 `json:"limit,omitempty" tfsdk:"limit"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Usage *int64 `json:"usage,omitempty" tfsdk:"usage"`
}

type MarketplaceOfferingResourceOptionsResponse struct {
	Options map[string]interface{} `json:"options,omitempty" tfsdk:"options"`

	Order *[]string `json:"order,omitempty" tfsdk:"order"`
}

type MarketplaceOfferingResourceOptionsOptionsResponse struct {
}

type MarketplaceOfferingRolesResponse struct {
	Name *string `json:"name,omitempty" tfsdk:"name"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceOfferingScreenshotsResponse struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`

	Image *string `json:"image,omitempty" tfsdk:"image"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Thumbnail *string `json:"thumbnail,omitempty" tfsdk:"thumbnail"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceOfferingSoftwareCatalogsResponse struct {
	Catalog *MarketplaceOfferingSoftwareCatalogsCatalogResponse `json:"catalog,omitempty" tfsdk:"catalog"`

	PackageCount *int64 `json:"package_count,omitempty" tfsdk:"package_count"`

	Partition *MarketplaceOfferingSoftwareCatalogsPartitionResponse `json:"partition,omitempty" tfsdk:"partition"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceOfferingSoftwareCatalogsCatalogResponse struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`

	Version *string `json:"version,omitempty" tfsdk:"version"`
}

type MarketplaceOfferingSoftwareCatalogsPartitionResponse struct {
	PartitionName *string `json:"partition_name,omitempty" tfsdk:"partition_name"`

	PriorityTier *int64 `json:"priority_tier,omitempty" tfsdk:"priority_tier"`

	Qos *string `json:"qos,omitempty" tfsdk:"qos"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceOfferingTagsResponse struct {
	Name *string `json:"name,omitempty" tfsdk:"name"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
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
