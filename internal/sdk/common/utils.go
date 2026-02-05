package common

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// BuildQueryFilters extracts filter values from a filters struct using reflection.
// It converts Terraform attribute values to query parameter strings based on tfsdk tags.
func BuildQueryFilters(filtersStruct interface{}) map[string]string {
	filters := make(map[string]string)
	if filtersStruct == nil {
		return filters
	}

	// Get the value, handling pointer types
	val := reflect.ValueOf(filtersStruct)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return filters
		}
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return filters
	}

	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tfsdkTag := field.Tag.Get("tfsdk")
		if tfsdkTag == "" || tfsdkTag == "-" {
			continue
		}

		fieldVal := val.Field(i).Interface()
		if attrVal, ok := fieldVal.(attr.Value); ok {
			if attrVal.IsNull() || attrVal.IsUnknown() {
				continue
			}
			switch v := attrVal.(type) {
			case types.String:
				filters[tfsdkTag] = v.ValueString()
			case types.Int64:
				filters[tfsdkTag] = fmt.Sprintf("%d", v.ValueInt64())
			case types.Bool:
				filters[tfsdkTag] = fmt.Sprintf("%t", v.ValueBool())
			case types.Float64:
				filters[tfsdkTag] = fmt.Sprintf("%f", v.ValueFloat64())
			}
		}
	}

	return filters
}

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
		var items []T
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
		if !diags.HasError() && len(items) > 0 {
			*target = items
		}
		return diags
	}

	var items []T
	diags := list.ElementsAs(ctx, &items, false)
	if !diags.HasError() && len(items) > 0 {
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
		var items []T
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
		if !diags.HasError() && len(items) > 0 {
			*target = items
		}
		return diags
	}

	var items []T
	diags := set.ElementsAs(ctx, &items, false)
	if !diags.HasError() && len(items) > 0 {
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
		var items []T
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

	var items []T
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
		var items []T
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

	var items []T
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

// ResolveResourceUUID extracts the resource UUID from a marketplace order response.
func ResolveResourceUUID(orderRes *OrderDetails) string {
	if orderRes == nil {
		return ""
	}
	if orderRes.ResourceUuid != nil && *orderRes.ResourceUuid != "" {
		return *orderRes.ResourceUuid
	}
	if orderRes.MarketplaceResourceUuid != nil && *orderRes.MarketplaceResourceUuid != "" {
		return *orderRes.MarketplaceResourceUuid
	}
	return ""
}

// WaitForOrder blocks until a marketplace order reaches the "done" state.
func WaitForOrder(ctx context.Context, c *client.Client, orderUUID string, timeout time.Duration) (*OrderDetails, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"pending", "executing", "created"},
		Target:  []string{"done"},
		Refresh: func() (interface{}, string, error) {
			var res OrderDetails
			err := c.GetURL(ctx, fmt.Sprintf("/api/marketplace-orders/%s/", orderUUID), &res)
			if err != nil {
				return nil, "", err
			}

			state := ""
			if res.State != nil {
				state = *res.State
			}
			if state == "erred" || state == "rejected" {
				msg := ""
				if res.ErrorMessage != nil {
					msg = *res.ErrorMessage
				}
				return &res, "failed", fmt.Errorf("order failed: %s", msg)
			}
			return &res, state, nil
		},
		Timeout:    timeout,
		Delay:      DefaultPollDelay,
		MinTimeout: DefaultPollMinTimeout,
	}

	rawResult, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return nil, err
	}

	return rawResult.(*OrderDetails), nil
}

// ResourceWithState defines the interface for resources that have a state and error message.
type ResourceWithState interface {
	GetState() string
	GetErrorMessage() string
}

