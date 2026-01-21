package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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

// OpenstackTenantApiResponse is the API response model.
type OpenstackTenantApiResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl                   *string                         `json:"access_url" tfsdk:"access_url"`
	AvailabilityZone            *string                         `json:"availability_zone" tfsdk:"availability_zone"`
	Created                     *string                         `json:"created" tfsdk:"created"`
	DefaultVolumeTypeName       *string                         `json:"default_volume_type_name" tfsdk:"default_volume_type_name"`
	ErrorMessage                *string                         `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback              *string                         `json:"error_traceback" tfsdk:"error_traceback"`
	ExternalNetworkId           *string                         `json:"external_network_id" tfsdk:"external_network_id"`
	InternalNetworkId           *string                         `json:"internal_network_id" tfsdk:"internal_network_id"`
	IsLimitBased                *bool                           `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                           `json:"is_usage_based" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string                         `json:"marketplace_category_name" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                         `json:"marketplace_category_uuid" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                         `json:"marketplace_offering_name" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                         `json:"marketplace_offering_uuid" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                         `json:"marketplace_plan_uuid" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                         `json:"marketplace_resource_state" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                         `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                         `json:"modified" tfsdk:"modified"`
	Quotas                      []OpenstackTenantQuotasResponse `json:"quotas" tfsdk:"quotas"`
	ResourceType                *string                         `json:"resource_type" tfsdk:"resource_type"`
	ServiceName                 *string                         `json:"service_name" tfsdk:"service_name"`
	ServiceSettings             *string                         `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                         `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                         `json:"service_settings_state" tfsdk:"service_settings_state"`
	SkipCreationOfDefaultRouter *bool                           `json:"skip_creation_of_default_router" tfsdk:"skip_creation_of_default_router"`
	Url                         *string                         `json:"url" tfsdk:"url"`
	UserPassword                *string                         `json:"user_password" tfsdk:"user_password"`
	UserUsername                *string                         `json:"user_username" tfsdk:"user_username"`
}

type OpenstackTenantQuotasResponse struct {
	Limit *int64 `json:"limit" tfsdk:"limit"`
	Usage *int64 `json:"usage" tfsdk:"usage"`
}

var openstacktenant_quotasAttrTypes = map[string]attr.Type{
	"limit": types.Int64Type,
	"name":  types.StringType,
	"usage": types.Int64Type,
}
var openstacktenant_quotasObjectType = types.ObjectType{
	AttrTypes: openstacktenant_quotasAttrTypes,
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
	SkipCreationOfDefaultRouter types.Bool   `tfsdk:"skip_creation_of_default_router"`
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
			"state": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "State",
			},
			"uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "UUID",
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
			"skip_creation_of_default_router": schema.BoolAttribute{
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
		var apiResp OpenstackTenantApiResponse

		err := d.client.GetByUUID(ctx, "/api/openstack-tenants/{uuid}/", data.UUID.ValueString(), &apiResp)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Tenant",
				"An error occurred while reading the Openstack Tenant by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, apiResp, &data)...)

	} else {
		// Filter by provided parameters
		var results []OpenstackTenantApiResponse

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

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *OpenstackTenantDataSource) mapResponseToModel(ctx context.Context, apiResp OpenstackTenantApiResponse, model *OpenstackTenantDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.AvailabilityZone = types.StringPointerValue(apiResp.AvailabilityZone)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.DefaultVolumeTypeName = types.StringPointerValue(apiResp.DefaultVolumeTypeName)
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
	listValQuotas, listDiagsQuotas := types.ListValueFrom(ctx, openstacktenant_quotasObjectType, apiResp.Quotas)
	diags.Append(listDiagsQuotas...)
	model.Quotas = listValQuotas
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.ServiceName = types.StringPointerValue(apiResp.ServiceName)
	model.ServiceSettings = types.StringPointerValue(apiResp.ServiceSettings)
	model.ServiceSettingsErrorMessage = types.StringPointerValue(apiResp.ServiceSettingsErrorMessage)
	model.ServiceSettingsState = types.StringPointerValue(apiResp.ServiceSettingsState)
	model.SkipCreationOfDefaultRouter = types.BoolPointerValue(apiResp.SkipCreationOfDefaultRouter)
	model.Url = types.StringPointerValue(apiResp.Url)
	model.UserPassword = types.StringPointerValue(apiResp.UserPassword)
	model.UserUsername = types.StringPointerValue(apiResp.UserUsername)

	return diags
}
