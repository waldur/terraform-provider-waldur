package customer

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type StructureCustomerCreateRequest struct {
	Abbreviation *string `json:"abbreviation,omitempty" tfsdk:"abbreviation"`

	AccessSubnets *string `json:"access_subnets,omitempty" tfsdk:"access_subnets"`

	AccountingStartDate *string `json:"accounting_start_date,omitempty" tfsdk:"accounting_start_date"`

	Address *string `json:"address,omitempty" tfsdk:"address"`

	AgreementNumber *string `json:"agreement_number,omitempty" tfsdk:"agreement_number"`

	ApartmentNr *string `json:"apartment_nr,omitempty" tfsdk:"apartment_nr"`

	Archived *bool `json:"archived,omitempty" tfsdk:"archived"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	BankAccount *string `json:"bank_account,omitempty" tfsdk:"bank_account"`

	BankName *string `json:"bank_name,omitempty" tfsdk:"bank_name"`

	Blocked *bool `json:"blocked,omitempty" tfsdk:"blocked"`

	City *string `json:"city,omitempty" tfsdk:"city"`

	ContactDetails *string `json:"contact_details,omitempty" tfsdk:"contact_details"`

	Country *string `json:"country,omitempty" tfsdk:"country"`

	DefaultTaxPercent *string `json:"default_tax_percent,omitempty" tfsdk:"default_tax_percent"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	DisplayBillingInfoInProjects *bool `json:"display_billing_info_in_projects,omitempty" tfsdk:"display_billing_info_in_projects"`

	Domain *string `json:"domain,omitempty" tfsdk:"domain"`

	Email *string `json:"email,omitempty" tfsdk:"email"`

	GracePeriodDays *int64 `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`

	Homepage *string `json:"homepage,omitempty" tfsdk:"homepage"`

	HouseNr *string `json:"house_nr,omitempty" tfsdk:"house_nr"`

	Household *string `json:"household,omitempty" tfsdk:"household"`

	Image *string `json:"image,omitempty" tfsdk:"image"`

	MaxServiceAccounts *int64 `json:"max_service_accounts,omitempty" tfsdk:"max_service_accounts"`

	Name *string `json:"name" tfsdk:"name"`

	NativeName *string `json:"native_name,omitempty" tfsdk:"native_name"`

	NotificationEmails *string `json:"notification_emails,omitempty" tfsdk:"notification_emails"`

	Parish *string `json:"parish,omitempty" tfsdk:"parish"`

	PhoneNumber *string `json:"phone_number,omitempty" tfsdk:"phone_number"`

	Postal *string `json:"postal,omitempty" tfsdk:"postal"`

	ProjectMetadataChecklist *string `json:"project_metadata_checklist,omitempty" tfsdk:"project_metadata_checklist"`

	ProjectSlugTemplate *string `json:"project_slug_template,omitempty" tfsdk:"project_slug_template"`

	RegistrationCode *string `json:"registration_code,omitempty" tfsdk:"registration_code"`

	Slug *string `json:"slug,omitempty" tfsdk:"slug"`

	SponsorNumber *int64 `json:"sponsor_number,omitempty" tfsdk:"sponsor_number"`

	Street *string `json:"street,omitempty" tfsdk:"street"`

	VatCode *string `json:"vat_code,omitempty" tfsdk:"vat_code"`
}

