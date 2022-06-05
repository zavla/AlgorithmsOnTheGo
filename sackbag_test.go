//Задача о рюкзаке.

package main

import (
	"reflect"
	"testing"
)

func Test_sackbag(t *testing.T) {
	type args struct {
		W      float64
		m      []float64
		c      []int
		factor int //you dicide when and how to scale
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantitems []int
	}{
		// TODO: Add test cases.
		{name: "1", args: args{W: 7, m: []float64{3, 2, 5, 1}, c: []int{10, 3, 18, 2}, factor: 1},
			want: 21, wantitems: []int{2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//BruteForce checks the correctness of sackbag.
			want1, want2 := BruteForceSackbag(tt.args.m, tt.args.c, tt.args.W)

			//test values are scaled to be ints
			mScaledInt := scaleToInts(tt.args.m, tt.args.factor)
			WScaled := int(tt.args.W) * tt.args.factor

			got, items := sackbag(WScaled, mScaledInt, tt.args.c)
			gotfloat := float64(got / tt.args.factor) //scale back the answer
			itemsfloat := scaleToFloats(items, 1.0/float64(tt.args.factor))

			if gotfloat != want1 ||
				!reflect.DeepEqual(itemsfloat, want2) {
				t.Errorf("sackbag() = %v, %v, want %v, %v", got, items, want1, want2)
			}
		})
	}
}
