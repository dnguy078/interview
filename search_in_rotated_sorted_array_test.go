package main

import "testing"

func Test_searchInRotated(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{
				nums:   []int{45, 61, 71, 72, 73, 0, 1, 21, 33, 37},
				target: 33,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInRotated(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInRotated() = %v, want %v", got, tt.want)
			}
		})
	}
}
