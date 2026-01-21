package datasources

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ConvertTFValue converts a Terraform attribute value to a Go interface{}.
// It handles types.String, types.Int64, types.Bool, types.Float64, types.List, and types.Object.
func ConvertTFValue(v attr.Value) interface{} {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	switch val := v.(type) {
	case types.String:
		return val.ValueString()
	case types.Int64:
		return val.ValueInt64()
	case types.Bool:
		return val.ValueBool()
	case types.Float64:
		return val.ValueFloat64()
	case types.List:
		items := make([]interface{}, len(val.Elements()))
		for i, elem := range val.Elements() {
			items[i] = ConvertTFValue(elem)
		}
		return items
	case types.Object:
		obj := make(map[string]interface{})
		for k, attr := range val.Attributes() {
			if converted := ConvertTFValue(attr); converted != nil {
				obj[k] = converted
			}
		}
		return obj
	case types.Map:
		obj := make(map[string]interface{})
		for k, attr := range val.Elements() {
			if converted := ConvertTFValue(attr); converted != nil {
				obj[k] = converted
			}
		}
		return obj
	}
	return nil
}

// APIListModel is a helper struct that can be used with types.ListValueFrom.
// It converts []interface{} from API to a slice of typed models.
type APIListModel[T any] struct {
	Items []T
}

// ConvertAPIListFrom converts a Go slice to a Terraform types.List using types.ListValueFrom.
// The modelSlice should be a pointer to a slice of structs with tfsdk tags.
func ConvertAPIListFrom(ctx context.Context, val interface{}, elemType attr.Type, modelSlice interface{}) (types.List, diag.Diagnostics) {
	arr, ok := val.([]interface{})
	if !ok || arr == nil || len(arr) == 0 {
		return types.ListNull(elemType), nil
	}

	// Use reflection to populate the model slice
	sliceVal := reflect.ValueOf(modelSlice).Elem()
	sliceType := sliceVal.Type()
	elemModelType := sliceType.Elem()

	for _, item := range arr {
		objMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		// Create a new element of the model type
		newElem := reflect.New(elemModelType).Elem()

		// Populate fields based on tfsdk tags
		for i := 0; i < elemModelType.NumField(); i++ {
			field := elemModelType.Field(i)
			tfsdkTag := field.Tag.Get("tfsdk")
			if tfsdkTag == "" || tfsdkTag == "-" {
				continue
			}

			apiVal, exists := objMap[tfsdkTag]
			if !exists {
				// Set null value for missing fields
				newElem.Field(i).Set(reflect.ValueOf(getNullValue(field.Type)))
				continue
			}

			// Convert API value to Terraform type
			tfVal := convertToTFType(apiVal, field.Type)
			newElem.Field(i).Set(reflect.ValueOf(tfVal))
		}

		sliceVal.Set(reflect.Append(sliceVal, newElem))
	}

	return types.ListValueFrom(ctx, elemType, modelSlice)
}

// ConvertAPIObjectFrom converts a Go map to a Terraform types.Object using types.ObjectValueFrom.
func ConvertAPIObjectFrom(ctx context.Context, val interface{}, attrTypes map[string]attr.Type, model interface{}) (types.Object, diag.Diagnostics) {
	objMap, ok := val.(map[string]interface{})
	if !ok || objMap == nil {
		return types.ObjectNull(attrTypes), nil
	}

	// Use reflection to populate the model
	modelVal := reflect.ValueOf(model).Elem()
	modelType := modelVal.Type()

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		tfsdkTag := field.Tag.Get("tfsdk")
		if tfsdkTag == "" || tfsdkTag == "-" {
			continue
		}

		apiVal, exists := objMap[tfsdkTag]
		if !exists {
			modelVal.Field(i).Set(reflect.ValueOf(getNullValue(field.Type)))
			continue
		}

		tfVal := convertToTFType(apiVal, field.Type)
		modelVal.Field(i).Set(reflect.ValueOf(tfVal))
	}

	return types.ObjectValueFrom(ctx, attrTypes, model)
}

