package offering

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

var _ list.ListResource = &MarketplaceOfferingList{}

type MarketplaceOfferingList struct {
	client *Client
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
				Description: "Offering attributes (JSON)",
				Optional:    true,
			},
			"billable": schema.BoolAttribute{
				Description: "Billable",
				Optional:    true,
			},
			"can_create_offering_user": schema.BoolAttribute{
				Description: "",
				Optional:    true,
			},
			"category_group_uuid": schema.StringAttribute{
				Description: "Category group UUID",
				Optional:    true,
			},
			"category_uuid": schema.StringAttribute{
				Description: "Category UUID",
				Optional:    true,
			},
			"created": schema.StringAttribute{
				Description: "Created after",
				Optional:    true,
			},
			"customer": schema.StringAttribute{
				Description: "Customer URL",
				Optional:    true,
			},
			"customer_uuid": schema.StringAttribute{
				Description: "Customer UUID",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description contains",
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
				Description: "Name",
				Optional:    true,
			},
			"name_exact": schema.StringAttribute{
				Description: "Name (exact)",
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
				Description: "Parent offering UUID",
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
				Description: "Shared",
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

	l.client = NewClient(client)
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
	listResult, err := l.client.ListMarketplaceOffering(ctx, filters)
	if err != nil {
		// Return error diagnostics
		resp.AddError("Failed to list resources", err.Error())
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Stream results
	stream.Results = func(push func(list.ListResult) bool) {
		for _, apiResp := range listResult {
			result := req.NewListResult(ctx)

			// Map item to model
			var data MarketplaceOfferingResourceModel
			model := &data

			var diags diag.Diagnostics

			data.UUID = types.StringPointerValue(apiResp.UUID)
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
			model.Created = types.StringPointerValue(apiResp.Created)
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

			// Set the resource state
			// For ListResource, we generally return the "Resource" state matching the main resource schema.
			// This allows `terraform plan` to correlate if using `terraform query`.

			setDiags := result.Resource.Set(ctx, &data)
			diags.Append(setDiags...)
			result.Diagnostics.Append(diags...)

			// Set Identity if possible (usually UUID)
			if !data.UUID.IsNull() && !data.UUID.IsUnknown() {
				result.Diagnostics.Append(result.Identity.Set(ctx, data.UUID.ValueString())...)
			}

			if !push(result) {
				return
			}
		}
	}
}
