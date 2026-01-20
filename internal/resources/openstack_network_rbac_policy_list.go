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

var _ list.ListResource = &OpenstackNetworkRbacPolicyList{}

type OpenstackNetworkRbacPolicyList struct {
	client *client.Client
}

func NewOpenstackNetworkRbacPolicyList() list.ListResource {
	return &OpenstackNetworkRbacPolicyList{}
}

func (l *OpenstackNetworkRbacPolicyList) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_network_rbac_policy"
}

func (l *OpenstackNetworkRbacPolicyList) ListResourceConfigSchema(ctx context.Context, req list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = schema.Schema{
		// Filter parameters can be added here if needed
		Attributes: map[string]schema.Attribute{},
	}
}

func (l *OpenstackNetworkRbacPolicyList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type OpenstackNetworkRbacPolicyListModel struct {
	// Add filter fields here if added to schema
}

func (l *OpenstackNetworkRbacPolicyList) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	var config OpenstackNetworkRbacPolicyListModel

	// Read config
	resp := req.Config.Get(ctx, &config)
	if resp.HasError() {
		stream.Results = list.ListResultsStreamDiagnostics(resp)
		return
	}

	// Call API
	var listResult []map[string]interface{}
	err := l.client.List(ctx, "/api/openstack-network-rbac-policies/", &listResult)
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
			var data OpenstackNetworkRbacPolicyResourceModel

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
			if val, ok := sourceMap["network"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.Network = types.StringValue(str)
				}
			} else {
				if data.Network.IsUnknown() {
					data.Network = types.StringNull()
				}
			}
			if val, ok := sourceMap["network_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.NetworkName = types.StringValue(str)
				}
			} else {
				if data.NetworkName.IsUnknown() {
					data.NetworkName = types.StringNull()
				}
			}
			if val, ok := sourceMap["policy_type"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.PolicyType = types.StringValue(str)
				}
			} else {
				if data.PolicyType.IsUnknown() {
					data.PolicyType = types.StringNull()
				}
			}
			if val, ok := sourceMap["target_tenant"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.TargetTenant = types.StringValue(str)
				}
			} else {
				if data.TargetTenant.IsUnknown() {
					data.TargetTenant = types.StringNull()
				}
			}
			if val, ok := sourceMap["target_tenant_name"]; ok && val != nil {
				if str, ok := val.(string); ok {
					data.TargetTenantName = types.StringValue(str)
				}
			} else {
				if data.TargetTenantName.IsUnknown() {
					data.TargetTenantName = types.StringNull()
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
