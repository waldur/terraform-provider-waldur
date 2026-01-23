package instance

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

var _ list.ListResource = &OpenstackInstanceList{}

type OpenstackInstanceList struct {
	client *Client
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
				Description: "Filter for attachment to volume UUID",
				Optional:    true,
			},
			"availability_zone_name": schema.StringAttribute{
				Description: "Availability zone name",
				Optional:    true,
			},
			"backend_id": schema.StringAttribute{
				Description: "Backend ID",
				Optional:    true,
			},
			"can_manage": schema.BoolAttribute{
				Description: "Can manage",
				Optional:    true,
			},
			"customer": schema.StringAttribute{
				Description: "Customer UUID",
				Optional:    true,
			},
			"customer_abbreviation": schema.StringAttribute{
				Description: "Customer abbreviation",
				Optional:    true,
			},
			"customer_name": schema.StringAttribute{
				Description: "Customer name",
				Optional:    true,
			},
			"customer_native_name": schema.StringAttribute{
				Description: "Customer native name",
				Optional:    true,
			},
			"customer_uuid": schema.StringAttribute{
				Description: "Customer UUID",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description",
				Optional:    true,
			},
			"external_ip": schema.StringAttribute{
				Description: "External IP",
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
			"page": schema.Int64Attribute{
				Description: "A page number within the paginated result set.",
				Optional:    true,
			},
			"page_size": schema.Int64Attribute{
				Description: "Number of results to return per page.",
				Optional:    true,
			},
			"project": schema.StringAttribute{
				Description: "Project UUID",
				Optional:    true,
			},
			"project_name": schema.StringAttribute{
				Description: "Project name",
				Optional:    true,
			},
			"project_uuid": schema.StringAttribute{
				Description: "Project UUID",
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
				Description: "Service settings name",
				Optional:    true,
			},
			"service_settings_uuid": schema.StringAttribute{
				Description: "Service settings UUID",
				Optional:    true,
			},
			"tenant": schema.StringAttribute{
				Description: "Tenant URL",
				Optional:    true,
			},
			"tenant_uuid": schema.StringAttribute{
				Description: "Tenant UUID",
				Optional:    true,
			},
			"uuid": schema.StringAttribute{
				Description: "UUID",
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

	l.client = NewClient(client)
}

type OpenstackInstanceListModel struct {
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
	listResult, err := l.client.ListOpenstackInstance(ctx, filters)
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
			var data OpenstackInstanceResourceModel
			model := &data

			var diags diag.Diagnostics

			data.UUID = types.StringPointerValue(apiResp.UUID)
			model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
			model.Action = types.StringPointerValue(apiResp.Action)
			model.AvailabilityZone = types.StringPointerValue(apiResp.AvailabilityZone)
			model.AvailabilityZoneName = types.StringPointerValue(apiResp.AvailabilityZoneName)
			model.BackendId = types.StringPointerValue(apiResp.BackendId)
			model.ConnectDirectlyToExternalNetwork = types.BoolPointerValue(apiResp.ConnectDirectlyToExternalNetwork)
			model.Cores = types.Int64PointerValue(apiResp.Cores)
			model.Created = types.StringPointerValue(apiResp.Created)
			model.Customer = types.StringPointerValue(apiResp.Customer)
			model.CustomerAbbreviation = types.StringPointerValue(apiResp.CustomerAbbreviation)
			model.CustomerName = types.StringPointerValue(apiResp.CustomerName)
			model.CustomerNativeName = types.StringPointerValue(apiResp.CustomerNativeName)
			model.CustomerUuid = types.StringPointerValue(apiResp.CustomerUuid)
			model.Description = types.StringPointerValue(apiResp.Description)
			model.Disk = types.Int64PointerValue(apiResp.Disk)
			model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
			model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
			model.ExternalAddress, _ = types.ListValueFrom(ctx, types.StringType, apiResp.ExternalAddress)
			model.ExternalIps, _ = types.ListValueFrom(ctx, types.StringType, apiResp.ExternalIps)
			model.FlavorDisk = types.Int64PointerValue(apiResp.FlavorDisk)
			model.FlavorName = types.StringPointerValue(apiResp.FlavorName)
			listValFloatingIps, listDiagsFloatingIps := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}}, apiResp.FloatingIps)
			diags.Append(listDiagsFloatingIps...)
			model.FloatingIps = listValFloatingIps
			model.HypervisorHostname = types.StringPointerValue(apiResp.HypervisorHostname)
			model.ImageName = types.StringPointerValue(apiResp.ImageName)
			model.InternalIps, _ = types.ListValueFrom(ctx, types.StringType, apiResp.InternalIps)
			model.IsLimitBased = types.BoolPointerValue(apiResp.IsLimitBased)
			model.IsUsageBased = types.BoolPointerValue(apiResp.IsUsageBased)
			model.KeyFingerprint = types.StringPointerValue(apiResp.KeyFingerprint)
			model.KeyName = types.StringPointerValue(apiResp.KeyName)
			model.Latitude = types.Float64PointerValue(apiResp.Latitude)
			model.Longitude = types.Float64PointerValue(apiResp.Longitude)
			model.MarketplaceCategoryName = types.StringPointerValue(apiResp.MarketplaceCategoryName)
			model.MarketplaceCategoryUuid = types.StringPointerValue(apiResp.MarketplaceCategoryUuid)
			model.MarketplaceOfferingName = types.StringPointerValue(apiResp.MarketplaceOfferingName)
			model.MarketplaceOfferingUuid = types.StringPointerValue(apiResp.MarketplaceOfferingUuid)
			model.MarketplacePlanUuid = types.StringPointerValue(apiResp.MarketplacePlanUuid)
			model.MarketplaceResourceState = types.StringPointerValue(apiResp.MarketplaceResourceState)
			model.MarketplaceResourceUuid = types.StringPointerValue(apiResp.MarketplaceResourceUuid)
			model.MinDisk = types.Int64PointerValue(apiResp.MinDisk)
			model.MinRam = types.Int64PointerValue(apiResp.MinRam)
			model.Modified = types.StringPointerValue(apiResp.Modified)
			model.Name = types.StringPointerValue(apiResp.Name)
			listValPorts, listDiagsPorts := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}}, apiResp.Ports)
			diags.Append(listDiagsPorts...)
			model.Ports = listValPorts
			model.Project = types.StringPointerValue(apiResp.Project)
			model.ProjectName = types.StringPointerValue(apiResp.ProjectName)
			model.ProjectUuid = types.StringPointerValue(apiResp.ProjectUuid)
			model.Ram = types.Int64PointerValue(apiResp.Ram)
			model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
			model.RuntimeState = types.StringPointerValue(apiResp.RuntimeState)
			listValSecurityGroups, listDiagsSecurityGroups := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}}, apiResp.SecurityGroups)
			diags.Append(listDiagsSecurityGroups...)
			model.SecurityGroups = listValSecurityGroups
			if apiResp.ServerGroup != nil {
				objValServerGroup, objDiagsServerGroup := types.ObjectValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
					"name":   types.StringType,
					"policy": types.StringType,
					"state":  types.StringType,
					"url":    types.StringType,
				}}.AttrTypes, *apiResp.ServerGroup)
				diags.Append(objDiagsServerGroup...)
				model.ServerGroup = objValServerGroup
			} else {
				model.ServerGroup = types.ObjectNull(types.ObjectType{AttrTypes: map[string]attr.Type{
					"name":   types.StringType,
					"policy": types.StringType,
					"state":  types.StringType,
					"url":    types.StringType,
				}}.AttrTypes)
			}
			model.ServiceName = types.StringPointerValue(apiResp.ServiceName)
			model.ServiceSettings = types.StringPointerValue(apiResp.ServiceSettings)
			model.ServiceSettingsErrorMessage = types.StringPointerValue(apiResp.ServiceSettingsErrorMessage)
			model.ServiceSettingsState = types.StringPointerValue(apiResp.ServiceSettingsState)
			model.ServiceSettingsUuid = types.StringPointerValue(apiResp.ServiceSettingsUuid)
			model.StartTime = types.StringPointerValue(apiResp.StartTime)
			model.State = types.StringPointerValue(apiResp.State)
			model.Tenant = types.StringPointerValue(apiResp.Tenant)
			model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
			model.Url = types.StringPointerValue(apiResp.Url)
			model.UserData = types.StringPointerValue(apiResp.UserData)
			listValVolumes, listDiagsVolumes := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}}, apiResp.Volumes)
			diags.Append(listDiagsVolumes...)
			model.Volumes = listValVolumes

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
