package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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

// OpenstackFlavorApiResponse is the API response model.
type OpenstackFlavorApiResponse struct {
	UUID *string `json:"uuid"`

	BackendId   *string `json:"backend_id" tfsdk:"backend_id"`
	Cores       *int64  `json:"cores" tfsdk:"cores"`
	Disk        *int64  `json:"disk" tfsdk:"disk"`
	DisplayName *string `json:"display_name" tfsdk:"display_name"`
	Name        *string `json:"name" tfsdk:"name"`
	Ram         *int64  `json:"ram" tfsdk:"ram"`
	Settings    *string `json:"settings" tfsdk:"settings"`
	Url         *string `json:"url" tfsdk:"url"`
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
		MarkdownDescription: "Openstack Flavor data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"cores": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Cores",
			},
			"cores__gte": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Cores gte",
			},
			"cores__lte": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Cores lte",
			},
			"disk": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Disk",
			},
			"disk__gte": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Disk gte",
			},
			"disk__lte": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Disk lte",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Name",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Name (exact)",
			},
			"name_iregex": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Name (regex)",
			},
			"offering_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Offering UUID",
			},
			"ram": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Ram",
			},
			"ram__gte": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Ram gte",
			},
			"ram__lte": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Ram lte",
			},
			"settings": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Settings URL",
			},
			"settings_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Settings UUID",
			},
			"tenant": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Tenant URL",
			},
			"tenant_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Tenant UUID",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of the backend",
			},
			"display_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the display",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Url",
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
		var apiResp OpenstackFlavorApiResponse

		err := d.client.GetByUUID(ctx, "/api/openstack-flavors/{uuid}/", data.UUID.ValueString(), &apiResp)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Flavor",
				"An error occurred while reading the Openstack Flavor by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, apiResp, &data)...)

	} else {
		// Filter by provided parameters
		var results []OpenstackFlavorApiResponse

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
				"Unable to List Openstack Flavor",
				"An error occurred while filtering Openstack Flavor: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Flavor Not Found",
				"No Openstack Flavor found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Flavors Found",
				fmt.Sprintf("Found %d Openstack Flavors with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *OpenstackFlavorDataSource) mapResponseToModel(ctx context.Context, apiResp OpenstackFlavorApiResponse, model *OpenstackFlavorDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Cores = types.Int64PointerValue(apiResp.Cores)
	model.Disk = types.Int64PointerValue(apiResp.Disk)
	model.DisplayName = types.StringPointerValue(apiResp.DisplayName)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.Ram = types.Int64PointerValue(apiResp.Ram)
	model.Settings = types.StringPointerValue(apiResp.Settings)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
