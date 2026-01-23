package customer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &StructureCustomerResource{}
var _ resource.ResourceWithImportState = &StructureCustomerResource{}

func NewStructureCustomerResource() resource.Resource {
	return &StructureCustomerResource{}
}

// StructureCustomerResource defines the resource implementation.
type StructureCustomerResource struct {
	client *Client
}

// StructureCustomerResourceModel describes the resource data model.
type StructureCustomerResourceModel struct {
	UUID                         types.String   `tfsdk:"id"`
	Abbreviation                 types.String   `tfsdk:"abbreviation"`
	AccessSubnets                types.String   `tfsdk:"access_subnets"`
	AccountingStartDate          types.String   `tfsdk:"accounting_start_date"`
	Address                      types.String   `tfsdk:"address"`
	AgreementNumber              types.String   `tfsdk:"agreement_number"`
	Archived                     types.Bool     `tfsdk:"archived"`
	BackendId                    types.String   `tfsdk:"backend_id"`
	BankAccount                  types.String   `tfsdk:"bank_account"`
	BankName                     types.String   `tfsdk:"bank_name"`
	Blocked                      types.Bool     `tfsdk:"blocked"`
	CallManagingOrganizationUuid types.String   `tfsdk:"call_managing_organization_uuid"`
	ContactDetails               types.String   `tfsdk:"contact_details"`
	Country                      types.String   `tfsdk:"country"`
	CountryName                  types.String   `tfsdk:"country_name"`
	Created                      types.String   `tfsdk:"created"`
	CustomerCredit               types.Float64  `tfsdk:"customer_credit"`
	CustomerUnallocatedCredit    types.Float64  `tfsdk:"customer_unallocated_credit"`
	DefaultTaxPercent            types.String   `tfsdk:"default_tax_percent"`
	Description                  types.String   `tfsdk:"description"`
	DisplayBillingInfoInProjects types.Bool     `tfsdk:"display_billing_info_in_projects"`
	DisplayName                  types.String   `tfsdk:"display_name"`
	Domain                       types.String   `tfsdk:"domain"`
	Email                        types.String   `tfsdk:"email"`
	GracePeriodDays              types.Int64    `tfsdk:"grace_period_days"`
	Homepage                     types.String   `tfsdk:"homepage"`
	Image                        types.String   `tfsdk:"image"`
	IsServiceProvider            types.Bool     `tfsdk:"is_service_provider"`
	Latitude                     types.Float64  `tfsdk:"latitude"`
	Longitude                    types.Float64  `tfsdk:"longitude"`
	MaxServiceAccounts           types.Int64    `tfsdk:"max_service_accounts"`
	Name                         types.String   `tfsdk:"name"`
	NativeName                   types.String   `tfsdk:"native_name"`
	NotificationEmails           types.String   `tfsdk:"notification_emails"`
	OrganizationGroups           types.List     `tfsdk:"organization_groups"`
	PaymentProfiles              types.List     `tfsdk:"payment_profiles"`
	PhoneNumber                  types.String   `tfsdk:"phone_number"`
	Postal                       types.String   `tfsdk:"postal"`
	ProjectMetadataChecklist     types.String   `tfsdk:"project_metadata_checklist"`
	ProjectsCount                types.Int64    `tfsdk:"projects_count"`
	RegistrationCode             types.String   `tfsdk:"registration_code"`
	ServiceProvider              types.String   `tfsdk:"service_provider"`
	ServiceProviderUuid          types.String   `tfsdk:"service_provider_uuid"`
	Slug                         types.String   `tfsdk:"slug"`
	SponsorNumber                types.Int64    `tfsdk:"sponsor_number"`
	Url                          types.String   `tfsdk:"url"`
	UsersCount                   types.Int64    `tfsdk:"users_count"`
	VatCode                      types.String   `tfsdk:"vat_code"`
	Timeouts                     timeouts.Value `tfsdk:"timeouts"`
}

func (r *StructureCustomerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_structure_customer"
}

