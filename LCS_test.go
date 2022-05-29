package main

import (
	"reflect"
	"testing"
)

func Test_LCS(t *testing.T) {
	type res struct {
		_1 int
		_2 []int
	}

	tests := []struct {
		name string
		arg1 []int
		arg2 []int
		want res
	}{
		{"1", []int{0, 2, 1, 3, 4}, []int{2, 1, 0, 3, 0},
			res{3, []int{2, 1, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_1, _2 := LCS(tt.arg1, tt.arg2)
			if _1 != tt.want._1 {
				t.Errorf("got %v want %v", _1, tt.want._1)
			}
			if !reflect.DeepEqual(_2, tt.want._2) {
				t.Errorf("_2 got %v want %v", _2, tt.want._2)
			}
		})
	}
}
