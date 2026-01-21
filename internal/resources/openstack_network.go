package resources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackNetworkResource{}
var _ resource.ResourceWithImportState = &OpenstackNetworkResource{}

func NewOpenstackNetworkResource() resource.Resource {
	return &OpenstackNetworkResource{}
}

// OpenstackNetworkResource defines the resource implementation.
type OpenstackNetworkResource struct {
	client *client.Client
}

// OpenstackNetworkApiResponse is the API response model.
type OpenstackNetworkApiResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl      *string                                `json:"access_url" tfsdk:"access_url"`
	BackendId      *string                                `json:"backend_id" tfsdk:"backend_id"`
	Created        *string                                `json:"created" tfsdk:"created"`
	Description    *string                                `json:"description" tfsdk:"description"`
	ErrorMessage   *string                                `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback *string                                `json:"error_traceback" tfsdk:"error_traceback"`
	IsExternal     *bool                                  `json:"is_external" tfsdk:"is_external"`
	Modified       *string                                `json:"modified" tfsdk:"modified"`
	Mtu            *int64                                 `json:"mtu" tfsdk:"mtu"`
	Name           *string                                `json:"name" tfsdk:"name"`
	RbacPolicies   []OpenstackNetworkRbacPoliciesResponse `json:"rbac_policies" tfsdk:"rbac_policies"`
	ResourceType   *string                                `json:"resource_type" tfsdk:"resource_type"`
	SegmentationId *int64                                 `json:"segmentation_id" tfsdk:"segmentation_id"`
	State          *string                                `json:"state" tfsdk:"state"`
	Subnets        []OpenstackNetworkSubnetsResponse      `json:"subnets" tfsdk:"subnets"`
	Tenant         *string                                `json:"tenant" tfsdk:"tenant"`
	TenantName     *string                                `json:"tenant_name" tfsdk:"tenant_name"`
	TenantUuid     *string                                `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Type           *string                                `json:"type" tfsdk:"type"`
	Url            *string                                `json:"url" tfsdk:"url"`
}

type OpenstackNetworkRbacPoliciesResponse struct {
	BackendId        *string `json:"backend_id" tfsdk:"backend_id"`
	Created          *string `json:"created" tfsdk:"created"`
	Network          *string `json:"network" tfsdk:"network"`
	NetworkName      *string `json:"network_name" tfsdk:"network_name"`
	PolicyType       *string `json:"policy_type" tfsdk:"policy_type"`
	TargetTenant     *string `json:"target_tenant" tfsdk:"target_tenant"`
	TargetTenantName *string `json:"target_tenant_name" tfsdk:"target_tenant_name"`
	Url              *string `json:"url" tfsdk:"url"`
}

type OpenstackNetworkSubnetsResponse struct {
	AllocationPools []OpenstackNetworkSubnetsAllocationPoolsResponse `json:"allocation_pools" tfsdk:"allocation_pools"`
	Cidr            *string                                          `json:"cidr" tfsdk:"cidr"`
	Description     *string                                          `json:"description" tfsdk:"description"`
	EnableDhcp      *bool                                            `json:"enable_dhcp" tfsdk:"enable_dhcp"`
	GatewayIp       *string                                          `json:"gateway_ip" tfsdk:"gateway_ip"`
	IpVersion       *int64                                           `json:"ip_version" tfsdk:"ip_version"`
	Name            *string                                          `json:"name" tfsdk:"name"`
}

type OpenstackNetworkSubnetsAllocationPoolsResponse struct {
	End   *string `json:"end" tfsdk:"end"`
	Start *string `json:"start" tfsdk:"start"`
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
				Computed:            true,
				MarkdownDescription: " ",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"is_external": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Defines whether this network is external (public) or internal (private)",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"mtu": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The maximum transmission unit (MTU) value to address fragmentation.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: " ",
			},
			"rbac_policies": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"backend_id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"created": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " ",
						},
						"network": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
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
							Optional:            true,
							Computed:            true,
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
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"segmentation_id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "VLAN ID for VLAN networks or tunnel ID for VXLAN/GRE networks",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"subnets": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allocation_pools": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"end": schema.StringAttribute{
										Optional:            true,
										Computed:            true,
										MarkdownDescription: "An IPv4 or IPv6 address.",
									},
									"start": schema.StringAttribute{
										Optional:            true,
										Computed:            true,
										MarkdownDescription: "An IPv4 or IPv6 address.",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: " ",
						},
						"cidr": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "IPv4 network address in CIDR format (e.g. 192.168.0.0/24)",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"enable_dhcp": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "If True, DHCP service will be enabled on this subnet",
						},
						"gateway_ip": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "IP address of the gateway for this subnet",
						},
						"ip_version": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "IP protocol version (4 or 6)",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"tenant": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Required path parameter for resource creation",
			},
			"tenant_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Network type, such as local, flat, vlan, vxlan, or gre",
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

	r.client = client
}

func (r *OpenstackNetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackNetworkResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to create resource
	var apiResp OpenstackNetworkApiResponse
	// Custom create operation via parent resource
	createPath := "/api/openstack-tenants/{uuid}/create_network/"
	createPath = strings.Replace(createPath, "{uuid}", data.Tenant.ValueString(), 1) // Prepare request body
	requestBody := map[string]interface{}{}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		if v := data.Description.ValueString(); v != "" {
			requestBody["description"] = v
		}
	}
	requestBody["name"] = data.Name.ValueString()
	err := r.client.Post(ctx, createPath, requestBody, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Openstack Network",
			"An error occurred while creating the Openstack Network: "+err.Error(),
		)
		return
	}
	// Extract UUID from response
	if apiResp.UUID != nil {
		data.UUID = types.StringPointerValue(apiResp.UUID)
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

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

	retrievePath := strings.Replace("/api/openstack-networks/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	var apiResp OpenstackNetworkApiResponse
	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &apiResp)
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

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

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

	// Prepare request body
	requestBody := map[string]interface{}{}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		if v := data.Description.ValueString(); v != "" {
			requestBody["description"] = v
		}
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		if v := data.Name.ValueString(); v != "" {
			requestBody["name"] = v
		}
	}

	// Call Waldur API to update resource
	var apiResp OpenstackNetworkApiResponse

	err := r.client.Update(ctx, "/api/openstack-networks/{uuid}/", data.UUID.ValueString(), requestBody, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Openstack Network",
			"An error occurred while updating the Openstack Network: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if apiResp.UUID != nil {
		data.UUID = types.StringPointerValue(apiResp.UUID)
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackNetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackNetworkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to delete resource
	err := r.client.DeleteByUUID(ctx, "/api/openstack-networks/{uuid}/", data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Openstack Network",
			"An error occurred while deleting the Openstack Network: "+err.Error(),
		)
		return
	}
}

func (r *OpenstackNetworkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *OpenstackNetworkResource) mapResponseToModel(ctx context.Context, apiResp OpenstackNetworkApiResponse, model *OpenstackNetworkResourceModel) diag.Diagnostics {
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
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.SegmentationId = types.Int64PointerValue(apiResp.SegmentationId)
	model.State = types.StringPointerValue(apiResp.State)
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
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantName = types.StringPointerValue(apiResp.TenantName)
	model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
	model.Type = types.StringPointerValue(apiResp.Type)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
