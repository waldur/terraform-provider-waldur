package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-provider-waldur/internal/client"
)

var _ list.ListResource = &MarketplaceOrderList{}

type MarketplaceOrderList struct {
	client *client.Client
}

func NewMarketplaceOrderList() list.ListResource {
	return &MarketplaceOrderList{}
}

func (l *MarketplaceOrderList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_order"
}

func (l *MarketplaceOrderList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"can_approve_as_consumer": schema.BoolAttribute{
				Description: "",
				Optional:    true,
			},
			"can_approve_as_provider": schema.BoolAttribute{
				Description: "",
				Optional:    true,
			},
			"category_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"created": schema.StringAttribute{
				Description: "Created after",
				Optional:    true,
			},
			"customer_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"modified": schema.StringAttribute{
				Description: "Modified after",
				Optional:    true,
			},
			"offering": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"offering_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"page": schema.Int64Attribute{
				Description: "A page number within the paginated result set.",
				Optional:    true,
			},
			"page_size": schema.Int64Attribute{
				Description: "Number of results to return per page.",
				Optional:    true,
			},
			"parent_offering_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"project_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"provider_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"query": schema.StringAttribute{
				Description: "Search by order UUID, slug, project name or resource name",
				Optional:    true,
			},
			"resource": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"resource_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"service_manager_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
		},
	}
}

func (l *MarketplaceOrderList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected List Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T", req.ProviderData),
		)
		return
	}

	l.client = client
}

type MarketplaceOrderListModel struct {
	CanApproveAsConsumer types.Bool   `tfsdk:"can_approve_as_consumer"`
	CanApproveAsProvider types.Bool   `tfsdk:"can_approve_as_provider"`
	CategoryUuid         types.String `tfsdk:"category_uuid"`
	Created              types.String `tfsdk:"created"`
	CustomerUuid         types.String `tfsdk:"customer_uuid"`
	Modified             types.String `tfsdk:"modified"`
	Offering             types.String `tfsdk:"offering"`
	OfferingUuid         types.String `tfsdk:"offering_uuid"`
	Page                 types.Int64  `tfsdk:"page"`
	PageSize             types.Int64  `tfsdk:"page_size"`
	ParentOfferingUuid   types.String `tfsdk:"parent_offering_uuid"`
	ProjectUuid          types.String `tfsdk:"project_uuid"`
	ProviderUuid         types.String `tfsdk:"provider_uuid"`
	Query                types.String `tfsdk:"query"`
	Resource             types.String `tfsdk:"resource"`
	ResourceUuid         types.String `tfsdk:"resource_uuid"`
	ServiceManagerUuid   types.String `tfsdk:"service_manager_uuid"`
}

func (l *MarketplaceOrderList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config MarketplaceOrderListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := make(map[string]string)
	if !config.CanApproveAsConsumer.IsNull() && !config.CanApproveAsConsumer.IsUnknown() {
		filters["can_approve_as_consumer"] = fmt.Sprintf("%t", config.CanApproveAsConsumer.ValueBool())
	}
	if !config.CanApproveAsProvider.IsNull() && !config.CanApproveAsProvider.IsUnknown() {
		filters["can_approve_as_provider"] = fmt.Sprintf("%t", config.CanApproveAsProvider.ValueBool())
	}
	if !config.CategoryUuid.IsNull() && !config.CategoryUuid.IsUnknown() {
		filters["category_uuid"] = config.CategoryUuid.ValueString()
	}
	if !config.Created.IsNull() && !config.Created.IsUnknown() {
		filters["created"] = config.Created.ValueString()
	}
	if !config.CustomerUuid.IsNull() && !config.CustomerUuid.IsUnknown() {
		filters["customer_uuid"] = config.CustomerUuid.ValueString()
	}
	if !config.Modified.IsNull() && !config.Modified.IsUnknown() {
		filters["modified"] = config.Modified.ValueString()
	}
	if !config.Offering.IsNull() && !config.Offering.IsUnknown() {
		filters["offering"] = config.Offering.ValueString()
	}
	if !config.OfferingUuid.IsNull() && !config.OfferingUuid.IsUnknown() {
		filters["offering_uuid"] = config.OfferingUuid.ValueString()
	}
	if !config.Page.IsNull() && !config.Page.IsUnknown() {
		filters["page"] = fmt.Sprintf("%d", config.Page.ValueInt64())
	}
	if !config.PageSize.IsNull() && !config.PageSize.IsUnknown() {
		filters["page_size"] = fmt.Sprintf("%d", config.PageSize.ValueInt64())
	}
	if !config.ParentOfferingUuid.IsNull() && !config.ParentOfferingUuid.IsUnknown() {
		filters["parent_offering_uuid"] = config.ParentOfferingUuid.ValueString()
	}
	if !config.ProjectUuid.IsNull() && !config.ProjectUuid.IsUnknown() {
		filters["project_uuid"] = config.ProjectUuid.ValueString()
	}
	if !config.ProviderUuid.IsNull() && !config.ProviderUuid.IsUnknown() {
		filters["provider_uuid"] = config.ProviderUuid.ValueString()
	}
	if !config.Query.IsNull() && !config.Query.IsUnknown() {
		filters["query"] = config.Query.ValueString()
	}
	if !config.Resource.IsNull() && !config.Resource.IsUnknown() {
		filters["resource"] = config.Resource.ValueString()
	}
	if !config.ResourceUuid.IsNull() && !config.ResourceUuid.IsUnknown() {
		filters["resource_uuid"] = config.ResourceUuid.ValueString()
	}
	if !config.ServiceManagerUuid.IsNull() && !config.ServiceManagerUuid.IsUnknown() {
		filters["service_manager_uuid"] = config.ServiceManagerUuid.ValueString()
	}

	// Call API
	var listResult []map[string]interface{}
	err := l.client.ListWithFilter(ctx, "/api/marketplace-orders/", filters, &listResult)
	if err != nil {
		// Return error diagnostics
		resp.AddError("Failed to list resources", err.Error())
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Stream results
	stream.Results = func(push func(list.ListResult) bool) {
		for _, item := range listResult {
			result := req.NewListResult(ctx)

			// Map item to model
			var data MarketplaceOrderResourceModel

			// Initialize lists and objects to empty/null if needed (or rely on map logic)

			sourceMap := item

			// Reuse the mapResponseFields logic (embedded here or via shared template)
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

			// Set the resource state
			// For ListResource, we generally return the "Resource" state matching the main resource schema.
			// This allows `terraform plan` to correlate if using `terraform query`.

			diags := result.Resource.Set(ctx, &data)
			result.Diagnostics.Append(diags...)

			// Set Identity if possible (usually UUID)
			if !data.UUID.IsNull() && !data.UUID.IsUnknown() {
				// Identity value must match what the resource uses for Import?
				// Typically implicit. For now just setting Resource is key.
				// result.Identity.Set(ctx, data.UUID.ValueString())
				// The doc says: "An error is returned if a list result in the stream contains a null identity"
				result.Diagnostics.Append(result.Identity.Set(ctx, data.UUID.ValueString())...)
			} else {
				// Try to fallback to "uuid" from map if model failed
				if uuid, ok := item["uuid"].(string); ok {
					result.Diagnostics.Append(result.Identity.Set(ctx, uuid)...)
				}
			}

			if !push(result) {
				return
			}
		}
	}
}
