package offering

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type MarketplaceOfferingFiltersModel struct {
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
	ParentUuid              types.String `tfsdk:"parent_uuid"`
	ProjectUuid             types.String `tfsdk:"project_uuid"`
	Query                   types.String `tfsdk:"query"`
	ResourceCustomerUuid    types.String `tfsdk:"resource_customer_uuid"`
	ResourceProjectUuid     types.String `tfsdk:"resource_project_uuid"`
	ScopeUuid               types.String `tfsdk:"scope_uuid"`
	ServiceManagerUuid      types.String `tfsdk:"service_manager_uuid"`
	Shared                  types.Bool   `tfsdk:"shared"`
	TagNamesAnd             types.String `tfsdk:"tag_names_and"`
	TagsAnd                 types.String `tfsdk:"tags_and"`
	UserHasConsent          types.Bool   `tfsdk:"user_has_consent"`
	UserHasOfferingUser     types.Bool   `tfsdk:"user_has_offering_user"`
	UuidList                types.String `tfsdk:"uuid_list"`
}

func (m *MarketplaceOfferingFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Marketplace Offering",
		Attributes: map[string]schema.Attribute{
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
			"tag_names_and": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Tag names with AND logic (comma-separated)",
			},
			"tags_and": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Tag UUIDs with AND logic (comma-separated)",
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
		},
	}
}

type MarketplaceOfferingModel struct {
	UUID                      types.String      `tfsdk:"id"`
	BackendId                 types.String      `tfsdk:"backend_id"`
	Billable                  types.Bool        `tfsdk:"billable"`
	BillingTypeClassification types.String      `tfsdk:"billing_type_classification"`
	Category                  types.String      `tfsdk:"category"`
	CategoryTitle             types.String      `tfsdk:"category_title"`
	CategoryUuid              types.String      `tfsdk:"category_uuid"`
	CitationCount             types.Int64       `tfsdk:"citation_count"`
	ComplianceChecklist       types.String      `tfsdk:"compliance_checklist"`
	Components                types.List        `tfsdk:"components"`
	Country                   types.String      `tfsdk:"country"`
	Created                   timetypes.RFC3339 `tfsdk:"created"`
	Customer                  types.String      `tfsdk:"customer"`
	DataciteDoi               types.String      `tfsdk:"datacite_doi"`
	Description               types.String      `tfsdk:"description"`
	Endpoints                 types.List        `tfsdk:"endpoints"`
	Files                     types.List        `tfsdk:"files"`
	FullDescription           types.String      `tfsdk:"full_description"`
	GettingStarted            types.String      `tfsdk:"getting_started"`
	GoogleCalendarIsPublic    types.Bool        `tfsdk:"google_calendar_is_public"`
	GoogleCalendarLink        types.String      `tfsdk:"google_calendar_link"`
	HasComplianceRequirements types.Bool        `tfsdk:"has_compliance_requirements"`
	Image                     types.String      `tfsdk:"image"`
	IntegrationGuide          types.String      `tfsdk:"integration_guide"`
	IsAccessible              types.Bool        `tfsdk:"is_accessible"`
	Latitude                  types.Float64     `tfsdk:"latitude"`
	Longitude                 types.Float64     `tfsdk:"longitude"`
	Name                      types.String      `tfsdk:"name"`
	Options                   types.Object      `tfsdk:"options"`
	OrderCount                types.Int64       `tfsdk:"order_count"`
	OrganizationGroups        types.List        `tfsdk:"organization_groups"`
	ParentDescription         types.String      `tfsdk:"parent_description"`
	ParentName                types.String      `tfsdk:"parent_name"`
	ParentUuid                types.String      `tfsdk:"parent_uuid"`
	Partitions                types.List        `tfsdk:"partitions"`
	PausedReason              types.String      `tfsdk:"paused_reason"`
	Plans                     types.List        `tfsdk:"plans"`
	PluginOptions             types.Object      `tfsdk:"plugin_options"`
	PrivacyPolicyLink         types.String      `tfsdk:"privacy_policy_link"`
	Project                   types.String      `tfsdk:"project"`
	PromotionCampaigns        types.List        `tfsdk:"promotion_campaigns"`
	Quotas                    types.List        `tfsdk:"quotas"`
	ResourceOptions           types.Object      `tfsdk:"resource_options"`
	Roles                     types.List        `tfsdk:"roles"`
	Scope                     types.String      `tfsdk:"scope"`
	ScopeErrorMessage         types.String      `tfsdk:"scope_error_message"`
	ScopeName                 types.String      `tfsdk:"scope_name"`
	ScopeState                types.String      `tfsdk:"scope_state"`
	ScopeUuid                 types.String      `tfsdk:"scope_uuid"`
	Screenshots               types.List        `tfsdk:"screenshots"`
	Shared                    types.Bool        `tfsdk:"shared"`
	Slug                      types.String      `tfsdk:"slug"`
	SoftwareCatalogs          types.List        `tfsdk:"software_catalogs"`
	State                     types.String      `tfsdk:"state"`
	Tags                      types.Set         `tfsdk:"tags"`
	Thumbnail                 types.String      `tfsdk:"thumbnail"`
	TotalCost                 types.Int64       `tfsdk:"total_cost"`
	TotalCostEstimated        types.Int64       `tfsdk:"total_cost_estimated"`
	TotalCustomers            types.Int64       `tfsdk:"total_customers"`
	Type                      types.String      `tfsdk:"type"`
	Url                       types.String      `tfsdk:"url"`
	UserHasConsent            types.Bool        `tfsdk:"user_has_consent"`
	VendorDetails             types.String      `tfsdk:"vendor_details"`
}

