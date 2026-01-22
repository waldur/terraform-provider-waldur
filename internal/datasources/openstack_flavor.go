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

		filters := make(map[string]string)
		if data.Filters != nil {
			type filterDef struct {
				name string
				val  attr.Value
			}
			filterDefs := []filterDef{
				{"cores", data.Filters.Cores},
				{"cores__gte", data.Filters.CoresGte},
				{"cores__lte", data.Filters.CoresLte},
				{"disk", data.Filters.Disk},
				{"disk__gte", data.Filters.DiskGte},
				{"disk__lte", data.Filters.DiskLte},
				{"name", data.Filters.Name},
				{"name_exact", data.Filters.NameExact},
				{"name_iregex", data.Filters.NameIregex},
				{"offering_uuid", data.Filters.OfferingUuid},
				{"ram", data.Filters.Ram},
				{"ram__gte", data.Filters.RamGte},
				{"ram__lte", data.Filters.RamLte},
				{"settings", data.Filters.Settings},
				{"settings_uuid", data.Filters.SettingsUuid},
				{"tenant", data.Filters.Tenant},
				{"tenant_uuid", data.Filters.TenantUuid},
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
