package common

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// UnknownIfNullModifier implements plan modification to set Null values to Unknown.
type UnknownIfNullModifier struct{}

func (m UnknownIfNullModifier) Description(ctx context.Context) string {
	return "Sets the attribute to Unknown if it is Null in the plan."
}

func (m UnknownIfNullModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m UnknownIfNullModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.PlanValue.IsNull() && !req.ConfigValue.IsUnknown() {
		resp.PlanValue = types.StringUnknown()
	}
}

func (m UnknownIfNullModifier) PlanModifyInt64(ctx context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	if req.PlanValue.IsNull() && !req.ConfigValue.IsUnknown() {
		resp.PlanValue = types.Int64Unknown()
	}
}

func (m UnknownIfNullModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	if req.PlanValue.IsNull() && !req.ConfigValue.IsUnknown() {
		resp.PlanValue = types.BoolUnknown()
	}
}

func (m UnknownIfNullModifier) PlanModifyFloat64(ctx context.Context, req planmodifier.Float64Request, resp *planmodifier.Float64Response) {
	if req.PlanValue.IsNull() && !req.ConfigValue.IsUnknown() {
		resp.PlanValue = types.Float64Unknown()
	}
}
