package main

import (
	"reflect"
	"testing"
)

func TestCodec_EncodeDecode(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want []string
	}{
		{
			name: "ok",
			strs: []string{"ab", "abcde"},
			want: []string{"ab", "abcde"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codec := &Codec{}
			got := codec.Decode(codec.Encode(tt.strs))
			if len(got) != len(tt.want) {
				t.Error("expected to be")
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected to be equal %+v, %+v", got, tt.want)
			}
		})
	}
}
