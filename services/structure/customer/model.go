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

func AffiliatedOrganizationType() types.ObjectType {
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
		"contract_sum":     types.StringType,
		"end_date":         types.StringType,
	}}
}

type StructureCustomerFiltersModel struct {
	Abbreviation                          types.String `tfsdk:"abbreviation"`
	AccountingIsRunning                   types.Bool   `tfsdk:"accounting_is_running"`
	AgreementNumber                       types.String `tfsdk:"agreement_number"`
	Archived                              types.Bool   `tfsdk:"archived"`
	BackendId                             types.String `tfsdk:"backend_id"`
	ContactDetails                        types.String `tfsdk:"contact_details"`
	CurrentUserHasProjectCreatePermission types.Bool   `tfsdk:"current_user_has_project_create_permission"`
	HasResources                          types.String `tfsdk:"has_resources"`
	IsCallManagingOrganization            types.Bool   `tfsdk:"is_call_managing_organization"`
	IsServiceProvider                     types.Bool   `tfsdk:"is_service_provider"`
	Name                                  types.String `tfsdk:"name"`
	NameExact                             types.String `tfsdk:"name_exact"`
	NativeName                            types.String `tfsdk:"native_name"`
	OrganizationGroupName                 types.String `tfsdk:"organization_group_name"`
	OwnedByCurrentUser                    types.Bool   `tfsdk:"owned_by_current_user"`
	Query                                 types.String `tfsdk:"query"`
	RegistrationCode                      types.String `tfsdk:"registration_code"`
	ServiceProviderUuid                   types.String `tfsdk:"service_provider_uuid"`
	Slug                                  types.String `tfsdk:"slug"`
	UserUuid                              types.String `tfsdk:"user_uuid"`
}

