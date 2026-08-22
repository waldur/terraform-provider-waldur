package project

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type StructureProjectCreateRequest struct {
	AffiliationUuid *string `json:"affiliation_uuid,omitempty" tfsdk:"affiliation_uuid"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Customer *string `json:"customer" tfsdk:"customer"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	DisplayCreditReports *bool `json:"display_credit_reports,omitempty" tfsdk:"display_credit_reports"`

	EndDate *string `json:"end_date,omitempty" tfsdk:"end_date"`

	GracePeriodDays *int64 `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`

	Image *string `json:"image,omitempty" tfsdk:"image"`

	IsIndustry *bool `json:"is_industry,omitempty" tfsdk:"is_industry"`

	Kind *string `json:"kind,omitempty" tfsdk:"kind"`

	Name *string `json:"name" tfsdk:"name"`

	OecdFos2007Code *string `json:"oecd_fos_2007_code,omitempty" tfsdk:"oecd_fos_2007_code"`

	ScienceSubDomain *string `json:"science_sub_domain,omitempty" tfsdk:"science_sub_domain"`

	Slug *string `json:"slug,omitempty" tfsdk:"slug"`

	StaffNotes *string `json:"staff_notes,omitempty" tfsdk:"staff_notes"`

	StartDate *string `json:"start_date,omitempty" tfsdk:"start_date"`

	Type *string `json:"type,omitempty" tfsdk:"type"`
}

