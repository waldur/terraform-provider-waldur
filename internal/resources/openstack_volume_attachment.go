package resources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-waldur-provider/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackVolumeAttachmentResource{}
var _ resource.ResourceWithImportState = &OpenstackVolumeAttachmentResource{}

func NewOpenstackVolumeAttachmentResource() resource.Resource {
	return &OpenstackVolumeAttachmentResource{}
}

// OpenstackVolumeAttachmentResource defines the resource implementation.
type OpenstackVolumeAttachmentResource struct {
	client *client.Client
}

// OpenstackVolumeAttachmentResourceModel describes the resource data model.
type OpenstackVolumeAttachmentResourceModel struct {
	UUID                    types.String   `tfsdk:"id"`
	AccessUrl               types.String   `tfsdk:"access_url"`
	Action                  types.String   `tfsdk:"action"`
	AvailabilityZone        types.String   `tfsdk:"availability_zone"`
	AvailabilityZoneName    types.String   `tfsdk:"availability_zone_name"`
	BackendId               types.String   `tfsdk:"backend_id"`
	Bootable                types.Bool     `tfsdk:"bootable"`
	Created                 types.String   `tfsdk:"created"`
	Description             types.String   `tfsdk:"description"`
	Device                  types.String   `tfsdk:"device"`
	ErrorMessage            types.String   `tfsdk:"error_message"`
	ErrorTraceback          types.String   `tfsdk:"error_traceback"`
	ExtendEnabled           types.Bool     `tfsdk:"extend_enabled"`
	Image                   types.String   `tfsdk:"image"`
	ImageMetadata           types.String   `tfsdk:"image_metadata"`
	ImageName               types.String   `tfsdk:"image_name"`
	Instance                types.String   `tfsdk:"instance"`
	InstanceMarketplaceUuid types.String   `tfsdk:"instance_marketplace_uuid"`
	InstanceName            types.String   `tfsdk:"instance_name"`
	Modified                types.String   `tfsdk:"modified"`
	Name                    types.String   `tfsdk:"name"`
	ResourceType            types.String   `tfsdk:"resource_type"`
	RuntimeState            types.String   `tfsdk:"runtime_state"`
	Size                    types.Int64    `tfsdk:"size"`
	SourceSnapshot          types.String   `tfsdk:"source_snapshot"`
	State                   types.String   `tfsdk:"state"`
	Tenant                  types.String   `tfsdk:"tenant"`
	TenantUuid              types.String   `tfsdk:"tenant_uuid"`
	Type                    types.String   `tfsdk:"type"`
	TypeName                types.String   `tfsdk:"type_name"`
	Url                     types.String   `tfsdk:"url"`
	Volume                  types.String   `tfsdk:"volume"`
	Timeouts                timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackVolumeAttachmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_volume_attachment"
}

func (r *OpenstackVolumeAttachmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "OpenstackVolumeAttachment resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Availability zone where this volume is located",
			},
			"availability_zone_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Volume ID in the OpenStack backend",
			},
			"bootable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Indicates if this volume can be used to boot an instance",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"device": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Link parameter",
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
				Optional:            true,
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
				MarkdownDescription: "",
			},
			"instance_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"runtime_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"size": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Size in MiB",
			},
			"source_snapshot": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Snapshot that this volume was created from, if any",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"tenant": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"type": schema.StringAttribute{
				Optional:            true,
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
			"volume": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Source resource UUID",
			},
		},

		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx, timeouts.Opts{
				Create: true,
				Update: true,
				Delete: true,
			}),
		},
	}
}

func (r *OpenstackVolumeAttachmentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	r.client = client
}

func (r *OpenstackVolumeAttachmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackVolumeAttachmentResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Link Plugin Create Logic
	sourceUUID := data.Volume.ValueString()
	linkPath := strings.Replace("/api/openstack-volumes/{uuid}/attach/", "{uuid}", sourceUUID, 1)

	requestBody := map[string]interface{}{
		"instance": data.Instance.ValueString(),
		"device":   data.Device.ValueString(),
	}

	var result map[string]interface{}
	err := r.client.Post(ctx, linkPath, requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError("Link Operation Failed", err.Error())
		return
	}

	// For Link resources, ID is composite of Source and Target UUIDs
	data.UUID = types.StringValue(sourceUUID + "/" + data.Instance.ValueString())

	// For Link resources, sourceMap is the Link result (often empty) or we might want to fetch Source?
	// But usually Link op returns nothing or small info.
	// MapResponseFields usually expects fields of the resource.
	// We'll set sourceMap to result for now.
	sourceMap := result
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
	if val, ok := sourceMap["availability_zone_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AvailabilityZoneName = types.StringValue(str)
		}
	} else {
		if data.AvailabilityZoneName.IsUnknown() {
			data.AvailabilityZoneName = types.StringNull()
		}
	}
	if val, ok := sourceMap["backend_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
		}
	} else {
		if data.BackendId.IsUnknown() {
			data.BackendId = types.StringNull()
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
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
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
	if val, ok := sourceMap["instance"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Instance = types.StringValue(str)
		}
	} else {
		if data.Instance.IsUnknown() {
			data.Instance = types.StringNull()
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
	if val, ok := sourceMap["runtime_state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.RuntimeState = types.StringValue(str)
		}
	} else {
		if data.RuntimeState.IsUnknown() {
			data.RuntimeState = types.StringNull()
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
	if val, ok := sourceMap["state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.State = types.StringValue(str)
		}
	} else {
		if data.State.IsUnknown() {
			data.State = types.StringNull()
		}
	}
	if val, ok := sourceMap["tenant"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Tenant = types.StringValue(str)
		}
	} else {
		if data.Tenant.IsUnknown() {
			data.Tenant = types.StringNull()
		}
	}
	if val, ok := sourceMap["tenant_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.TenantUuid = types.StringValue(str)
		}
	} else {
		if data.TenantUuid.IsUnknown() {
			data.TenantUuid = types.StringNull()
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
	if val, ok := sourceMap["volume"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Volume = types.StringValue(str)
		}
	} else {
		if data.Volume.IsUnknown() {
			data.Volume = types.StringNull()
		}
	}

	// Map filter parameters from response if available

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackVolumeAttachmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackVolumeAttachmentResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource
	var result map[string]interface{}
	// For Link resources, we read the Source resource and check if Target is linked
	parts := strings.Split(data.UUID.ValueString(), "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid Link ID", "Expected source_uuid/target_uuid")
		return
	}
	sourceUUID := parts[0]
	targetUUID := parts[1]

	sourcePath := strings.Replace("/api/openstack-volumes/{uuid}/", "{uuid}", sourceUUID, 1)
	err := r.client.GetByUUID(ctx, sourcePath, sourceUUID, &result)
	if err != nil {
		resp.Diagnostics.AddError("Unable to Read Source Resource", err.Error())
		return
	}

	// Check if target is present in the source resource
	// The LinkCheckKey usually points to a list or a single object/string
	if val, ok := result["instance"]; ok {
		found := false
		if strVal, ok := val.(string); ok {
			// Single value link (e.g. 1-to-1)
			// Check if it matches targetUUID or is a URL containing targetUUID
			if strings.Contains(strVal, targetUUID) {
				found = true
			}
		} else if listVal, ok := val.([]interface{}); ok {
			// List of links (e.g. 1-to-many)
			for _, item := range listVal {
				if itemMap, ok := item.(map[string]interface{}); ok {
					if uuid, ok := itemMap["uuid"].(string); ok && uuid == targetUUID {
						found = true
						break
					}
					if url, ok := itemMap["url"].(string); ok && strings.Contains(url, targetUUID) {
						found = true
						break
					}
				}
			}
		}

		if !found {
			// Link not found, remove from state
			resp.State.RemoveResource(ctx)
			return
		}
	} else {
		// Key not present
		resp.State.RemoveResource(ctx)
		return
	}

	// Should probably create a mock result map containing source and target UUIDs
	// to satisfy mapResponseFields if needed, but for now we just keep result which is Source
	// We need to ensure we don't overwrite data.UUID with Source UUID
	// So we delete "uuid" from result before mapping
	delete(result, "uuid")

	// Update UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
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
	if val, ok := sourceMap["availability_zone_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AvailabilityZoneName = types.StringValue(str)
		}
	} else {
		if data.AvailabilityZoneName.IsUnknown() {
			data.AvailabilityZoneName = types.StringNull()
		}
	}
	if val, ok := sourceMap["backend_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
		}
	} else {
		if data.BackendId.IsUnknown() {
			data.BackendId = types.StringNull()
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
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
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
	if val, ok := sourceMap["instance"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Instance = types.StringValue(str)
		}
	} else {
		if data.Instance.IsUnknown() {
			data.Instance = types.StringNull()
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
	if val, ok := sourceMap["runtime_state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.RuntimeState = types.StringValue(str)
		}
	} else {
		if data.RuntimeState.IsUnknown() {
			data.RuntimeState = types.StringNull()
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
	if val, ok := sourceMap["state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.State = types.StringValue(str)
		}
	} else {
		if data.State.IsUnknown() {
			data.State = types.StringNull()
		}
	}
	if val, ok := sourceMap["tenant"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Tenant = types.StringValue(str)
		}
	} else {
		if data.Tenant.IsUnknown() {
			data.Tenant = types.StringNull()
		}
	}
	if val, ok := sourceMap["tenant_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.TenantUuid = types.StringValue(str)
		}
	} else {
		if data.TenantUuid.IsUnknown() {
			data.TenantUuid = types.StringNull()
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
	if val, ok := sourceMap["volume"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Volume = types.StringValue(str)
		}
	} else {
		if data.Volume.IsUnknown() {
			data.Volume = types.StringNull()
		}
	}

	// Map filter parameters from response if available

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackVolumeAttachmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackVolumeAttachmentResourceModel
	var state OpenstackVolumeAttachmentResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	// Read current state to get the UUID (which is computed and not in plan)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Use UUID from state
	data.UUID = state.UUID
	// Prepare request body
	requestBody := map[string]interface{}{}
	if !data.Bootable.IsNull() && !data.Bootable.IsUnknown() {
		requestBody["bootable"] = data.Bootable.ValueBool()
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		if v := data.Description.ValueString(); v != "" {
			requestBody["description"] = v
		}
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		if v := data.Name.ValueString(); v != "" {
			requestBody["name"] = v
		}
	}

	// Call Waldur API to update resource
	var result map[string]interface{}

	err := r.client.Update(ctx, "/api/openstack-volumes/{uuid}/", data.UUID.ValueString(), requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update OpenstackVolumeAttachment",
			"An error occurred while updating the openstack_volume_attachment: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
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
	if val, ok := sourceMap["availability_zone_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AvailabilityZoneName = types.StringValue(str)
		}
	} else {
		if data.AvailabilityZoneName.IsUnknown() {
			data.AvailabilityZoneName = types.StringNull()
		}
	}
	if val, ok := sourceMap["backend_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
		}
	} else {
		if data.BackendId.IsUnknown() {
			data.BackendId = types.StringNull()
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
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
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
	if val, ok := sourceMap["instance"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Instance = types.StringValue(str)
		}
	} else {
		if data.Instance.IsUnknown() {
			data.Instance = types.StringNull()
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
	if val, ok := sourceMap["runtime_state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.RuntimeState = types.StringValue(str)
		}
	} else {
		if data.RuntimeState.IsUnknown() {
			data.RuntimeState = types.StringNull()
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
	if val, ok := sourceMap["state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.State = types.StringValue(str)
		}
	} else {
		if data.State.IsUnknown() {
			data.State = types.StringNull()
		}
	}
	if val, ok := sourceMap["tenant"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Tenant = types.StringValue(str)
		}
	} else {
		if data.Tenant.IsUnknown() {
			data.Tenant = types.StringNull()
		}
	}
	if val, ok := sourceMap["tenant_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.TenantUuid = types.StringValue(str)
		}
	} else {
		if data.TenantUuid.IsUnknown() {
			data.TenantUuid = types.StringNull()
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
	if val, ok := sourceMap["volume"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Volume = types.StringValue(str)
		}
	} else {
		if data.Volume.IsUnknown() {
			data.Volume = types.StringNull()
		}
	}

	// Map filter parameters from response if available

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackVolumeAttachmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackVolumeAttachmentResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Link Plugin Delete (Unlink)
	parts := strings.Split(data.UUID.ValueString(), "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid Link ID", "Expected source_uuid/target_uuid")
		return
	}
	sourceUUID := parts[0]
	unlinkPath := strings.Replace("/api/openstack-volumes/{uuid}/detach/", "{uuid}", sourceUUID, 1)

	err := r.client.Post(ctx, unlinkPath, nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Unlink Failed", err.Error())
		return
	}
}

func (r *OpenstackVolumeAttachmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import ID: source_uuid/target_uuid
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid Import ID", "Expected format: <volume_uuid>/<instance_uuid>")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("volume"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("instance"), parts[1])...)
}
