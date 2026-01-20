package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-provider-waldur/internal/client"
)

var _ list.ListResource = &MarketplaceResourceList{}

type MarketplaceResourceList struct {
	client *client.Client
}

func NewMarketplaceResourceList() list.ListResource {
	return &MarketplaceResourceList{}
}

func (l *MarketplaceResourceList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_marketplace_resource"
}

func (l *MarketplaceResourceList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		// Filter parameters can be added here if needed
		Attributes: map[string]schema.Attribute{},
	}
}

func (l *MarketplaceResourceList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type MarketplaceResourceListModel struct {
	// Add filter fields here if added to schema
}

func (l *MarketplaceResourceList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config MarketplaceResourceListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Call API
	var listResult []map[string]interface{}
	err := l.client.List(ctx, "/api/marketplace-resources/", &listResult)
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
			var data MarketplaceResourceResourceModel

			// Initialize lists and objects to empty/null if needed (or rely on map logic)

			sourceMap := item

			// Reuse the mapResponseFields logic (embedded here or via shared template)
			// Map response fields to data model
			_ = sourceMap
			if val, ok := sourceMap["available_actions"]; ok && val != nil {
				// List of strings (or flattened objects)
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if str, ok := item.(string); ok {
							items = append(items, types.StringValue(str))
						} else if obj, ok := item.(map[string]interface{}); ok {
							// Flattening logic: extract URL or UUID
							if url, ok := obj["url"].(string); ok {
								parts := strings.Split(strings.TrimRight(url, "/"), "/")
								uuid := parts[len(parts)-1]
								items = append(items, types.StringValue(uuid))
							} else if uuid, ok := obj["uuid"].(string); ok {
								items = append(items, types.StringValue(uuid))
							} else if name, ok := obj["name"].(string); ok {
								items = append(items, types.StringValue(name))
							}
						}
					}
					listVal, _ := types.ListValue(types.StringType, items)
					data.AvailableActions = listVal
				}
			} else {
				if data.AvailableActions.IsUnknown() {
					data.AvailableActions = types.ListNull(types.StringType)
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
			if val, ok := sourceMap["can_terminate"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.CanTerminate = types.BoolValue(b)
				}
			} else {
				if data.CanTerminate.IsUnknown() {
					data.CanTerminate = types.BoolNull()
				}
			}
			if val, ok := sourceMap["category_icon"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.CategoryIcon = types.StringValue(str)
				}
			} else {
				if data.CategoryIcon.IsUnknown() {
					data.CategoryIcon = types.StringNull()
				}
			}
			if val, ok := sourceMap["category_title"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.CategoryTitle = types.StringValue(str)
				}
			} else {
				if data.CategoryTitle.IsUnknown() {
					data.CategoryTitle = types.StringNull()
				}
			}
			if val, ok := sourceMap["category_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.CategoryUuid = types.StringValue(str)
				}
			} else {
				if data.CategoryUuid.IsUnknown() {
					data.CategoryUuid = types.StringNull()
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
			if val, ok := sourceMap["customer_slug"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.CustomerSlug = types.StringValue(str)
				}
			} else {
				if data.CustomerSlug.IsUnknown() {
					data.CustomerSlug = types.StringNull()
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
			if val, ok := sourceMap["downscaled"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.Downscaled = types.BoolValue(b)
				}
			} else {
				if data.Downscaled.IsUnknown() {
					data.Downscaled = types.BoolNull()
				}
			}
			if val, ok := sourceMap["effective_id"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.EffectiveId = types.StringValue(str)
				}
			} else {
				if data.EffectiveId.IsUnknown() {
					data.EffectiveId = types.StringNull()
				}
			}
			if val, ok := sourceMap["end_date"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.EndDate = types.StringValue(str)
				}
			} else {
				if data.EndDate.IsUnknown() {
					data.EndDate = types.StringNull()
				}
			}
			if val, ok := sourceMap["end_date_requested_by"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.EndDateRequestedBy = types.StringValue(str)
				}
			} else {
				if data.EndDateRequestedBy.IsUnknown() {
					data.EndDateRequestedBy = types.StringNull()
				}
			}
			if val, ok := sourceMap["endpoints"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"name": types.StringType,
								"url":  types.StringType,
							}
							attrValues := map[string]attr.Value{
								"name": func() attr.Value {
									if v, ok := objMap["name"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"url": func() attr.Value {
									if v, ok := objMap["url"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
							}
							objVal, _ := types.ObjectValue(attrTypes, attrValues)
							items = append(items, objVal)
						}
					}
					listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
						"name": types.StringType,
						"url":  types.StringType,
					}}, items)
					data.Endpoints = listVal
				}
			} else {
				if data.Endpoints.IsUnknown() {
					data.Endpoints = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"name": types.StringType,
						"url":  types.StringType,
					}})
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
			if val, ok := sourceMap["last_sync"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.LastSync = types.StringValue(str)
				}
			} else {
				if data.LastSync.IsUnknown() {
					data.LastSync = types.StringNull()
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
			if val, ok := sourceMap["offering_billable"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.OfferingBillable = types.BoolValue(b)
				}
			} else {
				if data.OfferingBillable.IsUnknown() {
					data.OfferingBillable = types.BoolNull()
				}
			}
			if val, ok := sourceMap["offering_description"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.OfferingDescription = types.StringValue(str)
				}
			} else {
				if data.OfferingDescription.IsUnknown() {
					data.OfferingDescription = types.StringNull()
				}
			}
			if val, ok := sourceMap["offering_image"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.OfferingImage = types.StringValue(str)
				}
			} else {
				if data.OfferingImage.IsUnknown() {
					data.OfferingImage = types.StringNull()
				}
			}
			if val, ok := sourceMap["offering_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.OfferingName = types.StringValue(str)
				}
			} else {
				if data.OfferingName.IsUnknown() {
					data.OfferingName = types.StringNull()
				}
			}
			if val, ok := sourceMap["offering_shared"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.OfferingShared = types.BoolValue(b)
				}
			} else {
				if data.OfferingShared.IsUnknown() {
					data.OfferingShared = types.BoolNull()
				}
			}
			if val, ok := sourceMap["offering_slug"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.OfferingSlug = types.StringValue(str)
				}
			} else {
				if data.OfferingSlug.IsUnknown() {
					data.OfferingSlug = types.StringNull()
				}
			}
			if val, ok := sourceMap["offering_thumbnail"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.OfferingThumbnail = types.StringValue(str)
				}
			} else {
				if data.OfferingThumbnail.IsUnknown() {
					data.OfferingThumbnail = types.StringNull()
				}
			}
			if val, ok := sourceMap["offering_type"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.OfferingType = types.StringValue(str)
				}
			} else {
				if data.OfferingType.IsUnknown() {
					data.OfferingType = types.StringNull()
				}
			}
			if val, ok := sourceMap["offering_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.OfferingUuid = types.StringValue(str)
				}
			} else {
				if data.OfferingUuid.IsUnknown() {
					data.OfferingUuid = types.StringNull()
				}
			}
			if val, ok := sourceMap["parent_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ParentName = types.StringValue(str)
				}
			} else {
				if data.ParentName.IsUnknown() {
					data.ParentName = types.StringNull()
				}
			}
			if val, ok := sourceMap["parent_offering_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ParentOfferingName = types.StringValue(str)
				}
			} else {
				if data.ParentOfferingName.IsUnknown() {
					data.ParentOfferingName = types.StringNull()
				}
			}
			if val, ok := sourceMap["parent_offering_slug"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ParentOfferingSlug = types.StringValue(str)
				}
			} else {
				if data.ParentOfferingSlug.IsUnknown() {
					data.ParentOfferingSlug = types.StringNull()
				}
			}
			if val, ok := sourceMap["parent_offering_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ParentOfferingUuid = types.StringValue(str)
				}
			} else {
				if data.ParentOfferingUuid.IsUnknown() {
					data.ParentOfferingUuid = types.StringNull()
				}
			}
			if val, ok := sourceMap["parent_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ParentUuid = types.StringValue(str)
				}
			} else {
				if data.ParentUuid.IsUnknown() {
					data.ParentUuid = types.StringNull()
				}
			}
			if val, ok := sourceMap["paused"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.Paused = types.BoolValue(b)
				}
			} else {
				if data.Paused.IsUnknown() {
					data.Paused = types.BoolNull()
				}
			}
			if val, ok := sourceMap["plan"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Plan = types.StringValue(str)
				}
			} else {
				if data.Plan.IsUnknown() {
					data.Plan = types.StringNull()
				}
			}
			if val, ok := sourceMap["plan_description"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.PlanDescription = types.StringValue(str)
				}
			} else {
				if data.PlanDescription.IsUnknown() {
					data.PlanDescription = types.StringNull()
				}
			}
			if val, ok := sourceMap["plan_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.PlanName = types.StringValue(str)
				}
			} else {
				if data.PlanName.IsUnknown() {
					data.PlanName = types.StringNull()
				}
			}
			if val, ok := sourceMap["plan_unit"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.PlanUnit = types.StringValue(str)
				}
			} else {
				if data.PlanUnit.IsUnknown() {
					data.PlanUnit = types.StringNull()
				}
			}
			if val, ok := sourceMap["plan_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.PlanUuid = types.StringValue(str)
				}
			} else {
				if data.PlanUuid.IsUnknown() {
					data.PlanUuid = types.StringNull()
				}
			}
			if val, ok := sourceMap["project_description"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ProjectDescription = types.StringValue(str)
				}
			} else {
				if data.ProjectDescription.IsUnknown() {
					data.ProjectDescription = types.StringNull()
				}
			}
			if val, ok := sourceMap["project_end_date"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ProjectEndDate = types.StringValue(str)
				}
			} else {
				if data.ProjectEndDate.IsUnknown() {
					data.ProjectEndDate = types.StringNull()
				}
			}
			if val, ok := sourceMap["project_end_date_requested_by"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ProjectEndDateRequestedBy = types.StringValue(str)
				}
			} else {
				if data.ProjectEndDateRequestedBy.IsUnknown() {
					data.ProjectEndDateRequestedBy = types.StringNull()
				}
			}
			if val, ok := sourceMap["project_slug"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ProjectSlug = types.StringValue(str)
				}
			} else {
				if data.ProjectSlug.IsUnknown() {
					data.ProjectSlug = types.StringNull()
				}
			}
			if val, ok := sourceMap["provider_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ProviderName = types.StringValue(str)
				}
			} else {
				if data.ProviderName.IsUnknown() {
					data.ProviderName = types.StringNull()
				}
			}
			if val, ok := sourceMap["provider_slug"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ProviderSlug = types.StringValue(str)
				}
			} else {
				if data.ProviderSlug.IsUnknown() {
					data.ProviderSlug = types.StringNull()
				}
			}
			if val, ok := sourceMap["provider_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ProviderUuid = types.StringValue(str)
				}
			} else {
				if data.ProviderUuid.IsUnknown() {
					data.ProviderUuid = types.StringNull()
				}
			}
			if val, ok := sourceMap["report"]; ok && val != nil {
				// List of objects
				if arr, ok := val.([]interface{}); ok {
					items := make([]attr.Value, 0, len(arr))
					for _, item := range arr {
						if objMap, ok := item.(map[string]interface{}); ok {
							attrTypes := map[string]attr.Type{
								"body":   types.StringType,
								"header": types.StringType,
							}
							attrValues := map[string]attr.Value{
								"body": func() attr.Value {
									if v, ok := objMap["body"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
								"header": func() attr.Value {
									if v, ok := objMap["header"].(string); ok {
										return types.StringValue(v)
									}
									return types.StringNull()
								}(),
							}
							objVal, _ := types.ObjectValue(attrTypes, attrValues)
							items = append(items, objVal)
						}
					}
					listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
						"body":   types.StringType,
						"header": types.StringType,
					}}, items)
					data.Report = listVal
				}
			} else {
				if data.Report.IsUnknown() {
					data.Report = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
						"body":   types.StringType,
						"header": types.StringType,
					}})
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
			if val, ok := sourceMap["resource_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.ResourceUuid = types.StringValue(str)
				}
			} else {
				if data.ResourceUuid.IsUnknown() {
					data.ResourceUuid = types.StringNull()
				}
			}
			if val, ok := sourceMap["restrict_member_access"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.RestrictMemberAccess = types.BoolValue(b)
				}
			} else {
				if data.RestrictMemberAccess.IsUnknown() {
					data.RestrictMemberAccess = types.BoolNull()
				}
			}
			if val, ok := sourceMap["scope"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Scope = types.StringValue(str)
				}
			} else {
				if data.Scope.IsUnknown() {
					data.Scope = types.StringNull()
				}
			}
			if val, ok := sourceMap["slug"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Slug = types.StringValue(str)
				}
			} else {
				if data.Slug.IsUnknown() {
					data.Slug = types.StringNull()
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
			if val, ok := sourceMap["url"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Url = types.StringValue(str)
				}
			} else {
				if data.Url.IsUnknown() {
					data.Url = types.StringNull()
				}
			}
			if val, ok := sourceMap["user_requires_reconsent"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.UserRequiresReconsent = types.BoolValue(b)
				}
			} else {
				if data.UserRequiresReconsent.IsUnknown() {
					data.UserRequiresReconsent = types.BoolNull()
				}
			}
			if val, ok := sourceMap["username"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Username = types.StringValue(str)
				}
			} else {
				if data.Username.IsUnknown() {
					data.Username = types.StringNull()
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
