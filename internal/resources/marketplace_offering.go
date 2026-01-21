package resources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MarketplaceOfferingResource{}
var _ resource.ResourceWithImportState = &MarketplaceOfferingResource{}

func NewMarketplaceOfferingResource() resource.Resource {
	return &MarketplaceOfferingResource{}
}

// MarketplaceOfferingResource defines the resource implementation.
type MarketplaceOfferingResource struct {
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
	Options                   *MarketplaceOfferingOptionsResponse             `json:"options" tfsdk:"options"`
	OrderCount                *int64                                          `json:"order_count" tfsdk:"order_count"`
	OrganizationGroups        []MarketplaceOfferingOrganizationGroupsResponse `json:"organization_groups" tfsdk:"organization_groups"`
	ParentDescription         *string                                         `json:"parent_description" tfsdk:"parent_description"`
	ParentName                *string                                         `json:"parent_name" tfsdk:"parent_name"`
	ParentUuid                *string                                         `json:"parent_uuid" tfsdk:"parent_uuid"`
	Partitions                []MarketplaceOfferingPartitionsResponse         `json:"partitions" tfsdk:"partitions"`
	PausedReason              *string                                         `json:"paused_reason" tfsdk:"paused_reason"`
	Plans                     []MarketplaceOfferingPlansResponse              `json:"plans" tfsdk:"plans"`
	PrivacyPolicyLink         *string                                         `json:"privacy_policy_link" tfsdk:"privacy_policy_link"`
	Quotas                    []MarketplaceOfferingQuotasResponse             `json:"quotas" tfsdk:"quotas"`
	ResourceOptions           *MarketplaceOfferingResourceOptionsResponse     `json:"resource_options" tfsdk:"resource_options"`
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
	OverageComponent   *string `json:"overage_component" tfsdk:"overage_component"`
	Type               *string `json:"type" tfsdk:"type"`
	UnitFactor         *int64  `json:"unit_factor" tfsdk:"unit_factor"`
}

type MarketplaceOfferingEndpointsResponse struct {
	Url *string `json:"url" tfsdk:"url"`
}

type MarketplaceOfferingFilesResponse struct {
	Created *string `json:"created" tfsdk:"created"`
	File    *string `json:"file" tfsdk:"file"`
}

type MarketplaceOfferingIntegrationStatusResponse struct {
	AgentType            *string `json:"agent_type" tfsdk:"agent_type"`
	LastRequestTimestamp *string `json:"last_request_timestamp" tfsdk:"last_request_timestamp"`
	ServiceName          *string `json:"service_name" tfsdk:"service_name"`
	Status               *string `json:"status" tfsdk:"status"`
}

type MarketplaceOfferingOptionsResponse struct {
	Order []string `json:"order" tfsdk:"order"`
}

type MarketplaceOfferingOrganizationGroupsResponse struct {
	CustomersCount *int64  `json:"customers_count" tfsdk:"customers_count"`
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
	Unit        *string `json:"unit" tfsdk:"unit"`
	UnitPrice   *string `json:"unit_price" tfsdk:"unit_price"`
}

type MarketplaceOfferingQuotasResponse struct {
	Limit *int64 `json:"limit" tfsdk:"limit"`
	Usage *int64 `json:"usage" tfsdk:"usage"`
}

type MarketplaceOfferingResourceOptionsResponse struct {
	Order []string `json:"order" tfsdk:"order"`
}

type MarketplaceOfferingRolesResponse struct {
	Url *string `json:"url" tfsdk:"url"`
}

type MarketplaceOfferingScreenshotsResponse struct {
	Created     *string `json:"created" tfsdk:"created"`
	Description *string `json:"description" tfsdk:"description"`
	Image       *string `json:"image" tfsdk:"image"`
	Thumbnail   *string `json:"thumbnail" tfsdk:"thumbnail"`
}

