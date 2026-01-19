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

	"github.com/waldur/terraform-waldur-provider/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MarketplaceOrderResource{}
var _ resource.ResourceWithImportState = &MarketplaceOrderResource{}

func NewMarketplaceOrderResource() resource.Resource {
	return &MarketplaceOrderResource{}
}

// MarketplaceOrderResource defines the resource implementation.
type MarketplaceOrderResource struct {
	client *client.Client
}

// MarketplaceOrderResourceModel describes the resource data model.
type MarketplaceOrderResourceModel struct {
	UUID                       types.String   `tfsdk:"id"`
	AcceptingTermsOfService    types.Bool     `tfsdk:"accepting_terms_of_service"`
	ActivationPrice            types.Float64  `tfsdk:"activation_price"`
	Attachment                 types.String   `tfsdk:"attachment"`
	BackendId                  types.String   `tfsdk:"backend_id"`
	CallbackUrl                types.String   `tfsdk:"callback_url"`
	CanTerminate               types.Bool     `tfsdk:"can_terminate"`
	CategoryIcon               types.String   `tfsdk:"category_icon"`
	CategoryTitle              types.String   `tfsdk:"category_title"`
	CategoryUuid               types.String   `tfsdk:"category_uuid"`
	CompletedAt                types.String   `tfsdk:"completed_at"`
	ConsumerReviewedAt         types.String   `tfsdk:"consumer_reviewed_at"`
	ConsumerReviewedBy         types.String   `tfsdk:"consumer_reviewed_by"`
	ConsumerReviewedByFullName types.String   `tfsdk:"consumer_reviewed_by_full_name"`
	ConsumerReviewedByUsername types.String   `tfsdk:"consumer_reviewed_by_username"`
	Cost                       types.String   `tfsdk:"cost"`
	Created                    types.String   `tfsdk:"created"`
	CreatedByCivilNumber       types.String   `tfsdk:"created_by_civil_number"`
	CreatedByFullName          types.String   `tfsdk:"created_by_full_name"`
	CreatedByUsername          types.String   `tfsdk:"created_by_username"`
	CustomerSlug               types.String   `tfsdk:"customer_slug"`
	ErrorMessage               types.String   `tfsdk:"error_message"`
	ErrorTraceback             types.String   `tfsdk:"error_traceback"`
	FixedPrice                 types.Float64  `tfsdk:"fixed_price"`
	Modified                   types.String   `tfsdk:"modified"`
	NewCostEstimate            types.String   `tfsdk:"new_cost_estimate"`
	NewPlanName                types.String   `tfsdk:"new_plan_name"`
	NewPlanUuid                types.String   `tfsdk:"new_plan_uuid"`
	Offering                   types.String   `tfsdk:"offering"`
	OfferingBillable           types.Bool     `tfsdk:"offering_billable"`
	OfferingDescription        types.String   `tfsdk:"offering_description"`
	OfferingImage              types.String   `tfsdk:"offering_image"`
	OfferingName               types.String   `tfsdk:"offering_name"`
	OfferingShared             types.Bool     `tfsdk:"offering_shared"`
	OfferingThumbnail          types.String   `tfsdk:"offering_thumbnail"`
	OfferingType               types.String   `tfsdk:"offering_type"`
	OfferingUuid               types.String   `tfsdk:"offering_uuid"`
	OldCostEstimate            types.String   `tfsdk:"old_cost_estimate"`
	OldPlanName                types.String   `tfsdk:"old_plan_name"`
	OldPlanUuid                types.String   `tfsdk:"old_plan_uuid"`
	Output                     types.String   `tfsdk:"output"`
	Plan                       types.String   `tfsdk:"plan"`
	PlanDescription            types.String   `tfsdk:"plan_description"`
	PlanName                   types.String   `tfsdk:"plan_name"`
	PlanUnit                   types.String   `tfsdk:"plan_unit"`
	PlanUuid                   types.String   `tfsdk:"plan_uuid"`
	Project                    types.String   `tfsdk:"project"`
	ProjectDescription         types.String   `tfsdk:"project_description"`
	ProjectSlug                types.String   `tfsdk:"project_slug"`
	ProviderName               types.String   `tfsdk:"provider_name"`
	ProviderReviewedAt         types.String   `tfsdk:"provider_reviewed_at"`
	ProviderReviewedBy         types.String   `tfsdk:"provider_reviewed_by"`
	ProviderReviewedByFullName types.String   `tfsdk:"provider_reviewed_by_full_name"`
	ProviderReviewedByUsername types.String   `tfsdk:"provider_reviewed_by_username"`
	ProviderSlug               types.String   `tfsdk:"provider_slug"`
	ProviderUuid               types.String   `tfsdk:"provider_uuid"`
	RequestComment             types.String   `tfsdk:"request_comment"`
	ResourceName               types.String   `tfsdk:"resource_name"`
	ResourceType               types.String   `tfsdk:"resource_type"`
	ResourceUuid               types.String   `tfsdk:"resource_uuid"`
	StartDate                  types.String   `tfsdk:"start_date"`
	State                      types.String   `tfsdk:"state"`
	TerminationComment         types.String   `tfsdk:"termination_comment"`
	Type                       types.String   `tfsdk:"type"`
	Url                        types.String   `tfsdk:"url"`
	Timeouts                   timeouts.Value `tfsdk:"timeouts"`
}

