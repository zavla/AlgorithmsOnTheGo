package main

import "fmt"

//A = [][]int матр смежности
func visit(v int) {
	fmt.Printf("%v ", v)
}

//нам понадобится вершины складывать в стек чтобы обходить вглубину
//-----------СТЕК--------------
type node struct {
	left *node
	val  int
}
type StackLinkedList struct {
	top *node
}

func (s *StackLinkedList) Pop() int {
	if s.top == nil {
		panic("empty stack")
	}
	t := s.top.val
	s.top = s.top.left
	return t
}
func (s *StackLinkedList) Push(el int) {
	t := &node{
		left: s.top,
		val:  el,
	}
	s.top = t
}
func (s *StackLinkedList) Empty() bool {
	if s.top == nil {
		return true
	}
	return false
}

//-----------------------
func DFS(A [][]int, start int, visit func(v int)) {
	//пройти по строке смежных вершин
	//использовать стек для ведения очередности обхода вершин
	//использовать список посещеных потому что граф имеет циклы
	stack := StackLinkedList{}
	stack.Push(start)

	visited := make(map[int]bool, len(A)) //map для константного поиска

	for !stack.Empty() {
		v := stack.Pop()
		//если граф ввиде ромба то нижняя вершина будет два раза помещена в стек, хотя она еще непосещена
		if _, ok := visited[v]; ok {
			continue
		}
		visited[v] = true
		visit(v)

		//ниже цикл только заносящий в стек
		for i := len(A) - 1; i >= 0; i-- {
			if i == v || A[v][i] == 0 { //если нет ребра из v -> i
				continue
			}
			//не ходить по кругу
			if _, ok := visited[i]; ok {
				continue
			}
			stack.Push(i) //следующая вершина для обработки i
		}
	}
}
