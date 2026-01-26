package offering

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
				MarkdownDescription: "Can create offering user",
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
	AccessUrl                 types.String      `tfsdk:"access_url"`
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
	CustomerName              types.String      `tfsdk:"customer_name"`
	CustomerUuid              types.String      `tfsdk:"customer_uuid"`
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
	IntegrationStatus         types.List        `tfsdk:"integration_status"`
	Latitude                  types.Float64     `tfsdk:"latitude"`
	Longitude                 types.Float64     `tfsdk:"longitude"`
	Name                      types.String      `tfsdk:"name"`
	OrderCount                types.Int64       `tfsdk:"order_count"`
	OrganizationGroups        types.List        `tfsdk:"organization_groups"`
	ParentDescription         types.String      `tfsdk:"parent_description"`
	ParentName                types.String      `tfsdk:"parent_name"`
	ParentUuid                types.String      `tfsdk:"parent_uuid"`
	Partitions                types.List        `tfsdk:"partitions"`
	PausedReason              types.String      `tfsdk:"paused_reason"`
	Plans                     types.List        `tfsdk:"plans"`
	PrivacyPolicyLink         types.String      `tfsdk:"privacy_policy_link"`
	Project                   types.String      `tfsdk:"project"`
	ProjectName               types.String      `tfsdk:"project_name"`
	ProjectUuid               types.String      `tfsdk:"project_uuid"`
	Quotas                    types.List        `tfsdk:"quotas"`
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
	Thumbnail                 types.String      `tfsdk:"thumbnail"`
	TotalCost                 types.Int64       `tfsdk:"total_cost"`
	TotalCostEstimated        types.Int64       `tfsdk:"total_cost_estimated"`
	TotalCustomers            types.Int64       `tfsdk:"total_customers"`
	Type                      types.String      `tfsdk:"type"`
	Url                       types.String      `tfsdk:"url"`
	VendorDetails             types.String      `tfsdk:"vendor_details"`
}

