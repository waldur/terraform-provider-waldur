package server_group

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
var _ datasource.DataSource = &OpenstackServerGroupDataSource{}

func NewOpenstackServerGroupDataSource() datasource.DataSource {
	return &OpenstackServerGroupDataSource{}
}

// OpenstackServerGroupDataSource defines the data source implementation.
type OpenstackServerGroupDataSource struct {
	client *Client
}

// OpenstackServerGroupFiltersModel contains the filter parameters for querying.
type OpenstackServerGroupFiltersModel struct {
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
	Tenant               types.String `tfsdk:"tenant"`
	TenantUuid           types.String `tfsdk:"tenant_uuid"`
	Uuid                 types.String `tfsdk:"uuid"`
}

type OpenstackServerGroupDataSourceModel struct {
	UUID           types.String                      `tfsdk:"id"`
	Filters        *OpenstackServerGroupFiltersModel `tfsdk:"filters"`
	AccessUrl      types.String                      `tfsdk:"access_url"`
	BackendId      types.String                      `tfsdk:"backend_id"`
	Created        types.String                      `tfsdk:"created"`
	Description    types.String                      `tfsdk:"description"`
	DisplayName    types.String                      `tfsdk:"display_name"`
	ErrorMessage   types.String                      `tfsdk:"error_message"`
	ErrorTraceback types.String                      `tfsdk:"error_traceback"`
	Instances      types.List                        `tfsdk:"instances"`
	Modified       types.String                      `tfsdk:"modified"`
	Name           types.String                      `tfsdk:"name"`
	Policy         types.String                      `tfsdk:"policy"`
	ResourceType   types.String                      `tfsdk:"resource_type"`
	State          types.String                      `tfsdk:"state"`
	Tenant         types.String                      `tfsdk:"tenant"`
	TenantName     types.String                      `tfsdk:"tenant_name"`
	TenantUuid     types.String                      `tfsdk:"tenant_uuid"`
	Url            types.String                      `tfsdk:"url"`
}

func (d *OpenstackServerGroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_server_group"
}

func (d *OpenstackServerGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Server Group data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Filter parameters for querying Openstack Server Group",
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
					"tenant": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Tenant URL",
					},
					"tenant_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Tenant UUID",
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
			"display_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the display",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error message",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error traceback",
			},
			"instances": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"backend_id": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Instance ID in the OpenStack backend",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the resource",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Instances",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the resource",
			},
			"policy": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Server group policy determining the rules for scheduling servers in this group",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource type",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "State",
			},
			"tenant": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Tenant",
			},
			"tenant_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the tenant",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the tenant",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Url",
			},
		},
	}
}

func (d *OpenstackServerGroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OpenstackServerGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackServerGroupDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		apiResp, err := d.client.GetOpenstackServerGroup(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Server Group",
				"An error occurred while reading the Openstack Server Group by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, *apiResp, &data)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_server_group.",
			)
			return
		}

		results, err := d.client.ListOpenstackServerGroup(ctx, filters)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Openstack Server Group",
				"An error occurred while filtering Openstack Server Group: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Server Group Not Found",
				"No Openstack Server Group found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Server Groups Found",
				fmt.Sprintf("Found %d Openstack Server Groups with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *OpenstackServerGroupDataSource) mapResponseToModel(ctx context.Context, apiResp OpenstackServerGroupResponse, model *OpenstackServerGroupDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.DisplayName = types.StringPointerValue(apiResp.DisplayName)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)

	{
		listValInstances, listDiagsInstances := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"backend_id": types.StringType,
			"name":       types.StringType,
		}}, apiResp.Instances)
		diags.Append(listDiagsInstances...)
		model.Instances = listValInstances
	}
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.Policy = types.StringPointerValue(apiResp.Policy)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.State = types.StringPointerValue(apiResp.State)
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantName = types.StringPointerValue(apiResp.TenantName)
	model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