func (m *StructureCustomerFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Structure Customer",
		Attributes: map[string]schema.Attribute{
			"abbreviation": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Abbreviation",
			},
			"accounting_is_running": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Filter by whether accounting is running.",
			},
			"agreement_number": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Agreement number",
			},
			"archived": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Archived",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "ID of the backend",
			},
			"contact_details": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Contact details",
			},
			"current_user_has_project_create_permission": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Return a list of customers where current user has project create permission.",
			},
			"has_resources": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Filter by customers with resources.",
			},
			"is_call_managing_organization": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Filter by customers that are call managing organizations.",
			},
			"is_service_provider": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Filter by customers that are service providers.",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name (exact)",
			},
			"native_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Native name",
			},
			"organization_group_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Organization group name",
			},
			"owned_by_current_user": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Return a list of customers where current user is owner.",
			},
			"query": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Filter by name, native name, abbreviation, domain, UUID, registration code or agreement number",
			},
			"registration_code": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Registration code",
			},
			"service_provider_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Filter by service provider UUID.",
			},
			"slug": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Slug",
			},
			"user_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Filter by user UUID.",
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
	ApartmentNr                  types.String      `tfsdk:"apartment_nr"`
	Archived                     types.Bool        `tfsdk:"archived"`
	BackendId                    types.String      `tfsdk:"backend_id"`
	BankAccount                  types.String      `tfsdk:"bank_account"`
	BankName                     types.String      `tfsdk:"bank_name"`
	BillingPriceEstimate         types.Object      `tfsdk:"billing_price_estimate"`
	Blocked                      types.Bool        `tfsdk:"blocked"`
	CallManagingOrganizationUuid types.String      `tfsdk:"call_managing_organization_uuid"`
	City                         types.String      `tfsdk:"city"`
	ContactDetails               types.String      `tfsdk:"contact_details"`
	Country                      types.String      `tfsdk:"country"`
	CountryName                  types.String      `tfsdk:"country_name"`
	CustomerCredit               types.String      `tfsdk:"customer_credit"`
	CustomerUnallocatedCredit    types.String      `tfsdk:"customer_unallocated_credit"`
	DefaultAffiliations          types.List        `tfsdk:"default_affiliations"`
	DefaultTaxPercent            types.String      `tfsdk:"default_tax_percent"`
	Description                  types.String      `tfsdk:"description"`
	DisplayBillingInfoInProjects types.Bool        `tfsdk:"display_billing_info_in_projects"`
	DisplayName                  types.String      `tfsdk:"display_name"`
	Domain                       types.String      `tfsdk:"domain"`
	Email                        types.String      `tfsdk:"email"`
	GracePeriodDays              types.Int64       `tfsdk:"grace_period_days"`
	HasActiveHelpdesk            types.Bool        `tfsdk:"has_active_helpdesk"`
	HasAffiliateLinks            types.Bool        `tfsdk:"has_affiliate_links"`
	Homepage                     types.String      `tfsdk:"homepage"`
	HouseNr                      types.String      `tfsdk:"house_nr"`
	Household                    types.String      `tfsdk:"household"`
	Image                        types.String      `tfsdk:"image"`
	IsServiceProvider            types.Bool        `tfsdk:"is_service_provider"`
	MaxServiceAccounts           types.Int64       `tfsdk:"max_service_accounts"`
	Name                         types.String      `tfsdk:"name"`
	NativeName                   types.String      `tfsdk:"native_name"`
	NotificationEmails           types.String      `tfsdk:"notification_emails"`
	OrganizationGroups           types.List        `tfsdk:"organization_groups"`
	Parish                       types.String      `tfsdk:"parish"`
	PaymentProfiles              types.List        `tfsdk:"payment_profiles"`
	PhoneNumber                  types.String      `tfsdk:"phone_number"`
	Postal                       types.String      `tfsdk:"postal"`
	ProjectMetadataChecklist     types.String      `tfsdk:"project_metadata_checklist"`
	ProjectSlugTemplate          types.String      `tfsdk:"project_slug_template"`
	ProjectsCount                types.Int64       `tfsdk:"projects_count"`
	RegistrationCode             types.String      `tfsdk:"registration_code"`
	ServiceProvider              types.String      `tfsdk:"service_provider"`
	ServiceProviderUuid          types.String      `tfsdk:"service_provider_uuid"`
	Slug                         types.String      `tfsdk:"slug"`
	SponsorNumber                types.Int64       `tfsdk:"sponsor_number"`
	Street                       types.String      `tfsdk:"street"`
	Url                          types.String      `tfsdk:"url"`
	UsersCount                   types.Int64       `tfsdk:"users_count"`
	VatCode                      types.String      `tfsdk:"vat_code"`
	State                        types.String      `tfsdk:"state"`
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

	model.ApartmentNr = common.StringPointerValue(apiResp.ApartmentNr)

	model.Archived = types.BoolPointerValue(apiResp.Archived)

	model.BackendId = common.StringPointerValue(apiResp.BackendId)

	model.BankAccount = common.StringPointerValue(apiResp.BankAccount)

	model.BankName = common.StringPointerValue(apiResp.BankName)

	if apiResp.BillingPriceEstimate != nil {
		valBillingPriceEstimate, diagsBillingPriceEstimate := types.ObjectValueFrom(ctx, BillingPriceEstimateType().AttrTypes, *apiResp.BillingPriceEstimate)
		diags.Append(diagsBillingPriceEstimate...)
		model.BillingPriceEstimate = valBillingPriceEstimate
	} else {
		model.BillingPriceEstimate = types.ObjectNull(BillingPriceEstimateType().AttrTypes)
	}

	model.Blocked = types.BoolPointerValue(apiResp.Blocked)

	model.CallManagingOrganizationUuid = common.StringPointerValue(apiResp.CallManagingOrganizationUuid)

	model.City = common.StringPointerValue(apiResp.City)

	model.ContactDetails = common.StringPointerValue(apiResp.ContactDetails)

	model.Country = common.StringPointerValue(apiResp.Country)

	model.CountryName = common.StringPointerValue(apiResp.CountryName)

	model.CustomerCredit = common.StringPointerValue(apiResp.CustomerCredit)

	model.CustomerUnallocatedCredit = common.StringPointerValue(apiResp.CustomerUnallocatedCredit)

	if apiResp.DefaultAffiliations != nil {
		valDefaultAffiliations, diagsDefaultAffiliations := types.ListValueFrom(ctx, AffiliatedOrganizationType(), apiResp.DefaultAffiliations)
		diags.Append(diagsDefaultAffiliations...)
		model.DefaultAffiliations = valDefaultAffiliations
	} else {
		model.DefaultAffiliations = types.ListNull(AffiliatedOrganizationType())
	}

	model.DefaultTaxPercent = common.StringPointerValue(apiResp.DefaultTaxPercent)

	model.Description = common.StringPointerValue(apiResp.Description)

	model.DisplayBillingInfoInProjects = types.BoolPointerValue(apiResp.DisplayBillingInfoInProjects)

	model.DisplayName = common.StringPointerValue(apiResp.DisplayName)

	model.Domain = common.StringPointerValue(apiResp.Domain)

	model.Email = common.StringPointerValue(apiResp.Email)

	model.GracePeriodDays = types.Int64PointerValue(apiResp.GracePeriodDays)

	model.HasActiveHelpdesk = types.BoolPointerValue(apiResp.HasActiveHelpdesk)

	model.HasAffiliateLinks = types.BoolPointerValue(apiResp.HasAffiliateLinks)

	model.Homepage = common.StringPointerValue(apiResp.Homepage)

	model.HouseNr = common.StringPointerValue(apiResp.HouseNr)

	model.Household = common.StringPointerValue(apiResp.Household)

	model.Image = common.StringPointerValue(apiResp.Image)

	model.IsServiceProvider = types.BoolPointerValue(apiResp.IsServiceProvider)

	model.MaxServiceAccounts = types.Int64PointerValue(apiResp.MaxServiceAccounts)

	model.Name = common.StringPointerValue(apiResp.Name)

	model.NativeName = common.StringPointerValue(apiResp.NativeName)

	model.NotificationEmails = common.StringPointerValue(apiResp.NotificationEmails)

	if apiResp.OrganizationGroups != nil {
		valOrganizationGroups, diagsOrganizationGroups := types.ListValueFrom(ctx, OrganizationGroupType(), apiResp.OrganizationGroups)
		diags.Append(diagsOrganizationGroups...)
		model.OrganizationGroups = valOrganizationGroups
	} else {
		model.OrganizationGroups = types.ListNull(OrganizationGroupType())
	}

	model.Parish = common.StringPointerValue(apiResp.Parish)

	if apiResp.PaymentProfiles != nil {
		valPaymentProfiles, diagsPaymentProfiles := types.ListValueFrom(ctx, PaymentProfileType(), apiResp.PaymentProfiles)
		diags.Append(diagsPaymentProfiles...)
		model.PaymentProfiles = valPaymentProfiles
	} else {
		model.PaymentProfiles = types.ListNull(PaymentProfileType())
	}

	model.PhoneNumber = common.StringPointerValue(apiResp.PhoneNumber)

	model.Postal = common.StringPointerValue(apiResp.Postal)

	model.ProjectMetadataChecklist = common.StringPointerValue(apiResp.ProjectMetadataChecklist)

	model.ProjectSlugTemplate = common.StringPointerValue(apiResp.ProjectSlugTemplate)

	model.ProjectsCount = types.Int64PointerValue(apiResp.ProjectsCount)

	model.RegistrationCode = common.StringPointerValue(apiResp.RegistrationCode)

	model.ServiceProvider = common.StringPointerValue(apiResp.ServiceProvider)

	model.ServiceProviderUuid = common.StringPointerValue(apiResp.ServiceProviderUuid)

	model.Slug = common.StringPointerValue(apiResp.Slug)

	model.SponsorNumber = types.Int64PointerValue(apiResp.SponsorNumber)

	model.Street = common.StringPointerValue(apiResp.Street)

	model.Url = common.StringPointerValue(apiResp.Url)

	model.UsersCount = types.Int64PointerValue(apiResp.UsersCount)

	model.VatCode = common.StringPointerValue(apiResp.VatCode)

	model.State = common.StringPointerValue(apiResp.State)

	return diags
}
