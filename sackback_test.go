//Задача о рюкзаке.

package main

import (
	"reflect"
	"testing"
)

func Test_sackback(t *testing.T) {
	type args struct {
		W int
		m []int
		c []int
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantitems []int
	}{
		// TODO: Add test cases.
		{name: "1", args: args{W: 7, m: []int{3, 1, 5, 1}, c: []int{10, 3, 18, 2}}, want: 24, wantitems: []int{1, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, items := sackback(tt.args.W, tt.args.m, tt.args.c); got != tt.want ||
				!reflect.DeepEqual(items, tt.wantitems) {
				t.Errorf("sackback() = %v, %v, want %v, %v", got, items, tt.want, tt.wantitems)
			}
		})
	}
}
