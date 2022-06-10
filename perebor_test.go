package main

import (
	"reflect"
	"testing"
)

func TestBruteForceSackbag(t *testing.T) {
	type args struct {
		m     []float64
		c     []int
		limit float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want2 []float64
	}{
		// TODO: Add test cases.
		{name: "1", args: args{m: []float64{5, 7, 4, 1}, c: []int{3, 6, 2, 1}, limit: 10},
			want: 7, want2: []float64{7, 1}},
		{name: "2", args: args{m: []float64{6, 15.8, 3.7, 7.7, 10.10, 17.7}, c: []int{10, 5, 23, 13, 8, 20}, limit: 23.7},
			want: 46, want2: []float64{6, 3.7, 7.7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got1, got2 := BruteForceSackbag(tt.args.m, tt.args.c, tt.args.limit); got1 != tt.want || !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("BruteForceSackbag() got1 %v, want %v; got2 %v, want2 %v", got1, tt.want, got2, tt.want2)
			}
		})
	}
}
