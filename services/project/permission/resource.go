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
var _ resource.Resource = &ProjectPermissionResource{}
var _ resource.ResourceWithImportState = &ProjectPermissionResource{}

func NewProjectPermissionResource() resource.Resource {
	return &ProjectPermissionResource{}
}

// ProjectPermissionResource defines the resource implementation.
type ProjectPermissionResource struct {
	client *ProjectPermissionClient
}

// ProjectPermissionResourceModel describes the resource data model.
type ProjectPermissionResourceModel struct {
	ProjectPermissionModel
	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

func (r *ProjectPermissionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_permission"
}

func (r *ProjectPermissionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Project Permission resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project Permission UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"expiration_time": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Expiration time of the role (RFC3339 format or null)"},
			"project": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.RequiresReplace(),
				}, MarkdownDescription: "UUID of the project"},
			"role": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.RequiresReplace(),
				}, MarkdownDescription: "Role name (e.g. PROJECT.ADMIN, PROJECT.MANAGER, PROJECT.MEMBER)"},
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

func (r *ProjectPermissionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.client = &ProjectPermissionClient{}
	if err := r.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			err.Error(),
		)
		return
	}
}

type ProjectPermissionMutateRequest struct {
	User           string  `json:"user"`
	Role           string  `json:"role"`
	ExpirationTime *string `json:"expiration_time,omitempty"`
}

type ProjectPermissionDeleteRequest struct {
	User string `json:"user"`
	Role string `json:"role"`
}

func (c *ProjectPermissionClient) AddUser(ctx context.Context, projectUUID string, req *ProjectPermissionMutateRequest) (*ProjectPermissionResponse, error) {
	var resp ProjectPermissionResponse
	err := c.Client.Post(ctx, fmt.Sprintf("/api/projects/%s/add_user/", projectUUID), req, &resp)
	return &resp, err
}

func (c *ProjectPermissionClient) UpdateUser(ctx context.Context, projectUUID string, req *ProjectPermissionMutateRequest) (*ProjectPermissionResponse, error) {
	var resp ProjectPermissionResponse
	err := c.Client.Post(ctx, fmt.Sprintf("/api/projects/%s/update_user/", projectUUID), req, &resp)
	return &resp, err
}

func (c *ProjectPermissionClient) DeleteUser(ctx context.Context, projectUUID string, req *ProjectPermissionDeleteRequest) error {
	return c.Client.Post(ctx, fmt.Sprintf("/api/projects/%s/delete_user/", projectUUID), req, nil)
}

func (c *ProjectPermissionClient) GetUserPermission(ctx context.Context, id string) (*ProjectPermissionResponse, error) {
	var apiResp ProjectPermissionResponse
	err := c.Client.Get(ctx, "/api/user-permissions/{uuid}/", id, &apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}

func (c *ProjectPermissionClient) ListUsers(ctx context.Context, projectUUID string, filter map[string]string) ([]ProjectPermissionResponse, error) {
	var listResult []ProjectPermissionResponse
	err := c.Client.List(ctx, fmt.Sprintf("/api/projects/%s/list_users/", projectUUID), filter, &listResult)
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

func (r *ProjectPermissionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var data ProjectPermissionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqBody := ProjectPermissionMutateRequest{
		User: data.User.ValueString(),
		Role: data.Role.ValueString(),
	}
	if !data.ExpirationTime.IsNull() && !data.ExpirationTime.IsUnknown() {
		reqBody.ExpirationTime = data.ExpirationTime.ValueStringPointer()
	}

	_, err := r.client.AddUser(ctx, data.Project.ValueString(), &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Project Permission",
			"An error occurred while creating the Project Permission: "+err.Error(),
		)
		return
	}

	// Fetch permission UUID via list_users
	filter := map[string]string{
		"user": data.User.ValueString(),
		"role": data.Role.ValueString(),
	}
	usersList, err := r.client.ListUsers(ctx, data.Project.ValueString(), filter)
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

func (r *ProjectPermissionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ProjectPermissionResourceModel

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
			"Unable to Read Project Permission",
			"An error occurred while reading the Project Permission: "+err.Error(),
		)
		return
	}

	// project, user, role are ForceNew and set by config — only refresh expiration_time
	// Normalize +00:00 -> Z so Terraform doesn't see a spurious diff
	data.ExpirationTime = common.StringPointerValue(normalizeUTCTime(apiResp.ExpirationTime))
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProjectPermissionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var data ProjectPermissionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqBody := ProjectPermissionMutateRequest{
		User: data.User.ValueString(),
		Role: data.Role.ValueString(),
	}
	if !data.ExpirationTime.IsNull() && !data.ExpirationTime.IsUnknown() {
		reqBody.ExpirationTime = data.ExpirationTime.ValueStringPointer()
	}

	_, err := r.client.UpdateUser(ctx, data.Project.ValueString(), &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Project Permission",
			"An error occurred while updating the Project Permission: "+err.Error(),
		)
		return
	}

	if data.ExpirationTime.IsUnknown() {
		data.ExpirationTime = types.StringNull()
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProjectPermissionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var data ProjectPermissionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqBody := ProjectPermissionDeleteRequest{
		User: data.User.ValueString(),
		Role: data.Role.ValueString(),
	}
	err := r.client.DeleteUser(ctx, data.Project.ValueString(), &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Project Permission",
			"An error occurred while deleting the Project Permission: "+err.Error(),
		)
		return
	}
}

func (r *ProjectPermissionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resp.Diagnostics.AddError("Import Not Supported", "Importing project permissions is not supported via this resource.")
}
