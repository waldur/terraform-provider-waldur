package volume

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackVolumeDataSource{}

func NewOpenstackVolumeDataSource() datasource.DataSource {
	return &OpenstackVolumeDataSource{}
}

type OpenstackVolumeDataSource struct {
	client *OpenstackVolumeClient
}

type OpenstackVolumeDataSourceModel struct {
	OpenstackVolumeModel
	Filters *OpenstackVolumeFiltersModel `tfsdk:"filters"`
}

func (d *OpenstackVolumeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_volume"
}

func (d *OpenstackVolumeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Volume data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Openstack Volume UUID",
			},
			"filters": (&OpenstackVolumeFiltersModel{}).GetSchema(),
			"availability_zone": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Availability zone where this volume is located"},
			"availability_zone_name": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Availability Zone Name"},
			"backend_id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Volume ID in the OpenStack backend"},
			"bootable": schema.BoolAttribute{
				Computed: true, MarkdownDescription: "Indicates if this volume can be used to boot an instance"},
			"customer": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Customer"},
			"description": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Description"},
			"device": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Name of volume as instance device e.g. /dev/vdb."},
			"error_message": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Error Message"},
			"extend_enabled": schema.BoolAttribute{
				Computed: true, MarkdownDescription: "Extend Enabled"},
			"image": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Image that this volume was created from, if any"},
			"image_metadata": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Metadata of the image this volume was created from"},
			"image_name": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Name of the image this volume was created from"},
			"instance": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Instance that this volume is attached to, if any"},
			"instance_marketplace_uuid": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Instance Marketplace Uuid"},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Marketplace Resource Uuid"},
			"name": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Name"},
			"project": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Project URL"},
			"resource_type": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Resource Type"},
			"runtime_state": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Runtime State"},
			"size": schema.Int64Attribute{
				Computed: true, MarkdownDescription: "Size in MiB",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
					int64validator.AtMost(2147483647),
				}},
			"source_snapshot": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Snapshot that this volume was created from, if any"},
			"state": schema.StringAttribute{
				Computed: true, MarkdownDescription: "State"},
			"tenant": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Tenant"},
			"tenant_uuid": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Tenant Uuid"},
			"type": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Type of the volume (e.g. SSD, HDD)"},
			"type_name": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Type Name"},
			"url": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Url"},
		},
	}
}

func (d *OpenstackVolumeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	d.client = &OpenstackVolumeClient{}
	if err := d.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			err.Error(),
		)
		return
	}
}

func (d *OpenstackVolumeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackVolumeDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		apiResp, err := d.client.Get(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Volume",
				"An error occurred while reading the Openstack Volume by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_volume.",
			)
			return
		}

		results, err := d.client.List(ctx, filters)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Openstack Volume",
				"An error occurred while filtering Openstack Volume: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Volume Not Found",
				"No Openstack Volume found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Volumes Found",
				fmt.Sprintf("Found %d Openstack Volumes with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, results[0])...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