func (r *StructureCustomerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Structure Customer resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"abbreviation": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Abbreviation",
			},
			"access_subnets": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Enter a comma separated list of IPv4 or IPv6 CIDR addresses from where connection to self-service is allowed.",
			},
			"accounting_start_date": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Accounting start date",
			},
			"address": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Address",
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
				MarkdownDescription: "Organization identifier in another application.",
			},
			"bank_account": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Bank account",
			},
			"bank_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name of the bank",
			},
			"blocked": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Blocked",
			},
			"call_managing_organization_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the call managing organization",
			},
			"contact_details": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Contact details",
			},
			"country": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Country code (ISO 3166-1 alpha-2)",
			},
			"country_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Human-readable country name",
			},
			"created": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Created",
			},
			"customer_credit": schema.Float64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Customer credit",
			},
			"customer_unallocated_credit": schema.Float64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Customer unallocated credit",
			},
			"default_tax_percent": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Default tax percent",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description of the resource",
			},
			"display_billing_info_in_projects": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Display billing info in projects",
			},
			"display_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Display name of the organization (includes native name if available)",
			},
			"domain": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Domain",
			},
			"email": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Email",
			},
			"grace_period_days": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Number of extra days after project end date before resources are terminated",
			},
			"homepage": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Homepage",
			},
			"image": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Image",
			},
			"is_service_provider": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Is service provider",
			},
			"latitude": schema.Float64Attribute{
				Optional:            true,
				MarkdownDescription: "Latitude",
			},
			"longitude": schema.Float64Attribute{
				Optional:            true,
				MarkdownDescription: "Longitude",
			},
			"max_service_accounts": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Maximum number of service accounts allowed",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name of the resource",
			},
			"native_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name of the native",
			},
			"notification_emails": schema.StringAttribute{
				Optional:            true,
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
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
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
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Payment profiles",
			},
			"phone_number": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Phone number",
			},
			"postal": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Postal",
			},
			"project_metadata_checklist": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Checklist to be used for project metadata validation in this organization",
			},
			"projects_count": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Number of projects in this organization",
			},
			"registration_code": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Registration code",
			},
			"service_provider": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Service provider",
			},
			"service_provider_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the service provider",
			},
			"slug": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
			},
			"sponsor_number": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "External ID of the sponsor covering the costs",
			},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Url",
			},
			"users_count": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Number of users with access to this organization",
			},
			"vat_code": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "VAT number",
			},
		},

		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx, timeouts.Opts{
				Create: true,
				Update: true,
				Delete: true,
			}),
		},
	}
}

func (r *StructureCustomerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	r.client = NewClient(client)
}