type StructureCustomerUpdateRequest struct {
	Abbreviation *string `json:"abbreviation,omitempty" tfsdk:"abbreviation"`

	AccessSubnets *string `json:"access_subnets,omitempty" tfsdk:"access_subnets"`

	AccountingStartDate *string `json:"accounting_start_date,omitempty" tfsdk:"accounting_start_date"`

	Address *string `json:"address,omitempty" tfsdk:"address"`

	AgreementNumber *string `json:"agreement_number,omitempty" tfsdk:"agreement_number"`

	ApartmentNr *string `json:"apartment_nr,omitempty" tfsdk:"apartment_nr"`

	Archived *bool `json:"archived,omitempty" tfsdk:"archived"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	BankAccount *string `json:"bank_account,omitempty" tfsdk:"bank_account"`

	BankName *string `json:"bank_name,omitempty" tfsdk:"bank_name"`

	Blocked *bool `json:"blocked,omitempty" tfsdk:"blocked"`

	City *string `json:"city,omitempty" tfsdk:"city"`

	ContactDetails *string `json:"contact_details,omitempty" tfsdk:"contact_details"`

	Country *string `json:"country,omitempty" tfsdk:"country"`

	DefaultTaxPercent *string `json:"default_tax_percent,omitempty" tfsdk:"default_tax_percent"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	DisplayBillingInfoInProjects *bool `json:"display_billing_info_in_projects,omitempty" tfsdk:"display_billing_info_in_projects"`

	Domain *string `json:"domain,omitempty" tfsdk:"domain"`

	Email *string `json:"email,omitempty" tfsdk:"email"`

	GracePeriodDays *int64 `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`

	Homepage *string `json:"homepage,omitempty" tfsdk:"homepage"`

	HouseNr *string `json:"house_nr,omitempty" tfsdk:"house_nr"`

	Household *string `json:"household,omitempty" tfsdk:"household"`

	Image *string `json:"image,omitempty" tfsdk:"image"`

	MaxServiceAccounts *int64 `json:"max_service_accounts,omitempty" tfsdk:"max_service_accounts"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	NativeName *string `json:"native_name,omitempty" tfsdk:"native_name"`

	NotificationEmails *string `json:"notification_emails,omitempty" tfsdk:"notification_emails"`

	Parish *string `json:"parish,omitempty" tfsdk:"parish"`

	PhoneNumber *string `json:"phone_number,omitempty" tfsdk:"phone_number"`

	Postal *string `json:"postal,omitempty" tfsdk:"postal"`

	ProjectMetadataChecklist *string `json:"project_metadata_checklist,omitempty" tfsdk:"project_metadata_checklist"`

	ProjectSlugTemplate *string `json:"project_slug_template,omitempty" tfsdk:"project_slug_template"`

	RegistrationCode *string `json:"registration_code,omitempty" tfsdk:"registration_code"`

	Slug *string `json:"slug,omitempty" tfsdk:"slug"`

	SponsorNumber *int64 `json:"sponsor_number,omitempty" tfsdk:"sponsor_number"`

	Street *string `json:"street,omitempty" tfsdk:"street"`

	VatCode *string `json:"vat_code,omitempty" tfsdk:"vat_code"`
}

type StructureCustomerResponse struct {
	UUID *string `json:"uuid"`

	Abbreviation *string `json:"abbreviation,omitempty" tfsdk:"abbreviation"`

	AccessSubnets *string `json:"access_subnets,omitempty" tfsdk:"access_subnets"`

	AccountingStartDate *string `json:"accounting_start_date,omitempty" tfsdk:"accounting_start_date"`

	Address *string `json:"address,omitempty" tfsdk:"address"`

	AgreementNumber *string `json:"agreement_number,omitempty" tfsdk:"agreement_number"`

	ApartmentNr *string `json:"apartment_nr,omitempty" tfsdk:"apartment_nr"`

	Archived *bool `json:"archived,omitempty" tfsdk:"archived"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	BankAccount *string `json:"bank_account,omitempty" tfsdk:"bank_account"`

	BankName *string `json:"bank_name,omitempty" tfsdk:"bank_name"`

	BillingPriceEstimate *StructureCustomerBillingPriceEstimateResponse `json:"billing_price_estimate,omitempty" tfsdk:"billing_price_estimate"`

	Blocked *bool `json:"blocked,omitempty" tfsdk:"blocked"`

	CallManagingOrganizationUuid *string `json:"call_managing_organization_uuid,omitempty" tfsdk:"call_managing_organization_uuid"`

	City *string `json:"city,omitempty" tfsdk:"city"`

	ContactDetails *string `json:"contact_details,omitempty" tfsdk:"contact_details"`

	Country *string `json:"country,omitempty" tfsdk:"country"`

	CountryName *string `json:"country_name,omitempty" tfsdk:"country_name"`

	CustomerCredit *string `json:"customer_credit,omitempty" tfsdk:"customer_credit"`

	CustomerUnallocatedCredit *string `json:"customer_unallocated_credit,omitempty" tfsdk:"customer_unallocated_credit"`

	DefaultAffiliations *[]common.AffiliatedOrganization `json:"default_affiliations,omitempty" tfsdk:"default_affiliations"`

	DefaultTaxPercent *string `json:"default_tax_percent,omitempty" tfsdk:"default_tax_percent"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	DisplayBillingInfoInProjects *bool `json:"display_billing_info_in_projects,omitempty" tfsdk:"display_billing_info_in_projects"`

	DisplayName *string `json:"display_name,omitempty" tfsdk:"display_name"`

	Domain *string `json:"domain,omitempty" tfsdk:"domain"`

	Email *string `json:"email,omitempty" tfsdk:"email"`

	GracePeriodDays *int64 `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`

	HasAffiliateLinks *bool `json:"has_affiliate_links,omitempty" tfsdk:"has_affiliate_links"`

	Homepage *string `json:"homepage,omitempty" tfsdk:"homepage"`

	HouseNr *string `json:"house_nr,omitempty" tfsdk:"house_nr"`

	Household *string `json:"household,omitempty" tfsdk:"household"`

	Image *string `json:"image,omitempty" tfsdk:"image"`

	IsServiceProvider *bool `json:"is_service_provider,omitempty" tfsdk:"is_service_provider"`

	MaxServiceAccounts *int64 `json:"max_service_accounts,omitempty" tfsdk:"max_service_accounts"`

	Name *string `json:"name" tfsdk:"name"`

	NativeName *string `json:"native_name,omitempty" tfsdk:"native_name"`

	NotificationEmails *string `json:"notification_emails,omitempty" tfsdk:"notification_emails"`

	OrganizationGroups *[]common.OrganizationGroup `json:"organization_groups,omitempty" tfsdk:"organization_groups"`

	Parish *string `json:"parish,omitempty" tfsdk:"parish"`

	PaymentProfiles *[]common.PaymentProfile `json:"payment_profiles,omitempty" tfsdk:"payment_profiles"`

	PhoneNumber *string `json:"phone_number,omitempty" tfsdk:"phone_number"`

	Postal *string `json:"postal,omitempty" tfsdk:"postal"`

	ProjectMetadataChecklist *string `json:"project_metadata_checklist,omitempty" tfsdk:"project_metadata_checklist"`

	ProjectSlugTemplate *string `json:"project_slug_template,omitempty" tfsdk:"project_slug_template"`

	ProjectsCount *int64 `json:"projects_count,omitempty" tfsdk:"projects_count"`

	RegistrationCode *string `json:"registration_code,omitempty" tfsdk:"registration_code"`

	ServiceProvider *string `json:"service_provider,omitempty" tfsdk:"service_provider"`

	ServiceProviderUuid *string `json:"service_provider_uuid,omitempty" tfsdk:"service_provider_uuid"`

	Slug *string `json:"slug,omitempty" tfsdk:"slug"`

	SponsorNumber *int64 `json:"sponsor_number,omitempty" tfsdk:"sponsor_number"`

	Street *string `json:"street,omitempty" tfsdk:"street"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	UsersCount *int64 `json:"users_count,omitempty" tfsdk:"users_count"`

	VatCode *string `json:"vat_code,omitempty" tfsdk:"vat_code"`

	State *string `json:"state,omitempty" tfsdk:"state"`
}

