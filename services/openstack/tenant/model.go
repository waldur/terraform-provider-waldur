package tenant

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

func OpenStackSecurityGroupRuleCreateRequestType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"cidr":         types.StringType,
		"description":  types.StringType,
		"direction":    types.StringType,
		"ethertype":    types.StringType,
		"from_port":    types.Int64Type,
		"protocol":     types.StringType,
		"remote_group": types.StringType,
		"to_port":      types.Int64Type,
	}}
}
func OpenStackTenantSecurityGroupRequestType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"description": types.StringType,
		"name":        types.StringType,
		"rules":       types.ListType{ElemType: OpenStackSecurityGroupRuleCreateRequestType()},
	}}
}
func QuotaType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"limit": types.Int64Type,
		"name":  types.StringType,
		"usage": types.Int64Type,
	}}
}

type OpenstackTenantFiltersModel struct {
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
	ServiceSettingsName  types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid  types.String `tfsdk:"service_settings_uuid"`
	Uuid                 types.String `tfsdk:"uuid"`
}

func (m *OpenstackTenantFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Openstack Tenant",
		Attributes: map[string]schema.Attribute{
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
			"external_ip": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "External IP",
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
			"service_settings_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Service settings name",
			},
			"service_settings_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Service settings UUID",
			},
			"uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "UUID",
			},
		},
	}
}

type OpenstackTenantModel struct {
	UUID                        types.String `tfsdk:"id"`
	AvailabilityZone            types.String `tfsdk:"availability_zone"`
	BackendId                   types.String `tfsdk:"backend_id"`
	Customer                    types.String `tfsdk:"customer"`
	DefaultVolumeTypeName       types.String `tfsdk:"default_volume_type_name"`
	Description                 types.String `tfsdk:"description"`
	ErrorMessage                types.String `tfsdk:"error_message"`
	ExternalNetworkId           types.String `tfsdk:"external_network_id"`
	InternalNetworkId           types.String `tfsdk:"internal_network_id"`
	MarketplaceResourceUuid     types.String `tfsdk:"marketplace_resource_uuid"`
	Name                        types.String `tfsdk:"name"`
	Project                     types.String `tfsdk:"project"`
	Quotas                      types.List   `tfsdk:"quotas"`
	ResourceType                types.String `tfsdk:"resource_type"`
	SkipCreationOfDefaultRouter types.Bool   `tfsdk:"skip_creation_of_default_router"`
	State                       types.String `tfsdk:"state"`
	Url                         types.String `tfsdk:"url"`
	UserPassword                types.String `tfsdk:"user_password"`
	UserUsername                types.String `tfsdk:"user_username"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackTenantModel) CopyFrom(ctx context.Context, apiResp OpenstackTenantResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)

	model.AvailabilityZone = common.StringPointerValue(apiResp.AvailabilityZone)

	model.BackendId = common.StringPointerValue(apiResp.BackendId)

	model.Customer = common.StringPointerValue(apiResp.Customer)

	model.DefaultVolumeTypeName = common.StringPointerValue(apiResp.DefaultVolumeTypeName)

	model.Description = common.StringPointerValue(apiResp.Description)

	model.ErrorMessage = common.StringPointerValue(apiResp.ErrorMessage)

	model.ExternalNetworkId = common.StringPointerValue(apiResp.ExternalNetworkId)

	model.InternalNetworkId = common.StringPointerValue(apiResp.InternalNetworkId)

	model.MarketplaceResourceUuid = common.StringPointerValue(apiResp.MarketplaceResourceUuid)

	model.Name = common.StringPointerValue(apiResp.Name)

	model.Project = common.StringPointerValue(apiResp.Project)

	if apiResp.Quotas != nil {
		valQuotas, diagsQuotas := types.ListValueFrom(ctx, QuotaType(), apiResp.Quotas)
		diags.Append(diagsQuotas...)
		model.Quotas = valQuotas
	} else {
		model.Quotas = types.ListNull(QuotaType())
	}

	model.ResourceType = common.StringPointerValue(apiResp.ResourceType)

	model.SkipCreationOfDefaultRouter = types.BoolPointerValue(apiResp.SkipCreationOfDefaultRouter)

	model.State = common.StringPointerValue(apiResp.State)

	model.Url = common.StringPointerValue(apiResp.Url)

	model.UserPassword = common.StringPointerValue(apiResp.UserPassword)

	model.UserUsername = common.StringPointerValue(apiResp.UserUsername)

	return diags
}
