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
var _ datasource.DataSource = &OpenstackImageDataSource{}

func NewOpenstackImageDataSource() datasource.DataSource {
	return &OpenstackImageDataSource{}
}

// OpenstackImageDataSource defines the data source implementation.
type OpenstackImageDataSource struct {
	client *client.Client
}

// OpenstackImageApiResponse is the API response model.
type OpenstackImageApiResponse struct {
	UUID *string `json:"uuid"`

	BackendId *string `json:"backend_id" tfsdk:"backend_id"`
	MinDisk   *int64  `json:"min_disk" tfsdk:"min_disk"`
	MinRam    *int64  `json:"min_ram" tfsdk:"min_ram"`
	Name      *string `json:"name" tfsdk:"name"`
	Settings  *string `json:"settings" tfsdk:"settings"`
	Url       *string `json:"url" tfsdk:"url"`
}

// OpenstackImageFiltersModel contains the filter parameters for querying.
type OpenstackImageFiltersModel struct {
	Name         types.String `tfsdk:"name"`
	NameExact    types.String `tfsdk:"name_exact"`
	OfferingUuid types.String `tfsdk:"offering_uuid"`
	Settings     types.String `tfsdk:"settings"`
	SettingsUuid types.String `tfsdk:"settings_uuid"`
	Tenant       types.String `tfsdk:"tenant"`
	TenantUuid   types.String `tfsdk:"tenant_uuid"`
}

// OpenstackImageDataSourceModel describes the data source data model.
type OpenstackImageDataSourceModel struct {
	UUID      types.String                `tfsdk:"id"`
	Filters   *OpenstackImageFiltersModel `tfsdk:"filters"`
	BackendId types.String                `tfsdk:"backend_id"`
	MinDisk   types.Int64                 `tfsdk:"min_disk"`
	MinRam    types.Int64                 `tfsdk:"min_ram"`
	Name      types.String                `tfsdk:"name"`
	Settings  types.String                `tfsdk:"settings"`
	Url       types.String                `tfsdk:"url"`
}

func (d *OpenstackImageDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_image"
}

func (d *OpenstackImageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Image data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Filter parameters for querying Openstack Image",
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name",
					},
					"name_exact": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name (exact)",
					},
					"offering_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Offering UUID",
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
			"min_disk": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Minimum disk size in MiB",
			},
			"min_ram": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Minimum memory size in MiB",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the resource",
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
		var apiResp OpenstackImageApiResponse

		err := d.client.GetByUUID(ctx, "/api/openstack-images/{uuid}/", data.UUID.ValueString(), &apiResp)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Image",
				"An error occurred while reading the Openstack Image by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, apiResp, &data)...)

	} else {
		// Filter by provided parameters
		var results []OpenstackImageApiResponse

		filters := make(map[string]string)
		if data.Filters != nil {
			type filterDef struct {
				name string
				val  attr.Value
			}
			filterDefs := []filterDef{
				{"name", data.Filters.Name},
				{"name_exact", data.Filters.NameExact},
				{"offering_uuid", data.Filters.OfferingUuid},
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
				"At least one filter parameter (or 'id') must be provided to lookup openstack_image.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/openstack-images/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Openstack Image",
				"An error occurred while filtering Openstack Image: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Image Not Found",
				"No Openstack Image found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Images Found",
				fmt.Sprintf("Found %d Openstack Images with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *OpenstackImageDataSource) mapResponseToModel(ctx context.Context, apiResp OpenstackImageApiResponse, model *OpenstackImageDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.MinDisk = types.Int64PointerValue(apiResp.MinDisk)
	model.MinRam = types.Int64PointerValue(apiResp.MinRam)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.Settings = types.StringPointerValue(apiResp.Settings)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