func (r *MarketplaceOrderResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_order"
}

func (r *MarketplaceOrderResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MarketplaceOrder resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"accepting_terms_of_service": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"activation_price": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"attachment": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"callback_url": schema.StringAttribute{
				Optional:            true,
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
			"category_uuid": schema.StringAttribute{
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
			"created": schema.StringAttribute{
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
			"modified": schema.StringAttribute{
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
			"offering": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
			"offering_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"offering_uuid": schema.StringAttribute{
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
				Optional:            true,
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
			"project": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
			"provider_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"request_comment": schema.StringAttribute{
				Optional:            true,
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
			"resource_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"start_date": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Enables delayed processing of resource provisioning order.",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"termination_comment": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
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

func (r *MarketplaceOrderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MarketplaceOrderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MarketplaceOrderResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Prepare request body
	requestBody := map[string]interface{}{}
	if !data.AcceptingTermsOfService.IsNull() && !data.AcceptingTermsOfService.IsUnknown() {
		requestBody["accepting_terms_of_service"] = data.AcceptingTermsOfService.ValueBool()
	}
	if !data.CallbackUrl.IsNull() && !data.CallbackUrl.IsUnknown() {
		if v := data.CallbackUrl.ValueString(); v != "" {
			requestBody["callback_url"] = v
		}
	}
	if !data.Offering.IsNull() && !data.Offering.IsUnknown() {
		if v := data.Offering.ValueString(); v != "" {
			requestBody["offering"] = v
		}
	}
	if !data.Plan.IsNull() && !data.Plan.IsUnknown() {
		if v := data.Plan.ValueString(); v != "" {
			requestBody["plan"] = v
		}
	}
	if !data.Project.IsNull() && !data.Project.IsUnknown() {
		if v := data.Project.ValueString(); v != "" {
			requestBody["project"] = v
		}
	}
	if !data.RequestComment.IsNull() && !data.RequestComment.IsUnknown() {
		if v := data.RequestComment.ValueString(); v != "" {
			requestBody["request_comment"] = v
		}
	}
	if !data.StartDate.IsNull() && !data.StartDate.IsUnknown() {
		if v := data.StartDate.ValueString(); v != "" {
			requestBody["start_date"] = v
		}
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		if v := data.Type.ValueString(); v != "" {
			requestBody["type"] = v
		}
	}

	// Call Waldur API to create resource
	var result map[string]interface{}
	err := r.client.Create(ctx, "/api/marketplace-orders/", requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create MarketplaceOrder",
			"An error occurred while creating the marketplace_order: "+err.Error(),
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
	if val, ok := sourceMap["accepting_terms_of_service"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.AcceptingTermsOfService = types.BoolValue(b)
		}
	} else {
		if data.AcceptingTermsOfService.IsUnknown() {
			data.AcceptingTermsOfService = types.BoolNull()
		}
	}
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
	if val, ok := sourceMap["category_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CategoryUuid = types.StringValue(str)
		}
	} else {
		if data.CategoryUuid.IsUnknown() {
			data.CategoryUuid = types.StringNull()
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
	if val, ok := sourceMap["created"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
	} else {
		if data.Created.IsUnknown() {
			data.Created = types.StringNull()
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
	if val, ok := sourceMap["modified"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Modified = types.StringValue(str)
		}
	} else {
		if data.Modified.IsUnknown() {
			data.Modified = types.StringNull()
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
	if val, ok := sourceMap["offering"]; ok && val != nil {
		if str, ok := val.(string); ok {
			// Normalize URL to UUID
			parts := strings.Split(strings.TrimRight(str, "/"), "/")
			uuid := parts[len(parts)-1]
			data.Offering = types.StringValue(uuid)
		} else {
			data.Offering = types.StringNull()
		}
	} else {
		if data.Offering.IsUnknown() {
			data.Offering = types.StringNull()
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
	if val, ok := sourceMap["offering_type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.OfferingType = types.StringValue(str)
		}
	} else {
		if data.OfferingType.IsUnknown() {
			data.OfferingType = types.StringNull()
		}
	}
	if val, ok := sourceMap["offering_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.OfferingUuid = types.StringValue(str)
		}
	} else {
		if data.OfferingUuid.IsUnknown() {
			data.OfferingUuid = types.StringNull()
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
	if val, ok := sourceMap["project"]; ok && val != nil {
		if str, ok := val.(string); ok {
			// Normalize URL to UUID
			parts := strings.Split(strings.TrimRight(str, "/"), "/")
			uuid := parts[len(parts)-1]
			data.Project = types.StringValue(uuid)
		} else {
			data.Project = types.StringNull()
		}
	} else {
		if data.Project.IsUnknown() {
			data.Project = types.StringNull()
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
	if val, ok := sourceMap["provider_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProviderUuid = types.StringValue(str)
		}
	} else {
		if data.ProviderUuid.IsUnknown() {
			data.ProviderUuid = types.StringNull()
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
	if val, ok := sourceMap["resource_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ResourceUuid = types.StringValue(str)
		}
	} else {
		if data.ResourceUuid.IsUnknown() {
			data.ResourceUuid = types.StringNull()
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
	if val, ok := sourceMap["state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.State = types.StringValue(str)
		}
	} else {
		if data.State.IsUnknown() {
			data.State = types.StringNull()
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

	// Map filter parameters from response if available

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MarketplaceOrderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MarketplaceOrderResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource
	var result map[string]interface{}

	retrievePath := strings.Replace("/api/marketplace-orders/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read MarketplaceOrder",
			"An error occurred while reading the marketplace_order: "+err.Error(),
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
	if val, ok := sourceMap["accepting_terms_of_service"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.AcceptingTermsOfService = types.BoolValue(b)
		}
	} else {
		if data.AcceptingTermsOfService.IsUnknown() {
			data.AcceptingTermsOfService = types.BoolNull()
		}
	}
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
	if val, ok := sourceMap["category_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CategoryUuid = types.StringValue(str)
		}
	} else {
		if data.CategoryUuid.IsUnknown() {
			data.CategoryUuid = types.StringNull()
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
	if val, ok := sourceMap["created"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
	} else {
		if data.Created.IsUnknown() {
			data.Created = types.StringNull()
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
	if val, ok := sourceMap["modified"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Modified = types.StringValue(str)
		}
	} else {
		if data.Modified.IsUnknown() {
			data.Modified = types.StringNull()
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
	if val, ok := sourceMap["offering"]; ok && val != nil {
		if str, ok := val.(string); ok {
			// Normalize URL to UUID
			parts := strings.Split(strings.TrimRight(str, "/"), "/")
			uuid := parts[len(parts)-1]
			data.Offering = types.StringValue(uuid)
		} else {
			data.Offering = types.StringNull()
		}
	} else {
		if data.Offering.IsUnknown() {
			data.Offering = types.StringNull()
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
	if val, ok := sourceMap["offering_type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.OfferingType = types.StringValue(str)
		}
	} else {
		if data.OfferingType.IsUnknown() {
			data.OfferingType = types.StringNull()
		}
	}
	if val, ok := sourceMap["offering_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.OfferingUuid = types.StringValue(str)
		}
	} else {
		if data.OfferingUuid.IsUnknown() {
			data.OfferingUuid = types.StringNull()
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
	if val, ok := sourceMap["project"]; ok && val != nil {
		if str, ok := val.(string); ok {
			// Normalize URL to UUID
			parts := strings.Split(strings.TrimRight(str, "/"), "/")
			uuid := parts[len(parts)-1]
			data.Project = types.StringValue(uuid)
		} else {
			data.Project = types.StringNull()
		}
	} else {
		if data.Project.IsUnknown() {
			data.Project = types.StringNull()
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
	if val, ok := sourceMap["provider_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProviderUuid = types.StringValue(str)
		}
	} else {
		if data.ProviderUuid.IsUnknown() {
			data.ProviderUuid = types.StringNull()
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
	if val, ok := sourceMap["resource_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ResourceUuid = types.StringValue(str)
		}
	} else {
		if data.ResourceUuid.IsUnknown() {
			data.ResourceUuid = types.StringNull()
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
	if val, ok := sourceMap["state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.State = types.StringValue(str)
		}
	} else {
		if data.State.IsUnknown() {
			data.State = types.StringNull()
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

	// Map filter parameters from response if available

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MarketplaceOrderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data MarketplaceOrderResourceModel
	var state MarketplaceOrderResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	// Read current state to get the UUID (which is computed and not in plan)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Use UUID from state
	data.UUID = state.UUID
	// Prepare request body
	requestBody := map[string]interface{}{}

	// Call Waldur API to update resource
	var result map[string]interface{}

	err := r.client.Update(ctx, "<no value>", data.UUID.ValueString(), requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update MarketplaceOrder",
			"An error occurred while updating the marketplace_order: "+err.Error(),
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
	if val, ok := sourceMap["accepting_terms_of_service"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.AcceptingTermsOfService = types.BoolValue(b)
		}
	} else {
		if data.AcceptingTermsOfService.IsUnknown() {
			data.AcceptingTermsOfService = types.BoolNull()
		}
	}
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
	if val, ok := sourceMap["category_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CategoryUuid = types.StringValue(str)
		}
	} else {
		if data.CategoryUuid.IsUnknown() {
			data.CategoryUuid = types.StringNull()
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
	if val, ok := sourceMap["created"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
	} else {
		if data.Created.IsUnknown() {
			data.Created = types.StringNull()
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
	if val, ok := sourceMap["modified"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Modified = types.StringValue(str)
		}
	} else {
		if data.Modified.IsUnknown() {
			data.Modified = types.StringNull()
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
	if val, ok := sourceMap["offering"]; ok && val != nil {
		if str, ok := val.(string); ok {
			// Normalize URL to UUID
			parts := strings.Split(strings.TrimRight(str, "/"), "/")
			uuid := parts[len(parts)-1]
			data.Offering = types.StringValue(uuid)
		} else {
			data.Offering = types.StringNull()
		}
	} else {
		if data.Offering.IsUnknown() {
			data.Offering = types.StringNull()
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
	if val, ok := sourceMap["offering_type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.OfferingType = types.StringValue(str)
		}
	} else {
		if data.OfferingType.IsUnknown() {
			data.OfferingType = types.StringNull()
		}
	}
	if val, ok := sourceMap["offering_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.OfferingUuid = types.StringValue(str)
		}
	} else {
		if data.OfferingUuid.IsUnknown() {
			data.OfferingUuid = types.StringNull()
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
	if val, ok := sourceMap["project"]; ok && val != nil {
		if str, ok := val.(string); ok {
			// Normalize URL to UUID
			parts := strings.Split(strings.TrimRight(str, "/"), "/")
			uuid := parts[len(parts)-1]
			data.Project = types.StringValue(uuid)
		} else {
			data.Project = types.StringNull()
		}
	} else {
		if data.Project.IsUnknown() {
			data.Project = types.StringNull()
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
	if val, ok := sourceMap["provider_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProviderUuid = types.StringValue(str)
		}
	} else {
		if data.ProviderUuid.IsUnknown() {
			data.ProviderUuid = types.StringNull()
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
	if val, ok := sourceMap["resource_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ResourceUuid = types.StringValue(str)
		}
	} else {
		if data.ResourceUuid.IsUnknown() {
			data.ResourceUuid = types.StringNull()
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
	if val, ok := sourceMap["state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.State = types.StringValue(str)
		}
	} else {
		if data.State.IsUnknown() {
			data.State = types.StringNull()
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

	// Map filter parameters from response if available

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MarketplaceOrderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MarketplaceOrderResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Call Waldur API to delete resource
	err := r.client.DeleteByUUID(ctx, "/api/marketplace-orders/{uuid}/", data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete MarketplaceOrder",
			"An error occurred while deleting the marketplace_order: "+err.Error(),
		)
		return
	}
}

func (r *MarketplaceOrderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