type MarketplaceOfferingSoftwareCatalogsResponse struct {
	Catalog      *MarketplaceOfferingSoftwareCatalogsCatalogResponse   `json:"catalog" tfsdk:"catalog"`
	PackageCount *int64                                                `json:"package_count" tfsdk:"package_count"`
	Partition    *MarketplaceOfferingSoftwareCatalogsPartitionResponse `json:"partition" tfsdk:"partition"`
}

type MarketplaceOfferingSoftwareCatalogsCatalogResponse struct {
	Description *string `json:"description" tfsdk:"description"`
	Version     *string `json:"version" tfsdk:"version"`
}

type MarketplaceOfferingSoftwareCatalogsPartitionResponse struct {
	PartitionName *string `json:"partition_name" tfsdk:"partition_name"`
	PriorityTier  *int64  `json:"priority_tier" tfsdk:"priority_tier"`
	Qos           *string `json:"qos" tfsdk:"qos"`
}

var marketplaceoffering_componentsAttrTypes = map[string]attr.Type{
	"article_code":         types.StringType,
	"billing_type":         types.StringType,
	"default_limit":        types.Int64Type,
	"description":          types.StringType,
	"is_boolean":           types.BoolType,
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
}
var marketplaceoffering_componentsObjectType = types.ObjectType{
	AttrTypes: marketplaceoffering_componentsAttrTypes,
}

var marketplaceoffering_endpointsAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"url":  types.StringType,
}
var marketplaceoffering_endpointsObjectType = types.ObjectType{
	AttrTypes: marketplaceoffering_endpointsAttrTypes,
}

var marketplaceoffering_filesAttrTypes = map[string]attr.Type{
	"created": types.StringType,
	"file":    types.StringType,
	"name":    types.StringType,
}
var marketplaceoffering_filesObjectType = types.ObjectType{
	AttrTypes: marketplaceoffering_filesAttrTypes,
}

var marketplaceoffering_integration_statusAttrTypes = map[string]attr.Type{
	"agent_type":             types.StringType,
	"last_request_timestamp": types.StringType,
	"service_name":           types.StringType,
	"status":                 types.StringType,
}
var marketplaceoffering_integration_statusObjectType = types.ObjectType{
	AttrTypes: marketplaceoffering_integration_statusAttrTypes,
}

var marketplaceoffering_optionsAttrTypes = map[string]attr.Type{
	"order": types.ListType{ElemType: types.StringType},
}
var marketplaceoffering_optionsObjectType = types.ObjectType{
	AttrTypes: marketplaceoffering_optionsAttrTypes,
}

var marketplaceoffering_organization_groupsAttrTypes = map[string]attr.Type{
	"customers_count": types.Int64Type,
	"name":            types.StringType,
	"parent":          types.StringType,
	"parent_name":     types.StringType,
	"parent_uuid":     types.StringType,
	"url":             types.StringType,
}
var marketplaceoffering_organization_groupsObjectType = types.ObjectType{
	AttrTypes: marketplaceoffering_organization_groupsAttrTypes,
}

var marketplaceoffering_partitionsAttrTypes = map[string]attr.Type{
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
}
var marketplaceoffering_partitionsObjectType = types.ObjectType{
	AttrTypes: marketplaceoffering_partitionsAttrTypes,
}

var marketplaceoffering_plansAttrTypes = map[string]attr.Type{
	"archived":     types.BoolType,
	"article_code": types.StringType,
	"backend_id":   types.StringType,
	"description":  types.StringType,
	"max_amount":   types.Int64Type,
	"name":         types.StringType,
	"unit":         types.StringType,
	"unit_price":   types.StringType,
}
var marketplaceoffering_plansObjectType = types.ObjectType{
	AttrTypes: marketplaceoffering_plansAttrTypes,
}

var marketplaceoffering_quotasAttrTypes = map[string]attr.Type{
	"limit": types.Int64Type,
	"name":  types.StringType,
	"usage": types.Int64Type,
}
var marketplaceoffering_quotasObjectType = types.ObjectType{
	AttrTypes: marketplaceoffering_quotasAttrTypes,
}

