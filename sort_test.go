package main

import (
	"reflect"
	"testing"
)

func Test_sort(t *testing.T) {
	type test struct {
		name string
		a    []int
		ret  []int
	}
	tests := []test{
		{name: "10", a: []int{0}, ret: []int{0}},

		{name: "9", a: []int{0, 0, 0, 1}, ret: []int{0, 0, 0, 1}},

		{name: "8", a: []int{1, 0, 0, 0}, ret: []int{0, 0, 0, 1}},

		{name: "7", a: []int{0, 0, 0, 0}, ret: []int{0, 0, 0, 0}},

		{name: "6", a: []int{1, 0, 0, 1}, ret: []int{0, 0, 1, 1}},

		{name: "1", a: []int{1, 3, 2}, ret: []int{1, 2, 3}},
		{name: "2", a: []int{2, 3, 5, 6, 7}, ret: []int{2, 3, 5, 6, 7}},

		{name: "3", a: []int{6, 3, 2, 1}, ret: []int{1, 2, 3, 6}},
		{name: "4", a: []int{3, 2, 1}, ret: []int{1, 2, 3}},
		{name: "5", a: []int{1, 2, 3}, ret: []int{1, 2, 3}},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {

			if ret := merge(v.a); !reflect.DeepEqual(ret, v.ret) {
				t.Errorf("got %v, want %v", ret, v.ret)
			}
			//quicksort
			ret := make([]int, len(v.a))
			copy(ret, v.a)
			sort_quick(ret)
			if !reflect.DeepEqual(ret, v.ret) {
				t.Errorf("sort_quick got %v, want %v", ret, v.ret)

			}
			//сортировка шелла
			ret = make([]int, len(v.a))
			copy(ret, v.a)
			sort_shell(ret)
			if !reflect.DeepEqual(ret, v.ret) {
				t.Errorf("sort_shell got %v want %v", ret, v.ret)
			}

		})
	}
}
