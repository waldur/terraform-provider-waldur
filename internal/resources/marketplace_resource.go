package resources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MarketplaceResourceResource{}
var _ resource.ResourceWithImportState = &MarketplaceResourceResource{}

func NewMarketplaceResourceResource() resource.Resource {
	return &MarketplaceResourceResource{}
}

// MarketplaceResourceResource defines the resource implementation.
type MarketplaceResourceResource struct {
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
	Url *string `json:"url" tfsdk:"url"`
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
	OverageComponent   *string `json:"overage_component" tfsdk:"overage_component"`
	Type               *string `json:"type" tfsdk:"type"`
	UnitFactor         *int64  `json:"unit_factor" tfsdk:"unit_factor"`
}

type MarketplaceResourceReportResponse struct {
	Body   *string `json:"body" tfsdk:"body"`
	Header *string `json:"header" tfsdk:"header"`
}

var marketplaceresource_endpointsAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"url":  types.StringType,
}
var marketplaceresource_endpointsObjectType = types.ObjectType{
	AttrTypes: marketplaceresource_endpointsAttrTypes,
}

var marketplaceresource_offering_componentsAttrTypes = map[string]attr.Type{
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
}
var marketplaceresource_offering_componentsObjectType = types.ObjectType{
	AttrTypes: marketplaceresource_offering_componentsAttrTypes,
}

var marketplaceresource_reportAttrTypes = map[string]attr.Type{
	"body":   types.StringType,
	"header": types.StringType,
}
var marketplaceresource_reportObjectType = types.ObjectType{
	AttrTypes: marketplaceresource_reportAttrTypes,
}

// MarketplaceResourceResourceModel describes the resource data model.
type MarketplaceResourceResourceModel struct {
	UUID                      types.String   `tfsdk:"id"`
	AvailableActions          types.List     `tfsdk:"available_actions"`
	BackendId                 types.String   `tfsdk:"backend_id"`
	CanTerminate              types.Bool     `tfsdk:"can_terminate"`
	CategoryIcon              types.String   `tfsdk:"category_icon"`
	CategoryTitle             types.String   `tfsdk:"category_title"`
	CategoryUuid              types.String   `tfsdk:"category_uuid"`
	Created                   types.String   `tfsdk:"created"`
	CustomerSlug              types.String   `tfsdk:"customer_slug"`
	Description               types.String   `tfsdk:"description"`
	Downscaled                types.Bool     `tfsdk:"downscaled"`
	EffectiveId               types.String   `tfsdk:"effective_id"`
	EndDate                   types.String   `tfsdk:"end_date"`
	EndDateRequestedBy        types.String   `tfsdk:"end_date_requested_by"`
	Endpoints                 types.List     `tfsdk:"endpoints"`
	ErrorMessage              types.String   `tfsdk:"error_message"`
	ErrorTraceback            types.String   `tfsdk:"error_traceback"`
	LastSync                  types.String   `tfsdk:"last_sync"`
	Modified                  types.String   `tfsdk:"modified"`
	Name                      types.String   `tfsdk:"name"`
	Offering                  types.String   `tfsdk:"offering"`
	OfferingBillable          types.Bool     `tfsdk:"offering_billable"`
	OfferingComponents        types.List     `tfsdk:"offering_components"`
	OfferingDescription       types.String   `tfsdk:"offering_description"`
	OfferingImage             types.String   `tfsdk:"offering_image"`
	OfferingName              types.String   `tfsdk:"offering_name"`
	OfferingShared            types.Bool     `tfsdk:"offering_shared"`
	OfferingSlug              types.String   `tfsdk:"offering_slug"`
	OfferingState             types.String   `tfsdk:"offering_state"`
	OfferingThumbnail         types.String   `tfsdk:"offering_thumbnail"`
	OfferingType              types.String   `tfsdk:"offering_type"`
	OfferingUuid              types.String   `tfsdk:"offering_uuid"`
	ParentName                types.String   `tfsdk:"parent_name"`
	ParentOfferingName        types.String   `tfsdk:"parent_offering_name"`
	ParentOfferingSlug        types.String   `tfsdk:"parent_offering_slug"`
	ParentOfferingUuid        types.String   `tfsdk:"parent_offering_uuid"`
	ParentUuid                types.String   `tfsdk:"parent_uuid"`
	Paused                    types.Bool     `tfsdk:"paused"`
	Plan                      types.String   `tfsdk:"plan"`
	PlanDescription           types.String   `tfsdk:"plan_description"`
	PlanName                  types.String   `tfsdk:"plan_name"`
	PlanUnit                  types.String   `tfsdk:"plan_unit"`
	PlanUuid                  types.String   `tfsdk:"plan_uuid"`
	ProjectDescription        types.String   `tfsdk:"project_description"`
	ProjectEndDate            types.String   `tfsdk:"project_end_date"`
	ProjectEndDateRequestedBy types.String   `tfsdk:"project_end_date_requested_by"`
	ProjectSlug               types.String   `tfsdk:"project_slug"`
	ProviderName              types.String   `tfsdk:"provider_name"`
	ProviderSlug              types.String   `tfsdk:"provider_slug"`
	ProviderUuid              types.String   `tfsdk:"provider_uuid"`
	Report                    types.List     `tfsdk:"report"`
	ResourceType              types.String   `tfsdk:"resource_type"`
	ResourceUuid              types.String   `tfsdk:"resource_uuid"`
	RestrictMemberAccess      types.Bool     `tfsdk:"restrict_member_access"`
	Scope                     types.String   `tfsdk:"scope"`
	Slug                      types.String   `tfsdk:"slug"`
	State                     types.String   `tfsdk:"state"`
	Url                       types.String   `tfsdk:"url"`
	UserRequiresReconsent     types.Bool     `tfsdk:"user_requires_reconsent"`
	Username                  types.String   `tfsdk:"username"`
	Timeouts                  timeouts.Value `tfsdk:"timeouts"`
}

