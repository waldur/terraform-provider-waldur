package customer

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type StructureCustomerCreateRequest struct {
	Abbreviation                 *string  `json:"abbreviation,omitempty"`
	AccessSubnets                *string  `json:"access_subnets,omitempty"`
	AccountingStartDate          *string  `json:"accounting_start_date,omitempty"`
	Address                      *string  `json:"address,omitempty"`
	AgreementNumber              *string  `json:"agreement_number,omitempty"`
	Archived                     *bool    `json:"archived,omitempty"`
	BackendId                    *string  `json:"backend_id,omitempty"`
	BankAccount                  *string  `json:"bank_account,omitempty"`
	BankName                     *string  `json:"bank_name,omitempty"`
	Blocked                      *bool    `json:"blocked,omitempty"`
	ContactDetails               *string  `json:"contact_details,omitempty"`
	Country                      *string  `json:"country,omitempty"`
	DefaultTaxPercent            *string  `json:"default_tax_percent,omitempty"`
	Description                  *string  `json:"description,omitempty"`
	DisplayBillingInfoInProjects *bool    `json:"display_billing_info_in_projects,omitempty"`
	Domain                       *string  `json:"domain,omitempty"`
	Email                        *string  `json:"email,omitempty"`
	GracePeriodDays              *int64   `json:"grace_period_days,omitempty"`
	Homepage                     *string  `json:"homepage,omitempty"`
	Image                        *string  `json:"image,omitempty"`
	Latitude                     *float64 `json:"latitude,omitempty"`
	Longitude                    *float64 `json:"longitude,omitempty"`
	MaxServiceAccounts           *int64   `json:"max_service_accounts,omitempty"`
	Name                         *string  `json:"name"`
	NativeName                   *string  `json:"native_name,omitempty"`
	NotificationEmails           *string  `json:"notification_emails,omitempty"`
	PhoneNumber                  *string  `json:"phone_number,omitempty"`
	Postal                       *string  `json:"postal,omitempty"`
	ProjectMetadataChecklist     *string  `json:"project_metadata_checklist,omitempty"`
	RegistrationCode             *string  `json:"registration_code,omitempty"`
	Slug                         *string  `json:"slug,omitempty"`
	SponsorNumber                *int64   `json:"sponsor_number,omitempty"`
	VatCode                      *string  `json:"vat_code,omitempty"`
}

type StructureCustomerUpdateRequest struct {
	Abbreviation                 *string  `json:"abbreviation,omitempty"`
	AccessSubnets                *string  `json:"access_subnets,omitempty"`
	AccountingStartDate          *string  `json:"accounting_start_date,omitempty"`
	Address                      *string  `json:"address,omitempty"`
	AgreementNumber              *string  `json:"agreement_number,omitempty"`
	Archived                     *bool    `json:"archived,omitempty"`
	BackendId                    *string  `json:"backend_id,omitempty"`
	BankAccount                  *string  `json:"bank_account,omitempty"`
	BankName                     *string  `json:"bank_name,omitempty"`
	Blocked                      *bool    `json:"blocked,omitempty"`
	ContactDetails               *string  `json:"contact_details,omitempty"`
	Country                      *string  `json:"country,omitempty"`
	DefaultTaxPercent            *string  `json:"default_tax_percent,omitempty"`
	Description                  *string  `json:"description,omitempty"`
	DisplayBillingInfoInProjects *bool    `json:"display_billing_info_in_projects,omitempty"`
	Domain                       *string  `json:"domain,omitempty"`
	Email                        *string  `json:"email,omitempty"`
	GracePeriodDays              *int64   `json:"grace_period_days,omitempty"`
	Homepage                     *string  `json:"homepage,omitempty"`
	Image                        *string  `json:"image,omitempty"`
	Latitude                     *float64 `json:"latitude,omitempty"`
	Longitude                    *float64 `json:"longitude,omitempty"`
	MaxServiceAccounts           *int64   `json:"max_service_accounts,omitempty"`
	Name                         *string  `json:"name,omitempty"`
	NativeName                   *string  `json:"native_name,omitempty"`
	NotificationEmails           *string  `json:"notification_emails,omitempty"`
	PhoneNumber                  *string  `json:"phone_number,omitempty"`
	Postal                       *string  `json:"postal,omitempty"`
	ProjectMetadataChecklist     *string  `json:"project_metadata_checklist,omitempty"`
	RegistrationCode             *string  `json:"registration_code,omitempty"`
	Slug                         *string  `json:"slug,omitempty"`
	SponsorNumber                *int64   `json:"sponsor_number,omitempty"`
	VatCode                      *string  `json:"vat_code,omitempty"`
}