var marketplaceoffering_resource_optionsAttrTypes = map[string]attr.Type{
	"order": types.ListType{ElemType: types.StringType},
}
var marketplaceoffering_resource_optionsObjectType = types.ObjectType{
	AttrTypes: marketplaceoffering_resource_optionsAttrTypes,
}

var marketplaceoffering_rolesAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"url":  types.StringType,
}
var marketplaceoffering_rolesObjectType = types.ObjectType{
	AttrTypes: marketplaceoffering_rolesAttrTypes,
}

var marketplaceoffering_screenshotsAttrTypes = map[string]attr.Type{
	"created":     types.StringType,
	"description": types.StringType,
	"image":       types.StringType,
	"name":        types.StringType,
	"thumbnail":   types.StringType,
}
var marketplaceoffering_screenshotsObjectType = types.ObjectType{
	AttrTypes: marketplaceoffering_screenshotsAttrTypes,
}

var marketplaceoffering_software_catalogsAttrTypes = map[string]attr.Type{
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
}
var marketplaceoffering_software_catalogsObjectType = types.ObjectType{
	AttrTypes: marketplaceoffering_software_catalogsAttrTypes,
}

var marketplaceofferingsoftwarecatalogs_catalogAttrTypes = map[string]attr.Type{
	"description": types.StringType,
	"name":        types.StringType,
	"version":     types.StringType,
}
var marketplaceofferingsoftwarecatalogs_catalogObjectType = types.ObjectType{
	AttrTypes: marketplaceofferingsoftwarecatalogs_catalogAttrTypes,
}

var marketplaceofferingsoftwarecatalogs_partitionAttrTypes = map[string]attr.Type{
	"partition_name": types.StringType,
	"priority_tier":  types.Int64Type,
	"qos":            types.StringType,
}
var marketplaceofferingsoftwarecatalogs_partitionObjectType = types.ObjectType{
	AttrTypes: marketplaceofferingsoftwarecatalogs_partitionAttrTypes,
}

