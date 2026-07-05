package bridge_link

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &IdentityBridgeLinkResource{}
var _ resource.ResourceWithImportState = &IdentityBridgeLinkResource{}

func NewIdentityBridgeLinkResource() resource.Resource {
	return &IdentityBridgeLinkResource{}
}

// IdentityBridgeLinkResource defines the resource implementation.
type IdentityBridgeLinkResource struct {
	client *IdentityBridgeLinkClient
}

// IdentityBridgeLinkResourceModel describes the resource data model.
type IdentityBridgeLinkResourceModel struct {
	IdentityBridgeLinkModel
	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

func (r *IdentityBridgeLinkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_identity_bridge_link"
}

func (r *IdentityBridgeLinkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Identity Bridge Link resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identity Bridge Link UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"address": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Address"},
			"affiliations": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{

					listplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "List of affiliations"},
			"birth_date": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Birth date"},
			"civil_number": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Civil number"},
			"country_of_residence": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Country of residence"},
			"eduperson_assurance": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{

					listplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "List of eduperson assurances"},
			"email": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Email address"},
			"first_name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "First name"},
			"gender": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Gender"},
			"identity_source": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Identity source"},
			"last_name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Last name"},
			"nationalities": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{

					listplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "List of nationalities"},
			"nationality": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Nationality"},
			"organization": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Organization name"},
			"organization_country": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Organization country"},
			"organization_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Organization type"},
			"personal_title": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Personal title"},
			"phone_number": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Phone number"},
			"place_of_birth": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.UseStateForUnknown(),
				}, MarkdownDescription: "Place of birth"},
			"source": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.RequiresReplace(),
				}, MarkdownDescription: "ISD source identifier (must match `^[a-z]+:[a-zA-Z0-9._-]+$`)"},
			"username": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{

					stringplanmodifier.RequiresReplace(),
				}, MarkdownDescription: "CUID / username of the federated user"},
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

func (r *IdentityBridgeLinkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.client = &IdentityBridgeLinkClient{}
	if err := r.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			err.Error(),
		)
		return
	}
}

type IdentityBridgeLinkDeleteRequest struct {
	Username string `json:"username"`
	Source   string `json:"source"`
}

func (c *IdentityBridgeLinkClient) Delete(ctx context.Context, req *IdentityBridgeLinkDeleteRequest) error {
	return c.Client.Post(ctx, "/api/identity-bridge/remove/", req, nil)
}

type IdentityBridgeUser struct {
	UUID       string   `json:"uuid"`
	ActiveIsds []string `json:"active_isds"`
	IsActive   bool     `json:"is_active"`
}