// convertToTFType converts an API value to the appropriate Terraform type.
func convertToTFType(val interface{}, targetType reflect.Type) interface{} {
	if val == nil {
		return getNullValue(targetType)
	}

	typeName := targetType.String()

	switch typeName {
	case "basetypes.StringValue", "types.String":
		if str, ok := val.(string); ok {
			return types.StringValue(str)
		}
		return types.StringNull()
	case "basetypes.Int64Value", "types.Int64":
		if num, ok := val.(float64); ok {
			return types.Int64Value(int64(num))
		}
		return types.Int64Null()
	case "basetypes.BoolValue", "types.Bool":
		if b, ok := val.(bool); ok {
			return types.BoolValue(b)
		}
		return types.BoolNull()
	case "basetypes.Float64Value", "types.Float64":
		if num, ok := val.(float64); ok {
			return types.Float64Value(num)
		}
		return types.Float64Null()
	default:
		return types.StringNull()
	}
}

// getNullValue returns the null value for a Terraform type.
func getNullValue(targetType reflect.Type) interface{} {
	typeName := targetType.String()

	switch typeName {
	case "basetypes.StringValue", "types.String":
		return types.StringNull()
	case "basetypes.Int64Value", "types.Int64":
		return types.Int64Null()
	case "basetypes.BoolValue", "types.Bool":
		return types.BoolNull()
	case "basetypes.Float64Value", "types.Float64":
		return types.Float64Null()
	default:
		return types.StringNull()
	}
}

// ConvertAPIValue converts a Go interface{} value (from API response) to Terraform attr.Value.
// It handles primitives, maps (objects), and slices (lists).
func ConvertAPIValue(ctx context.Context, val interface{}, attrType attr.Type) attr.Value {
	if val == nil {
		return getNullAttrValue(attrType)
	}

	switch t := attrType.(type) {
	case basetypes.StringType:
		if str, ok := val.(string); ok {
			return types.StringValue(str)
		}
		return types.StringNull()
	case basetypes.Int64Type:
		if num, ok := val.(float64); ok {
			return types.Int64Value(int64(num))
		}
		return types.Int64Null()
	case basetypes.BoolType:
		if b, ok := val.(bool); ok {
			return types.BoolValue(b)
		}
		return types.BoolNull()
	case basetypes.Float64Type:
		if num, ok := val.(float64); ok {
			return types.Float64Value(num)
		}
		return types.Float64Null()
	case basetypes.ListType:
		return ConvertAPIList(ctx, val, t.ElemType)
	case basetypes.ObjectType:
		return ConvertAPIObject(ctx, val, t.AttrTypes)
	}
	return nil
}

// ConvertAPIList converts a Go slice to a Terraform types.List.
func ConvertAPIList(ctx context.Context, val interface{}, elemType attr.Type) types.List {
	arr, ok := val.([]interface{})
	if !ok || arr == nil {
		return types.ListNull(elemType)
	}

	items := make([]attr.Value, 0, len(arr))
	for _, item := range arr {
		if elemVal := ConvertAPIValue(ctx, item, elemType); elemVal != nil {
			items = append(items, elemVal)
		}
	}

	listVal, _ := types.ListValue(elemType, items)
	return listVal
}

// ConvertAPIObject converts a Go map[string]interface{} to a Terraform types.Object.
func ConvertAPIObject(ctx context.Context, val interface{}, attrTypes map[string]attr.Type) types.Object {
	objMap, ok := val.(map[string]interface{})
	if !ok || objMap == nil {
		return types.ObjectNull(attrTypes)
	}

	attrValues := make(map[string]attr.Value)
	for key, attrType := range attrTypes {
		if v, exists := objMap[key]; exists {
			attrValues[key] = ConvertAPIValue(ctx, v, attrType)
		} else {
			attrValues[key] = getNullAttrValue(attrType)
		}
	}

	objVal, _ := types.ObjectValue(attrTypes, attrValues)
	return objVal
}

// getNullAttrValue returns the null attr.Value for the given attr.Type.
func getNullAttrValue(attrType attr.Type) attr.Value {
	switch t := attrType.(type) {
	case basetypes.StringType:
		return types.StringNull()
	case basetypes.Int64Type:
		return types.Int64Null()
	case basetypes.BoolType:
		return types.BoolNull()
	case basetypes.Float64Type:
		return types.Float64Null()
	case basetypes.ListType:
		return types.ListNull(t.ElemType)
	case basetypes.ObjectType:
		return types.ObjectNull(t.AttrTypes)
	default:
		return types.StringNull()
	}
}

// ConvertAPIString extracts a string from a value.
func ConvertAPIString(val interface{}) types.String {
	if val == nil {
		return types.StringNull()
	}
	if str, ok := val.(string); ok {
		return types.StringValue(str)
	}
	return types.StringNull()
}
