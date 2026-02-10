package customer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

func BillingPriceEstimateType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"current":     types.Float64Type,
		"tax":         types.Float64Type,
		"tax_current": types.Float64Type,
		"total":       types.Float64Type,
	}}
}
func OrganizationGroupType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"customers_count": types.Int64Type,
		"name":            types.StringType,
		"parent":          types.StringType,
		"parent_name":     types.StringType,
		"parent_uuid":     types.StringType,
		"url":             types.StringType,
		"uuid":            types.StringType,
	}}
}
func PaymentProfileType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"attributes":           PaymentProfileAttributesType(),
		"is_active":            types.BoolType,
		"name":                 types.StringType,
		"organization":         types.StringType,
		"organization_uuid":    types.StringType,
		"payment_type":         types.StringType,
		"payment_type_display": types.StringType,
		"url":                  types.StringType,
		"uuid":                 types.StringType,
	}}
}
func PaymentProfileAttributesType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"agreement_number": types.StringType,
		"contract_sum":     types.Int64Type,
		"end_date":         types.StringType,
	}}
}

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

func (m *StructureCustomerFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Structure Customer",
		Attributes: map[string]schema.Attribute{
			"abbreviation": schema.StringAttribute{
				Optional: true,
			},
			"agreement_number": schema.StringAttribute{
				Optional: true,
			},
			"archived": schema.BoolAttribute{
				Optional: true,
			},
			"backend_id": schema.StringAttribute{
				Optional: true,
			},
			"contact_details": schema.StringAttribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"name_exact": schema.StringAttribute{
				Optional: true,
			},
			"native_name": schema.StringAttribute{
				Optional: true,
			},
			"organization_group_name": schema.StringAttribute{
				Optional: true,
			},
			"owned_by_current_user": schema.BoolAttribute{
				Optional: true,
			},
			"query": schema.StringAttribute{
				Optional: true,
			},
			"registration_code": schema.StringAttribute{
				Optional: true,
			},
		},
	}
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
	BillingPriceEstimate         types.Object      `tfsdk:"billing_price_estimate"`
	Blocked                      types.Bool        `tfsdk:"blocked"`
	CallManagingOrganizationUuid types.String      `tfsdk:"call_managing_organization_uuid"`
	ContactDetails               types.String      `tfsdk:"contact_details"`
	Country                      types.String      `tfsdk:"country"`
	CountryName                  types.String      `tfsdk:"country_name"`
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
	model.Abbreviation = common.StringPointerValue(apiResp.Abbreviation)
	model.AccessSubnets = common.StringPointerValue(apiResp.AccessSubnets)
	valAccountingStartDate, diagsAccountingStartDate := timetypes.NewRFC3339PointerValue(apiResp.AccountingStartDate)
	diags.Append(diagsAccountingStartDate...)
	model.AccountingStartDate = valAccountingStartDate
	model.Address = common.StringPointerValue(apiResp.Address)
	model.AgreementNumber = common.StringPointerValue(apiResp.AgreementNumber)
	model.Archived = types.BoolPointerValue(apiResp.Archived)
	model.BackendId = common.StringPointerValue(apiResp.BackendId)
	model.BankAccount = common.StringPointerValue(apiResp.BankAccount)
	model.BankName = common.StringPointerValue(apiResp.BankName)
	if apiResp.BillingPriceEstimate != nil {
		objValBillingPriceEstimate, objDiagsBillingPriceEstimate := types.ObjectValueFrom(ctx, BillingPriceEstimateType().AttrTypes, *apiResp.BillingPriceEstimate)
		diags.Append(objDiagsBillingPriceEstimate...)
		model.BillingPriceEstimate = objValBillingPriceEstimate
	} else {
		model.BillingPriceEstimate = types.ObjectNull(BillingPriceEstimateType().AttrTypes)
	}
	model.Blocked = types.BoolPointerValue(apiResp.Blocked)
	model.CallManagingOrganizationUuid = common.StringPointerValue(apiResp.CallManagingOrganizationUuid)
	model.ContactDetails = common.StringPointerValue(apiResp.ContactDetails)
	model.Country = common.StringPointerValue(apiResp.Country)
	model.CountryName = common.StringPointerValue(apiResp.CountryName)
	model.CustomerCredit = types.Float64PointerValue(apiResp.CustomerCredit.Float64Ptr())
	model.CustomerUnallocatedCredit = types.Float64PointerValue(apiResp.CustomerUnallocatedCredit.Float64Ptr())
	model.DefaultTaxPercent = common.StringPointerValue(apiResp.DefaultTaxPercent)
	model.Description = common.StringPointerValue(apiResp.Description)
	model.DisplayBillingInfoInProjects = types.BoolPointerValue(apiResp.DisplayBillingInfoInProjects)
	model.DisplayName = common.StringPointerValue(apiResp.DisplayName)
	model.Domain = common.StringPointerValue(apiResp.Domain)
	model.Email = common.StringPointerValue(apiResp.Email)
	model.GracePeriodDays = types.Int64PointerValue(apiResp.GracePeriodDays)
	model.Homepage = common.StringPointerValue(apiResp.Homepage)
	model.Image = common.StringPointerValue(apiResp.Image)
	model.IsServiceProvider = types.BoolPointerValue(apiResp.IsServiceProvider)
	model.Latitude = types.Float64PointerValue(apiResp.Latitude.Float64Ptr())
	model.Longitude = types.Float64PointerValue(apiResp.Longitude.Float64Ptr())
	model.MaxServiceAccounts = types.Int64PointerValue(apiResp.MaxServiceAccounts)
	model.Name = common.StringPointerValue(apiResp.Name)
	model.NativeName = common.StringPointerValue(apiResp.NativeName)
	model.NotificationEmails = common.StringPointerValue(apiResp.NotificationEmails)

	if apiResp.OrganizationGroups != nil && len(*apiResp.OrganizationGroups) > 0 {
		listValOrganizationGroups, listDiagsOrganizationGroups := types.ListValueFrom(ctx, OrganizationGroupType(), apiResp.OrganizationGroups)
		diags.Append(listDiagsOrganizationGroups...)
		model.OrganizationGroups = listValOrganizationGroups
	} else {
		model.OrganizationGroups = types.ListNull(OrganizationGroupType())
	}

	if apiResp.PaymentProfiles != nil && len(*apiResp.PaymentProfiles) > 0 {
		listValPaymentProfiles, listDiagsPaymentProfiles := types.ListValueFrom(ctx, PaymentProfileType(), apiResp.PaymentProfiles)
		diags.Append(listDiagsPaymentProfiles...)
		model.PaymentProfiles = listValPaymentProfiles
	} else {
		model.PaymentProfiles = types.ListNull(PaymentProfileType())
	}
	model.PhoneNumber = common.StringPointerValue(apiResp.PhoneNumber)
	model.Postal = common.StringPointerValue(apiResp.Postal)
	model.ProjectMetadataChecklist = common.StringPointerValue(apiResp.ProjectMetadataChecklist)
	model.ProjectsCount = types.Int64PointerValue(apiResp.ProjectsCount)
	model.RegistrationCode = common.StringPointerValue(apiResp.RegistrationCode)
	model.ServiceProvider = common.StringPointerValue(apiResp.ServiceProvider)
	model.ServiceProviderUuid = common.StringPointerValue(apiResp.ServiceProviderUuid)
	model.Slug = common.StringPointerValue(apiResp.Slug)
	model.SponsorNumber = types.Int64PointerValue(apiResp.SponsorNumber)
	model.Url = common.StringPointerValue(apiResp.Url)
	model.UsersCount = types.Int64PointerValue(apiResp.UsersCount)
	model.VatCode = common.StringPointerValue(apiResp.VatCode)

	return diags
}
