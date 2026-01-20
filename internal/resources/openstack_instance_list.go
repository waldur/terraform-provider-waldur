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

var _ list.ListResource = &OpenstackInstanceList{}

type OpenstackInstanceList struct {
	client *client.Client
}

func NewOpenstackInstanceList() list.ListResource {
	return &OpenstackInstanceList{}
}

func (l *OpenstackInstanceList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_instance"
}

func (l *OpenstackInstanceList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"attach_volume_uuid": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"availability_zone_name": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
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
			"query": schema.StringAttribute{
				Description: "Search by name, internal IP, or external IP",
				Optional:    true,
			},
			"runtime_state": schema.StringAttribute{
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
			"tenant": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"tenant_uuid": schema.StringAttribute{
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

func (l *OpenstackInstanceList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type OpenstackInstanceListModel struct {
	// Add filter fields here if added to schema
	AttachVolumeUuid     types.String `tfsdk:"attach_volume_uuid"`
	AvailabilityZoneName types.String `tfsdk:"availability_zone_name"`
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
	Query                types.String `tfsdk:"query"`
	RuntimeState         types.String `tfsdk:"runtime_state"`
	ServiceSettingsName  types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid  types.String `tfsdk:"service_settings_uuid"`
	Tenant               types.String `tfsdk:"tenant"`
	TenantUuid           types.String `tfsdk:"tenant_uuid"`
	Uuid                 types.String `tfsdk:"uuid"`
}

func (l *OpenstackInstanceList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config OpenstackInstanceListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := make(map[string]string)
	if !config.AttachVolumeUuid.IsNull() && !config.AttachVolumeUuid.IsUnknown() {
		filters["attach_volume_uuid"] = config.AttachVolumeUuid.ValueString()
	}
	if !config.AvailabilityZoneName.IsNull() && !config.AvailabilityZoneName.IsUnknown() {
		filters["availability_zone_name"] = config.AvailabilityZoneName.ValueString()
	}
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
	if !config.Query.IsNull() && !config.Query.IsUnknown() {
		filters["query"] = config.Query.ValueString()
	}
	if !config.RuntimeState.IsNull() && !config.RuntimeState.IsUnknown() {
		filters["runtime_state"] = config.RuntimeState.ValueString()
	}
	if !config.ServiceSettingsName.IsNull() && !config.ServiceSettingsName.IsUnknown() {
		filters["service_settings_name"] = config.ServiceSettingsName.ValueString()
	}
	if !config.ServiceSettingsUuid.IsNull() && !config.ServiceSettingsUuid.IsUnknown() {
		filters["service_settings_uuid"] = config.ServiceSettingsUuid.ValueString()
	}
	if !config.Tenant.IsNull() && !config.Tenant.IsUnknown() {
		filters["tenant"] = config.Tenant.ValueString()
	}
	if !config.TenantUuid.IsNull() && !config.TenantUuid.IsUnknown() {
		filters["tenant_uuid"] = config.TenantUuid.ValueString()
	}
	if !config.Uuid.IsNull() && !config.Uuid.IsUnknown() {
		filters["uuid"] = config.Uuid.ValueString()
	}

	// Call API
	var listResult []map[string]interface{}
	err := l.client.ListWithFilter(ctx, "/api/openstack-instances/", filters, &listResult)
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
			var data OpenstackInstanceResourceModel

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
			if val, ok := sourceMap["action"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Action = types.StringValue(str)
				}
			} else {
				if data.Action.IsUnknown() {
					data.Action = types.StringNull()
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
			if val, ok := sourceMap["availability_zone_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.AvailabilityZoneName = types.StringValue(str)
				}
			} else {
				if data.AvailabilityZoneName.IsUnknown() {
					data.AvailabilityZoneName = types.StringNull()
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
			if val, ok := sourceMap["connect_directly_to_external_network"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.ConnectDirectlyToExternalNetwork = types.BoolValue(b)
				}
			} else {
				if data.ConnectDirectlyToExternalNetwork.IsUnknown() {
					data.ConnectDirectlyToExternalNetwork = types.BoolNull()
				}
			}
			if val, ok := sourceMap["cores"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.Cores = types.Int64Value(int64(num))
				}
			} else {
				if data.Cores.IsUnknown() {
					data.Cores = types.Int64Null()
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
			if val, ok := sourceMap["data_volume_size"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.DataVolumeSize = types.Int64Value(int64(num))
				}
			} else {
				if data.DataVolumeSize.IsUnknown() {
					data.DataVolumeSize = types.Int64Null()
				}
			}
			if val, ok := sourceMap["data_volume_type"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.DataVolumeType = types.StringValue(str)
				}
			} else {
				if data.DataVolumeType.IsUnknown() {
					data.DataVolumeType = types.StringNull()
				}
			}
			if val, ok := sourceMap["data_volumes"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"size":        types.Int64Type,
								"volume_type": types.StringType,
							}
							attrValues := map[string]attr.Value{
								"size": func() attr.Value {
									if v, ok := objMap["size"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"volume_type": func() attr.Value {
									if v, ok := objMap["volume_type"].(string); ok {
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
						"size":        types.Int64Type,
						"volume_type": types.StringType,
					}}, items)
					data.DataVolumes = listVal
				}
			} else {
				if data.DataVolumes.IsUnknown() {
					data.DataVolumes = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"size":        types.Int64Type,
						"volume_type": types.StringType,
					}})
				}
			}
			if val, ok := sourceMap["delete_volumes"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.DeleteVolumes = types.BoolValue(b)
				}
			} else {
				if data.DeleteVolumes.IsUnknown() {
					data.DeleteVolumes = types.BoolNull()
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
			if val, ok := sourceMap["disk"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.Disk = types.Int64Value(int64(num))
				}
			} else {
				if data.Disk.IsUnknown() {
					data.Disk = types.Int64Null()
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
			if val, ok := sourceMap["external_address"]; ok && val != nil {
				// List of strings (or flattened objects)
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if str, ok := item.(string); ok {
							items = append(items, types.StringValue(str))
						} else if obj, ok := item.(map[string]interface{}); ok {
							// Flattening logic: extract URL or UUID
							if url, ok := obj["url"].(string); ok {
								parts := strings.Split(strings.TrimRight(url, "/"), "/")
								uuid := parts[len(parts)-1]
								items = append(items, types.StringValue(uuid))
							} else if uuid, ok := obj["uuid"].(string); ok {
								items = append(items, types.StringValue(uuid))
							} else if name, ok := obj["name"].(string); ok {
								items = append(items, types.StringValue(name))
							}
						}
					}
					listVal, _ := types.ListValue(types.StringType, items)
					data.ExternalAddress = listVal
				}
			} else {
				if data.ExternalAddress.IsUnknown() {
					data.ExternalAddress = types.ListNull(types.StringType)
				}
			}
			if val, ok := sourceMap["external_ips"]; ok && val != nil {
				// List of strings (or flattened objects)
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if str, ok := item.(string); ok {
							items = append(items, types.StringValue(str))
						} else if obj, ok := item.(map[string]interface{}); ok {
							// Flattening logic: extract URL or UUID
							if url, ok := obj["url"].(string); ok {
								parts := strings.Split(strings.TrimRight(url, "/"), "/")
								uuid := parts[len(parts)-1]
								items = append(items, types.StringValue(uuid))
							} else if uuid, ok := obj["uuid"].(string); ok {
								items = append(items, types.StringValue(uuid))
							} else if name, ok := obj["name"].(string); ok {
								items = append(items, types.StringValue(name))
							}
						}
					}
					listVal, _ := types.ListValue(types.StringType, items)
					data.ExternalIps = listVal
				}
			} else {
				if data.ExternalIps.IsUnknown() {
					data.ExternalIps = types.ListNull(types.StringType)
				}
			}
			if val, ok := sourceMap["flavor"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Flavor = types.StringValue(str)
				}
			} else {
				if data.Flavor.IsUnknown() {
					data.Flavor = types.StringNull()
				}
			}
			if val, ok := sourceMap["flavor_disk"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.FlavorDisk = types.Int64Value(int64(num))
				}
			} else {
				if data.FlavorDisk.IsUnknown() {
					data.FlavorDisk = types.Int64Null()
				}
			}
			if val, ok := sourceMap["flavor_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.FlavorName = types.StringValue(str)
				}
			} else {
				if data.FlavorName.IsUnknown() {
					data.FlavorName = types.StringNull()
				}
			}
			if val, ok := sourceMap["floating_ips"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"ip_address": types.StringType,
								"subnet":     types.StringType,
								"url":        types.StringType,
								"address":    types.StringType,
								"port_fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"ip_address": types.StringType,
									"subnet_id":  types.StringType,
								}}},
								"port_mac_address":   types.StringType,
								"subnet_cidr":        types.StringType,
								"subnet_description": types.StringType,
								"subnet_name":        types.StringType,
								"subnet_uuid":        types.StringType,
							}
							attrValues := map[string]attr.Value{
								"ip_address": func() attr.Value {
									if v, ok := objMap["ip_address"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"subnet": func() attr.Value {
									if v, ok := objMap["subnet"].(string); ok {
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
								"address": func() attr.Value {
									if v, ok := objMap["address"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"port_fixed_ips": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"ip_address": types.StringType,
									"subnet_id":  types.StringType,
								}}}.ElemType),
								"port_mac_address": func() attr.Value {
									if v, ok := objMap["port_mac_address"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"subnet_cidr": func() attr.Value {
									if v, ok := objMap["subnet_cidr"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"subnet_description": func() attr.Value {
									if v, ok := objMap["subnet_description"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"subnet_name": func() attr.Value {
									if v, ok := objMap["subnet_name"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"subnet_uuid": func() attr.Value {
									if v, ok := objMap["subnet_uuid"].(string); ok {
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
						"ip_address": types.StringType,
						"subnet":     types.StringType,
						"url":        types.StringType,
						"address":    types.StringType,
						"port_fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"ip_address": types.StringType,
							"subnet_id":  types.StringType,
						}}},
						"port_mac_address":   types.StringType,
						"subnet_cidr":        types.StringType,
						"subnet_description": types.StringType,
						"subnet_name":        types.StringType,
						"subnet_uuid":        types.StringType,
					}}, items)
					data.FloatingIps = listVal
				}
			} else {
				if data.FloatingIps.IsUnknown() {
					data.FloatingIps = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"ip_address": types.StringType,
						"subnet":     types.StringType,
						"url":        types.StringType,
						"address":    types.StringType,
						"port_fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"ip_address": types.StringType,
							"subnet_id":  types.StringType,
						}}},
						"port_mac_address":   types.StringType,
						"subnet_cidr":        types.StringType,
						"subnet_description": types.StringType,
						"subnet_name":        types.StringType,
						"subnet_uuid":        types.StringType,
					}})
				}
			}
			if val, ok := sourceMap["hypervisor_hostname"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.HypervisorHostname = types.StringValue(str)
				}
			} else {
				if data.HypervisorHostname.IsUnknown() {
					data.HypervisorHostname = types.StringNull()
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
			if val, ok := sourceMap["image_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ImageName = types.StringValue(str)
				}
			} else {
				if data.ImageName.IsUnknown() {
					data.ImageName = types.StringNull()
				}
			}
			if val, ok := sourceMap["internal_ips"]; ok && val != nil {
				// List of strings (or flattened objects)
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if str, ok := item.(string); ok {
							items = append(items, types.StringValue(str))
						} else if obj, ok := item.(map[string]interface{}); ok {
							// Flattening logic: extract URL or UUID
							if url, ok := obj["url"].(string); ok {
								parts := strings.Split(strings.TrimRight(url, "/"), "/")
								uuid := parts[len(parts)-1]
								items = append(items, types.StringValue(uuid))
							} else if uuid, ok := obj["uuid"].(string); ok {
								items = append(items, types.StringValue(uuid))
							} else if name, ok := obj["name"].(string); ok {
								items = append(items, types.StringValue(name))
							}
						}
					}
					listVal, _ := types.ListValue(types.StringType, items)
					data.InternalIps = listVal
				}
			} else {
				if data.InternalIps.IsUnknown() {
					data.InternalIps = types.ListNull(types.StringType)
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
			if val, ok := sourceMap["key_fingerprint"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.KeyFingerprint = types.StringValue(str)
				}
			} else {
				if data.KeyFingerprint.IsUnknown() {
					data.KeyFingerprint = types.StringNull()
				}
			}
			if val, ok := sourceMap["key_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.KeyName = types.StringValue(str)
				}
			} else {
				if data.KeyName.IsUnknown() {
					data.KeyName = types.StringNull()
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
			if val, ok := sourceMap["min_disk"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.MinDisk = types.Int64Value(int64(num))
				}
			} else {
				if data.MinDisk.IsUnknown() {
					data.MinDisk = types.Int64Null()
				}
			}
			if val, ok := sourceMap["min_ram"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.MinRam = types.Int64Value(int64(num))
				}
			} else {
				if data.MinRam.IsUnknown() {
					data.MinRam = types.Int64Null()
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
			if val, ok := sourceMap["ports"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"ip_address": types.StringType,
									"subnet_id":  types.StringType,
								}}},
								"port":   types.StringType,
								"subnet": types.StringType,
								"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"mac_address": types.StringType,
								}}},
								"device_id":    types.StringType,
								"device_owner": types.StringType,
								"mac_address":  types.StringType,
								"security_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"access_url":                 types.StringType,
									"backend_id":                 types.StringType,
									"created":                    types.StringType,
									"customer":                   types.StringType,
									"customer_abbreviation":      types.StringType,
									"customer_name":              types.StringType,
									"customer_native_name":       types.StringType,
									"customer_uuid":              types.StringType,
									"description":                types.StringType,
									"error_message":              types.StringType,
									"error_traceback":            types.StringType,
									"is_limit_based":             types.BoolType,
									"is_usage_based":             types.BoolType,
									"marketplace_category_name":  types.StringType,
									"marketplace_category_uuid":  types.StringType,
									"marketplace_offering_name":  types.StringType,
									"marketplace_offering_uuid":  types.StringType,
									"marketplace_plan_uuid":      types.StringType,
									"marketplace_resource_state": types.StringType,
									"marketplace_resource_uuid":  types.StringType,
									"modified":                   types.StringType,
									"name":                       types.StringType,
									"project":                    types.StringType,
									"project_name":               types.StringType,
									"project_uuid":               types.StringType,
									"resource_type":              types.StringType,
									"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
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
									"service_name":                   types.StringType,
									"service_settings":               types.StringType,
									"service_settings_error_message": types.StringType,
									"service_settings_state":         types.StringType,
									"service_settings_uuid":          types.StringType,
									"state":                          types.StringType,
									"tenant":                         types.StringType,
									"tenant_name":                    types.StringType,
									"tenant_uuid":                    types.StringType,
									"url":                            types.StringType,
								}}},
								"subnet_cidr":        types.StringType,
								"subnet_description": types.StringType,
								"subnet_name":        types.StringType,
								"subnet_uuid":        types.StringType,
								"url":                types.StringType,
							}
							attrValues := map[string]attr.Value{
								"fixed_ips": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"ip_address": types.StringType,
									"subnet_id":  types.StringType,
								}}}.ElemType),
								"port": func() attr.Value {
									if v, ok := objMap["port"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"subnet": func() attr.Value {
									if v, ok := objMap["subnet"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"allowed_address_pairs": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"mac_address": types.StringType,
								}}}.ElemType),
								"device_id": func() attr.Value {
									if v, ok := objMap["device_id"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"device_owner": func() attr.Value {
									if v, ok := objMap["device_owner"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"mac_address": func() attr.Value {
									if v, ok := objMap["mac_address"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"security_groups": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"access_url":                 types.StringType,
									"backend_id":                 types.StringType,
									"created":                    types.StringType,
									"customer":                   types.StringType,
									"customer_abbreviation":      types.StringType,
									"customer_name":              types.StringType,
									"customer_native_name":       types.StringType,
									"customer_uuid":              types.StringType,
									"description":                types.StringType,
									"error_message":              types.StringType,
									"error_traceback":            types.StringType,
									"is_limit_based":             types.BoolType,
									"is_usage_based":             types.BoolType,
									"marketplace_category_name":  types.StringType,
									"marketplace_category_uuid":  types.StringType,
									"marketplace_offering_name":  types.StringType,
									"marketplace_offering_uuid":  types.StringType,
									"marketplace_plan_uuid":      types.StringType,
									"marketplace_resource_state": types.StringType,
									"marketplace_resource_uuid":  types.StringType,
									"modified":                   types.StringType,
									"name":                       types.StringType,
									"project":                    types.StringType,
									"project_name":               types.StringType,
									"project_uuid":               types.StringType,
									"resource_type":              types.StringType,
									"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
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
									"service_name":                   types.StringType,
									"service_settings":               types.StringType,
									"service_settings_error_message": types.StringType,
									"service_settings_state":         types.StringType,
									"service_settings_uuid":          types.StringType,
									"state":                          types.StringType,
									"tenant":                         types.StringType,
									"tenant_name":                    types.StringType,
									"tenant_uuid":                    types.StringType,
									"url":                            types.StringType,
								}}}.ElemType),
								"subnet_cidr": func() attr.Value {
									if v, ok := objMap["subnet_cidr"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"subnet_description": func() attr.Value {
									if v, ok := objMap["subnet_description"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"subnet_name": func() attr.Value {
									if v, ok := objMap["subnet_name"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"subnet_uuid": func() attr.Value {
									if v, ok := objMap["subnet_uuid"].(string); ok {
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
						"fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"ip_address": types.StringType,
							"subnet_id":  types.StringType,
						}}},
						"port":   types.StringType,
						"subnet": types.StringType,
						"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"mac_address": types.StringType,
						}}},
						"device_id":    types.StringType,
						"device_owner": types.StringType,
						"mac_address":  types.StringType,
						"security_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"access_url":                 types.StringType,
							"backend_id":                 types.StringType,
							"created":                    types.StringType,
							"customer":                   types.StringType,
							"customer_abbreviation":      types.StringType,
							"customer_name":              types.StringType,
							"customer_native_name":       types.StringType,
							"customer_uuid":              types.StringType,
							"description":                types.StringType,
							"error_message":              types.StringType,
							"error_traceback":            types.StringType,
							"is_limit_based":             types.BoolType,
							"is_usage_based":             types.BoolType,
							"marketplace_category_name":  types.StringType,
							"marketplace_category_uuid":  types.StringType,
							"marketplace_offering_name":  types.StringType,
							"marketplace_offering_uuid":  types.StringType,
							"marketplace_plan_uuid":      types.StringType,
							"marketplace_resource_state": types.StringType,
							"marketplace_resource_uuid":  types.StringType,
							"modified":                   types.StringType,
							"name":                       types.StringType,
							"project":                    types.StringType,
							"project_name":               types.StringType,
							"project_uuid":               types.StringType,
							"resource_type":              types.StringType,
							"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
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
							"service_name":                   types.StringType,
							"service_settings":               types.StringType,
							"service_settings_error_message": types.StringType,
							"service_settings_state":         types.StringType,
							"service_settings_uuid":          types.StringType,
							"state":                          types.StringType,
							"tenant":                         types.StringType,
							"tenant_name":                    types.StringType,
							"tenant_uuid":                    types.StringType,
							"url":                            types.StringType,
						}}},
						"subnet_cidr":        types.StringType,
						"subnet_description": types.StringType,
						"subnet_name":        types.StringType,
						"subnet_uuid":        types.StringType,
						"url":                types.StringType,
					}}, items)
					data.Ports = listVal
				}
			} else {
				if data.Ports.IsUnknown() {
					data.Ports = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"ip_address": types.StringType,
							"subnet_id":  types.StringType,
						}}},
						"port":   types.StringType,
						"subnet": types.StringType,
						"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"mac_address": types.StringType,
						}}},
						"device_id":    types.StringType,
						"device_owner": types.StringType,
						"mac_address":  types.StringType,
						"security_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"access_url":                 types.StringType,
							"backend_id":                 types.StringType,
							"created":                    types.StringType,
							"customer":                   types.StringType,
							"customer_abbreviation":      types.StringType,
							"customer_name":              types.StringType,
							"customer_native_name":       types.StringType,
							"customer_uuid":              types.StringType,
							"description":                types.StringType,
							"error_message":              types.StringType,
							"error_traceback":            types.StringType,
							"is_limit_based":             types.BoolType,
							"is_usage_based":             types.BoolType,
							"marketplace_category_name":  types.StringType,
							"marketplace_category_uuid":  types.StringType,
							"marketplace_offering_name":  types.StringType,
							"marketplace_offering_uuid":  types.StringType,
							"marketplace_plan_uuid":      types.StringType,
							"marketplace_resource_state": types.StringType,
							"marketplace_resource_uuid":  types.StringType,
							"modified":                   types.StringType,
							"name":                       types.StringType,
							"project":                    types.StringType,
							"project_name":               types.StringType,
							"project_uuid":               types.StringType,
							"resource_type":              types.StringType,
							"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
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
							"service_name":                   types.StringType,
							"service_settings":               types.StringType,
							"service_settings_error_message": types.StringType,
							"service_settings_state":         types.StringType,
							"service_settings_uuid":          types.StringType,
							"state":                          types.StringType,
							"tenant":                         types.StringType,
							"tenant_name":                    types.StringType,
							"tenant_uuid":                    types.StringType,
							"url":                            types.StringType,
						}}},
						"subnet_cidr":        types.StringType,
						"subnet_description": types.StringType,
						"subnet_name":        types.StringType,
						"subnet_uuid":        types.StringType,
						"url":                types.StringType,
					}})
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
			if val, ok := sourceMap["ram"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.Ram = types.Int64Value(int64(num))
				}
			} else {
				if data.Ram.IsUnknown() {
					data.Ram = types.Int64Null()
				}
			}
			if val, ok := sourceMap["release_floating_ips"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.ReleaseFloatingIps = types.BoolValue(b)
				}
			} else {
				if data.ReleaseFloatingIps.IsUnknown() {
					data.ReleaseFloatingIps = types.BoolNull()
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
			if val, ok := sourceMap["runtime_state"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.RuntimeState = types.StringValue(str)
				}
			} else {
				if data.RuntimeState.IsUnknown() {
					data.RuntimeState = types.StringNull()
				}
			}
			if val, ok := sourceMap["security_groups"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"url":         types.StringType,
								"description": types.StringType,
								"name":        types.StringType,
								"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"cidr":              types.StringType,
									"description":       types.StringType,
									"direction":         types.StringType,
									"ethertype":         types.StringType,
									"from_port":         types.Int64Type,
									"id":                types.Int64Type,
									"protocol":          types.StringType,
									"remote_group_name": types.StringType,
									"remote_group_uuid": types.StringType,
									"to_port":           types.Int64Type,
								}}},
								"state": types.StringType,
							}
							attrValues := map[string]attr.Value{
								"url": func() attr.Value {
									if v, ok := objMap["url"].(string); ok {
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
								"name": func() attr.Value {
									if v, ok := objMap["name"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"rules": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
									"cidr":              types.StringType,
									"description":       types.StringType,
									"direction":         types.StringType,
									"ethertype":         types.StringType,
									"from_port":         types.Int64Type,
									"id":                types.Int64Type,
									"protocol":          types.StringType,
									"remote_group_name": types.StringType,
									"remote_group_uuid": types.StringType,
									"to_port":           types.Int64Type,
								}}}.ElemType),
								"state": func() attr.Value {
									if v, ok := objMap["state"].(string); ok {
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
						"url":         types.StringType,
						"description": types.StringType,
						"name":        types.StringType,
						"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"cidr":              types.StringType,
							"description":       types.StringType,
							"direction":         types.StringType,
							"ethertype":         types.StringType,
							"from_port":         types.Int64Type,
							"id":                types.Int64Type,
							"protocol":          types.StringType,
							"remote_group_name": types.StringType,
							"remote_group_uuid": types.StringType,
							"to_port":           types.Int64Type,
						}}},
						"state": types.StringType,
					}}, items)
					data.SecurityGroups = listVal
				}
			} else {
				if data.SecurityGroups.IsUnknown() {
					data.SecurityGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"url":         types.StringType,
						"description": types.StringType,
						"name":        types.StringType,
						"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
							"cidr":              types.StringType,
							"description":       types.StringType,
							"direction":         types.StringType,
							"ethertype":         types.StringType,
							"from_port":         types.Int64Type,
							"id":                types.Int64Type,
							"protocol":          types.StringType,
							"remote_group_name": types.StringType,
							"remote_group_uuid": types.StringType,
							"to_port":           types.Int64Type,
						}}},
						"state": types.StringType,
					}})
				}
			}
			if val, ok := sourceMap["server_group"]; ok && val != nil {
				// Nested object
				if objMap, ok := val.(map[string]interface{}); ok {
					_ = objMap // Avoid unused variable if properties are empty
					attrTypes := map[string]attr.Type{
						"name":   types.StringType,
						"policy": types.StringType,
						"state":  types.StringType,
						"url":    types.StringType,
					}
					attrValues := map[string]attr.Value{
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"policy": func() attr.Value {
							if v, ok := objMap["policy"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"state": func() attr.Value {
							if v, ok := objMap["state"].(string); ok {
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
					data.ServerGroup = objVal
				}
			} else {
				if data.ServerGroup.IsUnknown() {
					data.ServerGroup = types.ObjectNull(map[string]attr.Type{
						"name":   types.StringType,
						"policy": types.StringType,
						"state":  types.StringType,
						"url":    types.StringType,
					})
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
			if val, ok := sourceMap["ssh_public_key"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.SshPublicKey = types.StringValue(str)
				}
			} else {
				if data.SshPublicKey.IsUnknown() {
					data.SshPublicKey = types.StringNull()
				}
			}
			if val, ok := sourceMap["start_time"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.StartTime = types.StringValue(str)
				}
			} else {
				if data.StartTime.IsUnknown() {
					data.StartTime = types.StringNull()
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
			if val, ok := sourceMap["system_volume_size"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.SystemVolumeSize = types.Int64Value(int64(num))
				}
			} else {
				if data.SystemVolumeSize.IsUnknown() {
					data.SystemVolumeSize = types.Int64Null()
				}
			}
			if val, ok := sourceMap["system_volume_type"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.SystemVolumeType = types.StringValue(str)
				}
			} else {
				if data.SystemVolumeType.IsUnknown() {
					data.SystemVolumeType = types.StringNull()
				}
			}
			if val, ok := sourceMap["tenant"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Tenant = types.StringValue(str)
				}
			} else {
				if data.Tenant.IsUnknown() {
					data.Tenant = types.StringNull()
				}
			}
			if val, ok := sourceMap["tenant_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.TenantUuid = types.StringValue(str)
				}
			} else {
				if data.TenantUuid.IsUnknown() {
					data.TenantUuid = types.StringNull()
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
			if val, ok := sourceMap["user_data"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.UserData = types.StringValue(str)
				}
			} else {
				if data.UserData.IsUnknown() {
					data.UserData = types.StringNull()
				}
			}
			if val, ok := sourceMap["volumes"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"bootable":                  types.BoolType,
								"device":                    types.StringType,
								"image_name":                types.StringType,
								"marketplace_resource_uuid": types.StringType,
								"name":                      types.StringType,
								"resource_type":             types.StringType,
								"size":                      types.Int64Type,
								"state":                     types.StringType,
								"type":                      types.StringType,
								"type_name":                 types.StringType,
								"url":                       types.StringType,
							}
							attrValues := map[string]attr.Value{
								"bootable": func() attr.Value {
									if v, ok := objMap["bootable"].(bool); ok {
										return types.BoolValue(v)
									}
									return types.BoolNull()
								}(),
								"device": func() attr.Value {
									if v, ok := objMap["device"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"image_name": func() attr.Value {
									if v, ok := objMap["image_name"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"marketplace_resource_uuid": func() attr.Value {
									if v, ok := objMap["marketplace_resource_uuid"].(string); ok {
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
								"resource_type": func() attr.Value {
									if v, ok := objMap["resource_type"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"size": func() attr.Value {
									if v, ok := objMap["size"].(float64); ok {
										return types.Int64Value(int64(v))
									}
									return types.Int64Null()
								}(),
								"state": func() attr.Value {
									if v, ok := objMap["state"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"type": func() attr.Value {
									if v, ok := objMap["type"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"type_name": func() attr.Value {
									if v, ok := objMap["type_name"].(string); ok {
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
						"bootable":                  types.BoolType,
						"device":                    types.StringType,
						"image_name":                types.StringType,
						"marketplace_resource_uuid": types.StringType,
						"name":                      types.StringType,
						"resource_type":             types.StringType,
						"size":                      types.Int64Type,
						"state":                     types.StringType,
						"type":                      types.StringType,
						"type_name":                 types.StringType,
						"url":                       types.StringType,
					}}, items)
					data.Volumes = listVal
				}
			} else {
				if data.Volumes.IsUnknown() {
					data.Volumes = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"bootable":                  types.BoolType,
						"device":                    types.StringType,
						"image_name":                types.StringType,
						"marketplace_resource_uuid": types.StringType,
						"name":                      types.StringType,
						"resource_type":             types.StringType,
						"size":                      types.Int64Type,
						"state":                     types.StringType,
						"type":                      types.StringType,
						"type_name":                 types.StringType,
						"url":                       types.StringType,
					}})
				}
			}

			// Map filter parameters from response if available
			if val, ok := sourceMap["attach_volume_uuid"]; ok && val != nil {
			}
			if val, ok := sourceMap["availability_zone_name"]; ok && val != nil {
			}
			if val, ok := sourceMap["backend_id"]; ok && val != nil {
			}
			if val, ok := sourceMap["can_manage"]; ok && val != nil {
			}
			if val, ok := sourceMap["customer"]; ok && val != nil {
			}
			if val, ok := sourceMap["customer_abbreviation"]; ok && val != nil {
			}
			if val, ok := sourceMap["customer_name"]; ok && val != nil {
			}
			if val, ok := sourceMap["customer_native_name"]; ok && val != nil {
			}
			if val, ok := sourceMap["customer_uuid"]; ok && val != nil {
			}
			if val, ok := sourceMap["description"]; ok && val != nil {
			}
			if val, ok := sourceMap["external_ip"]; ok && val != nil {
			}
			if val, ok := sourceMap["name"]; ok && val != nil {
			}
			if val, ok := sourceMap["name_exact"]; ok && val != nil {
			}
			if val, ok := sourceMap["page"]; ok && val != nil {
			}
			if val, ok := sourceMap["page_size"]; ok && val != nil {
			}
			if val, ok := sourceMap["project"]; ok && val != nil {
			}
			if val, ok := sourceMap["project_name"]; ok && val != nil {
			}
			if val, ok := sourceMap["project_uuid"]; ok && val != nil {
			}
			if val, ok := sourceMap["query"]; ok && val != nil {
			}
			if val, ok := sourceMap["runtime_state"]; ok && val != nil {
			}
			if val, ok := sourceMap["service_settings_name"]; ok && val != nil {
			}
			if val, ok := sourceMap["service_settings_uuid"]; ok && val != nil {
			}
			if val, ok := sourceMap["tenant"]; ok && val != nil {
			}
			if val, ok := sourceMap["tenant_uuid"]; ok && val != nil {
			}
			if val, ok := sourceMap["uuid"]; ok && val != nil {
			}

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
