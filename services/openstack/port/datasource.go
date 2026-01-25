package port

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackPortDataSource{}

func NewOpenstackPortDataSource() datasource.DataSource {
	return &OpenstackPortDataSource{}
}

type OpenstackPortDataSource struct {
	client *Client
}

type OpenstackPortDataSourceModel struct {
	OpenstackPortModel
	Filters *OpenstackPortFiltersModel `tfsdk:"filters"`
}

func (d *OpenstackPortDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_port"
}

func (d *OpenstackPortDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Port data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Filter parameters for querying Openstack Port",
				Attributes: map[string]schema.Attribute{
					"admin_state_up": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Admin state up",
					},
					"backend_id": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "ID of the backend",
					},
					"device_id": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "ID of the device",
					},
					"device_owner": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Device owner",
					},
					"exclude_subnet_uuids": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Exclude Subnet UUIDs (comma-separated)",
					},
					"fixed_ips": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Search by fixed IP",
					},
					"has_device_owner": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Has device owner",
					},
					"mac_address": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Mac address",
					},
					"name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name",
					},
					"name_exact": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name (exact)",
					},
					"network_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Search by network name",
					},
					"network_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Search by network UUID",
					},
					"query": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Search by name, MAC address or backend ID",
					},
					"status": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Status",
					},
					"tenant": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Tenant URL",
					},
					"tenant_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Tenant UUID",
					},
				},
			},
			"access_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Access url",
			},
			"admin_state_up": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Administrative state of the port. If down, port does not forward packets",
			},
			"allowed_address_pairs": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"mac_address": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Mac address",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Allowed address pairs",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Port ID in OpenStack",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Created",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Description of the resource",
			},
			"device_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of device (instance, router etc) to which this port is connected",
			},
			"device_owner": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Entity that uses this port (e.g. network:router_interface)",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error message",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error traceback",
			},
			"fixed_ips": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "IP address to assign to the port",
						},
						"subnet_id": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "ID of the subnet in which to assign the IP address",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Fixed ips",
			},
			"floating_ips": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "Floating ips",
			},
			"mac_address": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "MAC address of the port",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the resource",
			},
			"network": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Network to which this port belongs",
			},
			"network_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the network",
			},
			"network_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the network",
			},
			"port_security_enabled": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "If True, security groups and rules will be applied to this port",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource type",
			},
			"security_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the resource",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Url",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Security groups",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "State",
			},
			"status": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Port status in OpenStack (e.g. ACTIVE, DOWN)",
			},
			"tenant": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "OpenStack tenant this port belongs to",
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

func (d *OpenstackPortDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OpenstackPortDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackPortDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		apiResp, err := d.client.GetOpenstackPort(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Port",
				"An error occurred while reading the Openstack Port by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_port.",
			)
			return
		}

		results, err := d.client.ListOpenstackPort(ctx, filters)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Openstack Port",
				"An error occurred while filtering Openstack Port: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Port Not Found",
				"No Openstack Port found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Ports Found",
				fmt.Sprintf("Found %d Openstack Ports with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, results[0])...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
