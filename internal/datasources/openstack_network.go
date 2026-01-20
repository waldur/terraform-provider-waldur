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
var _ datasource.DataSource = &OpenstackNetworkDataSource{}

func NewOpenstackNetworkDataSource() datasource.DataSource {
	return &OpenstackNetworkDataSource{}
}

// OpenstackNetworkDataSource defines the data source implementation.
type OpenstackNetworkDataSource struct {
	client *client.Client
}

// OpenstackNetworkDataSourceModel describes the data source data model.
type OpenstackNetworkDataSourceModel struct {
	UUID                 types.String `tfsdk:"id"`
	BackendId            types.String `tfsdk:"backend_id"`
	CanManage            types.Bool   `tfsdk:"can_manage"`
	Customer             types.String `tfsdk:"customer"`
	CustomerAbbreviation types.String `tfsdk:"customer_abbreviation"`
	CustomerName         types.String `tfsdk:"customer_name"`
	CustomerNativeName   types.String `tfsdk:"customer_native_name"`
	CustomerUuid         types.String `tfsdk:"customer_uuid"`
	Description          types.String `tfsdk:"description"`
	DirectOnly           types.Bool   `tfsdk:"direct_only"`
	ExternalIp           types.String `tfsdk:"external_ip"`
	IsExternal           types.Bool   `tfsdk:"is_external"`
	Name                 types.String `tfsdk:"name"`
	NameExact            types.String `tfsdk:"name_exact"`
	Project              types.String `tfsdk:"project"`
	ProjectName          types.String `tfsdk:"project_name"`
	ProjectUuid          types.String `tfsdk:"project_uuid"`
	RbacOnly             types.Bool   `tfsdk:"rbac_only"`
	ServiceSettingsName  types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid  types.String `tfsdk:"service_settings_uuid"`
	State                types.String `tfsdk:"state"`
	Tenant               types.String `tfsdk:"tenant"`
	TenantUuid           types.String `tfsdk:"tenant_uuid"`
	Type                 types.String `tfsdk:"type"`
	Uuid                 types.String `tfsdk:"uuid"`
	AccessUrl            types.String `tfsdk:"access_url"`
	Created              types.String `tfsdk:"created"`
	ErrorMessage         types.String `tfsdk:"error_message"`
	ErrorTraceback       types.String `tfsdk:"error_traceback"`
	Modified             types.String `tfsdk:"modified"`
	Mtu                  types.Int64  `tfsdk:"mtu"`
	RbacPolicies         types.List   `tfsdk:"rbac_policies"`
	ResourceType         types.String `tfsdk:"resource_type"`
	SegmentationId       types.Int64  `tfsdk:"segmentation_id"`
	Subnets              types.List   `tfsdk:"subnets"`
	TenantName           types.String `tfsdk:"tenant_name"`
	Url                  types.String `tfsdk:"url"`
}

func (d *OpenstackNetworkDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_network"
}

func (d *OpenstackNetworkDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Network data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Backend ID",
			},
			"can_manage": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Can manage",
			},
			"customer": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer UUID",
			},
			"customer_abbreviation": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer abbreviation",
			},
			"customer_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer name",
			},
			"customer_native_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer native name",
			},
			"customer_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer UUID",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description",
			},
			"direct_only": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Direct only",
			},
			"external_ip": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "External IP",
			},
			"is_external": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name (exact)",
			},
			"project": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project UUID",
			},
			"project_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project name",
			},
			"project_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project UUID",
			},
			"rbac_only": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "RBAC only",
			},
			"service_settings_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Service settings name",
			},
			"service_settings_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Service settings UUID",
			},
			"state": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "State",
			},
			"tenant": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Tenant URL",
			},
			"tenant_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Tenant UUID",
			},
			"type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "UUID",
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
			"mtu": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The maximum transmission unit (MTU) value to address fragmentation.",
			},
			"rbac_policies": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"backend_id":         types.StringType,
					"created":            types.StringType,
					"network":            types.StringType,
					"network_name":       types.StringType,
					"policy_type":        types.StringType,
					"target_tenant":      types.StringType,
					"target_tenant_name": types.StringType,
					"url":                types.StringType,
				}}},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"segmentation_id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "VLAN ID for VLAN networks or tunnel ID for VXLAN/GRE networks",
			},
			"subnets": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"end":   types.StringType,
						"start": types.StringType,
					}}},
					"cidr":        types.StringType,
					"description": types.StringType,
					"enable_dhcp": types.BoolType,
					"gateway_ip":  types.StringType,
					"ip_version":  types.Int64Type,
					"name":        types.StringType,
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

