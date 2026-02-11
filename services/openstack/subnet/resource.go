package subnet

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

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
	client *OpenstackSubnetClient
}

// OpenstackSubnetResourceModel describes the resource data model.
type OpenstackSubnetResourceModel struct {
	OpenstackSubnetModel
	Timeouts timeouts.Value `tfsdk:"timeouts"`
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
				MarkdownDescription: "Openstack Subnet UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"allocation_pools": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"end": schema.StringAttribute{
							Required: true, MarkdownDescription: "An IPv4 or IPv6 address."},
						"start": schema.StringAttribute{
							Required: true, MarkdownDescription: "An IPv4 or IPv6 address."},
					},
				},
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.List{

					listplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Allocation Pools",
			},
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Backend Id"},
			"cidr": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Cidr"},
			"customer": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Customer"},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Description"},
			"disable_gateway": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{

					boolplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "If True, no gateway IP address will be allocated"},
			"dns_nameservers": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{

					listplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Dns Nameservers"},
			"enable_dhcp": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{

					boolplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "If True, DHCP service will be enabled on this subnet"},
			"error_message": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Error Message"},
			"gateway_ip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "IP address of the gateway for this subnet"},
			"host_routes": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"destination": schema.StringAttribute{
							Required: true, MarkdownDescription: "Destination"},
						"nexthop": schema.StringAttribute{
							Required: true, MarkdownDescription: "An IPv4 or IPv6 address."},
					},
				},
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.List{

					listplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Host Routes",
			},
			"ip_version": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{

					int64planmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "IP protocol version (4 or 6)"},
			"is_connected": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{

					boolplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Is subnet connected to the default tenant router."},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Marketplace Resource Uuid"},
			"name": schema.StringAttribute{
				Required: true, MarkdownDescription: "Name"},
			"network": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.RequiresReplace(),
				}, MarkdownDescription: "Network to which this subnet belongs"},
			"network_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Network Name"},
			"project": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Project"},
			"resource_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Resource Type"},
			"state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "State"},
			"tenant": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Tenant"},
			"tenant_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Tenant Name"},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Url"},
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

	r.client = &OpenstackSubnetClient{}
	if err := r.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			err.Error(),
		)
		return
	}
}

