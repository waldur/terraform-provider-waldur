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
var _ datasource.DataSource = &MarketplaceOfferingDataSource{}

func NewMarketplaceOfferingDataSource() datasource.DataSource {
	return &MarketplaceOfferingDataSource{}
}

// MarketplaceOfferingDataSource defines the data source implementation.
type MarketplaceOfferingDataSource struct {
	client *client.Client
}

// MarketplaceOfferingApiResponse is the API response model.
type MarketplaceOfferingApiResponse struct {
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
	Archived           *bool                                                `json:"archived" tfsdk:"archived"`
	ArticleCode        *string                                              `json:"article_code" tfsdk:"article_code"`
	BackendId          *string                                              `json:"backend_id" tfsdk:"backend_id"`
	Components         []MarketplaceOfferingPlansComponentsResponse         `json:"components" tfsdk:"components"`
	Description        *string                                              `json:"description" tfsdk:"description"`
	InitPrice          *float64                                             `json:"init_price" tfsdk:"init_price"`
	IsActive           *bool                                                `json:"is_active" tfsdk:"is_active"`
	MaxAmount          *int64                                               `json:"max_amount" tfsdk:"max_amount"`
	MinimalPrice       *float64                                             `json:"minimal_price" tfsdk:"minimal_price"`
	Name               *string                                              `json:"name" tfsdk:"name"`
	OrganizationGroups []MarketplaceOfferingPlansOrganizationGroupsResponse `json:"organization_groups" tfsdk:"organization_groups"`
	PlanType           *string                                              `json:"plan_type" tfsdk:"plan_type"`
	ResourcesCount     *int64                                               `json:"resources_count" tfsdk:"resources_count"`
	SwitchPrice        *float64                                             `json:"switch_price" tfsdk:"switch_price"`
	Unit               *string                                              `json:"unit" tfsdk:"unit"`
	UnitPrice          *string                                              `json:"unit_price" tfsdk:"unit_price"`
	Url                *string                                              `json:"url" tfsdk:"url"`
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

// MarketplaceOfferingDataSourceModel describes the data source data model.
type MarketplaceOfferingDataSourceModel struct {
	UUID                      types.String  `tfsdk:"id"`
	AccessibleViaCalls        types.Bool    `tfsdk:"accessible_via_calls"`
	AllowedCustomerUuid       types.String  `tfsdk:"allowed_customer_uuid"`
	Attributes                types.String  `tfsdk:"attributes"`
	Billable                  types.Bool    `tfsdk:"billable"`
	CanCreateOfferingUser     types.Bool    `tfsdk:"can_create_offering_user"`
	CategoryGroupUuid         types.String  `tfsdk:"category_group_uuid"`
	CategoryUuid              types.String  `tfsdk:"category_uuid"`
	Created                   types.String  `tfsdk:"created"`
	Customer                  types.String  `tfsdk:"customer"`
	CustomerUuid              types.String  `tfsdk:"customer_uuid"`
	Description               types.String  `tfsdk:"description"`
	HasActiveTermsOfService   types.Bool    `tfsdk:"has_active_terms_of_service"`
	HasTermsOfService         types.Bool    `tfsdk:"has_terms_of_service"`
	Keyword                   types.String  `tfsdk:"keyword"`
	Modified                  types.String  `tfsdk:"modified"`
	Name                      types.String  `tfsdk:"name"`
	NameExact                 types.String  `tfsdk:"name_exact"`
	OrganizationGroupUuid     types.String  `tfsdk:"organization_group_uuid"`
	ParentUuid                types.String  `tfsdk:"parent_uuid"`
	ProjectUuid               types.String  `tfsdk:"project_uuid"`
	Query                     types.String  `tfsdk:"query"`
	ResourceCustomerUuid      types.String  `tfsdk:"resource_customer_uuid"`
	ResourceProjectUuid       types.String  `tfsdk:"resource_project_uuid"`
	ScopeUuid                 types.String  `tfsdk:"scope_uuid"`
	ServiceManagerUuid        types.String  `tfsdk:"service_manager_uuid"`
	Shared                    types.Bool    `tfsdk:"shared"`
	State                     types.String  `tfsdk:"state"`
	Type                      types.String  `tfsdk:"type"`
	UserHasConsent            types.Bool    `tfsdk:"user_has_consent"`
	UserHasOfferingUser       types.Bool    `tfsdk:"user_has_offering_user"`
	UuidList                  types.String  `tfsdk:"uuid_list"`
	AccessUrl                 types.String  `tfsdk:"access_url"`
	BackendId                 types.String  `tfsdk:"backend_id"`
	BillingTypeClassification types.String  `tfsdk:"billing_type_classification"`
	Category                  types.String  `tfsdk:"category"`
	CategoryTitle             types.String  `tfsdk:"category_title"`
	CitationCount             types.Int64   `tfsdk:"citation_count"`
	ComplianceChecklist       types.String  `tfsdk:"compliance_checklist"`
	Components                types.List    `tfsdk:"components"`
	Country                   types.String  `tfsdk:"country"`
	DataciteDoi               types.String  `tfsdk:"datacite_doi"`
	Endpoints                 types.List    `tfsdk:"endpoints"`
	Files                     types.List    `tfsdk:"files"`
	FullDescription           types.String  `tfsdk:"full_description"`
	GettingStarted            types.String  `tfsdk:"getting_started"`
	GoogleCalendarIsPublic    types.Bool    `tfsdk:"google_calendar_is_public"`
	GoogleCalendarLink        types.String  `tfsdk:"google_calendar_link"`
	HasComplianceRequirements types.Bool    `tfsdk:"has_compliance_requirements"`
	Image                     types.String  `tfsdk:"image"`
	IntegrationGuide          types.String  `tfsdk:"integration_guide"`
	Latitude                  types.Float64 `tfsdk:"latitude"`
	Longitude                 types.Float64 `tfsdk:"longitude"`
	OrderCount                types.Int64   `tfsdk:"order_count"`
	OrganizationGroups        types.List    `tfsdk:"organization_groups"`
	ParentDescription         types.String  `tfsdk:"parent_description"`
	ParentName                types.String  `tfsdk:"parent_name"`
	Partitions                types.List    `tfsdk:"partitions"`
	PausedReason              types.String  `tfsdk:"paused_reason"`
	Plans                     types.List    `tfsdk:"plans"`
	PrivacyPolicyLink         types.String  `tfsdk:"privacy_policy_link"`
	PromotionCampaigns        types.List    `tfsdk:"promotion_campaigns"`
	Quotas                    types.List    `tfsdk:"quotas"`
	Roles                     types.List    `tfsdk:"roles"`
	Scope                     types.String  `tfsdk:"scope"`
	ScopeErrorMessage         types.String  `tfsdk:"scope_error_message"`
	ScopeName                 types.String  `tfsdk:"scope_name"`
	ScopeState                types.String  `tfsdk:"scope_state"`
	Screenshots               types.List    `tfsdk:"screenshots"`
	Slug                      types.String  `tfsdk:"slug"`
	SoftwareCatalogs          types.List    `tfsdk:"software_catalogs"`
	Thumbnail                 types.String  `tfsdk:"thumbnail"`
	TotalCost                 types.Int64   `tfsdk:"total_cost"`
	TotalCostEstimated        types.Int64   `tfsdk:"total_cost_estimated"`
	TotalCustomers            types.Int64   `tfsdk:"total_customers"`
	Url                       types.String  `tfsdk:"url"`
	VendorDetails             types.String  `tfsdk:"vendor_details"`
}

func (d *MarketplaceOfferingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_offering"
}

func (d *MarketplaceOfferingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Marketplace Offering data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"accessible_via_calls": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Accessible via calls",
			},
			"allowed_customer_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Allowed customer UUID",
			},
			"attributes": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Offering attributes (JSON)",
			},
			"billable": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Billable",
			},
			"can_create_offering_user": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Can create offering user",
			},
			"category_group_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Category group UUID",
			},
			"category_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Category UUID",
			},
			"created": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Created after",
			},
			"customer": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer URL",
			},
			"customer_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer UUID",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description contains",
			},
			"has_active_terms_of_service": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Has Active Terms of Service",
			},
			"has_terms_of_service": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Has Terms of Service",
			},
			"keyword": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Keyword",
			},
			"modified": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Modified after",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name (exact)",
			},
			"organization_group_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Organization group UUID",
			},
			"parent_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Parent offering UUID",
			},
			"project_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project UUID",
			},
			"query": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Search by offering name, slug or description",
			},
			"resource_customer_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Resource customer UUID",
			},
			"resource_project_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Resource project UUID",
			},
			"scope_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Scope UUID",
			},
			"service_manager_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Service manager UUID",
			},
			"shared": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Shared",
			},
			"state": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Offering state",
			},
			"type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Offering type",
			},
			"user_has_consent": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "User Has Consent",
			},
			"user_has_offering_user": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "User Has Offering User",
			},
			"uuid_list": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Comma-separated offering UUIDs",
			},
			"access_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Publicly accessible offering access URL",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of the backend",
			},
			"billing_type_classification": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Classify offering components by billing type. Returns 'limit_only', 'usage_only', or 'mixed'.",
			},
			"category": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Category",
			},
			"category_title": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Category title",
			},
			"citation_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of citations of a DOI",
			},
			"compliance_checklist": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Compliance checklist",
			},
			"components": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"article_code": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Article code",
						},
						"billing_type": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Billing type",
						},
						"default_limit": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Default limit",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Description of the resource",
						},
						"factor": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Factor",
						},
						"is_boolean": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "Is boolean",
						},
						"is_builtin": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Is builtin",
						},
						"is_prepaid": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "Is prepaid",
						},
						"limit_amount": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Limit amount",
						},
						"limit_period": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Limit period",
						},
						"max_available_limit": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Max available limit",
						},
						"max_prepaid_duration": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Max prepaid duration",
						},
						"max_value": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Max value",
						},
						"measured_unit": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Unit of measurement, for example, GB.",
						},
						"min_prepaid_duration": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Min prepaid duration",
						},
						"min_value": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Min value",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Display name for the measured unit, for example, Floating IP.",
						},
						"overage_component": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Overage component",
						},
						"type": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Unique internal name of the measured unit, for example floating_ip.",
						},
						"unit_factor": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "The conversion factor from backend units to measured_unit",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Components",
			},
			"country": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Country code (ISO 3166-1 alpha-2)",
			},
			"datacite_doi": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Datacite doi",
			},
			"endpoints": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the resource",
						},
						"url": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "URL of the access endpoint",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Endpoints",
			},
			"files": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Created",
						},
						"file": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "File",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the resource",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Files",
			},
			"full_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Full description",
			},
			"getting_started": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Getting started",
			},
			"google_calendar_is_public": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Google calendar is public",
			},
			"google_calendar_link": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Get the Google Calendar link for an offering.",
			},
			"has_compliance_requirements": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Has compliance requirements",
			},
			"image": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Image",
			},
			"integration_guide": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Integration guide",
			},
			"latitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "Latitude",
			},
			"longitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "Longitude",
			},
			"order_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Order count",
			},
			"organization_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"customers_count": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Number of customers in this organization group",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the resource",
						},
						"parent": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Parent",
						},
						"parent_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name of the parent organization group",
						},
						"parent_uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "UUID of the parent organization group",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Url",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Organization groups",
			},
			"parent_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Parent description",
			},
			"parent_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the parent",
			},
			"partitions": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cpu_bind": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Default task binding policy (SLURM cpu_bind)",
						},
						"def_cpu_per_gpu": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Default CPUs allocated per GPU",
						},
						"def_mem_per_cpu": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Default memory per CPU in MB",
						},
						"def_mem_per_gpu": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Default memory per GPU in MB",
						},
						"def_mem_per_node": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Default memory per node in MB",
						},
						"default_time": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Default time limit in minutes",
						},
						"exclusive_topo": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "Exclusive topology access required",
						},
						"exclusive_user": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "Exclusive user access required",
						},
						"grace_time": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Preemption grace time in seconds",
						},
						"max_cpus_per_node": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Maximum allocated CPUs per node",
						},
						"max_cpus_per_socket": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Maximum allocated CPUs per socket",
						},
						"max_mem_per_cpu": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Maximum memory per CPU in MB",
						},
						"max_mem_per_node": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Maximum memory per node in MB",
						},
						"max_nodes": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Maximum nodes per job",
						},
						"max_time": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Maximum time limit in minutes",
						},
						"min_nodes": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Minimum nodes per job",
						},
						"partition_name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the SLURM partition",
						},
						"priority_tier": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Priority tier for scheduling and preemption",
						},
						"qos": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Quality of Service (QOS) name",
						},
						"req_resv": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "Require reservation for job allocation",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Partitions",
			},
			"paused_reason": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Paused reason",
			},
			"plans": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"archived": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "Forbids creation of new resources.",
						},
						"article_code": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Article code",
						},
						"backend_id": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "ID of the backend",
						},
						"components": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"amount": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "Amount",
									},
									"discount_rate": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "Discount rate in percentage.",
									},
									"discount_threshold": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "Minimum amount to be eligible for discount.",
									},
									"future_price": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Future price",
									},
									"measured_unit": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Unit of measurement, for example, GB.",
									},
									"name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Display name for the measured unit, for example, Floating IP.",
									},
									"price": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Price",
									},
									"type": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Unique internal name of the measured unit, for example floating_ip.",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: "Components",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Description of the resource",
						},
						"init_price": schema.Float64Attribute{
							Computed:            true,
							MarkdownDescription: "Init price",
						},
						"is_active": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Is active",
						},
						"max_amount": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Maximum number of plans that could be active. Plan is disabled when maximum amount is reached.",
						},
						"minimal_price": schema.Float64Attribute{
							Computed:            true,
							MarkdownDescription: "Minimal price",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the resource",
						},
						"organization_groups": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"customers_count": schema.Int64Attribute{
										Computed:            true,
										MarkdownDescription: "Number of customers in this organization group",
									},
									"name": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Name of the resource",
									},
									"parent": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Parent",
									},
									"parent_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Name of the parent organization group",
									},
									"parent_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "UUID of the parent organization group",
									},
									"url": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Url",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: "Organization groups",
						},
						"plan_type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Plan type",
						},
						"resources_count": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Resources count",
						},
						"switch_price": schema.Float64Attribute{
							Computed:            true,
							MarkdownDescription: "Switch price",
						},
						"unit": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Unit",
						},
						"unit_price": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Unit price",
						},
						"url": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Url",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Plans",
			},
			"privacy_policy_link": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Privacy policy link",
			},
			"promotion_campaigns": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Description of the resource",
						},
						"discount": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Discount",
						},
						"discount_type": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Discount type",
						},
						"end_date": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "The last day the campaign is active.",
						},
						"months": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "How many months in a row should the related service (when activated) get special deal (0 for indefinitely until active)",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the resource",
						},
						"service_provider": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Service provider",
						},
						"start_date": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Starting from this date, the campaign is active.",
						},
						"stock": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Stock",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Promotion campaigns",
			},
			"quotas": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"limit": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Limit",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the resource",
						},
						"usage": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Usage",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Quotas",
			},
			"roles": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the resource",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Url",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Roles",
			},
			"scope": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Scope",
			},
			"scope_error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Scope error message",
			},
			"scope_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the scope",
			},
			"scope_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Scope state",
			},
			"screenshots": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Created",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Description of the resource",
						},
						"image": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Image",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the resource",
						},
						"thumbnail": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Thumbnail",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Screenshots",
			},
			"slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
			},
			"software_catalogs": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"catalog": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Optional:            true,
									MarkdownDescription: "Description of the resource",
								},
								"name": schema.StringAttribute{
									Optional:            true,
									MarkdownDescription: "Name of the resource",
								},
								"version": schema.StringAttribute{
									Optional:            true,
									MarkdownDescription: "Version",
								},
							},
							Computed:            true,
							MarkdownDescription: "Catalog",
						},
						"package_count": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Package count",
						},
						"partition": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"partition_name": schema.StringAttribute{
									Optional:            true,
									MarkdownDescription: "Name of the partition",
								},
								"priority_tier": schema.Int64Attribute{
									Optional:            true,
									MarkdownDescription: "Priority tier",
								},
								"qos": schema.StringAttribute{
									Optional:            true,
									MarkdownDescription: "Qos",
								},
							},
							Computed:            true,
							MarkdownDescription: "Partition",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Software catalogs",
			},
			"thumbnail": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Thumbnail",
			},
			"total_cost": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Total cost",
			},
			"total_cost_estimated": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Total cost estimated",
			},
			"total_customers": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Total customers",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Url",
			},
			"vendor_details": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Vendor details",
			},
		},
	}
}

