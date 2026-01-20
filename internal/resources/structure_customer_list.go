package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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
		// Filter parameters can be added here if needed
		Attributes: map[string]schema.Attribute{},
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
	// Add filter fields here if added to schema
}

func (l *StructureCustomerList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config StructureCustomerListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Call API
	var listResult []map[string]interface{}
	err := l.client.List(ctx, "/api/customers/", &listResult)
	if err != nil {
		// Return error diagnostics
		resp.AddError("Failed to list resources", err.Error())
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Stream results
	stream.Results = func(push func(list.ListResult) bool) {
		for _, item := range listResult {
			result := req.NewListResult(ctx)

			// Map item to model
			var data StructureCustomerResourceModel

			// Initialize lists and objects to empty/null if needed (or rely on map logic)

			sourceMap := item

			// Reuse the mapResponseFields logic (embedded here or via shared template)
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
							}
							attrValues := map[string]attr.Value{
								"attributes": types.ObjectNull(types.ObjectType{AttrTypes: map[string]attr.Type{
									"agreement_number": types.StringType,
									"contract_sum":     types.Int64Type,
									"end_date":         types.StringType,
								}}.AttrTypes),
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
					}}, items)
					data.PaymentProfiles = listVal
				}
			} else {
				if data.PaymentProfiles.IsUnknown() {
					data.PaymentProfiles = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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

			// Set the resource state
			// For ListResource, we generally return the "Resource" state matching the main resource schema.
			// This allows `terraform plan` to correlate if using `terraform query`.

			diags := result.Resource.Set(ctx, &data)
			result.Diagnostics.Append(diags...)

			// Set Identity if possible (usually UUID)
			if !data.UUID.IsNull() && !data.UUID.IsUnknown() {
				// Identity value must match what the resource uses for Import?
				// Typically implicit. For now just setting Resource is key.
				// result.Identity.Set(ctx, data.UUID.ValueString())
				// The doc says: "An error is returned if a list result in the stream contains a null identity"
				result.Diagnostics.Append(result.Identity.Set(ctx, data.UUID.ValueString())...)
			} else {
				// Try to fallback to "uuid" from map if model failed
				if uuid, ok := item["uuid"].(string); ok {
					result.Diagnostics.Append(result.Identity.Set(ctx, uuid)...)
				}
			}

			if !push(result) {
				return
			}
		}
	}
}
