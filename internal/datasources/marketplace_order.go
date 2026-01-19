package datasources

import (
	"context"
	"fmt"

	_ "github.com/hashicorp/terraform-plugin-framework/attr" // Used for object types
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-waldur-provider/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &MarketplaceOrderDataSource{}

func NewMarketplaceOrderDataSource() datasource.DataSource {
	return &MarketplaceOrderDataSource{}
}

// MarketplaceOrderDataSource defines the data source implementation.
type MarketplaceOrderDataSource struct {
	client *client.Client
}

// MarketplaceOrderDataSourceModel describes the data source data model.
type MarketplaceOrderDataSourceModel struct {
	UUID                       types.String  `tfsdk:"id"`
	CanApproveAsConsumer       types.Bool    `tfsdk:"can_approve_as_consumer"`
	CanApproveAsProvider       types.Bool    `tfsdk:"can_approve_as_provider"`
	CategoryUuid               types.String  `tfsdk:"category_uuid"`
	Created                    types.String  `tfsdk:"created"`
	CustomerUuid               types.String  `tfsdk:"customer_uuid"`
	Modified                   types.String  `tfsdk:"modified"`
	Offering                   types.String  `tfsdk:"offering"`
	OfferingSlug               types.String  `tfsdk:"offering_slug"`
	OfferingType               types.String  `tfsdk:"offering_type"`
	OfferingUuid               types.String  `tfsdk:"offering_uuid"`
	ParentOfferingUuid         types.String  `tfsdk:"parent_offering_uuid"`
	ProjectUuid                types.String  `tfsdk:"project_uuid"`
	ProviderUuid               types.String  `tfsdk:"provider_uuid"`
	Query                      types.String  `tfsdk:"query"`
	Resource                   types.String  `tfsdk:"resource"`
	ResourceUuid               types.String  `tfsdk:"resource_uuid"`
	ServiceManagerUuid         types.String  `tfsdk:"service_manager_uuid"`
	State                      types.String  `tfsdk:"state"`
	Type                       types.String  `tfsdk:"type"`
	ActivationPrice            types.Float64 `tfsdk:"activation_price"`
	Attachment                 types.String  `tfsdk:"attachment"`
	BackendId                  types.String  `tfsdk:"backend_id"`
	CallbackUrl                types.String  `tfsdk:"callback_url"`
	CanTerminate               types.Bool    `tfsdk:"can_terminate"`
	CategoryIcon               types.String  `tfsdk:"category_icon"`
	CategoryTitle              types.String  `tfsdk:"category_title"`
	CompletedAt                types.String  `tfsdk:"completed_at"`
	ConsumerReviewedAt         types.String  `tfsdk:"consumer_reviewed_at"`
	ConsumerReviewedBy         types.String  `tfsdk:"consumer_reviewed_by"`
	ConsumerReviewedByFullName types.String  `tfsdk:"consumer_reviewed_by_full_name"`
	ConsumerReviewedByUsername types.String  `tfsdk:"consumer_reviewed_by_username"`
	Cost                       types.String  `tfsdk:"cost"`
	CreatedByCivilNumber       types.String  `tfsdk:"created_by_civil_number"`
	CreatedByFullName          types.String  `tfsdk:"created_by_full_name"`
	CreatedByUsername          types.String  `tfsdk:"created_by_username"`
	CustomerSlug               types.String  `tfsdk:"customer_slug"`
	ErrorMessage               types.String  `tfsdk:"error_message"`
	ErrorTraceback             types.String  `tfsdk:"error_traceback"`
	FixedPrice                 types.Float64 `tfsdk:"fixed_price"`
	NewCostEstimate            types.String  `tfsdk:"new_cost_estimate"`
	NewPlanName                types.String  `tfsdk:"new_plan_name"`
	NewPlanUuid                types.String  `tfsdk:"new_plan_uuid"`
	OfferingBillable           types.Bool    `tfsdk:"offering_billable"`
	OfferingDescription        types.String  `tfsdk:"offering_description"`
	OfferingImage              types.String  `tfsdk:"offering_image"`
	OfferingName               types.String  `tfsdk:"offering_name"`
	OfferingShared             types.Bool    `tfsdk:"offering_shared"`
	OfferingThumbnail          types.String  `tfsdk:"offering_thumbnail"`
	OldCostEstimate            types.String  `tfsdk:"old_cost_estimate"`
	OldPlanName                types.String  `tfsdk:"old_plan_name"`
	OldPlanUuid                types.String  `tfsdk:"old_plan_uuid"`
	Output                     types.String  `tfsdk:"output"`
	Plan                       types.String  `tfsdk:"plan"`
	PlanDescription            types.String  `tfsdk:"plan_description"`
	PlanName                   types.String  `tfsdk:"plan_name"`
	PlanUnit                   types.String  `tfsdk:"plan_unit"`
	PlanUuid                   types.String  `tfsdk:"plan_uuid"`
	ProjectDescription         types.String  `tfsdk:"project_description"`
	ProjectSlug                types.String  `tfsdk:"project_slug"`
	ProviderName               types.String  `tfsdk:"provider_name"`
	ProviderReviewedAt         types.String  `tfsdk:"provider_reviewed_at"`
	ProviderReviewedBy         types.String  `tfsdk:"provider_reviewed_by"`
	ProviderReviewedByFullName types.String  `tfsdk:"provider_reviewed_by_full_name"`
	ProviderReviewedByUsername types.String  `tfsdk:"provider_reviewed_by_username"`
	ProviderSlug               types.String  `tfsdk:"provider_slug"`
	RequestComment             types.String  `tfsdk:"request_comment"`
	ResourceName               types.String  `tfsdk:"resource_name"`
	ResourceType               types.String  `tfsdk:"resource_type"`
	StartDate                  types.String  `tfsdk:"start_date"`
	TerminationComment         types.String  `tfsdk:"termination_comment"`
	Url                        types.String  `tfsdk:"url"`
}

func (d *MarketplaceOrderDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_order"
}

func (d *MarketplaceOrderDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MarketplaceOrder data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"can_approve_as_consumer": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"can_approve_as_provider": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"category_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"created": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Created after",
			},
			"customer_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"modified": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Modified after",
			},
			"offering": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"offering_slug": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Multiple values may be separated by commas.",
			},
			"offering_type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"offering_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"parent_offering_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"project_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"provider_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"query": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Search by order UUID, slug, project name or resource name",
			},
			"resource": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"resource_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"service_manager_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"state": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"activation_price": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"attachment": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"callback_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"can_terminate": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"category_icon": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"category_title": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"completed_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"consumer_reviewed_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"consumer_reviewed_by": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
			},
			"consumer_reviewed_by_full_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"consumer_reviewed_by_username": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
			},
			"cost": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"created_by_civil_number": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"created_by_full_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"created_by_username": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
			},
			"customer_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"fixed_price": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"new_cost_estimate": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"new_plan_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"new_plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"offering_billable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Purchase and usage is invoiced.",
			},
			"offering_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"offering_image": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"offering_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"offering_shared": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Accessible to all customers.",
			},
			"offering_thumbnail": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"old_cost_estimate": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"old_plan_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"old_plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"output": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"plan": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"plan_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"plan_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"plan_unit": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"project_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"project_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"provider_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"provider_reviewed_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"provider_reviewed_by": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
			},
			"provider_reviewed_by_full_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"provider_reviewed_by_username": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
			},
			"provider_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"request_comment": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"resource_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"start_date": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Enables delayed processing of resource provisioning order.",
			},
			"termination_comment": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
		},
	}
}

