package subnet

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackSubnetResource{}
var _ resource.ResourceWithImportState = &OpenstackSubnetResource{}

func NewOpenstackSubnetResource() resource.Resource {
	return &OpenstackSubnetResource{}
}

// OpenstackSubnetResource defines the resource implementation.
type OpenstackSubnetResource struct {
	client *Client
}

// OpenstackSubnetResourceModel describes the resource data model.
type OpenstackSubnetResourceModel struct {
	UUID            types.String   `tfsdk:"id"`
	AccessUrl       types.String   `tfsdk:"access_url"`
	AllocationPools types.List     `tfsdk:"allocation_pools"`
	BackendId       types.String   `tfsdk:"backend_id"`
	Cidr            types.String   `tfsdk:"cidr"`
	Created         types.String   `tfsdk:"created"`
	Description     types.String   `tfsdk:"description"`
	DisableGateway  types.Bool     `tfsdk:"disable_gateway"`
	DnsNameservers  types.List     `tfsdk:"dns_nameservers"`
	EnableDhcp      types.Bool     `tfsdk:"enable_dhcp"`
	ErrorMessage    types.String   `tfsdk:"error_message"`
	ErrorTraceback  types.String   `tfsdk:"error_traceback"`
	GatewayIp       types.String   `tfsdk:"gateway_ip"`
	HostRoutes      types.List     `tfsdk:"host_routes"`
	IpVersion       types.Int64    `tfsdk:"ip_version"`
	IsConnected     types.Bool     `tfsdk:"is_connected"`
	Modified        types.String   `tfsdk:"modified"`
	Name            types.String   `tfsdk:"name"`
	Network         types.String   `tfsdk:"network"`
	NetworkName     types.String   `tfsdk:"network_name"`
	ResourceType    types.String   `tfsdk:"resource_type"`
	State           types.String   `tfsdk:"state"`
	Tenant          types.String   `tfsdk:"tenant"`
	TenantName      types.String   `tfsdk:"tenant_name"`
	Url             types.String   `tfsdk:"url"`
	Timeouts        timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackSubnetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_subnet"
}

func (r *OpenstackSubnetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Subnet resource",

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
			"allocation_pools": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"end": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "An IPv4 or IPv6 address.",
						},
						"start": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "An IPv4 or IPv6 address.",
						},
					},
				},
				Optional:            true,
				MarkdownDescription: "Allocation pools",
			},
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "ID of the backend",
			},
			"cidr": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Cidr",
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
			"disable_gateway": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "If True, no gateway IP address will be allocated",
			},
			"dns_nameservers": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Optional:            true,
				MarkdownDescription: "Dns nameservers",
			},
			"enable_dhcp": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "If True, DHCP service will be enabled on this subnet",
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
			"gateway_ip": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "IP address of the gateway for this subnet",
			},
			"host_routes": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"destination": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Destination",
						},
						"nexthop": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "An IPv4 or IPv6 address.",
						},
					},
				},
				Optional:            true,
				MarkdownDescription: "Host routes",
			},
			"ip_version": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "IP protocol version (4 or 6)",
			},
			"is_connected": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Is subnet connected to the default tenant router.",
			},
			"modified": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name of the resource",
			},
			"network": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Required path parameter for resource creation",
			},
			"network_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the network",
			},
			"resource_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Resource type",
			},
			"state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "State",
			},
			"tenant": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Tenant",
			},
			"tenant_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the tenant",
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

