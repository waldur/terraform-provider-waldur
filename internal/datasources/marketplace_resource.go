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
var _ datasource.DataSource = &MarketplaceResourceDataSource{}

func NewMarketplaceResourceDataSource() datasource.DataSource {
	return &MarketplaceResourceDataSource{}
}

// MarketplaceResourceDataSource defines the data source implementation.
type MarketplaceResourceDataSource struct {
	client *client.Client
}

// MarketplaceResourceApiResponse is the API response model.
type MarketplaceResourceApiResponse struct {
	UUID *string `json:"uuid"`

	AvailableActions          []string                                        `json:"available_actions" tfsdk:"available_actions"`
	BackendId                 *string                                         `json:"backend_id" tfsdk:"backend_id"`
	CanTerminate              *bool                                           `json:"can_terminate" tfsdk:"can_terminate"`
	CategoryIcon              *string                                         `json:"category_icon" tfsdk:"category_icon"`
	CategoryTitle             *string                                         `json:"category_title" tfsdk:"category_title"`
	CategoryUuid              *string                                         `json:"category_uuid" tfsdk:"category_uuid"`
	Created                   *string                                         `json:"created" tfsdk:"created"`
	CustomerName              *string                                         `json:"customer_name" tfsdk:"customer_name"`
	CustomerSlug              *string                                         `json:"customer_slug" tfsdk:"customer_slug"`
	CustomerUuid              *string                                         `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description               *string                                         `json:"description" tfsdk:"description"`
	Downscaled                *bool                                           `json:"downscaled" tfsdk:"downscaled"`
	EffectiveId               *string                                         `json:"effective_id" tfsdk:"effective_id"`
	EndDate                   *string                                         `json:"end_date" tfsdk:"end_date"`
	EndDateRequestedBy        *string                                         `json:"end_date_requested_by" tfsdk:"end_date_requested_by"`
	Endpoints                 []MarketplaceResourceEndpointsResponse          `json:"endpoints" tfsdk:"endpoints"`
	ErrorMessage              *string                                         `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback            *string                                         `json:"error_traceback" tfsdk:"error_traceback"`
	IsLimitBased              *bool                                           `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased              *bool                                           `json:"is_usage_based" tfsdk:"is_usage_based"`
	LastSync                  *string                                         `json:"last_sync" tfsdk:"last_sync"`
	Modified                  *string                                         `json:"modified" tfsdk:"modified"`
	Name                      *string                                         `json:"name" tfsdk:"name"`
	Offering                  *string                                         `json:"offering" tfsdk:"offering"`
	OfferingBillable          *bool                                           `json:"offering_billable" tfsdk:"offering_billable"`
	OfferingComponents        []MarketplaceResourceOfferingComponentsResponse `json:"offering_components" tfsdk:"offering_components"`
	OfferingDescription       *string                                         `json:"offering_description" tfsdk:"offering_description"`
	OfferingImage             *string                                         `json:"offering_image" tfsdk:"offering_image"`
	OfferingName              *string                                         `json:"offering_name" tfsdk:"offering_name"`
	OfferingShared            *bool                                           `json:"offering_shared" tfsdk:"offering_shared"`
	OfferingSlug              *string                                         `json:"offering_slug" tfsdk:"offering_slug"`
	OfferingState             *string                                         `json:"offering_state" tfsdk:"offering_state"`
	OfferingThumbnail         *string                                         `json:"offering_thumbnail" tfsdk:"offering_thumbnail"`
	OfferingType              *string                                         `json:"offering_type" tfsdk:"offering_type"`
	OfferingUuid              *string                                         `json:"offering_uuid" tfsdk:"offering_uuid"`
	ParentName                *string                                         `json:"parent_name" tfsdk:"parent_name"`
	ParentOfferingName        *string                                         `json:"parent_offering_name" tfsdk:"parent_offering_name"`
	ParentOfferingSlug        *string                                         `json:"parent_offering_slug" tfsdk:"parent_offering_slug"`
	ParentOfferingUuid        *string                                         `json:"parent_offering_uuid" tfsdk:"parent_offering_uuid"`
	ParentUuid                *string                                         `json:"parent_uuid" tfsdk:"parent_uuid"`
	Paused                    *bool                                           `json:"paused" tfsdk:"paused"`
	Plan                      *string                                         `json:"plan" tfsdk:"plan"`
	PlanDescription           *string                                         `json:"plan_description" tfsdk:"plan_description"`
	PlanName                  *string                                         `json:"plan_name" tfsdk:"plan_name"`
	PlanUnit                  *string                                         `json:"plan_unit" tfsdk:"plan_unit"`
	PlanUuid                  *string                                         `json:"plan_uuid" tfsdk:"plan_uuid"`
	Project                   *string                                         `json:"project" tfsdk:"project"`
	ProjectDescription        *string                                         `json:"project_description" tfsdk:"project_description"`
	ProjectEndDate            *string                                         `json:"project_end_date" tfsdk:"project_end_date"`
	ProjectEndDateRequestedBy *string                                         `json:"project_end_date_requested_by" tfsdk:"project_end_date_requested_by"`
	ProjectName               *string                                         `json:"project_name" tfsdk:"project_name"`
	ProjectSlug               *string                                         `json:"project_slug" tfsdk:"project_slug"`
	ProjectUuid               *string                                         `json:"project_uuid" tfsdk:"project_uuid"`
	ProviderName              *string                                         `json:"provider_name" tfsdk:"provider_name"`
	ProviderSlug              *string                                         `json:"provider_slug" tfsdk:"provider_slug"`
	ProviderUuid              *string                                         `json:"provider_uuid" tfsdk:"provider_uuid"`
	Report                    []MarketplaceResourceReportResponse             `json:"report" tfsdk:"report"`
	ResourceType              *string                                         `json:"resource_type" tfsdk:"resource_type"`
	ResourceUuid              *string                                         `json:"resource_uuid" tfsdk:"resource_uuid"`
	RestrictMemberAccess      *bool                                           `json:"restrict_member_access" tfsdk:"restrict_member_access"`
	Scope                     *string                                         `json:"scope" tfsdk:"scope"`
	ServiceSettingsUuid       *string                                         `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	Slug                      *string                                         `json:"slug" tfsdk:"slug"`
	State                     *string                                         `json:"state" tfsdk:"state"`
	Url                       *string                                         `json:"url" tfsdk:"url"`
	UserRequiresReconsent     *bool                                           `json:"user_requires_reconsent" tfsdk:"user_requires_reconsent"`
	Username                  *string                                         `json:"username" tfsdk:"username"`
}

