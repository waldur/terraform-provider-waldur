package network_rbac_policy

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackNetworkRbacPolicyDataSource{}

func NewOpenstackNetworkRbacPolicyDataSource() datasource.DataSource {
	return &OpenstackNetworkRbacPolicyDataSource{}
}

type OpenstackNetworkRbacPolicyDataSource struct {
	client *Client
}

type OpenstackNetworkRbacPolicyDataSourceModel struct {
	OpenstackNetworkRbacPolicyModel
	Filters *OpenstackNetworkRbacPolicyFiltersModel `tfsdk:"filters"`
}

func (d *OpenstackNetworkRbacPolicyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_network_rbac_policy"
}

func (d *OpenstackNetworkRbacPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Network Rbac Policy data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Filter parameters for querying Openstack Network Rbac Policy",
				Attributes: map[string]schema.Attribute{
					"network": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Network URL",
					},
					"network_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Network UUID",
					},
					"policy_type": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Type of access granted - either shared access or external network access Allowed values: `access_as_external`, `access_as_shared`.",
						Validators: []validator.String{
							stringvalidator.OneOf("access_as_external", "access_as_shared"),
						},
					},
					"target_tenant": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Target tenant URL",
					},
					"target_tenant_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Target tenant UUID",
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
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of the backend",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Created",
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
		},
	}
}

func (d *OpenstackNetworkRbacPolicyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OpenstackNetworkRbacPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackNetworkRbacPolicyDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		apiResp, err := d.client.GetOpenstackNetworkRbacPolicy(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Network Rbac Policy",
				"An error occurred while reading the Openstack Network Rbac Policy by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_network_rbac_policy.",
			)
			return
		}

		results, err := d.client.ListOpenstackNetworkRbacPolicy(ctx, filters)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Openstack Network Rbac Policy",
				"An error occurred while filtering Openstack Network Rbac Policy: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Network Rbac Policy Not Found",
				"No Openstack Network Rbac Policy found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Network Rbac Policys Found",
				fmt.Sprintf("Found %d Openstack Network Rbac Policys with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, results[0])...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
