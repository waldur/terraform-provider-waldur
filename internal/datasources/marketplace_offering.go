package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &MarketplaceOfferingDataSource{}

func NewMarketplaceOfferingDataSource() datasource.DataSource {
	return &MarketplaceOfferingDataSource{}
}

// MarketplaceOfferingDataSource defines the data source implementation.
type MarketplaceOfferingDataSource struct {
	client *client.Client
}

// MarketplaceOfferingDataSourceModel describes the data source data model.
type MarketplaceOfferingDataSourceModel struct {
	UUID                      types.String  `tfsdk:"id"`
	AccessibleViaCalls        types.Bool    `tfsdk:"accessible_via_calls"`
	AllowedCustomerUuid       types.String  `tfsdk:"allowed_customer_uuid"`
	Attributes                types.String  `tfsdk:"attributes"`
	Billable                  types.Bool    `tfsdk:"billable"`
	CanCreateOfferingUser     types.Bool    `tfsdk:"can_create_offering_user"`
	CategoryGroupUuid         types.String  `tfsdk:"category_group_uuid"`
	CategoryUuid              types.String  `tfsdk:"category_uuid"`
	Created                   types.String  `tfsdk:"created"`
	Customer                  types.String  `tfsdk:"customer"`
	CustomerUuid              types.String  `tfsdk:"customer_uuid"`
	Description               types.String  `tfsdk:"description"`
	HasActiveTermsOfService   types.Bool    `tfsdk:"has_active_terms_of_service"`
	HasTermsOfService         types.Bool    `tfsdk:"has_terms_of_service"`
	Keyword                   types.String  `tfsdk:"keyword"`
	Modified                  types.String  `tfsdk:"modified"`
	Name                      types.String  `tfsdk:"name"`
	NameExact                 types.String  `tfsdk:"name_exact"`
	OrganizationGroupUuid     types.String  `tfsdk:"organization_group_uuid"`
	ParentUuid                types.String  `tfsdk:"parent_uuid"`
	ProjectUuid               types.String  `tfsdk:"project_uuid"`
	Query                     types.String  `tfsdk:"query"`
	ResourceCustomerUuid      types.String  `tfsdk:"resource_customer_uuid"`
	ResourceProjectUuid       types.String  `tfsdk:"resource_project_uuid"`
	ScopeUuid                 types.String  `tfsdk:"scope_uuid"`
	ServiceManagerUuid        types.String  `tfsdk:"service_manager_uuid"`
	Shared                    types.Bool    `tfsdk:"shared"`
	State                     types.String  `tfsdk:"state"`
	Type                      types.String  `tfsdk:"type"`
	UserHasConsent            types.Bool    `tfsdk:"user_has_consent"`
	UserHasOfferingUser       types.Bool    `tfsdk:"user_has_offering_user"`
	UuidList                  types.String  `tfsdk:"uuid_list"`
	AccessUrl                 types.String  `tfsdk:"access_url"`
	BackendId                 types.String  `tfsdk:"backend_id"`
	BillingTypeClassification types.String  `tfsdk:"billing_type_classification"`
	Category                  types.String  `tfsdk:"category"`
	CategoryTitle             types.String  `tfsdk:"category_title"`
	CitationCount             types.Int64   `tfsdk:"citation_count"`
	ComplianceChecklist       types.String  `tfsdk:"compliance_checklist"`
	Components                types.List    `tfsdk:"components"`
	Country                   types.String  `tfsdk:"country"`
	DataciteDoi               types.String  `tfsdk:"datacite_doi"`
	Endpoints                 types.List    `tfsdk:"endpoints"`
	Files                     types.List    `tfsdk:"files"`
	FullDescription           types.String  `tfsdk:"full_description"`
	GettingStarted            types.String  `tfsdk:"getting_started"`
	GoogleCalendarIsPublic    types.Bool    `tfsdk:"google_calendar_is_public"`
	GoogleCalendarLink        types.String  `tfsdk:"google_calendar_link"`
	HasComplianceRequirements types.Bool    `tfsdk:"has_compliance_requirements"`
	Image                     types.String  `tfsdk:"image"`
	IntegrationGuide          types.String  `tfsdk:"integration_guide"`
	Latitude                  types.Float64 `tfsdk:"latitude"`
	Longitude                 types.Float64 `tfsdk:"longitude"`
	OrderCount                types.Int64   `tfsdk:"order_count"`
	OrganizationGroups        types.List    `tfsdk:"organization_groups"`
	ParentDescription         types.String  `tfsdk:"parent_description"`
	ParentName                types.String  `tfsdk:"parent_name"`
	Partitions                types.List    `tfsdk:"partitions"`
	PausedReason              types.String  `tfsdk:"paused_reason"`
	Plans                     types.List    `tfsdk:"plans"`
	PrivacyPolicyLink         types.String  `tfsdk:"privacy_policy_link"`
	PromotionCampaigns        types.List    `tfsdk:"promotion_campaigns"`
	Quotas                    types.List    `tfsdk:"quotas"`
	Roles                     types.List    `tfsdk:"roles"`
	Scope                     types.String  `tfsdk:"scope"`
	ScopeErrorMessage         types.String  `tfsdk:"scope_error_message"`
	ScopeName                 types.String  `tfsdk:"scope_name"`
	ScopeState                types.String  `tfsdk:"scope_state"`
	Screenshots               types.List    `tfsdk:"screenshots"`
	Slug                      types.String  `tfsdk:"slug"`
	SoftwareCatalogs          types.List    `tfsdk:"software_catalogs"`
	Thumbnail                 types.String  `tfsdk:"thumbnail"`
	TotalCost                 types.Int64   `tfsdk:"total_cost"`
	TotalCostEstimated        types.Int64   `tfsdk:"total_cost_estimated"`
	TotalCustomers            types.Int64   `tfsdk:"total_customers"`
	Url                       types.String  `tfsdk:"url"`
	VendorDetails             types.String  `tfsdk:"vendor_details"`
}

