package resource

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

var _ list.ListResource = &MarketplaceResourceList{}

type MarketplaceResourceList struct {
	client *Client
}

func NewMarketplaceResourceList() list.ListResource {
	return &MarketplaceResourceList{}
}

func (l *MarketplaceResourceList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_resource"
}

func (l *MarketplaceResourceList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"backend_id": schema.StringAttribute{
				Description: "Backend ID",
				Optional:    true,
			},
			"category_uuid": schema.StringAttribute{
				Description: "Category UUID",
				Optional:    true,
			},
			"component_count": schema.Float64Attribute{
				Description: "Filter by exact number of components",
				Optional:    true,
			},
			"created": schema.StringAttribute{
				Description: "Created after",
				Optional:    true,
			},
			"customer": schema.StringAttribute{
				Description: "Customer URL",
				Optional:    true,
			},
			"customer_uuid": schema.StringAttribute{
				Description: "Customer UUID",
				Optional:    true,
			},
			"downscaled": schema.BoolAttribute{
				Description: "Downscaled",
				Optional:    true,
			},
			"has_terminate_date": schema.BoolAttribute{
				Description: "Has termination date",
				Optional:    true,
			},
			"is_attached": schema.BoolAttribute{
				Description: "Filter by attached state",
				Optional:    true,
			},
			"lexis_links_supported": schema.BoolAttribute{
				Description: "LEXIS links supported",
				Optional:    true,
			},
			"limit_based": schema.BoolAttribute{
				Description: "Filter by limit-based offerings",
				Optional:    true,
			},
			"limit_component_count": schema.Float64Attribute{
				Description: "Filter by exact number of limit-based components",
				Optional:    true,
			},
			"modified": schema.StringAttribute{
				Description: "Modified after",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name",
				Optional:    true,
			},
			"name_exact": schema.StringAttribute{
				Description: "Name (exact)",
				Optional:    true,
			},
			"offering": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"offering_billable": schema.BoolAttribute{
				Description: "Offering billable",
				Optional:    true,
			},
			"offering_shared": schema.BoolAttribute{
				Description: "Offering shared",
				Optional:    true,
			},
			"offering_type": schema.StringAttribute{
				Description: "Offering type",
				Optional:    true,
			},
			"only_limit_based": schema.BoolAttribute{
				Description: "Filter resources with only limit-based components",
				Optional:    true,
			},
			"only_usage_based": schema.BoolAttribute{
				Description: "Filter resources with only usage-based components",
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
			"paused": schema.BoolAttribute{
				Description: "Paused",
				Optional:    true,
			},
			"plan_uuid": schema.StringAttribute{
				Description: "Plan UUID",
				Optional:    true,
			},
			"project_name": schema.StringAttribute{
				Description: "Project name",
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
				Description: "Search by resource UUID, name, slug, backend ID, effective ID, IPs or hypervisor",
				Optional:    true,
			},
			"restrict_member_access": schema.BoolAttribute{
				Description: "Restrict member access",
				Optional:    true,
			},
			"runtime_state": schema.StringAttribute{
				Description: "Runtime state",
				Optional:    true,
			},
			"service_manager_uuid": schema.StringAttribute{
				Description: "Service manager UUID",
				Optional:    true,
			},
			"usage_based": schema.BoolAttribute{
				Description: "Filter by usage-based offerings",
				Optional:    true,
			},
			"visible_to_providers": schema.BoolAttribute{
				Description: "Include only resources visible to service providers",
				Optional:    true,
			},
			"visible_to_username": schema.StringAttribute{
				Description: "Visible to username",
				Optional:    true,
			},
		},
	}
}

func (l *MarketplaceResourceList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	l.client = NewClient(client)
}

