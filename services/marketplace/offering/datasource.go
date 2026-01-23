package offering

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &MarketplaceOfferingDataSource{}

func NewMarketplaceOfferingDataSource() datasource.DataSource {
	return &MarketplaceOfferingDataSource{}
}

// MarketplaceOfferingDataSource defines the data source implementation.
type MarketplaceOfferingDataSource struct {
	client *Client
}

// MarketplaceOfferingFiltersModel contains the filter parameters for querying.
type MarketplaceOfferingFiltersModel struct {
	AccessibleViaCalls      types.Bool   `tfsdk:"accessible_via_calls"`
	AllowedCustomerUuid     types.String `tfsdk:"allowed_customer_uuid"`
	Attributes              types.String `tfsdk:"attributes"`
	Billable                types.Bool   `tfsdk:"billable"`
	CanCreateOfferingUser   types.Bool   `tfsdk:"can_create_offering_user"`
	CategoryGroupUuid       types.String `tfsdk:"category_group_uuid"`
	CategoryUuid            types.String `tfsdk:"category_uuid"`
	Created                 types.String `tfsdk:"created"`
	Customer                types.String `tfsdk:"customer"`
	CustomerUuid            types.String `tfsdk:"customer_uuid"`
	Description             types.String `tfsdk:"description"`
	HasActiveTermsOfService types.Bool   `tfsdk:"has_active_terms_of_service"`
	HasTermsOfService       types.Bool   `tfsdk:"has_terms_of_service"`
	Keyword                 types.String `tfsdk:"keyword"`
	Modified                types.String `tfsdk:"modified"`
	Name                    types.String `tfsdk:"name"`
	NameExact               types.String `tfsdk:"name_exact"`
	OrganizationGroupUuid   types.String `tfsdk:"organization_group_uuid"`
	ParentUuid              types.String `tfsdk:"parent_uuid"`
	ProjectUuid             types.String `tfsdk:"project_uuid"`
	Query                   types.String `tfsdk:"query"`
	ResourceCustomerUuid    types.String `tfsdk:"resource_customer_uuid"`
	ResourceProjectUuid     types.String `tfsdk:"resource_project_uuid"`
	ScopeUuid               types.String `tfsdk:"scope_uuid"`
	ServiceManagerUuid      types.String `tfsdk:"service_manager_uuid"`
	Shared                  types.Bool   `tfsdk:"shared"`
	State                   types.String `tfsdk:"state"`
	Type                    types.String `tfsdk:"type"`
	UserHasConsent          types.Bool   `tfsdk:"user_has_consent"`
	UserHasOfferingUser     types.Bool   `tfsdk:"user_has_offering_user"`
	UuidList                types.String `tfsdk:"uuid_list"`
}

type MarketplaceOfferingDataSourceModel struct {
	UUID                      types.String                     `tfsdk:"id"`
	Filters                   *MarketplaceOfferingFiltersModel `tfsdk:"filters"`
	AccessUrl                 types.String                     `tfsdk:"access_url"`
	BackendId                 types.String                     `tfsdk:"backend_id"`
	Billable                  types.Bool                       `tfsdk:"billable"`
	BillingTypeClassification types.String                     `tfsdk:"billing_type_classification"`
	Category                  types.String                     `tfsdk:"category"`
	CategoryTitle             types.String                     `tfsdk:"category_title"`
	CategoryUuid              types.String                     `tfsdk:"category_uuid"`
	CitationCount             types.Int64                      `tfsdk:"citation_count"`
	ComplianceChecklist       types.String                     `tfsdk:"compliance_checklist"`
	Components                types.List                       `tfsdk:"components"`
	Country                   types.String                     `tfsdk:"country"`
	Created                   types.String                     `tfsdk:"created"`
	DataciteDoi               types.String                     `tfsdk:"datacite_doi"`
	Description               types.String                     `tfsdk:"description"`
	Endpoints                 types.List                       `tfsdk:"endpoints"`
	Files                     types.List                       `tfsdk:"files"`
	FullDescription           types.String                     `tfsdk:"full_description"`
	GettingStarted            types.String                     `tfsdk:"getting_started"`
	GoogleCalendarIsPublic    types.Bool                       `tfsdk:"google_calendar_is_public"`
	GoogleCalendarLink        types.String                     `tfsdk:"google_calendar_link"`
	HasComplianceRequirements types.Bool                       `tfsdk:"has_compliance_requirements"`
	Image                     types.String                     `tfsdk:"image"`
	IntegrationGuide          types.String                     `tfsdk:"integration_guide"`
	Latitude                  types.Float64                    `tfsdk:"latitude"`
	Longitude                 types.Float64                    `tfsdk:"longitude"`
	Name                      types.String                     `tfsdk:"name"`
	OrderCount                types.Int64                      `tfsdk:"order_count"`
	OrganizationGroups        types.List                       `tfsdk:"organization_groups"`
	ParentDescription         types.String                     `tfsdk:"parent_description"`
	ParentName                types.String                     `tfsdk:"parent_name"`
	ParentUuid                types.String                     `tfsdk:"parent_uuid"`
	Partitions                types.List                       `tfsdk:"partitions"`
	PausedReason              types.String                     `tfsdk:"paused_reason"`
	Plans                     types.List                       `tfsdk:"plans"`
	PrivacyPolicyLink         types.String                     `tfsdk:"privacy_policy_link"`
	PromotionCampaigns        types.List                       `tfsdk:"promotion_campaigns"`
	Quotas                    types.List                       `tfsdk:"quotas"`
	Roles                     types.List                       `tfsdk:"roles"`
	Scope                     types.String                     `tfsdk:"scope"`
	ScopeErrorMessage         types.String                     `tfsdk:"scope_error_message"`
	ScopeName                 types.String                     `tfsdk:"scope_name"`
	ScopeState                types.String                     `tfsdk:"scope_state"`
	ScopeUuid                 types.String                     `tfsdk:"scope_uuid"`
	Screenshots               types.List                       `tfsdk:"screenshots"`
	Shared                    types.Bool                       `tfsdk:"shared"`
	Slug                      types.String                     `tfsdk:"slug"`
	SoftwareCatalogs          types.List                       `tfsdk:"software_catalogs"`
	State                     types.String                     `tfsdk:"state"`
	Thumbnail                 types.String                     `tfsdk:"thumbnail"`
	TotalCost                 types.Int64                      `tfsdk:"total_cost"`
	TotalCostEstimated        types.Int64                      `tfsdk:"total_cost_estimated"`
	TotalCustomers            types.Int64                      `tfsdk:"total_customers"`
	Type                      types.String                     `tfsdk:"type"`
	Url                       types.String                     `tfsdk:"url"`
	UserHasConsent            types.Bool                       `tfsdk:"user_has_consent"`
	VendorDetails             types.String                     `tfsdk:"vendor_details"`
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
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Filter parameters for querying Marketplace Offering",
				Attributes: map[string]schema.Attribute{
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
						MarkdownDescription: "Offering state Allowed values: `Active`, `Archived`, `Draft`, `Paused`, `Unavailable`.",
						Validators: []validator.String{
							stringvalidator.OneOf("Active", "Archived", "Draft", "Paused", "Unavailable"),
						},
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
				},
			},
			"access_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Publicly accessible offering access URL",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of the backend",
			},
			"billable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Purchase and usage is invoiced.",
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
			"category_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the category",
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
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Created",
			},
			"datacite_doi": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Datacite doi",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Description of the resource",
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
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the resource",
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
			"parent_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the parent",
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
			"scope_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the scope",
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
			"shared": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Accessible to all customers.",
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
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "State",
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
			"type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Type",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Url",
			},
			"user_has_consent": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "User has consent",
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

	rawClient, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = NewClient(rawClient)
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
		apiResp, err := d.client.GetMarketplaceOffering(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Marketplace Offering",
				"An error occurred while reading the Marketplace Offering by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, *apiResp, &data)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup marketplace_offering.",
			)
			return
		}

		results, err := d.client.ListMarketplaceOffering(ctx, filters)
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