type StructureProjectUpdateRequest struct {
	AffiliationUuid *string `json:"affiliation_uuid,omitempty" tfsdk:"affiliation_uuid"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Customer *string `json:"customer,omitempty" tfsdk:"customer"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	DisplayCreditReports *bool `json:"display_credit_reports,omitempty" tfsdk:"display_credit_reports"`

	EndDate *string `json:"end_date,omitempty" tfsdk:"end_date"`

	GracePeriodDays *int64 `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`

	Image *string `json:"image,omitempty" tfsdk:"image"`

	IsIndustry *bool `json:"is_industry,omitempty" tfsdk:"is_industry"`

	Kind *string `json:"kind,omitempty" tfsdk:"kind"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	OecdFos2007Code *string `json:"oecd_fos_2007_code,omitempty" tfsdk:"oecd_fos_2007_code"`

	ScienceSubDomain *string `json:"science_sub_domain,omitempty" tfsdk:"science_sub_domain"`

	Slug *string `json:"slug,omitempty" tfsdk:"slug"`

	StaffNotes *string `json:"staff_notes,omitempty" tfsdk:"staff_notes"`

	StartDate *string `json:"start_date,omitempty" tfsdk:"start_date"`

	Type *string `json:"type,omitempty" tfsdk:"type"`
}

type StructureProjectResponse struct {
	UUID *string `json:"uuid"`

	Affiliation *StructureProjectAffiliationResponse `json:"affiliation,omitempty" tfsdk:"affiliation"`

	AffiliationCode *string `json:"affiliation_code,omitempty" tfsdk:"affiliation_code"`

	AffiliationName *string `json:"affiliation_name,omitempty" tfsdk:"affiliation_name"`

	AffiliationUuid *string `json:"affiliation_uuid,omitempty" tfsdk:"affiliation_uuid"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	BillingPriceEstimate *StructureProjectBillingPriceEstimateResponse `json:"billing_price_estimate,omitempty" tfsdk:"billing_price_estimate"`

	Customer *string `json:"customer" tfsdk:"customer"`

	CustomerDisplayBillingInfoInProjects *bool `json:"customer_display_billing_info_in_projects,omitempty" tfsdk:"customer_display_billing_info_in_projects"`

	CustomerGracePeriodDays *int64 `json:"customer_grace_period_days,omitempty" tfsdk:"customer_grace_period_days"`

	CustomerSlug *string `json:"customer_slug,omitempty" tfsdk:"customer_slug"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	DisplayCreditReports *bool `json:"display_credit_reports,omitempty" tfsdk:"display_credit_reports"`

	EffectiveEndDate *string `json:"effective_end_date,omitempty" tfsdk:"effective_end_date"`

	EndDate *string `json:"end_date,omitempty" tfsdk:"end_date"`

	EndDateRequestedBy *string `json:"end_date_requested_by,omitempty" tfsdk:"end_date_requested_by"`

	EndDateUpdatedAt *string `json:"end_date_updated_at,omitempty" tfsdk:"end_date_updated_at"`

	GracePeriodDays *int64 `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`

	Image *string `json:"image,omitempty" tfsdk:"image"`

	IsInGracePeriod *bool `json:"is_in_grace_period,omitempty" tfsdk:"is_in_grace_period"`

	IsIndustry *bool `json:"is_industry,omitempty" tfsdk:"is_industry"`

	IsRemoved *bool `json:"is_removed,omitempty" tfsdk:"is_removed"`

	Kind *string `json:"kind,omitempty" tfsdk:"kind"`

	MarketplaceResourceCount map[string]int64 `json:"marketplace_resource_count,omitempty" tfsdk:"marketplace_resource_count"`

	MaxServiceAccounts *int64 `json:"max_service_accounts,omitempty" tfsdk:"max_service_accounts"`

	Name *string `json:"name" tfsdk:"name"`

	OecdFos2007Code *string `json:"oecd_fos_2007_code,omitempty" tfsdk:"oecd_fos_2007_code"`

	OecdFos2007Label *string `json:"oecd_fos_2007_label,omitempty" tfsdk:"oecd_fos_2007_label"`

	ProjectCredit common.FlexibleNumber `json:"project_credit,omitempty" tfsdk:"project_credit"`

	ProjectMetadata *[]common.ProjectMetadataAnswer `json:"project_metadata,omitempty" tfsdk:"project_metadata"`

	ResourcesCount *int64 `json:"resources_count,omitempty" tfsdk:"resources_count"`

	ScienceDomainCode *string `json:"science_domain_code,omitempty" tfsdk:"science_domain_code"`

	ScienceDomainName *string `json:"science_domain_name,omitempty" tfsdk:"science_domain_name"`

	ScienceDomainUuid *string `json:"science_domain_uuid,omitempty" tfsdk:"science_domain_uuid"`

	ScienceSubDomain *string `json:"science_sub_domain,omitempty" tfsdk:"science_sub_domain"`

	ScienceSubDomainCode *string `json:"science_sub_domain_code,omitempty" tfsdk:"science_sub_domain_code"`

	ScienceSubDomainName *string `json:"science_sub_domain_name,omitempty" tfsdk:"science_sub_domain_name"`

	Slug *string `json:"slug,omitempty" tfsdk:"slug"`

	StaffNotes *string `json:"staff_notes,omitempty" tfsdk:"staff_notes"`

	StartDate *string `json:"start_date,omitempty" tfsdk:"start_date"`

	TerminationMetadata common.JSONStringMap `json:"termination_metadata,omitempty" tfsdk:"termination_metadata"`

	Type *string `json:"type,omitempty" tfsdk:"type"`

	TypeName *string `json:"type_name,omitempty" tfsdk:"type_name"`

	TypeUuid *string `json:"type_uuid,omitempty" tfsdk:"type_uuid"`

	Url *string `json:"url,omitempty" tfsdk:"url"`
}

type StructureProjectAffiliationResponse struct {
	Abbreviation *string `json:"abbreviation,omitempty" tfsdk:"abbreviation"`

	Address *string `json:"address,omitempty" tfsdk:"address"`

	Code *string `json:"code,omitempty" tfsdk:"code"`

	Country *string `json:"country,omitempty" tfsdk:"country"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	Email *string `json:"email,omitempty" tfsdk:"email"`

	Homepage *string `json:"homepage,omitempty" tfsdk:"homepage"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	ProjectsCount *int64 `json:"projects_count,omitempty" tfsdk:"projects_count"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type StructureProjectBillingPriceEstimateResponse struct {
	Current common.FlexibleNumber `json:"current,omitempty" tfsdk:"current"`

	Tax common.FlexibleNumber `json:"tax,omitempty" tfsdk:"tax"`

	TaxCurrent common.FlexibleNumber `json:"tax_current,omitempty" tfsdk:"tax_current"`

	Total common.FlexibleNumber `json:"total,omitempty" tfsdk:"total"`
}

type StructureProjectMarketplaceResourceCountResponse struct {
}

type StructureProjectProjectMetadataResponse struct {
	Question *string `json:"question,omitempty" tfsdk:"question"`

	QuestionType *string `json:"question_type,omitempty" tfsdk:"question_type"`

	QuestionUuid *string `json:"question_uuid,omitempty" tfsdk:"question_uuid"`
}

type StructureProjectTerminationMetadataResponse struct {
}

func (r *StructureProjectResponse) GetState() string {
	return "OK"
}

func (r *StructureProjectResponse) GetErrorMessage() string {
	return ""
}
