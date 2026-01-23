package network

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

var _ list.ListResource = &OpenstackNetworkList{}

type OpenstackNetworkList struct {
	client *Client
}

func NewOpenstackNetworkList() list.ListResource {
	return &OpenstackNetworkList{}
}

func (l *OpenstackNetworkList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_network"
}

func (l *OpenstackNetworkList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
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
			"direct_only": schema.BoolAttribute{
				Description: "Direct only",
				Optional:    true,
			},
			"external_ip": schema.StringAttribute{
				Description: "External IP",
				Optional:    true,
			},
			"is_external": schema.BoolAttribute{
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
			"rbac_only": schema.BoolAttribute{
				Description: "RBAC only",
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
			"type": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"uuid": schema.StringAttribute{
				Description: "UUID",
				Optional:    true,
			},
		},
	}
}

func (l *OpenstackNetworkList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type OpenstackNetworkListModel struct {
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
	Page                 types.Int64  `tfsdk:"page"`
	PageSize             types.Int64  `tfsdk:"page_size"`
	Project              types.String `tfsdk:"project"`
	ProjectName          types.String `tfsdk:"project_name"`
	ProjectUuid          types.String `tfsdk:"project_uuid"`
	RbacOnly             types.Bool   `tfsdk:"rbac_only"`
	ServiceSettingsName  types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid  types.String `tfsdk:"service_settings_uuid"`
	Tenant               types.String `tfsdk:"tenant"`
	TenantUuid           types.String `tfsdk:"tenant_uuid"`
	Type                 types.String `tfsdk:"type"`
	Uuid                 types.String `tfsdk:"uuid"`
}

func (l *OpenstackNetworkList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config OpenstackNetworkListModel

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
	if !config.DirectOnly.IsNull() && !config.DirectOnly.IsUnknown() {
		filters["direct_only"] = fmt.Sprintf("%t", config.DirectOnly.ValueBool())
	}
	if !config.ExternalIp.IsNull() && !config.ExternalIp.IsUnknown() {
		filters["external_ip"] = config.ExternalIp.ValueString()
	}
	if !config.IsExternal.IsNull() && !config.IsExternal.IsUnknown() {
		filters["is_external"] = fmt.Sprintf("%t", config.IsExternal.ValueBool())
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
	if !config.RbacOnly.IsNull() && !config.RbacOnly.IsUnknown() {
		filters["rbac_only"] = fmt.Sprintf("%t", config.RbacOnly.ValueBool())
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
	if !config.Type.IsNull() && !config.Type.IsUnknown() {
		filters["type"] = config.Type.ValueString()
	}
	if !config.Uuid.IsNull() && !config.Uuid.IsUnknown() {
		filters["uuid"] = config.Uuid.ValueString()
	}

	// Call API
	listResult, err := l.client.ListOpenstackNetwork(ctx, filters)
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
			var data OpenstackNetworkResourceModel
			model := &data

			var diags diag.Diagnostics

			data.UUID = types.StringPointerValue(apiResp.UUID)
			model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
			model.BackendId = types.StringPointerValue(apiResp.BackendId)
			model.Created = types.StringPointerValue(apiResp.Created)
			model.Description = types.StringPointerValue(apiResp.Description)
			model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
			model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
			model.IsExternal = types.BoolPointerValue(apiResp.IsExternal)
			model.Modified = types.StringPointerValue(apiResp.Modified)
			model.Mtu = types.Int64PointerValue(apiResp.Mtu)
			model.Name = types.StringPointerValue(apiResp.Name)

			{
				listValRbacPolicies, listDiagsRbacPolicies := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
					"backend_id":         types.StringType,
					"created":            types.StringType,
					"network":            types.StringType,
					"network_name":       types.StringType,
					"policy_type":        types.StringType,
					"target_tenant":      types.StringType,
					"target_tenant_name": types.StringType,
					"url":                types.StringType,
				}}, apiResp.RbacPolicies)
				diags.Append(listDiagsRbacPolicies...)
				model.RbacPolicies = listValRbacPolicies
			}
			model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
			model.SegmentationId = types.Int64PointerValue(apiResp.SegmentationId)
			model.State = types.StringPointerValue(apiResp.State)

			{
				listValSubnets, listDiagsSubnets := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
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
				}}, apiResp.Subnets)
				diags.Append(listDiagsSubnets...)
				model.Subnets = listValSubnets
			}
			model.Tenant = types.StringPointerValue(apiResp.Tenant)
			model.TenantName = types.StringPointerValue(apiResp.TenantName)
			model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
			model.Type = types.StringPointerValue(apiResp.Type)
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
