package volume

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackVolumeDataSource{}

func NewOpenstackVolumeDataSource() datasource.DataSource {
	return &OpenstackVolumeDataSource{}
}

type OpenstackVolumeDataSource struct {
	client *Client
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
				MarkdownDescription: "Resource UUID",
			},
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Filter parameters for querying Openstack Volume",
				Attributes: map[string]schema.Attribute{
					"attach_instance_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Filter for attachment to instance UUID",
					},
					"availability_zone_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Availability zone name",
					},
					"backend_id": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Backend ID",
					},
					"can_manage": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Can manage",
					},
					"customer": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Customer UUID",
					},
					"customer_abbreviation": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Customer abbreviation",
					},
					"customer_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Customer name",
					},
					"customer_native_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Customer native name",
					},
					"customer_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Customer UUID",
					},
					"description": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Description",
					},
					"external_ip": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "External IP",
					},
					"instance": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Instance URL",
					},
					"instance_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Instance UUID",
					},
					"name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name",
					},
					"name_exact": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name (exact)",
					},
					"project": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Project UUID",
					},
					"project_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Project name",
					},
					"project_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Project UUID",
					},
					"runtime_state": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Runtime state",
					},
					"service_settings_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Service settings name",
					},
					"service_settings_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Service settings UUID",
					},
					"snapshot": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Snapshot URL",
					},
					"snapshot_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Snapshot UUID",
					},
					"tenant": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Tenant URL",
					},
					"tenant_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Tenant UUID",
					},
					"uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "UUID",
					},
				},
			},
			"access_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Access url",
			},
			"action": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Action",
			},
			"availability_zone": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Availability zone where this volume is located",
			},
			"availability_zone_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the availability zone",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Volume ID in the OpenStack backend",
			},
			"bootable": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Indicates if this volume can be used to boot an instance",
			},
			"created": schema.StringAttribute{
				CustomType:          timetypes.RFC3339Type{},
				Computed:            true,
				MarkdownDescription: "Created",
			},
			"customer": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Customer",
			},
			"customer_abbreviation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Customer abbreviation",
			},
			"customer_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the customer",
			},
			"customer_native_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the customer native",
			},
			"customer_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the customer",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description of the resource",
			},
			"device": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of volume as instance device e.g. /dev/vdb.",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error message",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error traceback",
			},
			"extend_enabled": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Extend enabled",
			},
			"image": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Image that this volume was created from, if any",
			},
			"image_metadata": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Metadata of the image this volume was created from",
			},
			"image_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the image this volume was created from",
			},
			"instance": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Instance that this volume is attached to, if any",
			},
			"instance_marketplace_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the instance marketplace",
			},
			"instance_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the instance",
			},
			"is_limit_based": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Is limit based",
			},
			"is_usage_based": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Is usage based",
			},
			"marketplace_category_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the marketplace category",
			},
			"marketplace_category_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the marketplace category",
			},
			"marketplace_offering_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the marketplace offering",
			},
			"marketplace_offering_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the marketplace offering",
			},
			"marketplace_plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the marketplace plan",
			},
			"marketplace_resource_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Marketplace resource state",
			},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the marketplace resource",
			},
			"modified": schema.StringAttribute{
				CustomType:          timetypes.RFC3339Type{},
				Computed:            true,
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name of the resource",
			},
			"project": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project",
			},
			"project_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the project",
			},
			"project_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the project",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource type",
			},
			"runtime_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Runtime state",
			},
			"service_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the service",
			},
			"service_settings": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Service settings",
			},
			"service_settings_error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Service settings error message",
			},
			"service_settings_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Service settings state",
			},
			"service_settings_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the service settings",
			},
			"size": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Size in MiB",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
					int64validator.AtMost(2147483647),
				},
			},
			"source_snapshot": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Snapshot that this volume was created from, if any",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "State",
			},
			"tenant": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Tenant",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the tenant",
			},
			"type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Type of the volume (e.g. SSD, HDD)",
			},
			"type_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the type",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Url",
			},
		},
	}
}

func (d *OpenstackVolumeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OpenstackVolumeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackVolumeDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		apiResp, err := d.client.GetOpenstackVolume(ctx, data.UUID.ValueString())
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

		results, err := d.client.ListOpenstackVolume(ctx, filters)
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