func (d *MarketplaceOfferingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_offering"
}

func (d *MarketplaceOfferingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Marketplace Offering data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"accessible_via_calls": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Accessible via calls",
			},
			"allowed_customer_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Allowed customer UUID",
			},
			"attributes": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Offering attributes (JSON)",
			},
			"billable": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Billable",
			},
			"can_create_offering_user": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"category_group_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Category group UUID",
			},
			"category_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Category UUID",
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
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description contains",
			},
			"has_active_terms_of_service": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Has Active Terms of Service",
			},
			"has_terms_of_service": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Has Terms of Service",
			},
			"keyword": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Keyword",
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
			"organization_group_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Organization group UUID",
			},
			"parent_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Parent offering UUID",
			},
			"project_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project UUID",
			},
			"query": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Search by offering name, slug or description",
			},
			"resource_customer_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Resource customer UUID",
			},
			"resource_project_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Resource project UUID",
			},
			"scope_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Scope UUID",
			},
			"service_manager_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Service manager UUID",
			},
			"shared": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Shared",
			},
			"state": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Offering state",
			},
			"type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Offering type",
			},
			"user_has_consent": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "User Has Consent",
			},
			"user_has_offering_user": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "User Has Offering User",
			},
			"uuid_list": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Comma-separated offering UUIDs",
			},
			"access_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Publicly accessible offering access URL",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"billing_type_classification": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Classify offering components by billing type. Returns 'limit_only', 'usage_only', or 'mixed'.",
			},
			"category": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"category_title": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"citation_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of citations of a DOI",
			},
			"compliance_checklist": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"components": schema.ListAttribute{
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
			"country": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Country code (ISO 3166-1 alpha-2)",
			},
			"datacite_doi": schema.StringAttribute{
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
			"files": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"created": types.StringType,
					"file":    types.StringType,
					"name":    types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"full_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"getting_started": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"google_calendar_is_public": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"google_calendar_link": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Get the Google Calendar link for an offering.",
			},
			"has_compliance_requirements": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"image": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"integration_guide": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"latitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"longitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"order_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"organization_groups": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"customers_count": types.Int64Type,
					"name":            types.StringType,
					"parent":          types.StringType,
					"parent_name":     types.StringType,
					"parent_uuid":     types.StringType,
					"url":             types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"parent_description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"parent_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"partitions": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
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
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"paused_reason": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"plans": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"archived":     types.BoolType,
					"article_code": types.StringType,
					"backend_id":   types.StringType,
					"components": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"amount":             types.Int64Type,
						"discount_rate":      types.Int64Type,
						"discount_threshold": types.Int64Type,
						"future_price":       types.StringType,
						"measured_unit":      types.StringType,
						"name":               types.StringType,
						"price":              types.StringType,
						"type":               types.StringType,
					}}},
					"description":   types.StringType,
					"init_price":    types.Float64Type,
					"is_active":     types.BoolType,
					"max_amount":    types.Int64Type,
					"minimal_price": types.Float64Type,
					"name":          types.StringType,
					"organization_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"customers_count": types.Int64Type,
						"name":            types.StringType,
						"parent":          types.StringType,
						"parent_name":     types.StringType,
						"parent_uuid":     types.StringType,
						"url":             types.StringType,
					}}},
					"plan_type":       types.StringType,
					"resources_count": types.Int64Type,
					"switch_price":    types.Float64Type,
					"unit":            types.StringType,
					"unit_price":      types.StringType,
					"url":             types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"privacy_policy_link": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"promotion_campaigns": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"description":      types.StringType,
					"discount":         types.Int64Type,
					"discount_type":    types.StringType,
					"end_date":         types.StringType,
					"months":           types.Int64Type,
					"name":             types.StringType,
					"service_provider": types.StringType,
					"start_date":       types.StringType,
					"stock":            types.Int64Type,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"quotas": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"limit": types.Int64Type,
					"name":  types.StringType,
					"usage": types.Int64Type,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"roles": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"name": types.StringType,
					"url":  types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"scope": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"scope_error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"scope_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"scope_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"screenshots": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"created":     types.StringType,
					"description": types.StringType,
					"image":       types.StringType,
					"name":        types.StringType,
					"thumbnail":   types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
			},
			"software_catalogs": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
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
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"thumbnail": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"total_cost": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"total_cost_estimated": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"total_customers": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"vendor_details": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
		},
	}
}

