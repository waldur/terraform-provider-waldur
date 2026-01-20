package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-provider-waldur/internal/client"
)

var _ list.ListResource = &MarketplaceResourceList{}

type MarketplaceResourceList struct {
	client *client.Client
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
				Description: "",
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
				Description: "",
				Optional:    true,
			},
			"customer_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"downscaled": schema.BoolAttribute{
				Description: "",
				Optional:    true,
			},
			"has_terminate_date": schema.BoolAttribute{
				Description: "Has termination date",
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
				Description: "",
				Optional:    true,
			},
			"name_exact": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"offering": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"offering_billable": schema.BoolAttribute{
				Description: "",
				Optional:    true,
			},
			"offering_shared": schema.BoolAttribute{
				Description: "Offering shared",
				Optional:    true,
			},
			"offering_type": schema.StringAttribute{
				Description: "",
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
				Description: "",
				Optional:    true,
			},
			"plan_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"project_name": schema.StringAttribute{
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
				Description: "Search by resource UUID, name, slug, backend ID, effective ID, IPs or hypervisor",
				Optional:    true,
			},
			"restrict_member_access": schema.BoolAttribute{
				Description: "",
				Optional:    true,
			},
			"runtime_state": schema.StringAttribute{
				Description: "Runtime state",
				Optional:    true,
			},
			"service_manager_uuid": schema.StringAttribute{
				Description: "Service Manager UUID",
				Optional:    true,
			},
			"usage_based": schema.BoolAttribute{
				Description: "Filter by usage-based offerings",
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

	l.client = client
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
	if !config.VisibleToUsername.IsNull() && !config.VisibleToUsername.IsUnknown() {
		filters["visible_to_username"] = config.VisibleToUsername.ValueString()
	}

	// Call API
	var listResult []map[string]interface{}
	err := l.client.ListWithFilter(ctx, "/api/marketplace-resources/", filters, &listResult)
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
			var data MarketplaceResourceResourceModel

			// Initialize lists and objects to empty/null if needed (or rely on map logic)

			sourceMap := item

			// Reuse the mapResponseFields logic (embedded here or via shared template)
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
