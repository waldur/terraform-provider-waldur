package resource

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
var _ resource.Resource = &MarketplaceResourceResource{}
var _ resource.ResourceWithImportState = &MarketplaceResourceResource{}

func NewMarketplaceResourceResource() resource.Resource {
	return &MarketplaceResourceResource{}
}

// MarketplaceResourceResource defines the resource implementation.
type MarketplaceResourceResource struct {
	client *MarketplaceResourceClient
}

// MarketplaceResourceResourceModel describes the resource data model.
type MarketplaceResourceResourceModel struct {
	MarketplaceResourceModel
	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

func (r *MarketplaceResourceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_resource"
}

func (r *MarketplaceResourceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Marketplace Resource resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Marketplace Resource UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"available_actions": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Available actions",
			},
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "ID of the backend",
			},
			"backend_metadata": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Current action being performed",
					},
					"instance_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the backend instance",
					},
					"runtime_state": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Runtime state of the backend resource",
					},
					"state": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Backend resource state",
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Backend metadata",
			},
			"can_terminate": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Can terminate",
			},
			"category_icon": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Category icon",
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
			"creation_order": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"activation_price": schema.Float64Attribute{
						Computed: true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Activation price",
					},
					"attachment": schema.StringAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Attachment",
					},
					"backend_id": schema.StringAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "ID of the backend",
					},
					"callback_url": schema.StringAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Callback url",
					},
					"can_terminate": schema.BoolAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Can terminate",
					},
					"category_icon": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Category icon",
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
					"completed_at": schema.StringAttribute{
						CustomType: timetypes.RFC3339Type{},
						Computed:   true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Completed at",
					},
					"consumer_reviewed_at": schema.StringAttribute{
						CustomType: timetypes.RFC3339Type{},
						Computed:   true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Consumer reviewed at",
					},
					"consumer_reviewed_by": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"consumer_reviewed_by_full_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the consumer reviewed by full",
					},
					"consumer_reviewed_by_username": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"cost": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Cost",
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,12}(?:\.\d{0,10})?$`), ""),
						},
					},
					"created_by_civil_number": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Created by civil number",
					},
					"created_by_full_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the created by full",
					},
					"created_by_username": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"customer_slug": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Customer slug",
					},
					"error_message": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Error message",
					},
					"error_traceback": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Error traceback",
					},
					"fixed_price": schema.Float64Attribute{
						Computed: true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Fixed price",
					},
					"issue": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								MarkdownDescription: "Key",
							},
							"uuid": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								MarkdownDescription: "UUID of the Marketplace Resource",
							},
						},
						Computed: true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Issue",
					},
					"marketplace_resource_uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "UUID of the marketplace resource",
					},
					"new_cost_estimate": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "New cost estimate",
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,12}(?:\.\d{0,10})?$`), ""),
						},
					},
					"new_plan_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the new plan",
					},
					"new_plan_uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "UUID of the new plan",
					},
					"offering": schema.StringAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Offering",
					},
					"offering_billable": schema.BoolAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Purchase and usage is invoiced.",
					},
					"offering_description": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Offering description",
					},
					"offering_image": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Offering image",
					},
					"offering_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the offering",
					},
					"offering_shared": schema.BoolAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Accessible to all customers.",
					},
					"offering_thumbnail": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Offering thumbnail",
					},
					"offering_type": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Offering type",
					},
					"offering_uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "UUID of the offering",
					},
					"old_cost_estimate": schema.Float64Attribute{
						Computed: true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Old cost estimate",
					},
					"old_plan_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the old plan",
					},
					"old_plan_uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "UUID of the old plan",
					},
					"order_subtype": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Order subtype",
					},
					"output": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Output",
					},
					"plan": schema.StringAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Plan",
					},
					"plan_description": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Plan description",
					},
					"plan_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the plan",
					},
					"plan_unit": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Plan unit",
					},
					"plan_uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "UUID of the plan",
					},
					"project_description": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Project description",
					},
					"project_slug": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Project slug",
					},
					"provider_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the provider",
					},
					"provider_reviewed_at": schema.StringAttribute{
						CustomType: timetypes.RFC3339Type{},
						Computed:   true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Provider reviewed at",
					},
					"provider_reviewed_by": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"provider_reviewed_by_full_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the provider reviewed by full",
					},
					"provider_reviewed_by_username": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"provider_slug": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Provider slug",
					},
					"provider_uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "UUID of the provider",
					},
					"request_comment": schema.StringAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Request comment",
					},
					"resource_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the resource",
					},
					"resource_type": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Resource type",
					},
					"resource_uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "UUID of the resource",
					},
					"slug": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Slug",
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
						},
					},
					"start_date": schema.StringAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Enables delayed processing of resource provisioning order.",
					},
					"state": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "State",
					},
					"termination_comment": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Termination comment",
					},
					"type": schema.StringAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
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
					"uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "UUID of the Marketplace Resource",
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Creation order",
			},
			"customer_slug": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Customer slug",
			},
			"description": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Description of the Marketplace Resource",
			},
			"downscaled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Downscaled",
			},
			"effective_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "ID of the effective",
			},
			"end_date": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The date is inclusive. Once reached, a resource will be scheduled for termination.",
			},
			"end_date_requested_by": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "End date requested by",
			},
			"endpoints": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Name of the Marketplace Resource",
						},
						"url": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "URL of the access endpoint",
						},
						"uuid": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "UUID of the Marketplace Resource",
						},
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Endpoints",
			},
			"error_message": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Error message",
			},
			"error_traceback": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Error traceback",
			},
			"last_sync": schema.StringAttribute{
				CustomType: timetypes.RFC3339Type{},
				Computed:   true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Last sync",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the Marketplace Resource",
			},
			"offering": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Offering",
			},
			"offering_backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "ID of the offering backend",
			},
			"offering_billable": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Purchase and usage is invoiced.",
			},
			"offering_components": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"article_code": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Article code",
						},
						"billing_type": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Billing type",
							Validators: []validator.String{
								stringvalidator.OneOf("fixed", "usage", "limit", "one", "few"),
							},
						},
						"default_limit": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Default limit",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"description": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Description of the Marketplace Resource",
						},
						"factor": schema.Int64Attribute{
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Factor",
						},
						"is_boolean": schema.BoolAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Is boolean",
						},
						"is_builtin": schema.BoolAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Is builtin",
						},
						"is_prepaid": schema.BoolAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Is prepaid",
						},
						"limit_amount": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Limit amount",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"limit_period": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Limit period",
						},
						"max_available_limit": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Max available limit",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"max_prepaid_duration": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Max prepaid duration",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"max_value": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Max value",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"measured_unit": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Unit of measurement, for example, GB.",
						},
						"min_prepaid_duration": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Min prepaid duration",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"min_value": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Min value",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"name": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Display name for the measured unit, for example, Floating IP.",
						},
						"overage_component": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Overage component",
						},
						"type": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Unique internal name of the measured unit, for example floating_ip.",
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_\-\/:]+$`), ""),
							},
						},
						"unit_factor": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "The conversion factor from backend units to measured_unit",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"uuid": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "UUID of the Marketplace Resource",
						},
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Offering components",
			},
			"offering_description": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Offering description",
			},
			"offering_image": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Offering image",
			},
			"offering_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the offering",
			},
			"offering_shared": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Accessible to all customers.",
			},
			"offering_slug": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Offering slug",
			},
			"offering_state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Offering state",
			},
			"offering_thumbnail": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Offering thumbnail",
			},
			"offering_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Offering type",
			},
			"offering_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the offering",
			},
			"order_in_progress": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"activation_price": schema.Float64Attribute{
						Computed: true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Activation price",
					},
					"attachment": schema.StringAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Attachment",
					},
					"backend_id": schema.StringAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "ID of the backend",
					},
					"callback_url": schema.StringAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Callback url",
					},
					"can_terminate": schema.BoolAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Can terminate",
					},
					"category_icon": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Category icon",
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
					"completed_at": schema.StringAttribute{
						CustomType: timetypes.RFC3339Type{},
						Computed:   true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Completed at",
					},
					"consumer_reviewed_at": schema.StringAttribute{
						CustomType: timetypes.RFC3339Type{},
						Computed:   true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Consumer reviewed at",
					},
					"consumer_reviewed_by": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"consumer_reviewed_by_full_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the consumer reviewed by full",
					},
					"consumer_reviewed_by_username": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"cost": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Cost",
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,12}(?:\.\d{0,10})?$`), ""),
						},
					},
					"created_by_civil_number": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Created by civil number",
					},
					"created_by_full_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the created by full",
					},
					"created_by_username": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"customer_slug": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Customer slug",
					},
					"error_message": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Error message",
					},
					"error_traceback": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Error traceback",
					},
					"fixed_price": schema.Float64Attribute{
						Computed: true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Fixed price",
					},
					"issue": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								MarkdownDescription: "Key",
							},
							"uuid": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								MarkdownDescription: "UUID of the Marketplace Resource",
							},
						},
						Computed: true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Issue",
					},
					"marketplace_resource_uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "UUID of the marketplace resource",
					},
					"new_cost_estimate": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "New cost estimate",
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,12}(?:\.\d{0,10})?$`), ""),
						},
					},
					"new_plan_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the new plan",
					},
					"new_plan_uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "UUID of the new plan",
					},
					"offering": schema.StringAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Offering",
					},
					"offering_billable": schema.BoolAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Purchase and usage is invoiced.",
					},
					"offering_description": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Offering description",
					},
					"offering_image": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Offering image",
					},
					"offering_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the offering",
					},
					"offering_shared": schema.BoolAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Accessible to all customers.",
					},
					"offering_thumbnail": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Offering thumbnail",
					},
					"offering_type": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Offering type",
					},
					"offering_uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "UUID of the offering",
					},
					"old_cost_estimate": schema.Float64Attribute{
						Computed: true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Old cost estimate",
					},
					"old_plan_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the old plan",
					},
					"old_plan_uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "UUID of the old plan",
					},
					"order_subtype": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Order subtype",
					},
					"output": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Output",
					},
					"plan": schema.StringAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Plan",
					},
					"plan_description": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Plan description",
					},
					"plan_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the plan",
					},
					"plan_unit": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Plan unit",
					},
					"plan_uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "UUID of the plan",
					},
					"project_description": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Project description",
					},
					"project_slug": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Project slug",
					},
					"provider_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the provider",
					},
					"provider_reviewed_at": schema.StringAttribute{
						CustomType: timetypes.RFC3339Type{},
						Computed:   true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Provider reviewed at",
					},
					"provider_reviewed_by": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"provider_reviewed_by_full_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the provider reviewed by full",
					},
					"provider_reviewed_by_username": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"provider_slug": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Provider slug",
					},
					"provider_uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "UUID of the provider",
					},
					"request_comment": schema.StringAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Request comment",
					},
					"resource_name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name of the resource",
					},
					"resource_type": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Resource type",
					},
					"resource_uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "UUID of the resource",
					},
					"slug": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Slug",
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
						},
					},
					"start_date": schema.StringAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Enables delayed processing of resource provisioning order.",
					},
					"state": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "State",
					},
					"termination_comment": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Termination comment",
					},
					"type": schema.StringAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
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
					"uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "UUID of the Marketplace Resource",
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Order in progress",
			},
			"parent_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the parent",
			},
			"parent_offering_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the parent offering",
			},
			"parent_offering_slug": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Parent offering slug",
			},
			"parent_offering_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the parent offering",
			},
			"parent_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the parent",
			},
			"paused": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Paused",
			},
			"plan": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Plan",
			},
			"plan_description": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Plan description",
			},
			"plan_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the plan",
			},
			"plan_unit": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Plan unit",
			},
			"plan_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the plan",
			},
			"project": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Project",
			},
			"project_description": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Project description",
			},
			"project_end_date": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The date is inclusive. Once reached, all project resource will be scheduled for termination.",
			},
			"project_end_date_requested_by": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Project end date requested by",
			},
			"project_slug": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Project slug",
			},
			"provider_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the provider",
			},
			"provider_slug": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Provider slug",
			},
			"provider_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the provider",
			},
			"report": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"body": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Section body content",
						},
						"header": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Section header text",
						},
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Report",
			},
			"resource_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Resource type",
			},
			"resource_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the resource",
			},
			"restrict_member_access": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Restrict member access",
			},
			"scope": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Scope",
			},
			"slug": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
				},
			},
			"state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "State",
			},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Url",
			},
			"user_requires_reconsent": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Check if the current user needs to re-consent for this resource's offering.",
			},
			"username": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Username",
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

func (r *MarketplaceResourceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.client = &MarketplaceResourceClient{}
	if err := r.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			err.Error(),
		)
		return
	}
}

func (r *MarketplaceResourceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	resp.Diagnostics.AddError("Creation Not Supported", "This resource cannot be created via the API.")
}

func (r *MarketplaceResourceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MarketplaceResourceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource

	apiResp, err := r.client.Get(ctx, data.UUID.ValueString())
	if err != nil {
		if IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Marketplace Resource",
			"An error occurred while reading the Marketplace Resource: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MarketplaceResourceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var data MarketplaceResourceResourceModel
	var state MarketplaceResourceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := MarketplaceResourceUpdateRequest{}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {

		requestBody.Description = data.Description.ValueStringPointer()
	}
	if !data.EndDate.IsNull() && !data.EndDate.IsUnknown() {

		requestBody.EndDate = data.EndDate.ValueStringPointer()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {

		requestBody.Name = data.Name.ValueStringPointer()
	}

	apiResp, err := r.client.Update(ctx, data.UUID.ValueString(), &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Marketplace Resource",
			"An error occurred while updating the Marketplace Resource: "+err.Error(),
		)
		return
	}
	updateTimeout, diags := data.Timeouts.Update(ctx, common.DefaultUpdateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*MarketplaceResourceResponse, error) {
		return r.client.Get(ctx, data.UUID.ValueString())
	}, updateTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource update", err.Error())
		return
	}
	apiResp = newResp

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MarketplaceResourceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	resp.Diagnostics.AddError("Deletion Not Supported", "This resource cannot be deleted via the API.")
}

func (r *MarketplaceResourceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	uuid := req.ID
	if uuid == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Import ID cannot be empty. Please provide the UUID of the Marketplace Resource.",
		)
		return
	}

	tflog.Info(ctx, "Importing Marketplace Resource", map[string]interface{}{
		"uuid": uuid,
	})

	apiResp, err := r.client.Get(ctx, uuid)
	if err != nil {
		if IsNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("Marketplace Resource with UUID '%s' does not exist or is not accessible.", uuid),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to Import Marketplace Resource",
			fmt.Sprintf("An error occurred while fetching the Marketplace Resource: %s", err.Error()),
		)
		return
	}

	var data MarketplaceResourceResourceModel
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
