package volume_attachment

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

var _ list.ListResource = &OpenstackVolumeAttachmentList{}

type OpenstackVolumeAttachmentList struct {
	client *Client
}

func NewOpenstackVolumeAttachmentList() list.ListResource {
	return &OpenstackVolumeAttachmentList{}
}

func (l *OpenstackVolumeAttachmentList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_volume_attachment"
}

func (l *OpenstackVolumeAttachmentList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"attach_instance_uuid": schema.StringAttribute{
				Description: "Filter for attachment to instance UUID",
				Optional:    true,
			},
			"availability_zone_name": schema.StringAttribute{
				Description: "Availability zone name",
				Optional:    true,
			},
			"backend_id": schema.StringAttribute{
				Description: "Backend ID",
				Optional:    true,
			},
			"can_manage": schema.BoolAttribute{
				Description: "Can manage",
				Optional:    true,
			},
			"customer": schema.StringAttribute{
				Description: "Customer UUID",
				Optional:    true,
			},
			"customer_abbreviation": schema.StringAttribute{
				Description: "Customer abbreviation",
				Optional:    true,
			},
			"customer_name": schema.StringAttribute{
				Description: "Customer name",
				Optional:    true,
			},
			"customer_native_name": schema.StringAttribute{
				Description: "Customer native name",
				Optional:    true,
			},
			"customer_uuid": schema.StringAttribute{
				Description: "Customer UUID",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description",
				Optional:    true,
			},
			"external_ip": schema.StringAttribute{
				Description: "External IP",
				Optional:    true,
			},
			"instance": schema.StringAttribute{
				Description: "Instance URL",
				Optional:    true,
			},
			"instance_uuid": schema.StringAttribute{
				Description: "Instance UUID",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name",
				Optional:    true,
			},
			"name_exact": schema.StringAttribute{
				Description: "Name (exact)",
				Optional:    true,
			},
			"page": schema.Int64Attribute{
				Description: "A page number within the paginated result set.",
				Optional:    true,
			},
			"page_size": schema.Int64Attribute{
				Description: "Number of results to return per page.",
				Optional:    true,
			},
			"project": schema.StringAttribute{
				Description: "Project UUID",
				Optional:    true,
			},
			"project_name": schema.StringAttribute{
				Description: "Project name",
				Optional:    true,
			},
			"project_uuid": schema.StringAttribute{
				Description: "Project UUID",
				Optional:    true,
			},
			"runtime_state": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"service_settings_name": schema.StringAttribute{
				Description: "Service settings name",
				Optional:    true,
			},
			"service_settings_uuid": schema.StringAttribute{
				Description: "Service settings UUID",
				Optional:    true,
			},
			"snapshot": schema.StringAttribute{
				Description: "Snapshot URL",
				Optional:    true,
			},
			"snapshot_uuid": schema.StringAttribute{
				Description: "Snapshot UUID",
				Optional:    true,
			},
			"tenant": schema.StringAttribute{
				Description: "Tenant URL",
				Optional:    true,
			},
			"tenant_uuid": schema.StringAttribute{
				Description: "Tenant UUID",
				Optional:    true,
			},
			"uuid": schema.StringAttribute{
				Description: "UUID",
				Optional:    true,
			},
		},
	}
}

func (l *OpenstackVolumeAttachmentList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected List Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T", req.ProviderData),
		)
		return
	}

	l.client = NewClient(client)
}

