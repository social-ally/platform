package facebook

import (
	"encoding/json"
	"fmt"
	"net/url"
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
	return strings.Join(values, ",")
}

func formValues(value any) (url.Values, error) {
	values := url.Values{}
	structValue := reflect.Indirect(reflect.ValueOf(value))
	if !structValue.IsValid() || structValue.Kind() != reflect.Struct {
		return values, nil
	}
	structType := structValue.Type()
	for index := 0; index < structValue.NumField(); index++ {
		field := structValue.Field(index)
		if !field.CanInterface() || ((field.Kind() == reflect.Pointer || field.Kind() == reflect.Interface) && field.IsNil()) {
			continue
		}
		if field.Kind() == reflect.Interface {
			field = field.Elem()
		}
		name, _, _ := strings.Cut(structType.Field(index).Tag.Get("json"), ",")
		if name == "" || name == "-" {
			continue
		}
		fieldValue := field.Interface()
		switch field.Kind() {
		case reflect.Struct, reflect.Slice, reflect.Map:
			encoded, err := json.Marshal(fieldValue)
			if err != nil {
				return nil, err
			}
			values.Set(name, string(encoded))
		case reflect.Pointer:
			values.Set(name, fmt.Sprint(field.Elem().Interface()))
		default:
			values.Set(name, fmt.Sprint(fieldValue))
		}
	}
	return values, nil
}
