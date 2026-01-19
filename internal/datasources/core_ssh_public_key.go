package datasources

import (
	"context"
	"fmt"

	_ "github.com/hashicorp/terraform-plugin-framework/attr" // Used for object types
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-waldur-provider/internal/client"
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

// CoreSshPublicKeyDataSourceModel describes the data source data model.
type CoreSshPublicKeyDataSourceModel struct {
	UUID              types.String `tfsdk:"id"`
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
	PublicKey         types.String `tfsdk:"public_key"`
	Type              types.String `tfsdk:"type"`
	Url               types.String `tfsdk:"url"`
}

func (d *CoreSshPublicKeyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_core_ssh_public_key"
}

func (d *CoreSshPublicKeyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CoreSshPublicKey data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"created": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Created after",
			},
			"fingerprint_md5": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"fingerprint_sha256": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"fingerprint_sha512": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"is_shared": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"modified": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Modified after",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"user_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"public_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
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
		var item map[string]interface{}

		err := d.client.GetByUUID(ctx, "/api/keys/{uuid}/", data.UUID.ValueString(), &item)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read SshPublicKey",
				"An error occurred while reading the core_ssh_public_key by UUID: "+err.Error(),
			)
			return
		}

		// Extract data from single result
		if uuid, ok := item["uuid"].(string); ok {
			data.UUID = types.StringValue(uuid)
		}

		sourceMap := item
		// Map response fields to data model
		_ = sourceMap
		if val, ok := sourceMap["public_key"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.PublicKey = types.StringValue(str)
			}
		} else {
			if data.PublicKey.IsUnknown() {
				data.PublicKey = types.StringNull()
			}
		}
		if val, ok := sourceMap["type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Type = types.StringValue(str)
			}
		} else {
			if data.Type.IsUnknown() {
				data.Type = types.StringNull()
			}
		}
		if val, ok := sourceMap["url"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Url = types.StringValue(str)
			}
		} else {
			if data.Url.IsUnknown() {
				data.Url = types.StringNull()
			}
		}

		// Map filter parameters from response if available
		if val, ok := sourceMap["created"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Created = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["fingerprint_md5"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.FingerprintMd5 = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["fingerprint_sha256"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.FingerprintSha256 = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["fingerprint_sha512"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.FingerprintSha512 = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["is_shared"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsShared = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["modified"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Modified = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Name = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name_exact"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NameExact = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["user_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.UserUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Uuid = types.StringValue(str)
			}
		}

	} else {
		// Filter by provided parameters
		var results []map[string]interface{}

		filters := map[string]string{}
		if !data.Created.IsNull() {
			filters["created"] = data.Created.ValueString()
		}
		if !data.FingerprintMd5.IsNull() {
			filters["fingerprint_md5"] = data.FingerprintMd5.ValueString()
		}
		if !data.FingerprintSha256.IsNull() {
			filters["fingerprint_sha256"] = data.FingerprintSha256.ValueString()
		}
		if !data.FingerprintSha512.IsNull() {
			filters["fingerprint_sha512"] = data.FingerprintSha512.ValueString()
		}
		if !data.IsShared.IsNull() {
			filters["is_shared"] = fmt.Sprintf("%t", data.IsShared.ValueBool())
		}
		if !data.Modified.IsNull() {
			filters["modified"] = data.Modified.ValueString()
		}
		if !data.Name.IsNull() {
			filters["name"] = data.Name.ValueString()
		}
		if !data.NameExact.IsNull() {
			filters["name_exact"] = data.NameExact.ValueString()
		}
		if !data.UserUuid.IsNull() {
			filters["user_uuid"] = data.UserUuid.ValueString()
		}
		if !data.Uuid.IsNull() {
			filters["uuid"] = data.Uuid.ValueString()
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
				"Unable to List SshPublicKey",
				"An error occurred while filtering core_ssh_public_key: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"SshPublicKey Not Found",
				"No core_ssh_public_key found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple SshPublicKeys Found",
				fmt.Sprintf("Found %d core_ssh_public_keys with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		// Extract data from the single result
		if uuid, ok := results[0]["uuid"].(string); ok {
			data.UUID = types.StringValue(uuid)
		}
		sourceMap := results[0]
		// Map response fields to data model
		_ = sourceMap
		if val, ok := sourceMap["public_key"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.PublicKey = types.StringValue(str)
			}
		} else {
			if data.PublicKey.IsUnknown() {
				data.PublicKey = types.StringNull()
			}
		}
		if val, ok := sourceMap["type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Type = types.StringValue(str)
			}
		} else {
			if data.Type.IsUnknown() {
				data.Type = types.StringNull()
			}
		}
		if val, ok := sourceMap["url"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Url = types.StringValue(str)
			}
		} else {
			if data.Url.IsUnknown() {
				data.Url = types.StringNull()
			}
		}

		// Map filter parameters from response if available
		if val, ok := sourceMap["created"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Created = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["fingerprint_md5"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.FingerprintMd5 = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["fingerprint_sha256"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.FingerprintSha256 = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["fingerprint_sha512"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.FingerprintSha512 = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["is_shared"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsShared = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["modified"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Modified = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Name = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["name_exact"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NameExact = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["user_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.UserUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Uuid = types.StringValue(str)
			}
		}
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
