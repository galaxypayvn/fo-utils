package valid

import (
	"reflect"
)

func IsSlice(v interface{}) bool {
	if v == nil {
		return false
	}

	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Pointer {
		switch t.Kind() {
		case reflect.Slice, reflect.Array:
			return true
		default:
			return false
		}
	}

	return t.Elem().Kind() == reflect.Slice
}

// GetValue returns the value of v if it is not nil, otherwise it returns the zero value of type T.
func GetValue[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}
	return *v
}
