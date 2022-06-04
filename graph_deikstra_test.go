package main

import (
	"reflect"
	"testing"
)

var test1 = []struct {
	s, e rune
	cost int
}{
	{'A', 'B', 4},
	{'A', 'C', 1},
	{'A', 'D', 2},
	{'D', 'F', 1},
	{'C', 'E', 7},
	{'B', 'E', 2},
	{'F', 'E', 5},
}

func TestGraph_load(t *testing.T) {
	G := make(Graph)
	G.load(test1)
}
func TestString(t *testing.T) {
	G := make(Graph)
	G.load(test1)

	//t.Errorf("%v", G)
}
func TestGraph_getminvert(t *testing.T) {
	type args struct {
		D     map[rune]int
		notIn map[rune]bool
	}
	tests := []struct {
		name  string
		G     *Graph
		args  args
		want  rune
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.G.nearest_neigbour(tt.args.D, tt.args.notIn)
			if got != tt.want {
				t.Errorf("Graph.getminvert() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Graph.getminvert() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGraph_deikstra_shortest_path(t *testing.T) {
	type args struct {
		start rune
	}
	G := make(Graph)
	G.load(test1)

	tests := []struct {
		name string
		G    *Graph
		args args
		want maprune
	}{
		// TODO: Add test cases.
		{name: "1", G: &G, args: args{start: 'A'}, want: map[rune]int{
			'A': 0, 'B': 4, 'C': 1, 'D': 2, 'F': 3, 'E': 6,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maprune(tt.G.deikstra_shortest_path(tt.args.start)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.deikstra_shortest_path() = %v, want %v", got, tt.want)
			}
		})
	}
}
