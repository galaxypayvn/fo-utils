package valid

import (
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
