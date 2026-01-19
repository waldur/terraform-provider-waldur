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
var _ datasource.DataSource = &OpenstackFlavorDataSource{}

func NewOpenstackFlavorDataSource() datasource.DataSource {
	return &OpenstackFlavorDataSource{}
}

// OpenstackFlavorDataSource defines the data source implementation.
type OpenstackFlavorDataSource struct {
	client *client.Client
}

// OpenstackFlavorDataSourceModel describes the data source data model.
type OpenstackFlavorDataSourceModel struct {
	UUID         types.String `tfsdk:"id"`
	Cores        types.Int64  `tfsdk:"cores"`
	CoresGte     types.Int64  `tfsdk:"cores__gte"`
	CoresLte     types.Int64  `tfsdk:"cores__lte"`
	Disk         types.Int64  `tfsdk:"disk"`
	DiskGte      types.Int64  `tfsdk:"disk__gte"`
	DiskLte      types.Int64  `tfsdk:"disk__lte"`
	Name         types.String `tfsdk:"name"`
	NameExact    types.String `tfsdk:"name_exact"`
	NameIregex   types.String `tfsdk:"name_iregex"`
	OfferingUuid types.String `tfsdk:"offering_uuid"`
	Ram          types.Int64  `tfsdk:"ram"`
	RamGte       types.Int64  `tfsdk:"ram__gte"`
	RamLte       types.Int64  `tfsdk:"ram__lte"`
	Settings     types.String `tfsdk:"settings"`
	SettingsUuid types.String `tfsdk:"settings_uuid"`
	Tenant       types.String `tfsdk:"tenant"`
	TenantUuid   types.String `tfsdk:"tenant_uuid"`
	BackendId    types.String `tfsdk:"backend_id"`
	DisplayName  types.String `tfsdk:"display_name"`
	Url          types.String `tfsdk:"url"`
}

func (d *OpenstackFlavorDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_flavor"
}

func (d *OpenstackFlavorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "OpenstackFlavor data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"cores": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"cores__gte": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"cores__lte": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"disk": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"disk__gte": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"disk__lte": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"name_iregex": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"offering_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"ram": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"ram__gte": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: " ",
			},
			"ram__lte": schema.Int64Attribute{
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
			"display_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
		},
	}
}

func (d *OpenstackFlavorDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OpenstackFlavorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackFlavorDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var item map[string]interface{}

		err := d.client.GetByUUID(ctx, "/api/openstack-flavors/{uuid}/", data.UUID.ValueString(), &item)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Flavor",
				"An error occurred while reading the openstack_flavor by UUID: "+err.Error(),
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
		if val, ok := sourceMap["display_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.DisplayName = types.StringValue(str)
			}
		} else {
			if data.DisplayName.IsUnknown() {
				data.DisplayName = types.StringNull()
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
		if val, ok := sourceMap["cores"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Cores = types.Int64Value(int64(num))
			}
		}
		if val, ok := sourceMap["cores__gte"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.CoresGte = types.Int64Value(int64(num))
			}
		}
		if val, ok := sourceMap["cores__lte"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.CoresLte = types.Int64Value(int64(num))
			}
		}
		if val, ok := sourceMap["disk"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Disk = types.Int64Value(int64(num))
			}
		}
		if val, ok := sourceMap["disk__gte"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.DiskGte = types.Int64Value(int64(num))
			}
		}
		if val, ok := sourceMap["disk__lte"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.DiskLte = types.Int64Value(int64(num))
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
		if val, ok := sourceMap["name_iregex"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NameIregex = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["ram"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Ram = types.Int64Value(int64(num))
			}
		}
		if val, ok := sourceMap["ram__gte"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.RamGte = types.Int64Value(int64(num))
			}
		}
		if val, ok := sourceMap["ram__lte"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.RamLte = types.Int64Value(int64(num))
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
		if !data.Cores.IsNull() {
			filters["cores"] = fmt.Sprintf("%d", data.Cores.ValueInt64())
		}
		if !data.CoresGte.IsNull() {
			filters["cores__gte"] = fmt.Sprintf("%d", data.CoresGte.ValueInt64())
		}
		if !data.CoresLte.IsNull() {
			filters["cores__lte"] = fmt.Sprintf("%d", data.CoresLte.ValueInt64())
		}
		if !data.Disk.IsNull() {
			filters["disk"] = fmt.Sprintf("%d", data.Disk.ValueInt64())
		}
		if !data.DiskGte.IsNull() {
			filters["disk__gte"] = fmt.Sprintf("%d", data.DiskGte.ValueInt64())
		}
		if !data.DiskLte.IsNull() {
			filters["disk__lte"] = fmt.Sprintf("%d", data.DiskLte.ValueInt64())
		}
		if !data.Name.IsNull() {
			filters["name"] = data.Name.ValueString()
		}
		if !data.NameExact.IsNull() {
			filters["name_exact"] = data.NameExact.ValueString()
		}
		if !data.NameIregex.IsNull() {
			filters["name_iregex"] = data.NameIregex.ValueString()
		}
		if !data.OfferingUuid.IsNull() {
			filters["offering_uuid"] = data.OfferingUuid.ValueString()
		}
		if !data.Ram.IsNull() {
			filters["ram"] = fmt.Sprintf("%d", data.Ram.ValueInt64())
		}
		if !data.RamGte.IsNull() {
			filters["ram__gte"] = fmt.Sprintf("%d", data.RamGte.ValueInt64())
		}
		if !data.RamLte.IsNull() {
			filters["ram__lte"] = fmt.Sprintf("%d", data.RamLte.ValueInt64())
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
				"At least one filter parameter (or 'id') must be provided to lookup openstack_flavor.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/openstack-flavors/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Flavor",
				"An error occurred while filtering openstack_flavor: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Flavor Not Found",
				"No openstack_flavor found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Flavors Found",
				fmt.Sprintf("Found %d openstack_flavors with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
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
		if val, ok := sourceMap["display_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.DisplayName = types.StringValue(str)
			}
		} else {
			if data.DisplayName.IsUnknown() {
				data.DisplayName = types.StringNull()
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
		if val, ok := sourceMap["cores"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Cores = types.Int64Value(int64(num))
			}
		}
		if val, ok := sourceMap["cores__gte"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.CoresGte = types.Int64Value(int64(num))
			}
		}
		if val, ok := sourceMap["cores__lte"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.CoresLte = types.Int64Value(int64(num))
			}
		}
		if val, ok := sourceMap["disk"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Disk = types.Int64Value(int64(num))
			}
		}
		if val, ok := sourceMap["disk__gte"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.DiskGte = types.Int64Value(int64(num))
			}
		}
		if val, ok := sourceMap["disk__lte"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.DiskLte = types.Int64Value(int64(num))
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
		if val, ok := sourceMap["name_iregex"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.NameIregex = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["offering_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.OfferingUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["ram"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Ram = types.Int64Value(int64(num))
			}
		}
		if val, ok := sourceMap["ram__gte"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.RamGte = types.Int64Value(int64(num))
			}
		}
		if val, ok := sourceMap["ram__lte"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.RamLte = types.Int64Value(int64(num))
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
