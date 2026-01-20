package datasources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
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
				MarkdownDescription: " ",
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
				MarkdownDescription: " ",
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
			"available_actions": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
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
			"customer_slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"effective_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"end_date": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The date is inclusive. Once reached, a resource will be scheduled for termination.",
			},
			"end_date_requested_by": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"endpoints": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"name": types.StringType,
					"url":  types.StringType,
				}}},
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
			"offering_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"offering_thumbnail": schema.StringAttribute{
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
			"parent_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"plan": schema.StringAttribute{
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
			"report": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"body":   types.StringType,
					"header": types.StringType,
				}}},
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
			"scope": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
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
		var item map[string]interface{}

		err := d.client.GetByUUID(ctx, "/api/marketplace-resources/{uuid}/", data.UUID.ValueString(), &item)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Marketplace Resource",
				"An error occurred while reading the Marketplace Resource by UUID: "+err.Error(),
			)
			return
		}

		// Extract data from single result
		if uuid, ok := item["uuid"].(string); ok {
			data.UUID = types.StringValue(uuid)
		}

		sourceMap := item
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
		if val, ok := sourceMap["offering_components"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
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
						attrValues := map[string]attr.Value{
							"article_code": func() attr.Value {
								if v, ok := objMap["article_code"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"billing_type": func() attr.Value {
								if v, ok := objMap["billing_type"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"default_limit": func() attr.Value {
								if v, ok := objMap["default_limit"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"description": func() attr.Value {
								if v, ok := objMap["description"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"factor": func() attr.Value {
								if v, ok := objMap["factor"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"is_boolean": func() attr.Value {
								if v, ok := objMap["is_boolean"].(bool); ok {
									return types.BoolValue(v)
								}
								return types.BoolNull()
							}(),
							"is_builtin": func() attr.Value {
								if v, ok := objMap["is_builtin"].(bool); ok {
									return types.BoolValue(v)
								}
								return types.BoolNull()
							}(),
							"is_prepaid": func() attr.Value {
								if v, ok := objMap["is_prepaid"].(bool); ok {
									return types.BoolValue(v)
								}
								return types.BoolNull()
							}(),
							"limit_amount": func() attr.Value {
								if v, ok := objMap["limit_amount"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"limit_period": func() attr.Value {
								if v, ok := objMap["limit_period"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"max_available_limit": func() attr.Value {
								if v, ok := objMap["max_available_limit"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"max_prepaid_duration": func() attr.Value {
								if v, ok := objMap["max_prepaid_duration"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"max_value": func() attr.Value {
								if v, ok := objMap["max_value"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"measured_unit": func() attr.Value {
								if v, ok := objMap["measured_unit"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"min_prepaid_duration": func() attr.Value {
								if v, ok := objMap["min_prepaid_duration"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"min_value": func() attr.Value {
								if v, ok := objMap["min_value"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"name": func() attr.Value {
								if v, ok := objMap["name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"overage_component": func() attr.Value {
								if v, ok := objMap["overage_component"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"type": func() attr.Value {
								if v, ok := objMap["type"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"unit_factor": func() attr.Value {
								if v, ok := objMap["unit_factor"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
						}
						objVal, _ := types.ObjectValue(attrTypes, attrValues)
						items = append(items, objVal)
					}
				}
				listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
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
				}}, items)
				data.OfferingComponents = listVal
			}
		} else {
			if data.OfferingComponents.IsUnknown() {
				data.OfferingComponents = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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
				}})
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
		if val, ok := sourceMap["offering_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingState = types.StringValue(str)
			}
		} else {
			if data.OfferingState.IsUnknown() {
				data.OfferingState = types.StringNull()
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
		if val, ok := sourceMap["parent_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ParentUuid = types.StringValue(str)
			}
		} else {
			if data.ParentUuid.IsUnknown() {
				data.ParentUuid = types.StringNull()
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
		if val, ok := sourceMap["backend_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BackendId = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["category_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CategoryUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["component_count"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.ComponentCount = types.Float64Value(num)
			}
		}
		if val, ok := sourceMap["created"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Created = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Customer = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["downscaled"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.Downscaled = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["has_terminate_date"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.HasTerminateDate = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["is_attached"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsAttached = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["lexis_links_supported"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.LexisLinksSupported = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["limit_based"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.LimitBased = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["limit_component_count"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.LimitComponentCount = types.Float64Value(num)
			}
		}
		if val, ok := sourceMap["modified"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Modified = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Name = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name_exact"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NameExact = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Offering = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering_billable"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.OfferingBillable = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["offering_shared"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.OfferingShared = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["offering_slug"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingSlug = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering_type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingType = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["only_limit_based"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.OnlyLimitBased = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["only_usage_based"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.OnlyUsageBased = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["order_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OrderState = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["parent_offering_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ParentOfferingUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["paused"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.Paused = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["plan_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.PlanUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["provider_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProviderUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["query"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Query = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["restrict_member_access"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.RestrictMemberAccess = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["runtime_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.RuntimeState = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_manager_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceManagerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.State = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["usage_based"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.UsageBased = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["visible_to_providers"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.VisibleToProviders = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["visible_to_username"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.VisibleToUsername = types.StringValue(str)
			}
		}

	} else {
		// Filter by provided parameters
		var results []map[string]interface{}

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

		// Extract data from the single result
		if uuid, ok := results[0]["uuid"].(string); ok {
			data.UUID = types.StringValue(uuid)
		}
		sourceMap := results[0]
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
		if val, ok := sourceMap["offering_components"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
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
						attrValues := map[string]attr.Value{
							"article_code": func() attr.Value {
								if v, ok := objMap["article_code"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"billing_type": func() attr.Value {
								if v, ok := objMap["billing_type"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"default_limit": func() attr.Value {
								if v, ok := objMap["default_limit"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"description": func() attr.Value {
								if v, ok := objMap["description"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"factor": func() attr.Value {
								if v, ok := objMap["factor"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"is_boolean": func() attr.Value {
								if v, ok := objMap["is_boolean"].(bool); ok {
									return types.BoolValue(v)
								}
								return types.BoolNull()
							}(),
							"is_builtin": func() attr.Value {
								if v, ok := objMap["is_builtin"].(bool); ok {
									return types.BoolValue(v)
								}
								return types.BoolNull()
							}(),
							"is_prepaid": func() attr.Value {
								if v, ok := objMap["is_prepaid"].(bool); ok {
									return types.BoolValue(v)
								}
								return types.BoolNull()
							}(),
							"limit_amount": func() attr.Value {
								if v, ok := objMap["limit_amount"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"limit_period": func() attr.Value {
								if v, ok := objMap["limit_period"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"max_available_limit": func() attr.Value {
								if v, ok := objMap["max_available_limit"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"max_prepaid_duration": func() attr.Value {
								if v, ok := objMap["max_prepaid_duration"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"max_value": func() attr.Value {
								if v, ok := objMap["max_value"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"measured_unit": func() attr.Value {
								if v, ok := objMap["measured_unit"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"min_prepaid_duration": func() attr.Value {
								if v, ok := objMap["min_prepaid_duration"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"min_value": func() attr.Value {
								if v, ok := objMap["min_value"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"name": func() attr.Value {
								if v, ok := objMap["name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"overage_component": func() attr.Value {
								if v, ok := objMap["overage_component"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"type": func() attr.Value {
								if v, ok := objMap["type"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"unit_factor": func() attr.Value {
								if v, ok := objMap["unit_factor"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
						}
						objVal, _ := types.ObjectValue(attrTypes, attrValues)
						items = append(items, objVal)
					}
				}
				listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
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
				}}, items)
				data.OfferingComponents = listVal
			}
		} else {
			if data.OfferingComponents.IsUnknown() {
				data.OfferingComponents = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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
				}})
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
		if val, ok := sourceMap["offering_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingState = types.StringValue(str)
			}
		} else {
			if data.OfferingState.IsUnknown() {
				data.OfferingState = types.StringNull()
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
		if val, ok := sourceMap["parent_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ParentUuid = types.StringValue(str)
			}
		} else {
			if data.ParentUuid.IsUnknown() {
				data.ParentUuid = types.StringNull()
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
		if val, ok := sourceMap["backend_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BackendId = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["category_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CategoryUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["component_count"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.ComponentCount = types.Float64Value(num)
			}
		}
		if val, ok := sourceMap["created"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Created = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Customer = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["downscaled"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.Downscaled = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["has_terminate_date"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.HasTerminateDate = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["is_attached"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsAttached = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["lexis_links_supported"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.LexisLinksSupported = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["limit_based"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.LimitBased = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["limit_component_count"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.LimitComponentCount = types.Float64Value(num)
			}
		}
		if val, ok := sourceMap["modified"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Modified = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Name = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name_exact"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NameExact = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Offering = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering_billable"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.OfferingBillable = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["offering_shared"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.OfferingShared = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["offering_slug"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingSlug = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering_type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingType = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["only_limit_based"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.OnlyLimitBased = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["only_usage_based"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.OnlyUsageBased = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["order_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OrderState = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["parent_offering_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ParentOfferingUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["paused"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.Paused = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["plan_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.PlanUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["provider_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProviderUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["query"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Query = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["restrict_member_access"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.RestrictMemberAccess = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["runtime_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.RuntimeState = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_manager_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceManagerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.State = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["usage_based"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.UsageBased = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["visible_to_providers"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.VisibleToProviders = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["visible_to_username"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.VisibleToUsername = types.StringValue(str)
			}
		}
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
