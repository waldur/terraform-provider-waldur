package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &StructureCustomerDataSource{}

func NewStructureCustomerDataSource() datasource.DataSource {
	return &StructureCustomerDataSource{}
}

// StructureCustomerDataSource defines the data source implementation.
type StructureCustomerDataSource struct {
	client *client.Client
}

// StructureCustomerApiResponse is the API response model.
type StructureCustomerApiResponse struct {
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
	CustomerCredit               *float64                                      `json:"customer_credit" tfsdk:"customer_credit"`
	CustomerUnallocatedCredit    *float64                                      `json:"customer_unallocated_credit" tfsdk:"customer_unallocated_credit"`
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
	Latitude                     *float64                                      `json:"latitude" tfsdk:"latitude"`
	Longitude                    *float64                                      `json:"longitude" tfsdk:"longitude"`
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
}

type StructureCustomerPaymentProfilesAttributesResponse struct {
	AgreementNumber *string `json:"agreement_number" tfsdk:"agreement_number"`
	ContractSum     *int64  `json:"contract_sum" tfsdk:"contract_sum"`
	EndDate         *string `json:"end_date" tfsdk:"end_date"`
}

// StructureCustomerDataSourceModel describes the data source data model.
type StructureCustomerDataSourceModel struct {
	UUID                         types.String  `tfsdk:"id"`
	Abbreviation                 types.String  `tfsdk:"abbreviation"`
	AgreementNumber              types.String  `tfsdk:"agreement_number"`
	Archived                     types.Bool    `tfsdk:"archived"`
	BackendId                    types.String  `tfsdk:"backend_id"`
	ContactDetails               types.String  `tfsdk:"contact_details"`
	Name                         types.String  `tfsdk:"name"`
	NameExact                    types.String  `tfsdk:"name_exact"`
	NativeName                   types.String  `tfsdk:"native_name"`
	OrganizationGroupName        types.String  `tfsdk:"organization_group_name"`
	OrganizationGroupUuid        types.String  `tfsdk:"organization_group_uuid"`
	OwnedByCurrentUser           types.Bool    `tfsdk:"owned_by_current_user"`
	Query                        types.String  `tfsdk:"query"`
	RegistrationCode             types.String  `tfsdk:"registration_code"`
	AccessSubnets                types.String  `tfsdk:"access_subnets"`
	AccountingStartDate          types.String  `tfsdk:"accounting_start_date"`
	Address                      types.String  `tfsdk:"address"`
	BankAccount                  types.String  `tfsdk:"bank_account"`
	BankName                     types.String  `tfsdk:"bank_name"`
	Blocked                      types.Bool    `tfsdk:"blocked"`
	CallManagingOrganizationUuid types.String  `tfsdk:"call_managing_organization_uuid"`
	Country                      types.String  `tfsdk:"country"`
	CountryName                  types.String  `tfsdk:"country_name"`
	Created                      types.String  `tfsdk:"created"`
	CustomerCredit               types.Float64 `tfsdk:"customer_credit"`
	CustomerUnallocatedCredit    types.Float64 `tfsdk:"customer_unallocated_credit"`
	DefaultTaxPercent            types.String  `tfsdk:"default_tax_percent"`
	Description                  types.String  `tfsdk:"description"`
	DisplayBillingInfoInProjects types.Bool    `tfsdk:"display_billing_info_in_projects"`
	DisplayName                  types.String  `tfsdk:"display_name"`
	Domain                       types.String  `tfsdk:"domain"`
	Email                        types.String  `tfsdk:"email"`
	GracePeriodDays              types.Int64   `tfsdk:"grace_period_days"`
	Homepage                     types.String  `tfsdk:"homepage"`
	Image                        types.String  `tfsdk:"image"`
	IsServiceProvider            types.Bool    `tfsdk:"is_service_provider"`
	Latitude                     types.Float64 `tfsdk:"latitude"`
	Longitude                    types.Float64 `tfsdk:"longitude"`
	MaxServiceAccounts           types.Int64   `tfsdk:"max_service_accounts"`
	NotificationEmails           types.String  `tfsdk:"notification_emails"`
	OrganizationGroups           types.List    `tfsdk:"organization_groups"`
	PaymentProfiles              types.List    `tfsdk:"payment_profiles"`
	PhoneNumber                  types.String  `tfsdk:"phone_number"`
	Postal                       types.String  `tfsdk:"postal"`
	ProjectMetadataChecklist     types.String  `tfsdk:"project_metadata_checklist"`
	ProjectsCount                types.Int64   `tfsdk:"projects_count"`
	ServiceProvider              types.String  `tfsdk:"service_provider"`
	ServiceProviderUuid          types.String  `tfsdk:"service_provider_uuid"`
	Slug                         types.String  `tfsdk:"slug"`
	SponsorNumber                types.Int64   `tfsdk:"sponsor_number"`
	Url                          types.String  `tfsdk:"url"`
	UsersCount                   types.Int64   `tfsdk:"users_count"`
	VatCode                      types.String  `tfsdk:"vat_code"`
}

func (d *StructureCustomerDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_structure_customer"
}

