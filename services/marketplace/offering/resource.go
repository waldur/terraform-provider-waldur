package offering

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MarketplaceOfferingResource{}
var _ resource.ResourceWithImportState = &MarketplaceOfferingResource{}

func NewMarketplaceOfferingResource() resource.Resource {
	return &MarketplaceOfferingResource{}
}

// MarketplaceOfferingResource defines the resource implementation.
type MarketplaceOfferingResource struct {
	client *Client
}

// MarketplaceOfferingResourceModel describes the resource data model.
type MarketplaceOfferingResourceModel struct {
	MarketplaceOfferingModel
	Options         types.Object   `tfsdk:"options"`
	ResourceOptions types.Object   `tfsdk:"resource_options"`
	Timeouts        timeouts.Value `tfsdk:"timeouts"`
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
				MarkdownDescription: "Marketplace Offering UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"access_url": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Publicly accessible offering access URL",
			},
			"backend_id": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "ID of the backend",
			},
			"billable": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Purchase and usage is invoiced.",
			},
			"billing_type_classification": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Classify offering components by billing type. Returns 'limit_only', 'usage_only', or 'mixed'.",
			},
			"category": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Category",
			},
			"category_title": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Category title",
			},
			"category_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the category",
			},
			"citation_count": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Number of citations of a DOI",
			},
			"compliance_checklist": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
							Required:            true,
							MarkdownDescription: "Billing type",
							Validators: []validator.String{
								stringvalidator.OneOf("fixed", "usage", "limit", "one", "few"),
							},
						},
						"default_limit": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Default limit",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"description": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Description of the Marketplace Offering",
						},
						"is_boolean": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "Is boolean",
						},
						"is_prepaid": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "Is prepaid",
						},
						"limit_amount": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Limit amount",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"limit_period": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Limit period",
						},
						"max_available_limit": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Max available limit",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"max_prepaid_duration": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Max prepaid duration",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"max_value": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Max value",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"measured_unit": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Unit of measurement, for example, GB.",
						},
						"min_prepaid_duration": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Min prepaid duration",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"min_value": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Min value",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Display name for the measured unit, for example, Floating IP.",
						},
						"overage_component": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Overage component",
						},
						"type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Unique internal name of the measured unit, for example floating_ip.",
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_\-\/:]+$`), ""),
							},
						},
						"unit_factor": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "The conversion factor from backend units to measured_unit",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Components",
			},
			"country": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Country code (ISO 3166-1 alpha-2)",
			},
			"created": schema.StringAttribute{
				CustomType: timetypes.RFC3339Type{},
				Computed:   true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Created",
			},
			"customer": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Customer",
			},
			"datacite_doi": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Datacite doi",
			},
			"description": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Description of the Marketplace Offering",
			},
			"endpoints": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the Marketplace Offering",
						},
						"url": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "URL of the access endpoint",
						},
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
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
							MarkdownDescription: "Name of the Marketplace Offering",
						},
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Files",
			},
			"full_description": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Full description",
			},
			"getting_started": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Getting started",
			},
			"google_calendar_is_public": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Google calendar is public",
			},
			"google_calendar_link": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Get the Google Calendar link for an offering.",
			},
			"has_compliance_requirements": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Has compliance requirements",
			},
			"image": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Image",
			},
			"integration_guide": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Integration guide",
			},
			"integration_status": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"agent_type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Agent type",
						},
						"last_request_timestamp": schema.StringAttribute{
							CustomType:          timetypes.RFC3339Type{},
							Computed:            true,
							MarkdownDescription: "Last request timestamp",
						},
						"service_name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the service",
						},
						"status": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Status",
						},
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Integration status",
			},
			"latitude": schema.Float64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Latitude",
			},
			"longitude": schema.Float64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Longitude",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Name of the Marketplace Offering",
			},
			"options": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"order": schema.ListAttribute{
						ElementType:         types.StringType,
						Required:            true,
						MarkdownDescription: "Order",
					},
				},
				Optional: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Options",
			},
			"order_count": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
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
							MarkdownDescription: "Name of the Marketplace Offering",
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
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Organization groups",
			},
			"parent_description": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Parent description",
			},
			"parent_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the parent",
			},
			"parent_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the parent",
			},
			"partitions": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cpu_bind": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Default task binding policy (SLURM cpu_bind)",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"def_cpu_per_gpu": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Default CPUs allocated per GPU",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"def_mem_per_cpu": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Default memory per CPU in MB",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"def_mem_per_gpu": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Default memory per GPU in MB",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"def_mem_per_node": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Default memory per node in MB",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"default_time": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Default time limit in minutes",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
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
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"max_cpus_per_node": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Maximum allocated CPUs per node",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"max_cpus_per_socket": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Maximum allocated CPUs per socket",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"max_mem_per_cpu": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Maximum memory per CPU in MB",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"max_mem_per_node": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Maximum memory per node in MB",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"max_nodes": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Maximum nodes per job",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"max_time": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Maximum time limit in minutes",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"min_nodes": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Minimum nodes per job",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"partition_name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the SLURM partition",
						},
						"priority_tier": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Priority tier for scheduling and preemption",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(32767),
							},
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
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Partitions",
			},
			"paused_reason": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
						"description": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Description of the Marketplace Offering",
						},
						"max_amount": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Maximum number of plans that could be active. Plan is disabled when maximum amount is reached.",
							Validators: []validator.Int64{
								int64validator.AtLeast(1),
								int64validator.AtMost(32767),
							},
						},
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Name of the Marketplace Offering",
						},
						"unit": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Unit",
							Validators: []validator.String{
								stringvalidator.OneOf("month", "quarter", "half_month", "day", "hour", "quantity"),
							},
						},
						"unit_price": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Unit price",
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,12}(?:\.\d{0,10})?$`), ""),
							},
						},
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Plans",
			},
			"privacy_policy_link": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Privacy policy link",
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
							MarkdownDescription: "Name of the Marketplace Offering",
						},
						"usage": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Usage",
						},
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Quotas",
			},
			"resource_options": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"order": schema.ListAttribute{
						ElementType:         types.StringType,
						Required:            true,
						MarkdownDescription: "Order",
					},
				},
				Optional: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Resource options",
			},
			"roles": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the Marketplace Offering",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Url",
						},
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Roles",
			},
			"scope": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Scope",
			},
			"scope_error_message": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Scope error message",
			},
			"scope_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the scope",
			},
			"scope_state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Scope state",
			},
			"scope_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
							MarkdownDescription: "Description of the Marketplace Offering",
						},
						"image": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Image",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the Marketplace Offering",
						},
						"thumbnail": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Thumbnail",
						},
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Screenshots",
			},
			"shared": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Accessible to all customers.",
			},
			"slug": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
									Optional:            true,
									MarkdownDescription: "Description of the Marketplace Offering",
								},
								"name": schema.StringAttribute{
									Optional:            true,
									MarkdownDescription: "Name of the Marketplace Offering",
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
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Software catalogs",
			},
			"state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "State",
			},
			"thumbnail": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Thumbnail",
			},
			"total_cost": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Total cost",
			},
			"total_cost_estimated": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Total cost estimated",
			},
			"total_customers": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Total customers",
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Type",
			},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Url",
			},
			"vendor_details": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Vendor details",
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

	r.client = &Client{}
	if err := r.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			err.Error(),
		)
		return
	}
}

