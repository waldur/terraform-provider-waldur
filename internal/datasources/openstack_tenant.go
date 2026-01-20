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
var _ datasource.DataSource = &OpenstackTenantDataSource{}

func NewOpenstackTenantDataSource() datasource.DataSource {
	return &OpenstackTenantDataSource{}
}

// OpenstackTenantDataSource defines the data source implementation.
type OpenstackTenantDataSource struct {
	client *client.Client
}

// OpenstackTenantDataSourceModel describes the data source data model.
type OpenstackTenantDataSourceModel struct {
	UUID                        types.String `tfsdk:"id"`
	BackendId                   types.String `tfsdk:"backend_id"`
	CanManage                   types.Bool   `tfsdk:"can_manage"`
	Customer                    types.String `tfsdk:"customer"`
	CustomerAbbreviation        types.String `tfsdk:"customer_abbreviation"`
	CustomerName                types.String `tfsdk:"customer_name"`
	CustomerNativeName          types.String `tfsdk:"customer_native_name"`
	CustomerUuid                types.String `tfsdk:"customer_uuid"`
	Description                 types.String `tfsdk:"description"`
	ExternalIp                  types.String `tfsdk:"external_ip"`
	Name                        types.String `tfsdk:"name"`
	NameExact                   types.String `tfsdk:"name_exact"`
	Project                     types.String `tfsdk:"project"`
	ProjectName                 types.String `tfsdk:"project_name"`
	ProjectUuid                 types.String `tfsdk:"project_uuid"`
	ServiceSettingsName         types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid         types.String `tfsdk:"service_settings_uuid"`
	State                       types.String `tfsdk:"state"`
	Uuid                        types.String `tfsdk:"uuid"`
	AccessUrl                   types.String `tfsdk:"access_url"`
	AvailabilityZone            types.String `tfsdk:"availability_zone"`
	Created                     types.String `tfsdk:"created"`
	DefaultVolumeTypeName       types.String `tfsdk:"default_volume_type_name"`
	ErrorMessage                types.String `tfsdk:"error_message"`
	ErrorTraceback              types.String `tfsdk:"error_traceback"`
	ExternalNetworkId           types.String `tfsdk:"external_network_id"`
	InternalNetworkId           types.String `tfsdk:"internal_network_id"`
	IsLimitBased                types.Bool   `tfsdk:"is_limit_based"`
	IsUsageBased                types.Bool   `tfsdk:"is_usage_based"`
	MarketplaceCategoryName     types.String `tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     types.String `tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     types.String `tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     types.String `tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         types.String `tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    types.String `tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     types.String `tfsdk:"marketplace_resource_uuid"`
	Modified                    types.String `tfsdk:"modified"`
	Quotas                      types.List   `tfsdk:"quotas"`
	ResourceType                types.String `tfsdk:"resource_type"`
	ServiceName                 types.String `tfsdk:"service_name"`
	ServiceSettings             types.String `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage types.String `tfsdk:"service_settings_error_message"`
	ServiceSettingsState        types.String `tfsdk:"service_settings_state"`
	Url                         types.String `tfsdk:"url"`
	UserPassword                types.String `tfsdk:"user_password"`
	UserUsername                types.String `tfsdk:"user_username"`
}

func (d *OpenstackTenantDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_tenant"
}

func (d *OpenstackTenantDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Tenant data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"can_manage": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Can manage",
			},
			"customer": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"customer_abbreviation": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"customer_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"customer_native_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"customer_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"external_ip": schema.StringAttribute{
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
			"project": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"project_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"project_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"service_settings_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"service_settings_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"state": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"access_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"availability_zone": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Optional availability group. Will be used for all instances provisioned in this tenant",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"default_volume_type_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Volume type name to use when creating volumes.",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"external_network_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of external network connected to OpenStack tenant",
			},
			"internal_network_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of internal network in OpenStack tenant",
			},
			"is_limit_based": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"is_usage_based": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_category_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_category_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_offering_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_offering_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_resource_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"quotas": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"limit": types.Int64Type,
					"name":  types.StringType,
					"usage": types.Int64Type,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_settings": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_settings_error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_settings_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"user_password": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Password of the tenant user",
			},
			"user_username": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Username of the tenant user",
			},
		},
	}
}

func (d *OpenstackTenantDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OpenstackTenantDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackTenantDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var item map[string]interface{}

		err := d.client.GetByUUID(ctx, "/api/openstack-tenants/{uuid}/", data.UUID.ValueString(), &item)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Tenant",
				"An error occurred while reading the Openstack Tenant by UUID: "+err.Error(),
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
		if val, ok := sourceMap["created"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Created = types.StringValue(str)
			}
		} else {
			if data.Created.IsUnknown() {
				data.Created = types.StringNull()
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
		if val, ok := sourceMap["backend_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BackendId = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["can_manage"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanManage = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["customer"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Customer = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_abbreviation"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerAbbreviation = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_native_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerNativeName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Description = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["external_ip"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ExternalIp = types.StringValue(str)
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
		if val, ok := sourceMap["project"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Project = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_settings_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_settings_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.State = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Uuid = types.StringValue(str)
			}
		}

	} else {
		// Filter by provided parameters
		var results []map[string]interface{}

		filters := map[string]string{}
		if !data.BackendId.IsNull() {
			filters["backend_id"] = data.BackendId.ValueString()
		}
		if !data.CanManage.IsNull() {
			filters["can_manage"] = fmt.Sprintf("%t", data.CanManage.ValueBool())
		}
		if !data.Customer.IsNull() {
			filters["customer"] = data.Customer.ValueString()
		}
		if !data.CustomerAbbreviation.IsNull() {
			filters["customer_abbreviation"] = data.CustomerAbbreviation.ValueString()
		}
		if !data.CustomerName.IsNull() {
			filters["customer_name"] = data.CustomerName.ValueString()
		}
		if !data.CustomerNativeName.IsNull() {
			filters["customer_native_name"] = data.CustomerNativeName.ValueString()
		}
		if !data.CustomerUuid.IsNull() {
			filters["customer_uuid"] = data.CustomerUuid.ValueString()
		}
		if !data.Description.IsNull() {
			filters["description"] = data.Description.ValueString()
		}
		if !data.ExternalIp.IsNull() {
			filters["external_ip"] = data.ExternalIp.ValueString()
		}
		if !data.Name.IsNull() {
			filters["name"] = data.Name.ValueString()
		}
		if !data.NameExact.IsNull() {
			filters["name_exact"] = data.NameExact.ValueString()
		}
		if !data.Project.IsNull() {
			filters["project"] = data.Project.ValueString()
		}
		if !data.ProjectName.IsNull() {
			filters["project_name"] = data.ProjectName.ValueString()
		}
		if !data.ProjectUuid.IsNull() {
			filters["project_uuid"] = data.ProjectUuid.ValueString()
		}
		if !data.ServiceSettingsName.IsNull() {
			filters["service_settings_name"] = data.ServiceSettingsName.ValueString()
		}
		if !data.ServiceSettingsUuid.IsNull() {
			filters["service_settings_uuid"] = data.ServiceSettingsUuid.ValueString()
		}
		if !data.State.IsNull() {
			filters["state"] = data.State.ValueString()
		}
		if !data.Uuid.IsNull() {
			filters["uuid"] = data.Uuid.ValueString()
		}

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_tenant.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/openstack-tenants/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Openstack Tenant",
				"An error occurred while filtering Openstack Tenant: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Tenant Not Found",
				"No Openstack Tenant found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Tenants Found",
				fmt.Sprintf("Found %d Openstack Tenants with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
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
		if val, ok := sourceMap["created"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Created = types.StringValue(str)
			}
		} else {
			if data.Created.IsUnknown() {
				data.Created = types.StringNull()
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
		if val, ok := sourceMap["backend_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BackendId = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["can_manage"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanManage = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["customer"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Customer = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_abbreviation"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerAbbreviation = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_native_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerNativeName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Description = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["external_ip"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ExternalIp = types.StringValue(str)
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
		if val, ok := sourceMap["project"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Project = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_settings_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_settings_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.State = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Uuid = types.StringValue(str)
			}
		}
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
