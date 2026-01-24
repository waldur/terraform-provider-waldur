package offering

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// MarketplaceOffering Structs

type MarketplaceOfferingCreateRequest struct {
	AccessUrl           *string                           `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId           *string                           `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Billable            *bool                             `json:"billable,omitempty" tfsdk:"billable"`
	Category            *string                           `json:"category" tfsdk:"category"`
	ComplianceChecklist *string                           `json:"compliance_checklist,omitempty" tfsdk:"compliance_checklist"`
	Components          []common.OfferingComponentRequest `json:"components,omitempty" tfsdk:"components"`
	Country             *string                           `json:"country,omitempty" tfsdk:"country"`
	Customer            *string                           `json:"customer,omitempty" tfsdk:"customer"`
	DataciteDoi         *string                           `json:"datacite_doi,omitempty" tfsdk:"datacite_doi"`
	Description         *string                           `json:"description,omitempty" tfsdk:"description"`
	FullDescription     *string                           `json:"full_description,omitempty" tfsdk:"full_description"`
	GettingStarted      *string                           `json:"getting_started,omitempty" tfsdk:"getting_started"`
	Image               *string                           `json:"image,omitempty" tfsdk:"image"`
	IntegrationGuide    *string                           `json:"integration_guide,omitempty" tfsdk:"integration_guide"`
	Latitude            *float64                          `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude           *float64                          `json:"longitude,omitempty" tfsdk:"longitude"`
	Name                *string                           `json:"name" tfsdk:"name"`
	Options             *common.OfferingOptionsRequest    `json:"options,omitempty" tfsdk:"options"`
	Plans               []common.BaseProviderPlanRequest  `json:"plans,omitempty" tfsdk:"plans"`
	PrivacyPolicyLink   *string                           `json:"privacy_policy_link,omitempty" tfsdk:"privacy_policy_link"`
	ResourceOptions     *common.OfferingOptionsRequest    `json:"resource_options,omitempty" tfsdk:"resource_options"`
	Shared              *bool                             `json:"shared,omitempty" tfsdk:"shared"`
	Slug                *string                           `json:"slug,omitempty" tfsdk:"slug"`
	Thumbnail           *string                           `json:"thumbnail,omitempty" tfsdk:"thumbnail"`
	Type                *string                           `json:"type" tfsdk:"type"`
	VendorDetails       *string                           `json:"vendor_details,omitempty" tfsdk:"vendor_details"`
}

type MarketplaceOfferingResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl                 *string                                         `json:"access_url" tfsdk:"access_url"`
	BackendId                 *string                                         `json:"backend_id" tfsdk:"backend_id"`
	Billable                  *bool                                           `json:"billable" tfsdk:"billable"`
	BillingTypeClassification *string                                         `json:"billing_type_classification" tfsdk:"billing_type_classification"`
	Category                  *string                                         `json:"category" tfsdk:"category"`
	CategoryTitle             *string                                         `json:"category_title" tfsdk:"category_title"`
	CategoryUuid              *string                                         `json:"category_uuid" tfsdk:"category_uuid"`
	CitationCount             *int64                                          `json:"citation_count" tfsdk:"citation_count"`
	ComplianceChecklist       *string                                         `json:"compliance_checklist" tfsdk:"compliance_checklist"`
	Components                []MarketplaceOfferingComponentsResponse         `json:"components" tfsdk:"components"`
	Country                   *string                                         `json:"country" tfsdk:"country"`
	Created                   *string                                         `json:"created" tfsdk:"created"`
	Customer                  *string                                         `json:"customer" tfsdk:"customer"`
	CustomerName              *string                                         `json:"customer_name" tfsdk:"customer_name"`
	CustomerUuid              *string                                         `json:"customer_uuid" tfsdk:"customer_uuid"`
	DataciteDoi               *string                                         `json:"datacite_doi" tfsdk:"datacite_doi"`
	Description               *string                                         `json:"description" tfsdk:"description"`
	Endpoints                 []MarketplaceOfferingEndpointsResponse          `json:"endpoints" tfsdk:"endpoints"`
	Files                     []MarketplaceOfferingFilesResponse              `json:"files" tfsdk:"files"`
	FullDescription           *string                                         `json:"full_description" tfsdk:"full_description"`
	GettingStarted            *string                                         `json:"getting_started" tfsdk:"getting_started"`
	GoogleCalendarIsPublic    *bool                                           `json:"google_calendar_is_public" tfsdk:"google_calendar_is_public"`
	GoogleCalendarLink        *string                                         `json:"google_calendar_link" tfsdk:"google_calendar_link"`
	HasComplianceRequirements *bool                                           `json:"has_compliance_requirements" tfsdk:"has_compliance_requirements"`
	Image                     *string                                         `json:"image" tfsdk:"image"`
	IntegrationGuide          *string                                         `json:"integration_guide" tfsdk:"integration_guide"`
	IntegrationStatus         []MarketplaceOfferingIntegrationStatusResponse  `json:"integration_status" tfsdk:"integration_status"`
	Latitude                  *float64                                        `json:"latitude" tfsdk:"latitude"`
	Longitude                 *float64                                        `json:"longitude" tfsdk:"longitude"`
	Name                      *string                                         `json:"name" tfsdk:"name"`
	OrderCount                *int64                                          `json:"order_count" tfsdk:"order_count"`
	OrganizationGroups        []MarketplaceOfferingOrganizationGroupsResponse `json:"organization_groups" tfsdk:"organization_groups"`
	ParentDescription         *string                                         `json:"parent_description" tfsdk:"parent_description"`
	ParentName                *string                                         `json:"parent_name" tfsdk:"parent_name"`
	ParentUuid                *string                                         `json:"parent_uuid" tfsdk:"parent_uuid"`
	Partitions                []MarketplaceOfferingPartitionsResponse         `json:"partitions" tfsdk:"partitions"`
	PausedReason              *string                                         `json:"paused_reason" tfsdk:"paused_reason"`
	Plans                     []MarketplaceOfferingPlansResponse              `json:"plans" tfsdk:"plans"`
	PrivacyPolicyLink         *string                                         `json:"privacy_policy_link" tfsdk:"privacy_policy_link"`
	Project                   *string                                         `json:"project" tfsdk:"project"`
	ProjectName               *string                                         `json:"project_name" tfsdk:"project_name"`
	ProjectUuid               *string                                         `json:"project_uuid" tfsdk:"project_uuid"`
	PromotionCampaigns        []MarketplaceOfferingPromotionCampaignsResponse `json:"promotion_campaigns" tfsdk:"promotion_campaigns"`
	Quotas                    []MarketplaceOfferingQuotasResponse             `json:"quotas" tfsdk:"quotas"`
	Roles                     []MarketplaceOfferingRolesResponse              `json:"roles" tfsdk:"roles"`
	Scope                     *string                                         `json:"scope" tfsdk:"scope"`
	ScopeErrorMessage         *string                                         `json:"scope_error_message" tfsdk:"scope_error_message"`
	ScopeName                 *string                                         `json:"scope_name" tfsdk:"scope_name"`
	ScopeState                *string                                         `json:"scope_state" tfsdk:"scope_state"`
	ScopeUuid                 *string                                         `json:"scope_uuid" tfsdk:"scope_uuid"`
	Screenshots               []MarketplaceOfferingScreenshotsResponse        `json:"screenshots" tfsdk:"screenshots"`
	Shared                    *bool                                           `json:"shared" tfsdk:"shared"`
	Slug                      *string                                         `json:"slug" tfsdk:"slug"`
	SoftwareCatalogs          []MarketplaceOfferingSoftwareCatalogsResponse   `json:"software_catalogs" tfsdk:"software_catalogs"`
	State                     *string                                         `json:"state" tfsdk:"state"`
	Thumbnail                 *string                                         `json:"thumbnail" tfsdk:"thumbnail"`
	TotalCost                 *int64                                          `json:"total_cost" tfsdk:"total_cost"`
	TotalCostEstimated        *int64                                          `json:"total_cost_estimated" tfsdk:"total_cost_estimated"`
	TotalCustomers            *int64                                          `json:"total_customers" tfsdk:"total_customers"`
	Type                      *string                                         `json:"type" tfsdk:"type"`
	Url                       *string                                         `json:"url" tfsdk:"url"`
	UserHasConsent            *bool                                           `json:"user_has_consent" tfsdk:"user_has_consent"`
	VendorDetails             *string                                         `json:"vendor_details" tfsdk:"vendor_details"`
}

type MarketplaceOfferingComponentsResponse struct {
	ArticleCode        *string `json:"article_code" tfsdk:"article_code"`
	BillingType        *string `json:"billing_type" tfsdk:"billing_type"`
	DefaultLimit       *int64  `json:"default_limit" tfsdk:"default_limit"`
	Description        *string `json:"description" tfsdk:"description"`
	IsBoolean          *bool   `json:"is_boolean" tfsdk:"is_boolean"`
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
}

type MarketplaceOfferingEndpointsResponse struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
}

type MarketplaceOfferingFilesResponse struct {
	Created *string `json:"created" tfsdk:"created"`
	File    *string `json:"file" tfsdk:"file"`
	Name    *string `json:"name" tfsdk:"name"`
}

type MarketplaceOfferingIntegrationStatusResponse struct {
	AgentType            *string `json:"agent_type" tfsdk:"agent_type"`
	LastRequestTimestamp *string `json:"last_request_timestamp" tfsdk:"last_request_timestamp"`
	ServiceName          *string `json:"service_name" tfsdk:"service_name"`
	Status               *string `json:"status" tfsdk:"status"`
}

type MarketplaceOfferingOrganizationGroupsResponse struct {
	CustomersCount *int64  `json:"customers_count" tfsdk:"customers_count"`
	Name           *string `json:"name" tfsdk:"name"`
	Parent         *string `json:"parent" tfsdk:"parent"`
	ParentName     *string `json:"parent_name" tfsdk:"parent_name"`
	ParentUuid     *string `json:"parent_uuid" tfsdk:"parent_uuid"`
	Url            *string `json:"url" tfsdk:"url"`
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
}

type MarketplaceOfferingPlansResponse struct {
	Archived    *bool   `json:"archived" tfsdk:"archived"`
	ArticleCode *string `json:"article_code" tfsdk:"article_code"`
	BackendId   *string `json:"backend_id" tfsdk:"backend_id"`
	Description *string `json:"description" tfsdk:"description"`
	MaxAmount   *int64  `json:"max_amount" tfsdk:"max_amount"`
	Name        *string `json:"name" tfsdk:"name"`
	Unit        *string `json:"unit" tfsdk:"unit"`
	UnitPrice   *string `json:"unit_price" tfsdk:"unit_price"`
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
}

type MarketplaceOfferingQuotasResponse struct {
	Limit *int64  `json:"limit" tfsdk:"limit"`
	Name  *string `json:"name" tfsdk:"name"`
	Usage *int64  `json:"usage" tfsdk:"usage"`
}

type MarketplaceOfferingRolesResponse struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
}

type MarketplaceOfferingScreenshotsResponse struct {
	Created     *string `json:"created" tfsdk:"created"`
	Description *string `json:"description" tfsdk:"description"`
	Image       *string `json:"image" tfsdk:"image"`
	Name        *string `json:"name" tfsdk:"name"`
	Thumbnail   *string `json:"thumbnail" tfsdk:"thumbnail"`
}

type MarketplaceOfferingSoftwareCatalogsResponse struct {
	Catalog      *MarketplaceOfferingSoftwareCatalogsCatalogResponse   `json:"catalog" tfsdk:"catalog"`
	PackageCount *int64                                                `json:"package_count" tfsdk:"package_count"`
	Partition    *MarketplaceOfferingSoftwareCatalogsPartitionResponse `json:"partition" tfsdk:"partition"`
}

type MarketplaceOfferingSoftwareCatalogsCatalogResponse struct {
	Description *string `json:"description" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
	Version     *string `json:"version" tfsdk:"version"`
}

type MarketplaceOfferingSoftwareCatalogsPartitionResponse struct {
	PartitionName *string `json:"partition_name" tfsdk:"partition_name"`
	PriorityTier  *int64  `json:"priority_tier" tfsdk:"priority_tier"`
	Qos           *string `json:"qos" tfsdk:"qos"`
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