type MarketplaceResourceEndpointsResponse struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
}

type MarketplaceResourceOfferingComponentsResponse struct {
	ArticleCode        *string `json:"article_code" tfsdk:"article_code"`
	BillingType        *string `json:"billing_type" tfsdk:"billing_type"`
	DefaultLimit       *int64  `json:"default_limit" tfsdk:"default_limit"`
	Description        *string `json:"description" tfsdk:"description"`
	Factor             *int64  `json:"factor" tfsdk:"factor"`
	IsBoolean          *bool   `json:"is_boolean" tfsdk:"is_boolean"`
	IsBuiltin          *bool   `json:"is_builtin" tfsdk:"is_builtin"`
	IsPrepaid          *bool   `json:"is_prepaid" tfsdk:"is_prepaid"`
	LimitAmount        *int64  `json:"limit_amount" tfsdk:"limit_amount"`
	LimitPeriod        *string `json:"limit_period" tfsdk:"limit_period"`
	MaxAvailableLimit  *int64  `json:"max_available_limit" tfsdk:"max_available_limit"`
	MaxPrepaidDuration *int64  `json:"max_prepaid_duration" tfsdk:"max_prepaid_duration"`
	MaxValue           *int64  `json:"max_value" tfsdk:"max_value"`
	MeasuredUnit       *string `json:"measured_unit" tfsdk:"measured_unit"`
	MinPrepaidDuration *int64  `json:"min_prepaid_duration" tfsdk:"min_prepaid_duration"`
	MinValue           *int64  `json:"min_value" tfsdk:"min_value"`
	Name               *string `json:"name" tfsdk:"name"`
	OverageComponent   *string `json:"overage_component" tfsdk:"overage_component"`
	Type               *string `json:"type" tfsdk:"type"`
	UnitFactor         *int64  `json:"unit_factor" tfsdk:"unit_factor"`
}