type MarketplaceResourceListModel struct {
	BackendId            types.String  `tfsdk:"backend_id"`
	CategoryUuid         types.String  `tfsdk:"category_uuid"`
	ComponentCount       types.Float64 `tfsdk:"component_count"`
	Created              types.String  `tfsdk:"created"`
	Customer             types.String  `tfsdk:"customer"`
	CustomerUuid         types.String  `tfsdk:"customer_uuid"`
	Downscaled           types.Bool    `tfsdk:"downscaled"`
	HasTerminateDate     types.Bool    `tfsdk:"has_terminate_date"`
	IsAttached           types.Bool    `tfsdk:"is_attached"`
	LexisLinksSupported  types.Bool    `tfsdk:"lexis_links_supported"`
	LimitBased           types.Bool    `tfsdk:"limit_based"`
	LimitComponentCount  types.Float64 `tfsdk:"limit_component_count"`
	Modified             types.String  `tfsdk:"modified"`
	Name                 types.String  `tfsdk:"name"`
	NameExact            types.String  `tfsdk:"name_exact"`
	Offering             types.String  `tfsdk:"offering"`
	OfferingBillable     types.Bool    `tfsdk:"offering_billable"`
	OfferingShared       types.Bool    `tfsdk:"offering_shared"`
	OfferingType         types.String  `tfsdk:"offering_type"`
	OnlyLimitBased       types.Bool    `tfsdk:"only_limit_based"`
	OnlyUsageBased       types.Bool    `tfsdk:"only_usage_based"`
	Page                 types.Int64   `tfsdk:"page"`
	PageSize             types.Int64   `tfsdk:"page_size"`
	ParentOfferingUuid   types.String  `tfsdk:"parent_offering_uuid"`
	Paused               types.Bool    `tfsdk:"paused"`
	PlanUuid             types.String  `tfsdk:"plan_uuid"`
	ProjectName          types.String  `tfsdk:"project_name"`
	ProjectUuid          types.String  `tfsdk:"project_uuid"`
	ProviderUuid         types.String  `tfsdk:"provider_uuid"`
	Query                types.String  `tfsdk:"query"`
	RestrictMemberAccess types.Bool    `tfsdk:"restrict_member_access"`
	RuntimeState         types.String  `tfsdk:"runtime_state"`
	ServiceManagerUuid   types.String  `tfsdk:"service_manager_uuid"`
	UsageBased           types.Bool    `tfsdk:"usage_based"`
	VisibleToProviders   types.Bool    `tfsdk:"visible_to_providers"`
	VisibleToUsername    types.String  `tfsdk:"visible_to_username"`
}

