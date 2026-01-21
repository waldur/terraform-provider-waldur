package resources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
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

// OpenstackVolumeAttachmentApiResponse is the API response model.
type OpenstackVolumeAttachmentApiResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl               *string `json:"access_url" tfsdk:"access_url"`
	Action                  *string `json:"action" tfsdk:"action"`
	AvailabilityZone        *string `json:"availability_zone" tfsdk:"availability_zone"`
	AvailabilityZoneName    *string `json:"availability_zone_name" tfsdk:"availability_zone_name"`
	BackendId               *string `json:"backend_id" tfsdk:"backend_id"`
	Bootable                *bool   `json:"bootable" tfsdk:"bootable"`
	Created                 *string `json:"created" tfsdk:"created"`
	Description             *string `json:"description" tfsdk:"description"`
	Device                  *string `json:"device" tfsdk:"device"`
	ErrorMessage            *string `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback          *string `json:"error_traceback" tfsdk:"error_traceback"`
	ExtendEnabled           *bool   `json:"extend_enabled" tfsdk:"extend_enabled"`
	Image                   *string `json:"image" tfsdk:"image"`
	ImageMetadata           *string `json:"image_metadata" tfsdk:"image_metadata"`
	ImageName               *string `json:"image_name" tfsdk:"image_name"`
	Instance                *string `json:"instance" tfsdk:"instance"`
	InstanceMarketplaceUuid *string `json:"instance_marketplace_uuid" tfsdk:"instance_marketplace_uuid"`
	InstanceName            *string `json:"instance_name" tfsdk:"instance_name"`
	Modified                *string `json:"modified" tfsdk:"modified"`
	Name                    *string `json:"name" tfsdk:"name"`
	ResourceType            *string `json:"resource_type" tfsdk:"resource_type"`
	RuntimeState            *string `json:"runtime_state" tfsdk:"runtime_state"`
	Size                    *int64  `json:"size" tfsdk:"size"`
	SourceSnapshot          *string `json:"source_snapshot" tfsdk:"source_snapshot"`
	State                   *string `json:"state" tfsdk:"state"`
	Tenant                  *string `json:"tenant" tfsdk:"tenant"`
	TenantUuid              *string `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Type                    *string `json:"type" tfsdk:"type"`
	TypeName                *string `json:"type_name" tfsdk:"type_name"`
	Url                     *string `json:"url" tfsdk:"url"`
	Volume                  *string `json:"volume" tfsdk:"volume"`
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
		MarkdownDescription: "Openstack Volume Attachment resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"access_url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Access url",
			},
			"action": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Action",
			},
			"availability_zone": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Availability zone where this volume is located",
			},
			"availability_zone_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the availability zone",
			},
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Volume ID in the OpenStack backend",
			},
			"bootable": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Indicates if this volume can be used to boot an instance",
			},
			"created": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Created",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Description of the resource",
			},
			"device": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Link parameter",
			},
			"error_message": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Error message",
			},
			"error_traceback": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Error traceback",
			},
			"extend_enabled": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Extend enabled",
			},
			"image": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Image that this volume was created from, if any",
			},
			"image_metadata": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Metadata of the image this volume was created from",
			},
			"image_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the image this volume was created from",
			},
			"instance": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Instance that this volume is attached to, if any",
			},
			"instance_marketplace_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the instance marketplace",
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the instance",
			},
			"modified": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the resource",
			},
			"resource_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Resource type",
			},
			"runtime_state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Runtime state",
			},
			"size": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Size in MiB",
			},
			"source_snapshot": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Snapshot that this volume was created from, if any",
			},
			"state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "State",
			},
			"tenant": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Tenant",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the tenant",
			},
			"type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Type of the volume (e.g. SSD, HDD)",
			},
			"type_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the type",
			},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Url",
			},
			"volume": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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

	var apiResp OpenstackVolumeAttachmentApiResponse
	err := r.client.Post(ctx, linkPath, requestBody, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError("Link Operation Failed", err.Error())
		return
	}

	// For Link resources, ID is composite of Source and Target UUIDs
	data.UUID = types.StringValue(sourceUUID + "/" + data.Instance.ValueString())

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

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

	// For Link resources, we read the Source resource and check if Target is linked
	parts := strings.Split(data.UUID.ValueString(), "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid Link ID", "Expected source_uuid/target_uuid")
		return
	}
	sourceUUID := parts[0]
	targetUUID := parts[1]

	sourcePath := strings.Replace("/api/openstack-volumes/{uuid}/", "{uuid}", sourceUUID, 1)
	var result map[string]interface{}
	err := r.client.GetByUUID(ctx, sourcePath, sourceUUID, &result)
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}
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

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackVolumeAttachmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Link resources typically do not support update, as they are bindings.
	resp.Diagnostics.AddError("Update Not Supported", "Link resources cannot be updated.")
}

func (r *OpenstackVolumeAttachmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackVolumeAttachmentResourceModel
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

func (r *OpenstackVolumeAttachmentResource) mapResponseToModel(ctx context.Context, apiResp OpenstackVolumeAttachmentApiResponse, model *OpenstackVolumeAttachmentResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
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
	model.Volume = types.StringPointerValue(apiResp.Volume)

	return diags
}
