package network

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackNetworkDataSource{}

func NewOpenstackNetworkDataSource() datasource.DataSource {
	return &OpenstackNetworkDataSource{}
}

type OpenstackNetworkDataSource struct {
	client *OpenstackNetworkClient
}

type OpenstackNetworkDataSourceModel struct {
	OpenstackNetworkModel
	Filters *OpenstackNetworkFiltersModel `tfsdk:"filters"`
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
				MarkdownDescription: "Openstack Network UUID",
			},
			"filters": (&OpenstackNetworkFiltersModel{}).GetSchema(),
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of the backend",
			},
			"customer": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Customer",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Description of the Openstack Network",
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
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the marketplace resource",
			},
			"mtu": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The maximum transmission unit (MTU) value to address fragmentation.",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the Openstack Network",
			},
			"project": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project",
			},
			"rbac_policies": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"backend_id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "ID of the backend",
						},
						"network": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Network",
						},
						"network_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name of the network",
						},
						"policy_type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Type of access granted - either shared access or external network access",
						},
						"target_tenant": schema.StringAttribute{
							Computed:            true,
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
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "UUID of the Openstack Network",
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
										Computed:            true,
										MarkdownDescription: "An IPv4 or IPv6 address.",
									},
									"start": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "An IPv4 or IPv6 address.",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: "Allocation pools",
						},
						"cidr": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "IPv4 network address in CIDR format (e.g. 192.168.0.0/24)",
						},
						"description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Description of the Openstack Network",
						},
						"enable_dhcp": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "If True, DHCP service will be enabled on this subnet",
						},
						"gateway_ip": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "IP address of the gateway for this subnet",
						},
						"ip_version": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "IP protocol version (4 or 6)",
							Validators: []validator.Int64{
								int64validator.AtLeast(-32768),
								int64validator.AtMost(32767),
							},
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name of the Openstack Network",
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "UUID of the Openstack Network",
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

	d.client = &OpenstackNetworkClient{}
	if err := d.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			err.Error(),
		)
		return
	}
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
		apiResp, err := d.client.Get(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Network",
				"An error occurred while reading the Openstack Network by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_network.",
			)
			return
		}

		results, err := d.client.List(ctx, filters)
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

		resp.Diagnostics.Append(data.CopyFrom(ctx, results[0])...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
