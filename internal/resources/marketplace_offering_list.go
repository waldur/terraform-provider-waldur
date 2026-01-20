package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-provider-waldur/internal/client"
)

var _ list.ListResource = &MarketplaceOfferingList{}

type MarketplaceOfferingList struct {
	client *client.Client
}

func NewMarketplaceOfferingList() list.ListResource {
	return &MarketplaceOfferingList{}
}

func (l *MarketplaceOfferingList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_offering"
}

func (l *MarketplaceOfferingList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"accessible_via_calls": schema.BoolAttribute{
				Description: "Accessible via calls",
				Optional:    true,
			},
			"allowed_customer_uuid": schema.StringAttribute{
				Description: "Allowed customer UUID",
				Optional:    true,
			},
			"attributes": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"billable": schema.BoolAttribute{
				Description: "",
				Optional:    true,
			},
			"can_create_offering_user": schema.BoolAttribute{
				Description: "",
				Optional:    true,
			},
			"category_group_uuid": schema.StringAttribute{
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
			"customer": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"customer_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"has_active_terms_of_service": schema.BoolAttribute{
				Description: "Has Active Terms of Service",
				Optional:    true,
			},
			"has_terms_of_service": schema.BoolAttribute{
				Description: "Has Terms of Service",
				Optional:    true,
			},
			"keyword": schema.StringAttribute{
				Description: "Keyword",
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
			"page": schema.Int64Attribute{
				Description: "A page number within the paginated result set.",
				Optional:    true,
			},
			"page_size": schema.Int64Attribute{
				Description: "Number of results to return per page.",
				Optional:    true,
			},
			"parent_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"project_uuid": schema.StringAttribute{
				Description: "Project UUID",
				Optional:    true,
			},
			"query": schema.StringAttribute{
				Description: "Search by offering name, slug or description",
				Optional:    true,
			},
			"resource_customer_uuid": schema.StringAttribute{
				Description: "Resource customer UUID",
				Optional:    true,
			},
			"resource_project_uuid": schema.StringAttribute{
				Description: "Resource project UUID",
				Optional:    true,
			},
			"scope_uuid": schema.StringAttribute{
				Description: "Scope UUID",
				Optional:    true,
			},
			"service_manager_uuid": schema.StringAttribute{
				Description: "Service manager UUID",
				Optional:    true,
			},
			"shared": schema.BoolAttribute{
				Description: "",
				Optional:    true,
			},
			"user_has_consent": schema.BoolAttribute{
				Description: "User Has Consent",
				Optional:    true,
			},
			"user_has_offering_user": schema.BoolAttribute{
				Description: "User Has Offering User",
				Optional:    true,
			},
			"uuid_list": schema.StringAttribute{
				Description: "Comma-separated offering UUIDs",
				Optional:    true,
			},
		},
	}
}

