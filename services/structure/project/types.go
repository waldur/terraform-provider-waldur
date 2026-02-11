package project

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type StructureProjectCreateRequest struct {
	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Customer *string `json:"customer" tfsdk:"customer"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	EndDate *string `json:"end_date,omitempty" tfsdk:"end_date"`

	GracePeriodDays *int64 `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`

	Image *string `json:"image,omitempty" tfsdk:"image"`

	IsIndustry *bool `json:"is_industry,omitempty" tfsdk:"is_industry"`

	Kind *string `json:"kind,omitempty" tfsdk:"kind"`

	Name *string `json:"name" tfsdk:"name"`

	OecdFos2007Code *string `json:"oecd_fos_2007_code,omitempty" tfsdk:"oecd_fos_2007_code"`

	Slug *string `json:"slug,omitempty" tfsdk:"slug"`

	StaffNotes *string `json:"staff_notes,omitempty" tfsdk:"staff_notes"`

	StartDate *string `json:"start_date,omitempty" tfsdk:"start_date"`

	Type *string `json:"type,omitempty" tfsdk:"type"`
}

type StructureProjectUpdateRequest struct {
	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Customer *string `json:"customer,omitempty" tfsdk:"customer"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	EndDate *string `json:"end_date,omitempty" tfsdk:"end_date"`

	GracePeriodDays *int64 `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`

	Image *string `json:"image,omitempty" tfsdk:"image"`

	IsIndustry *bool `json:"is_industry,omitempty" tfsdk:"is_industry"`

	Kind *string `json:"kind,omitempty" tfsdk:"kind"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	OecdFos2007Code *string `json:"oecd_fos_2007_code,omitempty" tfsdk:"oecd_fos_2007_code"`

	Slug *string `json:"slug,omitempty" tfsdk:"slug"`

	StaffNotes *string `json:"staff_notes,omitempty" tfsdk:"staff_notes"`

	StartDate *string `json:"start_date,omitempty" tfsdk:"start_date"`

	Type *string `json:"type,omitempty" tfsdk:"type"`
}

type StructureProjectResponse struct {
	UUID *string `json:"uuid"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	BillingPriceEstimate *StructureProjectBillingPriceEstimateResponse `json:"billing_price_estimate,omitempty" tfsdk:"billing_price_estimate"`

	Customer *string `json:"customer" tfsdk:"customer"`

	CustomerDisplayBillingInfoInProjects *bool `json:"customer_display_billing_info_in_projects,omitempty" tfsdk:"customer_display_billing_info_in_projects"`

	CustomerSlug *string `json:"customer_slug,omitempty" tfsdk:"customer_slug"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	EndDate *string `json:"end_date,omitempty" tfsdk:"end_date"`

	EndDateRequestedBy *string `json:"end_date_requested_by,omitempty" tfsdk:"end_date_requested_by"`

	GracePeriodDays *int64 `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`

	Image *string `json:"image,omitempty" tfsdk:"image"`

	IsIndustry *bool `json:"is_industry,omitempty" tfsdk:"is_industry"`

	IsRemoved *bool `json:"is_removed,omitempty" tfsdk:"is_removed"`

	Kind *string `json:"kind,omitempty" tfsdk:"kind"`

	MaxServiceAccounts *int64 `json:"max_service_accounts,omitempty" tfsdk:"max_service_accounts"`

	Name *string `json:"name" tfsdk:"name"`

	OecdFos2007Code *string `json:"oecd_fos_2007_code,omitempty" tfsdk:"oecd_fos_2007_code"`

	OecdFos2007Label *string `json:"oecd_fos_2007_label,omitempty" tfsdk:"oecd_fos_2007_label"`

	ProjectCredit common.FlexibleNumber `json:"project_credit,omitempty" tfsdk:"project_credit"`

	ResourcesCount *int64 `json:"resources_count,omitempty" tfsdk:"resources_count"`

	Slug *string `json:"slug,omitempty" tfsdk:"slug"`

	StaffNotes *string `json:"staff_notes,omitempty" tfsdk:"staff_notes"`

	StartDate *string `json:"start_date,omitempty" tfsdk:"start_date"`

	Type *string `json:"type,omitempty" tfsdk:"type"`

	TypeName *string `json:"type_name,omitempty" tfsdk:"type_name"`

	TypeUuid *string `json:"type_uuid,omitempty" tfsdk:"type_uuid"`

	Url *string `json:"url,omitempty" tfsdk:"url"`
}

type StructureProjectBillingPriceEstimateResponse struct {
	Current common.FlexibleNumber `json:"current,omitempty" tfsdk:"current"`

	Tax common.FlexibleNumber `json:"tax,omitempty" tfsdk:"tax"`

	TaxCurrent common.FlexibleNumber `json:"tax_current,omitempty" tfsdk:"tax_current"`

	Total common.FlexibleNumber `json:"total,omitempty" tfsdk:"total"`
}

func (r *StructureProjectResponse) GetState() string {
	return "OK"
}

func (r *StructureProjectResponse) GetErrorMessage() string {
	return ""
}
