package common

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
