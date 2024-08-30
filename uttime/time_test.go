package uttime

import (
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestParseTimeLocation(t *testing.T) {
	type args struct {
		t            time.Time
		locationName string
	}
	tests := []struct {
		args args
		want time.Time
	}{
		// Test Case 1: Convert UTC to Asia/Ho_Chi_Minh
		{
			args: args{
				t:            time.Date(2024, time.August, 30, 12, 0, 0, 0, time.UTC),
				locationName: "Asia/Ho_Chi_Minh",
			},
			want: func() time.Time {
				loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
				return time.Date(2024, time.August, 30, 19, 0, 0, 0, loc)
			}(),
		},
		// Test Case 2: Convert UTC to America/New_York
		{
			args: args{
				t:            time.Date(2024, time.August, 30, 12, 0, 0, 0, time.UTC),
				locationName: "America/New_York",
			},
			want: func() time.Time {
				loc, _ := time.LoadLocation("America/New_York")
				return time.Date(2024, time.August, 30, 8, 0, 0, 0, loc)
			}(),
		},
		// Test Case 3: Invalid Timezone - should return original time
		{
			args: args{
				t:            time.Date(2024, time.August, 30, 12, 0, 0, 0, time.UTC),
				locationName: "Invalid/Timezone",
			},
			want: time.Date(2024, time.August, 30, 12, 0, 0, 0, time.UTC),
		},
		// Test Case 4: Convert UTC to UTC - should return the same time
		{
			args: args{
				t:            time.Date(2024, time.August, 30, 12, 0, 0, 0, time.UTC),
				locationName: "UTC",
			},
			want: time.Date(2024, time.August, 30, 12, 0, 0, 0, time.UTC),
		},
		// Test Case 5: Convert UTC to Europe/London during summer (BST)
		{
			args: args{
				t:            time.Date(2024, time.August, 30, 12, 0, 0, 0, time.UTC),
				locationName: "Europe/London",
			},
			want: func() time.Time {
				loc, _ := time.LoadLocation("Europe/London")
				return time.Date(2024, time.August, 30, 13, 0, 0, 0, loc)
			}(),
		},
	}

	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			if got := ParseTimeLocation(tt.args.t, tt.args.locationName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseTimeLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
