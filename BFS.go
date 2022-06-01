package main

//список для вершин которые обработать.
//в цикле взять из списка, во внутреннем цикле добавлять в список соседей данной вершины.

func BFS(A [][]int, start int, visit func(v int)) {
	visited := make(map[int]bool, len(A))

	todo := make([]int, 1, len(A))
	todo[0] = start

	for i := 0; i < len(todo); i++ {
		v := todo[i]
		if _, ok := visited[v]; ok {
			continue
		}
		visit(v)
		visited[v] = true

		for j := 0; j < len(A); j++ {
			if A[v][j] == 0 {
				continue
			}
			if _, ok := visited[j]; ok {
				continue
			}
			todo = append(todo, j)
		}
	}

}
