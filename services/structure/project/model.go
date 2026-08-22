package project

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

func AffiliationType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"abbreviation":   types.StringType,
		"address":        types.StringType,
		"code":           types.StringType,
		"country":        types.StringType,
		"description":    types.StringType,
		"email":          types.StringType,
		"homepage":       types.StringType,
		"name":           types.StringType,
		"projects_count": types.Int64Type,
		"url":            types.StringType,
		"uuid":           types.StringType,
	}}
}
func BillingPriceEstimateType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"current":     types.Float64Type,
		"tax":         types.Float64Type,
		"tax_current": types.Float64Type,
		"total":       types.Float64Type,
	}}
}
func ProjectMetadataAnswerType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"question":      types.StringType,
		"question_type": types.StringType,
		"question_uuid": types.StringType,
	}}
}

type StructureProjectFiltersModel struct {
	AccountingIsRunning     types.Bool   `tfsdk:"accounting_is_running"`
	AffiliationName         types.String `tfsdk:"affiliation_name"`
	BackendId               types.String `tfsdk:"backend_id"`
	CanAdmin                types.Bool   `tfsdk:"can_admin"`
	CanManage               types.Bool   `tfsdk:"can_manage"`
	ConcealFinishedProjects types.Bool   `tfsdk:"conceal_finished_projects"`
	Created                 types.String `tfsdk:"created"`
	CreatedBefore           types.String `tfsdk:"created_before"`
	CustomerAbbreviation    types.String `tfsdk:"customer_abbreviation"`
	CustomerName            types.String `tfsdk:"customer_name"`
	CustomerNativeName      types.String `tfsdk:"customer_native_name"`
	Description             types.String `tfsdk:"description"`
	HasAffiliation          types.Bool   `tfsdk:"has_affiliation"`
	IncludeTerminated       types.Bool   `tfsdk:"include_terminated"`
	IsRemoved               types.Bool   `tfsdk:"is_removed"`
	Modified                types.String `tfsdk:"modified"`
	ModifiedBefore          types.String `tfsdk:"modified_before"`
	Name                    types.String `tfsdk:"name"`
	NameExact               types.String `tfsdk:"name_exact"`
	Query                   types.String `tfsdk:"query"`
	ScienceDomainUuid       types.String `tfsdk:"science_domain_uuid"`
	ScienceSubDomainUuid    types.String `tfsdk:"science_sub_domain_uuid"`
	Slug                    types.String `tfsdk:"slug"`
	UserUuid                types.String `tfsdk:"user_uuid"`
	UserUuidWithActiveRole  types.String `tfsdk:"user_uuid_with_active_role"`
}

func (m *StructureProjectFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Structure Project",
		Attributes: map[string]schema.Attribute{
			"accounting_is_running": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Filter by whether accounting is running.",
			},
			"affiliation_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Affiliation name",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "ID of the backend",
			},
			"can_admin": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Return a list of projects where current user is admin.",
			},
			"can_manage": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Return a list of projects where current user is manager or a customer owner.",
			},
			"conceal_finished_projects": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Conceal finished projects",
			},
			"created": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Created after",
			},
			"created_before": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Created before",
			},
			"customer_abbreviation": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer abbreviation",
			},
			"customer_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer name",
			},
			"customer_native_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer native name",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description",
			},
			"has_affiliation": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Filter projects that have an affiliation.",
			},
			"include_terminated": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Include soft-deleted (terminated) projects. Only available to staff and support users, or users with organizational roles who can see their terminated projects.",
			},
			"is_removed": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Is removed",
			},
			"modified": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Modified after",
			},
			"modified_before": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Modified before",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name (exact)",
			},
			"query": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Filter by name, slug, UUID, backend ID or resource effective ID",
			},
			"science_domain_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Science domain UUID",
			},
			"science_sub_domain_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Science sub-domain UUID",
			},
			"slug": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Slug",
			},
			"user_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Filter by user UUID.",
			},
			"user_uuid_with_active_role": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Filter projects where the given user has a role.",
			},
		},
	}
}

