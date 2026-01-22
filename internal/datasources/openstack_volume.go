package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackVolumeDataSource{}

func NewOpenstackVolumeDataSource() datasource.DataSource {
	return &OpenstackVolumeDataSource{}
}

// OpenstackVolumeDataSource defines the data source implementation.
type OpenstackVolumeDataSource struct {
	client *client.Client
}

// OpenstackVolumeApiResponse is the API response model.
type OpenstackVolumeApiResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl                   *string `json:"access_url" tfsdk:"access_url"`
	Action                      *string `json:"action" tfsdk:"action"`
	AvailabilityZone            *string `json:"availability_zone" tfsdk:"availability_zone"`
	AvailabilityZoneName        *string `json:"availability_zone_name" tfsdk:"availability_zone_name"`
	BackendId                   *string `json:"backend_id" tfsdk:"backend_id"`
	Bootable                    *bool   `json:"bootable" tfsdk:"bootable"`
	Created                     *string `json:"created" tfsdk:"created"`
	Customer                    *string `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description                 *string `json:"description" tfsdk:"description"`
	Device                      *string `json:"device" tfsdk:"device"`
	ErrorMessage                *string `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback" tfsdk:"error_traceback"`
	ExtendEnabled               *bool   `json:"extend_enabled" tfsdk:"extend_enabled"`
	Image                       *string `json:"image" tfsdk:"image"`
	ImageMetadata               *string `json:"image_metadata" tfsdk:"image_metadata"`
	ImageName                   *string `json:"image_name" tfsdk:"image_name"`
	Instance                    *string `json:"instance" tfsdk:"instance"`
	InstanceMarketplaceUuid     *string `json:"instance_marketplace_uuid" tfsdk:"instance_marketplace_uuid"`
	InstanceName                *string `json:"instance_name" tfsdk:"instance_name"`
	IsLimitBased                *bool   `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased                *bool   `json:"is_usage_based" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string `json:"marketplace_category_name" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string `json:"marketplace_category_uuid" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string `json:"marketplace_offering_name" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string `json:"marketplace_offering_uuid" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string `json:"marketplace_plan_uuid" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string `json:"marketplace_resource_state" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string `json:"modified" tfsdk:"modified"`
	Name                        *string `json:"name" tfsdk:"name"`
	Project                     *string `json:"project" tfsdk:"project"`
	ProjectName                 *string `json:"project_name" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid" tfsdk:"project_uuid"`
	ResourceType                *string `json:"resource_type" tfsdk:"resource_type"`
	RuntimeState                *string `json:"runtime_state" tfsdk:"runtime_state"`
	ServiceName                 *string `json:"service_name" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	Size                        *int64  `json:"size" tfsdk:"size"`
	SourceSnapshot              *string `json:"source_snapshot" tfsdk:"source_snapshot"`
	State                       *string `json:"state" tfsdk:"state"`
	Tenant                      *string `json:"tenant" tfsdk:"tenant"`
	TenantUuid                  *string `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Type                        *string `json:"type" tfsdk:"type"`
	TypeName                    *string `json:"type_name" tfsdk:"type_name"`
	Url                         *string `json:"url" tfsdk:"url"`
}

