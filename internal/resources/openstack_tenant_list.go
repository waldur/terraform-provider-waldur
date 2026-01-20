package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-provider-waldur/internal/client"
)

var _ list.ListResource = &OpenstackTenantList{}

type OpenstackTenantList struct {
	client *client.Client
}

func NewOpenstackTenantList() list.ListResource {
	return &OpenstackTenantList{}
}

func (l *OpenstackTenantList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_tenant"
}

func (l *OpenstackTenantList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"backend_id": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"can_manage": schema.BoolAttribute{
				Description: "Can manage",
				Optional:    true,
			},
			"customer": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"customer_abbreviation": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"customer_name": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"customer_native_name": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"customer_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"external_ip": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"name_exact": schema.StringAttribute{
				Description: "",
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
			"project": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"project_name": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"project_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"service_settings_name": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"service_settings_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
		},
	}
}

func (l *OpenstackTenantList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type OpenstackTenantListModel struct {
	// Add filter fields here if added to schema
	BackendId            types.String `tfsdk:"backend_id"`
	CanManage            types.Bool   `tfsdk:"can_manage"`
	Customer             types.String `tfsdk:"customer"`
	CustomerAbbreviation types.String `tfsdk:"customer_abbreviation"`
	CustomerName         types.String `tfsdk:"customer_name"`
	CustomerNativeName   types.String `tfsdk:"customer_native_name"`
	CustomerUuid         types.String `tfsdk:"customer_uuid"`
	Description          types.String `tfsdk:"description"`
	ExternalIp           types.String `tfsdk:"external_ip"`
	Name                 types.String `tfsdk:"name"`
	NameExact            types.String `tfsdk:"name_exact"`
	Page                 types.Int64  `tfsdk:"page"`
	PageSize             types.Int64  `tfsdk:"page_size"`
	Project              types.String `tfsdk:"project"`
	ProjectName          types.String `tfsdk:"project_name"`
	ProjectUuid          types.String `tfsdk:"project_uuid"`
	ServiceSettingsName  types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid  types.String `tfsdk:"service_settings_uuid"`
	Uuid                 types.String `tfsdk:"uuid"`
}

func (l *OpenstackTenantList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config OpenstackTenantListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := make(map[string]string)
	if !config.BackendId.IsNull() && !config.BackendId.IsUnknown() {
		filters["backend_id"] = config.BackendId.ValueString()
	}
	if !config.CanManage.IsNull() && !config.CanManage.IsUnknown() {
		filters["can_manage"] = fmt.Sprintf("%t", config.CanManage.ValueBool())
	}
	if !config.Customer.IsNull() && !config.Customer.IsUnknown() {
		filters["customer"] = config.Customer.ValueString()
	}
	if !config.CustomerAbbreviation.IsNull() && !config.CustomerAbbreviation.IsUnknown() {
		filters["customer_abbreviation"] = config.CustomerAbbreviation.ValueString()
	}
	if !config.CustomerName.IsNull() && !config.CustomerName.IsUnknown() {
		filters["customer_name"] = config.CustomerName.ValueString()
	}
	if !config.CustomerNativeName.IsNull() && !config.CustomerNativeName.IsUnknown() {
		filters["customer_native_name"] = config.CustomerNativeName.ValueString()
	}
	if !config.CustomerUuid.IsNull() && !config.CustomerUuid.IsUnknown() {
		filters["customer_uuid"] = config.CustomerUuid.ValueString()
	}
	if !config.Description.IsNull() && !config.Description.IsUnknown() {
		filters["description"] = config.Description.ValueString()
	}
	if !config.ExternalIp.IsNull() && !config.ExternalIp.IsUnknown() {
		filters["external_ip"] = config.ExternalIp.ValueString()
	}
	if !config.Name.IsNull() && !config.Name.IsUnknown() {
		filters["name"] = config.Name.ValueString()
	}
	if !config.NameExact.IsNull() && !config.NameExact.IsUnknown() {
		filters["name_exact"] = config.NameExact.ValueString()
	}
	if !config.Page.IsNull() && !config.Page.IsUnknown() {
		filters["page"] = fmt.Sprintf("%d", config.Page.ValueInt64())
	}
	if !config.PageSize.IsNull() && !config.PageSize.IsUnknown() {
		filters["page_size"] = fmt.Sprintf("%d", config.PageSize.ValueInt64())
	}
	if !config.Project.IsNull() && !config.Project.IsUnknown() {
		filters["project"] = config.Project.ValueString()
	}
	if !config.ProjectName.IsNull() && !config.ProjectName.IsUnknown() {
		filters["project_name"] = config.ProjectName.ValueString()
	}
	if !config.ProjectUuid.IsNull() && !config.ProjectUuid.IsUnknown() {
		filters["project_uuid"] = config.ProjectUuid.ValueString()
	}
	if !config.ServiceSettingsName.IsNull() && !config.ServiceSettingsName.IsUnknown() {
		filters["service_settings_name"] = config.ServiceSettingsName.ValueString()
	}
	if !config.ServiceSettingsUuid.IsNull() && !config.ServiceSettingsUuid.IsUnknown() {
		filters["service_settings_uuid"] = config.ServiceSettingsUuid.ValueString()
	}
	if !config.Uuid.IsNull() && !config.Uuid.IsUnknown() {
		filters["uuid"] = config.Uuid.ValueString()
	}

	// Call API
	var listResult []map[string]interface{}
	err := l.client.ListWithFilter(ctx, "/api/openstack-tenants/", filters, &listResult)
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
			var data OpenstackTenantResourceModel

			// Initialize lists and objects to empty/null if needed (or rely on map logic)

			sourceMap := item

			// Reuse the mapResponseFields logic (embedded here or via shared template)
			// Map response fields to data model
			_ = sourceMap
			if val, ok := sourceMap["access_url"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.AccessUrl = types.StringValue(str)
				}
			} else {
				if data.AccessUrl.IsUnknown() {
					data.AccessUrl = types.StringNull()
				}
			}
			if val, ok := sourceMap["availability_zone"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.AvailabilityZone = types.StringValue(str)
				}
			} else {
				if data.AvailabilityZone.IsUnknown() {
					data.AvailabilityZone = types.StringNull()
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
			if val, ok := sourceMap["created"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Created = types.StringValue(str)
				}
			} else {
				if data.Created.IsUnknown() {
					data.Created = types.StringNull()
				}
			}
			if val, ok := sourceMap["customer"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Customer = types.StringValue(str)
				}
			} else {
				if data.Customer.IsUnknown() {
					data.Customer = types.StringNull()
				}
			}
			if val, ok := sourceMap["customer_abbreviation"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.CustomerAbbreviation = types.StringValue(str)
				}
			} else {
				if data.CustomerAbbreviation.IsUnknown() {
					data.CustomerAbbreviation = types.StringNull()
				}
			}
			if val, ok := sourceMap["customer_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.CustomerName = types.StringValue(str)
				}
			} else {
				if data.CustomerName.IsUnknown() {
					data.CustomerName = types.StringNull()
				}
			}
			if val, ok := sourceMap["customer_native_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.CustomerNativeName = types.StringValue(str)
				}
			} else {
				if data.CustomerNativeName.IsUnknown() {
					data.CustomerNativeName = types.StringNull()
				}
			}
			if val, ok := sourceMap["customer_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.CustomerUuid = types.StringValue(str)
				}
			} else {
				if data.CustomerUuid.IsUnknown() {
					data.CustomerUuid = types.StringNull()
				}
			}
			if val, ok := sourceMap["default_volume_type_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.DefaultVolumeTypeName = types.StringValue(str)
				}
			} else {
				if data.DefaultVolumeTypeName.IsUnknown() {
					data.DefaultVolumeTypeName = types.StringNull()
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
			if val, ok := sourceMap["error_message"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ErrorMessage = types.StringValue(str)
				}
			} else {
				if data.ErrorMessage.IsUnknown() {
					data.ErrorMessage = types.StringNull()
				}
			}
			if val, ok := sourceMap["error_traceback"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ErrorTraceback = types.StringValue(str)
				}
			} else {
				if data.ErrorTraceback.IsUnknown() {
					data.ErrorTraceback = types.StringNull()
				}
			}
			if val, ok := sourceMap["external_network_id"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ExternalNetworkId = types.StringValue(str)
				}
			} else {
				if data.ExternalNetworkId.IsUnknown() {
					data.ExternalNetworkId = types.StringNull()
				}
			}
			if val, ok := sourceMap["internal_network_id"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.InternalNetworkId = types.StringValue(str)
				}
			} else {
				if data.InternalNetworkId.IsUnknown() {
					data.InternalNetworkId = types.StringNull()
				}
			}
			if val, ok := sourceMap["is_limit_based"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.IsLimitBased = types.BoolValue(b)
				}
			} else {
				if data.IsLimitBased.IsUnknown() {
					data.IsLimitBased = types.BoolNull()
				}
			}
			if val, ok := sourceMap["is_usage_based"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.IsUsageBased = types.BoolValue(b)
				}
			} else {
				if data.IsUsageBased.IsUnknown() {
					data.IsUsageBased = types.BoolNull()
				}
			}
			if val, ok := sourceMap["marketplace_category_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.MarketplaceCategoryName = types.StringValue(str)
				}
			} else {
				if data.MarketplaceCategoryName.IsUnknown() {
					data.MarketplaceCategoryName = types.StringNull()
				}
			}
			if val, ok := sourceMap["marketplace_category_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.MarketplaceCategoryUuid = types.StringValue(str)
				}
			} else {
				if data.MarketplaceCategoryUuid.IsUnknown() {
					data.MarketplaceCategoryUuid = types.StringNull()
				}
			}
			if val, ok := sourceMap["marketplace_offering_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.MarketplaceOfferingName = types.StringValue(str)
				}
			} else {
				if data.MarketplaceOfferingName.IsUnknown() {
					data.MarketplaceOfferingName = types.StringNull()
				}
			}
			if val, ok := sourceMap["marketplace_offering_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.MarketplaceOfferingUuid = types.StringValue(str)
				}
			} else {
				if data.MarketplaceOfferingUuid.IsUnknown() {
					data.MarketplaceOfferingUuid = types.StringNull()
				}
			}
			if val, ok := sourceMap["marketplace_plan_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.MarketplacePlanUuid = types.StringValue(str)
				}
			} else {
				if data.MarketplacePlanUuid.IsUnknown() {
					data.MarketplacePlanUuid = types.StringNull()
				}
			}
			if val, ok := sourceMap["marketplace_resource_state"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.MarketplaceResourceState = types.StringValue(str)
				}
			} else {
				if data.MarketplaceResourceState.IsUnknown() {
					data.MarketplaceResourceState = types.StringNull()
				}
			}
			if val, ok := sourceMap["marketplace_resource_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.MarketplaceResourceUuid = types.StringValue(str)
				}
			} else {
				if data.MarketplaceResourceUuid.IsUnknown() {
					data.MarketplaceResourceUuid = types.StringNull()
				}
			}
			if val, ok := sourceMap["modified"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Modified = types.StringValue(str)
				}
			} else {
				if data.Modified.IsUnknown() {
					data.Modified = types.StringNull()
				}
			}
			if val, ok := sourceMap["offering"]; ok && val != nil {
				if str, ok := val.(string); ok {
					// Normalize URL to UUID
					parts := strings.Split(strings.TrimRight(str, "/"), "/")
					uuid := parts[len(parts)-1]
					data.Offering = types.StringValue(uuid)
				} else {
					data.Offering = types.StringNull()
				}
			} else {
				if data.Offering.IsUnknown() {
					data.Offering = types.StringNull()
				}
			}
			if val, ok := sourceMap["project"]; ok && val != nil {
				if str, ok := val.(string); ok {
					// Normalize URL to UUID
					parts := strings.Split(strings.TrimRight(str, "/"), "/")
					uuid := parts[len(parts)-1]
					data.Project = types.StringValue(uuid)
				} else {
					data.Project = types.StringNull()
				}
			} else {
				if data.Project.IsUnknown() {
					data.Project = types.StringNull()
				}
			}
			if val, ok := sourceMap["project_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ProjectName = types.StringValue(str)
				}
			} else {
				if data.ProjectName.IsUnknown() {
					data.ProjectName = types.StringNull()
				}
			}
			if val, ok := sourceMap["project_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ProjectUuid = types.StringValue(str)
				}
			} else {
				if data.ProjectUuid.IsUnknown() {
					data.ProjectUuid = types.StringNull()
				}
			}
			if val, ok := sourceMap["quotas"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"limit": types.Int64Type,
								"name":  types.StringType,
								"usage": types.Int64Type,
							}
							attrValues := map[string]attr.Value{
								"limit": func() attr.Value {
									if v, ok := objMap["limit"].(float64); ok {
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
								"usage": func() attr.Value {
									if v, ok := objMap["usage"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
							}
							objVal, _ := types.ObjectValue(attrTypes, attrValues)
							items = append(items, objVal)
						}
					}
					listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
						"limit": types.Int64Type,
						"name":  types.StringType,
						"usage": types.Int64Type,
					}}, items)
					data.Quotas = listVal
				}
			} else {
				if data.Quotas.IsUnknown() {
					data.Quotas = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"limit": types.Int64Type,
						"name":  types.StringType,
						"usage": types.Int64Type,
					}})
				}
			}
			if val, ok := sourceMap["resource_type"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ResourceType = types.StringValue(str)
				}
			} else {
				if data.ResourceType.IsUnknown() {
					data.ResourceType = types.StringNull()
				}
			}
			if val, ok := sourceMap["security_groups"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"description": types.StringType,
								"name":        types.StringType,
								"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"cidr":         types.StringType,
									"description":  types.StringType,
									"direction":    types.StringType,
									"ethertype":    types.StringType,
									"from_port":    types.Int64Type,
									"protocol":     types.StringType,
									"remote_group": types.StringType,
									"to_port":      types.Int64Type,
								}}},
							}
							attrValues := map[string]attr.Value{
								"description": func() attr.Value {
									if v, ok := objMap["description"].(string); ok {
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
								"rules": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"cidr":         types.StringType,
									"description":  types.StringType,
									"direction":    types.StringType,
									"ethertype":    types.StringType,
									"from_port":    types.Int64Type,
									"protocol":     types.StringType,
									"remote_group": types.StringType,
									"to_port":      types.Int64Type,
								}}}.ElemType),
							}
							objVal, _ := types.ObjectValue(attrTypes, attrValues)
							items = append(items, objVal)
						}
					}
					listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
						"description": types.StringType,
						"name":        types.StringType,
						"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"cidr":         types.StringType,
							"description":  types.StringType,
							"direction":    types.StringType,
							"ethertype":    types.StringType,
							"from_port":    types.Int64Type,
							"protocol":     types.StringType,
							"remote_group": types.StringType,
							"to_port":      types.Int64Type,
						}}},
					}}, items)
					data.SecurityGroups = listVal
				}
			} else {
				if data.SecurityGroups.IsUnknown() {
					data.SecurityGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"description": types.StringType,
						"name":        types.StringType,
						"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"cidr":         types.StringType,
							"description":  types.StringType,
							"direction":    types.StringType,
							"ethertype":    types.StringType,
							"from_port":    types.Int64Type,
							"protocol":     types.StringType,
							"remote_group": types.StringType,
							"to_port":      types.Int64Type,
						}}},
					}})
				}
			}
			if val, ok := sourceMap["service_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ServiceName = types.StringValue(str)
				}
			} else {
				if data.ServiceName.IsUnknown() {
					data.ServiceName = types.StringNull()
				}
			}
			if val, ok := sourceMap["service_settings"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ServiceSettings = types.StringValue(str)
				}
			} else {
				if data.ServiceSettings.IsUnknown() {
					data.ServiceSettings = types.StringNull()
				}
			}
			if val, ok := sourceMap["service_settings_error_message"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ServiceSettingsErrorMessage = types.StringValue(str)
				}
			} else {
				if data.ServiceSettingsErrorMessage.IsUnknown() {
					data.ServiceSettingsErrorMessage = types.StringNull()
				}
			}
			if val, ok := sourceMap["service_settings_state"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ServiceSettingsState = types.StringValue(str)
				}
			} else {
				if data.ServiceSettingsState.IsUnknown() {
					data.ServiceSettingsState = types.StringNull()
				}
			}
			if val, ok := sourceMap["service_settings_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ServiceSettingsUuid = types.StringValue(str)
				}
			} else {
				if data.ServiceSettingsUuid.IsUnknown() {
					data.ServiceSettingsUuid = types.StringNull()
				}
			}
			if val, ok := sourceMap["skip_connection_extnet"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.SkipConnectionExtnet = types.BoolValue(b)
				}
			} else {
				if data.SkipConnectionExtnet.IsUnknown() {
					data.SkipConnectionExtnet = types.BoolNull()
				}
			}
			if val, ok := sourceMap["skip_creation_of_default_router"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.SkipCreationOfDefaultRouter = types.BoolValue(b)
				}
			} else {
				if data.SkipCreationOfDefaultRouter.IsUnknown() {
					data.SkipCreationOfDefaultRouter = types.BoolNull()
				}
			}
			if val, ok := sourceMap["state"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.State = types.StringValue(str)
				}
			} else {
				if data.State.IsUnknown() {
					data.State = types.StringNull()
				}
			}
			if val, ok := sourceMap["subnet_cidr"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.SubnetCidr = types.StringValue(str)
				}
			} else {
				if data.SubnetCidr.IsUnknown() {
					data.SubnetCidr = types.StringNull()
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
			if val, ok := sourceMap["user_password"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.UserPassword = types.StringValue(str)
				}
			} else {
				if data.UserPassword.IsUnknown() {
					data.UserPassword = types.StringNull()
				}
			}
			if val, ok := sourceMap["user_username"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.UserUsername = types.StringValue(str)
				}
			} else {
				if data.UserUsername.IsUnknown() {
					data.UserUsername = types.StringNull()
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
