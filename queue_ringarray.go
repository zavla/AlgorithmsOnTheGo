package main

type queuering[T ~int] struct {
	begin int
	end   int
	n     int
	ar    []T
}

func NewQueuering[T ~int](n int) queuering[T] {
	return queuering[T]{
		begin: 0,
		end:   0,
		n:     n + 1,
		ar:    make([]T, n+1), //не разрешаем указателю end догнать begin, их равенство значит пустая очередь.
	}
}
func (q *queuering[T]) Empty() bool {
	if q.begin == q.end {
		return true
	}
	return false
}
func (q *queuering[T]) Put(element T) {

	if (q.end+1)%q.n == q.begin {
		//TODO: resize
		panic("oversize")

	}
	q.ar[q.end] = element
	q.end = (q.end + 1) % q.n

}
func (q *queuering[T]) Get() T {
	if q.begin == q.end {
		panic("get from empty string")
	}
	tmp := q.ar[q.begin]
	q.begin = (q.begin + 1) % q.n
	return tmp
}
