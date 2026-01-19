package resources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	client *client.Client
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
	Projects                     types.List     `tfsdk:"projects"`
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
		MarkdownDescription: "StructureCustomer resource",

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
				Computed:            true,
				MarkdownDescription: "",
			},
			"access_subnets": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Enter a comma separated list of IPv4 or IPv6 CIDR addresses from where connection to self-service is allowed.",
			},
			"accounting_start_date": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"address": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"agreement_number": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"archived": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Organization identifier in another application.",
			},
			"bank_account": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"bank_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"blocked": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"call_managing_organization_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"contact_details": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"country": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"country_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer_credit": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer_unallocated_credit": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"default_tax_percent": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"display_billing_info_in_projects": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"display_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"domain": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"email": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"grace_period_days": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Number of extra days after project end date before resources are terminated",
			},
			"homepage": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"image": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"is_service_provider": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"latitude": schema.Float64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"longitude": schema.Float64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"max_service_accounts": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Maximum number of service accounts allowed",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "",
			},
			"native_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"notification_emails": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Comma-separated list of notification email addresses",
			},
			"organization_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"customers_count": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "",
						},
						"parent": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "",
						},
						"parent_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
						"parent_uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "",
			},
			"payment_profiles": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"attributes": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"agreement_number": schema.StringAttribute{
									Optional:            true,
									Computed:            true,
									MarkdownDescription: "",
								},
								"contract_sum": schema.Int64Attribute{
									Optional:            true,
									Computed:            true,
									MarkdownDescription: "",
								},
								"end_date": schema.StringAttribute{
									Optional:            true,
									Computed:            true,
									MarkdownDescription: "",
								},
							},
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "",
						},
						"is_active": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "",
						},
						"organization": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "",
						},
						"organization_uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
						"payment_type": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "",
						},
						"payment_type_display": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "",
			},
			"phone_number": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"postal": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"project_metadata_checklist": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"projects": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"end_date": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "The date is inclusive. Once reached, all project resource will be scheduled for termination.",
						},
						"image": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "",
						},
						"resource_count": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "",
			},
			"projects_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"registration_code": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"service_provider": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"service_provider_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"slug": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"sponsor_number": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "External ID of the sponsor covering the costs",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"users_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"vat_code": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
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

func (r *StructureCustomerResource) convertTFValue(v attr.Value) interface{} {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	switch val := v.(type) {
	case types.String:
		return val.ValueString()
	case types.Int64:
		return val.ValueInt64()
	case types.Bool:
		return val.ValueBool()
	case types.Float64:
		return val.ValueFloat64()
	case types.List:
		items := make([]interface{}, len(val.Elements()))
		for i, elem := range val.Elements() {
			items[i] = r.convertTFValue(elem)
		}
		return items
	case types.Object:
		obj := make(map[string]interface{})
		for k, attr := range val.Attributes() {
			if converted := r.convertTFValue(attr); converted != nil {
				obj[k] = converted
			}
		}
		return obj
	}
	return nil
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

	r.client = client
}

