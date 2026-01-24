package common

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

// PopulateSliceField helps populating a slice field from a Terraform list.
func PopulateSliceField[T any](ctx context.Context, list types.List, target *[]T) {
	var items []T
	if diags := list.ElementsAs(ctx, &items, false); !diags.HasError() && len(items) > 0 {
		*target = items
	}
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
		Delay:      10 * time.Second,
		MinTimeout: 5 * time.Second,
	}

	rawResult, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return nil, err
	}

	return rawResult.(*OrderDetails), nil
}

// WaitForResource blocks until a resource reaches the "OK" state.
func WaitForResource[T any](ctx context.Context, getResource func(context.Context) (*T, error), timeout time.Duration) (*T, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"CREATION_SCHEDULED", "CREATING", "UPDATE_SCHEDULED", "UPDATING", "DELETION_SCHEDULED", "DELETING"},
		Target:  []string{"OK"},
		Refresh: func() (interface{}, string, error) {
			res, err := getResource(ctx)
			if err != nil {
				return nil, "", err
			}

			// Use reflection to check for "state" field
			val := reflect.ValueOf(res)
			if val.Kind() == reflect.Ptr {
				val = val.Elem()
			}

			state := "OK" // Default state if field not found
			if val.Kind() == reflect.Struct {
				field := val.FieldByName("State")
				if field.IsValid() && field.Kind() == reflect.String {
					state = field.String()
				}
			}

			if state == "ERRED" {
				// Try to find error_message
				msg := "resource in erred state"
				if val.Kind() == reflect.Struct {
					errField := val.FieldByName("ErrorMessage")
					if errField.IsValid() && errField.Kind() == reflect.String {
						msg = errField.String()
					} else {
						// Ptr to string?
						errField = val.FieldByName("ErrorMessage")
						if errField.IsValid() && errField.Kind() == reflect.Ptr && !errField.IsNil() {
							if errField.Elem().Kind() == reflect.String {
								msg = errField.Elem().String()
							}
						}
					}
				}
				return res, "failed", fmt.Errorf("%s", msg)
			}

			return res, state, nil
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 5 * time.Second,
	}

	rawResult, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return nil, err
	}

	return rawResult.(*T), nil
}

// WaitForDeletion blocks until a resource is gone (404).
func WaitForDeletion[T any](ctx context.Context, getResource func(context.Context) (*T, error), timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"CREATION_SCHEDULED", "CREATING", "UPDATE_SCHEDULED", "UPDATING", "DELETION_SCHEDULED", "DELETING", "OK"},
		Refresh: func() (interface{}, string, error) {
			res, err := getResource(ctx)
			if err != nil {
				if client.IsNotFoundError(err) {
					return nil, "", nil
				}
				return nil, "", err
			}

			// Use reflection to check for "state" field
			val := reflect.ValueOf(res)
			if val.Kind() == reflect.Ptr {
				val = val.Elem()
			}

			state := "DELETING" // Default state if field not found
			if val.Kind() == reflect.Struct {
				field := val.FieldByName("State")
				if field.IsValid() && field.Kind() == reflect.String {
					state = field.String()
				}
			}

			if state == "ERRED" {
				// Try to find error_message
				msg := "resource in erred state"
				if val.Kind() == reflect.Struct {
					errField := val.FieldByName("ErrorMessage")
					if errField.IsValid() && errField.Kind() == reflect.String {
						msg = errField.String()
					} else {
						// Ptr to string?
						errField = val.FieldByName("ErrorMessage")
						if errField.IsValid() && errField.Kind() == reflect.Ptr && !errField.IsNil() {
							if errField.Elem().Kind() == reflect.String {
								msg = errField.Elem().String()
							}
						}
					}
				}
				return res, "failed", fmt.Errorf("%s", msg)
			}

			return res, state, nil
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 5 * time.Second,
	}

	_, err := stateConf.WaitForStateContext(ctx)
	return err
}