func (l *MarketplaceOfferingList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type MarketplaceOfferingListModel struct {
	AccessibleViaCalls      types.Bool   `tfsdk:"accessible_via_calls"`
	AllowedCustomerUuid     types.String `tfsdk:"allowed_customer_uuid"`
	Attributes              types.String `tfsdk:"attributes"`
	Billable                types.Bool   `tfsdk:"billable"`
	CanCreateOfferingUser   types.Bool   `tfsdk:"can_create_offering_user"`
	CategoryGroupUuid       types.String `tfsdk:"category_group_uuid"`
	CategoryUuid            types.String `tfsdk:"category_uuid"`
	Created                 types.String `tfsdk:"created"`
	Customer                types.String `tfsdk:"customer"`
	CustomerUuid            types.String `tfsdk:"customer_uuid"`
	Description             types.String `tfsdk:"description"`
	HasActiveTermsOfService types.Bool   `tfsdk:"has_active_terms_of_service"`
	HasTermsOfService       types.Bool   `tfsdk:"has_terms_of_service"`
	Keyword                 types.String `tfsdk:"keyword"`
	Modified                types.String `tfsdk:"modified"`
	Name                    types.String `tfsdk:"name"`
	NameExact               types.String `tfsdk:"name_exact"`
	Page                    types.Int64  `tfsdk:"page"`
	PageSize                types.Int64  `tfsdk:"page_size"`
	ParentUuid              types.String `tfsdk:"parent_uuid"`
	ProjectUuid             types.String `tfsdk:"project_uuid"`
	Query                   types.String `tfsdk:"query"`
	ResourceCustomerUuid    types.String `tfsdk:"resource_customer_uuid"`
	ResourceProjectUuid     types.String `tfsdk:"resource_project_uuid"`
	ScopeUuid               types.String `tfsdk:"scope_uuid"`
	ServiceManagerUuid      types.String `tfsdk:"service_manager_uuid"`
	Shared                  types.Bool   `tfsdk:"shared"`
	UserHasConsent          types.Bool   `tfsdk:"user_has_consent"`
	UserHasOfferingUser     types.Bool   `tfsdk:"user_has_offering_user"`
	UuidList                types.String `tfsdk:"uuid_list"`
}

func (l *MarketplaceOfferingList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config MarketplaceOfferingListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := make(map[string]string)
	if !config.AccessibleViaCalls.IsNull() && !config.AccessibleViaCalls.IsUnknown() {
		filters["accessible_via_calls"] = fmt.Sprintf("%t", config.AccessibleViaCalls.ValueBool())
	}
	if !config.AllowedCustomerUuid.IsNull() && !config.AllowedCustomerUuid.IsUnknown() {
		filters["allowed_customer_uuid"] = config.AllowedCustomerUuid.ValueString()
	}
	if !config.Attributes.IsNull() && !config.Attributes.IsUnknown() {
		filters["attributes"] = config.Attributes.ValueString()
	}
	if !config.Billable.IsNull() && !config.Billable.IsUnknown() {
		filters["billable"] = fmt.Sprintf("%t", config.Billable.ValueBool())
	}
	if !config.CanCreateOfferingUser.IsNull() && !config.CanCreateOfferingUser.IsUnknown() {
		filters["can_create_offering_user"] = fmt.Sprintf("%t", config.CanCreateOfferingUser.ValueBool())
	}
	if !config.CategoryGroupUuid.IsNull() && !config.CategoryGroupUuid.IsUnknown() {
		filters["category_group_uuid"] = config.CategoryGroupUuid.ValueString()
	}
	if !config.CategoryUuid.IsNull() && !config.CategoryUuid.IsUnknown() {
		filters["category_uuid"] = config.CategoryUuid.ValueString()
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
	if !config.Description.IsNull() && !config.Description.IsUnknown() {
		filters["description"] = config.Description.ValueString()
	}
	if !config.HasActiveTermsOfService.IsNull() && !config.HasActiveTermsOfService.IsUnknown() {
		filters["has_active_terms_of_service"] = fmt.Sprintf("%t", config.HasActiveTermsOfService.ValueBool())
	}
	if !config.HasTermsOfService.IsNull() && !config.HasTermsOfService.IsUnknown() {
		filters["has_terms_of_service"] = fmt.Sprintf("%t", config.HasTermsOfService.ValueBool())
	}
	if !config.Keyword.IsNull() && !config.Keyword.IsUnknown() {
		filters["keyword"] = config.Keyword.ValueString()
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
	if !config.Page.IsNull() && !config.Page.IsUnknown() {
		filters["page"] = fmt.Sprintf("%d", config.Page.ValueInt64())
	}
	if !config.PageSize.IsNull() && !config.PageSize.IsUnknown() {
		filters["page_size"] = fmt.Sprintf("%d", config.PageSize.ValueInt64())
	}
	if !config.ParentUuid.IsNull() && !config.ParentUuid.IsUnknown() {
		filters["parent_uuid"] = config.ParentUuid.ValueString()
	}
	if !config.ProjectUuid.IsNull() && !config.ProjectUuid.IsUnknown() {
		filters["project_uuid"] = config.ProjectUuid.ValueString()
	}
	if !config.Query.IsNull() && !config.Query.IsUnknown() {
		filters["query"] = config.Query.ValueString()
	}
	if !config.ResourceCustomerUuid.IsNull() && !config.ResourceCustomerUuid.IsUnknown() {
		filters["resource_customer_uuid"] = config.ResourceCustomerUuid.ValueString()
	}
	if !config.ResourceProjectUuid.IsNull() && !config.ResourceProjectUuid.IsUnknown() {
		filters["resource_project_uuid"] = config.ResourceProjectUuid.ValueString()
	}
	if !config.ScopeUuid.IsNull() && !config.ScopeUuid.IsUnknown() {
		filters["scope_uuid"] = config.ScopeUuid.ValueString()
	}
	if !config.ServiceManagerUuid.IsNull() && !config.ServiceManagerUuid.IsUnknown() {
		filters["service_manager_uuid"] = config.ServiceManagerUuid.ValueString()
	}
	if !config.Shared.IsNull() && !config.Shared.IsUnknown() {
		filters["shared"] = fmt.Sprintf("%t", config.Shared.ValueBool())
	}
	if !config.UserHasConsent.IsNull() && !config.UserHasConsent.IsUnknown() {
		filters["user_has_consent"] = fmt.Sprintf("%t", config.UserHasConsent.ValueBool())
	}
	if !config.UserHasOfferingUser.IsNull() && !config.UserHasOfferingUser.IsUnknown() {
		filters["user_has_offering_user"] = fmt.Sprintf("%t", config.UserHasOfferingUser.ValueBool())
	}
	if !config.UuidList.IsNull() && !config.UuidList.IsUnknown() {
		filters["uuid_list"] = config.UuidList.ValueString()
	}

	// Call API
	var listResult []map[string]interface{}
	err := l.client.ListWithFilter(ctx, "/api/marketplace-provider-offerings/", filters, &listResult)
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
			var data MarketplaceOfferingResourceModel

			// Initialize lists and objects to empty/null if needed (or rely on map logic)

			sourceMap := item

			// Reuse the mapResponseFields logic (embedded here or via shared template)
			// Map response fields to data model
			_ = sourceMap
			if val, ok := sourceMap["access_url"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.AccessUrl = types.StringValue(str)
				}
			} else {
				if data.AccessUrl.IsUnknown() {
					data.AccessUrl = types.StringNull()
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
			if val, ok := sourceMap["billable"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.Billable = types.BoolValue(b)
				}
			} else {
				if data.Billable.IsUnknown() {
					data.Billable = types.BoolNull()
				}
			}
			if val, ok := sourceMap["category"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Category = types.StringValue(str)
				}
			} else {
				if data.Category.IsUnknown() {
					data.Category = types.StringNull()
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
			if val, ok := sourceMap["citation_count"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.CitationCount = types.Int64Value(int64(num))
				}
			} else {
				if data.CitationCount.IsUnknown() {
					data.CitationCount = types.Int64Null()
				}
			}
			if val, ok := sourceMap["compliance_checklist"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ComplianceChecklist = types.StringValue(str)
				}
			} else {
				if data.ComplianceChecklist.IsUnknown() {
					data.ComplianceChecklist = types.StringNull()
				}
			}
			if val, ok := sourceMap["components"]; ok && val != nil {
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
								"is_boolean":           types.BoolType,
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
								"is_boolean": func() attr.Value {
									if v, ok := objMap["is_boolean"].(bool); ok {
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
						"is_boolean":           types.BoolType,
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
					data.Components = listVal
				}
			} else {
				if data.Components.IsUnknown() {
					data.Components = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"article_code":         types.StringType,
						"billing_type":         types.StringType,
						"default_limit":        types.Int64Type,
						"description":          types.StringType,
						"is_boolean":           types.BoolType,
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
			if val, ok := sourceMap["country"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Country = types.StringValue(str)
				}
			} else {
				if data.Country.IsUnknown() {
					data.Country = types.StringNull()
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
			if val, ok := sourceMap["customer"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Customer = types.StringValue(str)
				}
			} else {
				if data.Customer.IsUnknown() {
					data.Customer = types.StringNull()
				}
			}
			if val, ok := sourceMap["datacite_doi"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.DataciteDoi = types.StringValue(str)
				}
			} else {
				if data.DataciteDoi.IsUnknown() {
					data.DataciteDoi = types.StringNull()
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
			if val, ok := sourceMap["files"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"created": types.StringType,
								"file":    types.StringType,
								"name":    types.StringType,
							}
							attrValues := map[string]attr.Value{
								"created": func() attr.Value {
									if v, ok := objMap["created"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"file": func() attr.Value {
									if v, ok := objMap["file"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"name": func() attr.Value {
									if v, ok := objMap["name"].(string); ok {
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
						"created": types.StringType,
						"file":    types.StringType,
						"name":    types.StringType,
					}}, items)
					data.Files = listVal
				}
			} else {
				if data.Files.IsUnknown() {
					data.Files = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"created": types.StringType,
						"file":    types.StringType,
						"name":    types.StringType,
					}})
				}
			}
			if val, ok := sourceMap["full_description"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.FullDescription = types.StringValue(str)
				}
			} else {
				if data.FullDescription.IsUnknown() {
					data.FullDescription = types.StringNull()
				}
			}
			if val, ok := sourceMap["getting_started"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.GettingStarted = types.StringValue(str)
				}
			} else {
				if data.GettingStarted.IsUnknown() {
					data.GettingStarted = types.StringNull()
				}
			}
			if val, ok := sourceMap["google_calendar_is_public"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.GoogleCalendarIsPublic = types.BoolValue(b)
				}
			} else {
				if data.GoogleCalendarIsPublic.IsUnknown() {
					data.GoogleCalendarIsPublic = types.BoolNull()
				}
			}
			if val, ok := sourceMap["google_calendar_link"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.GoogleCalendarLink = types.StringValue(str)
				}
			} else {
				if data.GoogleCalendarLink.IsUnknown() {
					data.GoogleCalendarLink = types.StringNull()
				}
			}
			if val, ok := sourceMap["has_compliance_requirements"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.HasComplianceRequirements = types.BoolValue(b)
				}
			} else {
				if data.HasComplianceRequirements.IsUnknown() {
					data.HasComplianceRequirements = types.BoolNull()
				}
			}
			if val, ok := sourceMap["image"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Image = types.StringValue(str)
				}
			} else {
				if data.Image.IsUnknown() {
					data.Image = types.StringNull()
				}
			}
			if val, ok := sourceMap["integration_guide"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.IntegrationGuide = types.StringValue(str)
				}
			} else {
				if data.IntegrationGuide.IsUnknown() {
					data.IntegrationGuide = types.StringNull()
				}
			}
			if val, ok := sourceMap["integration_status"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"agent_type":             types.StringType,
								"last_request_timestamp": types.StringType,
								"service_name":           types.StringType,
								"status":                 types.StringType,
							}
							attrValues := map[string]attr.Value{
								"agent_type": func() attr.Value {
									if v, ok := objMap["agent_type"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"last_request_timestamp": func() attr.Value {
									if v, ok := objMap["last_request_timestamp"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"service_name": func() attr.Value {
									if v, ok := objMap["service_name"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"status": func() attr.Value {
									if v, ok := objMap["status"].(string); ok {
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
						"agent_type":             types.StringType,
						"last_request_timestamp": types.StringType,
						"service_name":           types.StringType,
						"status":                 types.StringType,
					}}, items)
					data.IntegrationStatus = listVal
				}
			} else {
				if data.IntegrationStatus.IsUnknown() {
					data.IntegrationStatus = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"agent_type":             types.StringType,
						"last_request_timestamp": types.StringType,
						"service_name":           types.StringType,
						"status":                 types.StringType,
					}})
				}
			}
			if val, ok := sourceMap["latitude"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.Latitude = types.Float64Value(num)
				}
			} else {
				if data.Latitude.IsUnknown() {
					data.Latitude = types.Float64Null()
				}
			}
			if val, ok := sourceMap["longitude"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.Longitude = types.Float64Value(num)
				}
			} else {
				if data.Longitude.IsUnknown() {
					data.Longitude = types.Float64Null()
				}
			}
			if val, ok := sourceMap["options"]; ok && val != nil {
				// Nested object
				if objMap, ok := val.(map[string]interface{}); ok {
					_ = objMap // Avoid unused variable if properties are empty
					attrTypes := map[string]attr.Type{
						"order": types.ListType{ElemType: types.StringType},
					}
					attrValues := map[string]attr.Value{
						"order": types.ListNull(types.ListType{ElemType: types.StringType}.ElemType),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					data.Options = objVal
				}
			} else {
				if data.Options.IsUnknown() {
					data.Options = types.ObjectNull(map[string]attr.Type{
						"order": types.ListType{ElemType: types.StringType},
					})
				}
			}
			if val, ok := sourceMap["order_count"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.OrderCount = types.Int64Value(int64(num))
				}
			} else {
				if data.OrderCount.IsUnknown() {
					data.OrderCount = types.Int64Null()
				}
			}
			if val, ok := sourceMap["organization_groups"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"customers_count": types.Int64Type,
								"name":            types.StringType,
								"parent":          types.StringType,
								"parent_name":     types.StringType,
								"parent_uuid":     types.StringType,
								"url":             types.StringType,
							}
							attrValues := map[string]attr.Value{
								"customers_count": func() attr.Value {
									if v, ok := objMap["customers_count"].(float64); ok {
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
								"parent": func() attr.Value {
									if v, ok := objMap["parent"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"parent_name": func() attr.Value {
									if v, ok := objMap["parent_name"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"parent_uuid": func() attr.Value {
									if v, ok := objMap["parent_uuid"].(string); ok {
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
						"customers_count": types.Int64Type,
						"name":            types.StringType,
						"parent":          types.StringType,
						"parent_name":     types.StringType,
						"parent_uuid":     types.StringType,
						"url":             types.StringType,
					}}, items)
					data.OrganizationGroups = listVal
				}
			} else {
				if data.OrganizationGroups.IsUnknown() {
					data.OrganizationGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"customers_count": types.Int64Type,
						"name":            types.StringType,
						"parent":          types.StringType,
						"parent_name":     types.StringType,
						"parent_uuid":     types.StringType,
						"url":             types.StringType,
					}})
				}
			}
			if val, ok := sourceMap["parent_description"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ParentDescription = types.StringValue(str)
				}
			} else {
				if data.ParentDescription.IsUnknown() {
					data.ParentDescription = types.StringNull()
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
			if val, ok := sourceMap["parent_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ParentUuid = types.StringValue(str)
				}
			} else {
				if data.ParentUuid.IsUnknown() {
					data.ParentUuid = types.StringNull()
				}
			}
			if val, ok := sourceMap["partitions"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"cpu_bind":            types.Int64Type,
								"def_cpu_per_gpu":     types.Int64Type,
								"def_mem_per_cpu":     types.Int64Type,
								"def_mem_per_gpu":     types.Int64Type,
								"def_mem_per_node":    types.Int64Type,
								"default_time":        types.Int64Type,
								"exclusive_topo":      types.BoolType,
								"exclusive_user":      types.BoolType,
								"grace_time":          types.Int64Type,
								"max_cpus_per_node":   types.Int64Type,
								"max_cpus_per_socket": types.Int64Type,
								"max_mem_per_cpu":     types.Int64Type,
								"max_mem_per_node":    types.Int64Type,
								"max_nodes":           types.Int64Type,
								"max_time":            types.Int64Type,
								"min_nodes":           types.Int64Type,
								"partition_name":      types.StringType,
								"priority_tier":       types.Int64Type,
								"qos":                 types.StringType,
								"req_resv":            types.BoolType,
							}
							attrValues := map[string]attr.Value{
								"cpu_bind": func() attr.Value {
									if v, ok := objMap["cpu_bind"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"def_cpu_per_gpu": func() attr.Value {
									if v, ok := objMap["def_cpu_per_gpu"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"def_mem_per_cpu": func() attr.Value {
									if v, ok := objMap["def_mem_per_cpu"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"def_mem_per_gpu": func() attr.Value {
									if v, ok := objMap["def_mem_per_gpu"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"def_mem_per_node": func() attr.Value {
									if v, ok := objMap["def_mem_per_node"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"default_time": func() attr.Value {
									if v, ok := objMap["default_time"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"exclusive_topo": func() attr.Value {
									if v, ok := objMap["exclusive_topo"].(bool); ok {
										return types.BoolValue(v)
									}
									return types.BoolNull()
								}(),
								"exclusive_user": func() attr.Value {
									if v, ok := objMap["exclusive_user"].(bool); ok {
										return types.BoolValue(v)
									}
									return types.BoolNull()
								}(),
								"grace_time": func() attr.Value {
									if v, ok := objMap["grace_time"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"max_cpus_per_node": func() attr.Value {
									if v, ok := objMap["max_cpus_per_node"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"max_cpus_per_socket": func() attr.Value {
									if v, ok := objMap["max_cpus_per_socket"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"max_mem_per_cpu": func() attr.Value {
									if v, ok := objMap["max_mem_per_cpu"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"max_mem_per_node": func() attr.Value {
									if v, ok := objMap["max_mem_per_node"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"max_nodes": func() attr.Value {
									if v, ok := objMap["max_nodes"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"max_time": func() attr.Value {
									if v, ok := objMap["max_time"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"min_nodes": func() attr.Value {
									if v, ok := objMap["min_nodes"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"partition_name": func() attr.Value {
									if v, ok := objMap["partition_name"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"priority_tier": func() attr.Value {
									if v, ok := objMap["priority_tier"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"qos": func() attr.Value {
									if v, ok := objMap["qos"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"req_resv": func() attr.Value {
									if v, ok := objMap["req_resv"].(bool); ok {
										return types.BoolValue(v)
									}
									return types.BoolNull()
								}(),
							}
							objVal, _ := types.ObjectValue(attrTypes, attrValues)
							items = append(items, objVal)
						}
					}
					listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
						"cpu_bind":            types.Int64Type,
						"def_cpu_per_gpu":     types.Int64Type,
						"def_mem_per_cpu":     types.Int64Type,
						"def_mem_per_gpu":     types.Int64Type,
						"def_mem_per_node":    types.Int64Type,
						"default_time":        types.Int64Type,
						"exclusive_topo":      types.BoolType,
						"exclusive_user":      types.BoolType,
						"grace_time":          types.Int64Type,
						"max_cpus_per_node":   types.Int64Type,
						"max_cpus_per_socket": types.Int64Type,
						"max_mem_per_cpu":     types.Int64Type,
						"max_mem_per_node":    types.Int64Type,
						"max_nodes":           types.Int64Type,
						"max_time":            types.Int64Type,
						"min_nodes":           types.Int64Type,
						"partition_name":      types.StringType,
						"priority_tier":       types.Int64Type,
						"qos":                 types.StringType,
						"req_resv":            types.BoolType,
					}}, items)
					data.Partitions = listVal
				}
			} else {
				if data.Partitions.IsUnknown() {
					data.Partitions = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"cpu_bind":            types.Int64Type,
						"def_cpu_per_gpu":     types.Int64Type,
						"def_mem_per_cpu":     types.Int64Type,
						"def_mem_per_gpu":     types.Int64Type,
						"def_mem_per_node":    types.Int64Type,
						"default_time":        types.Int64Type,
						"exclusive_topo":      types.BoolType,
						"exclusive_user":      types.BoolType,
						"grace_time":          types.Int64Type,
						"max_cpus_per_node":   types.Int64Type,
						"max_cpus_per_socket": types.Int64Type,
						"max_mem_per_cpu":     types.Int64Type,
						"max_mem_per_node":    types.Int64Type,
						"max_nodes":           types.Int64Type,
						"max_time":            types.Int64Type,
						"min_nodes":           types.Int64Type,
						"partition_name":      types.StringType,
						"priority_tier":       types.Int64Type,
						"qos":                 types.StringType,
						"req_resv":            types.BoolType,
					}})
				}
			}
			if val, ok := sourceMap["paused_reason"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.PausedReason = types.StringValue(str)
				}
			} else {
				if data.PausedReason.IsUnknown() {
					data.PausedReason = types.StringNull()
				}
			}
			if val, ok := sourceMap["plans"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"archived":     types.BoolType,
								"article_code": types.StringType,
								"backend_id":   types.StringType,
								"description":  types.StringType,
								"max_amount":   types.Int64Type,
								"name":         types.StringType,
								"unit":         types.StringType,
								"unit_price":   types.StringType,
							}
							attrValues := map[string]attr.Value{
								"archived": func() attr.Value {
									if v, ok := objMap["archived"].(bool); ok {
										return types.BoolValue(v)
									}
									return types.BoolNull()
								}(),
								"article_code": func() attr.Value {
									if v, ok := objMap["article_code"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"backend_id": func() attr.Value {
									if v, ok := objMap["backend_id"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"description": func() attr.Value {
									if v, ok := objMap["description"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"max_amount": func() attr.Value {
									if v, ok := objMap["max_amount"].(float64); ok {
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
								"unit": func() attr.Value {
									if v, ok := objMap["unit"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"unit_price": func() attr.Value {
									if v, ok := objMap["unit_price"].(string); ok {
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
						"archived":     types.BoolType,
						"article_code": types.StringType,
						"backend_id":   types.StringType,
						"description":  types.StringType,
						"max_amount":   types.Int64Type,
						"name":         types.StringType,
						"unit":         types.StringType,
						"unit_price":   types.StringType,
					}}, items)
					data.Plans = listVal
				}
			} else {
				if data.Plans.IsUnknown() {
					data.Plans = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"archived":     types.BoolType,
						"article_code": types.StringType,
						"backend_id":   types.StringType,
						"description":  types.StringType,
						"max_amount":   types.Int64Type,
						"name":         types.StringType,
						"unit":         types.StringType,
						"unit_price":   types.StringType,
					}})
				}
			}
			if val, ok := sourceMap["privacy_policy_link"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.PrivacyPolicyLink = types.StringValue(str)
				}
			} else {
				if data.PrivacyPolicyLink.IsUnknown() {
					data.PrivacyPolicyLink = types.StringNull()
				}
			}
			if val, ok := sourceMap["quotas"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"limit": types.Int64Type,
								"name":  types.StringType,
								"usage": types.Int64Type,
							}
							attrValues := map[string]attr.Value{
								"limit": func() attr.Value {
									if v, ok := objMap["limit"].(float64); ok {
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
								"usage": func() attr.Value {
									if v, ok := objMap["usage"].(float64); ok {
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
						"limit": types.Int64Type,
						"name":  types.StringType,
						"usage": types.Int64Type,
					}}, items)
					data.Quotas = listVal
				}
			} else {
				if data.Quotas.IsUnknown() {
					data.Quotas = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"limit": types.Int64Type,
						"name":  types.StringType,
						"usage": types.Int64Type,
					}})
				}
			}
			if val, ok := sourceMap["resource_options"]; ok && val != nil {
				// Nested object
				if objMap, ok := val.(map[string]interface{}); ok {
					_ = objMap // Avoid unused variable if properties are empty
					attrTypes := map[string]attr.Type{
						"order": types.ListType{ElemType: types.StringType},
					}
					attrValues := map[string]attr.Value{
						"order": types.ListNull(types.ListType{ElemType: types.StringType}.ElemType),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					data.ResourceOptions = objVal
				}
			} else {
				if data.ResourceOptions.IsUnknown() {
					data.ResourceOptions = types.ObjectNull(map[string]attr.Type{
						"order": types.ListType{ElemType: types.StringType},
					})
				}
			}
			if val, ok := sourceMap["roles"]; ok && val != nil {
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
					data.Roles = listVal
				}
			} else {
				if data.Roles.IsUnknown() {
					data.Roles = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"name": types.StringType,
						"url":  types.StringType,
					}})
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
			if val, ok := sourceMap["scope_error_message"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ScopeErrorMessage = types.StringValue(str)
				}
			} else {
				if data.ScopeErrorMessage.IsUnknown() {
					data.ScopeErrorMessage = types.StringNull()
				}
			}
			if val, ok := sourceMap["scope_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ScopeName = types.StringValue(str)
				}
			} else {
				if data.ScopeName.IsUnknown() {
					data.ScopeName = types.StringNull()
				}
			}
			if val, ok := sourceMap["scope_state"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ScopeState = types.StringValue(str)
				}
			} else {
				if data.ScopeState.IsUnknown() {
					data.ScopeState = types.StringNull()
				}
			}
			if val, ok := sourceMap["scope_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ScopeUuid = types.StringValue(str)
				}
			} else {
				if data.ScopeUuid.IsUnknown() {
					data.ScopeUuid = types.StringNull()
				}
			}
			if val, ok := sourceMap["screenshots"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"created":     types.StringType,
								"description": types.StringType,
								"image":       types.StringType,
								"name":        types.StringType,
								"thumbnail":   types.StringType,
							}
							attrValues := map[string]attr.Value{
								"created": func() attr.Value {
									if v, ok := objMap["created"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"description": func() attr.Value {
									if v, ok := objMap["description"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"image": func() attr.Value {
									if v, ok := objMap["image"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"name": func() attr.Value {
									if v, ok := objMap["name"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"thumbnail": func() attr.Value {
									if v, ok := objMap["thumbnail"].(string); ok {
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
						"created":     types.StringType,
						"description": types.StringType,
						"image":       types.StringType,
						"name":        types.StringType,
						"thumbnail":   types.StringType,
					}}, items)
					data.Screenshots = listVal
				}
			} else {
				if data.Screenshots.IsUnknown() {
					data.Screenshots = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"created":     types.StringType,
						"description": types.StringType,
						"image":       types.StringType,
						"name":        types.StringType,
						"thumbnail":   types.StringType,
					}})
				}
			}
			if val, ok := sourceMap["shared"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.Shared = types.BoolValue(b)
				}
			} else {
				if data.Shared.IsUnknown() {
					data.Shared = types.BoolNull()
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
			if val, ok := sourceMap["software_catalogs"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"catalog": types.ObjectType{AttrTypes: map[string]attr.Type{
									"description": types.StringType,
									"name":        types.StringType,
									"version":     types.StringType,
								}},
								"package_count": types.Int64Type,
								"partition": types.ObjectType{AttrTypes: map[string]attr.Type{
									"partition_name": types.StringType,
									"priority_tier":  types.Int64Type,
									"qos":            types.StringType,
								}},
							}
							attrValues := map[string]attr.Value{
								"catalog": types.ObjectNull(types.ObjectType{AttrTypes: map[string]attr.Type{
									"description": types.StringType,
									"name":        types.StringType,
									"version":     types.StringType,
								}}.AttrTypes),
								"package_count": func() attr.Value {
									if v, ok := objMap["package_count"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"partition": types.ObjectNull(types.ObjectType{AttrTypes: map[string]attr.Type{
									"partition_name": types.StringType,
									"priority_tier":  types.Int64Type,
									"qos":            types.StringType,
								}}.AttrTypes),
							}
							objVal, _ := types.ObjectValue(attrTypes, attrValues)
							items = append(items, objVal)
						}
					}
					listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
						"catalog": types.ObjectType{AttrTypes: map[string]attr.Type{
							"description": types.StringType,
							"name":        types.StringType,
							"version":     types.StringType,
						}},
						"package_count": types.Int64Type,
						"partition": types.ObjectType{AttrTypes: map[string]attr.Type{
							"partition_name": types.StringType,
							"priority_tier":  types.Int64Type,
							"qos":            types.StringType,
						}},
					}}, items)
					data.SoftwareCatalogs = listVal
				}
			} else {
				if data.SoftwareCatalogs.IsUnknown() {
					data.SoftwareCatalogs = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"catalog": types.ObjectType{AttrTypes: map[string]attr.Type{
							"description": types.StringType,
							"name":        types.StringType,
							"version":     types.StringType,
						}},
						"package_count": types.Int64Type,
						"partition": types.ObjectType{AttrTypes: map[string]attr.Type{
							"partition_name": types.StringType,
							"priority_tier":  types.Int64Type,
							"qos":            types.StringType,
						}},
					}})
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
			if val, ok := sourceMap["thumbnail"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Thumbnail = types.StringValue(str)
				}
			} else {
				if data.Thumbnail.IsUnknown() {
					data.Thumbnail = types.StringNull()
				}
			}
			if val, ok := sourceMap["total_cost"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.TotalCost = types.Int64Value(int64(num))
				}
			} else {
				if data.TotalCost.IsUnknown() {
					data.TotalCost = types.Int64Null()
				}
			}
			if val, ok := sourceMap["total_cost_estimated"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.TotalCostEstimated = types.Int64Value(int64(num))
				}
			} else {
				if data.TotalCostEstimated.IsUnknown() {
					data.TotalCostEstimated = types.Int64Null()
				}
			}
			if val, ok := sourceMap["total_customers"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.TotalCustomers = types.Int64Value(int64(num))
				}
			} else {
				if data.TotalCustomers.IsUnknown() {
					data.TotalCustomers = types.Int64Null()
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
			if val, ok := sourceMap["vendor_details"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.VendorDetails = types.StringValue(str)
				}
			} else {
				if data.VendorDetails.IsUnknown() {
					data.VendorDetails = types.StringNull()
				}
			}

			// Map filter parameters from response if available

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
