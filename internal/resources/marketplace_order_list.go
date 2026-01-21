package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
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
				Description: "Can approve as consumer",
				Optional:    true,
			},
			"can_approve_as_provider": schema.BoolAttribute{
				Description: "Can approve as provider",
				Optional:    true,
			},
			"category_uuid": schema.StringAttribute{
				Description: "Category UUID",
				Optional:    true,
			},
			"created": schema.StringAttribute{
				Description: "Created after",
				Optional:    true,
			},
			"customer_uuid": schema.StringAttribute{
				Description: "Customer UUID",
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
				Description: "Offering UUID",
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
				Description: "Project UUID",
				Optional:    true,
			},
			"provider_uuid": schema.StringAttribute{
				Description: "Provider UUID",
				Optional:    true,
			},
			"query": schema.StringAttribute{
				Description: "Search by order UUID, slug, project name or resource name",
				Optional:    true,
			},
			"resource": schema.StringAttribute{
				Description: "Resource URL",
				Optional:    true,
			},
			"resource_name": schema.StringAttribute{
				Description: "Resource name",
				Optional:    true,
			},
			"resource_uuid": schema.StringAttribute{
				Description: "Resource UUID",
				Optional:    true,
			},
			"service_manager_uuid": schema.StringAttribute{
				Description: "Service manager UUID",
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
	ResourceName         types.String `tfsdk:"resource_name"`
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
	if !config.ResourceName.IsNull() && !config.ResourceName.IsUnknown() {
		filters["resource_name"] = config.ResourceName.ValueString()
	}
	if !config.ResourceUuid.IsNull() && !config.ResourceUuid.IsUnknown() {
		filters["resource_uuid"] = config.ResourceUuid.ValueString()
	}
	if !config.ServiceManagerUuid.IsNull() && !config.ServiceManagerUuid.IsUnknown() {
		filters["service_manager_uuid"] = config.ServiceManagerUuid.ValueString()
	}

	// Call API
	var listResult []MarketplaceOrderApiResponse
	err := l.client.ListWithFilter(ctx, "/api/marketplace-orders/", filters, &listResult)
	if err != nil {
		// Return error diagnostics
		resp.AddError("Failed to list resources", err.Error())
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Stream results
	stream.Results = func(push func(list.ListResult) bool) {
		for _, apiResp := range listResult {
			result := req.NewListResult(ctx)

			// Map item to model
			var data MarketplaceOrderResourceModel
			model := &data

			var diags diag.Diagnostics

			model.UUID = types.StringPointerValue(apiResp.UUID)
			model.AcceptingTermsOfService = types.BoolPointerValue(apiResp.AcceptingTermsOfService)
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

			// Set the resource state
			// For ListResource, we generally return the "Resource" state matching the main resource schema.
			// This allows `terraform plan` to correlate if using `terraform query`.

			setDiags := result.Resource.Set(ctx, &data)
			diags.Append(setDiags...)
			result.Diagnostics.Append(diags...)

			// Set Identity if possible (usually UUID)
			if !data.UUID.IsNull() && !data.UUID.IsUnknown() {
				result.Diagnostics.Append(result.Identity.Set(ctx, data.UUID.ValueString())...)
			}

			if !push(result) {
				return
			}
		}
	}
}
