package resource

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MarketplaceResourceResource{}
var _ resource.ResourceWithImportState = &MarketplaceResourceResource{}

func NewMarketplaceResourceResource() resource.Resource {
	return &MarketplaceResourceResource{}
}

// MarketplaceResourceResource defines the resource implementation.
type MarketplaceResourceResource struct {
	client *Client
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
				CustomType: types.ListType{ElemType: types.StringType},
				Computed:   true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Available actions",
			},
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "ID of the backend",
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
			"created": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Created",
			},
			"customer_slug": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Customer slug",
			},
			"description": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Description of the resource",
			},
			"downscaled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Downscaled",
			},
			"effective_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "ID of the effective",
			},
			"end_date": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The date is inclusive. Once reached, a resource will be scheduled for termination.",
			},
			"end_date_requested_by": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Endpoints",
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
			"last_sync": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Last sync",
			},
			"modified": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the resource",
			},
			"offering": schema.StringAttribute{
				Optional: true,
				Computed: true,
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
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Offering components",
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
			"offering_slug": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Offering slug",
			},
			"offering_state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Offering state",
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
			"parent_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the parent",
			},
			"parent_offering_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the parent offering",
			},
			"parent_offering_slug": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Parent offering slug",
			},
			"parent_offering_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the parent offering",
			},
			"parent_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the parent",
			},
			"paused": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Paused",
			},
			"plan": schema.StringAttribute{
				Optional: true,
				Computed: true,
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
			"project_description": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Project description",
			},
			"project_end_date": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The date is inclusive. Once reached, all project resource will be scheduled for termination.",
			},
			"project_end_date_requested_by": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Project end date requested by",
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
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Report",
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
			"restrict_member_access": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Restrict member access",
			},
			"scope": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Scope",
			},
			"slug": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
			},
			"state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "State",
			},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Url",
			},
			"user_requires_reconsent": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Check if the current user needs to re-consent for this resource's offering.",
			},
			"username": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Username",
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

	r.client = NewClient(client)
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

	apiResp, err := r.client.GetMarketplaceResource(ctx, data.UUID.ValueString())
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Marketplace Resource",
			"An error occurred while reading the Marketplace Resource: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

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

	requestBody := MarketplaceResourceUpdateRequest{
		Description: data.Description.ValueStringPointer(),
		EndDate:     data.EndDate.ValueStringPointer(),
		Name:        data.Name.ValueStringPointer(),
	}

	apiResp, err := r.client.UpdateMarketplaceResource(ctx, data.UUID.ValueString(), &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Marketplace Resource",
			"An error occurred while updating the Marketplace Resource: "+err.Error(),
		)
		return
	}

	updateTimeout, diags := data.Timeouts.Update(ctx, 30*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*MarketplaceResourceResponse, error) {
		return r.client.GetMarketplaceResource(ctx, data.UUID.ValueString())
	}, updateTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource update", err.Error())
		return
	}
	apiResp = newResp

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MarketplaceResourceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddError("Deletion Not Supported", "This resource cannot be deleted via the API.")
}

func (r *MarketplaceResourceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *MarketplaceResourceResource) mapResponseToModel(ctx context.Context, apiResp MarketplaceResourceResponse, model *MarketplaceResourceResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	listValAvailableActions, listDiagsAvailableActions := types.ListValueFrom(ctx, types.StringType, apiResp.AvailableActions)
	model.AvailableActions = listValAvailableActions
	diags.Append(listDiagsAvailableActions...)
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

	{
		listValEndpoints, listDiagsEndpoints := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"url":  types.StringType,
		}}, apiResp.Endpoints)
		diags.Append(listDiagsEndpoints...)
		model.Endpoints = listValEndpoints
	}
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.LastSync = types.StringPointerValue(apiResp.LastSync)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.Offering = types.StringPointerValue(apiResp.Offering)
	model.OfferingBillable = types.BoolPointerValue(apiResp.OfferingBillable)

	{
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
	}
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

	{
		listValReport, listDiagsReport := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"body":   types.StringType,
			"header": types.StringType,
		}}, apiResp.Report)
		diags.Append(listDiagsReport...)
		model.Report = listValReport
	}
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
