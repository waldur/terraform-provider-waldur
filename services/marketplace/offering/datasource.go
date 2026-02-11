package offering

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &MarketplaceOfferingDataSource{}

func NewMarketplaceOfferingDataSource() datasource.DataSource {
	return &MarketplaceOfferingDataSource{}
}

type MarketplaceOfferingDataSource struct {
	client *MarketplaceOfferingClient
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
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Backend Id",
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
				MarkdownDescription: "Category Title",
			},
			"category_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Category Uuid",
			},
			"citation_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of citations of a DOI",
			},
			"compliance_checklist": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Compliance Checklist",
			},
			"components": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"article_code": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Article Code",
						},
						"billing_type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Billing Type",

							Validators: []validator.String{
								stringvalidator.OneOf("fixed", "usage", "limit", "one", "few"),
							},
						},
						"default_limit": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Default Limit",

							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Description",
						},
						"factor": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Factor",
						},
						"is_boolean": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Is Boolean",
						},
						"is_builtin": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Is Builtin",
						},
						"is_prepaid": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Is Prepaid",
						},
						"limit_amount": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Limit Amount",

							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"limit_period": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Limit Period",
						},
						"max_available_limit": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Max Available Limit",

							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"max_prepaid_duration": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Max Prepaid Duration",

							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"max_value": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Max Value",

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
							MarkdownDescription: "Min Prepaid Duration",

							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"min_value": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Min Value",

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
							MarkdownDescription: "Overage Component",
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
							MarkdownDescription: "Uuid",
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
			"customer": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Customer",
			},
			"datacite_doi": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Datacite Doi",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Description",
			},
			"endpoints": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "URL of the access endpoint",
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Uuid",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Endpoints",
			},
			"files": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"file": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "File",
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Files",
			},
			"full_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Full Description",
			},
			"getting_started": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Getting Started",
			},
			"google_calendar_is_public": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Google Calendar Is Public",
			},
			"google_calendar_link": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Get the Google Calendar link for an offering.",
			},
			"has_compliance_requirements": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Has Compliance Requirements",
			},
			"image": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Image",
			},
			"integration_guide": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Integration Guide",
			},
			"is_accessible": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Is Accessible",
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
				MarkdownDescription: "Name",
			},
			"options": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"order": schema.ListAttribute{
						ElementType:         types.StringType,
						Computed:            true,
						MarkdownDescription: "Order",
					},
				},
				Computed:            true,
				MarkdownDescription: "Options",
			},
			"order_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Order Count",
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
							MarkdownDescription: "Name",
						},
						"parent": schema.StringAttribute{
							Computed:            true,
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
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Uuid",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Organization Groups",
			},
			"parent_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Parent Description",
			},
			"parent_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Parent Name",
			},
			"parent_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Parent Uuid",
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
							MarkdownDescription: "Uuid",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Partitions",
			},
			"paused_reason": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Paused Reason",
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
							MarkdownDescription: "Article Code",
						},
						"backend_id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Backend Id",
						},
						"components": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"amount": schema.Int64Attribute{
										Computed:            true,
										MarkdownDescription: "Amount",

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
										MarkdownDescription: "Future Price",

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
										MarkdownDescription: "Price",

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
							MarkdownDescription: "Components",
						},
						"description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Description",
						},
						"init_price": schema.Float64Attribute{
							Computed:            true,
							MarkdownDescription: "Init Price",
						},
						"is_active": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Is Active",
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
							MarkdownDescription: "Minimal Price",
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name",
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
										MarkdownDescription: "Name",
									},
									"parent": schema.StringAttribute{
										Computed:            true,
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
									"uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Uuid",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: "Organization Groups",
						},
						"plan_type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Plan Type",
						},
						"resources_count": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Resources Count",
						},
						"switch_price": schema.Float64Attribute{
							Computed:            true,
							MarkdownDescription: "Switch Price",
						},
						"unit": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Unit",

							Validators: []validator.String{
								stringvalidator.OneOf("month", "quarter", "half_month", "day", "hour", "quantity"),
							},
						},
						"unit_price": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Unit Price",

							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,12}(?:\.\d{0,10})?$`), ""),
							},
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Url",
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Uuid",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Plans",
			},
			"plugin_options": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"auto_approve_in_service_provider_projects": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "Skip approval of public offering belonging to the same organization under which the request is done",
					},
					"auto_approve_marketplace_script": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "If set to False, all orders require manual provider approval, including for service provider owners and staff",
					},
					"auto_approve_remote_orders": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "If set to True, an order can be processed without approval",
					},
					"backend_id_display_label": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Label used by UI for showing value of the backend_id",
					},
					"can_restore_resource": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "If set to True, resource can be restored.",
					},
					"conceal_billing_data": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "If set to True, pricing and components tab would be concealed.",
					},
					"create_orders_on_resource_option_change": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "If set to True, create orders when options of related resources are changed.",
					},
					"default_internal_network_mtu": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "If set, it will be used as a default MTU for the first network in a tenant",

						Validators: []validator.Int64{
							int64validator.AtLeast(68),
							int64validator.AtMost(9000),
						},
					},
					"default_resource_termination_offset_in_days": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "If set, it will be used as a default resource termination offset in days",

						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"deployment_mode": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Rancher deployment mode",
					},
					"disable_autoapprove": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "If set to True, orders for this offering will always require manual approval, overriding auto_approve_in_service_provider_projects",
					},
					"enable_display_of_order_actions_for_service_provider": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "Enable display of order actions for service provider",
					},
					"enable_issues_for_membership_changes": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "Enable issues for membership changes",
					},
					"enable_purchase_order_upload": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "If set to True, users will be able to upload purchase orders.",
					},
					"flavors_regex": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Regular expression to limit flavors list",
					},
					"heappe_cluster_id": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "HEAppE cluster id",
					},
					"heappe_local_base_path": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "HEAppE local base path",
					},
					"heappe_url": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "HEAppE url",
					},
					"heappe_username": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "HEAppE username",
					},
					"highlight_backend_id_display": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "Defines if backend_id should be shown more prominently by the UI",
					},
					"homedir_prefix": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "GLAuth homedir prefix",
					},
					"initial_primarygroup_number": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "GLAuth initial primary group number",

						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"initial_uidnumber": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "GLAuth initial uidnumber",

						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"initial_usergroup_number": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "GLAuth initial usergroup number",

						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"is_resource_termination_date_required": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "If set to True, resource termination date is required",
					},
					"latest_date_for_resource_termination": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "If set, it will be used as a latest date for resource termination",
					},
					"managed_rancher_load_balancer_data_volume_size_gb": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "Data volume size in GB for managed Rancher load balancer",
					},
					"managed_rancher_load_balancer_data_volume_type_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Data volume type name for managed Rancher load balancer",
					},
					"managed_rancher_load_balancer_flavor_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Flavor name for managed Rancher load balancer",
					},
					"managed_rancher_load_balancer_system_volume_size_gb": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "System volume size in GB for managed Rancher load balancer",
					},
					"managed_rancher_load_balancer_system_volume_type_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "System volume type name for managed Rancher load balancer",
					},
					"managed_rancher_server_data_volume_size_gb": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "Data volume size in GB for managed Rancher server",
					},
					"managed_rancher_server_data_volume_type_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Data volume type name for managed Rancher server",
					},
					"managed_rancher_server_flavor_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Flavor name for managed Rancher server instances",
					},
					"managed_rancher_server_system_volume_size_gb": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "System volume size in GB for managed Rancher server",
					},
					"managed_rancher_server_system_volume_type_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "System volume type name for managed Rancher server",
					},
					"managed_rancher_tenant_max_cpu": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "Max number of vCPUs for tenants",
					},
					"managed_rancher_tenant_max_disk": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "Max size of disk space for tenants (GB)",
					},
					"managed_rancher_tenant_max_ram": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "Max number of RAM for tenants (GB)",
					},
					"managed_rancher_worker_system_volume_size_gb": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "System volume size in GB for managed Rancher worker nodes",
					},
					"managed_rancher_worker_system_volume_type_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "System volume type name for managed Rancher worker nodes",
					},
					"max_instances": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "Default limit for number of instances in OpenStack tenant",

						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
					"max_resource_termination_offset_in_days": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "Maximum resource termination offset in days",

						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"max_security_groups": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "Default limit for number of security groups in OpenStack tenant",

						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
					"max_volumes": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "Default limit for number of volumes in OpenStack tenant",

						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
					"maximal_resource_count_per_project": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "Maximal number of offering resources allowed per project",
					},
					"minimal_team_count_for_provisioning": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "Minimal team count required for provisioning of resources",

						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
					"openstack_offering_uuid_list": schema.ListAttribute{
						ElementType:         types.StringType,
						Computed:            true,
						MarkdownDescription: "List of UUID of OpenStack offerings where tenant can be created",
					},
					"project_permanent_directory": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "HEAppE project permanent directory",
					},
					"require_purchase_order_upload": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "If set to True, users will be required to upload purchase orders.",
					},
					"required_team_role_for_provisioning": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required user role in a project for provisioning of resources",
					},
					"resource_expiration_threshold": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "Resource expiration threshold in days.",
					},
					"scratch_project_directory": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "HEAppE scratch project directory",
					},
					"service_provider_can_create_offering_user": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "Service provider can create offering user",
					},
					"slurm_periodic_policy_enabled": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "Enable SLURM periodic usage policy configuration. When enabled, allows configuring QoS-based threshold enforcement, carryover logic, and fairshare decay for site-agent managed SLURM offerings.",
					},
					"snapshot_size_limit_gb": schema.Int64Attribute{
						Computed:            true,
						MarkdownDescription: "Default limit for snapshot size in GB",

						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
					"storage_mode": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Storage mode for OpenStack offering",
					},
					"supports_downscaling": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "If set to True, it will be possible to downscale resources",
					},
					"supports_pausing": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "If set to True, it will be possible to pause resources",
					},
					"unique_resource_per_attribute": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Attribute name to enforce uniqueness per value. E.g., 'storage_data_type' ensures only one resource per storage type per project.",
					},
					"username_anonymized_prefix": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "GLAuth prefix for anonymized usernames",
					},
					"username_generation_policy": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "GLAuth username generation policy",
					},
				},
				Computed:            true,
				MarkdownDescription: "Plugin Options",
			},
			"privacy_policy_link": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Privacy Policy Link",
			},
			"project": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project",
			},
			"promotion_campaigns": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Description",
						},
						"discount": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Discount",

							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"discount_type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Discount Type",

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
							MarkdownDescription: "Name",
						},
						"service_provider": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Service Provider",
						},
						"start_date": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Starting from this date, the campaign is active.",
						},
						"stock": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Stock",

							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Uuid",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Promotion Campaigns",
			},
			"quotas": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"limit": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Limit",
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name",
						},
						"usage": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Usage",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Quotas",
			},
			"resource_options": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"order": schema.ListAttribute{
						ElementType:         types.StringType,
						Computed:            true,
						MarkdownDescription: "Order",
					},
				},
				Computed:            true,
				MarkdownDescription: "Resource Options",
			},
			"roles": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Url",
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Uuid",
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
				MarkdownDescription: "Scope Error Message",
			},
			"scope_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Scope Name",
			},
			"scope_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Scope State",
			},
			"scope_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Scope Uuid",
			},
			"screenshots": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Description",
						},
						"image": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Image",
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name",
						},
						"thumbnail": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Thumbnail",
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Uuid",
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
									MarkdownDescription: "Description",
								},
								"name": schema.StringAttribute{
									Computed:            true,
									MarkdownDescription: "Name",
								},
								"uuid": schema.StringAttribute{
									Computed:            true,
									MarkdownDescription: "Uuid",
								},
								"version": schema.StringAttribute{
									Computed:            true,
									MarkdownDescription: "Version",
								},
							},
							Computed:            true,
							MarkdownDescription: "Catalog",
						},
						"package_count": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Package Count",
						},
						"partition": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"partition_name": schema.StringAttribute{
									Computed:            true,
									MarkdownDescription: "Partition Name",
								},
								"priority_tier": schema.Int64Attribute{
									Computed:            true,
									MarkdownDescription: "Priority Tier",
								},
								"qos": schema.StringAttribute{
									Computed:            true,
									MarkdownDescription: "Qos",
								},
								"uuid": schema.StringAttribute{
									Computed:            true,
									MarkdownDescription: "Uuid",
								},
							},
							Computed:            true,
							MarkdownDescription: "Partition",
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Uuid",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Software Catalogs",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "State",
			},
			"tags": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name",
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Uuid",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Tags",
			},
			"thumbnail": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Thumbnail",
			},
			"total_cost": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Total Cost",
			},
			"total_cost_estimated": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Total Cost Estimated",
			},
			"total_customers": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Total Customers",
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
				MarkdownDescription: "User Has Consent",
			},
			"vendor_details": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Vendor Details",
			},
		},
	}
}

func (d *MarketplaceOfferingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	d.client = &MarketplaceOfferingClient{}
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
		apiResp, err := d.client.Get(ctx, data.UUID.ValueString())
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

		results, err := d.client.List(ctx, filters)
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
