package main

import (
	"fmt"
	"testing"
)

func TestLRUCache_Get(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{
				key: 6,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(2)
			this.Put(2, 1)
			this.Put(2, 2)
			this.Put(6, 2)

			fmt.Printf("expected 6, got %d", this.Get(2))

			if got := this.Get(tt.args.key); got != tt.want {
				t.Errorf("LRUCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
