package ssh_public_key

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CoreSshPublicKeyResource{}
var _ resource.ResourceWithImportState = &CoreSshPublicKeyResource{}

func NewCoreSshPublicKeyResource() resource.Resource {
	return &CoreSshPublicKeyResource{}
}

// CoreSshPublicKeyResource defines the resource implementation.
type CoreSshPublicKeyResource struct {
	client *CoreSshPublicKeyClient
}

// CoreSshPublicKeyResourceModel describes the resource data model.
type CoreSshPublicKeyResourceModel struct {
	CoreSshPublicKeyModel
	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

func (r *CoreSshPublicKeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_core_ssh_public_key"
}

func (r *CoreSshPublicKeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Core Ssh Public Key resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Core Ssh Public Key UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"fingerprint_md5": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Fingerprint Md5"},
			"fingerprint_sha256": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Fingerprint Sha256"},
			"fingerprint_sha512": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Fingerprint Sha512"},
			"is_shared": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{

					boolplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Is Shared"},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Name"},
			"public_key": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.RequiresReplace(),
				}, MarkdownDescription: "Public Key"},
			"type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Type"},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Url"},
			"user_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "User Uuid"},
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

func (r *CoreSshPublicKeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.client = &CoreSshPublicKeyClient{}
	if err := r.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			err.Error(),
		)
		return
	}
}

func (r *CoreSshPublicKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var data CoreSshPublicKeyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := CoreSshPublicKeyCreateRequest{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {

		requestBody.Name = data.Name.ValueStringPointer()
	}

	requestBody.PublicKey = data.PublicKey.ValueStringPointer()

	apiResp, err := r.client.Create(ctx, &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Core Ssh Public Key",
			"An error occurred while creating the Core Ssh Public Key: "+err.Error(),
		)
		return
	}
	data.UUID = types.StringPointerValue(apiResp.UUID)

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CoreSshPublicKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CoreSshPublicKeyResourceModel

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
			"Unable to Read Core Ssh Public Key",
			"An error occurred while reading the Core Ssh Public Key: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CoreSshPublicKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	resp.Diagnostics.AddError("Update Not Supported", "This resource cannot be updated via the API.")
}

func (r *CoreSshPublicKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var data CoreSshPublicKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.client.Delete(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Core Ssh Public Key",
			"An error occurred while deleting the Core Ssh Public Key: "+err.Error(),
		)
		return
	}
}

func (r *CoreSshPublicKeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	uuid := req.ID
	if uuid == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Import ID cannot be empty. Please provide the UUID of the Core Ssh Public Key.",
		)
		return
	}

	tflog.Info(ctx, "Importing Core Ssh Public Key", map[string]interface{}{
		"uuid": uuid,
	})

	apiResp, err := r.client.Get(ctx, uuid)
	if err != nil {
		if IsNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("Core Ssh Public Key with UUID '%s' does not exist or is not accessible.", uuid),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to Import Core Ssh Public Key",
			fmt.Sprintf("An error occurred while fetching the Core Ssh Public Key: %s", err.Error()),
		)
		return
	}

	var data CoreSshPublicKeyResourceModel
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
