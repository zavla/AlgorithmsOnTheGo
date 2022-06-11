package main

import (
	"testing"
)

func TestHeapLinkedNodes(t *testing.T) {
	t.Run("heap use", func(t *testing.T) {

		h := &heap[int]{}

		h.Add(1)

		got1 := h.GetMax()
		if got1 != 1 {
			t.Errorf("didn't got 1")
			return
		}
		h.RemoveMax()
		if h.root != nil {
			t.Errorf("remove last element error")
			return
		}

		h.Add(1)
		h.Add(2)
		penult := h.findPenultimate() //find last but one
		if penult.value != 1 {

			t.Error("findPenultimate != 1")
		}
		h.Add(0)
		if h.last.value != 0 {
			t.Error("last.value != 0")
		}
		penult1 := h.findPenultimate()
		if penult1.value != 1 {
			t.Error("findPenultimate1 != 1")
		}
		h.Add(-1)
		if h.last.value != -1 {
			t.Error("last.value != -1")
		}
		if h.root.left.left.value != -1 {
			t.Error("root.left.left.value != -1")
		}
		h.Add(3)
		if h.root.value != 3 {
			t.Error("root.value != 3")
		}
		h.Add(5)
		penult2 := h.findPenultimate()
		if penult2.value != 1 {
			t.Error("after adding 5 we expect 1 to be the last but one, penultimate")
		}
		// 		  5
		// 		/ 	\
		// 	   2	  3
		// 	  /	\ 	/
		//   -1  1	0
	})
}
