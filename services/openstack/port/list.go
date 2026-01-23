package port

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

var _ list.ListResource = &OpenstackPortList{}

type OpenstackPortList struct {
	client *Client
}

func NewOpenstackPortList() list.ListResource {
	return &OpenstackPortList{}
}

func (l *OpenstackPortList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_port"
}

func (l *OpenstackPortList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"admin_state_up": schema.BoolAttribute{
				Description: "",
				Optional:    true,
			},
			"backend_id": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"device_id": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"device_owner": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"exclude_subnet_uuids": schema.StringAttribute{
				Description: "Exclude Subnet UUIDs (comma-separated)",
				Optional:    true,
			},
			"fixed_ips": schema.StringAttribute{
				Description: "Search by fixed IP",
				Optional:    true,
			},
			"has_device_owner": schema.BoolAttribute{
				Description: "Has device owner",
				Optional:    true,
			},
			"mac_address": schema.StringAttribute{
				Description: "",
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
			"network_name": schema.StringAttribute{
				Description: "Search by network name",
				Optional:    true,
			},
			"network_uuid": schema.StringAttribute{
				Description: "Search by network UUID",
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
				Description: "Search by name, MAC address or backend ID",
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "",
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
		},
	}
}

func (l *OpenstackPortList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type OpenstackPortListModel struct {
	AdminStateUp       types.Bool   `tfsdk:"admin_state_up"`
	BackendId          types.String `tfsdk:"backend_id"`
	DeviceId           types.String `tfsdk:"device_id"`
	DeviceOwner        types.String `tfsdk:"device_owner"`
	ExcludeSubnetUuids types.String `tfsdk:"exclude_subnet_uuids"`
	FixedIps           types.String `tfsdk:"fixed_ips"`
	HasDeviceOwner     types.Bool   `tfsdk:"has_device_owner"`
	MacAddress         types.String `tfsdk:"mac_address"`
	Name               types.String `tfsdk:"name"`
	NameExact          types.String `tfsdk:"name_exact"`
	NetworkName        types.String `tfsdk:"network_name"`
	NetworkUuid        types.String `tfsdk:"network_uuid"`
	Page               types.Int64  `tfsdk:"page"`
	PageSize           types.Int64  `tfsdk:"page_size"`
	Query              types.String `tfsdk:"query"`
	Status             types.String `tfsdk:"status"`
	Tenant             types.String `tfsdk:"tenant"`
	TenantUuid         types.String `tfsdk:"tenant_uuid"`
}

func (l *OpenstackPortList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config OpenstackPortListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := make(map[string]string)
	if !config.AdminStateUp.IsNull() && !config.AdminStateUp.IsUnknown() {
		filters["admin_state_up"] = fmt.Sprintf("%t", config.AdminStateUp.ValueBool())
	}
	if !config.BackendId.IsNull() && !config.BackendId.IsUnknown() {
		filters["backend_id"] = config.BackendId.ValueString()
	}
	if !config.DeviceId.IsNull() && !config.DeviceId.IsUnknown() {
		filters["device_id"] = config.DeviceId.ValueString()
	}
	if !config.DeviceOwner.IsNull() && !config.DeviceOwner.IsUnknown() {
		filters["device_owner"] = config.DeviceOwner.ValueString()
	}
	if !config.ExcludeSubnetUuids.IsNull() && !config.ExcludeSubnetUuids.IsUnknown() {
		filters["exclude_subnet_uuids"] = config.ExcludeSubnetUuids.ValueString()
	}
	if !config.FixedIps.IsNull() && !config.FixedIps.IsUnknown() {
		filters["fixed_ips"] = config.FixedIps.ValueString()
	}
	if !config.HasDeviceOwner.IsNull() && !config.HasDeviceOwner.IsUnknown() {
		filters["has_device_owner"] = fmt.Sprintf("%t", config.HasDeviceOwner.ValueBool())
	}
	if !config.MacAddress.IsNull() && !config.MacAddress.IsUnknown() {
		filters["mac_address"] = config.MacAddress.ValueString()
	}
	if !config.Name.IsNull() && !config.Name.IsUnknown() {
		filters["name"] = config.Name.ValueString()
	}
	if !config.NameExact.IsNull() && !config.NameExact.IsUnknown() {
		filters["name_exact"] = config.NameExact.ValueString()
	}
	if !config.NetworkName.IsNull() && !config.NetworkName.IsUnknown() {
		filters["network_name"] = config.NetworkName.ValueString()
	}
	if !config.NetworkUuid.IsNull() && !config.NetworkUuid.IsUnknown() {
		filters["network_uuid"] = config.NetworkUuid.ValueString()
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
	if !config.Status.IsNull() && !config.Status.IsUnknown() {
		filters["status"] = config.Status.ValueString()
	}
	if !config.Tenant.IsNull() && !config.Tenant.IsUnknown() {
		filters["tenant"] = config.Tenant.ValueString()
	}
	if !config.TenantUuid.IsNull() && !config.TenantUuid.IsUnknown() {
		filters["tenant_uuid"] = config.TenantUuid.ValueString()
	}

	// Call API
	listResult, err := l.client.ListOpenstackPort(ctx, filters)
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
			var data OpenstackPortResourceModel
			model := &data

			var diags diag.Diagnostics

			data.UUID = types.StringPointerValue(apiResp.UUID)
			model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
			model.AdminStateUp = types.BoolPointerValue(apiResp.AdminStateUp)

			{
				listValAllowedAddressPairs, listDiagsAllowedAddressPairs := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
					"ip_address":  types.StringType,
					"mac_address": types.StringType,
				}}, apiResp.AllowedAddressPairs)
				diags.Append(listDiagsAllowedAddressPairs...)
				model.AllowedAddressPairs = listValAllowedAddressPairs
			}
			model.BackendId = types.StringPointerValue(apiResp.BackendId)
			model.Created = types.StringPointerValue(apiResp.Created)
			model.Description = types.StringPointerValue(apiResp.Description)
			model.DeviceId = types.StringPointerValue(apiResp.DeviceId)
			model.DeviceOwner = types.StringPointerValue(apiResp.DeviceOwner)
			model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
			model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)

			{
				listValFixedIps, listDiagsFixedIps := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
					"ip_address": types.StringType,
					"subnet_id":  types.StringType,
				}}, apiResp.FixedIps)
				diags.Append(listDiagsFixedIps...)
				model.FixedIps = listValFixedIps
			}
			model.FloatingIps, _ = types.ListValueFrom(ctx, types.StringType, apiResp.FloatingIps)
			model.MacAddress = types.StringPointerValue(apiResp.MacAddress)
			model.Modified = types.StringPointerValue(apiResp.Modified)
			model.Name = types.StringPointerValue(apiResp.Name)
			model.Network = types.StringPointerValue(apiResp.Network)
			model.NetworkName = types.StringPointerValue(apiResp.NetworkName)
			model.NetworkUuid = types.StringPointerValue(apiResp.NetworkUuid)
			model.PortSecurityEnabled = types.BoolPointerValue(apiResp.PortSecurityEnabled)
			model.ResourceType = types.StringPointerValue(apiResp.ResourceType)

			{
				listValSecurityGroups, listDiagsSecurityGroups := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
					"name": types.StringType,
				}}, apiResp.SecurityGroups)
				diags.Append(listDiagsSecurityGroups...)
				model.SecurityGroups = listValSecurityGroups
			}
			model.State = types.StringPointerValue(apiResp.State)
			model.Status = types.StringPointerValue(apiResp.Status)
			model.Tenant = types.StringPointerValue(apiResp.Tenant)
			model.TenantName = types.StringPointerValue(apiResp.TenantName)
			model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
			model.Url = types.StringPointerValue(apiResp.Url)

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