// MarketplaceOfferingResourceModel describes the resource data model.
type MarketplaceOfferingResourceModel struct {
	UUID                      types.String   `tfsdk:"id"`
	AccessUrl                 types.String   `tfsdk:"access_url"`
	BackendId                 types.String   `tfsdk:"backend_id"`
	Billable                  types.Bool     `tfsdk:"billable"`
	BillingTypeClassification types.String   `tfsdk:"billing_type_classification"`
	Category                  types.String   `tfsdk:"category"`
	CategoryTitle             types.String   `tfsdk:"category_title"`
	CategoryUuid              types.String   `tfsdk:"category_uuid"`
	CitationCount             types.Int64    `tfsdk:"citation_count"`
	ComplianceChecklist       types.String   `tfsdk:"compliance_checklist"`
	Components                types.List     `tfsdk:"components"`
	Country                   types.String   `tfsdk:"country"`
	Created                   types.String   `tfsdk:"created"`
	Customer                  types.String   `tfsdk:"customer"`
	DataciteDoi               types.String   `tfsdk:"datacite_doi"`
	Description               types.String   `tfsdk:"description"`
	Endpoints                 types.List     `tfsdk:"endpoints"`
	Files                     types.List     `tfsdk:"files"`
	FullDescription           types.String   `tfsdk:"full_description"`
	GettingStarted            types.String   `tfsdk:"getting_started"`
	GoogleCalendarIsPublic    types.Bool     `tfsdk:"google_calendar_is_public"`
	GoogleCalendarLink        types.String   `tfsdk:"google_calendar_link"`
	HasComplianceRequirements types.Bool     `tfsdk:"has_compliance_requirements"`
	Image                     types.String   `tfsdk:"image"`
	IntegrationGuide          types.String   `tfsdk:"integration_guide"`
	IntegrationStatus         types.List     `tfsdk:"integration_status"`
	Latitude                  types.Float64  `tfsdk:"latitude"`
	Longitude                 types.Float64  `tfsdk:"longitude"`
	Name                      types.String   `tfsdk:"name"`
	Options                   types.Object   `tfsdk:"options"`
	OrderCount                types.Int64    `tfsdk:"order_count"`
	OrganizationGroups        types.List     `tfsdk:"organization_groups"`
	ParentDescription         types.String   `tfsdk:"parent_description"`
	ParentName                types.String   `tfsdk:"parent_name"`
	ParentUuid                types.String   `tfsdk:"parent_uuid"`
	Partitions                types.List     `tfsdk:"partitions"`
	PausedReason              types.String   `tfsdk:"paused_reason"`
	Plans                     types.List     `tfsdk:"plans"`
	PrivacyPolicyLink         types.String   `tfsdk:"privacy_policy_link"`
	Quotas                    types.List     `tfsdk:"quotas"`
	ResourceOptions           types.Object   `tfsdk:"resource_options"`
	Roles                     types.List     `tfsdk:"roles"`
	Scope                     types.String   `tfsdk:"scope"`
	ScopeErrorMessage         types.String   `tfsdk:"scope_error_message"`
	ScopeName                 types.String   `tfsdk:"scope_name"`
	ScopeState                types.String   `tfsdk:"scope_state"`
	ScopeUuid                 types.String   `tfsdk:"scope_uuid"`
	Screenshots               types.List     `tfsdk:"screenshots"`
	Shared                    types.Bool     `tfsdk:"shared"`
	Slug                      types.String   `tfsdk:"slug"`
	SoftwareCatalogs          types.List     `tfsdk:"software_catalogs"`
	State                     types.String   `tfsdk:"state"`
	Thumbnail                 types.String   `tfsdk:"thumbnail"`
	TotalCost                 types.Int64    `tfsdk:"total_cost"`
	TotalCostEstimated        types.Int64    `tfsdk:"total_cost_estimated"`
	TotalCustomers            types.Int64    `tfsdk:"total_customers"`
	Type                      types.String   `tfsdk:"type"`
	Url                       types.String   `tfsdk:"url"`
	VendorDetails             types.String   `tfsdk:"vendor_details"`
	Timeouts                  timeouts.Value `tfsdk:"timeouts"`
}

func (r *MarketplaceOfferingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_offering"
}