func (r *IdentityBridgeLinkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var data IdentityBridgeLinkResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := IdentityBridgeLinkCreateRequest{
		Username: data.Username.ValueStringPointer(),
		Source:   data.Source.ValueStringPointer(),
	}
	if !data.Address.IsNull() && !data.Address.IsUnknown() {
		requestBody.Address = data.Address.ValueStringPointer()
	}
	if !data.Affiliations.IsNull() && !data.Affiliations.IsUnknown() {
		resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.Affiliations, &requestBody.Affiliations)...)
	}
	if !data.BirthDate.IsNull() && !data.BirthDate.IsUnknown() {
		requestBody.BirthDate = data.BirthDate.ValueStringPointer()
	}
	if !data.CivilNumber.IsNull() && !data.CivilNumber.IsUnknown() {
		requestBody.CivilNumber = data.CivilNumber.ValueStringPointer()
	}
	if !data.CountryOfResidence.IsNull() && !data.CountryOfResidence.IsUnknown() {
		requestBody.CountryOfResidence = data.CountryOfResidence.ValueStringPointer()
	}
	if !data.EdupersonAssurance.IsNull() && !data.EdupersonAssurance.IsUnknown() {
		resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.EdupersonAssurance, &requestBody.EdupersonAssurance)...)
	}
	if !data.Email.IsNull() && !data.Email.IsUnknown() {
		requestBody.Email = data.Email.ValueStringPointer()
	}
	if !data.FirstName.IsNull() && !data.FirstName.IsUnknown() {
		requestBody.FirstName = data.FirstName.ValueStringPointer()
	}
	if !data.Gender.IsNull() && !data.Gender.IsUnknown() {
		requestBody.Gender = data.Gender.ValueStringPointer()
	}
	if !data.IdentitySource.IsNull() && !data.IdentitySource.IsUnknown() {
		requestBody.IdentitySource = data.IdentitySource.ValueStringPointer()
	}
	if !data.LastName.IsNull() && !data.LastName.IsUnknown() {
		requestBody.LastName = data.LastName.ValueStringPointer()
	}
	if !data.Nationalities.IsNull() && !data.Nationalities.IsUnknown() {
		resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.Nationalities, &requestBody.Nationalities)...)
	}
	if !data.Nationality.IsNull() && !data.Nationality.IsUnknown() {
		requestBody.Nationality = data.Nationality.ValueStringPointer()
	}
	if !data.Organization.IsNull() && !data.Organization.IsUnknown() {
		requestBody.Organization = data.Organization.ValueStringPointer()
	}
	if !data.OrganizationCountry.IsNull() && !data.OrganizationCountry.IsUnknown() {
		requestBody.OrganizationCountry = data.OrganizationCountry.ValueStringPointer()
	}
	if !data.OrganizationType.IsNull() && !data.OrganizationType.IsUnknown() {
		requestBody.OrganizationType = data.OrganizationType.ValueStringPointer()
	}
	if !data.PersonalTitle.IsNull() && !data.PersonalTitle.IsUnknown() {
		requestBody.PersonalTitle = data.PersonalTitle.ValueStringPointer()
	}
	if !data.PhoneNumber.IsNull() && !data.PhoneNumber.IsUnknown() {
		requestBody.PhoneNumber = data.PhoneNumber.ValueStringPointer()
	}
	if !data.PlaceOfBirth.IsNull() && !data.PlaceOfBirth.IsUnknown() {
		requestBody.PlaceOfBirth = data.PlaceOfBirth.ValueStringPointer()
	}

	_, err := r.client.Create(ctx, &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Identity Bridge Link",
			"An error occurred while creating/pushing the Identity Bridge Link: "+err.Error(),
		)
		return
	}

	data.UUID = types.StringValue(data.Source.ValueString() + "/" + data.Username.ValueString())
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IdentityBridgeLinkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IdentityBridgeLinkResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource

	// The UUID/ID of the resource is source/username
	idParts := strings.Split(data.UUID.ValueString(), "/")
	if len(idParts) != 2 {
		resp.Diagnostics.AddError("Invalid Resource ID", "Resource ID must be in format source/username")
		return
	}
	source := idParts[0]
	username := idParts[1]

	// 1. Fetch user by username
	var users []IdentityBridgeUser
	err := r.client.Client.GetURL(ctx, "/api/users/?username="+username, &users)
	if err != nil {
		if IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Failed to query user", err.Error())
		return
	}

	if len(users) == 0 {
		resp.State.RemoveResource(ctx)
		return
	}

	user := users[0]

	// 2. Check if the source is active for this user
	sourceActive := false
	if user.ActiveIsds != nil {
		for _, isd := range user.ActiveIsds {
			if isd == source {
				sourceActive = true
				break
			}
		}
	} else {
		// If caller is not staff, they might not see active_isds.
		// We fallback to checking if the user is active at all.
		sourceActive = user.IsActive
	}

	if !sourceActive {
		resp.State.RemoveResource(ctx)
		return
	}

	// Read fields from prior state as Waldur doesn't expose bridge-specific pushed attributes directly on GET /api/users/ in a writable way
	// Keep the existing attributes in state

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IdentityBridgeLinkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Pushing attributes is idempotent, so update does the same push as Create
	var data IdentityBridgeLinkResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := IdentityBridgeLinkCreateRequest{
		Username: data.Username.ValueStringPointer(),
		Source:   data.Source.ValueStringPointer(),
	}
	if !data.Address.IsNull() && !data.Address.IsUnknown() {
		requestBody.Address = data.Address.ValueStringPointer()
	}
	if !data.Affiliations.IsNull() && !data.Affiliations.IsUnknown() {
		resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.Affiliations, &requestBody.Affiliations)...)
	}
	if !data.BirthDate.IsNull() && !data.BirthDate.IsUnknown() {
		requestBody.BirthDate = data.BirthDate.ValueStringPointer()
	}
	if !data.CivilNumber.IsNull() && !data.CivilNumber.IsUnknown() {
		requestBody.CivilNumber = data.CivilNumber.ValueStringPointer()
	}
	if !data.CountryOfResidence.IsNull() && !data.CountryOfResidence.IsUnknown() {
		requestBody.CountryOfResidence = data.CountryOfResidence.ValueStringPointer()
	}
	if !data.EdupersonAssurance.IsNull() && !data.EdupersonAssurance.IsUnknown() {
		resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.EdupersonAssurance, &requestBody.EdupersonAssurance)...)
	}
	if !data.Email.IsNull() && !data.Email.IsUnknown() {
		requestBody.Email = data.Email.ValueStringPointer()
	}
	if !data.FirstName.IsNull() && !data.FirstName.IsUnknown() {
		requestBody.FirstName = data.FirstName.ValueStringPointer()
	}
	if !data.Gender.IsNull() && !data.Gender.IsUnknown() {
		requestBody.Gender = data.Gender.ValueStringPointer()
	}
	if !data.IdentitySource.IsNull() && !data.IdentitySource.IsUnknown() {
		requestBody.IdentitySource = data.IdentitySource.ValueStringPointer()
	}
	if !data.LastName.IsNull() && !data.LastName.IsUnknown() {
		requestBody.LastName = data.LastName.ValueStringPointer()
	}
	if !data.Nationalities.IsNull() && !data.Nationalities.IsUnknown() {
		resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.Nationalities, &requestBody.Nationalities)...)
	}
	if !data.Nationality.IsNull() && !data.Nationality.IsUnknown() {
		requestBody.Nationality = data.Nationality.ValueStringPointer()
	}
	if !data.Organization.IsNull() && !data.Organization.IsUnknown() {
		requestBody.Organization = data.Organization.ValueStringPointer()
	}
	if !data.OrganizationCountry.IsNull() && !data.OrganizationCountry.IsUnknown() {
		requestBody.OrganizationCountry = data.OrganizationCountry.ValueStringPointer()
	}
	if !data.OrganizationType.IsNull() && !data.OrganizationType.IsUnknown() {
		requestBody.OrganizationType = data.OrganizationType.ValueStringPointer()
	}
	if !data.PersonalTitle.IsNull() && !data.PersonalTitle.IsUnknown() {
		requestBody.PersonalTitle = data.PersonalTitle.ValueStringPointer()
	}
	if !data.PhoneNumber.IsNull() && !data.PhoneNumber.IsUnknown() {
		requestBody.PhoneNumber = data.PhoneNumber.ValueStringPointer()
	}
	if !data.PlaceOfBirth.IsNull() && !data.PlaceOfBirth.IsUnknown() {
		requestBody.PlaceOfBirth = data.PlaceOfBirth.ValueStringPointer()
	}

	_, err := r.client.Create(ctx, &requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Identity Bridge Link",
			"An error occurred while updating/pushing the Identity Bridge Link: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IdentityBridgeLinkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var data IdentityBridgeLinkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	idParts := strings.Split(data.UUID.ValueString(), "/")
	if len(idParts) != 2 {
		resp.Diagnostics.AddError("Invalid Resource ID", "Resource ID must be in format source/username")
		return
	}
	source := idParts[0]
	username := idParts[1]

	reqBody := IdentityBridgeLinkDeleteRequest{
		Username: username,
		Source:   source,
	}

	err := r.client.Delete(ctx, &reqBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Identity Bridge Link",
			"An error occurred while deleting/removing the Identity Bridge Link: "+err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *IdentityBridgeLinkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	var data IdentityBridgeLinkResourceModel
	data.UUID = types.StringValue(req.ID)

	idParts := strings.Split(req.ID, "/")
	if len(idParts) != 2 {
		resp.Diagnostics.AddError("Invalid Resource ID", "Import ID must be in format source/username")
		return
	}
	data.Source = types.StringValue(idParts[0])
	data.Username = types.StringValue(idParts[1])

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