// CopyFrom maps the API response to the model fields.
func (model *MarketplaceOfferingModel) CopyFrom(ctx context.Context, apiResp MarketplaceOfferingResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Billable = types.BoolPointerValue(apiResp.Billable)
	model.BillingTypeClassification = types.StringPointerValue(apiResp.BillingTypeClassification)
	model.Category = types.StringPointerValue(apiResp.Category)
	model.CategoryTitle = types.StringPointerValue(apiResp.CategoryTitle)
	model.CategoryUuid = types.StringPointerValue(apiResp.CategoryUuid)
	model.CitationCount = types.Int64PointerValue(apiResp.CitationCount)
	model.ComplianceChecklist = types.StringPointerValue(apiResp.ComplianceChecklist)

	{
		listValComponents, listDiagsComponents := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
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
		}}, apiResp.Components)
		diags.Append(listDiagsComponents...)
		model.Components = listValComponents
	}
	model.Country = types.StringPointerValue(apiResp.Country)
	valCreated, diagsCreated := timetypes.NewRFC3339PointerValue(apiResp.Created)
	diags.Append(diagsCreated...)
	model.Created = valCreated
	model.Customer = types.StringPointerValue(apiResp.Customer)
	model.DataciteDoi = types.StringPointerValue(apiResp.DataciteDoi)
	model.Description = types.StringPointerValue(apiResp.Description)

	{
		listValEndpoints, listDiagsEndpoints := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"url":  types.StringType,
		}}, apiResp.Endpoints)
		diags.Append(listDiagsEndpoints...)
		model.Endpoints = listValEndpoints
	}

	{
		listValFiles, listDiagsFiles := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"created": types.StringType,
			"file":    types.StringType,
			"name":    types.StringType,
		}}, apiResp.Files)
		diags.Append(listDiagsFiles...)
		model.Files = listValFiles
	}
	model.FullDescription = types.StringPointerValue(apiResp.FullDescription)
	model.GettingStarted = types.StringPointerValue(apiResp.GettingStarted)
	model.GoogleCalendarIsPublic = types.BoolPointerValue(apiResp.GoogleCalendarIsPublic)
	model.GoogleCalendarLink = types.StringPointerValue(apiResp.GoogleCalendarLink)
	model.HasComplianceRequirements = types.BoolPointerValue(apiResp.HasComplianceRequirements)
	model.Image = types.StringPointerValue(apiResp.Image)
	model.IntegrationGuide = types.StringPointerValue(apiResp.IntegrationGuide)

	{
		listValIntegrationStatus, listDiagsIntegrationStatus := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"agent_type":             types.StringType,
			"last_request_timestamp": types.StringType,
			"service_name":           types.StringType,
			"status":                 types.StringType,
		}}, apiResp.IntegrationStatus)
		diags.Append(listDiagsIntegrationStatus...)
		model.IntegrationStatus = listValIntegrationStatus
	}
	model.Latitude = types.Float64PointerValue(apiResp.Latitude)
	model.Longitude = types.Float64PointerValue(apiResp.Longitude)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.OrderCount = types.Int64PointerValue(apiResp.OrderCount)

	{
		listValOrganizationGroups, listDiagsOrganizationGroups := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"customers_count": types.Int64Type,
			"name":            types.StringType,
			"parent":          types.StringType,
			"parent_name":     types.StringType,
			"parent_uuid":     types.StringType,
			"url":             types.StringType,
		}}, apiResp.OrganizationGroups)
		diags.Append(listDiagsOrganizationGroups...)
		model.OrganizationGroups = listValOrganizationGroups
	}
	model.ParentDescription = types.StringPointerValue(apiResp.ParentDescription)
	model.ParentName = types.StringPointerValue(apiResp.ParentName)
	model.ParentUuid = types.StringPointerValue(apiResp.ParentUuid)

	{
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
		}}, apiResp.Partitions)
		diags.Append(listDiagsPartitions...)
		model.Partitions = listValPartitions
	}
	model.PausedReason = types.StringPointerValue(apiResp.PausedReason)

	{
		listValPlans, listDiagsPlans := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"archived":     types.BoolType,
			"article_code": types.StringType,
			"backend_id":   types.StringType,
			"description":  types.StringType,
			"max_amount":   types.Int64Type,
			"name":         types.StringType,
			"unit":         types.StringType,
			"unit_price":   types.StringType,
		}}, apiResp.Plans)
		diags.Append(listDiagsPlans...)
		model.Plans = listValPlans
	}
	model.PrivacyPolicyLink = types.StringPointerValue(apiResp.PrivacyPolicyLink)

	{
		listValQuotas, listDiagsQuotas := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"limit": types.Int64Type,
			"name":  types.StringType,
			"usage": types.Int64Type,
		}}, apiResp.Quotas)
		diags.Append(listDiagsQuotas...)
		model.Quotas = listValQuotas
	}

	{
		listValRoles, listDiagsRoles := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"url":  types.StringType,
		}}, apiResp.Roles)
		diags.Append(listDiagsRoles...)
		model.Roles = listValRoles
	}
	model.Scope = types.StringPointerValue(apiResp.Scope)
	model.ScopeErrorMessage = types.StringPointerValue(apiResp.ScopeErrorMessage)
	model.ScopeName = types.StringPointerValue(apiResp.ScopeName)
	model.ScopeState = types.StringPointerValue(apiResp.ScopeState)
	model.ScopeUuid = types.StringPointerValue(apiResp.ScopeUuid)

	{
		listValScreenshots, listDiagsScreenshots := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"created":     types.StringType,
			"description": types.StringType,
			"image":       types.StringType,
			"name":        types.StringType,
			"thumbnail":   types.StringType,
		}}, apiResp.Screenshots)
		diags.Append(listDiagsScreenshots...)
		model.Screenshots = listValScreenshots
	}
	model.Shared = types.BoolPointerValue(apiResp.Shared)
	model.Slug = types.StringPointerValue(apiResp.Slug)

	{
		listValSoftwareCatalogs, listDiagsSoftwareCatalogs := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
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
		}}, apiResp.SoftwareCatalogs)
		diags.Append(listDiagsSoftwareCatalogs...)
		model.SoftwareCatalogs = listValSoftwareCatalogs
	}
	model.State = types.StringPointerValue(apiResp.State)
	model.Thumbnail = types.StringPointerValue(apiResp.Thumbnail)
	model.TotalCost = types.Int64PointerValue(apiResp.TotalCost)
	model.TotalCostEstimated = types.Int64PointerValue(apiResp.TotalCostEstimated)
	model.TotalCustomers = types.Int64PointerValue(apiResp.TotalCustomers)
	model.Type = types.StringPointerValue(apiResp.Type)
	model.Url = types.StringPointerValue(apiResp.Url)
	model.VendorDetails = types.StringPointerValue(apiResp.VendorDetails)

	return diags
}
