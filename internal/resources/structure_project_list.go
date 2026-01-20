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

var _ list.ListResource = &StructureProjectList{}

type StructureProjectList struct {
	client *client.Client
}

func NewStructureProjectList() list.ListResource {
	return &StructureProjectList{}
}

func (l *StructureProjectList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_structure_project"
}

func (l *StructureProjectList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		// Filter parameters can be added here if needed
		Attributes: map[string]schema.Attribute{},
	}
}

func (l *StructureProjectList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type StructureProjectListModel struct {
	// Add filter fields here if added to schema
}

func (l *StructureProjectList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config StructureProjectListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Call API
	var listResult []map[string]interface{}
	err := l.client.List(ctx, "/api/projects/", &listResult)
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
			var data StructureProjectResourceModel

			// Initialize lists and objects to empty/null if needed (or rely on map logic)

			sourceMap := item

			// Reuse the mapResponseFields logic (embedded here or via shared template)
			// Map response fields to data model
			_ = sourceMap
			if val, ok := sourceMap["backend_id"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.BackendId = types.StringValue(str)
				}
			} else {
				if data.BackendId.IsUnknown() {
					data.BackendId = types.StringNull()
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
			if val, ok := sourceMap["customer_display_billing_info_in_projects"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.CustomerDisplayBillingInfoInProjects = types.BoolValue(b)
				}
			} else {
				if data.CustomerDisplayBillingInfoInProjects.IsUnknown() {
					data.CustomerDisplayBillingInfoInProjects = types.BoolNull()
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
			if val, ok := sourceMap["grace_period_days"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.GracePeriodDays = types.Int64Value(int64(num))
				}
			} else {
				if data.GracePeriodDays.IsUnknown() {
					data.GracePeriodDays = types.Int64Null()
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
			if val, ok := sourceMap["is_industry"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.IsIndustry = types.BoolValue(b)
				}
			} else {
				if data.IsIndustry.IsUnknown() {
					data.IsIndustry = types.BoolNull()
				}
			}
			if val, ok := sourceMap["is_removed"]; ok && val != nil {
				if b, ok := val.(bool); ok {
					data.IsRemoved = types.BoolValue(b)
				}
			} else {
				if data.IsRemoved.IsUnknown() {
					data.IsRemoved = types.BoolNull()
				}
			}
			if val, ok := sourceMap["kind"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Kind = types.StringValue(str)
				}
			} else {
				if data.Kind.IsUnknown() {
					data.Kind = types.StringNull()
				}
			}
			if val, ok := sourceMap["max_service_accounts"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.MaxServiceAccounts = types.Int64Value(int64(num))
				}
			} else {
				if data.MaxServiceAccounts.IsUnknown() {
					data.MaxServiceAccounts = types.Int64Null()
				}
			}
			if val, ok := sourceMap["oecd_fos_2007_code"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.OecdFos2007Code = types.StringValue(str)
				}
			} else {
				if data.OecdFos2007Code.IsUnknown() {
					data.OecdFos2007Code = types.StringNull()
				}
			}
			if val, ok := sourceMap["oecd_fos_2007_label"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.OecdFos2007Label = types.StringValue(str)
				}
			} else {
				if data.OecdFos2007Label.IsUnknown() {
					data.OecdFos2007Label = types.StringNull()
				}
			}
			if val, ok := sourceMap["project_credit"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.ProjectCredit = types.Float64Value(num)
				}
			} else {
				if data.ProjectCredit.IsUnknown() {
					data.ProjectCredit = types.Float64Null()
				}
			}
			if val, ok := sourceMap["resources_count"]; ok && val != nil {
				if num, ok := val.(float64); ok {
					data.ResourcesCount = types.Int64Value(int64(num))
				}
			} else {
				if data.ResourcesCount.IsUnknown() {
					data.ResourcesCount = types.Int64Null()
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
			if val, ok := sourceMap["staff_notes"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.StaffNotes = types.StringValue(str)
				}
			} else {
				if data.StaffNotes.IsUnknown() {
					data.StaffNotes = types.StringNull()
				}
			}
			if val, ok := sourceMap["start_date"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.StartDate = types.StringValue(str)
				}
			} else {
				if data.StartDate.IsUnknown() {
					data.StartDate = types.StringNull()
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
			if val, ok := sourceMap["type_uuid"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.TypeUuid = types.StringValue(str)
				}
			} else {
				if data.TypeUuid.IsUnknown() {
					data.TypeUuid = types.StringNull()
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
