package network_rbac_policy

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackNetworkRbacPolicyResource{}
var _ resource.ResourceWithImportState = &OpenstackNetworkRbacPolicyResource{}

func NewOpenstackNetworkRbacPolicyResource() resource.Resource {
	return &OpenstackNetworkRbacPolicyResource{}
}

// OpenstackNetworkRbacPolicyResource defines the resource implementation.
type OpenstackNetworkRbacPolicyResource struct {
	client *Client
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
				MarkdownDescription: "Resource UUID (used as Terraform ID)",
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
				Optional:            true,
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

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	r.client = NewClient(client)
}

func (r *OpenstackNetworkRbacPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackNetworkRbacPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := OpenstackNetworkRbacPolicyCreateRequest{
		Network:      data.Network.ValueStringPointer(),
		PolicyType:   data.PolicyType.ValueStringPointer(),
		TargetTenant: data.TargetTenant.ValueStringPointer(),
	}

	apiResp, err := r.client.CreateOpenstackNetworkRbacPolicy(ctx, &requestBody)
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

	createTimeout, diags := data.Timeouts.Create(ctx, 30*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackNetworkRbacPolicyResponse, error) {
		return r.client.GetOpenstackNetworkRbacPolicy(ctx, data.UUID.ValueString())
	}, createTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource creation", err.Error())
		return
	}
	apiResp = newResp

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
			listResult, err := r.client.ListOpenstackNetworkRbacPolicy(ctx, filters)
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

	apiResp, err := r.client.GetOpenstackNetworkRbacPolicy(ctx, data.UUID.ValueString())
	if err != nil {
		if client.IsNotFoundError(err) {
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

	requestBody := OpenstackNetworkRbacPolicyUpdateRequest{
		Network:      data.Network.ValueStringPointer(),
		PolicyType:   data.PolicyType.ValueStringPointer(),
		TargetTenant: data.TargetTenant.ValueStringPointer(),
	}

	apiResp, err := r.client.UpdateOpenstackNetworkRbacPolicy(ctx, data.UUID.ValueString(), &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Openstack Network Rbac Policy",
			"An error occurred while updating the Openstack Network Rbac Policy: "+err.Error(),
		)
		return
	}

	updateTimeout, diags := data.Timeouts.Update(ctx, 30*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackNetworkRbacPolicyResponse, error) {
		return r.client.GetOpenstackNetworkRbacPolicy(ctx, data.UUID.ValueString())
	}, updateTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource update", err.Error())
		return
	}
	apiResp = newResp

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackNetworkRbacPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackNetworkRbacPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteOpenstackNetworkRbacPolicy(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Openstack Network Rbac Policy",
			"An error occurred while deleting the Openstack Network Rbac Policy: "+err.Error(),
		)
		return
	}

	deleteTimeout, diags := data.Timeouts.Delete(ctx, 10*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err = common.WaitForDeletion(ctx, func(ctx context.Context) (*OpenstackNetworkRbacPolicyResponse, error) {
		return r.client.GetOpenstackNetworkRbacPolicy(ctx, data.UUID.ValueString())
	}, deleteTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource deletion", err.Error())
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

func (r *OpenstackNetworkRbacPolicyResource) mapResponseToModel(ctx context.Context, apiResp OpenstackNetworkRbacPolicyResponse, model *OpenstackNetworkRbacPolicyResourceModel) diag.Diagnostics {
	return model.CopyFrom(ctx, apiResp)
}