func (d *MarketplaceOfferingDataSource) mapResponseToModel(ctx context.Context, apiResp MarketplaceOfferingResponse, model *MarketplaceOfferingDataSourceModel) diag.Diagnostics {
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

	{
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
	}
	model.Country = types.StringPointerValue(apiResp.Country)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.DataciteDoi = types.StringPointerValue(apiResp.DataciteDoi)
	model.Description = types.StringPointerValue(apiResp.Description)

	{
		listValEndpoints, listDiagsEndpoints := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"url":  types.StringType,
		}}, apiResp.Endpoints)
		diags.Append(listDiagsEndpoints...)
		model.Endpoints = listValEndpoints
	}

	{
		listValFiles, listDiagsFiles := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"created": types.StringType,
			"file":    types.StringType,
			"name":    types.StringType,
		}}, apiResp.Files)
		diags.Append(listDiagsFiles...)
		model.Files = listValFiles
	}
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

	{
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
	}
	model.ParentDescription = types.StringPointerValue(apiResp.ParentDescription)
	model.ParentName = types.StringPointerValue(apiResp.ParentName)
	model.ParentUuid = types.StringPointerValue(apiResp.ParentUuid)

	{
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
	}
	model.PausedReason = types.StringPointerValue(apiResp.PausedReason)

	{
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
	}
	model.PrivacyPolicyLink = types.StringPointerValue(apiResp.PrivacyPolicyLink)

	{
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
	}

	{
		listValQuotas, listDiagsQuotas := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"limit": types.Int64Type,
			"name":  types.StringType,
			"usage": types.Int64Type,
		}}, apiResp.Quotas)
		diags.Append(listDiagsQuotas...)
		model.Quotas = listValQuotas
	}

	{
		listValRoles, listDiagsRoles := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"url":  types.StringType,
		}}, apiResp.Roles)
		diags.Append(listDiagsRoles...)
		model.Roles = listValRoles
	}
	model.Scope = types.StringPointerValue(apiResp.Scope)
	model.ScopeErrorMessage = types.StringPointerValue(apiResp.ScopeErrorMessage)
	model.ScopeName = types.StringPointerValue(apiResp.ScopeName)
	model.ScopeState = types.StringPointerValue(apiResp.ScopeState)
	model.ScopeUuid = types.StringPointerValue(apiResp.ScopeUuid)

	{
		listValScreenshots, listDiagsScreenshots := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"created":     types.StringType,
			"description": types.StringType,
			"image":       types.StringType,
			"name":        types.StringType,
			"thumbnail":   types.StringType,
		}}, apiResp.Screenshots)
		diags.Append(listDiagsScreenshots...)
		model.Screenshots = listValScreenshots
	}
	model.Shared = types.BoolPointerValue(apiResp.Shared)
	model.Slug = types.StringPointerValue(apiResp.Slug)

	{
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
	}
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
