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
var _ datasource.DataSource = &OpenstackNetworkDataSource{}

func NewOpenstackNetworkDataSource() datasource.DataSource {
	return &OpenstackNetworkDataSource{}
}

// OpenstackNetworkDataSource defines the data source implementation.
type OpenstackNetworkDataSource struct {
	client *client.Client
}

// OpenstackNetworkApiResponse is the API response model.
type OpenstackNetworkApiResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl                   *string                                `json:"access_url" tfsdk:"access_url"`
	BackendId                   *string                                `json:"backend_id" tfsdk:"backend_id"`
	Created                     *string                                `json:"created" tfsdk:"created"`
	Customer                    *string                                `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation        *string                                `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName                *string                                `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName          *string                                `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerUuid                *string                                `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description                 *string                                `json:"description" tfsdk:"description"`
	ErrorMessage                *string                                `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback              *string                                `json:"error_traceback" tfsdk:"error_traceback"`
	IsExternal                  *bool                                  `json:"is_external" tfsdk:"is_external"`
	IsLimitBased                *bool                                  `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                                  `json:"is_usage_based" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string                                `json:"marketplace_category_name" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                                `json:"marketplace_category_uuid" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                                `json:"marketplace_offering_name" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                                `json:"marketplace_offering_uuid" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                                `json:"marketplace_plan_uuid" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                                `json:"marketplace_resource_state" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                                `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                                `json:"modified" tfsdk:"modified"`
	Mtu                         *int64                                 `json:"mtu" tfsdk:"mtu"`
	Name                        *string                                `json:"name" tfsdk:"name"`
	Project                     *string                                `json:"project" tfsdk:"project"`
	ProjectName                 *string                                `json:"project_name" tfsdk:"project_name"`
	ProjectUuid                 *string                                `json:"project_uuid" tfsdk:"project_uuid"`
	RbacPolicies                []OpenstackNetworkRbacPoliciesResponse `json:"rbac_policies" tfsdk:"rbac_policies"`
	ResourceType                *string                                `json:"resource_type" tfsdk:"resource_type"`
	SegmentationId              *int64                                 `json:"segmentation_id" tfsdk:"segmentation_id"`
	ServiceName                 *string                                `json:"service_name" tfsdk:"service_name"`
	ServiceSettings             *string                                `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                                `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                                `json:"service_settings_state" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                                `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	State                       *string                                `json:"state" tfsdk:"state"`
	Subnets                     []OpenstackNetworkSubnetsResponse      `json:"subnets" tfsdk:"subnets"`
	Tenant                      *string                                `json:"tenant" tfsdk:"tenant"`
	TenantName                  *string                                `json:"tenant_name" tfsdk:"tenant_name"`
	TenantUuid                  *string                                `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Type                        *string                                `json:"type" tfsdk:"type"`
	Url                         *string                                `json:"url" tfsdk:"url"`
}

type OpenstackNetworkRbacPoliciesResponse struct {
	BackendId        *string `json:"backend_id" tfsdk:"backend_id"`
	Created          *string `json:"created" tfsdk:"created"`
	Network          *string `json:"network" tfsdk:"network"`
	NetworkName      *string `json:"network_name" tfsdk:"network_name"`
	PolicyType       *string `json:"policy_type" tfsdk:"policy_type"`
	TargetTenant     *string `json:"target_tenant" tfsdk:"target_tenant"`
	TargetTenantName *string `json:"target_tenant_name" tfsdk:"target_tenant_name"`
	Url              *string `json:"url" tfsdk:"url"`
}

type OpenstackNetworkSubnetsResponse struct {
	AllocationPools []OpenstackNetworkSubnetsAllocationPoolsResponse `json:"allocation_pools" tfsdk:"allocation_pools"`
	Cidr            *string                                          `json:"cidr" tfsdk:"cidr"`
	Description     *string                                          `json:"description" tfsdk:"description"`
	EnableDhcp      *bool                                            `json:"enable_dhcp" tfsdk:"enable_dhcp"`
	GatewayIp       *string                                          `json:"gateway_ip" tfsdk:"gateway_ip"`
	IpVersion       *int64                                           `json:"ip_version" tfsdk:"ip_version"`
	Name            *string                                          `json:"name" tfsdk:"name"`
}

type OpenstackNetworkSubnetsAllocationPoolsResponse struct {
	End   *string `json:"end" tfsdk:"end"`
	Start *string `json:"start" tfsdk:"start"`
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
				Computed:            true,
				MarkdownDescription: "Backend ID",
			},
			"can_manage": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Can manage",
			},
			"customer": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer UUID",
			},
			"customer_abbreviation": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer abbreviation",
			},
			"customer_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer name",
			},
			"customer_native_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer native name",
			},
			"customer_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer UUID",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Description",
			},
			"direct_only": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Direct only",
			},
			"external_ip": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "External IP",
			},
			"is_external": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Name",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Name (exact)",
			},
			"project": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project UUID",
			},
			"project_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project name",
			},
			"project_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project UUID",
			},
			"rbac_only": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "RBAC only",
			},
			"service_settings_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Service settings name",
			},
			"service_settings_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Service settings UUID",
			},
			"state": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "State",
			},
			"tenant": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Tenant URL",
			},
			"tenant_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Tenant UUID",
			},
			"type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
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
		var apiResp OpenstackNetworkApiResponse

		err := d.client.GetByUUID(ctx, "/api/openstack-networks/{uuid}/", data.UUID.ValueString(), &apiResp)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Network",
				"An error occurred while reading the Openstack Network by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, apiResp, &data)...)

	} else {
		// Filter by provided parameters
		var results []OpenstackNetworkApiResponse

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

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *OpenstackNetworkDataSource) mapResponseToModel(ctx context.Context, apiResp OpenstackNetworkApiResponse, model *OpenstackNetworkDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Customer = types.StringPointerValue(apiResp.Customer)
	model.CustomerAbbreviation = types.StringPointerValue(apiResp.CustomerAbbreviation)
	model.CustomerName = types.StringPointerValue(apiResp.CustomerName)
	model.CustomerNativeName = types.StringPointerValue(apiResp.CustomerNativeName)
	model.CustomerUuid = types.StringPointerValue(apiResp.CustomerUuid)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.IsExternal = types.BoolPointerValue(apiResp.IsExternal)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.Mtu = types.Int64PointerValue(apiResp.Mtu)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.Project = types.StringPointerValue(apiResp.Project)
	model.ProjectName = types.StringPointerValue(apiResp.ProjectName)
	model.ProjectUuid = types.StringPointerValue(apiResp.ProjectUuid)
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
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.SegmentationId = types.Int64PointerValue(apiResp.SegmentationId)
	model.State = types.StringPointerValue(apiResp.State)
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
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantName = types.StringPointerValue(apiResp.TenantName)
	model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
	model.Type = types.StringPointerValue(apiResp.Type)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