// WaitForResource blocks until a resource reaches the "OK" state.
func WaitForResource[T ResourceWithState](ctx context.Context, getResource func(context.Context) (T, error), timeout time.Duration) (T, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"CREATION_SCHEDULED", "CREATING", "UPDATE_SCHEDULED", "UPDATING", "DELETION_SCHEDULED", "DELETING"},
		Target:  []string{"OK"},
		Refresh: func() (interface{}, string, error) {
			res, err := getResource(ctx)
			if err != nil {
				// return nil doesn't work for generic T if T is not nullable interface, but here T is ResourceWithState (interface), so nil is fine?
				// Actually T is a type parameter constrained by interface. We can return *new(T) or similar?
				// But we are returning (interface{}, ...). generic T is lost here.
				return nil, "", err
			}

			// Use interface methods directly on res
			state := res.GetState()

			if state == "ERRED" {
				msg := res.GetErrorMessage()
				if msg == "" {
					msg = "unknown error"
				}
				return res, "failed", fmt.Errorf("%s", msg)
			}

			return res, state, nil
		},
		Timeout:    timeout,
		Delay:      DefaultPollDelay,
		MinTimeout: DefaultPollMinTimeout,
	}

	rawResult, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		var zero T
		return zero, err
	}

	return rawResult.(T), nil
}

// WaitForDeletion blocks until a resource is gone (404).
func WaitForDeletion[T ResourceWithState](ctx context.Context, getResource func(context.Context) (T, error), timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"CREATION_SCHEDULED", "CREATING", "UPDATE_SCHEDULED", "UPDATING", "DELETION_SCHEDULED", "DELETING", "OK"},
		Refresh: func() (interface{}, string, error) {
			res, err := getResource(ctx)
			if err != nil {
				if IsNotFoundError(err) {
					return nil, "", nil
				}
				return nil, "", err
			}

			// Use interface methods
			state := res.GetState()

			if state == "ERRED" {
				msg := res.GetErrorMessage()
				if msg == "" {
					msg = "unknown error"
				}
				return res, "failed", fmt.Errorf("%s", msg)
			}

			return res, state, nil
		},
		Timeout:    timeout,
		Delay:      DefaultPollDelay,
		MinTimeout: DefaultPollMinTimeout,
	}

	_, err := stateConf.WaitForStateContext(ctx)
	return err
}

// IsNotFoundError checks if an error represents a 404 Not Found response
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "HTTP 404")
}

// StringToFloat64Ptr converts a string pointer to a types.Float64 value.
// This is used because the API returns decimal values as quoted strings (e.g., "11.00000").
func StringToFloat64Ptr(s *string) types.Float64 {
	if s == nil || *s == "" {
		return types.Float64Null()
	}
	f, err := strconv.ParseFloat(*s, 64)
	if err != nil {
		return types.Float64Null()
	}
	return types.Float64Value(f)
}

// StringPointerValue returns types.StringNull() if the pointer is nil or points to an empty string.
// This is useful for optional fields where the Waldur API returns "" instead of null.
func StringPointerValue(s *string) types.String {
	if s == nil || *s == "" {
		return types.StringNull()
	}
	return types.StringValue(*s)
}

// FlexibleNumber is a custom type that can unmarshal from both JSON numbers and strings.
// This is needed because the Waldur API is inconsistent: some decimal fields are returned
// as JSON numbers (e.g. 0) and others as quoted strings (e.g. "11.00000").
type FlexibleNumber float64

// UnmarshalJSON implements the json.Unmarshaler interface.
func (f *FlexibleNumber) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	// Try unmarshaling as a number first
	var n float64
	if err := json.Unmarshal(data, &n); err == nil {
		*f = FlexibleNumber(n)
		return nil
	}

	// If that fails, try unmarshaling as a string
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if s == "" {
		*f = 0
		return nil
	}

	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	*f = FlexibleNumber(val)
	return nil
}

// Float64Ptr returns a pointer to the float64 value.
func (f *FlexibleNumber) Float64Ptr() *float64 {
	if f == nil {
		return nil
	}
	v := float64(*f)
	return &v
}
