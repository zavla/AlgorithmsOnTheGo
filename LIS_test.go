//наибольшая возрастающая последовательность

//Largest Increasing Subsequance LIS

package main

import (
	"reflect"
	"testing"
)

func TestLIS(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 2, 6, 1, 4, 5}}, 4, nil},
		{"2", args{[]int{1, 2, 6, 1, 7, 1}}, 4, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := LIS(tt.args.a)
			if got != tt.want {
				t.Errorf("LIS() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LIS() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