type MarketplaceResourceReportResponse struct {
	Body   *string `json:"body" tfsdk:"body"`
	Header *string `json:"header" tfsdk:"header"`
}

// MarketplaceResourceDataSourceModel describes the data source data model.
type MarketplaceResourceDataSourceModel struct {
	UUID                      types.String  `tfsdk:"id"`
	BackendId                 types.String  `tfsdk:"backend_id"`
	CategoryUuid              types.String  `tfsdk:"category_uuid"`
	ComponentCount            types.Float64 `tfsdk:"component_count"`
	Created                   types.String  `tfsdk:"created"`
	Customer                  types.String  `tfsdk:"customer"`
	CustomerUuid              types.String  `tfsdk:"customer_uuid"`
	Downscaled                types.Bool    `tfsdk:"downscaled"`
	HasTerminateDate          types.Bool    `tfsdk:"has_terminate_date"`
	IsAttached                types.Bool    `tfsdk:"is_attached"`
	LexisLinksSupported       types.Bool    `tfsdk:"lexis_links_supported"`
	LimitBased                types.Bool    `tfsdk:"limit_based"`
	LimitComponentCount       types.Float64 `tfsdk:"limit_component_count"`
	Modified                  types.String  `tfsdk:"modified"`
	Name                      types.String  `tfsdk:"name"`
	NameExact                 types.String  `tfsdk:"name_exact"`
	Offering                  types.String  `tfsdk:"offering"`
	OfferingBillable          types.Bool    `tfsdk:"offering_billable"`
	OfferingShared            types.Bool    `tfsdk:"offering_shared"`
	OfferingSlug              types.String  `tfsdk:"offering_slug"`
	OfferingType              types.String  `tfsdk:"offering_type"`
	OfferingUuid              types.String  `tfsdk:"offering_uuid"`
	OnlyLimitBased            types.Bool    `tfsdk:"only_limit_based"`
	OnlyUsageBased            types.Bool    `tfsdk:"only_usage_based"`
	OrderState                types.String  `tfsdk:"order_state"`
	ParentOfferingUuid        types.String  `tfsdk:"parent_offering_uuid"`
	Paused                    types.Bool    `tfsdk:"paused"`
	PlanUuid                  types.String  `tfsdk:"plan_uuid"`
	ProjectName               types.String  `tfsdk:"project_name"`
	ProjectUuid               types.String  `tfsdk:"project_uuid"`
	ProviderUuid              types.String  `tfsdk:"provider_uuid"`
	Query                     types.String  `tfsdk:"query"`
	RestrictMemberAccess      types.Bool    `tfsdk:"restrict_member_access"`
	RuntimeState              types.String  `tfsdk:"runtime_state"`
	ServiceManagerUuid        types.String  `tfsdk:"service_manager_uuid"`
	State                     types.String  `tfsdk:"state"`
	UsageBased                types.Bool    `tfsdk:"usage_based"`
	VisibleToProviders        types.Bool    `tfsdk:"visible_to_providers"`
	VisibleToUsername         types.String  `tfsdk:"visible_to_username"`
	AvailableActions          types.List    `tfsdk:"available_actions"`
	CanTerminate              types.Bool    `tfsdk:"can_terminate"`
	CategoryIcon              types.String  `tfsdk:"category_icon"`
	CategoryTitle             types.String  `tfsdk:"category_title"`
	CustomerName              types.String  `tfsdk:"customer_name"`
	CustomerSlug              types.String  `tfsdk:"customer_slug"`
	Description               types.String  `tfsdk:"description"`
	EffectiveId               types.String  `tfsdk:"effective_id"`
	EndDate                   types.String  `tfsdk:"end_date"`
	EndDateRequestedBy        types.String  `tfsdk:"end_date_requested_by"`
	Endpoints                 types.List    `tfsdk:"endpoints"`
	ErrorMessage              types.String  `tfsdk:"error_message"`
	ErrorTraceback            types.String  `tfsdk:"error_traceback"`
	LastSync                  types.String  `tfsdk:"last_sync"`
	OfferingComponents        types.List    `tfsdk:"offering_components"`
	OfferingDescription       types.String  `tfsdk:"offering_description"`
	OfferingImage             types.String  `tfsdk:"offering_image"`
	OfferingName              types.String  `tfsdk:"offering_name"`
	OfferingState             types.String  `tfsdk:"offering_state"`
	OfferingThumbnail         types.String  `tfsdk:"offering_thumbnail"`
	ParentName                types.String  `tfsdk:"parent_name"`
	ParentOfferingName        types.String  `tfsdk:"parent_offering_name"`
	ParentOfferingSlug        types.String  `tfsdk:"parent_offering_slug"`
	ParentUuid                types.String  `tfsdk:"parent_uuid"`
	Plan                      types.String  `tfsdk:"plan"`
	PlanDescription           types.String  `tfsdk:"plan_description"`
	PlanName                  types.String  `tfsdk:"plan_name"`
	PlanUnit                  types.String  `tfsdk:"plan_unit"`
	Project                   types.String  `tfsdk:"project"`
	ProjectDescription        types.String  `tfsdk:"project_description"`
	ProjectEndDate            types.String  `tfsdk:"project_end_date"`
	ProjectEndDateRequestedBy types.String  `tfsdk:"project_end_date_requested_by"`
	ProjectSlug               types.String  `tfsdk:"project_slug"`
	ProviderName              types.String  `tfsdk:"provider_name"`
	ProviderSlug              types.String  `tfsdk:"provider_slug"`
	Report                    types.List    `tfsdk:"report"`
	ResourceType              types.String  `tfsdk:"resource_type"`
	ResourceUuid              types.String  `tfsdk:"resource_uuid"`
	Scope                     types.String  `tfsdk:"scope"`
	Slug                      types.String  `tfsdk:"slug"`
	Url                       types.String  `tfsdk:"url"`
	UserRequiresReconsent     types.Bool    `tfsdk:"user_requires_reconsent"`
	Username                  types.String  `tfsdk:"username"`
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
				MarkdownDescription: "Resource UUID",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Backend ID",
			},
			"category_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Category UUID",
			},
			"component_count": schema.Float64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Filter by exact number of components",
			},
			"created": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Created after",
			},
			"customer": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer URL",
			},
			"customer_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer UUID",
			},
			"downscaled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Downscaled",
			},
			"has_terminate_date": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Has termination date",
			},
			"is_attached": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Filter by attached state",
			},
			"lexis_links_supported": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "LEXIS links supported",
			},
			"limit_based": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Filter by limit-based offerings",
			},
			"limit_component_count": schema.Float64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Filter by exact number of limit-based components",
			},
			"modified": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Modified after",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Name",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Name (exact)",
			},
			"offering": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Offering",
			},
			"offering_billable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Offering billable",
			},
			"offering_shared": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Offering shared",
			},
			"offering_slug": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Multiple values may be separated by commas.",
			},
			"offering_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Offering type",
			},
			"offering_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Multiple values may be separated by commas.",
			},
			"only_limit_based": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Filter resources with only limit-based components",
			},
			"only_usage_based": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Filter resources with only usage-based components",
			},
			"order_state": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Order state",
			},
			"parent_offering_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "UUID of the parent offering",
			},
			"paused": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Paused",
			},
			"plan_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Plan UUID",
			},
			"project_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project name",
			},
			"project_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project UUID",
			},
			"provider_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Provider UUID",
			},
			"query": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Search by resource UUID, name, slug, backend ID, effective ID, IPs or hypervisor",
			},
			"restrict_member_access": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Restrict member access",
			},
			"runtime_state": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Runtime state",
			},
			"service_manager_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Service manager UUID",
			},
			"state": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource state",
			},
			"usage_based": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Filter by usage-based offerings",
			},
			"visible_to_providers": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Include only resources visible to service providers",
			},
			"visible_to_username": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Visible to username",
			},
			"available_actions": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Computed:            true,
				MarkdownDescription: "Available actions",
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
			"customer_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the customer",
			},
			"customer_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Customer slug",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Description of the resource",
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
			"endpoints": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"name": types.StringType,
					"url":  types.StringType,
				}}},
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
				Computed:            true,
				MarkdownDescription: "Last sync",
			},
			"offering_components": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
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
				}}},
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
			"offering_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering state",
			},
			"offering_thumbnail": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering thumbnail",
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
			"parent_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the parent",
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
			"report": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"body":   types.StringType,
					"header": types.StringType,
				}}},
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
			"scope": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Scope",
			},
			"slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
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

