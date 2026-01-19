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
var _ datasource.DataSource = &OpenstackSecurityGroupDataSource{}

func NewOpenstackSecurityGroupDataSource() datasource.DataSource {
	return &OpenstackSecurityGroupDataSource{}
}

// OpenstackSecurityGroupDataSource defines the data source implementation.
type OpenstackSecurityGroupDataSource struct {
	client *client.Client
}

// OpenstackSecurityGroupDataSourceModel describes the data source data model.
type OpenstackSecurityGroupDataSourceModel struct {
	UUID                 types.String `tfsdk:"id"`
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
	Project              types.String `tfsdk:"project"`
	ProjectName          types.String `tfsdk:"project_name"`
	ProjectUuid          types.String `tfsdk:"project_uuid"`
	Query                types.String `tfsdk:"query"`
	ServiceSettingsName  types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid  types.String `tfsdk:"service_settings_uuid"`
	State                types.String `tfsdk:"state"`
	Tenant               types.String `tfsdk:"tenant"`
	TenantUuid           types.String `tfsdk:"tenant_uuid"`
	Uuid                 types.String `tfsdk:"uuid"`
	AccessUrl            types.String `tfsdk:"access_url"`
	Created              types.String `tfsdk:"created"`
	ErrorMessage         types.String `tfsdk:"error_message"`
	ErrorTraceback       types.String `tfsdk:"error_traceback"`
	Modified             types.String `tfsdk:"modified"`
	ResourceType         types.String `tfsdk:"resource_type"`
	Rules                types.List   `tfsdk:"rules"`
	TenantName           types.String `tfsdk:"tenant_name"`
	Url                  types.String `tfsdk:"url"`
}

func (d *OpenstackSecurityGroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_security_group"
}

func (d *OpenstackSecurityGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "OpenstackSecurityGroup data source - lookup by name or UUID",

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
			"query": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Search by name or description",
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
			"tenant": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"tenant_uuid": schema.StringAttribute{
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
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"rules": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"cidr":              types.StringType,
					"description":       types.StringType,
					"direction":         types.StringType,
					"ethertype":         types.StringType,
					"from_port":         types.Int64Type,
					"id":                types.Int64Type,
					"protocol":          types.StringType,
					"remote_group":      types.StringType,
					"remote_group_name": types.StringType,
					"remote_group_uuid": types.StringType,
					"to_port":           types.Int64Type,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"tenant_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
		},
	}
}

