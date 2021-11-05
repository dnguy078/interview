package main

import "testing"

func Test_isValidParenthesis(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "ok",
			input: "{}",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidParenthesis(tt.input); got != tt.want {
				t.Errorf("isValidParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