func (r *MarketplaceOfferingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MarketplaceOfferingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := MarketplaceOfferingCreateRequest{
		AccessUrl:           data.AccessUrl.ValueStringPointer(),
		BackendId:           data.BackendId.ValueStringPointer(),
		Billable:            data.Billable.ValueBoolPointer(),
		Category:            data.Category.ValueStringPointer(),
		ComplianceChecklist: data.ComplianceChecklist.ValueStringPointer(),
		Country:             data.Country.ValueStringPointer(),
		Customer:            data.Customer.ValueStringPointer(),
		DataciteDoi:         data.DataciteDoi.ValueStringPointer(),
		Description:         data.Description.ValueStringPointer(),
		FullDescription:     data.FullDescription.ValueStringPointer(),
		GettingStarted:      data.GettingStarted.ValueStringPointer(),
		Image:               data.Image.ValueStringPointer(),
		IntegrationGuide:    data.IntegrationGuide.ValueStringPointer(),
		Latitude:            data.Latitude.ValueFloat64Pointer(),
		Longitude:           data.Longitude.ValueFloat64Pointer(),
		Name:                data.Name.ValueStringPointer(),
		PrivacyPolicyLink:   data.PrivacyPolicyLink.ValueStringPointer(),
		Shared:              data.Shared.ValueBoolPointer(),
		Slug:                data.Slug.ValueStringPointer(),
		Thumbnail:           data.Thumbnail.ValueStringPointer(),
		Type:                data.Type.ValueStringPointer(),
		VendorDetails:       data.VendorDetails.ValueStringPointer(),
	}
	resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.Components, &requestBody.Components)...)
	resp.Diagnostics.Append(common.PopulateOptionalObjectField(ctx, data.Options, &requestBody.Options)...)
	resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.Plans, &requestBody.Plans)...)
	resp.Diagnostics.Append(common.PopulateOptionalObjectField(ctx, data.ResourceOptions, &requestBody.ResourceOptions)...)

	apiResp, err := r.client.CreateMarketplaceOffering(ctx, &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Marketplace Offering",
			"An error occurred while creating the Marketplace Offering: "+err.Error(),
		)
		return
	}
	data.UUID = types.StringPointerValue(apiResp.UUID)

	createTimeout, diags := data.Timeouts.Create(ctx, common.DefaultCreateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*MarketplaceOfferingResponse, error) {
		return r.client.GetMarketplaceOffering(ctx, data.UUID.ValueString())
	}, createTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource creation", err.Error())
		return
	}
	apiResp = newResp

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

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

	apiResp, err := r.client.GetMarketplaceOffering(ctx, data.UUID.ValueString())
	if err != nil {
		if IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Marketplace Offering",
			"An error occurred while reading the Marketplace Offering: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

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

	err := r.client.DeleteMarketplaceOffering(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Marketplace Offering",
			"An error occurred while deleting the Marketplace Offering: "+err.Error(),
		)
		return
	}

	deleteTimeout, diags := data.Timeouts.Delete(ctx, common.DefaultDeleteTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err = common.WaitForDeletion(ctx, func(ctx context.Context) (*MarketplaceOfferingResponse, error) {
		return r.client.GetMarketplaceOffering(ctx, data.UUID.ValueString())
	}, deleteTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource deletion", err.Error())
		return
	}
}

func (r *MarketplaceOfferingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	uuid := req.ID
	if uuid == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Import ID cannot be empty. Please provide the UUID of the Marketplace Offering.",
		)
		return
	}

	tflog.Info(ctx, "Importing Marketplace Offering", map[string]interface{}{
		"uuid": uuid,
	})

	apiResp, err := r.client.GetMarketplaceOffering(ctx, uuid)
	if err != nil {
		if IsNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("Marketplace Offering with UUID '%s' does not exist or is not accessible.", uuid),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to Import Marketplace Offering",
			fmt.Sprintf("An error occurred while fetching the Marketplace Offering: %s", err.Error()),
		)
		return
	}

	var data MarketplaceOfferingResourceModel
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
