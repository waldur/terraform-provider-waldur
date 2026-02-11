package customer

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &StructureCustomerResource{}
var _ resource.ResourceWithImportState = &StructureCustomerResource{}

func NewStructureCustomerResource() resource.Resource {
	return &StructureCustomerResource{}
}

// StructureCustomerResource defines the resource implementation.
type StructureCustomerResource struct {
	client *StructureCustomerClient
}

// StructureCustomerResourceModel describes the resource data model.
type StructureCustomerResourceModel struct {
	StructureCustomerModel
	Timeouts timeouts.Value `tfsdk:"timeouts"`
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
				MarkdownDescription: "Structure Customer UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"abbreviation": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Abbreviation",
			},
			"access_subnets": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Enter a comma separated list of IPv4 or IPv6 CIDR addresses from where connection to self-service is allowed.",
			},
			"accounting_start_date": schema.StringAttribute{
				CustomType: timetypes.RFC3339Type{},

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Accounting Start Date",
			},
			"address": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Address",
			},
			"agreement_number": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Agreement Number",
			},
			"archived": schema.BoolAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Bool{

					boolplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Archived",
			},
			"backend_id": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Organization identifier in another application.",
			},
			"bank_account": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Bank Account",
			},
			"bank_name": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Bank Name",
			},
			"billing_price_estimate": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"current": schema.Float64Attribute{

						Computed: true,

						PlanModifiers: []planmodifier.Float64{

							float64planmodifier.UseStateForUnknown(),
						},

						MarkdownDescription: "Current",
					},
					"tax": schema.Float64Attribute{

						Computed: true,

						PlanModifiers: []planmodifier.Float64{

							float64planmodifier.UseStateForUnknown(),
						},

						MarkdownDescription: "Tax",
					},
					"tax_current": schema.Float64Attribute{

						Computed: true,

						PlanModifiers: []planmodifier.Float64{

							float64planmodifier.UseStateForUnknown(),
						},

						MarkdownDescription: "Tax Current",
					},
					"total": schema.Float64Attribute{

						Computed: true,

						PlanModifiers: []planmodifier.Float64{

							float64planmodifier.UseStateForUnknown(),
						},

						MarkdownDescription: "Total",
					},
				},

				Computed: true,

				PlanModifiers: []planmodifier.Object{

					objectplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Billing Price Estimate",
			},
			"blocked": schema.BoolAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Bool{

					boolplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Blocked",
			},
			"call_managing_organization_uuid": schema.StringAttribute{

				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Call Managing Organization Uuid",
			},
			"contact_details": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Contact Details",
			},
			"country": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Country code (ISO 3166-1 alpha-2)",
			},
			"country_name": schema.StringAttribute{

				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Human-readable country name",
			},
			"customer_credit": schema.Float64Attribute{

				Computed: true,

				PlanModifiers: []planmodifier.Float64{

					float64planmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Customer Credit",
			},
			"customer_unallocated_credit": schema.Float64Attribute{

				Computed: true,

				PlanModifiers: []planmodifier.Float64{

					float64planmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Customer Unallocated Credit",
			},
			"default_tax_percent": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Default Tax Percent",

				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,3}(?:\.\d{0,2})?$`), ""),
				},
			},
			"description": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Description",
			},
			"display_billing_info_in_projects": schema.BoolAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Bool{

					boolplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Display Billing Info In Projects",
			},
			"display_name": schema.StringAttribute{

				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Display name of the organization (includes native name if available)",
			},
			"domain": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Domain",
			},
			"email": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Email",
			},
			"grace_period_days": schema.Int64Attribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Int64{

					int64planmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Number of extra days after project end date before resources are terminated",

				Validators: []validator.Int64{
					int64validator.AtLeast(0),
					int64validator.AtMost(2147483647),
				},
			},
			"homepage": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Homepage",
			},
			"image": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Image",
			},
			"is_service_provider": schema.BoolAttribute{

				Computed: true,

				PlanModifiers: []planmodifier.Bool{

					boolplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Is Service Provider",
			},
			"latitude": schema.Float64Attribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Float64{

					float64planmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Latitude",
			},
			"longitude": schema.Float64Attribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Float64{

					float64planmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Longitude",
			},
			"max_service_accounts": schema.Int64Attribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Int64{

					int64planmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Maximum number of service accounts allowed",

				Validators: []validator.Int64{
					int64validator.AtLeast(0),
					int64validator.AtMost(32767),
				},
			},
			"name": schema.StringAttribute{

				Required: true,

				MarkdownDescription: "Name",
			},
			"native_name": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Native Name",
			},
			"notification_emails": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Comma-separated list of notification email addresses",
			},
			"organization_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"customers_count": schema.Int64Attribute{

							Computed: true,

							PlanModifiers: []planmodifier.Int64{

								int64planmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "Number of customers in this organization group",
						},
						"name": schema.StringAttribute{

							Optional: true,
							Computed: true,

							PlanModifiers: []planmodifier.String{

								stringplanmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "Name",
						},
						"parent": schema.StringAttribute{

							Optional: true,
							Computed: true,

							PlanModifiers: []planmodifier.String{

								stringplanmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "Parent",
						},
						"parent_name": schema.StringAttribute{

							Computed: true,

							PlanModifiers: []planmodifier.String{

								stringplanmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "Name of the parent organization group",
						},
						"parent_uuid": schema.StringAttribute{

							Computed: true,

							PlanModifiers: []planmodifier.String{

								stringplanmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "UUID of the parent organization group",
						},
						"url": schema.StringAttribute{

							Computed: true,

							PlanModifiers: []planmodifier.String{

								stringplanmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "Url",
						},
						"uuid": schema.StringAttribute{

							Computed: true,

							PlanModifiers: []planmodifier.String{

								stringplanmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "Uuid",
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

									Optional: true,
									Computed: true,

									PlanModifiers: []planmodifier.String{

										stringplanmodifier.UseStateForUnknown(),
									},

									MarkdownDescription: "Agreement Number",
								},
								"contract_sum": schema.Int64Attribute{

									Optional: true,
									Computed: true,

									PlanModifiers: []planmodifier.Int64{

										int64planmodifier.UseStateForUnknown(),
									},

									MarkdownDescription: "Contract Sum",
								},
								"end_date": schema.StringAttribute{

									Optional: true,
									Computed: true,

									PlanModifiers: []planmodifier.String{

										stringplanmodifier.UseStateForUnknown(),
									},

									MarkdownDescription: "End Date",
								},
							},

							Optional: true,
							Computed: true,

							PlanModifiers: []planmodifier.Object{

								objectplanmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "Attributes",
						},
						"is_active": schema.BoolAttribute{

							Optional: true,
							Computed: true,

							PlanModifiers: []planmodifier.Bool{

								boolplanmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "Is Active",
						},
						"name": schema.StringAttribute{

							Optional: true,
							Computed: true,

							PlanModifiers: []planmodifier.String{

								stringplanmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "Name",
						},
						"organization": schema.StringAttribute{

							Optional: true,
							Computed: true,

							PlanModifiers: []planmodifier.String{

								stringplanmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "Organization",
						},
						"organization_uuid": schema.StringAttribute{

							Computed: true,

							PlanModifiers: []planmodifier.String{

								stringplanmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "Organization Uuid",
						},
						"payment_type": schema.StringAttribute{

							Optional: true,
							Computed: true,

							PlanModifiers: []planmodifier.String{

								stringplanmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "Payment Type",

							Validators: []validator.String{
								stringvalidator.OneOf("fixed_price", "invoices", "payment_gw_monthly"),
							},
						},
						"payment_type_display": schema.StringAttribute{

							Computed: true,

							PlanModifiers: []planmodifier.String{

								stringplanmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "Payment Type Display",
						},
						"url": schema.StringAttribute{

							Computed: true,

							PlanModifiers: []planmodifier.String{

								stringplanmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "Url",
						},
						"uuid": schema.StringAttribute{

							Computed: true,

							PlanModifiers: []planmodifier.String{

								stringplanmodifier.UseStateForUnknown(),
							},

							MarkdownDescription: "Uuid",
						},
					},
				},

				Computed: true,

				PlanModifiers: []planmodifier.List{

					listplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Payment Profiles",
			},
			"phone_number": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Phone Number",
			},
			"postal": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Postal",
			},
			"project_metadata_checklist": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

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

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Registration Code",
			},
			"service_provider": schema.StringAttribute{

				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Service Provider",
			},
			"service_provider_uuid": schema.StringAttribute{

				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "Service Provider Uuid",
			},
			"slug": schema.StringAttribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",

				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
				},
			},
			"sponsor_number": schema.Int64Attribute{

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Int64{

					int64planmodifier.UseStateForUnknown(),
				},

				MarkdownDescription: "External ID of the sponsor covering the costs",

				Validators: []validator.Int64{
					int64validator.AtLeast(0),
					int64validator.AtMost(2147483647),
				},
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

				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				},

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

	r.client = &StructureCustomerClient{}
	if err := r.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			err.Error(),
		)
		return
	}
}

func (r *StructureCustomerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var data StructureCustomerResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := StructureCustomerCreateRequest{}
	if !data.Abbreviation.IsNull() && !data.Abbreviation.IsUnknown() {

		requestBody.Abbreviation = data.Abbreviation.ValueStringPointer()
	}
	if !data.AccessSubnets.IsNull() && !data.AccessSubnets.IsUnknown() {

		requestBody.AccessSubnets = data.AccessSubnets.ValueStringPointer()
	}
	if !data.AccountingStartDate.IsNull() && !data.AccountingStartDate.IsUnknown() {

		requestBody.AccountingStartDate = data.AccountingStartDate.ValueStringPointer()
	}
	if !data.Address.IsNull() && !data.Address.IsUnknown() {

		requestBody.Address = data.Address.ValueStringPointer()
	}
	if !data.AgreementNumber.IsNull() && !data.AgreementNumber.IsUnknown() {

		requestBody.AgreementNumber = data.AgreementNumber.ValueStringPointer()
	}
	if !data.Archived.IsNull() && !data.Archived.IsUnknown() {

		requestBody.Archived = data.Archived.ValueBoolPointer()
	}
	if !data.BackendId.IsNull() && !data.BackendId.IsUnknown() {

		requestBody.BackendId = data.BackendId.ValueStringPointer()
	}
	if !data.BankAccount.IsNull() && !data.BankAccount.IsUnknown() {

		requestBody.BankAccount = data.BankAccount.ValueStringPointer()
	}
	if !data.BankName.IsNull() && !data.BankName.IsUnknown() {

		requestBody.BankName = data.BankName.ValueStringPointer()
	}
	if !data.Blocked.IsNull() && !data.Blocked.IsUnknown() {

		requestBody.Blocked = data.Blocked.ValueBoolPointer()
	}
	if !data.ContactDetails.IsNull() && !data.ContactDetails.IsUnknown() {

		requestBody.ContactDetails = data.ContactDetails.ValueStringPointer()
	}
	if !data.Country.IsNull() && !data.Country.IsUnknown() {

		requestBody.Country = data.Country.ValueStringPointer()
	}
	if !data.DefaultTaxPercent.IsNull() && !data.DefaultTaxPercent.IsUnknown() {

		requestBody.DefaultTaxPercent = data.DefaultTaxPercent.ValueStringPointer()
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {

		requestBody.Description = data.Description.ValueStringPointer()
	}
	if !data.DisplayBillingInfoInProjects.IsNull() && !data.DisplayBillingInfoInProjects.IsUnknown() {

		requestBody.DisplayBillingInfoInProjects = data.DisplayBillingInfoInProjects.ValueBoolPointer()
	}
	if !data.Domain.IsNull() && !data.Domain.IsUnknown() {

		requestBody.Domain = data.Domain.ValueStringPointer()
	}
	if !data.Email.IsNull() && !data.Email.IsUnknown() {

		requestBody.Email = data.Email.ValueStringPointer()
	}
	if !data.GracePeriodDays.IsNull() && !data.GracePeriodDays.IsUnknown() {

		requestBody.GracePeriodDays = data.GracePeriodDays.ValueInt64Pointer()
	}
	if !data.Homepage.IsNull() && !data.Homepage.IsUnknown() {

		requestBody.Homepage = data.Homepage.ValueStringPointer()
	}
	if !data.Image.IsNull() && !data.Image.IsUnknown() {

		requestBody.Image = data.Image.ValueStringPointer()
	}
	if !data.Latitude.IsNull() && !data.Latitude.IsUnknown() {

		requestBody.Latitude = data.Latitude.ValueFloat64Pointer()
	}
	if !data.Longitude.IsNull() && !data.Longitude.IsUnknown() {

		requestBody.Longitude = data.Longitude.ValueFloat64Pointer()
	}
	if !data.MaxServiceAccounts.IsNull() && !data.MaxServiceAccounts.IsUnknown() {

		requestBody.MaxServiceAccounts = data.MaxServiceAccounts.ValueInt64Pointer()
	}

	requestBody.Name = data.Name.ValueStringPointer()
	if !data.NativeName.IsNull() && !data.NativeName.IsUnknown() {

		requestBody.NativeName = data.NativeName.ValueStringPointer()
	}
	if !data.NotificationEmails.IsNull() && !data.NotificationEmails.IsUnknown() {

		requestBody.NotificationEmails = data.NotificationEmails.ValueStringPointer()
	}
	if !data.PhoneNumber.IsNull() && !data.PhoneNumber.IsUnknown() {

		requestBody.PhoneNumber = data.PhoneNumber.ValueStringPointer()
	}
	if !data.Postal.IsNull() && !data.Postal.IsUnknown() {

		requestBody.Postal = data.Postal.ValueStringPointer()
	}
	if !data.ProjectMetadataChecklist.IsNull() && !data.ProjectMetadataChecklist.IsUnknown() {

		requestBody.ProjectMetadataChecklist = data.ProjectMetadataChecklist.ValueStringPointer()
	}
	if !data.RegistrationCode.IsNull() && !data.RegistrationCode.IsUnknown() {

		requestBody.RegistrationCode = data.RegistrationCode.ValueStringPointer()
	}
	if !data.Slug.IsNull() && !data.Slug.IsUnknown() {

		requestBody.Slug = data.Slug.ValueStringPointer()
	}
	if !data.SponsorNumber.IsNull() && !data.SponsorNumber.IsUnknown() {

		requestBody.SponsorNumber = data.SponsorNumber.ValueInt64Pointer()
	}
	if !data.VatCode.IsNull() && !data.VatCode.IsUnknown() {

		requestBody.VatCode = data.VatCode.ValueStringPointer()
	}

	apiResp, err := r.client.Create(ctx, &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Structure Customer",
			"An error occurred while creating the Structure Customer: "+err.Error(),
		)
		return
	}
	data.UUID = types.StringPointerValue(apiResp.UUID)

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
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

	apiResp, err := r.client.Get(ctx, data.UUID.ValueString())
	if err != nil {
		if IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Structure Customer",
			"An error occurred while reading the Structure Customer: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

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

	var apiResp *StructureCustomerResponse
	anyChanges := false
	requestBody := StructureCustomerUpdateRequest{}
	if !data.Abbreviation.IsNull() && !data.Abbreviation.IsUnknown() && !data.Abbreviation.Equal(state.Abbreviation) {
		anyChanges = true

		requestBody.Abbreviation = data.Abbreviation.ValueStringPointer()
	}
	if !data.AccessSubnets.IsNull() && !data.AccessSubnets.IsUnknown() && !data.AccessSubnets.Equal(state.AccessSubnets) {
		anyChanges = true

		requestBody.AccessSubnets = data.AccessSubnets.ValueStringPointer()
	}
	if !data.AccountingStartDate.IsNull() && !data.AccountingStartDate.IsUnknown() && !data.AccountingStartDate.Equal(state.AccountingStartDate) {
		anyChanges = true

		requestBody.AccountingStartDate = data.AccountingStartDate.ValueStringPointer()
	}
	if !data.Address.IsNull() && !data.Address.IsUnknown() && !data.Address.Equal(state.Address) {
		anyChanges = true

		requestBody.Address = data.Address.ValueStringPointer()
	}
	if !data.AgreementNumber.IsNull() && !data.AgreementNumber.IsUnknown() && !data.AgreementNumber.Equal(state.AgreementNumber) {
		anyChanges = true

		requestBody.AgreementNumber = data.AgreementNumber.ValueStringPointer()
	}
	if !data.Archived.IsNull() && !data.Archived.IsUnknown() && !data.Archived.Equal(state.Archived) {
		anyChanges = true

		requestBody.Archived = data.Archived.ValueBoolPointer()
	}
	if !data.BackendId.IsNull() && !data.BackendId.IsUnknown() && !data.BackendId.Equal(state.BackendId) {
		anyChanges = true

		requestBody.BackendId = data.BackendId.ValueStringPointer()
	}
	if !data.BankAccount.IsNull() && !data.BankAccount.IsUnknown() && !data.BankAccount.Equal(state.BankAccount) {
		anyChanges = true

		requestBody.BankAccount = data.BankAccount.ValueStringPointer()
	}
	if !data.BankName.IsNull() && !data.BankName.IsUnknown() && !data.BankName.Equal(state.BankName) {
		anyChanges = true

		requestBody.BankName = data.BankName.ValueStringPointer()
	}
	if !data.Blocked.IsNull() && !data.Blocked.IsUnknown() && !data.Blocked.Equal(state.Blocked) {
		anyChanges = true

		requestBody.Blocked = data.Blocked.ValueBoolPointer()
	}
	if !data.ContactDetails.IsNull() && !data.ContactDetails.IsUnknown() && !data.ContactDetails.Equal(state.ContactDetails) {
		anyChanges = true

		requestBody.ContactDetails = data.ContactDetails.ValueStringPointer()
	}
	if !data.Country.IsNull() && !data.Country.IsUnknown() && !data.Country.Equal(state.Country) {
		anyChanges = true

		requestBody.Country = data.Country.ValueStringPointer()
	}
	if !data.DefaultTaxPercent.IsNull() && !data.DefaultTaxPercent.IsUnknown() && !data.DefaultTaxPercent.Equal(state.DefaultTaxPercent) {
		anyChanges = true

		requestBody.DefaultTaxPercent = data.DefaultTaxPercent.ValueStringPointer()
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() && !data.Description.Equal(state.Description) {
		anyChanges = true

		requestBody.Description = data.Description.ValueStringPointer()
	}
	if !data.DisplayBillingInfoInProjects.IsNull() && !data.DisplayBillingInfoInProjects.IsUnknown() && !data.DisplayBillingInfoInProjects.Equal(state.DisplayBillingInfoInProjects) {
		anyChanges = true

		requestBody.DisplayBillingInfoInProjects = data.DisplayBillingInfoInProjects.ValueBoolPointer()
	}
	if !data.Domain.IsNull() && !data.Domain.IsUnknown() && !data.Domain.Equal(state.Domain) {
		anyChanges = true

		requestBody.Domain = data.Domain.ValueStringPointer()
	}
	if !data.Email.IsNull() && !data.Email.IsUnknown() && !data.Email.Equal(state.Email) {
		anyChanges = true

		requestBody.Email = data.Email.ValueStringPointer()
	}
	if !data.GracePeriodDays.IsNull() && !data.GracePeriodDays.IsUnknown() && !data.GracePeriodDays.Equal(state.GracePeriodDays) {
		anyChanges = true

		requestBody.GracePeriodDays = data.GracePeriodDays.ValueInt64Pointer()
	}
	if !data.Homepage.IsNull() && !data.Homepage.IsUnknown() && !data.Homepage.Equal(state.Homepage) {
		anyChanges = true

		requestBody.Homepage = data.Homepage.ValueStringPointer()
	}
	if !data.Image.IsNull() && !data.Image.IsUnknown() && !data.Image.Equal(state.Image) {
		anyChanges = true

		requestBody.Image = data.Image.ValueStringPointer()
	}
	if !data.Latitude.IsNull() && !data.Latitude.IsUnknown() && !data.Latitude.Equal(state.Latitude) {
		anyChanges = true

		requestBody.Latitude = data.Latitude.ValueFloat64Pointer()
	}
	if !data.Longitude.IsNull() && !data.Longitude.IsUnknown() && !data.Longitude.Equal(state.Longitude) {
		anyChanges = true

		requestBody.Longitude = data.Longitude.ValueFloat64Pointer()
	}
	if !data.MaxServiceAccounts.IsNull() && !data.MaxServiceAccounts.IsUnknown() && !data.MaxServiceAccounts.Equal(state.MaxServiceAccounts) {
		anyChanges = true

		requestBody.MaxServiceAccounts = data.MaxServiceAccounts.ValueInt64Pointer()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() && !data.Name.Equal(state.Name) {
		anyChanges = true

		requestBody.Name = data.Name.ValueStringPointer()
	}
	if !data.NativeName.IsNull() && !data.NativeName.IsUnknown() && !data.NativeName.Equal(state.NativeName) {
		anyChanges = true

		requestBody.NativeName = data.NativeName.ValueStringPointer()
	}
	if !data.NotificationEmails.IsNull() && !data.NotificationEmails.IsUnknown() && !data.NotificationEmails.Equal(state.NotificationEmails) {
		anyChanges = true

		requestBody.NotificationEmails = data.NotificationEmails.ValueStringPointer()
	}
	if !data.PhoneNumber.IsNull() && !data.PhoneNumber.IsUnknown() && !data.PhoneNumber.Equal(state.PhoneNumber) {
		anyChanges = true

		requestBody.PhoneNumber = data.PhoneNumber.ValueStringPointer()
	}
	if !data.Postal.IsNull() && !data.Postal.IsUnknown() && !data.Postal.Equal(state.Postal) {
		anyChanges = true

		requestBody.Postal = data.Postal.ValueStringPointer()
	}
	if !data.ProjectMetadataChecklist.IsNull() && !data.ProjectMetadataChecklist.IsUnknown() && !data.ProjectMetadataChecklist.Equal(state.ProjectMetadataChecklist) {
		anyChanges = true

		requestBody.ProjectMetadataChecklist = data.ProjectMetadataChecklist.ValueStringPointer()
	}
	if !data.RegistrationCode.IsNull() && !data.RegistrationCode.IsUnknown() && !data.RegistrationCode.Equal(state.RegistrationCode) {
		anyChanges = true

		requestBody.RegistrationCode = data.RegistrationCode.ValueStringPointer()
	}
	if !data.Slug.IsNull() && !data.Slug.IsUnknown() && !data.Slug.Equal(state.Slug) {
		anyChanges = true

		requestBody.Slug = data.Slug.ValueStringPointer()
	}
	if !data.SponsorNumber.IsNull() && !data.SponsorNumber.IsUnknown() && !data.SponsorNumber.Equal(state.SponsorNumber) {
		anyChanges = true

		requestBody.SponsorNumber = data.SponsorNumber.ValueInt64Pointer()
	}
	if !data.VatCode.IsNull() && !data.VatCode.IsUnknown() && !data.VatCode.Equal(state.VatCode) {
		anyChanges = true

		requestBody.VatCode = data.VatCode.ValueStringPointer()
	}

	if anyChanges {
		var err error
		apiResp, err = r.client.Update(ctx, data.UUID.ValueString(), &requestBody)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Update Structure Customer",
				"An error occurred while updating the Structure Customer: "+err.Error(),
			)
			return
		}
	}

	newResp, err := r.client.Get(ctx, data.UUID.ValueString())
	if err != nil {
		if IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Failed to Read Resource After Update", err.Error())
		return
	}
	apiResp = newResp

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StructureCustomerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var data StructureCustomerResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Structure Customer",
			"An error occurred while deleting the Structure Customer: "+err.Error(),
		)
		return
	}
}

func (r *StructureCustomerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	uuid := req.ID
	if uuid == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Import ID cannot be empty. Please provide the UUID of the Structure Customer.",
		)
		return
	}

	tflog.Info(ctx, "Importing Structure Customer", map[string]interface{}{
		"uuid": uuid,
	})

	apiResp, err := r.client.Get(ctx, uuid)
	if err != nil {
		if IsNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("Structure Customer with UUID '%s' does not exist or is not accessible.", uuid),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to Import Structure Customer",
			fmt.Sprintf("An error occurred while fetching the Structure Customer: %s", err.Error()),
		)
		return
	}

	var data StructureCustomerResourceModel
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
