package ssh_public_key

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &CoreSshPublicKeyDataSource{}

func NewCoreSshPublicKeyDataSource() datasource.DataSource {
	return &CoreSshPublicKeyDataSource{}
}

type CoreSshPublicKeyDataSource struct {
	client *Client
}

type CoreSshPublicKeyDataSourceModel struct {
	CoreSshPublicKeyModel
	Filters *CoreSshPublicKeyFiltersModel `tfsdk:"filters"`
}

func (d *CoreSshPublicKeyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_core_ssh_public_key"
}

func (d *CoreSshPublicKeyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Core Ssh Public Key data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Filter parameters for querying Core Ssh Public Key",
				Attributes: map[string]schema.Attribute{
					"created": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Created after",
					},
					"fingerprint_md5": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Fingerprint md5",
					},
					"fingerprint_sha256": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Fingerprint sha256",
					},
					"fingerprint_sha512": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Fingerprint sha512",
					},
					"is_shared": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Is shared",
					},
					"modified": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Modified after",
					},
					"name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name",
					},
					"name_exact": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name (exact)",
					},
					"user_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "User UUID",
					},
					"uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "UUID",
					},
				},
			},
			"fingerprint_md5": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Fingerprint md5",
			},
			"fingerprint_sha256": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Fingerprint sha256",
			},
			"fingerprint_sha512": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Fingerprint sha512",
			},
			"is_shared": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Is shared",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name of the resource",
			},
			"public_key": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Public key",
			},
			"type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Type",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Url",
			},
			"user_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the user",
			},
		},
	}
}

func (d *CoreSshPublicKeyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *CoreSshPublicKeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CoreSshPublicKeyDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		apiResp, err := d.client.GetCoreSshPublicKey(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Core Ssh Public Key",
				"An error occurred while reading the Core Ssh Public Key by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup core_ssh_public_key.",
			)
			return
		}

		results, err := d.client.ListCoreSshPublicKey(ctx, filters)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Core Ssh Public Key",
				"An error occurred while filtering Core Ssh Public Key: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Core Ssh Public Key Not Found",
				"No Core Ssh Public Key found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Core Ssh Public Keys Found",
				fmt.Sprintf("Found %d Core Ssh Public Keys with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, results[0])...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
