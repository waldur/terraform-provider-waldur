package offering

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

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
				MarkdownDescription: "Marketplace Offering UUID",
			},
			"filters": (&MarketplaceOfferingFiltersModel{}).GetSchema(),
			"access_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Publicly accessible offering access URL",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
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
				Computed:            true,
				MarkdownDescription: " ",
			},
			"components": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"article_code": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"billing_type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
							Validators: []validator.String{
								stringvalidator.OneOf("fixed", "usage", "limit", "one", "few"),
							},
						},
						"default_limit": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"factor": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"is_boolean": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"is_builtin": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"is_prepaid": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"limit_amount": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"limit_period": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"max_available_limit": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"max_prepaid_duration": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"max_value": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"measured_unit": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Unit of measurement, for example, GB.",
						},
						"min_prepaid_duration": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"min_value": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Display name for the measured unit, for example, Floating IP.",
						},
						"overage_component": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Unique internal name of the measured unit, for example floating_ip.",
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_\-\/:]+$`), ""),
							},
						},
						"unit_factor": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "The conversion factor from backend units to measured_unit",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"country": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Country code (ISO 3166-1 alpha-2)",
			},
			"created": schema.StringAttribute{
				CustomType:          timetypes.RFC3339Type{},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"customer": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"customer_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"customer_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"datacite_doi": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"endpoints": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "URL of the access endpoint",
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
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
							CustomType:          timetypes.RFC3339Type{},
							Computed:            true,
							MarkdownDescription: " ",
						},
						"file": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"full_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"getting_started": schema.StringAttribute{
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
				Computed:            true,
				MarkdownDescription: " ",
			},
			"integration_guide": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"is_accessible": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"latitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"longitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"name": schema.StringAttribute{
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
							Computed:            true,
							MarkdownDescription: " ",
						},
						"parent": schema.StringAttribute{
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
						"uuid": schema.StringAttribute{
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
							Computed:            true,
							MarkdownDescription: "Default task binding policy (SLURM cpu_bind)",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"def_cpu_per_gpu": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Default CPUs allocated per GPU",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"def_mem_per_cpu": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Default memory per CPU in MB",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"def_mem_per_gpu": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Default memory per GPU in MB",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"def_mem_per_node": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Default memory per node in MB",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"default_time": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Default time limit in minutes",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"exclusive_topo": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Exclusive topology access required",
						},
						"exclusive_user": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Exclusive user access required",
						},
						"grace_time": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Preemption grace time in seconds",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"max_cpus_per_node": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Maximum allocated CPUs per node",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"max_cpus_per_socket": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Maximum allocated CPUs per socket",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"max_mem_per_cpu": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Maximum memory per CPU in MB",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"max_mem_per_node": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Maximum memory per node in MB",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"max_nodes": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Maximum nodes per job",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"max_time": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Maximum time limit in minutes",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"min_nodes": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Minimum nodes per job",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"partition_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name of the SLURM partition",
						},
						"priority_tier": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Priority tier for scheduling and preemption",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(32767),
							},
						},
						"qos": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Quality of Service (QOS) name",
						},
						"req_resv": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Require reservation for job allocation",
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
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
							Computed:            true,
							MarkdownDescription: "Forbids creation of new resources.",
						},
						"article_code": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"backend_id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"components": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"amount": schema.Int64Attribute{
										Computed:            true,
										MarkdownDescription: " ",
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(2147483647),
										},
									},
									"discount_rate": schema.Int64Attribute{
										Computed:            true,
										MarkdownDescription: "Discount rate in percentage.",
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(2147483647),
										},
									},
									"discount_threshold": schema.Int64Attribute{
										Computed:            true,
										MarkdownDescription: "Minimum amount to be eligible for discount.",
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(2147483647),
										},
									},
									"future_price": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: " ",
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,12}(?:\.\d{0,10})?$`), ""),
										},
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
										Computed:            true,
										MarkdownDescription: " ",
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,12}(?:\.\d{0,10})?$`), ""),
										},
									},
									"type": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Unique internal name of the measured unit, for example floating_ip.",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: " ",
						},
						"description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"init_price": schema.Float64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"is_active": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"max_amount": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Maximum number of plans that could be active. Plan is disabled when maximum amount is reached.",
							Validators: []validator.Int64{
								int64validator.AtLeast(1),
								int64validator.AtMost(32767),
							},
						},
						"minimal_price": schema.Float64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"name": schema.StringAttribute{
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
										Computed:            true,
										MarkdownDescription: " ",
									},
									"parent": schema.StringAttribute{
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
									"uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: " ",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: " ",
						},
						"plan_type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"resources_count": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"switch_price": schema.Float64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"unit": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
							Validators: []validator.String{
								stringvalidator.OneOf("month", "quarter", "half_month", "day", "hour", "quantity"),
							},
						},
						"unit_price": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,12}(?:\.\d{0,10})?$`), ""),
							},
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"privacy_policy_link": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"project": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"project_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"project_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"promotion_campaigns": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"discount": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"discount_type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
							Validators: []validator.String{
								stringvalidator.OneOf("discount", "special_price"),
							},
						},
						"end_date": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The last day the campaign is active.",
						},
						"months": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "How many months in a row should the related service (when activated) get special deal (0 for indefinitely until active)",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"service_provider": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"start_date": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Starting from this date, the campaign is active.",
						},
						"stock": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"quotas": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"limit": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"usage": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"roles": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"uuid": schema.StringAttribute{
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
							CustomType:          timetypes.RFC3339Type{},
							Computed:            true,
							MarkdownDescription: " ",
						},
						"description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"image": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"thumbnail": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"shared": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Accessible to all customers.",
			},
			"slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
				},
			},
			"software_catalogs": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"catalog": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Computed:            true,
									MarkdownDescription: " ",
								},
								"name": schema.StringAttribute{
									Computed:            true,
									MarkdownDescription: " ",
								},
								"uuid": schema.StringAttribute{
									Computed:            true,
									MarkdownDescription: " ",
								},
								"version": schema.StringAttribute{
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
									Computed:            true,
									MarkdownDescription: " ",
								},
								"priority_tier": schema.Int64Attribute{
									Computed:            true,
									MarkdownDescription: " ",
								},
								"qos": schema.StringAttribute{
									Computed:            true,
									MarkdownDescription: " ",
								},
								"uuid": schema.StringAttribute{
									Computed:            true,
									MarkdownDescription: " ",
								},
							},
							Computed:            true,
							MarkdownDescription: " ",
						},
						"uuid": schema.StringAttribute{
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
			"tags": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"thumbnail": schema.StringAttribute{
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
				Computed:            true,
				MarkdownDescription: " ",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"user_has_consent": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"vendor_details": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
		},
	}
}

func (d *MarketplaceOfferingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	d.client = &Client{}
	if err := d.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			err.Error(),
		)
		return
	}
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
