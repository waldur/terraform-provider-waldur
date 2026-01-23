package tenant

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

var _ list.ListResource = &OpenstackTenantList{}

type OpenstackTenantList struct {
	client *Client
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
			"service_settings_name": schema.StringAttribute{
				Description: "Service settings name",
				Optional:    true,
			},
			"service_settings_uuid": schema.StringAttribute{
				Description: "Service settings UUID",
				Optional:    true,
			},
			"uuid": schema.StringAttribute{
				Description: "UUID",
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

	l.client = NewClient(client)
}

type OpenstackTenantListModel struct {
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
	listResult, err := l.client.ListOpenstackTenant(ctx, filters)
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
			var data OpenstackTenantResourceModel
			model := &data

			var diags diag.Diagnostics

			data.UUID = types.StringPointerValue(apiResp.UUID)
			model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
			model.AvailabilityZone = types.StringPointerValue(apiResp.AvailabilityZone)
			model.BackendId = types.StringPointerValue(apiResp.BackendId)
			model.Created = types.StringPointerValue(apiResp.Created)
			model.Customer = types.StringPointerValue(apiResp.Customer)
			model.CustomerAbbreviation = types.StringPointerValue(apiResp.CustomerAbbreviation)
			model.CustomerName = types.StringPointerValue(apiResp.CustomerName)
			model.CustomerNativeName = types.StringPointerValue(apiResp.CustomerNativeName)
			model.CustomerUuid = types.StringPointerValue(apiResp.CustomerUuid)
			model.DefaultVolumeTypeName = types.StringPointerValue(apiResp.DefaultVolumeTypeName)
			model.Description = types.StringPointerValue(apiResp.Description)
			model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
			model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
			model.ExternalNetworkId = types.StringPointerValue(apiResp.ExternalNetworkId)
			model.InternalNetworkId = types.StringPointerValue(apiResp.InternalNetworkId)
			model.IsLimitBased = types.BoolPointerValue(apiResp.IsLimitBased)
			model.IsUsageBased = types.BoolPointerValue(apiResp.IsUsageBased)
			model.MarketplaceCategoryName = types.StringPointerValue(apiResp.MarketplaceCategoryName)
			model.MarketplaceCategoryUuid = types.StringPointerValue(apiResp.MarketplaceCategoryUuid)
			model.MarketplaceOfferingName = types.StringPointerValue(apiResp.MarketplaceOfferingName)
			model.MarketplaceOfferingUuid = types.StringPointerValue(apiResp.MarketplaceOfferingUuid)
			model.MarketplacePlanUuid = types.StringPointerValue(apiResp.MarketplacePlanUuid)
			model.MarketplaceResourceState = types.StringPointerValue(apiResp.MarketplaceResourceState)
			model.MarketplaceResourceUuid = types.StringPointerValue(apiResp.MarketplaceResourceUuid)
			model.Modified = types.StringPointerValue(apiResp.Modified)
			model.Name = types.StringPointerValue(apiResp.Name)
			model.Project = types.StringPointerValue(apiResp.Project)
			model.ProjectName = types.StringPointerValue(apiResp.ProjectName)
			model.ProjectUuid = types.StringPointerValue(apiResp.ProjectUuid)

			{
				listValQuotas, listDiagsQuotas := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
					"limit": types.Int64Type,
					"name":  types.StringType,
					"usage": types.Int64Type,
				}}, apiResp.Quotas)
				diags.Append(listDiagsQuotas...)
				model.Quotas = listValQuotas
			}
			model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
			model.ServiceName = types.StringPointerValue(apiResp.ServiceName)
			model.ServiceSettings = types.StringPointerValue(apiResp.ServiceSettings)
			model.ServiceSettingsErrorMessage = types.StringPointerValue(apiResp.ServiceSettingsErrorMessage)
			model.ServiceSettingsState = types.StringPointerValue(apiResp.ServiceSettingsState)
			model.ServiceSettingsUuid = types.StringPointerValue(apiResp.ServiceSettingsUuid)
			model.SkipCreationOfDefaultRouter = types.BoolPointerValue(apiResp.SkipCreationOfDefaultRouter)
			model.State = types.StringPointerValue(apiResp.State)
			model.Url = types.StringPointerValue(apiResp.Url)
			model.UserPassword = types.StringPointerValue(apiResp.UserPassword)
			model.UserUsername = types.StringPointerValue(apiResp.UserUsername)

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
