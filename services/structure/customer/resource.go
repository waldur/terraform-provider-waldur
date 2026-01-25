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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
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
				CustomType:          timetypes.RFC3339Type{},
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
				CustomType: timetypes.RFC3339Type{},
				Computed:   true,
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
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,3}(?:\.\d{0,2})?$`), ""),
				},
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
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
					int64validator.AtMost(2147483647),
				},
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
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
					int64validator.AtMost(32767),
				},
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
							Validators: []validator.String{
								stringvalidator.OneOf("fixed_price", "invoices", "payment_gw_monthly"),
							},
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
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
				},
			},
			"sponsor_number": schema.Int64Attribute{
				Optional:            true,
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

	requestBody := StructureCustomerCreateRequest{
		Abbreviation:                 data.Abbreviation.ValueStringPointer(),
		AccessSubnets:                data.AccessSubnets.ValueStringPointer(),
		AccountingStartDate:          data.AccountingStartDate.ValueStringPointer(),
		Address:                      data.Address.ValueStringPointer(),
		AgreementNumber:              data.AgreementNumber.ValueStringPointer(),
		Archived:                     data.Archived.ValueBoolPointer(),
		BackendId:                    data.BackendId.ValueStringPointer(),
		BankAccount:                  data.BankAccount.ValueStringPointer(),
		BankName:                     data.BankName.ValueStringPointer(),
		Blocked:                      data.Blocked.ValueBoolPointer(),
		ContactDetails:               data.ContactDetails.ValueStringPointer(),
		Country:                      data.Country.ValueStringPointer(),
		DefaultTaxPercent:            data.DefaultTaxPercent.ValueStringPointer(),
		Description:                  data.Description.ValueStringPointer(),
		DisplayBillingInfoInProjects: data.DisplayBillingInfoInProjects.ValueBoolPointer(),
		Domain:                       data.Domain.ValueStringPointer(),
		Email:                        data.Email.ValueStringPointer(),
		GracePeriodDays:              data.GracePeriodDays.ValueInt64Pointer(),
		Homepage:                     data.Homepage.ValueStringPointer(),
		Image:                        data.Image.ValueStringPointer(),
		Latitude:                     data.Latitude.ValueFloat64Pointer(),
		Longitude:                    data.Longitude.ValueFloat64Pointer(),
		MaxServiceAccounts:           data.MaxServiceAccounts.ValueInt64Pointer(),
		Name:                         data.Name.ValueStringPointer(),
		NativeName:                   data.NativeName.ValueStringPointer(),
		NotificationEmails:           data.NotificationEmails.ValueStringPointer(),
		PhoneNumber:                  data.PhoneNumber.ValueStringPointer(),
		Postal:                       data.Postal.ValueStringPointer(),
		ProjectMetadataChecklist:     data.ProjectMetadataChecklist.ValueStringPointer(),
		RegistrationCode:             data.RegistrationCode.ValueStringPointer(),
		Slug:                         data.Slug.ValueStringPointer(),
		SponsorNumber:                data.SponsorNumber.ValueInt64Pointer(),
		VatCode:                      data.VatCode.ValueStringPointer(),
	}

	apiResp, err := r.client.CreateStructureCustomer(ctx, &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Structure Customer",
			"An error occurred while creating the Structure Customer: "+err.Error(),
		)
		return
	}
	data.UUID = types.StringPointerValue(apiResp.UUID)

	createTimeout, diags := data.Timeouts.Create(ctx, common.DefaultCreateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*StructureCustomerResponse, error) {
		return r.client.GetStructureCustomer(ctx, data.UUID.ValueString())
	}, createTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource creation", err.Error())
		return
	}
	apiResp = newResp

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

	requestBody := StructureCustomerUpdateRequest{
		Abbreviation:                 data.Abbreviation.ValueStringPointer(),
		AccessSubnets:                data.AccessSubnets.ValueStringPointer(),
		AccountingStartDate:          data.AccountingStartDate.ValueStringPointer(),
		Address:                      data.Address.ValueStringPointer(),
		AgreementNumber:              data.AgreementNumber.ValueStringPointer(),
		Archived:                     data.Archived.ValueBoolPointer(),
		BackendId:                    data.BackendId.ValueStringPointer(),
		BankAccount:                  data.BankAccount.ValueStringPointer(),
		BankName:                     data.BankName.ValueStringPointer(),
		Blocked:                      data.Blocked.ValueBoolPointer(),
		ContactDetails:               data.ContactDetails.ValueStringPointer(),
		Country:                      data.Country.ValueStringPointer(),
		DefaultTaxPercent:            data.DefaultTaxPercent.ValueStringPointer(),
		Description:                  data.Description.ValueStringPointer(),
		DisplayBillingInfoInProjects: data.DisplayBillingInfoInProjects.ValueBoolPointer(),
		Domain:                       data.Domain.ValueStringPointer(),
		Email:                        data.Email.ValueStringPointer(),
		GracePeriodDays:              data.GracePeriodDays.ValueInt64Pointer(),
		Homepage:                     data.Homepage.ValueStringPointer(),
		Image:                        data.Image.ValueStringPointer(),
		Latitude:                     data.Latitude.ValueFloat64Pointer(),
		Longitude:                    data.Longitude.ValueFloat64Pointer(),
		MaxServiceAccounts:           data.MaxServiceAccounts.ValueInt64Pointer(),
		Name:                         data.Name.ValueStringPointer(),
		NativeName:                   data.NativeName.ValueStringPointer(),
		NotificationEmails:           data.NotificationEmails.ValueStringPointer(),
		PhoneNumber:                  data.PhoneNumber.ValueStringPointer(),
		Postal:                       data.Postal.ValueStringPointer(),
		ProjectMetadataChecklist:     data.ProjectMetadataChecklist.ValueStringPointer(),
		RegistrationCode:             data.RegistrationCode.ValueStringPointer(),
		Slug:                         data.Slug.ValueStringPointer(),
		SponsorNumber:                data.SponsorNumber.ValueInt64Pointer(),
		VatCode:                      data.VatCode.ValueStringPointer(),
	}

	apiResp, err := r.client.UpdateStructureCustomer(ctx, data.UUID.ValueString(), &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Structure Customer",
			"An error occurred while updating the Structure Customer: "+err.Error(),
		)
		return
	}

	updateTimeout, diags := data.Timeouts.Update(ctx, common.DefaultUpdateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*StructureCustomerResponse, error) {
		return r.client.GetStructureCustomer(ctx, data.UUID.ValueString())
	}, updateTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource update", err.Error())
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

	err := r.client.DeleteStructureCustomer(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Structure Customer",
			"An error occurred while deleting the Structure Customer: "+err.Error(),
		)
		return
	}

	deleteTimeout, diags := data.Timeouts.Delete(ctx, common.DefaultDeleteTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err = common.WaitForDeletion(ctx, func(ctx context.Context) (*StructureCustomerResponse, error) {
		return r.client.GetStructureCustomer(ctx, data.UUID.ValueString())
	}, deleteTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource deletion", err.Error())
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

	apiResp, err := r.client.GetStructureCustomer(ctx, uuid)
	if err != nil {
		if client.IsNotFoundError(err) {
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
