package common

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// sanitizeObject replaces Unknown attributes with Null attributes to allow conversion to *T fields.
func sanitizeObject(ctx context.Context, obj types.Object) (types.Object, diag.Diagnostics) {
	attrs := obj.Attributes()
	newAttrs := make(map[string]attr.Value, len(attrs))
	for k, v := range attrs {
		if v.IsUnknown() {
			tfType := v.Type(ctx).TerraformType(ctx)
			val, err := v.Type(ctx).ValueFromTerraform(ctx, tftypes.NewValue(tfType, nil))
			if err != nil {
				return types.Object{}, diag.Diagnostics{diag.NewErrorDiagnostic("Sanitization Error", err.Error())}
			}
			newAttrs[k] = val
		} else {
			newAttrs[k] = v
		}
	}
	return types.ObjectValue(obj.AttributeTypes(ctx), newAttrs)
}

// PopulateSliceField helps populating a slice field from a Terraform list.
func PopulateSliceField[T any](ctx context.Context, list types.List, target *[]T) diag.Diagnostics {
	if list.IsNull() || list.IsUnknown() {
		return nil
	}

	// For lists of objects, use manual iteration and sanitization
	if _, ok := list.ElementType(ctx).(types.ObjectType); ok {
		items := make([]T, 0)
		var diags diag.Diagnostics
		for _, el := range list.Elements() {
			if obj, ok := el.(types.Object); ok {
				cleanObj, d := sanitizeObject(ctx, obj)
				diags.Append(d...)
				if d.HasError() {
					continue
				}

				var item T
				d = cleanObj.As(ctx, &item, basetypes.ObjectAsOptions{})
				diags.Append(d...)
				if !d.HasError() {
					items = append(items, item)
				}
			} else {
				diags.AddError("Conversion Error", "Expected Object element")
			}
		}
		if !diags.HasError() {
			*target = items
		}
		return diags
	}

	items := make([]T, 0)
	diags := list.ElementsAs(ctx, &items, false)
	if !diags.HasError() {
		*target = items
	}
	return diags
}

// PopulateSetField helps populating a slice field from a Terraform set.
func PopulateSetField[T any](ctx context.Context, set types.Set, target *[]T) diag.Diagnostics {
	if set.IsNull() || set.IsUnknown() {
		return nil
	}

	// For sets of objects
	if _, ok := set.ElementType(ctx).(types.ObjectType); ok {
		items := make([]T, 0)
		var diags diag.Diagnostics
		for _, el := range set.Elements() {
			if obj, ok := el.(types.Object); ok {
				cleanObj, d := sanitizeObject(ctx, obj)
				diags.Append(d...)
				if d.HasError() {
					continue
				}

				var item T
				d = cleanObj.As(ctx, &item, basetypes.ObjectAsOptions{})
				diags.Append(d...)
				if !d.HasError() {
					items = append(items, item)
				}
			} else {
				diags.AddError("Conversion Error", "Expected Object element")
			}
		}
		if !diags.HasError() {
			*target = items
		}
		return diags
	}

	items := make([]T, 0)
	diags := set.ElementsAs(ctx, &items, false)
	if !diags.HasError() {
		*target = items
	}
	return diags
}

// PopulateOptionalSliceField populates an optional slice field (**[]T) from a Terraform list.
func PopulateOptionalSliceField[T any](ctx context.Context, list types.List, target **[]T) diag.Diagnostics {
	if list.IsNull() || list.IsUnknown() {
		return nil
	}

	if _, ok := list.ElementType(ctx).(types.ObjectType); ok {
		items := make([]T, 0)
		var diags diag.Diagnostics
		for _, el := range list.Elements() {
			if obj, ok := el.(types.Object); ok {
				cleanObj, d := sanitizeObject(ctx, obj)
				diags.Append(d...)
				if d.HasError() {
					continue
				}

				var item T
				d = cleanObj.As(ctx, &item, basetypes.ObjectAsOptions{})
				diags.Append(d...)
				if !d.HasError() {
					items = append(items, item)
				}
			} else {
				diags.AddError("Conversion Error", "Expected Object element")
			}
		}
		if !diags.HasError() {
			*target = &items
		}
		return diags
	}

	items := make([]T, 0)
	diags := list.ElementsAs(ctx, &items, false)
	if !diags.HasError() {
		*target = &items
	}
	return diags
}

// PopulateOptionalSetField populates an optional slice field (**[]T) from a Terraform set.
func PopulateOptionalSetField[T any](ctx context.Context, set types.Set, target **[]T) diag.Diagnostics {
	if set.IsNull() || set.IsUnknown() {
		return nil
	}

	if _, ok := set.ElementType(ctx).(types.ObjectType); ok {
		items := make([]T, 0)
		var diags diag.Diagnostics
		for _, el := range set.Elements() {
			if obj, ok := el.(types.Object); ok {
				cleanObj, d := sanitizeObject(ctx, obj)
				diags.Append(d...)
				if d.HasError() {
					continue
				}

				var item T
				d = cleanObj.As(ctx, &item, basetypes.ObjectAsOptions{})
				diags.Append(d...)
				if !d.HasError() {
					items = append(items, item)
				}
			} else {
				diags.AddError("Conversion Error", "Expected Object element")
			}
		}
		if !diags.HasError() {
			*target = &items
		}
		return diags
	}

	items := make([]T, 0)
	diags := set.ElementsAs(ctx, &items, false)
	if !diags.HasError() {
		*target = &items
	}
	return diags
}

// PopulateOptionalObjectField populates an optional object field (**T) from a Terraform object.
func PopulateOptionalObjectField[T any](ctx context.Context, obj types.Object, target **T) diag.Diagnostics {
	if obj.IsNull() || obj.IsUnknown() {
		return nil
	}
	// Sanitize first
	cleanObj, diags := sanitizeObject(ctx, obj)
	if diags.HasError() {
		return diags
	}

	var item T
	d := cleanObj.As(ctx, &item, basetypes.ObjectAsOptions{})
	diags.Append(d...)
	if !diags.HasError() {
		*target = &item
	}
	return diags
}

// PopulateObjectField populates an object field (*T) from a Terraform object.
func PopulateObjectField[T any](ctx context.Context, obj types.Object, target *T) diag.Diagnostics {
	if obj.IsNull() || obj.IsUnknown() {
		return nil
	}
	// Sanitize first
	cleanObj, diags := sanitizeObject(ctx, obj)
	if diags.HasError() {
		return diags
	}

	d := cleanObj.As(ctx, target, basetypes.ObjectAsOptions{})
	diags.Append(d...)
	return diags
}

// PopulateMapField populates a map field (*map[string]T) from a Terraform map.
func PopulateMapField[T any](ctx context.Context, tfMap types.Map, target *map[string]T) diag.Diagnostics {
	if tfMap.IsNull() || tfMap.IsUnknown() {
		return nil
	}

	items := make(map[string]T)
	diags := tfMap.ElementsAs(ctx, &items, false)
	if !diags.HasError() {
		*target = items
	}
	return diags
}

// PopulateOptionalMapField populates an optional map field (*map[string]T) from a Terraform map.
func PopulateOptionalMapField[T any](ctx context.Context, tfMap types.Map, target *map[string]T) diag.Diagnostics {
	if tfMap.IsNull() || tfMap.IsUnknown() {
		return nil
	}

	items := make(map[string]T)
	diags := tfMap.ElementsAs(ctx, &items, false)
	if !diags.HasError() {
		*target = items
	}
	return diags
}
