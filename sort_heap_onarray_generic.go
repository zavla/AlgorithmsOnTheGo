package main

// heaparray is a generic heap based on array.

type hasLess[T any] interface {
	Less(el2 *T) bool // generic interface uses user instanciated Less method
}

// User will instanciate Less method on his type later.
// for example:
// type intLess int
// func (i intLess) Less(j *intLess) bool {
// 	if i < *j {
// 		return true
// 	}
// 	return false
// }

type heaparray[T hasLess[T]] struct {
	ar []*element[T]
}
type element[T hasLess[T]] struct {
	value T
}

func (e *element[T]) Less(el2 *element[T]) bool {
	if e.value.Less(&el2.value) {
		return true
	}
	return false
}
func NewHeapOnArray[T hasLess[T]]() heaparray[T] {
	h := heaparray[T]{}
	h.ar = make([]*element[T], 1, 10) //не используем 0 индекс , т.к. формула 2к, 2к+1 для индексов додерних узлов
	return h
}
func (h *heaparray[T]) Add(el T) {
	n := &element[T]{el}
	h.ar = append(h.ar, n)
	h.moveUp(len(h.ar) - 1)
}
func (h *heaparray[T]) RemoveMax() {
	li := len(h.ar) - 1
	if li <= 0 {
		return
	}
	h.ar[1] = h.ar[li] // используем слайс с индекса 1, не с 0
	h.ar[li] = nil
	h.ar = h.ar[:li]
	h.moveDown(1)
}
func (h *heaparray[T]) GetMax() T {
	return h.ar[1].value
}

// support functions
func (h *heaparray[T]) moveUp(ind int) {
	k := ind / 2
	for k >= 1 {

		if h.ar[k].Less(h.ar[ind]) {
			h.ar[k], h.ar[ind] = h.ar[ind], h.ar[k]
		}
		k, ind = k/2, k
	}
}
func (h *heaparray[T]) maxElementIndex(l, r int) int {
	if r >= len(h.ar) {
		return l
	}
	if !h.ar[l].Less(h.ar[r]) {
		return l
	}
	return r
}
func (h *heaparray[T]) moveDown(ind int) {
	l := 2 * ind
	r := l + 1
	for l < len(h.ar) {

		imax := h.maxElementIndex(l, r)

		if h.ar[imax].Less(h.ar[ind]) {
			return
		}

		h.ar[ind], h.ar[imax] = h.ar[imax], h.ar[ind]
		ind = imax
		l = 2 * ind
		r = l + 1
	}
}
