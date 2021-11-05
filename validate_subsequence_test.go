package main

import "testing"

func TestIsValidSubsequence(t *testing.T) {
	type args struct {
		array    []int
		sequence []int
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
			if got := IsValidSubsequence(tt.args.array, tt.args.sequence); got != tt.want {
				t.Errorf("IsValidSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
