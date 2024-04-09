package valid

import (
	"github.com/google/uuid"
	"reflect"
	"time"
)

func Bool(in *bool) bool {
	if in == nil {
		return false
	}
	return *in
}

func String(in *string) string {
	if in == nil {
		return ""
	}
	return *in
}

func Int(in *int) int {
	if in == nil {
		return 0
	}
	return *in
}

func Int64(in *int64) int64 {
	if in == nil {
		return 0
	}
	return *in
}

func Int32(in *int32) int32 {
	if in == nil {
		return 0
	}
	return *in
}

func Float64(in *float64) float64 {
	if in == nil {
		return 0
	}
	return *in
}

func Float32(in *float32) float32 {
	if in == nil {
		return 0
	}
	return *in
}

func Byte(in *byte) byte {
	if in == nil {
		return 0
	}
	return *in
}

func DayTime(in *time.Time) time.Time {
	if in == nil {
		return time.Time{}
	}
	return *in
}

func UUID(req *uuid.UUID) uuid.UUID {
	if req == nil {
		return uuid.Nil
	}
	return *req
}

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