type StructureProjectModel struct {
	UUID                                 types.String      `tfsdk:"id"`
	Affiliation                          types.Object      `tfsdk:"affiliation"`
	AffiliationCode                      types.String      `tfsdk:"affiliation_code"`
	AffiliationName                      types.String      `tfsdk:"affiliation_name"`
	AffiliationUuid                      types.String      `tfsdk:"affiliation_uuid"`
	BackendId                            types.String      `tfsdk:"backend_id"`
	BillingPriceEstimate                 types.Object      `tfsdk:"billing_price_estimate"`
	Customer                             types.String      `tfsdk:"customer"`
	CustomerDisplayBillingInfoInProjects types.Bool        `tfsdk:"customer_display_billing_info_in_projects"`
	CustomerGracePeriodDays              types.Int64       `tfsdk:"customer_grace_period_days"`
	CustomerSlug                         types.String      `tfsdk:"customer_slug"`
	Description                          types.String      `tfsdk:"description"`
	EffectiveEndDate                     types.String      `tfsdk:"effective_end_date"`
	EndDate                              types.String      `tfsdk:"end_date"`
	EndDateRequestedBy                   types.String      `tfsdk:"end_date_requested_by"`
	EndDateUpdatedAt                     timetypes.RFC3339 `tfsdk:"end_date_updated_at"`
	GracePeriodDays                      types.Int64       `tfsdk:"grace_period_days"`
	Image                                types.String      `tfsdk:"image"`
	IsInGracePeriod                      types.Bool        `tfsdk:"is_in_grace_period"`
	IsIndustry                           types.Bool        `tfsdk:"is_industry"`
	IsRemoved                            types.Bool        `tfsdk:"is_removed"`
	Kind                                 types.String      `tfsdk:"kind"`
	MarketplaceResourceCount             types.Map         `tfsdk:"marketplace_resource_count"`
	MaxServiceAccounts                   types.Int64       `tfsdk:"max_service_accounts"`
	Name                                 types.String      `tfsdk:"name"`
	OecdFos2007Code                      types.String      `tfsdk:"oecd_fos_2007_code"`
	OecdFos2007Label                     types.String      `tfsdk:"oecd_fos_2007_label"`
	ProjectCredit                        types.Float64     `tfsdk:"project_credit"`
	ProjectMetadata                      types.List        `tfsdk:"project_metadata"`
	ResourcesCount                       types.Int64       `tfsdk:"resources_count"`
	ScienceDomainCode                    types.String      `tfsdk:"science_domain_code"`
	ScienceDomainName                    types.String      `tfsdk:"science_domain_name"`
	ScienceDomainUuid                    types.String      `tfsdk:"science_domain_uuid"`
	ScienceSubDomain                     types.String      `tfsdk:"science_sub_domain"`
	ScienceSubDomainCode                 types.String      `tfsdk:"science_sub_domain_code"`
	ScienceSubDomainName                 types.String      `tfsdk:"science_sub_domain_name"`
	Slug                                 types.String      `tfsdk:"slug"`
	StaffNotes                           types.String      `tfsdk:"staff_notes"`
	StartDate                            types.String      `tfsdk:"start_date"`
	TerminationMetadata                  types.Map         `tfsdk:"termination_metadata"`
	Type                                 types.String      `tfsdk:"type"`
	TypeName                             types.String      `tfsdk:"type_name"`
	TypeUuid                             types.String      `tfsdk:"type_uuid"`
	Url                                  types.String      `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *StructureProjectModel) CopyFrom(ctx context.Context, apiResp StructureProjectResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)

	if apiResp.Affiliation != nil {
		valAffiliation, diagsAffiliation := types.ObjectValueFrom(ctx, AffiliationType().AttrTypes, *apiResp.Affiliation)
		diags.Append(diagsAffiliation...)
		model.Affiliation = valAffiliation
	} else {
		model.Affiliation = types.ObjectNull(AffiliationType().AttrTypes)
	}

	model.AffiliationCode = common.StringPointerValue(apiResp.AffiliationCode)

	model.AffiliationName = common.StringPointerValue(apiResp.AffiliationName)

	model.AffiliationUuid = common.StringPointerValue(apiResp.AffiliationUuid)

	model.BackendId = common.StringPointerValue(apiResp.BackendId)

	if apiResp.BillingPriceEstimate != nil {
		valBillingPriceEstimate, diagsBillingPriceEstimate := types.ObjectValueFrom(ctx, BillingPriceEstimateType().AttrTypes, *apiResp.BillingPriceEstimate)
		diags.Append(diagsBillingPriceEstimate...)
		model.BillingPriceEstimate = valBillingPriceEstimate
	} else {
		model.BillingPriceEstimate = types.ObjectNull(BillingPriceEstimateType().AttrTypes)
	}

	model.Customer = common.StringPointerValue(apiResp.Customer)

	model.CustomerDisplayBillingInfoInProjects = types.BoolPointerValue(apiResp.CustomerDisplayBillingInfoInProjects)

	model.CustomerGracePeriodDays = types.Int64PointerValue(apiResp.CustomerGracePeriodDays)

	model.CustomerSlug = common.StringPointerValue(apiResp.CustomerSlug)

	model.Description = common.StringPointerValue(apiResp.Description)

	model.EffectiveEndDate = common.StringPointerValue(apiResp.EffectiveEndDate)

	model.EndDate = common.StringPointerValue(apiResp.EndDate)

	model.EndDateRequestedBy = common.StringPointerValue(apiResp.EndDateRequestedBy)

	valEndDateUpdatedAt, diagsEndDateUpdatedAt := timetypes.NewRFC3339PointerValue(apiResp.EndDateUpdatedAt)
	diags.Append(diagsEndDateUpdatedAt...)
	model.EndDateUpdatedAt = valEndDateUpdatedAt

	model.GracePeriodDays = types.Int64PointerValue(apiResp.GracePeriodDays)

	model.Image = common.StringPointerValue(apiResp.Image)

	model.IsInGracePeriod = types.BoolPointerValue(apiResp.IsInGracePeriod)

	model.IsIndustry = types.BoolPointerValue(apiResp.IsIndustry)

	model.IsRemoved = types.BoolPointerValue(apiResp.IsRemoved)

	model.Kind = common.StringPointerValue(apiResp.Kind)

	if apiResp.MarketplaceResourceCount != nil {
		valMarketplaceResourceCount, diagsMarketplaceResourceCount := types.MapValueFrom(ctx, types.Int64Type, apiResp.MarketplaceResourceCount)
		diags.Append(diagsMarketplaceResourceCount...)
		model.MarketplaceResourceCount = valMarketplaceResourceCount
	} else {
		model.MarketplaceResourceCount = types.MapNull(types.Int64Type)
	}

	model.MaxServiceAccounts = types.Int64PointerValue(apiResp.MaxServiceAccounts)

	model.Name = common.StringPointerValue(apiResp.Name)

	model.OecdFos2007Code = common.StringPointerValue(apiResp.OecdFos2007Code)

	model.OecdFos2007Label = common.StringPointerValue(apiResp.OecdFos2007Label)

	model.ProjectCredit = types.Float64PointerValue(apiResp.ProjectCredit.Float64Ptr())

	if apiResp.ProjectMetadata != nil {
		valProjectMetadata, diagsProjectMetadata := types.ListValueFrom(ctx, ProjectMetadataAnswerType(), apiResp.ProjectMetadata)
		diags.Append(diagsProjectMetadata...)
		model.ProjectMetadata = valProjectMetadata
	} else {
		model.ProjectMetadata = types.ListNull(ProjectMetadataAnswerType())
	}

	model.ResourcesCount = types.Int64PointerValue(apiResp.ResourcesCount)

	model.ScienceDomainCode = common.StringPointerValue(apiResp.ScienceDomainCode)

	model.ScienceDomainName = common.StringPointerValue(apiResp.ScienceDomainName)

	model.ScienceDomainUuid = common.StringPointerValue(apiResp.ScienceDomainUuid)

	model.ScienceSubDomain = common.StringPointerValue(apiResp.ScienceSubDomain)

	model.ScienceSubDomainCode = common.StringPointerValue(apiResp.ScienceSubDomainCode)

	model.ScienceSubDomainName = common.StringPointerValue(apiResp.ScienceSubDomainName)

	model.Slug = common.StringPointerValue(apiResp.Slug)

	model.StaffNotes = common.StringPointerValue(apiResp.StaffNotes)

	model.StartDate = common.StringPointerValue(apiResp.StartDate)

	if apiResp.TerminationMetadata != nil {
		valTerminationMetadata, diagsTerminationMetadata := types.MapValueFrom(ctx, types.StringType, apiResp.TerminationMetadata)
		diags.Append(diagsTerminationMetadata...)
		model.TerminationMetadata = valTerminationMetadata
	} else {
		model.TerminationMetadata = types.MapNull(types.StringType)
	}

	model.Type = common.StringPointerValue(apiResp.Type)

	model.TypeName = common.StringPointerValue(apiResp.TypeName)

	model.TypeUuid = common.StringPointerValue(apiResp.TypeUuid)

	model.Url = common.StringPointerValue(apiResp.Url)

	return diags
}
