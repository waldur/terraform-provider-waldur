package project

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

func BillingPriceEstimateType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"current":     types.Float64Type,
		"tax":         types.Float64Type,
		"tax_current": types.Float64Type,
		"total":       types.Float64Type,
	}}
}

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
				Optional: true,
			},
			"can_admin": schema.BoolAttribute{
				Optional: true,
			},
			"can_manage": schema.BoolAttribute{
				Optional: true,
			},
			"conceal_finished_projects": schema.BoolAttribute{
				Optional: true,
			},
			"created": schema.StringAttribute{
				Optional: true,
			},
			"customer_abbreviation": schema.StringAttribute{
				Optional: true,
			},
			"customer_name": schema.StringAttribute{
				Optional: true,
			},
			"customer_native_name": schema.StringAttribute{
				Optional: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"include_terminated": schema.BoolAttribute{
				Optional: true,
			},
			"is_removed": schema.BoolAttribute{
				Optional: true,
			},
			"modified": schema.StringAttribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"name_exact": schema.StringAttribute{
				Optional: true,
			},
			"query": schema.StringAttribute{
				Optional: true,
			},
			"slug": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

type StructureProjectModel struct {
	UUID                                 types.String  `tfsdk:"id"`
	BackendId                            types.String  `tfsdk:"backend_id"`
	BillingPriceEstimate                 types.Object  `tfsdk:"billing_price_estimate"`
	Customer                             types.String  `tfsdk:"customer"`
	CustomerDisplayBillingInfoInProjects types.Bool    `tfsdk:"customer_display_billing_info_in_projects"`
	CustomerSlug                         types.String  `tfsdk:"customer_slug"`
	Description                          types.String  `tfsdk:"description"`
	EndDate                              types.String  `tfsdk:"end_date"`
	EndDateRequestedBy                   types.String  `tfsdk:"end_date_requested_by"`
	GracePeriodDays                      types.Int64   `tfsdk:"grace_period_days"`
	Image                                types.String  `tfsdk:"image"`
	IsIndustry                           types.Bool    `tfsdk:"is_industry"`
	IsRemoved                            types.Bool    `tfsdk:"is_removed"`
	Kind                                 types.String  `tfsdk:"kind"`
	MaxServiceAccounts                   types.Int64   `tfsdk:"max_service_accounts"`
	Name                                 types.String  `tfsdk:"name"`
	OecdFos2007Code                      types.String  `tfsdk:"oecd_fos_2007_code"`
	OecdFos2007Label                     types.String  `tfsdk:"oecd_fos_2007_label"`
	ProjectCredit                        types.Float64 `tfsdk:"project_credit"`
	ResourcesCount                       types.Int64   `tfsdk:"resources_count"`
	Slug                                 types.String  `tfsdk:"slug"`
	StaffNotes                           types.String  `tfsdk:"staff_notes"`
	StartDate                            types.String  `tfsdk:"start_date"`
	Type                                 types.String  `tfsdk:"type"`
	TypeName                             types.String  `tfsdk:"type_name"`
	TypeUuid                             types.String  `tfsdk:"type_uuid"`
	Url                                  types.String  `tfsdk:"url"`
}

// CopyFrom maps the API response to the model fields.
func (model *StructureProjectModel) CopyFrom(ctx context.Context, apiResp StructureProjectResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.BackendId = common.StringPointerValue(apiResp.BackendId)
	if apiResp.BillingPriceEstimate != nil {
		objValBillingPriceEstimate, objDiagsBillingPriceEstimate := types.ObjectValueFrom(ctx, BillingPriceEstimateType().AttrTypes, *apiResp.BillingPriceEstimate)
		diags.Append(objDiagsBillingPriceEstimate...)
		model.BillingPriceEstimate = objValBillingPriceEstimate
	} else {
		model.BillingPriceEstimate = types.ObjectNull(BillingPriceEstimateType().AttrTypes)
	}
	model.Customer = common.StringPointerValue(apiResp.Customer)
	model.CustomerDisplayBillingInfoInProjects = types.BoolPointerValue(apiResp.CustomerDisplayBillingInfoInProjects)
	model.CustomerSlug = common.StringPointerValue(apiResp.CustomerSlug)
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
