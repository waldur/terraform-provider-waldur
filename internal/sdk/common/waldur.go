package common

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ExtractUUIDFromURL extracts a UUID from a Waldur API URL.
// If the input is already a UUID (no slashes), it returns it unchanged.
// Example: "http://api.example.com/api/openstack-tenants/abc123/" -> "abc123"
func ExtractUUIDFromURL(urlOrUUID string) string {
	// If it doesn't contain a slash, assume it's already a UUID
	if !strings.Contains(urlOrUUID, "/") {
		return urlOrUUID
	}
	urlOrUUID = strings.TrimSuffix(urlOrUUID, "/")
	parts := strings.Split(urlOrUUID, "/")
	return parts[len(parts)-1]
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
