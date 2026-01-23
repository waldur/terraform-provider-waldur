package flavor

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackFlavorDataSource{}

func NewOpenstackFlavorDataSource() datasource.DataSource {
	return &OpenstackFlavorDataSource{}
}

// OpenstackFlavorDataSource defines the data source implementation.
type OpenstackFlavorDataSource struct {
	client *Client
}

// OpenstackFlavorFiltersModel contains the filter parameters for querying.
type OpenstackFlavorFiltersModel struct {
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
}

// OpenstackFlavorDataSourceModel describes the data source data model.
type OpenstackFlavorDataSourceModel struct {
	UUID        types.String                 `tfsdk:"id"`
	Filters     *OpenstackFlavorFiltersModel `tfsdk:"filters"`
	BackendId   types.String                 `tfsdk:"backend_id"`
	Cores       types.Int64                  `tfsdk:"cores"`
	Disk        types.Int64                  `tfsdk:"disk"`
	DisplayName types.String                 `tfsdk:"display_name"`
	Name        types.String                 `tfsdk:"name"`
	Ram         types.Int64                  `tfsdk:"ram"`
	Settings    types.String                 `tfsdk:"settings"`
	Url         types.String                 `tfsdk:"url"`
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
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Filter parameters for querying Openstack Flavor",
				Attributes: map[string]schema.Attribute{
					"cores": schema.Int64Attribute{
						Optional:            true,
						MarkdownDescription: "Cores",
					},
					"cores__gte": schema.Int64Attribute{
						Optional:            true,
						MarkdownDescription: "Cores gte",
					},
					"cores__lte": schema.Int64Attribute{
						Optional:            true,
						MarkdownDescription: "Cores lte",
					},
					"disk": schema.Int64Attribute{
						Optional:            true,
						MarkdownDescription: "Disk",
					},
					"disk__gte": schema.Int64Attribute{
						Optional:            true,
						MarkdownDescription: "Disk gte",
					},
					"disk__lte": schema.Int64Attribute{
						Optional:            true,
						MarkdownDescription: "Disk lte",
					},
					"name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name",
					},
					"name_exact": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name (exact)",
					},
					"name_iregex": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name (regex)",
					},
					"offering_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Offering UUID",
					},
					"ram": schema.Int64Attribute{
						Optional:            true,
						MarkdownDescription: "Ram",
					},
					"ram__gte": schema.Int64Attribute{
						Optional:            true,
						MarkdownDescription: "Ram gte",
					},
					"ram__lte": schema.Int64Attribute{
						Optional:            true,
						MarkdownDescription: "Ram lte",
					},
					"settings": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Settings URL",
					},
					"settings_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Settings UUID",
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
			"cores": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of cores in a VM",
			},
			"disk": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Root disk size in MiB",
			},
			"display_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the display",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the resource",
			},
			"ram": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Memory size in MiB",
			},
			"settings": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Settings",
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

func (d *OpenstackFlavorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackFlavorDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		apiResp, err := d.client.GetOpenstackFlavor(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Flavor",
				"An error occurred while reading the Openstack Flavor by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, *apiResp, &data)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_flavor.",
			)
			return
		}

		results, err := d.client.ListOpenstackFlavor(ctx, filters)
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

func (d *OpenstackFlavorDataSource) mapResponseToModel(ctx context.Context, apiResp OpenstackFlavorResponse, model *OpenstackFlavorDataSourceModel) diag.Diagnostics {
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
