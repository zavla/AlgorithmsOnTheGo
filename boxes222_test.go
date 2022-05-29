package main

import (
	"strconv"
	"testing"
)

func Test_numberOfDifferentPermutationsOFBoxes(t *testing.T) {
	var tests = []struct {
		name int
		arg  int
		want int
	}{
		{4, 4, 13},
		{5, 5, 24},
	}
	for _, tt := range tests {
		t.Run(strconv.FormatInt(int64(tt.name), 10), func(t *testing.T) {
			got := numberOfDifferentPermutationsOfBoxes(tt.arg)
			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
