package main

import "testing"

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

	})
}
