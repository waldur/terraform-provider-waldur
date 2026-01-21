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
var _ resource.Resource = &OpenstackPortResource{}
var _ resource.ResourceWithImportState = &OpenstackPortResource{}

func NewOpenstackPortResource() resource.Resource {
	return &OpenstackPortResource{}
}

// OpenstackPortResource defines the resource implementation.
type OpenstackPortResource struct {
	client *client.Client
}

// OpenstackPortApiResponse is the API response model.
type OpenstackPortApiResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl           *string                                    `json:"access_url" tfsdk:"access_url"`
	AdminStateUp        *bool                                      `json:"admin_state_up" tfsdk:"admin_state_up"`
	AllowedAddressPairs []OpenstackPortAllowedAddressPairsResponse `json:"allowed_address_pairs" tfsdk:"allowed_address_pairs"`
	BackendId           *string                                    `json:"backend_id" tfsdk:"backend_id"`
	Created             *string                                    `json:"created" tfsdk:"created"`
	Description         *string                                    `json:"description" tfsdk:"description"`
	DeviceId            *string                                    `json:"device_id" tfsdk:"device_id"`
	DeviceOwner         *string                                    `json:"device_owner" tfsdk:"device_owner"`
	ErrorMessage        *string                                    `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback      *string                                    `json:"error_traceback" tfsdk:"error_traceback"`
	FixedIps            []OpenstackPortFixedIpsResponse            `json:"fixed_ips" tfsdk:"fixed_ips"`
	FloatingIps         []string                                   `json:"floating_ips" tfsdk:"floating_ips"`
	MacAddress          *string                                    `json:"mac_address" tfsdk:"mac_address"`
	Modified            *string                                    `json:"modified" tfsdk:"modified"`
	Name                *string                                    `json:"name" tfsdk:"name"`
	Network             *string                                    `json:"network" tfsdk:"network"`
	NetworkName         *string                                    `json:"network_name" tfsdk:"network_name"`
	NetworkUuid         *string                                    `json:"network_uuid" tfsdk:"network_uuid"`
	PortSecurityEnabled *bool                                      `json:"port_security_enabled" tfsdk:"port_security_enabled"`
	ResourceType        *string                                    `json:"resource_type" tfsdk:"resource_type"`
	SecurityGroups      []OpenstackPortSecurityGroupsResponse      `json:"security_groups" tfsdk:"security_groups"`
	State               *string                                    `json:"state" tfsdk:"state"`
	Status              *string                                    `json:"status" tfsdk:"status"`
	TargetTenant        *string                                    `json:"target_tenant" tfsdk:"target_tenant"`
	Tenant              *string                                    `json:"tenant" tfsdk:"tenant"`
	TenantName          *string                                    `json:"tenant_name" tfsdk:"tenant_name"`
	TenantUuid          *string                                    `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Url                 *string                                    `json:"url" tfsdk:"url"`
}

type OpenstackPortAllowedAddressPairsResponse struct {
	IpAddress  *string `json:"ip_address" tfsdk:"ip_address"`
	MacAddress *string `json:"mac_address" tfsdk:"mac_address"`
}

type OpenstackPortFixedIpsResponse struct {
	IpAddress *string `json:"ip_address" tfsdk:"ip_address"`
	SubnetId  *string `json:"subnet_id" tfsdk:"subnet_id"`
}

type OpenstackPortSecurityGroupsResponse struct {
	Name *string `json:"name" tfsdk:"name"`
}

