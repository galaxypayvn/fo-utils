package valid

import (
	"reflect"
	"time"
)

/*
GetPointer returns a pointer to the input value of type T.
*/
func GetPointer[T any](input T) *T {
	if timeInput, ok := any(input).(time.Time); ok && timeInput.IsZero() {
		return nil
	}
	return &input
}

func GetNonPointer[T any](input *T) T {
	if input == nil {
		return reflect.Zero(reflect.TypeOf((*T)(nil)).Elem()).Interface().(T)
	}
	return *input
}
