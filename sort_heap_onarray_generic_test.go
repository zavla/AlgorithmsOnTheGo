package main

import (
	"math/rand"
	"testing"
)

type intLess int

func (i intLess) Less(j *intLess) bool {
	if i < *j {
		return true
	}
	return false
}

func TestHeapOnArray(t *testing.T) {
	t.Run("1", func(t *testing.T) {

		h := NewHeapOnArray[intLess]()
		h.Add(0)
		h.Add(1)
		h.Add(2)
		_2 := h.GetMax()
		if _2 != 2 {
			t.Error("want _2==2")
		}
		h.RemoveMax()
		_1 := h.GetMax()
		if _1 != 1 {
			t.Error("want _1==1")
		}
	})
}

func BenchmarkHeapOnArray_Add(b *testing.B) {
	h := NewHeapOnArray[intLess]()

	for i := 0; i < b.N; i++ {
		h.Add(intLess(rand.Intn(101)))
	}
	for i := 0; i < b.N; i++ {
		m := h.GetMax()
		_ = m
		//print(m, " ")
		h.RemoveMax()
	}
}
