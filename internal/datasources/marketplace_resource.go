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
	CustomerSlug              *string                                         `json:"customer_slug" tfsdk:"customer_slug"`
	Description               *string                                         `json:"description" tfsdk:"description"`
	Downscaled                *bool                                           `json:"downscaled" tfsdk:"downscaled"`
	EffectiveId               *string                                         `json:"effective_id" tfsdk:"effective_id"`
	EndDate                   *string                                         `json:"end_date" tfsdk:"end_date"`
	EndDateRequestedBy        *string                                         `json:"end_date_requested_by" tfsdk:"end_date_requested_by"`
	Endpoints                 []MarketplaceResourceEndpointsResponse          `json:"endpoints" tfsdk:"endpoints"`
	ErrorMessage              *string                                         `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback            *string                                         `json:"error_traceback" tfsdk:"error_traceback"`
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
	ProjectDescription        *string                                         `json:"project_description" tfsdk:"project_description"`
	ProjectEndDate            *string                                         `json:"project_end_date" tfsdk:"project_end_date"`
	ProjectEndDateRequestedBy *string                                         `json:"project_end_date_requested_by" tfsdk:"project_end_date_requested_by"`
	ProjectSlug               *string                                         `json:"project_slug" tfsdk:"project_slug"`
	ProviderName              *string                                         `json:"provider_name" tfsdk:"provider_name"`
	ProviderSlug              *string                                         `json:"provider_slug" tfsdk:"provider_slug"`
	ProviderUuid              *string                                         `json:"provider_uuid" tfsdk:"provider_uuid"`
	Report                    []MarketplaceResourceReportResponse             `json:"report" tfsdk:"report"`
	ResourceType              *string                                         `json:"resource_type" tfsdk:"resource_type"`
	ResourceUuid              *string                                         `json:"resource_uuid" tfsdk:"resource_uuid"`
	RestrictMemberAccess      *bool                                           `json:"restrict_member_access" tfsdk:"restrict_member_access"`
	Scope                     *string                                         `json:"scope" tfsdk:"scope"`
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

// MarketplaceResourceFiltersModel contains the filter parameters for querying.
type MarketplaceResourceFiltersModel struct {
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
	OfferingSlug         types.String  `tfsdk:"offering_slug"`
	OfferingType         types.String  `tfsdk:"offering_type"`
	OfferingUuid         types.String  `tfsdk:"offering_uuid"`
	OnlyLimitBased       types.Bool    `tfsdk:"only_limit_based"`
	OnlyUsageBased       types.Bool    `tfsdk:"only_usage_based"`
	OrderState           types.String  `tfsdk:"order_state"`
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
	State                types.String  `tfsdk:"state"`
	UsageBased           types.Bool    `tfsdk:"usage_based"`
	VisibleToProviders   types.Bool    `tfsdk:"visible_to_providers"`
	VisibleToUsername    types.String  `tfsdk:"visible_to_username"`
}

// MarketplaceResourceDataSourceModel describes the data source data model.
type MarketplaceResourceDataSourceModel struct {
	UUID                      types.String                     `tfsdk:"id"`
	Filters                   *MarketplaceResourceFiltersModel `tfsdk:"filters"`
	AvailableActions          types.List                       `tfsdk:"available_actions"`
	BackendId                 types.String                     `tfsdk:"backend_id"`
	CanTerminate              types.Bool                       `tfsdk:"can_terminate"`
	CategoryIcon              types.String                     `tfsdk:"category_icon"`
	CategoryTitle             types.String                     `tfsdk:"category_title"`
	CategoryUuid              types.String                     `tfsdk:"category_uuid"`
	Created                   types.String                     `tfsdk:"created"`
	CustomerSlug              types.String                     `tfsdk:"customer_slug"`
	Description               types.String                     `tfsdk:"description"`
	Downscaled                types.Bool                       `tfsdk:"downscaled"`
	EffectiveId               types.String                     `tfsdk:"effective_id"`
	EndDate                   types.String                     `tfsdk:"end_date"`
	EndDateRequestedBy        types.String                     `tfsdk:"end_date_requested_by"`
	Endpoints                 types.List                       `tfsdk:"endpoints"`
	ErrorMessage              types.String                     `tfsdk:"error_message"`
	ErrorTraceback            types.String                     `tfsdk:"error_traceback"`
	LastSync                  types.String                     `tfsdk:"last_sync"`
	Modified                  types.String                     `tfsdk:"modified"`
	Name                      types.String                     `tfsdk:"name"`
	Offering                  types.String                     `tfsdk:"offering"`
	OfferingBillable          types.Bool                       `tfsdk:"offering_billable"`
	OfferingComponents        types.List                       `tfsdk:"offering_components"`
	OfferingDescription       types.String                     `tfsdk:"offering_description"`
	OfferingImage             types.String                     `tfsdk:"offering_image"`
	OfferingName              types.String                     `tfsdk:"offering_name"`
	OfferingShared            types.Bool                       `tfsdk:"offering_shared"`
	OfferingSlug              types.String                     `tfsdk:"offering_slug"`
	OfferingState             types.String                     `tfsdk:"offering_state"`
	OfferingThumbnail         types.String                     `tfsdk:"offering_thumbnail"`
	OfferingType              types.String                     `tfsdk:"offering_type"`
	OfferingUuid              types.String                     `tfsdk:"offering_uuid"`
	ParentName                types.String                     `tfsdk:"parent_name"`
	ParentOfferingName        types.String                     `tfsdk:"parent_offering_name"`
	ParentOfferingSlug        types.String                     `tfsdk:"parent_offering_slug"`
	ParentOfferingUuid        types.String                     `tfsdk:"parent_offering_uuid"`
	ParentUuid                types.String                     `tfsdk:"parent_uuid"`
	Paused                    types.Bool                       `tfsdk:"paused"`
	Plan                      types.String                     `tfsdk:"plan"`
	PlanDescription           types.String                     `tfsdk:"plan_description"`
	PlanName                  types.String                     `tfsdk:"plan_name"`
	PlanUnit                  types.String                     `tfsdk:"plan_unit"`
	PlanUuid                  types.String                     `tfsdk:"plan_uuid"`
	ProjectDescription        types.String                     `tfsdk:"project_description"`
	ProjectEndDate            types.String                     `tfsdk:"project_end_date"`
	ProjectEndDateRequestedBy types.String                     `tfsdk:"project_end_date_requested_by"`
	ProjectSlug               types.String                     `tfsdk:"project_slug"`
	ProviderName              types.String                     `tfsdk:"provider_name"`
	ProviderSlug              types.String                     `tfsdk:"provider_slug"`
	ProviderUuid              types.String                     `tfsdk:"provider_uuid"`
	Report                    types.List                       `tfsdk:"report"`
	ResourceType              types.String                     `tfsdk:"resource_type"`
	ResourceUuid              types.String                     `tfsdk:"resource_uuid"`
	RestrictMemberAccess      types.Bool                       `tfsdk:"restrict_member_access"`
	Scope                     types.String                     `tfsdk:"scope"`
	Slug                      types.String                     `tfsdk:"slug"`
	State                     types.String                     `tfsdk:"state"`
	Url                       types.String                     `tfsdk:"url"`
	UserRequiresReconsent     types.Bool                       `tfsdk:"user_requires_reconsent"`
	Username                  types.String                     `tfsdk:"username"`
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
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Filter parameters for querying Marketplace Resource",
				Attributes: map[string]schema.Attribute{
					"backend_id": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Backend ID",
					},
					"category_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Category UUID",
					},
					"component_count": schema.Float64Attribute{
						Optional:            true,
						MarkdownDescription: "Filter by exact number of components",
					},
					"created": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Created after",
					},
					"customer": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Customer URL",
					},
					"customer_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Customer UUID",
					},
					"downscaled": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Downscaled",
					},
					"has_terminate_date": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Has termination date",
					},
					"is_attached": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Filter by attached state",
					},
					"lexis_links_supported": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "LEXIS links supported",
					},
					"limit_based": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Filter by limit-based offerings",
					},
					"limit_component_count": schema.Float64Attribute{
						Optional:            true,
						MarkdownDescription: "Filter by exact number of limit-based components",
					},
					"modified": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Modified after",
					},
					"name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name",
					},
					"name_exact": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name (exact)",
					},
					"offering": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Offering",
					},
					"offering_billable": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Offering billable",
					},
					"offering_shared": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Offering shared",
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
						MarkdownDescription: "Multiple values may be separated by commas.",
					},
					"only_limit_based": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Filter resources with only limit-based components",
					},
					"only_usage_based": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Filter resources with only usage-based components",
					},
					"order_state": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Order state",
					},
					"parent_offering_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "UUID of the parent offering",
					},
					"paused": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Paused",
					},
					"plan_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Plan UUID",
					},
					"project_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Project name",
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
						MarkdownDescription: "Search by resource UUID, name, slug, backend ID, effective ID, IPs or hypervisor",
					},
					"restrict_member_access": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Restrict member access",
					},
					"runtime_state": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Runtime state",
					},
					"service_manager_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Service manager UUID",
					},
					"state": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Resource state",
					},
					"usage_based": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Filter by usage-based offerings",
					},
					"visible_to_providers": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Include only resources visible to service providers",
					},
					"visible_to_username": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Visible to username",
					},
				},
			},
			"available_actions": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "Available actions",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of the backend",
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
			"category_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the category",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Created",
			},
			"customer_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Customer slug",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Description of the resource",
			},
			"downscaled": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Downscaled",
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
			"endpoints": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the resource",
						},
						"url": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "URL of the access endpoint",
						},
					},
				},
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
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the resource",
			},
			"offering": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering",
			},
			"offering_billable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Purchase and usage is invoiced.",
			},
			"offering_components": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"article_code": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Article code",
						},
						"billing_type": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Billing type",
						},
						"default_limit": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Default limit",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Description of the resource",
						},
						"factor": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Factor",
						},
						"is_boolean": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "Is boolean",
						},
						"is_builtin": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Is builtin",
						},
						"is_prepaid": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "Is prepaid",
						},
						"limit_amount": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Limit amount",
						},
						"limit_period": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Limit period",
						},
						"max_available_limit": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Max available limit",
						},
						"max_prepaid_duration": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Max prepaid duration",
						},
						"max_value": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Max value",
						},
						"measured_unit": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Unit of measurement, for example, GB.",
						},
						"min_prepaid_duration": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Min prepaid duration",
						},
						"min_value": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Min value",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Display name for the measured unit, for example, Floating IP.",
						},
						"overage_component": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Overage component",
						},
						"type": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Unique internal name of the measured unit, for example floating_ip.",
						},
						"unit_factor": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "The conversion factor from backend units to measured_unit",
						},
					},
				},
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
			"offering_shared": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Accessible to all customers.",
			},
			"offering_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering slug",
			},
			"offering_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering state",
			},
			"offering_thumbnail": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering thumbnail",
			},
			"offering_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Offering type",
			},
			"offering_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the offering",
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
			"parent_offering_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the parent offering",
			},
			"parent_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the parent",
			},
			"paused": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Paused",
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
			"provider_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the provider",
			},
			"report": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"body": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Section body content",
						},
						"header": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Section header text",
						},
					},
				},
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
			"restrict_member_access": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Restrict member access",
			},
			"scope": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Scope",
			},
			"slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "State",
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

		filters := make(map[string]string)
		if data.Filters != nil {
			type filterDef struct {
				name string
				val  attr.Value
			}
			filterDefs := []filterDef{
				{"backend_id", data.Filters.BackendId},
				{"category_uuid", data.Filters.CategoryUuid},
				{"component_count", data.Filters.ComponentCount},
				{"created", data.Filters.Created},
				{"customer", data.Filters.Customer},
				{"customer_uuid", data.Filters.CustomerUuid},
				{"downscaled", data.Filters.Downscaled},
				{"has_terminate_date", data.Filters.HasTerminateDate},
				{"is_attached", data.Filters.IsAttached},
				{"lexis_links_supported", data.Filters.LexisLinksSupported},
				{"limit_based", data.Filters.LimitBased},
				{"limit_component_count", data.Filters.LimitComponentCount},
				{"modified", data.Filters.Modified},
				{"name", data.Filters.Name},
				{"name_exact", data.Filters.NameExact},
				{"offering", data.Filters.Offering},
				{"offering_billable", data.Filters.OfferingBillable},
				{"offering_shared", data.Filters.OfferingShared},
				{"offering_slug", data.Filters.OfferingSlug},
				{"offering_type", data.Filters.OfferingType},
				{"offering_uuid", data.Filters.OfferingUuid},
				{"only_limit_based", data.Filters.OnlyLimitBased},
				{"only_usage_based", data.Filters.OnlyUsageBased},
				{"order_state", data.Filters.OrderState},
				{"parent_offering_uuid", data.Filters.ParentOfferingUuid},
				{"paused", data.Filters.Paused},
				{"plan_uuid", data.Filters.PlanUuid},
				{"project_name", data.Filters.ProjectName},
				{"project_uuid", data.Filters.ProjectUuid},
				{"provider_uuid", data.Filters.ProviderUuid},
				{"query", data.Filters.Query},
				{"restrict_member_access", data.Filters.RestrictMemberAccess},
				{"runtime_state", data.Filters.RuntimeState},
				{"service_manager_uuid", data.Filters.ServiceManagerUuid},
				{"state", data.Filters.State},
				{"usage_based", data.Filters.UsageBased},
				{"visible_to_providers", data.Filters.VisibleToProviders},
				{"visible_to_username", data.Filters.VisibleToUsername},
			}

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

	return diags
}
