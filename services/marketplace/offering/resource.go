package offering

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/waldur/terraform-provider-waldur/internal/client"
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
						},
						"default_limit": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Default limit",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Description of the resource",
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
						},
						"unit_factor": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "The conversion factor from backend units to measured_unit",
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
				Computed: true,
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
				MarkdownDescription: "Name of the resource",
			},
			"options": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"order": schema.ListAttribute{
						CustomType:          types.ListType{ElemType: types.StringType},
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
							MarkdownDescription: "Description of the resource",
						},
						"max_amount": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Maximum number of plans that could be active. Plan is disabled when maximum amount is reached.",
						},
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Name of the resource",
						},
						"unit": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Unit",
						},
						"unit_price": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Unit price",
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
							MarkdownDescription: "Name of the resource",
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
						CustomType:          types.ListType{ElemType: types.StringType},
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
							MarkdownDescription: "Name of the resource",
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

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	r.client = NewClient(client)
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
	{
		// Object array or other
		var items []common.OfferingComponentRequest
		diags := data.Components.ElementsAs(ctx, &items, false)
		resp.Diagnostics.Append(diags...)
		if !diags.HasError() {
			if !data.Components.IsNull() && !data.Components.IsUnknown() {
				requestBody.Components = &items
			}
		}
	}
	{
		var obj common.OfferingOptionsRequest
		if diags := data.Options.As(ctx, &obj, basetypes.ObjectAsOptions{}); !diags.HasError() {
			requestBody.Options = &obj
		}
	}
	{
		// Object array or other
		var items []common.BaseProviderPlanRequest
		diags := data.Plans.ElementsAs(ctx, &items, false)
		resp.Diagnostics.Append(diags...)
		if !diags.HasError() {
			if !data.Plans.IsNull() && !data.Plans.IsUnknown() {
				requestBody.Plans = &items
			}
		}
	}
	{
		var obj common.OfferingOptionsRequest
		if diags := data.ResourceOptions.As(ctx, &obj, basetypes.ObjectAsOptions{}); !diags.HasError() {
			requestBody.ResourceOptions = &obj
		}
	}

	apiResp, err := r.client.CreateMarketplaceOffering(ctx, &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Marketplace Offering",
			"An error occurred while creating the Marketplace Offering: "+err.Error(),
		)
		return
	}
	data.UUID = types.StringPointerValue(apiResp.UUID)

	createTimeout, diags := data.Timeouts.Create(ctx, 30*time.Minute)
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

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

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
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Marketplace Offering",
			"An error occurred while reading the Marketplace Offering: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

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

	deleteTimeout, diags := data.Timeouts.Delete(ctx, 10*time.Minute)
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

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *MarketplaceOfferingResource) mapResponseToModel(ctx context.Context, apiResp MarketplaceOfferingResponse, model *MarketplaceOfferingResourceModel) diag.Diagnostics {
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
		}}, apiResp.Components)
		diags.Append(listDiagsComponents...)
		model.Components = listValComponents
	}
	model.Country = types.StringPointerValue(apiResp.Country)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Customer = types.StringPointerValue(apiResp.Customer)
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

	{
		listValIntegrationStatus, listDiagsIntegrationStatus := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"agent_type":             types.StringType,
			"last_request_timestamp": types.StringType,
			"service_name":           types.StringType,
			"status":                 types.StringType,
		}}, apiResp.IntegrationStatus)
		diags.Append(listDiagsIntegrationStatus...)
		model.IntegrationStatus = listValIntegrationStatus
	}
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
			"description":  types.StringType,
			"max_amount":   types.Int64Type,
			"name":         types.StringType,
			"unit":         types.StringType,
			"unit_price":   types.StringType,
		}}, apiResp.Plans)
		diags.Append(listDiagsPlans...)
		model.Plans = listValPlans
	}
	model.PrivacyPolicyLink = types.StringPointerValue(apiResp.PrivacyPolicyLink)

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
	model.VendorDetails = types.StringPointerValue(apiResp.VendorDetails)

	return diags
}
