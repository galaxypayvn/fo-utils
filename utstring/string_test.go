package utstring

import "testing"

func TestTransformString(t *testing.T) {
	type args struct {
		input     string
		uppercase bool
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "Trim spaces and convert to lowercase",
			args: args{
				input:     "  Hello, World!  ",
				uppercase: false,
			},
			expected: "hello, world!",
		},
		{
			name: "Trim spaces and convert to uppercase",
			args: args{
				input:     "  Hello, World!  ",
				uppercase: true,
			},
			expected: "HELLO, WORLD!",
		},
		{
			name: "Remove non-spacing marks and convert to lowercase",
			args: args{
				input:     "Sửa chữa",
				uppercase: false,
			},
			expected: "sua chua",
		},
		{
			name: "Remove non-spacing marks",
			args: args{
				input:     "Ĥêllô, Wôrld!",
				uppercase: false,
			},
			expected: "hello, world!",
		},
		{
			name: "Replace 'Đ' and 'đ' with 'D' and 'd'",
			args: args{
				input:     "Đồng Nai",
				uppercase: false,
			},
			expected: "dong nai",
		},
		{
			name: "Invalid UTF-8 sequence",
			args: args{
				input:     "\x80", // Invalid UTF-8 sequence
				uppercase: false,
			},
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransformString(tt.args.input, tt.args.uppercase); got != tt.expected {
				t.Errorf("TransformString() = %v, expect %v", got, tt.expected)
			}
		})
	}
}

func Test_isMark(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMark(tt.args.r); got != tt.want {
				t.Errorf("isMark() = %v, want %v", got, tt.want)
			}
		})
	}
}
