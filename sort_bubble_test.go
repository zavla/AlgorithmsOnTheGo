package main

import (
	"reflect"
	"testing"
)

func Test_sort_bubble(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 3, 2}}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := make([]int, len(tt.args.a))
			copy(a, tt.args.a)
			sort_bubble(a)
			if !reflect.DeepEqual(a, tt.want) {
				t.Errorf("got %v, want %v", a, tt.want)
			}
		})
	}
}
