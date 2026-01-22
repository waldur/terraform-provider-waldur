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
var _ datasource.DataSource = &CoreSshPublicKeyDataSource{}

func NewCoreSshPublicKeyDataSource() datasource.DataSource {
	return &CoreSshPublicKeyDataSource{}
}

// CoreSshPublicKeyDataSource defines the data source implementation.
type CoreSshPublicKeyDataSource struct {
	client *client.Client
}

// CoreSshPublicKeyApiResponse is the API response model.
type CoreSshPublicKeyApiResponse struct {
	UUID *string `json:"uuid"`

	FingerprintMd5    *string `json:"fingerprint_md5" tfsdk:"fingerprint_md5"`
	FingerprintSha256 *string `json:"fingerprint_sha256" tfsdk:"fingerprint_sha256"`
	FingerprintSha512 *string `json:"fingerprint_sha512" tfsdk:"fingerprint_sha512"`
	IsShared          *bool   `json:"is_shared" tfsdk:"is_shared"`
	Name              *string `json:"name" tfsdk:"name"`
	PublicKey         *string `json:"public_key" tfsdk:"public_key"`
	Type              *string `json:"type" tfsdk:"type"`
	Url               *string `json:"url" tfsdk:"url"`
	UserUuid          *string `json:"user_uuid" tfsdk:"user_uuid"`
}

// CoreSshPublicKeyFiltersModel contains the filter parameters for querying.
type CoreSshPublicKeyFiltersModel struct {
	Created           types.String `tfsdk:"created"`
	FingerprintMd5    types.String `tfsdk:"fingerprint_md5"`
	FingerprintSha256 types.String `tfsdk:"fingerprint_sha256"`
	FingerprintSha512 types.String `tfsdk:"fingerprint_sha512"`
	IsShared          types.Bool   `tfsdk:"is_shared"`
	Modified          types.String `tfsdk:"modified"`
	Name              types.String `tfsdk:"name"`
	NameExact         types.String `tfsdk:"name_exact"`
	UserUuid          types.String `tfsdk:"user_uuid"`
	Uuid              types.String `tfsdk:"uuid"`
}

// CoreSshPublicKeyDataSourceModel describes the data source data model.
type CoreSshPublicKeyDataSourceModel struct {
	UUID              types.String                  `tfsdk:"id"`
	Filters           *CoreSshPublicKeyFiltersModel `tfsdk:"filters"`
	FingerprintMd5    types.String                  `tfsdk:"fingerprint_md5"`
	FingerprintSha256 types.String                  `tfsdk:"fingerprint_sha256"`
	FingerprintSha512 types.String                  `tfsdk:"fingerprint_sha512"`
	IsShared          types.Bool                    `tfsdk:"is_shared"`
	Name              types.String                  `tfsdk:"name"`
	PublicKey         types.String                  `tfsdk:"public_key"`
	Type              types.String                  `tfsdk:"type"`
	Url               types.String                  `tfsdk:"url"`
	UserUuid          types.String                  `tfsdk:"user_uuid"`
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
				Computed:            true,
				MarkdownDescription: "Name of the resource",
			},
			"public_key": schema.StringAttribute{
				Computed:            true,
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

func (d *CoreSshPublicKeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CoreSshPublicKeyDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var apiResp CoreSshPublicKeyApiResponse

		err := d.client.GetByUUID(ctx, "/api/keys/{uuid}/", data.UUID.ValueString(), &apiResp)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Core Ssh Public Key",
				"An error occurred while reading the Core Ssh Public Key by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, apiResp, &data)...)

	} else {
		// Filter by provided parameters
		var results []CoreSshPublicKeyApiResponse

		filters := make(map[string]string)
		if data.Filters != nil {
			type filterDef struct {
				name string
				val  attr.Value
			}
			filterDefs := []filterDef{
				{"created", data.Filters.Created},
				{"fingerprint_md5", data.Filters.FingerprintMd5},
				{"fingerprint_sha256", data.Filters.FingerprintSha256},
				{"fingerprint_sha512", data.Filters.FingerprintSha512},
				{"is_shared", data.Filters.IsShared},
				{"modified", data.Filters.Modified},
				{"name", data.Filters.Name},
				{"name_exact", data.Filters.NameExact},
				{"user_uuid", data.Filters.UserUuid},
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
				"At least one filter parameter (or 'id') must be provided to lookup core_ssh_public_key.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/keys/", filters, &results)
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

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *CoreSshPublicKeyDataSource) mapResponseToModel(ctx context.Context, apiResp CoreSshPublicKeyApiResponse, model *CoreSshPublicKeyDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.FingerprintMd5 = types.StringPointerValue(apiResp.FingerprintMd5)
	model.FingerprintSha256 = types.StringPointerValue(apiResp.FingerprintSha256)
	model.FingerprintSha512 = types.StringPointerValue(apiResp.FingerprintSha512)
	model.IsShared = types.BoolPointerValue(apiResp.IsShared)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.PublicKey = types.StringPointerValue(apiResp.PublicKey)
	model.Type = types.StringPointerValue(apiResp.Type)
	model.Url = types.StringPointerValue(apiResp.Url)
	model.UserUuid = types.StringPointerValue(apiResp.UserUuid)

	return diags
}
