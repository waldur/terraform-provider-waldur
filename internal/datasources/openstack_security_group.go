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
var _ datasource.DataSource = &OpenstackSecurityGroupDataSource{}

func NewOpenstackSecurityGroupDataSource() datasource.DataSource {
	return &OpenstackSecurityGroupDataSource{}
}

// OpenstackSecurityGroupDataSource defines the data source implementation.
type OpenstackSecurityGroupDataSource struct {
	client *client.Client
}

// OpenstackSecurityGroupApiResponse is the API response model.
type OpenstackSecurityGroupApiResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl      *string                               `json:"access_url" tfsdk:"access_url"`
	BackendId      *string                               `json:"backend_id" tfsdk:"backend_id"`
	Created        *string                               `json:"created" tfsdk:"created"`
	Description    *string                               `json:"description" tfsdk:"description"`
	ErrorMessage   *string                               `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback *string                               `json:"error_traceback" tfsdk:"error_traceback"`
	Modified       *string                               `json:"modified" tfsdk:"modified"`
	Name           *string                               `json:"name" tfsdk:"name"`
	ResourceType   *string                               `json:"resource_type" tfsdk:"resource_type"`
	Rules          []OpenstackSecurityGroupRulesResponse `json:"rules" tfsdk:"rules"`
	State          *string                               `json:"state" tfsdk:"state"`
	Tenant         *string                               `json:"tenant" tfsdk:"tenant"`
	TenantName     *string                               `json:"tenant_name" tfsdk:"tenant_name"`
	TenantUuid     *string                               `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Url            *string                               `json:"url" tfsdk:"url"`
}

type OpenstackSecurityGroupRulesResponse struct {
	Cidr            *string `json:"cidr" tfsdk:"cidr"`
	Description     *string `json:"description" tfsdk:"description"`
	Direction       *string `json:"direction" tfsdk:"direction"`
	Ethertype       *string `json:"ethertype" tfsdk:"ethertype"`
	FromPort        *int64  `json:"from_port" tfsdk:"from_port"`
	Id              *int64  `json:"id" tfsdk:"id"`
	Protocol        *string `json:"protocol" tfsdk:"protocol"`
	RemoteGroup     *string `json:"remote_group" tfsdk:"remote_group"`
	RemoteGroupName *string `json:"remote_group_name" tfsdk:"remote_group_name"`
	RemoteGroupUuid *string `json:"remote_group_uuid" tfsdk:"remote_group_uuid"`
	ToPort          *int64  `json:"to_port" tfsdk:"to_port"`
}

// OpenstackSecurityGroupFiltersModel contains the filter parameters for querying.
type OpenstackSecurityGroupFiltersModel struct {
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
	Query                types.String `tfsdk:"query"`
	ServiceSettingsName  types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid  types.String `tfsdk:"service_settings_uuid"`
	State                types.String `tfsdk:"state"`
	Tenant               types.String `tfsdk:"tenant"`
	TenantUuid           types.String `tfsdk:"tenant_uuid"`
	Uuid                 types.String `tfsdk:"uuid"`
}

// OpenstackSecurityGroupDataSourceModel describes the data source data model.
type OpenstackSecurityGroupDataSourceModel struct {
	UUID           types.String                        `tfsdk:"id"`
	Filters        *OpenstackSecurityGroupFiltersModel `tfsdk:"filters"`
	AccessUrl      types.String                        `tfsdk:"access_url"`
	BackendId      types.String                        `tfsdk:"backend_id"`
	Created        types.String                        `tfsdk:"created"`
	Description    types.String                        `tfsdk:"description"`
	ErrorMessage   types.String                        `tfsdk:"error_message"`
	ErrorTraceback types.String                        `tfsdk:"error_traceback"`
	Modified       types.String                        `tfsdk:"modified"`
	Name           types.String                        `tfsdk:"name"`
	ResourceType   types.String                        `tfsdk:"resource_type"`
	Rules          types.List                          `tfsdk:"rules"`
	State          types.String                        `tfsdk:"state"`
	Tenant         types.String                        `tfsdk:"tenant"`
	TenantName     types.String                        `tfsdk:"tenant_name"`
	TenantUuid     types.String                        `tfsdk:"tenant_uuid"`
	Url            types.String                        `tfsdk:"url"`
}

