package floating_ip

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackFloatingIpDataSource{}

func NewOpenstackFloatingIpDataSource() datasource.DataSource {
	return &OpenstackFloatingIpDataSource{}
}

type OpenstackFloatingIpDataSource struct {
	client *OpenstackFloatingIpClient
}

type OpenstackFloatingIpDataSourceModel struct {
	OpenstackFloatingIpModel
	Filters *OpenstackFloatingIpFiltersModel `tfsdk:"filters"`
}

func (d *OpenstackFloatingIpDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_floating_ip"
}

func (d *OpenstackFloatingIpDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Floating Ip data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Openstack Floating Ip UUID",
			},
			"filters": (&OpenstackFloatingIpFiltersModel{}).GetSchema(),
			"address": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The public IPv4 address of the floating IP",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Backend Id",
			},
			"backend_network_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of network in OpenStack where this floating IP is allocated",
			},
			"customer": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Customer",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Description",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error Message",
			},
			"external_address": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Optional address that maps to floating IP's address in external networks",
			},
			"instance_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Instance Name",
			},
			"instance_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Instance Url",
			},
			"instance_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Instance Uuid",
			},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Marketplace Resource Uuid",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name",
			},
			"port": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Port",
			},
			"port_fixed_ips": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "IP address to assign to the port",
						},
						"subnet_id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "ID of the subnet in which to assign the IP address",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Port Fixed Ips",
			},
			"project": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource Type",
			},
			"runtime_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Runtime State",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "State",
			},
			"tenant": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "OpenStack tenant this floating IP belongs to",
			},
			"tenant_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Tenant Name",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Tenant Uuid",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Url",
			},
		},
	}
}

func (d *OpenstackFloatingIpDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	d.client = &OpenstackFloatingIpClient{}
	if err := d.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			err.Error(),
		)
		return
	}
}

func (d *OpenstackFloatingIpDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackFloatingIpDataSourceModel

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
				"Unable to Read Openstack Floating Ip",
				"An error occurred while reading the Openstack Floating Ip by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_floating_ip.",
			)
			return
		}

		results, err := d.client.List(ctx, filters)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Openstack Floating Ip",
				"An error occurred while filtering Openstack Floating Ip: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Floating Ip Not Found",
				"No Openstack Floating Ip found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Floating Ips Found",
				fmt.Sprintf("Found %d Openstack Floating Ips with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, results[0])...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