func (d *OpenstackNetworkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OpenstackNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackNetworkDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var item map[string]interface{}

		err := d.client.GetByUUID(ctx, "/api/openstack-networks/{uuid}/", data.UUID.ValueString(), &item)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Network",
				"An error occurred while reading the Openstack Network by UUID: "+err.Error(),
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
		if val, ok := sourceMap["mtu"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Mtu = types.Int64Value(int64(num))
			}
		} else {
			if data.Mtu.IsUnknown() {
				data.Mtu = types.Int64Null()
			}
		}
		if val, ok := sourceMap["rbac_policies"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"backend_id":         types.StringType,
							"created":            types.StringType,
							"network":            types.StringType,
							"network_name":       types.StringType,
							"policy_type":        types.StringType,
							"target_tenant":      types.StringType,
							"target_tenant_name": types.StringType,
							"url":                types.StringType,
						}
						attrValues := map[string]attr.Value{
							"backend_id": func() attr.Value {
								if v, ok := objMap["backend_id"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"created": func() attr.Value {
								if v, ok := objMap["created"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"network": func() attr.Value {
								if v, ok := objMap["network"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"network_name": func() attr.Value {
								if v, ok := objMap["network_name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"policy_type": func() attr.Value {
								if v, ok := objMap["policy_type"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"target_tenant": func() attr.Value {
								if v, ok := objMap["target_tenant"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"target_tenant_name": func() attr.Value {
								if v, ok := objMap["target_tenant_name"].(string); ok {
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
					"backend_id":         types.StringType,
					"created":            types.StringType,
					"network":            types.StringType,
					"network_name":       types.StringType,
					"policy_type":        types.StringType,
					"target_tenant":      types.StringType,
					"target_tenant_name": types.StringType,
					"url":                types.StringType,
				}}, items)
				data.RbacPolicies = listVal
			}
		} else {
			if data.RbacPolicies.IsUnknown() {
				data.RbacPolicies = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"backend_id":         types.StringType,
					"created":            types.StringType,
					"network":            types.StringType,
					"network_name":       types.StringType,
					"policy_type":        types.StringType,
					"target_tenant":      types.StringType,
					"target_tenant_name": types.StringType,
					"url":                types.StringType,
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
		if val, ok := sourceMap["segmentation_id"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.SegmentationId = types.Int64Value(int64(num))
			}
		} else {
			if data.SegmentationId.IsUnknown() {
				data.SegmentationId = types.Int64Null()
			}
		}
		if val, ok := sourceMap["subnets"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"end":   types.StringType,
								"start": types.StringType,
							}}},
							"cidr":        types.StringType,
							"description": types.StringType,
							"enable_dhcp": types.BoolType,
							"gateway_ip":  types.StringType,
							"ip_version":  types.Int64Type,
							"name":        types.StringType,
						}
						attrValues := map[string]attr.Value{
							"allocation_pools": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"end":   types.StringType,
								"start": types.StringType,
							}}}.ElemType),
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
							"enable_dhcp": func() attr.Value {
								if v, ok := objMap["enable_dhcp"].(bool); ok {
									return types.BoolValue(v)
								}
								return types.BoolNull()
							}(),
							"gateway_ip": func() attr.Value {
								if v, ok := objMap["gateway_ip"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"ip_version": func() attr.Value {
								if v, ok := objMap["ip_version"].(float64); ok {
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
						}
						objVal, _ := types.ObjectValue(attrTypes, attrValues)
						items = append(items, objVal)
					}
				}
				listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
					"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"end":   types.StringType,
						"start": types.StringType,
					}}},
					"cidr":        types.StringType,
					"description": types.StringType,
					"enable_dhcp": types.BoolType,
					"gateway_ip":  types.StringType,
					"ip_version":  types.Int64Type,
					"name":        types.StringType,
				}}, items)
				data.Subnets = listVal
			}
		} else {
			if data.Subnets.IsUnknown() {
				data.Subnets = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"end":   types.StringType,
						"start": types.StringType,
					}}},
					"cidr":        types.StringType,
					"description": types.StringType,
					"enable_dhcp": types.BoolType,
					"gateway_ip":  types.StringType,
					"ip_version":  types.Int64Type,
					"name":        types.StringType,
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
		if val, ok := sourceMap["direct_only"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.DirectOnly = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["external_ip"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ExternalIp = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["is_external"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsExternal = types.BoolValue(b)
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
		if val, ok := sourceMap["rbac_only"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.RbacOnly = types.BoolValue(b)
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
		if val, ok := sourceMap["type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Type = types.StringValue(str)
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
		if !data.DirectOnly.IsNull() {
			filters["direct_only"] = fmt.Sprintf("%t", data.DirectOnly.ValueBool())
		}
		if !data.ExternalIp.IsNull() {
			filters["external_ip"] = data.ExternalIp.ValueString()
		}
		if !data.IsExternal.IsNull() {
			filters["is_external"] = fmt.Sprintf("%t", data.IsExternal.ValueBool())
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
		if !data.RbacOnly.IsNull() {
			filters["rbac_only"] = fmt.Sprintf("%t", data.RbacOnly.ValueBool())
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
		if !data.Type.IsNull() {
			filters["type"] = data.Type.ValueString()
		}
		if !data.Uuid.IsNull() {
			filters["uuid"] = data.Uuid.ValueString()
		}

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_network.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/openstack-networks/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Openstack Network",
				"An error occurred while filtering Openstack Network: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Network Not Found",
				"No Openstack Network found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Networks Found",
				fmt.Sprintf("Found %d Openstack Networks with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
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
		if val, ok := sourceMap["mtu"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Mtu = types.Int64Value(int64(num))
			}
		} else {
			if data.Mtu.IsUnknown() {
				data.Mtu = types.Int64Null()
			}
		}
		if val, ok := sourceMap["rbac_policies"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"backend_id":         types.StringType,
							"created":            types.StringType,
							"network":            types.StringType,
							"network_name":       types.StringType,
							"policy_type":        types.StringType,
							"target_tenant":      types.StringType,
							"target_tenant_name": types.StringType,
							"url":                types.StringType,
						}
						attrValues := map[string]attr.Value{
							"backend_id": func() attr.Value {
								if v, ok := objMap["backend_id"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"created": func() attr.Value {
								if v, ok := objMap["created"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"network": func() attr.Value {
								if v, ok := objMap["network"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"network_name": func() attr.Value {
								if v, ok := objMap["network_name"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"policy_type": func() attr.Value {
								if v, ok := objMap["policy_type"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"target_tenant": func() attr.Value {
								if v, ok := objMap["target_tenant"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"target_tenant_name": func() attr.Value {
								if v, ok := objMap["target_tenant_name"].(string); ok {
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
					"backend_id":         types.StringType,
					"created":            types.StringType,
					"network":            types.StringType,
					"network_name":       types.StringType,
					"policy_type":        types.StringType,
					"target_tenant":      types.StringType,
					"target_tenant_name": types.StringType,
					"url":                types.StringType,
				}}, items)
				data.RbacPolicies = listVal
			}
		} else {
			if data.RbacPolicies.IsUnknown() {
				data.RbacPolicies = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"backend_id":         types.StringType,
					"created":            types.StringType,
					"network":            types.StringType,
					"network_name":       types.StringType,
					"policy_type":        types.StringType,
					"target_tenant":      types.StringType,
					"target_tenant_name": types.StringType,
					"url":                types.StringType,
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
		if val, ok := sourceMap["segmentation_id"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.SegmentationId = types.Int64Value(int64(num))
			}
		} else {
			if data.SegmentationId.IsUnknown() {
				data.SegmentationId = types.Int64Null()
			}
		}
		if val, ok := sourceMap["subnets"]; ok && val != nil {
			// List of objects
			if arr, ok := val.([]interface{}); ok {
				items := make([]attr.Value, 0, len(arr))
				for _, item := range arr {
					if objMap, ok := item.(map[string]interface{}); ok {
						attrTypes := map[string]attr.Type{
							"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"end":   types.StringType,
								"start": types.StringType,
							}}},
							"cidr":        types.StringType,
							"description": types.StringType,
							"enable_dhcp": types.BoolType,
							"gateway_ip":  types.StringType,
							"ip_version":  types.Int64Type,
							"name":        types.StringType,
						}
						attrValues := map[string]attr.Value{
							"allocation_pools": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
								"end":   types.StringType,
								"start": types.StringType,
							}}}.ElemType),
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
							"enable_dhcp": func() attr.Value {
								if v, ok := objMap["enable_dhcp"].(bool); ok {
									return types.BoolValue(v)
								}
								return types.BoolNull()
							}(),
							"gateway_ip": func() attr.Value {
								if v, ok := objMap["gateway_ip"].(string); ok {
									return types.StringValue(v)
								}
								return types.StringNull()
							}(),
							"ip_version": func() attr.Value {
								if v, ok := objMap["ip_version"].(float64); ok {
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
						}
						objVal, _ := types.ObjectValue(attrTypes, attrValues)
						items = append(items, objVal)
					}
				}
				listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
					"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"end":   types.StringType,
						"start": types.StringType,
					}}},
					"cidr":        types.StringType,
					"description": types.StringType,
					"enable_dhcp": types.BoolType,
					"gateway_ip":  types.StringType,
					"ip_version":  types.Int64Type,
					"name":        types.StringType,
				}}, items)
				data.Subnets = listVal
			}
		} else {
			if data.Subnets.IsUnknown() {
				data.Subnets = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
						"end":   types.StringType,
						"start": types.StringType,
					}}},
					"cidr":        types.StringType,
					"description": types.StringType,
					"enable_dhcp": types.BoolType,
					"gateway_ip":  types.StringType,
					"ip_version":  types.Int64Type,
					"name":        types.StringType,
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
		if val, ok := sourceMap["direct_only"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.DirectOnly = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["external_ip"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ExternalIp = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["is_external"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsExternal = types.BoolValue(b)
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
		if val, ok := sourceMap["rbac_only"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.RbacOnly = types.BoolValue(b)
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
		if val, ok := sourceMap["type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Type = types.StringValue(str)
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