func (d *MarketplaceOfferingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *MarketplaceOfferingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data MarketplaceOfferingDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var item map[string]interface{}

		err := d.client.GetByUUID(ctx, "/api/marketplace-public-offerings/{uuid}/", data.UUID.ValueString(), &item)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Marketplace Offering",
				"An error occurred while reading the Marketplace Offering by UUID: "+err.Error(),
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
		if val, ok := sourceMap["billing_type_classification"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BillingTypeClassification = types.StringValue(str)
			}
		} else {
			if data.BillingTypeClassification.IsUnknown() {
				data.BillingTypeClassification = types.StringNull()
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
				data.Components = listVal
			}
		} else {
			if data.Components.IsUnknown() {
				data.Components = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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
		if val, ok := sourceMap["country"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Country = types.StringValue(str)
			}
		} else {
			if data.Country.IsUnknown() {
				data.Country = types.StringNull()
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
							"components": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"amount":             types.Int64Type,
								"discount_rate":      types.Int64Type,
								"discount_threshold": types.Int64Type,
								"future_price":       types.StringType,
								"measured_unit":      types.StringType,
								"name":               types.StringType,
								"price":              types.StringType,
								"type":               types.StringType,
							}}},
							"description":   types.StringType,
							"init_price":    types.Float64Type,
							"is_active":     types.BoolType,
							"max_amount":    types.Int64Type,
							"minimal_price": types.Float64Type,
							"name":          types.StringType,
							"organization_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"customers_count": types.Int64Type,
								"name":            types.StringType,
								"parent":          types.StringType,
								"parent_name":     types.StringType,
								"parent_uuid":     types.StringType,
								"url":             types.StringType,
							}}},
							"plan_type":       types.StringType,
							"resources_count": types.Int64Type,
							"switch_price":    types.Float64Type,
							"unit":            types.StringType,
							"unit_price":      types.StringType,
							"url":             types.StringType,
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
							"components": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"amount":             types.Int64Type,
								"discount_rate":      types.Int64Type,
								"discount_threshold": types.Int64Type,
								"future_price":       types.StringType,
								"measured_unit":      types.StringType,
								"name":               types.StringType,
								"price":              types.StringType,
								"type":               types.StringType,
							}}}.ElemType),
							"description": func() attr.Value {
								if v, ok := objMap["description"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"init_price": func() attr.Value {
								if v, ok := objMap["init_price"].(float64); ok {
									return types.Float64Value(v)
								}
								return types.Float64Null()
							}(),
							"is_active": func() attr.Value {
								if v, ok := objMap["is_active"].(bool); ok {
									return types.BoolValue(v)
								}
								return types.BoolNull()
							}(),
							"max_amount": func() attr.Value {
								if v, ok := objMap["max_amount"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"minimal_price": func() attr.Value {
								if v, ok := objMap["minimal_price"].(float64); ok {
									return types.Float64Value(v)
								}
								return types.Float64Null()
							}(),
							"name": func() attr.Value {
								if v, ok := objMap["name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"organization_groups": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"customers_count": types.Int64Type,
								"name":            types.StringType,
								"parent":          types.StringType,
								"parent_name":     types.StringType,
								"parent_uuid":     types.StringType,
								"url":             types.StringType,
							}}}.ElemType),
							"plan_type": func() attr.Value {
								if v, ok := objMap["plan_type"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"resources_count": func() attr.Value {
								if v, ok := objMap["resources_count"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"switch_price": func() attr.Value {
								if v, ok := objMap["switch_price"].(float64); ok {
									return types.Float64Value(v)
								}
								return types.Float64Null()
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
					"archived":     types.BoolType,
					"article_code": types.StringType,
					"backend_id":   types.StringType,
					"components": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"amount":             types.Int64Type,
						"discount_rate":      types.Int64Type,
						"discount_threshold": types.Int64Type,
						"future_price":       types.StringType,
						"measured_unit":      types.StringType,
						"name":               types.StringType,
						"price":              types.StringType,
						"type":               types.StringType,
					}}},
					"description":   types.StringType,
					"init_price":    types.Float64Type,
					"is_active":     types.BoolType,
					"max_amount":    types.Int64Type,
					"minimal_price": types.Float64Type,
					"name":          types.StringType,
					"organization_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"customers_count": types.Int64Type,
						"name":            types.StringType,
						"parent":          types.StringType,
						"parent_name":     types.StringType,
						"parent_uuid":     types.StringType,
						"url":             types.StringType,
					}}},
					"plan_type":       types.StringType,
					"resources_count": types.Int64Type,
					"switch_price":    types.Float64Type,
					"unit":            types.StringType,
					"unit_price":      types.StringType,
					"url":             types.StringType,
				}}, items)
				data.Plans = listVal
			}
		} else {
			if data.Plans.IsUnknown() {
				data.Plans = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"archived":     types.BoolType,
					"article_code": types.StringType,
					"backend_id":   types.StringType,
					"components": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"amount":             types.Int64Type,
						"discount_rate":      types.Int64Type,
						"discount_threshold": types.Int64Type,
						"future_price":       types.StringType,
						"measured_unit":      types.StringType,
						"name":               types.StringType,
						"price":              types.StringType,
						"type":               types.StringType,
					}}},
					"description":   types.StringType,
					"init_price":    types.Float64Type,
					"is_active":     types.BoolType,
					"max_amount":    types.Int64Type,
					"minimal_price": types.Float64Type,
					"name":          types.StringType,
					"organization_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"customers_count": types.Int64Type,
						"name":            types.StringType,
						"parent":          types.StringType,
						"parent_name":     types.StringType,
						"parent_uuid":     types.StringType,
						"url":             types.StringType,
					}}},
					"plan_type":       types.StringType,
					"resources_count": types.Int64Type,
					"switch_price":    types.Float64Type,
					"unit":            types.StringType,
					"unit_price":      types.StringType,
					"url":             types.StringType,
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
		if val, ok := sourceMap["promotion_campaigns"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"description":      types.StringType,
							"discount":         types.Int64Type,
							"discount_type":    types.StringType,
							"end_date":         types.StringType,
							"months":           types.Int64Type,
							"name":             types.StringType,
							"service_provider": types.StringType,
							"start_date":       types.StringType,
							"stock":            types.Int64Type,
						}
						attrValues := map[string]attr.Value{
							"description": func() attr.Value {
								if v, ok := objMap["description"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"discount": func() attr.Value {
								if v, ok := objMap["discount"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"discount_type": func() attr.Value {
								if v, ok := objMap["discount_type"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"end_date": func() attr.Value {
								if v, ok := objMap["end_date"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"months": func() attr.Value {
								if v, ok := objMap["months"].(float64); ok {
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
							"service_provider": func() attr.Value {
								if v, ok := objMap["service_provider"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"start_date": func() attr.Value {
								if v, ok := objMap["start_date"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"stock": func() attr.Value {
								if v, ok := objMap["stock"].(float64); ok {
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
					"description":      types.StringType,
					"discount":         types.Int64Type,
					"discount_type":    types.StringType,
					"end_date":         types.StringType,
					"months":           types.Int64Type,
					"name":             types.StringType,
					"service_provider": types.StringType,
					"start_date":       types.StringType,
					"stock":            types.Int64Type,
				}}, items)
				data.PromotionCampaigns = listVal
			}
		} else {
			if data.PromotionCampaigns.IsUnknown() {
				data.PromotionCampaigns = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"description":      types.StringType,
					"discount":         types.Int64Type,
					"discount_type":    types.StringType,
					"end_date":         types.StringType,
					"months":           types.Int64Type,
					"name":             types.StringType,
					"service_provider": types.StringType,
					"start_date":       types.StringType,
					"stock":            types.Int64Type,
				}})
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
		if val, ok := sourceMap["accessible_via_calls"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.AccessibleViaCalls = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["allowed_customer_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AllowedCustomerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["attributes"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Attributes = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["billable"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.Billable = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["can_create_offering_user"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanCreateOfferingUser = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["category_group_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CategoryGroupUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["category_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CategoryUuid = types.StringValue(str)
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
		if val, ok := sourceMap["description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Description = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["has_active_terms_of_service"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.HasActiveTermsOfService = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["has_terms_of_service"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.HasTermsOfService = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["keyword"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Keyword = types.StringValue(str)
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
		if val, ok := sourceMap["organization_group_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OrganizationGroupUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["parent_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ParentUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["query"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Query = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["resource_customer_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ResourceCustomerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["resource_project_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ResourceProjectUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["scope_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ScopeUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_manager_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceManagerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["shared"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.Shared = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.State = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Type = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["user_has_consent"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.UserHasConsent = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["user_has_offering_user"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.UserHasOfferingUser = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["uuid_list"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.UuidList = types.StringValue(str)
			}
		}

	} else {
		// Filter by provided parameters
		var results []map[string]interface{}

		filters := map[string]string{}
		if !data.AccessibleViaCalls.IsNull() {
			filters["accessible_via_calls"] = fmt.Sprintf("%t", data.AccessibleViaCalls.ValueBool())
		}
		if !data.AllowedCustomerUuid.IsNull() {
			filters["allowed_customer_uuid"] = data.AllowedCustomerUuid.ValueString()
		}
		if !data.Attributes.IsNull() {
			filters["attributes"] = data.Attributes.ValueString()
		}
		if !data.Billable.IsNull() {
			filters["billable"] = fmt.Sprintf("%t", data.Billable.ValueBool())
		}
		if !data.CanCreateOfferingUser.IsNull() {
			filters["can_create_offering_user"] = fmt.Sprintf("%t", data.CanCreateOfferingUser.ValueBool())
		}
		if !data.CategoryGroupUuid.IsNull() {
			filters["category_group_uuid"] = data.CategoryGroupUuid.ValueString()
		}
		if !data.CategoryUuid.IsNull() {
			filters["category_uuid"] = data.CategoryUuid.ValueString()
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
		if !data.Description.IsNull() {
			filters["description"] = data.Description.ValueString()
		}
		if !data.HasActiveTermsOfService.IsNull() {
			filters["has_active_terms_of_service"] = fmt.Sprintf("%t", data.HasActiveTermsOfService.ValueBool())
		}
		if !data.HasTermsOfService.IsNull() {
			filters["has_terms_of_service"] = fmt.Sprintf("%t", data.HasTermsOfService.ValueBool())
		}
		if !data.Keyword.IsNull() {
			filters["keyword"] = data.Keyword.ValueString()
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
		if !data.OrganizationGroupUuid.IsNull() {
			filters["organization_group_uuid"] = data.OrganizationGroupUuid.ValueString()
		}
		if !data.ParentUuid.IsNull() {
			filters["parent_uuid"] = data.ParentUuid.ValueString()
		}
		if !data.ProjectUuid.IsNull() {
			filters["project_uuid"] = data.ProjectUuid.ValueString()
		}
		if !data.Query.IsNull() {
			filters["query"] = data.Query.ValueString()
		}
		if !data.ResourceCustomerUuid.IsNull() {
			filters["resource_customer_uuid"] = data.ResourceCustomerUuid.ValueString()
		}
		if !data.ResourceProjectUuid.IsNull() {
			filters["resource_project_uuid"] = data.ResourceProjectUuid.ValueString()
		}
		if !data.ScopeUuid.IsNull() {
			filters["scope_uuid"] = data.ScopeUuid.ValueString()
		}
		if !data.ServiceManagerUuid.IsNull() {
			filters["service_manager_uuid"] = data.ServiceManagerUuid.ValueString()
		}
		if !data.Shared.IsNull() {
			filters["shared"] = fmt.Sprintf("%t", data.Shared.ValueBool())
		}
		if !data.State.IsNull() {
			filters["state"] = data.State.ValueString()
		}
		if !data.Type.IsNull() {
			filters["type"] = data.Type.ValueString()
		}
		if !data.UserHasConsent.IsNull() {
			filters["user_has_consent"] = fmt.Sprintf("%t", data.UserHasConsent.ValueBool())
		}
		if !data.UserHasOfferingUser.IsNull() {
			filters["user_has_offering_user"] = fmt.Sprintf("%t", data.UserHasOfferingUser.ValueBool())
		}
		if !data.UuidList.IsNull() {
			filters["uuid_list"] = data.UuidList.ValueString()
		}

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup marketplace_offering.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/marketplace-public-offerings/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Marketplace Offering",
				"An error occurred while filtering Marketplace Offering: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Marketplace Offering Not Found",
				"No Marketplace Offering found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Marketplace Offerings Found",
				fmt.Sprintf("Found %d Marketplace Offerings with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
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
		if val, ok := sourceMap["billing_type_classification"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BillingTypeClassification = types.StringValue(str)
			}
		} else {
			if data.BillingTypeClassification.IsUnknown() {
				data.BillingTypeClassification = types.StringNull()
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
				data.Components = listVal
			}
		} else {
			if data.Components.IsUnknown() {
				data.Components = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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
		if val, ok := sourceMap["country"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Country = types.StringValue(str)
			}
		} else {
			if data.Country.IsUnknown() {
				data.Country = types.StringNull()
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
							"components": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"amount":             types.Int64Type,
								"discount_rate":      types.Int64Type,
								"discount_threshold": types.Int64Type,
								"future_price":       types.StringType,
								"measured_unit":      types.StringType,
								"name":               types.StringType,
								"price":              types.StringType,
								"type":               types.StringType,
							}}},
							"description":   types.StringType,
							"init_price":    types.Float64Type,
							"is_active":     types.BoolType,
							"max_amount":    types.Int64Type,
							"minimal_price": types.Float64Type,
							"name":          types.StringType,
							"organization_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"customers_count": types.Int64Type,
								"name":            types.StringType,
								"parent":          types.StringType,
								"parent_name":     types.StringType,
								"parent_uuid":     types.StringType,
								"url":             types.StringType,
							}}},
							"plan_type":       types.StringType,
							"resources_count": types.Int64Type,
							"switch_price":    types.Float64Type,
							"unit":            types.StringType,
							"unit_price":      types.StringType,
							"url":             types.StringType,
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
							"components": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"amount":             types.Int64Type,
								"discount_rate":      types.Int64Type,
								"discount_threshold": types.Int64Type,
								"future_price":       types.StringType,
								"measured_unit":      types.StringType,
								"name":               types.StringType,
								"price":              types.StringType,
								"type":               types.StringType,
							}}}.ElemType),
							"description": func() attr.Value {
								if v, ok := objMap["description"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"init_price": func() attr.Value {
								if v, ok := objMap["init_price"].(float64); ok {
									return types.Float64Value(v)
								}
								return types.Float64Null()
							}(),
							"is_active": func() attr.Value {
								if v, ok := objMap["is_active"].(bool); ok {
									return types.BoolValue(v)
								}
								return types.BoolNull()
							}(),
							"max_amount": func() attr.Value {
								if v, ok := objMap["max_amount"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"minimal_price": func() attr.Value {
								if v, ok := objMap["minimal_price"].(float64); ok {
									return types.Float64Value(v)
								}
								return types.Float64Null()
							}(),
							"name": func() attr.Value {
								if v, ok := objMap["name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"organization_groups": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"customers_count": types.Int64Type,
								"name":            types.StringType,
								"parent":          types.StringType,
								"parent_name":     types.StringType,
								"parent_uuid":     types.StringType,
								"url":             types.StringType,
							}}}.ElemType),
							"plan_type": func() attr.Value {
								if v, ok := objMap["plan_type"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"resources_count": func() attr.Value {
								if v, ok := objMap["resources_count"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"switch_price": func() attr.Value {
								if v, ok := objMap["switch_price"].(float64); ok {
									return types.Float64Value(v)
								}
								return types.Float64Null()
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
					"archived":     types.BoolType,
					"article_code": types.StringType,
					"backend_id":   types.StringType,
					"components": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"amount":             types.Int64Type,
						"discount_rate":      types.Int64Type,
						"discount_threshold": types.Int64Type,
						"future_price":       types.StringType,
						"measured_unit":      types.StringType,
						"name":               types.StringType,
						"price":              types.StringType,
						"type":               types.StringType,
					}}},
					"description":   types.StringType,
					"init_price":    types.Float64Type,
					"is_active":     types.BoolType,
					"max_amount":    types.Int64Type,
					"minimal_price": types.Float64Type,
					"name":          types.StringType,
					"organization_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"customers_count": types.Int64Type,
						"name":            types.StringType,
						"parent":          types.StringType,
						"parent_name":     types.StringType,
						"parent_uuid":     types.StringType,
						"url":             types.StringType,
					}}},
					"plan_type":       types.StringType,
					"resources_count": types.Int64Type,
					"switch_price":    types.Float64Type,
					"unit":            types.StringType,
					"unit_price":      types.StringType,
					"url":             types.StringType,
				}}, items)
				data.Plans = listVal
			}
		} else {
			if data.Plans.IsUnknown() {
				data.Plans = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"archived":     types.BoolType,
					"article_code": types.StringType,
					"backend_id":   types.StringType,
					"components": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"amount":             types.Int64Type,
						"discount_rate":      types.Int64Type,
						"discount_threshold": types.Int64Type,
						"future_price":       types.StringType,
						"measured_unit":      types.StringType,
						"name":               types.StringType,
						"price":              types.StringType,
						"type":               types.StringType,
					}}},
					"description":   types.StringType,
					"init_price":    types.Float64Type,
					"is_active":     types.BoolType,
					"max_amount":    types.Int64Type,
					"minimal_price": types.Float64Type,
					"name":          types.StringType,
					"organization_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"customers_count": types.Int64Type,
						"name":            types.StringType,
						"parent":          types.StringType,
						"parent_name":     types.StringType,
						"parent_uuid":     types.StringType,
						"url":             types.StringType,
					}}},
					"plan_type":       types.StringType,
					"resources_count": types.Int64Type,
					"switch_price":    types.Float64Type,
					"unit":            types.StringType,
					"unit_price":      types.StringType,
					"url":             types.StringType,
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
		if val, ok := sourceMap["promotion_campaigns"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"description":      types.StringType,
							"discount":         types.Int64Type,
							"discount_type":    types.StringType,
							"end_date":         types.StringType,
							"months":           types.Int64Type,
							"name":             types.StringType,
							"service_provider": types.StringType,
							"start_date":       types.StringType,
							"stock":            types.Int64Type,
						}
						attrValues := map[string]attr.Value{
							"description": func() attr.Value {
								if v, ok := objMap["description"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"discount": func() attr.Value {
								if v, ok := objMap["discount"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"discount_type": func() attr.Value {
								if v, ok := objMap["discount_type"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"end_date": func() attr.Value {
								if v, ok := objMap["end_date"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"months": func() attr.Value {
								if v, ok := objMap["months"].(float64); ok {
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
							"service_provider": func() attr.Value {
								if v, ok := objMap["service_provider"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"start_date": func() attr.Value {
								if v, ok := objMap["start_date"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"stock": func() attr.Value {
								if v, ok := objMap["stock"].(float64); ok {
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
					"description":      types.StringType,
					"discount":         types.Int64Type,
					"discount_type":    types.StringType,
					"end_date":         types.StringType,
					"months":           types.Int64Type,
					"name":             types.StringType,
					"service_provider": types.StringType,
					"start_date":       types.StringType,
					"stock":            types.Int64Type,
				}}, items)
				data.PromotionCampaigns = listVal
			}
		} else {
			if data.PromotionCampaigns.IsUnknown() {
				data.PromotionCampaigns = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"description":      types.StringType,
					"discount":         types.Int64Type,
					"discount_type":    types.StringType,
					"end_date":         types.StringType,
					"months":           types.Int64Type,
					"name":             types.StringType,
					"service_provider": types.StringType,
					"start_date":       types.StringType,
					"stock":            types.Int64Type,
				}})
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
		if val, ok := sourceMap["accessible_via_calls"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.AccessibleViaCalls = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["allowed_customer_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AllowedCustomerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["attributes"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Attributes = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["billable"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.Billable = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["can_create_offering_user"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanCreateOfferingUser = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["category_group_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CategoryGroupUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["category_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CategoryUuid = types.StringValue(str)
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
		if val, ok := sourceMap["description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Description = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["has_active_terms_of_service"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.HasActiveTermsOfService = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["has_terms_of_service"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.HasTermsOfService = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["keyword"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Keyword = types.StringValue(str)
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
		if val, ok := sourceMap["organization_group_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OrganizationGroupUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["parent_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ParentUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["query"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Query = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["resource_customer_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ResourceCustomerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["resource_project_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ResourceProjectUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["scope_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ScopeUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_manager_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceManagerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["shared"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.Shared = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.State = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Type = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["user_has_consent"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.UserHasConsent = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["user_has_offering_user"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.UserHasOfferingUser = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["uuid_list"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.UuidList = types.StringValue(str)
			}
		}
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
