package valid

import (
	"github.com/google/uuid"
	"time"
)

func BoolPointer(b bool) *bool {
	return &b
}

func StringPointer(s string) *string {
	return &s
}

func IntPointer(i int) *int {
	return &i
}

func Int32Pointer(i int32) *int32 {
	return &i
}

func Int64Pointer(i int64) *int64 {
	return &i
}

func Float64Pointer(i float64) *float64 {
	return &i
}

func Float32Pointer(i float32) *float32 {
	return &i
}

func DayTimePointer(i time.Time) *time.Time {
	if i.IsZero() {
		return nil
	} else {
		return &i
	}
}

func UUIDPointer(i uuid.UUID) *uuid.UUID {
	return &i
}

/*
GetPointer returns a pointer to the input value of type T.
*/
func GetPointer[T any](input T) *T {
	if timeInput, ok := any(input).(time.Time); ok && timeInput.IsZero() {
		return nil
	}
	return &input
}
