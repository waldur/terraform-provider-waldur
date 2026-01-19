package resources

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	}
	return nil
}
