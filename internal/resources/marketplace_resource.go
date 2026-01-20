package resources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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
	OfferingDescription       types.String   `tfsdk:"offering_description"`
	OfferingImage             types.String   `tfsdk:"offering_image"`
	OfferingName              types.String   `tfsdk:"offering_name"`
	OfferingShared            types.Bool     `tfsdk:"offering_shared"`
	OfferingSlug              types.String   `tfsdk:"offering_slug"`
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
							MarkdownDescription: " ",
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
							MarkdownDescription: " ",
						},
						"header": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
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
				MarkdownDescription: " ",
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
	var result map[string]interface{}

	retrievePath := strings.Replace("/api/marketplace-resources/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read MarketplaceResource",
			"An error occurred while reading the marketplace_resource: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
	// Map response fields to data model
	_ = sourceMap
	if val, ok := sourceMap["available_actions"]; ok && val != nil {
		// List of strings (or flattened objects)
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if str, ok := item.(string); ok {
					items = append(items, types.StringValue(str))
				} else if obj, ok := item.(map[string]interface{}); ok {
					// Flattening logic: extract URL or UUID
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					} else if name, ok := obj["name"].(string); ok {
						items = append(items, types.StringValue(name))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.AvailableActions = listVal
		}
	} else {
		if data.AvailableActions.IsUnknown() {
			data.AvailableActions = types.ListNull(types.StringType)
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
	if val, ok := sourceMap["created"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
	} else {
		if data.Created.IsUnknown() {
			data.Created = types.StringNull()
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
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
		}
	}
	if val, ok := sourceMap["downscaled"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Downscaled = types.BoolValue(b)
		}
	} else {
		if data.Downscaled.IsUnknown() {
			data.Downscaled = types.BoolNull()
		}
	}
	if val, ok := sourceMap["effective_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.EffectiveId = types.StringValue(str)
		}
	} else {
		if data.EffectiveId.IsUnknown() {
			data.EffectiveId = types.StringNull()
		}
	}
	if val, ok := sourceMap["end_date"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.EndDate = types.StringValue(str)
		}
	} else {
		if data.EndDate.IsUnknown() {
			data.EndDate = types.StringNull()
		}
	}
	if val, ok := sourceMap["end_date_requested_by"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.EndDateRequestedBy = types.StringValue(str)
		}
	} else {
		if data.EndDateRequestedBy.IsUnknown() {
			data.EndDateRequestedBy = types.StringNull()
		}
	}
	if val, ok := sourceMap["endpoints"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"name": types.StringType,
						"url":  types.StringType,
					}
					attrValues := map[string]attr.Value{
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"name": types.StringType,
				"url":  types.StringType,
			}}, items)
			data.Endpoints = listVal
		}
	} else {
		if data.Endpoints.IsUnknown() {
			data.Endpoints = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"name": types.StringType,
				"url":  types.StringType,
			}})
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
	if val, ok := sourceMap["last_sync"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.LastSync = types.StringValue(str)
		}
	} else {
		if data.LastSync.IsUnknown() {
			data.LastSync = types.StringNull()
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
	if val, ok := sourceMap["offering_slug"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.OfferingSlug = types.StringValue(str)
		}
	} else {
		if data.OfferingSlug.IsUnknown() {
			data.OfferingSlug = types.StringNull()
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
	if val, ok := sourceMap["parent_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ParentName = types.StringValue(str)
		}
	} else {
		if data.ParentName.IsUnknown() {
			data.ParentName = types.StringNull()
		}
	}
	if val, ok := sourceMap["parent_offering_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ParentOfferingName = types.StringValue(str)
		}
	} else {
		if data.ParentOfferingName.IsUnknown() {
			data.ParentOfferingName = types.StringNull()
		}
	}
	if val, ok := sourceMap["parent_offering_slug"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ParentOfferingSlug = types.StringValue(str)
		}
	} else {
		if data.ParentOfferingSlug.IsUnknown() {
			data.ParentOfferingSlug = types.StringNull()
		}
	}
	if val, ok := sourceMap["parent_offering_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ParentOfferingUuid = types.StringValue(str)
		}
	} else {
		if data.ParentOfferingUuid.IsUnknown() {
			data.ParentOfferingUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["parent_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ParentUuid = types.StringValue(str)
		}
	} else {
		if data.ParentUuid.IsUnknown() {
			data.ParentUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["paused"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Paused = types.BoolValue(b)
		}
	} else {
		if data.Paused.IsUnknown() {
			data.Paused = types.BoolNull()
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
	if val, ok := sourceMap["project_description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProjectDescription = types.StringValue(str)
		}
	} else {
		if data.ProjectDescription.IsUnknown() {
			data.ProjectDescription = types.StringNull()
		}
	}
	if val, ok := sourceMap["project_end_date"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProjectEndDate = types.StringValue(str)
		}
	} else {
		if data.ProjectEndDate.IsUnknown() {
			data.ProjectEndDate = types.StringNull()
		}
	}
	if val, ok := sourceMap["project_end_date_requested_by"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProjectEndDateRequestedBy = types.StringValue(str)
		}
	} else {
		if data.ProjectEndDateRequestedBy.IsUnknown() {
			data.ProjectEndDateRequestedBy = types.StringNull()
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
	if val, ok := sourceMap["report"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"body":   types.StringType,
						"header": types.StringType,
					}
					attrValues := map[string]attr.Value{
						"body": func() attr.Value {
							if v, ok := objMap["body"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"header": func() attr.Value {
							if v, ok := objMap["header"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"body":   types.StringType,
				"header": types.StringType,
			}}, items)
			data.Report = listVal
		}
	} else {
		if data.Report.IsUnknown() {
			data.Report = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"body":   types.StringType,
				"header": types.StringType,
			}})
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
	if val, ok := sourceMap["restrict_member_access"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.RestrictMemberAccess = types.BoolValue(b)
		}
	} else {
		if data.RestrictMemberAccess.IsUnknown() {
			data.RestrictMemberAccess = types.BoolNull()
		}
	}
	if val, ok := sourceMap["scope"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Scope = types.StringValue(str)
		}
	} else {
		if data.Scope.IsUnknown() {
			data.Scope = types.StringNull()
		}
	}
	if val, ok := sourceMap["slug"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Slug = types.StringValue(str)
		}
	} else {
		if data.Slug.IsUnknown() {
			data.Slug = types.StringNull()
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
	if val, ok := sourceMap["url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Url = types.StringValue(str)
		}
	} else {
		if data.Url.IsUnknown() {
			data.Url = types.StringNull()
		}
	}
	if val, ok := sourceMap["user_requires_reconsent"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.UserRequiresReconsent = types.BoolValue(b)
		}
	} else {
		if data.UserRequiresReconsent.IsUnknown() {
			data.UserRequiresReconsent = types.BoolNull()
		}
	}
	if val, ok := sourceMap["username"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Username = types.StringValue(str)
		}
	} else {
		if data.Username.IsUnknown() {
			data.Username = types.StringNull()
		}
	}

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

	data.UUID = state.UUID

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
	var result map[string]interface{}

	err := r.client.Update(ctx, "/api/marketplace-resources/{uuid}/", data.UUID.ValueString(), requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update MarketplaceResource",
			"An error occurred while updating the marketplace_resource: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
	// Map response fields to data model
	_ = sourceMap
	if val, ok := sourceMap["available_actions"]; ok && val != nil {
		// List of strings (or flattened objects)
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if str, ok := item.(string); ok {
					items = append(items, types.StringValue(str))
				} else if obj, ok := item.(map[string]interface{}); ok {
					// Flattening logic: extract URL or UUID
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					} else if name, ok := obj["name"].(string); ok {
						items = append(items, types.StringValue(name))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.AvailableActions = listVal
		}
	} else {
		if data.AvailableActions.IsUnknown() {
			data.AvailableActions = types.ListNull(types.StringType)
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
	if val, ok := sourceMap["created"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
	} else {
		if data.Created.IsUnknown() {
			data.Created = types.StringNull()
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
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
		}
	}
	if val, ok := sourceMap["downscaled"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Downscaled = types.BoolValue(b)
		}
	} else {
		if data.Downscaled.IsUnknown() {
			data.Downscaled = types.BoolNull()
		}
	}
	if val, ok := sourceMap["effective_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.EffectiveId = types.StringValue(str)
		}
	} else {
		if data.EffectiveId.IsUnknown() {
			data.EffectiveId = types.StringNull()
		}
	}
	if val, ok := sourceMap["end_date"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.EndDate = types.StringValue(str)
		}
	} else {
		if data.EndDate.IsUnknown() {
			data.EndDate = types.StringNull()
		}
	}
	if val, ok := sourceMap["end_date_requested_by"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.EndDateRequestedBy = types.StringValue(str)
		}
	} else {
		if data.EndDateRequestedBy.IsUnknown() {
			data.EndDateRequestedBy = types.StringNull()
		}
	}
	if val, ok := sourceMap["endpoints"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"name": types.StringType,
						"url":  types.StringType,
					}
					attrValues := map[string]attr.Value{
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"name": types.StringType,
				"url":  types.StringType,
			}}, items)
			data.Endpoints = listVal
		}
	} else {
		if data.Endpoints.IsUnknown() {
			data.Endpoints = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"name": types.StringType,
				"url":  types.StringType,
			}})
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
	if val, ok := sourceMap["last_sync"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.LastSync = types.StringValue(str)
		}
	} else {
		if data.LastSync.IsUnknown() {
			data.LastSync = types.StringNull()
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
	if val, ok := sourceMap["offering_slug"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.OfferingSlug = types.StringValue(str)
		}
	} else {
		if data.OfferingSlug.IsUnknown() {
			data.OfferingSlug = types.StringNull()
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
	if val, ok := sourceMap["parent_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ParentName = types.StringValue(str)
		}
	} else {
		if data.ParentName.IsUnknown() {
			data.ParentName = types.StringNull()
		}
	}
	if val, ok := sourceMap["parent_offering_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ParentOfferingName = types.StringValue(str)
		}
	} else {
		if data.ParentOfferingName.IsUnknown() {
			data.ParentOfferingName = types.StringNull()
		}
	}
	if val, ok := sourceMap["parent_offering_slug"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ParentOfferingSlug = types.StringValue(str)
		}
	} else {
		if data.ParentOfferingSlug.IsUnknown() {
			data.ParentOfferingSlug = types.StringNull()
		}
	}
	if val, ok := sourceMap["parent_offering_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ParentOfferingUuid = types.StringValue(str)
		}
	} else {
		if data.ParentOfferingUuid.IsUnknown() {
			data.ParentOfferingUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["parent_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ParentUuid = types.StringValue(str)
		}
	} else {
		if data.ParentUuid.IsUnknown() {
			data.ParentUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["paused"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Paused = types.BoolValue(b)
		}
	} else {
		if data.Paused.IsUnknown() {
			data.Paused = types.BoolNull()
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
	if val, ok := sourceMap["project_description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProjectDescription = types.StringValue(str)
		}
	} else {
		if data.ProjectDescription.IsUnknown() {
			data.ProjectDescription = types.StringNull()
		}
	}
	if val, ok := sourceMap["project_end_date"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProjectEndDate = types.StringValue(str)
		}
	} else {
		if data.ProjectEndDate.IsUnknown() {
			data.ProjectEndDate = types.StringNull()
		}
	}
	if val, ok := sourceMap["project_end_date_requested_by"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProjectEndDateRequestedBy = types.StringValue(str)
		}
	} else {
		if data.ProjectEndDateRequestedBy.IsUnknown() {
			data.ProjectEndDateRequestedBy = types.StringNull()
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
	if val, ok := sourceMap["report"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"body":   types.StringType,
						"header": types.StringType,
					}
					attrValues := map[string]attr.Value{
						"body": func() attr.Value {
							if v, ok := objMap["body"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"header": func() attr.Value {
							if v, ok := objMap["header"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"body":   types.StringType,
				"header": types.StringType,
			}}, items)
			data.Report = listVal
		}
	} else {
		if data.Report.IsUnknown() {
			data.Report = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"body":   types.StringType,
				"header": types.StringType,
			}})
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
	if val, ok := sourceMap["restrict_member_access"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.RestrictMemberAccess = types.BoolValue(b)
		}
	} else {
		if data.RestrictMemberAccess.IsUnknown() {
			data.RestrictMemberAccess = types.BoolNull()
		}
	}
	if val, ok := sourceMap["scope"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Scope = types.StringValue(str)
		}
	} else {
		if data.Scope.IsUnknown() {
			data.Scope = types.StringNull()
		}
	}
	if val, ok := sourceMap["slug"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Slug = types.StringValue(str)
		}
	} else {
		if data.Slug.IsUnknown() {
			data.Slug = types.StringNull()
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
	if val, ok := sourceMap["url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Url = types.StringValue(str)
		}
	} else {
		if data.Url.IsUnknown() {
			data.Url = types.StringNull()
		}
	}
	if val, ok := sourceMap["user_requires_reconsent"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.UserRequiresReconsent = types.BoolValue(b)
		}
	} else {
		if data.UserRequiresReconsent.IsUnknown() {
			data.UserRequiresReconsent = types.BoolNull()
		}
	}
	if val, ok := sourceMap["username"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Username = types.StringValue(str)
		}
	} else {
		if data.Username.IsUnknown() {
			data.Username = types.StringNull()
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MarketplaceResourceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddError("Deletion Not Supported", "This resource cannot be deleted via the API.")
	return
}

func (r *MarketplaceResourceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