// OpenstackVolumeFiltersModel contains the filter parameters for querying.
type OpenstackVolumeFiltersModel struct {
	AttachInstanceUuid   types.String `tfsdk:"attach_instance_uuid"`
	AvailabilityZoneName types.String `tfsdk:"availability_zone_name"`
	BackendId            types.String `tfsdk:"backend_id"`
	CanManage            types.Bool   `tfsdk:"can_manage"`
	Customer             types.String `tfsdk:"customer"`
	CustomerAbbreviation types.String `tfsdk:"customer_abbreviation"`
	CustomerName         types.String `tfsdk:"customer_name"`
	CustomerNativeName   types.String `tfsdk:"customer_native_name"`
	CustomerUuid         types.String `tfsdk:"customer_uuid"`
	Description          types.String `tfsdk:"description"`
	ExternalIp           types.String `tfsdk:"external_ip"`
	Instance             types.String `tfsdk:"instance"`
	InstanceUuid         types.String `tfsdk:"instance_uuid"`
	Name                 types.String `tfsdk:"name"`
	NameExact            types.String `tfsdk:"name_exact"`
	Project              types.String `tfsdk:"project"`
	ProjectName          types.String `tfsdk:"project_name"`
	ProjectUuid          types.String `tfsdk:"project_uuid"`
	RuntimeState         types.String `tfsdk:"runtime_state"`
	ServiceSettingsName  types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid  types.String `tfsdk:"service_settings_uuid"`
	Snapshot             types.String `tfsdk:"snapshot"`
	SnapshotUuid         types.String `tfsdk:"snapshot_uuid"`
	State                types.String `tfsdk:"state"`
	Tenant               types.String `tfsdk:"tenant"`
	TenantUuid           types.String `tfsdk:"tenant_uuid"`
	Uuid                 types.String `tfsdk:"uuid"`
}

