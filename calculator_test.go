package main

import "testing"

func Test_calculate(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "ok",
			s:    "(1+(4+5+2)-3)+(6+8)",
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.s); got != tt.want {
				t.Errorf("calculate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