func (r *OpenstackSubnetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OpenstackSubnetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackSubnetResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := OpenstackSubnetCreateRequest{
		Cidr:           data.Cidr.ValueStringPointer(),
		Description:    data.Description.ValueStringPointer(),
		DisableGateway: data.DisableGateway.ValueBoolPointer(),
		GatewayIp:      data.GatewayIp.ValueStringPointer(),
		Name:           data.Name.ValueStringPointer(),
	}
	{
		// Object array or other
		var items []common.OpenStackSubNetAllocationPoolRequest
		diags := data.AllocationPools.ElementsAs(ctx, &items, false)
		resp.Diagnostics.Append(diags...)
		if !diags.HasError() {
			if !data.AllocationPools.IsNull() && !data.AllocationPools.IsUnknown() {
				requestBody.AllocationPools = &items
			}
		}
	}
	{
		var items []string
		diags := data.DnsNameservers.ElementsAs(ctx, &items, false)
		resp.Diagnostics.Append(diags...)
		if !diags.HasError() {
			if !data.DnsNameservers.IsNull() && !data.DnsNameservers.IsUnknown() {
				requestBody.DnsNameservers = &items
			}
		}
	}
	{
		// Object array or other
		var items []common.OpenStackStaticRouteRequest
		diags := data.HostRoutes.ElementsAs(ctx, &items, false)
		resp.Diagnostics.Append(diags...)
		if !diags.HasError() {
			if !data.HostRoutes.IsNull() && !data.HostRoutes.IsUnknown() {
				requestBody.HostRoutes = &items
			}
		}
	}

	apiResp, err := r.client.CreateOpenstackSubnet(ctx, data.Network.ValueString(), &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Openstack Subnet",
			"An error occurred while creating the Openstack Subnet: "+err.Error(),
		)
		return
	}
	data.UUID = types.StringPointerValue(apiResp.UUID)

	createTimeout, diags := data.Timeouts.Create(ctx, 30*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackSubnetResponse, error) {
		return r.client.GetOpenstackSubnet(ctx, data.UUID.ValueString())
	}, createTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource creation", err.Error())
		return
	}
	apiResp = newResp

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackSubnetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackSubnetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource

	apiResp, err := r.client.GetOpenstackSubnet(ctx, data.UUID.ValueString())
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Openstack Subnet",
			"An error occurred while reading the Openstack Subnet: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackSubnetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackSubnetResourceModel
	var state OpenstackSubnetResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := OpenstackSubnetUpdateRequest{
		Cidr:           data.Cidr.ValueStringPointer(),
		Description:    data.Description.ValueStringPointer(),
		DisableGateway: data.DisableGateway.ValueBoolPointer(),
		GatewayIp:      data.GatewayIp.ValueStringPointer(),
		Name:           data.Name.ValueStringPointer(),
	}
	{
		// Object array or other
		var items []common.OpenStackSubNetAllocationPoolRequest
		diags := data.AllocationPools.ElementsAs(ctx, &items, false)
		resp.Diagnostics.Append(diags...)
		if !diags.HasError() {
			if !data.AllocationPools.IsNull() && !data.AllocationPools.IsUnknown() {
				requestBody.AllocationPools = &items
			}
		}
	}
	{
		var items []string
		diags := data.DnsNameservers.ElementsAs(ctx, &items, false)
		resp.Diagnostics.Append(diags...)
		if !diags.HasError() {
			if !data.DnsNameservers.IsNull() && !data.DnsNameservers.IsUnknown() {
				requestBody.DnsNameservers = &items
			}
		}
	}
	{
		// Object array or other
		var items []common.OpenStackStaticRouteRequest
		diags := data.HostRoutes.ElementsAs(ctx, &items, false)
		resp.Diagnostics.Append(diags...)
		if !diags.HasError() {
			if !data.HostRoutes.IsNull() && !data.HostRoutes.IsUnknown() {
				requestBody.HostRoutes = &items
			}
		}
	}

	apiResp, err := r.client.UpdateOpenstackSubnet(ctx, data.UUID.ValueString(), &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Openstack Subnet",
			"An error occurred while updating the Openstack Subnet: "+err.Error(),
		)
		return
	}

	updateTimeout, diags := data.Timeouts.Update(ctx, 30*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackSubnetResponse, error) {
		return r.client.GetOpenstackSubnet(ctx, data.UUID.ValueString())
	}, updateTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource update", err.Error())
		return
	}
	apiResp = newResp

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackSubnetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackSubnetResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteOpenstackSubnet(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Openstack Subnet",
			"An error occurred while deleting the Openstack Subnet: "+err.Error(),
		)
		return
	}

	deleteTimeout, diags := data.Timeouts.Delete(ctx, 10*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err = common.WaitForDeletion(ctx, func(ctx context.Context) (*OpenstackSubnetResponse, error) {
		return r.client.GetOpenstackSubnet(ctx, data.UUID.ValueString())
	}, deleteTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource deletion", err.Error())
		return
	}
}

func (r *OpenstackSubnetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *OpenstackSubnetResource) mapResponseToModel(ctx context.Context, apiResp OpenstackSubnetResponse, model *OpenstackSubnetResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)

	{
		listValAllocationPools, listDiagsAllocationPools := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"end":   types.StringType,
			"start": types.StringType,
		}}, apiResp.AllocationPools)
		diags.Append(listDiagsAllocationPools...)
		model.AllocationPools = listValAllocationPools
	}
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Cidr = types.StringPointerValue(apiResp.Cidr)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.DisableGateway = types.BoolPointerValue(apiResp.DisableGateway)
	listValDnsNameservers, listDiagsDnsNameservers := types.ListValueFrom(ctx, types.StringType, apiResp.DnsNameservers)
	model.DnsNameservers = listValDnsNameservers
	diags.Append(listDiagsDnsNameservers...)
	model.EnableDhcp = types.BoolPointerValue(apiResp.EnableDhcp)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.GatewayIp = types.StringPointerValue(apiResp.GatewayIp)

	{
		listValHostRoutes, listDiagsHostRoutes := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
			"nexthop":     types.StringType,
		}}, apiResp.HostRoutes)
		diags.Append(listDiagsHostRoutes...)
		model.HostRoutes = listValHostRoutes
	}
	model.IpVersion = types.Int64PointerValue(apiResp.IpVersion)
	model.IsConnected = types.BoolPointerValue(apiResp.IsConnected)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.Network = types.StringPointerValue(apiResp.Network)
	model.NetworkName = types.StringPointerValue(apiResp.NetworkName)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.State = types.StringPointerValue(apiResp.State)
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantName = types.StringPointerValue(apiResp.TenantName)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
