//Задача о рюкзаке.

package main

import (
	"reflect"
	"testing"
)

//Тест использует BruteForceSackbag чтобы проверить решение динамикой.
func Test_sackbag(t *testing.T) {
	type args struct {
		W      float64
		m      []float64
		c      []int
		factor int //you dicide when and how to scale
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		//W - это лимит рюкзака по весу.
		//m - это веса предметов.
		//c - это стоимости предметов.
		//factor - число на которое надо умножить все веса чтобы они стали целыми числами.
		{name: "int values", args: args{W: 7, m: []float64{3, 2, 5, 1}, c: []int{10, 3, 18, 2}, factor: 1}},
		{name: "float values", args: args{W: 23.7, m: []float64{6, 15.8, 3.7, 7.7, 10.1, 17.7}, c: []int{10, 5, 23, 13, 8, 20}, factor: 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//BruteForce checks the correctness of sackbag.
			want1, want2 := BruteForceSackbag(tt.args.m, tt.args.c, tt.args.W)

			//test values are scaled to be ints
			mScaledInt := scaleToInts(tt.args.m, tt.args.factor)
			WScaled := int(tt.args.W * float64(tt.args.factor))

			got, items := sackbag(WScaled, mScaledInt, tt.args.c)
			itemsfloat := scaleToFloats(items, 1.0/float64(tt.args.factor)) //scale back the answer

			if float64(got) != want1 ||
				!reflect.DeepEqual(itemsfloat, want2) {
				t.Errorf("sackbag() = %v, %v, want %v, %v", got, items, want1, want2)
			}
		})
	}
}
