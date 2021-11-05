package main

import "testing"

func Test_sortWord(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{
				word: "bcad",
			},
			want: "abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortWord(tt.args.word); got != tt.want {
				t.Errorf("sortWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
