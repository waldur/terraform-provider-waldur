package security_group

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackSecurityGroupDataSource{}

func NewOpenstackSecurityGroupDataSource() datasource.DataSource {
	return &OpenstackSecurityGroupDataSource{}
}

type OpenstackSecurityGroupDataSource struct {
	client *Client
}

type OpenstackSecurityGroupDataSourceModel struct {
	OpenstackSecurityGroupModel
	Filters *OpenstackSecurityGroupFiltersModel `tfsdk:"filters"`
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
				CustomType:          timetypes.RFC3339Type{},
				Computed:            true,
				MarkdownDescription: "Created",
			},
			"description": schema.StringAttribute{
				Optional:            true,
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
				CustomType:          timetypes.RFC3339Type{},
				Computed:            true,
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Optional:            true,
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
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(65535),
							},
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
							Validators: []validator.Int64{
								int64validator.AtLeast(-2147483648),
								int64validator.AtMost(65535),
							},
						},
					},
				},
				Optional:            true,
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

func (d *OpenstackSecurityGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackSecurityGroupDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		apiResp, err := d.client.GetOpenstackSecurityGroup(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Security Group",
				"An error occurred while reading the Openstack Security Group by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_security_group.",
			)
			return
		}

		results, err := d.client.ListOpenstackSecurityGroup(ctx, filters)
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

		resp.Diagnostics.Append(data.CopyFrom(ctx, results[0])...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
