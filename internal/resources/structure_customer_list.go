package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-provider-waldur/internal/client"
)

var _ list.ListResource = &StructureCustomerList{}

type StructureCustomerList struct {
	client *client.Client
}

func NewStructureCustomerList() list.ListResource {
	return &StructureCustomerList{}
}

func (l *StructureCustomerList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_structure_customer"
}

func (l *StructureCustomerList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"abbreviation": schema.StringAttribute{
				Description: "Abbreviation",
				Optional:    true,
			},
			"agreement_number": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"archived": schema.BoolAttribute{
				Description: "",
				Optional:    true,
			},
			"backend_id": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"contact_details": schema.StringAttribute{
				Description: "Contact details",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name",
				Optional:    true,
			},
			"name_exact": schema.StringAttribute{
				Description: "Name (exact)",
				Optional:    true,
			},
			"native_name": schema.StringAttribute{
				Description: "Native name",
				Optional:    true,
			},
			"o": schema.StringAttribute{
				Description: "Which field to use when ordering the results.",
				Optional:    true,
			},
			"organization_group_name": schema.StringAttribute{
				Description: "Organization group name",
				Optional:    true,
			},
			"owned_by_current_user": schema.BoolAttribute{
				Description: "Return a list of customers where current user is owner.",
				Optional:    true,
			},
			"page": schema.Int64Attribute{
				Description: "A page number within the paginated result set.",
				Optional:    true,
			},
			"page_size": schema.Int64Attribute{
				Description: "Number of results to return per page.",
				Optional:    true,
			},
			"query": schema.StringAttribute{
				Description: "Filter by name, native name, abbreviation, domain, UUID, registration code or agreement number",
				Optional:    true,
			},
			"registration_code": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
		},
	}
}

func (l *StructureCustomerList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected List Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T", req.ProviderData),
		)
		return
	}

	l.client = client
}

type StructureCustomerListModel struct {
	Abbreviation          types.String `tfsdk:"abbreviation"`
	AgreementNumber       types.String `tfsdk:"agreement_number"`
	Archived              types.Bool   `tfsdk:"archived"`
	BackendId             types.String `tfsdk:"backend_id"`
	ContactDetails        types.String `tfsdk:"contact_details"`
	Name                  types.String `tfsdk:"name"`
	NameExact             types.String `tfsdk:"name_exact"`
	NativeName            types.String `tfsdk:"native_name"`
	O                     types.String `tfsdk:"o"`
	OrganizationGroupName types.String `tfsdk:"organization_group_name"`
	OwnedByCurrentUser    types.Bool   `tfsdk:"owned_by_current_user"`
	Page                  types.Int64  `tfsdk:"page"`
	PageSize              types.Int64  `tfsdk:"page_size"`
	Query                 types.String `tfsdk:"query"`
	RegistrationCode      types.String `tfsdk:"registration_code"`
}

func (l *StructureCustomerList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config StructureCustomerListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := make(map[string]string)
	if !config.Abbreviation.IsNull() && !config.Abbreviation.IsUnknown() {
		filters["abbreviation"] = config.Abbreviation.ValueString()
	}
	if !config.AgreementNumber.IsNull() && !config.AgreementNumber.IsUnknown() {
		filters["agreement_number"] = config.AgreementNumber.ValueString()
	}
	if !config.Archived.IsNull() && !config.Archived.IsUnknown() {
		filters["archived"] = fmt.Sprintf("%t", config.Archived.ValueBool())
	}
	if !config.BackendId.IsNull() && !config.BackendId.IsUnknown() {
		filters["backend_id"] = config.BackendId.ValueString()
	}
	if !config.ContactDetails.IsNull() && !config.ContactDetails.IsUnknown() {
		filters["contact_details"] = config.ContactDetails.ValueString()
	}
	if !config.Name.IsNull() && !config.Name.IsUnknown() {
		filters["name"] = config.Name.ValueString()
	}
	if !config.NameExact.IsNull() && !config.NameExact.IsUnknown() {
		filters["name_exact"] = config.NameExact.ValueString()
	}
	if !config.NativeName.IsNull() && !config.NativeName.IsUnknown() {
		filters["native_name"] = config.NativeName.ValueString()
	}
	if !config.O.IsNull() && !config.O.IsUnknown() {
		filters["o"] = config.O.ValueString()
	}
	if !config.OrganizationGroupName.IsNull() && !config.OrganizationGroupName.IsUnknown() {
		filters["organization_group_name"] = config.OrganizationGroupName.ValueString()
	}
	if !config.OwnedByCurrentUser.IsNull() && !config.OwnedByCurrentUser.IsUnknown() {
		filters["owned_by_current_user"] = fmt.Sprintf("%t", config.OwnedByCurrentUser.ValueBool())
	}
	if !config.Page.IsNull() && !config.Page.IsUnknown() {
		filters["page"] = fmt.Sprintf("%d", config.Page.ValueInt64())
	}
	if !config.PageSize.IsNull() && !config.PageSize.IsUnknown() {
		filters["page_size"] = fmt.Sprintf("%d", config.PageSize.ValueInt64())
	}
	if !config.Query.IsNull() && !config.Query.IsUnknown() {
		filters["query"] = config.Query.ValueString()
	}
	if !config.RegistrationCode.IsNull() && !config.RegistrationCode.IsUnknown() {
		filters["registration_code"] = config.RegistrationCode.ValueString()
	}

	// Call API
	var listResult []StructureCustomerApiResponse
	err := l.client.ListWithFilter(ctx, "/api/customers/", filters, &listResult)
	if err != nil {
		// Return error diagnostics
		resp.AddError("Failed to list resources", err.Error())
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Stream results
	stream.Results = func(push func(list.ListResult) bool) {
		for _, apiResp := range listResult {
			result := req.NewListResult(ctx)

			// Map item to model
			var data StructureCustomerResourceModel
			model := &data

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

			// Set the resource state
			// For ListResource, we generally return the "Resource" state matching the main resource schema.
			// This allows `terraform plan` to correlate if using `terraform query`.

			setDiags := result.Resource.Set(ctx, &data)
			diags.Append(setDiags...)
			result.Diagnostics.Append(diags...)

			// Set Identity if possible (usually UUID)
			if !data.UUID.IsNull() && !data.UUID.IsUnknown() {
				result.Diagnostics.Append(result.Identity.Set(ctx, data.UUID.ValueString())...)
			}

			if !push(result) {
				return
			}
		}
	}
}
