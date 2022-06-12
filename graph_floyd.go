package main

//for
//	for
//		for
//на входе матрица смежности
//на выходе матрица кратчайших расстояний.
//О(n^3)
func min1_2plus3(ij, ik, kj int) int {
	//0 в матрице смежности понимается как "нет пути"

	if ik == 0 || kj == 0 {
		return ij
	}
	if ij == 0 {
		return ik + kj
	}
	if ij < ik+kj {
		return ij
	}
	return ik + kj

}

func floyd(adj [][]int) [][]int {
	ret := make([][]int, len(adj))
	copy(ret, adj)
	//0 в матрице смежности понимается как "нет пути"

	for k := 0; k < len(adj); k++ {
		for i := 0; i < len(adj); i++ {
			for j := 0; j < len(adj); j++ {
				ret[i][j] = min1_2plus3(ret[i][j], ret[i][k], ret[k][j])
			}
		}
	}
	return ret

}
