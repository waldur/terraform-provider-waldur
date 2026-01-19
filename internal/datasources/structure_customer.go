package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
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
	Projects                     types.List    `tfsdk:"projects"`
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
		MarkdownDescription: "StructureCustomer data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"abbreviation": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"agreement_number": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"archived": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"contact_details": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"native_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"organization_group_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"organization_group_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "organization_group_uuid",
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
				MarkdownDescription: " ",
			},
			"access_subnets": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Enter a comma separated list of IPv4 or IPv6 CIDR addresses from where connection to self-service is allowed.",
			},
			"accounting_start_date": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"address": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"bank_account": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"bank_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"blocked": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"call_managing_organization_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"country": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"country_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"customer_credit": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"customer_unallocated_credit": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"default_tax_percent": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"display_billing_info_in_projects": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"display_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"domain": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"email": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"grace_period_days": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of extra days after project end date before resources are terminated",
			},
			"homepage": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"image": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"is_service_provider": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"latitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"longitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
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
				CustomType:          types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"customers_count": types.Int64Type, "name": types.StringType, "parent": types.StringType, "parent_name": types.StringType, "parent_uuid": types.StringType, "url": types.StringType}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"payment_profiles": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"attributes": types.ObjectType{AttrTypes: map[string]attr.Type{"agreement_number": types.StringType, "contract_sum": types.Int64Type, "end_date": types.StringType}}, "is_active": types.BoolType, "name": types.StringType, "organization": types.StringType, "organization_uuid": types.StringType, "payment_type": types.StringType, "payment_type_display": types.StringType, "url": types.StringType}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"phone_number": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"postal": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"project_metadata_checklist": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"projects": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"end_date": types.StringType, "image": types.StringType, "name": types.StringType, "resource_count": types.Int64Type, "url": types.StringType}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"projects_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_provider": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_provider_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"sponsor_number": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "External ID of the sponsor covering the costs",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"users_count": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: " ",
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
		var item map[string]interface{}

		err := d.client.GetByUUID(ctx, "/api/customers/{uuid}/", data.UUID.ValueString(), &item)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Customer",
				"An error occurred while reading the structure_customer by UUID: "+err.Error(),
			)
			return
		}

		// Extract data from single result
		if uuid, ok := item["uuid"].(string); ok {
			data.UUID = types.StringValue(uuid)
		}

		sourceMap := item
		// Map response fields to data model
		_ = sourceMap
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
		if val, ok := sourceMap["abbreviation"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Abbreviation = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["agreement_number"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AgreementNumber = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["archived"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.Archived = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["backend_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BackendId = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["contact_details"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ContactDetails = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Name = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name_exact"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NameExact = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["native_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NativeName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["organization_group_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OrganizationGroupName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["organization_group_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OrganizationGroupUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["owned_by_current_user"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.OwnedByCurrentUser = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["query"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Query = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["registration_code"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.RegistrationCode = types.StringValue(str)
			}
		}

	} else {
		// Filter by provided parameters
		var results []map[string]interface{}

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
				"Unable to List Customer",
				"An error occurred while filtering structure_customer: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Customer Not Found",
				"No structure_customer found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Customers Found",
				fmt.Sprintf("Found %d structure_customers with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		// Extract data from the single result
		if uuid, ok := results[0]["uuid"].(string); ok {
			data.UUID = types.StringValue(uuid)
		}
		sourceMap := results[0]
		// Map response fields to data model
		_ = sourceMap
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
		if val, ok := sourceMap["abbreviation"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Abbreviation = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["agreement_number"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AgreementNumber = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["archived"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.Archived = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["backend_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BackendId = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["contact_details"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ContactDetails = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Name = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name_exact"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NameExact = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["native_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NativeName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["organization_group_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OrganizationGroupName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["organization_group_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OrganizationGroupUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["owned_by_current_user"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.OwnedByCurrentUser = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["query"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Query = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["registration_code"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.RegistrationCode = types.StringValue(str)
			}
		}
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