type OpenstackVolumeAttachmentListModel struct {
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
	Page                 types.Int64  `tfsdk:"page"`
	PageSize             types.Int64  `tfsdk:"page_size"`
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

func (l *OpenstackVolumeAttachmentList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config OpenstackVolumeAttachmentListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Prepare filters
	filters := make(map[string]string)
	if !config.AttachInstanceUuid.IsNull() && !config.AttachInstanceUuid.IsUnknown() {
		filters["attach_instance_uuid"] = config.AttachInstanceUuid.ValueString()
	}
	if !config.AvailabilityZoneName.IsNull() && !config.AvailabilityZoneName.IsUnknown() {
		filters["availability_zone_name"] = config.AvailabilityZoneName.ValueString()
	}
	if !config.BackendId.IsNull() && !config.BackendId.IsUnknown() {
		filters["backend_id"] = config.BackendId.ValueString()
	}
	if !config.CanManage.IsNull() && !config.CanManage.IsUnknown() {
		filters["can_manage"] = fmt.Sprintf("%t", config.CanManage.ValueBool())
	}
	if !config.Customer.IsNull() && !config.Customer.IsUnknown() {
		filters["customer"] = config.Customer.ValueString()
	}
	if !config.CustomerAbbreviation.IsNull() && !config.CustomerAbbreviation.IsUnknown() {
		filters["customer_abbreviation"] = config.CustomerAbbreviation.ValueString()
	}
	if !config.CustomerName.IsNull() && !config.CustomerName.IsUnknown() {
		filters["customer_name"] = config.CustomerName.ValueString()
	}
	if !config.CustomerNativeName.IsNull() && !config.CustomerNativeName.IsUnknown() {
		filters["customer_native_name"] = config.CustomerNativeName.ValueString()
	}
	if !config.CustomerUuid.IsNull() && !config.CustomerUuid.IsUnknown() {
		filters["customer_uuid"] = config.CustomerUuid.ValueString()
	}
	if !config.Description.IsNull() && !config.Description.IsUnknown() {
		filters["description"] = config.Description.ValueString()
	}
	if !config.ExternalIp.IsNull() && !config.ExternalIp.IsUnknown() {
		filters["external_ip"] = config.ExternalIp.ValueString()
	}
	if !config.Instance.IsNull() && !config.Instance.IsUnknown() {
		filters["instance"] = config.Instance.ValueString()
	}
	if !config.InstanceUuid.IsNull() && !config.InstanceUuid.IsUnknown() {
		filters["instance_uuid"] = config.InstanceUuid.ValueString()
	}
	if !config.Name.IsNull() && !config.Name.IsUnknown() {
		filters["name"] = config.Name.ValueString()
	}
	if !config.NameExact.IsNull() && !config.NameExact.IsUnknown() {
		filters["name_exact"] = config.NameExact.ValueString()
	}
	if !config.Page.IsNull() && !config.Page.IsUnknown() {
		filters["page"] = fmt.Sprintf("%d", config.Page.ValueInt64())
	}
	if !config.PageSize.IsNull() && !config.PageSize.IsUnknown() {
		filters["page_size"] = fmt.Sprintf("%d", config.PageSize.ValueInt64())
	}
	if !config.Project.IsNull() && !config.Project.IsUnknown() {
		filters["project"] = config.Project.ValueString()
	}
	if !config.ProjectName.IsNull() && !config.ProjectName.IsUnknown() {
		filters["project_name"] = config.ProjectName.ValueString()
	}
	if !config.ProjectUuid.IsNull() && !config.ProjectUuid.IsUnknown() {
		filters["project_uuid"] = config.ProjectUuid.ValueString()
	}
	if !config.RuntimeState.IsNull() && !config.RuntimeState.IsUnknown() {
		filters["runtime_state"] = config.RuntimeState.ValueString()
	}
	if !config.ServiceSettingsName.IsNull() && !config.ServiceSettingsName.IsUnknown() {
		filters["service_settings_name"] = config.ServiceSettingsName.ValueString()
	}
	if !config.ServiceSettingsUuid.IsNull() && !config.ServiceSettingsUuid.IsUnknown() {
		filters["service_settings_uuid"] = config.ServiceSettingsUuid.ValueString()
	}
	if !config.Snapshot.IsNull() && !config.Snapshot.IsUnknown() {
		filters["snapshot"] = config.Snapshot.ValueString()
	}
	if !config.SnapshotUuid.IsNull() && !config.SnapshotUuid.IsUnknown() {
		filters["snapshot_uuid"] = config.SnapshotUuid.ValueString()
	}
	if !config.Tenant.IsNull() && !config.Tenant.IsUnknown() {
		filters["tenant"] = config.Tenant.ValueString()
	}
	if !config.TenantUuid.IsNull() && !config.TenantUuid.IsUnknown() {
		filters["tenant_uuid"] = config.TenantUuid.ValueString()
	}
	if !config.Uuid.IsNull() && !config.Uuid.IsUnknown() {
		filters["uuid"] = config.Uuid.ValueString()
	}

	// Call API
	listResult, err := l.client.ListOpenstackVolumeAttachment(ctx, filters)
	if err != nil {
		// Return error diagnostics
		resp.AddError("Failed to list resources", err.Error())
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Stream results
	stream.Results = func(push func(list.ListResult) bool) {
		for _, apiResp := range listResult {
			result := req.NewListResult(ctx)

			// Map item to model
			var data OpenstackVolumeAttachmentResourceModel
			model := &data

			var diags diag.Diagnostics

			data.UUID = types.StringPointerValue(apiResp.UUID)
			model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
			model.Action = types.StringPointerValue(apiResp.Action)
			model.AvailabilityZone = types.StringPointerValue(apiResp.AvailabilityZone)
			model.AvailabilityZoneName = types.StringPointerValue(apiResp.AvailabilityZoneName)
			model.BackendId = types.StringPointerValue(apiResp.BackendId)
			model.Bootable = types.BoolPointerValue(apiResp.Bootable)
			model.Created = types.StringPointerValue(apiResp.Created)
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
			model.Modified = types.StringPointerValue(apiResp.Modified)
			model.Name = types.StringPointerValue(apiResp.Name)
			model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
			model.RuntimeState = types.StringPointerValue(apiResp.RuntimeState)
			model.Size = types.Int64PointerValue(apiResp.Size)
			model.SourceSnapshot = types.StringPointerValue(apiResp.SourceSnapshot)
			model.State = types.StringPointerValue(apiResp.State)
			model.Tenant = types.StringPointerValue(apiResp.Tenant)
			model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
			model.Type = types.StringPointerValue(apiResp.Type)
			model.TypeName = types.StringPointerValue(apiResp.TypeName)
			model.Url = types.StringPointerValue(apiResp.Url)

			// Set the resource state
			// For ListResource, we generally return the "Resource" state matching the main resource schema.
			// This allows `terraform plan` to correlate if using `terraform query`.

			setDiags := result.Resource.Set(ctx, &data)
			diags.Append(setDiags...)
			result.Diagnostics.Append(diags...)

			// Set Identity if possible (usually UUID)
			if !data.UUID.IsNull() && !data.UUID.IsUnknown() {
				result.Diagnostics.Append(result.Identity.Set(ctx, data.UUID.ValueString())...)
			}

			if !push(result) {
				return
			}
		}
	}
}
