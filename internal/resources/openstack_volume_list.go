package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-provider-waldur/internal/client"
)

var _ list.ListResource = &OpenstackVolumeList{}

type OpenstackVolumeList struct {
	client *client.Client
}

func NewOpenstackVolumeList() list.ListResource {
	return &OpenstackVolumeList{}
}

func (l *OpenstackVolumeList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_volume"
}

func (l *OpenstackVolumeList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		// Filter parameters can be added here if needed
		Attributes: map[string]schema.Attribute{},
	}
}

func (l *OpenstackVolumeList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type OpenstackVolumeListModel struct {
	// Add filter fields here if added to schema
}

func (l *OpenstackVolumeList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config OpenstackVolumeListModel

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
			var data OpenstackVolumeResourceModel

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
			if val, ok := sourceMap["customer"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Customer = types.StringValue(str)
				}
			} else {
				if data.Customer.IsUnknown() {
					data.Customer = types.StringNull()
				}
			}
			if val, ok := sourceMap["customer_abbreviation"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.CustomerAbbreviation = types.StringValue(str)
				}
			} else {
				if data.CustomerAbbreviation.IsUnknown() {
					data.CustomerAbbreviation = types.StringNull()
				}
			}
			if val, ok := sourceMap["customer_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.CustomerName = types.StringValue(str)
				}
			} else {
				if data.CustomerName.IsUnknown() {
					data.CustomerName = types.StringNull()
				}
			}
			if val, ok := sourceMap["customer_native_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.CustomerNativeName = types.StringValue(str)
				}
			} else {
				if data.CustomerNativeName.IsUnknown() {
					data.CustomerNativeName = types.StringNull()
				}
			}
			if val, ok := sourceMap["customer_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.CustomerUuid = types.StringValue(str)
				}
			} else {
				if data.CustomerUuid.IsUnknown() {
					data.CustomerUuid = types.StringNull()
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
			if val, ok := sourceMap["offering"]; ok && val != nil {
				if str, ok := val.(string); ok {
					// Normalize URL to UUID
					parts := strings.Split(strings.TrimRight(str, "/"), "/")
					uuid := parts[len(parts)-1]
					data.Offering = types.StringValue(uuid)
				} else {
					data.Offering = types.StringNull()
				}
			} else {
				if data.Offering.IsUnknown() {
					data.Offering = types.StringNull()
				}
			}
			if val, ok := sourceMap["project"]; ok && val != nil {
				if str, ok := val.(string); ok {
					// Normalize URL to UUID
					parts := strings.Split(strings.TrimRight(str, "/"), "/")
					uuid := parts[len(parts)-1]
					data.Project = types.StringValue(uuid)
				} else {
					data.Project = types.StringNull()
				}
			} else {
				if data.Project.IsUnknown() {
					data.Project = types.StringNull()
				}
			}
			if val, ok := sourceMap["project_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ProjectName = types.StringValue(str)
				}
			} else {
				if data.ProjectName.IsUnknown() {
					data.ProjectName = types.StringNull()
				}
			}
			if val, ok := sourceMap["project_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ProjectUuid = types.StringValue(str)
				}
			} else {
				if data.ProjectUuid.IsUnknown() {
					data.ProjectUuid = types.StringNull()
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
			if val, ok := sourceMap["service_settings_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ServiceSettingsUuid = types.StringValue(str)
				}
			} else {
				if data.ServiceSettingsUuid.IsUnknown() {
					data.ServiceSettingsUuid = types.StringNull()
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
