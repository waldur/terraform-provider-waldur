package bridge_link

type IdentityBridgeLinkCreateRequest struct {
	Address *string `json:"address,omitempty" tfsdk:"address"`

	Affiliations *[]string `json:"affiliations,omitempty" tfsdk:"affiliations"`

	BirthDate *string `json:"birth_date,omitempty" tfsdk:"birth_date"`

	CivilNumber *string `json:"civil_number,omitempty" tfsdk:"civil_number"`

	CountryOfResidence *string `json:"country_of_residence,omitempty" tfsdk:"country_of_residence"`

	EdupersonAssurance *[]string `json:"eduperson_assurance,omitempty" tfsdk:"eduperson_assurance"`

	Email *string `json:"email,omitempty" tfsdk:"email"`

	FirstName *string `json:"first_name,omitempty" tfsdk:"first_name"`

	Gender *string `json:"gender,omitempty" tfsdk:"gender"`

	IdentitySource *string `json:"identity_source,omitempty" tfsdk:"identity_source"`

	LastName *string `json:"last_name,omitempty" tfsdk:"last_name"`

	Nationalities *[]string `json:"nationalities,omitempty" tfsdk:"nationalities"`

	Nationality *string `json:"nationality,omitempty" tfsdk:"nationality"`

	Organization *string `json:"organization,omitempty" tfsdk:"organization"`

	OrganizationCountry *string `json:"organization_country,omitempty" tfsdk:"organization_country"`

	OrganizationType *string `json:"organization_type,omitempty" tfsdk:"organization_type"`

	PersonalTitle *string `json:"personal_title,omitempty" tfsdk:"personal_title"`

	PhoneNumber *string `json:"phone_number,omitempty" tfsdk:"phone_number"`

	PlaceOfBirth *string `json:"place_of_birth,omitempty" tfsdk:"place_of_birth"`

	Source *string `json:"source" tfsdk:"source"`

	Username *string `json:"username" tfsdk:"username"`
}

type IdentityBridgeLinkUpdateRequest struct {
	Address *string `json:"address,omitempty" tfsdk:"address"`

	Affiliations *[]string `json:"affiliations,omitempty" tfsdk:"affiliations"`

	BirthDate *string `json:"birth_date,omitempty" tfsdk:"birth_date"`

	CivilNumber *string `json:"civil_number,omitempty" tfsdk:"civil_number"`

	CountryOfResidence *string `json:"country_of_residence,omitempty" tfsdk:"country_of_residence"`

	EdupersonAssurance *[]string `json:"eduperson_assurance,omitempty" tfsdk:"eduperson_assurance"`

	Email *string `json:"email,omitempty" tfsdk:"email"`

	FirstName *string `json:"first_name,omitempty" tfsdk:"first_name"`

	Gender *string `json:"gender,omitempty" tfsdk:"gender"`

	IdentitySource *string `json:"identity_source,omitempty" tfsdk:"identity_source"`

	LastName *string `json:"last_name,omitempty" tfsdk:"last_name"`

	Nationalities *[]string `json:"nationalities,omitempty" tfsdk:"nationalities"`

	Nationality *string `json:"nationality,omitempty" tfsdk:"nationality"`

	Organization *string `json:"organization,omitempty" tfsdk:"organization"`

	OrganizationCountry *string `json:"organization_country,omitempty" tfsdk:"organization_country"`

	OrganizationType *string `json:"organization_type,omitempty" tfsdk:"organization_type"`

	PersonalTitle *string `json:"personal_title,omitempty" tfsdk:"personal_title"`

	PhoneNumber *string `json:"phone_number,omitempty" tfsdk:"phone_number"`

	PlaceOfBirth *string `json:"place_of_birth,omitempty" tfsdk:"place_of_birth"`

	Source *string `json:"source" tfsdk:"source"`

	Username *string `json:"username" tfsdk:"username"`
}

type IdentityBridgeLinkResponse struct {
	UUID *string `json:"uuid"`

	IsCreated *bool `json:"created" tfsdk:"is_created"`

	UpdatedFields *[]string `json:"updated_fields,omitempty" tfsdk:"updated_fields"`

	UserUuid *string `json:"user_uuid,omitempty" tfsdk:"user_uuid"`
}

func (r *IdentityBridgeLinkResponse) GetState() string {
	return "OK"
}

func (r *IdentityBridgeLinkResponse) GetErrorMessage() string {
	return ""
}
