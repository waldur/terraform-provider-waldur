package offering

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &MarketplaceOfferingDataSource{}

func NewMarketplaceOfferingDataSource() datasource.DataSource {
	return &MarketplaceOfferingDataSource{}
}

type MarketplaceOfferingDataSource struct {
	client *Client
}

type MarketplaceOfferingDataSourceModel struct {
	MarketplaceOfferingModel
	Filters *MarketplaceOfferingFiltersModel `tfsdk:"filters"`
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
							CustomType:          timetypes.RFC3339Type{},
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
							CustomType:          timetypes.RFC3339Type{},
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

		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

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

		resp.Diagnostics.Append(data.CopyFrom(ctx, results[0])...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
