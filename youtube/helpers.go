package youtube

import (
	"fmt"
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
