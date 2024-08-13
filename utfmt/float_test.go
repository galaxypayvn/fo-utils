package utfmt

import (
	"strconv"
	"testing"
)

func TestFormatThousand(t *testing.T) {
	type args struct {
		amount   float64
		separate rune
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{amount: 1033232116564989.931, separate: 0},
			want: "1,033,232,116,564,989",
		},
		{
			args: args{amount: 123.45, separate: 0},
			want: "123",
		},
		{
			args: args{amount: 0, separate: 0},
			want: "0",
		},
		{
			args: args{amount: -1234567.89, separate: 0},
			want: "-1,234,567",
		},
		{
			args: args{amount: 999, separate: 0},
			want: "999",
		},
		{
			args: args{amount: 1000, separate: 0},
			want: "1,000",
		},
		{
			args: args{amount: 1000000.00, separate: 0},
			want: "1,000,000",
		},
		{
			args: args{amount: 1033232116564989.931, separate: '.'},
			want: "1.033.232.116.564.989",
		},
	}
	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			if got := FormatThousand(tt.args.amount, tt.args.separate); got != tt.want {
				t.Errorf("FormatThousand() = %v, want %v", got, tt.want)
			}
		})
	}
}
