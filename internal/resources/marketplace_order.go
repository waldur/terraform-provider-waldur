package resources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
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

// MarketplaceOrderApiResponse is the API response model.
type MarketplaceOrderApiResponse struct {
	UUID *string `json:"uuid"`

	AcceptingTermsOfService    *bool                  `json:"accepting_terms_of_service" tfsdk:"accepting_terms_of_service"`
	ActivationPrice            *float64               `json:"activation_price" tfsdk:"activation_price"`
	Attachment                 *string                `json:"attachment" tfsdk:"attachment"`
	Attributes                 map[string]interface{} `json:"attributes" tfsdk:"attributes"`
	BackendId                  *string                `json:"backend_id" tfsdk:"backend_id"`
	CallbackUrl                *string                `json:"callback_url" tfsdk:"callback_url"`
	CanTerminate               *bool                  `json:"can_terminate" tfsdk:"can_terminate"`
	CategoryIcon               *string                `json:"category_icon" tfsdk:"category_icon"`
	CategoryTitle              *string                `json:"category_title" tfsdk:"category_title"`
	CategoryUuid               *string                `json:"category_uuid" tfsdk:"category_uuid"`
	CompletedAt                *string                `json:"completed_at" tfsdk:"completed_at"`
	ConsumerReviewedAt         *string                `json:"consumer_reviewed_at" tfsdk:"consumer_reviewed_at"`
	ConsumerReviewedBy         *string                `json:"consumer_reviewed_by" tfsdk:"consumer_reviewed_by"`
	ConsumerReviewedByFullName *string                `json:"consumer_reviewed_by_full_name" tfsdk:"consumer_reviewed_by_full_name"`
	ConsumerReviewedByUsername *string                `json:"consumer_reviewed_by_username" tfsdk:"consumer_reviewed_by_username"`
	Cost                       *string                `json:"cost" tfsdk:"cost"`
	Created                    *string                `json:"created" tfsdk:"created"`
	CreatedByCivilNumber       *string                `json:"created_by_civil_number" tfsdk:"created_by_civil_number"`
	CreatedByFullName          *string                `json:"created_by_full_name" tfsdk:"created_by_full_name"`
	CreatedByUsername          *string                `json:"created_by_username" tfsdk:"created_by_username"`
	CustomerSlug               *string                `json:"customer_slug" tfsdk:"customer_slug"`
	ErrorMessage               *string                `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback             *string                `json:"error_traceback" tfsdk:"error_traceback"`
	FixedPrice                 *float64               `json:"fixed_price" tfsdk:"fixed_price"`
	Modified                   *string                `json:"modified" tfsdk:"modified"`
	NewCostEstimate            *string                `json:"new_cost_estimate" tfsdk:"new_cost_estimate"`
	NewPlanName                *string                `json:"new_plan_name" tfsdk:"new_plan_name"`
	NewPlanUuid                *string                `json:"new_plan_uuid" tfsdk:"new_plan_uuid"`
	Offering                   *string                `json:"offering" tfsdk:"offering"`
	OfferingBillable           *bool                  `json:"offering_billable" tfsdk:"offering_billable"`
	OfferingDescription        *string                `json:"offering_description" tfsdk:"offering_description"`
	OfferingImage              *string                `json:"offering_image" tfsdk:"offering_image"`
	OfferingName               *string                `json:"offering_name" tfsdk:"offering_name"`
	OfferingShared             *bool                  `json:"offering_shared" tfsdk:"offering_shared"`
	OfferingThumbnail          *string                `json:"offering_thumbnail" tfsdk:"offering_thumbnail"`
	OfferingType               *string                `json:"offering_type" tfsdk:"offering_type"`
	OfferingUuid               *string                `json:"offering_uuid" tfsdk:"offering_uuid"`
	OldCostEstimate            *float64               `json:"old_cost_estimate" tfsdk:"old_cost_estimate"`
	OldPlanName                *string                `json:"old_plan_name" tfsdk:"old_plan_name"`
	OldPlanUuid                *string                `json:"old_plan_uuid" tfsdk:"old_plan_uuid"`
	OrderSubtype               *string                `json:"order_subtype" tfsdk:"order_subtype"`
	Output                     *string                `json:"output" tfsdk:"output"`
	Plan                       *string                `json:"plan" tfsdk:"plan"`
	PlanDescription            *string                `json:"plan_description" tfsdk:"plan_description"`
	PlanName                   *string                `json:"plan_name" tfsdk:"plan_name"`
	PlanUnit                   *string                `json:"plan_unit" tfsdk:"plan_unit"`
	PlanUuid                   *string                `json:"plan_uuid" tfsdk:"plan_uuid"`
	Project                    *string                `json:"project" tfsdk:"project"`
	ProjectDescription         *string                `json:"project_description" tfsdk:"project_description"`
	ProjectSlug                *string                `json:"project_slug" tfsdk:"project_slug"`
	ProviderName               *string                `json:"provider_name" tfsdk:"provider_name"`
	ProviderReviewedAt         *string                `json:"provider_reviewed_at" tfsdk:"provider_reviewed_at"`
	ProviderReviewedBy         *string                `json:"provider_reviewed_by" tfsdk:"provider_reviewed_by"`
	ProviderReviewedByFullName *string                `json:"provider_reviewed_by_full_name" tfsdk:"provider_reviewed_by_full_name"`
	ProviderReviewedByUsername *string                `json:"provider_reviewed_by_username" tfsdk:"provider_reviewed_by_username"`
	ProviderSlug               *string                `json:"provider_slug" tfsdk:"provider_slug"`
	ProviderUuid               *string                `json:"provider_uuid" tfsdk:"provider_uuid"`
	RequestComment             *string                `json:"request_comment" tfsdk:"request_comment"`
	ResourceName               *string                `json:"resource_name" tfsdk:"resource_name"`
	ResourceType               *string                `json:"resource_type" tfsdk:"resource_type"`
	ResourceUuid               *string                `json:"resource_uuid" tfsdk:"resource_uuid"`
	Slug                       *string                `json:"slug" tfsdk:"slug"`
	StartDate                  *string                `json:"start_date" tfsdk:"start_date"`
	State                      *string                `json:"state" tfsdk:"state"`
	TerminationComment         *string                `json:"termination_comment" tfsdk:"termination_comment"`
	Type                       *string                `json:"type" tfsdk:"type"`
	Url                        *string                `json:"url" tfsdk:"url"`
}

type MarketplaceOrderAttributesResponse struct {
}

// MarketplaceOrderResourceModel describes the resource data model.
type MarketplaceOrderResourceModel struct {
	UUID                       types.String   `tfsdk:"id"`
	AcceptingTermsOfService    types.Bool     `tfsdk:"accepting_terms_of_service"`
	ActivationPrice            types.Float64  `tfsdk:"activation_price"`
	Attachment                 types.String   `tfsdk:"attachment"`
	Attributes                 types.Map      `tfsdk:"attributes"`
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
	OldCostEstimate            types.Float64  `tfsdk:"old_cost_estimate"`
	OldPlanName                types.String   `tfsdk:"old_plan_name"`
	OldPlanUuid                types.String   `tfsdk:"old_plan_uuid"`
	OrderSubtype               types.String   `tfsdk:"order_subtype"`
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
	Slug                       types.String   `tfsdk:"slug"`
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
		MarkdownDescription: "Marketplace Order resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"accepting_terms_of_service": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: " ",
			},
			"activation_price": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"attachment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: " ",
			},
			"attributes": schema.MapAttribute{
				ElementType: types.StringType,
				Required:    true,
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Order attributes",
			},
			"backend_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: " ",
			},
			"callback_url": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: " ",
			},
			"can_terminate": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"category_icon": schema.StringAttribute{
				Computed:            true,
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
			"completed_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"consumer_reviewed_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"consumer_reviewed_by": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
			},
			"consumer_reviewed_by_full_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"consumer_reviewed_by_username": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
			},
			"cost": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"created_by_civil_number": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"created_by_full_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"created_by_username": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
			},
			"customer_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"fixed_price": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"new_cost_estimate": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"new_plan_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"new_plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"offering": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: " ",
			},
			"offering_billable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Purchase and usage is invoiced.",
			},
			"offering_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"offering_image": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"offering_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"offering_shared": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Accessible to all customers.",
			},
			"offering_thumbnail": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"offering_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"offering_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"old_cost_estimate": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"old_plan_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"old_plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"order_subtype": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"output": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"plan": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: " ",
			},
			"plan_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"plan_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"plan_unit": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"project": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: " ",
			},
			"project_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"project_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"provider_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"provider_reviewed_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"provider_reviewed_by": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
			},
			"provider_reviewed_by_full_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"provider_reviewed_by_username": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Required. 128 characters or fewer. Lowercase letters, numbers and @/./+/-/_ characters",
			},
			"provider_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"provider_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"request_comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: " ",
			},
			"resource_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"resource_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
			},
			"start_date": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Enables delayed processing of resource provisioning order.",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"termination_comment": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: " ",
			},
			"url": schema.StringAttribute{
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
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to create resource
	var apiResp MarketplaceOrderApiResponse // Prepare request body
	requestBody := map[string]interface{}{}
	if !data.AcceptingTermsOfService.IsNull() && !data.AcceptingTermsOfService.IsUnknown() {
		requestBody["accepting_terms_of_service"] = data.AcceptingTermsOfService.ValueBool()
	}
	if v := ConvertTFValue(data.Attributes); v != nil {
		requestBody["attributes"] = v
	}
	if !data.CallbackUrl.IsNull() && !data.CallbackUrl.IsUnknown() {
		if v := data.CallbackUrl.ValueString(); v != "" {
			requestBody["callback_url"] = v
		}
	}
	requestBody["offering"] = data.Offering.ValueString()
	if !data.Plan.IsNull() && !data.Plan.IsUnknown() {
		if v := data.Plan.ValueString(); v != "" {
			requestBody["plan"] = v
		}
	}
	requestBody["project"] = data.Project.ValueString()
	if !data.RequestComment.IsNull() && !data.RequestComment.IsUnknown() {
		if v := data.RequestComment.ValueString(); v != "" {
			requestBody["request_comment"] = v
		}
	}
	if !data.Slug.IsNull() && !data.Slug.IsUnknown() {
		if v := data.Slug.ValueString(); v != "" {
			requestBody["slug"] = v
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
	err := r.client.Create(ctx, "/api/marketplace-orders/", requestBody, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Marketplace Order",
			"An error occurred while creating the Marketplace Order: "+err.Error(),
		)
		return
	}
	// Extract UUID from response
	if apiResp.UUID != nil {
		data.UUID = types.StringPointerValue(apiResp.UUID)
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

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

	retrievePath := strings.Replace("/api/marketplace-orders/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	var apiResp MarketplaceOrderApiResponse
	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &apiResp)
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Marketplace Order",
			"An error occurred while reading the Marketplace Order: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MarketplaceOrderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update Not Supported", "This resource cannot be updated via the API.")
}

func (r *MarketplaceOrderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MarketplaceOrderResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to delete resource
	err := r.client.DeleteByUUID(ctx, "/api/marketplace-orders/{uuid}/", data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Marketplace Order",
			"An error occurred while deleting the Marketplace Order: "+err.Error(),
		)
		return
	}
}

func (r *MarketplaceOrderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *MarketplaceOrderResource) mapResponseToModel(ctx context.Context, apiResp MarketplaceOrderApiResponse, model *MarketplaceOrderResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AcceptingTermsOfService = types.BoolPointerValue(apiResp.AcceptingTermsOfService)
	model.ActivationPrice = types.Float64PointerValue(apiResp.ActivationPrice)
	model.Attachment = types.StringPointerValue(apiResp.Attachment)
	if apiResp.Attributes != nil {
		mapValAttributes, mapDiagsAttributes := types.MapValueFrom(ctx, types.StringType, apiResp.Attributes)
		diags.Append(mapDiagsAttributes...)
		model.Attributes = mapValAttributes
	} else {
		model.Attributes = types.MapNull(types.StringType)
	}
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.CallbackUrl = types.StringPointerValue(apiResp.CallbackUrl)
	model.CanTerminate = types.BoolPointerValue(apiResp.CanTerminate)
	model.CategoryIcon = types.StringPointerValue(apiResp.CategoryIcon)
	model.CategoryTitle = types.StringPointerValue(apiResp.CategoryTitle)
	model.CategoryUuid = types.StringPointerValue(apiResp.CategoryUuid)
	model.CompletedAt = types.StringPointerValue(apiResp.CompletedAt)
	model.ConsumerReviewedAt = types.StringPointerValue(apiResp.ConsumerReviewedAt)
	model.ConsumerReviewedBy = types.StringPointerValue(apiResp.ConsumerReviewedBy)
	model.ConsumerReviewedByFullName = types.StringPointerValue(apiResp.ConsumerReviewedByFullName)
	model.ConsumerReviewedByUsername = types.StringPointerValue(apiResp.ConsumerReviewedByUsername)
	model.Cost = types.StringPointerValue(apiResp.Cost)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.CreatedByCivilNumber = types.StringPointerValue(apiResp.CreatedByCivilNumber)
	model.CreatedByFullName = types.StringPointerValue(apiResp.CreatedByFullName)
	model.CreatedByUsername = types.StringPointerValue(apiResp.CreatedByUsername)
	model.CustomerSlug = types.StringPointerValue(apiResp.CustomerSlug)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.FixedPrice = types.Float64PointerValue(apiResp.FixedPrice)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.NewCostEstimate = types.StringPointerValue(apiResp.NewCostEstimate)
	model.NewPlanName = types.StringPointerValue(apiResp.NewPlanName)
	model.NewPlanUuid = types.StringPointerValue(apiResp.NewPlanUuid)
	model.Offering = types.StringPointerValue(apiResp.Offering)
	model.OfferingBillable = types.BoolPointerValue(apiResp.OfferingBillable)
	model.OfferingDescription = types.StringPointerValue(apiResp.OfferingDescription)
	model.OfferingImage = types.StringPointerValue(apiResp.OfferingImage)
	model.OfferingName = types.StringPointerValue(apiResp.OfferingName)
	model.OfferingShared = types.BoolPointerValue(apiResp.OfferingShared)
	model.OfferingThumbnail = types.StringPointerValue(apiResp.OfferingThumbnail)
	model.OfferingType = types.StringPointerValue(apiResp.OfferingType)
	model.OfferingUuid = types.StringPointerValue(apiResp.OfferingUuid)
	model.OldCostEstimate = types.Float64PointerValue(apiResp.OldCostEstimate)
	model.OldPlanName = types.StringPointerValue(apiResp.OldPlanName)
	model.OldPlanUuid = types.StringPointerValue(apiResp.OldPlanUuid)
	model.OrderSubtype = types.StringPointerValue(apiResp.OrderSubtype)
	model.Output = types.StringPointerValue(apiResp.Output)
	model.Plan = types.StringPointerValue(apiResp.Plan)
	model.PlanDescription = types.StringPointerValue(apiResp.PlanDescription)
	model.PlanName = types.StringPointerValue(apiResp.PlanName)
	model.PlanUnit = types.StringPointerValue(apiResp.PlanUnit)
	model.PlanUuid = types.StringPointerValue(apiResp.PlanUuid)
	model.Project = types.StringPointerValue(apiResp.Project)
	model.ProjectDescription = types.StringPointerValue(apiResp.ProjectDescription)
	model.ProjectSlug = types.StringPointerValue(apiResp.ProjectSlug)
	model.ProviderName = types.StringPointerValue(apiResp.ProviderName)
	model.ProviderReviewedAt = types.StringPointerValue(apiResp.ProviderReviewedAt)
	model.ProviderReviewedBy = types.StringPointerValue(apiResp.ProviderReviewedBy)
	model.ProviderReviewedByFullName = types.StringPointerValue(apiResp.ProviderReviewedByFullName)
	model.ProviderReviewedByUsername = types.StringPointerValue(apiResp.ProviderReviewedByUsername)
	model.ProviderSlug = types.StringPointerValue(apiResp.ProviderSlug)
	model.ProviderUuid = types.StringPointerValue(apiResp.ProviderUuid)
	model.RequestComment = types.StringPointerValue(apiResp.RequestComment)
	model.ResourceName = types.StringPointerValue(apiResp.ResourceName)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.ResourceUuid = types.StringPointerValue(apiResp.ResourceUuid)
	model.Slug = types.StringPointerValue(apiResp.Slug)
	model.StartDate = types.StringPointerValue(apiResp.StartDate)
	model.State = types.StringPointerValue(apiResp.State)
	model.TerminationComment = types.StringPointerValue(apiResp.TerminationComment)
	model.Type = types.StringPointerValue(apiResp.Type)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
