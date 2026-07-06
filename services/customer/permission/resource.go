package permission

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CustomerPermissionResource{}
var _ resource.ResourceWithImportState = &CustomerPermissionResource{}

func NewCustomerPermissionResource() resource.Resource {
	return &CustomerPermissionResource{}
}

// CustomerPermissionResource defines the resource implementation.
type CustomerPermissionResource struct {
	client *CustomerPermissionClient
}

// CustomerPermissionResourceModel describes the resource data model.
type CustomerPermissionResourceModel struct {
	CustomerPermissionModel
	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

func (r *CustomerPermissionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_customer_permission"
}

func (r *CustomerPermissionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Customer Permission resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Customer Permission UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"customer": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.RequiresReplace(),
				}, MarkdownDescription: "UUID of the customer"},
			"expiration_time": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Expiration time of the role (RFC3339 format or null)"},
			"role": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.RequiresReplace(),
				}, MarkdownDescription: "Role name (e.g. PROJECT.ADMIN, CUSTOMER.OWNER)"},
			"user": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.RequiresReplace(),
				}, MarkdownDescription: "UUID of the user"},
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

func (r *CustomerPermissionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.client = &CustomerPermissionClient{}
	if err := r.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			err.Error(),
		)
		return
	}
}

type CustomerPermissionMutateRequest struct {
	User           string  `json:"user"`
	Role           string  `json:"role"`
	ExpirationTime *string `json:"expiration_time,omitempty"`
}

type CustomerPermissionDeleteRequest struct {
	User string `json:"user"`
	Role string `json:"role"`
}

func (c *CustomerPermissionClient) AddUser(ctx context.Context, scopeUUID string, req *CustomerPermissionMutateRequest) (*CustomerPermissionResponse, error) {
	var resp CustomerPermissionResponse
	err := c.Client.Post(ctx, fmt.Sprintf("/api/customers/%s/add_user/", scopeUUID), req, &resp)
	return &resp, err
}

func (c *CustomerPermissionClient) UpdateUser(ctx context.Context, scopeUUID string, req *CustomerPermissionMutateRequest) (*CustomerPermissionResponse, error) {
	var resp CustomerPermissionResponse
	err := c.Client.Post(ctx, fmt.Sprintf("/api/customers/%s/update_user/", scopeUUID), req, &resp)
	return &resp, err
}

func (c *CustomerPermissionClient) DeleteUser(ctx context.Context, scopeUUID string, req *CustomerPermissionDeleteRequest) error {
	return c.Client.Post(ctx, fmt.Sprintf("/api/customers/%s/delete_user/", scopeUUID), req, nil)
}

func (c *CustomerPermissionClient) GetUserPermission(ctx context.Context, id string) (*CustomerPermissionResponse, error) {
	var apiResp CustomerPermissionResponse
	err := c.Client.Get(ctx, "/api/user-permissions/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *CustomerPermissionClient) ListUsers(ctx context.Context, scopeUUID string, filter map[string]string) ([]CustomerPermissionResponse, error) {
	var listResult []CustomerPermissionResponse
	err := c.Client.List(ctx, fmt.Sprintf("/api/customers/%s/list_users/", scopeUUID), filter, &listResult)
	if err != nil {
		return nil, err
	}
	return listResult, nil
}

// normalizeUTCTime converts +00:00 timezone suffix to Z for consistent comparison.
func normalizeUTCTime(s *string) *string {
	if s == nil {
		return nil
	}
	v := strings.Replace(*s, "+00:00", "Z", 1)
	return &v
}

func (r *CustomerPermissionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var data CustomerPermissionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqBody := CustomerPermissionMutateRequest{
		User: data.User.ValueString(),
		Role: data.Role.ValueString(),
	}
	if !data.ExpirationTime.IsNull() && !data.ExpirationTime.IsUnknown() {
		reqBody.ExpirationTime = data.ExpirationTime.ValueStringPointer()
	}

	_, err := r.client.AddUser(ctx, data.Customer.ValueString(), &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Customer Permission",
			"An error occurred while creating the Customer Permission: "+err.Error(),
		)
		return
	}

	// Fetch permission UUID via list_users
	filter := map[string]string{
		"user": data.User.ValueString(),
		"role": data.Role.ValueString(),
	}
	usersList, err := r.client.ListUsers(ctx, data.Customer.ValueString(), filter)
	if err != nil || len(usersList) == 0 {
		errMsg := "unknown error"
		if err != nil {
			errMsg = err.Error()
		}
		resp.Diagnostics.AddError(
			"Unable to Find Permission UUID",
			"The permission was successfully created but could not be located via list_users: "+errMsg,
		)
		return
	}

	data.UUID = types.StringPointerValue(usersList[0].UUID)
	if data.ExpirationTime.IsUnknown() {
		data.ExpirationTime = types.StringNull()
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CustomerPermissionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CustomerPermissionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource

	apiResp, err := r.client.GetUserPermission(ctx, data.UUID.ValueString())
	if err != nil {
		if IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to Read Customer Permission",
			"An error occurred while reading the Customer Permission: "+err.Error(),
		)
		return
	}

	// scope, user, role are ForceNew and set by config — only refresh expiration_time
	// Normalize +00:00 -> Z so Terraform doesn't see a spurious diff
	data.ExpirationTime = common.StringPointerValue(normalizeUTCTime(apiResp.ExpirationTime))
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CustomerPermissionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var data CustomerPermissionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqBody := CustomerPermissionMutateRequest{
		User: data.User.ValueString(),
		Role: data.Role.ValueString(),
	}
	if !data.ExpirationTime.IsNull() && !data.ExpirationTime.IsUnknown() {
		reqBody.ExpirationTime = data.ExpirationTime.ValueStringPointer()
	}

	_, err := r.client.UpdateUser(ctx, data.Customer.ValueString(), &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Customer Permission",
			"An error occurred while updating the Customer Permission: "+err.Error(),
		)
		return
	}

	if data.ExpirationTime.IsUnknown() {
		data.ExpirationTime = types.StringNull()
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CustomerPermissionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var data CustomerPermissionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqBody := CustomerPermissionDeleteRequest{
		User: data.User.ValueString(),
		Role: data.Role.ValueString(),
	}
	err := r.client.DeleteUser(ctx, data.Customer.ValueString(), &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Customer Permission",
			"An error occurred while deleting the Customer Permission: "+err.Error(),
		)
		return
	}
}

func (r *CustomerPermissionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resp.Diagnostics.AddError("Import Not Supported", "Importing Customer Permission permissions is not supported via this resource.")
}
