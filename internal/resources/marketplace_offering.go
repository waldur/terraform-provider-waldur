package resources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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

// MarketplaceOfferingResourceModel describes the resource data model.
type MarketplaceOfferingResourceModel struct {
	UUID                      types.String   `tfsdk:"id"`
	AccessUrl                 types.String   `tfsdk:"access_url"`
	BackendId                 types.String   `tfsdk:"backend_id"`
	Billable                  types.Bool     `tfsdk:"billable"`
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
				MarkdownDescription: " ",
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
							MarkdownDescription: " ",
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
							MarkdownDescription: " ",
						},
						"parent_uuid": schema.StringAttribute{
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
				MarkdownDescription: " ",
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

	// Prepare request body
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

	// Call Waldur API to create resource
	var result map[string]interface{}
	err := r.client.Create(ctx, "/api/marketplace-provider-offerings/", requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create MarketplaceOffering",
			"An error occurred while creating the marketplace_offering: "+err.Error(),
		)
		return
	}
	// Extract UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
	// Map response fields to data model
	_ = sourceMap
	if val, ok := sourceMap["access_url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AccessUrl = types.StringValue(str)
		}
	} else {
		if data.AccessUrl.IsUnknown() {
			data.AccessUrl = types.StringNull()
		}
	}
	if val, ok := sourceMap["backend_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
		}
	} else {
		if data.BackendId.IsUnknown() {
			data.BackendId = types.StringNull()
		}
	}
	if val, ok := sourceMap["billable"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Billable = types.BoolValue(b)
		}
	} else {
		if data.Billable.IsUnknown() {
			data.Billable = types.BoolNull()
		}
	}
	if val, ok := sourceMap["category"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Category = types.StringValue(str)
		}
	} else {
		if data.Category.IsUnknown() {
			data.Category = types.StringNull()
		}
	}
	if val, ok := sourceMap["category_title"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CategoryTitle = types.StringValue(str)
		}
	} else {
		if data.CategoryTitle.IsUnknown() {
			data.CategoryTitle = types.StringNull()
		}
	}
	if val, ok := sourceMap["category_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CategoryUuid = types.StringValue(str)
		}
	} else {
		if data.CategoryUuid.IsUnknown() {
			data.CategoryUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["citation_count"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.CitationCount = types.Int64Value(int64(num))
		}
	} else {
		if data.CitationCount.IsUnknown() {
			data.CitationCount = types.Int64Null()
		}
	}
	if val, ok := sourceMap["compliance_checklist"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ComplianceChecklist = types.StringValue(str)
		}
	} else {
		if data.ComplianceChecklist.IsUnknown() {
			data.ComplianceChecklist = types.StringNull()
		}
	}
	if val, ok := sourceMap["components"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
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
					attrValues := map[string]attr.Value{
						"article_code": func() attr.Value {
							if v, ok := objMap["article_code"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"billing_type": func() attr.Value {
							if v, ok := objMap["billing_type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"default_limit": func() attr.Value {
							if v, ok := objMap["default_limit"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"description": func() attr.Value {
							if v, ok := objMap["description"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"is_boolean": func() attr.Value {
							if v, ok := objMap["is_boolean"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
						"is_prepaid": func() attr.Value {
							if v, ok := objMap["is_prepaid"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
						"limit_amount": func() attr.Value {
							if v, ok := objMap["limit_amount"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"limit_period": func() attr.Value {
							if v, ok := objMap["limit_period"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"max_available_limit": func() attr.Value {
							if v, ok := objMap["max_available_limit"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"max_prepaid_duration": func() attr.Value {
							if v, ok := objMap["max_prepaid_duration"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"max_value": func() attr.Value {
							if v, ok := objMap["max_value"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"measured_unit": func() attr.Value {
							if v, ok := objMap["measured_unit"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"min_prepaid_duration": func() attr.Value {
							if v, ok := objMap["min_prepaid_duration"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"min_value": func() attr.Value {
							if v, ok := objMap["min_value"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"overage_component": func() attr.Value {
							if v, ok := objMap["overage_component"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"type": func() attr.Value {
							if v, ok := objMap["type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"unit_factor": func() attr.Value {
							if v, ok := objMap["unit_factor"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}}, items)
			data.Components = listVal
		}
	} else {
		if data.Components.IsUnknown() {
			data.Components = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}})
		}
	}
	if val, ok := sourceMap["country"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Country = types.StringValue(str)
		}
	} else {
		if data.Country.IsUnknown() {
			data.Country = types.StringNull()
		}
	}
	if val, ok := sourceMap["created"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
	} else {
		if data.Created.IsUnknown() {
			data.Created = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Customer = types.StringValue(str)
		}
	} else {
		if data.Customer.IsUnknown() {
			data.Customer = types.StringNull()
		}
	}
	if val, ok := sourceMap["datacite_doi"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.DataciteDoi = types.StringValue(str)
		}
	} else {
		if data.DataciteDoi.IsUnknown() {
			data.DataciteDoi = types.StringNull()
		}
	}
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
		}
	}
	if val, ok := sourceMap["endpoints"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"name": types.StringType,
						"url":  types.StringType,
					}
					attrValues := map[string]attr.Value{
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"name": types.StringType,
				"url":  types.StringType,
			}}, items)
			data.Endpoints = listVal
		}
	} else {
		if data.Endpoints.IsUnknown() {
			data.Endpoints = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"name": types.StringType,
				"url":  types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["files"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"created": types.StringType,
						"file":    types.StringType,
						"name":    types.StringType,
					}
					attrValues := map[string]attr.Value{
						"created": func() attr.Value {
							if v, ok := objMap["created"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"file": func() attr.Value {
							if v, ok := objMap["file"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"created": types.StringType,
				"file":    types.StringType,
				"name":    types.StringType,
			}}, items)
			data.Files = listVal
		}
	} else {
		if data.Files.IsUnknown() {
			data.Files = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"created": types.StringType,
				"file":    types.StringType,
				"name":    types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["full_description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.FullDescription = types.StringValue(str)
		}
	} else {
		if data.FullDescription.IsUnknown() {
			data.FullDescription = types.StringNull()
		}
	}
	if val, ok := sourceMap["getting_started"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.GettingStarted = types.StringValue(str)
		}
	} else {
		if data.GettingStarted.IsUnknown() {
			data.GettingStarted = types.StringNull()
		}
	}
	if val, ok := sourceMap["google_calendar_is_public"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.GoogleCalendarIsPublic = types.BoolValue(b)
		}
	} else {
		if data.GoogleCalendarIsPublic.IsUnknown() {
			data.GoogleCalendarIsPublic = types.BoolNull()
		}
	}
	if val, ok := sourceMap["google_calendar_link"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.GoogleCalendarLink = types.StringValue(str)
		}
	} else {
		if data.GoogleCalendarLink.IsUnknown() {
			data.GoogleCalendarLink = types.StringNull()
		}
	}
	if val, ok := sourceMap["has_compliance_requirements"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.HasComplianceRequirements = types.BoolValue(b)
		}
	} else {
		if data.HasComplianceRequirements.IsUnknown() {
			data.HasComplianceRequirements = types.BoolNull()
		}
	}
	if val, ok := sourceMap["image"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Image = types.StringValue(str)
		}
	} else {
		if data.Image.IsUnknown() {
			data.Image = types.StringNull()
		}
	}
	if val, ok := sourceMap["integration_guide"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.IntegrationGuide = types.StringValue(str)
		}
	} else {
		if data.IntegrationGuide.IsUnknown() {
			data.IntegrationGuide = types.StringNull()
		}
	}
	if val, ok := sourceMap["integration_status"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"agent_type":             types.StringType,
						"last_request_timestamp": types.StringType,
						"service_name":           types.StringType,
						"status":                 types.StringType,
					}
					attrValues := map[string]attr.Value{
						"agent_type": func() attr.Value {
							if v, ok := objMap["agent_type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"last_request_timestamp": func() attr.Value {
							if v, ok := objMap["last_request_timestamp"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"service_name": func() attr.Value {
							if v, ok := objMap["service_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"status": func() attr.Value {
							if v, ok := objMap["status"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"agent_type":             types.StringType,
				"last_request_timestamp": types.StringType,
				"service_name":           types.StringType,
				"status":                 types.StringType,
			}}, items)
			data.IntegrationStatus = listVal
		}
	} else {
		if data.IntegrationStatus.IsUnknown() {
			data.IntegrationStatus = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"agent_type":             types.StringType,
				"last_request_timestamp": types.StringType,
				"service_name":           types.StringType,
				"status":                 types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["latitude"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Latitude = types.Float64Value(num)
		}
	} else {
		if data.Latitude.IsUnknown() {
			data.Latitude = types.Float64Null()
		}
	}
	if val, ok := sourceMap["longitude"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Longitude = types.Float64Value(num)
		}
	} else {
		if data.Longitude.IsUnknown() {
			data.Longitude = types.Float64Null()
		}
	}
	if val, ok := sourceMap["options"]; ok && val != nil {
		// Nested object
		if objMap, ok := val.(map[string]interface{}); ok {
			_ = objMap // Avoid unused variable if properties are empty
			attrTypes := map[string]attr.Type{
				"order": types.ListType{ElemType: types.StringType},
			}
			attrValues := map[string]attr.Value{
				"order": types.ListNull(types.ListType{ElemType: types.StringType}.ElemType),
			}
			objVal, _ := types.ObjectValue(attrTypes, attrValues)
			data.Options = objVal
		}
	} else {
		if data.Options.IsUnknown() {
			data.Options = types.ObjectNull(map[string]attr.Type{
				"order": types.ListType{ElemType: types.StringType},
			})
		}
	}
	if val, ok := sourceMap["order_count"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.OrderCount = types.Int64Value(int64(num))
		}
	} else {
		if data.OrderCount.IsUnknown() {
			data.OrderCount = types.Int64Null()
		}
	}
	if val, ok := sourceMap["organization_groups"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"customers_count": types.Int64Type,
						"name":            types.StringType,
						"parent":          types.StringType,
						"parent_name":     types.StringType,
						"parent_uuid":     types.StringType,
						"url":             types.StringType,
					}
					attrValues := map[string]attr.Value{
						"customers_count": func() attr.Value {
							if v, ok := objMap["customers_count"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"parent": func() attr.Value {
							if v, ok := objMap["parent"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"parent_name": func() attr.Value {
							if v, ok := objMap["parent_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"parent_uuid": func() attr.Value {
							if v, ok := objMap["parent_uuid"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"customers_count": types.Int64Type,
				"name":            types.StringType,
				"parent":          types.StringType,
				"parent_name":     types.StringType,
				"parent_uuid":     types.StringType,
				"url":             types.StringType,
			}}, items)
			data.OrganizationGroups = listVal
		}
	} else {
		if data.OrganizationGroups.IsUnknown() {
			data.OrganizationGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"customers_count": types.Int64Type,
				"name":            types.StringType,
				"parent":          types.StringType,
				"parent_name":     types.StringType,
				"parent_uuid":     types.StringType,
				"url":             types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["parent_description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ParentDescription = types.StringValue(str)
		}
	} else {
		if data.ParentDescription.IsUnknown() {
			data.ParentDescription = types.StringNull()
		}
	}
	if val, ok := sourceMap["parent_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ParentName = types.StringValue(str)
		}
	} else {
		if data.ParentName.IsUnknown() {
			data.ParentName = types.StringNull()
		}
	}
	if val, ok := sourceMap["parent_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ParentUuid = types.StringValue(str)
		}
	} else {
		if data.ParentUuid.IsUnknown() {
			data.ParentUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["partitions"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
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
					attrValues := map[string]attr.Value{
						"cpu_bind": func() attr.Value {
							if v, ok := objMap["cpu_bind"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"def_cpu_per_gpu": func() attr.Value {
							if v, ok := objMap["def_cpu_per_gpu"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"def_mem_per_cpu": func() attr.Value {
							if v, ok := objMap["def_mem_per_cpu"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"def_mem_per_gpu": func() attr.Value {
							if v, ok := objMap["def_mem_per_gpu"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"def_mem_per_node": func() attr.Value {
							if v, ok := objMap["def_mem_per_node"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"default_time": func() attr.Value {
							if v, ok := objMap["default_time"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"exclusive_topo": func() attr.Value {
							if v, ok := objMap["exclusive_topo"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
						"exclusive_user": func() attr.Value {
							if v, ok := objMap["exclusive_user"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
						"grace_time": func() attr.Value {
							if v, ok := objMap["grace_time"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"max_cpus_per_node": func() attr.Value {
							if v, ok := objMap["max_cpus_per_node"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"max_cpus_per_socket": func() attr.Value {
							if v, ok := objMap["max_cpus_per_socket"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"max_mem_per_cpu": func() attr.Value {
							if v, ok := objMap["max_mem_per_cpu"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"max_mem_per_node": func() attr.Value {
							if v, ok := objMap["max_mem_per_node"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"max_nodes": func() attr.Value {
							if v, ok := objMap["max_nodes"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"max_time": func() attr.Value {
							if v, ok := objMap["max_time"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"min_nodes": func() attr.Value {
							if v, ok := objMap["min_nodes"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"partition_name": func() attr.Value {
							if v, ok := objMap["partition_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"priority_tier": func() attr.Value {
							if v, ok := objMap["priority_tier"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"qos": func() attr.Value {
							if v, ok := objMap["qos"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"req_resv": func() attr.Value {
							if v, ok := objMap["req_resv"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}}, items)
			data.Partitions = listVal
		}
	} else {
		if data.Partitions.IsUnknown() {
			data.Partitions = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}})
		}
	}
	if val, ok := sourceMap["paused_reason"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.PausedReason = types.StringValue(str)
		}
	} else {
		if data.PausedReason.IsUnknown() {
			data.PausedReason = types.StringNull()
		}
	}
	if val, ok := sourceMap["plans"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"archived":     types.BoolType,
						"article_code": types.StringType,
						"backend_id":   types.StringType,
						"description":  types.StringType,
						"max_amount":   types.Int64Type,
						"name":         types.StringType,
						"unit":         types.StringType,
						"unit_price":   types.StringType,
					}
					attrValues := map[string]attr.Value{
						"archived": func() attr.Value {
							if v, ok := objMap["archived"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
						"article_code": func() attr.Value {
							if v, ok := objMap["article_code"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"backend_id": func() attr.Value {
							if v, ok := objMap["backend_id"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"description": func() attr.Value {
							if v, ok := objMap["description"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"max_amount": func() attr.Value {
							if v, ok := objMap["max_amount"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"unit": func() attr.Value {
							if v, ok := objMap["unit"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"unit_price": func() attr.Value {
							if v, ok := objMap["unit_price"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"archived":     types.BoolType,
				"article_code": types.StringType,
				"backend_id":   types.StringType,
				"description":  types.StringType,
				"max_amount":   types.Int64Type,
				"name":         types.StringType,
				"unit":         types.StringType,
				"unit_price":   types.StringType,
			}}, items)
			data.Plans = listVal
		}
	} else {
		if data.Plans.IsUnknown() {
			data.Plans = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"archived":     types.BoolType,
				"article_code": types.StringType,
				"backend_id":   types.StringType,
				"description":  types.StringType,
				"max_amount":   types.Int64Type,
				"name":         types.StringType,
				"unit":         types.StringType,
				"unit_price":   types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["privacy_policy_link"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.PrivacyPolicyLink = types.StringValue(str)
		}
	} else {
		if data.PrivacyPolicyLink.IsUnknown() {
			data.PrivacyPolicyLink = types.StringNull()
		}
	}
	if val, ok := sourceMap["quotas"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"limit": types.Int64Type,
						"name":  types.StringType,
						"usage": types.Int64Type,
					}
					attrValues := map[string]attr.Value{
						"limit": func() attr.Value {
							if v, ok := objMap["limit"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"usage": func() attr.Value {
							if v, ok := objMap["usage"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"limit": types.Int64Type,
				"name":  types.StringType,
				"usage": types.Int64Type,
			}}, items)
			data.Quotas = listVal
		}
	} else {
		if data.Quotas.IsUnknown() {
			data.Quotas = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"limit": types.Int64Type,
				"name":  types.StringType,
				"usage": types.Int64Type,
			}})
		}
	}
	if val, ok := sourceMap["resource_options"]; ok && val != nil {
		// Nested object
		if objMap, ok := val.(map[string]interface{}); ok {
			_ = objMap // Avoid unused variable if properties are empty
			attrTypes := map[string]attr.Type{
				"order": types.ListType{ElemType: types.StringType},
			}
			attrValues := map[string]attr.Value{
				"order": types.ListNull(types.ListType{ElemType: types.StringType}.ElemType),
			}
			objVal, _ := types.ObjectValue(attrTypes, attrValues)
			data.ResourceOptions = objVal
		}
	} else {
		if data.ResourceOptions.IsUnknown() {
			data.ResourceOptions = types.ObjectNull(map[string]attr.Type{
				"order": types.ListType{ElemType: types.StringType},
			})
		}
	}
	if val, ok := sourceMap["roles"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"name": types.StringType,
						"url":  types.StringType,
					}
					attrValues := map[string]attr.Value{
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"name": types.StringType,
				"url":  types.StringType,
			}}, items)
			data.Roles = listVal
		}
	} else {
		if data.Roles.IsUnknown() {
			data.Roles = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"name": types.StringType,
				"url":  types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["scope"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Scope = types.StringValue(str)
		}
	} else {
		if data.Scope.IsUnknown() {
			data.Scope = types.StringNull()
		}
	}
	if val, ok := sourceMap["scope_error_message"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ScopeErrorMessage = types.StringValue(str)
		}
	} else {
		if data.ScopeErrorMessage.IsUnknown() {
			data.ScopeErrorMessage = types.StringNull()
		}
	}
	if val, ok := sourceMap["scope_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ScopeName = types.StringValue(str)
		}
	} else {
		if data.ScopeName.IsUnknown() {
			data.ScopeName = types.StringNull()
		}
	}
	if val, ok := sourceMap["scope_state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ScopeState = types.StringValue(str)
		}
	} else {
		if data.ScopeState.IsUnknown() {
			data.ScopeState = types.StringNull()
		}
	}
	if val, ok := sourceMap["scope_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ScopeUuid = types.StringValue(str)
		}
	} else {
		if data.ScopeUuid.IsUnknown() {
			data.ScopeUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["screenshots"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"created":     types.StringType,
						"description": types.StringType,
						"image":       types.StringType,
						"name":        types.StringType,
						"thumbnail":   types.StringType,
					}
					attrValues := map[string]attr.Value{
						"created": func() attr.Value {
							if v, ok := objMap["created"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"description": func() attr.Value {
							if v, ok := objMap["description"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"image": func() attr.Value {
							if v, ok := objMap["image"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"thumbnail": func() attr.Value {
							if v, ok := objMap["thumbnail"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"created":     types.StringType,
				"description": types.StringType,
				"image":       types.StringType,
				"name":        types.StringType,
				"thumbnail":   types.StringType,
			}}, items)
			data.Screenshots = listVal
		}
	} else {
		if data.Screenshots.IsUnknown() {
			data.Screenshots = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"created":     types.StringType,
				"description": types.StringType,
				"image":       types.StringType,
				"name":        types.StringType,
				"thumbnail":   types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["shared"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Shared = types.BoolValue(b)
		}
	} else {
		if data.Shared.IsUnknown() {
			data.Shared = types.BoolNull()
		}
	}
	if val, ok := sourceMap["slug"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Slug = types.StringValue(str)
		}
	} else {
		if data.Slug.IsUnknown() {
			data.Slug = types.StringNull()
		}
	}
	if val, ok := sourceMap["software_catalogs"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
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
					attrValues := map[string]attr.Value{
						"catalog": types.ObjectNull(types.ObjectType{AttrTypes: map[string]attr.Type{
							"description": types.StringType,
							"name":        types.StringType,
							"version":     types.StringType,
						}}.AttrTypes),
						"package_count": func() attr.Value {
							if v, ok := objMap["package_count"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"partition": types.ObjectNull(types.ObjectType{AttrTypes: map[string]attr.Type{
							"partition_name": types.StringType,
							"priority_tier":  types.Int64Type,
							"qos":            types.StringType,
						}}.AttrTypes),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}}, items)
			data.SoftwareCatalogs = listVal
		}
	} else {
		if data.SoftwareCatalogs.IsUnknown() {
			data.SoftwareCatalogs = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}})
		}
	}
	if val, ok := sourceMap["state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.State = types.StringValue(str)
		}
	} else {
		if data.State.IsUnknown() {
			data.State = types.StringNull()
		}
	}
	if val, ok := sourceMap["thumbnail"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Thumbnail = types.StringValue(str)
		}
	} else {
		if data.Thumbnail.IsUnknown() {
			data.Thumbnail = types.StringNull()
		}
	}
	if val, ok := sourceMap["total_cost"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.TotalCost = types.Int64Value(int64(num))
		}
	} else {
		if data.TotalCost.IsUnknown() {
			data.TotalCost = types.Int64Null()
		}
	}
	if val, ok := sourceMap["total_cost_estimated"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.TotalCostEstimated = types.Int64Value(int64(num))
		}
	} else {
		if data.TotalCostEstimated.IsUnknown() {
			data.TotalCostEstimated = types.Int64Null()
		}
	}
	if val, ok := sourceMap["total_customers"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.TotalCustomers = types.Int64Value(int64(num))
		}
	} else {
		if data.TotalCustomers.IsUnknown() {
			data.TotalCustomers = types.Int64Null()
		}
	}
	if val, ok := sourceMap["type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Type = types.StringValue(str)
		}
	} else {
		if data.Type.IsUnknown() {
			data.Type = types.StringNull()
		}
	}
	if val, ok := sourceMap["url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Url = types.StringValue(str)
		}
	} else {
		if data.Url.IsUnknown() {
			data.Url = types.StringNull()
		}
	}
	if val, ok := sourceMap["vendor_details"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.VendorDetails = types.StringValue(str)
		}
	} else {
		if data.VendorDetails.IsUnknown() {
			data.VendorDetails = types.StringNull()
		}
	}

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
	var result map[string]interface{}

	retrievePath := strings.Replace("/api/marketplace-provider-offerings/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read MarketplaceOffering",
			"An error occurred while reading the marketplace_offering: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
	// Map response fields to data model
	_ = sourceMap
	if val, ok := sourceMap["access_url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AccessUrl = types.StringValue(str)
		}
	} else {
		if data.AccessUrl.IsUnknown() {
			data.AccessUrl = types.StringNull()
		}
	}
	if val, ok := sourceMap["backend_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
		}
	} else {
		if data.BackendId.IsUnknown() {
			data.BackendId = types.StringNull()
		}
	}
	if val, ok := sourceMap["billable"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Billable = types.BoolValue(b)
		}
	} else {
		if data.Billable.IsUnknown() {
			data.Billable = types.BoolNull()
		}
	}
	if val, ok := sourceMap["category"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Category = types.StringValue(str)
		}
	} else {
		if data.Category.IsUnknown() {
			data.Category = types.StringNull()
		}
	}
	if val, ok := sourceMap["category_title"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CategoryTitle = types.StringValue(str)
		}
	} else {
		if data.CategoryTitle.IsUnknown() {
			data.CategoryTitle = types.StringNull()
		}
	}
	if val, ok := sourceMap["category_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CategoryUuid = types.StringValue(str)
		}
	} else {
		if data.CategoryUuid.IsUnknown() {
			data.CategoryUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["citation_count"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.CitationCount = types.Int64Value(int64(num))
		}
	} else {
		if data.CitationCount.IsUnknown() {
			data.CitationCount = types.Int64Null()
		}
	}
	if val, ok := sourceMap["compliance_checklist"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ComplianceChecklist = types.StringValue(str)
		}
	} else {
		if data.ComplianceChecklist.IsUnknown() {
			data.ComplianceChecklist = types.StringNull()
		}
	}
	if val, ok := sourceMap["components"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
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
					attrValues := map[string]attr.Value{
						"article_code": func() attr.Value {
							if v, ok := objMap["article_code"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"billing_type": func() attr.Value {
							if v, ok := objMap["billing_type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"default_limit": func() attr.Value {
							if v, ok := objMap["default_limit"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"description": func() attr.Value {
							if v, ok := objMap["description"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"is_boolean": func() attr.Value {
							if v, ok := objMap["is_boolean"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
						"is_prepaid": func() attr.Value {
							if v, ok := objMap["is_prepaid"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
						"limit_amount": func() attr.Value {
							if v, ok := objMap["limit_amount"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"limit_period": func() attr.Value {
							if v, ok := objMap["limit_period"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"max_available_limit": func() attr.Value {
							if v, ok := objMap["max_available_limit"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"max_prepaid_duration": func() attr.Value {
							if v, ok := objMap["max_prepaid_duration"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"max_value": func() attr.Value {
							if v, ok := objMap["max_value"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"measured_unit": func() attr.Value {
							if v, ok := objMap["measured_unit"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"min_prepaid_duration": func() attr.Value {
							if v, ok := objMap["min_prepaid_duration"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"min_value": func() attr.Value {
							if v, ok := objMap["min_value"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"overage_component": func() attr.Value {
							if v, ok := objMap["overage_component"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"type": func() attr.Value {
							if v, ok := objMap["type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"unit_factor": func() attr.Value {
							if v, ok := objMap["unit_factor"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}}, items)
			data.Components = listVal
		}
	} else {
		if data.Components.IsUnknown() {
			data.Components = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}})
		}
	}
	if val, ok := sourceMap["country"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Country = types.StringValue(str)
		}
	} else {
		if data.Country.IsUnknown() {
			data.Country = types.StringNull()
		}
	}
	if val, ok := sourceMap["created"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
	} else {
		if data.Created.IsUnknown() {
			data.Created = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Customer = types.StringValue(str)
		}
	} else {
		if data.Customer.IsUnknown() {
			data.Customer = types.StringNull()
		}
	}
	if val, ok := sourceMap["datacite_doi"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.DataciteDoi = types.StringValue(str)
		}
	} else {
		if data.DataciteDoi.IsUnknown() {
			data.DataciteDoi = types.StringNull()
		}
	}
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
		}
	}
	if val, ok := sourceMap["endpoints"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"name": types.StringType,
						"url":  types.StringType,
					}
					attrValues := map[string]attr.Value{
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"name": types.StringType,
				"url":  types.StringType,
			}}, items)
			data.Endpoints = listVal
		}
	} else {
		if data.Endpoints.IsUnknown() {
			data.Endpoints = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"name": types.StringType,
				"url":  types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["files"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"created": types.StringType,
						"file":    types.StringType,
						"name":    types.StringType,
					}
					attrValues := map[string]attr.Value{
						"created": func() attr.Value {
							if v, ok := objMap["created"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"file": func() attr.Value {
							if v, ok := objMap["file"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"created": types.StringType,
				"file":    types.StringType,
				"name":    types.StringType,
			}}, items)
			data.Files = listVal
		}
	} else {
		if data.Files.IsUnknown() {
			data.Files = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"created": types.StringType,
				"file":    types.StringType,
				"name":    types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["full_description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.FullDescription = types.StringValue(str)
		}
	} else {
		if data.FullDescription.IsUnknown() {
			data.FullDescription = types.StringNull()
		}
	}
	if val, ok := sourceMap["getting_started"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.GettingStarted = types.StringValue(str)
		}
	} else {
		if data.GettingStarted.IsUnknown() {
			data.GettingStarted = types.StringNull()
		}
	}
	if val, ok := sourceMap["google_calendar_is_public"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.GoogleCalendarIsPublic = types.BoolValue(b)
		}
	} else {
		if data.GoogleCalendarIsPublic.IsUnknown() {
			data.GoogleCalendarIsPublic = types.BoolNull()
		}
	}
	if val, ok := sourceMap["google_calendar_link"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.GoogleCalendarLink = types.StringValue(str)
		}
	} else {
		if data.GoogleCalendarLink.IsUnknown() {
			data.GoogleCalendarLink = types.StringNull()
		}
	}
	if val, ok := sourceMap["has_compliance_requirements"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.HasComplianceRequirements = types.BoolValue(b)
		}
	} else {
		if data.HasComplianceRequirements.IsUnknown() {
			data.HasComplianceRequirements = types.BoolNull()
		}
	}
	if val, ok := sourceMap["image"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Image = types.StringValue(str)
		}
	} else {
		if data.Image.IsUnknown() {
			data.Image = types.StringNull()
		}
	}
	if val, ok := sourceMap["integration_guide"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.IntegrationGuide = types.StringValue(str)
		}
	} else {
		if data.IntegrationGuide.IsUnknown() {
			data.IntegrationGuide = types.StringNull()
		}
	}
	if val, ok := sourceMap["integration_status"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"agent_type":             types.StringType,
						"last_request_timestamp": types.StringType,
						"service_name":           types.StringType,
						"status":                 types.StringType,
					}
					attrValues := map[string]attr.Value{
						"agent_type": func() attr.Value {
							if v, ok := objMap["agent_type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"last_request_timestamp": func() attr.Value {
							if v, ok := objMap["last_request_timestamp"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"service_name": func() attr.Value {
							if v, ok := objMap["service_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"status": func() attr.Value {
							if v, ok := objMap["status"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"agent_type":             types.StringType,
				"last_request_timestamp": types.StringType,
				"service_name":           types.StringType,
				"status":                 types.StringType,
			}}, items)
			data.IntegrationStatus = listVal
		}
	} else {
		if data.IntegrationStatus.IsUnknown() {
			data.IntegrationStatus = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"agent_type":             types.StringType,
				"last_request_timestamp": types.StringType,
				"service_name":           types.StringType,
				"status":                 types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["latitude"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Latitude = types.Float64Value(num)
		}
	} else {
		if data.Latitude.IsUnknown() {
			data.Latitude = types.Float64Null()
		}
	}
	if val, ok := sourceMap["longitude"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Longitude = types.Float64Value(num)
		}
	} else {
		if data.Longitude.IsUnknown() {
			data.Longitude = types.Float64Null()
		}
	}
	if val, ok := sourceMap["options"]; ok && val != nil {
		// Nested object
		if objMap, ok := val.(map[string]interface{}); ok {
			_ = objMap // Avoid unused variable if properties are empty
			attrTypes := map[string]attr.Type{
				"order": types.ListType{ElemType: types.StringType},
			}
			attrValues := map[string]attr.Value{
				"order": types.ListNull(types.ListType{ElemType: types.StringType}.ElemType),
			}
			objVal, _ := types.ObjectValue(attrTypes, attrValues)
			data.Options = objVal
		}
	} else {
		if data.Options.IsUnknown() {
			data.Options = types.ObjectNull(map[string]attr.Type{
				"order": types.ListType{ElemType: types.StringType},
			})
		}
	}
	if val, ok := sourceMap["order_count"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.OrderCount = types.Int64Value(int64(num))
		}
	} else {
		if data.OrderCount.IsUnknown() {
			data.OrderCount = types.Int64Null()
		}
	}
	if val, ok := sourceMap["organization_groups"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"customers_count": types.Int64Type,
						"name":            types.StringType,
						"parent":          types.StringType,
						"parent_name":     types.StringType,
						"parent_uuid":     types.StringType,
						"url":             types.StringType,
					}
					attrValues := map[string]attr.Value{
						"customers_count": func() attr.Value {
							if v, ok := objMap["customers_count"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"parent": func() attr.Value {
							if v, ok := objMap["parent"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"parent_name": func() attr.Value {
							if v, ok := objMap["parent_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"parent_uuid": func() attr.Value {
							if v, ok := objMap["parent_uuid"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"customers_count": types.Int64Type,
				"name":            types.StringType,
				"parent":          types.StringType,
				"parent_name":     types.StringType,
				"parent_uuid":     types.StringType,
				"url":             types.StringType,
			}}, items)
			data.OrganizationGroups = listVal
		}
	} else {
		if data.OrganizationGroups.IsUnknown() {
			data.OrganizationGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"customers_count": types.Int64Type,
				"name":            types.StringType,
				"parent":          types.StringType,
				"parent_name":     types.StringType,
				"parent_uuid":     types.StringType,
				"url":             types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["parent_description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ParentDescription = types.StringValue(str)
		}
	} else {
		if data.ParentDescription.IsUnknown() {
			data.ParentDescription = types.StringNull()
		}
	}
	if val, ok := sourceMap["parent_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ParentName = types.StringValue(str)
		}
	} else {
		if data.ParentName.IsUnknown() {
			data.ParentName = types.StringNull()
		}
	}
	if val, ok := sourceMap["parent_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ParentUuid = types.StringValue(str)
		}
	} else {
		if data.ParentUuid.IsUnknown() {
			data.ParentUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["partitions"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
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
					attrValues := map[string]attr.Value{
						"cpu_bind": func() attr.Value {
							if v, ok := objMap["cpu_bind"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"def_cpu_per_gpu": func() attr.Value {
							if v, ok := objMap["def_cpu_per_gpu"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"def_mem_per_cpu": func() attr.Value {
							if v, ok := objMap["def_mem_per_cpu"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"def_mem_per_gpu": func() attr.Value {
							if v, ok := objMap["def_mem_per_gpu"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"def_mem_per_node": func() attr.Value {
							if v, ok := objMap["def_mem_per_node"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"default_time": func() attr.Value {
							if v, ok := objMap["default_time"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"exclusive_topo": func() attr.Value {
							if v, ok := objMap["exclusive_topo"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
						"exclusive_user": func() attr.Value {
							if v, ok := objMap["exclusive_user"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
						"grace_time": func() attr.Value {
							if v, ok := objMap["grace_time"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"max_cpus_per_node": func() attr.Value {
							if v, ok := objMap["max_cpus_per_node"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"max_cpus_per_socket": func() attr.Value {
							if v, ok := objMap["max_cpus_per_socket"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"max_mem_per_cpu": func() attr.Value {
							if v, ok := objMap["max_mem_per_cpu"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"max_mem_per_node": func() attr.Value {
							if v, ok := objMap["max_mem_per_node"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"max_nodes": func() attr.Value {
							if v, ok := objMap["max_nodes"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"max_time": func() attr.Value {
							if v, ok := objMap["max_time"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"min_nodes": func() attr.Value {
							if v, ok := objMap["min_nodes"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"partition_name": func() attr.Value {
							if v, ok := objMap["partition_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"priority_tier": func() attr.Value {
							if v, ok := objMap["priority_tier"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"qos": func() attr.Value {
							if v, ok := objMap["qos"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"req_resv": func() attr.Value {
							if v, ok := objMap["req_resv"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}}, items)
			data.Partitions = listVal
		}
	} else {
		if data.Partitions.IsUnknown() {
			data.Partitions = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}})
		}
	}
	if val, ok := sourceMap["paused_reason"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.PausedReason = types.StringValue(str)
		}
	} else {
		if data.PausedReason.IsUnknown() {
			data.PausedReason = types.StringNull()
		}
	}
	if val, ok := sourceMap["plans"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"archived":     types.BoolType,
						"article_code": types.StringType,
						"backend_id":   types.StringType,
						"description":  types.StringType,
						"max_amount":   types.Int64Type,
						"name":         types.StringType,
						"unit":         types.StringType,
						"unit_price":   types.StringType,
					}
					attrValues := map[string]attr.Value{
						"archived": func() attr.Value {
							if v, ok := objMap["archived"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
						"article_code": func() attr.Value {
							if v, ok := objMap["article_code"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"backend_id": func() attr.Value {
							if v, ok := objMap["backend_id"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"description": func() attr.Value {
							if v, ok := objMap["description"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"max_amount": func() attr.Value {
							if v, ok := objMap["max_amount"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"unit": func() attr.Value {
							if v, ok := objMap["unit"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"unit_price": func() attr.Value {
							if v, ok := objMap["unit_price"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"archived":     types.BoolType,
				"article_code": types.StringType,
				"backend_id":   types.StringType,
				"description":  types.StringType,
				"max_amount":   types.Int64Type,
				"name":         types.StringType,
				"unit":         types.StringType,
				"unit_price":   types.StringType,
			}}, items)
			data.Plans = listVal
		}
	} else {
		if data.Plans.IsUnknown() {
			data.Plans = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"archived":     types.BoolType,
				"article_code": types.StringType,
				"backend_id":   types.StringType,
				"description":  types.StringType,
				"max_amount":   types.Int64Type,
				"name":         types.StringType,
				"unit":         types.StringType,
				"unit_price":   types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["privacy_policy_link"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.PrivacyPolicyLink = types.StringValue(str)
		}
	} else {
		if data.PrivacyPolicyLink.IsUnknown() {
			data.PrivacyPolicyLink = types.StringNull()
		}
	}
	if val, ok := sourceMap["quotas"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"limit": types.Int64Type,
						"name":  types.StringType,
						"usage": types.Int64Type,
					}
					attrValues := map[string]attr.Value{
						"limit": func() attr.Value {
							if v, ok := objMap["limit"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"usage": func() attr.Value {
							if v, ok := objMap["usage"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"limit": types.Int64Type,
				"name":  types.StringType,
				"usage": types.Int64Type,
			}}, items)
			data.Quotas = listVal
		}
	} else {
		if data.Quotas.IsUnknown() {
			data.Quotas = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"limit": types.Int64Type,
				"name":  types.StringType,
				"usage": types.Int64Type,
			}})
		}
	}
	if val, ok := sourceMap["resource_options"]; ok && val != nil {
		// Nested object
		if objMap, ok := val.(map[string]interface{}); ok {
			_ = objMap // Avoid unused variable if properties are empty
			attrTypes := map[string]attr.Type{
				"order": types.ListType{ElemType: types.StringType},
			}
			attrValues := map[string]attr.Value{
				"order": types.ListNull(types.ListType{ElemType: types.StringType}.ElemType),
			}
			objVal, _ := types.ObjectValue(attrTypes, attrValues)
			data.ResourceOptions = objVal
		}
	} else {
		if data.ResourceOptions.IsUnknown() {
			data.ResourceOptions = types.ObjectNull(map[string]attr.Type{
				"order": types.ListType{ElemType: types.StringType},
			})
		}
	}
	if val, ok := sourceMap["roles"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"name": types.StringType,
						"url":  types.StringType,
					}
					attrValues := map[string]attr.Value{
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"name": types.StringType,
				"url":  types.StringType,
			}}, items)
			data.Roles = listVal
		}
	} else {
		if data.Roles.IsUnknown() {
			data.Roles = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"name": types.StringType,
				"url":  types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["scope"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Scope = types.StringValue(str)
		}
	} else {
		if data.Scope.IsUnknown() {
			data.Scope = types.StringNull()
		}
	}
	if val, ok := sourceMap["scope_error_message"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ScopeErrorMessage = types.StringValue(str)
		}
	} else {
		if data.ScopeErrorMessage.IsUnknown() {
			data.ScopeErrorMessage = types.StringNull()
		}
	}
	if val, ok := sourceMap["scope_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ScopeName = types.StringValue(str)
		}
	} else {
		if data.ScopeName.IsUnknown() {
			data.ScopeName = types.StringNull()
		}
	}
	if val, ok := sourceMap["scope_state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ScopeState = types.StringValue(str)
		}
	} else {
		if data.ScopeState.IsUnknown() {
			data.ScopeState = types.StringNull()
		}
	}
	if val, ok := sourceMap["scope_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ScopeUuid = types.StringValue(str)
		}
	} else {
		if data.ScopeUuid.IsUnknown() {
			data.ScopeUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["screenshots"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"created":     types.StringType,
						"description": types.StringType,
						"image":       types.StringType,
						"name":        types.StringType,
						"thumbnail":   types.StringType,
					}
					attrValues := map[string]attr.Value{
						"created": func() attr.Value {
							if v, ok := objMap["created"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"description": func() attr.Value {
							if v, ok := objMap["description"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"image": func() attr.Value {
							if v, ok := objMap["image"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"thumbnail": func() attr.Value {
							if v, ok := objMap["thumbnail"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"created":     types.StringType,
				"description": types.StringType,
				"image":       types.StringType,
				"name":        types.StringType,
				"thumbnail":   types.StringType,
			}}, items)
			data.Screenshots = listVal
		}
	} else {
		if data.Screenshots.IsUnknown() {
			data.Screenshots = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"created":     types.StringType,
				"description": types.StringType,
				"image":       types.StringType,
				"name":        types.StringType,
				"thumbnail":   types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["shared"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Shared = types.BoolValue(b)
		}
	} else {
		if data.Shared.IsUnknown() {
			data.Shared = types.BoolNull()
		}
	}
	if val, ok := sourceMap["slug"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Slug = types.StringValue(str)
		}
	} else {
		if data.Slug.IsUnknown() {
			data.Slug = types.StringNull()
		}
	}
	if val, ok := sourceMap["software_catalogs"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
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
					attrValues := map[string]attr.Value{
						"catalog": types.ObjectNull(types.ObjectType{AttrTypes: map[string]attr.Type{
							"description": types.StringType,
							"name":        types.StringType,
							"version":     types.StringType,
						}}.AttrTypes),
						"package_count": func() attr.Value {
							if v, ok := objMap["package_count"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"partition": types.ObjectNull(types.ObjectType{AttrTypes: map[string]attr.Type{
							"partition_name": types.StringType,
							"priority_tier":  types.Int64Type,
							"qos":            types.StringType,
						}}.AttrTypes),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}}, items)
			data.SoftwareCatalogs = listVal
		}
	} else {
		if data.SoftwareCatalogs.IsUnknown() {
			data.SoftwareCatalogs = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}})
		}
	}
	if val, ok := sourceMap["state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.State = types.StringValue(str)
		}
	} else {
		if data.State.IsUnknown() {
			data.State = types.StringNull()
		}
	}
	if val, ok := sourceMap["thumbnail"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Thumbnail = types.StringValue(str)
		}
	} else {
		if data.Thumbnail.IsUnknown() {
			data.Thumbnail = types.StringNull()
		}
	}
	if val, ok := sourceMap["total_cost"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.TotalCost = types.Int64Value(int64(num))
		}
	} else {
		if data.TotalCost.IsUnknown() {
			data.TotalCost = types.Int64Null()
		}
	}
	if val, ok := sourceMap["total_cost_estimated"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.TotalCostEstimated = types.Int64Value(int64(num))
		}
	} else {
		if data.TotalCostEstimated.IsUnknown() {
			data.TotalCostEstimated = types.Int64Null()
		}
	}
	if val, ok := sourceMap["total_customers"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.TotalCustomers = types.Int64Value(int64(num))
		}
	} else {
		if data.TotalCustomers.IsUnknown() {
			data.TotalCustomers = types.Int64Null()
		}
	}
	if val, ok := sourceMap["type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Type = types.StringValue(str)
		}
	} else {
		if data.Type.IsUnknown() {
			data.Type = types.StringNull()
		}
	}
	if val, ok := sourceMap["url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Url = types.StringValue(str)
		}
	} else {
		if data.Url.IsUnknown() {
			data.Url = types.StringNull()
		}
	}
	if val, ok := sourceMap["vendor_details"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.VendorDetails = types.StringValue(str)
		}
	} else {
		if data.VendorDetails.IsUnknown() {
			data.VendorDetails = types.StringNull()
		}
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MarketplaceOfferingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update Not Supported", "This resource cannot be updated via the API.")
	return
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
			"Unable to Delete MarketplaceOffering",
			"An error occurred while deleting the marketplace_offering: "+err.Error(),
		)
		return
	}
}

func (r *MarketplaceOfferingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
