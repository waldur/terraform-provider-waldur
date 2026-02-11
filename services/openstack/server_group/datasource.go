package server_group

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackServerGroupDataSource{}

func NewOpenstackServerGroupDataSource() datasource.DataSource {
	return &OpenstackServerGroupDataSource{}
}

type OpenstackServerGroupDataSource struct {
	client *OpenstackServerGroupClient
}

type OpenstackServerGroupDataSourceModel struct {
	OpenstackServerGroupModel
	Filters *OpenstackServerGroupFiltersModel `tfsdk:"filters"`
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
				MarkdownDescription: "Openstack Server Group UUID",
			},
			"filters": (&OpenstackServerGroupFiltersModel{}).GetSchema(),
			"backend_id": schema.StringAttribute{

				Computed: true,

				MarkdownDescription: "Backend Id",
			},
			"customer": schema.StringAttribute{

				Computed: true,

				MarkdownDescription: "Customer",
			},
			"description": schema.StringAttribute{

				Computed: true,

				MarkdownDescription: "Description",
			},
			"display_name": schema.StringAttribute{

				Computed: true,

				MarkdownDescription: "Display Name",
			},
			"error_message": schema.StringAttribute{

				Computed: true,

				MarkdownDescription: "Error Message",
			},
			"instances": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"backend_id": schema.StringAttribute{

							Computed: true,

							MarkdownDescription: "Instance ID in the OpenStack backend",
						},
						"name": schema.StringAttribute{

							Computed: true,

							MarkdownDescription: "Name",
						},
						"uuid": schema.StringAttribute{

							Computed: true,

							MarkdownDescription: "Uuid",
						},
					},
				},

				Computed: true,

				MarkdownDescription: "Instances",
			},
			"marketplace_resource_uuid": schema.StringAttribute{

				Computed: true,

				MarkdownDescription: "Marketplace Resource Uuid",
			},
			"name": schema.StringAttribute{

				Computed: true,

				MarkdownDescription: "Name",
			},
			"policy": schema.StringAttribute{

				Computed: true,

				MarkdownDescription: "Server group policy determining the rules for scheduling servers in this group",
			},
			"project": schema.StringAttribute{

				Computed: true,

				MarkdownDescription: "Project",
			},
			"resource_type": schema.StringAttribute{

				Computed: true,

				MarkdownDescription: "Resource Type",
			},
			"state": schema.StringAttribute{

				Computed: true,

				MarkdownDescription: "State",
			},
			"tenant": schema.StringAttribute{

				Computed: true,

				MarkdownDescription: "Tenant",
			},
			"tenant_name": schema.StringAttribute{

				Computed: true,

				MarkdownDescription: "Tenant Name",
			},
			"tenant_uuid": schema.StringAttribute{

				Computed: true,

				MarkdownDescription: "Tenant Uuid",
			},
			"url": schema.StringAttribute{

				Computed: true,

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

	d.client = &OpenstackServerGroupClient{}
	if err := d.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			err.Error(),
		)
		return
	}
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
		apiResp, err := d.client.Get(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Server Group",
				"An error occurred while reading the Openstack Server Group by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_server_group.",
			)
			return
		}

		results, err := d.client.List(ctx, filters)
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

		resp.Diagnostics.Append(data.CopyFrom(ctx, results[0])...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