func (d *StructureCustomerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Structure Customer data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"abbreviation": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Abbreviation",
			},
			"agreement_number": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Agreement number",
			},
			"archived": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Archived",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "ID of the backend",
			},
			"contact_details": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Contact details",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Name",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Name (exact)",
			},
			"native_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Native name",
			},
			"organization_group_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Organization group name",
			},
			"organization_group_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Organization group UUID",
			},
			"owned_by_current_user": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Return a list of customers where current user is owner.",
			},
			"query": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Filter by name, native name, abbreviation, domain, UUID, registration code or agreement number",
			},
			"registration_code": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Registration code",
			},
			"access_subnets": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Enter a comma separated list of IPv4 or IPv6 CIDR addresses from where connection to self-service is allowed.",
			},
			"accounting_start_date": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Accounting start date",
			},
			"address": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Address",
			},
			"bank_account": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Bank account",
			},
			"bank_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the bank",
			},
			"blocked": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Blocked",
			},
			"call_managing_organization_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the call managing organization",
			},
			"country": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Country code (ISO 3166-1 alpha-2)",
			},
			"country_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Human-readable country name",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Created",
			},
			"customer_credit": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "Customer credit",
			},
			"customer_unallocated_credit": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "Customer unallocated credit",
			},
			"default_tax_percent": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Default tax percent",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Description of the resource",
			},
			"display_billing_info_in_projects": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Display billing info in projects",
			},
			"display_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Display name of the organization (includes native name if available)",
			},
			"domain": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Domain",
			},
			"email": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Email",
			},
			"grace_period_days": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of extra days after project end date before resources are terminated",
			},
			"homepage": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Homepage",
			},
			"image": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Image",
			},
			"is_service_provider": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Is service provider",
			},
			"latitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "Latitude",
			},
			"longitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "Longitude",
			},
			"max_service_accounts": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Maximum number of service accounts allowed",
			},
			"notification_emails": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Comma-separated list of notification email addresses",
			},
			"organization_groups": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"customers_count": types.Int64Type,
					"name":            types.StringType,
					"parent":          types.StringType,
					"parent_name":     types.StringType,
					"parent_uuid":     types.StringType,
					"url":             types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: "Organization groups this customer belongs to",
			},
			"payment_profiles": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
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
				}}},
				Computed:            true,
				MarkdownDescription: "Payment profiles",
			},
			"phone_number": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Phone number",
			},
			"postal": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Postal",
			},
			"project_metadata_checklist": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Checklist to be used for project metadata validation in this organization",
			},
			"projects_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of projects in this organization",
			},
			"service_provider": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Service provider",
			},
			"service_provider_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the service provider",
			},
			"slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
			},
			"sponsor_number": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "External ID of the sponsor covering the costs",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Url",
			},
			"users_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of users with access to this organization",
			},
			"vat_code": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "VAT number",
			},
		},
	}
}

func (d *StructureCustomerDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
}

func (d *StructureCustomerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data StructureCustomerDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var apiResp StructureCustomerApiResponse

		err := d.client.GetByUUID(ctx, "/api/customers/{uuid}/", data.UUID.ValueString(), &apiResp)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Structure Customer",
				"An error occurred while reading the Structure Customer by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, apiResp, &data)...)

	} else {
		// Filter by provided parameters
		var results []StructureCustomerApiResponse

		filters := map[string]string{}
		if !data.Abbreviation.IsNull() {
			filters["abbreviation"] = data.Abbreviation.ValueString()
		}
		if !data.AgreementNumber.IsNull() {
			filters["agreement_number"] = data.AgreementNumber.ValueString()
		}
		if !data.Archived.IsNull() {
			filters["archived"] = fmt.Sprintf("%t", data.Archived.ValueBool())
		}
		if !data.BackendId.IsNull() {
			filters["backend_id"] = data.BackendId.ValueString()
		}
		if !data.ContactDetails.IsNull() {
			filters["contact_details"] = data.ContactDetails.ValueString()
		}
		if !data.Name.IsNull() {
			filters["name"] = data.Name.ValueString()
		}
		if !data.NameExact.IsNull() {
			filters["name_exact"] = data.NameExact.ValueString()
		}
		if !data.NativeName.IsNull() {
			filters["native_name"] = data.NativeName.ValueString()
		}
		if !data.OrganizationGroupName.IsNull() {
			filters["organization_group_name"] = data.OrganizationGroupName.ValueString()
		}
		if !data.OrganizationGroupUuid.IsNull() {
			filters["organization_group_uuid"] = data.OrganizationGroupUuid.ValueString()
		}
		if !data.OwnedByCurrentUser.IsNull() {
			filters["owned_by_current_user"] = fmt.Sprintf("%t", data.OwnedByCurrentUser.ValueBool())
		}
		if !data.Query.IsNull() {
			filters["query"] = data.Query.ValueString()
		}
		if !data.RegistrationCode.IsNull() {
			filters["registration_code"] = data.RegistrationCode.ValueString()
		}

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup structure_customer.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/customers/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Structure Customer",
				"An error occurred while filtering Structure Customer: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Structure Customer Not Found",
				"No Structure Customer found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Structure Customers Found",
				fmt.Sprintf("Found %d Structure Customers with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *StructureCustomerDataSource) mapResponseToModel(ctx context.Context, apiResp StructureCustomerApiResponse, model *StructureCustomerDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.Abbreviation = types.StringPointerValue(apiResp.Abbreviation)
	model.AccessSubnets = types.StringPointerValue(apiResp.AccessSubnets)
	model.AccountingStartDate = types.StringPointerValue(apiResp.AccountingStartDate)
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
	model.Created = types.StringPointerValue(apiResp.Created)
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
