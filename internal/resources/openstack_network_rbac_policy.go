package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackNetworkRbacPolicyResource{}
var _ resource.ResourceWithImportState = &OpenstackNetworkRbacPolicyResource{}

func NewOpenstackNetworkRbacPolicyResource() resource.Resource {
	return &OpenstackNetworkRbacPolicyResource{}
}

// OpenstackNetworkRbacPolicyResource defines the resource implementation.
type OpenstackNetworkRbacPolicyResource struct {
	client *client.Client
}

// OpenstackNetworkRbacPolicyResourceModel describes the resource data model.
type OpenstackNetworkRbacPolicyResourceModel struct {
	UUID             types.String   `tfsdk:"id"`
	BackendId        types.String   `tfsdk:"backend_id"`
	Created          types.String   `tfsdk:"created"`
	Network          types.String   `tfsdk:"network"`
	NetworkName      types.String   `tfsdk:"network_name"`
	PolicyType       types.String   `tfsdk:"policy_type"`
	TargetTenant     types.String   `tfsdk:"target_tenant"`
	TargetTenantName types.String   `tfsdk:"target_tenant_name"`
	Url              types.String   `tfsdk:"url"`
	Timeouts         timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackNetworkRbacPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_network_rbac_policy"
}

func (r *OpenstackNetworkRbacPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "OpenstackNetworkRbacPolicy resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"network": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: " ",
			},
			"network_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"policy_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Type of access granted - either shared access or external network access",
			},
			"target_tenant": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: " ",
			},
			"target_tenant_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
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

func (r *OpenstackNetworkRbacPolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OpenstackNetworkRbacPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackNetworkRbacPolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Prepare request body
	requestBody := map[string]interface{}{}
	if !data.Network.IsNull() && !data.Network.IsUnknown() {
		if v := data.Network.ValueString(); v != "" {
			requestBody["network"] = v
		}
	}
	if !data.PolicyType.IsNull() && !data.PolicyType.IsUnknown() {
		if v := data.PolicyType.ValueString(); v != "" {
			requestBody["policy_type"] = v
		}
	}
	if !data.TargetTenant.IsNull() && !data.TargetTenant.IsUnknown() {
		if v := data.TargetTenant.ValueString(); v != "" {
			requestBody["target_tenant"] = v
		}
	}

	// Call Waldur API to create resource
	var result map[string]interface{}
	err := r.client.Create(ctx, "/api/openstack-network-rbac-policies/", requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create OpenstackNetworkRbacPolicy",
			"An error occurred while creating the openstack_network_rbac_policy: "+err.Error(),
		)
		return
	}
	// Build composite ID from key fields
	compositeID := ""
	if v, ok := result["network"].(string); ok {
		compositeID += v
	} else if v, ok := result["network_uuid"].(string); ok {
		compositeID += v
	}
	compositeID += "/"
	if v, ok := result["target_tenant"].(string); ok {
		compositeID += v
	} else if v, ok := result["target_tenant_uuid"].(string); ok {
		compositeID += v
	}
	data.UUID = types.StringValue(compositeID)

	sourceMap := result
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

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackNetworkRbacPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackNetworkRbacPolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource
	var result map[string]interface{}
	// If UUID is unknown or contains slashes (composite key), try to look it up using composite keys
	if data.UUID.IsNull() || data.UUID.IsUnknown() || strings.Contains(data.UUID.ValueString(), "/") {
		filters := map[string]string{}
		if !data.Network.IsNull() {
			if v := data.Network.ValueString(); v != "" {
				filters["network"] = v
			}
		}
		if !data.TargetTenant.IsNull() {
			if v := data.TargetTenant.ValueString(); v != "" {
				filters["target_tenant"] = v
			}
		}

		expectedCount := 0
		expectedCount++
		expectedCount++

		if len(filters) == expectedCount {
			var listResult []map[string]interface{}
			// Construct list path from Retrieve path by removing {uuid}/
			listPath := strings.Replace("/api/openstack-network-rbac-policies/{uuid}/", "{uuid}/", "", 1)

			err := r.client.ListWithFilter(ctx, listPath, filters, &listResult)
			if err != nil {
				resp.Diagnostics.AddError("Failed to lookup resource by composite keys", err.Error())
				return
			}

			if len(listResult) == 1 {
				if uuid, ok := listResult[0]["uuid"].(string); ok {
					data.UUID = types.StringValue(uuid)
				}
			} else if len(listResult) > 1 {
				resp.Diagnostics.AddError("Ambiguous resource", fmt.Sprintf("Found %d resources matching composite keys", len(listResult)))
				return
			} else {
				resp.Diagnostics.AddError("Resource not found", "No resource found matching composite keys")
				return
			}
		}
	}

	retrievePath := strings.Replace("/api/openstack-network-rbac-policies/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read OpenstackNetworkRbacPolicy",
			"An error occurred while reading the openstack_network_rbac_policy: "+err.Error(),
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

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackNetworkRbacPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackNetworkRbacPolicyResourceModel
	var state OpenstackNetworkRbacPolicyResourceModel

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
	if !data.Network.IsNull() && !data.Network.IsUnknown() {
		if v := data.Network.ValueString(); v != "" {
			requestBody["network"] = v
		}
	}
	if !data.PolicyType.IsNull() && !data.PolicyType.IsUnknown() {
		if v := data.PolicyType.ValueString(); v != "" {
			requestBody["policy_type"] = v
		}
	}
	if !data.TargetTenant.IsNull() && !data.TargetTenant.IsUnknown() {
		if v := data.TargetTenant.ValueString(); v != "" {
			requestBody["target_tenant"] = v
		}
	}

	// Call Waldur API to update resource
	var result map[string]interface{}

	err := r.client.Update(ctx, "/api/openstack-network-rbac-policies/{uuid}/", data.UUID.ValueString(), requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update OpenstackNetworkRbacPolicy",
			"An error occurred while updating the openstack_network_rbac_policy: "+err.Error(),
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

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackNetworkRbacPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackNetworkRbacPolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Call Waldur API to delete resource
	err := r.client.DeleteByUUID(ctx, "/api/openstack-network-rbac-policies/{uuid}/", data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete OpenstackNetworkRbacPolicy",
			"An error occurred while deleting the openstack_network_rbac_policy: "+err.Error(),
		)
		return
	}
}

func (r *OpenstackNetworkRbacPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Parse composite ID: key1/key2/...
	parts := strings.Split(req.ID, "/")
	if len(parts) != 0 {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Expected format: <network>/<target_tenant>",
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("target_tenant"), parts[1])...)
}