func (d *OpenstackSecurityGroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_security_group"
}

func (d *OpenstackSecurityGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Security Group data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Filter parameters for querying Openstack Security Group",
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
					"query": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Search by name or description",
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
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error message",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error traceback",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the resource",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource type",
			},
			"rules": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cidr": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "CIDR notation for the source/destination network address range",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Description of the resource",
						},
						"direction": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Traffic direction - either 'ingress' (incoming) or 'egress' (outgoing)",
						},
						"ethertype": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "IP protocol version - either 'IPv4' or 'IPv6'",
						},
						"from_port": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Starting port number in the range (1-65535)",
						},
						"id": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Id",
						},
						"protocol": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "The network protocol (TCP, UDP, ICMP, or empty for any protocol)",
						},
						"remote_group": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Remote security group that this rule references, if any",
						},
						"remote_group_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name of the remote group",
						},
						"remote_group_uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "UUID of the remote group",
						},
						"to_port": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Ending port number in the range (1-65535)",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Rules",
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

func (d *OpenstackSecurityGroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OpenstackSecurityGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackSecurityGroupDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var apiResp OpenstackSecurityGroupApiResponse

		err := d.client.GetByUUID(ctx, "/api/openstack-security-groups/{uuid}/", data.UUID.ValueString(), &apiResp)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Security Group",
				"An error occurred while reading the Openstack Security Group by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, apiResp, &data)...)

	} else {
		// Filter by provided parameters
		var results []OpenstackSecurityGroupApiResponse

		filters := make(map[string]string)
		if data.Filters != nil {
			type filterDef struct {
				name string
				val  attr.Value
			}
			filterDefs := []filterDef{
				{"backend_id", data.Filters.BackendId},
				{"can_manage", data.Filters.CanManage},
				{"customer", data.Filters.Customer},
				{"customer_abbreviation", data.Filters.CustomerAbbreviation},
				{"customer_name", data.Filters.CustomerName},
				{"customer_native_name", data.Filters.CustomerNativeName},
				{"customer_uuid", data.Filters.CustomerUuid},
				{"description", data.Filters.Description},
				{"external_ip", data.Filters.ExternalIp},
				{"name", data.Filters.Name},
				{"name_exact", data.Filters.NameExact},
				{"project", data.Filters.Project},
				{"project_name", data.Filters.ProjectName},
				{"project_uuid", data.Filters.ProjectUuid},
				{"query", data.Filters.Query},
				{"service_settings_name", data.Filters.ServiceSettingsName},
				{"service_settings_uuid", data.Filters.ServiceSettingsUuid},
				{"state", data.Filters.State},
				{"tenant", data.Filters.Tenant},
				{"tenant_uuid", data.Filters.TenantUuid},
				{"uuid", data.Filters.Uuid},
			}

			for _, fd := range filterDefs {
				if fd.val.IsNull() || fd.val.IsUnknown() {
					continue
				}
				switch v := fd.val.(type) {
				case types.String:
					filters[fd.name] = v.ValueString()
				case types.Int64:
					filters[fd.name] = fmt.Sprintf("%d", v.ValueInt64())
				case types.Bool:
					filters[fd.name] = fmt.Sprintf("%t", v.ValueBool())
				case types.Float64:
					filters[fd.name] = fmt.Sprintf("%f", v.ValueFloat64())
				}
			}
		}

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_security_group.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/openstack-security-groups/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Openstack Security Group",
				"An error occurred while filtering Openstack Security Group: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Security Group Not Found",
				"No Openstack Security Group found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Security Groups Found",
				fmt.Sprintf("Found %d Openstack Security Groups with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *OpenstackSecurityGroupDataSource) mapResponseToModel(ctx context.Context, apiResp OpenstackSecurityGroupApiResponse, model *OpenstackSecurityGroupDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	listValRules, listDiagsRules := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
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
	}}, apiResp.Rules)
	diags.Append(listDiagsRules...)
	model.Rules = listValRules
	model.State = types.StringPointerValue(apiResp.State)
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantName = types.StringPointerValue(apiResp.TenantName)
	model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
