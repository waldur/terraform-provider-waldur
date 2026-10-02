package common

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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

// SetElementsFromState keeps the prior state of set elements that did not change.
//
// Attribute-level UseStateForUnknown does not work inside a set: the framework
// cannot tell which prior element a planned element corresponds to, so it pairs
// them by position and gives one element another's computed values. This
// modifier matches each planned element to the prior element whose Keys (the
// element's configurable attributes) hold the same values, and plans that prior
// element unchanged. An element with no match keeps its unknown computed values,
// which is what a real change needs.
type SetElementsFromState struct {
	Keys []string
}

func (m SetElementsFromState) Description(ctx context.Context) string {
	return "Keeps the prior state of set elements whose configurable attributes are unchanged."
}

func (m SetElementsFromState) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m SetElementsFromState) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() || req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}
	prior := req.StateValue.Elements()
	used := make([]bool, len(prior))
	planned := req.PlanValue.Elements()
	out := make([]attr.Value, len(planned))
	changed := false
	for i, el := range planned {
		out[i] = el
		obj, ok := el.(types.Object)
		if !ok || obj.IsNull() || obj.IsUnknown() {
			continue
		}
		for j, candidate := range prior {
			if used[j] {
				continue
			}
			priorObj, ok := candidate.(types.Object)
			if !ok || !m.sameKeys(obj, priorObj) {
				continue
			}
			used[j] = true
			out[i] = candidate
			changed = true
			break
		}
	}
	if !changed {
		return
	}
	v, diags := types.SetValue(req.PlanValue.ElementType(ctx), out)
	resp.Diagnostics.Append(diags...)
	if !diags.HasError() {
		resp.PlanValue = v
	}
}

// sameKeys reports whether every key known in the plan equals the prior value.
// At least one key has to be known, or there is nothing to match on.
func (m SetElementsFromState) sameKeys(plan, prior types.Object) bool {
	planAttrs, priorAttrs := plan.Attributes(), prior.Attributes()
	compared := 0
	for _, key := range m.Keys {
		pv, ok := planAttrs[key]
		if !ok || pv.IsUnknown() {
			continue
		}
		sv, ok := priorAttrs[key]
		if !ok || !pv.Equal(sv) {
			return false
		}
		compared++
	}
	return compared > 0
}
