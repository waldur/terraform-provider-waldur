package order

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MarketplaceOrderResource{}
var _ resource.ResourceWithImportState = &MarketplaceOrderResource{}

func NewMarketplaceOrderResource() resource.Resource {
	return &MarketplaceOrderResource{}
}

// MarketplaceOrderResource defines the resource implementation.
type MarketplaceOrderResource struct {
	client *Client
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
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Accepting terms of service",
			},
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
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Attachment",
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
				MarkdownDescription: "ID of the backend",
			},
			"callback_url": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
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
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Completed at",
			},
			"consumer_reviewed_at": schema.StringAttribute{
				Computed: true,
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
			},
			"created": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Created",
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
			"modified": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Modified",
			},
			"new_cost_estimate": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "New cost estimate",
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
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
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
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
				Computed: true,
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
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
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
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
			},
			"start_date": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
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

	r.client = NewClient(client)
}

func (r *MarketplaceOrderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MarketplaceOrderResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := MarketplaceOrderCreateRequest{
		AcceptingTermsOfService: data.AcceptingTermsOfService.ValueBoolPointer(),
		CallbackUrl:             data.CallbackUrl.ValueStringPointer(),
		Offering:                data.Offering.ValueStringPointer(),
		Plan:                    data.Plan.ValueStringPointer(),
		Project:                 data.Project.ValueStringPointer(),
		RequestComment:          data.RequestComment.ValueStringPointer(),
		Slug:                    data.Slug.ValueStringPointer(),
		StartDate:               data.StartDate.ValueStringPointer(),
		Type:                    data.Type.ValueStringPointer(),
	}
	{
		var mapItems map[string]interface{}
		if diags := data.Attributes.ElementsAs(ctx, &mapItems, false); !diags.HasError() && len(mapItems) > 0 {
			requestBody.Attributes = mapItems
		}
	}

	apiResp, err := r.client.CreateMarketplaceOrder(ctx, &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Marketplace Order",
			"An error occurred while creating the Marketplace Order: "+err.Error(),
		)
		return
	}
	data.UUID = types.StringPointerValue(apiResp.UUID)

	createTimeout, diags := data.Timeouts.Create(ctx, 30*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*MarketplaceOrderResponse, error) {
		return r.client.GetMarketplaceOrder(ctx, data.UUID.ValueString())
	}, createTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource creation", err.Error())
		return
	}
	apiResp = newResp

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

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

	apiResp, err := r.client.GetMarketplaceOrder(ctx, data.UUID.ValueString())
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

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

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

	err := r.client.DeleteMarketplaceOrder(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Marketplace Order",
			"An error occurred while deleting the Marketplace Order: "+err.Error(),
		)
		return
	}

	deleteTimeout, diags := data.Timeouts.Delete(ctx, 10*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err = common.WaitForDeletion(ctx, func(ctx context.Context) (*MarketplaceOrderResponse, error) {
		return r.client.GetMarketplaceOrder(ctx, data.UUID.ValueString())
	}, deleteTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource deletion", err.Error())
		return
	}
}

func (r *MarketplaceOrderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	uuid := req.ID
	if uuid == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Import ID cannot be empty. Please provide the UUID of the Marketplace Order.",
		)
		return
	}

	tflog.Info(ctx, "Importing Marketplace Order", map[string]interface{}{
		"uuid": uuid,
	})

	apiResp, err := r.client.GetMarketplaceOrder(ctx, uuid)
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("Marketplace Order with UUID '%s' does not exist or is not accessible.", uuid),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to Import Marketplace Order",
			fmt.Sprintf("An error occurred while fetching the Marketplace Order: %s", err.Error()),
		)
		return
	}

	var data MarketplaceOrderResourceModel
	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MarketplaceOrderResource) mapResponseToModel(ctx context.Context, apiResp MarketplaceOrderResponse, model *MarketplaceOrderResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.ActivationPrice = types.Float64PointerValue(apiResp.ActivationPrice)
	model.Attachment = types.StringPointerValue(apiResp.Attachment)
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