func (r *StructureCustomerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data StructureCustomerResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Prepare request body
	requestBody := map[string]interface{}{}
	if !data.Abbreviation.IsNull() && !data.Abbreviation.IsUnknown() {
		if v := data.Abbreviation.ValueString(); v != "" {
			requestBody["abbreviation"] = v
		}
	}
	if !data.AccessSubnets.IsNull() && !data.AccessSubnets.IsUnknown() {
		if v := data.AccessSubnets.ValueString(); v != "" {
			requestBody["access_subnets"] = v
		}
	}
	if !data.AccountingStartDate.IsNull() && !data.AccountingStartDate.IsUnknown() {
		if v := data.AccountingStartDate.ValueString(); v != "" {
			requestBody["accounting_start_date"] = v
		}
	}
	if !data.Address.IsNull() && !data.Address.IsUnknown() {
		if v := data.Address.ValueString(); v != "" {
			requestBody["address"] = v
		}
	}
	if !data.AgreementNumber.IsNull() && !data.AgreementNumber.IsUnknown() {
		if v := data.AgreementNumber.ValueString(); v != "" {
			requestBody["agreement_number"] = v
		}
	}
	if !data.Archived.IsNull() && !data.Archived.IsUnknown() {
		requestBody["archived"] = data.Archived.ValueBool()
	}
	if !data.BackendId.IsNull() && !data.BackendId.IsUnknown() {
		if v := data.BackendId.ValueString(); v != "" {
			requestBody["backend_id"] = v
		}
	}
	if !data.BankAccount.IsNull() && !data.BankAccount.IsUnknown() {
		if v := data.BankAccount.ValueString(); v != "" {
			requestBody["bank_account"] = v
		}
	}
	if !data.BankName.IsNull() && !data.BankName.IsUnknown() {
		if v := data.BankName.ValueString(); v != "" {
			requestBody["bank_name"] = v
		}
	}
	if !data.Blocked.IsNull() && !data.Blocked.IsUnknown() {
		requestBody["blocked"] = data.Blocked.ValueBool()
	}
	if !data.ContactDetails.IsNull() && !data.ContactDetails.IsUnknown() {
		if v := data.ContactDetails.ValueString(); v != "" {
			requestBody["contact_details"] = v
		}
	}
	if !data.Country.IsNull() && !data.Country.IsUnknown() {
		if v := data.Country.ValueString(); v != "" {
			requestBody["country"] = v
		}
	}
	if !data.DefaultTaxPercent.IsNull() && !data.DefaultTaxPercent.IsUnknown() {
		if v := data.DefaultTaxPercent.ValueString(); v != "" {
			requestBody["default_tax_percent"] = v
		}
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		if v := data.Description.ValueString(); v != "" {
			requestBody["description"] = v
		}
	}
	if !data.DisplayBillingInfoInProjects.IsNull() && !data.DisplayBillingInfoInProjects.IsUnknown() {
		requestBody["display_billing_info_in_projects"] = data.DisplayBillingInfoInProjects.ValueBool()
	}
	if !data.Domain.IsNull() && !data.Domain.IsUnknown() {
		if v := data.Domain.ValueString(); v != "" {
			requestBody["domain"] = v
		}
	}
	if !data.Email.IsNull() && !data.Email.IsUnknown() {
		if v := data.Email.ValueString(); v != "" {
			requestBody["email"] = v
		}
	}
	if !data.GracePeriodDays.IsNull() && !data.GracePeriodDays.IsUnknown() {
		requestBody["grace_period_days"] = data.GracePeriodDays.ValueInt64()
	}
	if !data.Homepage.IsNull() && !data.Homepage.IsUnknown() {
		if v := data.Homepage.ValueString(); v != "" {
			requestBody["homepage"] = v
		}
	}
	if !data.Image.IsNull() && !data.Image.IsUnknown() {
		if v := data.Image.ValueString(); v != "" {
			requestBody["image"] = v
		}
	}
	if !data.Latitude.IsNull() && !data.Latitude.IsUnknown() {
		requestBody["latitude"] = data.Latitude.ValueFloat64()
	}
	if !data.Longitude.IsNull() && !data.Longitude.IsUnknown() {
		requestBody["longitude"] = data.Longitude.ValueFloat64()
	}
	if !data.MaxServiceAccounts.IsNull() && !data.MaxServiceAccounts.IsUnknown() {
		requestBody["max_service_accounts"] = data.MaxServiceAccounts.ValueInt64()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		if v := data.Name.ValueString(); v != "" {
			requestBody["name"] = v
		}
	}
	if !data.NativeName.IsNull() && !data.NativeName.IsUnknown() {
		if v := data.NativeName.ValueString(); v != "" {
			requestBody["native_name"] = v
		}
	}
	if !data.NotificationEmails.IsNull() && !data.NotificationEmails.IsUnknown() {
		if v := data.NotificationEmails.ValueString(); v != "" {
			requestBody["notification_emails"] = v
		}
	}
	if !data.PhoneNumber.IsNull() && !data.PhoneNumber.IsUnknown() {
		if v := data.PhoneNumber.ValueString(); v != "" {
			requestBody["phone_number"] = v
		}
	}
	if !data.Postal.IsNull() && !data.Postal.IsUnknown() {
		if v := data.Postal.ValueString(); v != "" {
			requestBody["postal"] = v
		}
	}
	if !data.ProjectMetadataChecklist.IsNull() && !data.ProjectMetadataChecklist.IsUnknown() {
		if v := data.ProjectMetadataChecklist.ValueString(); v != "" {
			requestBody["project_metadata_checklist"] = v
		}
	}
	if !data.RegistrationCode.IsNull() && !data.RegistrationCode.IsUnknown() {
		if v := data.RegistrationCode.ValueString(); v != "" {
			requestBody["registration_code"] = v
		}
	}
	if !data.Slug.IsNull() && !data.Slug.IsUnknown() {
		if v := data.Slug.ValueString(); v != "" {
			requestBody["slug"] = v
		}
	}
	if !data.SponsorNumber.IsNull() && !data.SponsorNumber.IsUnknown() {
		requestBody["sponsor_number"] = data.SponsorNumber.ValueInt64()
	}
	if !data.VatCode.IsNull() && !data.VatCode.IsUnknown() {
		if v := data.VatCode.ValueString(); v != "" {
			requestBody["vat_code"] = v
		}
	}

	// Call Waldur API to create resource
	var result map[string]interface{}
	err := r.client.Create(ctx, "/api/customers/", requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create StructureCustomer",
			"An error occurred while creating the structure_customer: "+err.Error(),
		)
		return
	}
	// Extract UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
	// Map response fields to data model
	_ = sourceMap
	if val, ok := sourceMap["abbreviation"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Abbreviation = types.StringValue(str)
		}
	} else {
		if data.Abbreviation.IsUnknown() {
			data.Abbreviation = types.StringNull()
		}
	}
	if val, ok := sourceMap["access_subnets"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AccessSubnets = types.StringValue(str)
		}
	} else {
		if data.AccessSubnets.IsUnknown() {
			data.AccessSubnets = types.StringNull()
		}
	}
	if val, ok := sourceMap["accounting_start_date"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AccountingStartDate = types.StringValue(str)
		}
	} else {
		if data.AccountingStartDate.IsUnknown() {
			data.AccountingStartDate = types.StringNull()
		}
	}
	if val, ok := sourceMap["address"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Address = types.StringValue(str)
		}
	} else {
		if data.Address.IsUnknown() {
			data.Address = types.StringNull()
		}
	}
	if val, ok := sourceMap["agreement_number"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AgreementNumber = types.StringValue(str)
		}
	} else {
		if data.AgreementNumber.IsUnknown() {
			data.AgreementNumber = types.StringNull()
		}
	}
	if val, ok := sourceMap["archived"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Archived = types.BoolValue(b)
		}
	} else {
		if data.Archived.IsUnknown() {
			data.Archived = types.BoolNull()
		}
	}
	if val, ok := sourceMap["backend_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
		}
	} else {
		if data.BackendId.IsUnknown() {
			data.BackendId = types.StringNull()
		}
	}
	if val, ok := sourceMap["bank_account"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BankAccount = types.StringValue(str)
		}
	} else {
		if data.BankAccount.IsUnknown() {
			data.BankAccount = types.StringNull()
		}
	}
	if val, ok := sourceMap["bank_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BankName = types.StringValue(str)
		}
	} else {
		if data.BankName.IsUnknown() {
			data.BankName = types.StringNull()
		}
	}
	if val, ok := sourceMap["blocked"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Blocked = types.BoolValue(b)
		}
	} else {
		if data.Blocked.IsUnknown() {
			data.Blocked = types.BoolNull()
		}
	}
	if val, ok := sourceMap["call_managing_organization_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CallManagingOrganizationUuid = types.StringValue(str)
		}
	} else {
		if data.CallManagingOrganizationUuid.IsUnknown() {
			data.CallManagingOrganizationUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["contact_details"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ContactDetails = types.StringValue(str)
		}
	} else {
		if data.ContactDetails.IsUnknown() {
			data.ContactDetails = types.StringNull()
		}
	}
	if val, ok := sourceMap["country"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Country = types.StringValue(str)
		}
	} else {
		if data.Country.IsUnknown() {
			data.Country = types.StringNull()
		}
	}
	if val, ok := sourceMap["country_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CountryName = types.StringValue(str)
		}
	} else {
		if data.CountryName.IsUnknown() {
			data.CountryName = types.StringNull()
		}
	}
	if val, ok := sourceMap["created"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
	} else {
		if data.Created.IsUnknown() {
			data.Created = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_credit"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.CustomerCredit = types.Float64Value(num)
		}
	} else {
		if data.CustomerCredit.IsUnknown() {
			data.CustomerCredit = types.Float64Null()
		}
	}
	if val, ok := sourceMap["customer_unallocated_credit"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.CustomerUnallocatedCredit = types.Float64Value(num)
		}
	} else {
		if data.CustomerUnallocatedCredit.IsUnknown() {
			data.CustomerUnallocatedCredit = types.Float64Null()
		}
	}
	if val, ok := sourceMap["default_tax_percent"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.DefaultTaxPercent = types.StringValue(str)
		}
	} else {
		if data.DefaultTaxPercent.IsUnknown() {
			data.DefaultTaxPercent = types.StringNull()
		}
	}
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
		}
	}
	if val, ok := sourceMap["display_billing_info_in_projects"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.DisplayBillingInfoInProjects = types.BoolValue(b)
		}
	} else {
		if data.DisplayBillingInfoInProjects.IsUnknown() {
			data.DisplayBillingInfoInProjects = types.BoolNull()
		}
	}
	if val, ok := sourceMap["display_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.DisplayName = types.StringValue(str)
		}
	} else {
		if data.DisplayName.IsUnknown() {
			data.DisplayName = types.StringNull()
		}
	}
	if val, ok := sourceMap["domain"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Domain = types.StringValue(str)
		}
	} else {
		if data.Domain.IsUnknown() {
			data.Domain = types.StringNull()
		}
	}
	if val, ok := sourceMap["email"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Email = types.StringValue(str)
		}
	} else {
		if data.Email.IsUnknown() {
			data.Email = types.StringNull()
		}
	}
	if val, ok := sourceMap["grace_period_days"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.GracePeriodDays = types.Int64Value(int64(num))
		}
	} else {
		if data.GracePeriodDays.IsUnknown() {
			data.GracePeriodDays = types.Int64Null()
		}
	}
	if val, ok := sourceMap["homepage"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Homepage = types.StringValue(str)
		}
	} else {
		if data.Homepage.IsUnknown() {
			data.Homepage = types.StringNull()
		}
	}
	if val, ok := sourceMap["image"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Image = types.StringValue(str)
		}
	} else {
		if data.Image.IsUnknown() {
			data.Image = types.StringNull()
		}
	}
	if val, ok := sourceMap["is_service_provider"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.IsServiceProvider = types.BoolValue(b)
		}
	} else {
		if data.IsServiceProvider.IsUnknown() {
			data.IsServiceProvider = types.BoolNull()
		}
	}
	if val, ok := sourceMap["latitude"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Latitude = types.Float64Value(num)
		}
	} else {
		if data.Latitude.IsUnknown() {
			data.Latitude = types.Float64Null()
		}
	}
	if val, ok := sourceMap["longitude"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Longitude = types.Float64Value(num)
		}
	} else {
		if data.Longitude.IsUnknown() {
			data.Longitude = types.Float64Null()
		}
	}
	if val, ok := sourceMap["max_service_accounts"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.MaxServiceAccounts = types.Int64Value(int64(num))
		}
	} else {
		if data.MaxServiceAccounts.IsUnknown() {
			data.MaxServiceAccounts = types.Int64Null()
		}
	}
	if val, ok := sourceMap["native_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.NativeName = types.StringValue(str)
		}
	} else {
		if data.NativeName.IsUnknown() {
			data.NativeName = types.StringNull()
		}
	}
	if val, ok := sourceMap["notification_emails"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.NotificationEmails = types.StringValue(str)
		}
	} else {
		if data.NotificationEmails.IsUnknown() {
			data.NotificationEmails = types.StringNull()
		}
	}
	if val, ok := sourceMap["organization_groups"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"customers_count": types.Int64Type,
						"name":            types.StringType,
						"parent":          types.StringType,
						"parent_name":     types.StringType,
						"parent_uuid":     types.StringType,
						"url":             types.StringType,
					}
					attrValues := map[string]attr.Value{
						"customers_count": func() attr.Value {
							if v, ok := objMap["customers_count"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"parent": func() attr.Value {
							if v, ok := objMap["parent"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"parent_name": func() attr.Value {
							if v, ok := objMap["parent_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"parent_uuid": func() attr.Value {
							if v, ok := objMap["parent_uuid"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"customers_count": types.Int64Type,
				"name":            types.StringType,
				"parent":          types.StringType,
				"parent_name":     types.StringType,
				"parent_uuid":     types.StringType,
				"url":             types.StringType,
			}}, items)
			data.OrganizationGroups = listVal
		}
	} else {
		if data.OrganizationGroups.IsUnknown() {
			data.OrganizationGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"customers_count": types.Int64Type,
				"name":            types.StringType,
				"parent":          types.StringType,
				"parent_name":     types.StringType,
				"parent_uuid":     types.StringType,
				"url":             types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["payment_profiles"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"attributes":           types.ObjectType{AttrTypes: map[string]attr.Type{"agreement_number": types.StringType, "contract_sum": types.Int64Type, "end_date": types.StringType}},
						"is_active":            types.BoolType,
						"name":                 types.StringType,
						"organization":         types.StringType,
						"organization_uuid":    types.StringType,
						"payment_type":         types.StringType,
						"payment_type_display": types.StringType,
						"url":                  types.StringType,
					}
					attrValues := map[string]attr.Value{
						"attributes": types.ObjectNull(types.ObjectType{AttrTypes: map[string]attr.Type{"agreement_number": types.StringType, "contract_sum": types.Int64Type, "end_date": types.StringType}}.AttrTypes),
						"is_active": func() attr.Value {
							if v, ok := objMap["is_active"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"organization": func() attr.Value {
							if v, ok := objMap["organization"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"organization_uuid": func() attr.Value {
							if v, ok := objMap["organization_uuid"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"payment_type": func() attr.Value {
							if v, ok := objMap["payment_type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"payment_type_display": func() attr.Value {
							if v, ok := objMap["payment_type_display"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"attributes":           types.ObjectType{AttrTypes: map[string]attr.Type{"agreement_number": types.StringType, "contract_sum": types.Int64Type, "end_date": types.StringType}},
				"is_active":            types.BoolType,
				"name":                 types.StringType,
				"organization":         types.StringType,
				"organization_uuid":    types.StringType,
				"payment_type":         types.StringType,
				"payment_type_display": types.StringType,
				"url":                  types.StringType,
			}}, items)
			data.PaymentProfiles = listVal
		}
	} else {
		if data.PaymentProfiles.IsUnknown() {
			data.PaymentProfiles = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"attributes":           types.ObjectType{AttrTypes: map[string]attr.Type{"agreement_number": types.StringType, "contract_sum": types.Int64Type, "end_date": types.StringType}},
				"is_active":            types.BoolType,
				"name":                 types.StringType,
				"organization":         types.StringType,
				"organization_uuid":    types.StringType,
				"payment_type":         types.StringType,
				"payment_type_display": types.StringType,
				"url":                  types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["phone_number"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.PhoneNumber = types.StringValue(str)
		}
	} else {
		if data.PhoneNumber.IsUnknown() {
			data.PhoneNumber = types.StringNull()
		}
	}
	if val, ok := sourceMap["postal"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Postal = types.StringValue(str)
		}
	} else {
		if data.Postal.IsUnknown() {
			data.Postal = types.StringNull()
		}
	}
	if val, ok := sourceMap["project_metadata_checklist"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProjectMetadataChecklist = types.StringValue(str)
		}
	} else {
		if data.ProjectMetadataChecklist.IsUnknown() {
			data.ProjectMetadataChecklist = types.StringNull()
		}
	}
	if val, ok := sourceMap["projects"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"end_date":       types.StringType,
						"image":          types.StringType,
						"name":           types.StringType,
						"resource_count": types.Int64Type,
						"url":            types.StringType,
					}
					attrValues := map[string]attr.Value{
						"end_date": func() attr.Value {
							if v, ok := objMap["end_date"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"image": func() attr.Value {
							if v, ok := objMap["image"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"resource_count": func() attr.Value {
							if v, ok := objMap["resource_count"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"end_date":       types.StringType,
				"image":          types.StringType,
				"name":           types.StringType,
				"resource_count": types.Int64Type,
				"url":            types.StringType,
			}}, items)
			data.Projects = listVal
		}
	} else {
		if data.Projects.IsUnknown() {
			data.Projects = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"end_date":       types.StringType,
				"image":          types.StringType,
				"name":           types.StringType,
				"resource_count": types.Int64Type,
				"url":            types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["projects_count"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.ProjectsCount = types.Int64Value(int64(num))
		}
	} else {
		if data.ProjectsCount.IsUnknown() {
			data.ProjectsCount = types.Int64Null()
		}
	}
	if val, ok := sourceMap["registration_code"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.RegistrationCode = types.StringValue(str)
		}
	} else {
		if data.RegistrationCode.IsUnknown() {
			data.RegistrationCode = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_provider"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceProvider = types.StringValue(str)
		}
	} else {
		if data.ServiceProvider.IsUnknown() {
			data.ServiceProvider = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_provider_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceProviderUuid = types.StringValue(str)
		}
	} else {
		if data.ServiceProviderUuid.IsUnknown() {
			data.ServiceProviderUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["slug"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Slug = types.StringValue(str)
		}
	} else {
		if data.Slug.IsUnknown() {
			data.Slug = types.StringNull()
		}
	}
	if val, ok := sourceMap["sponsor_number"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.SponsorNumber = types.Int64Value(int64(num))
		}
	} else {
		if data.SponsorNumber.IsUnknown() {
			data.SponsorNumber = types.Int64Null()
		}
	}
	if val, ok := sourceMap["url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Url = types.StringValue(str)
		}
	} else {
		if data.Url.IsUnknown() {
			data.Url = types.StringNull()
		}
	}
	if val, ok := sourceMap["users_count"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.UsersCount = types.Int64Value(int64(num))
		}
	} else {
		if data.UsersCount.IsUnknown() {
			data.UsersCount = types.Int64Null()
		}
	}
	if val, ok := sourceMap["vat_code"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.VatCode = types.StringValue(str)
		}
	} else {
		if data.VatCode.IsUnknown() {
			data.VatCode = types.StringNull()
		}
	}

	// Map filter parameters from response if available

	// Save data into Terraform state
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
	var result map[string]interface{}

	retrievePath := strings.Replace("/api/customers/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read StructureCustomer",
			"An error occurred while reading the structure_customer: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
	// Map response fields to data model
	_ = sourceMap
	if val, ok := sourceMap["abbreviation"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Abbreviation = types.StringValue(str)
		}
	} else {
		if data.Abbreviation.IsUnknown() {
			data.Abbreviation = types.StringNull()
		}
	}
	if val, ok := sourceMap["access_subnets"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AccessSubnets = types.StringValue(str)
		}
	} else {
		if data.AccessSubnets.IsUnknown() {
			data.AccessSubnets = types.StringNull()
		}
	}
	if val, ok := sourceMap["accounting_start_date"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AccountingStartDate = types.StringValue(str)
		}
	} else {
		if data.AccountingStartDate.IsUnknown() {
			data.AccountingStartDate = types.StringNull()
		}
	}
	if val, ok := sourceMap["address"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Address = types.StringValue(str)
		}
	} else {
		if data.Address.IsUnknown() {
			data.Address = types.StringNull()
		}
	}
	if val, ok := sourceMap["agreement_number"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AgreementNumber = types.StringValue(str)
		}
	} else {
		if data.AgreementNumber.IsUnknown() {
			data.AgreementNumber = types.StringNull()
		}
	}
	if val, ok := sourceMap["archived"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Archived = types.BoolValue(b)
		}
	} else {
		if data.Archived.IsUnknown() {
			data.Archived = types.BoolNull()
		}
	}
	if val, ok := sourceMap["backend_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
		}
	} else {
		if data.BackendId.IsUnknown() {
			data.BackendId = types.StringNull()
		}
	}
	if val, ok := sourceMap["bank_account"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BankAccount = types.StringValue(str)
		}
	} else {
		if data.BankAccount.IsUnknown() {
			data.BankAccount = types.StringNull()
		}
	}
	if val, ok := sourceMap["bank_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BankName = types.StringValue(str)
		}
	} else {
		if data.BankName.IsUnknown() {
			data.BankName = types.StringNull()
		}
	}
	if val, ok := sourceMap["blocked"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Blocked = types.BoolValue(b)
		}
	} else {
		if data.Blocked.IsUnknown() {
			data.Blocked = types.BoolNull()
		}
	}
	if val, ok := sourceMap["call_managing_organization_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CallManagingOrganizationUuid = types.StringValue(str)
		}
	} else {
		if data.CallManagingOrganizationUuid.IsUnknown() {
			data.CallManagingOrganizationUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["contact_details"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ContactDetails = types.StringValue(str)
		}
	} else {
		if data.ContactDetails.IsUnknown() {
			data.ContactDetails = types.StringNull()
		}
	}
	if val, ok := sourceMap["country"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Country = types.StringValue(str)
		}
	} else {
		if data.Country.IsUnknown() {
			data.Country = types.StringNull()
		}
	}
	if val, ok := sourceMap["country_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CountryName = types.StringValue(str)
		}
	} else {
		if data.CountryName.IsUnknown() {
			data.CountryName = types.StringNull()
		}
	}
	if val, ok := sourceMap["created"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
	} else {
		if data.Created.IsUnknown() {
			data.Created = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_credit"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.CustomerCredit = types.Float64Value(num)
		}
	} else {
		if data.CustomerCredit.IsUnknown() {
			data.CustomerCredit = types.Float64Null()
		}
	}
	if val, ok := sourceMap["customer_unallocated_credit"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.CustomerUnallocatedCredit = types.Float64Value(num)
		}
	} else {
		if data.CustomerUnallocatedCredit.IsUnknown() {
			data.CustomerUnallocatedCredit = types.Float64Null()
		}
	}
	if val, ok := sourceMap["default_tax_percent"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.DefaultTaxPercent = types.StringValue(str)
		}
	} else {
		if data.DefaultTaxPercent.IsUnknown() {
			data.DefaultTaxPercent = types.StringNull()
		}
	}
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
		}
	}
	if val, ok := sourceMap["display_billing_info_in_projects"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.DisplayBillingInfoInProjects = types.BoolValue(b)
		}
	} else {
		if data.DisplayBillingInfoInProjects.IsUnknown() {
			data.DisplayBillingInfoInProjects = types.BoolNull()
		}
	}
	if val, ok := sourceMap["display_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.DisplayName = types.StringValue(str)
		}
	} else {
		if data.DisplayName.IsUnknown() {
			data.DisplayName = types.StringNull()
		}
	}
	if val, ok := sourceMap["domain"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Domain = types.StringValue(str)
		}
	} else {
		if data.Domain.IsUnknown() {
			data.Domain = types.StringNull()
		}
	}
	if val, ok := sourceMap["email"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Email = types.StringValue(str)
		}
	} else {
		if data.Email.IsUnknown() {
			data.Email = types.StringNull()
		}
	}
	if val, ok := sourceMap["grace_period_days"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.GracePeriodDays = types.Int64Value(int64(num))
		}
	} else {
		if data.GracePeriodDays.IsUnknown() {
			data.GracePeriodDays = types.Int64Null()
		}
	}
	if val, ok := sourceMap["homepage"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Homepage = types.StringValue(str)
		}
	} else {
		if data.Homepage.IsUnknown() {
			data.Homepage = types.StringNull()
		}
	}
	if val, ok := sourceMap["image"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Image = types.StringValue(str)
		}
	} else {
		if data.Image.IsUnknown() {
			data.Image = types.StringNull()
		}
	}
	if val, ok := sourceMap["is_service_provider"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.IsServiceProvider = types.BoolValue(b)
		}
	} else {
		if data.IsServiceProvider.IsUnknown() {
			data.IsServiceProvider = types.BoolNull()
		}
	}
	if val, ok := sourceMap["latitude"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Latitude = types.Float64Value(num)
		}
	} else {
		if data.Latitude.IsUnknown() {
			data.Latitude = types.Float64Null()
		}
	}
	if val, ok := sourceMap["longitude"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Longitude = types.Float64Value(num)
		}
	} else {
		if data.Longitude.IsUnknown() {
			data.Longitude = types.Float64Null()
		}
	}
	if val, ok := sourceMap["max_service_accounts"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.MaxServiceAccounts = types.Int64Value(int64(num))
		}
	} else {
		if data.MaxServiceAccounts.IsUnknown() {
			data.MaxServiceAccounts = types.Int64Null()
		}
	}
	if val, ok := sourceMap["native_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.NativeName = types.StringValue(str)
		}
	} else {
		if data.NativeName.IsUnknown() {
			data.NativeName = types.StringNull()
		}
	}
	if val, ok := sourceMap["notification_emails"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.NotificationEmails = types.StringValue(str)
		}
	} else {
		if data.NotificationEmails.IsUnknown() {
			data.NotificationEmails = types.StringNull()
		}
	}
	if val, ok := sourceMap["organization_groups"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"customers_count": types.Int64Type,
						"name":            types.StringType,
						"parent":          types.StringType,
						"parent_name":     types.StringType,
						"parent_uuid":     types.StringType,
						"url":             types.StringType,
					}
					attrValues := map[string]attr.Value{
						"customers_count": func() attr.Value {
							if v, ok := objMap["customers_count"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"parent": func() attr.Value {
							if v, ok := objMap["parent"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"parent_name": func() attr.Value {
							if v, ok := objMap["parent_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"parent_uuid": func() attr.Value {
							if v, ok := objMap["parent_uuid"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"customers_count": types.Int64Type,
				"name":            types.StringType,
				"parent":          types.StringType,
				"parent_name":     types.StringType,
				"parent_uuid":     types.StringType,
				"url":             types.StringType,
			}}, items)
			data.OrganizationGroups = listVal
		}
	} else {
		if data.OrganizationGroups.IsUnknown() {
			data.OrganizationGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"customers_count": types.Int64Type,
				"name":            types.StringType,
				"parent":          types.StringType,
				"parent_name":     types.StringType,
				"parent_uuid":     types.StringType,
				"url":             types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["payment_profiles"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"attributes":           types.ObjectType{AttrTypes: map[string]attr.Type{"agreement_number": types.StringType, "contract_sum": types.Int64Type, "end_date": types.StringType}},
						"is_active":            types.BoolType,
						"name":                 types.StringType,
						"organization":         types.StringType,
						"organization_uuid":    types.StringType,
						"payment_type":         types.StringType,
						"payment_type_display": types.StringType,
						"url":                  types.StringType,
					}
					attrValues := map[string]attr.Value{
						"attributes": types.ObjectNull(types.ObjectType{AttrTypes: map[string]attr.Type{"agreement_number": types.StringType, "contract_sum": types.Int64Type, "end_date": types.StringType}}.AttrTypes),
						"is_active": func() attr.Value {
							if v, ok := objMap["is_active"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"organization": func() attr.Value {
							if v, ok := objMap["organization"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"organization_uuid": func() attr.Value {
							if v, ok := objMap["organization_uuid"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"payment_type": func() attr.Value {
							if v, ok := objMap["payment_type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"payment_type_display": func() attr.Value {
							if v, ok := objMap["payment_type_display"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"attributes":           types.ObjectType{AttrTypes: map[string]attr.Type{"agreement_number": types.StringType, "contract_sum": types.Int64Type, "end_date": types.StringType}},
				"is_active":            types.BoolType,
				"name":                 types.StringType,
				"organization":         types.StringType,
				"organization_uuid":    types.StringType,
				"payment_type":         types.StringType,
				"payment_type_display": types.StringType,
				"url":                  types.StringType,
			}}, items)
			data.PaymentProfiles = listVal
		}
	} else {
		if data.PaymentProfiles.IsUnknown() {
			data.PaymentProfiles = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"attributes":           types.ObjectType{AttrTypes: map[string]attr.Type{"agreement_number": types.StringType, "contract_sum": types.Int64Type, "end_date": types.StringType}},
				"is_active":            types.BoolType,
				"name":                 types.StringType,
				"organization":         types.StringType,
				"organization_uuid":    types.StringType,
				"payment_type":         types.StringType,
				"payment_type_display": types.StringType,
				"url":                  types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["phone_number"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.PhoneNumber = types.StringValue(str)
		}
	} else {
		if data.PhoneNumber.IsUnknown() {
			data.PhoneNumber = types.StringNull()
		}
	}
	if val, ok := sourceMap["postal"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Postal = types.StringValue(str)
		}
	} else {
		if data.Postal.IsUnknown() {
			data.Postal = types.StringNull()
		}
	}
	if val, ok := sourceMap["project_metadata_checklist"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProjectMetadataChecklist = types.StringValue(str)
		}
	} else {
		if data.ProjectMetadataChecklist.IsUnknown() {
			data.ProjectMetadataChecklist = types.StringNull()
		}
	}
	if val, ok := sourceMap["projects"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"end_date":       types.StringType,
						"image":          types.StringType,
						"name":           types.StringType,
						"resource_count": types.Int64Type,
						"url":            types.StringType,
					}
					attrValues := map[string]attr.Value{
						"end_date": func() attr.Value {
							if v, ok := objMap["end_date"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"image": func() attr.Value {
							if v, ok := objMap["image"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"resource_count": func() attr.Value {
							if v, ok := objMap["resource_count"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"end_date":       types.StringType,
				"image":          types.StringType,
				"name":           types.StringType,
				"resource_count": types.Int64Type,
				"url":            types.StringType,
			}}, items)
			data.Projects = listVal
		}
	} else {
		if data.Projects.IsUnknown() {
			data.Projects = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"end_date":       types.StringType,
				"image":          types.StringType,
				"name":           types.StringType,
				"resource_count": types.Int64Type,
				"url":            types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["projects_count"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.ProjectsCount = types.Int64Value(int64(num))
		}
	} else {
		if data.ProjectsCount.IsUnknown() {
			data.ProjectsCount = types.Int64Null()
		}
	}
	if val, ok := sourceMap["registration_code"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.RegistrationCode = types.StringValue(str)
		}
	} else {
		if data.RegistrationCode.IsUnknown() {
			data.RegistrationCode = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_provider"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceProvider = types.StringValue(str)
		}
	} else {
		if data.ServiceProvider.IsUnknown() {
			data.ServiceProvider = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_provider_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceProviderUuid = types.StringValue(str)
		}
	} else {
		if data.ServiceProviderUuid.IsUnknown() {
			data.ServiceProviderUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["slug"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Slug = types.StringValue(str)
		}
	} else {
		if data.Slug.IsUnknown() {
			data.Slug = types.StringNull()
		}
	}
	if val, ok := sourceMap["sponsor_number"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.SponsorNumber = types.Int64Value(int64(num))
		}
	} else {
		if data.SponsorNumber.IsUnknown() {
			data.SponsorNumber = types.Int64Null()
		}
	}
	if val, ok := sourceMap["url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Url = types.StringValue(str)
		}
	} else {
		if data.Url.IsUnknown() {
			data.Url = types.StringNull()
		}
	}
	if val, ok := sourceMap["users_count"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.UsersCount = types.Int64Value(int64(num))
		}
	} else {
		if data.UsersCount.IsUnknown() {
			data.UsersCount = types.Int64Null()
		}
	}
	if val, ok := sourceMap["vat_code"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.VatCode = types.StringValue(str)
		}
	} else {
		if data.VatCode.IsUnknown() {
			data.VatCode = types.StringNull()
		}
	}

	// Map filter parameters from response if available

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StructureCustomerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data StructureCustomerResourceModel
	var state StructureCustomerResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	// Read current state to get the UUID (which is computed and not in plan)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Use UUID from state
	data.UUID = state.UUID
	// Prepare request body
	requestBody := map[string]interface{}{}
	if !data.Abbreviation.IsNull() && !data.Abbreviation.IsUnknown() {
		if v := data.Abbreviation.ValueString(); v != "" {
			requestBody["abbreviation"] = v
		}
	}
	if !data.AccessSubnets.IsNull() && !data.AccessSubnets.IsUnknown() {
		if v := data.AccessSubnets.ValueString(); v != "" {
			requestBody["access_subnets"] = v
		}
	}
	if !data.AccountingStartDate.IsNull() && !data.AccountingStartDate.IsUnknown() {
		if v := data.AccountingStartDate.ValueString(); v != "" {
			requestBody["accounting_start_date"] = v
		}
	}
	if !data.Address.IsNull() && !data.Address.IsUnknown() {
		if v := data.Address.ValueString(); v != "" {
			requestBody["address"] = v
		}
	}
	if !data.AgreementNumber.IsNull() && !data.AgreementNumber.IsUnknown() {
		if v := data.AgreementNumber.ValueString(); v != "" {
			requestBody["agreement_number"] = v
		}
	}
	if !data.Archived.IsNull() && !data.Archived.IsUnknown() {
		requestBody["archived"] = data.Archived.ValueBool()
	}
	if !data.BackendId.IsNull() && !data.BackendId.IsUnknown() {
		if v := data.BackendId.ValueString(); v != "" {
			requestBody["backend_id"] = v
		}
	}
	if !data.BankAccount.IsNull() && !data.BankAccount.IsUnknown() {
		if v := data.BankAccount.ValueString(); v != "" {
			requestBody["bank_account"] = v
		}
	}
	if !data.BankName.IsNull() && !data.BankName.IsUnknown() {
		if v := data.BankName.ValueString(); v != "" {
			requestBody["bank_name"] = v
		}
	}
	if !data.Blocked.IsNull() && !data.Blocked.IsUnknown() {
		requestBody["blocked"] = data.Blocked.ValueBool()
	}
	if !data.ContactDetails.IsNull() && !data.ContactDetails.IsUnknown() {
		if v := data.ContactDetails.ValueString(); v != "" {
			requestBody["contact_details"] = v
		}
	}
	if !data.Country.IsNull() && !data.Country.IsUnknown() {
		if v := data.Country.ValueString(); v != "" {
			requestBody["country"] = v
		}
	}
	if !data.DefaultTaxPercent.IsNull() && !data.DefaultTaxPercent.IsUnknown() {
		if v := data.DefaultTaxPercent.ValueString(); v != "" {
			requestBody["default_tax_percent"] = v
		}
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		if v := data.Description.ValueString(); v != "" {
			requestBody["description"] = v
		}
	}
	if !data.DisplayBillingInfoInProjects.IsNull() && !data.DisplayBillingInfoInProjects.IsUnknown() {
		requestBody["display_billing_info_in_projects"] = data.DisplayBillingInfoInProjects.ValueBool()
	}
	if !data.Domain.IsNull() && !data.Domain.IsUnknown() {
		if v := data.Domain.ValueString(); v != "" {
			requestBody["domain"] = v
		}
	}
	if !data.Email.IsNull() && !data.Email.IsUnknown() {
		if v := data.Email.ValueString(); v != "" {
			requestBody["email"] = v
		}
	}
	if !data.GracePeriodDays.IsNull() && !data.GracePeriodDays.IsUnknown() {
		requestBody["grace_period_days"] = data.GracePeriodDays.ValueInt64()
	}
	if !data.Homepage.IsNull() && !data.Homepage.IsUnknown() {
		if v := data.Homepage.ValueString(); v != "" {
			requestBody["homepage"] = v
		}
	}
	if !data.Image.IsNull() && !data.Image.IsUnknown() {
		if v := data.Image.ValueString(); v != "" {
			requestBody["image"] = v
		}
	}
	if !data.Latitude.IsNull() && !data.Latitude.IsUnknown() {
		requestBody["latitude"] = data.Latitude.ValueFloat64()
	}
	if !data.Longitude.IsNull() && !data.Longitude.IsUnknown() {
		requestBody["longitude"] = data.Longitude.ValueFloat64()
	}
	if !data.MaxServiceAccounts.IsNull() && !data.MaxServiceAccounts.IsUnknown() {
		requestBody["max_service_accounts"] = data.MaxServiceAccounts.ValueInt64()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		if v := data.Name.ValueString(); v != "" {
			requestBody["name"] = v
		}
	}
	if !data.NativeName.IsNull() && !data.NativeName.IsUnknown() {
		if v := data.NativeName.ValueString(); v != "" {
			requestBody["native_name"] = v
		}
	}
	if !data.NotificationEmails.IsNull() && !data.NotificationEmails.IsUnknown() {
		if v := data.NotificationEmails.ValueString(); v != "" {
			requestBody["notification_emails"] = v
		}
	}
	if !data.PhoneNumber.IsNull() && !data.PhoneNumber.IsUnknown() {
		if v := data.PhoneNumber.ValueString(); v != "" {
			requestBody["phone_number"] = v
		}
	}
	if !data.Postal.IsNull() && !data.Postal.IsUnknown() {
		if v := data.Postal.ValueString(); v != "" {
			requestBody["postal"] = v
		}
	}
	if !data.ProjectMetadataChecklist.IsNull() && !data.ProjectMetadataChecklist.IsUnknown() {
		if v := data.ProjectMetadataChecklist.ValueString(); v != "" {
			requestBody["project_metadata_checklist"] = v
		}
	}
	if !data.RegistrationCode.IsNull() && !data.RegistrationCode.IsUnknown() {
		if v := data.RegistrationCode.ValueString(); v != "" {
			requestBody["registration_code"] = v
		}
	}
	if !data.Slug.IsNull() && !data.Slug.IsUnknown() {
		if v := data.Slug.ValueString(); v != "" {
			requestBody["slug"] = v
		}
	}
	if !data.SponsorNumber.IsNull() && !data.SponsorNumber.IsUnknown() {
		requestBody["sponsor_number"] = data.SponsorNumber.ValueInt64()
	}
	if !data.VatCode.IsNull() && !data.VatCode.IsUnknown() {
		if v := data.VatCode.ValueString(); v != "" {
			requestBody["vat_code"] = v
		}
	}

	// Call Waldur API to update resource
	var result map[string]interface{}

	err := r.client.Update(ctx, "/api/customers/{uuid}/", data.UUID.ValueString(), requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update StructureCustomer",
			"An error occurred while updating the structure_customer: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
	// Map response fields to data model
	_ = sourceMap
	if val, ok := sourceMap["abbreviation"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Abbreviation = types.StringValue(str)
		}
	} else {
		if data.Abbreviation.IsUnknown() {
			data.Abbreviation = types.StringNull()
		}
	}
	if val, ok := sourceMap["access_subnets"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AccessSubnets = types.StringValue(str)
		}
	} else {
		if data.AccessSubnets.IsUnknown() {
			data.AccessSubnets = types.StringNull()
		}
	}
	if val, ok := sourceMap["accounting_start_date"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AccountingStartDate = types.StringValue(str)
		}
	} else {
		if data.AccountingStartDate.IsUnknown() {
			data.AccountingStartDate = types.StringNull()
		}
	}
	if val, ok := sourceMap["address"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Address = types.StringValue(str)
		}
	} else {
		if data.Address.IsUnknown() {
			data.Address = types.StringNull()
		}
	}
	if val, ok := sourceMap["agreement_number"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AgreementNumber = types.StringValue(str)
		}
	} else {
		if data.AgreementNumber.IsUnknown() {
			data.AgreementNumber = types.StringNull()
		}
	}
	if val, ok := sourceMap["archived"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Archived = types.BoolValue(b)
		}
	} else {
		if data.Archived.IsUnknown() {
			data.Archived = types.BoolNull()
		}
	}
	if val, ok := sourceMap["backend_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
		}
	} else {
		if data.BackendId.IsUnknown() {
			data.BackendId = types.StringNull()
		}
	}
	if val, ok := sourceMap["bank_account"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BankAccount = types.StringValue(str)
		}
	} else {
		if data.BankAccount.IsUnknown() {
			data.BankAccount = types.StringNull()
		}
	}
	if val, ok := sourceMap["bank_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BankName = types.StringValue(str)
		}
	} else {
		if data.BankName.IsUnknown() {
			data.BankName = types.StringNull()
		}
	}
	if val, ok := sourceMap["blocked"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Blocked = types.BoolValue(b)
		}
	} else {
		if data.Blocked.IsUnknown() {
			data.Blocked = types.BoolNull()
		}
	}
	if val, ok := sourceMap["call_managing_organization_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CallManagingOrganizationUuid = types.StringValue(str)
		}
	} else {
		if data.CallManagingOrganizationUuid.IsUnknown() {
			data.CallManagingOrganizationUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["contact_details"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ContactDetails = types.StringValue(str)
		}
	} else {
		if data.ContactDetails.IsUnknown() {
			data.ContactDetails = types.StringNull()
		}
	}
	if val, ok := sourceMap["country"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Country = types.StringValue(str)
		}
	} else {
		if data.Country.IsUnknown() {
			data.Country = types.StringNull()
		}
	}
	if val, ok := sourceMap["country_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CountryName = types.StringValue(str)
		}
	} else {
		if data.CountryName.IsUnknown() {
			data.CountryName = types.StringNull()
		}
	}
	if val, ok := sourceMap["created"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
	} else {
		if data.Created.IsUnknown() {
			data.Created = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_credit"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.CustomerCredit = types.Float64Value(num)
		}
	} else {
		if data.CustomerCredit.IsUnknown() {
			data.CustomerCredit = types.Float64Null()
		}
	}
	if val, ok := sourceMap["customer_unallocated_credit"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.CustomerUnallocatedCredit = types.Float64Value(num)
		}
	} else {
		if data.CustomerUnallocatedCredit.IsUnknown() {
			data.CustomerUnallocatedCredit = types.Float64Null()
		}
	}
	if val, ok := sourceMap["default_tax_percent"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.DefaultTaxPercent = types.StringValue(str)
		}
	} else {
		if data.DefaultTaxPercent.IsUnknown() {
			data.DefaultTaxPercent = types.StringNull()
		}
	}
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
		}
	}
	if val, ok := sourceMap["display_billing_info_in_projects"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.DisplayBillingInfoInProjects = types.BoolValue(b)
		}
	} else {
		if data.DisplayBillingInfoInProjects.IsUnknown() {
			data.DisplayBillingInfoInProjects = types.BoolNull()
		}
	}
	if val, ok := sourceMap["display_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.DisplayName = types.StringValue(str)
		}
	} else {
		if data.DisplayName.IsUnknown() {
			data.DisplayName = types.StringNull()
		}
	}
	if val, ok := sourceMap["domain"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Domain = types.StringValue(str)
		}
	} else {
		if data.Domain.IsUnknown() {
			data.Domain = types.StringNull()
		}
	}
	if val, ok := sourceMap["email"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Email = types.StringValue(str)
		}
	} else {
		if data.Email.IsUnknown() {
			data.Email = types.StringNull()
		}
	}
	if val, ok := sourceMap["grace_period_days"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.GracePeriodDays = types.Int64Value(int64(num))
		}
	} else {
		if data.GracePeriodDays.IsUnknown() {
			data.GracePeriodDays = types.Int64Null()
		}
	}
	if val, ok := sourceMap["homepage"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Homepage = types.StringValue(str)
		}
	} else {
		if data.Homepage.IsUnknown() {
			data.Homepage = types.StringNull()
		}
	}
	if val, ok := sourceMap["image"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Image = types.StringValue(str)
		}
	} else {
		if data.Image.IsUnknown() {
			data.Image = types.StringNull()
		}
	}
	if val, ok := sourceMap["is_service_provider"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.IsServiceProvider = types.BoolValue(b)
		}
	} else {
		if data.IsServiceProvider.IsUnknown() {
			data.IsServiceProvider = types.BoolNull()
		}
	}
	if val, ok := sourceMap["latitude"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Latitude = types.Float64Value(num)
		}
	} else {
		if data.Latitude.IsUnknown() {
			data.Latitude = types.Float64Null()
		}
	}
	if val, ok := sourceMap["longitude"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Longitude = types.Float64Value(num)
		}
	} else {
		if data.Longitude.IsUnknown() {
			data.Longitude = types.Float64Null()
		}
	}
	if val, ok := sourceMap["max_service_accounts"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.MaxServiceAccounts = types.Int64Value(int64(num))
		}
	} else {
		if data.MaxServiceAccounts.IsUnknown() {
			data.MaxServiceAccounts = types.Int64Null()
		}
	}
	if val, ok := sourceMap["native_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.NativeName = types.StringValue(str)
		}
	} else {
		if data.NativeName.IsUnknown() {
			data.NativeName = types.StringNull()
		}
	}
	if val, ok := sourceMap["notification_emails"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.NotificationEmails = types.StringValue(str)
		}
	} else {
		if data.NotificationEmails.IsUnknown() {
			data.NotificationEmails = types.StringNull()
		}
	}
	if val, ok := sourceMap["organization_groups"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"customers_count": types.Int64Type,
						"name":            types.StringType,
						"parent":          types.StringType,
						"parent_name":     types.StringType,
						"parent_uuid":     types.StringType,
						"url":             types.StringType,
					}
					attrValues := map[string]attr.Value{
						"customers_count": func() attr.Value {
							if v, ok := objMap["customers_count"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"parent": func() attr.Value {
							if v, ok := objMap["parent"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"parent_name": func() attr.Value {
							if v, ok := objMap["parent_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"parent_uuid": func() attr.Value {
							if v, ok := objMap["parent_uuid"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"customers_count": types.Int64Type,
				"name":            types.StringType,
				"parent":          types.StringType,
				"parent_name":     types.StringType,
				"parent_uuid":     types.StringType,
				"url":             types.StringType,
			}}, items)
			data.OrganizationGroups = listVal
		}
	} else {
		if data.OrganizationGroups.IsUnknown() {
			data.OrganizationGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"customers_count": types.Int64Type,
				"name":            types.StringType,
				"parent":          types.StringType,
				"parent_name":     types.StringType,
				"parent_uuid":     types.StringType,
				"url":             types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["payment_profiles"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"attributes":           types.ObjectType{AttrTypes: map[string]attr.Type{"agreement_number": types.StringType, "contract_sum": types.Int64Type, "end_date": types.StringType}},
						"is_active":            types.BoolType,
						"name":                 types.StringType,
						"organization":         types.StringType,
						"organization_uuid":    types.StringType,
						"payment_type":         types.StringType,
						"payment_type_display": types.StringType,
						"url":                  types.StringType,
					}
					attrValues := map[string]attr.Value{
						"attributes": types.ObjectNull(types.ObjectType{AttrTypes: map[string]attr.Type{"agreement_number": types.StringType, "contract_sum": types.Int64Type, "end_date": types.StringType}}.AttrTypes),
						"is_active": func() attr.Value {
							if v, ok := objMap["is_active"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"organization": func() attr.Value {
							if v, ok := objMap["organization"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"organization_uuid": func() attr.Value {
							if v, ok := objMap["organization_uuid"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"payment_type": func() attr.Value {
							if v, ok := objMap["payment_type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"payment_type_display": func() attr.Value {
							if v, ok := objMap["payment_type_display"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"attributes":           types.ObjectType{AttrTypes: map[string]attr.Type{"agreement_number": types.StringType, "contract_sum": types.Int64Type, "end_date": types.StringType}},
				"is_active":            types.BoolType,
				"name":                 types.StringType,
				"organization":         types.StringType,
				"organization_uuid":    types.StringType,
				"payment_type":         types.StringType,
				"payment_type_display": types.StringType,
				"url":                  types.StringType,
			}}, items)
			data.PaymentProfiles = listVal
		}
	} else {
		if data.PaymentProfiles.IsUnknown() {
			data.PaymentProfiles = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"attributes":           types.ObjectType{AttrTypes: map[string]attr.Type{"agreement_number": types.StringType, "contract_sum": types.Int64Type, "end_date": types.StringType}},
				"is_active":            types.BoolType,
				"name":                 types.StringType,
				"organization":         types.StringType,
				"organization_uuid":    types.StringType,
				"payment_type":         types.StringType,
				"payment_type_display": types.StringType,
				"url":                  types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["phone_number"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.PhoneNumber = types.StringValue(str)
		}
	} else {
		if data.PhoneNumber.IsUnknown() {
			data.PhoneNumber = types.StringNull()
		}
	}
	if val, ok := sourceMap["postal"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Postal = types.StringValue(str)
		}
	} else {
		if data.Postal.IsUnknown() {
			data.Postal = types.StringNull()
		}
	}
	if val, ok := sourceMap["project_metadata_checklist"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProjectMetadataChecklist = types.StringValue(str)
		}
	} else {
		if data.ProjectMetadataChecklist.IsUnknown() {
			data.ProjectMetadataChecklist = types.StringNull()
		}
	}
	if val, ok := sourceMap["projects"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"end_date":       types.StringType,
						"image":          types.StringType,
						"name":           types.StringType,
						"resource_count": types.Int64Type,
						"url":            types.StringType,
					}
					attrValues := map[string]attr.Value{
						"end_date": func() attr.Value {
							if v, ok := objMap["end_date"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"image": func() attr.Value {
							if v, ok := objMap["image"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"resource_count": func() attr.Value {
							if v, ok := objMap["resource_count"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"end_date":       types.StringType,
				"image":          types.StringType,
				"name":           types.StringType,
				"resource_count": types.Int64Type,
				"url":            types.StringType,
			}}, items)
			data.Projects = listVal
		}
	} else {
		if data.Projects.IsUnknown() {
			data.Projects = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"end_date":       types.StringType,
				"image":          types.StringType,
				"name":           types.StringType,
				"resource_count": types.Int64Type,
				"url":            types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["projects_count"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.ProjectsCount = types.Int64Value(int64(num))
		}
	} else {
		if data.ProjectsCount.IsUnknown() {
			data.ProjectsCount = types.Int64Null()
		}
	}
	if val, ok := sourceMap["registration_code"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.RegistrationCode = types.StringValue(str)
		}
	} else {
		if data.RegistrationCode.IsUnknown() {
			data.RegistrationCode = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_provider"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceProvider = types.StringValue(str)
		}
	} else {
		if data.ServiceProvider.IsUnknown() {
			data.ServiceProvider = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_provider_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceProviderUuid = types.StringValue(str)
		}
	} else {
		if data.ServiceProviderUuid.IsUnknown() {
			data.ServiceProviderUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["slug"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Slug = types.StringValue(str)
		}
	} else {
		if data.Slug.IsUnknown() {
			data.Slug = types.StringNull()
		}
	}
	if val, ok := sourceMap["sponsor_number"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.SponsorNumber = types.Int64Value(int64(num))
		}
	} else {
		if data.SponsorNumber.IsUnknown() {
			data.SponsorNumber = types.Int64Null()
		}
	}
	if val, ok := sourceMap["url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Url = types.StringValue(str)
		}
	} else {
		if data.Url.IsUnknown() {
			data.Url = types.StringNull()
		}
	}
	if val, ok := sourceMap["users_count"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.UsersCount = types.Int64Value(int64(num))
		}
	} else {
		if data.UsersCount.IsUnknown() {
			data.UsersCount = types.Int64Null()
		}
	}
	if val, ok := sourceMap["vat_code"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.VatCode = types.StringValue(str)
		}
	} else {
		if data.VatCode.IsUnknown() {
			data.VatCode = types.StringNull()
		}
	}

	// Map filter parameters from response if available

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StructureCustomerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data StructureCustomerResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Call Waldur API to delete resource
	err := r.client.DeleteByUUID(ctx, "/api/customers/{uuid}/", data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete StructureCustomer",
			"An error occurred while deleting the structure_customer: "+err.Error(),
		)
		return
	}
}

func (r *StructureCustomerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