func (d *MarketplaceOfferingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *MarketplaceOfferingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data MarketplaceOfferingDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var apiResp MarketplaceOfferingApiResponse

		err := d.client.GetByUUID(ctx, "/api/marketplace-public-offerings/{uuid}/", data.UUID.ValueString(), &apiResp)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Marketplace Offering",
				"An error occurred while reading the Marketplace Offering by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, apiResp, &data)...)

	} else {
		// Filter by provided parameters
		var results []MarketplaceOfferingApiResponse

		type filterDef struct {
			name string
			val  attr.Value
		}
		filterDefs := []filterDef{
			{"accessible_via_calls", data.AccessibleViaCalls},
			{"allowed_customer_uuid", data.AllowedCustomerUuid},
			{"attributes", data.Attributes},
			{"billable", data.Billable},
			{"can_create_offering_user", data.CanCreateOfferingUser},
			{"category_group_uuid", data.CategoryGroupUuid},
			{"category_uuid", data.CategoryUuid},
			{"created", data.Created},
			{"customer", data.Customer},
			{"customer_uuid", data.CustomerUuid},
			{"description", data.Description},
			{"has_active_terms_of_service", data.HasActiveTermsOfService},
			{"has_terms_of_service", data.HasTermsOfService},
			{"keyword", data.Keyword},
			{"modified", data.Modified},
			{"name", data.Name},
			{"name_exact", data.NameExact},
			{"organization_group_uuid", data.OrganizationGroupUuid},
			{"parent_uuid", data.ParentUuid},
			{"project_uuid", data.ProjectUuid},
			{"query", data.Query},
			{"resource_customer_uuid", data.ResourceCustomerUuid},
			{"resource_project_uuid", data.ResourceProjectUuid},
			{"scope_uuid", data.ScopeUuid},
			{"service_manager_uuid", data.ServiceManagerUuid},
			{"shared", data.Shared},
			{"state", data.State},
			{"type", data.Type},
			{"user_has_consent", data.UserHasConsent},
			{"user_has_offering_user", data.UserHasOfferingUser},
			{"uuid_list", data.UuidList},
		}

		filters := make(map[string]string)
		for _, fd := range filterDefs {
			if fd.val.IsNull() || fd.val.IsUnknown() {
				continue
			}
			switch v := fd.val.(type) {
			case types.String:
				filters[fd.name] = v.ValueString()
			case types.Int64:
				filters[fd.name] = fmt.Sprintf("%d", v.ValueInt64())
			case types.Bool:
				filters[fd.name] = fmt.Sprintf("%t", v.ValueBool())
			case types.Float64:
				filters[fd.name] = fmt.Sprintf("%f", v.ValueFloat64())
			}
		}

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup marketplace_offering.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/marketplace-public-offerings/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Marketplace Offering",
				"An error occurred while filtering Marketplace Offering: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Marketplace Offering Not Found",
				"No Marketplace Offering found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Marketplace Offerings Found",
				fmt.Sprintf("Found %d Marketplace Offerings with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *MarketplaceOfferingDataSource) mapResponseToModel(ctx context.Context, apiResp MarketplaceOfferingApiResponse, model *MarketplaceOfferingDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Billable = types.BoolPointerValue(apiResp.Billable)
	model.BillingTypeClassification = types.StringPointerValue(apiResp.BillingTypeClassification)
	model.Category = types.StringPointerValue(apiResp.Category)
	model.CategoryTitle = types.StringPointerValue(apiResp.CategoryTitle)
	model.CategoryUuid = types.StringPointerValue(apiResp.CategoryUuid)
	model.CitationCount = types.Int64PointerValue(apiResp.CitationCount)
	model.ComplianceChecklist = types.StringPointerValue(apiResp.ComplianceChecklist)
	listValComponents, listDiagsComponents := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"article_code":         types.StringType,
		"billing_type":         types.StringType,
		"default_limit":        types.Int64Type,
		"description":          types.StringType,
		"factor":               types.Int64Type,
		"is_boolean":           types.BoolType,
		"is_builtin":           types.BoolType,
		"is_prepaid":           types.BoolType,
		"limit_amount":         types.Int64Type,
		"limit_period":         types.StringType,
		"max_available_limit":  types.Int64Type,
		"max_prepaid_duration": types.Int64Type,
		"max_value":            types.Int64Type,
		"measured_unit":        types.StringType,
		"min_prepaid_duration": types.Int64Type,
		"min_value":            types.Int64Type,
		"name":                 types.StringType,
		"overage_component":    types.StringType,
		"type":                 types.StringType,
		"unit_factor":          types.Int64Type,
	}}, apiResp.Components)
	diags.Append(listDiagsComponents...)
	model.Components = listValComponents
	model.Country = types.StringPointerValue(apiResp.Country)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.DataciteDoi = types.StringPointerValue(apiResp.DataciteDoi)
	model.Description = types.StringPointerValue(apiResp.Description)
	listValEndpoints, listDiagsEndpoints := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"name": types.StringType,
		"url":  types.StringType,
	}}, apiResp.Endpoints)
	diags.Append(listDiagsEndpoints...)
	model.Endpoints = listValEndpoints
	listValFiles, listDiagsFiles := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"created": types.StringType,
		"file":    types.StringType,
		"name":    types.StringType,
	}}, apiResp.Files)
	diags.Append(listDiagsFiles...)
	model.Files = listValFiles
	model.FullDescription = types.StringPointerValue(apiResp.FullDescription)
	model.GettingStarted = types.StringPointerValue(apiResp.GettingStarted)
	model.GoogleCalendarIsPublic = types.BoolPointerValue(apiResp.GoogleCalendarIsPublic)
	model.GoogleCalendarLink = types.StringPointerValue(apiResp.GoogleCalendarLink)
	model.HasComplianceRequirements = types.BoolPointerValue(apiResp.HasComplianceRequirements)
	model.Image = types.StringPointerValue(apiResp.Image)
	model.IntegrationGuide = types.StringPointerValue(apiResp.IntegrationGuide)
	model.Latitude = types.Float64PointerValue(apiResp.Latitude)
	model.Longitude = types.Float64PointerValue(apiResp.Longitude)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.OrderCount = types.Int64PointerValue(apiResp.OrderCount)
	listValOrganizationGroups, listDiagsOrganizationGroups := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"customers_count": types.Int64Type,
		"name":            types.StringType,
		"parent":          types.StringType,
		"parent_name":     types.StringType,
		"parent_uuid":     types.StringType,
		"url":             types.StringType,
	}}, apiResp.OrganizationGroups)
	diags.Append(listDiagsOrganizationGroups...)
	model.OrganizationGroups = listValOrganizationGroups
	model.ParentDescription = types.StringPointerValue(apiResp.ParentDescription)
	model.ParentName = types.StringPointerValue(apiResp.ParentName)
	model.ParentUuid = types.StringPointerValue(apiResp.ParentUuid)
	listValPartitions, listDiagsPartitions := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"cpu_bind":            types.Int64Type,
		"def_cpu_per_gpu":     types.Int64Type,
		"def_mem_per_cpu":     types.Int64Type,
		"def_mem_per_gpu":     types.Int64Type,
		"def_mem_per_node":    types.Int64Type,
		"default_time":        types.Int64Type,
		"exclusive_topo":      types.BoolType,
		"exclusive_user":      types.BoolType,
		"grace_time":          types.Int64Type,
		"max_cpus_per_node":   types.Int64Type,
		"max_cpus_per_socket": types.Int64Type,
		"max_mem_per_cpu":     types.Int64Type,
		"max_mem_per_node":    types.Int64Type,
		"max_nodes":           types.Int64Type,
		"max_time":            types.Int64Type,
		"min_nodes":           types.Int64Type,
		"partition_name":      types.StringType,
		"priority_tier":       types.Int64Type,
		"qos":                 types.StringType,
		"req_resv":            types.BoolType,
	}}, apiResp.Partitions)
	diags.Append(listDiagsPartitions...)
	model.Partitions = listValPartitions
	model.PausedReason = types.StringPointerValue(apiResp.PausedReason)
	listValPlans, listDiagsPlans := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"archived":     types.BoolType,
		"article_code": types.StringType,
		"backend_id":   types.StringType,
		"components": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
			"amount":             types.Int64Type,
			"discount_rate":      types.Int64Type,
			"discount_threshold": types.Int64Type,
			"future_price":       types.StringType,
			"measured_unit":      types.StringType,
			"name":               types.StringType,
			"price":              types.StringType,
			"type":               types.StringType,
		}}},
		"description":   types.StringType,
		"init_price":    types.Float64Type,
		"is_active":     types.BoolType,
		"max_amount":    types.Int64Type,
		"minimal_price": types.Float64Type,
		"name":          types.StringType,
		"organization_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
			"customers_count": types.Int64Type,
			"name":            types.StringType,
			"parent":          types.StringType,
			"parent_name":     types.StringType,
			"parent_uuid":     types.StringType,
			"url":             types.StringType,
		}}},
		"plan_type":       types.StringType,
		"resources_count": types.Int64Type,
		"switch_price":    types.Float64Type,
		"unit":            types.StringType,
		"unit_price":      types.StringType,
		"url":             types.StringType,
	}}, apiResp.Plans)
	diags.Append(listDiagsPlans...)
	model.Plans = listValPlans
	model.PrivacyPolicyLink = types.StringPointerValue(apiResp.PrivacyPolicyLink)
	listValPromotionCampaigns, listDiagsPromotionCampaigns := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"description":      types.StringType,
		"discount":         types.Int64Type,
		"discount_type":    types.StringType,
		"end_date":         types.StringType,
		"months":           types.Int64Type,
		"name":             types.StringType,
		"service_provider": types.StringType,
		"start_date":       types.StringType,
		"stock":            types.Int64Type,
	}}, apiResp.PromotionCampaigns)
	diags.Append(listDiagsPromotionCampaigns...)
	model.PromotionCampaigns = listValPromotionCampaigns
	listValQuotas, listDiagsQuotas := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"limit": types.Int64Type,
		"name":  types.StringType,
		"usage": types.Int64Type,
	}}, apiResp.Quotas)
	diags.Append(listDiagsQuotas...)
	model.Quotas = listValQuotas
	listValRoles, listDiagsRoles := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"name": types.StringType,
		"url":  types.StringType,
	}}, apiResp.Roles)
	diags.Append(listDiagsRoles...)
	model.Roles = listValRoles
	model.Scope = types.StringPointerValue(apiResp.Scope)
	model.ScopeErrorMessage = types.StringPointerValue(apiResp.ScopeErrorMessage)
	model.ScopeName = types.StringPointerValue(apiResp.ScopeName)
	model.ScopeState = types.StringPointerValue(apiResp.ScopeState)
	model.ScopeUuid = types.StringPointerValue(apiResp.ScopeUuid)
	listValScreenshots, listDiagsScreenshots := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"created":     types.StringType,
		"description": types.StringType,
		"image":       types.StringType,
		"name":        types.StringType,
		"thumbnail":   types.StringType,
	}}, apiResp.Screenshots)
	diags.Append(listDiagsScreenshots...)
	model.Screenshots = listValScreenshots
	model.Shared = types.BoolPointerValue(apiResp.Shared)
	model.Slug = types.StringPointerValue(apiResp.Slug)
	listValSoftwareCatalogs, listDiagsSoftwareCatalogs := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"catalog": types.ObjectType{AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"name":        types.StringType,
			"version":     types.StringType,
		}},
		"package_count": types.Int64Type,
		"partition": types.ObjectType{AttrTypes: map[string]attr.Type{
			"partition_name": types.StringType,
			"priority_tier":  types.Int64Type,
			"qos":            types.StringType,
		}},
	}}, apiResp.SoftwareCatalogs)
	diags.Append(listDiagsSoftwareCatalogs...)
	model.SoftwareCatalogs = listValSoftwareCatalogs
	model.State = types.StringPointerValue(apiResp.State)
	model.Thumbnail = types.StringPointerValue(apiResp.Thumbnail)
	model.TotalCost = types.Int64PointerValue(apiResp.TotalCost)
	model.TotalCostEstimated = types.Int64PointerValue(apiResp.TotalCostEstimated)
	model.TotalCustomers = types.Int64PointerValue(apiResp.TotalCustomers)
	model.Type = types.StringPointerValue(apiResp.Type)
	model.Url = types.StringPointerValue(apiResp.Url)
	model.UserHasConsent = types.BoolPointerValue(apiResp.UserHasConsent)
	model.VendorDetails = types.StringPointerValue(apiResp.VendorDetails)

	return diags
}