func (r *StructureCustomerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data StructureCustomerResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var requestBody StructureCustomerCreateRequest // Prepare request body
	requestBody.Abbreviation = data.Abbreviation.ValueStringPointer()
	requestBody.AccessSubnets = data.AccessSubnets.ValueStringPointer()
	requestBody.AccountingStartDate = data.AccountingStartDate.ValueStringPointer()
	requestBody.Address = data.Address.ValueStringPointer()
	requestBody.AgreementNumber = data.AgreementNumber.ValueStringPointer()
	requestBody.Archived = data.Archived.ValueBoolPointer()
	requestBody.BackendId = data.BackendId.ValueStringPointer()
	requestBody.BankAccount = data.BankAccount.ValueStringPointer()
	requestBody.BankName = data.BankName.ValueStringPointer()
	requestBody.Blocked = data.Blocked.ValueBoolPointer()
	requestBody.ContactDetails = data.ContactDetails.ValueStringPointer()
	requestBody.Country = data.Country.ValueStringPointer()
	requestBody.DefaultTaxPercent = data.DefaultTaxPercent.ValueStringPointer()
	requestBody.Description = data.Description.ValueStringPointer()
	requestBody.DisplayBillingInfoInProjects = data.DisplayBillingInfoInProjects.ValueBoolPointer()
	requestBody.Domain = data.Domain.ValueStringPointer()
	requestBody.Email = data.Email.ValueStringPointer()
	requestBody.GracePeriodDays = data.GracePeriodDays.ValueInt64Pointer()
	requestBody.Homepage = data.Homepage.ValueStringPointer()
	requestBody.Image = data.Image.ValueStringPointer()
	requestBody.Latitude = data.Latitude.ValueFloat64Pointer()
	requestBody.Longitude = data.Longitude.ValueFloat64Pointer()
	requestBody.MaxServiceAccounts = data.MaxServiceAccounts.ValueInt64Pointer()
	requestBody.Name = data.Name.ValueStringPointer()
	requestBody.NativeName = data.NativeName.ValueStringPointer()
	requestBody.NotificationEmails = data.NotificationEmails.ValueStringPointer()
	requestBody.PhoneNumber = data.PhoneNumber.ValueStringPointer()
	requestBody.Postal = data.Postal.ValueStringPointer()
	requestBody.ProjectMetadataChecklist = data.ProjectMetadataChecklist.ValueStringPointer()
	requestBody.RegistrationCode = data.RegistrationCode.ValueStringPointer()
	requestBody.Slug = data.Slug.ValueStringPointer()
	requestBody.SponsorNumber = data.SponsorNumber.ValueInt64Pointer()
	requestBody.VatCode = data.VatCode.ValueStringPointer()

	apiResp, err := r.client.CreateStructureCustomer(ctx, &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Structure Customer",
			"An error occurred while creating the Structure Customer: "+err.Error(),
		)
		return
	}
	data.UUID = types.StringPointerValue(apiResp.UUID)

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StructureCustomerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data StructureCustomerResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource

	apiResp, err := r.client.GetStructureCustomer(ctx, data.UUID.ValueString())
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Structure Customer",
			"An error occurred while reading the Structure Customer: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StructureCustomerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data StructureCustomerResourceModel
	var state StructureCustomerResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var requestBody StructureCustomerUpdateRequest // Prepare request body
	requestBody.Abbreviation = data.Abbreviation.ValueStringPointer()
	requestBody.AccessSubnets = data.AccessSubnets.ValueStringPointer()
	requestBody.AccountingStartDate = data.AccountingStartDate.ValueStringPointer()
	requestBody.Address = data.Address.ValueStringPointer()
	requestBody.AgreementNumber = data.AgreementNumber.ValueStringPointer()
	requestBody.Archived = data.Archived.ValueBoolPointer()
	requestBody.BackendId = data.BackendId.ValueStringPointer()
	requestBody.BankAccount = data.BankAccount.ValueStringPointer()
	requestBody.BankName = data.BankName.ValueStringPointer()
	requestBody.Blocked = data.Blocked.ValueBoolPointer()
	requestBody.ContactDetails = data.ContactDetails.ValueStringPointer()
	requestBody.Country = data.Country.ValueStringPointer()
	requestBody.DefaultTaxPercent = data.DefaultTaxPercent.ValueStringPointer()
	requestBody.Description = data.Description.ValueStringPointer()
	requestBody.DisplayBillingInfoInProjects = data.DisplayBillingInfoInProjects.ValueBoolPointer()
	requestBody.Domain = data.Domain.ValueStringPointer()
	requestBody.Email = data.Email.ValueStringPointer()
	requestBody.GracePeriodDays = data.GracePeriodDays.ValueInt64Pointer()
	requestBody.Homepage = data.Homepage.ValueStringPointer()
	requestBody.Image = data.Image.ValueStringPointer()
	requestBody.Latitude = data.Latitude.ValueFloat64Pointer()
	requestBody.Longitude = data.Longitude.ValueFloat64Pointer()
	requestBody.MaxServiceAccounts = data.MaxServiceAccounts.ValueInt64Pointer()
	requestBody.Name = data.Name.ValueStringPointer()
	requestBody.NativeName = data.NativeName.ValueStringPointer()
	requestBody.NotificationEmails = data.NotificationEmails.ValueStringPointer()
	requestBody.PhoneNumber = data.PhoneNumber.ValueStringPointer()
	requestBody.Postal = data.Postal.ValueStringPointer()
	requestBody.ProjectMetadataChecklist = data.ProjectMetadataChecklist.ValueStringPointer()
	requestBody.RegistrationCode = data.RegistrationCode.ValueStringPointer()
	requestBody.Slug = data.Slug.ValueStringPointer()
	requestBody.SponsorNumber = data.SponsorNumber.ValueInt64Pointer()
	requestBody.VatCode = data.VatCode.ValueStringPointer()

	apiResp, err := r.client.UpdateStructureCustomer(ctx, data.UUID.ValueString(), &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Structure Customer",
			"An error occurred while updating the Structure Customer: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if apiResp.UUID != nil {
		data.UUID = types.StringPointerValue(apiResp.UUID)
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StructureCustomerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data StructureCustomerResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteStructureCustomer(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Structure Customer",
			"An error occurred while deleting the Structure Customer: "+err.Error(),
		)
		return
	}
}

func (r *StructureCustomerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *StructureCustomerResource) mapResponseToModel(ctx context.Context, apiResp StructureCustomerResponse, model *StructureCustomerResourceModel) diag.Diagnostics {
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