func (r *MarketplaceOfferingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Marketplace Offering resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"access_url": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Publicly accessible offering access URL",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"billable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Purchase and usage is invoiced.",
			},
			"billing_type_classification": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Classify offering components by billing type. Returns 'limit_only', 'usage_only', or 'mixed'.",
			},
			"category": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: " ",
			},
			"category_title": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"category_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"citation_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of citations of a DOI",
			},
			"compliance_checklist": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"components": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"article_code": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"billing_type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: " ",
						},
						"default_limit": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"is_boolean": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"is_prepaid": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"limit_amount": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"limit_period": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"max_available_limit": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"max_prepaid_duration": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"max_value": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"measured_unit": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Unit of measurement, for example, GB.",
						},
						"min_prepaid_duration": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"min_value": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Display name for the measured unit, for example, Floating IP.",
						},
						"overage_component": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Unique internal name of the measured unit, for example floating_ip.",
						},
						"unit_factor": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "The conversion factor from backend units to measured_unit",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"country": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Country code (ISO 3166-1 alpha-2)",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"customer": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"datacite_doi": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"endpoints": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"url": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "URL of the access endpoint",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"files": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"file": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"full_description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"getting_started": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"google_calendar_is_public": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"google_calendar_link": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Get the Google Calendar link for an offering.",
			},
			"has_compliance_requirements": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"image": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"integration_guide": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"integration_status": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"agent_type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"last_request_timestamp": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"service_name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"status": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"latitude": schema.Float64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"longitude": schema.Float64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: " ",
			},
			"options": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"order": schema.ListAttribute{
						CustomType:          types.ListType{ElemType: types.StringType},
						Required:            true,
						MarkdownDescription: " ",
					},
				},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"order_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
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
							Computed:            true,
							MarkdownDescription: " ",
						},
						"parent": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
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
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"parent_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"parent_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"parent_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"partitions": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cpu_bind": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Default task binding policy (SLURM cpu_bind)",
						},
						"def_cpu_per_gpu": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Default CPUs allocated per GPU",
						},
						"def_mem_per_cpu": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Default memory per CPU in MB",
						},
						"def_mem_per_gpu": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Default memory per GPU in MB",
						},
						"def_mem_per_node": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Default memory per node in MB",
						},
						"default_time": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Default time limit in minutes",
						},
						"exclusive_topo": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Exclusive topology access required",
						},
						"exclusive_user": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Exclusive user access required",
						},
						"grace_time": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Preemption grace time in seconds",
						},
						"max_cpus_per_node": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Maximum allocated CPUs per node",
						},
						"max_cpus_per_socket": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Maximum allocated CPUs per socket",
						},
						"max_mem_per_cpu": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Maximum memory per CPU in MB",
						},
						"max_mem_per_node": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Maximum memory per node in MB",
						},
						"max_nodes": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Maximum nodes per job",
						},
						"max_time": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Maximum time limit in minutes",
						},
						"min_nodes": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Minimum nodes per job",
						},
						"partition_name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Name of the SLURM partition",
						},
						"priority_tier": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Priority tier for scheduling and preemption",
						},
						"qos": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Quality of Service (QOS) name",
						},
						"req_resv": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Require reservation for job allocation",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"paused_reason": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"plans": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"archived": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Forbids creation of new resources.",
						},
						"article_code": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"backend_id": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"max_amount": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Maximum number of plans that could be active. Plan is disabled when maximum amount is reached.",
						},
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: " ",
						},
						"unit": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"unit_price": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"privacy_policy_link": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"quotas": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"limit": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"usage": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"resource_options": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"order": schema.ListAttribute{
						CustomType:          types.ListType{ElemType: types.StringType},
						Required:            true,
						MarkdownDescription: " ",
					},
				},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"roles": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"scope": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"scope_error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"scope_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"scope_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"scope_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"screenshots": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"image": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"thumbnail": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"shared": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Accessible to all customers.",
			},
			"slug": schema.StringAttribute{
				Optional:            true,
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
									Computed:            true,
									MarkdownDescription: " ",
								},
								"name": schema.StringAttribute{
									Optional:            true,
									Computed:            true,
									MarkdownDescription: " ",
								},
								"version": schema.StringAttribute{
									Optional:            true,
									Computed:            true,
									MarkdownDescription: " ",
								},
							},
							Computed:            true,
							MarkdownDescription: " ",
						},
						"package_count": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"partition": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"partition_name": schema.StringAttribute{
									Optional:            true,
									Computed:            true,
									MarkdownDescription: " ",
								},
								"priority_tier": schema.Int64Attribute{
									Optional:            true,
									Computed:            true,
									MarkdownDescription: " ",
								},
								"qos": schema.StringAttribute{
									Optional:            true,
									Computed:            true,
									MarkdownDescription: " ",
								},
							},
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"thumbnail": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"total_cost": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"total_cost_estimated": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"total_customers": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: " ",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"vendor_details": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
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

