package network_rbac_policy

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackNetworkRbacPolicyResource{}
var _ resource.ResourceWithImportState = &OpenstackNetworkRbacPolicyResource{}

func NewOpenstackNetworkRbacPolicyResource() resource.Resource {
	return &OpenstackNetworkRbacPolicyResource{}
}

// OpenstackNetworkRbacPolicyResource defines the resource implementation.
type OpenstackNetworkRbacPolicyResource struct {
	client *OpenstackNetworkRbacPolicyClient
}

// OpenstackNetworkRbacPolicyResourceModel describes the resource data model.
type OpenstackNetworkRbacPolicyResourceModel struct {
	OpenstackNetworkRbacPolicyModel
	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackNetworkRbacPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_network_rbac_policy"
}

func (r *OpenstackNetworkRbacPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Network Rbac Policy resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Openstack Network Rbac Policy UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "ID of the backend",
			},
			"created": schema.StringAttribute{
				CustomType: timetypes.RFC3339Type{},
				Computed:   true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Created",
			},
			"network": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Network",
			},
			"network_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the network",
			},
			"policy_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Type of access granted - either shared access or external network access",
			},
			"target_tenant": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Target tenant",
			},
			"target_tenant_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the target tenant",
			},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Url",
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

	r.client = &OpenstackNetworkRbacPolicyClient{}
	if err := r.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			err.Error(),
		)
		return
	}
}

func (r *OpenstackNetworkRbacPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackNetworkRbacPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := OpenstackNetworkRbacPolicyCreateRequest{}

	requestBody.Network = data.Network.ValueStringPointer()
	if !data.PolicyType.IsNull() && !data.PolicyType.IsUnknown() {

		requestBody.PolicyType = data.PolicyType.ValueStringPointer()
	}

	requestBody.TargetTenant = data.TargetTenant.ValueStringPointer()

	apiResp, err := r.client.Create(ctx, &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Openstack Network Rbac Policy",
			"An error occurred while creating the Openstack Network Rbac Policy: "+err.Error(),
		)
		return
	}
	// Build composite ID from key fields
	compositeID := ""
	if apiResp.Network != nil {
		val := *apiResp.Network
		if strings.Contains(val, "/") {
			parts := strings.Split(strings.TrimRight(val, "/"), "/")
			val = parts[len(parts)-1]
		}
		compositeID += val
	}
	compositeID += "/"
	if apiResp.TargetTenant != nil {
		val := *apiResp.TargetTenant
		if strings.Contains(val, "/") {
			parts := strings.Split(strings.TrimRight(val, "/"), "/")
			val = parts[len(parts)-1]
		}
		compositeID += val
	}
	data.UUID = types.StringValue(compositeID)

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

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
			listResult, err := r.client.List(ctx, filters)
			if err != nil {
				resp.Diagnostics.AddError("Failed to lookup resource by composite keys", err.Error())
				return
			}

			if len(listResult) == 1 {
				data.UUID = types.StringPointerValue(listResult[0].UUID)
			} else if len(listResult) > 1 {
				resp.Diagnostics.AddError("Ambiguous resource", fmt.Sprintf("Found %d resources matching composite keys", len(listResult)))
				return
			} else {
				resp.Diagnostics.AddError("Resource not found", "No resource found matching composite keys")
				return
			}
		}
	}

	apiResp, err := r.client.Get(ctx, data.UUID.ValueString())
	if err != nil {
		if IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Openstack Network Rbac Policy",
			"An error occurred while reading the Openstack Network Rbac Policy: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackNetworkRbacPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackNetworkRbacPolicyResourceModel
	var state OpenstackNetworkRbacPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := OpenstackNetworkRbacPolicyUpdateRequest{}
	if !data.Network.IsNull() && !data.Network.IsUnknown() {

		requestBody.Network = data.Network.ValueStringPointer()
	}
	if !data.PolicyType.IsNull() && !data.PolicyType.IsUnknown() {

		requestBody.PolicyType = data.PolicyType.ValueStringPointer()
	}
	if !data.TargetTenant.IsNull() && !data.TargetTenant.IsUnknown() {

		requestBody.TargetTenant = data.TargetTenant.ValueStringPointer()
	}

	apiResp, err := r.client.Update(ctx, data.UUID.ValueString(), &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Openstack Network Rbac Policy",
			"An error occurred while updating the Openstack Network Rbac Policy: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackNetworkRbacPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackNetworkRbacPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Openstack Network Rbac Policy",
			"An error occurred while deleting the Openstack Network Rbac Policy: "+err.Error(),
		)
		return
	}
}

func (r *OpenstackNetworkRbacPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	// Parse composite ID: key1/key2/...
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
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
