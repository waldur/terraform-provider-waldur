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
				MarkdownDescription: "Available actions",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of the backend",
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
				MarkdownDescription: "Backend metadata",
			},
			"can_terminate": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Can terminate",
			},
			"category_icon": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Category icon",
			},
			"category_title": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Category title",
			},
			"category_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the category",
			},
			"created": schema.StringAttribute{
				CustomType:          timetypes.RFC3339Type{},
				Computed:            true,
				MarkdownDescription: "Created",
			},
			"creation_order": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"activation_price": schema.Float64Attribute{
						Computed:            true,
						MarkdownDescription: "Activation price",
					},
					"attachment": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Attachment",
					},
					"backend_id": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "ID of the backend",
					},
					"callback_url": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Callback url",
					},
					"can_terminate": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "Can terminate",
					},
					"category_icon": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Category icon",
					},
					"category_title": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Category title",
					},
					"category_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "UUID of the category",
					},
					"completed_at": schema.StringAttribute{
						CustomType:          timetypes.RFC3339Type{},
						Computed:            true,
						MarkdownDescription: "Completed at",
					},
					"consumer_reviewed_at": schema.StringAttribute{
						CustomType:          timetypes.RFC3339Type{},
						Computed:            true,
						MarkdownDescription: "Consumer reviewed at",
					},
					"consumer_reviewed_by": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"consumer_reviewed_by_full_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the consumer reviewed by full",
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
					"created": schema.StringAttribute{
						CustomType:          timetypes.RFC3339Type{},
						Computed:            true,
						MarkdownDescription: "Created",
					},
					"created_by_civil_number": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Created by civil number",
					},
					"created_by_full_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the created by full",
					},
					"created_by_username": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"customer_slug": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Customer slug",
					},
					"error_message": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Error message",
					},
					"error_traceback": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Error traceback",
					},
					"fixed_price": schema.Float64Attribute{
						Computed:            true,
						MarkdownDescription: "Fixed price",
					},
					"issue": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Computed:            true,
								MarkdownDescription: "Key",
							},
							"uuid": schema.StringAttribute{
								Computed:            true,
								MarkdownDescription: "UUID of the Marketplace Resource",
							},
						},
						Computed:            true,
						MarkdownDescription: "Issue",
					},
					"marketplace_resource_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "UUID of the marketplace resource",
					},
					"modified": schema.StringAttribute{
						CustomType:          timetypes.RFC3339Type{},
						Computed:            true,
						MarkdownDescription: "Modified",
					},
					"new_cost_estimate": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "New cost estimate",
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,12}(?:\.\d{0,10})?$`), ""),
						},
					},
					"new_plan_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the new plan",
					},
					"new_plan_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "UUID of the new plan",
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
						MarkdownDescription: "Offering description",
					},
					"offering_image": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering image",
					},
					"offering_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the offering",
					},
					"offering_shared": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "Accessible to all customers.",
					},
					"offering_thumbnail": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering thumbnail",
					},
					"offering_type": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering type",
					},
					"offering_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "UUID of the offering",
					},
					"old_cost_estimate": schema.Float64Attribute{
						Computed:            true,
						MarkdownDescription: "Old cost estimate",
					},
					"old_plan_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the old plan",
					},
					"old_plan_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "UUID of the old plan",
					},
					"order_subtype": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Order subtype",
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
						MarkdownDescription: "Plan description",
					},
					"plan_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the plan",
					},
					"plan_unit": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Plan unit",
					},
					"plan_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "UUID of the plan",
					},
					"project_description": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Project description",
					},
					"project_slug": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Project slug",
					},
					"provider_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the provider",
					},
					"provider_reviewed_at": schema.StringAttribute{
						CustomType:          timetypes.RFC3339Type{},
						Computed:            true,
						MarkdownDescription: "Provider reviewed at",
					},
					"provider_reviewed_by": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"provider_reviewed_by_full_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the provider reviewed by full",
					},
					"provider_reviewed_by_username": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"provider_slug": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Provider slug",
					},
					"provider_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "UUID of the provider",
					},
					"request_comment": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Request comment",
					},
					"resource_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the resource",
					},
					"resource_type": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Resource type",
					},
					"resource_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "UUID of the resource",
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
						MarkdownDescription: "Termination comment",
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
						MarkdownDescription: "UUID of the Marketplace Resource",
					},
				},
				Computed:            true,
				MarkdownDescription: "Creation order",
			},
			"customer_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Customer slug",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Description of the Marketplace Resource",
			},
			"downscaled": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Downscaled",
			},
			"effective_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of the effective",
			},
			"end_date": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The date is inclusive. Once reached, a resource will be scheduled for termination.",
			},
			"end_date_requested_by": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "End date requested by",
			},
			"endpoints": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name of the Marketplace Resource",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "URL of the access endpoint",
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "UUID of the Marketplace Resource",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Endpoints",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error message",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error traceback",
			},
			"last_sync": schema.StringAttribute{
				CustomType:          timetypes.RFC3339Type{},
				Computed:            true,
				MarkdownDescription: "Last sync",
			},
			"modified": schema.StringAttribute{
				CustomType:          timetypes.RFC3339Type{},
				Computed:            true,
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the Marketplace Resource",
			},
			"offering": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering",
			},
			"offering_backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of the offering backend",
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
							MarkdownDescription: "Article code",
						},
						"billing_type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Billing type",
							Validators: []validator.String{
								stringvalidator.OneOf("fixed", "usage", "limit", "one", "few"),
							},
						},
						"default_limit": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Default limit",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Description of the Marketplace Resource",
						},
						"factor": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Factor",
						},
						"is_boolean": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Is boolean",
						},
						"is_builtin": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Is builtin",
						},
						"is_prepaid": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Is prepaid",
						},
						"limit_amount": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Limit amount",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"limit_period": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Limit period",
						},
						"max_available_limit": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Max available limit",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"max_prepaid_duration": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Max prepaid duration",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"max_value": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Max value",
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
							MarkdownDescription: "Min prepaid duration",
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(2147483647),
							},
						},
						"min_value": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Min value",
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
							MarkdownDescription: "Overage component",
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
							MarkdownDescription: "UUID of the Marketplace Resource",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Offering components",
			},
			"offering_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering description",
			},
			"offering_image": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering image",
			},
			"offering_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the offering",
			},
			"offering_shared": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Accessible to all customers.",
			},
			"offering_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering slug",
			},
			"offering_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering state",
			},
			"offering_thumbnail": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering thumbnail",
			},
			"offering_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering type",
			},
			"offering_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the offering",
			},
			"order_in_progress": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"activation_price": schema.Float64Attribute{
						Computed:            true,
						MarkdownDescription: "Activation price",
					},
					"attachment": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Attachment",
					},
					"backend_id": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "ID of the backend",
					},
					"callback_url": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Callback url",
					},
					"can_terminate": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "Can terminate",
					},
					"category_icon": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Category icon",
					},
					"category_title": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Category title",
					},
					"category_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "UUID of the category",
					},
					"completed_at": schema.StringAttribute{
						CustomType:          timetypes.RFC3339Type{},
						Computed:            true,
						MarkdownDescription: "Completed at",
					},
					"consumer_reviewed_at": schema.StringAttribute{
						CustomType:          timetypes.RFC3339Type{},
						Computed:            true,
						MarkdownDescription: "Consumer reviewed at",
					},
					"consumer_reviewed_by": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"consumer_reviewed_by_full_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the consumer reviewed by full",
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
					"created": schema.StringAttribute{
						CustomType:          timetypes.RFC3339Type{},
						Computed:            true,
						MarkdownDescription: "Created",
					},
					"created_by_civil_number": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Created by civil number",
					},
					"created_by_full_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the created by full",
					},
					"created_by_username": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"customer_slug": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Customer slug",
					},
					"error_message": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Error message",
					},
					"error_traceback": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Error traceback",
					},
					"fixed_price": schema.Float64Attribute{
						Computed:            true,
						MarkdownDescription: "Fixed price",
					},
					"issue": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Computed:            true,
								MarkdownDescription: "Key",
							},
							"uuid": schema.StringAttribute{
								Computed:            true,
								MarkdownDescription: "UUID of the Marketplace Resource",
							},
						},
						Computed:            true,
						MarkdownDescription: "Issue",
					},
					"marketplace_resource_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "UUID of the marketplace resource",
					},
					"modified": schema.StringAttribute{
						CustomType:          timetypes.RFC3339Type{},
						Computed:            true,
						MarkdownDescription: "Modified",
					},
					"new_cost_estimate": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "New cost estimate",
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,12}(?:\.\d{0,10})?$`), ""),
						},
					},
					"new_plan_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the new plan",
					},
					"new_plan_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "UUID of the new plan",
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
						MarkdownDescription: "Offering description",
					},
					"offering_image": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering image",
					},
					"offering_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the offering",
					},
					"offering_shared": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: "Accessible to all customers.",
					},
					"offering_thumbnail": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering thumbnail",
					},
					"offering_type": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Offering type",
					},
					"offering_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "UUID of the offering",
					},
					"old_cost_estimate": schema.Float64Attribute{
						Computed:            true,
						MarkdownDescription: "Old cost estimate",
					},
					"old_plan_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the old plan",
					},
					"old_plan_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "UUID of the old plan",
					},
					"order_subtype": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Order subtype",
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
						MarkdownDescription: "Plan description",
					},
					"plan_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the plan",
					},
					"plan_unit": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Plan unit",
					},
					"plan_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "UUID of the plan",
					},
					"project_description": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Project description",
					},
					"project_slug": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Project slug",
					},
					"provider_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the provider",
					},
					"provider_reviewed_at": schema.StringAttribute{
						CustomType:          timetypes.RFC3339Type{},
						Computed:            true,
						MarkdownDescription: "Provider reviewed at",
					},
					"provider_reviewed_by": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"provider_reviewed_by_full_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the provider reviewed by full",
					},
					"provider_reviewed_by_username": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
					},
					"provider_slug": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Provider slug",
					},
					"provider_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "UUID of the provider",
					},
					"request_comment": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Request comment",
					},
					"resource_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the resource",
					},
					"resource_type": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Resource type",
					},
					"resource_uuid": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "UUID of the resource",
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
						MarkdownDescription: "Termination comment",
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
						MarkdownDescription: "UUID of the Marketplace Resource",
					},
				},
				Computed:            true,
				MarkdownDescription: "Order in progress",
			},
			"parent_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the parent",
			},
			"parent_offering_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the parent offering",
			},
			"parent_offering_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Parent offering slug",
			},
			"parent_offering_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the parent offering",
			},
			"parent_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the parent",
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
				MarkdownDescription: "Plan description",
			},
			"plan_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the plan",
			},
			"plan_unit": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Plan unit",
			},
			"plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the plan",
			},
			"project": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project",
			},
			"project_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project description",
			},
			"project_end_date": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The date is inclusive. Once reached, all project resource will be scheduled for termination.",
			},
			"project_end_date_requested_by": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project end date requested by",
			},
			"project_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project slug",
			},
			"provider_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the provider",
			},
			"provider_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Provider slug",
			},
			"provider_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the provider",
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
				MarkdownDescription: "Resource type",
			},
			"resource_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the resource",
			},
			"restrict_member_access": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Restrict member access",
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
