package customer

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &StructureCustomerDataSource{}

func NewStructureCustomerDataSource() datasource.DataSource {
	return &StructureCustomerDataSource{}
}

type StructureCustomerDataSource struct {
	client *StructureCustomerClient
}

type StructureCustomerDataSourceModel struct {
	StructureCustomerModel
	Filters *StructureCustomerFiltersModel `tfsdk:"filters"`
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
				MarkdownDescription: "Structure Customer UUID",
			},
			"filters": (&StructureCustomerFiltersModel{}).GetSchema(),
			"abbreviation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Abbreviation",
			},
			"access_subnets": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Enter a comma separated list of IPv4 or IPv6 CIDR addresses from where connection to self-service is allowed.",
			},
			"accounting_start_date": schema.StringAttribute{
				CustomType:          timetypes.RFC3339Type{},
				Computed:            true,
				MarkdownDescription: "Accounting Start Date",
			},
			"address": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Address",
			},
			"agreement_number": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Agreement Number",
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
				MarkdownDescription: "Bank Account",
			},
			"bank_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Bank Name",
			},
			"billing_price_estimate": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"current": schema.Float64Attribute{
						Computed:            true,
						MarkdownDescription: "Current",
					},
					"tax": schema.Float64Attribute{
						Computed:            true,
						MarkdownDescription: "Tax",
					},
					"tax_current": schema.Float64Attribute{
						Computed:            true,
						MarkdownDescription: "Tax Current",
					},
					"total": schema.Float64Attribute{
						Computed:            true,
						MarkdownDescription: "Total",
					},
				},
				Computed:            true,
				MarkdownDescription: "Billing Price Estimate",
			},
			"blocked": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Blocked",
			},
			"call_managing_organization_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Call Managing Organization Uuid",
			},
			"contact_details": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Contact Details",
			},
			"country": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Country code (ISO 3166-1 alpha-2)",
			},
			"country_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Human-readable country name",
			},
			"customer_credit": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "Customer Credit",
			},
			"customer_unallocated_credit": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "Customer Unallocated Credit",
			},
			"default_tax_percent": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Default Tax Percent",

				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^-?\d{0,3}(?:\.\d{0,2})?$`), ""),
				},
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Description",
			},
			"display_billing_info_in_projects": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Display Billing Info In Projects",
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

				Validators: []validator.Int64{
					int64validator.AtLeast(0),
					int64validator.AtMost(2147483647),
				},
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
				MarkdownDescription: "Is Service Provider",
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

				Validators: []validator.Int64{
					int64validator.AtLeast(0),
					int64validator.AtMost(32767),
				},
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name",
			},
			"native_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Native Name",
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
							Computed:            true,
							MarkdownDescription: "Name",
						},
						"parent": schema.StringAttribute{
							Computed:            true,
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
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Uuid",
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
									Computed:            true,
									MarkdownDescription: "Agreement Number",
								},
								"contract_sum": schema.Int64Attribute{
									Computed:            true,
									MarkdownDescription: "Contract Sum",
								},
								"end_date": schema.StringAttribute{
									Computed:            true,
									MarkdownDescription: "End Date",
								},
							},
							Computed:            true,
							MarkdownDescription: "Attributes",
						},
						"is_active": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Is Active",
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name",
						},
						"organization": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Organization",
						},
						"organization_uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Organization Uuid",
						},
						"payment_type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Payment Type",

							Validators: []validator.String{
								stringvalidator.OneOf("fixed_price", "invoices", "payment_gw_monthly"),
							},
						},
						"payment_type_display": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Payment Type Display",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Url",
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Uuid",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Payment Profiles",
			},
			"phone_number": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Phone Number",
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
				MarkdownDescription: "Registration Code",
			},
			"service_provider": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Service Provider",
			},
			"service_provider_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Service Provider Uuid",
			},
			"slug": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",

				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
				},
			},
			"sponsor_number": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "External ID of the sponsor covering the costs",

				Validators: []validator.Int64{
					int64validator.AtLeast(0),
					int64validator.AtMost(2147483647),
				},
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

	d.client = &StructureCustomerClient{}
	if err := d.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			err.Error(),
		)
		return
	}
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
		apiResp, err := d.client.Get(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Structure Customer",
				"An error occurred while reading the Structure Customer by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup structure_customer.",
			)
			return
		}

		results, err := d.client.List(ctx, filters)
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

		resp.Diagnostics.Append(data.CopyFrom(ctx, results[0])...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