func (r *MarketplaceResourceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_resource"
}

func (r *MarketplaceResourceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Marketplace Resource resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"available_actions": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
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
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"customer_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"downscaled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"effective_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"end_date": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The date is inclusive. Once reached, a resource will be scheduled for termination.",
			},
			"end_date_requested_by": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"endpoints": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"url": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "URL of the access endpoint",
						},
					},
				},
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
			"last_sync": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"offering": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: " ",
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
							Computed:            true,
							MarkdownDescription: " ",
						},
						"billing_type": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"default_limit": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"factor": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"is_boolean": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"is_builtin": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"is_prepaid": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"limit_amount": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"limit_period": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"max_available_limit": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"max_prepaid_duration": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"max_value": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"measured_unit": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Unit of measurement, for example, GB.",
						},
						"min_prepaid_duration": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"min_value": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Display name for the measured unit, for example, Floating IP.",
						},
						"overage_component": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"type": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Unique internal name of the measured unit, for example floating_ip.",
						},
						"unit_factor": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "The conversion factor from backend units to measured_unit",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
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
			"offering_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"offering_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
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
			"parent_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"parent_offering_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"parent_offering_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"parent_offering_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"parent_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"paused": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"plan": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
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
			"project_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"project_end_date": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The date is inclusive. Once reached, all project resource will be scheduled for termination.",
			},
			"project_end_date_requested_by": schema.StringAttribute{
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
			"provider_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"provider_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"report": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"body": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Section body content",
						},
						"header": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Section header text",
						},
					},
				},
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
			"restrict_member_access": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"scope": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"slug": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"user_requires_reconsent": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Check if the current user needs to re-consent for this resource's offering.",
			},
			"username": schema.StringAttribute{
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

func (r *MarketplaceResourceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MarketplaceResourceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError("Creation Not Supported", "This resource cannot be created via the API.")
}

func (r *MarketplaceResourceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MarketplaceResourceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource

	retrievePath := strings.Replace("/api/marketplace-resources/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	var apiResp MarketplaceResourceApiResponse
	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Marketplace Resource",
			"An error occurred while reading the Marketplace Resource: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MarketplaceResourceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data MarketplaceResourceResourceModel
	var state MarketplaceResourceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Prepare request body
	requestBody := map[string]interface{}{}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		if v := data.Description.ValueString(); v != "" {
			requestBody["description"] = v
		}
	}
	if !data.EndDate.IsNull() && !data.EndDate.IsUnknown() {
		if v := data.EndDate.ValueString(); v != "" {
			requestBody["end_date"] = v
		}
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		if v := data.Name.ValueString(); v != "" {
			requestBody["name"] = v
		}
	}

	// Call Waldur API to update resource
	var apiResp MarketplaceResourceApiResponse

	err := r.client.Update(ctx, "/api/marketplace-resources/{uuid}/", data.UUID.ValueString(), requestBody, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Marketplace Resource",
			"An error occurred while updating the Marketplace Resource: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if apiResp.UUID != nil {
		data.UUID = types.StringPointerValue(apiResp.UUID)
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MarketplaceResourceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddError("Deletion Not Supported", "This resource cannot be deleted via the API.")
}

func (r *MarketplaceResourceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *MarketplaceResourceResource) mapResponseToModel(ctx context.Context, apiResp MarketplaceResourceApiResponse, model *MarketplaceResourceResourceModel) diag.Diagnostics {
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
	listValEndpoints, listDiagsEndpoints := types.ListValueFrom(ctx, marketplaceresource_endpointsObjectType, apiResp.Endpoints)
	diags.Append(listDiagsEndpoints...)
	model.Endpoints = listValEndpoints
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.LastSync = types.StringPointerValue(apiResp.LastSync)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	if apiResp.Offering != nil {
		parts := strings.Split(strings.TrimRight(*apiResp.Offering, "/"), "/")
		model.Offering = types.StringValue(parts[len(parts)-1])
	} else {
		model.Offering = types.StringNull()
	}
	model.OfferingBillable = types.BoolPointerValue(apiResp.OfferingBillable)
	listValOfferingComponents, listDiagsOfferingComponents := types.ListValueFrom(ctx, marketplaceresource_offering_componentsObjectType, apiResp.OfferingComponents)
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
	listValReport, listDiagsReport := types.ListValueFrom(ctx, marketplaceresource_reportObjectType, apiResp.Report)
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