type StructureCustomerResponse struct {
	UUID *string `json:"uuid"`

	Abbreviation                 *string                                       `json:"abbreviation" tfsdk:"abbreviation"`
	AccessSubnets                *string                                       `json:"access_subnets" tfsdk:"access_subnets"`
	AccountingStartDate          *string                                       `json:"accounting_start_date" tfsdk:"accounting_start_date"`
	Address                      *string                                       `json:"address" tfsdk:"address"`
	AgreementNumber              *string                                       `json:"agreement_number" tfsdk:"agreement_number"`
	Archived                     *bool                                         `json:"archived" tfsdk:"archived"`
	BackendId                    *string                                       `json:"backend_id" tfsdk:"backend_id"`
	BankAccount                  *string                                       `json:"bank_account" tfsdk:"bank_account"`
	BankName                     *string                                       `json:"bank_name" tfsdk:"bank_name"`
	Blocked                      *bool                                         `json:"blocked" tfsdk:"blocked"`
	CallManagingOrganizationUuid *string                                       `json:"call_managing_organization_uuid" tfsdk:"call_managing_organization_uuid"`
	ContactDetails               *string                                       `json:"contact_details" tfsdk:"contact_details"`
	Country                      *string                                       `json:"country" tfsdk:"country"`
	CountryName                  *string                                       `json:"country_name" tfsdk:"country_name"`
	Created                      *string                                       `json:"created" tfsdk:"created"`
	CustomerCredit               *common.FlexibleNumber                        `json:"customer_credit" tfsdk:"customer_credit"`
	CustomerUnallocatedCredit    *common.FlexibleNumber                        `json:"customer_unallocated_credit" tfsdk:"customer_unallocated_credit"`
	DefaultTaxPercent            *string                                       `json:"default_tax_percent" tfsdk:"default_tax_percent"`
	Description                  *string                                       `json:"description" tfsdk:"description"`
	DisplayBillingInfoInProjects *bool                                         `json:"display_billing_info_in_projects" tfsdk:"display_billing_info_in_projects"`
	DisplayName                  *string                                       `json:"display_name" tfsdk:"display_name"`
	Domain                       *string                                       `json:"domain" tfsdk:"domain"`
	Email                        *string                                       `json:"email" tfsdk:"email"`
	GracePeriodDays              *int64                                        `json:"grace_period_days" tfsdk:"grace_period_days"`
	Homepage                     *string                                       `json:"homepage" tfsdk:"homepage"`
	Image                        *string                                       `json:"image" tfsdk:"image"`
	IsServiceProvider            *bool                                         `json:"is_service_provider" tfsdk:"is_service_provider"`
	Latitude                     *common.FlexibleNumber                        `json:"latitude" tfsdk:"latitude"`
	Longitude                    *common.FlexibleNumber                        `json:"longitude" tfsdk:"longitude"`
	MaxServiceAccounts           *int64                                        `json:"max_service_accounts" tfsdk:"max_service_accounts"`
	Name                         *string                                       `json:"name" tfsdk:"name"`
	NativeName                   *string                                       `json:"native_name" tfsdk:"native_name"`
	NotificationEmails           *string                                       `json:"notification_emails" tfsdk:"notification_emails"`
	OrganizationGroups           []StructureCustomerOrganizationGroupsResponse `json:"organization_groups" tfsdk:"organization_groups"`
	PaymentProfiles              []StructureCustomerPaymentProfilesResponse    `json:"payment_profiles" tfsdk:"payment_profiles"`
	PhoneNumber                  *string                                       `json:"phone_number" tfsdk:"phone_number"`
	Postal                       *string                                       `json:"postal" tfsdk:"postal"`
	ProjectMetadataChecklist     *string                                       `json:"project_metadata_checklist" tfsdk:"project_metadata_checklist"`
	ProjectsCount                *int64                                        `json:"projects_count" tfsdk:"projects_count"`
	RegistrationCode             *string                                       `json:"registration_code" tfsdk:"registration_code"`
	ServiceProvider              *string                                       `json:"service_provider" tfsdk:"service_provider"`
	ServiceProviderUuid          *string                                       `json:"service_provider_uuid" tfsdk:"service_provider_uuid"`
	Slug                         *string                                       `json:"slug" tfsdk:"slug"`
	SponsorNumber                *int64                                        `json:"sponsor_number" tfsdk:"sponsor_number"`
	Url                          *string                                       `json:"url" tfsdk:"url"`
	UsersCount                   *int64                                        `json:"users_count" tfsdk:"users_count"`
	VatCode                      *string                                       `json:"vat_code" tfsdk:"vat_code"`
}

type StructureCustomerOrganizationGroupsResponse struct {
	CustomersCount *int64  `json:"customers_count" tfsdk:"customers_count"`
	Name           *string `json:"name" tfsdk:"name"`
	Parent         *string `json:"parent" tfsdk:"parent"`
	ParentName     *string `json:"parent_name" tfsdk:"parent_name"`
	ParentUuid     *string `json:"parent_uuid" tfsdk:"parent_uuid"`
	Url            *string `json:"url" tfsdk:"url"`
	Uuid           *string `json:"uuid" tfsdk:"uuid"`
}

type StructureCustomerPaymentProfilesResponse struct {
	Attributes         *StructureCustomerPaymentProfilesAttributesResponse `json:"attributes" tfsdk:"attributes"`
	IsActive           *bool                                               `json:"is_active" tfsdk:"is_active"`
	Name               *string                                             `json:"name" tfsdk:"name"`
	Organization       *string                                             `json:"organization" tfsdk:"organization"`
	OrganizationUuid   *string                                             `json:"organization_uuid" tfsdk:"organization_uuid"`
	PaymentType        *string                                             `json:"payment_type" tfsdk:"payment_type"`
	PaymentTypeDisplay *string                                             `json:"payment_type_display" tfsdk:"payment_type_display"`
	Url                *string                                             `json:"url" tfsdk:"url"`
	Uuid               *string                                             `json:"uuid" tfsdk:"uuid"`
}

type StructureCustomerPaymentProfilesAttributesResponse struct {
	AgreementNumber *string `json:"agreement_number" tfsdk:"agreement_number"`
	ContractSum     *int64  `json:"contract_sum" tfsdk:"contract_sum"`
	EndDate         *string `json:"end_date" tfsdk:"end_date"`
}

func (r *StructureCustomerResponse) GetState() string {
	return "OK"
}

func (r *StructureCustomerResponse) GetErrorMessage() string {
	return ""
}
