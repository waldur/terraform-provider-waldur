package subnet

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackSubnetDataSource{}

func NewOpenstackSubnetDataSource() datasource.DataSource {
	return &OpenstackSubnetDataSource{}
}

type OpenstackSubnetDataSource struct {
	client *Client
}

type OpenstackSubnetDataSourceModel struct {
	OpenstackSubnetModel
	Filters *OpenstackSubnetFiltersModel `tfsdk:"filters"`
}

func (d *OpenstackSubnetDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_subnet"
}

func (d *OpenstackSubnetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Subnet data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Filter parameters for querying Openstack Subnet",
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
					"enable_dhcp": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Enable dhcp",
					},
					"external_ip": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "External IP",
					},
					"ip_version": schema.Int64Attribute{
						Optional:            true,
						MarkdownDescription: "Ip version",
					},
					"name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name",
					},
					"name_exact": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name (exact)",
					},
					"network": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Network URL",
					},
					"network_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Network UUID",
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
				Optional:            true,
				MarkdownDescription: "Allocation pools",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of the backend",
			},
			"cidr": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Cidr",
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
			"disable_gateway": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "If True, no gateway IP address will be allocated",
			},
			"dns_nameservers": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Dns nameservers",
			},
			"enable_dhcp": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "If True, DHCP service will be enabled on this subnet",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error message",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error traceback",
			},
			"gateway_ip": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "IP address of the gateway for this subnet",
			},
			"host_routes": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"destination": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Destination",
						},
						"nexthop": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "An IPv4 or IPv6 address.",
						},
					},
				},
				Optional:            true,
				MarkdownDescription: "Host routes",
			},
			"ip_version": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "IP protocol version (4 or 6)",
			},
			"is_connected": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Is subnet connected to the default tenant router.",
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
			"network": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Network to which this subnet belongs",
			},
			"network_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the network",
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
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Url",
			},
		},
	}
}

func (d *OpenstackSubnetDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OpenstackSubnetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackSubnetDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		apiResp, err := d.client.GetOpenstackSubnet(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Subnet",
				"An error occurred while reading the Openstack Subnet by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_subnet.",
			)
			return
		}

		results, err := d.client.ListOpenstackSubnet(ctx, filters)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Openstack Subnet",
				"An error occurred while filtering Openstack Subnet: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Subnet Not Found",
				"No Openstack Subnet found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Subnets Found",
				fmt.Sprintf("Found %d Openstack Subnets with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, results[0])...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
