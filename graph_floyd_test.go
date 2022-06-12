package main

import (
	"reflect"
	"testing"
)

func Test_floyd(t *testing.T) {
	adj := [][]int{
		{0, 1, 2, 0, 0, 0},
		{1, 0, 0, 1, 1, 0},
		{2, 0, 0, 8, 1, 0},
		{0, 1, 8, 0, 0, 1},
		{0, 1, 1, 0, 0, 6},
		{0, 0, 0, 1, 6, 0},
	}
	got := floyd(adj)

	want := [][]int{
		{2, 1, 2, 2, 2, 3},
		{1, 2, 2, 1, 1, 2},
		{2, 2, 2, 3, 1, 4},
		{2, 1, 3, 2, 2, 1},
		{2, 1, 1, 2, 2, 3},
		{3, 2, 4, 1, 3, 2},
	}

	if !reflect.DeepEqual(got, want) {

		t.Errorf("got %v, want %v", got, want)
	}
}
