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
var _ datasource.DataSource = &OpenstackVolumeDataSource{}

func NewOpenstackVolumeDataSource() datasource.DataSource {
	return &OpenstackVolumeDataSource{}
}

// OpenstackVolumeDataSource defines the data source implementation.
type OpenstackVolumeDataSource struct {
	client *client.Client
}

// OpenstackVolumeDataSourceModel describes the data source data model.
type OpenstackVolumeDataSourceModel struct {
	UUID                        types.String `tfsdk:"id"`
	AttachInstanceUuid          types.String `tfsdk:"attach_instance_uuid"`
	AvailabilityZoneName        types.String `tfsdk:"availability_zone_name"`
	BackendId                   types.String `tfsdk:"backend_id"`
	CanManage                   types.Bool   `tfsdk:"can_manage"`
	Customer                    types.String `tfsdk:"customer"`
	CustomerAbbreviation        types.String `tfsdk:"customer_abbreviation"`
	CustomerName                types.String `tfsdk:"customer_name"`
	CustomerNativeName          types.String `tfsdk:"customer_native_name"`
	CustomerUuid                types.String `tfsdk:"customer_uuid"`
	Description                 types.String `tfsdk:"description"`
	ExternalIp                  types.String `tfsdk:"external_ip"`
	Instance                    types.String `tfsdk:"instance"`
	InstanceUuid                types.String `tfsdk:"instance_uuid"`
	Name                        types.String `tfsdk:"name"`
	NameExact                   types.String `tfsdk:"name_exact"`
	Project                     types.String `tfsdk:"project"`
	ProjectName                 types.String `tfsdk:"project_name"`
	ProjectUuid                 types.String `tfsdk:"project_uuid"`
	RuntimeState                types.String `tfsdk:"runtime_state"`
	ServiceSettingsName         types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid         types.String `tfsdk:"service_settings_uuid"`
	Snapshot                    types.String `tfsdk:"snapshot"`
	SnapshotUuid                types.String `tfsdk:"snapshot_uuid"`
	State                       types.String `tfsdk:"state"`
	Tenant                      types.String `tfsdk:"tenant"`
	TenantUuid                  types.String `tfsdk:"tenant_uuid"`
	Uuid                        types.String `tfsdk:"uuid"`
	AccessUrl                   types.String `tfsdk:"access_url"`
	Action                      types.String `tfsdk:"action"`
	AvailabilityZone            types.String `tfsdk:"availability_zone"`
	Bootable                    types.Bool   `tfsdk:"bootable"`
	Created                     types.String `tfsdk:"created"`
	Device                      types.String `tfsdk:"device"`
	ErrorMessage                types.String `tfsdk:"error_message"`
	ErrorTraceback              types.String `tfsdk:"error_traceback"`
	ExtendEnabled               types.Bool   `tfsdk:"extend_enabled"`
	Image                       types.String `tfsdk:"image"`
	ImageMetadata               types.String `tfsdk:"image_metadata"`
	ImageName                   types.String `tfsdk:"image_name"`
	InstanceMarketplaceUuid     types.String `tfsdk:"instance_marketplace_uuid"`
	InstanceName                types.String `tfsdk:"instance_name"`
	IsLimitBased                types.Bool   `tfsdk:"is_limit_based"`
	IsUsageBased                types.Bool   `tfsdk:"is_usage_based"`
	MarketplaceCategoryName     types.String `tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     types.String `tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     types.String `tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     types.String `tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         types.String `tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    types.String `tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     types.String `tfsdk:"marketplace_resource_uuid"`
	Modified                    types.String `tfsdk:"modified"`
	ResourceType                types.String `tfsdk:"resource_type"`
	ServiceName                 types.String `tfsdk:"service_name"`
	ServiceSettings             types.String `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage types.String `tfsdk:"service_settings_error_message"`
	ServiceSettingsState        types.String `tfsdk:"service_settings_state"`
	Size                        types.Int64  `tfsdk:"size"`
	SourceSnapshot              types.String `tfsdk:"source_snapshot"`
	Type                        types.String `tfsdk:"type"`
	TypeName                    types.String `tfsdk:"type_name"`
	Url                         types.String `tfsdk:"url"`
}

func (d *OpenstackVolumeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_volume"
}

func (d *OpenstackVolumeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "OpenstackVolume data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"attach_instance_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"availability_zone_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"can_manage": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Can manage",
			},
			"customer": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"customer_abbreviation": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"customer_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"customer_native_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"customer_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"external_ip": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"instance": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"instance_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"project": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"project_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"project_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"runtime_state": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"service_settings_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"service_settings_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"snapshot": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"snapshot_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"state": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"tenant": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"tenant_uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"uuid": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"access_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"action": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"availability_zone": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Availability zone where this volume is located",
			},
			"bootable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Indicates if this volume can be used to boot an instance",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"device": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of volume as instance device e.g. /dev/vdb.",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"extend_enabled": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "",
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
			"instance_marketplace_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"instance_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"is_limit_based": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"is_usage_based": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_category_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_category_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_offering_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_offering_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_resource_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"service_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"service_settings": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"service_settings_error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"service_settings_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"size": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Size in MiB",
			},
			"source_snapshot": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Snapshot that this volume was created from, if any",
			},
			"type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Type of the volume (e.g. SSD, HDD)",
			},
			"type_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
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
		var item map[string]interface{}

		err := d.client.GetByUUID(ctx, "/api/openstack-volumes/{uuid}/", data.UUID.ValueString(), &item)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Volume",
				"An error occurred while reading the openstack_volume by UUID: "+err.Error(),
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
		if val, ok := sourceMap["access_url"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AccessUrl = types.StringValue(str)
			}
		} else {
			if data.AccessUrl.IsUnknown() {
				data.AccessUrl = types.StringNull()
			}
		}
		if val, ok := sourceMap["action"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Action = types.StringValue(str)
			}
		} else {
			if data.Action.IsUnknown() {
				data.Action = types.StringNull()
			}
		}
		if val, ok := sourceMap["availability_zone"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AvailabilityZone = types.StringValue(str)
			}
		} else {
			if data.AvailabilityZone.IsUnknown() {
				data.AvailabilityZone = types.StringNull()
			}
		}
		if val, ok := sourceMap["bootable"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.Bootable = types.BoolValue(b)
			}
		} else {
			if data.Bootable.IsUnknown() {
				data.Bootable = types.BoolNull()
			}
		}
		if val, ok := sourceMap["created"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Created = types.StringValue(str)
			}
		} else {
			if data.Created.IsUnknown() {
				data.Created = types.StringNull()
			}
		}
		if val, ok := sourceMap["device"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Device = types.StringValue(str)
			}
		} else {
			if data.Device.IsUnknown() {
				data.Device = types.StringNull()
			}
		}
		if val, ok := sourceMap["error_message"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ErrorMessage = types.StringValue(str)
			}
		} else {
			if data.ErrorMessage.IsUnknown() {
				data.ErrorMessage = types.StringNull()
			}
		}
		if val, ok := sourceMap["error_traceback"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ErrorTraceback = types.StringValue(str)
			}
		} else {
			if data.ErrorTraceback.IsUnknown() {
				data.ErrorTraceback = types.StringNull()
			}
		}
		if val, ok := sourceMap["extend_enabled"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.ExtendEnabled = types.BoolValue(b)
			}
		} else {
			if data.ExtendEnabled.IsUnknown() {
				data.ExtendEnabled = types.BoolNull()
			}
		}
		if val, ok := sourceMap["image"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Image = types.StringValue(str)
			}
		} else {
			if data.Image.IsUnknown() {
				data.Image = types.StringNull()
			}
		}
		if val, ok := sourceMap["image_metadata"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ImageMetadata = types.StringValue(str)
			}
		} else {
			if data.ImageMetadata.IsUnknown() {
				data.ImageMetadata = types.StringNull()
			}
		}
		if val, ok := sourceMap["image_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ImageName = types.StringValue(str)
			}
		} else {
			if data.ImageName.IsUnknown() {
				data.ImageName = types.StringNull()
			}
		}
		if val, ok := sourceMap["instance_marketplace_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.InstanceMarketplaceUuid = types.StringValue(str)
			}
		} else {
			if data.InstanceMarketplaceUuid.IsUnknown() {
				data.InstanceMarketplaceUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["instance_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.InstanceName = types.StringValue(str)
			}
		} else {
			if data.InstanceName.IsUnknown() {
				data.InstanceName = types.StringNull()
			}
		}
		if val, ok := sourceMap["is_limit_based"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsLimitBased = types.BoolValue(b)
			}
		} else {
			if data.IsLimitBased.IsUnknown() {
				data.IsLimitBased = types.BoolNull()
			}
		}
		if val, ok := sourceMap["is_usage_based"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsUsageBased = types.BoolValue(b)
			}
		} else {
			if data.IsUsageBased.IsUnknown() {
				data.IsUsageBased = types.BoolNull()
			}
		}
		if val, ok := sourceMap["marketplace_category_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceCategoryName = types.StringValue(str)
			}
		} else {
			if data.MarketplaceCategoryName.IsUnknown() {
				data.MarketplaceCategoryName = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_category_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceCategoryUuid = types.StringValue(str)
			}
		} else {
			if data.MarketplaceCategoryUuid.IsUnknown() {
				data.MarketplaceCategoryUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_offering_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceOfferingName = types.StringValue(str)
			}
		} else {
			if data.MarketplaceOfferingName.IsUnknown() {
				data.MarketplaceOfferingName = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_offering_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceOfferingUuid = types.StringValue(str)
			}
		} else {
			if data.MarketplaceOfferingUuid.IsUnknown() {
				data.MarketplaceOfferingUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_plan_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplacePlanUuid = types.StringValue(str)
			}
		} else {
			if data.MarketplacePlanUuid.IsUnknown() {
				data.MarketplacePlanUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_resource_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceResourceState = types.StringValue(str)
			}
		} else {
			if data.MarketplaceResourceState.IsUnknown() {
				data.MarketplaceResourceState = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_resource_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceResourceUuid = types.StringValue(str)
			}
		} else {
			if data.MarketplaceResourceUuid.IsUnknown() {
				data.MarketplaceResourceUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["modified"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Modified = types.StringValue(str)
			}
		} else {
			if data.Modified.IsUnknown() {
				data.Modified = types.StringNull()
			}
		}
		if val, ok := sourceMap["resource_type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ResourceType = types.StringValue(str)
			}
		} else {
			if data.ResourceType.IsUnknown() {
				data.ResourceType = types.StringNull()
			}
		}
		if val, ok := sourceMap["service_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceName = types.StringValue(str)
			}
		} else {
			if data.ServiceName.IsUnknown() {
				data.ServiceName = types.StringNull()
			}
		}
		if val, ok := sourceMap["service_settings"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettings = types.StringValue(str)
			}
		} else {
			if data.ServiceSettings.IsUnknown() {
				data.ServiceSettings = types.StringNull()
			}
		}
		if val, ok := sourceMap["service_settings_error_message"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsErrorMessage = types.StringValue(str)
			}
		} else {
			if data.ServiceSettingsErrorMessage.IsUnknown() {
				data.ServiceSettingsErrorMessage = types.StringNull()
			}
		}
		if val, ok := sourceMap["service_settings_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsState = types.StringValue(str)
			}
		} else {
			if data.ServiceSettingsState.IsUnknown() {
				data.ServiceSettingsState = types.StringNull()
			}
		}
		if val, ok := sourceMap["size"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Size = types.Int64Value(int64(num))
			}
		} else {
			if data.Size.IsUnknown() {
				data.Size = types.Int64Null()
			}
		}
		if val, ok := sourceMap["source_snapshot"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.SourceSnapshot = types.StringValue(str)
			}
		} else {
			if data.SourceSnapshot.IsUnknown() {
				data.SourceSnapshot = types.StringNull()
			}
		}
		if val, ok := sourceMap["type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Type = types.StringValue(str)
			}
		} else {
			if data.Type.IsUnknown() {
				data.Type = types.StringNull()
			}
		}
		if val, ok := sourceMap["type_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.TypeName = types.StringValue(str)
			}
		} else {
			if data.TypeName.IsUnknown() {
				data.TypeName = types.StringNull()
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
		if val, ok := sourceMap["attach_instance_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AttachInstanceUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["availability_zone_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AvailabilityZoneName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["backend_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BackendId = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["can_manage"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanManage = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["customer"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Customer = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_abbreviation"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerAbbreviation = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_native_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerNativeName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Description = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["external_ip"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ExternalIp = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["instance"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Instance = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["instance_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.InstanceUuid = types.StringValue(str)
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
		if val, ok := sourceMap["project"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Project = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["runtime_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.RuntimeState = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_settings_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_settings_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["snapshot"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Snapshot = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["snapshot_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.SnapshotUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.State = types.StringValue(str)
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
		if val, ok := sourceMap["uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Uuid = types.StringValue(str)
			}
		}

	} else {
		// Filter by provided parameters
		var results []map[string]interface{}

		filters := map[string]string{}
		if !data.AttachInstanceUuid.IsNull() {
			filters["attach_instance_uuid"] = data.AttachInstanceUuid.ValueString()
		}
		if !data.AvailabilityZoneName.IsNull() {
			filters["availability_zone_name"] = data.AvailabilityZoneName.ValueString()
		}
		if !data.BackendId.IsNull() {
			filters["backend_id"] = data.BackendId.ValueString()
		}
		if !data.CanManage.IsNull() {
			filters["can_manage"] = fmt.Sprintf("%t", data.CanManage.ValueBool())
		}
		if !data.Customer.IsNull() {
			filters["customer"] = data.Customer.ValueString()
		}
		if !data.CustomerAbbreviation.IsNull() {
			filters["customer_abbreviation"] = data.CustomerAbbreviation.ValueString()
		}
		if !data.CustomerName.IsNull() {
			filters["customer_name"] = data.CustomerName.ValueString()
		}
		if !data.CustomerNativeName.IsNull() {
			filters["customer_native_name"] = data.CustomerNativeName.ValueString()
		}
		if !data.CustomerUuid.IsNull() {
			filters["customer_uuid"] = data.CustomerUuid.ValueString()
		}
		if !data.Description.IsNull() {
			filters["description"] = data.Description.ValueString()
		}
		if !data.ExternalIp.IsNull() {
			filters["external_ip"] = data.ExternalIp.ValueString()
		}
		if !data.Instance.IsNull() {
			filters["instance"] = data.Instance.ValueString()
		}
		if !data.InstanceUuid.IsNull() {
			filters["instance_uuid"] = data.InstanceUuid.ValueString()
		}
		if !data.Name.IsNull() {
			filters["name"] = data.Name.ValueString()
		}
		if !data.NameExact.IsNull() {
			filters["name_exact"] = data.NameExact.ValueString()
		}
		if !data.Project.IsNull() {
			filters["project"] = data.Project.ValueString()
		}
		if !data.ProjectName.IsNull() {
			filters["project_name"] = data.ProjectName.ValueString()
		}
		if !data.ProjectUuid.IsNull() {
			filters["project_uuid"] = data.ProjectUuid.ValueString()
		}
		if !data.RuntimeState.IsNull() {
			filters["runtime_state"] = data.RuntimeState.ValueString()
		}
		if !data.ServiceSettingsName.IsNull() {
			filters["service_settings_name"] = data.ServiceSettingsName.ValueString()
		}
		if !data.ServiceSettingsUuid.IsNull() {
			filters["service_settings_uuid"] = data.ServiceSettingsUuid.ValueString()
		}
		if !data.Snapshot.IsNull() {
			filters["snapshot"] = data.Snapshot.ValueString()
		}
		if !data.SnapshotUuid.IsNull() {
			filters["snapshot_uuid"] = data.SnapshotUuid.ValueString()
		}
		if !data.State.IsNull() {
			filters["state"] = data.State.ValueString()
		}
		if !data.Tenant.IsNull() {
			filters["tenant"] = data.Tenant.ValueString()
		}
		if !data.TenantUuid.IsNull() {
			filters["tenant_uuid"] = data.TenantUuid.ValueString()
		}
		if !data.Uuid.IsNull() {
			filters["uuid"] = data.Uuid.ValueString()
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
				"Unable to List Volume",
				"An error occurred while filtering openstack_volume: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Volume Not Found",
				"No openstack_volume found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Volumes Found",
				fmt.Sprintf("Found %d openstack_volumes with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
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
		if val, ok := sourceMap["access_url"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AccessUrl = types.StringValue(str)
			}
		} else {
			if data.AccessUrl.IsUnknown() {
				data.AccessUrl = types.StringNull()
			}
		}
		if val, ok := sourceMap["action"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Action = types.StringValue(str)
			}
		} else {
			if data.Action.IsUnknown() {
				data.Action = types.StringNull()
			}
		}
		if val, ok := sourceMap["availability_zone"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AvailabilityZone = types.StringValue(str)
			}
		} else {
			if data.AvailabilityZone.IsUnknown() {
				data.AvailabilityZone = types.StringNull()
			}
		}
		if val, ok := sourceMap["bootable"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.Bootable = types.BoolValue(b)
			}
		} else {
			if data.Bootable.IsUnknown() {
				data.Bootable = types.BoolNull()
			}
		}
		if val, ok := sourceMap["created"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Created = types.StringValue(str)
			}
		} else {
			if data.Created.IsUnknown() {
				data.Created = types.StringNull()
			}
		}
		if val, ok := sourceMap["device"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Device = types.StringValue(str)
			}
		} else {
			if data.Device.IsUnknown() {
				data.Device = types.StringNull()
			}
		}
		if val, ok := sourceMap["error_message"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ErrorMessage = types.StringValue(str)
			}
		} else {
			if data.ErrorMessage.IsUnknown() {
				data.ErrorMessage = types.StringNull()
			}
		}
		if val, ok := sourceMap["error_traceback"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ErrorTraceback = types.StringValue(str)
			}
		} else {
			if data.ErrorTraceback.IsUnknown() {
				data.ErrorTraceback = types.StringNull()
			}
		}
		if val, ok := sourceMap["extend_enabled"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.ExtendEnabled = types.BoolValue(b)
			}
		} else {
			if data.ExtendEnabled.IsUnknown() {
				data.ExtendEnabled = types.BoolNull()
			}
		}
		if val, ok := sourceMap["image"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Image = types.StringValue(str)
			}
		} else {
			if data.Image.IsUnknown() {
				data.Image = types.StringNull()
			}
		}
		if val, ok := sourceMap["image_metadata"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ImageMetadata = types.StringValue(str)
			}
		} else {
			if data.ImageMetadata.IsUnknown() {
				data.ImageMetadata = types.StringNull()
			}
		}
		if val, ok := sourceMap["image_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ImageName = types.StringValue(str)
			}
		} else {
			if data.ImageName.IsUnknown() {
				data.ImageName = types.StringNull()
			}
		}
		if val, ok := sourceMap["instance_marketplace_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.InstanceMarketplaceUuid = types.StringValue(str)
			}
		} else {
			if data.InstanceMarketplaceUuid.IsUnknown() {
				data.InstanceMarketplaceUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["instance_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.InstanceName = types.StringValue(str)
			}
		} else {
			if data.InstanceName.IsUnknown() {
				data.InstanceName = types.StringNull()
			}
		}
		if val, ok := sourceMap["is_limit_based"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsLimitBased = types.BoolValue(b)
			}
		} else {
			if data.IsLimitBased.IsUnknown() {
				data.IsLimitBased = types.BoolNull()
			}
		}
		if val, ok := sourceMap["is_usage_based"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.IsUsageBased = types.BoolValue(b)
			}
		} else {
			if data.IsUsageBased.IsUnknown() {
				data.IsUsageBased = types.BoolNull()
			}
		}
		if val, ok := sourceMap["marketplace_category_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceCategoryName = types.StringValue(str)
			}
		} else {
			if data.MarketplaceCategoryName.IsUnknown() {
				data.MarketplaceCategoryName = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_category_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceCategoryUuid = types.StringValue(str)
			}
		} else {
			if data.MarketplaceCategoryUuid.IsUnknown() {
				data.MarketplaceCategoryUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_offering_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceOfferingName = types.StringValue(str)
			}
		} else {
			if data.MarketplaceOfferingName.IsUnknown() {
				data.MarketplaceOfferingName = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_offering_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceOfferingUuid = types.StringValue(str)
			}
		} else {
			if data.MarketplaceOfferingUuid.IsUnknown() {
				data.MarketplaceOfferingUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_plan_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplacePlanUuid = types.StringValue(str)
			}
		} else {
			if data.MarketplacePlanUuid.IsUnknown() {
				data.MarketplacePlanUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_resource_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceResourceState = types.StringValue(str)
			}
		} else {
			if data.MarketplaceResourceState.IsUnknown() {
				data.MarketplaceResourceState = types.StringNull()
			}
		}
		if val, ok := sourceMap["marketplace_resource_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.MarketplaceResourceUuid = types.StringValue(str)
			}
		} else {
			if data.MarketplaceResourceUuid.IsUnknown() {
				data.MarketplaceResourceUuid = types.StringNull()
			}
		}
		if val, ok := sourceMap["modified"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Modified = types.StringValue(str)
			}
		} else {
			if data.Modified.IsUnknown() {
				data.Modified = types.StringNull()
			}
		}
		if val, ok := sourceMap["resource_type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ResourceType = types.StringValue(str)
			}
		} else {
			if data.ResourceType.IsUnknown() {
				data.ResourceType = types.StringNull()
			}
		}
		if val, ok := sourceMap["service_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceName = types.StringValue(str)
			}
		} else {
			if data.ServiceName.IsUnknown() {
				data.ServiceName = types.StringNull()
			}
		}
		if val, ok := sourceMap["service_settings"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettings = types.StringValue(str)
			}
		} else {
			if data.ServiceSettings.IsUnknown() {
				data.ServiceSettings = types.StringNull()
			}
		}
		if val, ok := sourceMap["service_settings_error_message"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsErrorMessage = types.StringValue(str)
			}
		} else {
			if data.ServiceSettingsErrorMessage.IsUnknown() {
				data.ServiceSettingsErrorMessage = types.StringNull()
			}
		}
		if val, ok := sourceMap["service_settings_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsState = types.StringValue(str)
			}
		} else {
			if data.ServiceSettingsState.IsUnknown() {
				data.ServiceSettingsState = types.StringNull()
			}
		}
		if val, ok := sourceMap["size"]; ok && val != nil {
			if num, ok := val.(float64); ok {
				data.Size = types.Int64Value(int64(num))
			}
		} else {
			if data.Size.IsUnknown() {
				data.Size = types.Int64Null()
			}
		}
		if val, ok := sourceMap["source_snapshot"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.SourceSnapshot = types.StringValue(str)
			}
		} else {
			if data.SourceSnapshot.IsUnknown() {
				data.SourceSnapshot = types.StringNull()
			}
		}
		if val, ok := sourceMap["type"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Type = types.StringValue(str)
			}
		} else {
			if data.Type.IsUnknown() {
				data.Type = types.StringNull()
			}
		}
		if val, ok := sourceMap["type_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.TypeName = types.StringValue(str)
			}
		} else {
			if data.TypeName.IsUnknown() {
				data.TypeName = types.StringNull()
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
		if val, ok := sourceMap["attach_instance_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AttachInstanceUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["availability_zone_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.AvailabilityZoneName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["backend_id"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.BackendId = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["can_manage"]; ok && val != nil {
			if b, ok := val.(bool); ok {
				data.CanManage = types.BoolValue(b)
			}
		}
		if val, ok := sourceMap["customer"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Customer = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_abbreviation"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerAbbreviation = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_native_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerNativeName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["customer_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.CustomerUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["description"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Description = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["external_ip"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ExternalIp = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["instance"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Instance = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["instance_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.InstanceUuid = types.StringValue(str)
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
		if val, ok := sourceMap["project"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Project = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["project_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ProjectUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["runtime_state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.RuntimeState = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_settings_name"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsName = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["service_settings_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.ServiceSettingsUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["snapshot"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Snapshot = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["snapshot_uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.SnapshotUuid = types.StringValue(str)
			}
		}
		if val, ok := sourceMap["state"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.State = types.StringValue(str)
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
		if val, ok := sourceMap["uuid"]; ok && val != nil {
			if str, ok := val.(string); ok {
				data.Uuid = types.StringValue(str)
			}
		}
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
