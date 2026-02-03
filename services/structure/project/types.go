package project

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type StructureProjectCreateRequest struct {
	BackendId       *string `json:"backend_id,omitempty"`
	Customer        *string `json:"customer"`
	Description     *string `json:"description,omitempty"`
	EndDate         *string `json:"end_date,omitempty"`
	GracePeriodDays *int64  `json:"grace_period_days,omitempty"`
	Image           *string `json:"image,omitempty"`
	IsIndustry      *bool   `json:"is_industry,omitempty"`
	Kind            *string `json:"kind,omitempty"`
	Name            *string `json:"name"`
	OecdFos2007Code *string `json:"oecd_fos_2007_code,omitempty"`
	Slug            *string `json:"slug,omitempty"`
	StaffNotes      *string `json:"staff_notes,omitempty"`
	StartDate       *string `json:"start_date,omitempty"`
	Type            *string `json:"type,omitempty"`
}

type StructureProjectUpdateRequest struct {
	BackendId       *string `json:"backend_id,omitempty"`
	Customer        *string `json:"customer,omitempty"`
	Description     *string `json:"description,omitempty"`
	EndDate         *string `json:"end_date,omitempty"`
	GracePeriodDays *int64  `json:"grace_period_days,omitempty"`
	Image           *string `json:"image,omitempty"`
	IsIndustry      *bool   `json:"is_industry,omitempty"`
	Kind            *string `json:"kind,omitempty"`
	Name            *string `json:"name,omitempty"`
	OecdFos2007Code *string `json:"oecd_fos_2007_code,omitempty"`
	Slug            *string `json:"slug,omitempty"`
	StaffNotes      *string `json:"staff_notes,omitempty"`
	StartDate       *string `json:"start_date,omitempty"`
	Type            *string `json:"type,omitempty"`
}

type StructureProjectResponse struct {
	UUID *string `json:"uuid"`

	BackendId                            *string                `json:"backend_id" tfsdk:"backend_id"`
	Created                              *string                `json:"created" tfsdk:"created"`
	Customer                             *string                `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation                 *string                `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerDisplayBillingInfoInProjects *bool                  `json:"customer_display_billing_info_in_projects" tfsdk:"customer_display_billing_info_in_projects"`
	CustomerName                         *string                `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName                   *string                `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerSlug                         *string                `json:"customer_slug" tfsdk:"customer_slug"`
	CustomerUuid                         *string                `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description                          *string                `json:"description" tfsdk:"description"`
	EndDate                              *string                `json:"end_date" tfsdk:"end_date"`
	EndDateRequestedBy                   *string                `json:"end_date_requested_by" tfsdk:"end_date_requested_by"`
	GracePeriodDays                      *int64                 `json:"grace_period_days" tfsdk:"grace_period_days"`
	Image                                *string                `json:"image" tfsdk:"image"`
	IsIndustry                           *bool                  `json:"is_industry" tfsdk:"is_industry"`
	IsRemoved                            *bool                  `json:"is_removed" tfsdk:"is_removed"`
	Kind                                 *string                `json:"kind" tfsdk:"kind"`
	MaxServiceAccounts                   *int64                 `json:"max_service_accounts" tfsdk:"max_service_accounts"`
	Name                                 *string                `json:"name" tfsdk:"name"`
	OecdFos2007Code                      *string                `json:"oecd_fos_2007_code" tfsdk:"oecd_fos_2007_code"`
	OecdFos2007Label                     *string                `json:"oecd_fos_2007_label" tfsdk:"oecd_fos_2007_label"`
	ProjectCredit                        *common.FlexibleNumber `json:"project_credit" tfsdk:"project_credit"`
	ResourcesCount                       *int64                 `json:"resources_count" tfsdk:"resources_count"`
	Slug                                 *string                `json:"slug" tfsdk:"slug"`
	StaffNotes                           *string                `json:"staff_notes" tfsdk:"staff_notes"`
	StartDate                            *string                `json:"start_date" tfsdk:"start_date"`
	Type                                 *string                `json:"type" tfsdk:"type"`
	TypeName                             *string                `json:"type_name" tfsdk:"type_name"`
	TypeUuid                             *string                `json:"type_uuid" tfsdk:"type_uuid"`
	Url                                  *string                `json:"url" tfsdk:"url"`
}

func (r *StructureProjectResponse) GetState() string {
	return "OK"
}

func (r *StructureProjectResponse) GetErrorMessage() string {
	return ""
}
