package customer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StructureCustomerFiltersModel struct {
	Abbreviation          types.String `tfsdk:"abbreviation"`
	AgreementNumber       types.String `tfsdk:"agreement_number"`
	Archived              types.Bool   `tfsdk:"archived"`
	BackendId             types.String `tfsdk:"backend_id"`
	ContactDetails        types.String `tfsdk:"contact_details"`
	Name                  types.String `tfsdk:"name"`
	NameExact             types.String `tfsdk:"name_exact"`
	NativeName            types.String `tfsdk:"native_name"`
	OrganizationGroupName types.String `tfsdk:"organization_group_name"`
	OwnedByCurrentUser    types.Bool   `tfsdk:"owned_by_current_user"`
	Query                 types.String `tfsdk:"query"`
	RegistrationCode      types.String `tfsdk:"registration_code"`
}

type StructureCustomerModel struct {
	UUID                         types.String      `tfsdk:"id"`
	Abbreviation                 types.String      `tfsdk:"abbreviation"`
	AccessSubnets                types.String      `tfsdk:"access_subnets"`
	AccountingStartDate          timetypes.RFC3339 `tfsdk:"accounting_start_date"`
	Address                      types.String      `tfsdk:"address"`
	AgreementNumber              types.String      `tfsdk:"agreement_number"`
	Archived                     types.Bool        `tfsdk:"archived"`
	BackendId                    types.String      `tfsdk:"backend_id"`
	BankAccount                  types.String      `tfsdk:"bank_account"`
	BankName                     types.String      `tfsdk:"bank_name"`
	Blocked                      types.Bool        `tfsdk:"blocked"`
	CallManagingOrganizationUuid types.String      `tfsdk:"call_managing_organization_uuid"`
	ContactDetails               types.String      `tfsdk:"contact_details"`
	Country                      types.String      `tfsdk:"country"`
	CountryName                  types.String      `tfsdk:"country_name"`
	Created                      timetypes.RFC3339 `tfsdk:"created"`
	CustomerCredit               types.Float64     `tfsdk:"customer_credit"`
	CustomerUnallocatedCredit    types.Float64     `tfsdk:"customer_unallocated_credit"`
	DefaultTaxPercent            types.String      `tfsdk:"default_tax_percent"`
	Description                  types.String      `tfsdk:"description"`
	DisplayBillingInfoInProjects types.Bool        `tfsdk:"display_billing_info_in_projects"`
	DisplayName                  types.String      `tfsdk:"display_name"`
	Domain                       types.String      `tfsdk:"domain"`
	Email                        types.String      `tfsdk:"email"`
	GracePeriodDays              types.Int64       `tfsdk:"grace_period_days"`
	Homepage                     types.String      `tfsdk:"homepage"`
	Image                        types.String      `tfsdk:"image"`
	IsServiceProvider            types.Bool        `tfsdk:"is_service_provider"`
	Latitude                     types.Float64     `tfsdk:"latitude"`
	Longitude                    types.Float64     `tfsdk:"longitude"`
	MaxServiceAccounts           types.Int64       `tfsdk:"max_service_accounts"`
	Name                         types.String      `tfsdk:"name"`
	NativeName                   types.String      `tfsdk:"native_name"`
	NotificationEmails           types.String      `tfsdk:"notification_emails"`
	OrganizationGroups           types.List        `tfsdk:"organization_groups"`
	PaymentProfiles              types.List        `tfsdk:"payment_profiles"`
	PhoneNumber                  types.String      `tfsdk:"phone_number"`
	Postal                       types.String      `tfsdk:"postal"`
	ProjectMetadataChecklist     types.String      `tfsdk:"project_metadata_checklist"`
	ProjectsCount                types.Int64       `tfsdk:"projects_count"`
	RegistrationCode             types.String      `tfsdk:"registration_code"`
	ServiceProvider              types.String      `tfsdk:"service_provider"`
	ServiceProviderUuid          types.String      `tfsdk:"service_provider_uuid"`
	Slug                         types.String      `tfsdk:"slug"`
	SponsorNumber                types.Int64       `tfsdk:"sponsor_number"`
	Url                          types.String      `tfsdk:"url"`
	UsersCount                   types.Int64       `tfsdk:"users_count"`
	VatCode                      types.String      `tfsdk:"vat_code"`
}

