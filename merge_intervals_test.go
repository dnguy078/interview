package main

import (
	"reflect"
	"testing"
)

func xTestMergeOverlappingIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "ok",
			args: args{
				intervals: [][]int{{2, 6}, {8, 10}, {15, 18}, {1, 3}},
			},
			want: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeOverlappingIntervals(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeOverlappingIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
