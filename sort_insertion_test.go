package main

import (
	"reflect"
	"testing"
)

func Test_sort_insertion(t *testing.T) {
	cases := []struct {
		name string
		arg  []int
		want []int
	}{
		{"1", []int{3, 4, 8, 5}, []int{3, 4, 5, 8}},
		{"2", []int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			arg := make([]int, len(v.arg))
			copy(arg, v.arg)
			sort_insertion(arg)
			if !reflect.DeepEqual(v.arg, arg) {
				t.Errorf("want %v, got %v", v.arg, arg)
			}
		})
	}

}