// CopyFrom maps the API response to the model fields.
func (model *MarketplaceOfferingModel) CopyFrom(ctx context.Context, apiResp MarketplaceOfferingResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.BackendId = common.StringPointerValue(apiResp.BackendId)
	model.Billable = types.BoolPointerValue(apiResp.Billable)
	model.BillingTypeClassification = common.StringPointerValue(apiResp.BillingTypeClassification)
	model.Category = common.StringPointerValue(apiResp.Category)
	model.CategoryTitle = common.StringPointerValue(apiResp.CategoryTitle)
	model.CategoryUuid = common.StringPointerValue(apiResp.CategoryUuid)
	model.CitationCount = types.Int64PointerValue(apiResp.CitationCount)
	model.ComplianceChecklist = common.StringPointerValue(apiResp.ComplianceChecklist)

	if apiResp.Components != nil && len(*apiResp.Components) > 0 {
		listValComponents, listDiagsComponents := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
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
			"uuid":                 types.StringType,
		}}, apiResp.Components)
		diags.Append(listDiagsComponents...)
		model.Components = listValComponents
	} else {
		model.Components = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			"uuid":                 types.StringType,
		}})
	}
	model.Country = common.StringPointerValue(apiResp.Country)
	valCreated, diagsCreated := timetypes.NewRFC3339PointerValue(apiResp.Created)
	diags.Append(diagsCreated...)
	model.Created = valCreated
	model.Customer = common.StringPointerValue(apiResp.Customer)
	model.DataciteDoi = common.StringPointerValue(apiResp.DataciteDoi)
	model.Description = common.StringPointerValue(apiResp.Description)

	if apiResp.Endpoints != nil && len(*apiResp.Endpoints) > 0 {
		listValEndpoints, listDiagsEndpoints := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"url":  types.StringType,
			"uuid": types.StringType,
		}}, apiResp.Endpoints)
		diags.Append(listDiagsEndpoints...)
		model.Endpoints = listValEndpoints
	} else {
		model.Endpoints = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"url":  types.StringType,
			"uuid": types.StringType,
		}})
	}

	if apiResp.Files != nil && len(*apiResp.Files) > 0 {
		listValFiles, listDiagsFiles := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"created": types.StringType,
			"file":    types.StringType,
			"name":    types.StringType,
		}}, apiResp.Files)
		diags.Append(listDiagsFiles...)
		model.Files = listValFiles
	} else {
		model.Files = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
			"created": types.StringType,
			"file":    types.StringType,
			"name":    types.StringType,
		}})
	}
	model.FullDescription = common.StringPointerValue(apiResp.FullDescription)
	model.GettingStarted = common.StringPointerValue(apiResp.GettingStarted)
	model.GoogleCalendarIsPublic = types.BoolPointerValue(apiResp.GoogleCalendarIsPublic)
	model.GoogleCalendarLink = common.StringPointerValue(apiResp.GoogleCalendarLink)
	model.HasComplianceRequirements = types.BoolPointerValue(apiResp.HasComplianceRequirements)
	model.Image = common.StringPointerValue(apiResp.Image)
	model.IntegrationGuide = common.StringPointerValue(apiResp.IntegrationGuide)
	model.IsAccessible = types.BoolPointerValue(apiResp.IsAccessible)
	model.Latitude = types.Float64PointerValue(apiResp.Latitude.Float64Ptr())
	model.Longitude = types.Float64PointerValue(apiResp.Longitude.Float64Ptr())
	model.Name = common.StringPointerValue(apiResp.Name)
	if apiResp.Options != nil {
		objValOptions, objDiagsOptions := types.ObjectValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"order": types.ListType{ElemType: types.StringType},
		}}.AttrTypes, *apiResp.Options)
		diags.Append(objDiagsOptions...)
		model.Options = objValOptions
	} else {
		model.Options = types.ObjectNull(types.ObjectType{AttrTypes: map[string]attr.Type{
			"order": types.ListType{ElemType: types.StringType},
		}}.AttrTypes)
	}
	model.OrderCount = types.Int64PointerValue(apiResp.OrderCount)

	if apiResp.OrganizationGroups != nil && len(*apiResp.OrganizationGroups) > 0 {
		listValOrganizationGroups, listDiagsOrganizationGroups := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"customers_count": types.Int64Type,
			"name":            types.StringType,
			"parent":          types.StringType,
			"parent_name":     types.StringType,
			"parent_uuid":     types.StringType,
			"url":             types.StringType,
			"uuid":            types.StringType,
		}}, apiResp.OrganizationGroups)
		diags.Append(listDiagsOrganizationGroups...)
		model.OrganizationGroups = listValOrganizationGroups
	} else {
		model.OrganizationGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
			"customers_count": types.Int64Type,
			"name":            types.StringType,
			"parent":          types.StringType,
			"parent_name":     types.StringType,
			"parent_uuid":     types.StringType,
			"url":             types.StringType,
			"uuid":            types.StringType,
		}})
	}
	model.ParentDescription = common.StringPointerValue(apiResp.ParentDescription)
	model.ParentName = common.StringPointerValue(apiResp.ParentName)
	model.ParentUuid = common.StringPointerValue(apiResp.ParentUuid)

	if apiResp.Partitions != nil && len(*apiResp.Partitions) > 0 {
		listValPartitions, listDiagsPartitions := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
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
			"uuid":                types.StringType,
		}}, apiResp.Partitions)
		diags.Append(listDiagsPartitions...)
		model.Partitions = listValPartitions
	} else {
		model.Partitions = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			"uuid":                types.StringType,
		}})
	}
	model.PausedReason = common.StringPointerValue(apiResp.PausedReason)

	if apiResp.Plans != nil && len(*apiResp.Plans) > 0 {
		listValPlans, listDiagsPlans := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
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
				"uuid":            types.StringType,
			}}},
			"plan_type":       types.StringType,
			"resources_count": types.Int64Type,
			"switch_price":    types.Float64Type,
			"unit":            types.StringType,
			"unit_price":      types.StringType,
			"url":             types.StringType,
			"uuid":            types.StringType,
		}}, apiResp.Plans)
		diags.Append(listDiagsPlans...)
		model.Plans = listValPlans
	} else {
		model.Plans = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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
				"uuid":            types.StringType,
			}}},
			"plan_type":       types.StringType,
			"resources_count": types.Int64Type,
			"switch_price":    types.Float64Type,
			"unit":            types.StringType,
			"unit_price":      types.StringType,
			"url":             types.StringType,
			"uuid":            types.StringType,
		}})
	}
	if apiResp.PluginOptions != nil {
		objValPluginOptions, objDiagsPluginOptions := types.ObjectValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"auto_approve_in_service_provider_projects":             types.BoolType,
			"auto_approve_marketplace_script":                       types.BoolType,
			"auto_approve_remote_orders":                            types.BoolType,
			"backend_id_display_label":                              types.StringType,
			"can_restore_resource":                                  types.BoolType,
			"conceal_billing_data":                                  types.BoolType,
			"create_orders_on_resource_option_change":               types.BoolType,
			"default_internal_network_mtu":                          types.Int64Type,
			"default_resource_termination_offset_in_days":           types.Int64Type,
			"deployment_mode":                                       types.StringType,
			"disable_autoapprove":                                   types.BoolType,
			"enable_display_of_order_actions_for_service_provider":  types.BoolType,
			"enable_issues_for_membership_changes":                  types.BoolType,
			"enable_purchase_order_upload":                          types.BoolType,
			"flavors_regex":                                         types.StringType,
			"heappe_cluster_id":                                     types.StringType,
			"heappe_local_base_path":                                types.StringType,
			"heappe_url":                                            types.StringType,
			"heappe_username":                                       types.StringType,
			"highlight_backend_id_display":                          types.BoolType,
			"homedir_prefix":                                        types.StringType,
			"initial_primarygroup_number":                           types.Int64Type,
			"initial_uidnumber":                                     types.Int64Type,
			"initial_usergroup_number":                              types.Int64Type,
			"is_resource_termination_date_required":                 types.BoolType,
			"latest_date_for_resource_termination":                  types.StringType,
			"managed_rancher_load_balancer_data_volume_size_gb":     types.Int64Type,
			"managed_rancher_load_balancer_data_volume_type_name":   types.StringType,
			"managed_rancher_load_balancer_flavor_name":             types.StringType,
			"managed_rancher_load_balancer_system_volume_size_gb":   types.Int64Type,
			"managed_rancher_load_balancer_system_volume_type_name": types.StringType,
			"managed_rancher_server_data_volume_size_gb":            types.Int64Type,
			"managed_rancher_server_data_volume_type_name":          types.StringType,
			"managed_rancher_server_flavor_name":                    types.StringType,
			"managed_rancher_server_system_volume_size_gb":          types.Int64Type,
			"managed_rancher_server_system_volume_type_name":        types.StringType,
			"managed_rancher_tenant_max_cpu":                        types.Int64Type,
			"managed_rancher_tenant_max_disk":                       types.Int64Type,
			"managed_rancher_tenant_max_ram":                        types.Int64Type,
			"managed_rancher_worker_system_volume_size_gb":          types.Int64Type,
			"managed_rancher_worker_system_volume_type_name":        types.StringType,
			"max_instances":                                         types.Int64Type,
			"max_resource_termination_offset_in_days":               types.Int64Type,
			"max_security_groups":                                   types.Int64Type,
			"max_volumes":                                           types.Int64Type,
			"maximal_resource_count_per_project":                    types.Int64Type,
			"minimal_team_count_for_provisioning":                   types.Int64Type,
			"openstack_offering_uuid_list":                          types.ListType{ElemType: types.StringType},
			"project_permanent_directory":                           types.StringType,
			"require_purchase_order_upload":                         types.BoolType,
			"required_team_role_for_provisioning":                   types.StringType,
			"resource_expiration_threshold":                         types.Int64Type,
			"scratch_project_directory":                             types.StringType,
			"service_provider_can_create_offering_user":             types.BoolType,
			"slurm_periodic_policy_enabled":                         types.BoolType,
			"snapshot_size_limit_gb":                                types.Int64Type,
			"storage_mode":                                          types.StringType,
			"supports_downscaling":                                  types.BoolType,
			"supports_pausing":                                      types.BoolType,
			"unique_resource_per_attribute":                         types.StringType,
			"username_anonymized_prefix":                            types.StringType,
			"username_generation_policy":                            types.StringType,
		}}.AttrTypes, *apiResp.PluginOptions)
		diags.Append(objDiagsPluginOptions...)
		model.PluginOptions = objValPluginOptions
	} else {
		model.PluginOptions = types.ObjectNull(types.ObjectType{AttrTypes: map[string]attr.Type{
			"auto_approve_in_service_provider_projects":             types.BoolType,
			"auto_approve_marketplace_script":                       types.BoolType,
			"auto_approve_remote_orders":                            types.BoolType,
			"backend_id_display_label":                              types.StringType,
			"can_restore_resource":                                  types.BoolType,
			"conceal_billing_data":                                  types.BoolType,
			"create_orders_on_resource_option_change":               types.BoolType,
			"default_internal_network_mtu":                          types.Int64Type,
			"default_resource_termination_offset_in_days":           types.Int64Type,
			"deployment_mode":                                       types.StringType,
			"disable_autoapprove":                                   types.BoolType,
			"enable_display_of_order_actions_for_service_provider":  types.BoolType,
			"enable_issues_for_membership_changes":                  types.BoolType,
			"enable_purchase_order_upload":                          types.BoolType,
			"flavors_regex":                                         types.StringType,
			"heappe_cluster_id":                                     types.StringType,
			"heappe_local_base_path":                                types.StringType,
			"heappe_url":                                            types.StringType,
			"heappe_username":                                       types.StringType,
			"highlight_backend_id_display":                          types.BoolType,
			"homedir_prefix":                                        types.StringType,
			"initial_primarygroup_number":                           types.Int64Type,
			"initial_uidnumber":                                     types.Int64Type,
			"initial_usergroup_number":                              types.Int64Type,
			"is_resource_termination_date_required":                 types.BoolType,
			"latest_date_for_resource_termination":                  types.StringType,
			"managed_rancher_load_balancer_data_volume_size_gb":     types.Int64Type,
			"managed_rancher_load_balancer_data_volume_type_name":   types.StringType,
			"managed_rancher_load_balancer_flavor_name":             types.StringType,
			"managed_rancher_load_balancer_system_volume_size_gb":   types.Int64Type,
			"managed_rancher_load_balancer_system_volume_type_name": types.StringType,
			"managed_rancher_server_data_volume_size_gb":            types.Int64Type,
			"managed_rancher_server_data_volume_type_name":          types.StringType,
			"managed_rancher_server_flavor_name":                    types.StringType,
			"managed_rancher_server_system_volume_size_gb":          types.Int64Type,
			"managed_rancher_server_system_volume_type_name":        types.StringType,
			"managed_rancher_tenant_max_cpu":                        types.Int64Type,
			"managed_rancher_tenant_max_disk":                       types.Int64Type,
			"managed_rancher_tenant_max_ram":                        types.Int64Type,
			"managed_rancher_worker_system_volume_size_gb":          types.Int64Type,
			"managed_rancher_worker_system_volume_type_name":        types.StringType,
			"max_instances":                                         types.Int64Type,
			"max_resource_termination_offset_in_days":               types.Int64Type,
			"max_security_groups":                                   types.Int64Type,
			"max_volumes":                                           types.Int64Type,
			"maximal_resource_count_per_project":                    types.Int64Type,
			"minimal_team_count_for_provisioning":                   types.Int64Type,
			"openstack_offering_uuid_list":                          types.ListType{ElemType: types.StringType},
			"project_permanent_directory":                           types.StringType,
			"require_purchase_order_upload":                         types.BoolType,
			"required_team_role_for_provisioning":                   types.StringType,
			"resource_expiration_threshold":                         types.Int64Type,
			"scratch_project_directory":                             types.StringType,
			"service_provider_can_create_offering_user":             types.BoolType,
			"slurm_periodic_policy_enabled":                         types.BoolType,
			"snapshot_size_limit_gb":                                types.Int64Type,
			"storage_mode":                                          types.StringType,
			"supports_downscaling":                                  types.BoolType,
			"supports_pausing":                                      types.BoolType,
			"unique_resource_per_attribute":                         types.StringType,
			"username_anonymized_prefix":                            types.StringType,
			"username_generation_policy":                            types.StringType,
		}}.AttrTypes)
	}
	model.PrivacyPolicyLink = common.StringPointerValue(apiResp.PrivacyPolicyLink)
	model.Project = common.StringPointerValue(apiResp.Project)

	if apiResp.PromotionCampaigns != nil && len(*apiResp.PromotionCampaigns) > 0 {
		listValPromotionCampaigns, listDiagsPromotionCampaigns := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"discount":         types.Int64Type,
			"discount_type":    types.StringType,
			"end_date":         types.StringType,
			"months":           types.Int64Type,
			"name":             types.StringType,
			"service_provider": types.StringType,
			"start_date":       types.StringType,
			"stock":            types.Int64Type,
			"uuid":             types.StringType,
		}}, apiResp.PromotionCampaigns)
		diags.Append(listDiagsPromotionCampaigns...)
		model.PromotionCampaigns = listValPromotionCampaigns
	} else {
		model.PromotionCampaigns = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"discount":         types.Int64Type,
			"discount_type":    types.StringType,
			"end_date":         types.StringType,
			"months":           types.Int64Type,
			"name":             types.StringType,
			"service_provider": types.StringType,
			"start_date":       types.StringType,
			"stock":            types.Int64Type,
			"uuid":             types.StringType,
		}})
	}

	if apiResp.Quotas != nil && len(*apiResp.Quotas) > 0 {
		listValQuotas, listDiagsQuotas := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"limit": types.Int64Type,
			"name":  types.StringType,
			"usage": types.Int64Type,
		}}, apiResp.Quotas)
		diags.Append(listDiagsQuotas...)
		model.Quotas = listValQuotas
	} else {
		model.Quotas = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
			"limit": types.Int64Type,
			"name":  types.StringType,
			"usage": types.Int64Type,
		}})
	}
	if apiResp.ResourceOptions != nil {
		objValResourceOptions, objDiagsResourceOptions := types.ObjectValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"order": types.ListType{ElemType: types.StringType},
		}}.AttrTypes, *apiResp.ResourceOptions)
		diags.Append(objDiagsResourceOptions...)
		model.ResourceOptions = objValResourceOptions
	} else {
		model.ResourceOptions = types.ObjectNull(types.ObjectType{AttrTypes: map[string]attr.Type{
			"order": types.ListType{ElemType: types.StringType},
		}}.AttrTypes)
	}

	if apiResp.Roles != nil && len(*apiResp.Roles) > 0 {
		listValRoles, listDiagsRoles := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"url":  types.StringType,
			"uuid": types.StringType,
		}}, apiResp.Roles)
		diags.Append(listDiagsRoles...)
		model.Roles = listValRoles
	} else {
		model.Roles = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"url":  types.StringType,
			"uuid": types.StringType,
		}})
	}
	model.Scope = common.StringPointerValue(apiResp.Scope)
	model.ScopeErrorMessage = common.StringPointerValue(apiResp.ScopeErrorMessage)
	model.ScopeName = common.StringPointerValue(apiResp.ScopeName)
	model.ScopeState = common.StringPointerValue(apiResp.ScopeState)
	model.ScopeUuid = common.StringPointerValue(apiResp.ScopeUuid)

	if apiResp.Screenshots != nil && len(*apiResp.Screenshots) > 0 {
		listValScreenshots, listDiagsScreenshots := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"created":     types.StringType,
			"description": types.StringType,
			"image":       types.StringType,
			"name":        types.StringType,
			"thumbnail":   types.StringType,
			"uuid":        types.StringType,
		}}, apiResp.Screenshots)
		diags.Append(listDiagsScreenshots...)
		model.Screenshots = listValScreenshots
	} else {
		model.Screenshots = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
			"created":     types.StringType,
			"description": types.StringType,
			"image":       types.StringType,
			"name":        types.StringType,
			"thumbnail":   types.StringType,
			"uuid":        types.StringType,
		}})
	}
	model.Shared = types.BoolPointerValue(apiResp.Shared)
	model.Slug = common.StringPointerValue(apiResp.Slug)

	if apiResp.SoftwareCatalogs != nil && len(*apiResp.SoftwareCatalogs) > 0 {
		listValSoftwareCatalogs, listDiagsSoftwareCatalogs := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"catalog": types.ObjectType{AttrTypes: map[string]attr.Type{
				"description": types.StringType,
				"name":        types.StringType,
				"uuid":        types.StringType,
				"version":     types.StringType,
			}},
			"package_count": types.Int64Type,
			"partition": types.ObjectType{AttrTypes: map[string]attr.Type{
				"partition_name": types.StringType,
				"priority_tier":  types.Int64Type,
				"qos":            types.StringType,
				"uuid":           types.StringType,
			}},
			"uuid": types.StringType,
		}}, apiResp.SoftwareCatalogs)
		diags.Append(listDiagsSoftwareCatalogs...)
		model.SoftwareCatalogs = listValSoftwareCatalogs
	} else {
		model.SoftwareCatalogs = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
			"catalog": types.ObjectType{AttrTypes: map[string]attr.Type{
				"description": types.StringType,
				"name":        types.StringType,
				"uuid":        types.StringType,
				"version":     types.StringType,
			}},
			"package_count": types.Int64Type,
			"partition": types.ObjectType{AttrTypes: map[string]attr.Type{
				"partition_name": types.StringType,
				"priority_tier":  types.Int64Type,
				"qos":            types.StringType,
				"uuid":           types.StringType,
			}},
			"uuid": types.StringType,
		}})
	}
	model.State = common.StringPointerValue(apiResp.State)
	if apiResp.Tags != nil && len(*apiResp.Tags) > 0 {
		setValTags, setDiagsTags := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"uuid": types.StringType,
		}}, apiResp.Tags)
		diags.Append(setDiagsTags...)
		model.Tags = setValTags
	} else {
		model.Tags = types.SetNull(types.ObjectType{AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"uuid": types.StringType,
		}})
	}
	model.Thumbnail = common.StringPointerValue(apiResp.Thumbnail)
	model.TotalCost = types.Int64PointerValue(apiResp.TotalCost)
	model.TotalCostEstimated = types.Int64PointerValue(apiResp.TotalCostEstimated)
	model.TotalCustomers = types.Int64PointerValue(apiResp.TotalCustomers)
	model.Type = common.StringPointerValue(apiResp.Type)
	model.Url = common.StringPointerValue(apiResp.Url)
	model.UserHasConsent = types.BoolPointerValue(apiResp.UserHasConsent)
	model.VendorDetails = common.StringPointerValue(apiResp.VendorDetails)

	return diags
}
