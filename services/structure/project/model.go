package project

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StructureProjectFiltersModel struct {
	BackendId               types.String `tfsdk:"backend_id"`
	CanAdmin                types.Bool   `tfsdk:"can_admin"`
	CanManage               types.Bool   `tfsdk:"can_manage"`
	ConcealFinishedProjects types.Bool   `tfsdk:"conceal_finished_projects"`
	Created                 types.String `tfsdk:"created"`
	CustomerAbbreviation    types.String `tfsdk:"customer_abbreviation"`
	CustomerName            types.String `tfsdk:"customer_name"`
	CustomerNativeName      types.String `tfsdk:"customer_native_name"`
	Description             types.String `tfsdk:"description"`
	IncludeTerminated       types.Bool   `tfsdk:"include_terminated"`
	IsRemoved               types.Bool   `tfsdk:"is_removed"`
	Modified                types.String `tfsdk:"modified"`
	Name                    types.String `tfsdk:"name"`
	NameExact               types.String `tfsdk:"name_exact"`
	Query                   types.String `tfsdk:"query"`
	Slug                    types.String `tfsdk:"slug"`
}

type StructureProjectModel struct {
	UUID                                 types.String      `tfsdk:"id"`
	BackendId                            types.String      `tfsdk:"backend_id"`
	Created                              timetypes.RFC3339 `tfsdk:"created"`
	Customer                             types.String      `tfsdk:"customer"`
	CustomerAbbreviation                 types.String      `tfsdk:"customer_abbreviation"`
	CustomerDisplayBillingInfoInProjects types.Bool        `tfsdk:"customer_display_billing_info_in_projects"`
	CustomerName                         types.String      `tfsdk:"customer_name"`
	CustomerNativeName                   types.String      `tfsdk:"customer_native_name"`
	CustomerSlug                         types.String      `tfsdk:"customer_slug"`
	CustomerUuid                         types.String      `tfsdk:"customer_uuid"`
	Description                          types.String      `tfsdk:"description"`
	EndDate                              types.String      `tfsdk:"end_date"`
	EndDateRequestedBy                   types.String      `tfsdk:"end_date_requested_by"`
	GracePeriodDays                      types.Int64       `tfsdk:"grace_period_days"`
	Image                                types.String      `tfsdk:"image"`
	IsIndustry                           types.Bool        `tfsdk:"is_industry"`
	IsRemoved                            types.Bool        `tfsdk:"is_removed"`
	Kind                                 types.String      `tfsdk:"kind"`
	MaxServiceAccounts                   types.Int64       `tfsdk:"max_service_accounts"`
	Name                                 types.String      `tfsdk:"name"`
	OecdFos2007Code                      types.String      `tfsdk:"oecd_fos_2007_code"`
	OecdFos2007Label                     types.String      `tfsdk:"oecd_fos_2007_label"`
	ProjectCredit                        types.Float64     `tfsdk:"project_credit"`
	ResourcesCount                       types.Int64       `tfsdk:"resources_count"`
	Slug                                 types.String      `tfsdk:"slug"`
	StaffNotes                           types.String      `tfsdk:"staff_notes"`
	StartDate                            types.String      `tfsdk:"start_date"`
	Type                                 types.String      `tfsdk:"type"`
	TypeName                             types.String      `tfsdk:"type_name"`
	TypeUuid                             types.String      `tfsdk:"type_uuid"`
	Url                                  types.String      `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *StructureProjectModel) CopyFrom(ctx context.Context, apiResp StructureProjectResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	valCreated, diagsCreated := timetypes.NewRFC3339PointerValue(apiResp.Created)
	diags.Append(diagsCreated...)
	model.Created = valCreated
	model.Customer = types.StringPointerValue(apiResp.Customer)
	model.CustomerDisplayBillingInfoInProjects = types.BoolPointerValue(apiResp.CustomerDisplayBillingInfoInProjects)
	model.CustomerSlug = types.StringPointerValue(apiResp.CustomerSlug)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.EndDate = types.StringPointerValue(apiResp.EndDate)
	model.EndDateRequestedBy = types.StringPointerValue(apiResp.EndDateRequestedBy)
	model.GracePeriodDays = types.Int64PointerValue(apiResp.GracePeriodDays)
	model.Image = types.StringPointerValue(apiResp.Image)
	model.IsIndustry = types.BoolPointerValue(apiResp.IsIndustry)
	model.IsRemoved = types.BoolPointerValue(apiResp.IsRemoved)
	model.Kind = types.StringPointerValue(apiResp.Kind)
	model.MaxServiceAccounts = types.Int64PointerValue(apiResp.MaxServiceAccounts)
	model.Name = types.StringPointerValue(apiResp.Name)
	model.OecdFos2007Code = types.StringPointerValue(apiResp.OecdFos2007Code)
	model.OecdFos2007Label = types.StringPointerValue(apiResp.OecdFos2007Label)
	model.ProjectCredit = types.Float64PointerValue(apiResp.ProjectCredit)
	model.ResourcesCount = types.Int64PointerValue(apiResp.ResourcesCount)
	model.Slug = types.StringPointerValue(apiResp.Slug)
	model.StaffNotes = types.StringPointerValue(apiResp.StaffNotes)
	model.StartDate = types.StringPointerValue(apiResp.StartDate)
	model.Type = types.StringPointerValue(apiResp.Type)
	model.TypeName = types.StringPointerValue(apiResp.TypeName)
	model.TypeUuid = types.StringPointerValue(apiResp.TypeUuid)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
