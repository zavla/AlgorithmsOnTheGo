package main

import (
	"reflect"
	"testing"
)

func Test_getn(t *testing.T) {
	type args struct {
		N int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"58", args{58, 1}, 8},
		{"123", args{123, 3}, 1},
		{"1234", args{1234, 3}, 2},
		{"123", args{123, 5}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getn(tt.args.N, tt.args.n); got != tt.want {
				t.Errorf("getn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sort_porazrad(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{123, 234, 178, 123, 12}}, []int{12, 123, 123, 178, 234}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort_porazrad(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sort_porazrad() = %v, want %v", got, tt.want)
			}
		})
	}
}
