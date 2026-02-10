package resource

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
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &MarketplaceResourceDataSource{}

func NewMarketplaceResourceDataSource() datasource.DataSource {
	return &MarketplaceResourceDataSource{}
}

type MarketplaceResourceDataSource struct {
	client *MarketplaceResourceClient
}

type MarketplaceResourceDataSourceModel struct {
	MarketplaceResourceModel
	Filters *MarketplaceResourceFiltersModel `tfsdk:"filters"`
}

func (d *MarketplaceResourceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_resource"
}

func (d *MarketplaceResourceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Marketplace Resource data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Marketplace Resource UUID",
			},
			"filters": (&MarketplaceResourceFiltersModel{}).GetSchema(),
			"available_actions": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "Available Actions",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Backend Id",
			},
			"backend_metadata": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Current action being performed",
					},
					"instance_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the backend instance",
					},
					"runtime_state": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Runtime state of the backend resource",
					},
					"state": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Backend resource state",
					},
				},
				Computed:            true,
				MarkdownDescription: "Backend Metadata",
			},
			"can_terminate": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Can Terminate",
			},
			"category_icon": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Category Icon",
			},
			"category_title": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Category Title",
			},
			"category_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Category Uuid",
			},
			"creation_order": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"activation_price": schema.Float64Attribute{
						Computed:            true,
						MarkdownDescription: "Activation Price",
					},
					"attachment": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Attachment",
					},
					"backend_id": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Backend Id",
					},
					"callback_url": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Callback Url",
					},
					"can_terminate": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "Can Terminate",
					},
					"category_icon": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Category Icon",
					},
					"category_title": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Category Title",
					},
					"category_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Category Uuid",
					},
					"completed_at": schema.StringAttribute{
						CustomType:          timetypes.RFC3339Type{},
						Computed:            true,
						MarkdownDescription: "Completed At",
					},
					"consumer_reviewed_at": schema.StringAttribute{
						CustomType:          timetypes.RFC3339Type{},
						Computed:            true,
						MarkdownDescription: "Consumer Reviewed At",
					},
					"consumer_reviewed_by": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"consumer_reviewed_by_full_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Consumer Reviewed By Full Name",
					},
					"consumer_reviewed_by_username": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"cost": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Cost",
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,12}(?:\.\d{0,10})?$`), ""),
						},
					},
					"created_by_civil_number": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Created By Civil Number",
					},
					"created_by_full_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Created By Full Name",
					},
					"created_by_username": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"customer_slug": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Customer Slug",
					},
					"error_message": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Error Message",
					},
					"fixed_price": schema.Float64Attribute{
						Computed:            true,
						MarkdownDescription: "Fixed Price",
					},
					"issue": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Computed:            true,
								MarkdownDescription: "Key",
							},
							"uuid": schema.StringAttribute{
								Computed:            true,
								MarkdownDescription: "Uuid",
							},
						},
						Computed:            true,
						MarkdownDescription: "Issue",
					},
					"marketplace_resource_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Marketplace Resource Uuid",
					},
					"new_cost_estimate": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "New Cost Estimate",
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,12}(?:\.\d{0,10})?$`), ""),
						},
					},
					"new_plan_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "New Plan Name",
					},
					"new_plan_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "New Plan Uuid",
					},
					"offering": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering",
					},
					"offering_billable": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "Purchase and usage is invoiced.",
					},
					"offering_description": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering Description",
					},
					"offering_image": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering Image",
					},
					"offering_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering Name",
					},
					"offering_shared": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "Accessible to all customers.",
					},
					"offering_thumbnail": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering Thumbnail",
					},
					"offering_type": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering Type",
					},
					"offering_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering Uuid",
					},
					"old_cost_estimate": schema.Float64Attribute{
						Computed:            true,
						MarkdownDescription: "Old Cost Estimate",
					},
					"old_plan_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Old Plan Name",
					},
					"old_plan_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Old Plan Uuid",
					},
					"order_subtype": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Order Subtype",
					},
					"output": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Output",
					},
					"plan": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Plan",
					},
					"plan_description": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Plan Description",
					},
					"plan_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Plan Name",
					},
					"plan_unit": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Plan Unit",
					},
					"plan_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Plan Uuid",
					},
					"project_description": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Project Description",
					},
					"project_slug": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Project Slug",
					},
					"provider_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Provider Name",
					},
					"provider_reviewed_at": schema.StringAttribute{
						CustomType:          timetypes.RFC3339Type{},
						Computed:            true,
						MarkdownDescription: "Provider Reviewed At",
					},
					"provider_reviewed_by": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"provider_reviewed_by_full_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Provider Reviewed By Full Name",
					},
					"provider_reviewed_by_username": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"provider_slug": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Provider Slug",
					},
					"provider_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Provider Uuid",
					},
					"request_comment": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Request Comment",
					},
					"resource_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Resource Name",
					},
					"resource_type": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Resource Type",
					},
					"resource_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Resource Uuid",
					},
					"slug": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Slug",
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
						},
					},
					"start_date": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Enables delayed processing of resource provisioning order.",
					},
					"state": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "State",
					},
					"termination_comment": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Termination Comment",
					},
					"type": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Type",
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
				Computed:            true,
				MarkdownDescription: "Creation Order",
			},
			"customer_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Customer Slug",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Description",
			},
			"downscaled": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Downscaled",
			},
			"effective_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Effective Id",
			},
			"end_date": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The date is inclusive. Once reached, a resource will be scheduled for termination.",
			},
			"end_date_requested_by": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "End Date Requested By",
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
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error Message",
			},
			"last_sync": schema.StringAttribute{
				CustomType:          timetypes.RFC3339Type{},
				Computed:            true,
				MarkdownDescription: "Last Sync",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name",
			},
			"offering": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering",
			},
			"offering_backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering Backend Id",
			},
			"offering_billable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Purchase and usage is invoiced.",
			},
			"offering_components": schema.ListNestedAttribute{
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
				MarkdownDescription: "Offering Components",
			},
			"offering_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering Description",
			},
			"offering_image": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering Image",
			},
			"offering_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering Name",
			},
			"offering_shared": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Accessible to all customers.",
			},
			"offering_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering Slug",
			},
			"offering_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering State",
			},
			"offering_thumbnail": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering Thumbnail",
			},
			"offering_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering Type",
			},
			"offering_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering Uuid",
			},
			"order_in_progress": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"activation_price": schema.Float64Attribute{
						Computed:            true,
						MarkdownDescription: "Activation Price",
					},
					"attachment": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Attachment",
					},
					"backend_id": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Backend Id",
					},
					"callback_url": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Callback Url",
					},
					"can_terminate": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "Can Terminate",
					},
					"category_icon": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Category Icon",
					},
					"category_title": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Category Title",
					},
					"category_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Category Uuid",
					},
					"completed_at": schema.StringAttribute{
						CustomType:          timetypes.RFC3339Type{},
						Computed:            true,
						MarkdownDescription: "Completed At",
					},
					"consumer_reviewed_at": schema.StringAttribute{
						CustomType:          timetypes.RFC3339Type{},
						Computed:            true,
						MarkdownDescription: "Consumer Reviewed At",
					},
					"consumer_reviewed_by": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"consumer_reviewed_by_full_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Consumer Reviewed By Full Name",
					},
					"consumer_reviewed_by_username": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"cost": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Cost",
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,12}(?:\.\d{0,10})?$`), ""),
						},
					},
					"created_by_civil_number": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Created By Civil Number",
					},
					"created_by_full_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Created By Full Name",
					},
					"created_by_username": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"customer_slug": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Customer Slug",
					},
					"error_message": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Error Message",
					},
					"fixed_price": schema.Float64Attribute{
						Computed:            true,
						MarkdownDescription: "Fixed Price",
					},
					"issue": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Computed:            true,
								MarkdownDescription: "Key",
							},
							"uuid": schema.StringAttribute{
								Computed:            true,
								MarkdownDescription: "Uuid",
							},
						},
						Computed:            true,
						MarkdownDescription: "Issue",
					},
					"marketplace_resource_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Marketplace Resource Uuid",
					},
					"new_cost_estimate": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "New Cost Estimate",
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,12}(?:\.\d{0,10})?$`), ""),
						},
					},
					"new_plan_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "New Plan Name",
					},
					"new_plan_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "New Plan Uuid",
					},
					"offering": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering",
					},
					"offering_billable": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "Purchase and usage is invoiced.",
					},
					"offering_description": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering Description",
					},
					"offering_image": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering Image",
					},
					"offering_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering Name",
					},
					"offering_shared": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "Accessible to all customers.",
					},
					"offering_thumbnail": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering Thumbnail",
					},
					"offering_type": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering Type",
					},
					"offering_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering Uuid",
					},
					"old_cost_estimate": schema.Float64Attribute{
						Computed:            true,
						MarkdownDescription: "Old Cost Estimate",
					},
					"old_plan_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Old Plan Name",
					},
					"old_plan_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Old Plan Uuid",
					},
					"order_subtype": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Order Subtype",
					},
					"output": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Output",
					},
					"plan": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Plan",
					},
					"plan_description": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Plan Description",
					},
					"plan_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Plan Name",
					},
					"plan_unit": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Plan Unit",
					},
					"plan_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Plan Uuid",
					},
					"project_description": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Project Description",
					},
					"project_slug": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Project Slug",
					},
					"provider_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Provider Name",
					},
					"provider_reviewed_at": schema.StringAttribute{
						CustomType:          timetypes.RFC3339Type{},
						Computed:            true,
						MarkdownDescription: "Provider Reviewed At",
					},
					"provider_reviewed_by": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"provider_reviewed_by_full_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Provider Reviewed By Full Name",
					},
					"provider_reviewed_by_username": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"provider_slug": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Provider Slug",
					},
					"provider_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Provider Uuid",
					},
					"request_comment": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Request Comment",
					},
					"resource_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Resource Name",
					},
					"resource_type": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Resource Type",
					},
					"resource_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Resource Uuid",
					},
					"slug": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Slug",
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
						},
					},
					"start_date": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Enables delayed processing of resource provisioning order.",
					},
					"state": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "State",
					},
					"termination_comment": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Termination Comment",
					},
					"type": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Type",
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
				Computed:            true,
				MarkdownDescription: "Order In Progress",
			},
			"parent_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Parent Name",
			},
			"parent_offering_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Parent Offering Name",
			},
			"parent_offering_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Parent Offering Slug",
			},
			"parent_offering_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Parent Offering Uuid",
			},
			"parent_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Parent Uuid",
			},
			"paused": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Paused",
			},
			"plan": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Plan",
			},
			"plan_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Plan Description",
			},
			"plan_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Plan Name",
			},
			"plan_unit": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Plan Unit",
			},
			"plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Plan Uuid",
			},
			"project": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project",
			},
			"project_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project Description",
			},
			"project_end_date": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The date is inclusive. Once reached, all project resource will be scheduled for termination.",
			},
			"project_end_date_requested_by": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project End Date Requested By",
			},
			"project_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project Slug",
			},
			"provider_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Provider Name",
			},
			"provider_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Provider Slug",
			},
			"provider_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Provider Uuid",
			},
			"report": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"body": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Section body content",
						},
						"header": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Section header text",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Report",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource Type",
			},
			"resource_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource Uuid",
			},
			"restrict_member_access": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Restrict Member Access",
			},
			"scope": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Scope",
			},
			"slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
				},
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "State",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Url",
			},
			"user_requires_reconsent": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Check if the current user needs to re-consent for this resource's offering.",
			},
			"username": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Username",
			},
		},
	}
}

func (d *MarketplaceResourceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	d.client = &MarketplaceResourceClient{}
	if err := d.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			err.Error(),
		)
		return
	}
}

func (d *MarketplaceResourceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data MarketplaceResourceDataSourceModel

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
				"Unable to Read Marketplace Resource",
				"An error occurred while reading the Marketplace Resource by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup marketplace_resource.",
			)
			return
		}

		results, err := d.client.List(ctx, filters)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Marketplace Resource",
				"An error occurred while filtering Marketplace Resource: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Marketplace Resource Not Found",
				"No Marketplace Resource found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Marketplace Resources Found",
				fmt.Sprintf("Found %d Marketplace Resources with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, results[0])...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
