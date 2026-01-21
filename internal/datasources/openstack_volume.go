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
		MarkdownDescription: "Openstack Volume data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"attach_instance_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Filter for attachment to instance UUID",
			},
			"availability_zone_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Availability zone name",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Backend ID",
			},
			"can_manage": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Can manage",
			},
			"customer": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer UUID",
			},
			"customer_abbreviation": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer abbreviation",
			},
			"customer_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer name",
			},
			"customer_native_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer native name",
			},
			"customer_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer UUID",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Description",
			},
			"external_ip": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "External IP",
			},
			"instance": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Instance URL",
			},
			"instance_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Instance UUID",
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
			"project": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project UUID",
			},
			"project_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project name",
			},
			"project_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project UUID",
			},
			"runtime_state": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Runtime state",
			},
			"service_settings_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Service settings name",
			},
			"service_settings_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Service settings UUID",
			},
			"snapshot": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Snapshot URL",
			},
			"snapshot_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Snapshot UUID",
			},
			"state": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "State",
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
			"uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "UUID",
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
			"bootable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Indicates if this volume can be used to boot an instance",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Created",
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
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource type",
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
