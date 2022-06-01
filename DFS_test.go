package main

import "testing"

var adj = [][]int{
	{0, 1, 2, 0, 0, 0},
	{1, 0, 0, 1, 1, 0},
	{2, 0, 0, 8, 1, 0},
	{0, 1, 8, 0, 0, 1},
	{0, 1, 1, 0, 0, 6},
	{0, 0, 0, 1, 6, 0},
}

type printintest testing.T

func (pt *printintest) visit(v int) {
	pt.Log(v)
}

func visit_test(t *testing.T, v int) {
	t.Log(v)
}
func TestStackLinkedList(t *testing.T) {
	s := StackLinkedList{}
	gotEmpty := s.Empty()
	if gotEmpty != true {
		t.Errorf("error in empty")
	}
	s.Push(2)
	s.Push(1)
	got1 := s.Pop()
	if got1 != 1 {
		t.Error("Pop()!=1")
	}
	got2 := s.Pop()
	if got2 != 2 {
		t.Error("Pop()!=2")
	}
	defer func() {
		if err := recover(); err != nil {
			//ok
		} else {
			t.Error("expects panic")
		}
	}()
	gotP := s.Pop() //expects panic
	_ = gotP
}

func TestDFS(t *testing.T) {
	type args struct {
		A     [][]int
		start int
		visit func(v int)
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"name", args{A: adj, start: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prn := printintest(*t)
			DFS(tt.args.A, tt.args.start, prn.visit)
		})
	}
}
