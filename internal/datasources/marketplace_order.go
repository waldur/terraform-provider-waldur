package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
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

// MarketplaceOrderApiResponse is the API response model.
type MarketplaceOrderApiResponse struct {
	UUID *string `json:"uuid"`

	ActivationPrice            *float64 `json:"activation_price" tfsdk:"activation_price"`
	Attachment                 *string  `json:"attachment" tfsdk:"attachment"`
	BackendId                  *string  `json:"backend_id" tfsdk:"backend_id"`
	CallbackUrl                *string  `json:"callback_url" tfsdk:"callback_url"`
	CanTerminate               *bool    `json:"can_terminate" tfsdk:"can_terminate"`
	CategoryIcon               *string  `json:"category_icon" tfsdk:"category_icon"`
	CategoryTitle              *string  `json:"category_title" tfsdk:"category_title"`
	CategoryUuid               *string  `json:"category_uuid" tfsdk:"category_uuid"`
	CompletedAt                *string  `json:"completed_at" tfsdk:"completed_at"`
	ConsumerReviewedAt         *string  `json:"consumer_reviewed_at" tfsdk:"consumer_reviewed_at"`
	ConsumerReviewedBy         *string  `json:"consumer_reviewed_by" tfsdk:"consumer_reviewed_by"`
	ConsumerReviewedByFullName *string  `json:"consumer_reviewed_by_full_name" tfsdk:"consumer_reviewed_by_full_name"`
	ConsumerReviewedByUsername *string  `json:"consumer_reviewed_by_username" tfsdk:"consumer_reviewed_by_username"`
	Cost                       *string  `json:"cost" tfsdk:"cost"`
	Created                    *string  `json:"created" tfsdk:"created"`
	CreatedByCivilNumber       *string  `json:"created_by_civil_number" tfsdk:"created_by_civil_number"`
	CreatedByFullName          *string  `json:"created_by_full_name" tfsdk:"created_by_full_name"`
	CreatedByUsername          *string  `json:"created_by_username" tfsdk:"created_by_username"`
	CustomerName               *string  `json:"customer_name" tfsdk:"customer_name"`
	CustomerSlug               *string  `json:"customer_slug" tfsdk:"customer_slug"`
	CustomerUuid               *string  `json:"customer_uuid" tfsdk:"customer_uuid"`
	ErrorMessage               *string  `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback             *string  `json:"error_traceback" tfsdk:"error_traceback"`
	FixedPrice                 *float64 `json:"fixed_price" tfsdk:"fixed_price"`
	MarketplaceResourceUuid    *string  `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                   *string  `json:"modified" tfsdk:"modified"`
	NewCostEstimate            *string  `json:"new_cost_estimate" tfsdk:"new_cost_estimate"`
	NewPlanName                *string  `json:"new_plan_name" tfsdk:"new_plan_name"`
	NewPlanUuid                *string  `json:"new_plan_uuid" tfsdk:"new_plan_uuid"`
	Offering                   *string  `json:"offering" tfsdk:"offering"`
	OfferingBillable           *bool    `json:"offering_billable" tfsdk:"offering_billable"`
	OfferingDescription        *string  `json:"offering_description" tfsdk:"offering_description"`
	OfferingImage              *string  `json:"offering_image" tfsdk:"offering_image"`
	OfferingName               *string  `json:"offering_name" tfsdk:"offering_name"`
	OfferingShared             *bool    `json:"offering_shared" tfsdk:"offering_shared"`
	OfferingThumbnail          *string  `json:"offering_thumbnail" tfsdk:"offering_thumbnail"`
	OfferingType               *string  `json:"offering_type" tfsdk:"offering_type"`
	OfferingUuid               *string  `json:"offering_uuid" tfsdk:"offering_uuid"`
	OldCostEstimate            *float64 `json:"old_cost_estimate" tfsdk:"old_cost_estimate"`
	OldPlanName                *string  `json:"old_plan_name" tfsdk:"old_plan_name"`
	OldPlanUuid                *string  `json:"old_plan_uuid" tfsdk:"old_plan_uuid"`
	OrderSubtype               *string  `json:"order_subtype" tfsdk:"order_subtype"`
	Output                     *string  `json:"output" tfsdk:"output"`
	Plan                       *string  `json:"plan" tfsdk:"plan"`
	PlanDescription            *string  `json:"plan_description" tfsdk:"plan_description"`
	PlanName                   *string  `json:"plan_name" tfsdk:"plan_name"`
	PlanUnit                   *string  `json:"plan_unit" tfsdk:"plan_unit"`
	PlanUuid                   *string  `json:"plan_uuid" tfsdk:"plan_uuid"`
	ProjectDescription         *string  `json:"project_description" tfsdk:"project_description"`
	ProjectName                *string  `json:"project_name" tfsdk:"project_name"`
	ProjectSlug                *string  `json:"project_slug" tfsdk:"project_slug"`
	ProjectUuid                *string  `json:"project_uuid" tfsdk:"project_uuid"`
	ProviderName               *string  `json:"provider_name" tfsdk:"provider_name"`
	ProviderReviewedAt         *string  `json:"provider_reviewed_at" tfsdk:"provider_reviewed_at"`
	ProviderReviewedBy         *string  `json:"provider_reviewed_by" tfsdk:"provider_reviewed_by"`
	ProviderReviewedByFullName *string  `json:"provider_reviewed_by_full_name" tfsdk:"provider_reviewed_by_full_name"`
	ProviderReviewedByUsername *string  `json:"provider_reviewed_by_username" tfsdk:"provider_reviewed_by_username"`
	ProviderSlug               *string  `json:"provider_slug" tfsdk:"provider_slug"`
	ProviderUuid               *string  `json:"provider_uuid" tfsdk:"provider_uuid"`
	RequestComment             *string  `json:"request_comment" tfsdk:"request_comment"`
	ResourceName               *string  `json:"resource_name" tfsdk:"resource_name"`
	ResourceType               *string  `json:"resource_type" tfsdk:"resource_type"`
	ResourceUuid               *string  `json:"resource_uuid" tfsdk:"resource_uuid"`
	Slug                       *string  `json:"slug" tfsdk:"slug"`
	StartDate                  *string  `json:"start_date" tfsdk:"start_date"`
	State                      *string  `json:"state" tfsdk:"state"`
	TerminationComment         *string  `json:"termination_comment" tfsdk:"termination_comment"`
	Type                       *string  `json:"type" tfsdk:"type"`
	Url                        *string  `json:"url" tfsdk:"url"`
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
	ResourceName               types.String  `tfsdk:"resource_name"`
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
	CustomerName               types.String  `tfsdk:"customer_name"`
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
	OldCostEstimate            types.Float64 `tfsdk:"old_cost_estimate"`
	OldPlanName                types.String  `tfsdk:"old_plan_name"`
	OldPlanUuid                types.String  `tfsdk:"old_plan_uuid"`
	OrderSubtype               types.String  `tfsdk:"order_subtype"`
	Output                     types.String  `tfsdk:"output"`
	Plan                       types.String  `tfsdk:"plan"`
	PlanDescription            types.String  `tfsdk:"plan_description"`
	PlanName                   types.String  `tfsdk:"plan_name"`
	PlanUnit                   types.String  `tfsdk:"plan_unit"`
	PlanUuid                   types.String  `tfsdk:"plan_uuid"`
	ProjectDescription         types.String  `tfsdk:"project_description"`
	ProjectName                types.String  `tfsdk:"project_name"`
	ProjectSlug                types.String  `tfsdk:"project_slug"`
	ProviderName               types.String  `tfsdk:"provider_name"`
	ProviderReviewedAt         types.String  `tfsdk:"provider_reviewed_at"`
	ProviderReviewedBy         types.String  `tfsdk:"provider_reviewed_by"`
	ProviderReviewedByFullName types.String  `tfsdk:"provider_reviewed_by_full_name"`
	ProviderReviewedByUsername types.String  `tfsdk:"provider_reviewed_by_username"`
	ProviderSlug               types.String  `tfsdk:"provider_slug"`
	RequestComment             types.String  `tfsdk:"request_comment"`
	ResourceType               types.String  `tfsdk:"resource_type"`
	Slug                       types.String  `tfsdk:"slug"`
	StartDate                  types.String  `tfsdk:"start_date"`
	TerminationComment         types.String  `tfsdk:"termination_comment"`
	Url                        types.String  `tfsdk:"url"`
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
				MarkdownDescription: "Resource UUID",
			},
			"can_approve_as_consumer": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Can approve as consumer",
			},
			"can_approve_as_provider": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Can approve as provider",
			},
			"category_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Category UUID",
			},
			"created": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Created after",
			},
			"customer_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer UUID",
			},
			"modified": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Modified after",
			},
			"offering": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Offering",
			},
			"offering_slug": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Multiple values may be separated by commas.",
			},
			"offering_type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Offering type",
			},
			"offering_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Offering UUID",
			},
			"parent_offering_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "UUID of the parent offering",
			},
			"project_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project UUID",
			},
			"provider_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Provider UUID",
			},
			"query": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Search by order UUID, slug, project name or resource name",
			},
			"resource": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Resource URL",
			},
			"resource_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Resource name",
			},
			"resource_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Resource UUID",
			},
			"service_manager_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Service manager UUID",
			},
			"state": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Order state",
			},
			"type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Order type",
			},
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
			"completed_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Completed at",
			},
			"consumer_reviewed_at": schema.StringAttribute{
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
			"customer_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the customer",
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
			"new_cost_estimate": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "New cost estimate",
			},
			"new_plan_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the new plan",
			},
			"new_plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the new plan",
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
			"project_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the project",
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
			"request_comment": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Request comment",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource type",
			},
			"slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Slug",
			},
			"start_date": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Enables delayed processing of resource provisioning order.",
			},
			"termination_comment": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Termination comment",
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
		var apiResp MarketplaceOrderApiResponse

		err := d.client.GetByUUID(ctx, "/api/marketplace-orders/{uuid}/", data.UUID.ValueString(), &apiResp)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Marketplace Order",
				"An error occurred while reading the Marketplace Order by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, apiResp, &data)...)

	} else {
		// Filter by provided parameters
		var results []MarketplaceOrderApiResponse

		type filterDef struct {
			name string
			val  attr.Value
		}
		filterDefs := []filterDef{
			{"can_approve_as_consumer", data.CanApproveAsConsumer},
			{"can_approve_as_provider", data.CanApproveAsProvider},
			{"category_uuid", data.CategoryUuid},
			{"created", data.Created},
			{"customer_uuid", data.CustomerUuid},
			{"modified", data.Modified},
			{"offering", data.Offering},
			{"offering_slug", data.OfferingSlug},
			{"offering_type", data.OfferingType},
			{"offering_uuid", data.OfferingUuid},
			{"parent_offering_uuid", data.ParentOfferingUuid},
			{"project_uuid", data.ProjectUuid},
			{"provider_uuid", data.ProviderUuid},
			{"query", data.Query},
			{"resource", data.Resource},
			{"resource_name", data.ResourceName},
			{"resource_uuid", data.ResourceUuid},
			{"service_manager_uuid", data.ServiceManagerUuid},
			{"state", data.State},
			{"type", data.Type},
		}

		filters := make(map[string]string)
		for _, fd := range filterDefs {
			if fd.val.IsNull() || fd.val.IsUnknown() {
				continue
			}
			switch v := fd.val.(type) {
			case types.String:
				filters[fd.name] = v.ValueString()
			case types.Int64:
				filters[fd.name] = fmt.Sprintf("%d", v.ValueInt64())
			case types.Bool:
				filters[fd.name] = fmt.Sprintf("%t", v.ValueBool())
			case types.Float64:
				filters[fd.name] = fmt.Sprintf("%f", v.ValueFloat64())
			}
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

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *MarketplaceOrderDataSource) mapResponseToModel(ctx context.Context, apiResp MarketplaceOrderApiResponse, model *MarketplaceOrderDataSourceModel) diag.Diagnostics {
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
	model.CustomerName = types.StringPointerValue(apiResp.CustomerName)
	model.CustomerSlug = types.StringPointerValue(apiResp.CustomerSlug)
	model.CustomerUuid = types.StringPointerValue(apiResp.CustomerUuid)
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
	model.ProjectName = types.StringPointerValue(apiResp.ProjectName)
	model.ProjectSlug = types.StringPointerValue(apiResp.ProjectSlug)
	model.ProjectUuid = types.StringPointerValue(apiResp.ProjectUuid)
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
