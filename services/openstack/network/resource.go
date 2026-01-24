package network

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackNetworkResource{}
var _ resource.ResourceWithImportState = &OpenstackNetworkResource{}

func NewOpenstackNetworkResource() resource.Resource {
	return &OpenstackNetworkResource{}
}

// OpenstackNetworkResource defines the resource implementation.
type OpenstackNetworkResource struct {
	client *Client
}

// OpenstackNetworkResourceModel describes the resource data model.
type OpenstackNetworkResourceModel struct {
	UUID           types.String   `tfsdk:"id"`
	AccessUrl      types.String   `tfsdk:"access_url"`
	BackendId      types.String   `tfsdk:"backend_id"`
	Created        types.String   `tfsdk:"created"`
	Description    types.String   `tfsdk:"description"`
	ErrorMessage   types.String   `tfsdk:"error_message"`
	ErrorTraceback types.String   `tfsdk:"error_traceback"`
	IsExternal     types.Bool     `tfsdk:"is_external"`
	Modified       types.String   `tfsdk:"modified"`
	Mtu            types.Int64    `tfsdk:"mtu"`
	Name           types.String   `tfsdk:"name"`
	RbacPolicies   types.List     `tfsdk:"rbac_policies"`
	ResourceType   types.String   `tfsdk:"resource_type"`
	SegmentationId types.Int64    `tfsdk:"segmentation_id"`
	State          types.String   `tfsdk:"state"`
	Subnets        types.List     `tfsdk:"subnets"`
	Tenant         types.String   `tfsdk:"tenant"`
	TenantName     types.String   `tfsdk:"tenant_name"`
	TenantUuid     types.String   `tfsdk:"tenant_uuid"`
	Type           types.String   `tfsdk:"type"`
	Url            types.String   `tfsdk:"url"`
	Timeouts       timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackNetworkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_network"
}

func (r *OpenstackNetworkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Network resource",

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
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "ID of the backend",
			},
			"created": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Created",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description of the resource",
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
			"is_external": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Defines whether this network is external (public) or internal (private)",
			},
			"modified": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Modified",
			},
			"mtu": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The maximum transmission unit (MTU) value to address fragmentation.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name of the resource",
			},
			"rbac_policies": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"backend_id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "ID of the backend",
						},
						"created": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Created",
						},
						"network": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Network",
						},
						"network_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name of the network",
						},
						"policy_type": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Type of access granted - either shared access or external network access",
						},
						"target_tenant": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Target tenant",
						},
						"target_tenant_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name of the target tenant",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Url",
						},
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Rbac policies",
			},
			"resource_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Resource type",
			},
			"segmentation_id": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "VLAN ID for VLAN networks or tunnel ID for VXLAN/GRE networks",
			},
			"state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "State",
			},
			"subnets": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allocation_pools": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"end": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "An IPv4 or IPv6 address.",
									},
									"start": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "An IPv4 or IPv6 address.",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: "Allocation pools",
						},
						"cidr": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "IPv4 network address in CIDR format (e.g. 192.168.0.0/24)",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Description of the resource",
						},
						"enable_dhcp": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "If True, DHCP service will be enabled on this subnet",
						},
						"gateway_ip": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "IP address of the gateway for this subnet",
						},
						"ip_version": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "IP protocol version (4 or 6)",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the resource",
						},
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Subnets",
			},
			"tenant": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Required path parameter for resource creation",
			},
			"tenant_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the tenant",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the tenant",
			},
			"type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Network type, such as local, flat, vlan, vxlan, or gre",
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