// OpenstackPortResourceModel describes the resource data model.
type OpenstackPortResourceModel struct {
	UUID                types.String   `tfsdk:"id"`
	AccessUrl           types.String   `tfsdk:"access_url"`
	AdminStateUp        types.Bool     `tfsdk:"admin_state_up"`
	AllowedAddressPairs types.List     `tfsdk:"allowed_address_pairs"`
	BackendId           types.String   `tfsdk:"backend_id"`
	Created             types.String   `tfsdk:"created"`
	Description         types.String   `tfsdk:"description"`
	DeviceId            types.String   `tfsdk:"device_id"`
	DeviceOwner         types.String   `tfsdk:"device_owner"`
	ErrorMessage        types.String   `tfsdk:"error_message"`
	ErrorTraceback      types.String   `tfsdk:"error_traceback"`
	FixedIps            types.List     `tfsdk:"fixed_ips"`
	FloatingIps         types.List     `tfsdk:"floating_ips"`
	MacAddress          types.String   `tfsdk:"mac_address"`
	Modified            types.String   `tfsdk:"modified"`
	Name                types.String   `tfsdk:"name"`
	Network             types.String   `tfsdk:"network"`
	NetworkName         types.String   `tfsdk:"network_name"`
	NetworkUuid         types.String   `tfsdk:"network_uuid"`
	PortSecurityEnabled types.Bool     `tfsdk:"port_security_enabled"`
	ResourceType        types.String   `tfsdk:"resource_type"`
	SecurityGroups      types.List     `tfsdk:"security_groups"`
	State               types.String   `tfsdk:"state"`
	Status              types.String   `tfsdk:"status"`
	TargetTenant        types.String   `tfsdk:"target_tenant"`
	Tenant              types.String   `tfsdk:"tenant"`
	TenantName          types.String   `tfsdk:"tenant_name"`
	TenantUuid          types.String   `tfsdk:"tenant_uuid"`
	Url                 types.String   `tfsdk:"url"`
	Timeouts            timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackPortResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_port"
}

func (r *OpenstackPortResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Port resource",

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
			"admin_state_up": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Administrative state of the port. If down, port does not forward packets",
			},
			"allowed_address_pairs": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"mac_address": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Port ID in OpenStack",
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
			"device_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of device (instance, router etc) to which this port is connected",
			},
			"device_owner": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Entity that uses this port (e.g. network:router_interface)",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"fixed_ips": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "IP address to assign to the port",
						},
						"subnet_id": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "ID of the subnet in which to assign the IP address",
						},
					},
				},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"floating_ips": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"mac_address": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "MAC address of the port",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: " ",
			},
			"network": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Network to which this port belongs",
			},
			"network_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"network_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"port_security_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "If True, security groups and rules will be applied to this port",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"security_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"status": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Port status in OpenStack (e.g. ACTIVE, DOWN)",
			},
			"target_tenant": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Target tenant for shared network port creation. If not specified, defaults to network's tenant.",
			},
			"tenant": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "OpenStack tenant this port belongs to",
			},
			"tenant_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"tenant_uuid": schema.StringAttribute{
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

