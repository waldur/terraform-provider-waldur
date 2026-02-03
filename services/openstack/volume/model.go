package volume

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

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
	Tenant               types.String `tfsdk:"tenant"`
	TenantUuid           types.String `tfsdk:"tenant_uuid"`
	Uuid                 types.String `tfsdk:"uuid"`
}

func (m *OpenstackVolumeFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
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
	}
}

type OpenstackVolumeModel struct {
	UUID                        types.String      `tfsdk:"id"`
	AccessUrl                   types.String      `tfsdk:"access_url"`
	Action                      types.String      `tfsdk:"action"`
	AvailabilityZone            types.String      `tfsdk:"availability_zone"`
	AvailabilityZoneName        types.String      `tfsdk:"availability_zone_name"`
	BackendId                   types.String      `tfsdk:"backend_id"`
	Bootable                    types.Bool        `tfsdk:"bootable"`
	Created                     timetypes.RFC3339 `tfsdk:"created"`
	Customer                    types.String      `tfsdk:"customer"`
	CustomerAbbreviation        types.String      `tfsdk:"customer_abbreviation"`
	CustomerName                types.String      `tfsdk:"customer_name"`
	CustomerNativeName          types.String      `tfsdk:"customer_native_name"`
	CustomerUuid                types.String      `tfsdk:"customer_uuid"`
	Description                 types.String      `tfsdk:"description"`
	Device                      types.String      `tfsdk:"device"`
	ErrorMessage                types.String      `tfsdk:"error_message"`
	ErrorTraceback              types.String      `tfsdk:"error_traceback"`
	ExtendEnabled               types.Bool        `tfsdk:"extend_enabled"`
	Image                       types.String      `tfsdk:"image"`
	ImageMetadata               types.String      `tfsdk:"image_metadata"`
	ImageName                   types.String      `tfsdk:"image_name"`
	Instance                    types.String      `tfsdk:"instance"`
	InstanceMarketplaceUuid     types.String      `tfsdk:"instance_marketplace_uuid"`
	InstanceName                types.String      `tfsdk:"instance_name"`
	IsLimitBased                types.Bool        `tfsdk:"is_limit_based"`
	IsUsageBased                types.Bool        `tfsdk:"is_usage_based"`
	Limits                      types.Map         `tfsdk:"limits"`
	MarketplaceCategoryName     types.String      `tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     types.String      `tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     types.String      `tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     types.String      `tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         types.String      `tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    types.String      `tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     types.String      `tfsdk:"marketplace_resource_uuid"`
	Modified                    timetypes.RFC3339 `tfsdk:"modified"`
	Name                        types.String      `tfsdk:"name"`
	Offering                    types.String      `tfsdk:"offering"`
	Plan                        types.String      `tfsdk:"plan"`
	Project                     types.String      `tfsdk:"project"`
	ProjectName                 types.String      `tfsdk:"project_name"`
	ProjectUuid                 types.String      `tfsdk:"project_uuid"`
	ResourceType                types.String      `tfsdk:"resource_type"`
	RuntimeState                types.String      `tfsdk:"runtime_state"`
	ServiceName                 types.String      `tfsdk:"service_name"`
	ServiceSettings             types.String      `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage types.String      `tfsdk:"service_settings_error_message"`
	ServiceSettingsState        types.String      `tfsdk:"service_settings_state"`
	ServiceSettingsUuid         types.String      `tfsdk:"service_settings_uuid"`
	Size                        types.Int64       `tfsdk:"size"`
	SourceSnapshot              types.String      `tfsdk:"source_snapshot"`
	State                       types.String      `tfsdk:"state"`
	Tenant                      types.String      `tfsdk:"tenant"`
	TenantUuid                  types.String      `tfsdk:"tenant_uuid"`
	Type                        types.String      `tfsdk:"type"`
	TypeName                    types.String      `tfsdk:"type_name"`
	Url                         types.String      `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *OpenstackVolumeModel) CopyFrom(ctx context.Context, apiResp OpenstackVolumeResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = common.StringPointerValue(apiResp.AccessUrl)
	model.Action = common.StringPointerValue(apiResp.Action)
	model.AvailabilityZone = common.StringPointerValue(apiResp.AvailabilityZone)
	model.AvailabilityZoneName = common.StringPointerValue(apiResp.AvailabilityZoneName)
	model.BackendId = common.StringPointerValue(apiResp.BackendId)
	model.Bootable = types.BoolPointerValue(apiResp.Bootable)
	valCreated, diagsCreated := timetypes.NewRFC3339PointerValue(apiResp.Created)
	diags.Append(diagsCreated...)
	model.Created = valCreated
	model.Customer = common.StringPointerValue(apiResp.Customer)
	model.CustomerAbbreviation = common.StringPointerValue(apiResp.CustomerAbbreviation)
	model.CustomerName = common.StringPointerValue(apiResp.CustomerName)
	model.CustomerNativeName = common.StringPointerValue(apiResp.CustomerNativeName)
	model.CustomerUuid = common.StringPointerValue(apiResp.CustomerUuid)
	model.Description = common.StringPointerValue(apiResp.Description)
	model.Device = common.StringPointerValue(apiResp.Device)
	model.ErrorMessage = common.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = common.StringPointerValue(apiResp.ErrorTraceback)
	model.ExtendEnabled = types.BoolPointerValue(apiResp.ExtendEnabled)
	model.Image = common.StringPointerValue(apiResp.Image)
	model.ImageMetadata = common.StringPointerValue(apiResp.ImageMetadata)
	model.ImageName = common.StringPointerValue(apiResp.ImageName)
	model.Instance = common.StringPointerValue(apiResp.Instance)
	model.InstanceMarketplaceUuid = common.StringPointerValue(apiResp.InstanceMarketplaceUuid)
	model.InstanceName = common.StringPointerValue(apiResp.InstanceName)
	model.IsLimitBased = types.BoolPointerValue(apiResp.IsLimitBased)
	model.IsUsageBased = types.BoolPointerValue(apiResp.IsUsageBased)
	model.MarketplaceCategoryName = common.StringPointerValue(apiResp.MarketplaceCategoryName)
	model.MarketplaceCategoryUuid = common.StringPointerValue(apiResp.MarketplaceCategoryUuid)
	model.MarketplaceOfferingName = common.StringPointerValue(apiResp.MarketplaceOfferingName)
	model.MarketplaceOfferingUuid = common.StringPointerValue(apiResp.MarketplaceOfferingUuid)
	model.MarketplacePlanUuid = common.StringPointerValue(apiResp.MarketplacePlanUuid)
	model.MarketplaceResourceState = common.StringPointerValue(apiResp.MarketplaceResourceState)
	model.MarketplaceResourceUuid = common.StringPointerValue(apiResp.MarketplaceResourceUuid)
	valModified, diagsModified := timetypes.NewRFC3339PointerValue(apiResp.Modified)
	diags.Append(diagsModified...)
	model.Modified = valModified
	model.Name = common.StringPointerValue(apiResp.Name)
	model.Project = common.StringPointerValue(apiResp.Project)
	model.ProjectName = common.StringPointerValue(apiResp.ProjectName)
	model.ProjectUuid = common.StringPointerValue(apiResp.ProjectUuid)
	model.ResourceType = common.StringPointerValue(apiResp.ResourceType)
	model.RuntimeState = common.StringPointerValue(apiResp.RuntimeState)
	model.ServiceName = common.StringPointerValue(apiResp.ServiceName)
	model.ServiceSettings = common.StringPointerValue(apiResp.ServiceSettings)
	model.ServiceSettingsErrorMessage = common.StringPointerValue(apiResp.ServiceSettingsErrorMessage)
	model.ServiceSettingsState = common.StringPointerValue(apiResp.ServiceSettingsState)
	model.ServiceSettingsUuid = common.StringPointerValue(apiResp.ServiceSettingsUuid)
	model.Size = types.Int64PointerValue(apiResp.Size)
	model.SourceSnapshot = common.StringPointerValue(apiResp.SourceSnapshot)
	model.State = common.StringPointerValue(apiResp.State)
	model.Tenant = common.StringPointerValue(apiResp.Tenant)
	model.TenantUuid = common.StringPointerValue(apiResp.TenantUuid)
	model.Type = common.StringPointerValue(apiResp.Type)
	model.TypeName = common.StringPointerValue(apiResp.TypeName)
	model.Url = common.StringPointerValue(apiResp.Url)

	return diags
}
