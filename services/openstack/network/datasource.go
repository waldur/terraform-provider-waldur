package network

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
var _ datasource.DataSource = &OpenstackNetworkDataSource{}

func NewOpenstackNetworkDataSource() datasource.DataSource {
	return &OpenstackNetworkDataSource{}
}

// OpenstackNetworkDataSource defines the data source implementation.
type OpenstackNetworkDataSource struct {
	client *Client
}

// OpenstackNetworkFiltersModel contains the filter parameters for querying.
type OpenstackNetworkFiltersModel struct {
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
}

// OpenstackNetworkDataSourceModel describes the data source data model.
type OpenstackNetworkDataSourceModel struct {
	UUID           types.String                  `tfsdk:"id"`
	Filters        *OpenstackNetworkFiltersModel `tfsdk:"filters"`
	AccessUrl      types.String                  `tfsdk:"access_url"`
	BackendId      types.String                  `tfsdk:"backend_id"`
	Created        types.String                  `tfsdk:"created"`
	Description    types.String                  `tfsdk:"description"`
	ErrorMessage   types.String                  `tfsdk:"error_message"`
	ErrorTraceback types.String                  `tfsdk:"error_traceback"`
	IsExternal     types.Bool                    `tfsdk:"is_external"`
	Modified       types.String                  `tfsdk:"modified"`
	Mtu            types.Int64                   `tfsdk:"mtu"`
	Name           types.String                  `tfsdk:"name"`
	RbacPolicies   types.List                    `tfsdk:"rbac_policies"`
	ResourceType   types.String                  `tfsdk:"resource_type"`
	SegmentationId types.Int64                   `tfsdk:"segmentation_id"`
	State          types.String                  `tfsdk:"state"`
	Subnets        types.List                    `tfsdk:"subnets"`
	Tenant         types.String                  `tfsdk:"tenant"`
	TenantName     types.String                  `tfsdk:"tenant_name"`
	TenantUuid     types.String                  `tfsdk:"tenant_uuid"`
	Type           types.String                  `tfsdk:"type"`
	Url            types.String                  `tfsdk:"url"`
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
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Filter parameters for querying Openstack Network",
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
						MarkdownDescription: "Is external",
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
						MarkdownDescription: "State Allowed values: `CREATING`, `CREATION_SCHEDULED`, `DELETING`, `DELETION_SCHEDULED`, `ERRED`, `OK`, `UPDATE_SCHEDULED`, `UPDATING`.",
						Validators: []validator.String{
							stringvalidator.OneOf("CREATING", "CREATION_SCHEDULED", "DELETING", "DELETION_SCHEDULED", "ERRED", "OK", "UPDATE_SCHEDULED", "UPDATING"),
						},
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
						MarkdownDescription: "Type",
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
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of the backend",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Created",
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
			"is_external": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Defines whether this network is external (public) or internal (private)",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Modified",
			},
			"mtu": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The maximum transmission unit (MTU) value to address fragmentation.",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the resource",
			},
			"rbac_policies": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"backend_id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "ID of the backend",
						},
						"created": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Created",
						},
						"network": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Network",
						},
						"network_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name of the network",
						},
						"policy_type": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Type of access granted - either shared access or external network access",
						},
						"target_tenant": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Target tenant",
						},
						"target_tenant_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name of the target tenant",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Url",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Rbac policies",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource type",
			},
			"segmentation_id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "VLAN ID for VLAN networks or tunnel ID for VXLAN/GRE networks",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "State",
			},
			"subnets": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allocation_pools": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"end": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "An IPv4 or IPv6 address.",
									},
									"start": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "An IPv4 or IPv6 address.",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: "Allocation pools",
						},
						"cidr": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "IPv4 network address in CIDR format (e.g. 192.168.0.0/24)",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Description of the resource",
						},
						"enable_dhcp": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "If True, DHCP service will be enabled on this subnet",
						},
						"gateway_ip": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "IP address of the gateway for this subnet",
						},
						"ip_version": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "IP protocol version (4 or 6)",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the resource",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Subnets",
			},
			"tenant": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "OpenStack tenant this network belongs to",
			},
			"tenant_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the tenant",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the tenant",
			},
			"type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Network type, such as local, flat, vlan, vxlan, or gre",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Url",
			},
		},
	}
}

func (d *OpenstackNetworkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OpenstackNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackNetworkDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		apiResp, err := d.client.GetOpenstackNetwork(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Network",
				"An error occurred while reading the Openstack Network by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, *apiResp, &data)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_network.",
			)
			return
		}

		results, err := d.client.ListOpenstackNetwork(ctx, filters)
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

func (d *OpenstackNetworkDataSource) mapResponseToModel(ctx context.Context, apiResp OpenstackNetworkResponse, model *OpenstackNetworkDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
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
