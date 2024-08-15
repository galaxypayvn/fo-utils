package utmail

import "testing"

func TestNormalize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Gmail lowercase", "test@gmail.com", "test@gmail.com"},
		{"Gmail uppercase", "TEST@GMAIL.COM", "test@gmail.com"},
		{"Gmail with dots", "t.e.s.t@gmail.com", "test@gmail.com"},
		{"Gmail with plus", "test+alias@gmail.com", "test@gmail.com"},
		{"Gmail with dots and plus", "t.e.s.t+alias@gmail.com", "test@gmail.com"},
		{"Googlemail alias", "test@googlemail.com", "test@gmail.com"},
		{"Hotmail lowercase", "test@hotmail.com", "test@hotmail.com"},
		{"Hotmail with plus", "test+alias@hotmail.com", "test@hotmail.com"},
		{"Hotmail with dots", "t.e.s.t@hotmail.com", "t.e.s.t@hotmail.com"},
		{"Live.com lowercase", "test@live.com", "test@live.com"},
		{"Live.com with dots", "t.e.s.t@live.com", "test@live.com"},
		{"Live.com with plus", "test+alias@live.com", "test@live.com"},
		{"Outlook lowercase", "test@outlook.com", "test@outlook.com"},
		{"Outlook with plus", "test+alias@outlook.com", "test@outlook.com"},
		{"Outlook with dots", "t.e.s.t@outlook.com", "t.e.s.t@outlook.com"},
		{"Non-normalizable domain", "test@example.com", "test@example.com"},
		{"Invalid email format", "testexample.com", "testexample.com"},
		{"Empty string", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeEmail(tt.input); got != tt.want {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