func (d *MarketplaceOrderDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
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
		var item map[string]interface{}

		err := d.client.GetByUUID(ctx, "/api/marketplace-orders/{uuid}/", data.UUID.ValueString(), &item)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Order",
				"An error occurred while reading the marketplace_order by UUID: "+err.Error(),
			)
			return
		}

		// Extract data from single result
		if uuid, ok := item["uuid"].(string); ok {
			data.UUID = types.StringValue(uuid)
		}

		sourceMap := item
		// Map response fields to data model
		_ = sourceMap
		if val, ok := sourceMap["activation_price"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.ActivationPrice = types.Float64Value(num)
			}
		} else {
			if data.ActivationPrice.IsUnknown() {
				data.ActivationPrice = types.Float64Null()
			}
		}
		if val, ok := sourceMap["attachment"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Attachment = types.StringValue(str)
			}
		} else {
			if data.Attachment.IsUnknown() {
				data.Attachment = types.StringNull()
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
		if val, ok := sourceMap["callback_url"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CallbackUrl = types.StringValue(str)
			}
		} else {
			if data.CallbackUrl.IsUnknown() {
				data.CallbackUrl = types.StringNull()
			}
		}
		if val, ok := sourceMap["can_terminate"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanTerminate = types.BoolValue(b)
			}
		} else {
			if data.CanTerminate.IsUnknown() {
				data.CanTerminate = types.BoolNull()
			}
		}
		if val, ok := sourceMap["category_icon"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CategoryIcon = types.StringValue(str)
			}
		} else {
			if data.CategoryIcon.IsUnknown() {
				data.CategoryIcon = types.StringNull()
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
		if val, ok := sourceMap["completed_at"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CompletedAt = types.StringValue(str)
			}
		} else {
			if data.CompletedAt.IsUnknown() {
				data.CompletedAt = types.StringNull()
			}
		}
		if val, ok := sourceMap["consumer_reviewed_at"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ConsumerReviewedAt = types.StringValue(str)
			}
		} else {
			if data.ConsumerReviewedAt.IsUnknown() {
				data.ConsumerReviewedAt = types.StringNull()
			}
		}
		if val, ok := sourceMap["consumer_reviewed_by"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ConsumerReviewedBy = types.StringValue(str)
			}
		} else {
			if data.ConsumerReviewedBy.IsUnknown() {
				data.ConsumerReviewedBy = types.StringNull()
			}
		}
		if val, ok := sourceMap["consumer_reviewed_by_full_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ConsumerReviewedByFullName = types.StringValue(str)
			}
		} else {
			if data.ConsumerReviewedByFullName.IsUnknown() {
				data.ConsumerReviewedByFullName = types.StringNull()
			}
		}
		if val, ok := sourceMap["consumer_reviewed_by_username"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ConsumerReviewedByUsername = types.StringValue(str)
			}
		} else {
			if data.ConsumerReviewedByUsername.IsUnknown() {
				data.ConsumerReviewedByUsername = types.StringNull()
			}
		}
		if val, ok := sourceMap["cost"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Cost = types.StringValue(str)
			}
		} else {
			if data.Cost.IsUnknown() {
				data.Cost = types.StringNull()
			}
		}
		if val, ok := sourceMap["created_by_civil_number"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CreatedByCivilNumber = types.StringValue(str)
			}
		} else {
			if data.CreatedByCivilNumber.IsUnknown() {
				data.CreatedByCivilNumber = types.StringNull()
			}
		}
		if val, ok := sourceMap["created_by_full_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CreatedByFullName = types.StringValue(str)
			}
		} else {
			if data.CreatedByFullName.IsUnknown() {
				data.CreatedByFullName = types.StringNull()
			}
		}
		if val, ok := sourceMap["created_by_username"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CreatedByUsername = types.StringValue(str)
			}
		} else {
			if data.CreatedByUsername.IsUnknown() {
				data.CreatedByUsername = types.StringNull()
			}
		}
		if val, ok := sourceMap["customer_slug"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerSlug = types.StringValue(str)
			}
		} else {
			if data.CustomerSlug.IsUnknown() {
				data.CustomerSlug = types.StringNull()
			}
		}
		if val, ok := sourceMap["error_message"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ErrorMessage = types.StringValue(str)
			}
		} else {
			if data.ErrorMessage.IsUnknown() {
				data.ErrorMessage = types.StringNull()
			}
		}
		if val, ok := sourceMap["error_traceback"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ErrorTraceback = types.StringValue(str)
			}
		} else {
			if data.ErrorTraceback.IsUnknown() {
				data.ErrorTraceback = types.StringNull()
			}
		}
		if val, ok := sourceMap["fixed_price"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.FixedPrice = types.Float64Value(num)
			}
		} else {
			if data.FixedPrice.IsUnknown() {
				data.FixedPrice = types.Float64Null()
			}
		}
		if val, ok := sourceMap["new_cost_estimate"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NewCostEstimate = types.StringValue(str)
			}
		} else {
			if data.NewCostEstimate.IsUnknown() {
				data.NewCostEstimate = types.StringNull()
			}
		}
		if val, ok := sourceMap["new_plan_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NewPlanName = types.StringValue(str)
			}
		} else {
			if data.NewPlanName.IsUnknown() {
				data.NewPlanName = types.StringNull()
			}
		}
		if val, ok := sourceMap["new_plan_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NewPlanUuid = types.StringValue(str)
			}
		} else {
			if data.NewPlanUuid.IsUnknown() {
				data.NewPlanUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["offering_billable"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.OfferingBillable = types.BoolValue(b)
			}
		} else {
			if data.OfferingBillable.IsUnknown() {
				data.OfferingBillable = types.BoolNull()
			}
		}
		if val, ok := sourceMap["offering_description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingDescription = types.StringValue(str)
			}
		} else {
			if data.OfferingDescription.IsUnknown() {
				data.OfferingDescription = types.StringNull()
			}
		}
		if val, ok := sourceMap["offering_image"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingImage = types.StringValue(str)
			}
		} else {
			if data.OfferingImage.IsUnknown() {
				data.OfferingImage = types.StringNull()
			}
		}
		if val, ok := sourceMap["offering_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingName = types.StringValue(str)
			}
		} else {
			if data.OfferingName.IsUnknown() {
				data.OfferingName = types.StringNull()
			}
		}
		if val, ok := sourceMap["offering_shared"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.OfferingShared = types.BoolValue(b)
			}
		} else {
			if data.OfferingShared.IsUnknown() {
				data.OfferingShared = types.BoolNull()
			}
		}
		if val, ok := sourceMap["offering_thumbnail"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingThumbnail = types.StringValue(str)
			}
		} else {
			if data.OfferingThumbnail.IsUnknown() {
				data.OfferingThumbnail = types.StringNull()
			}
		}
		if val, ok := sourceMap["old_cost_estimate"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OldCostEstimate = types.StringValue(str)
			}
		} else {
			if data.OldCostEstimate.IsUnknown() {
				data.OldCostEstimate = types.StringNull()
			}
		}
		if val, ok := sourceMap["old_plan_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OldPlanName = types.StringValue(str)
			}
		} else {
			if data.OldPlanName.IsUnknown() {
				data.OldPlanName = types.StringNull()
			}
		}
		if val, ok := sourceMap["old_plan_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OldPlanUuid = types.StringValue(str)
			}
		} else {
			if data.OldPlanUuid.IsUnknown() {
				data.OldPlanUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["output"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Output = types.StringValue(str)
			}
		} else {
			if data.Output.IsUnknown() {
				data.Output = types.StringNull()
			}
		}
		if val, ok := sourceMap["plan"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Plan = types.StringValue(str)
			}
		} else {
			if data.Plan.IsUnknown() {
				data.Plan = types.StringNull()
			}
		}
		if val, ok := sourceMap["plan_description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.PlanDescription = types.StringValue(str)
			}
		} else {
			if data.PlanDescription.IsUnknown() {
				data.PlanDescription = types.StringNull()
			}
		}
		if val, ok := sourceMap["plan_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.PlanName = types.StringValue(str)
			}
		} else {
			if data.PlanName.IsUnknown() {
				data.PlanName = types.StringNull()
			}
		}
		if val, ok := sourceMap["plan_unit"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.PlanUnit = types.StringValue(str)
			}
		} else {
			if data.PlanUnit.IsUnknown() {
				data.PlanUnit = types.StringNull()
			}
		}
		if val, ok := sourceMap["plan_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.PlanUuid = types.StringValue(str)
			}
		} else {
			if data.PlanUuid.IsUnknown() {
				data.PlanUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["project_description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectDescription = types.StringValue(str)
			}
		} else {
			if data.ProjectDescription.IsUnknown() {
				data.ProjectDescription = types.StringNull()
			}
		}
		if val, ok := sourceMap["project_slug"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectSlug = types.StringValue(str)
			}
		} else {
			if data.ProjectSlug.IsUnknown() {
				data.ProjectSlug = types.StringNull()
			}
		}
		if val, ok := sourceMap["provider_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProviderName = types.StringValue(str)
			}
		} else {
			if data.ProviderName.IsUnknown() {
				data.ProviderName = types.StringNull()
			}
		}
		if val, ok := sourceMap["provider_reviewed_at"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProviderReviewedAt = types.StringValue(str)
			}
		} else {
			if data.ProviderReviewedAt.IsUnknown() {
				data.ProviderReviewedAt = types.StringNull()
			}
		}
		if val, ok := sourceMap["provider_reviewed_by"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProviderReviewedBy = types.StringValue(str)
			}
		} else {
			if data.ProviderReviewedBy.IsUnknown() {
				data.ProviderReviewedBy = types.StringNull()
			}
		}
		if val, ok := sourceMap["provider_reviewed_by_full_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProviderReviewedByFullName = types.StringValue(str)
			}
		} else {
			if data.ProviderReviewedByFullName.IsUnknown() {
				data.ProviderReviewedByFullName = types.StringNull()
			}
		}
		if val, ok := sourceMap["provider_reviewed_by_username"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProviderReviewedByUsername = types.StringValue(str)
			}
		} else {
			if data.ProviderReviewedByUsername.IsUnknown() {
				data.ProviderReviewedByUsername = types.StringNull()
			}
		}
		if val, ok := sourceMap["provider_slug"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProviderSlug = types.StringValue(str)
			}
		} else {
			if data.ProviderSlug.IsUnknown() {
				data.ProviderSlug = types.StringNull()
			}
		}
		if val, ok := sourceMap["request_comment"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.RequestComment = types.StringValue(str)
			}
		} else {
			if data.RequestComment.IsUnknown() {
				data.RequestComment = types.StringNull()
			}
		}
		if val, ok := sourceMap["resource_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ResourceName = types.StringValue(str)
			}
		} else {
			if data.ResourceName.IsUnknown() {
				data.ResourceName = types.StringNull()
			}
		}
		if val, ok := sourceMap["resource_type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ResourceType = types.StringValue(str)
			}
		} else {
			if data.ResourceType.IsUnknown() {
				data.ResourceType = types.StringNull()
			}
		}
		if val, ok := sourceMap["start_date"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.StartDate = types.StringValue(str)
			}
		} else {
			if data.StartDate.IsUnknown() {
				data.StartDate = types.StringNull()
			}
		}
		if val, ok := sourceMap["termination_comment"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.TerminationComment = types.StringValue(str)
			}
		} else {
			if data.TerminationComment.IsUnknown() {
				data.TerminationComment = types.StringNull()
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

		// Map filter parameters from response if available
		if val, ok := sourceMap["can_approve_as_consumer"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanApproveAsConsumer = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["can_approve_as_provider"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanApproveAsProvider = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["category_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CategoryUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["created"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Created = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["modified"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Modified = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Offering = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering_slug"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingSlug = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering_type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingType = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["parent_offering_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ParentOfferingUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["provider_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProviderUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["query"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Query = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["resource"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Resource = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["resource_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ResourceUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_manager_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceManagerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.State = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Type = types.StringValue(str)
			}
		}

	} else {
		// Filter by provided parameters
		var results []map[string]interface{}

		filters := map[string]string{}
		if !data.CanApproveAsConsumer.IsNull() {
			filters["can_approve_as_consumer"] = fmt.Sprintf("%t", data.CanApproveAsConsumer.ValueBool())
		}
		if !data.CanApproveAsProvider.IsNull() {
			filters["can_approve_as_provider"] = fmt.Sprintf("%t", data.CanApproveAsProvider.ValueBool())
		}
		if !data.CategoryUuid.IsNull() {
			filters["category_uuid"] = data.CategoryUuid.ValueString()
		}
		if !data.Created.IsNull() {
			filters["created"] = data.Created.ValueString()
		}
		if !data.CustomerUuid.IsNull() {
			filters["customer_uuid"] = data.CustomerUuid.ValueString()
		}
		if !data.Modified.IsNull() {
			filters["modified"] = data.Modified.ValueString()
		}
		if !data.Offering.IsNull() {
			filters["offering"] = data.Offering.ValueString()
		}
		if !data.OfferingSlug.IsNull() {
			filters["offering_slug"] = data.OfferingSlug.ValueString()
		}
		if !data.OfferingType.IsNull() {
			filters["offering_type"] = data.OfferingType.ValueString()
		}
		if !data.OfferingUuid.IsNull() {
			filters["offering_uuid"] = data.OfferingUuid.ValueString()
		}
		if !data.ParentOfferingUuid.IsNull() {
			filters["parent_offering_uuid"] = data.ParentOfferingUuid.ValueString()
		}
		if !data.ProjectUuid.IsNull() {
			filters["project_uuid"] = data.ProjectUuid.ValueString()
		}
		if !data.ProviderUuid.IsNull() {
			filters["provider_uuid"] = data.ProviderUuid.ValueString()
		}
		if !data.Query.IsNull() {
			filters["query"] = data.Query.ValueString()
		}
		if !data.Resource.IsNull() {
			filters["resource"] = data.Resource.ValueString()
		}
		if !data.ResourceUuid.IsNull() {
			filters["resource_uuid"] = data.ResourceUuid.ValueString()
		}
		if !data.ServiceManagerUuid.IsNull() {
			filters["service_manager_uuid"] = data.ServiceManagerUuid.ValueString()
		}
		if !data.State.IsNull() {
			filters["state"] = data.State.ValueString()
		}
		if !data.Type.IsNull() {
			filters["type"] = data.Type.ValueString()
		}

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup marketplace_order.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/marketplace-orders/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Order",
				"An error occurred while filtering marketplace_order: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Order Not Found",
				"No marketplace_order found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Orders Found",
				fmt.Sprintf("Found %d marketplace_orders with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		// Extract data from the single result
		if uuid, ok := results[0]["uuid"].(string); ok {
			data.UUID = types.StringValue(uuid)
		}
		sourceMap := results[0]
		// Map response fields to data model
		_ = sourceMap
		if val, ok := sourceMap["activation_price"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.ActivationPrice = types.Float64Value(num)
			}
		} else {
			if data.ActivationPrice.IsUnknown() {
				data.ActivationPrice = types.Float64Null()
			}
		}
		if val, ok := sourceMap["attachment"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Attachment = types.StringValue(str)
			}
		} else {
			if data.Attachment.IsUnknown() {
				data.Attachment = types.StringNull()
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
		if val, ok := sourceMap["callback_url"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CallbackUrl = types.StringValue(str)
			}
		} else {
			if data.CallbackUrl.IsUnknown() {
				data.CallbackUrl = types.StringNull()
			}
		}
		if val, ok := sourceMap["can_terminate"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanTerminate = types.BoolValue(b)
			}
		} else {
			if data.CanTerminate.IsUnknown() {
				data.CanTerminate = types.BoolNull()
			}
		}
		if val, ok := sourceMap["category_icon"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CategoryIcon = types.StringValue(str)
			}
		} else {
			if data.CategoryIcon.IsUnknown() {
				data.CategoryIcon = types.StringNull()
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
		if val, ok := sourceMap["completed_at"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CompletedAt = types.StringValue(str)
			}
		} else {
			if data.CompletedAt.IsUnknown() {
				data.CompletedAt = types.StringNull()
			}
		}
		if val, ok := sourceMap["consumer_reviewed_at"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ConsumerReviewedAt = types.StringValue(str)
			}
		} else {
			if data.ConsumerReviewedAt.IsUnknown() {
				data.ConsumerReviewedAt = types.StringNull()
			}
		}
		if val, ok := sourceMap["consumer_reviewed_by"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ConsumerReviewedBy = types.StringValue(str)
			}
		} else {
			if data.ConsumerReviewedBy.IsUnknown() {
				data.ConsumerReviewedBy = types.StringNull()
			}
		}
		if val, ok := sourceMap["consumer_reviewed_by_full_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ConsumerReviewedByFullName = types.StringValue(str)
			}
		} else {
			if data.ConsumerReviewedByFullName.IsUnknown() {
				data.ConsumerReviewedByFullName = types.StringNull()
			}
		}
		if val, ok := sourceMap["consumer_reviewed_by_username"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ConsumerReviewedByUsername = types.StringValue(str)
			}
		} else {
			if data.ConsumerReviewedByUsername.IsUnknown() {
				data.ConsumerReviewedByUsername = types.StringNull()
			}
		}
		if val, ok := sourceMap["cost"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Cost = types.StringValue(str)
			}
		} else {
			if data.Cost.IsUnknown() {
				data.Cost = types.StringNull()
			}
		}
		if val, ok := sourceMap["created_by_civil_number"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CreatedByCivilNumber = types.StringValue(str)
			}
		} else {
			if data.CreatedByCivilNumber.IsUnknown() {
				data.CreatedByCivilNumber = types.StringNull()
			}
		}
		if val, ok := sourceMap["created_by_full_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CreatedByFullName = types.StringValue(str)
			}
		} else {
			if data.CreatedByFullName.IsUnknown() {
				data.CreatedByFullName = types.StringNull()
			}
		}
		if val, ok := sourceMap["created_by_username"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CreatedByUsername = types.StringValue(str)
			}
		} else {
			if data.CreatedByUsername.IsUnknown() {
				data.CreatedByUsername = types.StringNull()
			}
		}
		if val, ok := sourceMap["customer_slug"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerSlug = types.StringValue(str)
			}
		} else {
			if data.CustomerSlug.IsUnknown() {
				data.CustomerSlug = types.StringNull()
			}
		}
		if val, ok := sourceMap["error_message"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ErrorMessage = types.StringValue(str)
			}
		} else {
			if data.ErrorMessage.IsUnknown() {
				data.ErrorMessage = types.StringNull()
			}
		}
		if val, ok := sourceMap["error_traceback"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ErrorTraceback = types.StringValue(str)
			}
		} else {
			if data.ErrorTraceback.IsUnknown() {
				data.ErrorTraceback = types.StringNull()
			}
		}
		if val, ok := sourceMap["fixed_price"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.FixedPrice = types.Float64Value(num)
			}
		} else {
			if data.FixedPrice.IsUnknown() {
				data.FixedPrice = types.Float64Null()
			}
		}
		if val, ok := sourceMap["new_cost_estimate"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NewCostEstimate = types.StringValue(str)
			}
		} else {
			if data.NewCostEstimate.IsUnknown() {
				data.NewCostEstimate = types.StringNull()
			}
		}
		if val, ok := sourceMap["new_plan_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NewPlanName = types.StringValue(str)
			}
		} else {
			if data.NewPlanName.IsUnknown() {
				data.NewPlanName = types.StringNull()
			}
		}
		if val, ok := sourceMap["new_plan_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NewPlanUuid = types.StringValue(str)
			}
		} else {
			if data.NewPlanUuid.IsUnknown() {
				data.NewPlanUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["offering_billable"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.OfferingBillable = types.BoolValue(b)
			}
		} else {
			if data.OfferingBillable.IsUnknown() {
				data.OfferingBillable = types.BoolNull()
			}
		}
		if val, ok := sourceMap["offering_description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingDescription = types.StringValue(str)
			}
		} else {
			if data.OfferingDescription.IsUnknown() {
				data.OfferingDescription = types.StringNull()
			}
		}
		if val, ok := sourceMap["offering_image"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingImage = types.StringValue(str)
			}
		} else {
			if data.OfferingImage.IsUnknown() {
				data.OfferingImage = types.StringNull()
			}
		}
		if val, ok := sourceMap["offering_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingName = types.StringValue(str)
			}
		} else {
			if data.OfferingName.IsUnknown() {
				data.OfferingName = types.StringNull()
			}
		}
		if val, ok := sourceMap["offering_shared"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.OfferingShared = types.BoolValue(b)
			}
		} else {
			if data.OfferingShared.IsUnknown() {
				data.OfferingShared = types.BoolNull()
			}
		}
		if val, ok := sourceMap["offering_thumbnail"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingThumbnail = types.StringValue(str)
			}
		} else {
			if data.OfferingThumbnail.IsUnknown() {
				data.OfferingThumbnail = types.StringNull()
			}
		}
		if val, ok := sourceMap["old_cost_estimate"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OldCostEstimate = types.StringValue(str)
			}
		} else {
			if data.OldCostEstimate.IsUnknown() {
				data.OldCostEstimate = types.StringNull()
			}
		}
		if val, ok := sourceMap["old_plan_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OldPlanName = types.StringValue(str)
			}
		} else {
			if data.OldPlanName.IsUnknown() {
				data.OldPlanName = types.StringNull()
			}
		}
		if val, ok := sourceMap["old_plan_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OldPlanUuid = types.StringValue(str)
			}
		} else {
			if data.OldPlanUuid.IsUnknown() {
				data.OldPlanUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["output"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Output = types.StringValue(str)
			}
		} else {
			if data.Output.IsUnknown() {
				data.Output = types.StringNull()
			}
		}
		if val, ok := sourceMap["plan"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Plan = types.StringValue(str)
			}
		} else {
			if data.Plan.IsUnknown() {
				data.Plan = types.StringNull()
			}
		}
		if val, ok := sourceMap["plan_description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.PlanDescription = types.StringValue(str)
			}
		} else {
			if data.PlanDescription.IsUnknown() {
				data.PlanDescription = types.StringNull()
			}
		}
		if val, ok := sourceMap["plan_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.PlanName = types.StringValue(str)
			}
		} else {
			if data.PlanName.IsUnknown() {
				data.PlanName = types.StringNull()
			}
		}
		if val, ok := sourceMap["plan_unit"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.PlanUnit = types.StringValue(str)
			}
		} else {
			if data.PlanUnit.IsUnknown() {
				data.PlanUnit = types.StringNull()
			}
		}
		if val, ok := sourceMap["plan_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.PlanUuid = types.StringValue(str)
			}
		} else {
			if data.PlanUuid.IsUnknown() {
				data.PlanUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["project_description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectDescription = types.StringValue(str)
			}
		} else {
			if data.ProjectDescription.IsUnknown() {
				data.ProjectDescription = types.StringNull()
			}
		}
		if val, ok := sourceMap["project_slug"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectSlug = types.StringValue(str)
			}
		} else {
			if data.ProjectSlug.IsUnknown() {
				data.ProjectSlug = types.StringNull()
			}
		}
		if val, ok := sourceMap["provider_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProviderName = types.StringValue(str)
			}
		} else {
			if data.ProviderName.IsUnknown() {
				data.ProviderName = types.StringNull()
			}
		}
		if val, ok := sourceMap["provider_reviewed_at"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProviderReviewedAt = types.StringValue(str)
			}
		} else {
			if data.ProviderReviewedAt.IsUnknown() {
				data.ProviderReviewedAt = types.StringNull()
			}
		}
		if val, ok := sourceMap["provider_reviewed_by"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProviderReviewedBy = types.StringValue(str)
			}
		} else {
			if data.ProviderReviewedBy.IsUnknown() {
				data.ProviderReviewedBy = types.StringNull()
			}
		}
		if val, ok := sourceMap["provider_reviewed_by_full_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProviderReviewedByFullName = types.StringValue(str)
			}
		} else {
			if data.ProviderReviewedByFullName.IsUnknown() {
				data.ProviderReviewedByFullName = types.StringNull()
			}
		}
		if val, ok := sourceMap["provider_reviewed_by_username"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProviderReviewedByUsername = types.StringValue(str)
			}
		} else {
			if data.ProviderReviewedByUsername.IsUnknown() {
				data.ProviderReviewedByUsername = types.StringNull()
			}
		}
		if val, ok := sourceMap["provider_slug"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProviderSlug = types.StringValue(str)
			}
		} else {
			if data.ProviderSlug.IsUnknown() {
				data.ProviderSlug = types.StringNull()
			}
		}
		if val, ok := sourceMap["request_comment"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.RequestComment = types.StringValue(str)
			}
		} else {
			if data.RequestComment.IsUnknown() {
				data.RequestComment = types.StringNull()
			}
		}
		if val, ok := sourceMap["resource_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ResourceName = types.StringValue(str)
			}
		} else {
			if data.ResourceName.IsUnknown() {
				data.ResourceName = types.StringNull()
			}
		}
		if val, ok := sourceMap["resource_type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ResourceType = types.StringValue(str)
			}
		} else {
			if data.ResourceType.IsUnknown() {
				data.ResourceType = types.StringNull()
			}
		}
		if val, ok := sourceMap["start_date"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.StartDate = types.StringValue(str)
			}
		} else {
			if data.StartDate.IsUnknown() {
				data.StartDate = types.StringNull()
			}
		}
		if val, ok := sourceMap["termination_comment"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.TerminationComment = types.StringValue(str)
			}
		} else {
			if data.TerminationComment.IsUnknown() {
				data.TerminationComment = types.StringNull()
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

		// Map filter parameters from response if available
		if val, ok := sourceMap["can_approve_as_consumer"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanApproveAsConsumer = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["can_approve_as_provider"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanApproveAsProvider = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["category_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CategoryUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["created"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Created = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["modified"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Modified = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Offering = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering_slug"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingSlug = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering_type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingType = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["parent_offering_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ParentOfferingUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["provider_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProviderUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["query"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Query = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["resource"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Resource = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["resource_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ResourceUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_manager_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceManagerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.State = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Type = types.StringValue(str)
			}
		}
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