func (r *OpenstackSubnetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var data OpenstackSubnetResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := OpenstackSubnetCreateRequest{}
	if !data.Cidr.IsNull() && !data.Cidr.IsUnknown() {

		requestBody.Cidr = data.Cidr.ValueStringPointer()
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {

		requestBody.Description = data.Description.ValueStringPointer()
	}
	if !data.DisableGateway.IsNull() && !data.DisableGateway.IsUnknown() {

		requestBody.DisableGateway = data.DisableGateway.ValueBoolPointer()
	}
	if !data.GatewayIp.IsNull() && !data.GatewayIp.IsUnknown() {

		requestBody.GatewayIp = data.GatewayIp.ValueStringPointer()
	}

	requestBody.Name = data.Name.ValueStringPointer()
	resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.AllocationPools, &requestBody.AllocationPools)...)
	resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.DnsNameservers, &requestBody.DnsNameservers)...)
	resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.HostRoutes, &requestBody.HostRoutes)...)

	apiResp, err := r.client.Create(ctx, data.Network.ValueString(), &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Openstack Subnet",
			"An error occurred while creating the Openstack Subnet: "+err.Error(),
		)
		return
	}
	data.UUID = types.StringPointerValue(apiResp.UUID)
	createTimeout, diags := data.Timeouts.Create(ctx, common.DefaultCreateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackSubnetResponse, error) {
		return r.client.Get(ctx, data.UUID.ValueString())
	}, createTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource creation", err.Error())
		return
	}
	apiResp = newResp

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
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

	apiResp, err := r.client.Get(ctx, data.UUID.ValueString())
	if err != nil {
		if IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Openstack Subnet",
			"An error occurred while reading the Openstack Subnet: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

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

	var apiResp *OpenstackSubnetResponse
	updateTimeout, diags := data.Timeouts.Update(ctx, common.DefaultUpdateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	_ = updateTimeout
	anyChanges := false
	requestBody := OpenstackSubnetUpdateRequest{}
	if !data.AllocationPools.Equal(state.AllocationPools) {
		anyChanges = true
	}
	if !data.Cidr.IsNull() && !data.Cidr.IsUnknown() && !data.Cidr.Equal(state.Cidr) {
		anyChanges = true

		requestBody.Cidr = data.Cidr.ValueStringPointer()
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() && !data.Description.Equal(state.Description) {
		anyChanges = true

		requestBody.Description = data.Description.ValueStringPointer()
	}
	if !data.DisableGateway.IsNull() && !data.DisableGateway.IsUnknown() && !data.DisableGateway.Equal(state.DisableGateway) {
		anyChanges = true

		requestBody.DisableGateway = data.DisableGateway.ValueBoolPointer()
	}
	if !data.DnsNameservers.Equal(state.DnsNameservers) {
		anyChanges = true
	}
	if !data.GatewayIp.IsNull() && !data.GatewayIp.IsUnknown() && !data.GatewayIp.Equal(state.GatewayIp) {
		anyChanges = true

		requestBody.GatewayIp = data.GatewayIp.ValueStringPointer()
	}
	if !data.HostRoutes.Equal(state.HostRoutes) {
		anyChanges = true
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() && !data.Name.Equal(state.Name) {
		anyChanges = true

		requestBody.Name = data.Name.ValueStringPointer()
	}

	resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.AllocationPools, &requestBody.AllocationPools)...)
	resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.DnsNameservers, &requestBody.DnsNameservers)...)
	resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.HostRoutes, &requestBody.HostRoutes)...)

	if anyChanges {
		var err error
		apiResp, err = r.client.Update(ctx, data.UUID.ValueString(), &requestBody)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Update Openstack Subnet",
				"An error occurred while updating the Openstack Subnet: "+err.Error(),
			)
			return
		}
		// Wait for the resource to return to OK state
		newResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackSubnetResponse, error) {
			return r.client.Get(ctx, data.UUID.ValueString())
		}, updateTimeout)
		if err != nil {
			resp.Diagnostics.AddError("Wait for update failed", err.Error())
			return
		}
		apiResp = newResp
	}

	newResp, err := r.client.Get(ctx, data.UUID.ValueString())
	if err != nil {
		if IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Failed to Read Resource After Update", err.Error())
		return
	}
	apiResp = newResp

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackSubnetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var data OpenstackSubnetResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.client.Delete(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Openstack Subnet",
			"An error occurred while deleting the Openstack Subnet: "+err.Error(),
		)
		return
	}
	deleteTimeout, diags := data.Timeouts.Delete(ctx, common.DefaultDeleteTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err = common.WaitForDeletion(ctx, func(ctx context.Context) (*OpenstackSubnetResponse, error) {
		return r.client.Get(ctx, data.UUID.ValueString())
	}, deleteTimeout)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for resource deletion", err.Error())
		return
	}
}

func (r *OpenstackSubnetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	uuid := req.ID
	if uuid == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Import ID cannot be empty. Please provide the UUID of the Openstack Subnet.",
		)
		return
	}

	tflog.Info(ctx, "Importing Openstack Subnet", map[string]interface{}{
		"uuid": uuid,
	})

	apiResp, err := r.client.Get(ctx, uuid)
	if err != nil {
		if IsNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("Openstack Subnet with UUID '%s' does not exist or is not accessible.", uuid),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to Import Openstack Subnet",
			fmt.Sprintf("An error occurred while fetching the Openstack Subnet: %s", err.Error()),
		)
		return
	}

	var data OpenstackSubnetResourceModel
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
