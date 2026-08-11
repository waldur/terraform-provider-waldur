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

	CategoryUuid *string `json:"category_uuid,omitempty" tfsdk:"category_uuid"`

	CitationCount *int64 `json:"citation_count,omitempty" tfsdk:"citation_count"`

	ComplianceChecklist *string `json:"compliance_checklist,omitempty" tfsdk:"compliance_checklist"`

	Components *[]common.OfferingComponent `json:"components,omitempty" tfsdk:"components"`

	ConfigDriveDefault *bool `json:"config_drive_default,omitempty" tfsdk:"config_drive_default"`

	Country *string `json:"country,omitempty" tfsdk:"country"`

	Customer *string `json:"customer,omitempty" tfsdk:"customer"`

	DataciteDoi *string `json:"datacite_doi,omitempty" tfsdk:"datacite_doi"`

	DefaultAccessSubnets *[]common.NestedOfferingAccessSubnet `json:"default_access_subnets,omitempty" tfsdk:"default_access_subnets"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	DocumentationUrl *string `json:"documentation_url,omitempty" tfsdk:"documentation_url"`

	EffectiveAvailableLimits *[]string `json:"effective_available_limits,omitempty" tfsdk:"effective_available_limits"`

	Endpoints *[]common.NestedEndpoint `json:"endpoints,omitempty" tfsdk:"endpoints"`

	Files *[]common.NestedOfferingFile `json:"files,omitempty" tfsdk:"files"`

	FullDescription *string `json:"full_description,omitempty" tfsdk:"full_description"`

	GettingStarted *string `json:"getting_started,omitempty" tfsdk:"getting_started"`

	GoogleCalendarIsPublic *bool `json:"google_calendar_is_public,omitempty" tfsdk:"google_calendar_is_public"`

	GoogleCalendarLink *string `json:"google_calendar_link,omitempty" tfsdk:"google_calendar_link"`

	HasComplianceRequirements *bool `json:"has_compliance_requirements,omitempty" tfsdk:"has_compliance_requirements"`

	HelpdeskUrl *string `json:"helpdesk_url,omitempty" tfsdk:"helpdesk_url"`

	Image *string `json:"image,omitempty" tfsdk:"image"`

	IntegrationGuide *string `json:"integration_guide,omitempty" tfsdk:"integration_guide"`

	IsAccessible *bool `json:"is_accessible,omitempty" tfsdk:"is_accessible"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	OfferingGroup *string `json:"offering_group,omitempty" tfsdk:"offering_group"`

	OfferingGroupTitle *string `json:"offering_group_title,omitempty" tfsdk:"offering_group_title"`

	OfferingGroupUuid *string `json:"offering_group_uuid,omitempty" tfsdk:"offering_group_uuid"`

	OpenForProposals *bool `json:"open_for_proposals,omitempty" tfsdk:"open_for_proposals"`

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

	ProfileName *string `json:"profile_name,omitempty" tfsdk:"profile_name"`

	ProfileUuid *string `json:"profile_uuid,omitempty" tfsdk:"profile_uuid"`

	Project *string `json:"project,omitempty" tfsdk:"project"`

	PromotionCampaigns *[]common.NestedCampaign `json:"promotion_campaigns,omitempty" tfsdk:"promotion_campaigns"`

	QosProfiles *[]common.NestedQoS `json:"qos_profiles,omitempty" tfsdk:"qos_profiles"`

	Quotas *[]common.Quota `json:"quotas,omitempty" tfsdk:"quotas"`

	ResourceOptions *MarketplaceOfferingResourceOptionsResponse `json:"resource_options,omitempty" tfsdk:"resource_options"`

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

	MaxRenewalDuration *int64 `json:"max_renewal_duration,omitempty" tfsdk:"max_renewal_duration"`

	MaxValue *int64 `json:"max_value,omitempty" tfsdk:"max_value"`

	MeasuredUnit *string `json:"measured_unit,omitempty" tfsdk:"measured_unit"`

	MinPrepaidDuration *int64 `json:"min_prepaid_duration,omitempty" tfsdk:"min_prepaid_duration"`

	MinRenewalDuration *int64 `json:"min_renewal_duration,omitempty" tfsdk:"min_renewal_duration"`

	MinValue *int64 `json:"min_value,omitempty" tfsdk:"min_value"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	OfferingUuid *string `json:"offering_uuid,omitempty" tfsdk:"offering_uuid"`

	OverageComponent *string `json:"overage_component,omitempty" tfsdk:"overage_component"`

	PrepaidDurationStep *int64 `json:"prepaid_duration_step,omitempty" tfsdk:"prepaid_duration_step"`

	RenewalDurationStep *int64 `json:"renewal_duration_step,omitempty" tfsdk:"renewal_duration_step"`

	Type *string `json:"type,omitempty" tfsdk:"type"`

	UnitFactor *int64 `json:"unit_factor,omitempty" tfsdk:"unit_factor"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceOfferingDefaultAccessSubnetsResponse struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`

	Inet *string `json:"inet,omitempty" tfsdk:"inet"`

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
	CpuArch *string `json:"cpu_arch,omitempty" tfsdk:"cpu_arch"`

	CpuBind *int64 `json:"cpu_bind,omitempty" tfsdk:"cpu_bind"`

	DefCpuPerGpu *int64 `json:"def_cpu_per_gpu,omitempty" tfsdk:"def_cpu_per_gpu"`

	DefMemPerCpu *int64 `json:"def_mem_per_cpu,omitempty" tfsdk:"def_mem_per_cpu"`

	DefMemPerGpu *int64 `json:"def_mem_per_gpu,omitempty" tfsdk:"def_mem_per_gpu"`

	DefMemPerNode *int64 `json:"def_mem_per_node,omitempty" tfsdk:"def_mem_per_node"`

	DefaultTime *int64 `json:"default_time,omitempty" tfsdk:"default_time"`

	ExclusiveTopo *bool `json:"exclusive_topo,omitempty" tfsdk:"exclusive_topo"`

	ExclusiveUser *bool `json:"exclusive_user,omitempty" tfsdk:"exclusive_user"`

	GpuArch *string `json:"gpu_arch,omitempty" tfsdk:"gpu_arch"`

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

	QosOptions *[]common.NestedPartitionQoS `json:"qos_options,omitempty" tfsdk:"qos_options"`

	ReqResv *bool `json:"req_resv,omitempty" tfsdk:"req_resv"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceOfferingPartitionsQosOptionsResponse struct {
	IsDefault *bool `json:"is_default,omitempty" tfsdk:"is_default"`

	Qos *string `json:"qos,omitempty" tfsdk:"qos"`

	QosName *string `json:"qos_name,omitempty" tfsdk:"qos_name"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceOfferingPlansResponse struct {
	Archived *bool `json:"archived,omitempty" tfsdk:"archived"`

	ArticleCode *string `json:"article_code,omitempty" tfsdk:"article_code"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Components *[]common.NestedPlanComponent `json:"components,omitempty" tfsdk:"components"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	FuturePrices map[string]string `json:"future_prices,omitempty" tfsdk:"future_prices"`

	InitPrice common.FlexibleNumber `json:"init_price,omitempty" tfsdk:"init_price"`

	IsActive *bool `json:"is_active,omitempty" tfsdk:"is_active"`

	MaxAmount *int64 `json:"max_amount,omitempty" tfsdk:"max_amount"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	OrganizationGroups *[]common.OrganizationGroup `json:"organization_groups,omitempty" tfsdk:"organization_groups"`

	PlanType *string `json:"plan_type,omitempty" tfsdk:"plan_type"`

	Prices map[string]string `json:"prices,omitempty" tfsdk:"prices"`

	Quotas map[string]int64 `json:"quotas,omitempty" tfsdk:"quotas"`

	ResourcesCount *int64 `json:"resources_count,omitempty" tfsdk:"resources_count"`

	SwitchPrice common.FlexibleNumber `json:"switch_price,omitempty" tfsdk:"switch_price"`

	Unit *string `json:"unit,omitempty" tfsdk:"unit"`

	UnitPrice *string `json:"unit_price,omitempty" tfsdk:"unit_price"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceOfferingPlansComponentsResponse struct {
	Amount *int64 `json:"amount,omitempty" tfsdk:"amount"`

	DiscountAggregation *string `json:"discount_aggregation,omitempty" tfsdk:"discount_aggregation"`

	DiscountDescription *string `json:"discount_description,omitempty" tfsdk:"discount_description"`

	DiscountFormula *string `json:"discount_formula,omitempty" tfsdk:"discount_formula"`

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
	ActionOnUsageLimit *string `json:"action_on_usage_limit,omitempty" tfsdk:"action_on_usage_limit"`

	AutoApproveForRoles *[]string `json:"auto_approve_for_roles,omitempty" tfsdk:"auto_approve_for_roles"`

	AutoApproveInServiceProviderProjects *bool `json:"auto_approve_in_service_provider_projects,omitempty" tfsdk:"auto_approve_in_service_provider_projects"`

	AutoApproveMarketplaceScript *bool `json:"auto_approve_marketplace_script,omitempty" tfsdk:"auto_approve_marketplace_script"`

	AutoApproveRemoteOrders *bool `json:"auto_approve_remote_orders,omitempty" tfsdk:"auto_approve_remote_orders"`

	AutoOkResourceProjects *bool `json:"auto_ok_resource_projects,omitempty" tfsdk:"auto_ok_resource_projects"`

	BackendIdDisplayLabel *string `json:"backend_id_display_label,omitempty" tfsdk:"backend_id_display_label"`

	BillingSource *string `json:"billing_source,omitempty" tfsdk:"billing_source"`

	CanRestoreResource *bool `json:"can_restore_resource,omitempty" tfsdk:"can_restore_resource"`

	ConcealBillingData *bool `json:"conceal_billing_data,omitempty" tfsdk:"conceal_billing_data"`

	ConcealSubnetRestrictedResources *bool `json:"conceal_subnet_restricted_resources,omitempty" tfsdk:"conceal_subnet_restricted_resources"`

	CreateOrdersOnResourceOptionChange *bool `json:"create_orders_on_resource_option_change,omitempty" tfsdk:"create_orders_on_resource_option_change"`

	CreateOrdersOnResourceProjectChange *bool `json:"create_orders_on_resource_project_change,omitempty" tfsdk:"create_orders_on_resource_project_change"`

	DefaultInternalNetworkMtu *int64 `json:"default_internal_network_mtu,omitempty" tfsdk:"default_internal_network_mtu"`

	DefaultResourceTerminationOffsetInDays *int64 `json:"default_resource_termination_offset_in_days,omitempty" tfsdk:"default_resource_termination_offset_in_days"`

	DeploymentMode *string `json:"deployment_mode,omitempty" tfsdk:"deployment_mode"`

	DisableAutoapprove *bool `json:"disable_autoapprove,omitempty" tfsdk:"disable_autoapprove"`

	DisableGracePeriod *bool `json:"disable_grace_period,omitempty" tfsdk:"disable_grace_period"`

	DisabledResourceActions *[]string `json:"disabled_resource_actions,omitempty" tfsdk:"disabled_resource_actions"`

	EmitDisplayName *bool `json:"emit_display_name,omitempty" tfsdk:"emit_display_name"`

	EmitWaldurUsername *bool `json:"emit_waldur_username,omitempty" tfsdk:"emit_waldur_username"`

	EnableDisplayOfOrderActionsForServiceProvider *bool `json:"enable_display_of_order_actions_for_service_provider,omitempty" tfsdk:"enable_display_of_order_actions_for_service_provider"`

	EnableIssuesForMembershipChanges *bool `json:"enable_issues_for_membership_changes,omitempty" tfsdk:"enable_issues_for_membership_changes"`

	EnableMembershipSyncStatus *bool `json:"enable_membership_sync_status,omitempty" tfsdk:"enable_membership_sync_status"`

	EnablePosixAccount *bool `json:"enable_posix_account,omitempty" tfsdk:"enable_posix_account"`

	EnableProviderConsumerMessaging *bool `json:"enable_provider_consumer_messaging,omitempty" tfsdk:"enable_provider_consumer_messaging"`

	EnablePurchaseOrderUpload *bool `json:"enable_purchase_order_upload,omitempty" tfsdk:"enable_purchase_order_upload"`

	EnableResourceAccessSubnets *bool `json:"enable_resource_access_subnets,omitempty" tfsdk:"enable_resource_access_subnets"`

	EnableResourceEndDateChangeRequests *bool `json:"enable_resource_end_date_change_requests,omitempty" tfsdk:"enable_resource_end_date_change_requests"`

	EnableResourceProjects *bool `json:"enable_resource_projects,omitempty" tfsdk:"enable_resource_projects"`

	EnforceQos *bool `json:"enforce_qos,omitempty" tfsdk:"enforce_qos"`

	ExposeInferencePlayground *bool `json:"expose_inference_playground,omitempty" tfsdk:"expose_inference_playground"`

	FlavorsRegex *string `json:"flavors_regex,omitempty" tfsdk:"flavors_regex"`

	GidSource *string `json:"gid_source,omitempty" tfsdk:"gid_source"`

	HeappeClusterId *string `json:"heappe_cluster_id,omitempty" tfsdk:"heappe_cluster_id"`

	HeappeLocalBasePath *string `json:"heappe_local_base_path,omitempty" tfsdk:"heappe_local_base_path"`

	HeappeUrl *string `json:"heappe_url,omitempty" tfsdk:"heappe_url"`

	HeappeUsername *string `json:"heappe_username,omitempty" tfsdk:"heappe_username"`

	HighlightBackendIdDisplay *bool `json:"highlight_backend_id_display,omitempty" tfsdk:"highlight_backend_id_display"`

	HomedirPrefix *string `json:"homedir_prefix,omitempty" tfsdk:"homedir_prefix"`

	IsResourceTerminationDateRequired *bool `json:"is_resource_termination_date_required,omitempty" tfsdk:"is_resource_termination_date_required"`

	LatestDateForResourceTermination *string `json:"latest_date_for_resource_termination,omitempty" tfsdk:"latest_date_for_resource_termination"`

	LbaasEnabled *bool `json:"lbaas_enabled,omitempty" tfsdk:"lbaas_enabled"`

	LoginShell *string `json:"login_shell,omitempty" tfsdk:"login_shell"`

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

	NotifyAboutProviderConsumerMessages *bool `json:"notify_about_provider_consumer_messages,omitempty" tfsdk:"notify_about_provider_consumer_messages"`

	OfferingUserAutoDeletion *bool `json:"offering_user_auto_deletion,omitempty" tfsdk:"offering_user_auto_deletion"`

	OpenstackOfferingUuidList *[]string `json:"openstack_offering_uuid_list,omitempty" tfsdk:"openstack_offering_uuid_list"`

	ProjectPermanentDirectory *string `json:"project_permanent_directory,omitempty" tfsdk:"project_permanent_directory"`

	RequireEffectiveIdForHighlightedDisplay *bool `json:"require_effective_id_for_highlighted_display,omitempty" tfsdk:"require_effective_id_for_highlighted_display"`

	RequirePurchaseOrderUpload *bool `json:"require_purchase_order_upload,omitempty" tfsdk:"require_purchase_order_upload"`

	RequiredTeamRoleForProvisioning *string `json:"required_team_role_for_provisioning,omitempty" tfsdk:"required_team_role_for_provisioning"`

	ResourceExpirationThreshold *int64 `json:"resource_expiration_threshold,omitempty" tfsdk:"resource_expiration_threshold"`

	ResourceNamePattern *string `json:"resource_name_pattern,omitempty" tfsdk:"resource_name_pattern"`

	ResourceProjectRoleGroupTemplate *string `json:"resource_project_role_group_template,omitempty" tfsdk:"resource_project_role_group_template"`

	ResourceProjectRoleMap map[string]string `json:"resource_project_role_map,omitempty" tfsdk:"resource_project_role_map"`

	ResourceProjectsLimitPolicy *string `json:"resource_projects_limit_policy,omitempty" tfsdk:"resource_projects_limit_policy"`

	ResourceProjectsLimitsRequired *bool `json:"resource_projects_limits_required,omitempty" tfsdk:"resource_projects_limits_required"`

	ResourceRoleGroupTemplate *string `json:"resource_role_group_template,omitempty" tfsdk:"resource_role_group_template"`

	ResourceRoleMap map[string]string `json:"resource_role_map,omitempty" tfsdk:"resource_role_map"`

	ResourceSlugMaxLength *int64 `json:"resource_slug_max_length,omitempty" tfsdk:"resource_slug_max_length"`

	ResourceSlugTemplate *string `json:"resource_slug_template,omitempty" tfsdk:"resource_slug_template"`

	RestrictDeletionWithActiveResources *bool `json:"restrict_deletion_with_active_resources,omitempty" tfsdk:"restrict_deletion_with_active_resources"`

	RestrictedToRoles *[]string `json:"restricted_to_roles,omitempty" tfsdk:"restricted_to_roles"`

	ScratchProjectDirectory *string `json:"scratch_project_directory,omitempty" tfsdk:"scratch_project_directory"`

	ServiceProviderCanCreateOfferingUser *bool `json:"service_provider_can_create_offering_user,omitempty" tfsdk:"service_provider_can_create_offering_user"`

	SlurmPeriodicPolicyEnabled *bool `json:"slurm_periodic_policy_enabled,omitempty" tfsdk:"slurm_periodic_policy_enabled"`

	SnapshotSizeLimitGb *int64 `json:"snapshot_size_limit_gb,omitempty" tfsdk:"snapshot_size_limit_gb"`

	StorageMode *string `json:"storage_mode,omitempty" tfsdk:"storage_mode"`

	SupportsDownscaling *bool `json:"supports_downscaling,omitempty" tfsdk:"supports_downscaling"`

	SupportsPausing *bool `json:"supports_pausing,omitempty" tfsdk:"supports_pausing"`

	UidSource *string `json:"uid_source,omitempty" tfsdk:"uid_source"`

	UniqueResourcePerAttribute *string `json:"unique_resource_per_attribute,omitempty" tfsdk:"unique_resource_per_attribute"`

	UsagePollIntervalMinutes *int64 `json:"usage_poll_interval_minutes,omitempty" tfsdk:"usage_poll_interval_minutes"`

	UsernameAnonymizedPrefix *string `json:"username_anonymized_prefix,omitempty" tfsdk:"username_anonymized_prefix"`

	UsernameGenerationPolicy *string `json:"username_generation_policy,omitempty" tfsdk:"username_generation_policy"`
}

type MarketplaceOfferingPluginOptionsResourceProjectRoleMapResponse struct {
}

type MarketplaceOfferingPluginOptionsResourceRoleMapResponse struct {
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

type MarketplaceOfferingQosProfilesResponse struct {
	DefaultTime *int64 `json:"default_time,omitempty" tfsdk:"default_time"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	Flags *string `json:"flags,omitempty" tfsdk:"flags"`

	GraceTime *int64 `json:"grace_time,omitempty" tfsdk:"grace_time"`

	GrpTres *string `json:"grp_tres,omitempty" tfsdk:"grp_tres"`

	MaxNodes *int64 `json:"max_nodes,omitempty" tfsdk:"max_nodes"`

	MaxTime *int64 `json:"max_time,omitempty" tfsdk:"max_time"`

	MaxTresPerJob *string `json:"max_tres_per_job,omitempty" tfsdk:"max_tres_per_job"`

	MaxTresPerNode *string `json:"max_tres_per_node,omitempty" tfsdk:"max_tres_per_node"`

	MaxTresPerUser *string `json:"max_tres_per_user,omitempty" tfsdk:"max_tres_per_user"`

	MinNodes *int64 `json:"min_nodes,omitempty" tfsdk:"min_nodes"`

	MinTresPerJob *string `json:"min_tres_per_job,omitempty" tfsdk:"min_tres_per_job"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Priority *int64 `json:"priority,omitempty" tfsdk:"priority"`

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

type MarketplaceOfferingScreenshotsResponse struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`

	Image *string `json:"image,omitempty" tfsdk:"image"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Thumbnail *string `json:"thumbnail,omitempty" tfsdk:"thumbnail"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type MarketplaceOfferingSoftwareCatalogsResponse struct {
	Catalog *MarketplaceOfferingSoftwareCatalogsCatalogResponse `json:"catalog,omitempty" tfsdk:"catalog"`

	EnabledCpuFamily *[]string `json:"enabled_cpu_family,omitempty" tfsdk:"enabled_cpu_family"`

	EnabledCpuMicroarchitectures *[]string `json:"enabled_cpu_microarchitectures,omitempty" tfsdk:"enabled_cpu_microarchitectures"`

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
	CpuArch *string `json:"cpu_arch,omitempty" tfsdk:"cpu_arch"`

	GpuArch *string `json:"gpu_arch,omitempty" tfsdk:"gpu_arch"`

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
