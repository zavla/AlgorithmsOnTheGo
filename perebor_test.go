package main

import "testing"

func TestBruteForceMaxValue(t *testing.T) {
	type args struct {
		m     []int
		c     []int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "1", args: args{m: []int{5, 7, 4, 1}, c: []int{3, 6, 2, 1}, limit: 10}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BruteForceMaxValue(tt.args.m, tt.args.c, tt.args.limit); got != tt.want {
				t.Errorf("BruteForceMaxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