func (r *OpenstackPortResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OpenstackPortResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackPortResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to create resource
	var apiResp OpenstackPortApiResponse // Prepare request body
	requestBody := map[string]interface{}{}
	if !data.AllowedAddressPairs.IsNull() && !data.AllowedAddressPairs.IsUnknown() {
		if v := ConvertTFValue(data.AllowedAddressPairs); v != nil {
			requestBody["allowed_address_pairs"] = v
		}
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		if v := data.Description.ValueString(); v != "" {
			requestBody["description"] = v
		}
	}
	if !data.FixedIps.IsNull() && !data.FixedIps.IsUnknown() {
		if v := ConvertTFValue(data.FixedIps); v != nil {
			requestBody["fixed_ips"] = v
		}
	}
	if !data.MacAddress.IsNull() && !data.MacAddress.IsUnknown() {
		if v := data.MacAddress.ValueString(); v != "" {
			requestBody["mac_address"] = v
		}
	}
	requestBody["name"] = data.Name.ValueString()
	if !data.Network.IsNull() && !data.Network.IsUnknown() {
		if v := data.Network.ValueString(); v != "" {
			requestBody["network"] = v
		}
	}
	if !data.PortSecurityEnabled.IsNull() && !data.PortSecurityEnabled.IsUnknown() {
		requestBody["port_security_enabled"] = data.PortSecurityEnabled.ValueBool()
	}
	if !data.SecurityGroups.IsNull() && !data.SecurityGroups.IsUnknown() {
		if v := ConvertTFValue(data.SecurityGroups); v != nil {
			requestBody["security_groups"] = v
		}
	}
	if !data.TargetTenant.IsNull() && !data.TargetTenant.IsUnknown() {
		if v := data.TargetTenant.ValueString(); v != "" {
			requestBody["target_tenant"] = v
		}
	}
	err := r.client.Create(ctx, "/api/openstack-ports/", requestBody, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Openstack Port",
			"An error occurred while creating the Openstack Port: "+err.Error(),
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

func (r *OpenstackPortResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackPortResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource

	retrievePath := strings.Replace("/api/openstack-ports/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	var apiResp OpenstackPortApiResponse
	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Openstack Port",
			"An error occurred while reading the Openstack Port: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackPortResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackPortResourceModel
	var state OpenstackPortResourceModel
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
	if !data.SecurityGroups.IsNull() && !data.SecurityGroups.IsUnknown() {
		if v := ConvertTFValue(data.SecurityGroups); v != nil {
			requestBody["security_groups"] = v
		}
	}
	if !data.TargetTenant.IsNull() && !data.TargetTenant.IsUnknown() {
		if v := data.TargetTenant.ValueString(); v != "" {
			requestBody["target_tenant"] = v
		}
	}

	// Call Waldur API to update resource
	var apiResp OpenstackPortApiResponse

	err := r.client.Update(ctx, "/api/openstack-ports/{uuid}/", data.UUID.ValueString(), requestBody, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Openstack Port",
			"An error occurred while updating the Openstack Port: "+err.Error(),
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

func (r *OpenstackPortResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackPortResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to delete resource
	err := r.client.DeleteByUUID(ctx, "/api/openstack-ports/{uuid}/", data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Openstack Port",
			"An error occurred while deleting the Openstack Port: "+err.Error(),
		)
		return
	}
}

func (r *OpenstackPortResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *OpenstackPortResource) mapResponseToModel(ctx context.Context, apiResp OpenstackPortApiResponse, model *OpenstackPortResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.AdminStateUp = types.BoolPointerValue(apiResp.AdminStateUp)
	listValAllowedAddressPairs, listDiagsAllowedAddressPairs := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"ip_address":  types.StringType,
		"mac_address": types.StringType,
	}}, apiResp.AllowedAddressPairs)
	diags.Append(listDiagsAllowedAddressPairs...)
	model.AllowedAddressPairs = listValAllowedAddressPairs
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.DeviceId = types.StringPointerValue(apiResp.DeviceId)
	model.DeviceOwner = types.StringPointerValue(apiResp.DeviceOwner)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	listValFixedIps, listDiagsFixedIps := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"ip_address": types.StringType,
		"subnet_id":  types.StringType,
	}}, apiResp.FixedIps)
	diags.Append(listDiagsFixedIps...)
	model.FixedIps = listValFixedIps
	model.FloatingIps, _ = types.ListValueFrom(ctx, types.StringType, apiResp.FloatingIps)
	model.MacAddress = types.StringPointerValue(apiResp.MacAddress)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.Network = types.StringPointerValue(apiResp.Network)
	model.NetworkName = types.StringPointerValue(apiResp.NetworkName)
	model.NetworkUuid = types.StringPointerValue(apiResp.NetworkUuid)
	model.PortSecurityEnabled = types.BoolPointerValue(apiResp.PortSecurityEnabled)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	listValSecurityGroups, listDiagsSecurityGroups := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
		"name": types.StringType,
	}}, apiResp.SecurityGroups)
	diags.Append(listDiagsSecurityGroups...)
	model.SecurityGroups = listValSecurityGroups
	model.State = types.StringPointerValue(apiResp.State)
	model.Status = types.StringPointerValue(apiResp.Status)
	model.TargetTenant = types.StringPointerValue(apiResp.TargetTenant)
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantName = types.StringPointerValue(apiResp.TenantName)
	model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
