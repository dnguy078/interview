package main

import "testing"

func TestLongestPalindromicSubstring(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{
				"abad",
			},
			want: "aba",
		},
		{
			name: "long",
			args: args{
				"abad",
			},
			want: "aba",
		},
		{
			name: "single",
			args: args{
				"a",
			},
			want: "a",
		},
		{
			name: "double",
			args: args{
				"aa",
			},
			want: "aa",
		},
		{
			name: "double only one",
			args: args{
				"ab",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindromicSubstring(tt.args.str); got != tt.want {
				t.Errorf("LongestPalindromicSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
