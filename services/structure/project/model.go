package project

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
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

func (m *StructureProjectFiltersModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Filter parameters for querying Structure Project",
		Attributes: map[string]schema.Attribute{
			"backend_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "ID of the backend",
			},
			"can_admin": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Return a list of projects where current user is admin.",
			},
			"can_manage": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Return a list of projects where current user is manager or a customer owner.",
			},
			"conceal_finished_projects": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Conceal finished projects",
			},
			"created": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Created after",
			},
			"customer_abbreviation": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer abbreviation",
			},
			"customer_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer name",
			},
			"customer_native_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Customer native name",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description",
			},
			"include_terminated": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Include soft-deleted (terminated) projects. Only available to staff and support users, or users with organizational roles who can see their terminated projects.",
			},
			"is_removed": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Is removed",
			},
			"modified": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Modified after",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name (exact)",
			},
			"query": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Filter by name, slug, UUID, backend ID or resource effective ID",
			},
			"slug": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Slug",
			},
		},
	}
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
	model.BackendId = common.StringPointerValue(apiResp.BackendId)
	valCreated, diagsCreated := timetypes.NewRFC3339PointerValue(apiResp.Created)
	diags.Append(diagsCreated...)
	model.Created = valCreated
	model.Customer = common.StringPointerValue(apiResp.Customer)
	model.CustomerAbbreviation = common.StringPointerValue(apiResp.CustomerAbbreviation)
	model.CustomerDisplayBillingInfoInProjects = types.BoolPointerValue(apiResp.CustomerDisplayBillingInfoInProjects)
	model.CustomerName = common.StringPointerValue(apiResp.CustomerName)
	model.CustomerNativeName = common.StringPointerValue(apiResp.CustomerNativeName)
	model.CustomerSlug = common.StringPointerValue(apiResp.CustomerSlug)
	model.CustomerUuid = common.StringPointerValue(apiResp.CustomerUuid)
	model.Description = common.StringPointerValue(apiResp.Description)
	model.EndDate = common.StringPointerValue(apiResp.EndDate)
	model.EndDateRequestedBy = common.StringPointerValue(apiResp.EndDateRequestedBy)
	model.GracePeriodDays = types.Int64PointerValue(apiResp.GracePeriodDays)
	model.Image = common.StringPointerValue(apiResp.Image)
	model.IsIndustry = types.BoolPointerValue(apiResp.IsIndustry)
	model.IsRemoved = types.BoolPointerValue(apiResp.IsRemoved)
	model.Kind = common.StringPointerValue(apiResp.Kind)
	model.MaxServiceAccounts = types.Int64PointerValue(apiResp.MaxServiceAccounts)
	model.Name = common.StringPointerValue(apiResp.Name)
	model.OecdFos2007Code = common.StringPointerValue(apiResp.OecdFos2007Code)
	model.OecdFos2007Label = common.StringPointerValue(apiResp.OecdFos2007Label)
	model.ProjectCredit = types.Float64PointerValue(apiResp.ProjectCredit.Float64Ptr())
	model.ResourcesCount = types.Int64PointerValue(apiResp.ResourcesCount)
	model.Slug = common.StringPointerValue(apiResp.Slug)
	model.StaffNotes = common.StringPointerValue(apiResp.StaffNotes)
	model.StartDate = common.StringPointerValue(apiResp.StartDate)
	model.Type = common.StringPointerValue(apiResp.Type)
	model.TypeName = common.StringPointerValue(apiResp.TypeName)
	model.TypeUuid = common.StringPointerValue(apiResp.TypeUuid)
	model.Url = common.StringPointerValue(apiResp.Url)

	return diags
}
