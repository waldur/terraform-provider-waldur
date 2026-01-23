package customer

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &StructureCustomerDataSource{}

func NewStructureCustomerDataSource() datasource.DataSource {
	return &StructureCustomerDataSource{}
}

// StructureCustomerDataSource defines the data source implementation.
type StructureCustomerDataSource struct {
	client *Client
}

// StructureCustomerFiltersModel contains the filter parameters for querying.
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
	OrganizationGroupUuid types.String `tfsdk:"organization_group_uuid"`
	OwnedByCurrentUser    types.Bool   `tfsdk:"owned_by_current_user"`
	Query                 types.String `tfsdk:"query"`
	RegistrationCode      types.String `tfsdk:"registration_code"`
}

// StructureCustomerDataSourceModel describes the data source data model.
type StructureCustomerDataSourceModel struct {
	UUID                         types.String                   `tfsdk:"id"`
	Filters                      *StructureCustomerFiltersModel `tfsdk:"filters"`
	Abbreviation                 types.String                   `tfsdk:"abbreviation"`
	AccessSubnets                types.String                   `tfsdk:"access_subnets"`
	AccountingStartDate          types.String                   `tfsdk:"accounting_start_date"`
	Address                      types.String                   `tfsdk:"address"`
	AgreementNumber              types.String                   `tfsdk:"agreement_number"`
	Archived                     types.Bool                     `tfsdk:"archived"`
	BackendId                    types.String                   `tfsdk:"backend_id"`
	BankAccount                  types.String                   `tfsdk:"bank_account"`
	BankName                     types.String                   `tfsdk:"bank_name"`
	Blocked                      types.Bool                     `tfsdk:"blocked"`
	CallManagingOrganizationUuid types.String                   `tfsdk:"call_managing_organization_uuid"`
	ContactDetails               types.String                   `tfsdk:"contact_details"`
	Country                      types.String                   `tfsdk:"country"`
	CountryName                  types.String                   `tfsdk:"country_name"`
	Created                      types.String                   `tfsdk:"created"`
	CustomerCredit               types.Float64                  `tfsdk:"customer_credit"`
	CustomerUnallocatedCredit    types.Float64                  `tfsdk:"customer_unallocated_credit"`
	DefaultTaxPercent            types.String                   `tfsdk:"default_tax_percent"`
	Description                  types.String                   `tfsdk:"description"`
	DisplayBillingInfoInProjects types.Bool                     `tfsdk:"display_billing_info_in_projects"`
	DisplayName                  types.String                   `tfsdk:"display_name"`
	Domain                       types.String                   `tfsdk:"domain"`
	Email                        types.String                   `tfsdk:"email"`
	GracePeriodDays              types.Int64                    `tfsdk:"grace_period_days"`
	Homepage                     types.String                   `tfsdk:"homepage"`
	Image                        types.String                   `tfsdk:"image"`
	IsServiceProvider            types.Bool                     `tfsdk:"is_service_provider"`
	Latitude                     types.Float64                  `tfsdk:"latitude"`
	Longitude                    types.Float64                  `tfsdk:"longitude"`
	MaxServiceAccounts           types.Int64                    `tfsdk:"max_service_accounts"`
	Name                         types.String                   `tfsdk:"name"`
	NativeName                   types.String                   `tfsdk:"native_name"`
	NotificationEmails           types.String                   `tfsdk:"notification_emails"`
	OrganizationGroups           types.List                     `tfsdk:"organization_groups"`
	PaymentProfiles              types.List                     `tfsdk:"payment_profiles"`
	PhoneNumber                  types.String                   `tfsdk:"phone_number"`
	Postal                       types.String                   `tfsdk:"postal"`
	ProjectMetadataChecklist     types.String                   `tfsdk:"project_metadata_checklist"`
	ProjectsCount                types.Int64                    `tfsdk:"projects_count"`
	RegistrationCode             types.String                   `tfsdk:"registration_code"`
	ServiceProvider              types.String                   `tfsdk:"service_provider"`
	ServiceProviderUuid          types.String                   `tfsdk:"service_provider_uuid"`
	Slug                         types.String                   `tfsdk:"slug"`
	SponsorNumber                types.Int64                    `tfsdk:"sponsor_number"`
	Url                          types.String                   `tfsdk:"url"`
	UsersCount                   types.Int64                    `tfsdk:"users_count"`
	VatCode                      types.String                   `tfsdk:"vat_code"`
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
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Filter parameters for querying Structure Customer",
				Attributes: map[string]schema.Attribute{
					"abbreviation": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Abbreviation",
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
					"organization_group_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Organization group UUID",
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
				},
			},
			"abbreviation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Abbreviation",
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
			"agreement_number": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Agreement number",
			},
			"archived": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Archived",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Organization identifier in another application.",
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
			"contact_details": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Contact details",
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
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the resource",
			},
			"native_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the native",
			},
			"notification_emails": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Comma-separated list of notification email addresses",
			},
			"organization_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"customers_count": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Number of customers in this organization group",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the resource",
						},
						"parent": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Parent",
						},
						"parent_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name of the parent organization group",
						},
						"parent_uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "UUID of the parent organization group",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Url",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Organization groups this customer belongs to",
			},
			"payment_profiles": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"attributes": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"agreement_number": schema.StringAttribute{
									Optional:            true,
									MarkdownDescription: "Agreement number",
								},
								"contract_sum": schema.Int64Attribute{
									Optional:            true,
									MarkdownDescription: "Contract sum",
								},
								"end_date": schema.StringAttribute{
									Optional:            true,
									MarkdownDescription: "End date",
								},
							},
							Optional:            true,
							MarkdownDescription: "Attributes",
						},
						"is_active": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "Is active",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the resource",
						},
						"organization": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Organization",
						},
						"organization_uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "UUID of the organization",
						},
						"payment_type": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Payment type",
						},
						"payment_type_display": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Payment type display",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Url",
						},
					},
				},
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
			"registration_code": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Registration code",
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

	rawClient, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = NewClient(rawClient)
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
		apiResp, err := d.client.GetStructureCustomer(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Structure Customer",
				"An error occurred while reading the Structure Customer by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, *apiResp, &data)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup structure_customer.",
			)
			return
		}

		results, err := d.client.ListStructureCustomer(ctx, filters)
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

func (d *StructureCustomerDataSource) mapResponseToModel(ctx context.Context, apiResp StructureCustomerResponse, model *StructureCustomerDataSourceModel) diag.Diagnostics {
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