func (d *MarketplaceResourceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data MarketplaceResourceDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var apiResp MarketplaceResourceApiResponse

		err := d.client.GetByUUID(ctx, "/api/marketplace-resources/{uuid}/", data.UUID.ValueString(), &apiResp)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Marketplace Resource",
				"An error occurred while reading the Marketplace Resource by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, apiResp, &data)...)

	} else {
		// Filter by provided parameters
		var results []MarketplaceResourceApiResponse

		filters := map[string]string{}
		if !data.BackendId.IsNull() {
			filters["backend_id"] = data.BackendId.ValueString()
		}
		if !data.CategoryUuid.IsNull() {
			filters["category_uuid"] = data.CategoryUuid.ValueString()
		}
		if !data.ComponentCount.IsNull() {
			filters["component_count"] = fmt.Sprintf("%f", data.ComponentCount.ValueFloat64())
		}
		if !data.Created.IsNull() {
			filters["created"] = data.Created.ValueString()
		}
		if !data.Customer.IsNull() {
			filters["customer"] = data.Customer.ValueString()
		}
		if !data.CustomerUuid.IsNull() {
			filters["customer_uuid"] = data.CustomerUuid.ValueString()
		}
		if !data.Downscaled.IsNull() {
			filters["downscaled"] = fmt.Sprintf("%t", data.Downscaled.ValueBool())
		}
		if !data.HasTerminateDate.IsNull() {
			filters["has_terminate_date"] = fmt.Sprintf("%t", data.HasTerminateDate.ValueBool())
		}
		if !data.IsAttached.IsNull() {
			filters["is_attached"] = fmt.Sprintf("%t", data.IsAttached.ValueBool())
		}
		if !data.LexisLinksSupported.IsNull() {
			filters["lexis_links_supported"] = fmt.Sprintf("%t", data.LexisLinksSupported.ValueBool())
		}
		if !data.LimitBased.IsNull() {
			filters["limit_based"] = fmt.Sprintf("%t", data.LimitBased.ValueBool())
		}
		if !data.LimitComponentCount.IsNull() {
			filters["limit_component_count"] = fmt.Sprintf("%f", data.LimitComponentCount.ValueFloat64())
		}
		if !data.Modified.IsNull() {
			filters["modified"] = data.Modified.ValueString()
		}
		if !data.Name.IsNull() {
			filters["name"] = data.Name.ValueString()
		}
		if !data.NameExact.IsNull() {
			filters["name_exact"] = data.NameExact.ValueString()
		}
		if !data.Offering.IsNull() {
			filters["offering"] = data.Offering.ValueString()
		}
		if !data.OfferingBillable.IsNull() {
			filters["offering_billable"] = fmt.Sprintf("%t", data.OfferingBillable.ValueBool())
		}
		if !data.OfferingShared.IsNull() {
			filters["offering_shared"] = fmt.Sprintf("%t", data.OfferingShared.ValueBool())
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
		if !data.OnlyLimitBased.IsNull() {
			filters["only_limit_based"] = fmt.Sprintf("%t", data.OnlyLimitBased.ValueBool())
		}
		if !data.OnlyUsageBased.IsNull() {
			filters["only_usage_based"] = fmt.Sprintf("%t", data.OnlyUsageBased.ValueBool())
		}
		if !data.OrderState.IsNull() {
			filters["order_state"] = data.OrderState.ValueString()
		}
		if !data.ParentOfferingUuid.IsNull() {
			filters["parent_offering_uuid"] = data.ParentOfferingUuid.ValueString()
		}
		if !data.Paused.IsNull() {
			filters["paused"] = fmt.Sprintf("%t", data.Paused.ValueBool())
		}
		if !data.PlanUuid.IsNull() {
			filters["plan_uuid"] = data.PlanUuid.ValueString()
		}
		if !data.ProjectName.IsNull() {
			filters["project_name"] = data.ProjectName.ValueString()
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
		if !data.RestrictMemberAccess.IsNull() {
			filters["restrict_member_access"] = fmt.Sprintf("%t", data.RestrictMemberAccess.ValueBool())
		}
		if !data.RuntimeState.IsNull() {
			filters["runtime_state"] = data.RuntimeState.ValueString()
		}
		if !data.ServiceManagerUuid.IsNull() {
			filters["service_manager_uuid"] = data.ServiceManagerUuid.ValueString()
		}
		if !data.State.IsNull() {
			filters["state"] = data.State.ValueString()
		}
		if !data.UsageBased.IsNull() {
			filters["usage_based"] = fmt.Sprintf("%t", data.UsageBased.ValueBool())
		}
		if !data.VisibleToProviders.IsNull() {
			filters["visible_to_providers"] = fmt.Sprintf("%t", data.VisibleToProviders.ValueBool())
		}
		if !data.VisibleToUsername.IsNull() {
			filters["visible_to_username"] = data.VisibleToUsername.ValueString()
		}

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup marketplace_resource.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/marketplace-resources/", filters, &results)
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

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *MarketplaceResourceDataSource) mapResponseToModel(ctx context.Context, apiResp MarketplaceResourceApiResponse, model *MarketplaceResourceDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AvailableActions, _ = types.ListValueFrom(ctx, types.StringType, apiResp.AvailableActions)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.CanTerminate = types.BoolPointerValue(apiResp.CanTerminate)
	model.CategoryIcon = types.StringPointerValue(apiResp.CategoryIcon)
	model.CategoryTitle = types.StringPointerValue(apiResp.CategoryTitle)
	model.CategoryUuid = types.StringPointerValue(apiResp.CategoryUuid)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.CustomerName = types.StringPointerValue(apiResp.CustomerName)
	model.CustomerSlug = types.StringPointerValue(apiResp.CustomerSlug)
	model.CustomerUuid = types.StringPointerValue(apiResp.CustomerUuid)
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
	model.Project = types.StringPointerValue(apiResp.Project)
	model.ProjectDescription = types.StringPointerValue(apiResp.ProjectDescription)
	model.ProjectEndDate = types.StringPointerValue(apiResp.ProjectEndDate)
	model.ProjectEndDateRequestedBy = types.StringPointerValue(apiResp.ProjectEndDateRequestedBy)
	model.ProjectName = types.StringPointerValue(apiResp.ProjectName)
	model.ProjectSlug = types.StringPointerValue(apiResp.ProjectSlug)
	model.ProjectUuid = types.StringPointerValue(apiResp.ProjectUuid)
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

	return diags
}