type StructureCustomerBillingPriceEstimateResponse struct {
	Current common.FlexibleNumber `json:"current,omitempty" tfsdk:"current"`

	Tax common.FlexibleNumber `json:"tax,omitempty" tfsdk:"tax"`

	TaxCurrent common.FlexibleNumber `json:"tax_current,omitempty" tfsdk:"tax_current"`

	Total common.FlexibleNumber `json:"total,omitempty" tfsdk:"total"`
}

type StructureCustomerDefaultAffiliationsResponse struct {
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

type StructureCustomerOrganizationGroupsResponse struct {
	CustomersCount *int64 `json:"customers_count,omitempty" tfsdk:"customers_count"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Parent *string `json:"parent,omitempty" tfsdk:"parent"`

	ParentName *string `json:"parent_name,omitempty" tfsdk:"parent_name"`

	ParentUuid *string `json:"parent_uuid,omitempty" tfsdk:"parent_uuid"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type StructureCustomerPaymentProfilesResponse struct {
	Attributes *common.PaymentProfileAttributes `json:"attributes,omitempty" tfsdk:"attributes"`

	IsActive *bool `json:"is_active,omitempty" tfsdk:"is_active"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Organization *string `json:"organization,omitempty" tfsdk:"organization"`

	OrganizationUuid *string `json:"organization_uuid,omitempty" tfsdk:"organization_uuid"`

	PaymentType *string `json:"payment_type,omitempty" tfsdk:"payment_type"`

	PaymentTypeDisplay *string `json:"payment_type_display,omitempty" tfsdk:"payment_type_display"`

	Url *string `json:"url,omitempty" tfsdk:"url"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

type StructureCustomerPaymentProfilesAttributesResponse struct {
	AgreementNumber *string `json:"agreement_number,omitempty" tfsdk:"agreement_number"`

	ContractSum *string `json:"contract_sum,omitempty" tfsdk:"contract_sum"`

	EndDate *string `json:"end_date,omitempty" tfsdk:"end_date"`
}

func (r *StructureCustomerResponse) GetState() string {
	if r.State != nil {
		return *r.State
	}
	return "OK"
}

func (r *StructureCustomerResponse) GetErrorMessage() string {
	return ""
}