// OpenstackVolumeDataSourceModel describes the data source data model.
type OpenstackVolumeDataSourceModel struct {
	UUID                        types.String                 `tfsdk:"id"`
	Filters                     *OpenstackVolumeFiltersModel `tfsdk:"filters"`
	AccessUrl                   types.String                 `tfsdk:"access_url"`
	Action                      types.String                 `tfsdk:"action"`
	AvailabilityZone            types.String                 `tfsdk:"availability_zone"`
	AvailabilityZoneName        types.String                 `tfsdk:"availability_zone_name"`
	BackendId                   types.String                 `tfsdk:"backend_id"`
	Bootable                    types.Bool                   `tfsdk:"bootable"`
	Created                     types.String                 `tfsdk:"created"`
	Customer                    types.String                 `tfsdk:"customer"`
	CustomerAbbreviation        types.String                 `tfsdk:"customer_abbreviation"`
	CustomerName                types.String                 `tfsdk:"customer_name"`
	CustomerNativeName          types.String                 `tfsdk:"customer_native_name"`
	CustomerUuid                types.String                 `tfsdk:"customer_uuid"`
	Description                 types.String                 `tfsdk:"description"`
	Device                      types.String                 `tfsdk:"device"`
	ErrorMessage                types.String                 `tfsdk:"error_message"`
	ErrorTraceback              types.String                 `tfsdk:"error_traceback"`
	ExtendEnabled               types.Bool                   `tfsdk:"extend_enabled"`
	Image                       types.String                 `tfsdk:"image"`
	ImageMetadata               types.String                 `tfsdk:"image_metadata"`
	ImageName                   types.String                 `tfsdk:"image_name"`
	Instance                    types.String                 `tfsdk:"instance"`
	InstanceMarketplaceUuid     types.String                 `tfsdk:"instance_marketplace_uuid"`
	InstanceName                types.String                 `tfsdk:"instance_name"`
	IsLimitBased                types.Bool                   `tfsdk:"is_limit_based"`
	IsUsageBased                types.Bool                   `tfsdk:"is_usage_based"`
	MarketplaceCategoryName     types.String                 `tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     types.String                 `tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     types.String                 `tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     types.String                 `tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         types.String                 `tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    types.String                 `tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     types.String                 `tfsdk:"marketplace_resource_uuid"`
	Modified                    types.String                 `tfsdk:"modified"`
	Name                        types.String                 `tfsdk:"name"`
	Project                     types.String                 `tfsdk:"project"`
	ProjectName                 types.String                 `tfsdk:"project_name"`
	ProjectUuid                 types.String                 `tfsdk:"project_uuid"`
	ResourceType                types.String                 `tfsdk:"resource_type"`
	RuntimeState                types.String                 `tfsdk:"runtime_state"`
	ServiceName                 types.String                 `tfsdk:"service_name"`
	ServiceSettings             types.String                 `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage types.String                 `tfsdk:"service_settings_error_message"`
	ServiceSettingsState        types.String                 `tfsdk:"service_settings_state"`
	ServiceSettingsUuid         types.String                 `tfsdk:"service_settings_uuid"`
	Size                        types.Int64                  `tfsdk:"size"`
	SourceSnapshot              types.String                 `tfsdk:"source_snapshot"`
	State                       types.String                 `tfsdk:"state"`
	Tenant                      types.String                 `tfsdk:"tenant"`
	TenantUuid                  types.String                 `tfsdk:"tenant_uuid"`
	Type                        types.String                 `tfsdk:"type"`
	TypeName                    types.String                 `tfsdk:"type_name"`
	Url                         types.String                 `tfsdk:"url"`
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
					"state": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "State Allowed values: `CREATING`, `CREATION_SCHEDULED`, `DELETING`, `DELETION_SCHEDULED`, `ERRED`, `OK`, `UPDATE_SCHEDULED`, `UPDATING`.",
						Validators: []validator.String{
							stringvalidator.OneOf("CREATING", "CREATION_SCHEDULED", "DELETING", "DELETION_SCHEDULED", "ERRED", "OK", "UPDATE_SCHEDULED", "UPDATING"),
						},
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
				Computed:            true,
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
				Computed:            true,
				MarkdownDescription: "Indicates if this volume can be used to boot an instance",
			},
			"created": schema.StringAttribute{
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
				Computed:            true,
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
				Computed:            true,
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
				Computed:            true,
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the resource",
			},
			"project": schema.StringAttribute{
				Computed:            true,
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
				Computed:            true,
				MarkdownDescription: "Size in MiB",
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
				Computed:            true,
				MarkdownDescription: "Tenant",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the tenant",
			},
			"type": schema.StringAttribute{
				Computed:            true,
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

func (d *OpenstackVolumeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackVolumeDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var apiResp OpenstackVolumeApiResponse

		err := d.client.GetByUUID(ctx, "/api/openstack-volumes/{uuid}/", data.UUID.ValueString(), &apiResp)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Volume",
				"An error occurred while reading the Openstack Volume by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, apiResp, &data)...)

	} else {
		// Filter by provided parameters
		var results []OpenstackVolumeApiResponse

		filters := make(map[string]string)
		if data.Filters != nil {
			type filterDef struct {
				name string
				val  attr.Value
			}
			filterDefs := []filterDef{
				{"attach_instance_uuid", data.Filters.AttachInstanceUuid},
				{"availability_zone_name", data.Filters.AvailabilityZoneName},
				{"backend_id", data.Filters.BackendId},
				{"can_manage", data.Filters.CanManage},
				{"customer", data.Filters.Customer},
				{"customer_abbreviation", data.Filters.CustomerAbbreviation},
				{"customer_name", data.Filters.CustomerName},
				{"customer_native_name", data.Filters.CustomerNativeName},
				{"customer_uuid", data.Filters.CustomerUuid},
				{"description", data.Filters.Description},
				{"external_ip", data.Filters.ExternalIp},
				{"instance", data.Filters.Instance},
				{"instance_uuid", data.Filters.InstanceUuid},
				{"name", data.Filters.Name},
				{"name_exact", data.Filters.NameExact},
				{"project", data.Filters.Project},
				{"project_name", data.Filters.ProjectName},
				{"project_uuid", data.Filters.ProjectUuid},
				{"runtime_state", data.Filters.RuntimeState},
				{"service_settings_name", data.Filters.ServiceSettingsName},
				{"service_settings_uuid", data.Filters.ServiceSettingsUuid},
				{"snapshot", data.Filters.Snapshot},
				{"snapshot_uuid", data.Filters.SnapshotUuid},
				{"state", data.Filters.State},
				{"tenant", data.Filters.Tenant},
				{"tenant_uuid", data.Filters.TenantUuid},
				{"uuid", data.Filters.Uuid},
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
				"At least one filter parameter (or 'id') must be provided to lookup openstack_volume.",
			)
			return
		}

		err := d.client.ListWithFilter(ctx, "/api/openstack-volumes/", filters, &results)
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

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *OpenstackVolumeDataSource) mapResponseToModel(ctx context.Context, apiResp OpenstackVolumeApiResponse, model *OpenstackVolumeDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.Action = types.StringPointerValue(apiResp.Action)
	model.AvailabilityZone = types.StringPointerValue(apiResp.AvailabilityZone)
	model.AvailabilityZoneName = types.StringPointerValue(apiResp.AvailabilityZoneName)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Bootable = types.BoolPointerValue(apiResp.Bootable)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Customer = types.StringPointerValue(apiResp.Customer)
	model.CustomerAbbreviation = types.StringPointerValue(apiResp.CustomerAbbreviation)
	model.CustomerName = types.StringPointerValue(apiResp.CustomerName)
	model.CustomerNativeName = types.StringPointerValue(apiResp.CustomerNativeName)
	model.CustomerUuid = types.StringPointerValue(apiResp.CustomerUuid)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.Device = types.StringPointerValue(apiResp.Device)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.ExtendEnabled = types.BoolPointerValue(apiResp.ExtendEnabled)
	model.Image = types.StringPointerValue(apiResp.Image)
	model.ImageMetadata = types.StringPointerValue(apiResp.ImageMetadata)
	model.ImageName = types.StringPointerValue(apiResp.ImageName)
	model.Instance = types.StringPointerValue(apiResp.Instance)
	model.InstanceMarketplaceUuid = types.StringPointerValue(apiResp.InstanceMarketplaceUuid)
	model.InstanceName = types.StringPointerValue(apiResp.InstanceName)
	model.IsLimitBased = types.BoolPointerValue(apiResp.IsLimitBased)
	model.IsUsageBased = types.BoolPointerValue(apiResp.IsUsageBased)
	model.MarketplaceCategoryName = types.StringPointerValue(apiResp.MarketplaceCategoryName)
	model.MarketplaceCategoryUuid = types.StringPointerValue(apiResp.MarketplaceCategoryUuid)
	model.MarketplaceOfferingName = types.StringPointerValue(apiResp.MarketplaceOfferingName)
	model.MarketplaceOfferingUuid = types.StringPointerValue(apiResp.MarketplaceOfferingUuid)
	model.MarketplacePlanUuid = types.StringPointerValue(apiResp.MarketplacePlanUuid)
	model.MarketplaceResourceState = types.StringPointerValue(apiResp.MarketplaceResourceState)
	model.MarketplaceResourceUuid = types.StringPointerValue(apiResp.MarketplaceResourceUuid)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.Project = types.StringPointerValue(apiResp.Project)
	model.ProjectName = types.StringPointerValue(apiResp.ProjectName)
	model.ProjectUuid = types.StringPointerValue(apiResp.ProjectUuid)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.RuntimeState = types.StringPointerValue(apiResp.RuntimeState)
	model.ServiceName = types.StringPointerValue(apiResp.ServiceName)
	model.ServiceSettings = types.StringPointerValue(apiResp.ServiceSettings)
	model.ServiceSettingsErrorMessage = types.StringPointerValue(apiResp.ServiceSettingsErrorMessage)
	model.ServiceSettingsState = types.StringPointerValue(apiResp.ServiceSettingsState)
	model.ServiceSettingsUuid = types.StringPointerValue(apiResp.ServiceSettingsUuid)
	model.Size = types.Int64PointerValue(apiResp.Size)
	model.SourceSnapshot = types.StringPointerValue(apiResp.SourceSnapshot)
	model.State = types.StringPointerValue(apiResp.State)
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
	model.Type = types.StringPointerValue(apiResp.Type)
	model.TypeName = types.StringPointerValue(apiResp.TypeName)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
