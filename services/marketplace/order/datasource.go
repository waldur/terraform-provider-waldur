package order

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &MarketplaceOrderDataSource{}

func NewMarketplaceOrderDataSource() datasource.DataSource {
	return &MarketplaceOrderDataSource{}
}

type MarketplaceOrderDataSource struct {
	client *MarketplaceOrderClient
}

type MarketplaceOrderDataSourceModel struct {
	MarketplaceOrderModel
	Filters *MarketplaceOrderFiltersModel `tfsdk:"filters"`
}

func (d *MarketplaceOrderDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_order"
}

func (d *MarketplaceOrderDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Marketplace Order data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Marketplace Order UUID",
			},
			"filters": (&MarketplaceOrderFiltersModel{}).GetSchema(),
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
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",

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
		},
	}
}

func (d *MarketplaceOrderDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	d.client = &MarketplaceOrderClient{}
	if err := d.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			err.Error(),
		)
		return
	}
}

func (d *MarketplaceOrderDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data MarketplaceOrderDataSourceModel

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
				"Unable to Read Marketplace Order",
				"An error occurred while reading the Marketplace Order by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup marketplace_order.",
			)
			return
		}

		results, err := d.client.List(ctx, filters)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Marketplace Order",
				"An error occurred while filtering Marketplace Order: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Marketplace Order Not Found",
				"No Marketplace Order found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Marketplace Orders Found",
				fmt.Sprintf("Found %d Marketplace Orders with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, results[0])...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