func (r *OpenstackNetworkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OpenstackNetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackNetworkResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := OpenstackNetworkCreateRequest{
		Description: data.Description.ValueStringPointer(),
		Name:        data.Name.ValueStringPointer(),
	}

	apiResp, err := r.client.CreateOpenstackNetwork(ctx, data.Tenant.ValueString(), &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Openstack Network",
			"An error occurred while creating the Openstack Network: "+err.Error(),
		)
		return
	}
	data.UUID = types.StringPointerValue(apiResp.UUID)

	createTimeout, diags := data.Timeouts.Create(ctx, 30*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackNetworkResponse, error) {
		return r.client.GetOpenstackNetwork(ctx, data.UUID.ValueString())
	}, createTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource creation", err.Error())
		return
	}
	apiResp = newResp

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackNetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackNetworkResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource

	apiResp, err := r.client.GetOpenstackNetwork(ctx, data.UUID.ValueString())
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Openstack Network",
			"An error occurred while reading the Openstack Network: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackNetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackNetworkResourceModel
	var state OpenstackNetworkResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := OpenstackNetworkUpdateRequest{
		Description: data.Description.ValueStringPointer(),
		Name:        data.Name.ValueStringPointer(),
	}

	apiResp, err := r.client.UpdateOpenstackNetwork(ctx, data.UUID.ValueString(), &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Openstack Network",
			"An error occurred while updating the Openstack Network: "+err.Error(),
		)
		return
	}

	updateTimeout, diags := data.Timeouts.Update(ctx, 30*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackNetworkResponse, error) {
		return r.client.GetOpenstackNetwork(ctx, data.UUID.ValueString())
	}, updateTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource update", err.Error())
		return
	}
	apiResp = newResp

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackNetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackNetworkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteOpenstackNetwork(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Openstack Network",
			"An error occurred while deleting the Openstack Network: "+err.Error(),
		)
		return
	}

	deleteTimeout, diags := data.Timeouts.Delete(ctx, 10*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err = common.WaitForDeletion(ctx, func(ctx context.Context) (*OpenstackNetworkResponse, error) {
		return r.client.GetOpenstackNetwork(ctx, data.UUID.ValueString())
	}, deleteTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource deletion", err.Error())
		return
	}
}

func (r *OpenstackNetworkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	uuid := req.ID
	if uuid == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Import ID cannot be empty. Please provide the UUID of the Openstack Network.",
		)
		return
	}

	tflog.Info(ctx, "Importing Openstack Network", map[string]interface{}{
		"uuid": uuid,
	})

	apiResp, err := r.client.GetOpenstackNetwork(ctx, uuid)
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("Openstack Network with UUID '%s' does not exist or is not accessible.", uuid),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to Import Openstack Network",
			fmt.Sprintf("An error occurred while fetching the Openstack Network: %s", err.Error()),
		)
		return
	}

	var data OpenstackNetworkResourceModel
	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackNetworkResource) mapResponseToModel(ctx context.Context, apiResp OpenstackNetworkResponse, model *OpenstackNetworkResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.IsExternal = types.BoolPointerValue(apiResp.IsExternal)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.Mtu = types.Int64PointerValue(apiResp.Mtu)
	model.Name = types.StringPointerValue(apiResp.Name)

	{
		listValRbacPolicies, listDiagsRbacPolicies := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"backend_id":         types.StringType,
			"created":            types.StringType,
			"network":            types.StringType,
			"network_name":       types.StringType,
			"policy_type":        types.StringType,
			"target_tenant":      types.StringType,
			"target_tenant_name": types.StringType,
			"url":                types.StringType,
		}}, apiResp.RbacPolicies)
		diags.Append(listDiagsRbacPolicies...)
		model.RbacPolicies = listValRbacPolicies
	}
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.SegmentationId = types.Int64PointerValue(apiResp.SegmentationId)
	model.State = types.StringPointerValue(apiResp.State)

	{
		listValSubnets, listDiagsSubnets := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"allocation_pools": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"end":   types.StringType,
				"start": types.StringType,
			}}},
			"cidr":        types.StringType,
			"description": types.StringType,
			"enable_dhcp": types.BoolType,
			"gateway_ip":  types.StringType,
			"ip_version":  types.Int64Type,
			"name":        types.StringType,
		}}, apiResp.Subnets)
		diags.Append(listDiagsSubnets...)
		model.Subnets = listValSubnets
	}
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantName = types.StringPointerValue(apiResp.TenantName)
	model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
	model.Type = types.StringPointerValue(apiResp.Type)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