// CopyFrom maps the API response to the model fields.
func (model *StructureCustomerModel) CopyFrom(ctx context.Context, apiResp StructureCustomerResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.Abbreviation = types.StringPointerValue(apiResp.Abbreviation)
	model.AccessSubnets = types.StringPointerValue(apiResp.AccessSubnets)
	valAccountingStartDate, diagsAccountingStartDate := timetypes.NewRFC3339PointerValue(apiResp.AccountingStartDate)
	diags.Append(diagsAccountingStartDate...)
	model.AccountingStartDate = valAccountingStartDate
	model.Address = types.StringPointerValue(apiResp.Address)
	model.AgreementNumber = types.StringPointerValue(apiResp.AgreementNumber)
	model.Archived = types.BoolPointerValue(apiResp.Archived)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.BankAccount = types.StringPointerValue(apiResp.BankAccount)
	model.BankName = types.StringPointerValue(apiResp.BankName)
	model.Blocked = types.BoolPointerValue(apiResp.Blocked)
	model.CallManagingOrganizationUuid = types.StringPointerValue(apiResp.CallManagingOrganizationUuid)
	model.ContactDetails = types.StringPointerValue(apiResp.ContactDetails)
	model.Country = types.StringPointerValue(apiResp.Country)
	model.CountryName = types.StringPointerValue(apiResp.CountryName)
	valCreated, diagsCreated := timetypes.NewRFC3339PointerValue(apiResp.Created)
	diags.Append(diagsCreated...)
	model.Created = valCreated
	model.CustomerCredit = types.Float64PointerValue(apiResp.CustomerCredit)
	model.CustomerUnallocatedCredit = types.Float64PointerValue(apiResp.CustomerUnallocatedCredit)
	model.DefaultTaxPercent = types.StringPointerValue(apiResp.DefaultTaxPercent)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.DisplayBillingInfoInProjects = types.BoolPointerValue(apiResp.DisplayBillingInfoInProjects)
	model.DisplayName = types.StringPointerValue(apiResp.DisplayName)
	model.Domain = types.StringPointerValue(apiResp.Domain)
	model.Email = types.StringPointerValue(apiResp.Email)
	model.GracePeriodDays = types.Int64PointerValue(apiResp.GracePeriodDays)
	model.Homepage = types.StringPointerValue(apiResp.Homepage)
	model.Image = types.StringPointerValue(apiResp.Image)
	model.IsServiceProvider = types.BoolPointerValue(apiResp.IsServiceProvider)
	model.Latitude = types.Float64PointerValue(apiResp.Latitude)
	model.Longitude = types.Float64PointerValue(apiResp.Longitude)
	model.MaxServiceAccounts = types.Int64PointerValue(apiResp.MaxServiceAccounts)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.NativeName = types.StringPointerValue(apiResp.NativeName)
	model.NotificationEmails = types.StringPointerValue(apiResp.NotificationEmails)

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

	{
		listValPaymentProfiles, listDiagsPaymentProfiles := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"attributes": types.ObjectType{AttrTypes: map[string]attr.Type{
				"agreement_number": types.StringType,
				"contract_sum":     types.Int64Type,
				"end_date":         types.StringType,
			}},
			"is_active":            types.BoolType,
			"name":                 types.StringType,
			"organization":         types.StringType,
			"organization_uuid":    types.StringType,
			"payment_type":         types.StringType,
			"payment_type_display": types.StringType,
			"url":                  types.StringType,
		}}, apiResp.PaymentProfiles)
		diags.Append(listDiagsPaymentProfiles...)
		model.PaymentProfiles = listValPaymentProfiles
	}
	model.PhoneNumber = types.StringPointerValue(apiResp.PhoneNumber)
	model.Postal = types.StringPointerValue(apiResp.Postal)
	model.ProjectMetadataChecklist = types.StringPointerValue(apiResp.ProjectMetadataChecklist)
	model.ProjectsCount = types.Int64PointerValue(apiResp.ProjectsCount)
	model.RegistrationCode = types.StringPointerValue(apiResp.RegistrationCode)
	model.ServiceProvider = types.StringPointerValue(apiResp.ServiceProvider)
	model.ServiceProviderUuid = types.StringPointerValue(apiResp.ServiceProviderUuid)
	model.Slug = types.StringPointerValue(apiResp.Slug)
	model.SponsorNumber = types.Int64PointerValue(apiResp.SponsorNumber)
	model.Url = types.StringPointerValue(apiResp.Url)
	model.UsersCount = types.Int64PointerValue(apiResp.UsersCount)
	model.VatCode = types.StringPointerValue(apiResp.VatCode)

	return diags
}
