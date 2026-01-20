package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-provider-waldur/internal/client"
)

var _ list.ListResource = &OpenstackVolumeAttachmentList{}

type OpenstackVolumeAttachmentList struct {
	client *client.Client
}

func NewOpenstackVolumeAttachmentList() list.ListResource {
	return &OpenstackVolumeAttachmentList{}
}

func (l *OpenstackVolumeAttachmentList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_volume_attachment"
}

func (l *OpenstackVolumeAttachmentList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		// Filter parameters can be added here if needed
		Attributes: map[string]schema.Attribute{},
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

	l.client = client
}

type OpenstackVolumeAttachmentListModel struct {
	// Add filter fields here if added to schema
}

func (l *OpenstackVolumeAttachmentList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config OpenstackVolumeAttachmentListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Call API
	var listResult []map[string]interface{}
	err := l.client.List(ctx, "/api/openstack-volumes/", &listResult)
	if err != nil {
		// Return error diagnostics
		resp.AddError("Failed to list resources", err.Error())
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Stream results
	stream.Results = func(push func(list.ListResult) bool) {
		for _, item := range listResult {
			result := req.NewListResult(ctx)

			// Map item to model
			var data OpenstackVolumeAttachmentResourceModel

			// Initialize lists and objects to empty/null if needed (or rely on map logic)

			sourceMap := item

			// Reuse the mapResponseFields logic (embedded here or via shared template)
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

			// Map filter parameters from response if available

			// Set the resource state
			// For ListResource, we generally return the "Resource" state matching the main resource schema.
			// This allows `terraform plan` to correlate if using `terraform query`.

			diags := result.Resource.Set(ctx, &data)
			result.Diagnostics.Append(diags...)

			// Set Identity if possible (usually UUID)
			if !data.UUID.IsNull() && !data.UUID.IsUnknown() {
				// Identity value must match what the resource uses for Import?
				// Typically implicit. For now just setting Resource is key.
				// result.Identity.Set(ctx, data.UUID.ValueString())
				// The doc says: "An error is returned if a list result in the stream contains a null identity"
				result.Diagnostics.Append(result.Identity.Set(ctx, data.UUID.ValueString())...)
			} else {
				// Try to fallback to "uuid" from map if model failed
				if uuid, ok := item["uuid"].(string); ok {
					result.Diagnostics.Append(result.Identity.Set(ctx, uuid)...)
				}
			}

			if !push(result) {
				return
			}
		}
	}
}
