package tenant

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackTenantDataSource{}

func NewOpenstackTenantDataSource() datasource.DataSource {
	return &OpenstackTenantDataSource{}
}

// OpenstackTenantDataSource defines the data source implementation.
type OpenstackTenantDataSource struct {
	client *Client
}

// OpenstackTenantFiltersModel contains the filter parameters for querying.
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
	State                types.String `tfsdk:"state"`
	Uuid                 types.String `tfsdk:"uuid"`
}

type OpenstackTenantDataSourceModel struct {
	UUID                        types.String                 `tfsdk:"id"`
	Filters                     *OpenstackTenantFiltersModel `tfsdk:"filters"`
	AccessUrl                   types.String                 `tfsdk:"access_url"`
	AvailabilityZone            types.String                 `tfsdk:"availability_zone"`
	BackendId                   types.String                 `tfsdk:"backend_id"`
	Created                     types.String                 `tfsdk:"created"`
	Customer                    types.String                 `tfsdk:"customer"`
	CustomerAbbreviation        types.String                 `tfsdk:"customer_abbreviation"`
	CustomerName                types.String                 `tfsdk:"customer_name"`
	CustomerNativeName          types.String                 `tfsdk:"customer_native_name"`
	CustomerUuid                types.String                 `tfsdk:"customer_uuid"`
	DefaultVolumeTypeName       types.String                 `tfsdk:"default_volume_type_name"`
	Description                 types.String                 `tfsdk:"description"`
	ErrorMessage                types.String                 `tfsdk:"error_message"`
	ErrorTraceback              types.String                 `tfsdk:"error_traceback"`
	ExternalNetworkId           types.String                 `tfsdk:"external_network_id"`
	InternalNetworkId           types.String                 `tfsdk:"internal_network_id"`
	IsLimitBased                types.Bool                   `tfsdk:"is_limit_based"`
	IsUsageBased                types.Bool                   `tfsdk:"is_usage_based"`
	MarketplaceCategoryName     types.String                 `tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     types.String                 `tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     types.String                 `tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     types.String                 `tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         types.String                 `tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    types.String                 `tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     types.String                 `tfsdk:"marketplace_resource_uuid"`
	Modified                    types.String                 `tfsdk:"modified"`
	Name                        types.String                 `tfsdk:"name"`
	Project                     types.String                 `tfsdk:"project"`
	ProjectName                 types.String                 `tfsdk:"project_name"`
	ProjectUuid                 types.String                 `tfsdk:"project_uuid"`
	Quotas                      types.List                   `tfsdk:"quotas"`
	ResourceType                types.String                 `tfsdk:"resource_type"`
	ServiceName                 types.String                 `tfsdk:"service_name"`
	ServiceSettings             types.String                 `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage types.String                 `tfsdk:"service_settings_error_message"`
	ServiceSettingsState        types.String                 `tfsdk:"service_settings_state"`
	ServiceSettingsUuid         types.String                 `tfsdk:"service_settings_uuid"`
	SkipCreationOfDefaultRouter types.Bool                   `tfsdk:"skip_creation_of_default_router"`
	State                       types.String                 `tfsdk:"state"`
	Url                         types.String                 `tfsdk:"url"`
	UserPassword                types.String                 `tfsdk:"user_password"`
	UserUsername                types.String                 `tfsdk:"user_username"`
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
			"filters": schema.SingleNestedAttribute{
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
					"state": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "State Allowed values: `CREATING`, `CREATION_SCHEDULED`, `DELETING`, `DELETION_SCHEDULED`, `ERRED`, `OK`, `UPDATE_SCHEDULED`, `UPDATING`.",
						Validators: []validator.String{
							stringvalidator.OneOf("CREATING", "CREATION_SCHEDULED", "DELETING", "DELETION_SCHEDULED", "ERRED", "OK", "UPDATE_SCHEDULED", "UPDATING"),
						},
					},
					"uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "UUID",
					},
				},
			},
			"access_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Access url",
			},
			"availability_zone": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Optional availability group. Will be used for all instances provisioned in this tenant",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of tenant in the OpenStack backend",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Created",
			},
			"customer": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Customer",
			},
			"customer_abbreviation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Customer abbreviation",
			},
			"customer_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the customer",
			},
			"customer_native_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the customer native",
			},
			"customer_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the customer",
			},
			"default_volume_type_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Volume type name to use when creating volumes.",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Description of the resource",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error message",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error traceback",
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
				MarkdownDescription: "Is limit based",
			},
			"is_usage_based": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Is usage based",
			},
			"marketplace_category_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the marketplace category",
			},
			"marketplace_category_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the marketplace category",
			},
			"marketplace_offering_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the marketplace offering",
			},
			"marketplace_offering_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the marketplace offering",
			},
			"marketplace_plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the marketplace plan",
			},
			"marketplace_resource_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Marketplace resource state",
			},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the marketplace resource",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the resource",
			},
			"project": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project",
			},
			"project_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the project",
			},
			"project_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the project",
			},
			"quotas": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"limit": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Limit",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the resource",
						},
						"usage": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Usage",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Quotas",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource type",
			},
			"service_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the service",
			},
			"service_settings": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Service settings",
			},
			"service_settings_error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Service settings error message",
			},
			"service_settings_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Service settings state",
			},
			"service_settings_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the service settings",
			},
			"skip_creation_of_default_router": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Skip creation of default router",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "State",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Url",
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

	rawClient, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = NewClient(rawClient)
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
		apiResp, err := d.client.GetOpenstackTenant(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Tenant",
				"An error occurred while reading the Openstack Tenant by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, *apiResp, &data)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_tenant.",
			)
			return
		}

		results, err := d.client.ListOpenstackTenant(ctx, filters)
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

func (d *OpenstackTenantDataSource) mapResponseToModel(ctx context.Context, apiResp OpenstackTenantResponse, model *OpenstackTenantDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
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

	return diags
}