func (r *MarketplaceOfferingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MarketplaceOfferingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MarketplaceOfferingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to create resource
	var apiResp MarketplaceOfferingApiResponse // Prepare request body
	requestBody := map[string]interface{}{}
	if !data.AccessUrl.IsNull() && !data.AccessUrl.IsUnknown() {
		if v := data.AccessUrl.ValueString(); v != "" {
			requestBody["access_url"] = v
		}
	}
	if !data.BackendId.IsNull() && !data.BackendId.IsUnknown() {
		if v := data.BackendId.ValueString(); v != "" {
			requestBody["backend_id"] = v
		}
	}
	if !data.Billable.IsNull() && !data.Billable.IsUnknown() {
		requestBody["billable"] = data.Billable.ValueBool()
	}
	requestBody["category"] = data.Category.ValueString()
	if !data.ComplianceChecklist.IsNull() && !data.ComplianceChecklist.IsUnknown() {
		if v := data.ComplianceChecklist.ValueString(); v != "" {
			requestBody["compliance_checklist"] = v
		}
	}
	if !data.Components.IsNull() && !data.Components.IsUnknown() {
		if v := ConvertTFValue(data.Components); v != nil {
			requestBody["components"] = v
		}
	}
	if !data.Country.IsNull() && !data.Country.IsUnknown() {
		if v := data.Country.ValueString(); v != "" {
			requestBody["country"] = v
		}
	}
	if !data.Customer.IsNull() && !data.Customer.IsUnknown() {
		if v := data.Customer.ValueString(); v != "" {
			requestBody["customer"] = v
		}
	}
	if !data.DataciteDoi.IsNull() && !data.DataciteDoi.IsUnknown() {
		if v := data.DataciteDoi.ValueString(); v != "" {
			requestBody["datacite_doi"] = v
		}
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		if v := data.Description.ValueString(); v != "" {
			requestBody["description"] = v
		}
	}
	if !data.FullDescription.IsNull() && !data.FullDescription.IsUnknown() {
		if v := data.FullDescription.ValueString(); v != "" {
			requestBody["full_description"] = v
		}
	}
	if !data.GettingStarted.IsNull() && !data.GettingStarted.IsUnknown() {
		if v := data.GettingStarted.ValueString(); v != "" {
			requestBody["getting_started"] = v
		}
	}
	if !data.Image.IsNull() && !data.Image.IsUnknown() {
		if v := data.Image.ValueString(); v != "" {
			requestBody["image"] = v
		}
	}
	if !data.IntegrationGuide.IsNull() && !data.IntegrationGuide.IsUnknown() {
		if v := data.IntegrationGuide.ValueString(); v != "" {
			requestBody["integration_guide"] = v
		}
	}
	if !data.Latitude.IsNull() && !data.Latitude.IsUnknown() {
		requestBody["latitude"] = data.Latitude.ValueFloat64()
	}
	if !data.Longitude.IsNull() && !data.Longitude.IsUnknown() {
		requestBody["longitude"] = data.Longitude.ValueFloat64()
	}
	requestBody["name"] = data.Name.ValueString()
	if !data.Options.IsNull() && !data.Options.IsUnknown() {
		if v := ConvertTFValue(data.Options); v != nil {
			requestBody["options"] = v
		}
	}
	if !data.Plans.IsNull() && !data.Plans.IsUnknown() {
		if v := ConvertTFValue(data.Plans); v != nil {
			requestBody["plans"] = v
		}
	}
	if !data.PrivacyPolicyLink.IsNull() && !data.PrivacyPolicyLink.IsUnknown() {
		if v := data.PrivacyPolicyLink.ValueString(); v != "" {
			requestBody["privacy_policy_link"] = v
		}
	}
	if !data.ResourceOptions.IsNull() && !data.ResourceOptions.IsUnknown() {
		if v := ConvertTFValue(data.ResourceOptions); v != nil {
			requestBody["resource_options"] = v
		}
	}
	if !data.Shared.IsNull() && !data.Shared.IsUnknown() {
		requestBody["shared"] = data.Shared.ValueBool()
	}
	if !data.Slug.IsNull() && !data.Slug.IsUnknown() {
		if v := data.Slug.ValueString(); v != "" {
			requestBody["slug"] = v
		}
	}
	if !data.Thumbnail.IsNull() && !data.Thumbnail.IsUnknown() {
		if v := data.Thumbnail.ValueString(); v != "" {
			requestBody["thumbnail"] = v
		}
	}
	requestBody["type"] = data.Type.ValueString()
	if !data.VendorDetails.IsNull() && !data.VendorDetails.IsUnknown() {
		if v := data.VendorDetails.ValueString(); v != "" {
			requestBody["vendor_details"] = v
		}
	}
	err := r.client.Create(ctx, "/api/marketplace-provider-offerings/", requestBody, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Marketplace Offering",
			"An error occurred while creating the Marketplace Offering: "+err.Error(),
		)
		return
	}
	// Extract UUID from response
	if apiResp.UUID != nil {
		data.UUID = types.StringPointerValue(apiResp.UUID)
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MarketplaceOfferingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MarketplaceOfferingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource

	retrievePath := strings.Replace("/api/marketplace-provider-offerings/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	var apiResp MarketplaceOfferingApiResponse
	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Marketplace Offering",
			"An error occurred while reading the Marketplace Offering: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MarketplaceOfferingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update Not Supported", "This resource cannot be updated via the API.")
}