func (l *MarketplaceResourceList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config MarketplaceResourceListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := make(map[string]string)
	if !config.BackendId.IsNull() && !config.BackendId.IsUnknown() {
		filters["backend_id"] = config.BackendId.ValueString()
	}
	if !config.CategoryUuid.IsNull() && !config.CategoryUuid.IsUnknown() {
		filters["category_uuid"] = config.CategoryUuid.ValueString()
	}
	if !config.ComponentCount.IsNull() && !config.ComponentCount.IsUnknown() {
		filters["component_count"] = fmt.Sprintf("%f", config.ComponentCount.ValueFloat64())
	}
	if !config.Created.IsNull() && !config.Created.IsUnknown() {
		filters["created"] = config.Created.ValueString()
	}
	if !config.Customer.IsNull() && !config.Customer.IsUnknown() {
		filters["customer"] = config.Customer.ValueString()
	}
	if !config.CustomerUuid.IsNull() && !config.CustomerUuid.IsUnknown() {
		filters["customer_uuid"] = config.CustomerUuid.ValueString()
	}
	if !config.Downscaled.IsNull() && !config.Downscaled.IsUnknown() {
		filters["downscaled"] = fmt.Sprintf("%t", config.Downscaled.ValueBool())
	}
	if !config.HasTerminateDate.IsNull() && !config.HasTerminateDate.IsUnknown() {
		filters["has_terminate_date"] = fmt.Sprintf("%t", config.HasTerminateDate.ValueBool())
	}
	if !config.IsAttached.IsNull() && !config.IsAttached.IsUnknown() {
		filters["is_attached"] = fmt.Sprintf("%t", config.IsAttached.ValueBool())
	}
	if !config.LexisLinksSupported.IsNull() && !config.LexisLinksSupported.IsUnknown() {
		filters["lexis_links_supported"] = fmt.Sprintf("%t", config.LexisLinksSupported.ValueBool())
	}
	if !config.LimitBased.IsNull() && !config.LimitBased.IsUnknown() {
		filters["limit_based"] = fmt.Sprintf("%t", config.LimitBased.ValueBool())
	}
	if !config.LimitComponentCount.IsNull() && !config.LimitComponentCount.IsUnknown() {
		filters["limit_component_count"] = fmt.Sprintf("%f", config.LimitComponentCount.ValueFloat64())
	}
	if !config.Modified.IsNull() && !config.Modified.IsUnknown() {
		filters["modified"] = config.Modified.ValueString()
	}
	if !config.Name.IsNull() && !config.Name.IsUnknown() {
		filters["name"] = config.Name.ValueString()
	}
	if !config.NameExact.IsNull() && !config.NameExact.IsUnknown() {
		filters["name_exact"] = config.NameExact.ValueString()
	}
	if !config.Offering.IsNull() && !config.Offering.IsUnknown() {
		filters["offering"] = config.Offering.ValueString()
	}
	if !config.OfferingBillable.IsNull() && !config.OfferingBillable.IsUnknown() {
		filters["offering_billable"] = fmt.Sprintf("%t", config.OfferingBillable.ValueBool())
	}
	if !config.OfferingShared.IsNull() && !config.OfferingShared.IsUnknown() {
		filters["offering_shared"] = fmt.Sprintf("%t", config.OfferingShared.ValueBool())
	}
	if !config.OfferingType.IsNull() && !config.OfferingType.IsUnknown() {
		filters["offering_type"] = config.OfferingType.ValueString()
	}
	if !config.OnlyLimitBased.IsNull() && !config.OnlyLimitBased.IsUnknown() {
		filters["only_limit_based"] = fmt.Sprintf("%t", config.OnlyLimitBased.ValueBool())
	}
	if !config.OnlyUsageBased.IsNull() && !config.OnlyUsageBased.IsUnknown() {
		filters["only_usage_based"] = fmt.Sprintf("%t", config.OnlyUsageBased.ValueBool())
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
	if !config.Paused.IsNull() && !config.Paused.IsUnknown() {
		filters["paused"] = fmt.Sprintf("%t", config.Paused.ValueBool())
	}
	if !config.PlanUuid.IsNull() && !config.PlanUuid.IsUnknown() {
		filters["plan_uuid"] = config.PlanUuid.ValueString()
	}
	if !config.ProjectName.IsNull() && !config.ProjectName.IsUnknown() {
		filters["project_name"] = config.ProjectName.ValueString()
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
	if !config.RestrictMemberAccess.IsNull() && !config.RestrictMemberAccess.IsUnknown() {
		filters["restrict_member_access"] = fmt.Sprintf("%t", config.RestrictMemberAccess.ValueBool())
	}
	if !config.RuntimeState.IsNull() && !config.RuntimeState.IsUnknown() {
		filters["runtime_state"] = config.RuntimeState.ValueString()
	}
	if !config.ServiceManagerUuid.IsNull() && !config.ServiceManagerUuid.IsUnknown() {
		filters["service_manager_uuid"] = config.ServiceManagerUuid.ValueString()
	}
	if !config.UsageBased.IsNull() && !config.UsageBased.IsUnknown() {
		filters["usage_based"] = fmt.Sprintf("%t", config.UsageBased.ValueBool())
	}
	if !config.VisibleToProviders.IsNull() && !config.VisibleToProviders.IsUnknown() {
		filters["visible_to_providers"] = fmt.Sprintf("%t", config.VisibleToProviders.ValueBool())
	}
	if !config.VisibleToUsername.IsNull() && !config.VisibleToUsername.IsUnknown() {
		filters["visible_to_username"] = config.VisibleToUsername.ValueString()
	}

	// Call API
	listResult, err := l.client.ListMarketplaceResource(ctx, filters)
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
			var data MarketplaceResourceResourceModel
			model := &data

			var diags diag.Diagnostics

			data.UUID = types.StringPointerValue(apiResp.UUID)
			model.AvailableActions, _ = types.ListValueFrom(ctx, types.StringType, apiResp.AvailableActions)
			model.BackendId = types.StringPointerValue(apiResp.BackendId)
			model.CanTerminate = types.BoolPointerValue(apiResp.CanTerminate)
			model.CategoryIcon = types.StringPointerValue(apiResp.CategoryIcon)
			model.CategoryTitle = types.StringPointerValue(apiResp.CategoryTitle)
			model.CategoryUuid = types.StringPointerValue(apiResp.CategoryUuid)
			model.Created = types.StringPointerValue(apiResp.Created)
			model.CustomerSlug = types.StringPointerValue(apiResp.CustomerSlug)
			model.Description = types.StringPointerValue(apiResp.Description)
			model.Downscaled = types.BoolPointerValue(apiResp.Downscaled)
			model.EffectiveId = types.StringPointerValue(apiResp.EffectiveId)
			model.EndDate = types.StringPointerValue(apiResp.EndDate)
			model.EndDateRequestedBy = types.StringPointerValue(apiResp.EndDateRequestedBy)
			listValEndpoints, listDiagsEndpoints := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
				"name": types.StringType,
				"url":  types.StringType,
			}}, apiResp.Endpoints)
			diags.Append(listDiagsEndpoints...)
			model.Endpoints = listValEndpoints
			model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
			model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
			model.LastSync = types.StringPointerValue(apiResp.LastSync)
			model.Modified = types.StringPointerValue(apiResp.Modified)
			model.Name = types.StringPointerValue(apiResp.Name)
			model.Offering = types.StringPointerValue(apiResp.Offering)
			model.OfferingBillable = types.BoolPointerValue(apiResp.OfferingBillable)
			listValOfferingComponents, listDiagsOfferingComponents := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
				"article_code":         types.StringType,
				"billing_type":         types.StringType,
				"default_limit":        types.Int64Type,
				"description":          types.StringType,
				"factor":               types.Int64Type,
				"is_boolean":           types.BoolType,
				"is_builtin":           types.BoolType,
				"is_prepaid":           types.BoolType,
				"limit_amount":         types.Int64Type,
				"limit_period":         types.StringType,
				"max_available_limit":  types.Int64Type,
				"max_prepaid_duration": types.Int64Type,
				"max_value":            types.Int64Type,
				"measured_unit":        types.StringType,
				"min_prepaid_duration": types.Int64Type,
				"min_value":            types.Int64Type,
				"name":                 types.StringType,
				"overage_component":    types.StringType,
				"type":                 types.StringType,
				"unit_factor":          types.Int64Type,
			}}, apiResp.OfferingComponents)
			diags.Append(listDiagsOfferingComponents...)
			model.OfferingComponents = listValOfferingComponents
			model.OfferingDescription = types.StringPointerValue(apiResp.OfferingDescription)
			model.OfferingImage = types.StringPointerValue(apiResp.OfferingImage)
			model.OfferingName = types.StringPointerValue(apiResp.OfferingName)
			model.OfferingShared = types.BoolPointerValue(apiResp.OfferingShared)
			model.OfferingSlug = types.StringPointerValue(apiResp.OfferingSlug)
			model.OfferingState = types.StringPointerValue(apiResp.OfferingState)
			model.OfferingThumbnail = types.StringPointerValue(apiResp.OfferingThumbnail)
			model.OfferingType = types.StringPointerValue(apiResp.OfferingType)
			model.OfferingUuid = types.StringPointerValue(apiResp.OfferingUuid)
			model.ParentName = types.StringPointerValue(apiResp.ParentName)
			model.ParentOfferingName = types.StringPointerValue(apiResp.ParentOfferingName)
			model.ParentOfferingSlug = types.StringPointerValue(apiResp.ParentOfferingSlug)
			model.ParentOfferingUuid = types.StringPointerValue(apiResp.ParentOfferingUuid)
			model.ParentUuid = types.StringPointerValue(apiResp.ParentUuid)
			model.Paused = types.BoolPointerValue(apiResp.Paused)
			model.Plan = types.StringPointerValue(apiResp.Plan)
			model.PlanDescription = types.StringPointerValue(apiResp.PlanDescription)
			model.PlanName = types.StringPointerValue(apiResp.PlanName)
			model.PlanUnit = types.StringPointerValue(apiResp.PlanUnit)
			model.PlanUuid = types.StringPointerValue(apiResp.PlanUuid)
			model.ProjectDescription = types.StringPointerValue(apiResp.ProjectDescription)
			model.ProjectEndDate = types.StringPointerValue(apiResp.ProjectEndDate)
			model.ProjectEndDateRequestedBy = types.StringPointerValue(apiResp.ProjectEndDateRequestedBy)
			model.ProjectSlug = types.StringPointerValue(apiResp.ProjectSlug)
			model.ProviderName = types.StringPointerValue(apiResp.ProviderName)
			model.ProviderSlug = types.StringPointerValue(apiResp.ProviderSlug)
			model.ProviderUuid = types.StringPointerValue(apiResp.ProviderUuid)
			listValReport, listDiagsReport := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
				"body":   types.StringType,
				"header": types.StringType,
			}}, apiResp.Report)
			diags.Append(listDiagsReport...)
			model.Report = listValReport
			model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
			model.ResourceUuid = types.StringPointerValue(apiResp.ResourceUuid)
			model.RestrictMemberAccess = types.BoolPointerValue(apiResp.RestrictMemberAccess)
			model.Scope = types.StringPointerValue(apiResp.Scope)
			model.Slug = types.StringPointerValue(apiResp.Slug)
			model.State = types.StringPointerValue(apiResp.State)
			model.Url = types.StringPointerValue(apiResp.Url)
			model.UserRequiresReconsent = types.BoolPointerValue(apiResp.UserRequiresReconsent)
			model.Username = types.StringPointerValue(apiResp.Username)

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
