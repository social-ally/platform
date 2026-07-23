package youtube

import (
	"fmt"
	"reflect"
	"strings"
)

func stringValue(value any) string {
	switch typed := value.(type) {
	case []string:
		return strings.Join(typed, ",")
	case []any:
		values := make([]string, len(typed))
		for i, value := range typed {
			values[i] = fmt.Sprint(value)
		}
		return strings.Join(values, ",")
	default:
		reflected := reflect.ValueOf(value)
		if reflected.IsValid() && (reflected.Kind() == reflect.Slice || reflected.Kind() == reflect.Array) {
			values := make([]string, reflected.Len())
			for index := range values {
				values[index] = fmt.Sprint(reflected.Index(index).Interface())
			}
			return strings.Join(values, ",")
		}
		return fmt.Sprint(value)
	}
}
func scopeValue(scopes []Scope) string {
	values := make([]string, len(scopes))
	for i, scope := range scopes {
		values[i] = string(scope)
	}
	return strings.Join(values, " ")
}