func (r *MarketplaceOfferingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MarketplaceOfferingResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to delete resource
	err := r.client.DeleteByUUID(ctx, "/api/marketplace-provider-offerings/{uuid}/", data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Marketplace Offering",
			"An error occurred while deleting the Marketplace Offering: "+err.Error(),
		)
		return
	}
}

func (r *MarketplaceOfferingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *MarketplaceOfferingResource) mapResponseToModel(ctx context.Context, apiResp MarketplaceOfferingApiResponse, model *MarketplaceOfferingResourceModel) diag.Diagnostics {
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
	listValComponents, listDiagsComponents := types.ListValueFrom(ctx, marketplaceoffering_componentsObjectType, apiResp.Components)
	diags.Append(listDiagsComponents...)
	model.Components = listValComponents
	model.Country = types.StringPointerValue(apiResp.Country)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Customer = types.StringPointerValue(apiResp.Customer)
	model.DataciteDoi = types.StringPointerValue(apiResp.DataciteDoi)
	model.Description = types.StringPointerValue(apiResp.Description)
	listValEndpoints, listDiagsEndpoints := types.ListValueFrom(ctx, marketplaceoffering_endpointsObjectType, apiResp.Endpoints)
	diags.Append(listDiagsEndpoints...)
	model.Endpoints = listValEndpoints
	listValFiles, listDiagsFiles := types.ListValueFrom(ctx, marketplaceoffering_filesObjectType, apiResp.Files)
	diags.Append(listDiagsFiles...)
	model.Files = listValFiles
	model.FullDescription = types.StringPointerValue(apiResp.FullDescription)
	model.GettingStarted = types.StringPointerValue(apiResp.GettingStarted)
	model.GoogleCalendarIsPublic = types.BoolPointerValue(apiResp.GoogleCalendarIsPublic)
	model.GoogleCalendarLink = types.StringPointerValue(apiResp.GoogleCalendarLink)
	model.HasComplianceRequirements = types.BoolPointerValue(apiResp.HasComplianceRequirements)
	model.Image = types.StringPointerValue(apiResp.Image)
	model.IntegrationGuide = types.StringPointerValue(apiResp.IntegrationGuide)
	listValIntegrationStatus, listDiagsIntegrationStatus := types.ListValueFrom(ctx, marketplaceoffering_integration_statusObjectType, apiResp.IntegrationStatus)
	diags.Append(listDiagsIntegrationStatus...)
	model.IntegrationStatus = listValIntegrationStatus
	model.Latitude = types.Float64PointerValue(apiResp.Latitude)
	model.Longitude = types.Float64PointerValue(apiResp.Longitude)
	if apiResp.Options != nil {
		objValOptions, objDiagsOptions := types.ObjectValueFrom(ctx, marketplaceoffering_optionsAttrTypes, *apiResp.Options)
		diags.Append(objDiagsOptions...)
		model.Options = objValOptions
	} else {
		model.Options = types.ObjectNull(marketplaceoffering_optionsAttrTypes)
	}
	model.OrderCount = types.Int64PointerValue(apiResp.OrderCount)
	listValOrganizationGroups, listDiagsOrganizationGroups := types.ListValueFrom(ctx, marketplaceoffering_organization_groupsObjectType, apiResp.OrganizationGroups)
	diags.Append(listDiagsOrganizationGroups...)
	model.OrganizationGroups = listValOrganizationGroups
	model.ParentDescription = types.StringPointerValue(apiResp.ParentDescription)
	model.ParentName = types.StringPointerValue(apiResp.ParentName)
	model.ParentUuid = types.StringPointerValue(apiResp.ParentUuid)
	listValPartitions, listDiagsPartitions := types.ListValueFrom(ctx, marketplaceoffering_partitionsObjectType, apiResp.Partitions)
	diags.Append(listDiagsPartitions...)
	model.Partitions = listValPartitions
	model.PausedReason = types.StringPointerValue(apiResp.PausedReason)
	listValPlans, listDiagsPlans := types.ListValueFrom(ctx, marketplaceoffering_plansObjectType, apiResp.Plans)
	diags.Append(listDiagsPlans...)
	model.Plans = listValPlans
	model.PrivacyPolicyLink = types.StringPointerValue(apiResp.PrivacyPolicyLink)
	listValQuotas, listDiagsQuotas := types.ListValueFrom(ctx, marketplaceoffering_quotasObjectType, apiResp.Quotas)
	diags.Append(listDiagsQuotas...)
	model.Quotas = listValQuotas
	if apiResp.ResourceOptions != nil {
		objValResourceOptions, objDiagsResourceOptions := types.ObjectValueFrom(ctx, marketplaceoffering_resource_optionsAttrTypes, *apiResp.ResourceOptions)
		diags.Append(objDiagsResourceOptions...)
		model.ResourceOptions = objValResourceOptions
	} else {
		model.ResourceOptions = types.ObjectNull(marketplaceoffering_resource_optionsAttrTypes)
	}
	listValRoles, listDiagsRoles := types.ListValueFrom(ctx, marketplaceoffering_rolesObjectType, apiResp.Roles)
	diags.Append(listDiagsRoles...)
	model.Roles = listValRoles
	model.Scope = types.StringPointerValue(apiResp.Scope)
	model.ScopeErrorMessage = types.StringPointerValue(apiResp.ScopeErrorMessage)
	model.ScopeName = types.StringPointerValue(apiResp.ScopeName)
	model.ScopeState = types.StringPointerValue(apiResp.ScopeState)
	model.ScopeUuid = types.StringPointerValue(apiResp.ScopeUuid)
	listValScreenshots, listDiagsScreenshots := types.ListValueFrom(ctx, marketplaceoffering_screenshotsObjectType, apiResp.Screenshots)
	diags.Append(listDiagsScreenshots...)
	model.Screenshots = listValScreenshots
	model.Shared = types.BoolPointerValue(apiResp.Shared)
	model.Slug = types.StringPointerValue(apiResp.Slug)
	listValSoftwareCatalogs, listDiagsSoftwareCatalogs := types.ListValueFrom(ctx, marketplaceoffering_software_catalogsObjectType, apiResp.SoftwareCatalogs)
	diags.Append(listDiagsSoftwareCatalogs...)
	model.SoftwareCatalogs = listValSoftwareCatalogs
	model.State = types.StringPointerValue(apiResp.State)
	model.Thumbnail = types.StringPointerValue(apiResp.Thumbnail)
	model.TotalCost = types.Int64PointerValue(apiResp.TotalCost)
	model.TotalCostEstimated = types.Int64PointerValue(apiResp.TotalCostEstimated)
	model.TotalCustomers = types.Int64PointerValue(apiResp.TotalCustomers)
	model.Type = types.StringPointerValue(apiResp.Type)
	model.Url = types.StringPointerValue(apiResp.Url)
	model.VendorDetails = types.StringPointerValue(apiResp.VendorDetails)

	return diags
}
