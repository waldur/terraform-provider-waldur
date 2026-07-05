package bridge_link

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type IdentityBridgeLinkModel struct {
	UUID                types.String `tfsdk:"id"`
	Address             types.String `tfsdk:"address"`
	Affiliations        types.List   `tfsdk:"affiliations"`
	BirthDate           types.String `tfsdk:"birth_date"`
	CivilNumber         types.String `tfsdk:"civil_number"`
	CountryOfResidence  types.String `tfsdk:"country_of_residence"`
	EdupersonAssurance  types.List   `tfsdk:"eduperson_assurance"`
	Email               types.String `tfsdk:"email"`
	FirstName           types.String `tfsdk:"first_name"`
	Gender              types.String `tfsdk:"gender"`
	IdentitySource      types.String `tfsdk:"identity_source"`
	LastName            types.String `tfsdk:"last_name"`
	Nationalities       types.List   `tfsdk:"nationalities"`
	Nationality         types.String `tfsdk:"nationality"`
	Organization        types.String `tfsdk:"organization"`
	OrganizationCountry types.String `tfsdk:"organization_country"`
	OrganizationType    types.String `tfsdk:"organization_type"`
	PersonalTitle       types.String `tfsdk:"personal_title"`
	PhoneNumber         types.String `tfsdk:"phone_number"`
	PlaceOfBirth        types.String `tfsdk:"place_of_birth"`
	Source              types.String `tfsdk:"source"`
	Username            types.String `tfsdk:"username"`
}

// CopyFrom maps the API response to the model fields.
func (model *IdentityBridgeLinkModel) CopyFrom(ctx context.Context, apiResp IdentityBridgeLinkResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)

	model.Address = common.StringPointerValue(apiResp.Address)

	if apiResp.Affiliations != nil {
		valAffiliations, diagsAffiliations := types.ListValueFrom(ctx, types.StringType, apiResp.Affiliations)
		diags.Append(diagsAffiliations...)
		model.Affiliations = valAffiliations
	} else {
		model.Affiliations = types.ListNull(types.StringType)
	}

	model.BirthDate = common.StringPointerValue(apiResp.BirthDate)

	model.CivilNumber = common.StringPointerValue(apiResp.CivilNumber)

	model.CountryOfResidence = common.StringPointerValue(apiResp.CountryOfResidence)

	if apiResp.EdupersonAssurance != nil {
		valEdupersonAssurance, diagsEdupersonAssurance := types.ListValueFrom(ctx, types.StringType, apiResp.EdupersonAssurance)
		diags.Append(diagsEdupersonAssurance...)
		model.EdupersonAssurance = valEdupersonAssurance
	} else {
		model.EdupersonAssurance = types.ListNull(types.StringType)
	}

	model.Email = common.StringPointerValue(apiResp.Email)

	model.FirstName = common.StringPointerValue(apiResp.FirstName)

	model.Gender = common.StringPointerValue(apiResp.Gender)

	model.IdentitySource = common.StringPointerValue(apiResp.IdentitySource)

	model.LastName = common.StringPointerValue(apiResp.LastName)

	if apiResp.Nationalities != nil {
		valNationalities, diagsNationalities := types.ListValueFrom(ctx, types.StringType, apiResp.Nationalities)
		diags.Append(diagsNationalities...)
		model.Nationalities = valNationalities
	} else {
		model.Nationalities = types.ListNull(types.StringType)
	}

	model.Nationality = common.StringPointerValue(apiResp.Nationality)

	model.Organization = common.StringPointerValue(apiResp.Organization)

	model.OrganizationCountry = common.StringPointerValue(apiResp.OrganizationCountry)

	model.OrganizationType = common.StringPointerValue(apiResp.OrganizationType)

	model.PersonalTitle = common.StringPointerValue(apiResp.PersonalTitle)

	model.PhoneNumber = common.StringPointerValue(apiResp.PhoneNumber)

	model.PlaceOfBirth = common.StringPointerValue(apiResp.PlaceOfBirth)

	model.Source = common.StringPointerValue(apiResp.Source)

	model.Username = common.StringPointerValue(apiResp.Username)

	return diags
}