func (d *OpenstackSecurityGroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OpenstackSecurityGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackSecurityGroupDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var item map[string]interface{}

		err := d.client.GetByUUID(ctx, "/api/openstack-security-groups/{uuid}/", data.UUID.ValueString(), &item)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read SecurityGroup",
				"An error occurred while reading the openstack_security_group by UUID: "+err.Error(),
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
		if val, ok := sourceMap["created"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Created = types.StringValue(str)
			}
		} else {
			if data.Created.IsUnknown() {
				data.Created = types.StringNull()
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
		if val, ok := sourceMap["modified"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Modified = types.StringValue(str)
			}
		} else {
			if data.Modified.IsUnknown() {
				data.Modified = types.StringNull()
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
		if val, ok := sourceMap["rules"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"cidr":              types.StringType,
							"description":       types.StringType,
							"direction":         types.StringType,
							"ethertype":         types.StringType,
							"from_port":         types.Int64Type,
							"id":                types.Int64Type,
							"protocol":          types.StringType,
							"remote_group":      types.StringType,
							"remote_group_name": types.StringType,
							"remote_group_uuid": types.StringType,
							"to_port":           types.Int64Type,
						}
						attrValues := map[string]attr.Value{
							"cidr": func() attr.Value {
								if v, ok := objMap["cidr"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"description": func() attr.Value {
								if v, ok := objMap["description"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"direction": func() attr.Value {
								if v, ok := objMap["direction"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"ethertype": func() attr.Value {
								if v, ok := objMap["ethertype"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"from_port": func() attr.Value {
								if v, ok := objMap["from_port"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"id": func() attr.Value {
								if v, ok := objMap["id"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"protocol": func() attr.Value {
								if v, ok := objMap["protocol"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"remote_group": func() attr.Value {
								if v, ok := objMap["remote_group"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"remote_group_name": func() attr.Value {
								if v, ok := objMap["remote_group_name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"remote_group_uuid": func() attr.Value {
								if v, ok := objMap["remote_group_uuid"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"to_port": func() attr.Value {
								if v, ok := objMap["to_port"].(float64); ok {
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
					"cidr":              types.StringType,
					"description":       types.StringType,
					"direction":         types.StringType,
					"ethertype":         types.StringType,
					"from_port":         types.Int64Type,
					"id":                types.Int64Type,
					"protocol":          types.StringType,
					"remote_group":      types.StringType,
					"remote_group_name": types.StringType,
					"remote_group_uuid": types.StringType,
					"to_port":           types.Int64Type,
				}}, items)
				data.Rules = listVal
			}
		} else {
			if data.Rules.IsUnknown() {
				data.Rules = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"cidr":              types.StringType,
					"description":       types.StringType,
					"direction":         types.StringType,
					"ethertype":         types.StringType,
					"from_port":         types.Int64Type,
					"id":                types.Int64Type,
					"protocol":          types.StringType,
					"remote_group":      types.StringType,
					"remote_group_name": types.StringType,
					"remote_group_uuid": types.StringType,
					"to_port":           types.Int64Type,
				}})
			}
		}
		if val, ok := sourceMap["tenant_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.TenantName = types.StringValue(str)
			}
		} else {
			if data.TenantName.IsUnknown() {
				data.TenantName = types.StringNull()
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

		// Map filter parameters from response if available
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
		if val, ok := sourceMap["query"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Query = types.StringValue(str)
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
		if val, ok := sourceMap["tenant"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Tenant = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["tenant_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.TenantUuid = types.StringValue(str)
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
		if !data.Query.IsNull() {
			filters["query"] = data.Query.ValueString()
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
		if !data.Tenant.IsNull() {
			filters["tenant"] = data.Tenant.ValueString()
		}
		if !data.TenantUuid.IsNull() {
			filters["tenant_uuid"] = data.TenantUuid.ValueString()
		}
		if !data.Uuid.IsNull() {
			filters["uuid"] = data.Uuid.ValueString()
		}

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_security_group.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/openstack-security-groups/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List SecurityGroup",
				"An error occurred while filtering openstack_security_group: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"SecurityGroup Not Found",
				"No openstack_security_group found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple SecurityGroups Found",
				fmt.Sprintf("Found %d openstack_security_groups with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
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
		if val, ok := sourceMap["created"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Created = types.StringValue(str)
			}
		} else {
			if data.Created.IsUnknown() {
				data.Created = types.StringNull()
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
		if val, ok := sourceMap["modified"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Modified = types.StringValue(str)
			}
		} else {
			if data.Modified.IsUnknown() {
				data.Modified = types.StringNull()
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
		if val, ok := sourceMap["rules"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"cidr":              types.StringType,
							"description":       types.StringType,
							"direction":         types.StringType,
							"ethertype":         types.StringType,
							"from_port":         types.Int64Type,
							"id":                types.Int64Type,
							"protocol":          types.StringType,
							"remote_group":      types.StringType,
							"remote_group_name": types.StringType,
							"remote_group_uuid": types.StringType,
							"to_port":           types.Int64Type,
						}
						attrValues := map[string]attr.Value{
							"cidr": func() attr.Value {
								if v, ok := objMap["cidr"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"description": func() attr.Value {
								if v, ok := objMap["description"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"direction": func() attr.Value {
								if v, ok := objMap["direction"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"ethertype": func() attr.Value {
								if v, ok := objMap["ethertype"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"from_port": func() attr.Value {
								if v, ok := objMap["from_port"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"id": func() attr.Value {
								if v, ok := objMap["id"].(float64); ok {
									return types.Int64Value(int64(v))
								}
								return types.Int64Null()
							}(),
							"protocol": func() attr.Value {
								if v, ok := objMap["protocol"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"remote_group": func() attr.Value {
								if v, ok := objMap["remote_group"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"remote_group_name": func() attr.Value {
								if v, ok := objMap["remote_group_name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"remote_group_uuid": func() attr.Value {
								if v, ok := objMap["remote_group_uuid"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"to_port": func() attr.Value {
								if v, ok := objMap["to_port"].(float64); ok {
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
					"cidr":              types.StringType,
					"description":       types.StringType,
					"direction":         types.StringType,
					"ethertype":         types.StringType,
					"from_port":         types.Int64Type,
					"id":                types.Int64Type,
					"protocol":          types.StringType,
					"remote_group":      types.StringType,
					"remote_group_name": types.StringType,
					"remote_group_uuid": types.StringType,
					"to_port":           types.Int64Type,
				}}, items)
				data.Rules = listVal
			}
		} else {
			if data.Rules.IsUnknown() {
				data.Rules = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"cidr":              types.StringType,
					"description":       types.StringType,
					"direction":         types.StringType,
					"ethertype":         types.StringType,
					"from_port":         types.Int64Type,
					"id":                types.Int64Type,
					"protocol":          types.StringType,
					"remote_group":      types.StringType,
					"remote_group_name": types.StringType,
					"remote_group_uuid": types.StringType,
					"to_port":           types.Int64Type,
				}})
			}
		}
		if val, ok := sourceMap["tenant_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.TenantName = types.StringValue(str)
			}
		} else {
			if data.TenantName.IsUnknown() {
				data.TenantName = types.StringNull()
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

		// Map filter parameters from response if available
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
		if val, ok := sourceMap["query"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Query = types.StringValue(str)
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
		if val, ok := sourceMap["tenant"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Tenant = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["tenant_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.TenantUuid = types.StringValue(str)
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
