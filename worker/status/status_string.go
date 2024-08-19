// Code generated by "stringer -type=Status"; DO NOT EDIT.

package workerstatus

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Success-200]
	_ = x[Drop-500]
	_ = x[Retry-400]
	_ = x[FailReproduce-302]
}

const (
	_Status_name_0 = "Success"
	_Status_name_1 = "FailReproduce"
	_Status_name_2 = "Retry"
	_Status_name_3 = "Drop"
)

func (i Status) String() string {
	switch {
	case i == 200:
		return _Status_name_0
	case i == 302:
		return _Status_name_1
	case i == 400:
		return _Status_name_2
	case i == 500:
		return _Status_name_3
	default:
		return "Status(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}