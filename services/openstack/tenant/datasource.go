package tenant

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackTenantDataSource{}

func NewOpenstackTenantDataSource() datasource.DataSource {
	return &OpenstackTenantDataSource{}
}

type OpenstackTenantDataSource struct {
	client *OpenstackTenantClient
}

type OpenstackTenantDataSourceModel struct {
	OpenstackTenantModel
	Filters *OpenstackTenantFiltersModel `tfsdk:"filters"`
}

func (d *OpenstackTenantDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_tenant"
}

func (d *OpenstackTenantDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Tenant data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Openstack Tenant UUID",
			},
			"filters": (&OpenstackTenantFiltersModel{}).GetSchema(),
			"availability_zone": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Optional availability group. Will be used for all instances provisioned in this tenant",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of tenant in the OpenStack backend",
			},
			"customer": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Customer",
			},
			"default_volume_type_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Volume type name to use when creating volumes.",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Description",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error Message",
			},
			"external_network_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of external network connected to OpenStack tenant",
			},
			"internal_network_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of internal network in OpenStack tenant",
			},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Marketplace Resource Uuid",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name",
			},
			"project": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project URL",
			},
			"quotas": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"limit": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Limit",
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name",
						},
						"usage": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Usage",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "Quotas",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource Type",
			},
			"skip_creation_of_default_router": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Skip Creation Of Default Router",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "State",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Url",
			},
			"user_password": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Password of the tenant user",
				Sensitive:           true,
			},
			"user_username": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Username of the tenant user",
			},
		},
	}
}

func (d *OpenstackTenantDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	d.client = &OpenstackTenantClient{}
	if err := d.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			err.Error(),
		)
		return
	}
}

func (d *OpenstackTenantDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackTenantDataSourceModel

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
				"Unable to Read Openstack Tenant",
				"An error occurred while reading the Openstack Tenant by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_tenant.",
			)
			return
		}

		results, err := d.client.List(ctx, filters)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Openstack Tenant",
				"An error occurred while filtering Openstack Tenant: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Tenant Not Found",
				"No Openstack Tenant found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Tenants Found",
				fmt.Sprintf("Found %d Openstack Tenants with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, results[0])...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
