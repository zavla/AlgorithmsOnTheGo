//Волновой алгоритм находит длинну кратчайшего пути но в невзвешенном графе.

package main

import (
	"reflect"
	"testing"
)

func Test_graph_volnovoi_poisk(t *testing.T) {
	type args struct {
		adj   [][]int
		start int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1",
			args: args{adj: [][]int{
				{0, 1, 1, 0, 0, 0},
				{0, 0, 0, 1, 0, 0},
				{1, 0, 0, 0, 1, 1},
				{0, 1, 0, 0, 1, 0},
				{0, 0, 1, 1, 0, 1},
				{0, 0, 1, 0, 1, 0},
			},
				start: 0},
			want: []int{0, 1, 1, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := graph_volnovoy_poisk_shortest_path(tt.args.adj, tt.args.start)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}
