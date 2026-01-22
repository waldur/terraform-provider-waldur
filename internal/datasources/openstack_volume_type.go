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
var _ datasource.DataSource = &OpenstackVolumeTypeDataSource{}

func NewOpenstackVolumeTypeDataSource() datasource.DataSource {
	return &OpenstackVolumeTypeDataSource{}
}

// OpenstackVolumeTypeDataSource defines the data source implementation.
type OpenstackVolumeTypeDataSource struct {
	client *client.Client
}

// OpenstackVolumeTypeApiResponse is the API response model.
type OpenstackVolumeTypeApiResponse struct {
	UUID *string `json:"uuid"`

	Description *string `json:"description" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
	Settings    *string `json:"settings" tfsdk:"settings"`
	Url         *string `json:"url" tfsdk:"url"`
}

// OpenstackVolumeTypeFiltersModel contains the filter parameters for querying.
type OpenstackVolumeTypeFiltersModel struct {
	Name         types.String `tfsdk:"name"`
	NameExact    types.String `tfsdk:"name_exact"`
	OfferingUuid types.String `tfsdk:"offering_uuid"`
	Settings     types.String `tfsdk:"settings"`
	SettingsUuid types.String `tfsdk:"settings_uuid"`
	Tenant       types.String `tfsdk:"tenant"`
	TenantUuid   types.String `tfsdk:"tenant_uuid"`
}

// OpenstackVolumeTypeDataSourceModel describes the data source data model.
type OpenstackVolumeTypeDataSourceModel struct {
	UUID        types.String                     `tfsdk:"id"`
	Filters     *OpenstackVolumeTypeFiltersModel `tfsdk:"filters"`
	Description types.String                     `tfsdk:"description"`
	Name        types.String                     `tfsdk:"name"`
	Settings    types.String                     `tfsdk:"settings"`
	Url         types.String                     `tfsdk:"url"`
}

func (d *OpenstackVolumeTypeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_volume_type"
}

func (d *OpenstackVolumeTypeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Volume Type data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Filter parameters for querying Openstack Volume Type",
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
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Description of the resource",
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

func (d *OpenstackVolumeTypeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OpenstackVolumeTypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackVolumeTypeDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var apiResp OpenstackVolumeTypeApiResponse

		err := d.client.GetByUUID(ctx, "/api/openstack-volume-types/{uuid}/", data.UUID.ValueString(), &apiResp)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Volume Type",
				"An error occurred while reading the Openstack Volume Type by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, apiResp, &data)...)

	} else {
		// Filter by provided parameters
		var results []OpenstackVolumeTypeApiResponse

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
				"At least one filter parameter (or 'id') must be provided to lookup openstack_volume_type.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/openstack-volume-types/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Openstack Volume Type",
				"An error occurred while filtering Openstack Volume Type: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Volume Type Not Found",
				"No Openstack Volume Type found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Volume Types Found",
				fmt.Sprintf("Found %d Openstack Volume Types with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *OpenstackVolumeTypeDataSource) mapResponseToModel(ctx context.Context, apiResp OpenstackVolumeTypeApiResponse, model *OpenstackVolumeTypeDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.Settings = types.StringPointerValue(apiResp.Settings)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
