package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackImageDataSource{}

func NewOpenstackImageDataSource() datasource.DataSource {
	return &OpenstackImageDataSource{}
}

// OpenstackImageDataSource defines the data source implementation.
type OpenstackImageDataSource struct {
	client *client.Client
}

// OpenstackImageDataSourceModel describes the data source data model.
type OpenstackImageDataSourceModel struct {
	UUID         types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	NameExact    types.String `tfsdk:"name_exact"`
	OfferingUuid types.String `tfsdk:"offering_uuid"`
	Settings     types.String `tfsdk:"settings"`
	SettingsUuid types.String `tfsdk:"settings_uuid"`
	Tenant       types.String `tfsdk:"tenant"`
	TenantUuid   types.String `tfsdk:"tenant_uuid"`
	BackendId    types.String `tfsdk:"backend_id"`
	MinDisk      types.Int64  `tfsdk:"min_disk"`
	MinRam       types.Int64  `tfsdk:"min_ram"`
	Url          types.String `tfsdk:"url"`
}

func (d *OpenstackImageDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_image"
}

func (d *OpenstackImageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "OpenstackImage data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"offering_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"settings": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"settings_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"tenant": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"tenant_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"min_disk": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Minimum disk size in MiB",
			},
			"min_ram": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Minimum memory size in MiB",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
		},
	}
}

func (d *OpenstackImageDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OpenstackImageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackImageDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var item map[string]interface{}

		err := d.client.GetByUUID(ctx, "/api/openstack-images/{uuid}/", data.UUID.ValueString(), &item)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Image",
				"An error occurred while reading the openstack_image by UUID: "+err.Error(),
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
		if val, ok := sourceMap["backend_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BackendId = types.StringValue(str)
			}
		} else {
			if data.BackendId.IsUnknown() {
				data.BackendId = types.StringNull()
			}
		}
		if val, ok := sourceMap["min_disk"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.MinDisk = types.Int64Value(int64(num))
			}
		} else {
			if data.MinDisk.IsUnknown() {
				data.MinDisk = types.Int64Null()
			}
		}
		if val, ok := sourceMap["min_ram"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.MinRam = types.Int64Value(int64(num))
			}
		} else {
			if data.MinRam.IsUnknown() {
				data.MinRam = types.Int64Null()
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
		if val, ok := sourceMap["offering_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["settings"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Settings = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["settings_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.SettingsUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["tenant"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Tenant = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["tenant_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.TenantUuid = types.StringValue(str)
			}
		}

	} else {
		// Filter by provided parameters
		var results []map[string]interface{}

		filters := map[string]string{}
		if !data.Name.IsNull() {
			filters["name"] = data.Name.ValueString()
		}
		if !data.NameExact.IsNull() {
			filters["name_exact"] = data.NameExact.ValueString()
		}
		if !data.OfferingUuid.IsNull() {
			filters["offering_uuid"] = data.OfferingUuid.ValueString()
		}
		if !data.Settings.IsNull() {
			filters["settings"] = data.Settings.ValueString()
		}
		if !data.SettingsUuid.IsNull() {
			filters["settings_uuid"] = data.SettingsUuid.ValueString()
		}
		if !data.Tenant.IsNull() {
			filters["tenant"] = data.Tenant.ValueString()
		}
		if !data.TenantUuid.IsNull() {
			filters["tenant_uuid"] = data.TenantUuid.ValueString()
		}

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_image.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/openstack-images/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Image",
				"An error occurred while filtering openstack_image: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Image Not Found",
				"No openstack_image found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Images Found",
				fmt.Sprintf("Found %d openstack_images with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
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
		if val, ok := sourceMap["backend_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BackendId = types.StringValue(str)
			}
		} else {
			if data.BackendId.IsUnknown() {
				data.BackendId = types.StringNull()
			}
		}
		if val, ok := sourceMap["min_disk"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.MinDisk = types.Int64Value(int64(num))
			}
		} else {
			if data.MinDisk.IsUnknown() {
				data.MinDisk = types.Int64Null()
			}
		}
		if val, ok := sourceMap["min_ram"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.MinRam = types.Int64Value(int64(num))
			}
		} else {
			if data.MinRam.IsUnknown() {
				data.MinRam = types.Int64Null()
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
		if val, ok := sourceMap["offering_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["settings"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Settings = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["settings_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.SettingsUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["tenant"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Tenant = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["tenant_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.TenantUuid = types.StringValue(str)
			}
		}
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
